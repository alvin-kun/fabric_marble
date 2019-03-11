# 链码功能简介：
## 该链码能够实现对账户的查询、转账和删除账户的功能，并且整合完善资产管理应用链码的功能，能够让用户在分类账上创建资产，并通过指定的函数实现对资产的修改与查询。

### Init函数：初始化两个账户，账户分别为a、b，对应的金额分别为100、200
#### 函数流程：
> * 判断参数个数是否为4
> * 获取args[0]的值并赋给A
> * strconv.Atoi(args[1]) 转化为int型整数，返回bval，err
> * 判断err
> * 获取args[2]的值赋给B
> * strconv.Atoi(arg3[3]) 转换为整数，返回bval，err
> * 判断err
> * 将A的状态值记录到分布式账本中
> * 判断err
> * return shim.Success(nil)

### Invoke函数：应用程序具备三个不同的分支功能：find、payment、delete分别实现查询、转账和删除的功能，根据交易参数定位到不同的分支进行处理。
#### 函数流程
> * 获取函数名称与参数列表
> * 判断函数名称并调用相应的函数

### Invoke调用的函数实现：
#### find函数：根据给定的账号查询对应的状态信息
#### 函数流程
> * 判断参数是否为1个
> * 根据传入的参数调用GetState查询状态， aval，err 为接收的返回值
> * 如果返回err不为空，则返回错误
> * 如果返回的状态为空，则返回错误
> * 如果无错误，返回查询到的值

#### payment函数：根据指定的两个账户名称及余额，实现转账
> * 判断参数是否为3
> * 获取两个账户名称（args[0]与args[1])值，赋给两个变量
> * 调用GetState获取a账户状态，avalsByte，err为返回值
> *    * 判断有无错误（err不为空，avalsByte为空）
> * 类型转换：aval，_ = strconv.Atoi(string(avalsByte))
> * 调用GetState获取b账户状态，bvalsByte，err为返回值
> * * 判断有无错误（err不为空，bvalsByte为空）
> * 类型转换：bval，_ = strconv.Atoi(string(bvalsByte))
> * 将要转账的数额进行类型转换：x, err = strconv.Atoi(args[2])
> * 判断err是否为空
> * aval，bval执行转账操作
> * 记录状态，err = PutState(a, []byte(strconv.Itoa(aval)))
> * * Itoa: 将整数转换为十进制字符串形式
> * * 判断有无错误
> * 记录状态，err = PutState(b, []byte(strconv.Itoa(bval)))
> * return shim.Success(nil)

#### delAccount函数: 根据指定的名称删除对应的实体信息
> * 判断参数个数是否为1
> * 调用DelState方法，err接收返回值
> * 如果err不为空，返回错误
> * 返回成功 shim.Success(nil)

### set函数：设置指定账户的值