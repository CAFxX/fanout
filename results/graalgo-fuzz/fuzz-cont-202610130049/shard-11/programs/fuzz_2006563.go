package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint8; F1 string; F2 uint8 }
type GenS4 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS5 struct { F0 string }
func genFn0(p0 bool, p1 string, p2 bool) bool {
    v6 := GenS1{F0: 42}
    p2 = false
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (!((!p2) && false))
}

func main() {
    var v1 float64 = ((1.5 - 1.5) - 3.14159)
    var v2 int = 32767
    v3 := make(map[string]uint32)
    var v4 float32 = 3.14159
    v5 := make(map[int]int)
    var v7 int = v2
    fmt.Println(v7, v2, v1)
    if (65535 <= 566) {
        fmt.Println(v2)
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int64 = -9223372036854775808
            v10 := GenS2{F0: 32767, F1: 2147483648, F2: 100}
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    v11 := GenS2{F0: v2, F1: 4294967295, F2: 40}
    v12 := GenS3{F0: 230, F1: "hello", F2: 78}
    v4 = 1.5
    v13 := v5[(v7 / (((v7 + v7)) | 1))]
    if true {
        if false {
            v1 = v1
            v13 = v7
        } else {
            v4 = -2.25
            v14 := []int{(v13 ^ (v7 * v7)), (v7 & v7)}
            for fi15 := 0; fi15 < 7; fi15++ {
                if true {
                    for fi16 := 0; fi16 < 20; fi16++ {
                        var v17 uint64 = 1
                        if len(v14) > 0 {
                            v14[0] = fi16
                        }
                        _ = fi16
                        _ = v17
                    }
                    v18 := GenS2{F0: 0, F1: 1, F2: -120}
                    var v19 float64 = v1
                    _ = v18
                    _ = v19
                }
                _ = fi15
            }
            _ = v14
        }
        _ = len(v3)
    } else {
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 int64 = 0
            var v22 int64 = v21
            v23 := GenS4{F0: 688, F1: 0, F2: 263}
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        switch (v2 | -9223372036854775808) {
        case 1:
            v24 := GenS5{F0: ("\n\t" + ("\n\t" + "\n\t"))}
            _ = v24
        default:
            var v25 uint8 = 43
            _ = v25
        }
        _ = len(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
