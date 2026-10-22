package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 int8; F1 int8; F2 int8 }
type GenS7 struct { F0 int; F1 int16; F2 int64 }

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{}
    var v3 int8 = 100
    fmt.Println(v3)
    _ = len(v2)
    for fi4 := 0; fi4 < 17; fi4++ {
        v5 := GenS1{F0: 721, F1: 18446744073709551615}
        _ = fi4
        _ = v5
    }
    v3 = v3
    var v6 float64 = (1.5 - 0.0)
    var v7 uint = 1
    v6 = (((v6 / v6) - v6) * 1.5)
    var v8 uint64 = 990
    v6 = float64(3.14159)
    if (-837 >= -32768) {
        if (0 == 1) {
            for fi9 := 0; fi9 < 1; fi9++ {
                if true {
                    v6 = v6
                } else {
                    var v10 int8 = v3
                    _ = v10
                }
                var v11 int = 2147483647
                _ = fi9
                _ = v11
            }
        } else {
            _ = len(v1)
            switch 9223372036854775807 {
            case 1:
                for fi12 := 0; fi12 < 9; fi12++ {
                    delete(v1, (("\n\t" + (("" + "hello") + ("foo bar" + ""))) + "hello"))
                    v8 = v8
                    v8 = v8
                    _ = fi12
                }
            case 42:
                v8 = 307
            default:
                v13 := GenS1{F0: 255, F1: 148}
                _ = v13
            }
            fmt.Println(v6, v8)
        }
        v14 := GenS2{F0: (v7 / (((v7 / (((v7 - (v7 ^ v7))) | 1))) | 1))}
        _ = v14
    } else {
        v1[("hello" + ("\n\t" + ""))] = ("foo bar" + ("foo bar" + "a"))
        for fi15 := 0; fi15 < 11; fi15++ {
            if len(v2) > 0 {
                v2[0] = uint32(v7)
            }
            fmt.Println(v8)
            _ = fi15
        }
        if true {
            for fi16 := 0; fi16 < 16; fi16++ {
                v7 = (^(v7))
                switch fi16 {
                case 4294967295:
                    fmt.Println(fi16, fi16)
                default:
                    var v17 float64 = v6
                    _ = v17
                }
                _ = fi16
            }
            for fi18 := 0; fi18 < 11; fi18++ {
                if true {
                    var v19 int16 = -868
                    if true {
                        var v20 bool = ((-(v3)) == v3)
                        if (v7 < (^(v7))) {
                            var v21 int64 = 9223372036854775807
                            _ = v21
                        }
                        _ = v20
                    }
                    _ = v19
                }
                _ = fi18
            }
            v6 += (v6 / float64(3.14159))
        }
    }
    v6 -= ((v6 * (0.1 - 0.1)) * -2.25)
    switch "hello" {
    case "hello":
        for fi22 := 0; fi22 < 2; fi22++ {
            for fi23 := 0; fi23 < 1; fi23++ {
                var v24 bool = true
                var v25 int16 = 169
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi22
        }
    case "foo bar":
        if false {
            fmt.Println(v6)
        }
    default:
        var v26 uint16 = 1
        _ = v26
    }
    delete(v1, (("\n\t" + "a") + ("\n\t" + (("a" + "hello") + "hello"))))
    var v27 bool = false
    v1[((("foo bar" + (("\n\t" + "") + "foo bar")) + ((("" + "a") + "a") + "\n\t")) + "")] = (("" + "\n\t") + ((("\n\t" + "hello") + "hello") + ""))
    if ((v7 | 83) == v7) {
        v28 := GenS3{F0: -32768}
        v27 = (!false)
        _ = v28
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        _ = len(v1)
        if ((false && ((-1 >= 0) && (!(1 > -986)))) || (((v6 - v6) <= v6) || (false || false))) {
            v30 := GenS4{F0: v3, F1: 9223372036854775807}
            _ = v30
        } else {
            switch fi29 {
            case -772:
                v7 = (v7 / ((2) | 1))
            case 0:
                v31 := GenS4{F0: v3, F1: -472}
                _ = v31
            case 17:
                for fi32 := 0; fi32 < 9; fi32++ {
                    v33 := GenS4{F0: v3, F1: int64(v7)}
                    var v34 float64 = v6
                    var v35 float64 = 0.0
                    _ = fi32
                    _ = v33
                    _ = v34
                    _ = v35
                }
            default:
                var v36 float64 = 0.1
                _ = v36
            }
            if v27 {
                v3 -= int8(fi29)
                if v27 {
                    if true {
                        v37 := GenS5{F0: 42, F1: v7}
                        v38 := GenS6{F0: (v3 + 127), F1: (v3 & (v3 >> ((v7) & 63))), F2: (v3 % ((-48) | 1))}
                        _ = v37
                        _ = v38
                    } else {
                        for fi39 := 0; fi39 < 11; fi39++ {
                            var v40 uint32 = 2147483647
                            _ = fi39
                            _ = v40
                        }
                    }
                    if (0.1 >= float64(0.1)) {
                        v41 := GenS6{F0: 1, F1: v3, F2: 127}
                        _ = v41
                    }
                }
                v42 := GenS7{F0: (^(fi29)), F1: -32768, F2: -485}
                _ = v42
            } else {
                var v43 string = (("" + "hello") + ("hello" + "a"))
                v44 := GenS1{F0: v7, F1: (v7 / (((v7 >> ((v7) & 63))) | 1))}
                _ = v43
                _ = v44
            }
            if false {
                var v45 string = ("a" + (("\n\t" + "hello") + ("hello" + "hello")))
                for fi46 := 0; fi46 < 9; fi46++ {
                    if (true || false) {
                        if v27 {
                            var v47 uint16 = 255
                            if (((91 != 127) || false) || (-191 == -9223372036854775808)) {
                                v8 -= (v8 / (((v8 % (((v8 >> ((v7) & 63))) | 1))) | 1))
                                v48 := GenS1{F0: v7, F1: v7}
                                switch v45 {
                                case "":
                                    var v49 int16 = 32767
                                    _ = v49
                                case "a":
                                    v47 = (v47 & 255)
                                default:
                                    v3 -= int8(v7)
                                }
                                _ = v48
                            } else {
                                v2 = append(v2, 648)
                                var v50 bool = (!v27)
                                _ = v50
                            }
                            var v51 float64 = ((v6 / v6) * (v6 / float64(-2.25)))
                            var v52 int = fi29
                            var v53 int16 = -32768
                            _ = v47
                            _ = v51
                            _ = v52
                            _ = v53
                        } else {
                            v6 = ((0.0 + v6) - -2.25)
                        }
                    }
                    _ = fi46
                }
                _ = v45
            }
        }
        if len(v2) > 0 {
            v2[0] = 497
        }
        _ = fi29
    }
    v27 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
