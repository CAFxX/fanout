package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint32 }
func genFn0() uint {
    var v5 uint8 = 2
    v5 -= (v5 ^ v5)
    v5 -= v5
    _ = v5
    return 206
}
func genFn1(p0 uint8, p1 string, p2 string) string {
    p2 = ""
    switch p1 {
    case "hello":
        var v6 int64 = 42
        _ = v6
    case "a":
        for fi7 := 0; fi7 < 4; fi7++ {
            v8 := GenS1{F0: 42, F1: 2147483647, F2: -251}
            var v9 float32 = (1.5 + 1.5)
            _ = fi7
            _ = v8
            _ = v9
        }
    case "\n\t":
        p1 = (("foo bar" + (p2 + (p2 + p2))) + p1)
    default:
        v10 := []int{32767, 0, 127, -109}
        _ = v10
    }
    var v11 uint32 = 42
    if (635 <= -165) {
        fmt.Println(p1)
    }
    v11 += v11
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p2
}

func main() {
    var v1 int32 = 42
    var v2 uint = 42
    var v3 uint32 = 2147483647
    var v4 int8 = -128
    switch (("hello" + (("hello" + ("hello" + "\n\t")) + "")) + "") {
    case "foo bar":
        fmt.Println(v2)
    default:
        fmt.Println(v3)
    }
    v1 = (-(v1))
    var v12 int32 = v1
    var v13 uint16 = 255
    var v14 float64 = 0.1
    var v15 uint = (^(v2))
    v14 = v14
    v16 := GenS2{F0: 42}
    v17 := []int{}
    if (((1 != 2) || ((true || true) && ((v4 & (^(v4))) == (^(v4))))) && true) {
        var v18 uint32 = v3
        _ = v18
    }
    if (!(3.14159 < 1e10)) {
        if true {
            v17 = append(v17, -2147483648)
            v19 := GenS3{F0: (v13 >> ((2) & 63)), F1: 1, F2: (v3 & (v3 / (((v3 * 2)) | 1)))}
            _ = v19
        }
    } else {
        v3 = (v3 | 1)
        v2 = v2
    }
    if len(v17) > 0 {
        v17[0] = 9223372036854775807
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        v21 := GenS2{F0: 0}
        v4 = v4
        v15 -= 255
        _ = fi20
        _ = v21
    }
    fmt.Println(v4, v4, v3)
    if false {
        fmt.Println(v14)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
