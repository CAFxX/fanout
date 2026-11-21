package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 bool }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 bool; F1 uint8; F2 int16 }
type GenS4 struct { F0 string; F1 string; F2 uint8 }
type GenS5 struct { F0 string; F1 string; F2 int16 }
type GenS6 struct { F0 uint16 }
func genFn0() int8 {
    switch 128 {
    case -1:
        var v7 string = "foo bar"
        _ = v7
    case 2147483647:
        v8 := GenS1{F0: 42, F1: 2, F2: (-2147483648 > 1)}
        _ = v8
    case 32767:
        var v9 int64 = 9223372036854775807
        _ = v9
    }
    var v10 int64 = 1
    if (false && true) {
        v11 := []int{2147483647, -1, 128}
        fmt.Println(v10)
        _ = v11
    } else {
        v12 := []int{-9223372036854775808, -1}
        var v13 int16 = 32767
        for fi14 := 0; fi14 < 12; fi14++ {
            v12 = append(v12, (fi14 / ((fi14) | 1)))
            _ = fi14
        }
        _ = v12
        _ = v13
    }
    var v15 int16 = 1
    _ = v10
    _ = v15
    return -1
}

func main() {
    v1 := []string{"foo bar", "\n\t", "hello"}
    v2 := make(map[string]string)
    v3 := []string{"foo bar"}
    var v4 int64 = 42
    v5 := make(map[int]string)
    v6 := make(map[uint32]bool)
    for fi16 := 0; fi16 < 1; fi16++ {
        for fi17 := 0; fi17 < 2; fi17++ {
            v18 := GenS1{F0: (v4 << ((2) & 63)), F1: 255, F2: true}
            if (42 >= 42) {
                if false {
                    fmt.Println(v4)
                    _ = len(v5)
                    v19 := GenS1{F0: 1, F1: 18446744073709551615, F2: (42 > 42)}
                    v4 -= (v4 % ((v4) | 1))
                    _ = v19
                }
                fmt.Println(fi16, fi16)
                var v20 int32 = 42
                _ = v20
            }
            switch "a" {
            case "a":
                v21 := GenS2{F0: "foo bar", F1: -106}
                _ = v21
            case "foo bar":
                v22 := GenS2{F0: "a", F1: 1}
                _ = v22
            default:
                var v23 int8 = 1
                _ = v23
            }
            _ = fi17
            _ = v18
        }
        var v24 string = "hello"
        _ = fi16
        _ = v24
    }
    v4 -= v4
    _ = len(v1)
    if false {
        switch (("a" + "foo bar") + "\n\t") {
        case "hello":
            switch "a" {
            case "foo bar":
                v4 -= (v4 % ((0) | 1))
            default:
                for fi25 := 0; fi25 < 7; fi25++ {
                    var v26 uint16 = 71
                    fmt.Println(v26, fi25, v26)
                    fmt.Println(fi25, v26)
                    _ = fi25
                    _ = v26
                }
            }
        }
    }
    if true {
        v4 += (v4 * v4)
        for fi27 := 0; fi27 < 11; fi27++ {
            var v28 int8 = -1
            v29 := GenS2{F0: "hello", F1: 42}
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    var v30 int = -984
    for fi31 := 0; fi31 < 18; fi31++ {
        switch fi31 {
        case 0:
            var v32 float64 = ((3.14159 - float64((((0.1 / float32(((3.14159 / 1.5) * ((0.0 + ((0.1 * 1e10) - 0.0)) / -2.25)))) - 0.1) - 1e10))) - (float64(((0.1 - 1.5) * -2.25)) + -2.25))
            _ = v32
        default:
            _ = len(v3)
        }
        var v33 uint = 1
        if (3.14159 != (float32((3.14159 - 3.14159)) / (-2.25 + 1e10))) {
            for fi34 := 0; fi34 < 6; fi34++ {
                switch ("\n\t" + (("\n\t" + "") + "hello")) {
                case "a":
                    v30 -= v30
                case "":
                    v35 := GenS3{F0: ((v4 / (((v4 & v4)) | 1)) >= v4), F1: 9, F2: 32767}
                    _ = v35
                case "\n\t":
                    v36 := GenS4{F0: "", F1: "\n\t", F2: 128}
                    _ = v36
                default:
                    var v37 float32 = 0.0
                    _ = v37
                }
                var v38 int = int(v33)
                v30 += v30
                if false {
                    var v39 string = (("\n\t" + "foo bar") + "")
                    _ = v39
                } else {
                    fmt.Println(fi31, v38, v33)
                    var v40 int32 = 2147483647
                    _ = v40
                }
                fmt.Println(v4)
                var v41 bool = (128 == 103)
                _ = fi34
                _ = v38
                _ = v41
            }
        } else {
            v33 = v33
            var v42 int8 = 66
            _ = v42
        }
        _ = fi31
        _ = v33
    }
    var v43 uint16 = 0
    if (-1 != int32(v4)) {
        var v44 float32 = 0.1
        v44 = (((float32(0.0) / v44) * 0.0) - v44)
        _ = v44
    } else {
        var v45 float64 = 0.0
        if (-974 < 0) {
            var v46 string = "\n\t"
            _ = v46
        } else {
            _ = len(v5)
            for fi47 := 0; fi47 < 19; fi47++ {
                v48 := GenS5{F0: "\n\t", F1: "\n\t", F2: 42}
                v4 -= v4
                _ = fi47
                _ = v48
            }
            if true {
                for fi49 := 0; fi49 < 18; fi49++ {
                    var v50 uint32 = 65535
                    v51 := GenS3{F0: (128 < 255), F1: uint8(v50), F2: 1}
                    _ = fi49
                    _ = v50
                    _ = v51
                }
                if (42 <= 42) {
                    for fi52 := 0; fi52 < 1; fi52++ {
                        var v53 bool = true
                        _ = fi52
                        _ = v53
                    }
                } else {
                    v2["\n\t"] = "\n\t"
                    switch ("a" + "hello") {
                    case "a":
                        for fi54 := 0; fi54 < 13; fi54++ {
                            v4 -= v4
                            if (v30 == v30) {
                                if true {
                                    fmt.Println(v4)
                                } else {
                                    var v55 uint64 = 0
                                    for fi56 := 0; fi56 < 18; fi56++ {
                                        v43 = (v43 - v43)
                                        var v57 string = "a"
                                        _ = fi56
                                        _ = v57
                                    }
                                    if false {
                                        if (v43 != v43) {
                                            if false {
                                                v30 += -134
                                            }
                                            var v58 uint32 = 2147483647
                                            v59 := GenS2{F0: ("\n\t" + "a"), F1: 100}
                                            _ = v58
                                            _ = v59
                                        }
                                    }
                                    _ = v55
                                }
                                var v60 int8 = 0
                                var v61 uint32 = 4294967295
                                _ = v60
                                _ = v61
                            }
                            _ = fi54
                        }
                    default:
                        if false {
                            for fi62 := 0; fi62 < 9; fi62++ {
                                v4 = (-(v4))
                                if false {
                                    v63 := GenS6{F0: v43}
                                    _ = v63
                                } else {
                                    _ = len(v6)
                                    for fi64 := 0; fi64 < 12; fi64++ {
                                        var v65 float32 = 3.14159
                                        v66 := GenS5{F0: ("a" + ""), F1: ("" + ""), F2: 42}
                                        _ = fi64
                                        _ = v65
                                        _ = v66
                                    }
                                }
                                _ = fi62
                            }
                        }
                    }
                }
                fmt.Println(v43, v30, v30)
            }
        }
        _ = v45
    }
    v4 = (v4 + v4)
    var v67 uint64 = 953
    fmt.Println(v30, v4, v43)
    v43 -= v43
    if true {
        v68 := GenS2{F0: "", F1: 0}
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        _ = v68
    }
    v67 = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v67
}
