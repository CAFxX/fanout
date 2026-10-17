package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 string; F2 int }
type GenS3 struct { F0 uint }
func genFn0(p0 int8, p1 string, p2 int64) int8 {
    p0 = (p0 + (p0 % (((p0 % ((-128) | 1))) | 1)))
    v5 := []int{9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -83
}
func genFn1() int32 {
    var v6 float32 = (0.1 + (0.1 + 0.1))
    var v7 int = 32767
    switch (v7 % ((v7) | 1)) {
    case -272:
        var v8 uint8 = 113
        _ = v8
    }
    v9 := []int{v7, (v7 - v7), (v7 % (((v7 >> ((4294967295) & 63))) | 1))}
    _ = v6
    _ = v7
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{}
    var v3 int64 = 1
    var v4 uint16 = 2
    var v10 float32 = ((1e10 * (0.1 - float32(1e10))) + (float32((0.0 + 0.0)) + 1.5))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v1)
    v11 := GenS1{F0: 408}
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 int = fi12
        fmt.Println(fi12, v13, v10)
        _ = fi12
        _ = v13
    }
    if (!false) {
        if (1000 < 705) {
            var v14 float64 = ((3.14159 / 1.0) - float64(v10))
            v14 -= (v14 - 0.0)
            v15 := GenS2{F0: -44, F1: "a", F2: 304}
            _ = v14
            _ = v15
        }
        v10 -= v10
    }
    if false {
        for fi16 := 0; fi16 < 17; fi16++ {
            switch "a" {
            case "\n\t":
                var v17 int32 = genFn1()
                _ = v17
            default:
                v18 := v1["a"]
                _ = v18
            }
            var v19 uint8 = 106
            _ = fi16
            _ = v19
        }
        switch 32767 {
        case 710:
            switch "" {
            case "hello":
                _ = len(v1)
            case "a":
                v20 := GenS1{F0: 42}
                _ = v20
            case "":
                v21 := GenS1{F0: 4294967295}
                _ = v21
            }
        case -677:
            v4 += v4
        default:
            var v22 string = ("hello" + ((("" + "hello") + "\n\t") + "hello"))
            _ = v22
        }
    }
    v23 := GenS3{F0: 255}
    v10 = float32(((((-2.25 - 1.5) - (-2.25 / float64(v10))) * -2.25) * -2.25))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
