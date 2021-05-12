# kerbalwzygo
个人使用Golang进行后端开发时封装的一些工具对象

**这里封装都是写工具的初级原型, 具体融合到项目中时可能需要做一些修改**

**懒得封装到兼容性很高, 我个人认为纯粹调包开发非常无聊而且包内有些功能也用不上, 抽空自己造一造轮子其实挺有意思的**

----
- ### u_csv
    + func ReadCSV(filepath string) ([][]string, error) 读取CSV文件, 去除了空行

- ### u_excel
    + ExcelIllegalCharactersRe 非法字符串正则匹配器
    + ExcelSheet  分页数据对象
    + func (obj *ExcelSheet) Len() int 获取数据行数
    + func (obj *ExcelSheet) SetSafeLimit(n int) 设置单页安全数据行数上限值(最大值=1048576)
    + func (obj *ExcelSheet) Safe() []ExcelSheet 将分页数据对象进行安全转换
    + func MakeExcelFp(data ...ExcelSheet) (*excelize.File, error) 创建Excel文件对象
    + func SafeMakeExcelFp(data ...ExcelSheet) (*excelize.File, error) 安全的创建Excel文件对象

- ### u_file
    + func PathOk(path string) (bool, error) 判断文件或者文件夹路径是否正常
    + func ValidFileUTF8(filepath string, checkLines int) (bool, error) 验证文件是否能被UTF8解码
    + func ListDirFiles(dirPath, suffix string) ([]string, error) 获取目录下是所有文件的绝对路径(不含文件夹, 并且可以通过suffix过滤, 当suffix为空字符串或者"*"时表示配匹所有文件尾缀)
    
- ### u_jwt
    + type CustomJWTClaims struct 自定义JWT数据载体
    + func CreateJWTToken(claims CustomJWTClaims, salt []byte) (string, error) 创建JWT-Token字符串
    + func ParseJWTToken(tokenStr string, salt []byte) (*CustomJWTClaims, error) 解析JWT-Token字符串
    + func RefreshJWTToken(tokenStr string, salt []byte, survivalTime time.Duration) (string, error) 刷新JWT-Token字符串;

- ### u_logger
    + type Level int 日志级别类型
    + type XLogger struct 日志记录器结构体, 继承了标准库的log.Logger
    + func (obj *XLogger) SetLevel(level Level) 设置日志级别
    + func (obj *XLogger) Level() Level 获取日志级别
    + func (obj *XLogger) Debug(msg ...interface{}) 输出Debug级别的日志
    + func (obj *XLogger) Info(msg ...interface{}) 输出Info级别的日志
    + func (obj *XLogger) Warn(msg ...interface{}) 输出Warn级别的日志
    + func (obj *XLogger) Error(msg ...interface{}) 输出Error级别的日志
    + func GetLogger() *XLogger 获取日志记录器对象,单例模式,默认格式与输出
    
- ### u_rotate_file
    + type RotateFileWriter struct 循环文件写入器, 可以帮助我们自己实现循环文件日志
    + func (obj *RotateFileWriter) Write(p []byte) (n int, err error) 往文件写入数据
    + func NewRotateFileWriter(fileName, dirPath string, maxCount int, maxSize int64) *RotateFileWriter 创建新的循环文件写入对象;
  
- ### u_string
    + func BytesMD5Hash(data []byte) string 获取字节数组的MD5签名
    + func StringMD5Hash(data string) string 获取字符串的MD5签名
    + func MultiStringMD5Hash(data ...string) string 获取多个字符串的MD5签名
    + func StringContainsHan(data string) bool 检查字符串是否包含了中文字符
    + func StringContainsSpace(data string) 检查字符串是否包含了空白字符
    + func SafeSliceString(data string, start, end int) (string, error) 安全的字符串切片
  
- ### u_time
    + BJS 北京时间时区
    + func NowTimestamp() int64 当前时间戳
    + func Timestamp2Datetime(timestamp int64, tz *time.Location) time.Time 时间戳转时间对象
    + func Timestamp2DatetimeStr(timestamp int64, tz *time.Location) string 时间戳转时间字符串
    + func BJSNowDatetimeStr() string 当前北京时间字符串
    + func UTCNowDatetimeStr() string 当前UTC时间字符串
    + func BJSTodayDateStr() string 当前北京时间日期字符串
    + func UTCTodayDateStr() string 当前UTC时间日期字符串
    + func Time2BJS(value time.Time) time.Time 时间对象转换到北京时区;
    
    
  

