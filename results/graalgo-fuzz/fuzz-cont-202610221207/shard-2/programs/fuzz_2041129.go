package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS6 struct { F0 uint }
func genFn0(p0 uint8) uint32 {
    p0 -= p0
    var v4 int32 = 0
    v5 := []int{1}
    fmt.Println(p0, v4)
    switch (("" + "\n\t") + ("a" + ("\n\t" + "a"))) {
    case "":
        v4 = (^(v4))
    case "foo bar":
        switch "" {
        case "\n\t":
            v4 += (^(v4))
        case "a":
            var v6 uint32 = 0
            _ = v6
        }
    }
    _ = p0
    _ = v4
    _ = v5
    return 2
}
func genFn1(p0 uint16, p1 bool) int {
    p1 = (0 > 128)
    for fi7 := 0; fi7 < 4; fi7++ {
        if p1 {
            var v8 bool = p1
            var v9 uint32 = 1
            var v10 int32 = 418
            _ = v10
            _ = v8
            _ = v9
        }
        v11 := []int{32768}
        var v12 int = fi7
        _ = fi7
        _ = v11
        _ = v12
    }
    v13 := []int{-9223372036854775808, 269, -2147483648}
    p0 += (p0 >> ((2) & 63))
    _ = p0
    _ = p1
    _ = v13
    return 4294967295
}

func main() {
    var v1 int64 = 1
    var v2 float32 = (3.14159 / float32((-2.25 / (-2.25 * (1e10 / 3.14159)))))
    v3 := make(map[string]int)
    if (!false) {
        var v14 int16 = 0
        if (-1 <= -1) {
            var v15 uint = 255
            _ = v15
        }
        _ = v14
    } else {
        switch -9223372036854775808 {
        case 0:
            var v16 float32 = -2.25
            _ = v16
        case -2147483648:
            for fi17 := 0; fi17 < 3; fi17++ {
                for fi18 := 0; fi18 < 1; fi18++ {
                    var v19 float64 = (0.0 + 0.1)
                    v20 := GenS1{F0: false}
                    v19 -= ((float64(v2) - v19) + v19)
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                for fi21 := 0; fi21 < 6; fi21++ {
                    v22 := GenS2{F0: false}
                    switch "\n\t" {
                    case "foo bar":
                        v2 = (1e10 - v2)
                    }
                    _ = fi21
                    _ = v22
                }
                _ = fi17
            }
        }
        if false {
            var v23 uint32 = 1
            if ("a" == "foo bar") {
                v23 -= (^(v23))
                switch -961 {
                case 554:
                    v2 += v2
                case 600:
                    for fi24 := 0; fi24 < 6; fi24++ {
                        var v25 float64 = ((-2.25 / -2.25) - float64(v2))
                        for fi26 := 0; fi26 < 15; fi26++ {
                            v27 := []int{fi26, fi24}
                            switch (fi26 | (fi24 | genFn1(0, (4294967295 >= (fi26 * fi24))))) {
                            case 42:
                                v2 = 0.0
                            case -9223372036854775808:
                                v28 := GenS2{F0: false}
                                _ = v28
                            default:
                                v23 -= genFn0(42)
                            }
                            _ = fi26
                            _ = v27
                        }
                        _ = fi24
                        _ = v25
                    }
                case 32767:
                    switch "foo bar" {
                    case "a":
                        if (!false) {
                            v1 += v1
                        }
                    case "foo bar":
                        switch 1 {
                        case 0:
                            var v29 int8 = 1
                            _ = v29
                        }
                    }
                default:
                    v30 := []int{32767, -902, 0, -2147483648}
                    _ = v30
                }
            } else {
                fmt.Println(v23)
                var v31 int = 32768
                v2 = float32(3.14159)
                _ = v31
            }
            _ = v23
        } else {
            for fi32 := 0; fi32 < 16; fi32++ {
                for fi33 := 0; fi33 < 9; fi33++ {
                    if false {
                        var v34 int8 = 1
                        if (!(-2.25 > (v2 - v2))) {
                            var v35 bool = true
                            _ = v35
                        } else {
                            var v36 int64 = 0
                            if (1 >= 255) {
                                v36 += (-(v36))
                                var v37 uint64 = 4294967295
                                var v38 uint8 = 1
                                var v39 int16 = 1000
                                var v40 uint = 835
                                var v41 string = "\n\t"
                                if false {
                                    v3[((v41 + "") + v41)] = genFn1(347, ((v38 & (^(v38))) > (v38 ^ (v38 & v38))))
                                }
                                _ = v37
                                _ = v38
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                            v2 -= (0.0 - float32(1.5))
                            _ = v36
                        }
                        _ = v34
                    } else {
                        if (genFn1(672, ((v1 * (v1 + v1)) > v1)) != 32767) {
                            v42 := []int{9223372036854775807, fi32, (fi32 << ((810) & 63))}
                            _ = v42
                        } else {
                            var v43 float64 = 1e10
                            if true {
                                if (true || true) {
                                    v44 := GenS3{F0: (v1 >> ((255) & 63)), F1: 36}
                                    v43 += v43
                                    var v45 bool = (!("foo bar" < ("foo bar" + "\n\t")))
                                    _ = v44
                                    _ = v45
                                }
                            } else {
                                var v46 int = genFn1(255, true)
                                _ = v46
                            }
                            fmt.Println(v2)
                            var v47 int = (fi32 + fi32)
                            _ = v43
                            _ = v47
                        }
                        var v48 uint32 = 0
                        v2 -= (0.1 / -2.25)
                        _ = v48
                    }
                    _ = fi33
                }
                _ = fi32
            }
        }
    }
    if false {
        switch "" {
        case "\n\t":
            v49 := GenS4{F0: -1}
            _ = v49
        default:
            if false {
                for fi50 := 0; fi50 < 1; fi50++ {
                    if false {
                        switch "hello" {
                        case "a":
                            for fi51 := 0; fi51 < 17; fi51++ {
                                var v52 int16 = 32767
                                var v53 int = 9223372036854775807
                                var v54 int16 = v52
                                _ = fi51
                                _ = v52
                                _ = v53
                                _ = v54
                            }
                        case "\n\t":
                            if (v1 == (v1 % ((0) | 1))) {
                                v55 := GenS1{F0: true}
                                v56 := GenS1{F0: true}
                                _ = len(v3)
                                _ = v55
                                _ = v56
                            }
                        default:
                            v1 = (v1 >> ((499) & 63))
                        }
                        var v57 float32 = 0.0
                        _ = v57
                    }
                    var v58 int32 = -2147483648
                    v58 = (v58 ^ v58)
                    _ = fi50
                    _ = v58
                }
            }
        }
        for fi59 := 0; fi59 < 8; fi59++ {
            if (-1 == -1) {
                for fi60 := 0; fi60 < 17; fi60++ {
                    v1 += 839
                    v61 := GenS4{F0: 42}
                    _ = fi60
                    _ = v61
                }
                var v62 int8 = 1
                v62 += v62
                _ = v62
            } else {
                var v63 uint8 = 196
                switch genFn1(899, false) {
                case 128:
                    fmt.Println(v1, v2, v2)
                case 1:
                    switch "foo bar" {
                    case "":
                        if false {
                            fmt.Println(v2)
                            fmt.Println(v1)
                            for fi64 := 0; fi64 < 9; fi64++ {
                                var v65 uint8 = 0
                                var v66 int32 = 42
                                var v67 int = (fi59 << ((42) & 63))
                                _ = fi64
                                _ = v65
                                _ = v66
                                _ = v67
                            }
                        }
                    case "a":
                        v63 = (v63 << ((618) & 63))
                    case "hello":
                        v3[((("hello" + "foo bar") + ("a" + "")) + (("hello" + "hello") + "\n\t"))] = -829
                    default:
                        v68 := GenS5{F0: 127, F1: 1, F2: 65535}
                        _ = v68
                    }
                default:
                    v69 := GenS3{F0: (-(v1)), F1: 42}
                    _ = v69
                }
                _ = v63
            }
            _ = fi59
        }
        fmt.Println(v2, v1)
    }
    for fi70 := 0; fi70 < 18; fi70++ {
        var v71 uint64 = 4294967295
        var v72 uint32 = uint32(v71)
        _ = fi70
        _ = v71
        _ = v72
    }
    v1 -= 786
    for fi73 := 0; fi73 < 1; fi73++ {
        var v74 int32 = 878
        v75 := []int{fi73, fi73, 409, (fi73 << ((0) & 63))}
        v2 = v2
        _ = fi73
        _ = v74
        _ = v75
    }
    for fi76 := 0; fi76 < 10; fi76++ {
        v3[""] = 1
        switch genFn1(42, (42 > 254)) {
        case 368:
            for fi77 := 0; fi77 < 15; fi77++ {
                if (32767 <= 1) {
                    v78 := []int{2147483647, (fi77 % ((9223372036854775807) | 1))}
                    fmt.Println(v2, fi76)
                    _ = v78
                } else {
                    if true {
                        if false {
                            var v79 string = "a"
                            v3[v79] = fi76
                            switch fi77 {
                            case -2147483648:
                                v79 = (v79 + v79)
                            }
                            _ = v79
                        }
                        var v80 int = fi77
                        _ = v80
                    } else {
                        var v81 string = ("\n\t" + "\n\t")
                        if (true || true) {
                            var v82 int = genFn1(42, (!true))
                            v83 := []int{fi76, fi76}
                            if true {
                                fmt.Println(v2, v82)
                                v2 += (v2 * 1.5)
                                if len(v83) >= 2 {
                                    v83 = v83[1:len(v83)]
                                }
                            } else {
                                if false {
                                    _ = len(v3)
                                    var v84 uint16 = 42
                                    v84 = (v84 * 2)
                                    _ = v84
                                }
                            }
                            v2 -= v2
                            fmt.Println(fi76)
                            switch ((v81 + v81) + "foo bar") {
                            case "\n\t":
                                v82 += (-(fi76))
                            case "":
                                v85 := GenS6{F0: 300}
                                _ = v85
                            case "a":
                                v86 := GenS3{F0: v1, F1: 130}
                                _ = v86
                            default:
                                v83 = append(v83, fi77)
                            }
                            _ = v82
                            _ = v83
                        }
                        v81 = ((v81 + v81) + "foo bar")
                        _ = v81
                    }
                    var v87 uint32 = 2147483648
                    v88 := []int{127, fi77, 4294967295, genFn1(42, ((-9223372036854775808 < v1) && (v2 > v2)))}
                    _ = v87
                    _ = v88
                }
                var v89 uint8 = 140
                v90 := GenS2{F0: (0 < 1)}
                _ = fi77
                _ = v89
                _ = v90
            }
        }
        fmt.Println(v2, fi76, v1)
        _ = fi76
    }
    v1 = 251
    v91 := []int{-643, genFn1(255, true), 1}
    var v92 uint16 = 32767
    var v93 uint64 = 750
    v92 = 788
    var v94 uint32 = genFn0(0)
    v91 = append(v91, 4294967295)
    fmt.Println(v92, v2)
    for fi95 := 0; fi95 < 15; fi95++ {
        if (0 != 127) {
            if true {
                v96 := v3[("hello" + "foo bar")]
                _ = v96
            }
        } else {
            var v97 float64 = 1e10
            v98 := GenS1{F0: true}
            var v99 int16 = 1
            _ = v97
            _ = v98
            _ = v99
        }
        _ = fi95
    }
    v100 := GenS5{F0: 2, F1: v93, F2: v92}
    for fi101 := 0; fi101 < 15; fi101++ {
        for fi102 := 0; fi102 < 1; fi102++ {
            v93 = (v93 & v93)
            v1 += v1
            var v103 uint64 = 0
            _ = fi102
            _ = v103
        }
        switch "foo bar" {
        case "":
            v91 = append(v91, (fi101 << ((4294967295) & 63)))
        case "hello":
            for fi104 := 0; fi104 < 2; fi104++ {
                _ = len(v3)
                _ = len(v3)
                var v105 float64 = -2.25
                _ = fi104
                _ = v105
            }
        case "foo bar":
            v92 = (v92 & 0)
        default:
            var v106 int64 = v1
            _ = v106
        }
        fmt.Println(v1, v1)
        _ = fi101
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v91:", len(v91), v91)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v91
    _ = v92
    _ = v93
    _ = v94
}
