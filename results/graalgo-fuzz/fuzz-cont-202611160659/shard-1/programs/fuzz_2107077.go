package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint32; F1 bool; F2 uint }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 uint32; F1 uint8; F2 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 bool; F1 bool }
type GenS10 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int64; F1 bool }
type GenS13 struct { F0 int16 }
func genFn0(p0 int8, p1 int, p2 uint8) int16 {
    p1 += p1
    p1 -= (p1 - p1)
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 4294967295
    v3 := []string{"\n\t"}
    v4 := make(map[int]bool)
    var v5 string = "hello"
    var v6 uint64 = 2
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: (2147483647 > 0)}
        _ = fi7
        _ = v8
    }
    v6 = 2
    fmt.Println(v5)
    for fi9 := 0; fi9 < 7; fi9++ {
        _ = len(v1)
        var v10 bool = false
        _ = fi9
        _ = v10
    }
    var v11 string = (("foo bar" + (v5 + "\n\t")) + "")
    if true {
        for fi12 := 0; fi12 < 12; fi12++ {
            v4[fi12] = (false && (1 < 42))
            _ = fi12
        }
        if false {
            v2 -= v2
            if (0.1 == (1.5 / float64(1.5))) {
                for fi13 := 0; fi13 < 5; fi13++ {
                    var v14 string = "\n\t"
                    var v15 int8 = 47
                    v15 -= (v15 >> ((255) & 63))
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v16 := GenS2{F0: 678, F1: v5}
                switch 42 {
                case 0:
                    for fi17 := 0; fi17 < 14; fi17++ {
                        if (57 <= 255) {
                            v5 = (((v11 + "") + v5) + "a")
                        } else {
                            if (!(0 <= 0)) {
                                var v18 bool = false
                                v19 := GenS3{F0: v2, F1: true, F2: 65535}
                                switch (v11 + (("foo bar" + "foo bar") + v5)) {
                                case "":
                                    v18 = (1e10 > 1.5)
                                }
                                _ = v18
                                _ = v19
                            }
                        }
                        switch (fi17 << ((972) & 63)) {
                        case 32768:
                            fmt.Println(fi17)
                        case -2147483648:
                            v20 := GenS4{F0: ("hello" + ("\n\t" + (((v11 + v5) + "a") + ((v11 + v5) + "\n\t")))), F1: fi17}
                            _ = v20
                        default:
                            for fi21 := 0; fi21 < 15; fi21++ {
                                v6 += (v6 + v6)
                                switch (v5 + "\n\t") {
                                case "\n\t":
                                    _ = len(v4)
                                    _ = len(v1)
                                    var v22 uint64 = v6
                                    _ = v22
                                case "":
                                    v11 = v11
                                case "a":
                                    v2 -= 1
                                }
                                v5 = v11
                                var v23 bool = (true && ((181 >= 255) && true))
                                v23 = (false || (!((0 <= -1) && true)))
                                v11 = "hello"
                                _ = fi21
                                _ = v23
                            }
                        }
                        var v24 uint8 = 0
                        _ = fi17
                        _ = v24
                    }
                case -957:
                    if true {
                        var v25 float64 = 3.14159
                        var v26 uint16 = 42
                        v27 := GenS1{F0: (-2147483648 > 0)}
                        _ = v25
                        _ = v26
                        _ = v27
                    } else {
                        if (!(!true)) {
                            v6 = v6
                            switch (v5 + v11) {
                            case "foo bar":
                                var v28 float64 = (((0.1 * float64(-2.25)) / -2.25) * -2.25)
                                _ = v28
                            case "":
                                if false {
                                    v2 -= v2
                                    v29 := GenS5{F0: 343, F1: 130, F2: true}
                                    switch "\n\t" {
                                    case "":
                                        v30 := GenS5{F0: 874, F1: 121, F2: (0 <= 255)}
                                        _ = v30
                                    }
                                    _ = v29
                                }
                            case "\n\t":
                                v11 = ""
                            }
                            if true {
                                delete(v4, -525)
                                var v31 int8 = 1
                                _ = v31
                            }
                        }
                        v32 := GenS2{F0: 32767, F1: (v5 + v11)}
                        fmt.Println(v11, v5, v11)
                        _ = v32
                    }
                case 32768:
                    v6 += v6
                default:
                    if false {
                        if false {
                            var v33 int32 = 1
                            var v34 uint = 946
                            _ = v33
                            _ = v34
                        }
                    }
                }
                _ = v16
            }
            v11 = "foo bar"
        } else {
            for fi35 := 0; fi35 < 8; fi35++ {
                if false {
                    delete(v4, -868)
                    switch (fi35 >> ((2) & 63)) {
                    case -751:
                        var v36 bool = (false && (!false))
                        _ = v36
                    case 32768:
                        fmt.Println(v6)
                    case 378:
                        fmt.Println(fi35, fi35)
                    }
                }
                _ = fi35
            }
            _ = len(v1)
            v37 := GenS4{F0: v5, F1: -9223372036854775808}
            _ = v37
        }
        if ((^(v6)) > v6) {
            var v38 uint32 = v2
            switch 128 {
            case -895:
                fmt.Println(v11, v5)
            case 32767:
                v39 := GenS1{F0: (-9223372036854775808 != -1)}
                _ = v39
            }
            v40 := GenS4{F0: v11, F1: 42}
            _ = v38
            _ = v40
        } else {
            v2 -= v2
        }
    }
    var v41 float64 = -2.25
    _ = len(v3)
    var v42 uint16 = 46
    for fi43 := 0; fi43 < 7; fi43++ {
        if true {
            v44 := v4[int(v6)]
            var v45 float64 = float64(3.14159)
            if v44 {
                fmt.Println(v11, v45)
                var v46 uint16 = v42
                _ = v46
            }
            _ = v44
            _ = v45
        } else {
            fmt.Println(fi43, v5)
            v2 = (v2 << ((65535) & 63))
            switch fi43 {
            case 42:
                for fi47 := 0; fi47 < 7; fi47++ {
                    var v48 int32 = 0
                    if false {
                        v49 := GenS1{F0: (false && true)}
                        _ = v49
                    } else {
                        if true {
                            _ = len(v1)
                            var v50 int64 = 9223372036854775807
                            var v51 uint64 = (v6 % ((2) | 1))
                            v52 := GenS6{F0: 42}
                            _ = v50
                            _ = v51
                            _ = v52
                        } else {
                            v53 := GenS6{F0: (fi43 + 529)}
                            var v54 int32 = v48
                            if ((1e10 / 1.0) != 3.14159) {
                                v55 := GenS5{F0: 569, F1: 128, F2: (fi43 != fi47)}
                                var v56 float64 = v41
                                switch v5 {
                                case "":
                                    v41 = v56
                                case "\n\t":
                                    var v57 uint32 = v2
                                    _ = v57
                                }
                                if (v11 > v11) {
                                    var v58 int16 = 0
                                    v42 -= (v42 + v42)
                                    _ = v58
                                } else {
                                    _ = len(v3)
                                    v56 += 1.5
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch (fi47 << ((1) & 63)) {
                                    case 42:
                                        if true {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v59 int16 = 928
                                            _ = v59
                                        } else {
                                            v4[0] = false
                                            if (!(true && ((int(v42) > fi47) || (-1 != 127)))) {
                                                fmt.Println(v2)
                                            } else {
                                                v54 += (v54 >> ((255) & 63))
                                                switch ((v5 + "hello") + v11) {
                                                case "a":
                                                    v54 = v48
                                                    var v60 int32 = (-(v54))
                                                    _ = v60
                                                }
                                            }
                                        }
                                    case -1:
                                        _ = len(v4)
                                    case 127:
                                        var v61 uint32 = 571
                                        _ = v61
                                    default:
                                        switch fi43 {
                                        case 909:
                                            var v62 float32 = (-2.25 - 0.1)
                                            _ = v62
                                        case 9223372036854775807:
                                            if true {
                                                v54 = v48
                                                if ((624 != 415) || ("\n\t" > v11)) {
                                                    if true {
                                                        var v63 int64 = -9223372036854775808
                                                        fmt.Println(fi43, v63, v6)
                                                        if true {
                                                            var v64 uint32 = (^(v2))
                                                            _ = v64
                                                        } else {
                                                            if true {
                                                                if (v6 <= v6) {
                                                                    if true {
                                                                        v3 = append(v3, ((v5 + v5) + ("foo bar" + v11)))
                                                                        if false {
                                                                            var v65 int = -1
                                                                            _ = v65
                                                                        } else {
                                                                            v41 -= 1.5
                                                                        }
                                                                        v6 += v6
                                                                    } else {
                                                                        if (!(false && true)) {
                                                                            var v66 int16 = 1000
                                                                            _ = v66
                                                                        }
                                                                        v54 = (v54 / ((-1) | 1))
                                                                    }
                                                                    var v67 uint = 42
                                                                    _ = v67
                                                                }
                                                                v54 += v48
                                                                v11 = ((v5 + v5) + "")
                                                            }
                                                            if false {
                                                                v2 += (v2 & 4294967295)
                                                                v3 = append(v3, ("" + "\n\t"))
                                                                v68 := GenS7{F0: false}
                                                                _ = v68
                                                            } else {
                                                                v69 := GenS8{F0: true}
                                                                _ = v69
                                                            }
                                                            v70 := GenS9{F0: true, F1: false}
                                                            _ = v70
                                                        }
                                                        _ = v63
                                                    } else {
                                                        var v71 int8 = -128
                                                        v11 = v5
                                                        _ = v71
                                                    }
                                                    var v72 bool = (false || false)
                                                    v73 := GenS10{F0: v2, F1: (^(v6)), F2: (v42 - 971)}
                                                    v3 = append(v3, (((v11 + "") + "foo bar") + "\n\t"))
                                                    _ = v72
                                                    _ = v73
                                                }
                                            } else {
                                                var v74 int32 = (v48 >> ((255) & 63))
                                                if ((65535 >= 1) && (!false)) {
                                                    v5 = "foo bar"
                                                    fmt.Println(v5, v41)
                                                    v75 := GenS8{F0: false}
                                                    _ = v75
                                                }
                                                v76 := v4[fi47]
                                                _ = v74
                                                _ = v76
                                            }
                                        }
                                        v5 = "hello"
                                    }
                                    switch (fi43 - (fi47 | 9223372036854775807)) {
                                    case 44:
                                        fmt.Println(v48, v56, v41)
                                    case 56:
                                        v77 := GenS3{F0: 65535, F1: (true || true), F2: 633}
                                        _ = v77
                                    default:
                                        var v78 float64 = v56
                                        _ = v78
                                    }
                                }
                                v79 := GenS11{F0: (^(v2))}
                                _ = len(v4)
                                _ = v55
                                _ = v56
                                _ = v79
                            }
                            _ = v53
                            _ = v54
                        }
                        v80 := GenS12{F0: -1, F1: ((1000 > 0) || false)}
                        switch int(v48) {
                        case 127:
                            var v81 string = ("" + ((v5 + "") + "\n\t"))
                            _ = v81
                        }
                        _ = v80
                    }
                    fmt.Println(v48)
                    var v82 uint16 = v42
                    if (true || true) {
                        v3 = append(v3, "")
                        var v83 float64 = (v41 * (float64(0.1) + 1.5))
                        v2 = v2
                        _ = v83
                    }
                    _ = fi47
                    _ = v48
                    _ = v82
                }
            case -1:
                v84 := GenS11{F0: (^(v2))}
                _ = v84
            case 457:
                v85 := GenS2{F0: fi43, F1: v11}
                _ = v85
            default:
                v86 := GenS3{F0: 887, F1: (uint64(v2) <= (v6 << ((644) & 63))), F2: 42}
                _ = v86
            }
        }
        _ = fi43
    }
    var v87 uint32 = (v2 >> ((18446744073709551615) & 63))
    switch v11 {
    case "hello":
        v5 = v11
    default:
        var v88 float64 = v41
        _ = v88
    }
    v89 := GenS3{F0: v87, F1: (!false), F2: 1}
    var v90 float32 = 0.0
    v91 := GenS13{F0: 940}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v87:", v87)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v87
    _ = v89
    _ = v90
    _ = v91
}
