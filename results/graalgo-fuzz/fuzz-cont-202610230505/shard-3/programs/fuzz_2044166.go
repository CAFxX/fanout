package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int }
type GenS2 struct { F0 uint32; F1 int8; F2 int64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int32, p1 int64) string {
    for fi4 := 0; fi4 < 6; fi4++ {
        if (158 <= 1) {
            fmt.Println(fi4, p1, fi4)
            var v5 int8 = 0
            _ = v5
        }
        switch "foo bar" {
        case "\n\t":
            v6 := []int{(fi4 >> ((144) & 63)), (fi4 >> ((2) & 63)), fi4}
            _ = v6
        case "foo bar":
            var v7 uint16 = 666
            _ = v7
        }
        _ = fi4
    }
    var v8 uint64 = 123
    p0 = p0
    fmt.Println(v8)
    _ = p0
    _ = p1
    _ = v8
    return "hello"
}
func genFn1(p0 string, p1 int, p2 uint16) int64 {
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: p0, F1: (p0 + "\n\t"), F2: (p1 * fi9)}
        fmt.Println(p2)
        _ = fi9
        _ = v10
    }
    if false {
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := []int{}
            if (93 != -2147483648) {
                for fi13 := 0; fi13 < 14; fi13++ {
                    v14 := GenS1{F0: p0, F1: ("hello" + p0), F2: (fi11 / (((-(p1))) | 1))}
                    var v15 int8 = -128
                    v16 := GenS2{F0: 2, F1: v15, F2: 0}
                    if false {
                        p2 = (p2 & p2)
                        var v17 uint64 = 4294967295
                        _ = v17
                    }
                    if (!false) {
                        v18 := GenS2{F0: 42, F1: (^(v15)), F2: -884}
                        _ = v18
                    } else {
                        var v19 int32 = 320
                        v15 -= -128
                        _ = v19
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
            }
            _ = fi11
            _ = v12
        }
        p1 += 9223372036854775807
        p0 = "hello"
    } else {
        var v20 uint8 = 42
        fmt.Println(p0, p2)
        var v21 int8 = -1
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    return 9223372036854775807
}

func main() {
    v1 := []uint32{2, 1, 42}
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v22 uint32 = 612
    var v23 int32 = 207
    v24 := GenS2{F0: v22, F1: 42, F2: -1}
    var v25 int64 = 0
    if (1 != 1) {
        v23 = 0
    }
    v26 := GenS2{F0: 247, F1: 16, F2: genFn1("a", 4294967295, 1)}
    if ((1 < 25) || false) {
        var v27 uint8 = 114
        fmt.Println(v23)
        _ = v27
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        switch (fi28 + (fi28 & (fi28 + (fi28 * (fi28 / (((fi28 % ((fi28) | 1))) | 1)))))) {
        case -398:
            switch ("a" + "foo bar") {
            case "foo bar":
                fmt.Println(v22)
            case "":
                if true {
                    if false {
                        v22 = v22
                    }
                    switch (fi28 | fi28) {
                    case 9223372036854775807:
                        for fi29 := 0; fi29 < 17; fi29++ {
                            var v30 uint32 = 2147483647
                            v31 := GenS2{F0: (v30 * v22), F1: -1, F2: (v25 * (-(v25)))}
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    case 342:
                        if (!(42 == (v22 + (v22 * v22)))) {
                            var v32 bool = (0.1 == (1.5 - (-2.25 - -2.25)))
                            _ = v32
                        } else {
                            var v33 uint32 = v22
                            _ = v33
                        }
                    }
                }
            default:
                v34 := GenS2{F0: (v22 & v22), F1: 98, F2: 9223372036854775807}
                _ = v34
            }
        }
        _ = fi28
    }
    fmt.Println(v25)
    v25 = (v25 ^ v25)
    v35 := GenS3{F0: 1}
    var v36 int64 = (-(v25))
    v22 -= (v22 << ((4294967295) & 63))
    v23 = (-(v23))
    var v37 uint32 = v22
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v37
}
