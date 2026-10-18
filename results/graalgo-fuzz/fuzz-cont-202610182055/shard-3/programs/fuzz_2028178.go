package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint }
func genFn0(p0 bool, p1 uint64, p2 uint64) int64 {
    if true {
        v5 := []int{32767, 2147483647, 2147483647, 350}
        switch -2147483648 {
        case -9223372036854775808:
            var v6 int64 = 0
            _ = v6
        case -1:
            v7 := GenS1{F0: "\n\t", F1: (p2 ^ 42), F2: 18446744073709551615}
            _ = v7
        default:
            for fi8 := 0; fi8 < 2; fi8++ {
                v5 = append(v5, (fi8 + (fi8 % (((fi8 / ((fi8) | 1))) | 1))))
                p1 -= (^(p2))
                _ = fi8
            }
        }
        _ = v5
    } else {
        p1 += p2
        var v9 float32 = 3.14159
        v9 += float32(0.1)
        _ = v9
    }
    if (p0 || (-32768 < 258)) {
        v10 := GenS1{F0: ("hello" + ("" + "\n\t")), F1: (^(p2)), F2: 65535}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int32 = 1
    var v2 string = "foo bar"
    v3 := make(map[uint32]uint32)
    var v4 int32 = v1
    var v11 uint32 = 4294967295
    v11 -= 253
    _ = len(v3)
    fmt.Println(v11)
    switch 637 {
    case 0:
        var v12 uint = 255
        _ = v12
    default:
        var v13 int16 = -930
        _ = v13
    }
    switch "" {
    case "foo bar":
        switch v2 {
        case "a":
            if true {
                if false {
                    switch v2 {
                    case "hello":
                        var v14 uint64 = 30
                        _ = v14
                    case "\n\t":
                        v1 = v4
                    }
                }
                switch (((("\n\t" + v2) + v2) + "foo bar") + v2) {
                case "hello":
                    v15 := GenS1{F0: v2, F1: 0, F2: 0}
                    _ = v15
                case "":
                    v16 := GenS1{F0: v2, F1: 859, F2: 1}
                    _ = v16
                default:
                    var v17 int8 = 0
                    _ = v17
                }
            } else {
                var v18 uint32 = (v11 ^ (v11 >> ((1) & 63)))
                _ = v18
            }
        case "\n\t":
            var v19 uint32 = (^(v11))
            _ = v19
        case "hello":
            var v20 float32 = ((0.1 / (3.14159 - -2.25)) / 1.0)
            _ = v20
        }
    case "\n\t":
        var v21 uint = 42
        _ = v21
    case "a":
        fmt.Println(v1)
    }
    v22 := []int{32768, 4294967295}
    v11 = 0
    v11 = v11
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 float32 = 0.0
        _ = len(v3)
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
