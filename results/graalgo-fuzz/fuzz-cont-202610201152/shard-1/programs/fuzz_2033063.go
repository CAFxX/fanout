package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint }
type GenS2 struct { F0 int64; F1 int; F2 bool }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    v4 := []int{716, -65}
    var v5 int64 = -318
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int16 = -689
        var v8 float64 = -2.25
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 bool = false
    var v10 uint = 220
    if (2147483647 < 2) {
        v9 = v9
        v9 = false
    }
    fmt.Println(v5)
    v11 := GenS1{F0: 127, F1: (v5 - v5), F2: (v10 - (v10 / ((v10) | 1)))}
    switch -1 {
    case 1:
        fmt.Println(v10)
    case 677:
        fmt.Println(v9, v5, v5)
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 int = 4294967295
        for fi14 := 0; fi14 < 13; fi14++ {
            fmt.Println(v5, v5, v9)
            _ = len(v4)
            switch (("" + "") + ("" + "hello")) {
            case "a":
                v9 = true
            case "":
                delete(v3, "foo bar")
            case "hello":
                var v15 uint = (v10 >> ((42) & 63))
                _ = v15
            default:
                var v16 float64 = 0.1
                _ = v16
            }
            v9 = v9
            _ = fi14
        }
        switch "a" {
        case "foo bar":
            v9 = (-1 > 211)
        case "":
            var v17 int16 = 1000
            _ = v17
        case "\n\t":
            var v18 uint64 = 0
            _ = v18
        default:
            switch "" {
            case "a":
                v4 = append(v4, 42)
            case "foo bar":
                v13 -= (fi12 >> (((v10 + v10)) & 63))
            default:
                v1["foo bar"] = -539
            }
        }
        _ = fi12
        _ = v13
    }
    v10 -= 2
    var v19 uint64 = 796
    for fi20 := 0; fi20 < 11; fi20++ {
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 uint64 = (^(v19))
            v23 := GenS1{F0: -1, F1: -1, F2: v10}
            if v9 {
                v24 := GenS2{F0: -9223372036854775808, F1: fi21, F2: false}
                var v25 uint64 = 0
                var v26 int64 = v5
                _ = v24
                _ = v25
                _ = v26
            }
            _ = fi21
            _ = v22
            _ = v23
        }
        v9 = ((61 == 152) || false)
        switch 457 {
        case -116:
            for fi27 := 0; fi27 < 3; fi27++ {
                if v9 {
                    var v28 uint16 = 255
                    _ = v28
                } else {
                    v29 := GenS3{F0: "\n\t"}
                    _ = len(v4)
                    _ = v29
                }
                v30 := GenS1{F0: 100, F1: (v5 + -205), F2: uint(v19)}
                _ = fi27
                _ = v30
            }
        case 174:
            v31 := GenS2{F0: 117, F1: (fi20 >> ((v10) & 63)), F2: (false || (1e10 == 1.5))}
            _ = v31
        default:
            var v32 string = ""
            _ = v32
        }
        _ = fi20
    }
    var v33 int64 = (v5 << ((v10) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v9
}
