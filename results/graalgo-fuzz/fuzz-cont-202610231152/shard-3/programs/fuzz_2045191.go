package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint16; F1 bool; F2 int32 }
type GenS3 struct { F0 int64; F1 bool }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 245
    var v3 bool = (1000 < 0)
    v4 := make(map[int]string)
    v3 = v3
    v5 := []int{4294967295}
    v6 := GenS1{F0: (v2 * v2), F1: 152}
    v2 -= v2
    switch 624 {
    case -44:
        for fi7 := 0; fi7 < 13; fi7++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if false {
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 int = (-(fi7))
                    v10 := GenS2{F0: 255, F1: (false || v3), F2: 0}
                    if (v3 && v3) {
                        var v11 bool = (0.1 == 3.14159)
                        v12 := GenS3{F0: 1, F1: (v3 && (0.1 != (3.14159 - 0.1)))}
                        _ = v11
                        _ = v12
                    } else {
                        var v13 bool = false
                        _ = v13
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                if true {
                    v14 := GenS2{F0: 65535, F1: ((^(v2)) != v2), F2: -2147483648}
                    v15 := GenS1{F0: v2, F1: 1}
                    var v16 uint32 = 408
                    _ = v14
                    _ = v15
                    _ = v16
                } else {
                    var v17 float32 = 3.14159
                    _ = v17
                }
                v2 += (^(v2))
            }
            v3 = (false || (832 >= 833))
            _ = fi7
        }
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v3 = false
        var v19 int16 = -1
        _ = fi18
        _ = v19
    }
    if v3 {
        fmt.Println(v3, v3)
    }
    delete(v4, 32768)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
