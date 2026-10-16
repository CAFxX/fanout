package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint; F2 string }
type GenS4 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 int8; F1 int64; F2 int8 }
func genFn0() uint32 {
    var v7 int8 = 74
    fmt.Println(v7)
    var v8 uint64 = 4294967295
    v9 := GenS1{F0: true}
    _ = v7
    _ = v8
    _ = v9
    return 744
}
func genFn1(p0 int8, p1 int8) int {
    p1 = 127
    if false {
        if true {
            v10 := []int{42, -238, -9223372036854775808, int(p0)}
            p0 = p1
            _ = v10
        } else {
            fmt.Println(p0, p1)
            var v11 uint = 2
            fmt.Println(v11)
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 20; fi12++ {
            v13 := []int{51}
            _ = len(v13)
            for fi14 := 0; fi14 < 16; fi14++ {
                p1 += 127
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
        if true {
            if (!(p1 == p1)) {
                var v15 int64 = 823
                var v16 float32 = 1e10
                v15 = v15
                _ = v15
                _ = v16
            }
            v17 := GenS1{F0: true}
            _ = v17
        }
        fmt.Println(p0)
    }
    p0 += (-(p1))
    p0 = 74
    _ = p0
    _ = p1
    return -1
}

func main() {
    v1 := []uint32{}
    v2 := []int{32767, -9223372036854775808, -164}
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 248
    var v5 uint16 = 32767
    v6 := make(map[string]bool)
    var v18 int = 9223372036854775807
    v3[42] = 400
    var v19 uint16 = v5
    switch "a" {
    case "":
        var v20 float64 = (-2.25 / ((3.14159 / -2.25) + 0.0))
        _ = v20
    case "\n\t":
        for fi21 := 0; fi21 < 16; fi21++ {
            switch 127 {
            case -184:
                if ((((1e10 / (0.0 + ((-2.25 + float64((((3.14159 + 3.14159) * 1.5) + 1.5))) / (1e10 / 1e10)))) < 1e10) || false) || (6 >= 87)) {
                    var v22 uint = 18446744073709551615
                    _ = v22
                }
            default:
                if true {
                    v23 := GenS1{F0: false}
                    for fi24 := 0; fi24 < 2; fi24++ {
                        if ((^(v19)) >= (v19 << ((0) & 63))) {
                            v2 = append(v2, v18)
                            if (2 != 1) {
                                v4 = v4
                            } else {
                                fmt.Println(v5, fi24)
                                var v25 int8 = -1
                                fmt.Println(v4)
                                v5 -= 32767
                                _ = v25
                            }
                        } else {
                            var v26 int16 = -1
                            if (((!false) || false) || true) {
                                v27 := GenS2{F0: 0, F1: 1000, F2: 18446744073709551615}
                                var v28 int = (^(fi21))
                                _ = v27
                                _ = v28
                            }
                            _ = v26
                        }
                        v19 = 0
                        v18 -= 4294967295
                        _ = fi24
                    }
                    _ = v23
                } else {
                    for fi29 := 0; fi29 < 2; fi29++ {
                        if false {
                            v18 = (-(v18))
                        }
                        _ = fi29
                    }
                    if true {
                        v1 = append(v1, genFn0())
                        fmt.Println(v19)
                    } else {
                        var v30 int32 = 2147483647
                        v31 := GenS3{F0: 4294967295, F1: 458, F2: ("\n\t" + "hello")}
                        _ = v30
                        _ = v31
                    }
                    v2 = append(v2, (fi21 >> ((42) & 63)))
                }
            }
            for fi32 := 0; fi32 < 13; fi32++ {
                var v33 string = "hello"
                var v34 uint8 = v4
                v5 = (v5 + 65535)
                _ = fi32
                _ = v33
                _ = v34
            }
            if false {
                if (!true) {
                    if true {
                        if (65535 <= 1) {
                            _ = len(v1)
                            var v35 string = ("\n\t" + "\n\t")
                            fmt.Println(v4)
                            _ = v35
                        }
                        var v36 uint64 = 4294967295
                        _ = v36
                    } else {
                        v4 = v4
                    }
                    v37 := GenS1{F0: false}
                    fmt.Println(v5, v5, v5)
                    _ = v37
                } else {
                    switch ("hello" + ("\n\t" + "")) {
                    case "a":
                        var v38 bool = (0 < 66)
                        _ = v38
                    case "foo bar":
                        var v39 int = genFn1(42, 0)
                        _ = v39
                    case "hello":
                        fmt.Println(v18, v18)
                    }
                    for fi40 := 0; fi40 < 3; fi40++ {
                        v41 := GenS3{F0: 876, F1: 255, F2: (("" + "a") + ("hello" + ""))}
                        v42 := GenS4{F0: -1, F1: 2, F2: v4}
                        var v43 uint64 = 59
                        _ = fi40
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                }
                if (false || (534 > 302)) {
                    v18 = fi21
                    v4 += 2
                } else {
                    v19 = 2
                }
            } else {
                v5 = v5
                var v44 float32 = 0.1
                _ = v44
            }
            _ = fi21
        }
    default:
        for fi45 := 0; fi45 < 20; fi45++ {
            if true {
                v46 := GenS5{F0: "\n\t", F1: 42}
                v47 := GenS1{F0: (32767 != 0)}
                if (false || true) {
                    v4 = v4
                    _ = len(v2)
                } else {
                    for fi48 := 0; fi48 < 9; fi48++ {
                        if ((v5 - v5) != v19) {
                            var v49 uint = 92
                            var v50 uint = v49
                            _ = v49
                            _ = v50
                        } else {
                            v51 := GenS5{F0: "foo bar", F1: 42}
                            _ = v51
                        }
                        var v52 uint8 = 2
                        delete(v6, ((("\n\t" + "a") + "a") + ("" + "a")))
                        _ = fi48
                        _ = v52
                    }
                    v5 = v19
                }
                _ = v46
                _ = v47
            }
            _ = fi45
        }
    }
    var v53 string = "foo bar"
    var v54 int8 = -128
    v54 = (v54 | v54)
    v55 := GenS3{F0: 2, F1: 42, F2: v53}
    v56 := GenS6{F0: (v54 | 127), F1: 1, F2: -128}
    var v57 int16 = int16(v19)
    var v58 uint32 = 65535
    v18 = (v18 & (v18 ^ (v18 << ((18446744073709551615) & 63))))
    _ = len(v2)
    fmt.Println(v18, v18, v54)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v6
}
