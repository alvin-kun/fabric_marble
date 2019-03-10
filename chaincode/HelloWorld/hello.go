package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func main() {
	err := shim.Start(new(HelloChaincode))
	if err != nil {
		fmt.Printf("Start chaincode fail: %v", err)
	}
}

type HelloChaincode struct {
}

// 实例化/升级链码时被自动调用
// -c '{"Args":["Hello","World"]'
func (t *HelloChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Start instantiate chaincode....")

	_, args := stub.GetFunctionAndParameters()
	// 判断参数长度是否为2个
	if len(args) != 2 {
		return shim.Error("The amount of parameters was error")
	}

	fmt.Println("Saving data......")

	// 通过调用PutState方法将数据保存在账本中
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error("Encounter error in saving data...")
	}

	fmt.Println("Instantiate success")

	return shim.Success(nil)

}

// 对账本数据进行操作时被自动调用(query, invoke)
func (t *HelloChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// 获取调用链码时传递的参数内容(包括要调用的函数名及参数)
	fun, args := stub.GetFunctionAndParameters()

	// 客户意图
	if fun == "query" {
		return query(stub, args)
	}

	return shim.Error("Invalid operation, specified function cannot be implemented")
}

func query(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// 检查传递的参数个数是否为1
	if len(args) != 1 {
		return shim.Error("Parameter error，msut and can only specify the corresponding Key")
	}

	// 根据指定的Key调用GetState方法查询数据
	result, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("An error occured while querying for data based on the specified " + args[0] )
	}
	if result == nil {
		return shim.Error("No data was queried based on the specified " + args[0])
	}

	// 返回查询结果
	return shim.Success(result)
}
