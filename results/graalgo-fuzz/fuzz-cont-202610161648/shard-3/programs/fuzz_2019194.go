package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 bool }
type GenS2 struct { F0 int8; F1 bool; F2 int8 }
func genFn0(p0 uint32, p1 uint32) uint16 {
    v7 := []int{9223372036854775807}
    switch "hello" {
    case "\n\t":
        fmt.Println(p0, p0)
    case "":
        var v8 float64 = 0.0
        _ = v8
    default:
        switch "\n\t" {
        case "foo bar":
            p1 = p0
        case "a":
            v9 := GenS1{F0: 4294967295, F1: -1, F2: (!((4294967295 < 0) && (("\n\t" + "") == "foo bar")))}
            _ = v9
        case "":
            fmt.Println(p1)
        }
    }
    fmt.Println(p1, p1)
    p0 = p0
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1(p0 uint8, p1 int16) uint64 {
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int = 2147483647
        _ = fi10
        _ = v11
    }
    v12 := []int{-127, 9, 9223372036854775807, 32768}
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 float32 = (0.0 * 3.14159)
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        if true {
            var v16 float32 = 3.14159
            v17 := GenS2{F0: -128, F1: (false && true), F2: -1}
            p0 += p0
            _ = v16
            _ = v17
        } else {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            p0 += p0
            _ = fi18
        }
        switch "" {
        case "foo bar":
            p0 = (^(p0))
        case "hello":
            for fi19 := 0; fi19 < 6; fi19++ {
                fmt.Println(p1, fi15)
                switch (fi15 & 32768) {
                case 486:
                    p0 += 55
                case -2147483648:
                    fmt.Println(fi15, p0)
                case 42:
                    v20 := GenS2{F0: -1, F1: (false && ((339 > 2147483648) && false)), F2: 127}
                    _ = v20
                }
                var v21 int16 = p1
                _ = fi19
                _ = v21
            }
        default:
            v22 := GenS2{F0: 100, F1: true, F2: -1}
            _ = v22
        }
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = v12
    return 0
}

func main() {
    var v1 int64 = -901
    v2 := []uint32{2147483647, 2, 0, 2}
    v3 := []uint32{42, 2147483648, 2147483648}
    v4 := []string{"\n\t", "\n\t", ("" + "hello")}
    var v5 uint32 = 2147483647
    var v6 uint16 = 65535
    switch 32768 {
    case 128:
        fmt.Println(v1, v1)
    case -2147483648:
        var v23 float64 = (-2.25 + -2.25)
        _ = v23
    }
    fmt.Println(v6)
    if len(v4) > 0 {
        v4[0] = "a"
    }
    var v24 uint32 = (^(v5))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v6, v5, v1)
    var v25 int8 = -27
    var v26 int16 = 224
    v27 := GenS2{F0: v25, F1: true, F2: v25}
    v25 -= v25
    v5 += 592
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
