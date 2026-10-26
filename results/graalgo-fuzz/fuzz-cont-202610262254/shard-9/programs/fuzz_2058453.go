package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint16 = 1
    v2 := []uint32{2, 840, 2, 2147483647}
    v3 := []string{}
    v1 -= v1
    var v4 int64 = 9223372036854775807
    v5 := GenS1{F0: v4, F1: 4294967295}
    fmt.Println(v1, v1)
    fmt.Println(v4, v1)
    v6 := GenS2{F0: v1, F1: 0}
    for fi7 := 0; fi7 < 17; fi7++ {
        switch "" {
        case "foo bar":
            v3 = append(v3, "hello")
        default:
            v4 = (v4 ^ v4)
        }
        _ = fi7
    }
    if (!(!true)) {
        var v8 float32 = 1e10
        if ((!(!true)) || ((-2.25 + v8) < (3.14159 - -2.25))) {
            _ = len(v2)
            var v9 uint = 18446744073709551615
            var v10 int = 42
            _ = v10
            _ = v9
        } else {
            v1 -= v1
            for fi11 := 0; fi11 < 9; fi11++ {
                for fi12 := 0; fi12 < 14; fi12++ {
                    var v13 string = "\n\t"
                    _ = fi12
                    _ = v13
                }
                if (false || (!(1e10 > 0.1))) {
                    switch "foo bar" {
                    case "foo bar":
                        v14 := GenS1{F0: (v4 * (^(v4))), F1: (fi11 ^ 42)}
                        _ = v14
                    case "":
                        var v15 int16 = 304
                        _ = v15
                    }
                    fmt.Println(fi11)
                    for fi16 := 0; fi16 < 10; fi16++ {
                        v1 += (v1 + (^(v1)))
                        if false {
                            var v17 int = fi16
                            v4 -= (v4 ^ 9223372036854775807)
                            _ = v17
                        }
                        _ = fi16
                    }
                }
                var v18 uint8 = 147
                _ = fi11
                _ = v18
            }
            if (!false) {
                v19 := GenS1{F0: -9223372036854775808, F1: -9223372036854775808}
                _ = v19
            } else {
                if false {
                    v2 = append(v2, 925)
                } else {
                    v8 = (0.1 - ((v8 / 3.14159) / ((v8 / 1.5) * v8)))
                    if true {
                        if (-1 != (v4 * v4)) {
                            if (0.1 >= ((1.5 - 1e10) / float64(v8))) {
                                var v20 uint64 = 1
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v8 -= 0.0
                                _ = v20
                            }
                        }
                        if (false && true) {
                            var v21 uint16 = 255
                            _ = v21
                        } else {
                            v22 := GenS2{F0: 42, F1: 127}
                            var v23 uint64 = 441
                            _ = v22
                            _ = v23
                        }
                        var v24 int32 = -2147483648
                        _ = v24
                    }
                    if len(v2) > 0 {
                        v2[0] = uint32(v1)
                    }
                }
                var v25 float32 = v8
                var v26 int32 = 0
                _ = v25
                _ = v26
            }
        }
        var v27 uint32 = 330
        _ = v27
        _ = v8
    } else {
        var v28 int32 = -744
        v29 := GenS1{F0: (v4 + -920), F1: 32768}
        if (4294967295 >= 1) {
            if (53 >= -1) {
                fmt.Println(v4, v28)
                fmt.Println(v1)
            } else {
                if false {
                    var v30 uint64 = 0
                    v28 -= (v28 * v28)
                    v31 := GenS3{F0: (v30 & v30)}
                    _ = v30
                    _ = v31
                } else {
                    if true {
                        var v32 uint16 = (v1 - (v1 & v1))
                        var v33 float64 = float64(1.5)
                        var v34 int64 = -9223372036854775808
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                }
            }
            fmt.Println(v28, v1, v1)
        }
        _ = v28
        _ = v29
    }
    v1 = v1
    v1 = v1
    var v35 int = 9223372036854775807
    switch "a" {
    case "hello":
        v35 = v35
    case "a":
        if true {
            for fi36 := 0; fi36 < 10; fi36++ {
                var v37 uint8 = 255
                _ = fi36
                _ = v37
            }
            _ = len(v2)
        }
    case "\n\t":
        fmt.Println(v4, v35)
    default:
        var v38 uint8 = 40
        _ = v38
    }
    var v39 uint64 = 42
    v39 += v39
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
