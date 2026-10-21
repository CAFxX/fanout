package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 int; F1 uint; F2 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint16; F1 int; F2 int16 }

func main() {
    var v1 uint = 0
    var v2 int16 = 815
    var v3 uint8 = 157
    var v4 bool = true
    var v5 uint = (v1 / (((v1 << ((v1) & 63))) | 1))
    v6 := make(map[int]bool)
    v4 = (-857 <= -998)
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 float64 = 1.5
            fmt.Println(fi7, fi8, v5)
            var v10 string = "hello"
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 uint32 = 0
        _ = fi7
        _ = v11
    }
    switch ("\n\t" + ("\n\t" + "")) {
    case "a":
        var v12 int8 = 1
        _ = v12
    case "":
        var v13 int64 = 1
        _ = v13
    case "hello":
        var v14 int64 = 1
        _ = v14
    default:
        v15 := GenS1{F0: ""}
        _ = v15
    }
    if true {
        for fi16 := 0; fi16 < 8; fi16++ {
            for fi17 := 0; fi17 < 8; fi17++ {
                v18 := []int{703}
                _ = fi17
                _ = v18
            }
            switch "" {
            case "\n\t":
                var v19 int32 = 42
                _ = v19
            case "a":
                v20 := []int{fi16, fi16}
                _ = v20
            default:
                if v4 {
                    v21 := GenS2{F0: 2147483647, F1: (v5 & v5)}
                    _ = v21
                } else {
                    v22 := GenS2{F0: uint32(v3), F1: v1}
                    _ = v22
                }
            }
            switch (fi16 % (((fi16 | fi16)) | 1)) {
            case 9223372036854775807:
                v23 := []int{-693}
                _ = v23
            case 702:
                switch (-(fi16)) {
                case 42:
                    var v24 uint64 = 1
                    _ = v24
                default:
                    for fi25 := 0; fi25 < 12; fi25++ {
                        fmt.Println(v1, fi16, v1)
                        v26 := []int{1}
                        v2 -= v2
                        _ = fi25
                        _ = v26
                    }
                }
            }
            _ = fi16
        }
    } else {
        v4 = true
        if v4 {
            for fi27 := 0; fi27 < 4; fi27++ {
                for fi28 := 0; fi28 < 1; fi28++ {
                    var v29 uint = 0
                    var v30 int = (fi27 & fi27)
                    v30 -= 32767
                    fmt.Println(v1, fi27)
                    v31 := []int{fi28, (fi27 | (v30 | 32767))}
                    v32 := v6[int(v2)]
                    _ = fi28
                    _ = v29
                    _ = v30
                    _ = v31
                    _ = v32
                }
                var v33 uint64 = 42
                fmt.Println(v3)
                _ = fi27
                _ = v33
            }
            switch "foo bar" {
            case "foo bar":
                if v4 {
                    var v34 uint = (v5 << ((2) & 63))
                    var v35 int = 765
                    _ = v34
                    _ = v35
                } else {
                    switch 0 {
                    case 2147483647:
                        var v36 uint8 = (^(v3))
                        _ = v36
                    case 837:
                        v3 = v3
                    case -1:
                        v5 -= (v1 & uint(v3))
                    default:
                        fmt.Println(v2, v5)
                    }
                }
            case "\n\t":
                for fi37 := 0; fi37 < 16; fi37++ {
                    var v38 string = (("foo bar" + "foo bar") + ("\n\t" + ("\n\t" + "")))
                    v1 -= v5
                    _ = fi37
                    _ = v38
                }
            }
        } else {
            v4 = v4
            var v39 float32 = 1e10
            _ = v39
        }
        if v4 {
            var v40 float32 = (-2.25 - -2.25)
            _ = v40
        } else {
            for fi41 := 0; fi41 < 10; fi41++ {
                if ((-2.25 != 0.1) || true) {
                    var v42 uint8 = 2
                    _ = v42
                } else {
                    switch (("hello" + "hello") + "foo bar") {
                    case "hello":
                        if (((!v4) || (136 == 332)) && v4) {
                            var v43 uint64 = 769
                            switch 9223372036854775807 {
                            case 863:
                                v44 := []int{}
                                _ = v44
                            case 127:
                                v45 := GenS3{F0: (65535 > 2147483647), F1: v43}
                                _ = v45
                            case 2147483647:
                                for fi46 := 0; fi46 < 1; fi46++ {
                                    v47 := GenS1{F0: "a"}
                                    var v48 uint16 = 255
                                    var v49 float32 = float32((1.5 + -2.25))
                                    _ = fi46
                                    _ = v47
                                    _ = v48
                                    _ = v49
                                }
                            default:
                                var v50 bool = v4
                                _ = v50
                            }
                            _ = v43
                        } else {
                            for fi51 := 0; fi51 < 4; fi51++ {
                                v52 := []int{fi41, (fi51 << ((v1) & 63)), 136, (fi41 * (fi41 % (((^(fi51))) | 1)))}
                                delete(v6, fi51)
                                var v53 int = (fi41 / ((-1) | 1))
                                _ = fi51
                                _ = v52
                                _ = v53
                            }
                            var v54 int16 = (v2 & v2)
                            var v55 uint8 = v3
                            _ = v54
                            _ = v55
                        }
                    case "a":
                        v4 = v4
                    default:
                        if false {
                            _ = len(v6)
                        }
                    }
                    for fi56 := 0; fi56 < 3; fi56++ {
                        if true {
                            v57 := GenS2{F0: 399, F1: v1}
                            var v58 int32 = 42
                            v59 := GenS3{F0: v4, F1: 2}
                            _ = v57
                            _ = v58
                            _ = v59
                        } else {
                            var v60 uint16 = 614
                            v2 = (v2 | v2)
                            _ = v60
                        }
                        delete(v6, -948)
                        v61 := GenS4{F0: v2}
                        v3 -= v3
                        if (0 == 65535) {
                            switch (fi56 ^ (fi41 << (((v1 % ((2) | 1))) & 63))) {
                            case -2147483648:
                                v62 := GenS3{F0: (!false), F1: 299}
                                _ = v62
                            default:
                                switch 96 {
                                case 9223372036854775807:
                                    if false {
                                        v63 := v6[(-(fi56))]
                                        var v64 uint16 = 42
                                        var v65 uint = v5
                                        _ = v63
                                        _ = v64
                                        _ = v65
                                    } else {
                                        switch "a" {
                                        case "hello":
                                            var v66 int = fi41
                                            _ = v66
                                        }
                                        v67 := GenS5{F0: -1, F1: v4}
                                        switch (("\n\t" + "a") + ("\n\t" + ("\n\t" + "\n\t"))) {
                                        case "foo bar":
                                            v68 := GenS1{F0: ""}
                                            v69 := GenS6{F0: (fi41 / ((fi56) | 1)), F1: v1, F2: (v3 + v3)}
                                            _ = v68
                                            _ = v69
                                        case "hello":
                                            v1 -= v1
                                            var v70 uint32 = 0
                                            _ = len(v6)
                                            _ = v70
                                        case "a":
                                            switch ("foo bar" + "a") {
                                            case "hello":
                                                delete(v6, fi41)
                                            case "foo bar":
                                                v71 := GenS7{F0: v1}
                                                _ = v71
                                            }
                                        default:
                                            var v72 int = fi41
                                            _ = v72
                                        }
                                        v73 := GenS4{F0: (v2 << (((v5 + 4294967295)) & 63))}
                                        _ = v67
                                        _ = v73
                                    }
                                case -2147483648:
                                    var v74 uint = (v5 << ((v5) & 63))
                                    _ = v74
                                case -9223372036854775808:
                                    v75 := []int{-2147483648, (fi41 ^ int(v3)), (^(fi56))}
                                    _ = v75
                                }
                            }
                        }
                        _ = fi56
                        _ = v61
                    }
                }
                if v4 {
                    var v76 int16 = v2
                    if (-70 == -10) {
                        for fi77 := 0; fi77 < 6; fi77++ {
                            var v78 string = "a"
                            v79 := GenS3{F0: (601 == 1), F1: 42}
                            v80 := []int{32768, fi41, fi77}
                            _ = fi77
                            _ = v78
                            _ = v79
                            _ = v80
                        }
                    }
                    _ = v76
                }
                _ = fi41
            }
            var v81 uint32 = 2
            _ = v81
        }
    }
    var v82 uint32 = 502
    var v83 bool = true
    var v84 string = ""
    v85 := GenS1{F0: (("a" + v84) + v84)}
    var v86 int8 = 1
    if v4 {
        v84 = v84
        v87 := []int{}
        _ = v87
    }
    switch 9223372036854775807 {
    case 32768:
        var v88 uint = (v1 % (((v1 / ((v1) | 1))) | 1))
        _ = v88
    default:
        var v89 int32 = 2147483647
        _ = v89
    }
    v90 := GenS8{F0: 443, F1: 269, F2: (v2 & v2)}
    v91 := GenS3{F0: false, F1: 4294967295}
    v5 = (v5 >> (((v5 | v1)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v82
    _ = v83
    _ = v84
    _ = v85
    _ = v86
    _ = v90
    _ = v91
}
