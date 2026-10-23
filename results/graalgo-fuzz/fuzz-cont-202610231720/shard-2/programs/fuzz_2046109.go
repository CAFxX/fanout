package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint, p1 int64) uint32 {
    var v5 int8 = -126
    var v6 float64 = (1e10 + 0.1)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 484
}
func genFn1() bool {
    if false {
        switch "hello" {
        case "":
            v7 := GenS1{F0: 610, F1: 128, F2: 820}
            _ = v7
        case "\n\t":
            switch "foo bar" {
            case "hello":
                for fi8 := 0; fi8 < 14; fi8++ {
                    for fi9 := 0; fi9 < 15; fi9++ {
                        var v10 bool = ((false || false) && false)
                        fmt.Println(v10)
                        _ = fi9
                        _ = v10
                    }
                    if true {
                        switch "\n\t" {
                        case "hello":
                            if true {
                                fmt.Println(fi8)
                                var v11 uint = 1
                                _ = v11
                            }
                        case "a":
                            for fi12 := 0; fi12 < 14; fi12++ {
                                var v13 bool = (true || false)
                                _ = fi12
                                _ = v13
                            }
                        case "\n\t":
                            var v14 float64 = ((1e10 / 3.14159) / -2.25)
                            _ = v14
                        }
                    }
                    if true {
                        var v15 int16 = 1
                        if ("hello" <= "foo bar") {
                            v16 := []int{(fi8 + 128)}
                            switch "hello" {
                            case "foo bar":
                                v15 -= (v15 % ((v15) | 1))
                            case "\n\t":
                                v17 := GenS1{F0: 4294967295, F1: 0, F2: 2}
                                _ = v17
                            default:
                                v15 -= v15
                            }
                            _ = v16
                        }
                        _ = v15
                    } else {
                        v18 := []int{(fi8 >> ((365) & 63)), -191, 127, fi8}
                        fmt.Println(fi8)
                        _ = v18
                    }
                    _ = fi8
                }
            case "a":
                v19 := GenS2{F0: 2147483647, F1: 256}
                _ = v19
            case "":
                if (-32768 > 0) {
                    var v20 uint16 = 2
                    v21 := []int{32768, 32767, 0}
                    _ = v20
                    _ = v21
                }
            }
        case "hello":
            v22 := GenS3{F0: 879, F1: 42}
            _ = v22
        }
        var v23 float32 = (3.14159 - float32(1.5))
        _ = v23
    }
    var v24 uint8 = 150
    for fi25 := 0; fi25 < 20; fi25++ {
        v26 := []int{(fi25 + (fi25 & -381)), (fi25 * fi25), fi25}
        _ = fi25
        _ = v26
    }
    fmt.Println(v24)
    v27 := []int{42}
    _ = v24
    _ = v27
    return ((267 >= 837) || true)
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{2, 2, 42}
    v3 := []string{}
    v4 := make(map[int]uint32)
    var v28 string = "\n\t"
    var v29 int = 2147483647
    fmt.Println(v28, v28)
    v4[v29] = 770
    delete(v1, v29)
    if len(v3) > 0 {
        v3[0] = v28
    }
    var v30 float64 = 3.14159
    var v31 float32 = 0.0
    switch 399 {
    case -109:
        _ = len(v4)
    case 127:
        for fi32 := 0; fi32 < 2; fi32++ {
            v30 += v30
            if (true || false) {
                v1[(fi32 - v29)] = (!(false || (!true)))
                if true {
                    v28 = v28
                    v30 = ((0.0 / v30) * 1.5)
                    if true {
                        fmt.Println(v31)
                        var v33 int32 = 42
                        _ = v33
                    }
                }
            } else {
                v34 := GenS2{F0: -2147483648, F1: 1}
                v35 := GenS4{F0: 47}
                _ = v34
                _ = v35
            }
            var v36 int = (fi32 & 9223372036854775807)
            _ = fi32
            _ = v36
        }
    }
    switch ("" + v28) {
    case "\n\t":
        var v37 uint64 = 1
        _ = v37
    }
    v30 -= v30
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
