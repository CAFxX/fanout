package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint16 }
func genFn0(p0 uint32) uint16 {
    var v6 int16 = 42
    switch "\n\t" {
    case "a":
        var v7 uint32 = (p0 % ((p0) | 1))
        _ = v7
    default:
        if (float64(float32(float64(float32(float64((1e10 + -2.25)))))) <= (1.5 + 1e10)) {
            var v8 bool = false
            for fi9 := 0; fi9 < 8; fi9++ {
                p0 = 2
                p0 += 183
                var v10 int32 = 756
                _ = fi9
                _ = v10
            }
            fmt.Println(v6, v8)
            _ = v8
        } else {
            v6 = v6
        }
    }
    var v11 string = "hello"
    for fi12 := 0; fi12 < 16; fi12++ {
        p0 += 459
        v6 -= v6
        var v13 uint8 = 0
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v11
    _ = v6
    return 862
}

func main() {
    var v1 string = "hello"
    var v2 uint8 = 255
    v3 := []string{"foo bar", v1}
    var v4 int16 = -678
    var v5 int8 = 100
    var v14 uint16 = genFn0(2)
    v15 := GenS1{F0: (^(v14)), F1: -2147483648, F2: genFn0(65535)}
    v14 = 32767
    if true {
        _ = len(v3)
    } else {
        fmt.Println(v14, v5)
        fmt.Println(v14, v5, v4)
    }
    var v16 int8 = 1
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 int64 = 0
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        if (false && false) {
            if (42 != 4294967295) {
                if (v1 < v1) {
                    fmt.Println(v4, v16, v1)
                    v16 -= v16
                    var v20 uint32 = 2
                    _ = v20
                }
            } else {
                var v21 uint8 = (v2 | (v2 / (((v2 / ((v2) | 1))) | 1)))
                fmt.Println(fi19, v5, v16)
                v14 -= (^(v14))
                _ = v21
            }
            var v22 int64 = -1
            _ = v22
        } else {
            v1 = v1
        }
        var v23 uint16 = 65535
        var v24 int64 = 9223372036854775807
        _ = fi19
        _ = v23
        _ = v24
    }
    v5 -= (v5 - (v16 / (((v5 + 124)) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
