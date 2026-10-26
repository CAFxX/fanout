package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint16 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int; F1 int32 }
func genFn0() uint16 {
    var v6 uint16 = 2
    var v7 int16 = 32767
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 uint32, p1 int32, p2 uint64) string {
    var v8 uint32 = 2
    var v9 uint32 = (p0 % ((v8) | 1))
    v10 := GenS1{F0: (p1 & -2147483648), F1: 32767, F2: genFn0()}
    fmt.Println(p1, p2, p2)
    switch 42 {
    case 42:
        var v11 int = -9223372036854775808
        _ = v11
    default:
        var v12 bool = (false || (((("hello" + "foo bar") + ((("" + (("a" + "") + "foo bar")) + "foo bar") + ("foo bar" + ""))) + "foo bar") < "a"))
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return ("" + ("a" + ("a" + "hello")))
}

func main() {
    v1 := []string{("hello" + ""), "hello", ""}
    var v2 uint64 = 740
    v3 := []int{186, -39, 2147483647, 127}
    var v4 uint32 = 2147483647
    var v5 float32 = 0.0
    v13 := GenS1{F0: 42, F1: 0, F2: 1}
    var v14 uint32 = 2
    fmt.Println(v4, v5)
    if (true && (int8(v14) >= 127)) {
        fmt.Println(v4)
    }
    switch "" {
    case "foo bar":
        switch "hello" {
        case "foo bar":
            _ = len(v1)
        default:
            for fi15 := 0; fi15 < 10; fi15++ {
                fmt.Println(v5, v2, v4)
                v16 := GenS2{F0: "a", F1: false}
                var v17 uint = 255
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    default:
        v4 += v14
    }
    v18 := GenS3{F0: -281, F1: 0}
    var v19 uint64 = (v2 << ((uint(v14)) & 63))
    if true {
        for fi20 := 0; fi20 < 4; fi20++ {
            if (true || false) {
                var v21 uint16 = 422
                _ = v21
            } else {
                fmt.Println(fi20, v14)
                v5 -= 1e10
                var v22 uint32 = v4
                _ = v22
            }
            v2 += uint64(v4)
            v23 := GenS1{F0: -2147483648, F1: 1, F2: genFn0()}
            _ = fi20
            _ = v23
        }
    } else {
        var v24 int16 = 972
        _ = v24
    }
    v2 -= v19
    var v25 bool = ((4294967295 == -2147483648) && (true || false))
    for fi26 := 0; fi26 < 1; fi26++ {
        fmt.Println(v2)
        if len(v3) > 0 {
            v3[0] = 32767
        }
        _ = fi26
    }
    v19 += v19
    fmt.Println(v2, v25)
    var v27 int = 0
    var v28 int8 = 3
    v14 += v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
