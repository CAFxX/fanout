package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 int32; F1 int64 }
func genFn0(p0 int16) uint16 {
    v6 := []int{-678, 395, -13, 2147483647}
    _ = len(v6)
    if true {
        var v7 bool = (p0 > p0)
        _ = v7
    }
    var v8 int64 = -487
    _ = p0
    _ = v6
    _ = v8
    return 0
}
func genFn1(p0 uint32) uint64 {
    v9 := GenS1{F0: "hello", F1: 42}
    if true {
        switch 127 {
        case 1:
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 uint8 = 255
                fmt.Println(fi10)
                _ = fi10
                _ = v11
            }
        case 2147483647:
            if (-818 < 42) {
                fmt.Println(p0)
            } else {
                p0 -= p0
                var v12 bool = (0 < 9223372036854775807)
                _ = v12
            }
        case 4294967295:
            fmt.Println(p0)
        }
        var v13 int32 = -27
        _ = v13
    }
    fmt.Println(p0)
    p0 += p0
    _ = p0
    _ = v9
    return 148
}

func main() {
    v1 := []uint32{0, 460, 22, 2147483648}
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]string)
    var v5 uint16 = 62
    if len(v1) > 0 {
        v1[0] = 0
    }
    v1 = append(v1, 4294967295)
    v5 = v5
    v14 := GenS2{F0: (v5 / (((v5 + genFn0(-1))) | 1)), F1: 1, F2: 0}
    var v15 uint64 = 18446744073709551615
    if (v15 != (v15 | v15)) {
        v16 := GenS3{F0: true, F1: v5}
        var v17 int16 = 32767
        fmt.Println(v17, v5, v15)
        _ = v16
        _ = v17
    } else {
        if true {
            v15 = 4294967295
            if true {
                v15 += (v15 * v15)
            }
            var v18 int64 = 0
            _ = v18
        }
    }
    var v19 uint8 = 255
    v20 := GenS4{F0: 2147483647, F1: 9223372036854775807}
    v21 := GenS2{F0: 255, F1: (v15 + genFn1(507)), F2: 0}
    for fi22 := 0; fi22 < 7; fi22++ {
        switch (("foo bar" + "") + ("a" + "\n\t")) {
        case "hello":
            v23 := v3[85]
            _ = v23
        }
        _ = fi22
    }
    v15 += (^(v15))
    var v24 float32 = (float32(0.1) / -2.25)
    v15 += (v15 * (v15 * (v15 + (v15 + v15))))
    fmt.Println(v19)
    v25 := GenS2{F0: (v5 >> ((2) & 63)), F1: (^(v15)), F2: -794}
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 float32 = v24
        if false {
            fmt.Println(v27, fi26, v24)
        } else {
            v28 := GenS3{F0: (false && (2 < uint(v19))), F1: v5}
            _ = v28
        }
        v1 = append(v1, uint32(v19))
        _ = fi26
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
