package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int16; F2 int32 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 string, p1 string) int32 {
    var v7 uint16 = 2
    v7 = v7
    var v8 int8 = -128
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}
func genFn1() int {
    switch "foo bar" {
    case "a":
        v9 := GenS1{F0: 42}
        _ = v9
    case "hello":
        var v10 float32 = 1.5
        _ = v10
    case "":
        var v11 uint16 = 337
        _ = v11
    default:
        var v12 int64 = -111
        _ = v12
    }
    var v13 uint32 = 2
    v13 = (v13 % ((v13) | 1))
    _ = v13
    return 4294967295
}

func main() {
    var v1 uint64 = 2
    var v2 float64 = (1e10 - 1e10)
    var v3 uint = 42
    var v4 uint32 = uint32(v3)
    v5 := []int{-540}
    var v6 uint8 = 182
    v1 -= 2
    if (!false) {
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 string = "a"
            v16 := GenS1{F0: 32767}
            _ = fi14
            _ = v15
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            v4 += v4
            v5 = append(v5, 42)
            var v18 int = (fi17 * fi17)
            _ = fi17
            _ = v18
        }
    }
    if true {
        var v19 int8 = 127
        _ = v19
    }
    v20 := GenS1{F0: 65535}
    var v21 uint32 = v4
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 int16 = 1
        _ = fi22
        _ = v23
    }
    switch "\n\t" {
    case "foo bar":
        if true {
            var v24 uint8 = (v6 >> ((597) & 63))
            var v25 bool = true
            _ = len(v5)
            _ = v24
            _ = v25
        }
    case "hello":
        if (230 < 42) {
            for fi26 := 0; fi26 < 19; fi26++ {
                var v27 int32 = -1
                var v28 uint16 = 905
                _ = fi26
                _ = v27
                _ = v28
            }
            v29 := GenS1{F0: uint16(v6)}
            _ = v29
        } else {
            var v30 uint32 = (v21 ^ v4)
            if (32767 <= 1000) {
                var v31 int64 = 9223372036854775807
                v32 := GenS2{F0: (v3 | v3)}
                _ = v31
                _ = v32
            }
            _ = v30
        }
    default:
        fmt.Println(v4, v3, v1)
    }
    if (!true) {
        v33 := GenS3{F0: 0, F1: 32767, F2: 2147483647}
        if len(v5) > 0 {
            v5[0] = 127
        }
        v4 = v4
        _ = v33
    }
    var v34 uint16 = 255
    v35 := GenS4{F0: v34}
    v36 := GenS4{F0: v34}
    var v37 uint64 = v1
    v3 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
