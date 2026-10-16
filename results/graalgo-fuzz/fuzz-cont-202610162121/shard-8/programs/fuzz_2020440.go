package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint, p1 int) uint64 {
    var v4 uint16 = 32767
    v5 := GenS1{F0: -1, F1: 8, F2: -9223372036854775808}
    var v6 string = ("hello" + "\n\t")
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 bool, p1 uint16, p2 int64) uint32 {
    var v7 int64 = p2
    for fi8 := 0; fi8 < 5; fi8++ {
        p2 -= v7
        for fi9 := 0; fi9 < 14; fi9++ {
            fmt.Println(fi9)
            v10 := []int{(fi8 * 616), (fi9 / ((2147483647) | 1)), fi9}
            p2 += (p2 >> ((1) & 63))
            _ = fi9
            _ = v10
        }
        v11 := []int{(fi8 * (fi8 - fi8)), (fi8 ^ (fi8 / (((fi8 & fi8)) | 1))), (fi8 << ((18446744073709551615) & 63)), fi8}
        _ = fi8
        _ = v11
    }
    v12 := []int{-9223372036854775808, 2147483647, 752}
    switch 128 {
    case 714:
        v12 = append(v12, 1)
    default:
        var v13 int16 = -767
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    return 799
}

func main() {
    var v1 uint16 = 2
    var v2 float64 = 3.14159
    var v3 bool = false
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 uint8 = 128
        _ = fi14
        _ = v15
    }
    var v16 uint32 = 1
    v17 := GenS2{F0: 1, F1: 260}
    fmt.Println(v2, v16, v3)
    for fi18 := 0; fi18 < 12; fi18++ {
        if (false && v3) {
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 float64 = v2
                v21 := GenS1{F0: -128, F1: 0, F2: 9223372036854775807}
                _ = fi19
                _ = v20
                _ = v21
            }
            v22 := GenS1{F0: 100, F1: 0, F2: 407}
            _ = v22
        }
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 uint8 = 127
            _ = fi23
            _ = v24
        }
        _ = fi18
    }
    v25 := []int{-2147483648, 661}
    if (!v3) {
        switch "hello" {
        case "foo bar":
            fmt.Println(v3)
        case "\n\t":
            var v26 uint8 = 255
            _ = v26
        case "hello":
            for fi27 := 0; fi27 < 16; fi27++ {
                fmt.Println(v3)
                var v28 uint32 = genFn1(true, v1, 9223372036854775807)
                var v29 int8 = 5
                _ = fi27
                _ = v28
                _ = v29
            }
        }
    }
    v3 = v3
    v16 = (v16 >> ((4294967295) & 63))
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
}
