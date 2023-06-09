# Неделя 1. ДЗ

Необходимо написать отдельный пакет, в котором будет реализован in-memory cache со следующими методами:

- `Set(key string, value interface{})` - запись значения `value` в кеш по ключу `key`
- `Get(key string)`
- `Delete(key)`

```go
func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId := cache.Get("userId")

	fmt.Println(userId)
}
```

Пакет должен быть импортируемой библиотекой, которую можно установить себе в проект с помощью `go get -u <название модуля>`.

💡 Также должен быть оформлен README файл с описанием проекта и инструкцией по применению.
