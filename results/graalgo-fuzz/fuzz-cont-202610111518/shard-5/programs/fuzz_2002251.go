package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    var v3 float32 = (((0.0 - (3.14159 * ((0.0 + 1e10) * 1e10))) * (0.0 * 1e10)) / 1.5)
    var v4 uint = 65535
    v5 := GenS1{F0: true, F1: 27}
    var v6 string = "foo bar"
    fmt.Println(v6)
    var v7 float64 = 0.0
    for fi8 := 0; fi8 < 1; fi8++ {
        switch (v6 + "") {
        case "a":
            var v9 uint64 = 1
            _ = v9
        default:
            v10 := GenS2{F0: 127, F1: 608}
            _ = v10
        }
        switch (fi8 + fi8) {
        case 32768:
            v11 := GenS1{F0: (true && false), F1: 42}
            _ = v11
        case 127:
            var v12 int64 = 9223372036854775807
            _ = v12
        default:
            if (0.0 <= (v7 * float64(((float32(v7) * v3) + (3.14159 - float32(v7)))))) {
                v13 := GenS1{F0: (!true), F1: 42}
                var v14 uint64 = 18446744073709551615
                _ = v13
                _ = v14
            } else {
                v15 := v2[(fi8 | (^(fi8)))]
                v16 := []int{128, v15, (v15 / (((^(fi8))) | 1)), fi8}
                if (-128 >= -128) {
                    for fi17 := 0; fi17 < 17; fi17++ {
                        var v18 int32 = -26
                        v6 = v6
                        var v19 int32 = (v18 % (((^(v18))) | 1))
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    v20 := GenS2{F0: 0, F1: 0}
                    v7 += v7
                    _ = v20
                } else {
                    v15 = 4294967295
                }
                _ = v15
                _ = v16
            }
        }
        if false {
            v21 := GenS2{F0: 127, F1: int32(fi8)}
            _ = v21
        } else {
            if (!(false && (134 != -1))) {
                for fi22 := 0; fi22 < 12; fi22++ {
                    var v23 uint16 = 563
                    var v24 float64 = (((float64(v3) * v7) - v7) + v7)
                    if false {
                        var v25 int32 = -789
                        _ = v25
                    }
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                v26 := GenS3{F0: 343, F1: 42, F2: fi8}
                _ = v26
            } else {
                v27 := []int{-97, -2147483648}
                _ = v27
            }
        }
        _ = fi8
    }
    fmt.Println(v6)
    if (32767 == 0) {
        var v28 uint64 = 4294967295
        _ = v28
    }
    switch ("foo bar" + v6) {
    case "hello":
        v3 = ((0.1 + (v3 / (v3 * 1e10))) / float32((v7 / 1.5)))
    case "foo bar":
        var v29 string = "a"
        _ = v29
    default:
        fmt.Println(v3, v3, v4)
    }
    if (!false) {
        v6 = ("hello" + ((v6 + v6) + "a"))
        fmt.Println(v4, v3, v7)
    }
    v3 -= (v3 + 1e10)
    fmt.Println(v7, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
