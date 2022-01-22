package logging

type HttpLog struct {
	Headers      map[string][]string
	RequestBody  string
	ResponseBody string
	IP           string
	Method       string
	Status       int
	URI          string
}
