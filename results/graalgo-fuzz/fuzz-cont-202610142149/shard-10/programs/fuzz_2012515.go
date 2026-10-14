package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int; F1 uint; F2 uint32 }
type GenS4 struct { F0 int; F1 int16 }

func main() {
    var v1 uint16 = 255
    var v2 float64 = 0.1
    var v3 int8 = -83
    v4 := make(map[uint32]int)
    var v5 float32 = 1.5
    var v6 int64 = 1
    var v7 uint64 = 18446744073709551615
    v8 := GenS1{F0: (v1 / ((204) | 1)), F1: 32767}
    var v9 float32 = 0.0
    v10 := GenS2{F0: (v7 * 42), F1: 255}
    v11 := []int{59, 2147483647}
    var v12 int16 = 1000
    if false {
        v13 := v4[2147483648]
        _ = v13
    } else {
        if false {
            if ((false || (true && (-39 >= (v3 % ((v3) | 1))))) || (v5 < (1e10 * (((v9 / (v9 * 0.1)) / v9) - v9)))) {
                if false {
                    v14 := GenS3{F0: 1, F1: 539, F2: 2147483647}
                    var v15 uint64 = 0
                    _ = v14
                    _ = v15
                }
                for fi16 := 0; fi16 < 6; fi16++ {
                    v4[599] = fi16
                    v17 := GenS4{F0: (fi16 * (fi16 / ((128) | 1))), F1: -1}
                    _ = fi16
                    _ = v17
                }
            }
            fmt.Println(v6)
        }
        v9 = float32(float64(3.14159))
    }
    _ = len(v11)
    v18 := GenS3{F0: -2147483648, F1: 4294967295, F2: 2147483648}
    var v19 bool = true
    fmt.Println(v12, v9)
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
