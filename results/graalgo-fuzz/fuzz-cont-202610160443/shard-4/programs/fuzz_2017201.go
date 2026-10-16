package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 int; F2 uint16 }

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 748
    var v3 string = ("foo bar" + "a")
    var v4 float32 = 1.5
    var v5 int32 = 1
    v6 := make(map[int]uint32)
    switch 2147483647 {
    case -43:
        v2 -= v2
    case 32767:
        var v7 uint64 = 0
        _ = v7
    }
    v2 -= (v2 >> ((42) & 63))
    for fi8 := 0; fi8 < 11; fi8++ {
        v6[fi8] = 242
        for fi9 := 0; fi9 < 11; fi9++ {
            v3 = (v3 + "hello")
            v10 := []int{(fi9 * 9223372036854775807), fi9, fi9}
            v4 -= float32(float64(v4))
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v4 = ((v4 - float32((-2.25 * 1e10))) - 0.0)
    if true {
        v11 := []int{-2147483648, -68, 2147483647, 1}
        v12 := GenS1{F0: (v2 + v2), F1: 1}
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 17; fi13++ {
            v3 = v3
            _ = fi13
        }
        var v14 float32 = v4
        _ = v14
    }
    switch -2147483648 {
    case -414:
        var v15 int32 = v5
        _ = v15
    case -9223372036854775808:
        v16 := GenS2{F0: false, F1: 127, F2: 18446744073709551615}
        _ = v16
    case 32768:
        v3 = v3
    default:
        var v17 int32 = v5
        _ = v17
    }
    if false {
        v18 := GenS3{F0: false, F1: 42}
        _ = v18
    }
    var v19 int64 = 9223372036854775807
    v2 = v2
    var v20 int16 = 1000
    v21 := GenS4{F0: false}
    var v22 int = 9223372036854775807
    v23 := GenS5{F0: 1, F1: v22, F2: 835}
    fmt.Println(v2, v20, v19)
    for fi24 := 0; fi24 < 10; fi24++ {
        v3 = ("\n\t" + v3)
        if true {
            var v25 float64 = -2.25
            for fi26 := 0; fi26 < 7; fi26++ {
                var v27 int64 = (v19 >> ((823) & 63))
                _ = fi26
                _ = v27
            }
            var v28 uint8 = 2
            _ = v25
            _ = v28
        }
        _ = fi24
    }
    fmt.Println(v22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
