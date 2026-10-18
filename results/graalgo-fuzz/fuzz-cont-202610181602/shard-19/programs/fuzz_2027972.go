package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint16; F2 int64 }
type GenS3 struct { F0 int8; F1 uint; F2 uint }
func genFn0(p0 uint, p1 int16, p2 int16) int16 {
    var v6 uint = p0
    var v7 int16 = -1
    switch ("" + "\n\t") {
    case "foo bar":
        var v8 bool = true
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 int = 2147483647
    v2 := []string{"a", "a", "hello"}
    var v3 float32 = 1e10
    var v4 uint64 = 0
    var v5 bool = (!true)
    _ = len(v2)
    for fi9 := 0; fi9 < 12; fi9++ {
        v3 = v3
        for fi10 := 0; fi10 < 20; fi10++ {
            v3 = float32((((((1e10 * 1e10) + (((1e10 - 1.5) * ((((3.14159 * ((((1.5 * ((((1e10 + float64((1e10 / float32(3.14159)))) + 0.1) / 1.5) + (0.0 + 1.5))) / 0.1) / float64(v3)) / 1.5)) + (-2.25 / float64(float32(1e10)))) * 1e10) / 0.1)) + -2.25)) / 1.5) * 1.5) - ((1e10 - float64((float32((0.1 / 1.5)) - v3))) / 1.0)))
            _ = fi10
        }
        _ = fi9
    }
    var v11 uint32 = 1
    for fi12 := 0; fi12 < 13; fi12++ {
        fmt.Println(v3)
        fmt.Println(v1)
        if ("\n\t" != (((("\n\t" + "hello") + "hello") + "a") + "\n\t")) {
            var v13 uint = 0
            v1 = (v1 | (v1 * (fi12 / ((fi12) | 1))))
            _ = v13
        } else {
            for fi14 := 0; fi14 < 9; fi14++ {
                var v15 int = fi14
                if v5 {
                    v16 := GenS1{F0: 42}
                    v17 := GenS2{F0: "hello", F1: 503, F2: -104}
                    var v18 uint8 = 1
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    var v19 string = "foo bar"
                    _ = v19
                }
                _ = fi14
                _ = v15
            }
            if v5 {
                for fi20 := 0; fi20 < 8; fi20++ {
                    fmt.Println(fi20, v5)
                    v2 = append(v2, "\n\t")
                    v21 := GenS1{F0: 88}
                    _ = fi20
                    _ = v21
                }
            } else {
                switch "foo bar" {
                case "hello":
                    switch (v1 % ((1) | 1)) {
                    case 0:
                        for fi22 := 0; fi22 < 9; fi22++ {
                            fmt.Println(fi12, v11, fi12)
                            v1 = (v1 << ((255) & 63))
                            _ = fi22
                        }
                    case 42:
                        switch 0 {
                        case 2147483647:
                            var v23 int32 = -507
                            _ = v23
                        case 127:
                            fmt.Println(v4, v3, v4)
                        case 128:
                            fmt.Println(v1, fi12, v11)
                        }
                    case 128:
                        for fi24 := 0; fi24 < 15; fi24++ {
                            v11 -= 2147483647
                            var v25 int16 = genFn0(65535, -32768, 0)
                            _ = fi24
                            _ = v25
                        }
                    default:
                        if (42 < 42) {
                            var v26 int64 = -10
                            _ = v26
                        }
                    }
                }
                switch "\n\t" {
                case "a":
                    switch "\n\t" {
                    case "":
                        var v27 uint64 = v4
                        _ = v27
                    case "\n\t":
                        var v28 uint8 = 127
                        _ = v28
                    case "a":
                        switch "a" {
                        case "":
                            v29 := GenS3{F0: 42, F1: 255, F2: 255}
                            _ = v29
                        case "foo bar":
                            v4 += (v4 | v4)
                        }
                    default:
                        switch ("" + "\n\t") {
                        case "":
                            v30 := GenS2{F0: "hello", F1: 255, F2: 606}
                            _ = v30
                        case "a":
                            if v5 {
                                var v31 int64 = 357
                                _ = v31
                            }
                        default:
                            if v5 {
                                v5 = v5
                                _ = len(v2)
                            } else {
                                if (!true) {
                                    fmt.Println(v3, v4)
                                    v32 := GenS2{F0: "\n\t", F1: 140, F2: 0}
                                    _ = v32
                                }
                                v33 := GenS3{F0: -98, F1: 66, F2: 1}
                                _ = v33
                            }
                        }
                    }
                case "hello":
                    switch (fi12 << ((255) & 63)) {
                    case 127:
                        switch v1 {
                        case 42:
                            fmt.Println(v1)
                        default:
                            fmt.Println(v3, fi12, fi12)
                        }
                    }
                default:
                    v2 = append(v2, "")
                }
                v3 = ((v3 * (((v3 / float32(float64(v3))) / v3) * v3)) - -2.25)
            }
        }
        _ = fi12
    }
    fmt.Println(v1, v1)
    switch v1 {
    case 127:
        v4 += (v4 + (v4 - 4294967295))
    default:
        var v34 uint = 18446744073709551615
        _ = v34
    }
    v2 = append(v2, "hello")
    fmt.Println(v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
