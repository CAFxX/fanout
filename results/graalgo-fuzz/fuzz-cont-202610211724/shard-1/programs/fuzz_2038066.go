package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint64 }
func genFn0(p0 bool) int8 {
    var v7 int32 = -1
    var v8 float64 = 1e10
    v7 = v7
    _ = p0
    _ = v7
    _ = v8
    return -128
}
func genFn1(p0 int16, p1 int16, p2 uint8) uint64 {
    var v9 uint32 = 65535
    var v10 int16 = (p1 / ((p0) | 1))
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 int32 = 42
    var v2 int64 = -9223372036854775808
    var v3 bool = false
    var v4 int8 = 68
    var v5 bool = (v3 || (0 != 42))
    var v6 uint8 = 1
    v1 -= v1
    v11 := []int{int(v4)}
    if (true || (-1 != -32768)) {
        var v12 uint16 = 65535
        switch -763 {
        case -2147483648:
            v1 += (v1 * -741)
        case 32768:
            for fi13 := 0; fi13 < 20; fi13++ {
                fmt.Println(v1, v1)
                var v14 uint32 = uint32(v12)
                var v15 uint = 18446744073709551615
                _ = fi13
                _ = v14
                _ = v15
            }
        case 128:
            var v16 float32 = float32(0.1)
            _ = v16
        default:
            for fi17 := 0; fi17 < 18; fi17++ {
                v18 := GenS1{F0: -2147483648, F1: 0, F2: 212}
                _ = fi17
                _ = v18
            }
        }
        _ = v12
    }
    switch 144 {
    case 32767:
        fmt.Println(v4, v6, v5)
    default:
        for fi19 := 0; fi19 < 6; fi19++ {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = fi19
        }
    }
    switch 1 {
    case 431:
        v1 = 42
    case 127:
        fmt.Println(v3)
    case 458:
        _ = len(v11)
    }
    if false {
        v20 := GenS2{F0: int32(v6), F1: 1}
        var v21 float32 = float32((0.1 * 3.14159))
        _ = v20
        _ = v21
    }
    fmt.Println(v2, v1, v5)
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 float64 = ((0.0 / 0.1) - 1.5)
        for fi24 := 0; fi24 < 11; fi24++ {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = fi24
        }
        switch ("\n\t" + "foo bar") {
        case "":
            v11 = append(v11, (fi22 + (-(fi22))))
        }
        _ = fi22
        _ = v23
    }
    v25 := GenS3{F0: -1, F1: 107}
    fmt.Println(v6)
    for fi26 := 0; fi26 < 1; fi26++ {
        v4 = int8(v2)
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        switch "foo bar" {
        case "\n\t":
            if len(v11) > 0 {
                v11[0] = fi26
            }
        case "hello":
            for fi27 := 0; fi27 < 15; fi27++ {
                v4 -= -1
                var v28 bool = v3
                _ = fi27
                _ = v28
            }
        }
        _ = fi26
    }
    v4 = (v4 - int8(v1))
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
