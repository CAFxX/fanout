package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 float64 = (1.5 - 0.1)
    var v4 float64 = float64(0.1)
    var v5 float64 = -2.25
    v6 := v2[-899]
    v7 := GenS1{F0: v6}
    v8 := []int{0, 0, 32767}
    for fi9 := 0; fi9 < 10; fi9++ {
        if (0 >= 348) {
            v3 += 0.1
            var v10 uint = 1
            _ = v10
        } else {
            v5 -= (v3 + (float64(0.0) / (v4 - v5)))
        }
        if true {
            var v11 float64 = v4
            _ = v11
        } else {
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 uint64 = 1
                _ = fi12
                _ = v13
            }
            switch "hello" {
            case "hello":
                v5 = float64(0.0)
            case "a":
                v8 = append(v8, fi9)
            default:
                fmt.Println(fi9, v6, v3)
            }
            v2[(fi9 & 42)] = v6
        }
        _ = fi9
    }
    var v14 int16 = 1
    var v15 float32 = (float32(v5) - float32(((0.1 + v4) + 3.14159)))
    switch 32767 {
    case 127:
        v14 = (v14 / (((v14 << ((744) & 63))) | 1))
    case 523:
        fmt.Println(v15, v5)
    case 128:
        _ = len(v8)
    default:
        v16 := GenS2{F0: 0, F1: 793}
        _ = v16
    }
    fmt.Println(v14)
    v5 -= 3.14159
    v15 -= v15
    switch ("hello" + (("\n\t" + ("" + "a")) + ("" + (("" + "\n\t") + "")))) {
    case "":
        switch -752 {
        case 32768:
            v17 := GenS3{F0: 1}
            _ = v17
        case -2147483648:
            var v18 uint64 = 1
            _ = v18
        default:
            v5 = v5
        }
    case "foo bar":
        for fi19 := 0; fi19 < 19; fi19++ {
            if len(v8) > 0 {
                v8[0] = 32768
            }
            fmt.Println(v4)
            _ = fi19
        }
    case "a":
        var v20 float64 = v3
        _ = v20
    }
    v21 := GenS2{F0: -2147483648, F1: 32768}
    for fi22 := 0; fi22 < 11; fi22++ {
        fmt.Println(v4, v3, fi22)
        delete(v2, (fi22 | (-(fi22))))
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
