package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 bool; F1 uint64; F2 string }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 bool; F1 bool }
func genFn0() string {
    var v6 uint64 = 4294967295
    v6 += 2
    _ = v6
    return "a"
}

func main() {
    v1 := []int{-2147483648, -2147483648, -9223372036854775808, 128}
    v2 := []int{1, 1, 458, -759}
    v3 := make(map[uint32]string)
    var v4 float32 = 3.14159
    v5 := []string{"hello"}
    for fi7 := 0; fi7 < 1; fi7++ {
        v4 = v4
        for fi8 := 0; fi8 < 2; fi8++ {
            v4 -= v4
            v4 = 0.0
            _ = fi8
        }
        v9 := GenS1{F0: 28, F1: fi7}
        _ = fi7
        _ = v9
    }
    v4 -= float32(3.14159)
    var v10 int64 = 9223372036854775807
    v11 := GenS1{F0: 1, F1: -2147483648}
    v12 := GenS2{F0: true, F1: 42, F2: ""}
    if (2147483647 != -952) {
        v10 -= v10
        v13 := GenS2{F0: (-111 >= 23), F1: 42, F2: ""}
        if (!((v10 % ((v10) | 1)) >= v10)) {
            var v14 int64 = v10
            var v15 int64 = v14
            v16 := GenS2{F0: true, F1: 18446744073709551615, F2: ((("\n\t" + ("\n\t" + ("a" + "hello"))) + ("foo bar" + (("" + (("a" + ("foo bar" + "foo bar")) + ("a" + "foo bar"))) + ""))) + "\n\t")}
            _ = v14
            _ = v15
            _ = v16
        } else {
            if false {
                v17 := GenS1{F0: 1, F1: 0}
                v18 := GenS3{F0: 413, F1: 78}
                if false {
                    v4 += 0.0
                    var v19 int32 = 789
                    v10 = -9223372036854775808
                    _ = v19
                } else {
                    var v20 int16 = 1
                    for fi21 := 0; fi21 < 15; fi21++ {
                        _ = len(v5)
                        var v22 uint = 42
                        var v23 uint64 = 0
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    if true {
                        if ((v20 + 1000) <= (v20 << ((255) & 63))) {
                            delete(v3, 42)
                        }
                        v3[65535] = ""
                    }
                    _ = v20
                }
                _ = v17
                _ = v18
            } else {
                if ((v10 << ((2) & 63)) > (v10 ^ (-(v10)))) {
                    v10 = 1
                    if false {
                        if (!((false || false) || (1 != 255))) {
                            if (0 >= -32768) {
                                switch "\n\t" {
                                case "\n\t":
                                    v10 += 9223372036854775807
                                case "hello":
                                    switch 128 {
                                    case -9223372036854775808:
                                        v3[708] = "a"
                                    case 0:
                                        v2 = append(v2, -9223372036854775808)
                                    case 4294967295:
                                        delete(v3, 4294967295)
                                    default:
                                        if false {
                                            v10 = (v10 & v10)
                                            v10 = v10
                                        }
                                    }
                                case "foo bar":
                                    for fi24 := 0; fi24 < 14; fi24++ {
                                        fmt.Println(v4, v10)
                                        if true {
                                            var v25 uint8 = 201
                                            _ = v25
                                        } else {
                                            for fi26 := 0; fi26 < 14; fi26++ {
                                                v3[2147483647] = ("\n\t" + "foo bar")
                                                if (true && false) {
                                                    v10 -= 1
                                                    var v27 int16 = -32768
                                                    _ = len(v5)
                                                    _ = v27
                                                }
                                                _ = fi26
                                            }
                                            v10 += 362
                                        }
                                        _ = fi24
                                    }
                                }
                                if len(v1) > 0 {
                                    v1[0] = -9223372036854775808
                                }
                                var v28 uint = 4294967295
                                _ = v28
                            } else {
                                var v29 uint = 65535
                                _ = v29
                            }
                            if true {
                                var v30 int = 82
                                fmt.Println(v30, v4)
                                _ = v30
                            } else {
                                var v31 uint16 = 32767
                                for fi32 := 0; fi32 < 1; fi32++ {
                                    for fi33 := 0; fi33 < 2; fi33++ {
                                        _ = len(v5)
                                        switch "" {
                                        case "foo bar":
                                            fmt.Println(v4)
                                        }
                                        _ = fi33
                                    }
                                    switch "foo bar" {
                                    case "a":
                                        if (3.14159 < float64(v4)) {
                                            if (-1 < 0) {
                                                v4 = 0.1
                                                if false {
                                                    var v34 uint16 = (v31 % (((v31 / ((v31) | 1))) | 1))
                                                    if true {
                                                        for fi35 := 0; fi35 < 12; fi35++ {
                                                            fmt.Println(fi32, v10)
                                                            _ = len(v5)
                                                            _ = fi35
                                                        }
                                                        v36 := GenS2{F0: (!(4294967295 >= 65535)), F1: 306, F2: "foo bar"}
                                                        _ = v36
                                                    } else {
                                                        v10 += v10
                                                        fmt.Println(v34)
                                                    }
                                                    _ = v34
                                                }
                                            } else {
                                                v37 := GenS2{F0: (("" + "") >= "\n\t"), F1: 0, F2: ("hello" + ("foo bar" + (("" + ("" + "foo bar")) + "foo bar")))}
                                                v31 -= 1
                                                v38 := GenS2{F0: (!((v31 > v31) || (!((0 != 18446744073709551615) && false)))), F1: 2, F2: "foo bar"}
                                                _ = v37
                                                _ = v38
                                            }
                                            for fi39 := 0; fi39 < 16; fi39++ {
                                                v40 := GenS1{F0: 232, F1: (fi32 & fi39)}
                                                _ = len(v3)
                                                v41 := GenS2{F0: true, F1: 0, F2: "\n\t"}
                                                _ = fi39
                                                _ = v40
                                                _ = v41
                                            }
                                            var v42 float64 = float64(0.1)
                                            _ = v42
                                        } else {
                                            var v43 uint32 = 1
                                            var v44 int16 = 1
                                            v45 := GenS4{F0: false, F1: false}
                                            _ = v43
                                            _ = v44
                                            _ = v45
                                        }
                                    case "\n\t":
                                        v10 = (v10 & v10)
                                    default:
                                        var v46 int = 836
                                        _ = v46
                                    }
                                    for fi47 := 0; fi47 < 19; fi47++ {
                                        var v48 uint64 = 1
                                        _ = len(v3)
                                        var v49 bool = ("hello" <= "")
                                        _ = fi47
                                        _ = v48
                                        _ = v49
                                    }
                                    _ = fi32
                                }
                                _ = v31
                            }
                        }
                        fmt.Println(v10, v10)
                    } else {
                        for fi50 := 0; fi50 < 8; fi50++ {
                            for fi51 := 0; fi51 < 4; fi51++ {
                                var v52 uint16 = 65535
                                var v53 int64 = v10
                                delete(v3, 93)
                                _ = fi51
                                _ = v52
                                _ = v53
                            }
                            if true {
                                var v54 uint16 = 32767
                                fmt.Println(v10, v54)
                                v55 := GenS4{F0: true, F1: true}
                                _ = v54
                                _ = v55
                            } else {
                                v4 = ((v4 - -2.25) + v4)
                            }
                            _ = len(v3)
                            _ = fi50
                        }
                    }
                    v10 = 0
                } else {
                    var v56 uint = 2
                    _ = v56
                }
                v2 = append(v2, 692)
            }
        }
        _ = v13
    }
    fmt.Println(v10)
    v57 := GenS3{F0: 42, F1: -128}
    v10 += v10
    v58 := GenS4{F0: (!(1 < 86)), F1: false}
    fmt.Println(v10, v4)
    var v59 float32 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v59
}
