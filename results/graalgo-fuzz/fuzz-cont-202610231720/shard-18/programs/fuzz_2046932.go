package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []int{32768, 42, 128}
    v2 := []int{}
    v3 := make(map[uint32]string)
    v4 := []string{("" + ("a" + "a")), "foo bar", "hello", "a"}
    v5 := make(map[string]bool)
    var v6 float64 = (0.0 / 1.5)
    var v7 int16 = 1000
    var v8 int16 = 42
    var v9 uint64 = 4294967295
    var v10 int32 = 2147483647
    if false {
        fmt.Println(v7, v7)
        v11 := GenS1{F0: 127}
        _ = v11
    } else {
        var v12 int64 = 0
        if true {
            if true {
                var v13 int64 = 1
                fmt.Println(v9, v7, v7)
                for fi14 := 0; fi14 < 15; fi14++ {
                    if (!(255 != 255)) {
                        switch "hello" {
                        case "\n\t":
                            v9 = v9
                        default:
                            v13 = 1
                        }
                        switch "a" {
                        case "hello":
                            v10 = v10
                        case "foo bar":
                            for fi15 := 0; fi15 < 12; fi15++ {
                                v8 = 42
                                if true {
                                    v13 -= (v12 * v12)
                                    if false {
                                        var v16 float64 = ((-2.25 * 1.5) * ((float64(1.5) - (v6 - (1e10 + v6))) * v6))
                                        _ = v16
                                    }
                                    var v17 uint8 = 2
                                    var v18 float32 = (0.1 * 1.5)
                                    _ = v17
                                    _ = v18
                                }
                                var v19 uint64 = v9
                                _ = fi15
                                _ = v19
                            }
                        case "a":
                            for fi20 := 0; fi20 < 16; fi20++ {
                                v12 -= v13
                                v13 = (v13 + (-(v13)))
                                if false {
                                    var v21 uint16 = 2
                                    var v22 uint16 = (v21 >> ((899) & 63))
                                    var v23 string = ("" + "hello")
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                                _ = fi20
                            }
                        default:
                            var v24 int = fi14
                            _ = v24
                        }
                    } else {
                        fmt.Println(v10, v12, v6)
                    }
                    v13 += (v13 - (v12 + v12))
                    _ = fi14
                }
                _ = v13
            }
        }
        var v25 uint64 = (v9 / ((2) | 1))
        _ = v12
        _ = v25
    }
    if true {
        fmt.Println(v8)
        if true {
            v26 := GenS1{F0: 128}
            for fi27 := 0; fi27 < 17; fi27++ {
                v7 -= v8
                var v28 int16 = -1
                v6 -= float64(1.5)
                _ = fi27
                _ = v28
            }
            v29 := GenS1{F0: 2}
            _ = v26
            _ = v29
        }
        var v30 uint32 = 404
        _ = v30
    }
    var v31 uint8 = 128
    var v32 int64 = 9223372036854775807
    switch 128 {
    case -577:
        fmt.Println(v32, v6, v32)
    default:
        v10 = (v10 + (v10 + (v10 & v10)))
    }
    var v33 float64 = ((float64(0.0) + v6) + 0.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
