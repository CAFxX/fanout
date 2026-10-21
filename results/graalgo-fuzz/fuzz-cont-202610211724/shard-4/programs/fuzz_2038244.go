package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int32 }

func main() {
    v1 := []int{}
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]uint32)
    if false {
        if true {
            v2 += (-(v2))
            var v4 uint = 18446744073709551615
            switch "\n\t" {
            case "a":
                if true {
                    for fi5 := 0; fi5 < 3; fi5++ {
                        switch "foo bar" {
                        case "foo bar":
                            if false {
                                v4 += v4
                            } else {
                                fmt.Println(v2, v4, fi5)
                                var v6 uint = v4
                                _ = v6
                            }
                        case "\n\t":
                            var v7 int32 = -2147483648
                            _ = v7
                        case "hello":
                            fmt.Println(v4, v2, v4)
                        }
                        v2 += v2
                        var v8 int16 = 32767
                        _ = fi5
                        _ = v8
                    }
                    switch 42 {
                    case 1:
                        v3[2147483647] = 189
                    case -2147483648:
                        var v9 int16 = -32768
                        _ = v9
                    case 9223372036854775807:
                        var v10 uint = 332
                        _ = v10
                    default:
                        for fi11 := 0; fi11 < 9; fi11++ {
                            var v12 int32 = -2147483648
                            v13 := GenS1{F0: 0}
                            v14 := GenS2{F0: 15, F1: (v12 * v12)}
                            _ = fi11
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                    }
                    delete(v3, 334)
                }
            case "\n\t":
                if (977 == 1) {
                    var v15 float64 = (0.1 + -2.25)
                    if ((v2 - v2) >= 42) {
                        if true {
                            fmt.Println(v15, v2, v15)
                            if true {
                                v4 = (v4 + v4)
                            }
                            fmt.Println(v2)
                        }
                    }
                    switch -1 {
                    case -636:
                        for fi16 := 0; fi16 < 11; fi16++ {
                            if false {
                                v17 := GenS1{F0: 255}
                                v18 := GenS1{F0: 0}
                                _ = v17
                                _ = v18
                            }
                            if len(v1) > 0 {
                                v1[0] = (-(fi16))
                            }
                            _ = fi16
                        }
                    }
                    _ = v15
                } else {
                    var v19 uint16 = 32767
                    _ = v19
                }
            }
            _ = v4
        } else {
            v3[-2147483648] = 42
        }
        v2 -= (v2 | 848)
    }
    switch ("a" + ("a" + "hello")) {
    case "\n\t":
        v3[-548] = 2
    case "":
        if (3.14159 <= 1.5) {
            if true {
                switch ("foo bar" + "\n\t") {
                case "":
                    v2 = (v2 / (((v2 / ((337) | 1))) | 1))
                default:
                    for fi20 := 0; fi20 < 18; fi20++ {
                        v2 = 357
                        switch (fi20 & (^(fi20))) {
                        case 127:
                            for fi21 := 0; fi21 < 17; fi21++ {
                                v3[(fi21 ^ 0)] = 42
                                v2 = v2
                                v2 -= (v2 ^ -961)
                                _ = fi21
                            }
                        case 128:
                            var v22 int64 = v2
                            _ = v22
                        case 32767:
                            var v23 int = 32767
                            _ = v23
                        }
                        v24 := GenS2{F0: 94, F1: -820}
                        _ = fi20
                        _ = v24
                    }
                }
            }
            var v25 int8 = 10
            _ = v25
        } else {
            if false {
                v2 -= (v2 & 9223372036854775807)
                if (false && ((2 <= 4294967295) || false)) {
                    _ = len(v3)
                } else {
                    v2 = v2
                }
            }
        }
    default:
        var v26 float64 = (0.0 / 1.0)
        _ = v26
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v27 int64 = (v2 ^ (v2 * (^(v2))))
    fmt.Println(v2, v27)
    if true {
        var v28 float32 = 0.1
        var v29 float64 = (float64(v28) / 1e10)
        _ = v28
        _ = v29
    } else {
        v2 += v2
        _ = len(v1)
    }
    v30 := GenS1{F0: 127}
    if true {
        if (255 != 1) {
            var v31 uint64 = 4294967295
            _ = v31
        } else {
            var v32 uint32 = 0
            var v33 int8 = -1
            _ = v32
            _ = v33
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
}
