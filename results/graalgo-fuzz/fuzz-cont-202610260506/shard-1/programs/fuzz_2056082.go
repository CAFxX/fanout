package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 string; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int8; F1 int32; F2 bool }
type GenS5 struct { F0 string; F1 int; F2 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int; F1 int8; F2 uint64 }
func genFn0(p0 uint8) uint8 {
    v5 := GenS1{F0: ("foo bar" + "hello")}
    var v6 float64 = 1e10
    v6 += (v6 - (float64(1e10) * 0.1))
    p0 = p0
    fmt.Println(p0, v6)
    _ = p0
    _ = v5
    _ = v6
    return p0
}
func genFn1(p0 int16) bool {
    p0 -= p0
    for fi7 := 0; fi7 < 9; fi7++ {
        switch "a" {
        case "foo bar":
            var v8 int64 = -153
            _ = v8
        }
        v9 := []int{}
        _ = fi7
        _ = v9
    }
    switch "\n\t" {
    case "a":
        switch "" {
        case "a":
            p0 = (p0 * (p0 | p0))
        }
    case "\n\t":
        if true {
            var v10 float64 = 1e10
            var v11 float32 = 1.5
            var v12 uint16 = 32767
            _ = v10
            _ = v11
            _ = v12
        }
    default:
        v13 := []int{127}
        _ = v13
    }
    _ = p0
    return true
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    v3 := []string{}
    v4 := make(map[uint32]string)
    var v14 int = 0
    switch ("hello" + "a") {
    case "foo bar":
        var v15 uint16 = 1
        _ = v15
    case "":
        v16 := GenS2{F0: 0, F1: "hello", F2: 667}
        _ = v16
    case "hello":
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := GenS3{F0: 2147483647, F1: -1}
            _ = fi17
            _ = v18
        }
    }
    v14 = v14
    var v19 float64 = 0.1
    var v20 int8 = -89
    v19 -= 1.5
    if false {
        v21 := GenS3{F0: 2147483648, F1: (v20 / ((v20) | 1))}
        v14 -= (-(v14))
        fmt.Println(v19)
        _ = v21
    } else {
        switch v14 {
        case 2147483647:
            if ((18446744073709551615 <= 0) || (1e10 < v19)) {
                for fi22 := 0; fi22 < 1; fi22++ {
                    fmt.Println(v19)
                    _ = fi22
                }
                for fi23 := 0; fi23 < 4; fi23++ {
                    v24 := GenS2{F0: 94, F1: (("" + (("foo bar" + ("foo bar" + "")) + "")) + ("" + "hello")), F2: 209}
                    v3 = append(v3, ("a" + "a"))
                    switch (-(fi23)) {
                    case 128:
                        for fi25 := 0; fi25 < 11; fi25++ {
                            v26 := GenS2{F0: -663, F1: "a", F2: 2147483648}
                            if (((v19 + v19) * (-2.25 - (1.5 - 1e10))) > 3.14159) {
                                v14 -= fi25
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = ("\n\t" + "")
                                }
                                var v27 int32 = 4
                                _ = v27
                            }
                            _ = fi25
                            _ = v26
                        }
                    case -1:
                        fmt.Println(v14, v19, v14)
                    default:
                        fmt.Println(fi23, v19, v19)
                    }
                    _ = fi23
                    _ = v24
                }
            }
        case 863:
            v20 = (v20 >> ((564) & 63))
        }
        var v28 int = (v14 - v14)
        v14 = 9223372036854775807
        _ = v28
    }
    var v29 int16 = -213
    v19 = v19
    v19 -= v19
    v30 := GenS3{F0: 981, F1: v20}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        var v32 float64 = v19
        switch "foo bar" {
        case "":
            var v33 int64 = 1
            _ = v33
        case "foo bar":
            v20 -= v20
        default:
            if (!(1 <= -9223372036854775808)) {
                v14 = fi31
                _ = len(v3)
            }
        }
        if (((("\n\t" + "foo bar") + ("\n\t" + "a")) + "hello") > "a") {
            if (v20 >= v20) {
                switch (fi31 ^ 127) {
                case 127:
                    v19 += (1.5 + float64(3.14159))
                }
            } else {
                for fi34 := 0; fi34 < 17; fi34++ {
                    v32 = v32
                    fmt.Println(fi31)
                    _ = fi34
                }
            }
        } else {
            if false {
                var v35 int8 = v20
                if (!("hello" < "hello")) {
                    switch ("foo bar" + "\n\t") {
                    case "":
                        var v36 uint16 = 861
                        _ = v36
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    v37 := GenS3{F0: 42, F1: -1}
                    for fi38 := 0; fi38 < 12; fi38++ {
                        if true {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if false {
                                if (v19 <= float64(float32(float64((3.14159 * (-2.25 - 1e10)))))) {
                                    v3 = append(v3, ("hello" + "a"))
                                } else {
                                    v39 := GenS1{F0: "\n\t"}
                                    var v40 uint8 = 255
                                    _ = v39
                                    _ = v40
                                }
                                v32 = ((0.0 - v32) + float64(float32(v19)))
                            }
                            var v41 uint64 = 4294967295
                            _ = v41
                        }
                        if true {
                            var v42 uint16 = 2
                            if (!(!false)) {
                                fmt.Println(v42, v42, fi38)
                                var v43 uint32 = 2147483647
                                v44 := GenS4{F0: v35, F1: 1, F2: false}
                                v45 := GenS3{F0: 0, F1: v35}
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v42
                        } else {
                            switch "a" {
                            case "\n\t":
                                v46 := GenS1{F0: "\n\t"}
                                _ = v46
                            case "":
                                v47 := GenS5{F0: ("\n\t" + "hello"), F1: fi38, F2: 2}
                                _ = v47
                            case "a":
                                if (831 > 2) {
                                    v14 = -77
                                }
                                fmt.Println(v35, v20, v20)
                                if false {
                                    if (((967 == 18446744073709551615) || true) || (-2.25 > 0.0)) {
                                        v35 += (v35 + v20)
                                        if true {
                                            fmt.Println(v20, fi31)
                                            switch fi38 {
                                            case -2147483648:
                                                var v48 int16 = (v29 % (((-(v29))) | 1))
                                                _ = v48
                                            default:
                                                var v49 uint8 = 255
                                                _ = v49
                                            }
                                            v14 = (v14 << ((65535) & 63))
                                            switch "hello" {
                                            case "":
                                                v50 := GenS2{F0: v29, F1: "hello", F2: 741}
                                                _ = v50
                                            }
                                            var v51 float32 = 1e10
                                            _ = v51
                                        } else {
                                            if (!true) {
                                                var v52 bool = (false || (0.1 <= 3.14159))
                                                switch ("a" + "foo bar") {
                                                case "hello":
                                                    _ = len(v2)
                                                case "\n\t":
                                                    var v53 uint8 = 76
                                                    _ = v53
                                                case "foo bar":
                                                    var v54 int8 = (v35 ^ (v35 % (((v35 >> ((0) & 63))) | 1)))
                                                    _ = v54
                                                default:
                                                    v19 = v19
                                                }
                                                _ = v52
                                            } else {
                                                var v55 float32 = -2.25
                                                v56 := GenS4{F0: v35, F1: -882, F2: (true && true)}
                                                if false {
                                                    _ = len(v3)
                                                } else {
                                                    v57 := GenS6{F0: 1}
                                                    v32 = 0.1
                                                    var v58 uint = 0
                                                    _ = v57
                                                    _ = v58
                                                }
                                                _ = v55
                                                _ = v56
                                            }
                                        }
                                        var v59 uint16 = 2
                                        _ = v59
                                    } else {
                                        switch "hello" {
                                        case "foo bar":
                                            v29 = v29
                                        case "\n\t":
                                            v14 = fi38
                                        default:
                                            if len(v3) > 0 {
                                                v3[0] = "hello"
                                            }
                                        }
                                        var v60 string = ("hello" + ("" + "hello"))
                                        var v61 bool = (((v35 != v35) && ((v35 - 100) < (v35 * (v20 | (v20 + (-(v20))))))) && true)
                                        _ = v60
                                        _ = v61
                                    }
                                } else {
                                    v32 -= v32
                                    v62 := GenS2{F0: v29, F1: "hello", F2: 779}
                                    v32 -= float64(0.0)
                                    _ = v62
                                }
                                if ((0.1 * 0.0) != 0.1) {
                                    var v63 float32 = float32(((v32 - 0.1) + v19))
                                    _ = v63
                                } else {
                                    v29 = (^(v29))
                                    fmt.Println(v20, v19)
                                    v64 := GenS4{F0: v35, F1: 2147483647, F2: (false || (0 != -32768))}
                                    _ = v64
                                }
                                var v65 int64 = 1
                                _ = len(v4)
                                switch 128 {
                                case -1:
                                    var v66 uint64 = 42
                                    _ = v66
                                default:
                                    v67 := v1[42]
                                    _ = v67
                                }
                                if true {
                                    var v68 uint32 = 122
                                    v29 += (v29 * v29)
                                    v29 = 0
                                    _ = v68
                                } else {
                                    fmt.Println(v35, fi38)
                                    _ = len(v3)
                                    v32 -= 1.5
                                }
                                var v69 bool = (!false)
                                _ = v65
                                _ = v69
                            default:
                                v70 := GenS2{F0: v29, F1: "a", F2: 414}
                                _ = v70
                            }
                            var v71 int8 = (v20 ^ v35)
                            var v72 uint32 = 42
                            _ = v71
                            _ = v72
                        }
                        if true {
                            v73 := GenS5{F0: "a", F1: 9223372036854775807, F2: 0}
                            v74 := GenS7{F0: (v20 - v20)}
                            if false {
                                v35 += 1
                            }
                            fmt.Println(v14, v14, fi38)
                            v14 = fi31
                            fmt.Println(v35, v35)
                            var v75 float64 = float64((-2.25 + 1e10))
                            _ = v73
                            _ = v74
                            _ = v75
                        }
                        _ = fi38
                    }
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                    _ = v37
                }
                if len(v3) > 0 {
                    v3[0] = ""
                }
                _ = v35
            } else {
                for fi76 := 0; fi76 < 19; fi76++ {
                    var v77 int = fi76
                    _ = fi76
                    _ = v77
                }
                var v78 bool = ("\n\t" > "foo bar")
                for fi79 := 0; fi79 < 9; fi79++ {
                    v80 := GenS8{F0: 32767, F1: (v20 ^ v20), F2: 1}
                    v81 := GenS2{F0: (v29 + 42), F1: "", F2: 1}
                    v32 = v19
                    switch (v14 >> ((255) & 63)) {
                    case 4294967295:
                        v19 -= v19
                    default:
                        var v82 float32 = 0.0
                        _ = v82
                    }
                    _ = fi79
                    _ = v80
                    _ = v81
                }
                _ = v78
            }
            _ = len(v1)
        }
        _ = fi31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
