package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint, p1 bool) uint {
    p0 -= p0
    if (p1 && (18446744073709551615 != 42)) {
        v5 := GenS1{F0: 2, F1: 665}
        _ = v5
    }
    _ = p0
    _ = p1
    return 255
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[string]bool)
    v3 := []uint32{718, 2147483648}
    v4 := []int{-1, -1, 1, 1}
    var v6 int = 9223372036854775807
    var v7 int8 = 100
    if (!true) {
        fmt.Println(v6)
        fmt.Println(v1, v1, v1)
        v1 = (^(v1))
    }
    v6 = v6
    v6 -= v6
    v1 = v1
    fmt.Println(v1, v6)
    for fi8 := 0; fi8 < 11; fi8++ {
        fmt.Println(v6, v1, v1)
        switch "foo bar" {
        case "a":
            fmt.Println(v7, v7)
        case "foo bar":
            v6 += 0
        case "hello":
            v9 := GenS1{F0: 42, F1: -591}
            _ = v9
        }
        switch v6 {
        case 9223372036854775807:
            if false {
                if false {
                    v7 = (v7 | 42)
                    v10 := GenS1{F0: 2147483647, F1: 42}
                    for fi11 := 0; fi11 < 8; fi11++ {
                        var v12 float32 = (1e10 * 0.0)
                        switch v6 {
                        case 35:
                            var v13 int16 = 979
                            _ = v13
                        case -941:
                            if true {
                                v14 := GenS1{F0: 42, F1: -1}
                                var v15 uint = genFn0(18446744073709551615, false)
                                _ = v14
                                _ = v15
                            } else {
                                switch "" {
                                case "":
                                    var v16 int32 = 2147483647
                                    _ = v16
                                case "hello":
                                    v17 := v2[("foo bar" + (("\n\t" + (("\n\t" + ((("foo bar" + "a") + ("\n\t" + "\n\t")) + "")) + ((("" + "hello") + "foo bar") + ("a" + "\n\t")))) + "a"))]
                                    _ = v17
                                case "foo bar":
                                    var v18 int = v6
                                    _ = v18
                                }
                                v19 := GenS1{F0: 3, F1: -9223372036854775808}
                                if false {
                                    v1 = (v1 << ((523) & 63))
                                }
                                _ = v19
                            }
                        case 42:
                            v20 := GenS2{F0: (v7 | -73)}
                            _ = v20
                        }
                        switch 0 {
                        case 0:
                            fmt.Println(v12, fi8, v7)
                        case -9223372036854775808:
                            v6 = v6
                        default:
                            v21 := GenS2{F0: (v7 ^ v7)}
                            _ = v21
                        }
                        switch "hello" {
                        case "foo bar":
                            v1 = 4294967295
                        default:
                            v22 := GenS2{F0: (v7 + (-(v7)))}
                            _ = v22
                        }
                        _ = fi11
                        _ = v12
                    }
                    _ = v10
                }
                _ = len(v3)
                var v23 bool = false
                _ = v23
            }
        case 32768:
            var v24 int8 = 127
            _ = v24
        }
        _ = fi8
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
