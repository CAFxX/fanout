package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS4 struct { F0 string; F1 uint16 }
func genFn0(p0 uint32) uint {
    v5 := GenS1{F0: 20}
    p0 = p0
    v6 := []int{42}
    _ = p0
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 string, p1 string, p2 int32) int8 {
    p1 = p1
    p2 -= 42
    if (4294967295 == 42) {
        var v7 float32 = 0.1
        if ((1e10 + 1.5) > 0.1) {
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 int32 = p2
                _ = fi8
                _ = v9
            }
            var v10 float64 = 0.0
            _ = v10
        } else {
            v11 := []int{-9223372036854775808, -501, 42}
            for fi12 := 0; fi12 < 14; fi12++ {
                p2 -= p2
                _ = fi12
            }
            p1 = p0
            _ = v11
        }
        _ = v7
    }
    var v13 string = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return -32
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 61
    v3 := make(map[string]string)
    var v4 int64 = 9223372036854775807
    var v14 int64 = (^(v4))
    var v15 int64 = v14
    v16 := v1["hello"]
    fmt.Println(v16, v2)
    v4 = v14
    v15 -= 0
    v2 = 65535
    v4 += (v15 - -1)
    switch ("hello" + "foo bar") {
    case "\n\t":
        fmt.Println(v2, v15)
    case "hello":
        v17 := GenS2{F0: 4294967295}
        _ = v17
    default:
        fmt.Println(v4, v14, v4)
    }
    v18 := GenS3{F0: 2, F1: 604, F2: -32768}
    v19 := GenS4{F0: "\n\t", F1: 411}
    fmt.Println(v15, v14)
    v15 -= v15
    if false {
        var v20 int32 = 0
        v21 := GenS2{F0: 797}
        _ = v20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
