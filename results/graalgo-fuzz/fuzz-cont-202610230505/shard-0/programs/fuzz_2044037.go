package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{713}
    var v3 uint64 = 233
    var v4 uint16 = 1
    var v5 float32 = float32((-2.25 - 0.1))
    var v6 uint32 = 60
    var v7 int8 = -128
    if (uint(v4) < 2) {
        for fi8 := 0; fi8 < 6; fi8++ {
            if true {
                fmt.Println(v4, v6, v7)
                for fi9 := 0; fi9 < 5; fi9++ {
                    var v10 int64 = 9223372036854775807
                    if false {
                        if (v3 == v3) {
                            var v11 int64 = -1
                            _ = v11
                        }
                        fmt.Println(v10, v5)
                        var v12 float32 = float32(1.5)
                        _ = v12
                    }
                    _ = fi9
                    _ = v10
                }
                v4 = 915
            } else {
                var v13 float32 = v5
                _ = v13
            }
            v14 := GenS1{F0: 32767, F1: 65535, F2: 193}
            _ = fi8
            _ = v14
        }
        fmt.Println(v6, v7)
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 int64 = 9223372036854775807
        _ = len(v1)
        switch fi15 {
        case 9223372036854775807:
            for fi17 := 0; fi17 < 12; fi17++ {
                v2 = append(v2, 2147483647)
                var v18 uint = 255
                _ = fi17
                _ = v18
            }
        case 32767:
            v1[fi15] = (!((true && (v16 < (-(v16)))) && (v16 >= 9223372036854775807)))
        case 2147483647:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi15
        _ = v16
    }
    var v19 bool = (((v5 + -2.25) + v5) <= v5)
    delete(v1, -605)
    for fi20 := 0; fi20 < 6; fi20++ {
        fmt.Println(v5, v19, v5)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v21 bool = (!v19)
        _ = fi20
        _ = v21
    }
    v19 = (-1 > 42)
    if (27 > int64(v4)) {
        v7 += v7
        v19 = ("" > "")
    } else {
        v22 := GenS2{F0: 4294967295}
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        v19 = v19
        _ = v22
    }
    var v23 uint8 = 128
    var v24 uint32 = v6
    fmt.Println(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
