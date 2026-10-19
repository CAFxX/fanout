package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 bool) uint32 {
    var v6 uint16 = 2
    if p0 {
        if false {
            v6 += (^(v6))
            switch -1 {
            case -2147483648:
                if (!(-8 != -128)) {
                    var v7 uint8 = 128
                    p0 = (3.14159 != 1e10)
                    _ = v7
                }
            }
            v8 := GenS1{F0: 2, F1: (v6 + (v6 - v6))}
            _ = v8
        }
        for fi9 := 0; fi9 < 19; fi9++ {
            v6 -= (v6 << ((255) & 63))
            for fi10 := 0; fi10 < 7; fi10++ {
                p0 = ((0.0 != (float64(float32(((0.0 + 1e10) + 1.5))) / 1e10)) && p0)
                fmt.Println(v6, fi10)
                v6 += (v6 & 65535)
                switch fi9 {
                case 9223372036854775807:
                    v11 := GenS1{F0: 18446744073709551615, F1: 509}
                    _ = v11
                }
                v6 += v6
                p0 = (1000 > -832)
                v12 := []int{fi10, (fi10 >> ((58) & 63)), fi10}
                _ = fi10
                _ = v12
            }
            p0 = (-32768 < 0)
            _ = fi9
        }
    } else {
        if ("a" < "") {
            var v13 int32 = -1
            var v14 int64 = 42
            _ = v13
            _ = v14
        }
    }
    v6 -= v6
    v15 := []int{-496, 9223372036854775807, -1, 128}
    v15 = append(v15, -116)
    _ = p0
    _ = v15
    _ = v6
    return 4294967295
}
func genFn1(p0 int, p1 int32) uint {
    p1 -= p1
    fmt.Println(p1, p1)
    p1 -= p1
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 800
}

func main() {
    var v1 int8 = -128
    var v2 uint32 = 1
    var v3 uint8 = 140
    v4 := []uint32{v2}
    v5 := []int{-753}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v16 := GenS1{F0: 400, F1: 2}
    v2 = (v2 * 425)
    if true {
        var v17 bool = true
        var v18 uint8 = 104
        _ = v17
        _ = v18
    } else {
        _ = len(v4)
        v19 := GenS2{F0: 0}
        _ = v19
    }
    v1 += 100
    v2 = genFn0(false)
    v2 = (^(v2))
    var v20 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
