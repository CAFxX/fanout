package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int8; F1 bool }
func genFn0(p0 int, p1 int32) uint {
    v4 := []int{}
    p1 -= p1
    p0 += (p0 ^ p0)
    p0 = 32768
    var v5 float32 = (1e10 - 1.5)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1(p0 bool) bool {
    var v6 float32 = ((((3.14159 / ((3.14159 - -2.25) + ((0.0 * (-2.25 + (float32((((3.14159 + 3.14159) / float64(3.14159)) / 3.14159)) * 0.0))) + float32(-2.25)))) + float32(float64((3.14159 / 1.5)))) + 1.5) + 0.1)
    p0 = p0
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v6
    return p0
}

func main() {
    v1 := []int{-2147483648, -438, 0}
    v2 := []string{}
    v3 := make(map[int]uint32)
    v9 := GenS1{F0: "a", F1: 33}
    var v10 uint8 = 169
    v10 += v10
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            if (2 != 2) {
                fmt.Println(fi11)
                v10 -= 42
                v10 += v10
                var v13 int32 = 42
                _ = v13
            }
            var v14 int = -9223372036854775808
            var v15 int = 42
            _ = fi12
            _ = v14
            _ = v15
        }
        _ = fi11
    }
    var v16 uint64 = 72
    var v17 int64 = 9223372036854775807
    var v18 int64 = (v17 * v17)
    v10 = v10
    if (false && ((float64(-2.25) - (0.1 * 1.5)) <= 1.5)) {
        for fi19 := 0; fi19 < 1; fi19++ {
            v20 := GenS2{F0: -106, F1: true}
            if (127 == -128) {
                fmt.Println(v17)
                v1 = append(v1, fi19)
                v17 -= 440
            }
            var v21 uint = genFn0(fi19, -179)
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        v18 += 9223372036854775807
        v16 = v16
    }
    var v22 uint64 = v16
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v9
}
