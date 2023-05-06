# errors
简单易用、可携带业务错误码的错误处理包.  基于 `pkg/errors` 加入了更丰富的功能，更适用在微服务体系下.

# quick start
1. cd ./example
2. go run ./cmd
3. 如下请求查看返回值和日志信息
   ```shell
   // 参数错误
   curl --location --request GET 'http://127.0.0.1:8080/user?uid=a'

   // 业务错误
   curl --location --request GET 'http://127.0.0.1:8080/user?uid=1'
   curl --location --request GET 'http://127.0.0.1:8080/user?uid=99'

   // 内部错误
   curl --location --request GET 'http://127.0.0.1:8080/user?uid=10000'
   ```

# errors generate
1. 安装 `go install github.com/go-leo/errors/codegen`
2. 定义错误码文件
   ```go
    package code
    //go:generate codegen -type=int

    // base: base errors.
    const (
        // ErrUnknown - 500: Internal server error.
        ErrUnknown int = iota + 100001

        // ErrBind - 400: Error occurred while binding the request body to the struct.
        ErrBind

        // ErrValidation - 400: Validation failed.
        ErrValidation
    )
   ```
3. 实现错误注册方法
   ```go
    package code

    import (
        "net/http"

        "github.com/go-leo/errors"
        "golang.org/x/exp/slices"
    )

    // ErrCode implements `panda/pkg/errors`.Coder interface.
    type ErrCode struct {
        // C refers to the code of the ErrCode.
        C int `json:"code,omitempty"`

        // HTTP status that should be used for the associated error code.
        HTTP int `json:"http,omitempty"`

        // External (user) facing error text.
        Ext string `json:"msg,omitempty"`

        // Ref specify the reference document.
        Ref string `json:"ref,omitempty"`
    }

    var _ errors.Coder = &ErrCode{}

    // Code returns the integer code of ErrCode.
    func (coder ErrCode) Code() int {
        return coder.C
    }

    // String implements stringer. String returns the external error message,
    // if any.
    func (coder ErrCode) String() string {
        return coder.Ext
    }

    // Reference returns the reference document.
    func (coder ErrCode) Reference() string {
        return coder.Ref
    }

    // HTTPStatus returns the associated HTTP status code, if any. Otherwise,
    // returns 200.
    func (coder ErrCode) HTTPStatus() int {
        if coder.HTTP == 0 {
            return http.StatusInternalServerError
        }

        return coder.HTTP
    }

    //nolint: unparam // .
    func register(code int, httpStatus int, message string, refs ...string) {
        found := slices.Contains([]int{200, 400, 401, 403, 404, 500}, httpStatus)
        if !found {
            panic("http code not in `200, 400, 401, 403, 404, 500`")
        }

        var reference string
        if len(refs) > 0 {
            reference = refs[0]
        }

        coder := &ErrCode{
            C:    code,
            HTTP: httpStatus,
            Ext:  message,
            Ref:  reference,
        }

        errors.MustRegister(coder)
    }
   ```
4. 生成错误码文件
    ```shell
    go generate ./...
    ```