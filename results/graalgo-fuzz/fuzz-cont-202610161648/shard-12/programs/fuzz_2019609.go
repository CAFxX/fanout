package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int64; F1 int }

func main() {
    v1 := []uint32{}
    v2 := make(map[string]bool)
    v3 := []string{}
    v4 := []string{("a" + (("a" + "hello") + "")), "hello", ("a" + "a"), (("" + (("a" + "foo bar") + "")) + "\n\t")}
    v5 := GenS1{F0: -1}
    v6 := GenS2{F0: 79, F1: 16}
    var v7 int8 = 127
    v7 = 6
    v2["a"] = (!(0 >= -1))
    v7 = 100
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(v7, fi8)
        if (!true) {
            v7 = (v7 + (v7 % (((v7 / (((v7 & -1)) | 1))) | 1)))
        } else {
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 float32 = 0.0
                v10 -= 1e10
                v10 = 0.0
                _ = fi9
                _ = v10
            }
        }
        fmt.Println(v7, v7)
        _ = fi8
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 float64 = float64(1e10)
        v7 = v7
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 uint32 = 65535
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        fmt.Println(fi15, fi15)
        _ = fi15
    }
    if (42 == 42) {
        if (666 == -889) {
            v16 := GenS3{F0: 639, F1: 127}
            _ = v16
        }
        for fi17 := 0; fi17 < 11; fi17++ {
            if (("\n\t" >= ("" + "hello")) && true) {
                var v18 int32 = 922
                if true {
                    _ = len(v4)
                    var v19 uint8 = 1
                    _ = v19
                }
                _ = v18
            }
            _ = fi17
        }
    }
    switch -430 {
    case 42:
        var v20 float64 = 0.0
        _ = v20
    }
    var v21 float64 = float64(1.5)
    fmt.Println(v21)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
