package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 uint32; F2 string }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 uint8; F1 uint32; F2 uint }
type GenS7 struct { F0 uint8; F1 uint8 }
type GenS8 struct { F0 int8; F1 int }
type GenS9 struct { F0 string; F1 uint16 }
type GenS10 struct { F0 uint8; F1 string }
type GenS11 struct { F0 uint8; F1 int16; F2 uint }
type GenS12 struct { F0 uint; F1 int32 }
type GenS13 struct { F0 string; F1 int8; F2 int32 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = -2.25
    var v3 uint64 = 0
    var v4 int8 = 42
    if ((^(v3)) < 2) {
        for fi5 := 0; fi5 < 6; fi5++ {
            v3 = (^(v3))
            v6 := GenS1{F0: 555, F1: fi5}
            var v7 int32 = 2147483647
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := GenS1{F0: 1, F1: 1}
        var v9 bool = true
        _ = v8
        _ = v9
    } else {
        v10 := []int{32767}
        v3 = (v3 << ((42) & 63))
        delete(v1, -9223372036854775808)
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v3 = 828
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := GenS1{F0: 0, F1: (fi11 % ((1) | 1))}
            v14 := []int{fi12, fi12, fi11}
            switch 127 {
            case 32767:
                if true {
                    var v15 bool = false
                    if v15 {
                        v16 := GenS1{F0: 255, F1: fi12}
                        _ = v16
                    }
                    var v17 uint64 = 2
                    _ = v15
                    _ = v17
                } else {
                    fmt.Println(fi12)
                }
                v18 := v1[fi12]
                _ = v18
            case -1:
                switch 32767 {
                case -2147483648:
                    if true {
                        var v19 int64 = 9223372036854775807
                        var v20 float64 = 0.0
                        var v21 int64 = (v19 | v19)
                        fmt.Println(v21)
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                }
            case -2147483648:
                if true {
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    _ = len(v14)
                    switch fi11 {
                    case 1:
                        switch ("foo bar" + ("a" + "hello")) {
                        case "a":
                            switch fi12 {
                            case -364:
                                _ = len(v1)
                            case 56:
                                fmt.Println(v4, v2, v4)
                            default:
                                var v22 uint8 = 29
                                _ = v22
                            }
                        case "hello":
                            v23 := GenS2{F0: -2147483648, F1: -1}
                            _ = v23
                        case "foo bar":
                            fmt.Println(v4, fi12, v2)
                        default:
                            if (1 != 172) {
                                v24 := GenS3{F0: "hello"}
                                var v25 int8 = (-(v4))
                                _ = v24
                                _ = v25
                            } else {
                                switch ((("a" + ("" + ("foo bar" + ((("\n\t" + "a") + "foo bar") + ("foo bar" + (("\n\t" + "") + ("hello" + "\n\t"))))))) + (("a" + ("hello" + ("" + (("a" + ("foo bar" + "hello")) + "\n\t")))) + "hello")) + ("foo bar" + "foo bar")) {
                                case "\n\t":
                                    if true {
                                        if true {
                                            if true {
                                                switch fi11 {
                                                case 799:
                                                    if (v4 != (v4 + (-(v4)))) {
                                                        var v26 float64 = (1.5 + (((0.0 * (float64(v2) / 1.0)) * -2.25) + (((0.0 * -2.25) / 1e10) * float64((3.14159 - v2)))))
                                                        v27 := GenS1{F0: 96, F1: (fi11 << ((4294967295) & 63))}
                                                        _ = v26
                                                        _ = v27
                                                    } else {
                                                        v3 -= v3
                                                    }
                                                case -310:
                                                    fmt.Println(v4)
                                                case 0:
                                                    switch "a" {
                                                    case "hello":
                                                        if true {
                                                            if true {
                                                                v2 -= 0.0
                                                                v4 += (v4 * (v4 & (v4 / ((v4) | 1))))
                                                                v2 = 1e10
                                                            } else {
                                                                var v28 uint16 = 42
                                                                _ = v28
                                                            }
                                                        } else {
                                                            v4 = v4
                                                            if true {
                                                                v29 := GenS3{F0: "hello"}
                                                                if len(v14) > 0 {
                                                                    v14[0] = 0
                                                                }
                                                                v4 += v4
                                                                _ = v29
                                                            }
                                                            _ = len(v14)
                                                            switch fi11 {
                                                            case -1:
                                                                if true {
                                                                    delete(v1, fi11)
                                                                    v14 = append(v14, 32767)
                                                                    var v30 int64 = 1
                                                                    _ = v30
                                                                } else {
                                                                    if false {
                                                                        var v31 float64 = float64(v2)
                                                                        _ = v31
                                                                    }
                                                                    switch ("\n\t" + "") {
                                                                    case "a":
                                                                        delete(v1, fi12)
                                                                    case "":
                                                                        v2 += v2
                                                                    case "\n\t":
                                                                        v14 = append(v14, 1)
                                                                    }
                                                                    if true {
                                                                        if true {
                                                                            var v32 bool = true
                                                                            _ = v32
                                                                        }
                                                                        v3 = (v3 % ((v3) | 1))
                                                                    } else {
                                                                        fmt.Println(v3)
                                                                        if ((0.1 + (-2.25 * (3.14159 + 0.1))) == -2.25) {
                                                                            if ((v4 / ((1) | 1)) != 42) {
                                                                                var v33 bool = (!true)
                                                                                v2 += v2
                                                                                if v33 {
                                                                                    v2 += v2
                                                                                }
                                                                                var v34 bool = v33
                                                                                _ = v33
                                                                                _ = v34
                                                                            }
                                                                            if false {
                                                                                v2 += v2
                                                                            }
                                                                            v35 := GenS4{F0: v4, F1: 2, F2: "\n\t"}
                                                                            var v36 int8 = 127
                                                                            _ = v35
                                                                            _ = v36
                                                                        }
                                                                        fmt.Println(fi12, fi12)
                                                                    }
                                                                }
                                                            case 1:
                                                                v3 += (v3 / (((v3 & 614)) | 1))
                                                            }
                                                        }
                                                    case "\n\t":
                                                        var v37 int64 = 560
                                                        switch "" {
                                                        case "\n\t":
                                                            if false {
                                                                var v38 int64 = v37
                                                                _ = v38
                                                            }
                                                        case "hello":
                                                            fmt.Println(fi12)
                                                        }
                                                        v39 := GenS3{F0: ""}
                                                        v4 = (v4 | (v4 & 100))
                                                        if (761 == 65535) {
                                                            v37 = (v37 / ((42) | 1))
                                                            v37 = (^(v37))
                                                            var v40 int = (fi11 | -809)
                                                            _ = v40
                                                        } else {
                                                            v3 += (v3 & 4294967295)
                                                            v41 := GenS2{F0: 0, F1: 129}
                                                            var v42 float64 = 3.14159
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                        v14 = append(v14, (fi11 ^ fi11))
                                                        _ = v37
                                                        _ = v39
                                                    case "foo bar":
                                                        if (false && true) {
                                                            if (676 > 65535) {
                                                                var v43 int16 = -148
                                                                if false {
                                                                    v44 := GenS5{F0: -1, F1: fi12}
                                                                    var v45 int32 = 0
                                                                    if (126 <= 42) {
                                                                        v2 += v2
                                                                        fmt.Println(fi12, v45)
                                                                        var v46 uint16 = 29
                                                                        var v47 string = "a"
                                                                        v48 := GenS6{F0: 202, F1: 0, F2: 1}
                                                                        _ = v46
                                                                        _ = v47
                                                                        _ = v48
                                                                    } else {
                                                                        v49 := GenS7{F0: 128, F1: 2}
                                                                        if (false || (true && (128 <= 245))) {
                                                                            v45 += 1
                                                                            if len(v14) >= 2 {
                                                                                v14 = v14[1:len(v14)]
                                                                            }
                                                                            var v50 int16 = 0
                                                                            v14 = append(v14, 0)
                                                                            _ = v50
                                                                        } else {
                                                                            v45 += v45
                                                                        }
                                                                        var v51 float64 = 1e10
                                                                        _ = v49
                                                                        _ = v51
                                                                    }
                                                                    _ = v44
                                                                    _ = v45
                                                                } else {
                                                                    var v52 float64 = 1.5
                                                                    _ = v52
                                                                }
                                                                switch (fi12 >> ((72) & 63)) {
                                                                case -9223372036854775808:
                                                                    v43 += (v43 - v43)
                                                                default:
                                                                    v53 := GenS8{F0: -1, F1: fi11}
                                                                    _ = v53
                                                                }
                                                                _ = v43
                                                            } else {
                                                                v54 := GenS4{F0: (v4 / (((v4 ^ -1)) | 1)), F1: 568, F2: ("" + ("" + "a"))}
                                                                v55 := GenS8{F0: -128, F1: 792}
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                            v14 = append(v14, (fi12 * 32767))
                                                            fmt.Println(fi12)
                                                        } else {
                                                            fmt.Println(v3, fi11, fi12)
                                                            v56 := GenS3{F0: ("a" + ("hello" + "\n\t"))}
                                                            _ = v56
                                                        }
                                                    default:
                                                        if true {
                                                            var v57 int64 = 1
                                                            _ = v57
                                                        } else {
                                                            fmt.Println(v3, fi12, fi11)
                                                        }
                                                    }
                                                }
                                            } else {
                                                v3 += 307
                                                if false {
                                                    var v58 int64 = 0
                                                    v59 := GenS9{F0: "", F1: 459}
                                                    _ = v58
                                                    _ = v59
                                                } else {
                                                    v60 := GenS10{F0: 127, F1: "\n\t"}
                                                    _ = v60
                                                }
                                            }
                                        } else {
                                            switch (fi12 % (((fi11 ^ -312)) | 1)) {
                                            case 4294967295:
                                                var v61 uint64 = (v3 << ((1) & 63))
                                                _ = v61
                                            case 32767:
                                                var v62 string = ("foo bar" + "hello")
                                                if len(v14) > 0 {
                                                    v14[0] = -752
                                                }
                                                switch v62 {
                                                case "":
                                                    v3 = (v3 & (v3 - v3))
                                                case "a":
                                                    if true {
                                                        fmt.Println(v4, v62, v2)
                                                        v3 -= v3
                                                    } else {
                                                        v63 := GenS4{F0: (v4 - (v4 % ((v4) | 1))), F1: 42, F2: v62}
                                                        v3 = v3
                                                        var v64 uint64 = v3
                                                        v4 = 0
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                case "hello":
                                                    v65 := GenS5{F0: -9223372036854775808, F1: (fi11 | fi11)}
                                                    _ = v65
                                                }
                                                _ = v62
                                            case -158:
                                                v66 := GenS2{F0: -189, F1: 633}
                                                _ = v66
                                            }
                                            v67 := GenS2{F0: 2147483647, F1: -1}
                                            switch 9223372036854775807 {
                                            case 32767:
                                                v2 -= 1e10
                                            case 1:
                                                fmt.Println(v2, fi12, v2)
                                            case 32768:
                                                var v68 uint = 4294967295
                                                _ = v68
                                            default:
                                                var v69 float64 = ((-2.25 / 0.1) / (3.14159 - 0.0))
                                                _ = v69
                                            }
                                            _ = v67
                                        }
                                        fmt.Println(fi11, v3, fi11)
                                    }
                                case "hello":
                                    switch -2147483648 {
                                    case 1:
                                        fmt.Println(fi11, v4)
                                    }
                                }
                                var v70 int32 = 2147483647
                                var v71 string = "a"
                                _ = v70
                                _ = v71
                            }
                        }
                    case 809:
                        switch 9223372036854775807 {
                        case 9223372036854775807:
                            switch 32767 {
                            case -479:
                                var v72 uint = 0
                                var v73 int16 = 1000
                                _ = v72
                                _ = v73
                            default:
                                v2 = v2
                            }
                        default:
                            if true {
                                v74 := GenS11{F0: uint8(v3), F1: 1000, F2: 488}
                                _ = v74
                            }
                        }
                    }
                }
            default:
                var v75 uint64 = (v3 / ((v3) | 1))
                _ = v75
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        for fi76 := 0; fi76 < 16; fi76++ {
            fmt.Println(v3, v2, fi11)
            if true {
                _ = len(v1)
                fmt.Println(v4)
            } else {
                v77 := []int{fi76, (fi76 ^ fi11)}
                _ = len(v77)
                var v78 uint8 = 188
                _ = v77
                _ = v78
            }
            if (0 > 1) {
                switch (fi11 | 128) {
                case -739:
                    fmt.Println(v3, v3, v2)
                default:
                    var v79 int16 = -1
                    _ = v79
                }
                var v80 uint8 = 42
                _ = v80
            } else {
                if ((!(!true)) || true) {
                    var v81 string = ("a" + "hello")
                    fmt.Println(v4, v2, v2)
                    _ = v81
                } else {
                    switch -9223372036854775808 {
                    case 32767:
                        _ = len(v1)
                    case 103:
                        if true {
                            v82 := GenS4{F0: v4, F1: 42, F2: (("" + ("" + "")) + "\n\t")}
                            v83 := GenS3{F0: (((("\n\t" + "") + ("foo bar" + ("hello" + ("hello" + ((("foo bar" + ("a" + (("\n\t" + ("" + "hello")) + "a"))) + "foo bar") + "a"))))) + ("hello" + ("foo bar" + "foo bar"))) + "a")}
                            _ = v82
                            _ = v83
                        }
                    }
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(fi11, fi76, v4)
                    case "foo bar":
                        v3 -= (v3 << ((18446744073709551615) & 63))
                        v4 = v4
                        v84 := GenS3{F0: ("\n\t" + "foo bar")}
                        fmt.Println(fi76, fi76, v3)
                        _ = v84
                    case "":
                        fmt.Println(v4)
                    default:
                        switch fi11 {
                        case 128:
                            v2 += 0.0
                            v3 = (v3 % (((v3 ^ (v3 % ((1) | 1)))) | 1))
                            var v85 int32 = 1
                            _ = v85
                        }
                    }
                }
                fmt.Println(v2, fi76, v3)
                var v86 bool = false
                _ = v86
            }
            _ = fi76
        }
        _ = fi11
    }
    v87 := GenS12{F0: 2, F1: 42}
    for fi88 := 0; fi88 < 10; fi88++ {
        var v89 int16 = 739
        var v90 uint = 65535
        v91 := []int{(fi88 * fi88)}
        _ = fi88
        _ = v89
        _ = v90
        _ = v91
    }
    for fi92 := 0; fi92 < 14; fi92++ {
        v3 = (^(v3))
        v93 := GenS11{F0: 2, F1: -574, F2: 4294967295}
        v3 -= (^(v3))
        _ = fi92
        _ = v93
    }
    var v94 int64 = -1
    switch -2147483648 {
    case 1:
        switch -1 {
        case 1:
            for fi95 := 0; fi95 < 9; fi95++ {
                var v96 int = (fi95 - (fi95 | (fi95 * (fi95 >> ((194) & 63)))))
                _ = fi95
                _ = v96
            }
        case 32767:
            v97 := []int{-337}
            _ = v97
        case 128:
            v3 -= (v3 & v3)
        }
    case 738:
        for fi98 := 0; fi98 < 5; fi98++ {
            v99 := GenS13{F0: ("hello" + "hello"), F1: v4, F2: 383}
            delete(v1, -2147483648)
            v2 = 1.5
            _ = fi98
            _ = v99
        }
    case -9223372036854775808:
        for fi100 := 0; fi100 < 18; fi100++ {
            var v101 int = fi100
            v2 = float32(-2.25)
            _ = fi100
            _ = v101
        }
    }
    var v102 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v102:", v102)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v87
    _ = v94
}
