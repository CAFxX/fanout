package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int32; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 bool; F2 int32 }
type GenS6 struct { F0 int; F1 uint64 }
type GenS7 struct { F0 int64; F1 int64 }
type GenS8 struct { F0 bool; F1 uint64 }
type GenS9 struct { F0 uint32 }
func genFn0(p0 int16, p1 int8, p2 int64) int64 {
    var v6 float64 = 3.14159
    v7 := []int{42, 578, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -1
}
func genFn1() bool {
    var v8 float32 = float32(1e10)
    var v9 uint = 0
    var v10 string = ""
    if (-1 == -184) {
        v9 = v9
    } else {
        var v11 uint32 = 0
        _ = v11
    }
    _ = v10
    _ = v8
    _ = v9
    return ((42 > 42) && true)
}

func main() {
    v1 := []string{"foo bar", ("foo bar" + "hello"), "foo bar", "\n\t"}
    var v2 uint64 = 4294967295
    var v3 uint32 = 2147483648
    var v4 uint32 = v3
    v5 := make(map[uint32]int)
    var v12 float64 = ((0.0 - 0.0) + -2.25)
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 int8 = -123
        if true {
            if false {
                v15 := GenS1{F0: fi13, F1: int64(v4)}
                _ = v15
            }
        }
        switch fi13 {
        case 290:
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 uint32 = v3
                v18 := GenS1{F0: fi16, F1: -1}
                v19 := GenS2{F0: 42}
                fmt.Println(v3, v17)
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
        case 127:
            if (!(true || true)) {
                v20 := GenS3{F0: 179, F1: 2147483647, F2: 0}
                if false {
                    var v21 float32 = -2.25
                    var v22 float64 = 1e10
                    _ = v21
                    _ = v22
                } else {
                    v14 = (v14 >> ((586) & 63))
                    v4 -= (v3 - 472)
                }
                v23 := GenS4{F0: (fi13 - 1)}
                _ = v20
                _ = v23
            }
        }
        _ = fi13
        _ = v14
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v12)
    for fi24 := 0; fi24 < 20; fi24++ {
        switch "foo bar" {
        case "foo bar":
            for fi25 := 0; fi25 < 3; fi25++ {
                if false {
                    var v26 int16 = -466
                    if (true && ((v26 < (v26 << ((42) & 63))) || (v12 >= v12))) {
                        v27 := v5[603]
                        v28 := GenS4{F0: 128}
                        _ = v27
                        _ = v28
                    } else {
                        fmt.Println(v12)
                    }
                    _ = v26
                }
                _ = fi25
            }
        case "\n\t":
            v29 := v5[(v4 / (((v4 >> ((255) & 63))) | 1))]
            _ = v29
        default:
            if (((-2147483648 <= 1) || (("foo bar" <= "a") || false)) || true) {
                v4 -= (v4 & v3)
                switch "\n\t" {
                case "a":
                    var v30 uint = 65535
                    _ = v30
                }
            } else {
                for fi31 := 0; fi31 < 7; fi31++ {
                    v32 := GenS4{F0: (fi31 | (fi24 >> ((255) & 63)))}
                    var v33 int64 = -9223372036854775808
                    if false {
                        var v34 uint16 = 2
                        if ((^(v33)) >= v33) {
                            v35 := GenS5{F0: -325, F1: (!(true && (!(!true)))), F2: int32(v33)}
                            _ = v35
                        } else {
                            v36 := GenS1{F0: 42, F1: (-(v33))}
                            v37 := GenS6{F0: fi31, F1: 2}
                            _ = v36
                            _ = v37
                        }
                        _ = v34
                    } else {
                        if true {
                            var v38 int = (fi31 ^ (fi24 ^ 42))
                            if ((v33 | (v33 / ((v33) | 1))) >= v33) {
                                switch (-(v38)) {
                                case -296:
                                    v3 -= (v3 * (v3 - 0))
                                case 128:
                                    v39 := GenS4{F0: fi31}
                                    _ = v39
                                }
                                delete(v5, v4)
                                v40 := GenS5{F0: 160, F1: true, F2: 135}
                                _ = v40
                            } else {
                                switch 0 {
                                case 127:
                                    v1 = append(v1, ("a" + ("hello" + "foo bar")))
                                case 2147483647:
                                    v41 := GenS7{F0: v33, F1: v33}
                                    _ = v41
                                case 9223372036854775807:
                                    if (42 == -50) {
                                        var v42 int32 = -1
                                        v43 := GenS3{F0: int16(v33), F1: v42, F2: 127}
                                        _ = v42
                                        _ = v43
                                    }
                                default:
                                    switch (fi31 / (((-(fi24))) | 1)) {
                                    case -2147483648:
                                        var v44 float32 = 3.14159
                                        _ = v44
                                    case 4294967295:
                                        fmt.Println(fi24, fi31)
                                    default:
                                        var v45 uint8 = 0
                                        _ = v45
                                    }
                                }
                            }
                            _ = v38
                        } else {
                            v4 += 4294967295
                            v46 := GenS8{F0: true, F1: v2}
                            _ = v46
                        }
                        var v47 int8 = 116
                        v3 -= 610
                        _ = v47
                    }
                    if true {
                        if false {
                            v33 -= genFn0(758, 1, int64(v2))
                            switch "foo bar" {
                            case "hello":
                                if len(v1) > 0 {
                                    v1[0] = "\n\t"
                                }
                            case "\n\t":
                                fmt.Println(v4, fi31)
                            case "":
                                v3 = (v3 - v4)
                            default:
                                if (false && true) {
                                    switch 4294967295 {
                                    case 133:
                                        if false {
                                            var v48 uint32 = 2
                                            var v49 uint64 = (v2 / ((649) | 1))
                                            if true {
                                                switch "" {
                                                case "\n\t":
                                                    v4 += v3
                                                default:
                                                    var v50 int16 = 42
                                                    _ = v50
                                                }
                                                if (v33 > v33) {
                                                    v3 += 4294967295
                                                } else {
                                                    v51 := GenS9{F0: (v3 - (v48 - v4))}
                                                    _ = v51
                                                }
                                                v5[(v3 & 156)] = (fi24 | (fi31 & fi31))
                                                var v52 bool = (-128 == 42)
                                                _ = v52
                                            } else {
                                                v49 -= (^(v2))
                                            }
                                            _ = v48
                                            _ = v49
                                        }
                                    case -88:
                                        switch (fi24 & -588) {
                                        case 32768:
                                            v2 += v2
                                        }
                                    default:
                                        var v53 uint16 = 255
                                        _ = v53
                                    }
                                    var v54 int16 = 1
                                    _ = v54
                                }
                            }
                            var v55 uint64 = v2
                            v56 := GenS2{F0: 2147483647}
                            _ = v55
                            _ = v56
                        }
                        v4 -= 161
                        switch fi24 {
                        case -2147483648:
                            switch (fi31 / (((fi24 % (((fi31 | fi24)) | 1))) | 1)) {
                            case -2147483648:
                                if ((128 <= 255) || true) {
                                    var v57 uint16 = 555
                                    var v58 float64 = v12
                                    var v59 uint16 = 336
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                }
                            case 128:
                                fmt.Println(v2, fi31, v12)
                                _ = len(v5)
                                fmt.Println(v12)
                                var v60 bool = false
                                var v61 uint16 = 32767
                                _ = v60
                                _ = v61
                            case 84:
                                v12 += v12
                            default:
                                v62 := GenS8{F0: true, F1: (v2 % ((981) | 1))}
                                _ = v62
                            }
                        }
                    } else {
                        if false {
                            v33 = v33
                            v2 = v2
                            switch 495 {
                            case -2147483648:
                                v63 := GenS6{F0: fi24, F1: v2}
                                _ = v63
                            case 119:
                                var v64 int64 = v33
                                _ = v64
                            default:
                                if true {
                                    v2 = (v2 & (^(v2)))
                                }
                            }
                        }
                    }
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            }
        }
        var v65 float32 = ((-2.25 / 1e10) + 1e10)
        var v66 int = (fi24 & (fi24 / (((fi24 ^ -2147483648)) | 1)))
        _ = fi24
        _ = v65
        _ = v66
    }
    var v67 string = "hello"
    if (false && false) {
        v2 = (^(v2))
        switch ((v67 + v67) + v67) {
        case "":
            var v68 bool = false
            _ = v68
        case "foo bar":
            fmt.Println(v67)
        case "hello":
            var v69 int32 = 778
            _ = v69
        }
        var v70 float32 = (0.1 + 1e10)
        _ = v70
    } else {
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        for fi71 := 0; fi71 < 7; fi71++ {
            v12 -= -2.25
            _ = fi71
        }
    }
    if ("\n\t" == v67) {
        for fi72 := 0; fi72 < 15; fi72++ {
            var v73 string = (v67 + v67)
            _ = fi72
            _ = v73
        }
        fmt.Println(v2)
    } else {
        var v74 float64 = v12
        var v75 uint16 = 0
        _ = v74
        _ = v75
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v67:", v67)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v67
}
