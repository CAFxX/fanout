package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int; F1 int32; F2 int8 }
type GenS4 struct { F0 uint32 }
func genFn0() int32 {
    var v5 int8 = -14
    v6 := GenS1{F0: 2147483647, F1: 1}
    v7 := GenS2{F0: true, F1: 52}
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1() uint8 {
    var v8 uint32 = 613
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS3{F0: fi10, F1: -1, F2: 127}
            v8 = (v8 / ((435) | 1))
            v12 := []int{(fi9 % ((4294967295) | 1)), 42, fi10}
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 8; fi13++ {
            if (!(42 <= 722)) {
                switch (fi13 + (fi9 - fi13)) {
                case 1:
                    var v14 uint64 = 452
                    _ = v14
                case 9223372036854775807:
                    var v15 uint8 = 34
                    _ = v15
                case -795:
                    v16 := []int{-264, -1, -2147483648, (fi13 * 201)}
                    _ = v16
                }
                var v17 bool = (526 < 2)
                v18 := []int{(^(fi13))}
                _ = v17
                _ = v18
            }
            if false {
                var v19 float32 = (((float32(1.5) / 1e10) * (-2.25 + 3.14159)) / (-2.25 * 3.14159))
                if false {
                    v20 := []int{(-(fi13)), fi13, 402, (fi13 >> ((42) & 63))}
                    var v21 uint64 = 4294967295
                    _ = v20
                    _ = v21
                }
                v8 = (^(v8))
                _ = v19
            }
            _ = fi13
        }
        _ = fi9
    }
    switch 13 {
    case 42:
        if false {
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := []int{(fi22 ^ fi22), (fi22 * fi22), fi22, fi22}
                if true {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        if ((910 < (v8 + (v8 & v8))) || false) {
                            v25 := GenS4{F0: (v8 & 65535)}
                            _ = v25
                        }
                        var v26 int32 = 1
                        var v27 int8 = 127
                        var v28 float64 = float64(((1.5 * 1e10) * (0.0 * float32(1.5))))
                        v29 := GenS4{F0: v8}
                        _ = fi24
                        _ = v26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                }
                _ = fi22
                _ = v23
            }
        }
    case 9223372036854775807:
        v8 = (v8 - 2147483648)
    case 32768:
        fmt.Println(v8)
    }
    v30 := []int{}
    _ = v30
    _ = v8
    return 127
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = -1
    var v3 string = "\n\t"
    var v4 int32 = 42
    if (2 != 4294967295) {
        switch -733 {
        case -9223372036854775808:
            if ((float64((float32(((-2.25 * 1e10) * (((0.1 - 1.5) * 1.5) + (1e10 * 3.14159)))) * 1e10)) - (0.0 - (0.1 / float64(3.14159)))) != 0.1) {
                v4 = 2147483647
            }
        case -370:
            _ = len(v1)
        }
    }
    switch v3 {
    case "hello":
        var v31 float64 = (float64(1e10) * 0.0)
        _ = v31
    case "\n\t":
        v4 = (^(v4))
    default:
        v3 = v3
    }
    v4 = -858
    v32 := []int{}
    v3 = ("a" + v3)
    switch v3 {
    case "a":
        fmt.Println(v3, v4, v3)
    case "hello":
        var v33 float64 = 0.0
        _ = v33
    default:
        for fi34 := 0; fi34 < 17; fi34++ {
            var v35 bool = (fi34 >= fi34)
            for fi36 := 0; fi36 < 14; fi36++ {
                var v37 uint = 791
                if v35 {
                    v38 := GenS1{F0: 426, F1: v4}
                    v3 = v3
                    v39 := GenS3{F0: fi36, F1: 1, F2: 100}
                    _ = v38
                    _ = v39
                }
                _ = fi36
                _ = v37
            }
            for fi40 := 0; fi40 < 15; fi40++ {
                switch (fi40 >> ((18446744073709551615) & 63)) {
                case 32767:
                    v2 = -128
                case 42:
                    if (fi34 >= (fi34 & (fi34 | fi34))) {
                        v32 = append(v32, (-(fi34)))
                        var v41 int64 = 252
                        v2 -= v2
                        _ = v41
                    }
                default:
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v3)
                    case "a":
                        v4 = genFn0()
                    }
                }
                _ = fi40
            }
            _ = fi34
            _ = v35
        }
    }
    var v42 uint64 = 18446744073709551615
    v2 += (v2 % ((v2) | 1))
    v4 += (v4 << ((0) & 63))
    var v43 float32 = (((-2.25 - 0.1) + 3.14159) * 0.0)
    var v44 uint64 = v42
    v45 := GenS4{F0: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
