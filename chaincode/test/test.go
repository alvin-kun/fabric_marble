//链码实现资产管理

package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Println("启动 SimpleChaincode 时发生错误：%s", err)
	}
}

//Init函数：初始化数据状态
//获取参数，使用GetStringArgs函数传递给调用链码的所需参数
//检查合法性，检查参数数量是否为2个，如果不是，则返回错误信息
//利用两个参数，调用Putstate方法向账本中写入状态，如果有错误则返回shim.Error(),否则返回nil(shim.Success)
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("初始化的参数只能为2个，分别代表名称与状态数据")
	}
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error("再保存状态时出项错误")
	}
	return shim.Success(nil)
}

//Invoke函数：验证函数名称为set或get，并调用链式代码应用程序函数，通过shim.Success或shim.Error函数返回响应。
//获取参数名与参数
//对获取的参数名称进行判断，如果未set，则调用set方法，反之调用get
//set/get函数返回两个值（result，err)
//如果err部位空则返回错误
//err为空则返回[]byte（result）
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fun, args := stub.GetFunctionAndParameters()

	var result string
	var err error
	if fun == "set" {
		result, err = set(stub, args)
	} else {
		result, err = get(stub, args)
	}
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(result))
}

//实现set函数，修改资产
//检查参数个数是否为2
//利用PutState方法将状态写入
//如果成功，则返回要写入的状态，失败返回错误：fmt.Errorf("...")
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("给定的参数个数不符合要求")

	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	return string(args[0]), nil
}

//实现get函数，查询函数
//接受参数并判断个数是否为1个
//调用GetState方法，返回并接收两个返回值(value,err) 判断err及value是否为空 return "", fmt.Errorf("......")
//返回值 return string(value), nil
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("给定的参数不符合要求")
	}
	result, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("获取数据发生错误")
	}
	if result == nil {
		return "", fmt.Errorf("根据 %s 没有获取到相应的数据", args[0])
	}
	return string(result), nil
}
