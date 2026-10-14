package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 uint32; F1 bool; F2 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint8; F1 int32; F2 int }
func genFn0(p0 int, p1 uint32, p2 string) bool {
    var v7 int32 = 2147483647
    v8 := GenS1{F0: 1}
    if true {
        fmt.Println(p2, p2, p0)
        var v9 uint64 = 324
        if true {
            v10 := []int{32767, p0, int(v9), p0}
            _ = v10
        } else {
            var v11 float64 = ((-2.25 * -2.25) * 3.14159)
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 uint32 = 793
                var v14 int64 = 1
                _ = fi12
                _ = v13
                _ = v14
            }
            v15 := []int{p0, (p0 / ((128) | 1))}
            _ = v11
            _ = v15
        }
        _ = v9
    } else {
        var v16 uint64 = 42
        if true {
            p0 = -844
            v17 := GenS2{F0: 2147483647}
            v18 := []int{p0}
            _ = v17
            _ = v18
        }
        if false {
            v19 := []int{2147483647, (p0 | p0), (p0 | int(v16))}
            var v20 uint64 = v16
            switch "" {
            case "a":
                for fi21 := 0; fi21 < 19; fi21++ {
                    for fi22 := 0; fi22 < 16; fi22++ {
                        var v23 int = (^(fi21))
                        _ = fi22
                        _ = v23
                    }
                    _ = fi21
                }
            case "foo bar":
                v16 = v16
            }
            _ = v19
            _ = v20
        }
        _ = v16
    }
    var v24 int32 = (^(v7))
    var v25 uint16 = 38
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    _ = v7
    _ = v8
    return false
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{198, 2147483647}
    v3 := []int{1}
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]int)
    v26 := GenS3{F0: 64}
    v27 := GenS4{F0: 56, F1: 687}
    var v28 int = 0
    v28 -= (v28 / ((4294967295) | 1))
    switch ("foo bar" + "foo bar") {
    case "a":
        v29 := GenS4{F0: 255, F1: v28}
        _ = v29
    case "foo bar":
        v30 := GenS5{F0: 42, F1: false, F2: 0}
        _ = v30
    case "":
        switch "a" {
        case "hello":
            var v31 uint16 = 858
            _ = v31
        default:
            if (false || (!(!(18446744073709551615 == 18446744073709551615)))) {
                v28 = v28
            } else {
                if false {
                    v28 -= -2147483648
                    fmt.Println(v28)
                    if true {
                        v28 -= (v28 ^ (v28 % (((v28 << ((4294967295) & 63))) | 1)))
                        var v32 int8 = 45
                        var v33 float32 = -2.25
                        _ = v32
                        _ = v33
                    }
                }
            }
        }
    }
    fmt.Println(v28)
    v34 := GenS6{F0: (false || ((true || false) && (v28 <= (v28 / ((v28) | 1)))))}
    fmt.Println(v28)
    fmt.Println(v28)
    for fi35 := 0; fi35 < 10; fi35++ {
        _ = len(v4)
        _ = fi35
    }
    if len(v2) > 0 {
        v2[0] = 2
    }
    var v36 uint16 = 202
    switch "hello" {
    case "foo bar":
        v37 := GenS1{F0: 42}
        _ = v37
    case "\n\t":
        fmt.Println(v28, v28)
    default:
        v38 := GenS7{F0: 128, F1: 396, F2: v28}
        _ = v38
    }
    v36 -= 2
    v36 += (v36 + 0)
    fmt.Println(v28)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
