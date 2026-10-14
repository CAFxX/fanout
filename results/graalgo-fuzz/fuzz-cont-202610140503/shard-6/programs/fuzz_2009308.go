package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 int8 }
func genFn0(p0 int16, p1 uint) uint16 {
    var v4 bool = (!(1.5 != 1.5))
    switch "" {
    case "":
        var v5 int = 352
        _ = v5
    default:
        v6 := []int{-933, -2147483648, 127}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 282
}

func main() {
    var v1 int64 = 1
    var v2 string = "\n\t"
    var v3 uint64 = 984
    fmt.Println(v1, v2, v3)
    v7 := GenS1{F0: 100, F1: v3, F2: 128}
    if (true || false) {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS2{F0: 18446744073709551615, F1: (1e10 < (0.1 + 0.1))}
            for fi10 := 0; fi10 < 4; fi10++ {
                v2 = v2
                var v11 uint64 = v3
                v12 := []int{fi8, (fi10 % ((fi10) | 1))}
                _ = fi10
                _ = v11
                _ = v12
            }
            switch fi8 {
            case -1:
                for fi13 := 0; fi13 < 8; fi13++ {
                    v1 = 99
                    _ = fi13
                }
            }
            _ = fi8
            _ = v9
        }
    } else {
        var v14 bool = (float32(3.14159) <= 1.5)
        var v15 uint = 2
        _ = v14
        _ = v15
    }
    v16 := []int{}
    switch 4294967295 {
    case -888:
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := GenS3{F0: 0}
            _ = fi17
            _ = v18
        }
    case 679:
        if (uint32(v3) <= 2147483647) {
            var v19 float64 = (float64(float32(1e10)) / 0.1)
            _ = v19
        } else {
            for fi20 := 0; fi20 < 7; fi20++ {
                for fi21 := 0; fi21 < 10; fi21++ {
                    v3 -= v3
                    v3 = v3
                    _ = fi21
                }
                _ = fi20
            }
            var v22 int16 = 1000
            if len(v16) > 0 {
                v16[0] = 0
            }
            _ = v22
        }
    case 101:
        for fi23 := 0; fi23 < 4; fi23++ {
            v2 = v2
            fmt.Println(fi23)
            fmt.Println(v1, fi23, fi23)
            _ = fi23
        }
    }
    v24 := GenS4{F0: v1, F1: 0}
    switch 0 {
    case -2147483648:
        var v25 string = v2
        _ = v25
    case 97:
        var v26 uint32 = 2
        _ = v26
    case 893:
        v2 = v2
    default:
        if false {
            var v27 uint64 = v3
            _ = v27
        }
    }
    fmt.Println(v1, v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v7
}
