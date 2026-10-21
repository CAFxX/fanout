package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint16 }
type GenS2 struct { F0 int; F1 int8; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint64; F2 string }
type GenS4 struct { F0 uint }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    var v4 float64 = ((0.0 * 1e10) - -2.25)
    var v5 int8 = 69
    fmt.Println(v1, v5, v1)
    v5 += 72
    switch "a" {
    case "":
        fmt.Println(v4)
    case "\n\t":
        switch ("a" + ("\n\t" + "a")) {
        case "hello":
            if true {
                for fi6 := 0; fi6 < 12; fi6++ {
                    v7 := GenS1{F0: (1 < -97), F1: true, F2: 42}
                    v1 = v1
                    _ = fi6
                    _ = v7
                }
                delete(v3, ((("a" + "hello") + "a") + "foo bar"))
                fmt.Println(v1, v5)
            }
        default:
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := GenS1{F0: (!(!false)), F1: false, F2: 2}
                _ = fi8
                _ = v9
            }
        }
    case "foo bar":
        if ((1 != 42) && true) {
            var v10 int16 = 1
            var v11 int64 = (^(v1))
            _ = v10
            _ = v11
        } else {
            switch ("\n\t" + "") {
            case "":
                for fi12 := 0; fi12 < 20; fi12++ {
                    v1 = (v1 | (v1 & -1))
                    v13 := GenS1{F0: (-1 < 1), F1: true, F2: 32767}
                    v14 := []int{128, (fi12 / ((826) | 1))}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            case "hello":
                if false {
                    fmt.Println(v5, v4, v4)
                    v15 := GenS2{F0: 104, F1: 42, F2: 2147483647}
                    _ = v15
                } else {
                    delete(v2, "foo bar")
                }
            default:
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 uint = 0
                    v5 = 0
                    _ = fi16
                    _ = v17
                }
            }
            var v18 uint32 = 219
            _ = v18
        }
    default:
        v19 := []int{127}
        _ = v19
    }
    v3[""] = 128
    var v20 int = 103
    switch "foo bar" {
    case "a":
        v5 -= (^(v5))
    default:
        v5 = (v5 / ((int8(v1)) | 1))
    }
    v21 := GenS3{F0: 764, F1: 772, F2: ("\n\t" + ("a" + ""))}
    fmt.Println(v4)
    v5 = 97
    v1 = (v1 ^ -9223372036854775808)
    v4 = float64(float32((((v4 / v4) - v4) / v4)))
    _ = len(v2)
    if true {
        switch (v20 & (-(v20))) {
        case -9223372036854775808:
            if true {
                var v22 uint = 2
                if ((float32((-2.25 + 0.1)) - (3.14159 - ((1.5 + 3.14159) + 3.14159))) <= 0.0) {
                    if (1e10 < (v4 + float64(float32(3.14159)))) {
                        v23 := []int{32768, (v20 >> ((v22) & 63)), (v20 * v20)}
                        for fi24 := 0; fi24 < 17; fi24++ {
                            for fi25 := 0; fi25 < 15; fi25++ {
                                v26 := GenS1{F0: (false && true), F1: false, F2: 2}
                                _ = len(v23)
                                var v27 bool = false
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                            if ("hello" >= (("" + "foo bar") + "foo bar")) {
                                v28 := GenS1{F0: true, F1: ("a" > "hello"), F2: 1}
                                v1 = v1
                                _ = v28
                            } else {
                                fmt.Println(v22, v20)
                                switch 32767 {
                                case 2147483647:
                                    var v29 float32 = (-2.25 / (float32(v4) / 1e10))
                                    _ = v29
                                case 181:
                                    v23 = append(v23, (v20 % ((v20) | 1)))
                                case 127:
                                    fmt.Println(v5)
                                default:
                                    fmt.Println(v20, v4)
                                }
                                v22 = 18446744073709551615
                            }
                            if false {
                                v30 := GenS4{F0: 708}
                                var v31 int = -9223372036854775808
                                switch (("" + ("a" + "a")) + "a") {
                                case "hello":
                                    var v32 float64 = v4
                                    _ = v32
                                case "\n\t":
                                    if (0 >= 18446744073709551615) {
                                        v20 = (v31 % ((32768) | 1))
                                        var v33 uint = v22
                                        _ = v33
                                    }
                                }
                                _ = v30
                                _ = v31
                            } else {
                                for fi34 := 0; fi34 < 7; fi34++ {
                                    var v35 uint32 = 1
                                    switch (fi34 % (((fi34 + (^(v20)))) | 1)) {
                                    case 1:
                                        v36 := GenS2{F0: (-(fi34)), F1: 42, F2: -1}
                                        _ = v36
                                    case 42:
                                        switch (fi24 / ((776) | 1)) {
                                        case -834:
                                            v23 = append(v23, (v20 ^ -486))
                                            v5 += (v5 + (-(v5)))
                                            fmt.Println(v35)
                                        default:
                                            v20 = v20
                                        }
                                    case -1:
                                        v22 += v22
                                        v22 += v22
                                    default:
                                        if true {
                                            v5 = v5
                                            var v37 bool = ((!false) || (true && (v35 == v35)))
                                            v38 := v2["hello"]
                                            _ = v37
                                            _ = v38
                                        }
                                    }
                                    v39 := v2[("hello" + "")]
                                    _ = fi34
                                    _ = v35
                                    _ = v39
                                }
                                if len(v23) > 0 {
                                    v23[0] = v20
                                }
                            }
                            _ = fi24
                        }
                        _ = v23
                    }
                    var v40 uint32 = 938
                    _ = v40
                }
                var v41 uint = 42
                _ = v22
                _ = v41
            } else {
                var v42 uint16 = 905
                v1 -= (v1 << ((432) & 63))
                v20 -= (v20 / (((^(v20))) | 1))
                _ = v42
            }
        default:
            switch v20 {
            case 32768:
                v43 := []int{(v20 | 127), (v20 / (((^(v20))) | 1)), (^(v20))}
                _ = v43
            case 522:
                v44 := GenS4{F0: 255}
                _ = v44
            default:
                switch "" {
                case "foo bar":
                    for fi45 := 0; fi45 < 10; fi45++ {
                        var v46 float32 = (0.0 * 1e10)
                        var v47 int16 = 32767
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                default:
                    v48 := v3[("\n\t" + "foo bar")]
                    _ = v48
                }
            }
        }
    }
    v49 := GenS4{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v49
    _ = v5
}
