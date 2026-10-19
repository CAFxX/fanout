package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint64; F1 string; F2 int64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int16; F1 int16 }
type GenS12 struct { F0 uint32; F1 uint32; F2 uint16 }
func genFn0(p0 uint64, p1 bool) uint {
    p1 = p1
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint16 = 2
    var v2 bool = true
    var v3 int32 = 2147483647
    v4 := make(map[uint32]string)
    var v5 string = ""
    v6 := []int{1, -897, 128, -1}
    v7 := GenS1{F0: 0}
    if (!((0.0 == (1.5 / (1e10 + 3.14159))) && ((-1 <= 0) || (v2 && v2)))) {
        v5 = "a"
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 float32 = -2.25
            v9 -= v9
            var v10 int = (fi8 >> ((42) & 63))
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    if (0.0 <= 0.0) {
        var v11 uint = 255
        v5 = v5
        fmt.Println(v2, v3, v2)
        _ = v11
    }
    if (v2 && v2) {
        if v2 {
            var v12 uint = 4294967295
            _ = v12
        }
    } else {
        v13 := GenS2{F0: 651}
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := GenS3{F0: v1, F1: 42}
        if len(v6) > 0 {
            v6[0] = (fi14 & fi14)
        }
        v16 := GenS4{F0: 127, F1: 42}
        _ = fi14
        _ = v15
        _ = v16
    }
    if false {
        v17 := GenS5{F0: 42, F1: 340}
        var v18 uint32 = 0
        _ = v17
        _ = v18
    }
    if (128 < 97) {
        var v19 int32 = int32(v1)
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 int16 = 232
            if v2 {
                switch "foo bar" {
                case "":
                    for fi22 := 0; fi22 < 19; fi22++ {
                        switch ("hello" + (v5 + v5)) {
                        case "hello":
                            v23 := GenS6{F0: (v3 + v3)}
                            _ = v23
                        }
                        if (!v2) {
                            v24 := GenS2{F0: (^(v21))}
                            v5 = "foo bar"
                            v19 = v3
                            _ = v24
                        }
                        _ = fi22
                    }
                case "a":
                    var v25 int64 = 42
                    _ = v25
                default:
                    for fi26 := 0; fi26 < 19; fi26++ {
                        v1 -= (v1 - v1)
                        var v27 uint64 = 0
                        _ = len(v4)
                        if (9223372036854775807 != 9223372036854775807) {
                            fmt.Println(fi26)
                        } else {
                            v4[42] = "hello"
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        fmt.Println(v19, fi20)
                        _ = fi26
                        _ = v27
                    }
                }
                switch fi20 {
                case -696:
                    if v2 {
                        fmt.Println(v2, v19)
                        v4[65535] = "a"
                        for fi28 := 0; fi28 < 4; fi28++ {
                            var v29 string = v5
                            v30 := GenS7{F0: v3}
                            fmt.Println(v19, v3, v5)
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    }
                case -9223372036854775808:
                    for fi31 := 0; fi31 < 9; fi31++ {
                        if v2 {
                            var v32 bool = (0.1 != 1.5)
                            _ = v32
                        }
                        var v33 int32 = (v19 * v3)
                        _ = fi31
                        _ = v33
                    }
                case 1:
                    if true {
                        var v34 uint8 = 128
                        _ = v34
                    }
                default:
                    var v35 bool = ((0.0 / 0.1) > 3.14159)
                    _ = v35
                }
                var v36 uint = 255
                _ = v36
            } else {
                var v37 uint64 = 1
                var v38 int8 = 42
                _ = v37
                _ = v38
            }
            if ((1e10 * 0.1) != ((1e10 / 1e10) - -2.25)) {
                fmt.Println(v19, v2)
                if false {
                    if (!(v2 || true)) {
                        v3 += -2147483648
                    } else {
                        v39 := GenS8{F0: 516, F1: "foo bar", F2: 9223372036854775807}
                        _ = v39
                    }
                } else {
                    var v40 int16 = (v21 - (v21 ^ 97))
                    for fi41 := 0; fi41 < 2; fi41++ {
                        v4[4294967295] = (v5 + "hello")
                        fmt.Println(v21)
                        _ = fi41
                    }
                    v42 := GenS7{F0: (^(v19))}
                    _ = v40
                    _ = v42
                }
            }
            _ = fi20
            _ = v21
        }
        _ = v19
    } else {
        var v43 uint64 = 979
        switch 128 {
        case -161:
            if v2 {
                switch ("" + (("a" + v5) + v5)) {
                case "foo bar":
                    switch "a" {
                    case "a":
                        for fi44 := 0; fi44 < 14; fi44++ {
                            for fi45 := 0; fi45 < 2; fi45++ {
                                _ = len(v4)
                                v46 := GenS9{F0: 251}
                                v2 = v2
                                _ = fi45
                                _ = v46
                            }
                            v47 := v4[424]
                            _ = fi44
                            _ = v47
                        }
                    case "foo bar":
                        if false {
                            var v48 string = (("a" + ("hello" + v5)) + "\n\t")
                            switch v48 {
                            case "hello":
                                _ = len(v4)
                            case "":
                                if v2 {
                                    v49 := GenS10{F0: -1}
                                    v6 = append(v6, -577)
                                    for fi50 := 0; fi50 < 3; fi50++ {
                                        var v51 int8 = 127
                                        _ = fi50
                                        _ = v51
                                    }
                                    _ = v49
                                }
                            default:
                                for fi52 := 0; fi52 < 10; fi52++ {
                                    v5 = "hello"
                                    v3 -= 42
                                    _ = fi52
                                }
                            }
                            _ = len(v6)
                            _ = v48
                        }
                    case "hello":
                        v53 := GenS7{F0: 42}
                        _ = v53
                    }
                }
            }
        case 904:
            if v2 {
                v54 := GenS10{F0: 281}
                var v55 uint8 = 89
                var v56 float64 = 0.0
                _ = v54
                _ = v55
                _ = v56
            }
        default:
            if len(v6) > 0 {
                v6[0] = 128
            }
        }
        v6 = append(v6, -529)
        _ = v43
    }
    if v2 {
        v2 = ((v1 ^ 32767) == v1)
        v5 = v5
    }
    var v57 uint16 = (v1 + v1)
    v58 := GenS11{F0: -1, F1: -475}
    v2 = v2
    switch 42 {
    case 1:
        var v59 string = (v5 + "foo bar")
        _ = v59
    case -903:
        for fi60 := 0; fi60 < 7; fi60++ {
            var v61 uint = 0
            for fi62 := 0; fi62 < 9; fi62++ {
                v57 = 2
                v63 := GenS12{F0: 1, F1: 1, F2: (^(v57))}
                v64 := GenS5{F0: 4294967295, F1: 2}
                _ = fi62
                _ = v63
                _ = v64
            }
            var v65 int = 127
            _ = fi60
            _ = v61
            _ = v65
        }
    }
    if (v2 || (v2 && false)) {
        var v66 uint64 = 895
        v1 = 32767
        _ = v66
    }
    var v67 bool = (!(!(("" + "") < "foo bar")))
    _ = len(v4)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
    _ = v67
    _ = v7
}
