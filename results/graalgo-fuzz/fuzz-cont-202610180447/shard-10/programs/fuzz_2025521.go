package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 uint; F1 uint64; F2 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 0
    v3 := make(map[string]string)
    switch 700 {
    case -9223372036854775808:
        v4 := GenS1{F0: 473, F1: 1, F2: 2147483648}
        _ = v4
    default:
        v5 := []int{}
        _ = v5
    }
    v6 := GenS2{F0: 833, F1: 2}
    var v7 uint = 18446744073709551615
    if true {
        for fi8 := 0; fi8 < 1; fi8++ {
            for fi9 := 0; fi9 < 10; fi9++ {
                v10 := v1[779]
                v7 = (v7 >> ((v7) & 63))
                if false {
                    if false {
                        var v11 int16 = -1
                        v12 := v1[2147483648]
                        _ = v11
                        _ = v12
                    } else {
                        var v13 int64 = 42
                        var v14 string = "a"
                        var v15 float64 = (3.14159 + float64(0.1))
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                }
                v3["\n\t"] = "a"
                if (!v10) {
                    switch "" {
                    case "a":
                        var v16 int64 = 703
                        _ = v16
                    case "\n\t":
                        v17 := []int{-604, (-(fi9))}
                        _ = v17
                    default:
                        var v18 int64 = -253
                        _ = v18
                    }
                    v7 -= (v7 << ((v7) & 63))
                    fmt.Println(v7, fi8, v10)
                    v10 = v10
                    var v19 string = "\n\t"
                    _ = v19
                } else {
                    var v20 float32 = 0.1
                    var v21 int8 = 0
                    _ = v20
                    _ = v21
                }
                var v22 int64 = -526
                switch (fi9 >> ((652) & 63)) {
                case 9223372036854775807:
                    switch (^(fi9)) {
                    case -9223372036854775808:
                        v23 := GenS1{F0: 0, F1: 943, F2: v2}
                        _ = v23
                    case -138:
                        fmt.Println(v22, v10, fi8)
                        switch ("" + ("a" + "foo bar")) {
                        case "foo bar":
                            v24 := GenS2{F0: (v7 | (v7 % ((v7) | 1))), F1: (v7 << ((v7) & 63))}
                            _ = v24
                        case "hello":
                            v25 := v3["hello"]
                            _ = v25
                        default:
                            v10 = v10
                        }
                        v2 = (v2 ^ v2)
                    case -2147483648:
                        v26 := GenS1{F0: 18446744073709551615, F1: 912, F2: (^(v2))}
                        _ = v26
                    }
                case -9223372036854775808:
                    switch (fi8 * fi9) {
                    case 9223372036854775807:
                        var v27 float32 = 1.5
                        _ = v27
                    }
                case 4294967295:
                    v2 = v2
                }
                v28 := GenS1{F0: 0, F1: -1, F2: v2}
                v7 -= (v7 >> ((uint(v2)) & 63))
                _ = fi9
                _ = v10
                _ = v22
                _ = v28
            }
            _ = fi8
        }
    }
    var v29 int = 128
    _ = len(v1)
    v30 := GenS3{F0: int16(v7), F1: 42}
    if true {
        v2 -= v2
    }
    v2 = (v2 | v2)
    if (!(false && (0 == 2))) {
        v31 := []int{v29, (v29 * (v29 ^ 2147483647))}
        var v32 uint8 = 42
        v33 := GenS4{F0: (^(v32)), F1: 2}
        _ = v31
        _ = v32
        _ = v33
    }
    if false {
        v34 := GenS2{F0: (v7 ^ (v7 - uint(v2))), F1: (v7 % ((256) | 1))}
        var v35 int = (v29 + v29)
        var v36 float32 = (-2.25 * -2.25)
        _ = v34
        _ = v35
        _ = v36
    }
    var v37 string = "hello"
    _ = len(v3)
    v38 := GenS5{F0: (^(v7)), F1: 18446744073709551615, F2: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v6
    _ = v7
}
