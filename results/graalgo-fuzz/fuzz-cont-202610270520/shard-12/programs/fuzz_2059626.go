package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 int8; F2 uint }
type GenS5 struct { F0 uint64; F1 bool; F2 int16 }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int8 }
func genFn0(p0 uint8) uint16 {
    v7 := GenS1{F0: 42}
    var v8 uint64 = 945
    v9 := []int{-750}
    var v10 uint8 = p0
    v11 := GenS2{F0: -36}
    _ = p0
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 779
}
func genFn1(p0 int32, p1 bool, p2 int32) int {
    var v12 bool = false
    switch "" {
    case "foo bar":
        v13 := GenS3{F0: 1}
        _ = v13
    case "":
        var v14 float32 = 3.14159
        _ = v14
    case "hello":
        p2 = p0
    default:
        switch ("foo bar" + ("hello" + "hello")) {
        case "hello":
            var v15 uint64 = 836
            _ = v15
        case "\n\t":
            if (0 >= -1) {
                var v16 float64 = 0.0
                p2 += (p2 ^ (p0 * (p2 % ((1) | 1))))
                var v17 int64 = 0
                _ = v16
                _ = v17
            } else {
                if (!true) {
                    v18 := GenS2{F0: 42}
                    v19 := []int{}
                    _ = v18
                    _ = v19
                }
            }
        default:
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 bool = true
                v22 := GenS2{F0: -9223372036854775808}
                v23 := []int{}
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
        }
    }
    v12 = (!v12)
    v24 := GenS2{F0: -9223372036854775808}
    fmt.Println(v12, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v24
    return -9223372036854775808
}

func main() {
    var v1 string = (("foo bar" + "") + "hello")
    v2 := make(map[uint32]int)
    var v3 int64 = -77
    var v4 float32 = 1e10
    var v5 int64 = v3
    var v6 uint16 = 510
    v25 := GenS4{F0: 2147483647, F1: 100, F2: 1}
    fmt.Println(v4, v3)
    if (false || false) {
        switch -645 {
        case 191:
            v26 := []int{genFn1(234, (!false), 1), 455, 1}
            _ = v26
        case 127:
            v6 -= (^(v6))
        }
        v27 := GenS3{F0: v5}
        if (117 == 5) {
            switch "hello" {
            case "a":
                v1 = v1
            case "hello":
                for fi28 := 0; fi28 < 3; fi28++ {
                    for fi29 := 0; fi29 < 19; fi29++ {
                        v30 := GenS3{F0: 421}
                        v31 := []int{fi28, genFn1(-780, (-869 <= int16(fi28)), 1), -9223372036854775808, (^(fi28))}
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = fi28
                }
            default:
                v32 := GenS2{F0: (v5 & (v3 >> ((18446744073709551615) & 63)))}
                _ = v32
            }
        } else {
            var v33 float32 = (v4 / v4)
            var v34 float32 = float32(((3.14159 / -2.25) - 0.0))
            for fi35 := 0; fi35 < 20; fi35++ {
                fmt.Println(v6)
                fmt.Println(v34)
                switch fi35 {
                case 42:
                    v36 := GenS4{F0: 723, F1: 127, F2: 2}
                    _ = v36
                case 2147483647:
                    v37 := v2[0]
                    _ = v37
                }
                _ = fi35
            }
            _ = v33
            _ = v34
        }
        _ = v27
    } else {
        var v38 float32 = 0.0
        _ = v38
    }
    var v39 int = 32767
    v6 += v6
    if (42 != 1) {
        v40 := GenS5{F0: 18446744073709551615, F1: (false && (1 == 2147483648)), F2: 32767}
        var v41 float32 = ((1.5 * (v4 * v4)) + v4)
        fmt.Println(v1)
        _ = v40
        _ = v41
    } else {
        for fi42 := 0; fi42 < 11; fi42++ {
            switch (v39 << ((65535) & 63)) {
            case 128:
                switch (-(v39)) {
                case 32768:
                    if (!true) {
                        if (42 >= 63) {
                            v43 := []int{(fi42 & 994), -443, genFn1(49, true, -1), (v39 + (fi42 % ((782) | 1)))}
                            fmt.Println(v6, v5, v4)
                            switch (fi42 - v39) {
                            case 1:
                                v39 = 4294967295
                            }
                            _ = v43
                        } else {
                            v1 = v1
                            var v44 int32 = -2147483648
                            v39 += (fi42 | -2147483648)
                            _ = v44
                        }
                    } else {
                        for fi45 := 0; fi45 < 13; fi45++ {
                            v39 = (v39 - fi45)
                            switch v1 {
                            case "foo bar":
                                var v46 int8 = -1
                                _ = v46
                            case "":
                                v1 = ((("hello" + "foo bar") + v1) + (((v1 + ("\n\t" + ((v1 + v1) + (v1 + ((v1 + "foo bar") + v1))))) + v1) + (v1 + "\n\t")))
                            case "a":
                                switch genFn1(42, (false && (false || (2147483648 <= 2147483647))), -1) {
                                case 0:
                                    v2[4294967295] = (-(fi45))
                                    if true {
                                        if ((false || false) || true) {
                                            var v47 string = v1
                                            _ = v47
                                        }
                                        if true {
                                            var v48 uint32 = 42
                                            if false {
                                                var v49 int64 = v5
                                                _ = v49
                                            } else {
                                                fmt.Println(v1, v1, v48)
                                            }
                                            switch (fi42 * v39) {
                                            case 32768:
                                                var v50 string = (v1 + v1)
                                                _ = v50
                                            case -2147483648:
                                                v6 = (v6 * v6)
                                            case 9223372036854775807:
                                                v5 = v3
                                            }
                                            if true {
                                                var v51 int = (fi42 ^ (v39 << ((18446744073709551615) & 63)))
                                                var v52 uint32 = (v48 & (v48 * v48))
                                                _ = len(v2)
                                                _ = v51
                                                _ = v52
                                            } else {
                                                delete(v2, v48)
                                                v53 := v2[uint32(v6)]
                                                _ = v53
                                            }
                                            v54 := []int{(-(v39))}
                                            _ = v48
                                            _ = v54
                                        } else {
                                            var v55 uint32 = 2147483647
                                            switch v39 {
                                            case 686:
                                                var v56 uint16 = v6
                                                _ = v56
                                            case 32767:
                                                var v57 uint16 = v6
                                                _ = v57
                                            default:
                                                if true {
                                                    v6 += (v6 & 562)
                                                    v39 = (fi42 << ((42) & 63))
                                                } else {
                                                    var v58 float32 = (0.0 * 0.0)
                                                    if true {
                                                        v5 = (v3 / ((v5) | 1))
                                                    } else {
                                                        v5 = 1
                                                        v59 := v2[(v55 ^ (v55 << ((889) & 63)))]
                                                        _ = v59
                                                    }
                                                    _ = v58
                                                }
                                            }
                                            _ = v55
                                        }
                                    }
                                    if false {
                                        if ("" >= ("" + v1)) {
                                            if (false || (true || false)) {
                                                v60 := []int{42, fi45}
                                                _ = v60
                                            }
                                            v61 := GenS6{F0: -32768, F1: 624}
                                            _ = v61
                                        }
                                        v62 := GenS7{F0: true}
                                        var v63 float64 = (-2.25 / 1.5)
                                        if (!(!((^(v6)) > v6))) {
                                            v64 := []int{332, int(v5)}
                                            v1 = ("a" + v1)
                                            var v65 int = v39
                                            _ = v64
                                            _ = v65
                                        } else {
                                            v66 := GenS8{F0: 2, F1: 0, F2: 65535}
                                            var v67 float64 = v63
                                            fmt.Println(fi42)
                                            v68 := []int{-9223372036854775808}
                                            if true {
                                                delete(v2, 1)
                                                v63 -= 3.14159
                                                var v69 int16 = -1
                                                _ = v69
                                            } else {
                                                fmt.Println(fi42, v67, v1)
                                                var v70 int16 = 1
                                                _ = v70
                                            }
                                            if true {
                                                if len(v68) >= 2 {
                                                    v68 = v68[1:len(v68)]
                                                }
                                                v5 += (v3 >> ((1) & 63))
                                                if true {
                                                    if (-32768 == -1) {
                                                        var v71 int8 = int8(v6)
                                                        v72 := GenS9{F0: 140}
                                                        v68 = append(v68, 0)
                                                        _ = v71
                                                        _ = v72
                                                    } else {
                                                        var v73 uint64 = 413
                                                        _ = v73
                                                    }
                                                }
                                                fmt.Println(v1)
                                            } else {
                                                v39 = genFn1(42, (!(true || true)), 42)
                                                switch v1 {
                                                case "foo bar":
                                                    v74 := GenS10{F0: 741}
                                                    _ = v74
                                                }
                                            }
                                            v75 := GenS9{F0: 59}
                                            var v76 int64 = v3
                                            v77 := v2[uint32(v6)]
                                            if true {
                                                var v78 int8 = 42
                                                var v79 uint32 = 4294967295
                                                var v80 uint32 = v79
                                                _ = v78
                                                _ = v79
                                                _ = v80
                                            }
                                            v63 = float64(((1.5 - v4) / ((3.14159 * (v4 / v4)) - (v4 + v4))))
                                            v68 = append(v68, fi45)
                                            v2[42] = 42
                                            var v81 float64 = v63
                                            v4 = 0.1
                                            var v82 int16 = -32768
                                            switch (v1 + "a") {
                                            case "":
                                                v4 += 3.14159
                                            case "foo bar":
                                                switch (v1 + ("hello" + "\n\t")) {
                                                case "a":
                                                    switch "foo bar" {
                                                    case "a":
                                                        v4 = v4
                                                    default:
                                                        if true {
                                                            v1 = v1
                                                            switch v1 {
                                                            case "foo bar":
                                                                v6 = v6
                                                            }
                                                        } else {
                                                            var v83 bool = (1 == 64)
                                                            _ = v83
                                                        }
                                                    }
                                                default:
                                                    var v84 int64 = (^(v3))
                                                    _ = v84
                                                }
                                                if len(v68) >= 2 {
                                                    v68 = v68[1:len(v68)]
                                                }
                                            case "hello":
                                                v5 -= v5
                                            default:
                                                v85 := GenS1{F0: 417}
                                                _ = v85
                                            }
                                            if len(v68) >= 2 {
                                                v68 = v68[1:len(v68)]
                                            }
                                            v63 -= -2.25
                                            _ = v66
                                            _ = v67
                                            _ = v68
                                            _ = v75
                                            _ = v76
                                            _ = v77
                                            _ = v81
                                            _ = v82
                                        }
                                        v86 := GenS3{F0: v5}
                                        _ = v62
                                        _ = v63
                                        _ = v86
                                    }
                                    fmt.Println(fi42, fi45, v4)
                                }
                            }
                            var v87 int16 = -32768
                            _ = fi45
                            _ = v87
                        }
                        switch 32767 {
                        case 32768:
                            v3 += v5
                        case 822:
                            fmt.Println(fi42)
                        default:
                            switch "hello" {
                            case "a":
                                fmt.Println(v4, v6, v5)
                            case "\n\t":
                                fmt.Println(fi42)
                            }
                        }
                    }
                default:
                    var v88 uint16 = v6
                    _ = v88
                }
            default:
                v1 = "a"
            }
            var v89 uint32 = 2
            _ = fi42
            _ = v89
        }
        v90 := GenS8{F0: 1, F1: 0, F2: 0}
        _ = v90
    }
    if false {
        switch -9223372036854775808 {
        case 749:
            fmt.Println(v39, v6)
        case -9223372036854775808:
            for fi91 := 0; fi91 < 5; fi91++ {
                v39 -= 32767
                _ = fi91
            }
        case 4294967295:
            for fi92 := 0; fi92 < 17; fi92++ {
                v1 = v1
                for fi93 := 0; fi93 < 8; fi93++ {
                    v2[906] = fi92
                    _ = fi93
                }
                var v94 int64 = (v5 % ((-1) | 1))
                _ = fi92
                _ = v94
            }
        }
        v95 := GenS11{F0: 127}
        _ = v95
    }
    v5 += (v5 - v3)
    var v96 uint = 255
    var v97 int64 = 150
    v3 = v5
    var v98 uint32 = 65535
    v99 := v2[v98]
    switch v1 {
    case "hello":
        fmt.Println(v39, v99, v4)
    }
    v100 := GenS12{F0: 42}
    v1 = ""
    var v101 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v2
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v96
    _ = v97
    _ = v98
    _ = v99
}
