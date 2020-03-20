package twse

var (
	twse *Service
)

func init() {
	client := GetClient()

	var err error
	twse, err = New(client)
	if err != nil {
		panic(err)
	}
}
