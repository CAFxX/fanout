package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int32 }

func main() {
    var v1 int = 128
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    if (!(481 != 2)) {
        for fi4 := 0; fi4 < 1; fi4++ {
            v5 := GenS1{F0: 482, F1: 65535, F2: -145}
            var v6 string = "\n\t"
            _ = fi4
            _ = v5
            _ = v6
        }
        var v7 bool = true
        var v8 int32 = -1
        _ = v7
        _ = v8
    }
    fmt.Println(v1)
    v1 = 128
    v1 = (v1 >> ((255) & 63))
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int16 = 1000
        switch (v1 << ((2) & 63)) {
        case 9223372036854775807:
            v11 := []int{v1, (v1 / ((4294967295) | 1)), (v1 * fi9)}
            _ = v11
        case 899:
            if true {
                for fi12 := 0; fi12 < 16; fi12++ {
                    v1 = -842
                    _ = fi12
                }
                v13 := GenS1{F0: 794, F1: 0, F2: 1}
                for fi14 := 0; fi14 < 10; fi14++ {
                    switch "" {
                    case "hello":
                        if false {
                            v15 := []int{-90, fi9, 32767}
                            var v16 int8 = 76
                            _ = v15
                            _ = v16
                        }
                    case "":
                        var v17 int64 = 42
                        _ = v17
                    default:
                        fmt.Println(v1, fi14)
                    }
                    _ = fi14
                }
                _ = v13
            }
        default:
            var v18 float32 = (0.1 / (3.14159 / 0.1))
            _ = v18
        }
        _ = fi9
        _ = v10
    }
    v1 += v1
    if (1 >= 529) {
        var v19 int = 1
        _ = v19
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        switch "hello" {
        case "hello":
            fmt.Println(v1, v1)
        }
        var v21 uint16 = 32767
        _ = fi20
        _ = v21
    }
    var v22 float32 = ((3.14159 + (-2.25 / (-2.25 * (-2.25 - 1e10)))) - 0.1)
    if true {
        for fi23 := 0; fi23 < 20; fi23++ {
            if (18446744073709551615 >= 4294967295) {
                var v24 uint32 = 880
                v22 = float32((0.0 - 0.0))
                switch "" {
                case "a":
                    for fi25 := 0; fi25 < 16; fi25++ {
                        var v26 int = fi23
                        var v27 int16 = 42
                        var v28 int8 = 1
                        _ = fi25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                case "hello":
                    var v29 uint8 = 231
                    _ = v29
                case "\n\t":
                    fmt.Println(v1, fi23)
                }
                _ = v24
            }
            _ = fi23
        }
        v30 := GenS2{F0: "foo bar"}
        _ = v30
    } else {
        var v31 int64 = -9223372036854775808
        _ = v31
    }
    v32 := GenS3{F0: 65535, F1: -2147483648}
    var v33 uint = 4294967295
    if false {
        fmt.Println(v1)
    }
    fmt.Println(v33, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v33
}
