package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int64 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int8; F1 int }
func genFn0() uint64 {
    var v4 uint32 = 4294967295
    switch ("a" + ("foo bar" + "hello")) {
    case "a":
        fmt.Println(v4)
    case "\n\t":
        v5 := []int{2147483647, 128, 4294967295}
        _ = v5
    }
    v4 += v4
    v4 = 4294967295
    switch "a" {
    case "a":
        v4 = v4
    case "":
        v6 := []int{2147483647, 403, -545, 413}
        _ = v6
    }
    _ = v4
    return 1
}
func genFn1(p0 uint16, p1 int, p2 uint) uint16 {
    v7 := GenS1{F0: (^(p1)), F1: -32768, F2: -1}
    var v8 uint32 = 54
    fmt.Println(p1, p0)
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            if false {
                var v11 float64 = (float64((-2.25 * (((0.0 - 1.5) / 1.0) * 0.1))) + (-2.25 / -2.25))
                v12 := GenS1{F0: fi10, F1: -177, F2: 0}
                _ = v11
                _ = v12
            }
            _ = fi10
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (^(p0))
}

func main() {
    v1 := []string{("\n\t" + "\n\t"), ("a" + "a"), "foo bar"}
    v2 := []uint32{901, 65535}
    v3 := make(map[int]int)
    v3[-1] = -9223372036854775808
    if true {
        if (false || false) {
            var v13 string = ""
            fmt.Println(v13)
            _ = v13
        }
        var v14 uint = 65535
        fmt.Println(v14)
        _ = v14
    }
    var v15 int32 = -766
    v15 = (v15 - v15)
    for fi16 := 0; fi16 < 2; fi16++ {
        if true {
            switch (fi16 / ((42) | 1)) {
            case 127:
                if (524 >= 75) {
                    fmt.Println(fi16, fi16)
                    v17 := GenS1{F0: -925, F1: 521, F2: 1}
                    if (0 >= -32768) {
                        v18 := GenS1{F0: 42, F1: -32768, F2: 9223372036854775807}
                        _ = v18
                    }
                    _ = v17
                } else {
                    v15 = v15
                    v15 -= v15
                    fmt.Println(v15)
                }
            case 2147483647:
                v19 := GenS1{F0: fi16, F1: 32767, F2: 9223372036854775807}
                _ = v19
            }
            var v20 uint = 2
            var v21 uint8 = 128
            _ = v20
            _ = v21
        } else {
            for fi22 := 0; fi22 < 13; fi22++ {
                v15 -= 42
                if true {
                    fmt.Println(fi16, fi22)
                }
                v15 = (-(v15))
                _ = fi22
            }
            v15 = v15
            v15 -= (v15 ^ v15)
        }
        var v23 bool = (0 == 621)
        v24 := GenS2{F0: 18446744073709551615, F1: v23}
        _ = fi16
        _ = v23
        _ = v24
    }
    _ = len(v3)
    switch "a" {
    case "foo bar":
        v25 := GenS2{F0: 2, F1: true}
        _ = v25
    case "\n\t":
        var v26 uint64 = 2
        _ = v26
    default:
        v27 := v3[32767]
        _ = v27
    }
    v1 = append(v1, ("a" + "\n\t"))
    v28 := GenS3{F0: 1, F1: 127}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v3
}
