package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint16; F1 uint; F2 uint32 }
func genFn0() int64 {
    var v4 int64 = 1
    v4 = (v4 / ((42) | 1))
    v4 = (^(v4))
    if (3.14159 > float32(1.5)) {
        v5 := GenS1{F0: true}
        v6 := GenS1{F0: (true && false)}
        var v7 float32 = -2.25
        _ = v5
        _ = v6
        _ = v7
    }
    if (1 >= 41) {
        if (false && false) {
            var v8 float64 = 0.1
            for fi9 := 0; fi9 < 12; fi9++ {
                v10 := GenS2{F0: 1}
                _ = fi9
                _ = v10
            }
            v8 -= v8
            _ = v8
        } else {
            fmt.Println(v4)
        }
    } else {
        switch "" {
        case "":
            v11 := GenS3{F0: 1, F1: 4294967295}
            _ = v11
        case "\n\t":
            var v12 int = 4294967295
            _ = v12
        case "a":
            if ((0.0 + (float32(1.5) / float32(-2.25))) >= 1.5) {
                v13 := []int{32768, 0, 1, 0}
                var v14 uint8 = 23
                _ = v13
                _ = v14
            }
        default:
            var v15 int16 = -1
            _ = v15
        }
    }
    _ = v4
    return (v4 | v4)
}
func genFn1() int32 {
    var v16 uint64 = 4294967295
    var v17 uint32 = 969
    _ = v16
    _ = v17
    return 0
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    v1 = 42
    var v18 uint8 = 0
    v19 := GenS3{F0: (v18 + (v18 & (v18 / ((v18) | 1)))), F1: 32768}
    fmt.Println(v1, v1)
    v18 -= 128
    var v20 uint64 = 18446744073709551615
    if true {
        var v21 uint64 = (v20 ^ v20)
        _ = v21
    }
    v22 := GenS4{F0: 0, F1: 0, F2: 4294967295}
    v18 += (v18 ^ 42)
    if (!false) {
        v1 = 2147483647
        switch -807 {
        case -1:
            v1 -= genFn1()
        }
    } else {
        v23 := GenS3{F0: 127, F1: 32768}
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
}
