package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int8; F1 int }
func genFn0() uint32 {
    if (65535 > 71) {
        var v5 int8 = 0
        v5 = (v5 << ((992) & 63))
        _ = v5
    } else {
        switch (("foo bar" + "\n\t") + "hello") {
        case "":
            var v6 float64 = (-2.25 + (-2.25 - (((0.1 - 0.0) - float64(0.0)) - ((1e10 * 1.5) - 1.5))))
            _ = v6
        case "foo bar":
            switch "foo bar" {
            case "\n\t":
                v7 := []int{149, 9223372036854775807}
                _ = v7
            case "":
                v8 := []int{176, -635, 2147483647}
                _ = v8
            case "foo bar":
                var v9 bool = true
                _ = v9
            }
        default:
            v10 := GenS1{F0: false, F1: 65535}
            _ = v10
        }
        var v11 float32 = 3.14159
        _ = v11
    }
    if (0.0 != 0.0) {
        v12 := GenS2{F0: 4294967295, F1: -2147483648}
        var v13 int32 = -1
        if (!(!(1 >= 0))) {
            v13 = -377
        } else {
            v14 := GenS3{F0: -38, F1: 1}
            for fi15 := 0; fi15 < 18; fi15++ {
                v13 -= v13
                v16 := GenS4{F0: 42, F1: fi15}
                _ = fi15
                _ = v16
            }
            switch ((("a" + "\n\t") + "\n\t") + ("foo bar" + "hello")) {
            case "":
                v17 := []int{-9223372036854775808}
                _ = v17
            case "hello":
                var v18 string = (("a" + "a") + "hello")
                _ = v18
            default:
                v13 += (v13 >> ((0) & 63))
            }
            _ = v14
        }
        _ = v12
        _ = v13
    } else {
        v19 := GenS2{F0: 1, F1: 1}
        _ = v19
    }
    return 65535
}

func main() {
    var v1 uint32 = 934
    var v2 float32 = 0.0
    v3 := make(map[int]int)
    var v4 int = 252
    v1 = (v1 * 65535)
    if (true || (false && false)) {
        var v20 int8 = 42
        v21 := GenS3{F0: 100, F1: 42}
        v22 := []int{v4}
        _ = v20
        _ = v21
        _ = v22
    }
    v23 := GenS3{F0: 42, F1: 127}
    v1 -= (v1 & (v1 & (v1 << ((1) & 63))))
    v24 := GenS4{F0: 127, F1: v4}
    if (true || (-1 <= 1)) {
        fmt.Println(v1)
        v25 := GenS3{F0: 27, F1: 2}
        _ = v25
    }
    v4 += (v4 << ((4294967295) & 63))
    v26 := []int{v4}
    if true {
        var v27 uint32 = 42
        _ = len(v26)
        v2 = v2
        _ = v27
    } else {
        v1 = v1
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
    }
    var v28 uint64 = 18446744073709551615
    v2 = v2
    v2 -= float32(0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v28
    _ = v3
    _ = v4
}
