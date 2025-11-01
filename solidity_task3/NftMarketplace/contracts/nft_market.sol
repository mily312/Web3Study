// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;
import "@openzeppelin/contracts/interfaces/IERC20.sol";
import "@openzeppelin/contracts/interfaces/IERC721.sol";

contract Market {
    //通过构造函数把erc20和erc721地址传过来
    IERC20 public erc20;
    IERC721 public erc721;

    struct Order {
        address seller;
        uint256 tokenId;
        uint256 price;
    }

    // 状态变量，存储在链上
    mapping(uint256 => Order) public orderOfId; // token id to Order
    Order[] public orders;
    mapping(uint256 => uint256) idToOrderIndex; // token id to index of orders

    bytes4 internal constant MAGIC_ON_ERC721_RECEIVED = 0x150b7a02;

    //event
    event Deal(address seller, address buyer, uint256 tokenId, uint256 price);
    event CreateOrder(address seller, uint256 tokenId, uint256 price);
    event CancleOrder(address seller, uint256 tokenId);
    event UpdOrder(
        address seller,
        uint256 tokenId,
        uint256 currentPrice,
        uint256 newPrice
    );

    constructor(address _erc20, address _erc721) {
        require(_erc20 != address(0), "zero address");
        require(_erc721 != address(0), "zero address");

        erc20 = IERC20(_erc20);
        erc721 = IERC721(_erc721);
    }

    // 购买nft
    function buy(uint256 _tokenId) external {
        address seller = orderOfId[_tokenId].seller;
        address buyer = msg.sender;
        uint256 price = orderOfId[_tokenId].price;

        // 买家将erc20代币转给卖家
        require(
            erc20.transferFrom(buyer, seller, price),
            "buy nft not success"
        );
        // nft 从合约中转给buyer
        erc721.safeTransferFrom(address(this), buyer, _tokenId);

        //下架
        removeOrder(_tokenId);

        emit Deal(seller, buyer, _tokenId, price);
    }

    // 取消订单
    function cancleOrder(uint256 _tokenId) external {
        address seller = orderOfId[_tokenId].seller;
        require(seller == msg.sender, "you can not cancle this order");

        erc721.safeTransferFrom(address(this), seller, _tokenId);

        //下架
        removeOrder(_tokenId);

        emit CancleOrder(seller, _tokenId);
    }

    // 更改订单价格
    function updOrderPrice(uint256 _tokeId, uint256 _price) external {
        address seller = orderOfId[_tokeId].seller;
        require(seller == msg.sender, "you can not update this order price");

        //取出当前价格，并将映射中存储得价格进行更改
        uint256 currentPrice = orderOfId[_tokeId].price;
        orderOfId[_tokeId].price = _price;

        // 数组中价格也需要进行更改
        uint256 orderIndex = idToOrderIndex[_tokeId];
        orders[orderIndex].price = _price;

        emit UpdOrder(seller, _tokeId, currentPrice, _price);
    }

    // 卖家上架售卖nft

    // 挂单时market自动执行该函数 erc721合约调用safetransferFrom函数时，会自动执行该函数
    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external returns (bytes4) {
        uint256 price = toUint256(data, 0);
        require(price > 0, "price must greater than 0");

        //上架
        Order memory order = Order(from, tokenId, price);
        orderOfId[tokenId] = order;
        orders.push(order);
        idToOrderIndex[tokenId] = orders.length - 1;

        emit CreateOrder(from, tokenId, price);

        // 返回值写死
        return MAGIC_ON_ERC721_RECEIVED;
        //return this.onERC721Received.selector;
    }

    // 下架
    function removeOrder(uint256 _tokenId) internal {
        // 从三个状态变量中移除

        //数组移除
        uint256 index = idToOrderIndex[_tokenId];
        uint256 lastIndex = orders.length - 1;

        if (index != lastIndex) {
            //将数组中得最后一个元素移到要删除的元素位置中
            Order memory lastOrder = orders[lastIndex];
            orders[index] = lastOrder;
            idToOrderIndex[lastOrder.tokenId] = index;
        }
        orders.pop();

        //映射中移除
        delete orderOfId[_tokenId];
        delete idToOrderIndex[_tokenId];
    }

    //格式转换，将价格转为uint256
    // https://stackoverflow.com/questions/63252057/
    // how-to-use-bytestouint-function-in-solidity-the-one-with-assembly
    // 0000000000000000000000000000000000000000000000000001c6bf52634000
    function toUint256(
        bytes memory _bytes,
        uint256 _start
    ) public pure returns (uint256) {
        require(_start + 32 >= _start, "market:toUint256_overflow");
        require(_bytes.length >= _start + 32, "Market: toUint256 outOfBounds");

        uint256 tempUint;

        assembly {
            tempUint := mload(add(add(_bytes, 0x20), _start))
        }

        return tempUint;
    }

    function getOrdersLength() external view returns (uint256) {
        return orders.length;
    }

    function getAllOrders() external view returns (Order[] memory) {
        return orders;
    }

    // 获取我的所有订单
    function getMyAllOrders() external view returns (Order[] memory) {
        Order[] memory myOrders = new Order[](orders.length);
        uint256 index;

        for (uint256 i = 0; i < orders.length; i++) {
            Order memory order = orders[i];
            if (msg.sender == order.seller) {
                myOrders[index] = order;
                index++;
            }
        }

        return myOrders;
    }

    // 是否上架 
    function isListed(uint256 _tokenId) public view returns (bool) {
        return orderOfId[_tokenId].seller != address(0);
    }
}
