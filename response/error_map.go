package response

type ResCode int

// 各业务边界起始代码
const (
	SystemCode = 10000 + iota*100
	ItemCode
	KmsCode
	DmsCode
	ErpCode
	WmsCode
)

// 系统相关 错误码10000开始
const (
	DataBaseConnectionError   = SystemCode + iota //数据库连接接错误
	DataBaseOperationError                        //数据库操作错误
	UserInfoNotExist                              //用户信息不存在
	UsernameOrPasswordError                       // 用户名/密码错误
	NameExisted                                   //名称重复
	ParamMissing                                  //参数缺失
	ParamFormatError                              //参数格式错误
	AuthExpire                                    //权限验证过期，请重新登录
	AuthParamMissing                              //权限验证参数缺失
	AuthError                                     //权限验证错误
	LoginMethodError                              //登录方式错误，此账号被限制只能在web端或者app端登录
	FileCacheError                                //文件缓存错误
	FileImportError                               //文件导入错误
	BaseUnitExisted                               //基本单位信息已存在
	FileImportErrorCode                           //导入文件中存在不合规商品代码
	FileImportWithoutMFG                          //导入文件中存在缺失生产日期的商品
	FileImportStorageError                        //导入文件中库位信息出错
	OperatorDesignated                            //已设置过操作员,请检查
	FileImportDateFormatError                     //导入文件日期格式错误
	//TODO 只有制表状态下可以删除(未使用)
	Deleted                   //已被删除
	FileImportLackNecessary   //导入文件,缺少必要的字段,请检查
	DataNotChanged            //数据未发生修改,请检查
	FileExportError           //文件导出错误
	NoDataToExport            //没有可以导出的数据
	BillNotExist              //单据信息不存在
	FormUnitFormat            //表单中单位数据不合规
	ConfigInfoExisted         //配置信息已存在
	FormNotComplete           //表格信息未填写完整，请检查
	CanNotModified            //只有特定状态可进行修改，该表格状态已发生改变，请刷新页面
	InfoNotExist              //查询的信息不存在
	CodeMissing               //代码缺失,请补充
	AccountDisabled           //账号已被禁用
	BaseInfoNotAllowModified  //基础信息不允许变更
	InfoModifiedNotExist      //修改的信息不存在
	CanNotModifiedOrderType   //当前状态无法修改订单类型
	OrderTypeNotSelect        //订单类型未选择，无法同步
	BillNotAdd                //单据明细还未添加,请添加
	OperationFailed           //操作失败
	Synchronized              //已同步，请检查
	BillNotAllowModified      //单据不可发生该变更
	BillStatusNotMatch        //单据状态不匹配,请刷新页面或者重新进入后重试
	HouseApplyNotInUnfinished //申请表不处于部分上架状态
	HouseApplyInInUnfinished  //申请表还未完全上架
	ClassInfoNotExist         //分类信息不存在
	ClassExisted              //分类名称或code已存在，请更换
	PubParamMissing           //公共请求参数缺失
	SignCheckFail             //sign校验失败
	InvalidAppId              //无效app_id
	SynchronizedError         //同步系统发生错误
	SynchronizedNoData        //同步无数据
	CanNotSynchronized        //当前状态无法同步
	BillUpload                //单据已上传，请勿重复操作
	BillNotFoundAMAItem       //单据中存在不归属与ama系统的商品，无法同步
	UserBillNotMatch          //TODO 非批发用户单据无法同步，请更换主体(原系统未使用)
	//TODO 批发用户单据无法同步，请更换主体(原系统未使用)
	BillNotProcessed       //单据还有关联单据未处理,无法同步
	NeedConfirmOrderType   //请先确定好订单类型，再进行此操作
	SynchronizedAuthExpire //同步权限过期，需要重新授权，即将请跳转页面进行授权操作后，再行操作\ TODO 是否启用通用的权限过期代码
	OldPasswordError       //旧密码错误，查询不到用户信息
	CanNotBind             //无法绑定，该账号已绑定过微信号
	//经营属性代码已存在，请重新输入代码 // TODO 是否启用通用代码
	//参数异常 TODO 是否启用通用代码（暂定弃用）
	DataDeletedNotExist //该删除数据不存在
	UnknownError        //其他系统错误
)

// 商品相关 错误码10100开始
const (
	ItemNotExist           = ItemCode + iota //商品信息不存在
	ItemBarCodeExisted                       //条码已被分配,请检查
	ItemBarCodeNotExisted                    //条码不存在,请检查
	CodeExisted                              //代码已被分配,请检查
	ItemCanNotSynchronized                   //当前商品状态无法同步
	CodeException                            //代码已被分配到已删除商品
	StorageInItemNotExist                    //入库商品信息不存在
	NotItemToInventory                       //当前库存内没有可以盘点的商品
	FeeExisted                               //重复插入商品配送服务费
	UnitError                                //导入表格中有商品的单位信息不合规,请检查
	ItemListEmpty                            //商品列表为空,请检查
	ItemInfoIncomplete                       //存在未输入数量/单位/生产日期/库位的商品
	ItemNotDeleted                           //该商品尚有库存，无法删除
	FormExistItemNotMFG                      //表单中存在未填写生产日期的商品,请检查
	ItemNumError                             //该code的商品数量不对
)

// 绩效考核相关 错误码10300开始
const (
	UnfinishedExam        = KmsCode + iota //您还有未完成的考核
	UserDeptNotSet                         //用户部门未设置,无法参与考核
	LackQuestion                           //题库中题目不足,无法完成考核,请联系题库管理员
	ExamEnd                                //考核已结束,无法提交答案
	UnfinishedPreQuestion                  //题目不允许有前置空选项,请检查
	NotAnswer                              //题目不允许没有答案
	AnswerExistNull                        //正确答案不允许有空选项,请检查
	NoAnswer                               //您没有作答任何题目
	ExamDeptToMany                         //考核部门太多了,请检查
)

// 车辆相关 错误码10400开始
const (
	CarInfoNotExist    = DmsCode + iota //车辆信息不存在
	DriverInfoNotExist                  //驾驶员信息不存在
)

// 门店相关 错误码10500开始
const (
	StoreCodeMissing = ErpCode + iota //门店code缺失
	IsDC                              //该门店已是配送中心，不可添加上游配送中心
	// TODO 门店状态由非配送中心变为配送中心时，upstream_center_id和upstream_center_name必须为空
	DCCanNotSelf        //上游配送中心不可为自己
	DCNotMatch          //上游配送中心信息与已有门店信息不匹配
	DCNotExist          //该上游配送中心id不存在
	NotDC               //选择的上游配送中心不是配送中心
	DefaultHouseExisted //该门店已经有默认仓库
)

// 仓储相关 错误码10600开始
const (
	// Pallet     //已设置过托盘 TODO 原系统未使用
	PalletBusy          = WmsCode + iota //托盘忙碌中
	PickOrQCUnfinished                   //只有拣货完成和质检完成状态下可以设置
	StorageNotExist                      //库位信息不存在
	PalletNotExist                       //托盘信息不存在
	CanNotDeleteStorage                  //库位上还有商品,无法删除该库位
	//选择的托盘不处于空闲状态,请重新选择 TODO 语义重复
	PalletNotModified //托盘与原纪录一致,信息未发生改变
	//托盘上没有商品可供上架
	//入库申请表已被领取
	//现有库存无法满足申请表需求,请检查
	//入库申请表信息不存在
	//商品信息已超期，无法收货
	//质检表商品为空，无法匹配库位
	//质检表不可发生该变更
	//质检表状态不匹配
	//质检表中存在未输入生产日期/过期日期/质检数量的商品
	//该托盘没有需要上架的任务
	//质检表尚未选择托盘,不能审核
	//质检表信息有未填写项,请检查
	//质检表信息不存在,请检查
	//只有制表状态的质检表才可操作
	//出库申请表没有可分拣的商品
	//仓库中没有可以分拣的商品
	//出库申请表不存在
	//申请主体未选择
	//申请单中所有商品的数量为0，请填写
	//当前状态下,拣货单不可发生该变更
	//申请表中有商品仓位未选择
	//表单中没有商品可以拆分
	//申请表不存在多个顶级库位分类的商品
	//只有制表状态才可以拆分
	//该条调仓记录已被操作
	//没有可以合并的拣货单
	//拣货单中没有需要合并的数据
	//拣货单中已有同批次同库位商品,请检查
	//当前状态下,调仓单不可发生该变更,请检查
	//调仓单已生成,等待调仓员操作中
	//只有制表状态下单据可以被领取
	//只有拣货完成状态下单据可以被质检
	//拣货单未设置拣货员，请设置后再开始拣货
	//只有拣货中的单据可以回滚操作
	//盘点数和原纪录一致,本次未做修改
	//此次盘点修改了批次信息,修改后的批次库存中已有记录,请检查
	//同一备货单内所有商品的到货时间应该统一，请进行拆单处理或修改到货日期
	//备货单不存在可拆分的商品
	//结算单中预计金额和实际结算金额不匹配，请检查
	//先添加商品,再修改主体
	//没有足够的空闲出库待发区,请释放出库待发区
)

// 供应商相关 错误码10700开始
const ()

func (r ResCode) Msg() string {
	if value, ok := codeMsgMap[r]; ok {
		return value
	}
	return codeMsgMap[UnknownError]
}

var codeMsgMap = map[ResCode]string{
	CodeSuccess:               "success",
	DataBaseConnectionError:   "数据库连接错误",
	DataBaseOperationError:    "数据库操作错误",
	UserInfoNotExist:          "用户信息不存在",
	UsernameOrPasswordError:   "用户名/密码错误",
	NameExisted:               "名称重复",
	ParamMissing:              "参数缺失",
	ParamFormatError:          "参数格式错误",
	AuthExpire:                "权限验证过期，请重新登录",
	AuthParamMissing:          "权限验证参数缺失",
	AuthError:                 "权限验证错误",
	LoginMethodError:          "登录方式错误，此账号被限制只能在web端或者app端登录",
	FileCacheError:            "文件缓存错误",
	FileImportError:           "文件导入错误",
	BaseUnitExisted:           "基本单位信息已存在",
	FileImportErrorCode:       "导入文件中存在不合规商品代码",
	FileImportWithoutMFG:      "导入文件中存在缺失生产日期的商品",
	FileImportStorageError:    "导入文件中库位信息出错",
	OperatorDesignated:        "已设置过操作员,请检查",
	FileImportDateFormatError: "导入文件日期格式错误",
	Deleted:                   "已被删除",
	FileImportLackNecessary:   "导入文件,缺少必要的字段,请检查",
	DataNotChanged:            "数据未发生修改,请检查",
	FileExportError:           "文件导出错误",
	NoDataToExport:            "没有可以导出的数据",
	BillNotExist:              "单据信息不存在",
	FormUnitFormat:            "表单中单位数据不合规",
	ConfigInfoExisted:         "配置信息已存在",
	FormNotComplete:           "表格信息未填写完整，请检查",
	CanNotModified:            "只有特定状态可进行修改，该表格状态已发生改变，请刷新页面",
	InfoNotExist:              "查询的信息不存在",
	CodeMissing:               "代码缺失,请补充",
	AccountDisabled:           "账号已被禁用",
	BaseInfoNotAllowModified:  "基础信息不允许变更",
	InfoModifiedNotExist:      "修改的信息不存在",
	CanNotModifiedOrderType:   "当前状态无法修改订单类型",
	OrderTypeNotSelect:        "订单类型未选择，无法同步",
	BillNotAdd:                "单据明细还未添加,请添加",
	OperationFailed:           "操作失败",
	Synchronized:              "已同步，请检查",
	BillNotAllowModified:      "单据不可发生该变更",
	BillStatusNotMatch:        "单据状态不匹配,请刷新页面或者重新进入后重试",
	HouseApplyNotInUnfinished: "申请表不处于部分上架状态",
	HouseApplyInInUnfinished:  "申请表还未完全上架",
	ClassInfoNotExist:         "分类信息不存在",
	ClassExisted:              "分类名称或code已存在，请更换",
	PubParamMissing:           "公共请求参数缺失",
	SignCheckFail:             "sign校验失败",
	InvalidAppId:              "无效app_id",
	SynchronizedError:         "同步系统发生错误",
	SynchronizedNoData:        "同步无数据",
	CanNotSynchronized:        "当前状态无法同步",
	BillUpload:                "单据已上传，请勿重复操作",
	BillNotFoundAMAItem:       "单据中存在不归属与ama系统的商品，无法同步",
	BillNotProcessed:          "单据还有关联单据未处理,无法同步",
	NeedConfirmOrderType:      "请先确定好订单类型，再进行此操作",
	SynchronizedAuthExpire:    "同步权限过期，需要重新授权，即将请跳转页面进行授权操作后，再行操作",
	OldPasswordError:          "旧密码错误，查询不到用户信息",
	CanNotBind:                "无法绑定，该账号已绑定过微信号",
	DataDeletedNotExist:       "该删除数据不存在",
	UnknownError:              "未知错误",

	ItemNotExist:          "商品不存在",
	ItemBarCodeExisted:    "条码重复",
	ItemBarCodeNotExisted: "条码不存在",
}
