package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint64; F1 string; F2 uint8 }
func genFn0(p0 uint16) uint32 {
    v4 := GenS1{F0: 2147483647, F1: -633}
    var v5 string = "a"
    if (((((1.5 + 1e10) - ((0.0 / float32(1.5)) + (((1e10 - 0.1) * 1.5) + 3.14159))) * -2.25) + -2.25) != 1e10) {
        fmt.Println(v5, v5)
        var v6 uint = 255
        _ = v6
    } else {
        v7 := []int{}
        for fi8 := 0; fi8 < 10; fi8++ {
            for fi9 := 0; fi9 < 7; fi9++ {
                var v10 int32 = 298
                _ = fi9
                _ = v10
            }
            var v11 int16 = 594
            _ = fi8
            _ = v11
        }
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int = fi12
            _ = fi12
            _ = v13
        }
        _ = v7
    }
    _ = p0
    _ = v4
    _ = v5
    return 2147483648
}

func main() {
    var v1 int16 = 456
    v2 := []string{}
    var v3 uint64 = 18446744073709551615
    v3 += v3
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint16 = 1
        v1 += v1
        fmt.Println(fi14, v1)
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: 370, F1: 4294967295}
    _ = len(v2)
    if (1 < 1) {
        v3 = (v3 ^ v3)
    }
    var v17 int16 = v1
    var v18 bool = (!(false && (127 == int8(v17))))
    if false {
        v3 = (^(v3))
        if (v3 < (^(v3))) {
            fmt.Println(v3, v18, v17)
        } else {
            var v19 int16 = (^(v1))
            for fi20 := 0; fi20 < 13; fi20++ {
                v18 = (true || v18)
                v17 += v1
                _ = fi20
            }
            var v21 uint8 = 232
            _ = v19
            _ = v21
        }
    } else {
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    }
    if (0 <= 0) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if false {
        for fi22 := 0; fi22 < 8; fi22++ {
            v17 -= v17
            v23 := GenS3{F0: 255, F1: fi22}
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 1; fi24++ {
            v25 := GenS2{F0: 42, F1: (v3 - (v3 & 4294967295))}
            v3 = (v3 << ((4294967295) & 63))
            _ = fi24
            _ = v25
        }
    }
    var v26 int16 = (^(v17))
    v1 -= 369
    v27 := GenS4{F0: v3, F1: "\n\t", F2: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
