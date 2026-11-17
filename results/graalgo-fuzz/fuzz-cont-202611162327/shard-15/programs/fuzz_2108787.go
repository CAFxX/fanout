package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int8; F1 string; F2 int8 }
type GenS10 struct { F0 uint16; F1 int8; F2 int32 }
type GenS11 struct { F0 string; F1 uint16 }
func genFn0() bool {
    if true {
        var v4 uint16 = 32767
        v4 += 348
        v4 += (v4 / ((v4) | 1))
        _ = v4
    } else {
        var v5 uint16 = 32767
        _ = v5
    }
    if true {
        v6 := []int{9223372036854775807, 1}
        var v7 uint64 = 18446744073709551615
        _ = v6
        _ = v7
    }
    if false {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := []int{fi8, fi8, (fi8 / (((-(fi8))) | 1)), 32767}
            var v10 uint8 = 128
            switch ("a" + "foo bar") {
            case "a":
                v10 = v10
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := []int{592, 42, 0}
        _ = v11
    }
    v12 := GenS1{F0: 233}
    var v13 int32 = -767
    _ = v12
    _ = v13
    return (false || false)
}
func genFn1() bool {
    v14 := GenS2{F0: -2147483648, F1: 0}
    v15 := GenS3{F0: false, F1: 300}
    _ = v14
    _ = v15
    return (!true)
}

func main() {
    var v1 int = -551
    var v2 int16 = 32767
    var v3 string = ("foo bar" + "foo bar")
    for fi16 := 0; fi16 < 9; fi16++ {
        if false {
            v3 = ((v3 + v3) + v3)
            var v17 int8 = -128
            fmt.Println(v17)
            _ = v17
        }
        if false {
            if ((42 >= 42) && true) {
                if (false && ((-128 < 100) && false)) {
                    v18 := GenS3{F0: (-2147483648 != 0), F1: 9}
                    _ = v18
                }
                fmt.Println(v2, v3)
            } else {
                switch v3 {
                case "a":
                    var v19 uint8 = 0
                    _ = v19
                case "hello":
                    v20 := []int{-1, (v1 * (fi16 % (((^(v1))) | 1))), v1, 1}
                    _ = v20
                }
                switch fi16 {
                case 0:
                    switch v3 {
                    case "a":
                        fmt.Println(v1, fi16)
                    case "foo bar":
                        v21 := GenS4{F0: v3}
                        _ = v21
                    default:
                        for fi22 := 0; fi22 < 19; fi22++ {
                            v23 := GenS5{F0: 2147483647}
                            v24 := []int{(-(fi22)), -1}
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    }
                case -270:
                    var v25 int32 = 2147483647
                    _ = v25
                default:
                    v2 = v2
                }
            }
            for fi26 := 0; fi26 < 4; fi26++ {
                v2 = (v2 % ((v2) | 1))
                switch v1 {
                case 4294967295:
                    var v27 uint16 = 65535
                    _ = v27
                case 127:
                    var v28 int8 = -128
                    if true {
                        var v29 string = ""
                        _ = v29
                    }
                    var v30 bool = (1 == 42)
                    _ = v28
                    _ = v30
                }
                var v31 string = ((("a" + v3) + ((v3 + "") + v3)) + v3)
                _ = fi26
                _ = v31
            }
            v32 := GenS3{F0: false, F1: 574}
            _ = v32
        } else {
            if true {
                v33 := GenS3{F0: false, F1: 65535}
                v3 = (v3 + v3)
                _ = v33
            }
            var v34 string = v3
            var v35 uint32 = 0
            _ = v34
            _ = v35
        }
        _ = fi16
    }
    var v36 int32 = -2147483648
    if true {
        if false {
            v37 := GenS6{F0: v36}
            v2 = 32767
            switch v1 {
            case 42:
                fmt.Println(v1, v2, v2)
            case -841:
                v3 = (v3 + "a")
            default:
                v38 := []int{v1, (v1 / ((977) | 1)), v1, (v1 >> ((42) & 63))}
                _ = v38
            }
            _ = v37
        } else {
            v39 := GenS1{F0: 1}
            switch (-(v1)) {
            case -568:
                fmt.Println(v36)
            case 32767:
                v40 := GenS7{F0: 0, F1: 2, F2: 18446744073709551615}
                _ = v40
            case -134:
                var v41 uint8 = 255
                _ = v41
            }
            for fi42 := 0; fi42 < 13; fi42++ {
                v43 := []int{(fi42 / ((v1) | 1)), fi42, (fi42 + int(v36))}
                if len(v43) > 0 {
                    v43[0] = fi42
                }
                _ = fi42
                _ = v43
            }
            _ = v39
        }
        v44 := []int{32767}
        _ = v44
    }
    for fi45 := 0; fi45 < 13; fi45++ {
        switch v1 {
        case -1:
            var v46 uint32 = 65535
            _ = v46
        case 2147483647:
            switch (v1 >> ((919) & 63)) {
            case 0:
                var v47 string = ("hello" + v3)
                _ = v47
            }
        case 32767:
            for fi48 := 0; fi48 < 14; fi48++ {
                fmt.Println(fi45, fi48, v36)
                v49 := []int{fi45, (v1 & fi45)}
                var v50 uint32 = 2
                _ = fi48
                _ = v49
                _ = v50
            }
        default:
            for fi51 := 0; fi51 < 5; fi51++ {
                v1 += (fi45 ^ (v1 / ((fi45) | 1)))
                switch v3 {
                case "a":
                    if true {
                        var v52 string = v3
                        var v53 string = (v3 + ("" + "hello"))
                        var v54 int64 = -1
                        v55 := GenS5{F0: 42}
                        if (!("" > v3)) {
                            v56 := GenS5{F0: (v36 | 83)}
                            v57 := GenS8{F0: v36}
                            v58 := GenS3{F0: ("foo bar" != v52), F1: 634}
                            _ = v56
                            _ = v57
                            _ = v58
                        }
                        switch 128 {
                        case 32767:
                            v52 = "foo bar"
                        case 127:
                            if false {
                                v3 = (v3 + v3)
                            }
                        case 9223372036854775807:
                            var v59 int8 = -128
                            _ = v59
                        default:
                            v60 := GenS5{F0: (v36 / (((v36 % (((v36 ^ (v36 % ((v36) | 1)))) | 1))) | 1))}
                            _ = v60
                        }
                        _ = v52
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                case "":
                    if (!true) {
                        v61 := []int{(v1 - fi45), 42}
                        _ = v61
                    }
                case "\n\t":
                    var v62 float32 = 1e10
                    _ = v62
                default:
                    if false {
                        v2 = v2
                        fmt.Println(fi51, fi45, v2)
                    }
                }
                _ = fi51
            }
        }
        _ = fi45
    }
    if (!((((1.5 + 3.14159) * 0.0) / 0.1) == -2.25)) {
        v63 := GenS1{F0: 724}
        v36 = -1
        _ = v63
    } else {
        v64 := GenS9{F0: -128, F1: ((v3 + v3) + v3), F2: -80}
        _ = v64
    }
    v2 -= (v2 << ((2) & 63))
    if false {
        for fi65 := 0; fi65 < 18; fi65++ {
            if (0 > 65535) {
                var v66 uint8 = 127
                var v67 int64 = -394
                _ = v66
                _ = v67
            } else {
                for fi68 := 0; fi68 < 16; fi68++ {
                    switch (((v3 + v3) + "a") + "a") {
                    case "foo bar":
                        var v69 uint16 = 1
                        _ = v69
                    }
                    if true {
                        var v70 uint64 = 18446744073709551615
                        var v71 uint8 = 128
                        if (!true) {
                            fmt.Println(fi65)
                            switch (^(v1)) {
                            case -19:
                                v72 := GenS10{F0: 255, F1: 107, F2: -2147483648}
                                _ = v72
                            case 42:
                                var v73 int16 = 1000
                                _ = v73
                            }
                            var v74 int = (fi65 << ((4294967295) & 63))
                            switch ("\n\t" + "foo bar") {
                            case "a":
                                v75 := GenS3{F0: true, F1: 1}
                                _ = v75
                            }
                            _ = v74
                        } else {
                            v76 := GenS8{F0: v36}
                            v1 += fi65
                            if (true || false) {
                                var v77 uint64 = 236
                                v78 := []int{}
                                _ = v77
                                _ = v78
                            } else {
                                v79 := GenS6{F0: (^(v36))}
                                var v80 uint32 = 145
                                _ = v79
                                _ = v80
                            }
                            _ = v76
                        }
                        _ = v70
                        _ = v71
                    } else {
                        switch v3 {
                        case "\n\t":
                            v2 = v2
                        case "a":
                            var v81 string = v3
                            _ = v81
                        default:
                            v2 = v2
                        }
                        fmt.Println(v2)
                        v82 := []int{int(v2), 2147483647}
                        _ = v82
                    }
                    v3 = v3
                    _ = fi68
                }
                var v83 uint16 = 901
                v84 := []int{(-(v1)), (fi65 % ((int(v36)) | 1))}
                _ = v83
                _ = v84
            }
            _ = fi65
        }
        v36 -= (v36 & (v36 << ((255) & 63)))
        fmt.Println(v36)
    }
    v36 = v36
    var v85 uint8 = 127
    v3 = v3
    for fi86 := 0; fi86 < 11; fi86++ {
        v85 = 128
        v87 := GenS4{F0: "foo bar"}
        switch v1 {
        case 328:
            for fi88 := 0; fi88 < 6; fi88++ {
                if (!false) {
                    v89 := []int{127, (v1 + (-(fi86))), 2147483647, fi88}
                    _ = v89
                } else {
                    fmt.Println(v1, fi86, v1)
                }
                _ = fi88
            }
        case -1:
            fmt.Println(v1)
        case 12:
            v90 := []int{573, (v1 % ((4294967295) | 1)), v1}
            _ = v90
        }
        _ = fi86
        _ = v87
    }
    switch v1 {
    case 9223372036854775807:
        if ((368 <= 0) || ((v36 & (v36 << ((0) & 63))) >= (v36 | v36))) {
            fmt.Println(v3)
            var v91 float64 = 1.5
            v92 := GenS11{F0: v3, F1: 32767}
            _ = v91
            _ = v92
        }
    default:
        var v93 bool = true
        _ = v93
    }
    fmt.Println(v3)
    v3 = v3
    if true {
        v94 := []int{(v1 & (-(v1))), int(v36), (^(v1))}
        var v95 bool = false
        _ = v94
        _ = v95
    }
    v96 := []int{v1, -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v85:", v85)
    fmt.Println("slicelen v96:", len(v96), v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v85
    _ = v96
}
