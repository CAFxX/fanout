package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int = 127
    v2 := make(map[int]bool)
    var v3 bool = (true || (42 >= 752))
    var v4 int64 = 243
    var v5 uint32 = 2
    var v6 int16 = -32768
    var v7 uint16 = 32767
    v3 = v3
    var v8 uint = 817
    for fi9 := 0; fi9 < 13; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(v8)
            _ = fi10
        }
        if true {
            for fi11 := 0; fi11 < 6; fi11++ {
                if (!(v3 || v3)) {
                    if true {
                        v12 := GenS1{F0: "a", F1: -2147483648}
                        _ = v12
                    }
                    v2[570] = true
                    v6 -= v6
                } else {
                    v4 = (-(v4))
                    if ((v8 | 18446744073709551615) != v8) {
                        fmt.Println(v1, v6, v8)
                        var v13 int32 = -1
                        _ = v13
                    }
                }
                v5 += 0
                _ = fi11
            }
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 int8 = -128
                _ = fi14
                _ = v15
            }
            v16 := GenS2{F0: 2}
            _ = v16
        } else {
            v5 -= (v5 & v5)
        }
        _ = fi9
    }
    var v17 uint64 = 718
    delete(v2, v1)
    if (1e10 >= ((float64((float32(-2.25) - 1e10)) * -2.25) / 1.0)) {
        fmt.Println(v6)
    } else {
        switch 127 {
        case 595:
            switch "" {
            case "hello":
                for fi18 := 0; fi18 < 2; fi18++ {
                    var v19 int32 = 42
                    if (!false) {
                        v19 += 42
                        var v20 bool = v3
                        _ = v20
                    }
                    v21 := GenS2{F0: 91}
                    _ = fi18
                    _ = v19
                    _ = v21
                }
            case "foo bar":
                for fi22 := 0; fi22 < 4; fi22++ {
                    for fi23 := 0; fi23 < 12; fi23++ {
                        fmt.Println(v17, v3)
                        fmt.Println(fi22)
                        v24 := []int{-35, (fi23 | fi23)}
                        _ = fi23
                        _ = v24
                    }
                    v25 := v2[fi22]
                    var v26 int = (v1 * (-(v1)))
                    _ = fi22
                    _ = v25
                    _ = v26
                }
            default:
                fmt.Println(v8, v1, v6)
            }
        case -1:
            switch "hello" {
            case "a":
                var v27 string = ("a" + "a")
                _ = v27
            case "hello":
                var v28 uint64 = (v17 >> (((v8 >> ((uint(v5)) & 63))) & 63))
                _ = v28
            }
        default:
            v5 -= v5
        }
        v29 := GenS3{F0: uint8(v5)}
        _ = v29
    }
    v30 := []int{v1, v1, v1}
    var v31 int16 = (v6 % ((42) | 1))
    var v32 string = ("a" + "hello")
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
