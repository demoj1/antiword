# Зависимости
Установка ubuntu: `apt-get install antiword`

# Что это такое?
Http сервер для конвертации doc файлов в простой текст.

За подробностями идем [читать](http://www.winfield.demon.nl/).
# Как использовать?
Ловит post запросы на корень с параметром `data`.
Параметр `data` должен содержать файл формата `doc`.

Код для отправки файла:
```go
b, _ := ioutil.ReadFile("test.doc")

r, _ := http.PostForm(host, url.Values{
	"data": []string{string(b)},
})

ans, _ := ioutil.ReadAll(r.Body)
fmt.Println(string(ans))
```
В ответ приходит текст из файла.