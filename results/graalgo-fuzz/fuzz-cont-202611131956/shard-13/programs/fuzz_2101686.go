package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int16 }
type GenS2 struct { F0 int32; F1 uint8; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 int64; F2 uint32 }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 0.0
    var v3 bool = true
    var v4 bool = ((!(true && (!(v3 || (100 < 42))))) && (-1 == -77))
    v1[42] = 1
    for fi5 := 0; fi5 < 10; fi5++ {
        switch ("foo bar" + "\n\t") {
        case "hello":
            _ = len(v1)
        }
        if v4 {
            for fi6 := 0; fi6 < 4; fi6++ {
                switch ("\n\t" + "hello") {
                case "foo bar":
                    fmt.Println(v3, fi5)
                }
                if false {
                    if (1.5 != 1.5) {
                        var v7 int = fi6
                        if (2 < 0) {
                            v3 = v3
                            var v8 uint16 = 1
                            _ = v8
                        }
                        var v9 float64 = (3.14159 / 1e10)
                        _ = v7
                        _ = v9
                    } else {
                        v4 = v3
                        v2 += v2
                        if v3 {
                            if false {
                                v10 := GenS1{F0: 24, F1: fi6, F2: -32768}
                                switch fi5 {
                                case 2147483647:
                                    _ = len(v1)
                                case 9223372036854775807:
                                    v3 = ((!false) && v4)
                                }
                                if (!(false || true)) {
                                    v11 := GenS2{F0: -2147483648, F1: 2, F2: v3}
                                    if v4 {
                                        v2 -= 0.1
                                        v2 -= v2
                                        v12 := []int{(fi5 | (fi5 * fi5))}
                                        _ = v12
                                    } else {
                                        v2 = 0.0
                                        v3 = (127 > 2)
                                    }
                                    v3 = true
                                    _ = v11
                                } else {
                                    var v13 uint = 42
                                    _ = v13
                                }
                                _ = v10
                            }
                        } else {
                            var v14 string = ""
                            if (427 > -1) {
                                switch fi6 {
                                case 0:
                                    fmt.Println(v2)
                                case -232:
                                    v2 = float32(1e10)
                                }
                                v2 += v2
                                switch ((("foo bar" + ("foo bar" + "\n\t")) + (v14 + (((((((("hello" + v14) + "a") + (("" + "") + ((("hello" + v14) + ((((v14 + (v14 + (("foo bar" + v14) + "foo bar"))) + "\n\t") + "a") + v14)) + ((v14 + (v14 + "\n\t")) + "a")))) + "hello") + ((v14 + "\n\t") + ((v14 + (v14 + ((((v14 + "") + (v14 + v14)) + "a") + ""))) + "a"))) + v14) + v14) + v14))) + v14) {
                                case "\n\t":
                                    v2 = (v2 - v2)
                                    var v15 float32 = (float32(-2.25) * (3.14159 * 1.5))
                                    var v16 int8 = -126
                                    _ = v15
                                    _ = v16
                                }
                            }
                            var v17 int64 = -9223372036854775808
                            v18 := GenS1{F0: 1, F1: int(v17), F2: -1}
                            delete(v1, 42)
                            _ = v14
                            _ = v17
                            _ = v18
                        }
                    }
                    v19 := []int{}
                    if len(v19) > 0 {
                        v19[0] = (fi6 - (fi5 & fi5))
                    }
                    _ = v19
                } else {
                    var v20 uint16 = 0
                    var v21 int32 = 957
                    _ = v20
                    _ = v21
                }
                delete(v1, 2)
                _ = fi6
            }
            if (v4 || v3) {
                for fi22 := 0; fi22 < 7; fi22++ {
                    if v4 {
                        fmt.Println(fi22, fi22)
                    } else {
                        fmt.Println(v3, v4)
                    }
                    fmt.Println(fi5, v4)
                    _ = fi22
                }
                switch 9223372036854775807 {
                case 422:
                    if false {
                        v23 := []int{fi5}
                        var v24 int64 = 1
                        fmt.Println(fi5, fi5)
                        _ = v23
                        _ = v24
                    }
                default:
                    v25 := GenS1{F0: 42, F1: 32768, F2: -32768}
                    _ = v25
                }
            } else {
                v26 := []int{(fi5 ^ (fi5 / ((128) | 1)))}
                switch fi5 {
                case 128:
                    fmt.Println(v3, v4)
                default:
                    v26 = append(v26, fi5)
                }
                _ = v26
            }
            v4 = v3
        } else {
            for fi27 := 0; fi27 < 11; fi27++ {
                fmt.Println(fi27)
                v28 := GenS3{F0: -591}
                _ = len(v1)
                _ = fi27
                _ = v28
            }
            if (("a" >= ((("a" + "\n\t") + "\n\t") + "")) || (2147483647 > 65535)) {
                if v4 {
                    var v29 int32 = -2147483648
                    _ = v29
                }
                switch "hello" {
                case "hello":
                    v30 := []int{(fi5 >> ((18446744073709551615) & 63)), (fi5 - (^(fi5))), (fi5 * fi5), 42}
                    _ = v30
                case "":
                    v3 = (-128 != -121)
                }
            }
            v31 := GenS1{F0: 128, F1: fi5, F2: -1}
            _ = v31
        }
        var v32 uint8 = 1
        _ = fi5
        _ = v32
    }
    v33 := GenS1{F0: 43, F1: 0, F2: 42}
    if (0 != 0) {
        v34 := GenS4{F0: ("" + "\n\t"), F1: 0, F2: 65535}
        v35 := GenS2{F0: -1, F1: 209, F2: (!false)}
        _ = v34
        _ = v35
    } else {
        for fi36 := 0; fi36 < 2; fi36++ {
            var v37 int16 = -32768
            switch fi36 {
            case 4294967295:
                v38 := GenS2{F0: 445, F1: 86, F2: (!(539 <= 42))}
                _ = v38
            case -1:
                if (v37 < 32767) {
                    v39 := GenS5{F0: 220}
                    _ = v39
                }
            default:
                var v40 uint64 = 2
                _ = v40
            }
            var v41 uint8 = 2
            _ = fi36
            _ = v37
            _ = v41
        }
    }
    v3 = false
    for fi42 := 0; fi42 < 11; fi42++ {
        v3 = (!v4)
        var v43 int64 = -1
        if v3 {
            var v44 uint64 = 2
            if v4 {
                switch ("hello" + (("hello" + "") + "hello")) {
                case "foo bar":
                    v44 = v44
                case "hello":
                    v45 := GenS3{F0: (fi42 + fi42)}
                    _ = v45
                default:
                    fmt.Println(v43, v4)
                }
            }
            _ = v44
        } else {
            v43 -= (v43 + (-(v43)))
            var v46 int16 = 1
            _ = v46
        }
        _ = fi42
        _ = v43
    }
    var v47 float32 = float32(-2.25)
    fmt.Println(v2, v47, v4)
    if false {
        for fi48 := 0; fi48 < 15; fi48++ {
            var v49 int32 = 0
            v50 := []int{fi48, fi48}
            _ = fi48
            _ = v49
            _ = v50
        }
        v51 := []int{-662, 32, -673, -181}
        v52 := GenS5{F0: 16}
        _ = v51
        _ = v52
    } else {
        var v53 int32 = -2147483648
        _ = v53
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v47
}
