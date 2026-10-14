package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 bool; F1 int16 }
func genFn0(p0 int32, p1 string, p2 int64) uint16 {
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int8 = -1
        p0 = p0
        _ = fi7
        _ = v8
    }
    var v9 int = 32768
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 2
}
func genFn1(p0 int64, p1 uint32) string {
    if (42 == 1) {
        v10 := GenS1{F0: -128}
        var v11 string = ("foo bar" + ("foo bar" + "\n\t"))
        var v12 int8 = -72
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := []int{-2147483648}
    if true {
        p0 = (p0 ^ (p0 << ((4294967295) & 63)))
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = v13
    return "a"
}

func main() {
    var v1 uint16 = 65535
    var v2 int64 = 9223372036854775807
    var v3 bool = false
    v4 := []uint32{540, 2147483647, 763}
    var v5 int32 = 2147483647
    var v6 string = "\n\t"
    for fi14 := 0; fi14 < 19; fi14++ {
        if (!(-2.25 != 3.14159)) {
            v1 += (^(v1))
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 float32 = float32(3.14159)
            v17 := GenS1{F0: -91}
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    var v18 string = "a"
    for fi19 := 0; fi19 < 7; fi19++ {
        if (v3 || true) {
            v20 := GenS1{F0: 1}
            _ = v20
        }
        v4 = append(v4, 42)
        fmt.Println(v6)
        _ = fi19
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v21 int32 = -557
    var v22 uint16 = v1
    v23 := GenS2{F0: 128, F1: 42}
    switch 127 {
    case -986:
        var v24 uint = 2
        _ = v24
    default:
        v25 := GenS3{F0: v3, F1: 1}
        _ = v25
    }
    var v26 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
