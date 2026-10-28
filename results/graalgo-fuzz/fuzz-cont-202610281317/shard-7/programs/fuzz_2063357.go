package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint }
func genFn0(p0 int8, p1 bool) uint64 {
    fmt.Println(p1, p1)
    var v7 float32 = 3.14159
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: -812, F1: 42}
        _ = fi8
        _ = v9
    }
    v7 += 0.0
    _ = p0
    _ = p1
    _ = v7
    return 0
}
func genFn1(p0 uint64) uint64 {
    v10 := GenS2{F0: 792}
    switch "hello" {
    case "\n\t":
        v11 := []int{}
        _ = v11
    case "a":
        fmt.Println(p0)
    case "hello":
        fmt.Println(p0)
    }
    p0 += p0
    var v12 uint16 = 65535
    v13 := GenS3{F0: false, F1: 2147483647}
    _ = p0
    _ = v10
    _ = v12
    _ = v13
    return (p0 - uint64(v12))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -677
    var v3 uint32 = 0
    var v4 uint64 = 18446744073709551615
    var v5 float32 = 1.5
    v6 := []string{"hello", ("a" + "hello"), ""}
    v2 = (v2 % (((v2 - -1)) | 1))
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 int8 = 42
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 uint16 = 2
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 2; fi18++ {
            _ = len(v1)
            _ = fi18
        }
        _ = fi14
        _ = v15
    }
    v3 += v3
    v3 -= (v3 % ((2147483648) | 1))
    var v19 float32 = v5
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (!(2 != (v4 >> ((0) & 63)))) {
        fmt.Println(v4, v5, v2)
    }
    var v20 uint = 0
    var v21 float64 = 3.14159
    v1[(v3 / ((v3) | 1))] = false
    v2 -= v2
    v19 -= (((float32((v21 - -2.25)) / v19) * float32((v21 / 1e10))) / ((1.5 * -2.25) + v19))
    v22 := GenS4{F0: (v20 << (((v20 | v20)) & 63))}
    var v23 float32 = v19
    v24 := GenS2{F0: v2}
    switch -2147483648 {
    case 202:
        for fi25 := 0; fi25 < 11; fi25++ {
            v21 += (float64(v19) - 1e10)
            _ = fi25
        }
    case -2147483648:
        if (!true) {
            var v26 int16 = -24
            for fi27 := 0; fi27 < 3; fi27++ {
                v26 = (v26 << (((v20 ^ (v20 + v20))) & 63))
                _ = fi27
            }
            v28 := GenS4{F0: uint(v4)}
            _ = v26
            _ = v28
        }
    case -1:
        v2 = (v2 * (v2 % ((v2) | 1)))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
