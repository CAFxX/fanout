package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 bool; F2 int32 }
type GenS3 struct { F0 string; F1 uint32; F2 int }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 bool; F2 string }
type GenS7 struct { F0 uint16; F1 int8 }
type GenS8 struct { F0 int16; F1 bool; F2 string }
type GenS9 struct { F0 uint; F1 bool }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS12 struct { F0 uint8 }

func main() {
    v1 := make(map[int]bool)
    var v2 bool = false
    var v3 string = ""
    var v4 string = v3
    if v2 {
        if false {
            v3 = "hello"
            var v5 uint8 = 255
            var v6 string = (v3 + ("" + ""))
            _ = v5
            _ = v6
        } else {
            if (871 == 570) {
                v4 = (v4 + v3)
            }
            switch 4294967295 {
            case 57:
                for fi7 := 0; fi7 < 19; fi7++ {
                    fmt.Println(v3, fi7)
                    _ = fi7
                }
            case 32767:
                fmt.Println(v2, v2)
            case 2147483647:
                var v8 uint = 255
                _ = v8
            default:
                var v9 int8 = 127
                _ = v9
            }
        }
        v10 := GenS1{F0: 923}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 1; fi11++ {
            if v2 {
                for fi12 := 0; fi12 < 19; fi12++ {
                    v4 = "hello"
                    var v13 int64 = -9223372036854775808
                    var v14 int32 = 394
                    if ((-13 != 100) || false) {
                        if v2 {
                            var v15 float32 = (-2.25 - -2.25)
                            var v16 uint = 480
                            _ = v15
                            _ = v16
                        }
                        var v17 int8 = 89
                        var v18 int = (fi12 - (fi12 * 128))
                        _ = v17
                        _ = v18
                    }
                    v19 := GenS2{F0: 848, F1: ((false && ((v14 - (v14 | v14)) == (v14 & -919))) && v2), F2: v14}
                    v20 := GenS1{F0: 42}
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v19
                    _ = v20
                }
                v2 = (!v2)
                for fi21 := 0; fi21 < 16; fi21++ {
                    if true {
                        v22 := GenS2{F0: -32768, F1: true, F2: 0}
                        v23 := []int{}
                        if (v2 && v2) {
                            if ((!v2) || false) {
                                v4 = "\n\t"
                                if (true && v2) {
                                    v1[fi11] = (65535 <= 42)
                                    var v24 int16 = 32767
                                    var v25 int16 = (v24 >> ((4294967295) & 63))
                                    var v26 uint64 = 42
                                    var v27 uint8 = 226
                                    var v28 int32 = 837
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                } else {
                                    v29 := GenS2{F0: -32768, F1: (2147483647 == 1), F2: -865}
                                    _ = v29
                                }
                            }
                            if (int64(fi21) < -1) {
                                v3 = ((v4 + v4) + (v4 + (v3 + v4)))
                            } else {
                                var v30 uint32 = 1
                                var v31 uint8 = 255
                                _ = v30
                                _ = v31
                            }
                            v3 = v3
                            var v32 uint64 = 189
                            if v2 {
                                if false {
                                    var v33 int = fi21
                                    if (fi21 <= (fi11 / ((fi21) | 1))) {
                                        v2 = (42 > 0)
                                        v34 := v1[fi21]
                                        var v35 uint16 = 2
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v2 = false
                                        v36 := v1[(fi21 ^ 0)]
                                        _ = v36
                                    }
                                    _ = v33
                                }
                            }
                            _ = v32
                        } else {
                            v3 = "a"
                            v23 = append(v23, fi21)
                        }
                        _ = v22
                        _ = v23
                    } else {
                        if false {
                            if true {
                                var v37 string = v3
                                var v38 float32 = (0.1 / -2.25)
                                _ = v37
                                _ = v38
                            } else {
                                v39 := GenS3{F0: "a", F1: 0, F2: fi21}
                                _ = v39
                            }
                            v40 := GenS4{F0: -128, F1: 42}
                            var v41 int64 = 0
                            v41 = -111
                            var v42 uint = 0
                            _ = v40
                            _ = v41
                            _ = v42
                        } else {
                            switch v4 {
                            case "foo bar":
                                if v2 {
                                    var v43 uint32 = 0
                                    switch fi11 {
                                    case 42:
                                        var v44 int = fi21
                                        _ = v44
                                    case -369:
                                        v3 = ((v4 + "hello") + v4)
                                        switch (fi21 ^ fi11) {
                                        case 590:
                                            var v45 bool = false
                                            _ = v45
                                        case 2147483647:
                                            var v46 int32 = 2147483647
                                            _ = v46
                                        }
                                        _ = len(v1)
                                    case -884:
                                        var v47 uint64 = 4294967295
                                        _ = v47
                                    }
                                    v48 := []int{(fi21 + 631), fi21, fi11}
                                    _ = v43
                                    _ = v48
                                } else {
                                    fmt.Println(fi11, v4, v3)
                                    var v49 float32 = 1.5
                                    _ = v49
                                }
                            case "\n\t":
                                v4 = v4
                            case "hello":
                                var v50 int16 = 435
                                if v2 {
                                    if (0.0 > -2.25) {
                                        switch (v3 + (v4 + ("\n\t" + v3))) {
                                        case "\n\t":
                                            v50 = 1
                                        case "":
                                            var v51 float32 = (1.5 / -2.25)
                                            v2 = true
                                            v50 += 446
                                            v52 := GenS4{F0: 1, F1: 509}
                                            _ = v51
                                            _ = v52
                                        default:
                                            v53 := []int{fi21, 498, -9223372036854775808, (-(fi21))}
                                            _ = v53
                                        }
                                        v54 := []int{fi11}
                                        if (209 == 42) {
                                            v54 = append(v54, fi21)
                                            switch v4 {
                                            case "":
                                                if false {
                                                    v55 := GenS5{F0: -128}
                                                    var v56 string = v3
                                                    _ = v55
                                                    _ = v56
                                                }
                                            case "a":
                                                if (v3 >= (v4 + v4)) {
                                                    v3 = v3
                                                    if true {
                                                        fmt.Println(v4, v3)
                                                        v3 = (((v3 + v4) + ("" + v3)) + v3)
                                                    }
                                                }
                                                fmt.Println(v3, v3, fi21)
                                            case "hello":
                                                var v57 float64 = -2.25
                                                var v58 bool = true
                                                _ = v57
                                                _ = v58
                                            }
                                        } else {
                                            var v59 int64 = -1
                                            _ = len(v54)
                                            _ = v59
                                        }
                                        _ = v54
                                    } else {
                                        v2 = v2
                                        v60 := GenS6{F0: 200, F1: false, F2: "\n\t"}
                                        fmt.Println(v4, v4, v4)
                                        _ = v60
                                    }
                                    v2 = ((!true) || false)
                                    v61 := GenS7{F0: 255, F1: 100}
                                    _ = v61
                                }
                                v3 = "\n\t"
                                _ = v50
                            default:
                                v62 := GenS8{F0: 42, F1: (!false), F2: "a"}
                                _ = v62
                            }
                            v63 := GenS1{F0: 2}
                            if ((!((-112 < 42) && true)) || (1e10 == 0.1)) {
                                delete(v1, (fi11 ^ 42))
                                switch (-(fi21)) {
                                case 0:
                                    switch fi11 {
                                    case 127:
                                        switch fi11 {
                                        case -1:
                                            v4 = ("foo bar" + ((v4 + ("hello" + v3)) + v3))
                                        case 653:
                                            var v64 uint64 = 306
                                            _ = v64
                                        default:
                                            v65 := GenS9{F0: 65535, F1: v2}
                                            _ = v65
                                        }
                                    case 4294967295:
                                        var v66 int16 = -1
                                        _ = v66
                                    case -9223372036854775808:
                                        var v67 int16 = 595
                                        _ = v67
                                    default:
                                        if false {
                                            switch "" {
                                            case "\n\t":
                                                var v68 float32 = 0.0
                                                _ = v68
                                            default:
                                                fmt.Println(fi21, v4)
                                            }
                                            var v69 string = v3
                                            _ = v69
                                        }
                                    }
                                case 340:
                                    v70 := GenS10{F0: -373}
                                    _ = v70
                                default:
                                    var v71 uint32 = 65535
                                    _ = v71
                                }
                                fmt.Println(v3)
                            } else {
                                v4 = "hello"
                                var v72 bool = v2
                                _ = v72
                            }
                            _ = v63
                        }
                        v3 = v4
                    }
                    v1[(fi21 + (fi11 - -82))] = true
                    _ = fi21
                }
            } else {
                v73 := GenS10{F0: -133}
                _ = v73
            }
            for fi74 := 0; fi74 < 8; fi74++ {
                switch 727 {
                case 42:
                    v4 = (v3 + v3)
                case 1:
                    v75 := GenS11{F0: 2147483647, F1: 178, F2: 1}
                    _ = v75
                }
                _ = fi74
            }
            v4 = (((((v4 + (v4 + "a")) + ("" + v4)) + v4) + v4) + v4)
            _ = fi11
        }
        var v76 uint16 = 32767
        switch 32768 {
        case 32768:
            v76 -= (v76 >> ((409) & 63))
        }
        _ = v76
    }
    var v77 uint16 = 42
    var v78 int32 = -431
    for fi79 := 0; fi79 < 12; fi79++ {
        v80 := []int{fi79, 128}
        var v81 string = v4
        for fi82 := 0; fi82 < 9; fi82++ {
            switch (-(fi82)) {
            case 1:
                fmt.Println(v78, v78, v3)
            case 32767:
                switch (fi82 | 1) {
                case 42:
                    if len(v80) > 0 {
                        v80[0] = -92
                    }
                default:
                    var v83 float32 = ((0.1 * (3.14159 - 3.14159)) - 0.1)
                    _ = v83
                }
            case 2147483647:
                var v84 int8 = -53
                _ = v84
            }
            _ = len(v1)
            var v85 uint16 = v77
            if len(v80) >= 2 {
                v80 = v80[1:len(v80)]
            }
            _ = fi82
            _ = v85
        }
        _ = fi79
        _ = v80
        _ = v81
    }
    var v86 int32 = (v78 % (((^(v78))) | 1))
    v77 = (v77 % (((v77 / ((v77) | 1))) | 1))
    v2 = true
    v1[2147483647] = (v2 && v2)
    for fi87 := 0; fi87 < 2; fi87++ {
        v88 := GenS12{F0: 128}
        v89 := []int{}
        v4 = ""
        _ = fi87
        _ = v88
        _ = v89
    }
    fmt.Println(v2, v78, v2)
    v90 := []int{-150, 32768, 9223372036854775807}
    v77 -= (v77 - (v77 & v77))
    var v91 uint64 = 778
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v86:", v86)
    fmt.Println("slicelen v90:", len(v90), v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v77
    _ = v78
    _ = v86
    _ = v90
    _ = v91
}
