package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 bool }
type GenS2 struct { F0 uint; F1 bool; F2 int8 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint32 }
func genFn0(p0 int) uint {
    switch ("\n\t" + "hello") {
    case "hello":
        var v6 uint64 = 506
        _ = v6
    case "a":
        var v7 int64 = 0
        _ = v7
    default:
        switch "hello" {
        case "a":
            var v8 uint32 = 42
            _ = v8
        default:
            v9 := []int{-715}
            _ = v9
        }
    }
    v10 := GenS1{F0: 0, F1: 965, F2: (65535 < 42)}
    p0 -= -28
    p0 -= (p0 - (p0 | p0))
    _ = p0
    _ = v10
    return 374
}
func genFn1() uint32 {
    var v11 uint32 = 0
    for fi12 := 0; fi12 < 14; fi12++ {
        v11 = (v11 & 638)
        var v13 float32 = ((0.0 + float32(0.0)) + 1.5)
        _ = fi12
        _ = v13
    }
    v14 := []int{-2147483648, 32768, 74, -641}
    _ = v11
    _ = v14
    return v11
}

func main() {
    var v1 float32 = (-2.25 + float32((3.14159 / (-2.25 - float64(1e10)))))
    var v2 int16 = -1
    var v3 uint16 = 1
    v4 := []int{int(v3)}
    v5 := make(map[uint32]bool)
    v1 += float32(float64((1.5 - v1)))
    var v15 int64 = -1
    var v16 int64 = (v15 ^ v15)
    v17 := GenS2{F0: uint(v3), F1: false, F2: 100}
    var v18 float64 = 1.5
    v3 = (v3 * v3)
    for fi19 := 0; fi19 < 2; fi19++ {
        v4 = append(v4, 746)
        v4 = append(v4, fi19)
        var v20 uint32 = 2
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        var v22 uint8 = 42
        v15 = -9223372036854775808
        v23 := GenS3{F0: 42, F1: 1, F2: 2147483647}
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS2{F0: genFn0(532), F1: (v3 <= v3), F2: 100}
    switch "\n\t" {
    case "\n\t":
        v2 -= v2
    case "a":
        v3 += (v3 | (v3 >> ((genFn0(32767)) & 63)))
    case "foo bar":
        fmt.Println(v18, v2, v2)
    default:
        if false {
            var v25 float32 = v1
            v16 = (v16 % ((v15) | 1))
            fmt.Println(v15)
            _ = v25
        }
    }
    var v26 string = ("hello" + ("a" + "hello"))
    var v27 string = "\n\t"
    for fi28 := 0; fi28 < 20; fi28++ {
        for fi29 := 0; fi29 < 1; fi29++ {
            v1 = v1
            if (!false) {
                var v30 int64 = v16
                var v31 int32 = 1
                switch 2147483647 {
                case 127:
                    v16 = 930
                }
                _ = len(v4)
                if (!true) {
                    v1 = (v1 + (3.14159 + (0.1 * 0.1)))
                } else {
                    v3 -= (v3 | v3)
                    var v32 int8 = 0
                    _ = v32
                }
                _ = v30
                _ = v31
            }
            var v33 int8 = -105
            v16 += v15
            var v34 uint8 = 127
            _ = fi29
            _ = v33
            _ = v34
        }
        var v35 uint16 = v3
        _ = fi28
        _ = v35
    }
    var v36 uint16 = v3
    var v37 uint16 = v36
    for fi38 := 0; fi38 < 9; fi38++ {
        v39 := GenS2{F0: genFn0(fi38), F1: true, F2: 27}
        _ = fi38
        _ = v39
    }
    v37 = 476
    v40 := GenS2{F0: 718, F1: (-506 < 1), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v5
}
