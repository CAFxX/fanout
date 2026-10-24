package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int64 }
func genFn0(p0 uint64) int64 {
    p0 += 0
    p0 = (^(p0))
    v6 := GenS1{F0: 0}
    switch "" {
    case "\n\t":
        if true {
            var v7 uint = 667
            for fi8 := 0; fi8 < 11; fi8++ {
                fmt.Println(v7, v7)
                var v9 string = "hello"
                for fi10 := 0; fi10 < 9; fi10++ {
                    switch (fi10 >> ((594) & 63)) {
                    case 128:
                        p0 = (p0 + (p0 * p0))
                    case 942:
                        switch 222 {
                        case 0:
                            var v11 int64 = 9223372036854775807
                            _ = v11
                        }
                        p0 -= p0
                    case 1:
                        var v12 int16 = 42
                        _ = v12
                    }
                    v13 := []int{fi8, fi10, (fi10 | (fi8 - (fi10 * fi10)))}
                    _ = fi10
                    _ = v13
                }
                _ = fi8
                _ = v9
            }
            var v14 uint32 = 2
            _ = v14
            _ = v7
        }
    case "":
        v15 := GenS1{F0: -869}
        _ = v15
    default:
        if false {
            p0 += (p0 + (p0 & (p0 >> ((0) & 63))))
            v16 := []int{486}
            _ = v16
        }
    }
    fmt.Println(p0)
    _ = p0
    _ = v6
    return -92
}
func genFn1(p0 uint32, p1 uint) int8 {
    var v17 uint16 = 42
    v18 := []int{-26, 0}
    _ = len(v18)
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 42
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[string]uint32)
    var v3 float64 = 0.0
    var v4 uint64 = 130
    var v5 uint = 1
    for fi19 := 0; fi19 < 1; fi19++ {
        v20 := []int{(v1 >> (((^(v5))) & 63)), (^(fi19))}
        v3 = (((v3 - v3) - 3.14159) + 1e10)
        _ = fi19
        _ = v20
    }
    var v21 int32 = 2147483647
    v22 := GenS1{F0: (-(v21))}
    var v23 uint8 = 2
    if true {
        v24 := []int{-2147483648, (v1 % (((v1 - v1)) | 1))}
        _ = v24
    }
    v25 := GenS1{F0: v21}
    v21 = (v21 - v21)
    if (true && (0 == 789)) {
        v26 := GenS2{F0: -1, F1: 0}
        _ = v26
    } else {
        var v27 int = (v1 << (((v5 >> (((v5 << ((uint(v4)) & 63))) & 63))) & 63))
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
