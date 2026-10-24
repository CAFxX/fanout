package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int16) int64 {
    v5 := []int{}
    var v6 float32 = -2.25
    for fi7 := 0; fi7 < 7; fi7++ {
        switch "a" {
        case "\n\t":
            fmt.Println(p0)
        case "foo bar":
            var v8 int16 = (^(p0))
            _ = v8
        case "a":
            switch "" {
            case "foo bar":
                v6 -= v6
            case "\n\t":
                var v9 int8 = -128
                _ = v9
            case "hello":
                fmt.Println(v6)
            default:
                if (!(1e10 > (v6 + ((((0.1 * v6) - 0.1) * (v6 + (v6 / v6))) / v6)))) {
                    v5 = append(v5, int(p0))
                    if (true || true) {
                        if (2147483648 == 4294967295) {
                            switch fi7 {
                            case -9223372036854775808:
                                for fi10 := 0; fi10 < 2; fi10++ {
                                    fmt.Println(p0, v6)
                                    _ = fi10
                                }
                            }
                            for fi11 := 0; fi11 < 5; fi11++ {
                                v6 -= (float32(1.5) * (v6 + v6))
                                _ = fi11
                            }
                        }
                        p0 = (p0 - p0)
                        v6 -= v6
                    }
                } else {
                    if false {
                        switch fi7 {
                        case 0:
                            var v12 int = fi7
                            _ = v12
                        case 32768:
                            for fi13 := 0; fi13 < 7; fi13++ {
                                v6 = v6
                                if ((int8(p0) <= 0) || true) {
                                    if false {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    var v14 int = int(p0)
                                    if len(v5) > 0 {
                                        v5[0] = (-(fi13))
                                    }
                                    _ = v14
                                } else {
                                    switch "foo bar" {
                                    case "foo bar":
                                        if (1e10 > 0.1) {
                                            v5 = append(v5, fi13)
                                        }
                                    case "hello":
                                        var v15 int16 = p0
                                        _ = v15
                                    case "":
                                        var v16 string = (("hello" + "hello") + "foo bar")
                                        _ = v16
                                    }
                                    v17 := GenS1{F0: 2147483647, F1: -911, F2: 127}
                                    v18 := GenS1{F0: 698, F1: 873, F2: 2}
                                    fmt.Println(fi7, v6)
                                    v19 := GenS1{F0: 2, F1: -9223372036854775808, F2: 255}
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                }
                                var v20 bool = (-2147483648 < 1)
                                v21 := GenS1{F0: 4294967295, F1: 807, F2: 128}
                                _ = fi13
                                _ = v20
                                _ = v21
                            }
                        case 1:
                            v22 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: 207}
                            _ = v22
                        default:
                            if true {
                                switch "foo bar" {
                                case "a":
                                    if true {
                                        if false {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                    }
                                default:
                                    if (-128 >= 107) {
                                        var v23 uint8 = 1
                                        fmt.Println(v23, v6)
                                        _ = v23
                                    } else {
                                        if false {
                                            v24 := GenS1{F0: 584, F1: 0, F2: 0}
                                            _ = v24
                                        }
                                    }
                                }
                            } else {
                                p0 -= p0
                                p0 += p0
                                v5 = append(v5, 42)
                            }
                        }
                        v25 := GenS2{F0: 1}
                        _ = v25
                    }
                }
            }
        default:
            var v26 string = "a"
            _ = v26
        }
        fmt.Println(p0)
        v6 = v6
        _ = fi7
    }
    _ = p0
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []uint32{568, 2}
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 831
    var v27 float32 = 0.0
    fmt.Println(v4, v27)
    var v28 float32 = v27
    v29 := v3[1]
    v1 -= (v1 * (v1 * (v1 * -937)))
    v4 -= (v4 >> ((255) & 63))
    fmt.Println(v4, v4, v1)
    fmt.Println(v27)
    switch (("\n\t" + ("\n\t" + "\n\t")) + "foo bar") {
    case "foo bar":
        v2 = append(v2, v29)
    case "":
        var v30 float64 = 0.1
        _ = v30
    }
    v31 := GenS2{F0: 2}
    v27 -= float32(float64(3.14159))
    if false {
        delete(v3, v29)
    }
    fmt.Println(v29, v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v27 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
}
