package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 uint }
func genFn0(p0 string) uint {
    fmt.Println(p0)
    v6 := GenS1{F0: 288, F1: (1 == 119)}
    for fi7 := 0; fi7 < 1; fi7++ {
        p0 = p0
        _ = fi7
    }
    switch p0 {
    case "\n\t":
        fmt.Println(p0)
    case "":
        var v8 uint16 = 65535
        _ = v8
    default:
        v9 := []int{32768, 32768}
        _ = v9
    }
    p0 = (("" + (p0 + p0)) + "\n\t")
    _ = p0
    _ = v6
    return 42
}
func genFn1(p0 int64, p1 int16) bool {
    switch 0 {
    case 0:
        p0 -= p0
    case -1:
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := []int{(fi10 / ((fi10) | 1)), (fi10 >> ((921) & 63)), 128, fi10}
            var v12 uint64 = 4294967295
            for fi13 := 0; fi13 < 8; fi13++ {
                fmt.Println(v12)
                v12 -= (^(v12))
                _ = fi13
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    case -45:
        var v14 int8 = 100
        _ = v14
    default:
        v15 := GenS2{F0: 2}
        _ = v15
    }
    v16 := GenS3{F0: 2, F1: 612}
    if (255 >= 42) {
        v17 := []int{2147483647, 4294967295}
        for fi18 := 0; fi18 < 17; fi18++ {
            for fi19 := 0; fi19 < 2; fi19++ {
                var v20 uint16 = 32767
                var v21 int16 = p1
                _ = fi19
                _ = v20
                _ = v21
            }
            p1 = (-(p1))
            _ = fi18
        }
        _ = v17
    } else {
        if (42 > genFn0(("" + "\n\t"))) {
            if (((1 < -1) && true) || true) {
                var v22 string = "hello"
                p0 -= p0
                for fi23 := 0; fi23 < 12; fi23++ {
                    v22 = (v22 + v22)
                    _ = fi23
                }
                _ = v22
            }
            p0 = (p0 << ((300) & 63))
        }
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    _ = v16
    return true
}

func main() {
    var v1 int64 = -1
    v2 := []uint32{2147483647, 42, 65535}
    v3 := make(map[string]string)
    var v4 uint16 = 662
    var v5 int = -121
    v24 := GenS3{F0: 4294967295, F1: 808}
    var v25 uint32 = 42
    for fi26 := 0; fi26 < 17; fi26++ {
        if len(v2) > 0 {
            v2[0] = (v25 & (v25 % ((2) | 1)))
        }
        _ = fi26
    }
    fmt.Println(v25)
    var v27 uint64 = 4294967295
    v5 -= 1
    fmt.Println(v1, v25)
    v1 = (v1 ^ v1)
    for fi28 := 0; fi28 < 15; fi28++ {
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 int16 = 1
            v25 += (v25 - v25)
            var v31 bool = ((32767 == (v4 + (v4 - v4))) && true)
            v32 := GenS3{F0: uint32(v4), F1: genFn0("\n\t")}
            fmt.Println(v27, v1)
            var v33 uint32 = v25
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
            _ = v33
        }
        switch ("a" + ("a" + ("" + (("" + ("" + ("\n\t" + ""))) + "")))) {
        case "":
            var v34 int16 = -530
            _ = v34
        default:
            var v35 float64 = (1e10 + float64((-2.25 - 3.14159)))
            _ = v35
        }
        v1 = (v1 - 42)
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
