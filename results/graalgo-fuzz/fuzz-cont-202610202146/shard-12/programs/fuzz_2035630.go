package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 uint16; F2 int64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint8; F1 int64; F2 int64 }
type GenS9 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS10 struct { F0 string; F1 uint32; F2 int8 }
type GenS11 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS12 struct { F0 uint; F1 int; F2 int16 }
type GenS13 struct { F0 uint16; F1 int16; F2 int }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 562
    var v3 bool = false
    var v4 uint16 = 32767
    v5 := make(map[uint32]int)
    var v6 int64 = 42
    var v7 uint64 = 611
    if v3 {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 int = (fi8 % (((fi8 % (((^(fi8))) | 1))) | 1))
            switch (fi8 >> ((4294967295) & 63)) {
            case 42:
                v6 -= v6
            case 127:
                for fi10 := 0; fi10 < 11; fi10++ {
                    switch (("foo bar" + "") + ("a" + "")) {
                    case "hello":
                        delete(v1, fi10)
                    case "\n\t":
                        v9 -= fi8
                        var v11 uint32 = 2
                        _ = v11
                    case "a":
                        if v3 {
                            v7 -= 0
                        } else {
                            v12 := GenS1{F0: (v4 - 2), F1: "\n\t", F2: 61}
                            v4 += 1
                            v9 -= v9
                            _ = v12
                        }
                    default:
                        var v13 uint8 = 2
                        _ = v13
                    }
                    fmt.Println(v2, v2)
                    v14 := []int{fi10, (fi10 + (^(fi8)))}
                    _ = fi10
                    _ = v14
                }
            case -9223372036854775808:
                switch int(v6) {
                case 32768:
                    v15 := GenS2{F0: 75}
                    _ = v15
                case -9223372036854775808:
                    var v16 float64 = 0.0
                    _ = v16
                case 4294967295:
                    v7 = 153
                default:
                    for fi17 := 0; fi17 < 1; fi17++ {
                        if (1 > 0) {
                            v4 = 4
                            v18 := GenS1{F0: 751, F1: "\n\t", F2: 0}
                            if true {
                                if (v3 && (-1 >= v2)) {
                                    v4 += (v4 * 2)
                                    var v19 uint16 = v4
                                    var v20 int16 = 42
                                    v21 := GenS3{F0: 99}
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                } else {
                                    var v22 float32 = (-2.25 - (1.5 - -2.25))
                                    var v23 uint32 = 0
                                    switch v9 {
                                    case 32768:
                                        var v24 bool = true
                                        _ = v24
                                    case 0:
                                        v7 = 42
                                    default:
                                        if (v2 < v2) {
                                            fmt.Println(v7, v23, v2)
                                            var v25 uint8 = 55
                                            var v26 uint16 = 255
                                            _ = v25
                                            _ = v26
                                        }
                                    }
                                    _ = v22
                                    _ = v23
                                }
                                if v3 {
                                    v27 := GenS2{F0: 127}
                                    _ = v27
                                }
                            }
                            v28 := GenS3{F0: v7}
                            _ = v18
                            _ = v28
                        }
                        _ = fi17
                    }
                }
            }
            _ = fi8
            _ = v9
        }
        var v29 int64 = v6
        var v30 int32 = (v2 + (v2 & 0))
        _ = v29
        _ = v30
    } else {
        var v31 float64 = (((3.14159 / 1.0) * 1.5) + 0.0)
        if (v31 > v31) {
            var v32 int = 32768
            _ = v32
        } else {
            v2 -= (v2 % ((v2) | 1))
            if false {
                _ = len(v1)
                for fi33 := 0; fi33 < 14; fi33++ {
                    var v34 int32 = v2
                    _ = fi33
                    _ = v34
                }
            } else {
                fmt.Println(v31)
            }
            for fi35 := 0; fi35 < 13; fi35++ {
                var v36 float64 = 1.5
                var v37 int32 = v2
                v37 = (^(v2))
                _ = fi35
                _ = v36
                _ = v37
            }
        }
        v38 := GenS4{F0: (((("foo bar" + "\n\t") + "foo bar") + ("" + "\n\t")) + ("" + "")), F1: (-1 != int8(v4))}
        _ = v31
        _ = v38
    }
    v4 = v4
    fmt.Println(v6)
    var v39 bool = (!(v7 >= (v7 + v7)))
    if ((v2 ^ -404) >= v2) {
        var v40 uint = 65535
        for fi41 := 0; fi41 < 5; fi41++ {
            fmt.Println(v4, v3)
            v42 := []int{-388, (fi41 + fi41)}
            if false {
                v4 += (v4 << ((v40) & 63))
                if len(v42) > 0 {
                    v42[0] = fi41
                }
                if true {
                    fmt.Println(v2, v7, v3)
                    var v43 int = fi41
                    v44 := GenS5{F0: 88}
                    _ = v43
                    _ = v44
                } else {
                    v2 = v2
                }
            } else {
                for fi45 := 0; fi45 < 19; fi45++ {
                    v46 := GenS6{F0: 1, F1: v4, F2: -1}
                    _ = fi45
                    _ = v46
                }
            }
            _ = fi41
            _ = v42
        }
        var v47 int16 = 1000
        _ = v40
        _ = v47
    }
    v7 -= v7
    v48 := v1[9223372036854775807]
    var v49 int32 = 648
    for fi50 := 0; fi50 < 9; fi50++ {
        for fi51 := 0; fi51 < 18; fi51++ {
            if true {
                if true {
                    v52 := []int{4294967295, fi50}
                    var v53 float64 = 0.1
                    if false {
                        if true {
                            v52 = append(v52, 128)
                            v7 = v7
                            _ = len(v5)
                        } else {
                            if v39 {
                                switch (v48 + v48) {
                                case "foo bar":
                                    if v3 {
                                        switch v48 {
                                        case "":
                                            v54 := GenS7{F0: -32768}
                                            _ = v54
                                        case "a":
                                            if v3 {
                                                v55 := GenS1{F0: v4, F1: v48, F2: 255}
                                                var v56 bool = (!true)
                                                var v57 string = "a"
                                                _ = v55
                                                _ = v56
                                                _ = v57
                                            }
                                        case "foo bar":
                                            var v58 float32 = 3.14159
                                            _ = v58
                                        }
                                    } else {
                                        v2 -= v49
                                        _ = len(v5)
                                        var v59 int8 = 42
                                        _ = len(v52)
                                        _ = v59
                                    }
                                }
                                var v60 int = -421
                                v7 = v7
                                _ = v60
                            }
                        }
                        switch "hello" {
                        case "":
                            v48 = v48
                        case "hello":
                            switch "\n\t" {
                            case "\n\t":
                                v49 = -1
                            case "":
                                fmt.Println(v49, v48)
                            case "hello":
                                switch int(v2) {
                                case 1:
                                    fmt.Println(v6)
                                case 365:
                                    _ = len(v52)
                                case 128:
                                    switch v48 {
                                    case "":
                                        v4 = 909
                                    case "a":
                                        v52 = append(v52, -12)
                                    }
                                default:
                                    v53 = v53
                                }
                            }
                            switch fi51 {
                            case 1:
                                v61 := GenS7{F0: int16(fi50)}
                                _ = v61
                            case 42:
                                v62 := GenS8{F0: 75, F1: -1, F2: -408}
                                _ = v62
                            }
                        case "foo bar":
                            v4 = 181
                        }
                        fmt.Println(v39, v2, v4)
                    }
                    _ = v52
                    _ = v53
                }
                v63 := GenS9{F0: 335, F1: 42, F2: (v7 / (((v7 % (((v7 | v7)) | 1))) | 1))}
                v4 = v4
                _ = v63
            }
            var v64 int16 = 0
            switch fi51 {
            case 32768:
                if v39 {
                    v65 := GenS2{F0: 127}
                    switch 32767 {
                    case 32768:
                        if ((fi50 + (fi50 ^ (fi50 - (fi50 * (fi50 + -570))))) == 2147483647) {
                            var v66 bool = true
                            var v67 uint16 = (^(v4))
                            _ = v66
                            _ = v67
                        }
                        if v3 {
                            if (v3 || (fi51 >= (fi50 & fi51))) {
                                v68 := GenS8{F0: 2, F1: v6, F2: (-(v6))}
                                v7 = (v7 + 518)
                                if true {
                                    fmt.Println(fi50, v48, v48)
                                    v69 := []int{fi51}
                                    v7 -= 2
                                    _ = v69
                                }
                                var v70 uint = 65535
                                _ = v68
                                _ = v70
                            } else {
                                v2 = v49
                                var v71 string = "\n\t"
                                if v39 {
                                    v3 = (!(fi50 >= -2147483648))
                                    var v72 int32 = (v49 & int32(v4))
                                    if (v39 || false) {
                                        if (true && v3) {
                                            switch ("\n\t" + v71) {
                                            case "":
                                                v73 := []int{(-(fi50))}
                                                _ = v73
                                            case "\n\t":
                                                switch v71 {
                                                case "a":
                                                    v71 = "\n\t"
                                                case "":
                                                    v74 := GenS10{F0: v48, F1: 565, F2: 0}
                                                    fmt.Println(v72)
                                                    _ = v74
                                                case "\n\t":
                                                    v49 = 1
                                                default:
                                                    var v75 uint8 = 127
                                                    _ = v75
                                                }
                                            default:
                                                v5[4294967295] = (fi50 - 4294967295)
                                                v49 += (v2 / ((v49) | 1))
                                                var v76 uint32 = 333
                                                var v77 float32 = (1.5 - float32(1e10))
                                                switch (fi51 >> ((uint(v7)) & 63)) {
                                                case 2147483647:
                                                    v78 := GenS11{F0: uint32(v7), F1: uint(v4), F2: 1}
                                                    _ = v78
                                                }
                                                v3 = (!v39)
                                                _ = v76
                                                _ = v77
                                            }
                                            fmt.Println(v6, v64, v6)
                                        } else {
                                            var v79 uint8 = 189
                                            v80 := []int{int(v64), (fi50 + fi50), fi51, (fi50 * int(v7))}
                                            _ = v79
                                            _ = v80
                                        }
                                        v81 := GenS12{F0: 0, F1: (fi51 ^ fi51), F2: (-(v64))}
                                        _ = v81
                                    } else {
                                        v6 = v6
                                        v82 := GenS9{F0: (v64 | v64), F1: 127, F2: uint64(v4)}
                                        v48 = ("foo bar" + (("" + v71) + v48))
                                        _ = v82
                                    }
                                    v83 := []int{}
                                    v72 = (v49 << ((0) & 63))
                                    _ = v72
                                    _ = v83
                                } else {
                                    if true {
                                        fmt.Println(fi51, v2, v2)
                                    } else {
                                        v84 := GenS13{F0: (^(v4)), F1: v64, F2: (fi50 | (fi51 & (fi51 + (fi50 ^ 32767))))}
                                        v1[fi51] = v71
                                        _ = v84
                                    }
                                }
                                _ = v71
                            }
                            fmt.Println(v49)
                            var v85 uint8 = 59
                            _ = v85
                        } else {
                            var v86 int32 = (v49 / ((v49) | 1))
                            if (((!false) || v39) && (0.0 != (((0.0 * (0.1 * float64((3.14159 + 0.0)))) + 0.1) * (0.1 + float64(1e10))))) {
                                if ((-2.25 >= ((3.14159 / 1.5) - (0.0 / -2.25))) || v39) {
                                    v49 = (v49 << ((4294967295) & 63))
                                    switch (fi50 ^ (fi51 - fi50)) {
                                    case 4294967295:
                                        v48 = v48
                                    default:
                                        fmt.Println(v64)
                                    }
                                    v87 := GenS7{F0: -1}
                                    var v88 uint32 = 1
                                    _ = v87
                                    _ = v88
                                } else {
                                    v89 := []int{fi50, (fi51 % ((330) | 1)), fi50, (fi51 + (fi51 / ((int(v86)) | 1)))}
                                    if false {
                                        v39 = v3
                                        v39 = ((v49 * 0) != (v49 << ((4294967295) & 63)))
                                        v4 = v4
                                    }
                                    _ = v89
                                }
                            } else {
                                var v90 int = -9223372036854775808
                                _ = v90
                            }
                            _ = v86
                        }
                        var v91 int32 = v2
                        _ = v91
                    }
                    if true {
                        v92 := GenS7{F0: (v64 - v64)}
                        switch (fi51 - int(v6)) {
                        case 2147483647:
                            fmt.Println(v2, v64)
                        }
                        v3 = false
                        _ = v92
                    } else {
                        var v93 string = ("foo bar" + v48)
                        if true {
                            v94 := []int{(fi50 >> ((571) & 63)), fi50}
                            _ = v94
                        } else {
                            v95 := GenS7{F0: v64}
                            var v96 int8 = -1
                            var v97 float32 = 1e10
                            _ = v95
                            _ = v96
                            _ = v97
                        }
                        var v98 int8 = 42
                        _ = v93
                        _ = v98
                    }
                    _ = v65
                } else {
                    v99 := []int{fi50, fi51}
                    v39 = (v39 && (v48 >= v48))
                    var v100 int16 = v64
                    _ = v100
                    _ = v99
                }
            case 127:
                v6 += (v6 ^ -9223372036854775808)
            }
            _ = fi51
            _ = v64
        }
        _ = fi50
    }
    fmt.Println(v49, v6, v6)
    v6 = (^(v6))
    v101 := []int{-1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v101:", len(v101), v101)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
