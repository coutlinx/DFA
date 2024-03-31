# DFA敏感词检测算法
### 在运行本项目前，请确保你的本地包中有gorm和xlsx包，以下是安装命令
```shell
 go get "github.com/tealeg/xlsx"
 
 go get "gorm.io/gorm"
 
 #或者直接
 go mod tidy
```

## 使用说明
### SearchWords方法
>  ``` go 
>  // SearchWords DAF算法检查敏感词
> func SearchWords(text string) bool {
>   return SensitiveTire.search(text)
> }
> // text为一句话或者一个词，如果里面包含敏感词就返回true，不包括则就false
>  ```
### SearchWordsOne方法
>  ``` go 
> // SearchWordsOne 查询第一个敏感词
> func SearchWordsOne(text string) (string, bool) {
> return SensitiveTire.searchOne(text)
> }
> // text为一句话或者一个词，如果里面包含敏感词就返回查到的第一个敏感词和true，不包括则就返回空字符串和false
>  ```

### SearchWordAll方法
> ``` go
> // SearchWordsAll查询全部敏感词
> func SearchWordsAll(text string) []string {
> 	return SensitiveTire.searchAll(text)
> }
> > // text为一句话或者一个词，如果里面包含敏感词就返回查到的全部敏感词，不包括则就返回空数组
> ```
> 
> 