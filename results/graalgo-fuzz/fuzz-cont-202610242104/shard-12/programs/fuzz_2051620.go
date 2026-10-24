package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint64 }
type GenS2 struct { F0 bool; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint; F1 int32; F2 int }
type GenS4 struct { F0 int }
func genFn0(p0 string, p1 int64, p2 bool) uint8 {
    v7 := []int{4294967295, 127, 1}
    if true {
        var v8 string = p0
        _ = v8
    } else {
        var v9 bool = (42 <= 18446744073709551615)
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 127
}
func genFn1(p0 int8, p1 uint8, p2 uint32) int8 {
    v10 := []int{-2147483648, 128, -9223372036854775808}
    _ = len(v10)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 2147483647
    v3 := make(map[string]string)
    var v4 int = (^(v2))
    v5 := make(map[int]int)
    var v6 int16 = -1
    v11 := []int{(-(v4)), v4, (v4 + -2147483648)}
    fmt.Println(v6)
    var v12 int32 = -1
    switch "" {
    case "a":
        _ = len(v11)
    }
    v5[v2] = v4
    delete(v1, (v4 % ((2147483647) | 1)))
    if false {
        if true {
            if len(v11) > 0 {
                v11[0] = 32767
            }
            v6 = (v6 ^ (v6 + v6))
            var v13 int8 = 127
            _ = v13
        } else {
            v2 = v4
            for fi14 := 0; fi14 < 15; fi14++ {
                if ((v2 % (((-(fi14))) | 1)) == 520) {
                    var v15 int64 = -9223372036854775808
                    v15 = (-(v15))
                    _ = v15
                }
                switch (v2 | (fi14 * (-(v2)))) {
                case -560:
                    v12 = (-(v12))
                case 127:
                    v16 := GenS1{F0: ("" + ("" + "foo bar")), F1: 1, F2: 689}
                    _ = v16
                case 745:
                    if false {
                        v12 -= (v12 | (v12 - (v12 >> ((146) & 63))))
                        v17 := GenS1{F0: ("" + (("a" + "foo bar") + "hello")), F1: 899, F2: 2}
                        switch "\n\t" {
                        case "hello":
                            if (!false) {
                                var v18 int32 = v12
                                fmt.Println(v4, v12)
                                _ = v18
                            } else {
                                fmt.Println(v4, v6)
                                var v19 int8 = -128
                                _ = v19
                            }
                        case "\n\t":
                            var v20 string = ""
                            _ = v20
                        }
                        _ = v17
                    }
                }
                switch (v2 | (v4 * v4)) {
                case 32767:
                    if (true || (!(489 == 42))) {
                        v4 -= fi14
                        var v21 int8 = 116
                        _ = v21
                    }
                case 32768:
                    for fi22 := 0; fi22 < 3; fi22++ {
                        if ((false || true) || ((fi22 * v4) == 42)) {
                            switch ("foo bar" + ("a" + "a")) {
                            case "":
                                v12 = v12
                            case "foo bar":
                                var v23 int32 = v12
                                _ = v23
                            case "hello":
                                v24 := GenS2{F0: (true && ((v6 >= int16(fi22)) || false)), F1: (v12 | 1), F2: 1}
                                _ = v24
                            }
                            switch "hello" {
                            case "a":
                                var v25 float32 = 1.5
                                _ = v25
                            }
                            v2 = fi14
                            switch fi14 {
                            case -422:
                                v2 += int(v6)
                            case -2147483648:
                                if (!true) {
                                    if (2147483647 > 0) {
                                        switch v4 {
                                        case -568:
                                            fmt.Println(v4, fi14)
                                        default:
                                            v26 := GenS1{F0: ((("\n\t" + (("\n\t" + "") + "a")) + ("hello" + "a")) + "foo bar"), F1: 128, F2: 0}
                                            fmt.Println(fi14)
                                            _ = v26
                                        }
                                        var v27 string = ""
                                        var v28 uint32 = 2147483647
                                        if (!false) {
                                            switch (v27 + v27) {
                                            case "a":
                                                v6 -= v6
                                            case "\n\t":
                                                var v29 int16 = v6
                                                _ = v29
                                            default:
                                                var v30 int32 = v12
                                                _ = v30
                                            }
                                            v2 = (-(fi22))
                                        } else {
                                            v31 := GenS1{F0: v27, F1: 9223372036854775807, F2: 4294967295}
                                            switch v4 {
                                            case 128:
                                                var v32 float64 = 3.14159
                                                _ = v32
                                            default:
                                                fmt.Println(v2, v4)
                                            }
                                            v2 -= (fi22 & v2)
                                            _ = v31
                                        }
                                        _ = v27
                                        _ = v28
                                    }
                                    var v33 int32 = 1
                                    if len(v11) > 0 {
                                        v11[0] = -2147483648
                                    }
                                    _ = v33
                                }
                            case -9223372036854775808:
                                var v34 int64 = int64(v6)
                                _ = v34
                            default:
                                v35 := v1[(-(v4))]
                                _ = v35
                            }
                            if len(v11) > 0 {
                                v11[0] = (-(v2))
                            }
                        } else {
                            v12 += 2147483647
                            v36 := GenS2{F0: (32768 >= fi22), F1: (v12 * v12), F2: 983}
                            _ = v36
                        }
                        v2 = v4
                        _ = len(v5)
                        if (false || (!false)) {
                            if ((false && (0.1 <= (0.0 * 0.0))) && false) {
                                var v37 int = (v2 - (fi22 - v4))
                                _ = v37
                            } else {
                                fmt.Println(fi22, fi22, v2)
                            }
                        }
                        _ = fi22
                    }
                }
                _ = fi14
            }
        }
        var v38 bool = false
        _ = v38
    } else {
        var v39 int16 = (v6 % ((v6) | 1))
        _ = v39
    }
    v40 := GenS2{F0: false, F1: (v12 % ((v12) | 1)), F2: 1}
    v41 := GenS3{F0: 2, F1: 0, F2: 0}
    v42 := GenS4{F0: -944}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
