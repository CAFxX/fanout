package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int8; F1 string }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint16; F1 uint32 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 string; F1 int32 }
func genFn0(p0 string, p1 int) uint8 {
    var v6 uint32 = 912
    p0 = ""
    var v7 int8 = -128
    v6 = (^(v6))
    p0 = "a"
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 uint32) int16 {
    if (2147483647 < 520) {
        if (-2147483648 < 0) {
            v8 := GenS1{F0: false}
            var v9 bool = false
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: true}
            p0 = (p0 | 4294967295)
            switch "a" {
            case "":
                var v11 uint32 = (^(p0))
                _ = v11
            }
            _ = v10
        }
        if false {
            v12 := GenS2{F0: -749}
            v13 := GenS3{F0: 227, F1: 255}
            switch -395 {
            case 4294967295:
                p0 -= (p0 % ((891) | 1))
            case -110:
                switch 1 {
                case 1:
                    v14 := GenS3{F0: 1, F1: 4294967295}
                    _ = v14
                case 9223372036854775807:
                    p0 = 991
                case 32767:
                    p0 += p0
                default:
                    v15 := GenS2{F0: 768}
                    _ = v15
                }
            default:
                p0 += (p0 << ((18446744073709551615) & 63))
            }
            _ = v12
            _ = v13
        } else {
            var v16 int32 = 42
            _ = v16
        }
        switch -9223372036854775808 {
        case -1:
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 int8 = -128
                _ = fi17
                _ = v18
            }
        case -680:
            fmt.Println(p0)
        case -2147483648:
            p0 += p0
        }
    }
    if false {
        v19 := []int{-9223372036854775808, -184, 128}
        if false {
            var v20 uint8 = genFn0(("hello" + ""), 620)
            _ = v20
        } else {
            p0 -= (^(p0))
            var v21 int = 312
            _ = v21
        }
        if ("a" < "foo bar") {
            if ((p0 << ((721) & 63)) > 2147483648) {
                _ = len(v19)
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
            } else {
                var v22 uint16 = 803
                v22 += 2
                _ = v22
            }
        } else {
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            v23 := GenS4{F0: 48}
            var v24 int64 = 0
            _ = v23
            _ = v24
        }
        _ = v19
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        for fi26 := 0; fi26 < 16; fi26++ {
            v27 := GenS5{F0: -25, F1: true}
            var v28 uint8 = 2
            var v29 float32 = ((3.14159 / (1.5 * 0.1)) / 0.1)
            if (0 > -9223372036854775808) {
                v30 := []int{(fi25 | (fi26 ^ int(p0))), (^(fi25)), (^(fi25)), (fi26 & (fi26 - fi25))}
                if false {
                    v29 = float32((0.1 - 1.5))
                } else {
                    if ("" >= "hello") {
                        v29 -= v29
                        var v31 int = 127
                        var v32 bool = ((1.5 - 3.14159) != 3.14159)
                        v33 := GenS6{F0: 2147483647, F1: v28}
                        var v34 string = (("\n\t" + "a") + "foo bar")
                        _ = v31
                        _ = v32
                        _ = v33
                        _ = v34
                    } else {
                        var v35 uint = 18446744073709551615
                        _ = v35
                    }
                }
                var v36 float64 = (-2.25 - 1e10)
                _ = v30
                _ = v36
            } else {
                fmt.Println(v28)
                if true {
                    v37 := []int{42, 2147483647}
                    fmt.Println(fi26, v28)
                    _ = v37
                }
                fmt.Println(fi25, v29)
                switch ("hello" + ("\n\t" + "")) {
                case "foo bar":
                    var v38 string = ""
                    _ = v38
                case "":
                    var v39 uint64 = 1
                    _ = v39
                case "hello":
                    v28 = 0
                default:
                    v28 = (v28 % (((^(v28))) | 1))
                    v40 := GenS7{F0: ""}
                    v41 := GenS3{F0: 65535, F1: 0}
                    switch "a" {
                    case "a":
                        v42 := GenS4{F0: 1}
                        _ = v42
                    case "":
                        var v43 string = ((("hello" + "a") + ("\n\t" + (((("" + ((("a" + "a") + "\n\t") + "foo bar")) + "hello") + "hello") + "a"))) + "\n\t")
                        _ = v43
                    default:
                        switch "hello" {
                        case "hello":
                            v44 := GenS8{F0: 127, F1: "hello"}
                            _ = v44
                        default:
                            var v45 float32 = v29
                            _ = v45
                        }
                        if false {
                            if true {
                                switch fi26 {
                                case 941:
                                    var v46 float64 = 3.14159
                                    var v47 int64 = -1
                                    _ = v46
                                    _ = v47
                                case 33:
                                    var v48 uint8 = 0
                                    _ = v48
                                case -61:
                                    v49 := []int{fi26, fi26}
                                    _ = v49
                                default:
                                    fmt.Println(v29, v29)
                                }
                            } else {
                                fmt.Println(fi25)
                                p0 += (p0 % (((^(p0))) | 1))
                                var v50 uint8 = genFn0("foo bar", 499)
                                _ = v50
                            }
                            if (uint8(p0) != (v28 * v28)) {
                                var v51 uint16 = 415
                                var v52 uint64 = 137
                                _ = v51
                                _ = v52
                            }
                            var v53 float64 = 3.14159
                            v54 := []int{}
                            v28 = 42
                            v29 += 1e10
                            var v55 int = fi25
                            fmt.Println(fi25, v55)
                            _ = v53
                            _ = v54
                            _ = v55
                        }
                        v56 := GenS9{F0: 32767}
                        var v57 int32 = int32(p0)
                        _ = v56
                        _ = v57
                    }
                    var v58 float64 = (-2.25 + 3.14159)
                    v59 := []int{-41, (fi25 + (fi26 - (fi25 % ((fi26) | 1)))), fi26, 128}
                    v60 := GenS1{F0: true}
                    v58 = v58
                    _ = v40
                    _ = v41
                    _ = v58
                    _ = v59
                    _ = v60
                }
                v61 := GenS1{F0: (1 < uint16(p0))}
                if false {
                    fmt.Println(p0, p0)
                    if true {
                        v62 := GenS10{F0: 2, F1: (p0 + p0)}
                        _ = v62
                    } else {
                        v63 := []int{(fi25 | (fi26 | fi26)), (fi26 - fi25), (fi26 | -761)}
                        _ = v63
                    }
                    fmt.Println(v29, p0)
                } else {
                    var v64 string = "a"
                    _ = v64
                }
                _ = v61
            }
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
        if (190 != 75) {
            switch (fi25 >> ((117) & 63)) {
            case 1:
                fmt.Println(p0, fi25)
            default:
                v65 := GenS11{F0: 32767}
                _ = v65
            }
            if true {
                var v66 int32 = 1
                var v67 int8 = -87
                v68 := []int{(fi25 - 42), (fi25 >> ((18446744073709551615) & 63)), (fi25 >> ((42) & 63)), (fi25 / ((128) | 1))}
                _ = v66
                _ = v67
                _ = v68
            } else {
                v69 := []int{fi25, (^(fi25)), 241, 32767}
                fmt.Println(p0)
                fmt.Println(fi25, fi25)
                _ = v69
            }
            if false {
                var v70 uint = 42
                v71 := []int{fi25, fi25, (fi25 >> (((v70 * v70)) & 63)), (fi25 * fi25)}
                _ = v70
                _ = v71
            } else {
                p0 = 2
                p0 = (p0 / (((p0 & 2)) | 1))
                for fi72 := 0; fi72 < 6; fi72++ {
                    p0 = (p0 << ((1) & 63))
                    p0 += (p0 & p0)
                    p0 = p0
                    _ = fi72
                }
            }
        }
        switch fi25 {
        case -2147483648:
            for fi73 := 0; fi73 < 13; fi73++ {
                v74 := []int{fi73, fi73, (fi25 & (-(fi73)))}
                v75 := GenS1{F0: (!false)}
                _ = fi73
                _ = v74
                _ = v75
            }
        case 127:
            for fi76 := 0; fi76 < 13; fi76++ {
                p0 += p0
                var v77 uint32 = 65535
                _ = fi76
                _ = v77
            }
        }
        _ = fi25
    }
    for fi78 := 0; fi78 < 11; fi78++ {
        p0 += (p0 - p0)
        v79 := GenS2{F0: 1}
        p0 = p0
        _ = fi78
        _ = v79
    }
    _ = p0
    return -32768
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 45
    var v3 int = 2147483647
    v4 := make(map[string]uint32)
    var v5 int64 = 0
    var v80 int64 = v2
    v81 := []int{(v3 * 42), (v3 | -1)}
    if (3.14159 >= 3.14159) {
        if (v3 > (v3 - (^(v3)))) {
            v82 := v1[v3]
            _ = v82
        }
        v83 := GenS4{F0: 42}
        _ = v83
    }
    var v84 int64 = (v5 - (v5 << ((87) & 63)))
    var v85 float32 = 0.1
    switch "foo bar" {
    case "foo bar":
        if (99 > 91) {
            switch v3 {
            case 2147483647:
                var v86 uint8 = 255
                _ = v86
            case 585:
                fmt.Println(v80, v2, v85)
            }
        } else {
            fmt.Println(v5)
            v87 := GenS2{F0: 42}
            v1[-953] = 930
            _ = v87
        }
    default:
        v5 += 0
    }
    v88 := GenS12{F0: "a", F1: -468}
    var v89 int16 = 1000
    v90 := GenS7{F0: ""}
    v2 += v80
    for fi91 := 0; fi91 < 17; fi91++ {
        if true {
            v80 = (v80 | 42)
        }
        _ = fi91
    }
    v84 = (v5 / ((v80) | 1))
    v89 += 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    fmt.Println("slicelen v81:", len(v81), v81)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v80
    _ = v81
    _ = v84
    _ = v85
    _ = v88
    _ = v89
    _ = v90
}
