package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 bool; F1 bool; F2 int64 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 bool }
func genFn0(p0 string, p1 bool, p2 uint8) int8 {
    if p1 {
        v5 := []int{42, 42}
        if (p1 && p1) {
            var v6 uint16 = 825
            p2 -= (p2 ^ 255)
            fmt.Println(v6, v6)
            _ = v6
        }
        _ = v5
    } else {
        v7 := GenS1{F0: 0, F1: 2147483648}
        _ = v7
    }
    v8 := GenS1{F0: 42, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 100
}

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    var v4 uint64 = 501
    switch 128 {
    case 127:
        if (145 >= 1) {
            var v9 bool = (false && false)
            if v9 {
                for fi10 := 0; fi10 < 11; fi10++ {
                    var v11 bool = v9
                    _ = fi10
                    _ = v11
                }
                var v12 int16 = -32768
                v4 = 18446744073709551615
                _ = v12
            } else {
                v13 := []int{32767, 32767, 1}
                _ = v13
            }
            v9 = true
            _ = v9
        }
    case 9223372036854775807:
        v14 := GenS1{F0: 330, F1: 652}
        _ = v14
    case 128:
        v15 := []int{42, 306}
        _ = v15
    default:
        switch "a" {
        case "foo bar":
            v16 := GenS2{F0: true, F1: ((v4 | 18446744073709551615) >= (^(v4))), F2: 42}
            _ = v16
        case "":
            if (0.1 != (((float32(1e10) - ((-2.25 * (0.1 / 3.14159)) + 1e10)) - (0.0 / 3.14159)) + 0.0)) {
                var v17 uint16 = 1
                _ = v17
            }
        }
    }
    v1 = 1
    v18 := GenS1{F0: 42, F1: 508}
    if false {
        delete(v2, 1)
    } else {
        v19 := GenS3{F0: 42, F1: -73}
        _ = v19
    }
    fmt.Println(v1)
    for fi20 := 0; fi20 < 7; fi20++ {
        switch "foo bar" {
        case "\n\t":
            v21 := GenS3{F0: 18446744073709551615, F1: -1}
            _ = v21
        case "foo bar":
            var v22 uint64 = (v4 + v4)
            _ = v22
        default:
            for fi23 := 0; fi23 < 16; fi23++ {
                v24 := GenS4{F0: (("\n\t" + "") >= ("foo bar" + ("a" + ("\n\t" + ("foo bar" + "")))))}
                fmt.Println(fi23, v1)
                _ = fi23
                _ = v24
            }
        }
        _ = fi20
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        v1 += (v1 | 2147483647)
        v4 += 0
        _ = fi25
    }
    v1 = 2147483647
    v4 = 42
    var v26 uint8 = 110
    var v27 float32 = 1e10
    v27 = (1e10 - 1.5)
    var v28 uint64 = uint64(v26)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
