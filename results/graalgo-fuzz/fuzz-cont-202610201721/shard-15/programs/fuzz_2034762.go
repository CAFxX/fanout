package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int64; F2 int64 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint64 }

func main() {
    var v1 int = 0
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    var v4 int8 = 127
    v5 := make(map[int]string)
    var v6 int32 = 2147483647
    switch (("hello" + (("\n\t" + "hello") + ("a" + "foo bar"))) + "\n\t") {
    case "\n\t":
        var v7 uint32 = 2147483648
        _ = v7
    case "hello":
        v8 := []int{9223372036854775807}
        _ = v8
    case "a":
        v4 += -128
    }
    if false {
        for fi9 := 0; fi9 < 12; fi9++ {
            v1 = v1
            v1 -= (-(fi9))
            _ = fi9
        }
        switch (v1 & v1) {
        case -9223372036854775808:
            if false {
                v6 = v6
            } else {
                fmt.Println(v6, v4, v4)
                fmt.Println(v4, v4)
            }
        }
    }
    var v10 uint16 = 0
    switch 984 {
    case 9223372036854775807:
        var v11 uint = 0
        _ = v11
    case 42:
        switch v1 {
        case -744:
            v12 := GenS1{F0: (^(v10))}
            _ = v12
        case 127:
            if (!true) {
                v6 += (v6 | int32(v1))
                v13 := GenS2{F0: "", F1: -272, F2: -356}
                _ = v13
            } else {
                fmt.Println(v1)
                switch -1 {
                case 1:
                    fmt.Println(v6)
                case 923:
                    v14 := GenS2{F0: "a", F1: 949, F2: -1}
                    _ = v14
                default:
                    if (1 >= 2) {
                        fmt.Println(v6, v4)
                        v15 := GenS3{F0: v4, F1: (^(v10)), F2: 2}
                        _ = v15
                    } else {
                        v1 = (^(v1))
                        var v16 float64 = 3.14159
                        _ = v16
                    }
                }
            }
        default:
            v17 := v5[v1]
            _ = v17
        }
    case 32767:
        v10 += v10
    }
    v18 := []int{(v1 * (v1 / ((v1) | 1))), int(v4)}
    delete(v3, 0)
    v19 := GenS3{F0: v4, F1: (v10 << ((65535) & 63)), F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
