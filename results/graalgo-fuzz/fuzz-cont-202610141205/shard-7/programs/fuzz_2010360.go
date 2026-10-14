package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS3 struct { F0 string; F1 int16; F2 int16 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 int; F1 string; F2 int8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int64, p1 int16) int64 {
    p1 = p1
    v6 := GenS1{F0: 42, F1: 42, F2: 0}
    v7 := []int{4294967295}
    var v8 uint64 = 634
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 int32) uint32 {
    if true {
        v9 := GenS1{F0: 128, F1: 32767, F2: 77}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := []int{(fi10 | fi10), (fi10 & (fi10 - 128))}
            v11 = append(v11, -990)
            _ = fi10
            _ = v11
        }
    }
    p0 -= (p0 - p0)
    p0 -= (-(p0))
    fmt.Println(p0)
    _ = p0
    return 190
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{4294967295, 42, 571, 298}
    var v3 bool = (-207 == -1)
    v4 := []uint32{42, 0, 2147483648}
    v5 := make(map[uint32]bool)
    if false {
        v3 = v3
        v3 = false
    }
    if len(v4) > 0 {
        v4[0] = 1
    }
    v3 = v3
    v12 := GenS1{F0: 255, F1: 2, F2: 42}
    v13 := GenS2{F0: 2, F1: 0, F2: 648}
    v3 = (v3 || v3)
    if (!("\n\t" < "foo bar")) {
        if (!v3) {
            if v3 {
                v3 = (v3 && (!v3))
            } else {
                var v14 int64 = 9223372036854775807
                v3 = v3
                _ = v14
            }
            _ = len(v2)
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            switch ("\n\t" + "hello") {
            case "a":
                switch "a" {
                case "hello":
                    if false {
                        for fi16 := 0; fi16 < 6; fi16++ {
                            var v17 uint64 = 816
                            _ = fi16
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 16; fi18++ {
                            switch "\n\t" {
                            case "a":
                                _ = len(v1)
                            case "foo bar":
                                fmt.Println(fi15)
                            default:
                                if true {
                                    if false {
                                        v19 := GenS1{F0: 0, F1: 429, F2: 1}
                                        var v20 uint32 = 448
                                        _ = v19
                                        _ = v20
                                    }
                                } else {
                                    fmt.Println(fi15, v3, fi18)
                                }
                            }
                            _ = fi18
                        }
                        if (genFn1(425) <= 42) {
                            if true {
                                _ = len(v2)
                                v3 = (1 <= 952)
                                for fi21 := 0; fi21 < 17; fi21++ {
                                    var v22 float64 = 0.0
                                    fmt.Println(fi21, v3)
                                    if (!(4294967295 == 832)) {
                                        v22 += 1e10
                                        var v23 int16 = -1
                                        fmt.Println(v3, fi21, v22)
                                        fmt.Println(fi15)
                                        _ = v23
                                    }
                                    _ = fi21
                                    _ = v22
                                }
                            }
                            for fi24 := 0; fi24 < 12; fi24++ {
                                var v25 int32 = 536
                                _ = fi24
                                _ = v25
                            }
                            if (519 > 4) {
                                if v3 {
                                    v3 = false
                                }
                                for fi26 := 0; fi26 < 2; fi26++ {
                                    if v3 {
                                        v3 = true
                                    } else {
                                        _ = len(v2)
                                        fmt.Println(fi15)
                                    }
                                    var v27 uint16 = 223
                                    _ = fi26
                                    _ = v27
                                }
                            }
                        }
                    }
                }
            case "foo bar":
                if (!true) {
                    fmt.Println(fi15)
                }
            case "hello":
                if v3 {
                    fmt.Println(v3, v3)
                }
            }
            for fi28 := 0; fi28 < 20; fi28++ {
                switch fi28 {
                case 128:
                    fmt.Println(fi28)
                case 246:
                    v3 = (!(1 > 509))
                }
                if (18446744073709551615 <= 2) {
                    var v29 float32 = 1.5
                    v29 += float32(1.5)
                    v30 := GenS2{F0: 0, F1: 42, F2: 1}
                    if (false || v3) {
                        v29 += v29
                        v3 = true
                        if true {
                            delete(v1, (("a" + "hello") + ""))
                        } else {
                            var v31 int8 = -119
                            v32 := GenS2{F0: 2, F1: -1, F2: 2}
                            var v33 int16 = -1
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                    } else {
                        v34 := GenS3{F0: "hello", F1: 32767, F2: 133}
                        if len(v4) > 0 {
                            v4[0] = 65535
                        }
                        if (4294967295 != 255) {
                            var v35 int16 = 1
                            _ = v35
                        } else {
                            var v36 uint64 = 4294967295
                            fmt.Println(fi15, fi15)
                            if ((3.14159 < (0.1 * -2.25)) || (!(1e10 != 3.14159))) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v37 := GenS1{F0: 243, F1: 65535, F2: 1}
                                _ = v37
                            } else {
                                var v38 float64 = 0.1
                                if false {
                                    v3 = v3
                                    var v39 int8 = 127
                                    if true {
                                        var v40 int32 = 113
                                        switch "foo bar" {
                                        case "hello":
                                            v38 = v38
                                        default:
                                            fmt.Println(fi15, v29, v36)
                                        }
                                        v41 := GenS4{F0: ("" + "a"), F1: 42}
                                        _ = v40
                                        _ = v41
                                    } else {
                                        switch (fi15 / ((-1) | 1)) {
                                        case -9223372036854775808:
                                            var v42 string = ""
                                            _ = v42
                                        case 115:
                                            v38 -= (v38 + v38)
                                        default:
                                            v29 -= v29
                                        }
                                        var v43 bool = v3
                                        var v44 bool = v43
                                        var v45 bool = (v39 <= 100)
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    }
                                    _ = v39
                                } else {
                                    var v46 string = "a"
                                    _ = v46
                                }
                                _ = v38
                            }
                            _ = v36
                        }
                        _ = v34
                    }
                    _ = v29
                    _ = v30
                } else {
                    switch fi28 {
                    case 42:
                        v4 = append(v4, 2147483648)
                    case 0:
                        fmt.Println(v3, fi28)
                    default:
                        v3 = ((-9223372036854775808 < 948) && (255 == 1))
                    }
                    fmt.Println(fi28, v3)
                    v3 = false
                }
                _ = fi28
            }
            _ = fi15
        }
        if v3 {
            v3 = v3
        }
    } else {
        v3 = (v3 || v3)
        _ = len(v1)
        if v3 {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            v47 := GenS5{F0: 127, F1: 0}
            var v48 uint32 = 635
            _ = v47
            _ = v48
        }
    }
    if (!v3) {
        var v49 string = ""
        fmt.Println(v49, v49)
        for fi50 := 0; fi50 < 5; fi50++ {
            var v51 int64 = 0
            var v52 int64 = v51
            if true {
                if v3 {
                    var v53 string = v49
                    for fi54 := 0; fi54 < 11; fi54++ {
                        fmt.Println(v49, fi50, fi50)
                        v51 += 42
                        v55 := GenS3{F0: "", F1: 752, F2: 1}
                        _ = fi54
                        _ = v55
                    }
                    var v56 uint64 = 0
                    _ = v53
                    _ = v56
                } else {
                    var v57 uint32 = 1
                    var v58 int64 = (v52 >> ((2) & 63))
                    v52 += v58
                    _ = v57
                    _ = v58
                }
                switch "hello" {
                case "\n\t":
                    v59 := GenS6{F0: fi50, F1: v49, F2: -128}
                    _ = v59
                default:
                    fmt.Println(v52)
                }
                switch "a" {
                case "hello":
                    for fi60 := 0; fi60 < 4; fi60++ {
                        fmt.Println(v49)
                        var v61 uint8 = 168
                        _ = fi60
                        _ = v61
                    }
                case "a":
                    var v62 string = (v49 + v49)
                    _ = v62
                }
            } else {
                fmt.Println(v49)
                v63 := GenS2{F0: 42, F1: 0, F2: 1}
                v52 -= (v51 + v51)
                _ = v63
            }
            _ = fi50
            _ = v51
            _ = v52
        }
        _ = v49
    } else {
        switch -233 {
        case 127:
            v3 = v3
        case 579:
            fmt.Println(v3)
        case 4294967295:
            var v64 uint64 = 1
            _ = v64
        default:
            switch ((("foo bar" + "foo bar") + "a") + ("" + "\n\t")) {
            case "":
                var v65 bool = v3
                _ = v65
            case "foo bar":
                for fi66 := 0; fi66 < 7; fi66++ {
                    switch fi66 {
                    case 128:
                        v67 := GenS4{F0: ((("\n\t" + "") + "hello") + ("a" + "\n\t")), F1: 924}
                        _ = v67
                    case -9223372036854775808:
                        fmt.Println(fi66, fi66)
                    case 244:
                        v3 = v3
                    }
                    _ = fi66
                }
            }
        }
        var v68 int8 = 42
        if false {
            v69 := v5[2147483647]
            var v70 float32 = ((0.0 * 0.0) - 0.0)
            _ = v69
            _ = v70
        }
        _ = v68
    }
    delete(v5, 953)
    var v71 uint16 = 2
    if (0 > 88) {
        _ = len(v1)
        v72 := GenS1{F0: 115, F1: 1, F2: (v71 & 621)}
        for fi73 := 0; fi73 < 12; fi73++ {
            fmt.Println(v3, fi73)
            var v74 uint64 = 4294967295
            _ = fi73
            _ = v74
        }
        _ = v72
    } else {
        _ = len(v1)
    }
    switch 128 {
    case -225:
        var v75 bool = v3
        _ = v75
    case 498:
        if (610 < 951) {
            var v76 int16 = 42
            _ = v76
        } else {
            var v77 float32 = 1e10
            _ = v77
        }
    }
    var v78 int64 = 42
    switch "hello" {
    case "foo bar":
        switch 493 {
        case -158:
            if true {
                switch 9223372036854775807 {
                case 128:
                    v78 -= genFn0((v78 * 42), 1000)
                case 2147483647:
                    v71 = (v71 | (v71 * (v71 >> ((102) & 63))))
                case -803:
                    v71 += v71
                default:
                    v71 = 65535
                }
                var v79 int32 = -1
                _ = v79
            } else {
                _ = len(v5)
                if ((v71 - 65535) != v71) {
                    for fi80 := 0; fi80 < 4; fi80++ {
                        var v81 int = fi80
                        _ = fi80
                        _ = v81
                    }
                    v82 := GenS2{F0: 612, F1: 2147483647, F2: (v71 / ((942) | 1))}
                    _ = v82
                }
            }
        case 4294967295:
            fmt.Println(v78, v78, v3)
        default:
            v83 := GenS3{F0: ("\n\t" + (("foo bar" + "") + "foo bar")), F1: -32768, F2: 670}
            _ = v83
        }
    default:
        var v84 bool = v3
        _ = v84
    }
    var v85 int16 = 1
    v86 := GenS7{F0: 42}
    var v87 uint64 = 18446744073709551615
    v88 := GenS8{F0: 525, F1: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v71:", v71)
    fmt.Println("v78:", v78)
    fmt.Println("v85:", v85)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v71
    _ = v78
    _ = v85
    _ = v86
    _ = v87
    _ = v88
}
