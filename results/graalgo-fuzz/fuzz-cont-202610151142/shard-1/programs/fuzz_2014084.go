package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int16) int16 {
    if (846 >= 42) {
        v6 := GenS1{F0: 9223372036854775807, F1: 18446744073709551615, F2: -128}
        _ = v6
    }
    p0 -= p0
    switch "" {
    case "hello":
        for fi7 := 0; fi7 < 5; fi7++ {
            switch "a" {
            case "hello":
                for fi8 := 0; fi8 < 11; fi8++ {
                    v9 := GenS1{F0: -427, F1: 0, F2: 127}
                    var v10 string = "foo bar"
                    var v11 bool = true
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
            case "foo bar":
                p0 = (^(p0))
            case "a":
                fmt.Println(fi7, fi7)
            default:
                var v12 uint32 = 778
                _ = v12
            }
            _ = fi7
        }
    }
    _ = p0
    return p0
}
func genFn1(p0 int16, p1 int, p2 uint) string {
    fmt.Println(p2, p2, p0)
    if (-1 < -128) {
        p0 -= p0
    } else {
        switch "hello" {
        case "foo bar":
            v13 := []int{}
            _ = v13
        case "\n\t":
            p2 = 0
        case "a":
            if false {
                fmt.Println(p0, p2)
                fmt.Println(p1, p1)
            }
        }
        v14 := []int{(p1 ^ (p1 >> (((p2 / ((p2) | 1))) & 63)))}
        _ = v14
    }
    var v15 int64 = 0
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{("" + ""), "a", ("foo bar" + ("hello" + "foo bar")), "\n\t"}
    var v3 float64 = 0.0
    v4 := []string{((("foo bar" + "\n\t") + "") + ""), "foo bar"}
    v5 := make(map[int]bool)
    v4 = append(v4, "foo bar")
    for fi16 := 0; fi16 < 2; fi16++ {
        v3 = 3.14159
        _ = fi16
    }
    var v17 int16 = -80
    v18 := GenS1{F0: 138, F1: 1, F2: 42}
    fmt.Println(v17, v3)
    fmt.Println(v17, v17)
    v19 := GenS1{F0: -9223372036854775808, F1: 1, F2: 38}
    var v20 uint64 = 615
    var v21 int16 = (v17 >> ((42) & 63))
    v5[1] = (uint16(v20) > 444)
    for fi22 := 0; fi22 < 14; fi22++ {
        switch (fi22 + fi22) {
        case 32767:
            fmt.Println(fi22, v21, v17)
        case 32768:
            var v23 int64 = -9223372036854775808
            _ = v23
        }
        _ = fi22
    }
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    var v24 int16 = v21
    var v25 uint16 = 0
    for fi26 := 0; fi26 < 8; fi26++ {
        v27 := GenS2{F0: 2, F1: 18446744073709551615}
        _ = fi26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
