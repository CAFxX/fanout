package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool; F1 int; F2 uint32 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 string }

func main() {
    var v1 float64 = (1e10 + float64(3.14159))
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    for fi4 := 0; fi4 < 12; fi4++ {
        switch fi4 {
        case 4294967295:
            v1 += (1.5 - v1)
        }
        _ = fi4
    }
    v5 := GenS1{F0: 42, F1: -889}
    fmt.Println(v1)
    v6 := GenS1{F0: 847, F1: 32767}
    v7 := GenS1{F0: 32767, F1: 9223372036854775807}
    v1 = v1
    v1 -= (v1 * 0.0)
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := v2[2147483647]
        v10 := []int{(fi8 * (fi8 / ((fi8) | 1)))}
        for fi11 := 0; fi11 < 14; fi11++ {
            v9 = 42
            var v12 string = "hello"
            _ = len(v3)
            if (!((fi8 >= (fi8 << ((620) & 63))) && (32767 > 65535))) {
                switch ((("" + ((v12 + (v12 + v12)) + v12)) + (v12 + (v12 + v12))) + (((v12 + "") + "hello") + v12)) {
                case "hello":
                    v13 := GenS2{F0: 1}
                    _ = v13
                default:
                    if false {
                        if false {
                            switch fi11 {
                            case 4294967295:
                                if true {
                                    v14 := v2[(v9 % (((v9 + 881)) | 1))]
                                    v14 = (v14 | v9)
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    _ = v14
                                } else {
                                    var v15 int = (fi8 | 87)
                                    var v16 int32 = 42
                                    if true {
                                        fmt.Println(v1)
                                    } else {
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                        var v17 uint64 = 4294967295
                                        _ = v17
                                    }
                                    _ = v15
                                    _ = v16
                                }
                            case 2147483647:
                                if false {
                                    switch ("hello" + v12) {
                                    case "hello":
                                        v18 := GenS1{F0: 903, F1: fi8}
                                        if (-32768 > 0) {
                                            v19 := GenS1{F0: 65535, F1: fi8}
                                            var v20 int8 = 127
                                            v9 += v9
                                            fmt.Println(v20, fi11)
                                            if true {
                                                v21 := GenS1{F0: 2, F1: (-(fi8))}
                                                switch (fi8 / (((^(fi11))) | 1)) {
                                                case 4294967295:
                                                    var v22 bool = ((-2.25 + float64(3.14159)) <= 0.1)
                                                    _ = v22
                                                case -1:
                                                    switch (v12 + ((v12 + (v12 + "hello")) + v12)) {
                                                    case "hello":
                                                        var v23 int8 = -128
                                                        _ = v23
                                                    case "foo bar":
                                                        switch v12 {
                                                        case "\n\t":
                                                            var v24 uint = 1
                                                            _ = v24
                                                        case "":
                                                            v20 = v20
                                                        case "a":
                                                            v20 += v20
                                                        default:
                                                            var v25 uint8 = 1
                                                            _ = v25
                                                        }
                                                    case "\n\t":
                                                        if (v20 <= -26) {
                                                            var v26 bool = true
                                                            v27 := GenS2{F0: (fi11 << ((42) & 63))}
                                                            _ = v26
                                                            _ = v27
                                                        }
                                                    default:
                                                        if len(v10) > 0 {
                                                            v10[0] = (-(fi8))
                                                        }
                                                    }
                                                }
                                                _ = v21
                                            } else {
                                                var v28 string = "a"
                                                fmt.Println(fi8, v1, v28)
                                                _ = v28
                                            }
                                            _ = v19
                                            _ = v20
                                        }
                                        var v29 int = fi8
                                        v12 = ""
                                        var v30 int8 = 50
                                        v12 = "foo bar"
                                        v3[(v9 ^ v9)] = fi11
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                        v31 := GenS1{F0: 1, F1: 523}
                                        v9 -= v9
                                        _ = v18
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    case "":
                                        if true {
                                            if true {
                                                switch v12 {
                                                case "hello":
                                                    if (255 > 63) {
                                                        switch "a" {
                                                        case "foo bar":
                                                            v32 := GenS1{F0: 898, F1: (fi11 << ((18446744073709551615) & 63))}
                                                            _ = v32
                                                        case "hello":
                                                            v9 -= (v9 * 625)
                                                        default:
                                                            var v33 int = (fi11 - 701)
                                                            _ = v33
                                                        }
                                                    }
                                                case "foo bar":
                                                    fmt.Println(fi8, v9)
                                                }
                                            }
                                            if (!(42 < 0)) {
                                                v9 -= (^(v9))
                                                fmt.Println(v12, v9)
                                                var v34 uint8 = 2
                                                _ = v34
                                            } else {
                                                v12 = (v12 + "hello")
                                                _ = len(v3)
                                                v35 := v3[(v9 ^ v9)]
                                                _ = v35
                                            }
                                            fmt.Println(v1, v9, v1)
                                            var v36 bool = false
                                            var v37 uint16 = 255
                                            _ = v36
                                            _ = v37
                                        }
                                    case "a":
                                        if false {
                                            v2[2] = 2147483648
                                        } else {
                                            if (-2.25 > v1) {
                                                v3[v9] = fi11
                                            }
                                            if true {
                                                var v38 string = ((v12 + v12) + v12)
                                                switch "hello" {
                                                case "\n\t":
                                                    v9 = 58
                                                default:
                                                    switch fi11 {
                                                    case 32767:
                                                        if true {
                                                            fmt.Println(v9, v12)
                                                            v1 += (1e10 / (v1 - v1))
                                                            v39 := GenS2{F0: fi8}
                                                            _ = v39
                                                        } else {
                                                            switch (fi11 ^ (fi8 - (fi11 / ((1) | 1)))) {
                                                            case -1:
                                                                v1 = float64(float32((3.14159 * v1)))
                                                            case 2147483647:
                                                                fmt.Println(v9)
                                                                if (!(v9 <= v9)) {
                                                                    var v40 float64 = (v1 - ((v1 / v1) - v1))
                                                                    _ = v40
                                                                } else {
                                                                    var v41 int = (fi11 >> ((1) & 63))
                                                                    _ = v41
                                                                }
                                                                fmt.Println(v38, fi11, fi8)
                                                            case 42:
                                                                v9 -= 2147483647
                                                            default:
                                                                v12 = ("a" + (v12 + (((("foo bar" + (v12 + v12)) + ("" + v12)) + "a") + ((v38 + v38) + "foo bar"))))
                                                            }
                                                            v38 = "foo bar"
                                                        }
                                                        _ = len(v3)
                                                        var v42 int8 = 42
                                                        _ = v42
                                                    case 1:
                                                        v9 += 65535
                                                    case 9223372036854775807:
                                                        var v43 int8 = 42
                                                        _ = v43
                                                    default:
                                                        v9 = (v9 - 88)
                                                    }
                                                }
                                                v44 := GenS3{F0: 255}
                                                _ = v38
                                                _ = v44
                                            } else {
                                                switch v12 {
                                                case "hello":
                                                    if false {
                                                        v45 := GenS2{F0: (fi11 - fi8)}
                                                        _ = v45
                                                    }
                                                default:
                                                    fmt.Println(fi11, v1, fi11)
                                                }
                                                v12 = ("\n\t" + v12)
                                            }
                                            v9 += 654
                                        }
                                    }
                                    v9 = v9
                                }
                            case -212:
                                var v46 uint8 = 127
                                _ = v46
                            }
                        } else {
                            var v47 uint32 = 42
                            _ = v47
                        }
                        var v48 uint64 = 4294967295
                        v3[(v9 ^ v9)] = (fi8 >> ((0) & 63))
                        _ = v48
                    }
                }
                if (((v12 + v12) + v12) > v12) {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    switch fi8 {
                    case 884:
                        var v49 uint16 = 65535
                        _ = v49
                    case 9223372036854775807:
                        if (!true) {
                            v9 -= (^(v9))
                        } else {
                            var v50 int16 = 1
                            if true {
                                fmt.Println(v12)
                                if (3.14159 == (v1 * v1)) {
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                }
                                if len(v10) > 0 {
                                    v10[0] = (fi11 + (fi11 >> ((255) & 63)))
                                }
                            }
                            var v51 int16 = v50
                            v52 := GenS3{F0: 42}
                            v50 += v50
                            v50 = (v50 + v50)
                            v53 := GenS2{F0: 721}
                            v54 := GenS2{F0: (fi8 / ((fi11) | 1))}
                            var v55 uint64 = 42
                            fmt.Println(v51, v50, fi8)
                            var v56 int64 = -1
                            if (true || false) {
                                v55 = (v55 - uint64(v9))
                                if true {
                                    var v57 int64 = (v56 << ((822) & 63))
                                    v58 := GenS2{F0: 1}
                                    v59 := GenS2{F0: fi11}
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                } else {
                                    var v60 uint16 = 1
                                    var v61 uint8 = 255
                                    var v62 bool = (255 == 255)
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                }
                                if false {
                                    v63 := GenS4{F0: v55, F1: ""}
                                    var v64 uint = uint(v55)
                                    fmt.Println(v12, v56, v1)
                                    _ = v63
                                    _ = v64
                                }
                            } else {
                                if (v56 > (-(v56))) {
                                    fmt.Println(v55, v56)
                                    v65 := GenS5{F0: v56}
                                    _ = v65
                                } else {
                                    var v66 int8 = -1
                                    fmt.Println(fi8)
                                    var v67 int = -359
                                    var v68 int8 = 127
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                }
                                var v69 int16 = v50
                                if false {
                                    if false {
                                        v55 -= (v55 | 2)
                                        switch (fi11 >> ((386) & 63)) {
                                        case 108:
                                            v70 := v2[2]
                                            _ = v70
                                        case 1:
                                            v3[(v9 % (((v9 >> ((1) & 63))) | 1))] = fi8
                                        }
                                    } else {
                                        var v71 bool = (false || true)
                                        fmt.Println(v71, v51)
                                        _ = v71
                                    }
                                    var v72 int = 4294967295
                                    if (((v12 + v12) + "foo bar") <= (v12 + ("hello" + "\n\t"))) {
                                        v73 := GenS6{F0: 100}
                                        var v74 uint = 999
                                        switch fi11 {
                                        case -1:
                                            var v75 uint8 = 31
                                            _ = v75
                                        case 4294967295:
                                            switch (v72 - (fi11 | fi8)) {
                                            case -9223372036854775808:
                                                if true {
                                                    if false {
                                                        fmt.Println(v55, v74, v9)
                                                        v10 = append(v10, fi8)
                                                    } else {
                                                        if false {
                                                            var v76 float32 = 1.5
                                                            v77 := GenS1{F0: 1, F1: fi8}
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                    }
                                                    v78 := GenS3{F0: uint(v55)}
                                                    _ = v78
                                                } else {
                                                    v79 := GenS6{F0: -67}
                                                    v80 := GenS7{F0: (v74 << ((4294967295) & 63))}
                                                    _ = v79
                                                    _ = v80
                                                }
                                            default:
                                                v81 := GenS8{F0: false, F1: fi11, F2: (v9 / (((v9 << (((v74 >> ((v74) & 63))) & 63))) | 1))}
                                                _ = v81
                                            }
                                        case 252:
                                            var v82 string = ((v12 + ("foo bar" + "")) + "")
                                            _ = v82
                                        default:
                                            v1 = v1
                                        }
                                        _ = v73
                                        _ = v74
                                    }
                                    _ = v72
                                } else {
                                    if len(v10) > 0 {
                                        v10[0] = (fi11 >> ((18446744073709551615) & 63))
                                    }
                                    v83 := GenS9{F0: (v50 ^ v51)}
                                    if false {
                                        var v84 bool = true
                                        if (-2147483648 >= 220) {
                                            if len(v10) > 0 {
                                                v10[0] = -1
                                            }
                                        }
                                        delete(v3, v9)
                                        v85 := v2[v9]
                                        _ = v84
                                        _ = v85
                                    } else {
                                        var v86 int32 = 1
                                        v3[v9] = fi8
                                        _ = v86
                                    }
                                    _ = v83
                                }
                                _ = v69
                            }
                            var v87 uint = 0
                            _ = v50
                            _ = v51
                            _ = v52
                            _ = v53
                            _ = v54
                            _ = v55
                            _ = v56
                            _ = v87
                        }
                    default:
                        if false {
                            _ = len(v10)
                        } else {
                            var v88 uint32 = 1
                            _ = v88
                        }
                    }
                } else {
                    fmt.Println(fi11, v1)
                }
            } else {
                if (32767 > 0) {
                    var v89 int16 = -542
                    _ = v89
                }
            }
            fmt.Println(v12, v1, v1)
            _ = fi11
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi90 := 0; fi90 < 13; fi90++ {
        var v91 uint64 = 973
        _ = fi90
        _ = v91
    }
    v1 += ((v1 * (1e10 + -2.25)) - v1)
    v92 := GenS10{F0: "foo bar"}
    var v93 uint16 = 811
    v94 := GenS1{F0: 0, F1: -862}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v92
    _ = v93
    _ = v94
}
