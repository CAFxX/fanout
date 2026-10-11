package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = 1e10
    var v3 float32 = 1e10
    var v4 float64 = 0.0
    var v5 int8 = 0
    var v6 bool = (!true)
    var v7 int16 = 1000
    var v8 float32 = v2
    v9 := []int{9223372036854775807, 656, 1, 32767}
    var v10 int = int(v5)
    for fi11 := 0; fi11 < 16; fi11++ {
        switch (fi11 >> ((1) & 63)) {
        case 796:
            v8 -= (((((3.14159 / 0.1) - v3) / v8) * (v2 * 1e10)) - 0.1)
        }
        v6 = (871 > (v7 / (((v7 / ((v7) | 1))) | 1)))
        if (true || v6) {
            v3 = (float32(float64((v3 * 0.0))) - v8)
            var v12 float64 = (float64(1e10) * v4)
            _ = v12
        }
        _ = fi11
    }
    v13 := GenS1{F0: 42, F1: 574, F2: -1}
    switch "\n\t" {
    case "foo bar":
        v2 = 1.5
    case "a":
        fmt.Println(v7)
    case "":
        fmt.Println(v3, v6, v10)
    default:
        v14 := GenS2{F0: 95}
        _ = v14
    }
    if v6 {
        var v15 int32 = 1
        switch ("" + "hello") {
        case "":
            v16 := GenS1{F0: 0, F1: 2, F2: v10}
            _ = v16
        case "a":
            v3 -= (v3 - v2)
        case "hello":
            fmt.Println(v6)
        }
        switch "hello" {
        case "a":
            var v17 uint32 = 186
            _ = v17
        case "hello":
            v10 -= (v10 & v10)
        default:
            if v6 {
                switch (v10 + v10) {
                case 9223372036854775807:
                    switch ((("" + ("\n\t" + "\n\t")) + ("" + "a")) + "") {
                    case "a":
                        if len(v9) > 0 {
                            v9[0] = v10
                        }
                    case "":
                        var v18 uint64 = 33
                        _ = v18
                    }
                }
                _ = len(v9)
            }
        }
        _ = v15
    } else {
        if true {
            var v19 uint32 = 406
            if (!(v8 == v2)) {
                switch "foo bar" {
                case "hello":
                    v7 += v7
                }
                if ((42 >= 42) || true) {
                    switch -1 {
                    case 790:
                        _ = len(v1)
                    default:
                        for fi20 := 0; fi20 < 2; fi20++ {
                            var v21 int16 = (^(v7))
                            _ = fi20
                            _ = v21
                        }
                    }
                } else {
                    for fi22 := 0; fi22 < 12; fi22++ {
                        var v23 string = "a"
                        _ = len(v1)
                        var v24 uint32 = 42
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    var v25 uint32 = v19
                    delete(v1, (^(v10)))
                    _ = v25
                }
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 int8 = 1
                    v19 = 2147483647
                    _ = fi26
                    _ = v27
                }
            }
            var v28 int64 = 9223372036854775807
            _ = v19
            _ = v28
        }
        if v6 {
            var v29 uint32 = 2147483648
            _ = v29
        }
        var v30 int16 = v7
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
