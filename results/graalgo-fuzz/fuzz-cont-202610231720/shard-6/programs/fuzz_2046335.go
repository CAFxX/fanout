package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int32; F1 int16 }
func genFn0() int {
    var v7 uint8 = 57
    if true {
        var v8 int16 = 0
        _ = v8
    }
    switch "\n\t" {
    case "hello":
        v9 := GenS1{F0: 18446744073709551615, F1: 759, F2: 1}
        _ = v9
    }
    var v10 uint = 4
    var v11 uint64 = 774
    _ = v10
    _ = v11
    _ = v7
    return 981
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 798
    var v3 uint = 588
    var v4 int64 = -647
    var v5 float32 = 1e10
    v6 := make(map[uint32]bool)
    v12 := []int{(v2 - v2), (v2 & (v2 / ((genFn0()) | 1))), 914, -792}
    var v13 int8 = 42
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 uint8 = 255
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(fi16, v4, v2)
        var v17 uint32 = 4294967295
        switch ("a" + "a") {
        case "\n\t":
            v17 -= (v17 - v17)
        case "":
            if (((1000 < 0) && (2 > 255)) && true) {
                switch "" {
                case "hello":
                    v2 = (fi16 << ((16) & 63))
                default:
                    switch "a" {
                    case "foo bar":
                        _ = len(v1)
                    case "":
                        v6[(^(v17))] = false
                    case "a":
                        v18 := GenS1{F0: v3, F1: 0, F2: v4}
                        _ = v18
                    }
                }
            } else {
                fmt.Println(v4)
            }
        case "hello":
            var v19 int = fi16
            _ = v19
        default:
            for fi20 := 0; fi20 < 2; fi20++ {
                v21 := GenS2{F0: 127, F1: 251}
                v5 = v5
                _ = fi20
                _ = v21
            }
        }
        _ = fi16
        _ = v17
    }
    var v22 bool = true
    v23 := GenS3{F0: 42, F1: 42}
    v5 = 3.14159
    var v24 bool = v22
    var v25 int = (v2 / ((genFn0()) | 1))
    v25 = genFn0()
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
