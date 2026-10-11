package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 uint32; F1 int; F2 int16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint; F1 uint32 }
type GenS10 struct { F0 uint16; F1 int8 }
type GenS11 struct { F0 int; F1 int8 }

func main() {
    var v1 float32 = float32((((0.0 * float64((float32((1.5 * (float64((3.14159 + ((0.0 * 0.0) / 1.0))) - 1.5))) / -2.25))) - (3.14159 - 1e10)) - 0.1))
    var v2 uint = 18446744073709551615
    var v3 uint8 = 0
    var v4 string = "foo bar"
    v5 := make(map[uint32]bool)
    var v6 int8 = -124
    var v7 uint16 = 637
    v8 := []int{-140, int(v3)}
    var v9 int32 = 0
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 uint8 = 255
        fmt.Println(v11)
        if (0 != (v6 >> (((v2 ^ 2)) & 63))) {
            v12 := GenS1{F0: false}
            if ((v7 / (((v7 << ((1) & 63))) | 1)) != (v7 / ((v7) | 1))) {
                fmt.Println(v9)
                for fi13 := 0; fi13 < 15; fi13++ {
                    if true {
                        if true {
                            v7 = (v7 % ((v7) | 1))
                            _ = len(v5)
                            var v14 int64 = 1
                            _ = v14
                        } else {
                            var v15 float32 = ((0.1 - 1.5) / ((v1 - ((0.1 / v1) - (v1 * 3.14159))) - v1))
                            _ = v15
                        }
                        fmt.Println(v2)
                        v16 := GenS2{F0: -1}
                        _ = v16
                    } else {
                        if false {
                            var v17 int = (fi13 + fi10)
                            v5[0] = true
                            _ = v17
                        } else {
                            fmt.Println(v11, v6)
                        }
                        if (42 == 32767) {
                            _ = len(v5)
                        }
                    }
                    _ = fi13
                }
            }
            var v18 float32 = v1
            _ = v12
            _ = v18
        } else {
            if true {
                for fi19 := 0; fi19 < 16; fi19++ {
                    var v20 int8 = v6
                    delete(v5, 1)
                    switch "hello" {
                    case "":
                        if false {
                            fmt.Println(fi10, v9, fi19)
                            switch (fi19 >> ((2) & 63)) {
                            case -615:
                                v21 := GenS2{F0: 672}
                                _ = v21
                            case 4294967295:
                                v22 := GenS1{F0: (!((!(18446744073709551615 > 0)) && false))}
                                _ = v22
                            case 127:
                                fmt.Println(v4)
                            }
                        } else {
                            fmt.Println(v20)
                            v23 := GenS3{F0: fi10, F1: (-1 <= -1), F2: 42}
                            v24 := GenS4{F0: v7}
                            _ = v23
                            _ = v24
                        }
                    default:
                        var v25 int16 = 0
                        _ = v25
                    }
                    _ = fi19
                    _ = v20
                }
            } else {
                switch "hello" {
                case "\n\t":
                    _ = len(v5)
                case "a":
                    fmt.Println(v3, v11, v9)
                case "":
                    if false {
                        if (float32(1e10) != v1) {
                            for fi26 := 0; fi26 < 1; fi26++ {
                                v2 = 18446744073709551615
                                v27 := GenS4{F0: v7}
                                var v28 bool = false
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                            v7 = 449
                        } else {
                            if (0 > 2147483648) {
                                v4 = v4
                            } else {
                                v4 = ((v4 + "foo bar") + v4)
                                var v29 uint64 = 0
                                _ = v29
                            }
                        }
                        var v30 int8 = v6
                        if false {
                            var v31 int32 = v9
                            v8 = append(v8, fi10)
                            _ = v31
                        } else {
                            _ = len(v5)
                            if false {
                                v30 = v6
                                for fi32 := 0; fi32 < 12; fi32++ {
                                    if len(v8) > 0 {
                                        v8[0] = fi32
                                    }
                                    _ = fi32
                                }
                            }
                        }
                        _ = v30
                    }
                }
                fmt.Println(v1)
                if ((true || (int64(v9) >= int64(v7))) && (!true)) {
                    switch (^(fi10)) {
                    case 4294967295:
                        v33 := GenS5{F0: (v6 * int8(v11))}
                        _ = v33
                    case -512:
                        v34 := GenS6{F0: (v3 > v11), F1: 0}
                        _ = v34
                    default:
                        var v35 uint8 = v3
                        _ = v35
                    }
                }
            }
        }
        _ = fi10
        _ = v11
    }
    if ((0.0 / v1) != (float32(-2.25) / v1)) {
        _ = len(v8)
        v6 += (v6 - (-(v6)))
        v9 += (^(v9))
    }
    var v36 uint32 = 1
    if len(v8) > 0 {
        v8[0] = 109
    }
    var v37 int = 32767
    v6 = v6
    if true {
        v3 = (v3 + v3)
        if true {
            var v38 int16 = 0
            v39 := GenS1{F0: (false || true)}
            _ = v38
            _ = v39
        } else {
            var v40 int16 = int16(v3)
            _ = v40
        }
        v41 := GenS4{F0: v7}
        _ = v41
    }
    for fi42 := 0; fi42 < 14; fi42++ {
        if (-798 != 0) {
            v7 += (v7 * v7)
            var v43 uint8 = v3
            _ = v43
        }
        switch v37 {
        case 1:
            var v44 float64 = float64(v1)
            _ = v44
        case -2147483648:
            v6 += (-(v6))
        case -1:
            var v45 int32 = 0
            _ = v45
        }
        for fi46 := 0; fi46 < 13; fi46++ {
            var v47 uint8 = v3
            if (v1 == float32(1e10)) {
                v4 = v4
                if true {
                    v5[(^(v36))] = true
                    fmt.Println(v2, v3)
                    v48 := GenS1{F0: (v47 != 255)}
                    v3 = (v3 - v47)
                    var v49 uint64 = 152
                    if (true || (-1 >= (v9 & (v9 % (((v9 << ((65535) & 63))) | 1))))) {
                        fmt.Println(v37, v9)
                    } else {
                        v4 = "\n\t"
                        if ((!((true && true) && (int64(v2) == 9223372036854775807))) && false) {
                            var v50 int = fi42
                            v49 += (v49 + (v49 * (v49 - (v49 / ((v49) | 1)))))
                            _ = v50
                        }
                        v4 = "\n\t"
                        var v51 uint64 = (v49 - (v49 | (^(v49))))
                        _ = v51
                    }
                    if true {
                        var v52 int32 = int32(v37)
                        v53 := GenS7{F0: (v36 >> ((v2) & 63)), F1: -117, F2: -499}
                        if false {
                            v2 += (^(v2))
                            if true {
                                v54 := GenS1{F0: ((v6 - (v6 & (v6 * 98))) == v6)}
                                v5[(v36 * v36)] = (!true)
                                _ = v54
                            }
                            switch (v4 + "") {
                            case "hello":
                                var v55 float64 = 0.1
                                _ = v55
                            default:
                                if true {
                                    if false {
                                        if (0.1 == float64(v1)) {
                                            if false {
                                                switch v37 {
                                                case -1:
                                                    v56 := GenS6{F0: true, F1: v7}
                                                    _ = v56
                                                case 42:
                                                    v57 := GenS2{F0: 32767}
                                                    _ = v57
                                                default:
                                                    var v58 int = v37
                                                    _ = v58
                                                }
                                                fmt.Println(v49)
                                                v59 := GenS8{F0: false}
                                                v9 += (v9 >> ((v2) & 63))
                                                _ = v59
                                            }
                                        } else {
                                            v60 := GenS8{F0: false}
                                            _ = v60
                                        }
                                    }
                                    v61 := v5[(v36 & 4294967295)]
                                    var v62 int32 = v9
                                    _ = v61
                                    _ = v62
                                } else {
                                    if (true && ((v3 >> (((^(v2))) & 63)) == (v3 << ((uint(v36)) & 63)))) {
                                        v8 = append(v8, fi46)
                                    } else {
                                        var v63 uint64 = 1
                                        _ = v63
                                    }
                                }
                            }
                        }
                        _ = v52
                        _ = v53
                    }
                    _ = v48
                    _ = v49
                }
                var v64 int = (fi42 * (^(fi42)))
                _ = v64
            } else {
                if ((1.5 * v1) != float32(0.0)) {
                    switch ("hello" + v4) {
                    case "":
                        fmt.Println(v6, v3)
                    }
                    var v65 int64 = -746
                    _ = v65
                } else {
                    v66 := v5[(v36 >> (((v2 + (v2 % (((v2 * (^(v2)))) | 1)))) & 63))]
                    _ = v66
                }
                var v67 int = (fi46 / (((fi46 >> ((4294967295) & 63))) | 1))
                _ = v67
            }
            _ = fi46
            _ = v47
        }
        _ = fi42
    }
    var v68 bool = true
    if (!true) {
        v69 := v5[uint32(v7)]
        for fi70 := 0; fi70 < 17; fi70++ {
            var v71 int8 = (v6 - (v6 & v6))
            for fi72 := 0; fi72 < 5; fi72++ {
                var v73 uint16 = 32767
                if (uint16(v3) <= (v73 / ((42) | 1))) {
                    var v74 bool = true
                    _ = v74
                }
                v3 -= (v3 >> (((v2 << ((v2) & 63))) & 63))
                _ = fi72
                _ = v73
            }
            if (411 != 42) {
                _ = len(v8)
                v8 = append(v8, fi70)
            }
            _ = fi70
            _ = v71
        }
        _ = v69
    } else {
        if (v36 == 2) {
            if false {
                if v68 {
                    fmt.Println(v6, v1, v7)
                }
                switch (v37 / (((v37 + v37)) | 1)) {
                case 1:
                    var v75 int32 = v9
                    _ = v75
                case 8:
                    delete(v5, 2)
                default:
                    var v76 int8 = -81
                    _ = v76
                }
                v77 := GenS2{F0: 716}
                _ = v77
            } else {
                for fi78 := 0; fi78 < 7; fi78++ {
                    if v68 {
                        if len(v8) > 0 {
                            v8[0] = (v37 - fi78)
                        }
                        for fi79 := 0; fi79 < 11; fi79++ {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            var v80 int = fi78
                            _ = fi79
                            _ = v80
                        }
                    } else {
                        v9 -= v9
                        for fi81 := 0; fi81 < 10; fi81++ {
                            v82 := GenS9{F0: v2, F1: (v36 & v36)}
                            fmt.Println(v36, v4, v2)
                            _ = fi81
                            _ = v82
                        }
                        v83 := v5[v36]
                        _ = v83
                    }
                    v3 = 216
                    v4 = "a"
                    _ = fi78
                }
            }
            for fi84 := 0; fi84 < 15; fi84++ {
                for fi85 := 0; fi85 < 14; fi85++ {
                    var v86 uint16 = v7
                    var v87 int32 = v9
                    if v68 {
                        switch v4 {
                        case "":
                            v88 := GenS10{F0: (v7 - (v7 + v7)), F1: v6}
                            _ = v88
                        case "hello":
                            var v89 int64 = 581
                            _ = v89
                        case "foo bar":
                            var v90 uint32 = v36
                            _ = v90
                        }
                        v6 -= (v6 ^ v6)
                        var v91 uint = 42
                        _ = v91
                    } else {
                        v92 := GenS4{F0: v7}
                        _ = v92
                    }
                    _ = fi85
                    _ = v86
                    _ = v87
                }
                _ = fi84
            }
            v3 -= uint8(v7)
        }
    }
    v37 = v37
    v7 = v7
    v93 := GenS11{F0: (^(v37)), F1: v6}
    if (v1 > v1) {
        for fi94 := 0; fi94 < 20; fi94++ {
            v37 += v37
            var v95 float32 = (v1 / float32(1e10))
            _ = fi94
            _ = v95
        }
        var v96 int8 = (v6 + v6)
        _ = v96
    }
    var v97 float32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v7
    _ = v8
    _ = v9
    _ = v93
    _ = v97
}
