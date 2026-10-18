package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int8 = 1
    var v2 float64 = ((1e10 * 1e10) + 1.5)
    var v3 float64 = (v2 * ((v2 - (v2 + (((float64(0.1) - (float64(-2.25) + (0.1 - float64(3.14159)))) - v2) - v2))) - (v2 + 1e10)))
    var v4 string = "a"
    if (v3 <= (v2 - (v3 / v3))) {
        var v5 bool = false
        _ = v5
    }
    if false {
        fmt.Println(v3, v2)
        v6 := []int{-264, -9223372036854775808}
        _ = v6
    }
    var v7 uint32 = 2
    if false {
        v2 = 3.14159
        var v8 float64 = v2
        _ = v8
    }
    var v9 uint = 0
    v2 = float64(-2.25)
    v3 -= v3
    var v10 uint8 = 127
    v11 := []int{2147483647, 127, 4294967295}
    for fi12 := 0; fi12 < 9; fi12++ {
        if (!((!(true && true)) && (519 == -44))) {
            v11 = append(v11, 0)
            v13 := GenS1{F0: -32768, F1: -309, F2: int32(v1)}
            if (false || true) {
                if ("a" > v4) {
                    v14 := GenS2{F0: (v7 + 140), F1: 42}
                    _ = v14
                }
            }
            _ = v13
        }
        _ = fi12
    }
    v15 := GenS1{F0: 1000, F1: -424, F2: 42}
    v9 -= 2
    v1 = v1
    if false {
        v3 = (v2 * v3)
        var v16 uint16 = 1
        _ = v16
    } else {
        v3 += float64((-2.25 - 1.5))
        v17 := GenS3{F0: (^(v10))}
        var v18 float32 = (-2.25 * 1e10)
        _ = v17
        _ = v18
    }
    var v19 uint8 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
