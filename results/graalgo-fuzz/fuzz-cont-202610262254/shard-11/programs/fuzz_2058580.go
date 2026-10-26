package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint8 }
func genFn0(p0 int32, p1 uint64, p2 int64) uint64 {
    v6 := []int{399, -2147483648, 42, 872}
    if true {
        var v7 int32 = int32(p1)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v7
    } else {
        p0 = p0
        v8 := GenS1{F0: "a", F1: p1, F2: 128}
        p0 = 1
        _ = v8
    }
    switch "a" {
    case "foo bar":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "":
        p1 = p1
    default:
        switch 0 {
        case -590:
            var v9 uint = 295
            _ = v9
        default:
            var v10 uint32 = 65535
            _ = v10
        }
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p1 / (((p1 / (((^(p1))) | 1))) | 1))
}
func genFn1(p0 int32) bool {
    fmt.Println(p0)
    v11 := []int{-9223372036854775808, -1, 978}
    switch 32768 {
    case -9223372036854775808:
        switch int(p0) {
        case 582:
            var v12 uint16 = 32767
            _ = v12
        case 4294967295:
            fmt.Println(p0)
        default:
            p0 = (p0 + 0)
        }
    }
    var v13 bool = false
    v13 = ((767 != 0) || v13)
    _ = p0
    _ = v11
    _ = v13
    return (((false || false) && v13) && ((1 >= -13) || (2 > 1)))
}

func main() {
    var v1 uint8 = 0
    var v2 float64 = (float64(1.5) - (0.0 / float64(-2.25)))
    v3 := []int{-2147483648, 128, -2147483648}
    var v4 float32 = ((3.14159 + float32(float64(float32(0.0)))) - (3.14159 + (1.5 / 3.14159)))
    var v5 string = ("\n\t" + "foo bar")
    v1 -= v1
    fmt.Println(v4, v5, v5)
    fmt.Println(v2, v4)
    var v14 uint32 = 521
    var v15 int64 = 127
    for fi16 := 0; fi16 < 7; fi16++ {
        v3 = append(v3, fi16)
        fmt.Println(v15)
        fmt.Println(v1, v14)
        _ = fi16
    }
    var v17 uint32 = 132
    var v18 uint32 = v17
    switch 42 {
    case 42:
        v5 = ("" + v5)
    case 757:
        if true {
            switch 128 {
            case 316:
                fmt.Println(v5, v2, v15)
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            var v19 uint8 = 54
            _ = v19
        } else {
            switch "" {
            case "\n\t":
                for fi20 := 0; fi20 < 16; fi20++ {
                    for fi21 := 0; fi21 < 10; fi21++ {
                        v15 -= v15
                        _ = fi21
                    }
                    _ = fi20
                }
            case "":
                fmt.Println(v2, v4, v5)
            case "hello":
                var v22 int64 = (v15 & v15)
                _ = v22
            default:
                var v23 uint16 = 573
                _ = v23
            }
        }
    default:
        var v24 float32 = 0.1
        _ = v24
    }
    var v25 float32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
