package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint16, p1 int) int64 {
    if false {
        for fi7 := 0; fi7 < 14; fi7++ {
            switch fi7 {
            case 42:
                if ((!false) || (false || false)) {
                    var v8 int8 = 100
                    if false {
                        v9 := []int{(fi7 % ((1) | 1)), -2147483648, (^(p1))}
                        if len(v9) > 0 {
                            v9[0] = (fi7 % (((p1 + p1)) | 1))
                        }
                        v10 := GenS1{F0: (-(p1)), F1: (fi7 >> ((255) & 63))}
                        _ = v10
                        _ = v9
                    }
                    _ = v8
                } else {
                    fmt.Println(p1)
                }
            default:
                p0 += (p0 % (((p0 + (^(p0)))) | 1))
            }
            _ = fi7
        }
        var v11 uint64 = 144
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := GenS1{F0: (fi12 + 32767), F1: fi12}
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    v14 := GenS2{F0: 2, F1: 4294967295}
    _ = p0
    _ = p1
    _ = v14
    return 1
}
func genFn1() int {
    var v15 int32 = 485
    fmt.Println(v15)
    v15 -= (^(v15))
    v16 := []int{0, -2147483648, 42}
    for fi17 := 0; fi17 < 1; fi17++ {
        switch "\n\t" {
        case "foo bar":
            v16 = append(v16, 475)
        case "\n\t":
            v15 = (v15 | v15)
        case "hello":
            for fi18 := 0; fi18 < 10; fi18++ {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                var v19 bool = (true || (-2.25 > float32(1e10)))
                var v20 uint = 65535
                _ = fi18
                _ = v19
                _ = v20
            }
        default:
            for fi21 := 0; fi21 < 20; fi21++ {
                if len(v16) > 0 {
                    v16[0] = fi17
                }
                v22 := GenS1{F0: (fi17 << ((4294967295) & 63)), F1: (^(fi17))}
                v23 := GenS1{F0: -9223372036854775808, F1: 32768}
                _ = fi21
                _ = v22
                _ = v23
            }
        }
        v15 -= v15
        _ = fi17
    }
    _ = v15
    _ = v16
    return 127
}

func main() {
    v1 := []string{"a", (((("a" + "foo bar") + "foo bar") + "foo bar") + ("\n\t" + "foo bar"))}
    v2 := make(map[int]uint32)
    v3 := []int{32768}
    var v4 int16 = 1
    v5 := []string{}
    v6 := make(map[uint32]uint32)
    for fi24 := 0; fi24 < 7; fi24++ {
        if (false && (1 != 42)) {
            v25 := GenS2{F0: 4294967295, F1: 835}
            _ = v25
        }
        v4 -= -32768
        fmt.Println(fi24)
        _ = fi24
    }
    var v26 float64 = (1.5 - (0.0 - ((0.1 * ((0.0 / -2.25) - 1e10)) * (1.5 - float64(0.1)))))
    fmt.Println(v4, v26)
    var v27 int64 = -1
    switch genFn1() {
    case 769:
        switch "a" {
        case "":
            switch -2147483648 {
            case -912:
                v28 := GenS1{F0: -9223372036854775808, F1: 636}
                _ = v28
            }
        case "hello":
            var v29 float64 = float64(1.5)
            _ = v29
        default:
            v4 -= int16(v27)
        }
    case 32767:
        if true {
            var v30 int16 = v4
            v26 = 0.1
            _ = v30
        } else {
            v26 -= v26
            v31 := GenS3{F0: 18446744073709551615}
            if len(v5) > 0 {
                v5[0] = "foo bar"
            }
            _ = v31
        }
    }
    fmt.Println(v27, v4)
    var v32 bool = (1.5 >= float32(float64(1.5)))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    delete(v6, 2)
    var v33 uint64 = 42
    var v34 int8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
