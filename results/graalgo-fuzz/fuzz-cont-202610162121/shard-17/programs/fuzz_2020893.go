package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int16; F2 int64 }
func genFn0(p0 int, p1 int64) int16 {
    var v4 int16 = 42
    for fi5 := 0; fi5 < 10; fi5++ {
        switch fi5 {
        case 2147483647:
            fmt.Println(v4, v4)
        case -983:
            v6 := GenS1{F0: 4294967295}
            _ = v6
        case -1:
            switch -694 {
            case 0:
                var v7 int8 = -128
                _ = v7
            case 4294967295:
                p0 = fi5
            default:
                var v8 uint = 4294967295
                _ = v8
            }
        default:
            for fi9 := 0; fi9 < 10; fi9++ {
                if true {
                    if (194 > 18446744073709551615) {
                        switch (fi9 - fi9) {
                        case 1:
                            if false {
                                switch ("a" + "foo bar") {
                                case "\n\t":
                                    var v10 int8 = 1
                                    _ = v10
                                case "hello":
                                    fmt.Println(p0, fi5)
                                case "":
                                    switch "" {
                                    case "\n\t":
                                        if ("hello" != ((("hello" + "\n\t") + "hello") + "foo bar")) {
                                            var v11 uint64 = 42
                                            p0 -= fi9
                                            switch 66 {
                                            case -790:
                                                v12 := []int{p0, fi5, p0, int(p1)}
                                                _ = v12
                                            default:
                                                p0 += 9223372036854775807
                                            }
                                            _ = v11
                                        }
                                    case "foo bar":
                                        var v13 float32 = 3.14159
                                        _ = v13
                                    case "":
                                        p1 -= (p1 - p1)
                                    }
                                default:
                                    v14 := []int{}
                                    _ = v14
                                }
                            } else {
                                v15 := GenS2{F0: 356, F1: v4, F2: p1}
                                switch (fi5 / ((fi9) | 1)) {
                                case 42:
                                    v16 := GenS2{F0: 1, F1: v4, F2: 1}
                                    _ = v16
                                default:
                                    var v17 int16 = v4
                                    _ = v17
                                }
                                _ = v15
                            }
                            v18 := []int{}
                            var v19 uint = 255
                            _ = v18
                            _ = v19
                        case -896:
                            switch fi5 {
                            case 4294967295:
                                if (2 < 255) {
                                    var v20 float32 = -2.25
                                    _ = v20
                                } else {
                                    var v21 float64 = ((0.0 * 0.1) / 1e10)
                                    var v22 uint64 = 374
                                    v23 := GenS2{F0: 1, F1: v4, F2: -324}
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                            }
                        case -51:
                            var v24 uint64 = 209
                            _ = v24
                        default:
                            var v25 uint16 = 42
                            _ = v25
                        }
                    } else {
                        fmt.Println(p0)
                    }
                } else {
                    p0 -= fi5
                    fmt.Println(v4, fi9, p1)
                    var v26 uint64 = 441
                    _ = v26
                }
                _ = fi9
            }
        }
        _ = fi5
    }
    v27 := []int{p0}
    var v28 uint32 = 65535
    p1 += p1
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    _ = v4
    return -335
}

func main() {
    var v1 int = 4294967295
    v2 := make(map[string]int)
    var v3 int64 = -1
    if false {
        _ = len(v2)
        v29 := []int{v1}
        switch "hello" {
        case "":
            v30 := GenS1{F0: 297}
            _ = v30
        case "\n\t":
            var v31 int16 = int16(v1)
            _ = v31
        }
        _ = v29
    }
    v32 := GenS1{F0: 4294967295}
    var v33 int16 = 32767
    v3 = v3
    if false {
        v33 -= (-(v33))
        v3 += v3
    }
    v1 = v1
    v1 = v1
    v34 := v2[""]
    v1 -= (v34 - 128)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
