package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string }
func genFn0(p0 int8, p1 uint32, p2 uint16) uint16 {
    switch "" {
    case "":
        v6 := GenS1{F0: 18, F1: 885}
        _ = v6
    case "a":
        var v7 uint = 255
        _ = v7
    case "foo bar":
        switch 98 {
        case -724:
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := GenS2{F0: -2147483648}
                _ = fi8
                _ = v9
            }
        }
    default:
        p2 -= (^(p2))
    }
    switch "" {
    case "":
        var v10 int64 = -539
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 613
}
func genFn1(p0 int16, p1 int16, p2 uint64) bool {
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint16 = 0
        _ = fi11
        _ = v12
    }
    fmt.Println(p0)
    if true {
        var v13 int = 32767
        p1 += (p0 / ((p1) | 1))
        if true {
            var v14 int8 = -87
            p1 = 42
            p2 -= (p2 | 0)
            _ = v14
        }
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return true
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = "foo bar"
    var v3 uint32 = 888
    v4 := []int{int(v3), -148, -9223372036854775808, -2147483648}
    var v5 bool = ((0 < 0) || (v3 != (v3 / ((v3) | 1))))
    var v15 uint32 = (v3 + (v3 % ((65535) | 1)))
    switch 42 {
    case 32767:
        v15 += 65535
    case 127:
        fmt.Println(v2, v2, v5)
    case 514:
        fmt.Println(v2)
    }
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    v2 = "a"
    var v16 int8 = 42
    _ = len(v1)
    v16 = -128
    v5 = false
    for fi17 := 0; fi17 < 2; fi17++ {
        delete(v1, fi17)
        var v18 uint64 = 498
        _ = fi17
        _ = v18
    }
    v19 := GenS3{F0: ("\n\t" + "")}
    v15 = (^(v15))
    v20 := v1[42]
    v20 = ((9223372036854775807 == 0) || (((v20 && (!(((!(32767 != 1000)) && (v5 || v20)) && (808 == 32767)))) || (!true)) || v5))
    var v21 uint16 = 463
    v4 = append(v4, 32767)
    v15 -= uint32(v21)
    for fi22 := 0; fi22 < 4; fi22++ {
        if v5 {
            fmt.Println(v2, v15)
        }
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 int64 = 42
            _ = fi23
            _ = v24
        }
        fmt.Println(fi22, v15)
        _ = fi22
    }
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
