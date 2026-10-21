package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint; F2 int64 }
type GenS4 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint32) uint32 {
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 float64 = float64(1.5)
        v6 = 0.0
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 1000}
    if true {
        if true {
            p0 -= (^(p0))
        }
        v8 := GenS1{F0: 0}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS2{F0: 9223372036854775807, F1: 65535, F2: 102}
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := []int{fi11, fi11, fi11}
            p0 = p0
            _ = fi11
            _ = v12
        }
        p0 = p0
    }
    v13 := GenS3{F0: 42, F1: 440, F2: -9223372036854775808}
    for fi14 := 0; fi14 < 12; fi14++ {
        if true {
            if false {
                var v15 int = (fi14 * (^(fi14)))
                v16 := GenS4{F0: 186, F1: 1}
                fmt.Println(v15, p0, v15)
                _ = v15
                _ = v16
            }
        } else {
            var v17 int64 = -1
            var v18 bool = true
            p0 += (p0 >> ((0) & 63))
            _ = v17
            _ = v18
        }
        var v19 uint16 = 65535
        _ = fi14
        _ = v19
    }
    _ = p0
    _ = v13
    _ = v7
    return 2
}

func main() {
    v1 := []string{}
    v2 := []uint32{42}
    var v3 uint8 = 198
    var v4 uint32 = 4294967295
    var v20 int8 = 0
    v3 = (v3 / (((v3 + (v3 - (^(v3))))) | 1))
    v21 := GenS2{F0: 152, F1: (v4 & v4), F2: v20}
    v22 := GenS3{F0: genFn0((v4 / (((v4 / (((^(v4))) | 1))) | 1))), F1: 65535, F2: 1}
    for fi23 := 0; fi23 < 2; fi23++ {
        v4 = v4
        _ = len(v2)
        fmt.Println(fi23, v4)
        _ = fi23
    }
    fmt.Println(v3, v3)
    v20 -= (v20 + (v20 + v20))
    fmt.Println(v4, v3, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
