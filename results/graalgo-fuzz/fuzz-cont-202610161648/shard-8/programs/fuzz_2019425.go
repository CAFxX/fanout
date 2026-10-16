package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 int64; F1 string; F2 uint16 }
type GenS8 struct { F0 int16; F1 uint32 }
type GenS9 struct { F0 int64; F1 bool; F2 int16 }
func genFn0(p0 int64) int8 {
    v5 := GenS1{F0: "\n\t"}
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            p0 += p0
            var v8 string = ("a" + ("a" + ("a" + "a")))
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 5; fi9++ {
            p0 -= (p0 | (p0 % (((p0 >> ((212) & 63))) | 1)))
            if ("\n\t" == "foo bar") {
                v10 := GenS1{F0: ""}
                _ = v10
            }
            if false {
                var v11 uint = 2
                v12 := GenS2{F0: 1, F1: 100}
                v13 := GenS3{F0: 1}
                if false {
                    if false {
                        var v14 string = ""
                        p0 = p0
                        v15 := GenS2{F0: 1, F1: 127}
                        switch (fi6 / ((1) | 1)) {
                        case 42:
                            switch "\n\t" {
                            case "hello":
                                v11 -= (v11 >> (((v11 * 1)) & 63))
                            case "foo bar":
                                var v16 int = -720
                                _ = v16
                            case "":
                                v17 := GenS2{F0: -966, F1: 1}
                                _ = v17
                            default:
                                fmt.Println(v14, p0)
                            }
                        default:
                            p0 -= p0
                        }
                        if true {
                            var v18 int8 = 100
                            switch ("hello" + "") {
                            case "hello":
                                fmt.Println(v14, fi6)
                            case "\n\t":
                                v11 = v11
                            }
                            var v19 float64 = 1.5
                            _ = v18
                            _ = v19
                        } else {
                            switch (fi6 >> ((342) & 63)) {
                            case -246:
                                v20 := []int{fi9}
                                var v21 uint64 = 4294967295
                                if true {
                                    p0 = p0
                                    if false {
                                        v22 := GenS1{F0: v14}
                                        _ = v22
                                    } else {
                                        v20 = append(v20, (fi9 >> ((v11) & 63)))
                                    }
                                    fmt.Println(p0, v11)
                                } else {
                                    var v23 bool = true
                                    var v24 uint = v11
                                    var v25 int8 = 127
                                    var v26 int8 = v25
                                    fmt.Println(v23)
                                    v20 = append(v20, fi6)
                                    v21 += (v21 - 18446744073709551615)
                                    v27 := GenS4{F0: 258}
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                                _ = v20
                                _ = v21
                            case 0:
                                v28 := GenS5{F0: (^(p0))}
                                _ = v28
                            case 748:
                                v11 += (v11 << ((923) & 63))
                            default:
                                var v29 int = fi6
                                _ = v29
                            }
                            v30 := GenS6{F0: v11, F1: 0}
                            _ = v30
                        }
                        var v31 string = v14
                        _ = v14
                        _ = v15
                        _ = v31
                    }
                } else {
                    if false {
                        v32 := GenS7{F0: p0, F1: "", F2: 1}
                        _ = v32
                    }
                    var v33 int8 = 1
                    _ = v33
                }
                _ = v11
                _ = v12
                _ = v13
            }
            var v34 uint32 = 49
            _ = fi9
            _ = v34
        }
        _ = fi6
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        for fi36 := 0; fi36 < 9; fi36++ {
            p0 = p0
            _ = fi36
        }
        fmt.Println(fi35)
        _ = fi35
    }
    _ = p0
    _ = v5
    return 0
}
func genFn1(p0 int64, p1 uint32, p2 int64) int {
    var v37 uint64 = 18446744073709551615
    p1 -= 42
    var v38 bool = false
    var v39 uint16 = 5
    _ = p0
    _ = p1
    _ = p2
    _ = v37
    _ = v38
    _ = v39
    return 32768
}

func main() {
    var v1 int = 32767
    v2 := make(map[uint32]bool)
    var v3 float64 = 0.0
    var v4 int16 = 335
    v4 = 696
    v40 := GenS1{F0: ("" + (("" + "hello") + (((("\n\t" + "hello") + ("\n\t" + "\n\t")) + "foo bar") + "")))}
    for fi41 := 0; fi41 < 13; fi41++ {
        v1 = fi41
        if false {
            v42 := v2[905]
            var v43 float32 = 1.5
            v44 := GenS8{F0: (v4 + -1), F1: 4294967295}
            _ = v42
            _ = v43
            _ = v44
        } else {
            fmt.Println(v3, v3)
            var v45 uint64 = 18446744073709551615
            _ = v45
        }
        var v46 uint64 = 0
        _ = fi41
        _ = v46
    }
    switch "a" {
    case "hello":
        if true {
            if (65535 < 2) {
                v47 := []int{v1}
                _ = v47
            } else {
                switch (v1 % ((42) | 1)) {
                case -529:
                    v1 = v1
                default:
                    v3 += v3
                }
            }
            fmt.Println(v4, v1, v1)
            if (255 <= 255) {
                var v48 uint64 = 1
                v3 += v3
                _ = v48
            } else {
                var v49 uint = 65535
                if false {
                    _ = len(v2)
                } else {
                    if (false || ((!(192 < 102)) && (1 <= 0))) {
                        v50 := []int{}
                        v51 := v2[2147483647]
                        _ = v50
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 15; fi52++ {
                        switch 835 {
                        case 42:
                            for fi53 := 0; fi53 < 3; fi53++ {
                                v49 -= 972
                                _ = fi53
                            }
                        default:
                            var v54 bool = true
                            _ = v54
                        }
                        if ((v1 ^ (fi52 - (v1 >> ((2) & 63)))) != (v1 ^ (-(fi52)))) {
                            v49 += v49
                        }
                        _ = fi52
                    }
                }
                _ = v49
            }
        }
    case "\n\t":
        if ((v1 ^ (v1 & genFn1(-1, 2, 9223372036854775807))) > v1) {
            var v55 uint64 = 2
            v3 += v3
            switch (v1 * 42) {
            case -432:
                v56 := GenS9{F0: 9223372036854775807, F1: false, F2: 0}
                _ = v56
            }
            _ = v55
        } else {
            switch (v1 ^ (v1 * (v1 << ((255) & 63)))) {
            case -2147483648:
                var v57 int32 = 42
                _ = v57
            case 32768:
                var v58 int16 = -32768
                _ = v58
            case 502:
                var v59 uint8 = 2
                _ = v59
            default:
                var v60 string = ""
                _ = v60
            }
        }
    }
    v3 = v3
    var v61 uint32 = 2147483647
    if (true && true) {
        fmt.Println(v61, v61, v3)
    }
    for fi62 := 0; fi62 < 5; fi62++ {
        switch 1 {
        case -9223372036854775808:
            var v63 int16 = 547
            _ = v63
        case -105:
            if false {
                for fi64 := 0; fi64 < 10; fi64++ {
                    v65 := v2[4294967295]
                    switch ("a" + "hello") {
                    case "hello":
                        var v66 uint32 = v61
                        _ = v66
                    case "\n\t":
                        if (((("hello" + (("foo bar" + "a") + "")) + (("foo bar" + "hello") + "foo bar")) + ("\n\t" + (("" + "a") + ("hello" + "foo bar")))) > "a") {
                            var v67 int8 = 106
                            switch "a" {
                            case "":
                                fmt.Println(fi62, fi62, v61)
                            default:
                                fmt.Println(v1)
                            }
                            fmt.Println(v65, v61, v67)
                            _ = v67
                        } else {
                            var v68 uint = 255
                            fmt.Println(fi62, fi64, v3)
                            delete(v2, 2147483648)
                            _ = v68
                        }
                    case "a":
                        v69 := GenS3{F0: v61}
                        _ = v69
                    }
                    _ = fi64
                    _ = v65
                }
                var v70 int8 = -1
                v4 = v4
                _ = v70
            }
        case 128:
            for fi71 := 0; fi71 < 3; fi71++ {
                var v72 bool = false
                v72 = (2 < 2)
                switch 1 {
                case 412:
                    var v73 uint8 = 0
                    _ = v73
                default:
                    if v72 {
                        var v74 bool = false
                        var v75 int16 = v4
                        if (939 != 18446744073709551615) {
                            v4 -= v75
                            var v76 int = fi71
                            v77 := GenS8{F0: 432, F1: v61}
                            _ = v76
                            _ = v77
                        } else {
                            if v72 {
                                fmt.Println(v3, v72)
                            } else {
                                fmt.Println(fi62, v3)
                                var v78 int = fi71
                                fmt.Println(v75)
                                _ = v78
                            }
                            v72 = v74
                            if (-127 > 105) {
                                v1 -= fi71
                                v61 += 42
                                fmt.Println(v3, v3, v4)
                            } else {
                                if true {
                                    if ((v61 * (^(v61))) >= (v61 >> ((331) & 63))) {
                                        if (((^(v61)) != (^(v61))) && v72) {
                                            v4 = v75
                                            fmt.Println(v72, v61, fi71)
                                            var v79 int16 = v4
                                            _ = v79
                                        }
                                    }
                                    var v80 int = (v1 / ((genFn1(-9223372036854775808, (v61 >> ((255) & 63)), -981)) | 1))
                                    fmt.Println(v74)
                                    v80 -= (fi71 & v80)
                                    _ = v80
                                }
                                if (!true) {
                                    if false {
                                        var v81 uint8 = 223
                                        switch int(v61) {
                                        case 32767:
                                            var v82 uint32 = (v61 * v61)
                                            _ = v82
                                        case 2147483647:
                                            if false {
                                                v83 := GenS2{F0: 1, F1: genFn0(1)}
                                                var v84 uint64 = 42
                                                _ = v83
                                                _ = v84
                                            } else {
                                                v85 := v2[v61]
                                                v86 := []int{(fi71 % (((fi62 << ((18446744073709551615) & 63))) | 1)), 128, 2147483647, (v1 * (fi71 - v1))}
                                                var v87 uint = 135
                                                _ = v85
                                                _ = v86
                                                _ = v87
                                            }
                                        case 1:
                                            v1 = (fi62 >> ((2) & 63))
                                        }
                                        _ = v81
                                    }
                                    var v88 float32 = (0.0 - float32(v3))
                                    v89 := GenS7{F0: 0, F1: ("a" + ("\n\t" + ("foo bar" + ("\n\t" + (((("\n\t" + "foo bar") + ("\n\t" + (("a" + "") + ("foo bar" + "")))) + "hello") + ""))))), F2: 42}
                                    _ = v88
                                    _ = v89
                                }
                            }
                        }
                        _ = v74
                        _ = v75
                    }
                }
                _ = fi71
                _ = v72
            }
        }
        switch ("foo bar" + "foo bar") {
        case "\n\t":
            var v90 string = ("hello" + "a")
            _ = v90
        }
        _ = fi62
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v61
}
