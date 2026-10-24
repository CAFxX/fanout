package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int16; F2 string }
func genFn0(p0 uint, p1 bool) uint16 {
    var v5 uint64 = uint64(p0)
    switch 128 {
    case -2147483648:
        p0 -= p0
    case 777:
        var v6 uint32 = 1
        _ = v6
    case 128:
        if (!(p1 || (!((p0 % ((p0) | 1)) <= p0)))) {
            for fi7 := 0; fi7 < 7; fi7++ {
                v8 := GenS1{F0: 9223372036854775807}
                for fi9 := 0; fi9 < 5; fi9++ {
                    fmt.Println(v5)
                    _ = fi9
                }
                fmt.Println(fi7, fi7, v5)
                _ = fi7
                _ = v8
            }
            var v10 float32 = 1.5
            fmt.Println(v10)
            _ = v10
        }
    default:
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := GenS1{F0: -973}
            p0 += 2
            _ = fi11
            _ = v12
        }
    }
    var v13 uint8 = 230
    _ = p0
    _ = p1
    _ = v13
    _ = v5
    return 255
}

func main() {
    var v1 uint8 = 2
    var v2 uint8 = 0
    var v3 uint16 = 1
    var v4 float32 = ((3.14159 / float32(3.14159)) + 3.14159)
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 float64 = -2.25
        _ = fi14
        _ = v15
    }
    var v16 uint8 = (v2 + 0)
    fmt.Println(v2, v3)
    var v17 string = "\n\t"
    v18 := GenS2{F0: 0, F1: -117, F2: v17}
    switch ((v17 + v17) + "hello") {
    case "\n\t":
        v16 += v16
    case "hello":
        v19 := GenS1{F0: 0}
        _ = v19
    case "a":
        var v20 int8 = 117
        _ = v20
    }
    v4 += (((3.14159 - (float32(float64(v4)) / v4)) + v4) + 1.5)
    fmt.Println(v2, v16, v3)
    v2 = 1
    v21 := []int{-2147483648, 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
