package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 int32; F1 uint16 }
func genFn0(p0 string) int32 {
    var v6 int64 = -953
    if true {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 int32 = -1
            v8 = v8
            v9 := GenS1{F0: -1, F1: (fi7 << ((2) & 63))}
            _ = fi7
            _ = v8
            _ = v9
        }
        fmt.Println(v6, v6)
    }
    p0 = "foo bar"
    if (false && ((0.0 > (-2.25 - 0.0)) || false)) {
        if (-9223372036854775808 <= 323) {
            if ("a" <= p0) {
                fmt.Println(v6)
                fmt.Println(v6)
            } else {
                var v10 int8 = 127
                _ = v10
            }
        }
    }
    _ = p0
    _ = v6
    return -2147483648
}
func genFn1() uint8 {
    switch ("\n\t" + ((("hello" + "foo bar") + "foo bar") + (("\n\t" + "hello") + ((("hello" + "\n\t") + "a") + "foo bar")))) {
    case "\n\t":
        if ((42 > 2147483647) || (!(-128 == 127))) {
            var v11 bool = false
            if v11 {
                switch 42 {
                case -909:
                    v12 := GenS2{F0: 32768, F1: -2147483648}
                    _ = v12
                case -1:
                    if v11 {
                        v11 = true
                        v13 := []int{-2147483648, 2147483647, 2147483647, 659}
                        _ = v13
                    } else {
                        var v14 uint32 = 42
                        if (true && v11) {
                            v15 := GenS3{F0: 18446744073709551615, F1: 255, F2: (v14 | v14)}
                            _ = v15
                        } else {
                            var v16 uint32 = v14
                            switch 4294967295 {
                            case -2147483648:
                                v16 = (v16 & 2)
                            case -569:
                                v14 = v14
                            default:
                                var v17 uint8 = 80
                                _ = v17
                            }
                            v11 = ((0 < int8(v14)) || v11)
                            _ = v16
                        }
                        v11 = (-2.25 == 0.1)
                        _ = v14
                    }
                default:
                    fmt.Println(v11)
                }
                switch ("\n\t" + "a") {
                case "hello":
                    var v18 float64 = (0.0 + (1.5 - (0.1 / 1.0)))
                    _ = v18
                case "a":
                    fmt.Println(v11)
                case "foo bar":
                    v11 = (!(!v11))
                default:
                    v11 = v11
                }
                for fi19 := 0; fi19 < 1; fi19++ {
                    fmt.Println(fi19)
                    _ = fi19
                }
            }
            for fi20 := 0; fi20 < 8; fi20++ {
                v21 := []int{-2147483648, -2147483648, (fi20 | fi20), fi20}
                _ = fi20
                _ = v21
            }
            _ = v11
        }
    case "foo bar":
        var v22 string = "foo bar"
        _ = v22
    default:
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 uint16 = 880
            var v25 float64 = 3.14159
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    var v26 int64 = 471
    var v27 float32 = 1e10
    switch "hello" {
    case "":
        v28 := []int{}
        _ = v28
    case "\n\t":
        var v29 int = 32768
        _ = v29
    case "a":
        var v30 uint32 = 113
        _ = v30
    default:
        v26 = v26
    }
    v26 += v26
    _ = v26
    _ = v27
    return 128
}

func main() {
    var v1 int64 = -550
    var v2 int16 = 42
    var v3 uint64 = 18446744073709551615
    v4 := make(map[int]int)
    var v5 float32 = 0.0
    fmt.Println(v3, v2, v1)
    v31 := []int{128, -2147483648, 0, -1}
    v32 := GenS4{F0: 4294967295, F1: 688}
    if len(v31) >= 2 {
        v31 = v31[1:len(v31)]
    }
    v5 -= (-2.25 - v5)
    var v33 int8 = 78
    switch -2147483648 {
    case -9223372036854775808:
        var v34 uint16 = 2
        _ = v34
    case 127:
        v5 -= 1e10
    case 546:
        var v35 uint16 = 797
        _ = v35
    default:
        if true {
            v31 = append(v31, 9223372036854775807)
        }
    }
    var v36 int64 = v1
    v2 += (v2 + (v2 / ((v2) | 1)))
    var v37 uint8 = 144
    switch "a" {
    case "foo bar":
        v38 := v4[42]
        _ = v38
    }
    v3 += (v3 % ((v3) | 1))
    switch "\n\t" {
    case "foo bar":
        v3 = (v3 * 1)
    case "":
        var v39 int32 = 1
        _ = v39
    default:
        v40 := GenS5{F0: 1, F1: 971}
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
