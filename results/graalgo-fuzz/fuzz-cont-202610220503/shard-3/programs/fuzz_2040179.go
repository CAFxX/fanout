package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 string; F2 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 int64; F2 int }
type GenS6 struct { F0 int16; F1 uint; F2 int }
type GenS7 struct { F0 uint8; F1 bool; F2 int16 }
func genFn0(p0 uint32, p1 string) uint64 {
    fmt.Println(p0, p1)
    v4 := []int{-1, 1, -226}
    v5 := GenS1{F0: 42, F1: ""}
    p1 = "a"
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 899
}

func main() {
    var v1 float64 = (-2.25 * float64(0.1))
    v2 := make(map[uint32]string)
    var v3 int16 = -32768
    fmt.Println(v3, v3)
    for fi6 := 0; fi6 < 7; fi6++ {
        if true {
            for fi7 := 0; fi7 < 18; fi7++ {
                if true {
                    if false {
                        v3 += 1
                        if ("\n\t" <= ("foo bar" + "hello")) {
                            if true {
                                v8 := v2[2147483648]
                                if true {
                                    v1 = 1.5
                                    v9 := GenS1{F0: 835, F1: "\n\t"}
                                    var v10 int64 = 199
                                    _ = v10
                                    _ = v9
                                } else {
                                    v8 = v8
                                }
                                _ = v8
                            }
                        } else {
                            v11 := GenS2{F0: genFn0(887, (("hello" + "") + "foo bar"))}
                            if (false && true) {
                                v12 := GenS2{F0: genFn0(4294967295, "\n\t")}
                                _ = v12
                            }
                            _ = v11
                        }
                        v1 = v1
                        v3 += 42
                        v1 += (3.14159 - (-2.25 * 0.0))
                        var v13 bool = (!false)
                        v14 := []int{fi6, -877, (fi6 * fi6), fi7}
                        v15 := GenS3{F0: true, F1: ("a" + "foo bar"), F2: 255}
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    var v16 int = fi7
                    v17 := GenS4{F0: 255}
                    _ = v16
                    _ = v17
                } else {
                    if (((2147483648 != 4294967295) && false) || false) {
                        if true {
                            var v18 float64 = v1
                            _ = v18
                        } else {
                            v1 -= (v1 + v1)
                        }
                    } else {
                        if (38 >= 42) {
                            v3 = v3
                            v1 -= ((float64((-2.25 + 1.5)) / v1) * ((v1 + v1) - -2.25))
                            var v19 uint64 = 1
                            _ = v19
                        } else {
                            if (199 == 1) {
                                var v20 uint = 2
                                _ = v20
                            } else {
                                v1 -= (v1 + 3.14159)
                                switch fi7 {
                                case -877:
                                    var v21 int = (fi7 + fi6)
                                    _ = v21
                                default:
                                    var v22 uint64 = 1
                                    _ = v22
                                }
                                v23 := v2[65535]
                                _ = v23
                            }
                            delete(v2, 0)
                        }
                        v3 -= (v3 & 1)
                    }
                    v24 := []int{fi7, 42}
                    v25 := v2[287]
                    _ = v24
                    _ = v25
                }
                var v26 int16 = (v3 & v3)
                _ = fi7
                _ = v26
            }
            v27 := v2[42]
            _ = v27
        }
        _ = fi6
    }
    v3 = (^(v3))
    if true {
        var v28 int64 = -242
        v3 = 546
        for fi29 := 0; fi29 < 7; fi29++ {
            if true {
                var v30 float64 = ((0.1 / (v1 * ((v1 / (v1 - (1.5 * 1e10))) / 1.0))) + 1.5)
                v30 = v1
                _ = v30
            }
            _ = fi29
        }
        _ = v28
    } else {
        v31 := []int{128, -9223372036854775808, 124, 42}
        fmt.Println(v3, v1)
        fmt.Println(v1)
        _ = v31
    }
    v32 := GenS5{F0: 42, F1: 9223372036854775807, F2: -2147483648}
    var v33 int64 = -1
    delete(v2, 65535)
    switch "foo bar" {
    case "\n\t":
        v3 -= (v3 % ((-1) | 1))
    case "foo bar":
        v33 = 790
    }
    if true {
        if true {
            for fi34 := 0; fi34 < 12; fi34++ {
                if (!true) {
                    v35 := v2[4294967295]
                    v36 := GenS6{F0: int16(v33), F1: 246, F2: 127}
                    v3 -= v3
                    _ = v35
                    _ = v36
                }
                switch "a" {
                case "a":
                    delete(v2, 4294967295)
                }
                v1 -= (v1 - 3.14159)
                _ = fi34
            }
            v1 = (((v1 * (v1 * (3.14159 / v1))) / float64((0.0 - -2.25))) * (v1 / (-2.25 * 1e10)))
        } else {
            v37 := GenS7{F0: 40, F1: ((!true) && ((v3 / ((v3) | 1)) != v3)), F2: 1000}
            var v38 uint8 = 235
            _ = v37
            _ = v38
        }
        var v39 string = ""
        if true {
            var v40 bool = true
            for fi41 := 0; fi41 < 17; fi41++ {
                v40 = v40
                _ = fi41
            }
            _ = v40
        }
        _ = v39
    }
    switch 42 {
    case -232:
        var v42 bool = (0 >= 1)
        _ = v42
    default:
        var v43 uint8 = 239
        _ = v43
    }
    var v44 uint64 = 4294967295
    v1 = (v1 + v1)
    for fi45 := 0; fi45 < 9; fi45++ {
        if true {
            for fi46 := 0; fi46 < 15; fi46++ {
                var v47 uint8 = 1
                _ = fi46
                _ = v47
            }
            for fi48 := 0; fi48 < 18; fi48++ {
                delete(v2, 610)
                if (-2.25 >= 1e10) {
                    switch "" {
                    case "hello":
                        fmt.Println(fi45, fi45, v3)
                    }
                }
                v44 += 4294967295
                var v49 uint64 = v44
                v50 := GenS7{F0: 128, F1: false, F2: -814}
                _ = fi48
                _ = v49
                _ = v50
            }
            var v51 uint8 = 2
            _ = v51
        }
        _ = fi45
    }
    v1 += v1
    v3 = v3
    fmt.Println(v44)
    var v52 float32 = 1e10
    v44 = (v44 & 18446744073709551615)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v44:", v44)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v44
    _ = v52
}
