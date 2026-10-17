package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 string }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 int8; F1 int32; F2 uint16 }
func genFn0(p0 uint8, p1 int8) int64 {
    v6 := GenS1{F0: 762, F1: 214, F2: "a"}
    v7 := GenS2{F0: 4294967295, F1: 2}
    p0 = 124
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 876
}
func genFn1() uint8 {
    var v8 int64 = 0
    var v9 float32 = float32(-2.25)
    switch "hello" {
    case "hello":
        v10 := GenS1{F0: 2147483647, F1: -2147483648, F2: "hello"}
        _ = v10
    case "a":
        v11 := GenS2{F0: 1, F1: 42}
        _ = v11
    }
    v12 := []int{0, 128}
    v9 = 1.5
    _ = v12
    _ = v8
    _ = v9
    return 156
}

func main() {
    var v1 uint = 0
    var v2 uint = v1
    var v3 uint32 = 328
    v4 := []string{}
    v5 := []string{"\n\t"}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v2)
    for fi13 := 0; fi13 < 11; fi13++ {
        v2 += (^(v1))
        for fi14 := 0; fi14 < 12; fi14++ {
            switch "" {
            case "hello":
                if (-2147483648 > 0) {
                    if true {
                        if false {
                            if false {
                                var v15 bool = false
                                v4 = append(v4, (("hello" + "\n\t") + (("a" + "") + "foo bar")))
                                var v16 int8 = 100
                                v17 := GenS2{F0: 4294967295, F1: 4294967295}
                                var v18 int = -731
                                _ = v15
                                _ = v16
                                _ = v17
                                _ = v18
                            } else {
                                var v19 uint = v2
                                v19 = v2
                                _ = v19
                            }
                            fmt.Println(fi14)
                            v1 -= v2
                        } else {
                            if len(v4) > 0 {
                                v4[0] = "a"
                            }
                            v3 = v3
                            v1 = (^(v1))
                        }
                    }
                    switch 32767 {
                    case 2147483647:
                        v20 := GenS3{F0: (("foo bar" + (("foo bar" + ("a" + ("a" + ((("\n\t" + "foo bar") + ("foo bar" + ((("hello" + "\n\t") + "hello") + "foo bar"))) + "foo bar")))) + ("\n\t" + ""))) + "foo bar")}
                        _ = v20
                    case -622:
                        fmt.Println(v3, v1, v2)
                    default:
                        v21 := GenS3{F0: "\n\t"}
                        _ = v21
                    }
                    v1 -= (v1 * (v1 | (v1 >> ((v1) & 63))))
                    fmt.Println(v2, v3)
                    if false {
                        v1 += v1
                        v22 := GenS1{F0: 65535, F1: -2147483648, F2: "\n\t"}
                        if (false && true) {
                            if (!((^(v3)) < (v3 % ((0) | 1)))) {
                                v2 = v1
                                var v23 uint64 = 97
                                _ = v23
                            }
                            if (!false) {
                                if len(v4) > 0 {
                                    v4[0] = "hello"
                                }
                                if false {
                                    var v24 string = ("a" + "a")
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    var v25 int32 = -1
                                    _ = v24
                                    _ = v25
                                }
                                v4 = append(v4, "hello")
                            } else {
                                v1 -= v1
                                if true {
                                    if true {
                                        _ = len(v5)
                                    }
                                    var v26 int16 = 1000
                                    v3 -= v3
                                    _ = v26
                                }
                                v27 := GenS2{F0: 1, F1: 1}
                                _ = v27
                            }
                            var v28 float64 = 3.14159
                            _ = v28
                        }
                        _ = v22
                    } else {
                        v2 -= uint(v3)
                        if true {
                            switch -1 {
                            case -9223372036854775808:
                                v29 := GenS3{F0: ""}
                                _ = v29
                            case -2147483648:
                                v1 += v1
                            case 560:
                                fmt.Println(fi13, fi13)
                            }
                        }
                        v1 = (v2 % (((v2 ^ (^(v1)))) | 1))
                    }
                } else {
                    v3 -= (^(v3))
                    var v30 int32 = 0
                    _ = v30
                }
            case "foo bar":
                if len(v5) > 0 {
                    v5[0] = "\n\t"
                }
            }
            if (!false) {
                var v31 uint32 = 1
                if ((!false) && false) {
                    v1 += v1
                    if (2147483647 != -1) {
                        var v32 int = fi14
                        _ = v32
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            v33 := GenS4{F0: (fi14 << (((^(v2))) & 63)), F1: 61}
                            _ = v33
                        case "hello":
                            v1 = v2
                        }
                        if (607 > 42) {
                            v4 = append(v4, "")
                        }
                        v31 = v31
                    }
                }
                v2 = v2
                _ = v31
            } else {
                switch fi14 {
                case -9223372036854775808:
                    switch fi14 {
                    case 1:
                        v34 := GenS3{F0: ""}
                        _ = v34
                    case 127:
                        var v35 float64 = (((1.5 / (-2.25 - 0.0)) - (-2.25 / 1.5)) * 1e10)
                        _ = v35
                    case 465:
                        v36 := GenS5{F0: 0, F1: -1, F2: 0}
                        v4 = append(v4, "a")
                        v3 -= 731
                        _ = v36
                    }
                case 127:
                    fmt.Println(fi13, fi13)
                }
                var v37 float64 = (0.0 + (3.14159 - ((3.14159 * -2.25) - 0.1)))
                _ = v37
            }
            _ = fi14
        }
        _ = fi13
    }
    fmt.Println(v2, v2)
    fmt.Println(v1, v2, v2)
    v2 = v2
    v5 = append(v5, "")
    if true {
        var v38 string = ("\n\t" + "foo bar")
        v1 -= (v2 ^ 4294967295)
        fmt.Println(v3)
        _ = v38
    } else {
        for fi39 := 0; fi39 < 16; fi39++ {
            fmt.Println(v3, fi39, v1)
            var v40 int32 = 221
            v41 := GenS4{F0: 2147483647, F1: 161}
            _ = fi39
            _ = v40
            _ = v41
        }
        v1 = (^(v2))
    }
    var v42 int32 = 2147483647
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
