package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint) bool {
    v4 := GenS1{F0: 32767}
    if (true && ((("hello" + "foo bar") + "\n\t") > "hello")) {
        p0 = (p0 >> (((p0 * 488)) & 63))
        var v5 uint = (p0 | (p0 * p0))
        v6 := GenS2{F0: "hello"}
        _ = v5
        _ = v6
    } else {
        fmt.Println(p0)
        p0 += p0
        var v7 float32 = 0.1
        _ = v7
    }
    _ = p0
    _ = v4
    return ((false || true) || true)
}

func main() {
    v1 := []string{"hello", "\n\t"}
    var v2 string = ("hello" + "")
    var v3 float32 = (-2.25 - 0.0)
    var v8 int64 = 0
    fmt.Println(v8, v8)
    fmt.Println(v3)
    _ = len(v1)
    v1 = append(v1, (v2 + v2))
    if (((v8 / ((-1) | 1)) == (v8 * -1)) || false) {
        if (true && true) {
            var v9 int32 = -2147483648
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            _ = v9
        } else {
            v8 -= (v8 ^ (-(v8)))
            var v10 int8 = -128
            _ = v10
        }
        fmt.Println(v3)
    } else {
        var v11 uint8 = 1
        v3 += v3
        _ = v11
    }
    v12 := GenS3{F0: 0}
    v8 += -9223372036854775808
    var v13 uint64 = 343
    switch 32767 {
    case -2147483648:
        v1 = append(v1, v2)
    default:
        v2 = v2
    }
    var v14 int16 = 32767
    var v15 bool = false
    fmt.Println(v14, v8, v3)
    var v16 int16 = v14
    v14 = 1000
    v13 = (v13 >> ((421) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
}
