package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0(p0 string, p1 int, p2 uint32) uint64 {
    switch p1 {
    case 0:
        fmt.Println(p0, p1)
    case 9223372036854775807:
        fmt.Println(p0, p2, p2)
    }
    v4 := []int{p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 18446744073709551615
}

func main() {
    var v1 string = ("" + ((("" + "") + "foo bar") + "hello"))
    v2 := make(map[string]string)
    v3 := make(map[uint32]uint32)
    var v5 int = 128
    for fi6 := 0; fi6 < 1; fi6++ {
        switch fi6 {
        case 32768:
            v7 := GenS1{F0: genFn0("a", fi6, 42), F1: ("a" + "foo bar"), F2: 2}
            _ = v7
        case 127:
            fmt.Println(v5)
        case -706:
            v5 = (v5 | (v5 * fi6))
        default:
            fmt.Println(v5, fi6, v5)
        }
        var v8 int8 = 1
        _ = fi6
        _ = v8
    }
    switch 372 {
    case 128:
        var v9 int64 = -1
        _ = v9
    case 42:
        switch v1 {
        case "\n\t":
            fmt.Println(v5)
        }
    default:
        v10 := []int{(v5 / ((v5) | 1)), v5}
        _ = v10
    }
    v11 := GenS2{F0: 127}
    var v12 int = (^(v5))
    v3[1] = 2147483648
    var v13 uint16 = 255
    v5 += 1
    if (4294967295 != 42) {
        var v14 int = -9223372036854775808
        _ = v14
    }
    v15 := []int{(v12 >> ((861) & 63)), 32767, (v5 ^ (v5 - (v12 / ((v5) | 1))))}
    var v16 uint = 1
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 int = (^(v5))
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        v20 := GenS2{F0: 42}
        for fi21 := 0; fi21 < 7; fi21++ {
            v1 = v1
            _ = fi21
        }
        if true {
            v5 += (fi19 ^ v5)
            if ((!true) || (355 >= 0)) {
                v12 -= v5
                if false {
                    var v22 bool = (1000 >= -730)
                    var v23 uint64 = 1
                    _ = v22
                    _ = v23
                }
            } else {
                if ((!(-1 != 1)) && true) {
                    for fi24 := 0; fi24 < 4; fi24++ {
                        if (!(!(true && (true || false)))) {
                            _ = len(v15)
                        }
                        switch (fi24 & fi19) {
                        case 2147483647:
                            v12 += fi19
                            v25 := GenS1{F0: 443, F1: v1, F2: 391}
                            if false {
                                _ = len(v3)
                                fmt.Println(fi24)
                                switch v1 {
                                case "a":
                                    fmt.Println(fi24, v5, fi24)
                                    v26 := GenS3{F0: 0, F1: 42}
                                    _ = v26
                                case "":
                                    var v27 int16 = -612
                                    _ = v27
                                case "foo bar":
                                    _ = len(v2)
                                }
                            } else {
                                var v28 bool = false
                                switch "foo bar" {
                                case "foo bar":
                                    v28 = v28
                                }
                                _ = v28
                            }
                            var v29 uint16 = 596
                            _ = v25
                            _ = v29
                        case 42:
                            v12 -= (fi24 << ((v16) & 63))
                        }
                        _ = fi24
                    }
                    fmt.Println(v13, v5)
                }
            }
        } else {
            var v30 float64 = (-2.25 + (1.5 * ((3.14159 / -2.25) - 1e10)))
            v12 = 127
            _ = v30
        }
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v5
}
