package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint }
type GenS2 struct { F0 int64; F1 uint64; F2 int }
func genFn0(p0 int64, p1 int32) uint32 {
    p1 = -898
    p1 = -57
    p1 += -1
    v4 := []int{-2147483648, 989, 128, -424}
    var v5 int32 = (-(p1))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2147483647
}

func main() {
    v1 := []int{-2147483648}
    v2 := []int{}
    v3 := make(map[int]bool)
    if true {
        switch -9223372036854775808 {
        case 130:
            for fi6 := 0; fi6 < 16; fi6++ {
                for fi7 := 0; fi7 < 18; fi7++ {
                    fmt.Println(fi7, fi6)
                    v8 := GenS1{F0: 2, F1: 2147483647, F2: 1}
                    if false {
                        delete(v3, (fi7 - fi7))
                    } else {
                        var v9 uint16 = 65535
                        _ = v9
                    }
                    v10 := GenS1{F0: genFn0(-1, 0), F1: 65535, F2: 250}
                    _ = fi7
                    _ = v10
                    _ = v8
                }
                switch 0 {
                case 128:
                    for fi11 := 0; fi11 < 2; fi11++ {
                        switch ("foo bar" + "foo bar") {
                        case "hello":
                            var v12 int = -1
                            _ = v12
                        case "a":
                            var v13 int = 1
                            _ = v13
                        }
                        switch "\n\t" {
                        case "\n\t":
                            var v14 uint32 = 594
                            v14 = v14
                            switch fi6 {
                            case 9223372036854775807:
                                if false {
                                    v2 = append(v2, fi11)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            case 127:
                                if true {
                                    var v15 bool = false
                                    var v16 int = (fi11 << ((0) & 63))
                                    _ = v15
                                    _ = v16
                                } else {
                                    if false {
                                        var v17 uint64 = 0
                                        fmt.Println(v17, fi6, v17)
                                        _ = v17
                                    } else {
                                        switch ("" + "\n\t") {
                                        case "foo bar":
                                            var v18 int16 = 42
                                            _ = v18
                                        }
                                        fmt.Println(fi6, fi11, fi11)
                                        var v19 int32 = 2147483647
                                        _ = v19
                                    }
                                    v14 = (v14 & (^(v14)))
                                }
                                _ = len(v2)
                                v14 -= (v14 % ((0) | 1))
                            }
                            _ = v14
                        case "":
                            switch (fi6 * fi11) {
                            case -10:
                                var v20 float64 = (3.14159 / 1.5)
                                _ = v20
                            case 32767:
                                var v21 int32 = 1
                                _ = v21
                            default:
                                var v22 int16 = 0
                                _ = v22
                            }
                        case "foo bar":
                            var v23 uint8 = 189
                            _ = v23
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi11
                    }
                case 9223372036854775807:
                    for fi24 := 0; fi24 < 3; fi24++ {
                        var v25 float64 = 0.0
                        _ = fi24
                        _ = v25
                    }
                case -466:
                    for fi26 := 0; fi26 < 20; fi26++ {
                        var v27 uint64 = 42
                        _ = fi26
                        _ = v27
                    }
                }
                _ = fi6
            }
        case -2147483648:
            var v28 uint64 = 256
            _ = v28
        case 267:
            var v29 uint16 = 276
            _ = v29
        default:
            for fi30 := 0; fi30 < 13; fi30++ {
                var v31 int64 = -9223372036854775808
                _ = fi30
                _ = v31
            }
        }
    }
    var v32 int8 = 100
    switch -2147483648 {
    case 0:
        v32 = 42
    case 42:
        v33 := GenS1{F0: 65535, F1: 2147483647, F2: 467}
        _ = v33
    case 1:
        var v34 uint16 = 0
        _ = v34
    default:
        var v35 int8 = (v32 & v32)
        _ = v35
    }
    var v36 uint16 = 0
    v37 := v3[-146]
    _ = len(v2)
    delete(v3, -593)
    fmt.Println(v32, v32, v37)
    for fi38 := 0; fi38 < 2; fi38++ {
        v32 += v32
        _ = fi38
    }
    _ = len(v3)
    var v39 uint = 0
    v40 := GenS2{F0: 25, F1: 773, F2: -2147483648}
    v32 -= (v32 % ((v32) | 1))
    var v41 string = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v37
    _ = v39
    _ = v40
    _ = v41
}
