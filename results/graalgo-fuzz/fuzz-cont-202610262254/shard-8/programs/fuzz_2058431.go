package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 bool }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool }
func genFn0(p0 int16, p1 int32) string {
    p1 += p1
    if false {
        var v5 float32 = -2.25
        _ = v5
    } else {
        v6 := GenS1{F0: "hello", F1: "hello", F2: true}
        var v7 uint64 = 0
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    return "a"
}

func main() {
    var v1 uint8 = 255
    var v2 uint32 = 0
    v3 := make(map[string]int)
    var v4 uint = 255
    fmt.Println(v1)
    for fi8 := 0; fi8 < 11; fi8++ {
        v2 += (^(v2))
        _ = fi8
    }
    v1 += v1
    fmt.Println(v4, v4)
    for fi9 := 0; fi9 < 16; fi9++ {
        if (!(1 > -9223372036854775808)) {
            v10 := GenS2{F0: uint8(v4), F1: v2}
            v2 = (v2 - 2147483647)
            _ = v10
        } else {
            if false {
                var v11 float64 = (0.0 - (float64(float32(0.1)) * float64(1.5)))
                var v12 int16 = int16(v1)
                fmt.Println(v2, v11)
                _ = v11
                _ = v12
            } else {
                fmt.Println(v1, v4)
                switch 2147483647 {
                case 4294967295:
                    var v13 uint = v4
                    _ = v13
                case 32767:
                    if false {
                        var v14 int = fi9
                        _ = v14
                    }
                case -899:
                    switch (fi9 + (fi9 - fi9)) {
                    case -9223372036854775808:
                        v15 := GenS3{F0: 0, F1: 1, F2: 758}
                        _ = v15
                    }
                default:
                    switch -9223372036854775808 {
                    case 2147483647:
                        var v16 uint16 = 65535
                        _ = v16
                    }
                }
            }
            for fi17 := 0; fi17 < 15; fi17++ {
                var v18 bool = (-1 <= 32767)
                var v19 bool = v18
                v2 = 2147483647
                _ = fi17
                _ = v18
                _ = v19
            }
            v2 = 0
        }
        _ = fi9
    }
    var v20 uint32 = (^(v2))
    var v21 uint32 = uint32(v1)
    v1 = (v1 + 255)
    v22 := []int{}
    var v23 int32 = 2147483647
    v21 = 838
    for fi24 := 0; fi24 < 17; fi24++ {
        v1 -= (v1 << ((v4) & 63))
        var v25 bool = (3.14159 < 1e10)
        _ = fi24
        _ = v25
    }
    fmt.Println(v20, v1, v1)
    for fi26 := 0; fi26 < 6; fi26++ {
        fmt.Println(v2, v1)
        fmt.Println(v2, v1, v23)
        _ = fi26
    }
    if len(v22) > 0 {
        v22[0] = 2147483647
    }
    v27 := GenS4{F0: false}
    v28 := GenS5{F0: (v23 > (v23 % ((42) | 1)))}
    v21 = 785
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
