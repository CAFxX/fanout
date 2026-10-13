package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 int8; F1 uint; F2 int16 }
func genFn0(p0 int8, p1 uint64) uint64 {
    v4 := GenS1{F0: 42, F1: 65535, F2: (^(p1))}
    v5 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p1
}

func main() {
    var v1 float32 = (1e10 - 1.5)
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    switch (("\n\t" + "") + "a") {
    case "\n\t":
        var v6 int = -2147483648
        _ = v6
    case "":
        var v7 float32 = 1e10
        _ = v7
    }
    if (4 > 42) {
        v1 -= v1
        v8 := GenS1{F0: 1, F1: 714, F2: 0}
        _ = v8
    }
    if (42 > -2147483648) {
        v1 = ((v1 - (v1 * ((float32(float64(v1)) * v1) * (v1 - (v1 - 1e10))))) + v1)
    }
    switch "a" {
    case "":
        for fi9 := 0; fi9 < 4; fi9++ {
            v1 -= v1
            for fi10 := 0; fi10 < 19; fi10++ {
                if false {
                    v11 := v3[4294967295]
                    _ = v11
                } else {
                    v12 := GenS2{F0: 42, F1: 32767, F2: 4294967295}
                    v1 += 3.14159
                    fmt.Println(fi10, v1, v1)
                    _ = v12
                }
                _ = fi10
            }
            _ = fi9
        }
    case "\n\t":
        v13 := GenS3{F0: -184, F1: -1}
        _ = v13
    case "a":
        fmt.Println(v1)
    }
    delete(v3, 0)
    if true {
        for fi14 := 0; fi14 < 17; fi14++ {
            v15 := GenS4{F0: 209, F1: 0}
            v16 := GenS5{F0: 127, F1: 4294967295, F2: 1000}
            _ = fi14
            _ = v15
            _ = v16
        }
    } else {
        var v17 bool = (false || (32768 <= 1))
        var v18 uint16 = 333
        var v19 uint8 = 1
        _ = v17
        _ = v18
        _ = v19
    }
    v20 := GenS3{F0: -32768, F1: -92}
    var v21 uint16 = 916
    v22 := v2[-128]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
