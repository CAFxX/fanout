package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint32 }

func main() {
    var v1 int16 = -540
    var v2 int = -757
    var v3 bool = false
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]bool)
    var v6 uint16 = 32767
    if (0 < -1) {
        for fi7 := 0; fi7 < 7; fi7++ {
            for fi8 := 0; fi8 < 8; fi8++ {
                v2 = (fi8 | (^(v2)))
                fmt.Println(fi8, v1, fi8)
                v9 := []int{fi8, fi8, (fi8 - 726)}
                _ = fi8
                _ = v9
            }
            delete(v5, (v2 * fi7))
            _ = fi7
        }
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 float32 = -2.25
            var v12 uint32 = 2147483647
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    v13 := GenS1{F0: (v1 >> ((255) & 63)), F1: v6, F2: 643}
    v2 += v2
    var v14 int = 4294967295
    fmt.Println(v6, v1, v6)
    var v15 uint64 = 2
    v16 := v5[v14]
    if (v3 || (-651 >= -1)) {
        var v17 uint16 = v6
        var v18 int = -474
        var v19 int8 = 42
        _ = v17
        _ = v18
        _ = v19
    } else {
        switch "foo bar" {
        case "\n\t":
            var v20 int = (v2 >> ((42) & 63))
            _ = v20
        }
        var v21 string = "hello"
        _ = v21
    }
    v14 += int(v15)
    if (2 == 191) {
        var v22 uint64 = v15
        if v3 {
            v23 := []int{v14, (^(v14)), v14}
            for fi24 := 0; fi24 < 16; fi24++ {
                v6 -= v6
                if v16 {
                    v25 := GenS1{F0: v1, F1: v6, F2: 128}
                    delete(v5, (v14 % ((fi24) | 1)))
                    v26 := GenS1{F0: v1, F1: 42, F2: 65535}
                    _ = v25
                    _ = v26
                }
                _ = fi24
            }
            v22 -= (v15 + v22)
            _ = v23
        } else {
            fmt.Println(v16)
            var v27 float64 = (((1e10 + 3.14159) / 0.1) / (float64((0.0 + (1e10 - 1.5))) * -2.25))
            v27 = v27
            _ = v27
        }
        v5[v14] = true
        _ = v22
    } else {
        var v28 bool = (!false)
        var v29 float64 = 0.0
        var v30 int32 = -1
        _ = v28
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
