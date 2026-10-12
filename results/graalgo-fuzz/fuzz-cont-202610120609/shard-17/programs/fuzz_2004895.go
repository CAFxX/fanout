package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 string; F2 uint64 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int) bool {
    p0 += p0
    fmt.Println(p0)
    p0 -= p0
    v5 := GenS1{F0: -2147483648}
    _ = p0
    _ = v5
    return (1 < 2)
}

func main() {
    var v1 uint16 = 957
    v2 := []int{321}
    var v3 int64 = -166
    var v4 uint32 = 4294967295
    fmt.Println(v3)
    var v6 int = 9223372036854775807
    if false {
        fmt.Println(v6, v1, v1)
        switch 32767 {
        case -898:
            v7 := GenS2{F0: false}
            _ = v7
        case 2147483647:
            for fi8 := 0; fi8 < 1; fi8++ {
                switch "foo bar" {
                case "":
                    var v9 float64 = (-2.25 + ((-2.25 + (0.0 * -2.25)) * 0.1))
                    _ = v9
                case "hello":
                    switch (("a" + ("foo bar" + "foo bar")) + "hello") {
                    case "hello":
                        var v10 int = (^(fi8))
                        _ = v10
                    case "a":
                        v11 := GenS2{F0: false}
                        _ = v11
                    case "foo bar":
                        var v12 uint16 = (v1 % (((^(v1))) | 1))
                        _ = v12
                    }
                case "a":
                    if false {
                        _ = len(v2)
                        var v13 uint = 42
                        _ = v13
                    }
                default:
                    var v14 int8 = 100
                    _ = v14
                }
                var v15 int32 = int32(v1)
                _ = fi8
                _ = v15
            }
        default:
            var v16 uint32 = (v4 - (v4 / ((2) | 1)))
            _ = v16
        }
        var v17 int64 = (v3 | (^(v3)))
        _ = v17
    } else {
        var v18 string = ("" + "")
        _ = v18
    }
    var v19 int32 = 807
    switch "\n\t" {
    case "\n\t":
        v4 -= (v4 - 886)
    default:
        v20 := GenS2{F0: false}
        _ = v20
    }
    v1 += v1
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 13; fi22++ {
            v23 := GenS3{F0: 322, F1: "", F2: 2}
            _ = fi22
            _ = v23
        }
        v24 := GenS4{F0: v19}
        _ = fi21
        _ = v24
    }
    switch "foo bar" {
    case "":
        var v25 int = v6
        _ = v25
    case "hello":
        if (3.14159 >= (float64(1e10) * (((3.14159 / -2.25) * float64(1e10)) + 0.1))) {
            var v26 int64 = 0
            v6 += (v6 * v6)
            _ = v26
        }
    case "a":
        var v27 float32 = ((float32(0.1) - 1e10) / (3.14159 / 0.1))
        _ = v27
    }
    v2 = append(v2, (-(v6)))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
