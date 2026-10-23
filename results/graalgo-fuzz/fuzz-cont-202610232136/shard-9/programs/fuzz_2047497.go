package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 string, p1 uint64, p2 uint8) string {
    v4 := GenS1{F0: 931}
    var v5 int64 = 42
    switch 1 {
    case 1:
        fmt.Println(p0)
    case -883:
        v6 := GenS2{F0: 42, F1: p2, F2: (v5 << ((0) & 63))}
        _ = v6
    case -2147483648:
        v7 := GenS3{F0: (p2 | 178)}
        _ = v7
    }
    fmt.Println(p1, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p0
}
func genFn1(p0 bool, p1 uint8, p2 uint8) uint8 {
    var v8 int16 = -580
    v9 := []int{}
    var v10 int32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return p1
}

func main() {
    v1 := []string{"foo bar", "a"}
    v2 := make(map[int]int)
    var v3 float64 = ((float64((0.1 + (0.0 - 0.1))) * 1e10) / (float64(-2.25) + 1e10))
    v3 = v3
    for fi11 := 0; fi11 < 3; fi11++ {
        v3 = v3
        _ = fi11
    }
    v12 := GenS3{F0: 128}
    _ = len(v2)
    switch "\n\t" {
    case "foo bar":
        var v13 int32 = 2147483647
        _ = v13
    case "":
        var v14 int64 = 962
        _ = v14
    default:
        v3 = 0.1
    }
    var v15 uint8 = 1
    var v16 uint16 = 261
    var v17 bool = false
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    v3 = v3
    v15 -= (v15 & (v15 - v15))
    v18 := GenS1{F0: 42}
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := GenS1{F0: 1000}
        fmt.Println(fi19)
        for fi21 := 0; fi21 < 13; fi21++ {
            v3 += 0.0
            _ = fi21
        }
        _ = fi19
        _ = v20
    }
    var v22 int64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
}
