package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{4294967295}
    v3 := make(map[int]int)
    v4 := []int{}
    var v5 uint16 = 165
    var v6 float64 = 0.1
    var v7 int32 = 42
    for fi8 := 0; fi8 < 5; fi8++ {
        switch "a" {
        case "hello":
            for fi9 := 0; fi9 < 6; fi9++ {
                fmt.Println(v7, v6, fi8)
                _ = fi9
            }
        }
        switch "a" {
        case "a":
            v5 = v5
        }
        var v10 uint64 = 1
        _ = fi8
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        fmt.Println(fi11, v6, v5)
        _ = fi11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 float64 = (v6 * v6)
        switch 896 {
        case -1:
            fmt.Println(fi12, fi12)
        }
        var v14 uint16 = (^(v5))
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(v7, v6, v5)
    switch "hello" {
    case "\n\t":
        v6 -= 0.0
    case "a":
        v15 := v3[-52]
        _ = v15
    case "foo bar":
        switch 128 {
        case 898:
            if (false && ((false || (!true)) && (-9223372036854775808 == 9223372036854775807))) {
                if true {
                    var v16 float32 = (3.14159 + 1e10)
                    var v17 uint = 4294967295
                    _ = v16
                    _ = v17
                } else {
                    var v18 int32 = (v7 * 54)
                    delete(v3, 0)
                    _ = v18
                }
                var v19 float64 = 1e10
                _ = v19
            }
        case -1:
            v20 := GenS1{F0: 609, F1: (32767 < 42)}
            _ = v20
        case 9223372036854775807:
            for fi21 := 0; fi21 < 4; fi21++ {
                if (1.5 == (1.5 + -2.25)) {
                    v6 -= ((3.14159 / v6) / 1.0)
                    switch 32767 {
                    case -9223372036854775808:
                        v7 += (v7 ^ v7)
                    case 128:
                        v5 = v5
                    case 858:
                        var v22 bool = true
                        _ = v22
                    }
                    v3[fi21] = -80
                } else {
                    var v23 int8 = 79
                    _ = v23
                }
                var v24 int = fi21
                _ = len(v1)
                _ = fi21
                _ = v24
            }
        default:
            _ = len(v3)
        }
    }
    fmt.Println(v6, v7, v7)
    fmt.Println(v5, v6)
    switch "" {
    case "\n\t":
        if true {
            v25 := GenS1{F0: 64, F1: false}
            var v26 int64 = -1
            v26 = v26
            _ = v25
            _ = v26
        }
    case "a":
        if (true && false) {
            v6 = float64(1.5)
            for fi27 := 0; fi27 < 16; fi27++ {
                v28 := GenS1{F0: 0, F1: (2 > 65535)}
                for fi29 := 0; fi29 < 3; fi29++ {
                    var v30 uint = 57
                    _ = fi29
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
        }
    case "hello":
        switch "\n\t" {
        case "\n\t":
            _ = len(v4)
        case "foo bar":
            for fi31 := 0; fi31 < 9; fi31++ {
                fmt.Println(v6, v7)
                var v32 int64 = int64(v7)
                var v33 int8 = 0
                _ = fi31
                _ = v32
                _ = v33
            }
        case "hello":
            v7 = v7
        }
    }
    v6 = v6
    v6 = (float64((3.14159 - 0.1)) + v6)
    for fi34 := 0; fi34 < 7; fi34++ {
        fmt.Println(v7)
        _ = fi34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
