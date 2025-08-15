package gerror_test

import (
    "encoding/json"
    "errors"
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/camry/g/v2/gerrors/gcode"
    "github.com/camry/g/v2/gerrors/gerror"
)

func nilError() error {
    return nil
}

func Test_Nil(t *testing.T) {
    assert.NotNil(t, gerror.New(""))
    assert.Nil(t, gerror.Wrap(nilError(), "test"))
}

func Test_New(t *testing.T) {
    err1 := gerror.New("1")
    assert.NotNil(t, err1)
    assert.Equal(t, err1.Error(), "1")

    err2 := gerror.Newf("%d", 1)
    assert.NotNil(t, err2)
    assert.Equal(t, err2.Error(), "1")

    err3 := gerror.NewSkipf(1, "%d", 1)
    assert.NotNil(t, err3)
    assert.Equal(t, err3.Error(), "1")
}

func Test_Wrap(t *testing.T) {
    err1 := errors.New("1")
    err1 = gerror.Wrap(err1, "2")
    err1 = gerror.Wrap(err1, "3")
    assert.NotNil(t, err1)
    assert.Equal(t, err1.Error(), "3: 2: 1")

    err2 := gerror.New("1")
    err2 = gerror.Wrap(err2, "")
    assert.NotNil(t, err2)
    assert.Equal(t, err2.Error(), "1")
}

func Test_Wrapf(t *testing.T) {
    err1 := errors.New("1")
    err1 = gerror.Wrapf(err1, "%d", 2)
    err1 = gerror.Wrapf(err1, "%d", 3)
    assert.NotNil(t, err1)
    assert.Equal(t, err1.Error(), "3: 2: 1")

    err2 := gerror.New("1")
    err2 = gerror.Wrapf(err2, "")
    assert.NotNil(t, err2, nil)
    assert.Equal(t, err2.Error(), "1")
}

func Test_WrapSkip(t *testing.T) {
    err1 := errors.New("1")
    err1 = gerror.WrapSkip(1, err1, "2")
    err1 = gerror.WrapSkip(1, err1, "3")
    assert.NotNil(t, err1, nil)
    assert.Equal(t, err1.Error(), "3: 2: 1")

    err2 := gerror.New("1")
    err2 = gerror.WrapSkip(1, err2, "")
    assert.NotNil(t, err2, nil)
    assert.Equal(t, err2.Error(), "1")
}

func Test_WrapSkipf(t *testing.T) {
    err1 := errors.New("1")
    err1 = gerror.WrapSkipf(1, err1, "2")
    err1 = gerror.WrapSkipf(1, err1, "3")
    assert.NotNil(t, err1, nil)
    assert.Equal(t, err1.Error(), "3: 2: 1")

    err2 := gerror.New("1")
    err2 = gerror.WrapSkipf(1, err2, "")
    assert.NotNil(t, err2, nil)
    assert.Equal(t, err2.Error(), "1")
}

func Test_Cause(t *testing.T) {
    err := errors.New("1")
    assert.Equal(t, gerror.Cause(err), err)

    err1 := errors.New("1")
    err1 = gerror.Wrap(err1, "2")
    err1 = gerror.Wrap(err1, "3")
    assert.Equal(t, gerror.Cause(err1).Error(), "1")

    err2 := gerror.New("1")
    assert.Equal(t, gerror.Cause(err2).Error(), "1")

    err3 := gerror.New("1")
    err3 = gerror.Wrap(err3, "2")
    err3 = gerror.Wrap(err3, "3")
    assert.Equal(t, gerror.Cause(err3).Error(), "1")
}

func Test_Format(t *testing.T) {
    err1 := errors.New("1")
    err1 = gerror.Wrap(err1, "2")
    err1 = gerror.Wrap(err1, "3")
    assert.NotNil(t, err1)
    assert.Equal(t, fmt.Sprintf("%s", err1), "3: 2: 1")
    assert.Equal(t, fmt.Sprintf("%v", err1), "3: 2: 1")

    err2 := gerror.New("1")
    err2 = gerror.Wrap(err2, "2")
    err2 = gerror.Wrap(err2, "3")
    assert.NotNil(t, err2, nil)
    assert.Equal(t, fmt.Sprintf("%-s", err2), "3")
    assert.Equal(t, fmt.Sprintf("%-v", err2), "3")
}

func Test_Stack(t *testing.T) {
    err := errors.New("1")
    assert.Equal(t, fmt.Sprintf("%+v", err), "1")

    err1 := errors.New("1")
    err1 = gerror.Wrap(err1, "2")
    err1 = gerror.Wrap(err1, "3")
    assert.NotNil(t, err1, nil)
    // fmt.Printf("%+v", err1)

    err2 := gerror.New("1")
    assert.NotNil(t, fmt.Sprintf("%+v", err2), "1")
    // fmt.Printf("%+v", err2)

    err3 := gerror.New("1")
    err3 = gerror.Wrap(err3, "2")
    err3 = gerror.Wrap(err3, "3")
    assert.NotNil(t, err3, nil)
    // fmt.Printf("%+v", err3)
}

func Test_Current(t *testing.T) {
    err := errors.New("1")
    err = gerror.Wrap(err, "2")
    err = gerror.Wrap(err, "3")
    assert.Equal(t, err.Error(), "3: 2: 1")
    assert.Equal(t, gerror.Current(err).Error(), "3")
}

func Test_Unwrap(t *testing.T) {
    err := errors.New("1")
    err = gerror.Wrap(err, "2")
    err = gerror.Wrap(err, "3")
    assert.Equal(t, err.Error(), "3: 2: 1")

    err = gerror.Unwrap(err)
    assert.Equal(t, err.Error(), "2: 1")

    err = gerror.Unwrap(err)
    assert.Equal(t, err.Error(), "1")

    err = gerror.Unwrap(err)
    assert.Nil(t, err)
}

func Test_Code(t *testing.T) {
    err1 := errors.New("123")
    assert.Equal(t, gerror.Code(err1).Code(), -1)
    assert.Equal(t, err1.Error(), "123")

    err2 := gerror.NewCode(gcode.CodeUnknown, "123")
    assert.Equal(t, gerror.Code(err2), gcode.CodeUnknown)
    assert.Equal(t, err2.Error(), "123")

    err3 := gerror.NewCodef(gcode.New(1, "", nil), "%s", "123")
    assert.Equal(t, gerror.Code(err3).Code(), 1)
    assert.Equal(t, err3.Error(), "123")

    err4 := gerror.NewCodeSkip(gcode.New(1, "", nil), 0, "123")
    assert.Equal(t, gerror.Code(err4).Code(), 1)
    assert.Equal(t, err4.Error(), "123")

    err5 := gerror.NewCodeSkipf(gcode.New(1, "", nil), 0, "%s", "123")
    assert.Equal(t, gerror.Code(err5).Code(), 1)
    assert.Equal(t, err5.Error(), "123")

    err6 := errors.New("1")
    err6 = gerror.Wrap(err6, "2")
    err6 = gerror.WrapCode(gcode.New(1, "", nil), err6, "3")
    assert.Equal(t, gerror.Code(err6).Code(), 1)
    assert.Equal(t, err6.Error(), "3: 2: 1")

    err7 := errors.New("1")
    err7 = gerror.Wrap(err7, "2")
    err7 = gerror.WrapCodef(gcode.New(1, "", nil), err7, "%s", "3")
    assert.Equal(t, gerror.Code(err7).Code(), 1)
    assert.Equal(t, err7.Error(), "3: 2: 1")

    err8 := errors.New("1")
    err8 = gerror.Wrap(err8, "2")
    err8 = gerror.WrapCodeSkip(gcode.New(1, "", nil), 100, err8, "3")
    assert.Equal(t, gerror.Code(err8).Code(), 1)
    assert.Equal(t, err8.Error(), "3: 2: 1")

    err9 := errors.New("1")
    err9 = gerror.Wrap(err9, "2")
    err9 = gerror.WrapCodeSkipf(gcode.New(1, "", nil), 100, err9, "%s", "3")
    assert.Equal(t, gerror.Code(err9).Code(), 1)
    assert.Equal(t, err9.Error(), "3: 2: 1")
}

func Test_SetCode(t *testing.T) {
    err := gerror.New("123")
    assert.Equal(t, gerror.Code(err).Code(), -1)
    assert.Equal(t, err.Error(), "123")

    err.(*gerror.Error).SetCode(gcode.CodeValidationFailed)
    assert.Equal(t, gerror.Code(err), gcode.CodeValidationFailed)
    assert.Equal(t, err.Error(), "123")
}

func Test_Json(t *testing.T) {
    err := gerror.Wrap(gerror.New("1"), "2")
    b, e := json.Marshal(err)
    assert.Equal(t, e, nil)
    assert.Equal(t, string(b), `"2: 1"`)
}

func Test_HasStack(t *testing.T) {
    err1 := errors.New("1")
    err2 := gerror.New("1")
    assert.Equal(t, gerror.HasStack(err1), false)
    assert.Equal(t, gerror.HasStack(err2), true)
}

func Test_Equal(t *testing.T) {
    err1 := errors.New("1")
    err2 := errors.New("1")
    err3 := gerror.New("1")
    err4 := gerror.New("4")
    assert.Equal(t, gerror.Equal(err1, err2), false)
    assert.Equal(t, gerror.Equal(err1, err3), true)
    assert.Equal(t, gerror.Equal(err2, err3), true)
    assert.Equal(t, gerror.Equal(err3, err4), false)
    assert.Equal(t, gerror.Equal(err1, err4), false)

}

func Test_Is(t *testing.T) {
    err1 := errors.New("1")
    err2 := gerror.Wrap(err1, "2")
    err2 = gerror.Wrap(err2, "3")
    assert.Equal(t, gerror.Is(err2, err1), true)

}

func Test_HashError(t *testing.T) {
    err1 := errors.New("1")
    err2 := gerror.Wrap(err1, "2")
    err2 = gerror.Wrap(err2, "3")
    assert.Equal(t, gerror.HasError(err2, err1), true)

}

func Test_HashCode(t *testing.T) {
    err1 := errors.New("1")
    err2 := gerror.WrapCode(gcode.CodeNotAuthorized, err1, "2")
    err3 := gerror.Wrap(err2, "3")
    err4 := gerror.Wrap(err3, "4")
    assert.Equal(t, gerror.HasCode(err1, gcode.CodeNotAuthorized), false)
    assert.Equal(t, gerror.HasCode(err2, gcode.CodeNotAuthorized), true)
    assert.Equal(t, gerror.HasCode(err3, gcode.CodeNotAuthorized), true)
    assert.Equal(t, gerror.HasCode(err4, gcode.CodeNotAuthorized), true)
}
