package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint, p1 int32, p2 uint64) bool {
    v6 := GenS1{F0: 42, F1: 255}
    switch "\n\t" {
    case "hello":
        v7 := []int{1}
        _ = v7
    }
    var v8 int32 = 0
    if (false || true) {
        var v9 uint16 = 2
        _ = v9
    }
    var v10 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v8
    return true
}
func genFn1(p0 string, p1 uint8) uint32 {
    switch p0 {
    case "hello":
        fmt.Println(p1, p0)
    default:
        if true {
            var v11 int64 = 670
            v11 -= v11
            for fi12 := 0; fi12 < 18; fi12++ {
                switch -2147483648 {
                case 2147483647:
                    v11 += v11
                case 1:
                    var v13 int = (fi12 & (fi12 % ((fi12) | 1)))
                    _ = v13
                case -2147483648:
                    var v14 uint8 = (^(p1))
                    _ = v14
                }
                _ = fi12
            }
            _ = v11
        }
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        v16 := GenS2{F0: 1000, F1: fi15}
        fmt.Println(p1)
        v17 := GenS3{F0: 255}
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 2147483648
}

func main() {
    var v1 int = 42
    var v2 uint8 = 2
    var v3 string = ""
    var v4 int16 = 1
    var v5 int8 = 11
    fmt.Println(v4, v3, v4)
    for fi18 := 0; fi18 < 2; fi18++ {
        v19 := []int{fi18, (v1 % (((v1 / ((42) | 1))) | 1))}
        for fi20 := 0; fi20 < 8; fi20++ {
            v2 -= (v2 * 127)
            if (true && false) {
                fmt.Println(fi20, v2)
                if true {
                    var v21 uint = 42
                    _ = v21
                } else {
                    v3 = (v3 + v3)
                    switch (v1 & 1) {
                    case -449:
                        v22 := GenS1{F0: (v1 / ((1) | 1)), F1: 32767}
                        _ = v22
                    }
                }
            } else {
                if ((v4 & v4) <= v4) {
                    if false {
                        fmt.Println(v4, v1)
                        v1 = 128
                        if (1e10 == (-2.25 / (3.14159 - 0.1))) {
                            _ = len(v19)
                        }
                    }
                    var v23 uint32 = 1
                    v19 = append(v19, (fi18 & -9223372036854775808))
                    _ = v23
                } else {
                    if (96 < uint32(v2)) {
                        v24 := GenS4{F0: 65535, F1: 658}
                        fmt.Println(v3, v2, fi18)
                        switch "" {
                        case "foo bar":
                            if len(v19) > 0 {
                                v19[0] = (fi20 % (((fi20 - (v1 / (((fi20 - 529)) | 1)))) | 1))
                            }
                        case "":
                            if true {
                                v25 := GenS3{F0: 42}
                                var v26 int64 = int64(v1)
                                v5 = int8(v1)
                                _ = v25
                                _ = v26
                            } else {
                                v3 = v3
                                fmt.Println(v5, v2, v3)
                                v3 = v3
                            }
                        default:
                            if len(v19) > 0 {
                                v19[0] = (fi20 ^ (-(fi20)))
                            }
                            var v27 float32 = (float32(((1e10 - float64((-2.25 * (1e10 / 1.5)))) + float64(0.0))) / (((1e10 + ((3.14159 + 1e10) / 1.5)) - (-2.25 * 3.14159)) * 1e10))
                            _ = v27
                        }
                        _ = v24
                    } else {
                        if ((v4 / ((849) | 1)) <= v4) {
                            if true {
                                fmt.Println(v4)
                            }
                            var v28 float64 = (1e10 / (1e10 + -2.25))
                            var v29 uint = 4294967295
                            _ = v28
                            _ = v29
                        }
                        switch "a" {
                        case "a":
                            fmt.Println(v3, fi20, fi18)
                        case "":
                            var v30 uint64 = 18446744073709551615
                            _ = v30
                        default:
                            v4 = (v4 % ((v4) | 1))
                        }
                    }
                }
                v5 -= 1
                switch v1 {
                case 128:
                    var v31 int8 = v5
                    _ = v31
                }
            }
            v1 += fi20
            _ = fi20
        }
        fmt.Println(fi18, v5, v1)
        _ = fi18
        _ = v19
    }
    v4 -= 42
    if true {
        fmt.Println(v2, v1, v2)
        for fi32 := 0; fi32 < 5; fi32++ {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v2, v5)
            case "a":
                v33 := []int{-1, fi32}
                _ = v33
            default:
                fmt.Println(v2)
            }
            if true {
                var v34 uint32 = 1
                _ = v34
            } else {
                var v35 uint32 = 407
                switch 511 {
                case 552:
                    var v36 uint16 = 664
                    _ = v36
                case 1:
                    v37 := []int{v1, (fi32 | fi32), (fi32 / (((^(fi32))) | 1))}
                    _ = v37
                }
                _ = v35
            }
            _ = fi32
        }
        fmt.Println(v1, v5, v4)
    }
    var v38 float32 = 1.5
    var v39 uint8 = v2
    v5 = int8(v4)
    var v40 float64 = (0.1 / 0.1)
    if (9223372036854775807 >= -9223372036854775808) {
        var v41 int64 = -9223372036854775808
        _ = v41
    }
    var v42 uint8 = 73
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v5
}
