package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int16; F1 int64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    var v4 int64 = -9223372036854775808
    fmt.Println(v4)
    v5 := GenS1{F0: -9223372036854775808}
    if false {
        var v6 uint16 = 839
        v7 := GenS2{F0: v4, F1: false}
        fmt.Println(v4, v4)
        _ = v6
        _ = v7
    } else {
        v4 = (v4 - (v4 | v4))
    }
    v4 += v4
    v8 := GenS3{F0: 42, F1: 128}
    v9 := []int{}
    for fi10 := 0; fi10 < 6; fi10++ {
        switch fi10 {
        case -2147483648:
            v4 = (v4 * 234)
        }
        _ = fi10
    }
    v11 := GenS2{F0: v4, F1: ((!false) && (27 != 90))}
    switch "\n\t" {
    case "\n\t":
        switch 32767 {
        case 32768:
            if false {
                for fi12 := 0; fi12 < 4; fi12++ {
                    var v13 int64 = -918
                    var v14 int16 = 596
                    delete(v2, 161)
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                fmt.Println(v4)
            } else {
                var v15 float64 = (0.0 * (0.0 / -2.25))
                fmt.Println(v4)
                v4 = (v4 & (-(v4)))
                _ = v15
            }
        default:
            var v16 string = ("hello" + "a")
            _ = v16
        }
    }
    switch 32767 {
    case -722:
        switch 127 {
        case 42:
            v4 += -1
        case 127:
            for fi17 := 0; fi17 < 11; fi17++ {
                fmt.Println(v4)
                _ = fi17
            }
        case 1:
            v18 := GenS4{F0: 2, F1: 42}
            _ = v18
        default:
            v19 := GenS5{F0: -247, F1: -1}
            _ = v19
        }
    case 695:
        if true {
            if ((3.14159 == -2.25) && false) {
                v20 := GenS3{F0: 73, F1: 42}
                _ = v20
            } else {
                v21 := GenS3{F0: 1, F1: 42}
                v4 = v4
                _ = v21
            }
        }
    case -9223372036854775808:
        v1[4294967295] = 4294967295
    }
    switch "" {
    case "hello":
        var v22 uint32 = 65535
        _ = v22
    case "\n\t":
        for fi23 := 0; fi23 < 3; fi23++ {
            v4 += v4
            _ = len(v9)
            _ = fi23
        }
    case "foo bar":
        var v24 float64 = float64(0.0)
        _ = v24
    }
    delete(v1, -9223372036854775808)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
