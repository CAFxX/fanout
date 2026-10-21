package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int }
type GenS2 struct { F0 int64; F1 int64; F2 uint }
type GenS3 struct { F0 string }
func genFn0(p0 int) int32 {
    for fi7 := 0; fi7 < 1; fi7++ {
        switch (^(p0)) {
        case 2147483647:
            v8 := []int{-445, (fi7 | p0), (p0 << ((255) & 63)), (p0 + fi7)}
            _ = v8
        default:
            var v9 int8 = 127
            _ = v9
        }
        v10 := GenS1{F0: (^(fi7)), F1: -959, F2: p0}
        p0 -= (fi7 / ((fi7) | 1))
        _ = fi7
        _ = v10
    }
    p0 = (p0 << ((18446744073709551615) & 63))
    v11 := GenS1{F0: p0, F1: 0, F2: (p0 - 2147483647)}
    v12 := GenS1{F0: 847, F1: 9223372036854775807, F2: p0}
    _ = p0
    _ = v11
    _ = v12
    return 1
}
func genFn1(p0 int8, p1 int8) int32 {
    switch (("hello" + "foo bar") + ((("a" + ("" + "a")) + "") + "\n\t")) {
    case "\n\t":
        var v13 int16 = 0
        _ = v13
    default:
        fmt.Println(p0, p0)
    }
    switch "\n\t" {
    case "\n\t":
        v14 := []int{787, 2147483647}
        _ = v14
    default:
        v15 := []int{32768, -775, 127}
        _ = v15
    }
    v16 := GenS1{F0: 426, F1: -1, F2: -204}
    p0 = (p0 | -128)
    _ = p0
    _ = p1
    _ = v16
    return 146
}

func main() {
    v1 := []int{42, 32768}
    v2 := make(map[uint32]int)
    v3 := make(map[string]bool)
    var v4 uint32 = 42
    var v5 int8 = -48
    v6 := []string{"a"}
    delete(v2, v4)
    _ = len(v6)
    switch "" {
    case "":
        if true {
            var v17 string = "a"
            var v18 int16 = 954
            var v19 int64 = 0
            _ = v17
            _ = v18
            _ = v19
        }
    default:
        if true {
            v20 := GenS2{F0: -9223372036854775808, F1: 42, F2: 42}
            _ = v20
        }
    }
    var v21 int32 = 42
    v4 = v4
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 uint16 = 577
        if false {
            fmt.Println(fi22, v21, fi22)
            if false {
                if ((v21 >> ((0) & 63)) == genFn0((fi22 / (((fi22 * 0)) | 1)))) {
                    var v24 int = (-(fi22))
                    _ = v24
                } else {
                    var v25 uint8 = 128
                    _ = v25
                }
                v26 := GenS3{F0: "a"}
                var v27 float64 = float64(0.1)
                _ = v26
                _ = v27
            }
            v28 := GenS3{F0: ""}
            _ = v28
        } else {
            for fi29 := 0; fi29 < 12; fi29++ {
                var v30 int = (fi22 - (fi29 & fi22))
                switch (fi29 >> ((4294967295) & 63)) {
                case -1:
                    var v31 float32 = 0.0
                    _ = v31
                case 188:
                    if len(v1) > 0 {
                        v1[0] = 137
                    }
                case 42:
                    switch "\n\t" {
                    case "a":
                        v32 := GenS1{F0: fi22, F1: 42, F2: 32768}
                        _ = v32
                    default:
                        delete(v3, "\n\t")
                    }
                }
                _ = len(v6)
                _ = fi29
                _ = v30
            }
        }
        _ = fi22
        _ = v23
    }
    fmt.Println(v4, v5)
    var v33 uint = 18446744073709551615
    if true {
        var v34 int8 = (v5 | v5)
        var v35 uint16 = 1
        _ = v34
        _ = v35
    }
    v4 += (v4 | 141)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
