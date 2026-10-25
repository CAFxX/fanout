package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 bool }
func genFn0(p0 uint16, p1 uint64) string {
    p1 = (^(p1))
    if ((!false) || ((float64(float32(1e10)) >= 1.5) || (-399 < 0))) {
        var v4 float64 = 1e10
        _ = v4
    }
    _ = p0
    _ = p1
    return "\n\t"
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]int)
    v3 := []int{128, 0, -1, -709}
    v5 := GenS1{F0: 668, F1: 0, F2: 4294967295}
    var v6 uint8 = 0
    if (890 >= 18446744073709551615) {
        v6 += v6
        fmt.Println(v6)
        for fi7 := 0; fi7 < 18; fi7++ {
            if false {
                v8 := GenS2{F0: 2}
                v6 += v6
                _ = v8
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            v9 := GenS2{F0: 18446744073709551615}
            _ = fi7
            _ = v9
        }
    }
    var v10 uint32 = 2
    v11 := GenS3{F0: 1, F1: true}
    v10 = (v10 - v10)
    var v12 uint8 = v6
    fmt.Println(v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
