package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 bool; F1 int64; F2 uint }
type GenS5 struct { F0 uint8; F1 uint; F2 uint16 }
func genFn0(p0 uint32, p1 bool) uint16 {
    var v5 uint16 = 65535
    v6 := GenS1{F0: -127}
    var v7 string = "hello"
    for fi8 := 0; fi8 < 2; fi8++ {
        v5 -= v5
        var v9 int8 = 90
        v5 += (v5 | (v5 & 648))
        _ = fi8
        _ = v9
    }
    if (true && (-128 < -128)) {
        if p1 {
            fmt.Println(v7)
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (v5 ^ v5)
}
func genFn1(p0 int8) int8 {
    var v10 float64 = 3.14159
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := []int{97}
        if (false && (!true)) {
            switch "hello" {
            case "a":
                v13 := GenS2{F0: 1000}
                _ = v13
            case "hello":
                for fi14 := 0; fi14 < 14; fi14++ {
                    var v15 uint64 = 710
                    _ = fi14
                    _ = v15
                }
            }
            if (false && true) {
                var v16 int8 = (p0 / ((42) | 1))
                v10 -= (3.14159 * 0.0)
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = v16
            }
            if false {
                v12 = append(v12, (-(fi11)))
                if (!(-208 == 0)) {
                    var v17 uint8 = 42
                    if true {
                        v10 -= (1.5 + 0.0)
                    } else {
                        v18 := GenS3{F0: p0, F1: 42}
                        if (887 != 2147483647) {
                            v12 = append(v12, (^(fi11)))
                        } else {
                            fmt.Println(v17, fi11)
                            switch ("hello" + ("" + "\n\t")) {
                            case "foo bar":
                                var v19 float64 = float64((float32(v10) / 1e10))
                                _ = v19
                            }
                        }
                        if false {
                            v10 -= 0.0
                            if len(v12) > 0 {
                                v12[0] = (fi11 / (((fi11 * (fi11 >> ((2) & 63)))) | 1))
                            }
                            v17 += (v17 << ((899) & 63))
                        }
                        _ = v18
                    }
                    _ = v17
                } else {
                    var v20 uint8 = 0
                    fmt.Println(p0, v20)
                    _ = v20
                }
            } else {
                if false {
                    if true {
                        if (!(1 > 18446744073709551615)) {
                            var v21 int8 = p0
                            if (0 <= 47) {
                                for fi22 := 0; fi22 < 19; fi22++ {
                                    var v23 int = (-(fi11))
                                    v21 = int8(fi11)
                                    var v24 float64 = -2.25
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                }
                                v25 := GenS2{F0: 0}
                                _ = v25
                            } else {
                                v26 := GenS3{F0: (v21 + (v21 << ((4294967295) & 63))), F1: 2}
                                _ = v26
                            }
                            _ = v21
                        } else {
                            var v27 uint32 = 237
                            var v28 int64 = -187
                            fmt.Println(p0)
                            _ = v27
                            _ = v28
                        }
                        p0 = 127
                        if true {
                            if true {
                                var v29 float64 = v10
                                _ = v29
                            }
                            var v30 int = fi11
                            _ = v30
                        }
                    } else {
                        v10 = (v10 - v10)
                    }
                    switch fi11 {
                    case -91:
                        v12 = append(v12, fi11)
                    case -830:
                        if (698 > 1) {
                            _ = len(v12)
                        }
                    case 941:
                        v10 = -2.25
                    }
                }
                v10 = -2.25
                fmt.Println(fi11, v10)
            }
        } else {
            p0 -= p0
            p0 -= -1
        }
        var v31 int16 = 42
        _ = fi11
        _ = v12
        _ = v31
    }
    _ = p0
    _ = v10
    return 114
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 32768
    v3 := []int{}
    v4 := []string{"\n\t", "foo bar"}
    v2 = 42
    var v32 float32 = 1.5
    v32 = (float32((0.1 - -2.25)) * ((((v32 / float32(-2.25)) * ((3.14159 / v32) * float32(3.14159))) - v32) * v32))
    v33 := GenS4{F0: ((-2147483648 < 1) || true), F1: 60, F2: 4294967295}
    var v34 uint = 4294967295
    for fi35 := 0; fi35 < 8; fi35++ {
        var v36 uint64 = 592
        _ = fi35
        _ = v36
    }
    fmt.Println(v32, v32)
    v2 = 236
    for fi37 := 0; fi37 < 13; fi37++ {
        v38 := GenS5{F0: 255, F1: (v34 / ((18446744073709551615) | 1)), F2: 980}
        v39 := v1["a"]
        for fi40 := 0; fi40 < 15; fi40++ {
            fmt.Println(v2)
            delete(v1, v39)
            _ = fi40
        }
        _ = fi37
        _ = v38
        _ = v39
    }
    if ((!(2 != 127)) && (127 <= 16)) {
        var v41 uint32 = 2
        _ = v41
    }
    v2 -= 2147483647
    var v42 int32 = int32(v2)
    switch "\n\t" {
    case "\n\t":
        var v43 float64 = (1.5 + 1e10)
        _ = v43
    }
    var v44 int64 = -273
    for fi45 := 0; fi45 < 7; fi45++ {
        fmt.Println(fi45, v32, v44)
        switch 590 {
        case -9223372036854775808:
            var v46 uint16 = 927
            _ = v46
        default:
            var v47 uint16 = 1
            _ = v47
        }
        _ = fi45
    }
    if true {
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v42
    _ = v44
}
