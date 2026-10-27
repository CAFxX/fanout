package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 uint32; F1 int64 }
type GenS8 struct { F0 uint32; F1 bool; F2 uint16 }
func genFn0() int32 {
    if (true && (true || false)) {
        for fi4 := 0; fi4 < 19; fi4++ {
            fmt.Println(fi4)
            v5 := []int{(fi4 & fi4), fi4, 42, (fi4 + (fi4 + (^(fi4))))}
            var v6 uint8 = 128
            _ = fi4
            _ = v5
            _ = v6
        }
    }
    var v7 float64 = (1.5 * (((-2.25 - (0.1 + 0.0)) / 3.14159) * ((0.0 - (((0.0 - 1e10) - 3.14159) + 1.5)) - 0.1)))
    v7 -= ((v7 + 1.5) - v7)
    if (42 == 42) {
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 uint32 = 2
            v10 := []int{127, fi8, fi8}
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := []int{0, -2147483648, 9223372036854775807, -9223372036854775808}
        _ = v11
    }
    _ = v7
    return 2147483647
}
func genFn1(p0 string) int8 {
    p0 = p0
    if (true && (42 < -1)) {
        p0 = p0
    } else {
        if true {
            var v12 uint16 = 42
            for fi13 := 0; fi13 < 18; fi13++ {
                for fi14 := 0; fi14 < 1; fi14++ {
                    var v15 int8 = 0
                    var v16 float32 = (0.0 * 1.5)
                    if false {
                        v17 := []int{fi14, 128, fi14, fi13}
                        _ = v17
                    } else {
                        v12 -= (v12 * 950)
                        var v18 int8 = 42
                        _ = v18
                    }
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                switch "hello" {
                case "foo bar":
                    var v19 int8 = 0
                    _ = v19
                case "hello":
                    var v20 string = ((((p0 + (p0 + "")) + (p0 + "a")) + "\n\t") + ((("hello" + "a") + p0) + (((p0 + "") + "hello") + "foo bar")))
                    _ = v20
                case "a":
                    v21 := GenS1{F0: (v12 % (((^(v12))) | 1)), F1: -32768, F2: -406}
                    _ = v21
                }
                switch "a" {
                case "hello":
                    v22 := GenS1{F0: v12, F1: 32767, F2: 0}
                    _ = v22
                case "foo bar":
                    if true {
                        var v23 bool = (("" + "") > ((p0 + (p0 + p0)) + p0))
                        p0 = ""
                        if (-638 <= -207) {
                            for fi24 := 0; fi24 < 8; fi24++ {
                                var v25 int16 = -32768
                                var v26 int32 = 2147483647
                                var v27 int8 = 0
                                switch (("hello" + ("\n\t" + ("\n\t" + p0))) + (p0 + p0)) {
                                case "":
                                    var v28 float32 = 0.1
                                    _ = v28
                                case "hello":
                                    v29 := []int{569, fi24, (fi13 / ((1) | 1)), (fi13 - (fi24 * (fi24 | (fi24 % ((1) | 1)))))}
                                    _ = v29
                                }
                                var v30 uint8 = 128
                                fmt.Println(fi24, v12, fi13)
                                _ = fi24
                                _ = v25
                                _ = v26
                                _ = v27
                                _ = v30
                            }
                        } else {
                            fmt.Println(v12, fi13, v23)
                            for fi31 := 0; fi31 < 20; fi31++ {
                                v23 = v23
                                switch ("foo bar" + (p0 + "a")) {
                                case "":
                                    var v32 uint8 = 183
                                    _ = v32
                                case "hello":
                                    v33 := GenS2{F0: -1, F1: "hello"}
                                    _ = v33
                                }
                                fmt.Println(v23, v23)
                                _ = fi31
                            }
                        }
                        _ = v23
                    } else {
                        if ((fi13 + (-(fi13))) >= fi13) {
                            switch fi13 {
                            case 4294967295:
                                v34 := []int{(fi13 >> ((0) & 63)), (^(fi13))}
                                _ = v34
                            case 943:
                                fmt.Println(fi13)
                            case 127:
                                v35 := GenS3{F0: true}
                                _ = v35
                            }
                            v12 += (v12 - (v12 % ((v12) | 1)))
                        } else {
                            for fi36 := 0; fi36 < 19; fi36++ {
                                p0 = p0
                                p0 = p0
                                _ = fi36
                            }
                        }
                        switch "a" {
                        case "foo bar":
                            fmt.Println(p0, v12)
                        case "hello":
                            if (((32767 <= (v12 - v12)) || true) || (!true)) {
                                if (!true) {
                                    fmt.Println(fi13, fi13, fi13)
                                    switch (p0 + ((p0 + p0) + (p0 + ((("a" + "\n\t") + p0) + (((p0 + p0) + p0) + p0))))) {
                                    case "\n\t":
                                        switch p0 {
                                        case "":
                                            fmt.Println(fi13)
                                        }
                                    case "hello":
                                        var v37 int = (fi13 + fi13)
                                        _ = v37
                                    case "foo bar":
                                        v38 := GenS4{F0: -1, F1: 0}
                                        _ = v38
                                    }
                                    v39 := GenS4{F0: 1, F1: 744}
                                    _ = v39
                                }
                                v40 := []int{fi13, (-(fi13)), fi13, fi13}
                                _ = v40
                            } else {
                                var v41 int64 = -114
                                fmt.Println(v41)
                                v42 := GenS5{F0: fi13}
                                _ = v41
                                _ = v42
                            }
                        }
                        v43 := []int{(fi13 | (fi13 << ((65535) & 63))), fi13}
                        _ = v43
                    }
                }
                _ = fi13
            }
            _ = v12
        } else {
            v44 := GenS1{F0: 32767, F1: -32768, F2: 0}
            _ = v44
        }
    }
    _ = p0
    return 0
}

func main() {
    var v1 int8 = 19
    var v2 bool = true
    var v3 uint8 = 0
    var v45 uint8 = (v3 % (((v3 & (v3 / ((255) | 1)))) | 1))
    v46 := []int{-1, -2147483648, -312}
    var v47 float32 = 0.0
    for fi48 := 0; fi48 < 10; fi48++ {
        if ((0 > 0) && (!v2)) {
            v46 = append(v46, (-(fi48)))
            var v49 float64 = -2.25
            _ = v49
        }
        if v2 {
            fmt.Println(v47)
            if true {
                v50 := GenS5{F0: 666}
                var v51 int64 = 1
                _ = v50
                _ = v51
            } else {
                if (false && (2 <= 105)) {
                    var v52 int32 = 1
                    v52 += (^(v52))
                    _ = v52
                } else {
                    fmt.Println(fi48)
                }
                if false {
                    for fi53 := 0; fi53 < 5; fi53++ {
                        if len(v46) >= 2 {
                            v46 = v46[1:len(v46)]
                        }
                        _ = fi53
                    }
                    var v54 uint32 = 2147483648
                    _ = v54
                } else {
                    var v55 float32 = 1e10
                    _ = v55
                }
                v56 := GenS4{F0: -2147483648, F1: 1}
                _ = v56
            }
            for fi57 := 0; fi57 < 9; fi57++ {
                var v58 int16 = -32768
                _ = fi57
                _ = v58
            }
        } else {
            v3 = 42
        }
        _ = fi48
    }
    v2 = v2
    switch 32767 {
    case 890:
        v3 = (v3 * (v3 + v3))
    default:
        if v2 {
            v47 = 0.0
        } else {
            if true {
                var v59 uint8 = (v3 & (v45 & (^(v45))))
                v59 = 255
                _ = v59
            }
        }
    }
    for fi60 := 0; fi60 < 11; fi60++ {
        var v61 int8 = v1
        if v2 {
            var v62 uint = 18446744073709551615
            v47 -= (v47 + (float32((0.0 - (float64(v47) + (3.14159 * float64((3.14159 / float32((float64(((v47 / 1e10) + (v47 * float32(1.5)))) + 1e10)))))))) - 0.1))
            if (!(!v2)) {
                switch 81 {
                case 2147483647:
                    if false {
                        if len(v46) > 0 {
                            v46[0] = (fi60 / (((fi60 / ((9223372036854775807) | 1))) | 1))
                        }
                        v3 = (v45 % ((127) | 1))
                    }
                case 980:
                    var v63 int8 = genFn1("")
                    _ = v63
                }
                var v64 uint8 = 2
                var v65 bool = ((((1.5 - 1e10) - (-2.25 + -2.25)) + (-2.25 / 1e10)) > (0.0 * -2.25))
                _ = v64
                _ = v65
            } else {
                fmt.Println(v2, v45, v47)
                var v66 string = (("hello" + "a") + ("" + "\n\t"))
                var v67 int = -2147483648
                _ = v66
                _ = v67
            }
            _ = v62
        } else {
            v46 = append(v46, fi60)
        }
        var v68 uint32 = 2
        _ = fi60
        _ = v61
        _ = v68
    }
    v69 := GenS6{F0: 37, F1: 18446744073709551615}
    fmt.Println(v1)
    v2 = (-288 <= -2147483648)
    fmt.Println(v45, v1)
    for fi70 := 0; fi70 < 9; fi70++ {
        switch fi70 {
        case 4294967295:
            v1 = 0
        case 127:
            if ((fi70 >> ((4294967295) & 63)) <= fi70) {
                var v71 bool = (!(("a" + "") < ""))
                _ = v71
            } else {
                if (!true) {
                    var v72 int8 = v1
                    _ = v72
                } else {
                    v45 = v45
                    for fi73 := 0; fi73 < 19; fi73++ {
                        var v74 bool = false
                        if v74 {
                            fmt.Println(v2, v2, v3)
                            var v75 int = fi70
                            _ = v75
                        } else {
                            fmt.Println(fi70)
                        }
                        _ = fi73
                        _ = v74
                    }
                }
                v47 = (v47 - 0.1)
                v76 := GenS7{F0: 2147483647, F1: -340}
                _ = v76
            }
        }
        _ = fi70
    }
    var v77 int = -982
    v3 = (v45 << ((904) & 63))
    if (genFn0() == -1) {
        fmt.Println(v1)
        fmt.Println(v3, v2)
    } else {
        if (1000 >= -162) {
            v78 := GenS7{F0: 2147483648, F1: -987}
            v3 = v45
            switch 20 {
            case 9223372036854775807:
                _ = len(v46)
            default:
                v46 = append(v46, 0)
            }
            _ = v78
        } else {
            for fi79 := 0; fi79 < 4; fi79++ {
                if ((v47 * (v47 * (3.14159 * (v47 + ((1e10 / (1.5 / (v47 / ((((0.1 + -2.25) * (((v47 + 3.14159) / float32((-2.25 + 0.0))) - (v47 - (float32(0.0) / 1.0)))) * v47) - v47)))) / ((((-2.25 * float32((1e10 / (3.14159 * 1.5)))) / 1.0) * float32(0.1)) / (((0.0 + (v47 / ((float32(0.0) / v47) * 0.0))) - v47) / (-2.25 / v47)))))))) != v47) {
                    var v80 uint64 = 2
                    _ = v80
                } else {
                    for fi81 := 0; fi81 < 2; fi81++ {
                        if len(v46) >= 2 {
                            v46 = v46[1:len(v46)]
                        }
                        var v82 int64 = 42
                        v83 := GenS8{F0: 1, F1: v2, F2: 42}
                        _ = fi81
                        _ = v82
                        _ = v83
                    }
                }
                switch "" {
                case "a":
                    v45 -= 127
                }
                fmt.Println(v3, v1, v45)
                _ = fi79
            }
        }
    }
    v77 = v77
    if (0 != -9223372036854775808) {
        for fi84 := 0; fi84 < 17; fi84++ {
            for fi85 := 0; fi85 < 6; fi85++ {
                if len(v46) >= 2 {
                    v46 = v46[1:len(v46)]
                }
                v1 += genFn1("\n\t")
                _ = fi85
            }
            switch v77 {
            case 723:
                fmt.Println(v3)
            case -9223372036854775808:
                var v86 int32 = -823
                _ = v86
            case 32767:
                fmt.Println(v45)
            default:
                if v2 {
                    var v87 int32 = -20
                    _ = v87
                }
            }
            v2 = ((v45 - (v3 << ((1) & 63))) > v3)
            _ = fi84
        }
        if len(v46) >= 2 {
            v46 = v46[1:len(v46)]
        }
        v88 := GenS1{F0: 190, F1: 32767, F2: 1}
        _ = v88
    } else {
        fmt.Println(v47)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v47
    _ = v69
    _ = v77
}
