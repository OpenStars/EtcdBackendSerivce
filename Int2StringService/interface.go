package Int2StringService

type Int2StringServiceIf interface {
	PutData(key int64, value string) error
	GetData(key int64) (string, error)
}
