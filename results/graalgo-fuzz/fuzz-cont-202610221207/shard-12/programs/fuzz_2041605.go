package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint, p1 string) bool {
    v6 := GenS1{F0: (float32((1.5 - (-2.25 * 0.0))) != 0.1), F1: -32768}
    switch p1 {
    case "hello":
        p1 = p1
    case "a":
        var v7 string = ("" + p1)
        _ = v7
    case "foo bar":
        if true {
            p0 = 0
            var v8 bool = true
            v9 := GenS2{F0: -1}
            _ = v8
            _ = v9
        }
    default:
        v10 := GenS3{F0: -1}
        _ = v10
    }
    v11 := GenS4{F0: 255, F1: 160}
    p0 -= p0
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return true
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    var v4 string = "hello"
    var v5 uint8 = 183
    switch "foo bar" {
    case "\n\t":
        if (false || (((-75 < -1) && (true || true)) && (!(171 != 1)))) {
            v5 = (v5 >> ((481) & 63))
            fmt.Println(v5, v5)
        } else {
            fmt.Println(v5, v4)
            if false {
                v12 := GenS1{F0: true, F1: 42}
                v4 = v4
                _ = v12
            } else {
                var v13 int32 = 42
                v14 := GenS1{F0: false, F1: 42}
                _ = v13
                _ = v14
            }
            if (true || (-1 < 42)) {
                fmt.Println(v4)
            }
        }
    default:
        var v15 float64 = 0.0
        _ = v15
    }
    var v16 int = 127
    if false {
        v17 := v3[(v16 << ((2) & 63))]
        fmt.Println(v16, v4, v4)
        _ = v17
    }
    v5 = 128
    v18 := GenS2{F0: 746}
    var v19 uint64 = 1
    v20 := GenS3{F0: (v16 << ((18446744073709551615) & 63))}
    if (v5 != (v5 & 254)) {
        v21 := []int{42, (v16 << ((2) & 63)), v16}
        _ = v21
    } else {
        v4 = (("\n\t" + ("" + ("" + "hello"))) + v4)
        var v22 float32 = 0.0
        var v23 uint = 65535
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := []int{}
        _ = fi24
        _ = v25
    }
    switch -720 {
    case 128:
        var v26 int64 = 0
        _ = v26
    case 420:
        v27 := []int{(v16 >> ((uint(v5)) & 63)), (v16 * v16)}
        _ = v27
    default:
        var v28 string = "foo bar"
        _ = v28
    }
    var v29 int32 = 0
    if false {
        for fi30 := 0; fi30 < 14; fi30++ {
            if (!(!true)) {
                v31 := GenS4{F0: v5, F1: 0}
                _ = v31
            }
            var v32 float32 = (1.5 * ((float32(1.5) * ((float32(((3.14159 + (1.5 * (((0.0 + 1e10) * 0.0) * (((0.0 - 3.14159) + (0.1 / (((1.5 + 1e10) / 1e10) - ((((0.1 / -2.25) / 0.1) / -2.25) / 1e10)))) * (0.0 * (0.1 * 3.14159)))))) - 3.14159)) + -2.25) / 3.14159)) / 3.14159))
            var v33 uint8 = v5
            _ = fi30
            _ = v32
            _ = v33
        }
        v4 = v4
        v34 := GenS3{F0: (v16 & (-(v16)))}
        _ = v34
    }
    var v35 uint8 = (v5 / (((v5 >> ((318) & 63))) | 1))
    for fi36 := 0; fi36 < 12; fi36++ {
        var v37 uint64 = v19
        var v38 uint64 = (v37 | v37)
        _ = fi36
        _ = v37
        _ = v38
    }
    v39 := GenS2{F0: 42}
    v5 += 1
    fmt.Println(v16, v29, v5)
    for fi40 := 0; fi40 < 15; fi40++ {
        delete(v1, 657)
        _ = fi40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v5
}
