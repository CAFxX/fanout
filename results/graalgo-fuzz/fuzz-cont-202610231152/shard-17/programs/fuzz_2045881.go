package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 string; F1 int8; F2 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 uint }
func genFn0(p0 bool, p1 uint16, p2 int8) uint64 {
    v7 := []int{}
    var v8 float32 = (3.14159 + ((-2.25 + (1.5 / 1e10)) + (float32(3.14159) + -2.25)))
    switch 127 {
    case 128:
        var v9 int8 = p2
        _ = v9
    case 368:
        if true {
            v8 += (v8 * float32(3.14159))
            v8 = (v8 * v8)
            var v10 uint16 = 65535
            _ = v10
        } else {
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := GenS1{F0: -31, F1: 255}
                var v13 uint8 = 231
                var v14 uint16 = (p1 % (((p1 * 32767)) | 1))
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
    default:
        fmt.Println(v8, p1, p0)
    }
    p0 = (!(1 <= 783))
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 565
}
func genFn1(p0 uint) uint8 {
    if ((-9223372036854775808 >= 32768) && false) {
        fmt.Println(p0)
        p0 = p0
    } else {
        for fi15 := 0; fi15 < 9; fi15++ {
            p0 -= p0
            p0 = 4294967295
            _ = fi15
        }
        var v16 int8 = -50
        p0 += (p0 << ((1) & 63))
        _ = v16
    }
    var v17 int64 = -1
    v18 := GenS2{F0: "\n\t", F1: 1, F2: 169}
    _ = p0
    _ = v17
    _ = v18
    return 1
}

func main() {
    var v1 uint16 = 1
    v2 := []uint32{120, 639, 655}
    var v3 int = -9223372036854775808
    v4 := make(map[string]bool)
    v5 := make(map[int]string)
    var v6 float64 = (0.0 / 0.1)
    switch (-(v3)) {
    case 4294967295:
        fmt.Println(v6, v1)
    default:
        var v19 bool = (true && false)
        _ = v19
    }
    var v20 int = v3
    var v21 float32 = -2.25
    _ = len(v4)
    var v22 int16 = 32767
    v23 := GenS2{F0: "a", F1: 127, F2: 127}
    switch "hello" {
    case "hello":
        v24 := v5[(v3 * v3)]
        _ = v24
    case "a":
        switch (v3 / (((v3 | v20)) | 1)) {
        case 42:
            fmt.Println(v1, v22, v6)
        }
    }
    v20 = (v20 >> ((0) & 63))
    v2 = append(v2, 2147483647)
    for fi25 := 0; fi25 < 14; fi25++ {
        var v26 uint16 = 42
        var v27 string = ("" + "foo bar")
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        v1 += (v1 * v1)
        v29 := GenS3{F0: v20}
        var v30 int8 = 42
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 int16 = (v22 | v22)
    v32 := GenS4{F0: 0, F1: 0}
    var v33 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
