package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int, p1 int32) string {
    v7 := []int{p0, p0}
    var v8 int32 = (p1 >> ((4294967295) & 63))
    if (215 >= 1) {
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 float64 = (0.0 - 3.14159)
            _ = fi9
            _ = v10
        }
    } else {
        var v11 uint = 1
        var v12 uint64 = 18446744073709551615
        for fi13 := 0; fi13 < 9; fi13++ {
            _ = len(v7)
            p1 = 724
            v12 = (v12 & 2)
            _ = fi13
        }
        _ = v11
        _ = v12
    }
    p1 = -242
    for fi14 := 0; fi14 < 12; fi14++ {
        if true {
            if false {
                switch "" {
                case "foo bar":
                    switch "" {
                    case "a":
                        var v15 bool = (!false)
                        _ = v15
                    case "hello":
                        fmt.Println(fi14)
                    case "\n\t":
                        fmt.Println(p1, p0)
                    }
                }
                for fi16 := 0; fi16 < 15; fi16++ {
                    switch "" {
                    case "foo bar":
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    default:
                        switch "foo bar" {
                        case "hello":
                            switch "\n\t" {
                            case "\n\t":
                                var v17 int = 9223372036854775807
                                _ = v17
                            }
                        default:
                            _ = len(v7)
                        }
                    }
                    var v18 string = "a"
                    _ = fi16
                    _ = v18
                }
                v19 := GenS1{F0: 65535, F1: 65535}
                _ = v19
            }
            switch 42 {
            case 387:
                var v20 uint = 18446744073709551615
                _ = v20
            default:
                if (!(!false)) {
                    v7 = append(v7, (fi14 >> ((1) & 63)))
                    p1 = -778
                }
            }
        } else {
            for fi21 := 0; fi21 < 15; fi21++ {
                p1 = (^(p1))
                v8 += 2147483647
                _ = fi21
            }
            for fi22 := 0; fi22 < 10; fi22++ {
                v7 = append(v7, 32768)
                p1 = p1
                _ = fi22
            }
            var v23 int = (fi14 + p0)
            _ = v23
        }
        var v24 float64 = float64(((1e10 / 3.14159) / 1.5))
        _ = fi14
        _ = v24
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return "\n\t"
}
func genFn1(p0 int16, p1 uint8) uint8 {
    v25 := []int{-631, 9223372036854775807}
    for fi26 := 0; fi26 < 11; fi26++ {
        p1 = p1
        var v27 float32 = 0.1
        var v28 uint16 = 42
        _ = fi26
        _ = v27
        _ = v28
    }
    _ = len(v25)
    var v29 uint16 = 0
    _ = p0
    _ = p1
    _ = v25
    _ = v29
    return 249
}

func main() {
    var v1 float32 = -2.25
    var v2 float32 = 1e10
    v3 := []int{127, 127}
    var v4 int8 = 42
    var v5 uint32 = 363
    var v6 bool = true
    v4 -= v4
    var v30 int = -9223372036854775808
    var v31 uint8 = 234
    if v6 {
        switch v30 {
        case 42:
            v32 := GenS1{F0: 65535, F1: 42}
            _ = v32
        default:
            if ((14 < genFn1(-32768, 1)) && ("\n\t" != "hello")) {
                v33 := GenS1{F0: 705, F1: 255}
                _ = len(v3)
                v1 += v2
                _ = v33
            }
        }
    } else {
        var v34 int64 = 128
        v4 = v4
        _ = v34
    }
    v6 = v6
    var v35 uint64 = 1
    fmt.Println(v30)
    var v36 int64 = -614
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
