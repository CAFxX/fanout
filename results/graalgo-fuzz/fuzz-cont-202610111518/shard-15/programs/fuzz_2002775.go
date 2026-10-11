package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int32; F1 int16; F2 uint }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 float64 = (0.1 + (-2.25 / (1e10 + ((1e10 + 1e10) / ((((((3.14159 + float64(-2.25)) * (0.0 * ((float64(((0.1 * 1e10) - -2.25)) - float64((1.5 * -2.25))) * ((((((0.0 * (3.14159 / 1e10)) / 0.1) + 3.14159) + 1.5) + 3.14159) - -2.25)))) * (-2.25 + 1e10)) + ((((3.14159 - 0.1) / (0.1 * (3.14159 * -2.25))) * (float64(1.5) * (-2.25 - float64(1.5)))) * -2.25)) - (-2.25 - 1.5)) - 0.1)))))
    var v2 uint8 = 127
    var v3 bool = false
    var v4 uint16 = 2
    var v5 bool = (false && (2 <= 446))
    v6 := GenS1{F0: 42, F1: v4}
    v3 = v5
    if (1 != v2) {
        if (-112 >= 0) {
            var v7 uint64 = uint64(v4)
            var v8 uint = 42
            _ = v7
            _ = v8
        }
        fmt.Println(v2, v4, v1)
        v1 += v1
    }
    switch -9223372036854775808 {
    case 32768:
        if (!(!(499 == 820))) {
            v9 := GenS2{F0: ((v4 < (v4 | 255)) && v3), F1: 2147483647}
            for fi10 := 0; fi10 < 15; fi10++ {
                switch "\n\t" {
                case "foo bar":
                    for fi11 := 0; fi11 < 8; fi11++ {
                        var v12 uint32 = 976
                        v12 += v12
                        v2 = 255
                        _ = fi11
                        _ = v12
                    }
                case "hello":
                    if (-1 >= 0) {
                        fmt.Println(v2, v4)
                    }
                default:
                    if (0 == 1) {
                        if (("hello" + "\n\t") <= "foo bar") {
                            fmt.Println(fi10)
                            var v13 int8 = 0
                            _ = v13
                        }
                        v14 := []int{}
                        fmt.Println(v1, v4, v2)
                        _ = v14
                    }
                }
                _ = fi10
            }
            _ = v9
        }
    case -92:
        for fi15 := 0; fi15 < 11; fi15++ {
            fmt.Println(v2)
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 uint = 18446744073709551615
                if false {
                    var v18 uint8 = (^(v2))
                    _ = v18
                } else {
                    v19 := []int{fi16, (fi16 / ((fi15) | 1)), (-(fi16)), (fi15 | (^(fi15)))}
                    v20 := GenS3{F0: -2147483648, F1: 1, F2: 636}
                    v2 -= v2
                    _ = v19
                    _ = v20
                }
                var v21 uint64 = 0
                _ = fi16
                _ = v17
                _ = v21
            }
            if (((-2.25 - 3.14159) / 3.14159) <= 1e10) {
                switch 32768 {
                case 2147483647:
                    var v22 string = "hello"
                    _ = v22
                case 4294967295:
                    var v23 uint64 = 0
                    _ = v23
                default:
                    v3 = (-1 >= 730)
                }
                if v5 {
                    fmt.Println(v5, v5)
                    for fi24 := 0; fi24 < 18; fi24++ {
                        v1 = ((1e10 - v1) + float64(1.5))
                        v5 = v3
                        _ = fi24
                    }
                } else {
                    for fi25 := 0; fi25 < 6; fi25++ {
                        switch (-(fi15)) {
                        case -680:
                            fmt.Println(fi15, v1)
                        case 2147483647:
                            v26 := GenS4{F0: v2}
                            _ = v26
                        default:
                            v27 := []int{(fi25 + (fi15 & -2147483648)), 128}
                            _ = v27
                        }
                        _ = fi25
                    }
                }
            } else {
                var v28 uint = 1
                var v29 uint = v28
                _ = v28
                _ = v29
            }
            _ = fi15
        }
    case 0:
        fmt.Println(v1)
    }
    v2 = 115
    v3 = v5
    v30 := []int{4294967295}
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
