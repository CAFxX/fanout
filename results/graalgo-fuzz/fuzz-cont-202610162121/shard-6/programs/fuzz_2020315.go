package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint64 }
func genFn0() bool {
    if false {
        var v7 int16 = 0
        for fi8 := 0; fi8 < 12; fi8++ {
            v7 += (v7 % ((v7) | 1))
            for fi9 := 0; fi9 < 20; fi9++ {
                fmt.Println(fi8, fi8)
                switch ("" + "foo bar") {
                case "a":
                    switch "" {
                    case "hello":
                        switch ("a" + "foo bar") {
                        case "":
                            v10 := []int{615, (^(fi8)), fi9}
                            _ = v10
                        case "\n\t":
                            var v11 float32 = (0.0 + (3.14159 - 0.0))
                            if (false && (!true)) {
                                var v12 string = "foo bar"
                                _ = v12
                            }
                            fmt.Println(v7, v7, v11)
                            _ = v11
                        }
                    case "a":
                        if true {
                            v7 += (v7 ^ (-(v7)))
                            v13 := GenS1{F0: 2, F1: true, F2: 4294967295}
                            var v14 uint = 2
                            v15 := []int{(fi8 | fi8)}
                            v7 = (v7 + v7)
                            _ = v13
                            _ = v14
                            _ = v15
                        } else {
                            v7 = 568
                            var v16 uint32 = 956
                            var v17 uint32 = 42
                            _ = v16
                            _ = v17
                        }
                    case "\n\t":
                        v7 = v7
                    }
                }
                switch (fi8 | (fi8 / (((fi8 % (((fi8 / ((fi9) | 1))) | 1))) | 1))) {
                case -2147483648:
                    var v18 uint64 = 2
                    _ = v18
                }
                _ = fi9
            }
            var v19 uint = 601
            _ = fi8
            _ = v19
        }
        fmt.Println(v7)
        _ = v7
    } else {
        var v20 uint8 = 42
        _ = v20
    }
    var v21 uint32 = 2
    switch 9223372036854775807 {
    case 4294967295:
        v21 += v21
    case 734:
        fmt.Println(v21)
    default:
        if false {
            var v22 int = 4294967295
            _ = v22
        }
    }
    _ = v21
    return true
}
func genFn1(p0 int32, p1 int16) uint64 {
    if false {
        switch 4294967295 {
        case 127:
            switch 0 {
            case -2147483648:
                p0 -= -1
            case 128:
                fmt.Println(p0)
            }
        case 4294967295:
            fmt.Println(p0, p1)
        }
        var v23 string = "\n\t"
        _ = v23
    } else {
        v24 := []int{-2147483648}
        _ = v24
    }
    var v25 int16 = int16(p0)
    switch -658 {
    case -2147483648:
        v26 := []int{128, 9223372036854775807, -112, 128}
        _ = v26
    default:
        fmt.Println(p1, p1, p0)
    }
    _ = p0
    _ = p1
    _ = v25
    return 2
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    var v3 bool = true
    v4 := []string{}
    v5 := []uint32{}
    var v6 uint32 = 611
    var v27 float64 = ((1e10 * 0.1) + 3.14159)
    fmt.Println(v27, v3, v3)
    v6 += v6
    var v28 int32 = -1
    var v29 string = ("" + "hello")
    for fi30 := 0; fi30 < 6; fi30++ {
        switch -558 {
        case -152:
            v29 = "a"
        }
        var v31 uint32 = (v6 | v6)
        fmt.Println(v29, fi30, v28)
        _ = fi30
        _ = v31
    }
    v5 = append(v5, 2147483647)
    v32 := GenS1{F0: 0, F1: (-1 == 59), F2: 18446744073709551615}
    if len(v5) > 0 {
        v5[0] = v6
    }
    var v33 int8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
