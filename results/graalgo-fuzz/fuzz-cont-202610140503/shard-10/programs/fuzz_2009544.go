package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint8 }
func genFn0(p0 bool, p1 bool, p2 bool) uint {
    p2 = false
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 float32 = 1.5
        _ = fi4
        _ = v5
    }
    var v6 int8 = 127
    p0 = (float64((1e10 + 3.14159)) < float64(1e10))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 638
}
func genFn1(p0 int, p1 string, p2 string) int64 {
    switch "\n\t" {
    case "\n\t":
        for fi7 := 0; fi7 < 4; fi7++ {
            v8 := GenS1{F0: ((float64(1.5) - (1.5 - ((((-2.25 - -2.25) - ((float64((1.5 / 1.5)) - 1.5) * float64(0.1))) + 0.0) + 0.0))) > (float64(3.14159) + (float64((-2.25 * ((-2.25 * 1.5) - 3.14159))) + float64(3.14159))))}
            switch fi7 {
            case -9223372036854775808:
                var v9 float32 = float32(float64(3.14159))
                _ = v9
            default:
                p2 = p2
            }
            _ = fi7
            _ = v8
        }
    case "a":
        var v10 int = (^(p0))
        _ = v10
    }
    if (2 < 18446744073709551615) {
        v11 := GenS2{F0: 100, F1: "\n\t"}
        _ = v11
    } else {
        var v12 uint64 = 0
        _ = v12
    }
    v13 := []int{p0, (p0 % (((p0 + (p0 / ((p0) | 1)))) | 1)), p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return -822
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{0}
    v3 := []string{("" + "a"), "a", ""}
    v14 := GenS1{F0: (!(!true))}
    var v15 int32 = -2147483648
    var v16 float32 = (0.0 * 1e10)
    fmt.Println(v15)
    v1[4294967295] = -505
    v17 := GenS3{F0: 2}
    fmt.Println(v16, v16)
    v1[572] = 127
    fmt.Println(v15)
    var v18 uint = 94
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
