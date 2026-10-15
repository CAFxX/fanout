package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int32; F2 int8 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 int16; F1 uint16; F2 int32 }
type GenS7 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS8 struct { F0 int16; F1 uint32; F2 int32 }
type GenS9 struct { F0 uint; F1 int32; F2 int32 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint32; F1 string }
type GenS12 struct { F0 uint64; F1 int8 }
type GenS13 struct { F0 uint; F1 uint8; F2 int16 }
type GenS14 struct { F0 int8; F1 int32; F2 int }
func genFn0(p0 uint64) int8 {
    switch 32767 {
    case 42:
        var v6 uint64 = p0
        _ = v6
    case 32767:
        if (true || false) {
            if true {
                var v7 bool = (1 == 145)
                var v8 int32 = 42
                fmt.Println(p0, p0)
                _ = v7
                _ = v8
            } else {
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 bool = true
                    _ = fi9
                    _ = v10
                }
            }
        }
    case -2147483648:
        p0 += (p0 + p0)
    }
    v11 := GenS1{F0: 0}
    p0 -= p0
    v12 := []int{}
    _ = p0
    _ = v11
    _ = v12
    return -104
}
func genFn1(p0 uint64, p1 bool, p2 int8) bool {
    fmt.Println(p1, p1)
    switch "hello" {
    case "foo bar":
        v13 := GenS2{F0: true, F1: 4294967295, F2: p1}
        _ = v13
    case "hello":
        var v14 bool = p1
        _ = v14
    case "":
        if (3.14159 >= 1e10) {
            if p1 {
                if p1 {
                    p0 = (p0 + (p0 | (p0 - (p0 << ((0) & 63)))))
                    v15 := GenS1{F0: 34}
                    var v16 uint32 = 0
                    _ = v15
                    _ = v16
                } else {
                    v17 := GenS3{F0: 198}
                    var v18 float64 = 0.1
                    var v19 uint8 = 58
                    _ = v17
                    _ = v18
                    _ = v19
                }
            } else {
                fmt.Println(p1)
            }
            for fi20 := 0; fi20 < 8; fi20++ {
                for fi21 := 0; fi21 < 9; fi21++ {
                    if p1 {
                        switch ("" + "foo bar") {
                        case "\n\t":
                            fmt.Println(p2)
                            v22 := []int{fi21, fi21, fi21, (-(fi21))}
                            _ = v22
                        default:
                            var v23 uint16 = 870
                            _ = v23
                        }
                        switch "\n\t" {
                        case "foo bar":
                            if (true && p1) {
                                if p1 {
                                    p0 = (p0 | (p0 + 2))
                                    v24 := GenS4{F0: (fi20 >> ((760) & 63)), F1: 1, F2: (p2 * p2)}
                                    fmt.Println(fi21, p2, fi20)
                                    fmt.Println(fi21)
                                    _ = v24
                                }
                                var v25 int = (fi21 | -9223372036854775808)
                                _ = v25
                            }
                        case "hello":
                            fmt.Println(fi20)
                        case "a":
                            p1 = (!p1)
                        }
                        switch ("\n\t" + ("hello" + ("a" + ""))) {
                        case "a":
                            var v26 int8 = (p2 % (((-(p2))) | 1))
                            _ = v26
                        case "hello":
                            p0 -= (p0 & p0)
                        default:
                            p0 = p0
                        }
                        p1 = (-1 == genFn0((p0 & (^(p0)))))
                    } else {
                        if p1 {
                            if true {
                                var v27 uint32 = 65535
                                v27 = (v27 | v27)
                                _ = v27
                            }
                            var v28 float32 = (float32((0.1 - -2.25)) + 1.5)
                            v29 := GenS2{F0: true, F1: 1, F2: p1}
                            _ = v28
                            _ = v29
                        } else {
                            v30 := GenS1{F0: 214}
                            p0 = (p0 & p0)
                            var v31 string = ""
                            _ = v30
                            _ = v31
                        }
                        fmt.Println(fi21, fi21, fi21)
                        p2 = (p2 | (p2 - p2))
                    }
                    var v32 int8 = int8(fi21)
                    _ = fi21
                    _ = v32
                }
                var v33 string = ("\n\t" + "foo bar")
                if true {
                    v34 := GenS2{F0: p1, F1: 255, F2: true}
                    for fi35 := 0; fi35 < 15; fi35++ {
                        v36 := []int{fi20, (fi20 & fi20)}
                        var v37 uint8 = 179
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    _ = v34
                }
                _ = fi20
                _ = v33
            }
        }
    default:
        v38 := GenS3{F0: 255}
        _ = v38
    }
    v39 := GenS4{F0: -99, F1: 1, F2: (^(p2))}
    _ = p0
    _ = p1
    _ = p2
    _ = v39
    return (2147483648 >= 494)
}

func main() {
    var v1 uint8 = 23
    var v2 uint32 = 777
    var v3 int64 = 774
    v4 := []int{}
    var v5 int32 = 2147483647
    var v40 uint64 = 282
    if true {
        v41 := GenS5{F0: 510, F1: 32767}
        _ = v41
    }
    switch 9223372036854775807 {
    case -1:
        switch 2147483647 {
        case -36:
            for fi42 := 0; fi42 < 19; fi42++ {
                switch "hello" {
                case "":
                    var v43 bool = (!false)
                    _ = v43
                case "foo bar":
                    v4 = append(v4, (fi42 % (((fi42 | 128)) | 1)))
                case "hello":
                    if true {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        if false {
                            v4 = append(v4, fi42)
                            if true {
                                v44 := GenS2{F0: true, F1: 2, F2: false}
                                v4 = append(v4, fi42)
                                _ = v44
                            }
                            fmt.Println(v3, v5, v2)
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if false {
                                v1 -= 128
                                if (-125 == v5) {
                                    if (true || false) {
                                        var v45 int16 = 32767
                                        _ = v45
                                    } else {
                                        v46 := GenS1{F0: v1}
                                        v40 = v40
                                        v47 := GenS5{F0: uint32(v1), F1: 2}
                                        _ = v46
                                        _ = v47
                                    }
                                } else {
                                    var v48 float32 = (1e10 - 0.1)
                                    v5 = (-(v5))
                                    _ = v48
                                }
                            }
                            var v49 uint8 = uint8(v40)
                            _ = v49
                        }
                        v1 = (v1 & 255)
                        switch fi42 {
                        case 1:
                            switch "hello" {
                            case "":
                                for fi50 := 0; fi50 < 18; fi50++ {
                                    v51 := GenS6{F0: -323, F1: 42, F2: (v5 >> ((456) & 63))}
                                    if ((v40 / ((1) | 1)) != v40) {
                                        v52 := GenS7{F0: 922, F1: 2, F2: 423}
                                        v5 += v5
                                        _ = v52
                                    }
                                    if true {
                                        var v53 int8 = 0
                                        switch "foo bar" {
                                        case "\n\t":
                                            v54 := GenS8{F0: -647, F1: (v2 >> ((4294967295) & 63)), F2: v5}
                                            _ = v54
                                        case "":
                                            var v55 int16 = 0
                                            _ = v55
                                        case "a":
                                            var v56 uint32 = v2
                                            _ = v56
                                        default:
                                            v5 = v5
                                        }
                                        v5 = v5
                                        _ = v53
                                    } else {
                                        var v57 string = "foo bar"
                                        var v58 float32 = float32(((((0.1 * 1.5) / -2.25) * 1.5) / 1.0))
                                        v59 := GenS9{F0: 2, F1: -2147483648, F2: -1}
                                        _ = v57
                                        _ = v58
                                        _ = v59
                                    }
                                    _ = fi50
                                    _ = v51
                                }
                            case "foo bar":
                                v3 -= (v3 | v3)
                            case "hello":
                                for fi60 := 0; fi60 < 14; fi60++ {
                                    v1 = v1
                                    var v61 float64 = 0.0
                                    _ = fi60
                                    _ = v61
                                }
                            default:
                                for fi62 := 0; fi62 < 18; fi62++ {
                                    v2 = 0
                                    v63 := GenS10{F0: (("a" + "") + "")}
                                    if len(v4) > 0 {
                                        v4[0] = (fi62 / ((fi62) | 1))
                                    }
                                    v2 = 783
                                    _ = fi62
                                    _ = v63
                                }
                            }
                        case 128:
                            v40 += 18446744073709551615
                        default:
                            fmt.Println(v2)
                        }
                    }
                default:
                    v64 := GenS11{F0: (v2 * (^(v2))), F1: "hello"}
                    _ = v64
                }
                v3 -= -542
                _ = fi42
            }
        case 2147483647:
            fmt.Println(v2, v3)
        default:
            var v65 int16 = 1
            _ = v65
        }
    case -9223372036854775808:
        if true {
            v66 := GenS12{F0: (v40 + (v40 * (v40 >> ((42) & 63)))), F1: 42}
            _ = v66
        }
    case 0:
        v3 += (v3 ^ (v3 - (v3 & v3)))
    default:
        fmt.Println(v5, v1)
    }
    fmt.Println(v2)
    fmt.Println(v1, v1, v1)
    var v67 bool = (2147483647 < -1)
    if true {
        v68 := GenS13{F0: 460, F1: (v1 << ((0) & 63)), F2: -32768}
        _ = v68
    } else {
        v5 += v5
    }
    v67 = (v40 != (v40 ^ (v40 * 42)))
    v2 += 1
    v69 := GenS14{F0: -1, F1: -247, F2: -177}
    v70 := GenS12{F0: v40, F1: 24}
    for fi71 := 0; fi71 < 10; fi71++ {
        v1 += (^(v1))
        for fi72 := 0; fi72 < 19; fi72++ {
            if len(v4) > 0 {
                v4[0] = fi71
            }
            _ = fi72
        }
        var v73 int32 = -293
        _ = fi71
        _ = v73
    }
    if (0 >= (v40 + 835)) {
        var v74 bool = v67
        _ = v74
    } else {
        v67 = (!v67)
        v1 = uint8(v2)
    }
    var v75 int32 = v5
    v5 += (v75 - v75)
    v75 = 1
    v5 = (v75 & 444)
    if true {
        v67 = v67
        v1 = v1
    } else {
        switch "" {
        case "\n\t":
            switch -382 {
            case 32768:
                v75 = v5
            default:
                var v76 int8 = 100
                _ = v76
            }
        case "foo bar":
            v2 = v2
        case "":
            var v77 uint8 = v1
            _ = v77
        }
        for fi78 := 0; fi78 < 3; fi78++ {
            switch "a" {
            case "a":
                switch -1 {
                case 4294967295:
                    if (((14 > -86) || true) || v67) {
                        fmt.Println(v1)
                    }
                }
            }
            _ = fi78
        }
        switch 42 {
        case 127:
            var v79 uint32 = (v2 | 65535)
            _ = v79
        default:
            for fi80 := 0; fi80 < 12; fi80++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi81 := 0; fi81 < 11; fi81++ {
                    v75 -= (v75 << ((uint(v2)) & 63))
                    _ = fi81
                }
                _ = fi80
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v67
    _ = v69
    _ = v70
    _ = v75
}
