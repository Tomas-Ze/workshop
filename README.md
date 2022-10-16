# workshop

## Used packages
`go get github.com/go-chi/chi/v5`

`go get github.com/stretchr/testify/assert`


## Table testing
```
import tAssert "github.com/stretchr/testify/assert

func TestTableDriven(t *testing.T) {
	cases := []struct {
		Name string

		GivenValue int
		Expected   int
	}{
		{
			Name:       "value one",
			GivenValue: 1,
			Expected:   1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			assert := tAssert.New(t)
			assert.Equal(tt.Expected, tt.GivenValue)
		})
	}
}
```

## Request and response mock
```
responseWriterMock := &mocks.ResponseWriter{}

req := &http.Request{
    Body: io.NopCloser(bytes.NewBuffer(nil)),
}
```
