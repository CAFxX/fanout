package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int16; F2 string }
type GenS3 struct { F0 uint8; F1 int; F2 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 uint8; F2 int64 }
type GenS6 struct { F0 int; F1 uint64; F2 int }
type GenS7 struct { F0 uint8 }

func main() {
    var v1 bool = true
    v2 := []int{-2147483648, -594}
    var v3 int32 = -2147483648
    v4 := make(map[string]string)
    var v5 string = ((("hello" + "\n\t") + ("hello" + ("\n\t" + "\n\t"))) + "a")
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 1; fi7++ {
        fmt.Println(v1, v1, v5)
        v8 := GenS1{F0: true}
        for fi9 := 0; fi9 < 6; fi9++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v5 = ("hello" + "\n\t")
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    if ((true || true) || (!v1)) {
        v10 := v4[""]
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 float64 = 3.14159
            _ = fi11
            _ = v12
        }
        if (1 < 32767) {
            if (81 < 340) {
                v13 := GenS1{F0: v1}
                v3 = v3
                _ = v13
            }
            var v14 int64 = -349
            _ = v14
        } else {
            switch 2147483647 {
            case 909:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 0:
                fmt.Println(v10, v5)
            default:
                var v15 string = ("foo bar" + ("foo bar" + (("foo bar" + ("\n\t" + "\n\t")) + "foo bar")))
                _ = v15
            }
            for fi16 := 0; fi16 < 20; fi16++ {
                for fi17 := 0; fi17 < 18; fi17++ {
                    v18 := GenS2{F0: -847, F1: -1, F2: ((("foo bar" + v5) + v5) + "foo bar")}
                    fmt.Println(v10, fi17, v3)
                    _ = fi17
                    _ = v18
                }
                v19 := GenS2{F0: fi16, F1: 42, F2: (v5 + "a")}
                _ = fi16
                _ = v19
            }
            if false {
                if v1 {
                    v20 := v4["a"]
                    switch v10 {
                    case "":
                        fmt.Println(v20, v1, v1)
                    case "hello":
                        v20 = v20
                    case "a":
                        if len(v2) > 0 {
                            v2[0] = -867
                        }
                    default:
                        var v21 int8 = 0
                        _ = v21
                    }
                    if (1 != 686) {
                        if (65535 < 947) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                    _ = v20
                }
                if v1 {
                    for fi22 := 0; fi22 < 5; fi22++ {
                        v23 := GenS3{F0: 109, F1: 9223372036854775807, F2: -41}
                        _ = fi22
                        _ = v23
                    }
                    v10 = "\n\t"
                    v24 := GenS2{F0: 0, F1: -138, F2: "\n\t"}
                    _ = v24
                } else {
                    v3 = (v3 | v3)
                }
            }
        }
        _ = v10
    }
    if v1 {
        if v1 {
            if v1 {
                var v25 int64 = -691
                _ = v25
            } else {
                for fi26 := 0; fi26 < 5; fi26++ {
                    switch -9223372036854775808 {
                    case 127:
                        v4[v5] = (v5 + "a")
                    default:
                        fmt.Println(v5, v1, v5)
                    }
                    v1 = (18446744073709551615 == 255)
                    _ = fi26
                }
                v5 = v5
                var v27 int8 = -128
                _ = v27
            }
            var v28 int32 = 1
            fmt.Println(v3)
            _ = v28
        }
    } else {
        var v29 int8 = 88
        v30 := GenS2{F0: -280, F1: 32767, F2: (v5 + (("hello" + v5) + ("hello" + "foo bar")))}
        _ = v29
        _ = v30
    }
    if v1 {
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 uint64 = 2
            var v33 uint16 = 816
            _ = fi31
            _ = v32
            _ = v33
        }
        v3 = (v3 ^ v3)
        for fi34 := 0; fi34 < 1; fi34++ {
            v35 := GenS2{F0: 0, F1: 480, F2: v5}
            switch ("foo bar" + (v5 + v5)) {
            case "a":
                if (float32(-2.25) >= ((((0.1 * float32(0.0)) / -2.25) / ((-2.25 + (0.1 / float32(float64(0.1)))) / 1.5)) / -2.25)) {
                    v36 := GenS2{F0: (fi34 + (fi34 >> ((444) & 63))), F1: -880, F2: (v5 + "a")}
                    _ = v36
                } else {
                    v1 = (255 != 0)
                    if v1 {
                        delete(v6, 161)
                        v6[856] = 700
                    } else {
                        v1 = true
                        if (2 >= 246) {
                            v6[65535] = 358
                        } else {
                            var v37 uint32 = 2
                            _ = v37
                        }
                    }
                }
            }
            _ = fi34
            _ = v35
        }
    } else {
        v1 = (!(2147483647 > 127))
        v1 = false
        var v38 uint64 = 707
        _ = v38
    }
    var v39 uint16 = 42
    v40 := GenS3{F0: 0, F1: 32768, F2: -1}
    v1 = (0.1 != (float64(3.14159) * ((-2.25 / 1e10) + float64(0.1))))
    v39 = (v39 % ((v39) | 1))
    for fi41 := 0; fi41 < 20; fi41++ {
        if v1 {
            var v42 uint64 = 2
            _ = v42
        }
        _ = fi41
    }
    if ((!((v39 % ((v39) | 1)) != (v39 % (((v39 ^ 255)) | 1)))) && v1) {
        v1 = v1
        if ((42 == 128) || (!(v5 != ""))) {
            v3 = v3
            if len(v2) > 0 {
                v2[0] = 602
            }
        } else {
            if v1 {
                var v43 bool = (954 >= 732)
                _ = v43
            }
            if (v1 && v1) {
                v44 := GenS2{F0: 289, F1: -319, F2: ("hello" + v5)}
                _ = v44
            } else {
                v3 += v3
                v39 -= 1
                var v45 int = -1
                _ = v45
            }
        }
    } else {
        v46 := GenS4{F0: 9223372036854775807}
        for fi47 := 0; fi47 < 4; fi47++ {
            var v48 uint = 273
            v1 = (-1 < -177)
            var v49 int16 = -23
            _ = fi47
            _ = v48
            _ = v49
        }
        _ = v46
    }
    var v50 bool = (860 > 18446744073709551615)
    _ = len(v2)
    if true {
        v51 := GenS4{F0: 9223372036854775807}
        _ = v51
    } else {
        for fi52 := 0; fi52 < 6; fi52++ {
            if (18446744073709551615 >= 42) {
                if (!true) {
                    fmt.Println(v50)
                    v4[v5] = v5
                } else {
                    var v53 string = "hello"
                    v5 = v53
                    for fi54 := 0; fi54 < 20; fi54++ {
                        v55 := GenS5{F0: v53, F1: uint8(v39), F2: 228}
                        if ((v50 || v1) && v50) {
                            var v56 uint64 = 0
                            switch 1 {
                            case -2147483648:
                                v39 = v39
                            case 1:
                                switch (v53 + ("\n\t" + "hello")) {
                                case "a":
                                    var v57 uint32 = 121
                                    _ = v57
                                case "foo bar":
                                    delete(v4, (v5 + v53))
                                default:
                                    var v58 int16 = 1
                                    _ = v58
                                }
                            case 606:
                                v6[0] = 1
                            }
                            if v50 {
                                if v50 {
                                    v59 := v4[v53]
                                    delete(v4, ("a" + (v59 + "a")))
                                    var v60 int16 = -1
                                    switch "a" {
                                    case "hello":
                                        v2 = append(v2, 0)
                                    default:
                                        if len(v2) > 0 {
                                            v2[0] = fi54
                                        }
                                    }
                                    _ = v59
                                    _ = v60
                                }
                                v61 := v6[435]
                                _ = v61
                            } else {
                                if (42 < -58) {
                                    v62 := GenS6{F0: 772, F1: 18446744073709551615, F2: fi54}
                                    _ = v62
                                } else {
                                    var v63 uint8 = 26
                                    var v64 int64 = -9223372036854775808
                                    _ = v63
                                    _ = v64
                                }
                                switch (fi52 - fi52) {
                                case 9223372036854775807:
                                    if false {
                                        switch ((v53 + v5) + "") {
                                        case "":
                                            v65 := GenS4{F0: 9223372036854775807}
                                            _ = v65
                                        }
                                        var v66 int64 = 0
                                        switch fi54 {
                                        case -2147483648:
                                            var v67 int16 = 32767
                                            _ = v67
                                        }
                                        _ = v66
                                    } else {
                                        fmt.Println(fi52, v3, v39)
                                        v50 = (!(!(687 > v56)))
                                        v5 = v53
                                        _ = len(v2)
                                    }
                                case 128:
                                    if (51 != 1) {
                                        v68 := GenS7{F0: 1}
                                        fmt.Println(v3, v39)
                                        var v69 uint8 = 42
                                        v3 += (v3 & v3)
                                        _ = len(v4)
                                        _ = v68
                                        _ = v69
                                    }
                                    var v70 int8 = 0
                                    _ = v70
                                default:
                                    v3 = (v3 << ((255) & 63))
                                }
                            }
                            _ = v56
                        } else {
                            v71 := GenS2{F0: fi54, F1: -79, F2: ("" + v53)}
                            var v72 int64 = 0
                            _ = v71
                            _ = v72
                        }
                        v1 = v50
                        _ = fi54
                        _ = v55
                    }
                    _ = v53
                }
            }
            _ = fi52
        }
        v50 = (!v50)
    }
    v73 := v4["\n\t"]
    delete(v6, 434)
    switch 32767 {
    case -746:
        v74 := GenS6{F0: 639, F1: 42, F2: 9223372036854775807}
        _ = v74
    case 5:
        if false {
            for fi75 := 0; fi75 < 19; fi75++ {
                v4[v5] = v5
                _ = fi75
            }
        }
    case 1:
        for fi76 := 0; fi76 < 14; fi76++ {
            if v50 {
                v77 := GenS2{F0: fi76, F1: 42, F2: v73}
                if (false && (128 < 128)) {
                    switch v73 {
                    case "a":
                        fmt.Println(v50, v1)
                    case "":
                        v5 = (v73 + ("a" + v73))
                    case "\n\t":
                        switch ("" + "\n\t") {
                        case "\n\t":
                            var v78 uint16 = 255
                            _ = v78
                        default:
                            var v79 int32 = (^(v3))
                            _ = v79
                        }
                    default:
                        v5 = v5
                    }
                    v5 = v73
                }
                _ = v77
            }
            _ = fi76
        }
    default:
        var v80 int = 9223372036854775807
        _ = v80
    }
    for fi81 := 0; fi81 < 5; fi81++ {
        var v82 int64 = -9223372036854775808
        fmt.Println(v82, fi81)
        var v83 float32 = -2.25
        _ = fi81
        _ = v82
        _ = v83
    }
    var v84 uint16 = v39
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v73:", v73)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v50
    _ = v6
    _ = v73
    _ = v84
}
