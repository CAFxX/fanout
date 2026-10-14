package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint16; F1 bool; F2 bool }
type GenS3 struct { F0 uint64; F1 uint }

func main() {
    v1 := []string{"", "", "foo bar"}
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v4 int64 = 9223372036854775807
    var v5 float64 = 1e10
    var v6 uint16 = 295
    var v7 string = "\n\t"
    for fi8 := 0; fi8 < 19; fi8++ {
        switch v7 {
        case "a":
            if false {
                switch fi8 {
                case 394:
                    v6 -= 1
                case 1:
                    v9 := GenS1{F0: 255, F1: 2}
                    _ = v9
                default:
                    var v10 uint64 = 18446744073709551615
                    _ = v10
                }
                fmt.Println(v7)
                var v11 bool = (fi8 == fi8)
                _ = v11
            } else {
                if false {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        switch 128 {
                        case 127:
                            v1 = append(v1, "foo bar")
                        case 0:
                            if true {
                                if true {
                                    v4 += v4
                                    var v13 uint8 = 255
                                    var v14 bool = (false || true)
                                    _ = v13
                                    _ = v14
                                }
                                v15 := v3[fi12]
                                _ = v15
                            }
                        }
                        _ = fi12
                    }
                }
                fmt.Println(v5)
            }
        case "\n\t":
            v16 := v2[("" + v7)]
            _ = v16
        default:
            for fi17 := 0; fi17 < 2; fi17++ {
                v7 = "a"
                if true {
                    var v18 int64 = (v4 / ((v4) | 1))
                    _ = v18
                }
                _ = fi17
            }
        }
        if ((6 != 217) || (0.0 > 1.5)) {
            switch fi8 {
            case 127:
                v19 := GenS1{F0: 462, F1: 779}
                _ = v19
            case 9223372036854775807:
                var v20 uint8 = 156
                _ = v20
            default:
                if false {
                    if false {
                        var v21 bool = true
                        v22 := GenS1{F0: 255, F1: 4294967295}
                        _ = v21
                        _ = v22
                    }
                }
            }
        }
        _ = fi8
    }
    v6 = v6
    switch -2147483648 {
    case 979:
        switch 9223372036854775807 {
        case 127:
            var v23 float32 = 0.1
            _ = v23
        case -1:
            v24 := GenS2{F0: v6, F1: (float64(0.0) >= ((3.14159 * (1e10 - v5)) - v5)), F2: (42 <= 294)}
            _ = v24
        }
    }
    v7 = v7
    v1 = append(v1, "hello")
    if false {
        var v25 uint32 = 234
        _ = v25
    } else {
        if (32767 < v6) {
            v26 := GenS3{F0: 42, F1: 18446744073709551615}
            _ = v26
        }
    }
    var v27 int16 = 1000
    var v28 float32 = (1e10 * 1.5)
    fmt.Println(v5)
    delete(v2, "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
