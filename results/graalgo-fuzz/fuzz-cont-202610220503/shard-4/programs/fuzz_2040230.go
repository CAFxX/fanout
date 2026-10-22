package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 string; F1 bool; F2 uint8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = (float64(float32((3.14159 * (0.1 - -2.25)))) - (((0.1 / (-2.25 * 1e10)) + (3.14159 * (1e10 * 0.1))) - 1.5))
    v3 := []int{-7, 878}
    var v4 uint8 = 47
    var v5 uint8 = v4
    if true {
        v5 += v4
        var v6 uint64 = 1
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 int = (fi7 ^ -528)
            v9 := GenS1{F0: 65535}
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v6
    }
    v4 = v4
    v4 = 0
    if false {
        fmt.Println(v2)
    } else {
        fmt.Println(v2)
        var v10 uint32 = 2
        v1[-537] = 2
        _ = v10
    }
    v11 := GenS1{F0: 255}
    var v12 uint16 = 92
    var v13 uint16 = 0
    if (4294967295 != 32767) {
        if (!false) {
            v14 := GenS1{F0: v13}
            var v15 uint16 = uint16(v4)
            _ = v14
            _ = v15
        } else {
            v16 := v1[-1]
            _ = v16
        }
        v17 := GenS1{F0: (v12 - v13)}
        _ = v17
    } else {
        if (v13 > v12) {
            v18 := GenS1{F0: v13}
            if (uint64(v5) < 179) {
                v5 = v5
                for fi19 := 0; fi19 < 7; fi19++ {
                    for fi20 := 0; fi20 < 6; fi20++ {
                        var v21 uint32 = 0
                        v22 := GenS2{F0: 1, F1: 2147483647}
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    _ = fi19
                }
            }
            _ = v18
        }
        var v23 int = int(v13)
        fmt.Println(v5, v13)
        _ = v23
    }
    if (v12 < v13) {
        for fi24 := 0; fi24 < 6; fi24++ {
            fmt.Println(fi24)
            v25 := GenS3{F0: ("hello" + "\n\t"), F1: (-16 >= 42), F2: v4}
            _ = fi24
            _ = v25
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
