package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 int16; F1 string }
type GenS9 struct { F0 uint; F1 int64; F2 int32 }
type GenS10 struct { F0 uint32 }
func genFn0(p0 int64, p1 uint16) int {
    if false {
        fmt.Println(p0, p1)
        for fi6 := 0; fi6 < 1; fi6++ {
            for fi7 := 0; fi7 < 2; fi7++ {
                fmt.Println(p1, p1)
                if (false && ((true && (false || (-2147483648 != int32(p0)))) || false)) {
                    var v8 bool = false
                    p1 = 255
                    switch fi6 {
                    case 127:
                        switch "foo bar" {
                        case "a":
                            if (100 <= 100) {
                                switch ("" + (("foo bar" + "a") + ("hello" + ("\n\t" + "")))) {
                                case "foo bar":
                                    var v9 int32 = 1
                                    _ = v9
                                case "hello":
                                    p0 = (p0 * (p0 + -9223372036854775808))
                                case "\n\t":
                                    switch "" {
                                    case "foo bar":
                                        var v10 int32 = 2147483647
                                        _ = v10
                                    case "":
                                        v8 = v8
                                    }
                                }
                                v11 := GenS1{F0: 0, F1: 919, F2: -9223372036854775808}
                                _ = v11
                            } else {
                                v12 := GenS1{F0: 255, F1: 2, F2: (p0 | (p0 - p0))}
                                switch "foo bar" {
                                case "\n\t":
                                    v13 := []int{fi7, (fi7 >> ((18446744073709551615) & 63))}
                                    _ = v13
                                }
                                _ = v12
                            }
                        case "":
                            switch "\n\t" {
                            case "\n\t":
                                v14 := GenS2{F0: 4294967295}
                                _ = v14
                            case "hello":
                                p0 = (p0 ^ p0)
                            default:
                                v8 = ((!(v8 || (127 == 127))) && true)
                            }
                        }
                    default:
                        switch "a" {
                        case "\n\t":
                            fmt.Println(fi7, p1, fi7)
                        case "foo bar":
                            if true {
                                p0 = (p0 & p0)
                                if (v8 && (-63 >= (p0 - (p0 ^ p0)))) {
                                    p0 += (p0 - p0)
                                    var v15 uint16 = (p1 * 2)
                                    fmt.Println(fi7)
                                    var v16 float64 = -2.25
                                    var v17 int16 = -1
                                    switch "\n\t" {
                                    case "foo bar":
                                        switch (-(fi6)) {
                                        case 254:
                                            fmt.Println(p0, v15, fi7)
                                        case 32767:
                                            v18 := []int{}
                                            _ = v18
                                        }
                                    case "a":
                                        v19 := []int{(-(fi6)), int(v17), fi6}
                                        _ = v19
                                    case "\n\t":
                                        switch (fi7 % ((fi7) | 1)) {
                                        case 1:
                                            switch (fi6 - 937) {
                                            case 832:
                                                v17 += (v17 / ((v17) | 1))
                                            case 42:
                                                v20 := GenS3{F0: p0, F1: 128, F2: (v17 + int16(v15))}
                                                _ = v20
                                            default:
                                                v8 = (v8 && true)
                                            }
                                        default:
                                            p1 += (p1 | p1)
                                        }
                                    }
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                }
                                p0 += 0
                            } else {
                                v8 = (127 <= 2)
                                v21 := []int{(fi7 | fi6), (^(fi7)), (fi7 - (fi6 / ((fi6) | 1))), (^(fi6))}
                                p1 += (p1 + p1)
                                _ = v21
                            }
                        }
                        switch 1 {
                        case 32768:
                            if v8 {
                                var v22 uint8 = 255
                                fmt.Println(p1, v8, v22)
                                var v23 int32 = 1
                                _ = v22
                                _ = v23
                            } else {
                                p0 += p0
                                switch "hello" {
                                case "\n\t":
                                    v24 := []int{fi7, -2147483648, -1}
                                    _ = v24
                                case "foo bar":
                                    var v25 uint32 = 2147483647
                                    var v26 uint8 = 2
                                    _ = v25
                                    _ = v26
                                case "":
                                    fmt.Println(fi6, fi6)
                                }
                            }
                        case 4294967295:
                            v27 := []int{32768, (fi6 + fi7), 2147483647, fi7}
                            _ = v27
                        }
                        switch fi6 {
                        case 52:
                            var v28 string = (("foo bar" + "a") + (("" + ("" + "a")) + "foo bar"))
                            _ = v28
                        case 9223372036854775807:
                            var v29 uint = 18446744073709551615
                            _ = v29
                        case 128:
                            var v30 string = "a"
                            _ = v30
                        }
                        fmt.Println(fi6)
                    }
                    var v31 int64 = p0
                    _ = v31
                    _ = v8
                } else {
                    v32 := GenS4{F0: p1, F1: (fi7 / (((fi6 << ((592) & 63))) | 1))}
                    p1 += p1
                    p0 = (p0 & 59)
                    v33 := []int{9223372036854775807, (^(fi7)), fi7, (fi7 | fi6)}
                    _ = v32
                    _ = v33
                }
                _ = fi7
            }
            _ = fi6
        }
        p0 += (p0 % ((-9223372036854775808) | 1))
    } else {
        if true {
            v34 := GenS1{F0: 605, F1: 274, F2: (-(p0))}
            _ = v34
        }
        var v35 float32 = float32(0.1)
        _ = v35
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        v37 := []int{(-(fi36)), -1}
        p0 = -1
        _ = fi36
        _ = v37
    }
    var v38 string = "a"
    if (2147483647 <= -1) {
        if ((p0 & p0) != (-(p0))) {
            v39 := []int{-557, 127, 127, 127}
            _ = v39
        } else {
            switch ((v38 + (("" + "\n\t") + "")) + ("hello" + ("a" + v38))) {
            case "foo bar":
                if (false || false) {
                    for fi40 := 0; fi40 < 19; fi40++ {
                        for fi41 := 0; fi41 < 11; fi41++ {
                            v38 = (v38 + "foo bar")
                            p0 = (-(p0))
                            _ = fi41
                        }
                        fmt.Println(fi40, v38, p1)
                        var v42 float32 = 0.1
                        _ = fi40
                        _ = v42
                    }
                }
            case "":
                v38 = v38
            case "hello":
                var v43 int64 = (p0 ^ (p0 + (p0 - p0)))
                _ = v43
            default:
                v38 = ((v38 + "hello") + "foo bar")
            }
            if true {
                if (("foo bar" + (v38 + v38)) <= v38) {
                    switch 32767 {
                    case 127:
                        v44 := []int{int(p1), 128, 85, 369}
                        _ = v44
                    case 446:
                        v45 := []int{42}
                        _ = v45
                    case 386:
                        var v46 int = 511
                        _ = v46
                    default:
                        p0 = (p0 - (p0 + 42))
                    }
                    var v47 int = 0
                    _ = v47
                }
                var v48 uint64 = 4294967295
                _ = v48
            } else {
                p1 -= (p1 >> ((707) & 63))
                v38 = v38
                for fi49 := 0; fi49 < 2; fi49++ {
                    v50 := GenS5{F0: ("foo bar" + v38), F1: true}
                    v51 := []int{fi49, 1}
                    _ = fi49
                    _ = v50
                    _ = v51
                }
            }
            v52 := GenS6{F0: 4294967295}
            _ = v52
        }
        if true {
            if (693 < (p1 + (p1 & p1))) {
                var v53 uint = 448
                _ = v53
            }
            for fi54 := 0; fi54 < 13; fi54++ {
                v55 := GenS7{F0: fi54, F1: 1}
                _ = fi54
                _ = v55
            }
            v56 := []int{42}
            _ = v56
        } else {
            v57 := []int{9223372036854775807, 4294967295, 9223372036854775807}
            fmt.Println(v38, p1)
            fmt.Println(v38, p0, v38)
            _ = v57
        }
        var v58 uint32 = 4294967295
        _ = v58
    }
    v59 := []int{-126}
    _ = p0
    _ = p1
    _ = v38
    _ = v59
    return 9223372036854775807
}
func genFn1() int64 {
    var v60 uint16 = 255
    var v61 int64 = -508
    v62 := []int{}
    v61 = (v61 + -1)
    for fi63 := 0; fi63 < 6; fi63++ {
        for fi64 := 0; fi64 < 19; fi64++ {
            if false {
                if false {
                    v65 := GenS4{F0: 255, F1: (fi63 ^ (^(fi64)))}
                    if false {
                        if len(v62) >= 2 {
                            v62 = v62[1:len(v62)]
                        }
                        if true {
                            v61 = (^(v61))
                        }
                        v66 := GenS3{F0: (v61 / (((v61 - (-(v61)))) | 1)), F1: 255, F2: 0}
                        if false {
                            v60 -= 2
                            var v67 int64 = v61
                            _ = v67
                        } else {
                            var v68 int8 = -1
                            if len(v62) >= 2 {
                                v62 = v62[1:len(v62)]
                            }
                            v61 = -9223372036854775808
                            _ = v68
                        }
                        v60 = (v60 ^ (v60 + v60))
                        v60 -= 65535
                        _ = v66
                    } else {
                        var v69 uint32 = 0
                        var v70 int8 = -128
                        var v71 int32 = 1
                        switch (fi63 ^ (fi64 ^ (fi63 << ((2) & 63)))) {
                        case 88:
                            switch "\n\t" {
                            case "hello":
                                var v72 uint = 255
                                _ = v72
                            case "foo bar":
                                v62 = append(v62, -953)
                            case "a":
                                var v73 float32 = 3.14159
                                _ = v73
                            default:
                                v74 := GenS5{F0: "foo bar", F1: false}
                                if (true && true) {
                                    var v75 uint32 = (^(v69))
                                    _ = v75
                                } else {
                                    var v76 int = fi64
                                    var v77 int16 = 290
                                    _ = v76
                                    _ = v77
                                }
                                _ = v74
                            }
                        case 1:
                            v78 := GenS1{F0: 0, F1: 65535, F2: v61}
                            _ = v78
                        default:
                            var v79 int = (-(fi63))
                            _ = v79
                        }
                        _ = v69
                        _ = v70
                        _ = v71
                    }
                    fmt.Println(fi63, fi64, fi63)
                    v80 := GenS5{F0: "a", F1: false}
                    var v81 int16 = 42
                    _ = v65
                    _ = v80
                    _ = v81
                } else {
                    fmt.Println(v61, v60)
                    var v82 int64 = -499
                    fmt.Println(v61, fi64)
                    _ = v82
                }
                v60 += 65535
                v83 := GenS8{F0: -32768, F1: (("" + ((("\n\t" + "hello") + "foo bar") + (("" + "") + ((("hello" + "\n\t") + "hello") + "a")))) + "a")}
                _ = v83
            }
            switch fi63 {
            case -1:
                switch "\n\t" {
                case "a":
                    v60 = (v60 - 65535)
                }
            default:
                v60 = 42
            }
            fmt.Println(v60, fi64)
            v84 := GenS5{F0: ((("" + "foo bar") + (("hello" + "") + "a")) + ("" + (("\n\t" + (("a" + "foo bar") + "a")) + (("" + (("hello" + "a") + ("a" + ("hello" + "")))) + "")))), F1: false}
            v85 := GenS6{F0: 709}
            _ = fi64
            _ = v84
            _ = v85
        }
        _ = fi63
    }
    _ = v60
    _ = v61
    _ = v62
    return v61
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 0
    var v3 int32 = 2147483647
    var v4 bool = (1 <= 255)
    v5 := make(map[uint32]uint32)
    for fi86 := 0; fi86 < 10; fi86++ {
        var v87 float32 = 3.14159
        if false {
            fmt.Println(v3, fi86, fi86)
            var v88 uint8 = 127
            _ = v88
        }
        var v89 bool = (!false)
        _ = fi86
        _ = v87
        _ = v89
    }
    var v90 int32 = v3
    var v91 float32 = float32(0.1)
    v5[0] = 0
    v3 += v3
    v1[(("" + "\n\t") + "hello")] = -9223372036854775808
    v2 = (-(v2))
    v92 := GenS9{F0: 1, F1: -9223372036854775808, F2: v90}
    v93 := []int{}
    v94 := GenS10{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    fmt.Println("slicelen v93:", len(v93), v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v90
    _ = v91
    _ = v92
    _ = v93
    _ = v94
}
