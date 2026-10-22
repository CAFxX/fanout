package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 int32; F1 int32; F2 int }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint64) int16 {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint = 1
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int64 = 42
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    v11 := []int{9223372036854775807, 128}
    fmt.Println(p0)
    if true {
        p0 = p0
    } else {
        p0 -= p0
    }
    _ = p0
    _ = v11
    return -552
}

func main() {
    var v1 int8 = 100
    v2 := []string{"a", ""}
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]string)
    var v5 uint16 = 42
    v6 := []string{(("foo bar" + "hello") + ""), "a"}
    fmt.Println(v5)
    var v12 uint32 = 245
    var v13 int64 = 42
    var v14 uint16 = v5
    var v15 bool = false
    var v16 uint16 = v14
    var v17 float32 = -2.25
    switch "foo bar" {
    case "\n\t":
        v16 = (^(v16))
    case "foo bar":
        for fi18 := 0; fi18 < 2; fi18++ {
            switch "hello" {
            case "hello":
                v16 += (v5 | (v16 % ((1) | 1)))
            case "a":
                for fi19 := 0; fi19 < 17; fi19++ {
                    v13 = (^(v13))
                    v20 := GenS1{F0: true, F1: 2, F2: 1}
                    v1 = -65
                    if (!(!(uint64(v14) > 18446744073709551615))) {
                        v21 := GenS1{F0: v15, F1: 2147483647, F2: 18446744073709551615}
                        v15 = (371 == (v14 - (^(v14))))
                        v2 = append(v2, ("hello" + ("a" + "hello")))
                        v16 += 581
                        _ = v21
                    } else {
                        v3[(v12 + 493)] = v12
                        if (v5 > 2) {
                            v13 += (v13 + v13)
                            var v22 uint8 = 0
                            if v15 {
                                v16 = 743
                                _ = len(v6)
                                fmt.Println(v22, v13, v14)
                                if (false || true) {
                                    var v23 string = ("a" + ("hello" + ("foo bar" + "a")))
                                    v22 = (v22 % (((^(v22))) | 1))
                                    v22 = v22
                                    if len(v2) > 0 {
                                        v2[0] = v23
                                    }
                                    if ((((3.14159 - float64(v17)) / float64(v17)) / -2.25) > (0.0 - 1.5)) {
                                        var v24 int64 = 658
                                        v25 := GenS2{F0: -619, F1: 251}
                                        var v26 uint64 = 42
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    } else {
                                        fmt.Println(fi19, v22)
                                        fmt.Println(v14)
                                        v27 := GenS3{F0: (v13 % ((0) | 1)), F1: v15}
                                        _ = v27
                                    }
                                    _ = v23
                                }
                            } else {
                                if (-11 > v1) {
                                    if len(v6) > 0 {
                                        v6[0] = "hello"
                                    }
                                    switch (((("foo bar" + "") + "a") + (("foo bar" + "a") + "hello")) + "\n\t") {
                                    case "\n\t":
                                        v17 = v17
                                    case "hello":
                                        v1 = v1
                                    }
                                    if v15 {
                                        var v28 uint32 = (v12 / ((v12) | 1))
                                        var v29 int32 = 2147483647
                                        delete(v4, 588)
                                        _ = v28
                                        _ = v29
                                    }
                                } else {
                                    v30 := GenS1{F0: v15, F1: (v12 / ((v12) | 1)), F2: 4294967295}
                                    var v31 uint8 = 42
                                    if v15 {
                                        fmt.Println(v16)
                                    }
                                    _ = v30
                                    _ = v31
                                }
                            }
                            _ = v22
                        } else {
                            v32 := GenS4{F0: "a", F1: 4294967295}
                            var v33 uint32 = uint32(v14)
                            switch (fi18 + (fi19 | (-(fi19)))) {
                            case 449:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                v34 := GenS5{F0: 18446744073709551615, F1: 797}
                                _ = v34
                            }
                            _ = v32
                            _ = v33
                        }
                        v5 -= (^(v5))
                    }
                    _ = fi19
                    _ = v20
                }
            default:
                v12 = (v12 * v12)
            }
            v1 = (^(v1))
            for fi35 := 0; fi35 < 2; fi35++ {
                v36 := GenS6{F0: int32(v1), F1: 1, F2: -9223372036854775808}
                var v37 float32 = 0.0
                if v15 {
                    var v38 float64 = 3.14159
                    _ = v38
                }
                var v39 bool = (((1.5 / ((-2.25 * 0.0) - v17)) <= (v17 - 0.1)) && ("hello" > "a"))
                _ = fi35
                _ = v36
                _ = v37
                _ = v39
            }
            _ = fi18
        }
    case "hello":
        if v15 {
            switch "a" {
            case "hello":
                switch (("\n\t" + "foo bar") + "") {
                case "foo bar":
                    var v40 int8 = (v1 % (((v1 >> ((42) & 63))) | 1))
                    _ = v40
                case "hello":
                    var v41 bool = ((("a" + "foo bar") + ("\n\t" + "foo bar")) != (("" + "\n\t") + "hello"))
                    _ = v41
                default:
                    var v42 string = "foo bar"
                    _ = v42
                }
            case "":
                v43 := GenS7{F0: (-(v1)), F1: 42}
                _ = v43
            case "\n\t":
                fmt.Println(v15)
            default:
                if (4294967295 <= 42) {
                    var v44 float64 = (-2.25 / 0.1)
                    for fi45 := 0; fi45 < 9; fi45++ {
                        v46 := GenS8{F0: int32(v1), F1: 122}
                        if (uint8(v5) != 42) {
                            v47 := GenS5{F0: 0, F1: 42}
                            for fi48 := 0; fi48 < 8; fi48++ {
                                v49 := v3[v12]
                                _ = fi48
                                _ = v49
                            }
                            _ = v47
                        } else {
                            switch ((("\n\t" + "") + "foo bar") + "hello") {
                            case "foo bar":
                                v15 = (9223372036854775807 == -1)
                            default:
                                switch "hello" {
                                case "\n\t":
                                    var v50 uint8 = 0
                                    _ = v50
                                default:
                                    fmt.Println(v1, v1, v16)
                                }
                            }
                        }
                        v14 += 0
                        _ = fi45
                        _ = v46
                    }
                    _ = v44
                }
            }
            v14 = v5
        } else {
            for fi51 := 0; fi51 < 18; fi51++ {
                if len(v6) > 0 {
                    v6[0] = "hello"
                }
                _ = fi51
            }
        }
    default:
        var v52 float64 = -2.25
        _ = v52
    }
    switch -1 {
    case -2147483648:
        fmt.Println(v1, v1)
    default:
        fmt.Println(v15)
    }
    v53 := GenS2{F0: int32(v12), F1: 967}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
}
