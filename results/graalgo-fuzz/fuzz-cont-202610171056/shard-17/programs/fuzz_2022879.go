package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int8 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16; F1 bool; F2 int16 }
type GenS9 struct { F0 int64; F1 int64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 bool; F1 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = ((((0.1 / (1e10 - 1.5)) + float64(0.1)) - 0.0) / float64(1e10))
    var v3 string = ("a" + "foo bar")
    var v4 int = 1
    switch "hello" {
    case "a":
        v2 = 1.5
    default:
        v5 := []int{v4, (v4 | v4)}
        _ = v5
    }
    fmt.Println(v2, v2, v2)
    if true {
        var v6 uint = 4294967295
        v7 := []int{42, (v4 / ((v4) | 1)), (v4 - -2147483648), 32767}
        fmt.Println(v6, v2, v6)
        _ = v6
        _ = v7
    }
    if false {
        fmt.Println(v4, v4, v4)
        if (true && true) {
            v8 := []int{(v4 & v4), (v4 * v4)}
            _ = v8
        } else {
            v9 := GenS1{F0: 1, F1: v3, F2: int8(v4)}
            delete(v1, 2)
            v10 := GenS2{F0: (v3 + ((((v3 + (((v3 + v3) + v3) + "a")) + (v3 + "a")) + "foo bar") + "foo bar")), F1: 231}
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := []int{(fi11 ^ (v4 % ((128) | 1)))}
            for fi13 := 0; fi13 < 3; fi13++ {
                if (2 > 42) {
                    var v14 int8 = -79
                    _ = v14
                } else {
                    var v15 int8 = 69
                    if (2 < 4294967295) {
                        var v16 uint16 = 0
                        v2 = -2.25
                        v2 -= (v2 - v2)
                        _ = v16
                    } else {
                        var v17 float32 = float32((v2 * (0.1 / ((v2 + (1e10 - v2)) - (float64(1.5) - float64(-2.25))))))
                        if (true || (true || true)) {
                            switch "\n\t" {
                            case "":
                                var v18 int64 = 707
                                _ = v18
                            case "a":
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            }
                        } else {
                            v3 = "a"
                        }
                        v17 = v17
                        var v19 uint8 = 127
                        if true {
                            v12 = append(v12, -1)
                            v20 := GenS1{F0: 456, F1: "", F2: (v15 % ((v15) | 1))}
                            fmt.Println(fi13, v17)
                            if (!true) {
                                switch fi11 {
                                case -304:
                                    v21 := v1[42]
                                    _ = v21
                                case -2147483648:
                                    fmt.Println(fi11, v15, v17)
                                    v22 := GenS2{F0: ((((("hello" + v3) + v3) + (v3 + ("foo bar" + (v3 + ("" + ""))))) + v3) + "\n\t"), F1: 1}
                                    v23 := GenS1{F0: 42, F1: v3, F2: (v15 >> ((42) & 63))}
                                    _ = v22
                                    _ = v23
                                case 581:
                                    if false {
                                        var v24 int32 = -343
                                        var v25 int16 = 1000
                                        switch (fi11 & -9223372036854775808) {
                                        case -449:
                                            if false {
                                                var v26 int8 = v15
                                                _ = v26
                                            } else {
                                                fmt.Println(v19, fi13)
                                                v3 = "a"
                                            }
                                        case -1:
                                            v27 := GenS1{F0: 980, F1: v3, F2: (v15 >> ((2) & 63))}
                                            _ = v27
                                        case -806:
                                            switch (v4 - (^(v4))) {
                                            case 127:
                                                v28 := GenS1{F0: 515, F1: (v3 + "foo bar"), F2: 127}
                                                _ = v28
                                            case -567:
                                                v29 := GenS3{F0: 2}
                                                _ = v29
                                            default:
                                                var v30 uint8 = (^(v19))
                                                _ = v30
                                            }
                                        default:
                                            v4 = (-(fi11))
                                        }
                                        switch (fi11 + fi13) {
                                        case 0:
                                            v31 := GenS2{F0: v3, F1: 32767}
                                            _ = v31
                                        case 2147483647:
                                            var v32 bool = ((!false) && false)
                                            _ = v32
                                        case -192:
                                            var v33 uint16 = 968
                                            _ = v33
                                        }
                                        _ = v24
                                        _ = v25
                                    } else {
                                        if (!false) {
                                            v19 -= 127
                                            if ((false && (!false)) || false) {
                                                var v34 string = "hello"
                                                _ = v34
                                            }
                                        }
                                        var v35 float32 = v17
                                        _ = v35
                                    }
                                default:
                                    v4 = (v4 + v4)
                                }
                                if (2147483647 > 860) {
                                    switch fi11 {
                                    case 32768:
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        var v36 int = v4
                                        switch v3 {
                                        case "hello":
                                            v37 := GenS1{F0: 957, F1: (v3 + "foo bar"), F2: 100}
                                            _ = v37
                                        }
                                        if true {
                                            if (false && ((float64(0.1) / -2.25) == (v2 / v2))) {
                                                v38 := GenS4{F0: false, F1: false}
                                                v17 -= v17
                                                _ = v38
                                            } else {
                                                switch "hello" {
                                                case "":
                                                    if (true || true) {
                                                        var v39 float64 = -2.25
                                                        if (true || false) {
                                                            var v40 int8 = (v15 / ((-128) | 1))
                                                            var v41 int32 = 74
                                                            _ = v40
                                                            _ = v41
                                                        } else {
                                                            _ = len(v12)
                                                            v42 := GenS4{F0: false, F1: false}
                                                            fmt.Println(v4, v17, fi11)
                                                            _ = v42
                                                        }
                                                        v3 = (v3 + ("a" + v3))
                                                        v17 = v17
                                                        _ = v39
                                                    }
                                                default:
                                                    delete(v1, 42)
                                                }
                                            }
                                            if (!false) {
                                                if (false && ((-1 < 1) && (593 == -1))) {
                                                    v15 += 6
                                                } else {
                                                    var v43 float64 = (v2 * v2)
                                                    v3 = "foo bar"
                                                    v15 -= v15
                                                    _ = v43
                                                }
                                                v44 := GenS1{F0: 65535, F1: v3, F2: (v15 / ((v15) | 1))}
                                                var v45 int64 = -1
                                                v36 = fi13
                                                _ = v44
                                                _ = v45
                                            }
                                            _ = len(v1)
                                            v46 := GenS5{F0: 2147483647, F1: 1}
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v47 float64 = v2
                                                _ = v47
                                            }
                                            _ = v46
                                        }
                                        _ = v36
                                    }
                                } else {
                                    v2 = float64(float32(((float64(v17) - 0.1) - (3.14159 - v2))))
                                    if (2 > 32767) {
                                        v17 += float32((v2 / float64((float32(0.0) + v17))))
                                        if (305 >= 866) {
                                            var v48 int16 = 0
                                            _ = v48
                                        }
                                    } else {
                                        switch "" {
                                        case "a":
                                            _ = len(v12)
                                        case "foo bar":
                                            if (((v17 - (0.1 - (v17 - float32((0.0 + v2))))) + v17) >= 3.14159) {
                                                delete(v1, 2147483648)
                                                var v49 int16 = 655
                                                var v50 string = v3
                                                var v51 int64 = 0
                                                if false {
                                                    switch (-(v4)) {
                                                    case -1:
                                                        var v52 int8 = int8(fi13)
                                                        _ = v52
                                                    }
                                                    v53 := GenS4{F0: false, F1: (v19 <= 1)}
                                                    var v54 bool = (!(false || (true || ((^(v19)) == v19))))
                                                    _ = v53
                                                    _ = v54
                                                }
                                                _ = v49
                                                _ = v50
                                                _ = v51
                                            } else {
                                                if true {
                                                    fmt.Println(v3, v2, v15)
                                                    if false {
                                                        switch v3 {
                                                        case "":
                                                            var v55 string = (("" + v3) + "a")
                                                            _ = v55
                                                        case "\n\t":
                                                            if len(v12) >= 2 {
                                                                v12 = v12[1:len(v12)]
                                                            }
                                                        }
                                                        v2 = (1e10 - v2)
                                                        v56 := GenS1{F0: 42, F1: v3, F2: (^(v15))}
                                                        var v57 uint8 = (v19 * (^(v19)))
                                                        var v58 int16 = 42
                                                        v59 := GenS4{F0: false, F1: (42 < -9223372036854775808)}
                                                        v17 += 1e10
                                                        var v60 uint8 = v19
                                                        var v61 uint32 = 2147483648
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                        _ = v59
                                                        _ = v60
                                                        _ = v61
                                                    } else {
                                                        var v62 uint16 = 1
                                                        switch int(v15) {
                                                        case -164:
                                                            v63 := GenS6{F0: (v15 >> ((682) & 63))}
                                                            _ = v63
                                                        }
                                                        _ = v62
                                                    }
                                                }
                                            }
                                        case "\n\t":
                                            var v64 uint = 0
                                            _ = v64
                                        default:
                                            var v65 int16 = 32767
                                            v19 -= (v19 ^ v19)
                                            var v66 uint = 255
                                            _ = v65
                                            _ = v66
                                        }
                                        v67 := v1[1]
                                        v4 += fi11
                                        _ = v67
                                    }
                                    v2 += float64((-2.25 + -2.25))
                                }
                                var v68 uint = 255
                                v69 := GenS7{F0: 473}
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                _ = v68
                                _ = v69
                            }
                            fmt.Println(v19, fi11, v4)
                            _ = v20
                        } else {
                            switch v3 {
                            case "\n\t":
                                if (255 >= v19) {
                                    if true {
                                        v15 = v15
                                    } else {
                                        v70 := GenS8{F0: 42, F1: true, F2: 42}
                                        _ = v70
                                    }
                                    var v71 int64 = 963
                                    fmt.Println(v17, fi13)
                                    v19 = (v19 * v19)
                                    _ = v71
                                } else {
                                    fmt.Println(fi11, fi11, v15)
                                    _ = len(v1)
                                    fmt.Println(v4, v17)
                                }
                                if (18446744073709551615 == 100) {
                                    _ = len(v1)
                                    if false {
                                        v3 = v3
                                        var v72 uint = 0
                                        v2 = float64(((((((0.0 * v17) - 3.14159) + (float32(-2.25) / -2.25)) + 3.14159) * 1.5) + (v17 - 1.5)))
                                        var v73 int16 = -1
                                        _ = v72
                                        _ = v73
                                    }
                                } else {
                                    if false {
                                        fmt.Println(fi11, fi11, v19)
                                    }
                                }
                            case "foo bar":
                                v17 = ((v17 * v17) + 3.14159)
                            default:
                                switch (fi11 - (v4 << ((1) & 63))) {
                                case -936:
                                    var v74 float64 = (v2 * v2)
                                    _ = v74
                                case 746:
                                    v75 := GenS9{F0: -716, F1: -9223372036854775808}
                                    _ = v75
                                }
                            }
                            v1[4294967295] = 4294967295
                            v3 = ((v3 + "a") + "hello")
                        }
                        _ = v17
                        _ = v19
                    }
                    switch 128 {
                    case 42:
                        switch ((v3 + v3) + ("hello" + "hello")) {
                        case "a":
                            if false {
                                v76 := GenS6{F0: 42}
                                fmt.Println(v15, v15, v2)
                                _ = v76
                            } else {
                                if false {
                                    v3 = v3
                                    v3 = ((v3 + "foo bar") + ((v3 + (v3 + ("\n\t" + "\n\t"))) + "\n\t"))
                                } else {
                                    fmt.Println(v4, fi13)
                                    var v77 float64 = v2
                                    v78 := v1[2147483648]
                                    v3 = v3
                                    switch "foo bar" {
                                    case "":
                                        var v79 int = fi13
                                        v2 -= 1e10
                                        _ = v79
                                    default:
                                        var v80 int32 = -542
                                        _ = v80
                                    }
                                    _ = v77
                                    _ = v78
                                }
                                v12 = append(v12, (fi13 | (fi13 << ((65535) & 63))))
                            }
                        }
                    case -9223372036854775808:
                        _ = len(v12)
                    default:
                        if ((true || ((true || true) || true)) && true) {
                            var v81 uint32 = 2147483647
                            v4 = (fi11 - (fi13 % (((fi11 % ((fi13) | 1))) | 1)))
                            _ = v81
                        } else {
                            v82 := GenS10{F0: 65535}
                            switch (-(v4)) {
                            case 0:
                                v4 = fi11
                            case -390:
                                v83 := v1[42]
                                _ = v83
                            case 128:
                                var v84 int32 = 2147483647
                                _ = v84
                            default:
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            }
                            var v85 int16 = -32768
                            _ = v82
                            _ = v85
                        }
                    }
                    _ = v15
                }
                v86 := GenS11{F0: true, F1: 9223372036854775807}
                var v87 float64 = float64(float32(1.5))
                _ = fi13
                _ = v86
                _ = v87
            }
            _ = fi11
            _ = v12
        }
    }
    if true {
        v88 := GenS6{F0: 96}
        _ = v88
    }
    for fi89 := 0; fi89 < 3; fi89++ {
        var v90 uint = 2
        v91 := []int{}
        switch -416 {
        case -755:
            _ = len(v1)
        case -675:
            v3 = v3
        default:
            if (!true) {
                switch fi89 {
                case -663:
                    var v92 uint64 = 2
                    _ = v92
                }
                v93 := GenS1{F0: v90, F1: ("\n\t" + "hello"), F2: 56}
                _ = v93
            }
        }
        _ = fi89
        _ = v90
        _ = v91
    }
    var v94 float64 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v94
}
