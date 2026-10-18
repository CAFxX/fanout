package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 string }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int; F1 uint32; F2 int64 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string; F1 int32; F2 int32 }
type GenS9 struct { F0 uint16; F1 int8; F2 uint32 }
func genFn0() uint32 {
    v5 := GenS1{F0: 32767, F1: -1, F2: ("hello" + "a")}
    var v6 uint64 = 547
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS2{F0: 2, F1: -929}
        v9 := GenS1{F0: 1, F1: -236, F2: "a"}
        fmt.Println(v6)
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int = 128
    var v2 uint32 = 2147483647
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    v10 := GenS3{F0: ((true || true) && (0 == 65535)), F1: "a"}
    switch "hello" {
    case "":
        switch v1 {
        case 32767:
            if true {
                fmt.Println(v2, v1)
                v11 := []int{(v1 + (v1 * (v1 ^ 127))), v1, v1, (^(v1))}
                if len(v11) > 0 {
                    v11[0] = (v1 & (v1 >> ((722) & 63)))
                }
                _ = v11
            } else {
                var v12 uint = 266
                if true {
                    if true {
                        var v13 uint32 = 2147483648
                        _ = v13
                    } else {
                        for fi14 := 0; fi14 < 4; fi14++ {
                            v15 := GenS2{F0: (v12 ^ 0), F1: -32768}
                            v16 := GenS4{F0: (^(fi14)), F1: (^(v2)), F2: 1}
                            var v17 int = (v1 | v1)
                            _ = fi14
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        v2 = 835
                        if (v12 >= v12) {
                            v3[(v1 + -307)] = "a"
                            for fi18 := 0; fi18 < 20; fi18++ {
                                v12 += v12
                                _ = fi18
                            }
                        }
                    }
                }
                v1 += -410
                _ = v12
            }
        case 4294967295:
            var v19 uint32 = v2
            _ = v19
        }
    default:
        if false {
            var v20 uint32 = (^(v2))
            for fi21 := 0; fi21 < 2; fi21++ {
                for fi22 := 0; fi22 < 3; fi22++ {
                    if ((^(v20)) > v20) {
                        if false {
                            v23 := GenS4{F0: fi22, F1: (v20 - (v2 % (((v20 + (v20 * (v2 - 850)))) | 1))), F2: -622}
                            _ = v23
                        } else {
                            fmt.Println(v1, fi21)
                            switch (("foo bar" + "") + "") {
                            case "":
                                if (true && false) {
                                    v20 = (v2 * (^(v2)))
                                    var v24 uint8 = 1
                                    fmt.Println(v24, v20, fi21)
                                    _ = v24
                                } else {
                                    switch "\n\t" {
                                    case "a":
                                        delete(v4, fi21)
                                    case "\n\t":
                                        v1 += fi21
                                    }
                                    switch ("hello" + "\n\t") {
                                    case "hello":
                                        if true {
                                            var v25 int32 = 2147483647
                                            _ = v25
                                        }
                                    case "foo bar":
                                        var v26 uint64 = 145
                                        _ = v26
                                    case "\n\t":
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v27 bool = (!(2 > 234))
                                            _ = v27
                                        }
                                    default:
                                        fmt.Println(v1, v2, fi21)
                                    }
                                }
                            default:
                                fmt.Println(v1)
                            }
                            v2 = v2
                        }
                        fmt.Println(fi21, v20, v20)
                        v1 -= v1
                        v2 -= (v20 - 244)
                    } else {
                        var v28 bool = false
                        _ = v28
                    }
                    v29 := v3[128]
                    switch v29 {
                    case "foo bar":
                        if ((526 > 887) && false) {
                            var v30 uint16 = 0
                            delete(v3, 738)
                            v31 := []int{}
                            _ = v30
                            _ = v31
                        }
                    }
                    _ = fi22
                    _ = v29
                }
                for fi32 := 0; fi32 < 14; fi32++ {
                    if ((0 < 371) || false) {
                        v33 := GenS5{F0: 42, F1: -1}
                        _ = v33
                    } else {
                        v20 = 2147483647
                    }
                    v34 := GenS6{F0: fi21, F1: "foo bar"}
                    v35 := GenS6{F0: (-(fi32)), F1: "hello"}
                    _ = fi32
                    _ = v34
                    _ = v35
                }
                _ = fi21
            }
            if ("hello" >= "\n\t") {
                var v36 bool = false
                v37 := GenS6{F0: v1, F1: ("foo bar" + "")}
                if true {
                    v38 := GenS1{F0: 1000, F1: -1, F2: "\n\t"}
                    _ = v38
                }
                _ = v36
                _ = v37
            } else {
                fmt.Println(v20, v1, v1)
            }
            _ = v20
        } else {
            fmt.Println(v2)
            v1 = v1
        }
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        fmt.Println(v1, fi39)
        _ = fi39
    }
    switch "foo bar" {
    case "\n\t":
        if true {
            if (1 == 183) {
                v40 := []int{1, (v1 << ((270) & 63))}
                _ = v40
            } else {
                switch "a" {
                case "a":
                    switch v1 {
                    case 167:
                        switch v1 {
                        case -2147483648:
                            var v41 uint32 = 0
                            _ = v41
                        case 1:
                            for fi42 := 0; fi42 < 1; fi42++ {
                                v43 := GenS6{F0: v1, F1: "hello"}
                                v1 -= (fi42 | -1)
                                for fi44 := 0; fi44 < 10; fi44++ {
                                    fmt.Println(fi42, fi44, v1)
                                    v1 = fi42
                                    v45 := GenS2{F0: 883, F1: 32767}
                                    var v46 bool = false
                                    _ = fi44
                                    _ = v45
                                    _ = v46
                                }
                                _ = fi42
                                _ = v43
                            }
                        case 32767:
                            if false {
                                if (3.14159 > -2.25) {
                                    if false {
                                        fmt.Println(v1, v2)
                                    } else {
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v47 uint32 = 2147483647
                                            _ = v47
                                        default:
                                            if (!false) {
                                                if true {
                                                    v48 := []int{(v1 & v1), (v1 << ((0) & 63))}
                                                    v1 = v1
                                                    v49 := GenS7{F0: -11}
                                                    _ = v48
                                                    _ = v49
                                                }
                                            } else {
                                                v2 = (^(v2))
                                            }
                                        }
                                    }
                                }
                                var v50 float64 = (1.5 + 0.0)
                                switch v1 {
                                case 1:
                                    var v51 float32 = 3.14159
                                    _ = v51
                                }
                                _ = v50
                            } else {
                                if false {
                                    _ = len(v4)
                                    v1 += v1
                                    var v52 uint16 = 2
                                    _ = v52
                                } else {
                                    v53 := []int{}
                                    var v54 float32 = (float32((1e10 + float64((((((0.0 * (3.14159 - (1e10 - 3.14159))) + 0.0) / (0.1 - 1e10)) + 1.5) * 1.5)))) - 3.14159)
                                    v55 := v3[v1]
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                }
                                var v56 uint32 = v2
                                _ = v56
                            }
                        default:
                            var v57 float32 = 0.1
                            _ = v57
                        }
                    case 42:
                        v2 = (v2 + 2147483648)
                    }
                case "":
                    v58 := GenS8{F0: "a", F1: 1, F2: -650}
                    _ = v58
                case "foo bar":
                    if (!false) {
                        v59 := GenS9{F0: 1, F1: -128, F2: (v2 & (v2 % ((genFn0()) | 1)))}
                        _ = v59
                    } else {
                        var v60 int16 = -32768
                        v61 := GenS5{F0: v60, F1: 9223372036854775807}
                        _ = len(v3)
                        _ = v60
                        _ = v61
                    }
                }
                if true {
                    var v62 uint16 = 274
                    var v63 int = (v1 & 32768)
                    switch v1 {
                    case 9223372036854775807:
                        v63 -= 1
                    }
                    _ = v62
                    _ = v63
                }
                for fi64 := 0; fi64 < 9; fi64++ {
                    var v65 uint64 = 2
                    _ = fi64
                    _ = v65
                }
            }
        }
    case "hello":
        v1 = (v1 - 9223372036854775807)
    case "":
        fmt.Println(v1, v2)
    default:
        var v66 int32 = 2147483647
        _ = v66
    }
    v67 := GenS3{F0: (0.0 == -2.25), F1: "hello"}
    var v68 int32 = 433
    var v69 int8 = -1
    v69 += v69
    v70 := []int{1, v1, v1, v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("slicelen v70:", len(v70), v70)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v67
    _ = v68
    _ = v69
    _ = v70
}
