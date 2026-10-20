package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint64 }
func genFn0() int16 {
    var v4 uint = 2
    v4 = 388
    _ = v4
    return -1
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float64 = 3.14159
    v3 := []uint32{619}
    v1 -= v1
    v5 := GenS1{F0: -655, F1: 18446744073709551615, F2: 4294967295}
    fmt.Println(v1, v1)
    v1 = (v1 * v1)
    fmt.Println(v2, v1)
    var v6 uint64 = 4294967295
    switch -2147483648 {
    case 127:
        v2 = ((v2 + v2) - (-2.25 + float64(3.14159)))
    case 128:
        v3 = append(v3, 1)
    default:
        var v7 float64 = 1e10
        _ = v7
    }
    var v8 int16 = 42
    v1 -= (v1 << ((18446744073709551615) & 63))
    if ((("hello" + ("hello" + "hello")) + "hello") < ("foo bar" + "foo bar")) {
        if false {
            fmt.Println(v8, v1, v8)
            fmt.Println(v8, v6, v2)
        } else {
            var v9 string = "foo bar"
            if len(v3) > 0 {
                v3[0] = 65535
            }
            switch -9223372036854775808 {
            case 169:
                var v10 int16 = -735
                _ = v10
            case -2147483648:
                var v11 uint32 = 432
                _ = v11
            default:
                var v12 int = -608
                _ = v12
            }
            _ = v9
        }
    }
    if (!(false || false)) {
        v13 := GenS1{F0: (v8 | v8), F1: 65535, F2: v6}
        var v14 int64 = -853
        _ = v13
        _ = v14
    }
    var v15 int = 260
    switch v15 {
    case 32768:
        var v16 bool = ((!false) || ("foo bar" > "foo bar"))
        _ = v16
    default:
        var v17 float32 = ((0.0 / 1.0) + 0.1)
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v8
}
