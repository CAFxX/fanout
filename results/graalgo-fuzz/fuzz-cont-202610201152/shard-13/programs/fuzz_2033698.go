package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int8; F1 int32; F2 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 uint64; F2 int16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint; F1 int32; F2 int32 }
type GenS10 struct { F0 int16; F1 string; F2 int64 }
type GenS11 struct { F0 bool }
func genFn0(p0 uint8, p1 int16) int8 {
    fmt.Println(p1, p0)
    p1 = (^(p1))
    v6 := GenS1{F0: (p1 - p1), F1: 0, F2: uint16(p0)}
    p0 -= (p0 & p0)
    switch -9223372036854775808 {
    case -9223372036854775808:
        switch -2147483648 {
        case 608:
            var v7 int16 = p1
            _ = v7
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 100
}

func main() {
    var v1 bool = true
    v2 := make(map[string]int)
    var v3 uint16 = 42
    v4 := make(map[int]int)
    var v5 int = -40
    var v8 int64 = 0
    v9 := []int{(v5 & 1), v5, v5}
    for fi10 := 0; fi10 < 1; fi10++ {
        v3 = (v3 ^ (v3 | v3))
        switch v5 {
        case 32767:
            if true {
                if len(v9) > 0 {
                    v9[0] = (v5 / ((fi10) | 1))
                }
                var v11 string = "hello"
                _ = v11
            }
        }
        switch "foo bar" {
        case "a":
            switch -2147483648 {
            case 32768:
                v3 = v3
            }
        case "":
            if (42 > 32767) {
                for fi12 := 0; fi12 < 19; fi12++ {
                    v1 = v1
                    if ((1e10 + (1e10 / 1e10)) < (((3.14159 + 0.0) - float64(0.1)) + 0.1)) {
                        var v13 uint = 1
                        var v14 string = "\n\t"
                        v15 := GenS1{F0: 0, F1: 1, F2: 1}
                        v16 := GenS1{F0: 32767, F1: 1, F2: (v3 - (v3 / (((v3 & (^(v3)))) | 1)))}
                        var v17 int64 = 9223372036854775807
                        _ = v13
                        _ = v14
                        _ = v15
                        _ = v16
                        _ = v17
                    } else {
                        v1 = true
                        v18 := GenS2{F0: 0, F1: ((!true) || v1)}
                        if v1 {
                            v19 := GenS3{F0: fi12, F1: "\n\t"}
                            if (!false) {
                                var v20 float32 = 0.0
                                v20 = float32(3.14159)
                                switch v5 {
                                case 9223372036854775807:
                                    if ((-50 >= -47) || (v20 != 0.0)) {
                                        if v1 {
                                            if ((false || (-1 > -28)) && v1) {
                                                var v21 uint = 18446744073709551615
                                                _ = v21
                                            }
                                            if (2 < 978) {
                                                v20 -= float32(0.0)
                                            }
                                            fmt.Println(v3, v5, v1)
                                        }
                                        v5 = (fi12 << ((4294967295) & 63))
                                    }
                                case 32768:
                                    switch "foo bar" {
                                    case "a":
                                        if v1 {
                                            fmt.Println(v3)
                                            var v22 uint32 = 2147483647
                                            var v23 uint = 65535
                                            _ = v22
                                            _ = v23
                                        } else {
                                            if false {
                                                switch (fi12 | (-(fi10))) {
                                                case 32767:
                                                    fmt.Println(v8, fi12)
                                                case 128:
                                                    v24 := GenS4{F0: 100, F1: 0, F2: (v8 | (v8 ^ (-(v8))))}
                                                    _ = v24
                                                case -1:
                                                    v25 := v4[fi12]
                                                    _ = v25
                                                default:
                                                    switch "" {
                                                    case "":
                                                        v26 := GenS1{F0: 1000, F1: 42, F2: v3}
                                                        var v27 uint32 = 0
                                                        v28 := GenS2{F0: 0, F1: v1}
                                                        _ = v26
                                                        _ = v27
                                                        _ = v28
                                                    case "hello":
                                                        if (0 < 15) {
                                                            fmt.Println(fi10)
                                                            v29 := GenS5{F0: 42}
                                                            v30 := GenS2{F0: 0, F1: v1}
                                                            _ = v29
                                                            _ = v30
                                                        }
                                                    case "foo bar":
                                                        var v31 uint16 = (v3 % ((0) | 1))
                                                        _ = v31
                                                    default:
                                                        v8 = (^(v8))
                                                    }
                                                }
                                                v32 := GenS5{F0: 65535}
                                                _ = v32
                                            }
                                            if ((!(false && v1)) && ((v3 >> ((1) & 63)) <= (v3 ^ v3))) {
                                                var v33 int = fi12
                                                _ = len(v9)
                                                if v1 {
                                                    fmt.Println(v3, v5, fi10)
                                                    if (!((!v1) || false)) {
                                                        var v34 int64 = v8
                                                        v34 = int64(v5)
                                                        _ = v34
                                                    }
                                                }
                                                _ = v33
                                            }
                                        }
                                    case "\n\t":
                                        fmt.Println(v3)
                                    case "foo bar":
                                        var v35 int64 = 0
                                        _ = v35
                                    default:
                                        var v36 string = ""
                                        _ = v36
                                    }
                                default:
                                    v3 -= v3
                                }
                                _ = v20
                            } else {
                                delete(v4, v5)
                                _ = len(v9)
                            }
                            _ = v19
                        } else {
                            v5 += 32768
                        }
                        _ = v18
                    }
                    v37 := GenS6{F0: 201}
                    _ = fi12
                    _ = v37
                }
                switch 1 {
                case 575:
                    for fi38 := 0; fi38 < 15; fi38++ {
                        var v39 uint64 = 1
                        if (float64((-2.25 - ((1e10 / ((1.5 * (((((-2.25 + 3.14159) + 1e10) + ((0.1 - 1e10) / (((3.14159 / 1.5) * ((1e10 / ((3.14159 / -2.25) / ((1.5 + 3.14159) + float32(0.1)))) + (1e10 / -2.25))) - (0.1 - ((0.0 - -2.25) / 1e10))))) * 0.1) * float32(float64((float32((3.14159 * ((((3.14159 / 1.0) * float64((-2.25 - 0.0))) / 1.0) + -2.25))) / 3.14159))))) / -2.25)) * -2.25))) < 1e10) {
                            fmt.Println(v3, v8, v8)
                            v9 = append(v9, -1)
                        } else {
                            v40 := GenS5{F0: uint(v3)}
                            v41 := GenS5{F0: 2}
                            switch 9223372036854775807 {
                            case 128:
                                if (42 == 255) {
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v42 uint = 4294967295
                                        _ = v42
                                    case "a":
                                        if v1 {
                                            var v43 int8 = 23
                                            _ = v43
                                        } else {
                                            var v44 float32 = 0.1
                                            _ = v44
                                        }
                                    }
                                    v3 += 686
                                    v5 -= (^(v5))
                                }
                            case 127:
                                var v45 uint32 = 951
                                _ = v45
                            case 4294967295:
                                var v46 int16 = 42
                                _ = v46
                            }
                            _ = v40
                            _ = v41
                        }
                        v8 = v8
                        _ = fi38
                        _ = v39
                    }
                case 42:
                    if ((3.14159 * 3.14159) > -2.25) {
                        v47 := v4[32767]
                        _ = v47
                    } else {
                        v1 = v1
                    }
                }
                v48 := GenS2{F0: 187, F1: (!v1)}
                _ = v48
            } else {
                if true {
                    v2["\n\t"] = v5
                    fmt.Println(v5, v3)
                    var v49 string = "a"
                    _ = v49
                } else {
                    v50 := GenS6{F0: 71}
                    _ = v50
                }
            }
        case "hello":
            var v51 int32 = 1
            _ = v51
        default:
            _ = len(v2)
        }
        _ = fi10
    }
    var v52 int64 = v8
    var v53 bool = false
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    v54 := GenS7{F0: v8, F1: 18446744073709551615, F2: 32767}
    if ((241 < 127) || (4294967295 <= 1)) {
        var v55 uint8 = 42
        if v1 {
            for fi56 := 0; fi56 < 17; fi56++ {
                var v57 bool = ((120 > -1) && v1)
                _ = fi56
                _ = v57
            }
            _ = len(v4)
            switch v5 {
            case 42:
                v55 += 42
            case 110:
                if ((v55 / (((v55 - (v55 & v55))) | 1)) <= (^(v55))) {
                    var v58 uint8 = 0
                    v59 := GenS8{F0: ((329 < 255) || v1)}
                    _ = v58
                    _ = v59
                } else {
                    v60 := v2["a"]
                    _ = v60
                }
            default:
                for fi61 := 0; fi61 < 12; fi61++ {
                    v62 := GenS9{F0: 2, F1: int32(v55), F2: -194}
                    if len(v9) > 0 {
                        v9[0] = (fi61 | (^(fi61)))
                    }
                    _ = fi61
                    _ = v62
                }
            }
        } else {
            v1 = v53
            if (int(v55) < (v5 % (((v5 >> ((283) & 63))) | 1))) {
                v63 := GenS3{F0: (v5 * 2147483647), F1: ("a" + ("\n\t" + "foo bar"))}
                if v1 {
                    v64 := GenS6{F0: 128}
                    _ = v64
                }
                var v65 int32 = -2147483648
                _ = v63
                _ = v65
            }
        }
        for fi66 := 0; fi66 < 4; fi66++ {
            switch (("hello" + "a") + ("hello" + "hello")) {
            case "a":
                v67 := GenS10{F0: 42, F1: "foo bar", F2: (-(v8))}
                _ = v67
            case "":
                v68 := GenS2{F0: 18446744073709551615, F1: v53}
                _ = v68
            case "hello":
                var v69 int = fi66
                _ = v69
            default:
                if (((v8 == 0) && (!(v53 && (!(862 == 2147483647))))) || (900 > 18446744073709551615)) {
                    for fi70 := 0; fi70 < 20; fi70++ {
                        v5 = fi66
                        var v71 uint8 = (v55 << ((2) & 63))
                        _ = fi70
                        _ = v71
                    }
                } else {
                    var v72 uint16 = 32767
                    if v53 {
                        v52 = v52
                    } else {
                        v73 := GenS11{F0: (float64((float32(-2.25) + 0.1)) != (1e10 - (0.1 + 3.14159)))}
                        fmt.Println(fi66, v53)
                        _ = v73
                    }
                    var v74 int32 = -1
                    _ = v72
                    _ = v74
                }
            }
            v55 -= 54
            _ = fi66
        }
        _ = v55
    }
    v75 := GenS1{F0: -276, F1: 2147483647, F2: v3}
    v9 = append(v9, (^(v5)))
    var v76 bool = (!true)
    fmt.Println(v5, v53)
    var v77 int16 = 0
    if (255 == 128) {
        v78 := GenS1{F0: v77, F1: 1, F2: 290}
        for fi79 := 0; fi79 < 2; fi79++ {
            for fi80 := 0; fi80 < 19; fi80++ {
                var v81 int16 = v77
                v82 := v4[-1]
                fmt.Println(v3, v82, v82)
                _ = fi80
                _ = v81
                _ = v82
            }
            _ = fi79
        }
        v2[("foo bar" + ("\n\t" + ((("\n\t" + ("\n\t" + "foo bar")) + ("" + (("\n\t" + ("a" + "\n\t")) + (("foo bar" + ("foo bar" + "hello")) + "a")))) + ("foo bar" + "a"))))] = (v5 * 4294967295)
        _ = v78
    } else {
        fmt.Println(v76, v53)
    }
    v2[("hello" + ((("foo bar" + ("\n\t" + "")) + "foo bar") + ("foo bar" + "")))] = (v5 % (((v5 << ((878) & 63))) | 1))
    v52 -= v52
    if v53 {
        fmt.Println(v77, v1)
        if len(v9) > 0 {
            v9[0] = 0
        }
        v8 = 1
    } else {
        v52 += 1
    }
    if (!v53) {
        var v83 uint8 = 127
        _ = v83
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v75
    _ = v76
    _ = v77
    _ = v8
    _ = v9
}
