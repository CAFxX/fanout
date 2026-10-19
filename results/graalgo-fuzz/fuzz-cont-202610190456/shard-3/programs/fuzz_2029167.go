package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 string) uint8 {
    var v7 float64 = ((1.5 / (1e10 + 0.1)) / 1e10)
    var v8 uint32 = 0
    var v9 string = p0
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 255
}
func genFn1() uint16 {
    var v10 int64 = 42
    v11 := GenS1{F0: 0, F1: v10, F2: 42}
    _ = v10
    _ = v11
    return 65535
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 100
    var v3 int8 = v2
    var v4 int32 = 826
    var v5 uint32 = 2147483648
    var v6 int16 = -803
    v4 = 1
    if false {
        v12 := GenS1{F0: 5, F1: 151, F2: (v6 << ((42) & 63))}
        fmt.Println(v4)
        _ = v12
    } else {
        for fi13 := 0; fi13 < 12; fi13++ {
            fmt.Println(v3, v6, v4)
            _ = fi13
        }
        delete(v1, -9223372036854775808)
        var v14 uint = 42
        _ = v14
    }
    var v15 uint32 = v5
    fmt.Println(v2, v15, v3)
    v5 = (v5 + 42)
    v15 = v15
    fmt.Println(v2, v5)
    v15 -= v15
    v16 := []int{32767, 828, -1}
    var v17 int32 = v4
    v1[27] = "\n\t"
    v18 := GenS1{F0: 127, F1: 232, F2: int16(v5)}
    fmt.Println(v4, v15)
    var v19 uint = 135
    for fi20 := 0; fi20 < 16; fi20++ {
        switch (("a" + "foo bar") + "") {
        case "":
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 int16 = 960
                v23 := GenS1{F0: 128, F1: 1, F2: (v6 - v6)}
                var v24 int64 = 9223372036854775807
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
        default:
            v2 = (^(v2))
        }
        var v25 float64 = 3.14159
        _ = len(v1)
        _ = fi20
        _ = v25
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        if ((!false) || ((v17 & v4) != v17)) {
            v1[fi26] = "hello"
        }
        _ = fi26
    }
    v27 := GenS2{F0: 2}
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
