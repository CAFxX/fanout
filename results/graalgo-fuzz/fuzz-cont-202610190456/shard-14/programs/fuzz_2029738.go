package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int64; F1 uint64 }
func genFn0() uint64 {
    v6 := []int{1, -1}
    if false {
        var v7 int16 = -390
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: 0, F1: "a"}
            v7 += (v7 & 1)
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 0, F1: "\n\t"}
        _ = v10
        _ = v7
    } else {
        v11 := GenS2{F0: 1, F1: 18446744073709551615}
        _ = v11
    }
    var v12 int64 = -1
    var v13 int32 = -393
    v13 = 968
    _ = v12
    _ = v13
    _ = v6
    return 4294967295
}
func genFn1(p0 uint) uint16 {
    v14 := []int{}
    v15 := GenS1{F0: -104, F1: (("foo bar" + "hello") + "")}
    switch 32767 {
    case 127:
        p0 = 476
    case 119:
        switch "\n\t" {
        case "foo bar":
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        case "":
            p0 += p0
        default:
            fmt.Println(p0)
        }
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(p0)
    case "a":
        switch -1 {
        case -2147483648:
            v16 := GenS2{F0: 799, F1: 2}
            _ = v16
        case 199:
            var v17 int16 = 1
            _ = v17
        }
    default:
        p0 = (p0 & p0)
    }
    _ = p0
    _ = v14
    _ = v15
    return 75
}

func main() {
    var v1 uint64 = 0
    var v2 uint64 = (v1 & (v1 << ((0) & 63)))
    var v3 int16 = -32768
    v4 := make(map[uint32]int)
    var v5 uint8 = uint8(v2)
    v3 = (v3 % ((1) | 1))
    v18 := GenS2{F0: -504, F1: 206}
    v2 -= uint64(v5)
    if (v5 == v5) {
        v19 := GenS1{F0: 42, F1: "foo bar"}
        for fi20 := 0; fi20 < 16; fi20++ {
            v2 = v2
            v1 = (v1 ^ v2)
            var v21 uint64 = (^(v2))
            _ = fi20
            _ = v21
        }
        v3 = (v3 * -32768)
        _ = v19
    }
    fmt.Println(v2, v3)
    var v22 int32 = 198
    if true {
        for fi23 := 0; fi23 < 5; fi23++ {
            var v24 float64 = 1.5
            v25 := []int{fi23, fi23, fi23, (fi23 ^ fi23)}
            _ = fi23
            _ = v24
            _ = v25
        }
        switch 4294967295 {
        case 32768:
            var v26 int = -68
            _ = v26
        case 32767:
            switch 0 {
            case -9223372036854775808:
                _ = len(v4)
            default:
                v27 := []int{32768}
                _ = v27
            }
        default:
            _ = len(v4)
        }
        for fi28 := 0; fi28 < 11; fi28++ {
            fmt.Println(fi28, v22, v1)
            _ = fi28
        }
    }
    v5 = (v5 * (v5 * (v5 + (v5 % (((^(v5))) | 1)))))
    var v29 int32 = v22
    v30 := []int{-255}
    for fi31 := 0; fi31 < 7; fi31++ {
        var v32 uint16 = 65535
        var v33 uint8 = 127
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
}
