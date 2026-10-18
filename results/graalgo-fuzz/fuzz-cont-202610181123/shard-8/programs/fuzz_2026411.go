package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32 }
func genFn0() int64 {
    var v7 float64 = float64(3.14159)
    for fi8 := 0; fi8 < 20; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    _ = v7
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = -830
    var v3 string = ""
    var v4 string = (v3 + (v3 + v3))
    var v5 uint16 = 1
    v6 := []int{1}
    if true {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        switch (v2 - (v2 >> ((4294967295) & 63))) {
        case 2147483647:
            fmt.Println(v4, v4)
        case 1:
            for fi9 := 0; fi9 < 8; fi9++ {
                v1[v2] = 481
                _ = fi9
            }
        default:
            var v10 uint16 = (v5 & 255)
            _ = v10
        }
    }
    v11 := GenS1{F0: v4, F1: v2, F2: -308}
    v2 = v2
    var v12 uint32 = 650
    v13 := GenS1{F0: "a", F1: (v2 % ((666) | 1)), F2: int32(v2)}
    if false {
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 uint32 = 2147483647
            switch "hello" {
            case "foo bar":
                switch (fi14 + (-(v2))) {
                case 222:
                    fmt.Println(v4)
                case -166:
                    v16 := GenS2{F0: 2, F1: 2, F2: 0}
                    _ = v16
                case -913:
                    fmt.Println(v2, v12)
                default:
                    var v17 uint = 815
                    _ = v17
                }
            case "":
                var v18 int64 = 9223372036854775807
                _ = v18
            default:
                v19 := GenS2{F0: 0, F1: 2, F2: -306}
                _ = v19
            }
            _ = fi14
            _ = v15
        }
        for fi20 := 0; fi20 < 19; fi20++ {
            v21 := GenS3{F0: 1}
            for fi22 := 0; fi22 < 8; fi22++ {
                _ = len(v6)
                v23 := GenS1{F0: "hello", F1: fi20, F2: 638}
                _ = fi22
                _ = v23
            }
            _ = fi20
            _ = v21
        }
        v12 += (v12 / ((2147483647) | 1))
    } else {
        switch "a" {
        case "hello":
            v24 := GenS4{F0: 0}
            _ = v24
        }
    }
    var v25 int32 = 42
    var v26 uint = 42
    fmt.Println(v25, v2, v4)
    fmt.Println(v5, v3, v25)
    v27 := v1[(v2 * 4294967295)]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
