package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 int16 }
func genFn0(p0 bool) uint32 {
    for fi4 := 0; fi4 < 12; fi4++ {
        p0 = false
        var v5 float64 = -2.25
        var v6 string = "a"
        _ = fi4
        _ = v5
        _ = v6
    }
    switch -804 {
    case -524:
        p0 = true
    case 2147483647:
        fmt.Println(p0)
    case 32768:
        v7 := []int{-319}
        _ = v7
    default:
        v8 := GenS1{F0: 225, F1: ("" + "\n\t"), F2: 42}
        _ = v8
    }
    switch ("\n\t" + "foo bar") {
    case "a":
        switch ("hello" + "") {
        case "foo bar":
            for fi9 := 0; fi9 < 15; fi9++ {
                p0 = p0
                var v10 int16 = 741
                v11 := []int{32768, (fi9 / ((fi9) | 1)), fi9, fi9}
                _ = fi9
                _ = v10
                _ = v11
            }
        case "a":
            p0 = p0
        }
    }
    if ("\n\t" != "foo bar") {
        v12 := []int{127, 0, 750}
        p0 = (true && p0)
        if ((2 < 0) || p0) {
            for fi13 := 0; fi13 < 15; fi13++ {
                v14 := GenS1{F0: 1, F1: "\n\t", F2: 739}
                if (-9223372036854775808 >= -1) {
                    if ((!true) || p0) {
                        if p0 {
                            for fi15 := 0; fi15 < 1; fi15++ {
                                var v16 uint32 = 268
                                p0 = ((fi13 | (^(fi15))) <= -9223372036854775808)
                                v17 := GenS1{F0: 213, F1: "\n\t", F2: 65535}
                                var v18 int8 = 3
                                _ = fi15
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                        }
                        p0 = ((42 != -1) || (-551 == 1))
                        _ = len(v12)
                    }
                    var v19 float32 = 1e10
                    _ = v19
                } else {
                    if false {
                        var v20 int8 = -128
                        _ = v20
                    } else {
                        v21 := GenS1{F0: 2, F1: ("hello" + "hello"), F2: 200}
                        _ = v21
                    }
                    v22 := GenS1{F0: 115, F1: "foo bar", F2: 510}
                    var v23 string = "foo bar"
                    _ = v22
                    _ = v23
                }
                v24 := GenS2{F0: 9223372036854775807, F1: 1}
                _ = fi13
                _ = v14
                _ = v24
            }
        } else {
            p0 = (!(0.0 <= 3.14159))
            p0 = (3.14159 >= 0.0)
        }
        _ = v12
    }
    switch "hello" {
    case "hello":
        fmt.Println(p0)
    }
    _ = p0
    return 42
}
func genFn1(p0 int32) int32 {
    fmt.Println(p0)
    if true {
        for fi25 := 0; fi25 < 18; fi25++ {
            p0 += 0
            v26 := GenS1{F0: 142, F1: "a", F2: 579}
            if (!("" == "a")) {
                if (!("hello" > "\n\t")) {
                    var v27 int16 = 1
                    _ = v27
                } else {
                    var v28 uint8 = 127
                    var v29 bool = false
                    _ = v28
                    _ = v29
                }
                v30 := []int{127, fi25, 32767}
                var v31 int8 = -128
                _ = v30
                _ = v31
            }
            _ = fi25
            _ = v26
        }
        v32 := GenS2{F0: 0, F1: 0}
        _ = v32
    }
    p0 = p0
    v33 := []int{2147483647, 128, -820, -1}
    _ = p0
    _ = v33
    return p0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{}
    var v3 uint32 = 1
    for fi34 := 0; fi34 < 2; fi34++ {
        switch "" {
        case "foo bar":
            var v35 int32 = 0
            _ = v35
        default:
            v3 += v3
        }
        _ = fi34
    }
    v36 := GenS3{F0: 42}
    if true {
        if true {
            if true {
                var v37 uint32 = v3
                v3 = (^(v37))
                _ = v37
            } else {
                switch (("hello" + "\n\t") + "a") {
                case "":
                    if true {
                        if true {
                            if (v3 > 325) {
                                fmt.Println(v3)
                                _ = len(v1)
                            }
                        }
                        var v38 float32 = float32(0.0)
                        _ = len(v2)
                        _ = v38
                    }
                case "foo bar":
                    var v39 int = -2147483648
                    _ = v39
                default:
                    if false {
                        v40 := GenS2{F0: 1, F1: 2}
                        _ = v40
                    } else {
                        v41 := GenS1{F0: 0, F1: ("\n\t" + ("hello" + "hello")), F2: 164}
                        if (9223372036854775807 > 0) {
                            v3 = (^(v3))
                            fmt.Println(v3)
                        } else {
                            for fi42 := 0; fi42 < 18; fi42++ {
                                var v43 int8 = 100
                                _ = fi42
                                _ = v43
                            }
                        }
                        v3 -= v3
                        _ = v41
                    }
                }
                v44 := GenS4{F0: 1}
                v3 = v3
                _ = v44
            }
            v3 -= 65535
        }
        v45 := GenS5{F0: 2147483647, F1: -2147483648}
        _ = v45
    }
    if false {
        var v46 bool = false
        v3 = (v3 & 1)
        _ = v46
    } else {
        for fi47 := 0; fi47 < 10; fi47++ {
            var v48 bool = true
            _ = fi47
            _ = v48
        }
    }
    var v49 float32 = -2.25
    v1[genFn0(true)] = 743
    fmt.Println(v49, v49)
    v3 = (v3 * 557)
    var v50 int8 = 48
    v2 = append(v2, (("" + ("hello" + ("\n\t" + "foo bar"))) + (("\n\t" + "\n\t") + "hello")))
    fmt.Println(v50)
    switch "foo bar" {
    case "a":
        switch 820 {
        case -159:
            v3 = v3
        default:
            for fi51 := 0; fi51 < 19; fi51++ {
                v2 = append(v2, ("foo bar" + "a"))
                _ = fi51
            }
        }
    case "foo bar":
        fmt.Println(v49)
    case "":
        var v52 float64 = ((float64((-2.25 / 0.1)) * 3.14159) - float64((v49 + 0.0)))
        _ = v52
    default:
        switch "" {
        case "foo bar":
            _ = len(v2)
        }
    }
    if (false || (!false)) {
        if true {
            if (939 < 2) {
                v3 -= (v3 * (v3 >> ((42) & 63)))
                var v53 float32 = (((v49 / float32((1e10 * 1e10))) + v49) / (v49 / 1e10))
                _ = v53
            } else {
                v54 := GenS6{F0: -558}
                _ = v54
            }
            if false {
                var v55 uint64 = 42
                _ = v55
            }
        }
        fmt.Println(v49, v49, v50)
    }
    if true {
        v56 := GenS2{F0: 1, F1: 4294967295}
        if (!true) {
            fmt.Println(v3, v50, v50)
            v57 := v1[(v3 ^ (v3 & (v3 & 4294967295)))]
            _ = v57
        } else {
            v3 -= (^(v3))
            v49 = (float32(((3.14159 + -2.25) + (1.5 / (0.1 * float64(0.1))))) / -2.25)
        }
        _ = v56
    } else {
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        switch -478 {
        case 42:
            v50 -= (v50 + (v50 | (v50 * 4)))
        }
        if (0 == 32767) {
            for fi58 := 0; fi58 < 11; fi58++ {
                var v59 int16 = 280
                var v60 int8 = (v50 % (((v50 ^ (v50 + v50))) | 1))
                _ = fi58
                _ = v59
                _ = v60
            }
            for fi61 := 0; fi61 < 3; fi61++ {
                switch fi61 {
                case 354:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 1:
                    var v62 int16 = -669
                    _ = v62
                case 408:
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                var v63 int8 = v50
                var v64 int = 1
                _ = fi61
                _ = v63
                _ = v64
            }
            if false {
                v50 = v50
                fmt.Println(v50, v50, v50)
                if len(v2) > 0 {
                    v2[0] = ("" + "")
                }
            }
        } else {
            v65 := GenS3{F0: -415}
            var v66 int8 = (-(v50))
            _ = v65
            _ = v66
        }
    }
    var v67 bool = (-1 != 197)
    v68 := GenS2{F0: -1, F1: 79}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v49
    _ = v50
    _ = v67
    _ = v68
}
