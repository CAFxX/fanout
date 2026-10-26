package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0(p0 uint) int8 {
    var v7 uint = (p0 << (((p0 + (^(p0)))) & 63))
    v8 := GenS1{F0: ("foo bar" + "")}
    _ = p0
    _ = v7
    _ = v8
    return -128
}
func genFn1(p0 uint8) string {
    v9 := GenS1{F0: "hello"}
    v10 := []int{-635, -2147483648, 0}
    fmt.Println(p0)
    for fi11 := 0; fi11 < 2; fi11++ {
        p0 += p0
        for fi12 := 0; fi12 < 19; fi12++ {
            p0 = (p0 * (p0 + p0))
            fmt.Println(fi12)
            fmt.Println(fi11, p0)
            _ = len(v10)
            _ = fi12
        }
        v13 := GenS1{F0: ""}
        _ = fi11
        _ = v13
    }
    _ = p0
    _ = v10
    _ = v9
    return ("\n\t" + "hello")
}

func main() {
    var v1 int16 = 903
    var v2 uint64 = 107
    v3 := make(map[string]string)
    var v4 uint8 = 1
    var v5 uint32 = 1
    v6 := make(map[string]uint32)
    v14 := GenS2{F0: 1, F1: (1 == -874), F2: 230}
    for fi15 := 0; fi15 < 17; fi15++ {
        v16 := []int{(fi15 & fi15), -17, -9223372036854775808, (fi15 & (fi15 ^ (-(fi15))))}
        if (157 > 279) {
            if (!true) {
                v17 := GenS2{F0: fi15, F1: false, F2: 128}
                _ = v17
            } else {
                if (v4 != v4) {
                    var v18 string = ("foo bar" + "\n\t")
                    v5 -= (v5 | v5)
                    _ = v18
                }
                var v19 float64 = (1e10 - float64(1e10))
                var v20 string = "a"
                _ = v19
                _ = v20
            }
        } else {
            v16 = append(v16, 753)
        }
        _ = fi15
        _ = v16
    }
    var v21 bool = (v4 == 128)
    v22 := []int{-1, 77, 9223372036854775807}
    var v23 uint8 = v4
    v5 -= 0
    v5 = 2147483647
    fmt.Println(v4, v5)
    fmt.Println(v5)
    if len(v22) > 0 {
        v22[0] = 719
    }
    var v24 float32 = float32((-2.25 + 3.14159))
    v25 := GenS2{F0: -112, F1: v21, F2: 1}
    if len(v22) > 0 {
        v22[0] = -315
    }
    var v26 int = int(v23)
    v27 := GenS2{F0: (v26 & -2147483648), F1: (-9223372036854775808 >= v26), F2: 127}
    if (255 == 255) {
        if ((!(((("foo bar" + (("\n\t" + "a") + (("foo bar" + "\n\t") + (("\n\t" + "") + ("\n\t" + "foo bar"))))) + "foo bar") + "") < ("" + "foo bar"))) || ((3.14159 / 3.14159) >= float64((3.14159 * (v24 - ((1e10 / v24) / 0.1)))))) {
            var v28 uint8 = 42
            v29 := GenS1{F0: ""}
            if (v21 || (!v21)) {
                v23 -= 1
                var v30 int = (-(v26))
                _ = v30
            }
            _ = v28
            _ = v29
        } else {
            _ = len(v6)
            if v21 {
                if v21 {
                    for fi31 := 0; fi31 < 19; fi31++ {
                        for fi32 := 0; fi32 < 19; fi32++ {
                            fmt.Println(v2, v23)
                            _ = fi32
                        }
                        for fi33 := 0; fi33 < 12; fi33++ {
                            if (0.0 <= ((0.0 * (1.5 + (1.5 - ((0.0 + 1e10) * 0.1)))) + 1e10)) {
                                var v34 uint32 = (v5 + (v5 & (v5 ^ v5)))
                                v6[("foo bar" + "\n\t")] = 42
                                fmt.Println(v34)
                                _ = v34
                            }
                            v5 = (v5 & 780)
                            _ = fi33
                        }
                        _ = fi31
                    }
                    _ = len(v6)
                    var v35 float64 = 3.14159
                    _ = v35
                }
                if ((v24 / v24) < v24) {
                    var v36 uint64 = (v2 / ((v2) | 1))
                    fmt.Println(v21, v5, v4)
                    v1 -= (v1 * (v1 << ((42) & 63)))
                    _ = v36
                }
                var v37 uint64 = (v2 >> ((772) & 63))
                _ = v37
            } else {
                var v38 uint16 = 255
                _ = v38
            }
            if (!(1 <= -1)) {
                switch ("a" + "foo bar") {
                case "a":
                    v39 := GenS3{F0: v2, F1: 2}
                    _ = v39
                case "":
                    if true {
                        var v40 int16 = (v1 % ((v1) | 1))
                        if len(v22) > 0 {
                            v22[0] = -9223372036854775808
                        }
                        switch (^(v26)) {
                        case 128:
                            v40 += v40
                        case 628:
                            v5 -= v5
                        case 32767:
                            var v41 int32 = -268
                            _ = v41
                        default:
                            var v42 bool = (!v21)
                            _ = v42
                        }
                        _ = v40
                    } else {
                        v43 := GenS1{F0: "foo bar"}
                        _ = v43
                    }
                }
            }
        }
        v44 := GenS1{F0: ("foo bar" + (("" + "") + "\n\t"))}
        v2 = (v2 + 829)
        _ = v44
    } else {
        for fi45 := 0; fi45 < 20; fi45++ {
            v46 := GenS3{F0: v2, F1: 255}
            _ = fi45
            _ = v46
        }
        v47 := GenS1{F0: "a"}
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
