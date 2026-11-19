package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int; F1 int32; F2 int32 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint }
type GenS4 struct { F0 int16 }
func genFn0(p0 int32, p1 uint64, p2 int16) int32 {
    var v7 string = "hello"
    var v8 uint64 = 699
    if (0 > 2) {
        fmt.Println(v7)
        v9 := []int{2147483647, 42, -1, -9223372036854775808}
        v10 := GenS1{F0: -726, F1: 128}
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (-(p0))
}
func genFn1() uint8 {
    var v11 string = "foo bar"
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 uint8 = 2
        v11 = v11
        _ = fi12
        _ = v13
    }
    var v14 uint16 = 481
    _ = v11
    _ = v14
    return 201
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 4294967295
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    v5 := []string{"hello", "hello", "a", "hello"}
    var v6 float32 = -2.25
    switch "hello" {
    case "hello":
        for fi15 := 0; fi15 < 7; fi15++ {
            for fi16 := 0; fi16 < 15; fi16++ {
                v6 = ((0.0 + 0.0) + (v6 + v6))
                _ = fi16
            }
            var v17 bool = (!(true && false))
            var v18 string = ("hello" + "hello")
            _ = fi15
            _ = v17
            _ = v18
        }
    }
    v6 -= v6
    v6 = 0.1
    v2 = (v2 % ((v2) | 1))
    switch 1 {
    case -459:
        var v19 int16 = 0
        _ = v19
    default:
        _ = len(v5)
    }
    fmt.Println(v2)
    delete(v3, 580)
    v20 := GenS2{F0: int(v2), F1: 0, F2: 2147483647}
    if (!false) {
        v21 := GenS3{F0: (v2 * 18446744073709551615), F1: 65535, F2: 255}
        fmt.Println(v2, v6)
        var v22 int = -510
        _ = v21
        _ = v22
    } else {
        v2 += 18446744073709551615
        v23 := GenS4{F0: 42}
        v5 = append(v5, ("" + "\n\t"))
        _ = v23
    }
    v24 := GenS3{F0: (v2 + (v2 | (^(v2)))), F1: 32767, F2: 575}
    if (false || (!false)) {
        switch "a" {
        case "foo bar":
            v25 := v4[-25]
            _ = v25
        }
        _ = len(v1)
        if (1 >= 42) {
            fmt.Println(v6)
        }
    }
    var v26 int = -1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v26)
    fmt.Println(v2)
    var v27 int64 = int64(v2)
    for fi28 := 0; fi28 < 12; fi28++ {
        _ = len(v5)
        _ = fi28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
