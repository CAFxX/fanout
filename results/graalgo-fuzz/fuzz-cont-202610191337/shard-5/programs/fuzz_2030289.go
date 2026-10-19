package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int64 }
type GenS2 struct { F0 int; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint8; F1 string; F2 int32 }
func genFn0(p0 string, p1 uint, p2 bool) bool {
    v7 := []int{-1, 128, 42}
    p0 = ((p0 + "hello") + p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return false
}
func genFn1() string {
    for fi8 := 0; fi8 < 10; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int16 = 0
        v10 -= (^(v10))
        v10 = v10
        _ = fi9
        _ = v10
    }
    var v11 uint32 = 704
    _ = v11
    return ("\n\t" + "foo bar")
}

func main() {
    var v1 uint = 0
    v2 := make(map[string]uint32)
    v3 := []string{"hello", "\n\t", "foo bar"}
    var v4 float64 = (-2.25 * 1e10)
    var v5 uint = 255
    v6 := make(map[string]uint32)
    v12 := GenS1{F0: (!false), F1: 127, F2: -1}
    var v13 uint64 = 186
    v13 = (^(v13))
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 uint8 = 42
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 float64 = 1.5
            _ = fi16
            _ = v17
        }
        switch fi14 {
        case -1:
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
        }
        _ = fi14
        _ = v15
    }
    if (!false) {
        v6[("hello" + ("a" + ""))] = 2
    }
    var v18 int16 = -1
    fmt.Println(v4, v18)
    switch (("hello" + (("hello" + (((("a" + "") + "a") + "") + "\n\t")) + "a")) + "\n\t") {
    case "":
        for fi19 := 0; fi19 < 7; fi19++ {
            v6[""] = 529
            if (((("foo bar" + "foo bar") > "foo bar") || (32767 == (v18 | 32767))) && true) {
                var v20 uint8 = 127
                var v21 int64 = 42
                for fi22 := 0; fi22 < 13; fi22++ {
                    v4 = v4
                    _ = fi22
                }
                _ = v20
                _ = v21
            } else {
                for fi23 := 0; fi23 < 18; fi23++ {
                    switch "" {
                    case "\n\t":
                        var v24 float32 = ((((0.1 * -2.25) + 1.5) * float32(0.0)) - (1.5 - (3.14159 - 0.0)))
                        _ = v24
                    case "":
                        v25 := GenS1{F0: true, F1: 100, F2: 1}
                        _ = v25
                    default:
                        v26 := GenS1{F0: (uint64(v1) < v13), F1: 0, F2: 9223372036854775807}
                        _ = v26
                    }
                    var v27 uint64 = v13
                    var v28 int32 = 1
                    _ = fi23
                    _ = v27
                    _ = v28
                }
                switch "\n\t" {
                case "foo bar":
                    var v29 float32 = 1.5
                    _ = v29
                case "hello":
                    if false {
                        if (false && (127 < 42)) {
                            switch "\n\t" {
                            case "hello":
                                var v30 uint8 = 127
                                _ = v30
                            case "":
                                for fi31 := 0; fi31 < 18; fi31++ {
                                    switch "a" {
                                    case "\n\t":
                                        var v32 uint32 = 42
                                        _ = v32
                                    case "hello":
                                        delete(v6, "\n\t")
                                    }
                                    var v33 int = -1
                                    v34 := GenS1{F0: (!true), F1: 0, F2: 0}
                                    var v35 string = "foo bar"
                                    _ = fi31
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                            case "\n\t":
                                fmt.Println(v18)
                            }
                            var v36 uint64 = v13
                            if false {
                                v37 := GenS2{F0: 0, F1: -1, F2: 0}
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v37
                            } else {
                                var v38 uint8 = 2
                                fmt.Println(v38, v5)
                                for fi39 := 0; fi39 < 15; fi39++ {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                    _ = fi39
                                }
                                _ = v38
                            }
                            _ = v36
                        } else {
                            v40 := v6[""]
                            switch "" {
                            case "":
                                switch "hello" {
                                case "":
                                    v41 := GenS1{F0: (!(v13 != 18446744073709551615)), F1: -7, F2: -36}
                                    _ = v41
                                case "\n\t":
                                    if ((v5 & (v5 ^ v1)) == 18446744073709551615) {
                                        fmt.Println(v13)
                                        if len(v3) > 0 {
                                            v3[0] = ""
                                        }
                                    }
                                case "hello":
                                    if len(v3) > 0 {
                                        v3[0] = ("a" + ("" + "a"))
                                    }
                                }
                            case "\n\t":
                                var v42 uint32 = 375
                                _ = v42
                            default:
                                var v43 int = 127
                                _ = v43
                            }
                            _ = v40
                        }
                    } else {
                        v44 := GenS1{F0: (65535 != uint16(v1)), F1: -128, F2: 42}
                        _ = v44
                    }
                }
            }
            v13 += (v13 >> (((v1 - (v1 % ((18446744073709551615) | 1)))) & 63))
            _ = fi19
        }
    case "\n\t":
        var v45 string = "a"
        _ = v45
    }
    v46 := v2["\n\t"]
    var v47 int32 = -2147483648
    for fi48 := 0; fi48 < 19; fi48++ {
        var v49 int16 = (v18 + v18)
        _ = fi48
        _ = v49
    }
    for fi50 := 0; fi50 < 2; fi50++ {
        v46 -= 439
        v13 += v13
        _ = fi50
    }
    v51 := GenS3{F0: 128, F1: ("a" + "\n\t"), F2: -571}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v51
    _ = v6
}
