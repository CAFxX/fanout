package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 int }

func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]string)
    if false {
        if true {
            var v6 uint = 786
            var v7 float64 = 1e10
            _ = v6
            _ = v7
        }
    }
    fmt.Println(v1)
    v8 := []int{}
    for fi9 := 0; fi9 < 9; fi9++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        if false {
            var v10 uint = 42
            _ = v10
        } else {
            for fi11 := 0; fi11 < 7; fi11++ {
                switch "a" {
                case "a":
                    fmt.Println(fi11, fi11, v1)
                default:
                    v12 := GenS1{F0: 2}
                    _ = v12
                }
                _ = fi11
            }
            v1 = v1
            v8 = append(v8, fi9)
        }
        v13 := GenS1{F0: 2}
        _ = fi9
        _ = v13
    }
    v1 += (v1 | v1)
    if (false && false) {
        v3[0] = 0
        if false {
            v14 := GenS2{F0: 2147483647}
            if false {
                var v15 uint = 828
                _ = v15
            } else {
                var v16 uint32 = 42
                v17 := GenS3{F0: 0}
                _ = v16
                _ = v17
            }
            _ = v14
        }
    }
    v18 := GenS3{F0: 4294967295}
    v1 = v1
    if ((251 > 127) || (2 != 3)) {
        v19 := GenS1{F0: 18446744073709551615}
        _ = v19
    } else {
        v1 = v1
    }
    v1 = (v1 ^ v1)
    var v20 float64 = (3.14159 - float64((0.0 * 3.14159)))
    v21 := GenS4{F0: 6, F1: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
