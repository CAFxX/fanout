package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 string; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int64; F1 int8; F2 int16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 bool; F1 int64 }
type GenS10 struct { F0 int8; F1 int16; F2 int16 }
type GenS11 struct { F0 int32; F1 uint64 }
func genFn0() string {
    var v7 int64 = 469
    v7 += (v7 ^ (-(v7)))
    v8 := GenS1{F0: 127, F1: 123}
    var v9 string = ("\n\t" + (("\n\t" + "foo bar") + "\n\t"))
    var v10 string = (v9 + "\n\t")
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return ("" + v9)
}
func genFn1() bool {
    v11 := GenS1{F0: 2, F1: 82}
    if ((-1 == 2147483647) && true) {
        var v12 string = ""
        v13 := GenS1{F0: 2, F1: 127}
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := []int{(fi14 | 607), 983, fi14}
            var v16 uint64 = 4294967295
            if ((!((42 != fi14) || (-2.25 > (1e10 - ((1.5 + 1.5) + 0.0))))) && false) {
                if true {
                    v16 = 18446744073709551615
                } else {
                    fmt.Println(fi14, v16)
                    if (!false) {
                        v17 := GenS1{F0: 0, F1: 133}
                        switch (("hello" + "hello") + "foo bar") {
                        case "\n\t":
                            var v18 int = (fi14 ^ (fi14 + fi14))
                            _ = v18
                        case "hello":
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                        case "foo bar":
                            for fi19 := 0; fi19 < 16; fi19++ {
                                fmt.Println(fi14, fi14, v16)
                                _ = fi19
                            }
                        }
                        v16 -= 1
                        _ = v17
                    } else {
                        if (2147483648 <= 333) {
                            switch fi14 {
                            case 32768:
                                v20 := GenS1{F0: 722, F1: 2}
                                _ = v20
                            case 9223372036854775807:
                                var v21 uint16 = 42
                                _ = v21
                            case 4294967295:
                                v16 = v16
                            default:
                                v22 := GenS2{F0: (((("a" + "hello") + "\n\t") + "hello") + ("\n\t" + "a")), F1: 2147483647, F2: 0}
                                _ = v22
                            }
                            var v23 bool = true
                            switch "a" {
                            case "foo bar":
                                v16 += 782
                            }
                            _ = v23
                        } else {
                            v24 := GenS3{F0: 4294967295, F1: 2, F2: 1}
                            v25 := GenS2{F0: "a", F1: 582, F2: 5}
                            _ = v24
                            _ = v25
                        }
                    }
                    fmt.Println(v16, fi14)
                }
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        v26 := GenS2{F0: "a", F1: -2, F2: 45}
        var v27 string = "a"
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        if true {
            v29 := GenS4{F0: 1, F1: 127}
            fmt.Println(fi28)
            for fi30 := 0; fi30 < 2; fi30++ {
                if ((-97 == 0) && false) {
                    v31 := GenS1{F0: 2, F1: 127}
                    if false {
                        switch "" {
                        case "\n\t":
                            var v32 uint32 = 42
                            _ = v32
                        case "foo bar":
                            var v33 uint8 = 217
                            _ = v33
                        case "hello":
                            var v34 uint = 1
                            var v35 uint16 = 65535
                            if (0 == 127) {
                                v36 := GenS5{F0: 1000}
                                v37 := GenS1{F0: 2, F1: 42}
                                v35 = 812
                                _ = v36
                                _ = v37
                            }
                            if (!(42 != 100)) {
                                v38 := []int{2147483647, (-(fi28))}
                                fmt.Println(fi30)
                                var v39 uint16 = (v35 | v35)
                                _ = v38
                                _ = v39
                            }
                            var v40 uint16 = v35
                            _ = v34
                            _ = v35
                            _ = v40
                        default:
                            v41 := []int{fi30, 37, (fi30 << ((576) & 63))}
                            _ = v41
                        }
                        v42 := GenS6{F0: 2}
                        var v43 int64 = -65
                        var v44 int64 = (-(v43))
                        var v45 int = (fi28 * (fi30 >> ((810) & 63)))
                        v46 := GenS2{F0: "hello", F1: -1, F2: 127}
                        var v47 int64 = (v44 ^ v44)
                        v44 = (v44 | 9223372036854775807)
                        if false {
                            v48 := []int{(v45 ^ (fi28 >> ((255) & 63))), v45}
                            _ = v48
                        } else {
                            v45 += 32768
                            if true {
                                v49 := []int{32768}
                                _ = v49
                            } else {
                                var v50 int16 = -1
                                v43 = (v43 * (^(v47)))
                                _ = v50
                            }
                            v43 += (v44 << ((235) & 63))
                        }
                        _ = v42
                        _ = v43
                        _ = v44
                        _ = v45
                        _ = v46
                        _ = v47
                    } else {
                        v51 := []int{-2147483648, -368, (fi30 % (((fi30 | fi30)) | 1))}
                        var v52 int = 0
                        _ = v51
                        _ = v52
                    }
                    _ = v31
                } else {
                    v53 := GenS7{F0: 1, F1: 100, F2: int16(fi30)}
                    switch ("a" + "a") {
                    case "a":
                        var v54 uint16 = 65535
                        _ = v54
                    default:
                        v55 := []int{(-(fi28))}
                        _ = v55
                    }
                    if (42 >= 42) {
                        v56 := GenS8{F0: false}
                        switch (("" + "a") + "\n\t") {
                        case "foo bar":
                            var v57 int = fi30
                            _ = v57
                        default:
                            var v58 int16 = -32768
                            _ = v58
                        }
                        var v59 float32 = 1e10
                        _ = v56
                        _ = v59
                    }
                    _ = v53
                }
                _ = fi30
            }
            _ = v29
        }
        _ = fi28
    }
    _ = v11
    return (((false || false) && false) || (true || false))
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 636
    var v3 uint8 = 163
    var v4 float32 = -2.25
    var v5 uint = 1
    v6 := make(map[int]int)
    v60 := []int{int(v2), 42, -1}
    for fi61 := 0; fi61 < 18; fi61++ {
        for fi62 := 0; fi62 < 15; fi62++ {
            v63 := GenS9{F0: ((v2 / (((v2 % (((v2 + 661)) | 1))) | 1)) == 42), F1: -1}
            if (false || (v4 >= float32((0.0 + 3.14159)))) {
                v4 += (v4 + v4)
                fmt.Println(v3, v2)
                switch fi61 {
                case 570:
                    v64 := GenS10{F0: -1, F1: v2, F2: v2}
                    _ = v64
                case 2147483647:
                    if (65535 != 2147483647) {
                        fmt.Println(v4, v2)
                        if (!true) {
                            v65 := GenS11{F0: 13, F1: 2}
                            _ = v65
                        } else {
                            if (true && true) {
                                switch (fi62 << (((v5 - v5)) & 63)) {
                                case 1:
                                    v66 := v1[int(v5)]
                                    _ = v66
                                default:
                                    var v67 uint64 = 2
                                    _ = v67
                                }
                                var v68 float64 = 1.5
                                v69 := GenS4{F0: -2147483648, F1: (v3 - uint8(v5))}
                                _ = v68
                                _ = v69
                            }
                        }
                    } else {
                        fmt.Println(v3, fi62, v2)
                        if true {
                            if false {
                                v3 = (v3 - uint8(v5))
                                switch -873 {
                                case 2147483647:
                                    var v70 string = "a"
                                    _ = v70
                                default:
                                    v60 = append(v60, 2147483647)
                                }
                                if true {
                                    switch (("hello" + ("" + "foo bar")) + "") {
                                    case "\n\t":
                                        v71 := GenS3{F0: 2147483648, F1: v3, F2: 2}
                                        _ = v71
                                    case "foo bar":
                                        v72 := GenS10{F0: -128, F1: 249, F2: int16(fi61)}
                                        _ = v72
                                    case "a":
                                        v73 := GenS1{F0: 539, F1: (v3 | 255)}
                                        _ = v73
                                    default:
                                        v74 := v6[(^(fi61))]
                                        _ = v74
                                    }
                                    var v75 int32 = 0
                                    if (((2 >= 2147483648) || ((fi62 % ((2147483647) | 1)) != fi62)) && true) {
                                        _ = len(v60)
                                        if (true || (true && (!false))) {
                                            v75 = (v75 | (v75 ^ (^(v75))))
                                            var v76 uint8 = uint8(v5)
                                            v76 += (v3 >> (((v5 % ((v5) | 1))) & 63))
                                            if (true && true) {
                                                v5 = v5
                                            } else {
                                                var v77 uint = 1
                                                var v78 uint32 = 4294967295
                                                var v79 float64 = 0.1
                                                _ = v77
                                                _ = v78
                                                _ = v79
                                            }
                                            if len(v60) > 0 {
                                                v60[0] = (fi61 | (-(fi61)))
                                            }
                                            _ = v76
                                        }
                                        _ = len(v60)
                                    }
                                    _ = v75
                                }
                            }
                        }
                        v4 += (v4 + v4)
                    }
                case 1:
                    var v80 string = ("" + "hello")
                    _ = v80
                default:
                    var v81 uint32 = 2147483648
                    _ = v81
                }
                v82 := GenS1{F0: 2, F1: 255}
                var v83 int32 = 127
                v84 := GenS5{F0: (v2 & -32768)}
                _ = v82
                _ = v83
                _ = v84
            }
            _ = fi62
            _ = v63
        }
        var v85 bool = true
        if (!((v5 + v5) == uint(v3))) {
            _ = len(v60)
            v86 := GenS8{F0: (true || v85)}
            _ = v86
        }
        _ = fi61
        _ = v85
    }
    var v87 uint16 = 2
    for fi88 := 0; fi88 < 11; fi88++ {
        var v89 int = 0
        if true {
            fmt.Println(fi88, v5, v89)
            v5 = (v5 * (v5 % ((373) | 1)))
            v3 = v3
        }
        var v90 string = "foo bar"
        _ = fi88
        _ = v89
        _ = v90
    }
    v91 := GenS7{F0: 9223372036854775807, F1: -123, F2: -1}
    v4 = 0.1
    v92 := GenS11{F0: -451, F1: 18446744073709551615}
    var v93 string = "\n\t"
    var v94 uint8 = (v3 ^ 42)
    v5 = v5
    for fi95 := 0; fi95 < 20; fi95++ {
        var v96 bool = true
        fmt.Println(v87, v3)
        _ = fi95
        _ = v96
    }
    if ((!true) && ((2 <= 2) || (!false))) {
        _ = len(v1)
        var v97 int32 = 1
        _ = v97
    }
    switch "foo bar" {
    case "":
        v93 = "a"
    default:
        if false {
            if len(v60) >= 2 {
                v60 = v60[1:len(v60)]
            }
            var v98 float32 = float32(1e10)
            _ = v98
        } else {
            var v99 uint = (v5 | 2)
            v100 := GenS1{F0: 1, F1: (v3 % ((160) | 1))}
            _ = v100
            _ = v99
        }
    }
    for fi101 := 0; fi101 < 11; fi101++ {
        var v102 int8 = -94
        _ = fi101
        _ = v102
    }
    fmt.Println(v4, v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v60:", len(v60), v60)
    fmt.Println("v87:", v87)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v87
    _ = v91
    _ = v92
    _ = v93
    _ = v94
}
