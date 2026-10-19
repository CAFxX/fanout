package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int32; F2 int64 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "\n\t"
    v3 := make(map[int]int)
    var v4 float64 = (3.14159 * 0.1)
    v2 = ((v2 + "foo bar") + v2)
    switch "foo bar" {
    case "a":
        v2 = (v2 + "foo bar")
    case "hello":
        if false {
            fmt.Println(v4, v4)
            var v5 int32 = -1
            for fi6 := 0; fi6 < 7; fi6++ {
                v5 = (-(v5))
                _ = fi6
            }
            _ = v5
        }
    default:
        var v7 string = "foo bar"
        _ = v7
    }
    if ((!(-9223372036854775808 == 9223372036854775807)) || (1 > 704)) {
        if false {
            var v8 uint16 = 42
            _ = v8
        } else {
            v9 := GenS1{F0: 65535, F1: 2, F2: 33}
            v2 = v2
            var v10 int16 = 204
            _ = v10
            _ = v9
        }
    } else {
        if false {
            if (685 != 977) {
                var v11 int = -63
                _ = v11
            }
        } else {
            v12 := GenS2{F0: 127, F1: 1, F2: -9223372036854775808}
            _ = v12
        }
        var v13 int = 42
        for fi14 := 0; fi14 < 1; fi14++ {
            for fi15 := 0; fi15 < 14; fi15++ {
                fmt.Println(v2)
                if (!(33 >= 24)) {
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v2, v2, fi15)
                    case "a":
                        v16 := GenS3{F0: 9223372036854775807}
                        delete(v1, ((v2 + v2) + (v2 + v2)))
                        _ = v16
                    case "hello":
                        var v17 int8 = 42
                        _ = v17
                    }
                }
                v2 = (v2 + "a")
                _ = fi15
            }
            var v18 uint = 766
            _ = fi14
            _ = v18
        }
        _ = v13
    }
    v1[v2] = 4294967295
    var v19 int64 = 9223372036854775807
    fmt.Println(v19)
    v20 := []int{-2147483648, 4294967295, 232}
    _ = len(v20)
    v2 = (v2 + v2)
    v21 := GenS3{F0: (v19 ^ (v19 % ((v19) | 1)))}
    v4 -= (v4 - float64(0.0))
    if (v19 < v19) {
        v4 += v4
        v22 := GenS2{F0: 128, F1: 0, F2: v19}
        _ = v22
    }
    _ = len(v3)
    var v23 uint32 = 767
    if true {
        v24 := GenS3{F0: (v19 * v19)}
        v2 = "a"
        v2 = v2
        _ = v24
    }
    var v25 int64 = (v19 & (v19 % ((v19) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v25
    _ = v3
    _ = v4
}
