package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int; F1 string; F2 int32 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64; F1 int; F2 string }
func genFn0() uint32 {
    if true {
        for fi4 := 0; fi4 < 12; fi4++ {
            switch ("" + "a") {
            case "hello":
                var v5 uint32 = 1
                _ = v5
            default:
                fmt.Println(fi4)
            }
            v6 := []int{fi4, (fi4 / (((fi4 / ((fi4) | 1))) | 1))}
            var v7 int = fi4
            _ = fi4
            _ = v6
            _ = v7
        }
        v8 := []int{846, 9223372036854775807, 4294967295}
        if len(v8) > 0 {
            v8[0] = -9223372036854775808
        }
        _ = v8
    }
    var v9 uint = 255
    for fi10 := 0; fi10 < 12; fi10++ {
        switch "foo bar" {
        case "foo bar":
            switch "\n\t" {
            case "a":
                v9 = v9
            case "foo bar":
                var v11 int = (-(fi10))
                _ = v11
            default:
                v9 -= (v9 ^ 4294967295)
            }
        case "a":
            switch "\n\t" {
            case "\n\t":
                v12 := GenS1{F0: fi10}
                _ = v12
            case "hello":
                switch fi10 {
                case 80:
                    var v13 uint64 = 805
                    _ = v13
                default:
                    fmt.Println(fi10, v9)
                }
            }
        default:
            var v14 uint = v9
            _ = v14
        }
        var v15 int = (fi10 ^ 128)
        var v16 uint64 = 18446744073709551615
        _ = fi10
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        v18 := GenS2{F0: 0, F1: 9223372036854775807}
        _ = fi17
        _ = v18
    }
    var v19 float64 = (float64(0.1) + 1.5)
    _ = v19
    _ = v9
    return 4294967295
}
func genFn1(p0 uint32, p1 uint, p2 string) uint {
    p2 = p2
    if true {
        switch ("" + p2) {
        case "hello":
            var v20 uint8 = 1
            _ = v20
        default:
            var v21 uint32 = 65535
            _ = v21
        }
        var v22 uint = p1
        var v23 uint32 = 1
        _ = v22
        _ = v23
    }
    fmt.Println(p2, p0)
    for fi24 := 0; fi24 < 18; fi24++ {
        for fi25 := 0; fi25 < 18; fi25++ {
            p0 = (p0 ^ (p0 << (((p1 << ((p1) & 63))) & 63)))
            fmt.Println(p0, p1, p2)
            var v26 int8 = 100
            _ = fi25
            _ = v26
        }
        var v27 uint64 = 18446744073709551615
        _ = fi24
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 uint64 = 42
    var v2 bool = true
    v3 := []uint32{1}
    if v2 {
        if (v2 || v2) {
            v28 := GenS3{F0: -9223372036854775808, F1: "\n\t", F2: 864}
            var v29 int = 497
            _ = v28
            _ = v29
        } else {
            if v2 {
                if (v2 && false) {
                    v1 -= 1
                    _ = len(v3)
                } else {
                    if (v2 || (100 != 42)) {
                        var v30 uint32 = 4294967295
                        _ = v30
                    } else {
                        v1 -= (v1 / (((v1 >> ((79) & 63))) | 1))
                        var v31 string = "a"
                        v32 := GenS2{F0: 1, F1: 793}
                        _ = v31
                        _ = v32
                    }
                    var v33 uint16 = 221
                    for fi34 := 0; fi34 < 9; fi34++ {
                        v35 := GenS1{F0: fi34}
                        for fi36 := 0; fi36 < 12; fi36++ {
                            v33 += (v33 ^ v33)
                            var v37 int64 = -1
                            _ = fi36
                            _ = v37
                        }
                        switch ("foo bar" + "foo bar") {
                        case "hello":
                            v2 = v2
                        case "a":
                            switch (((("hello" + ("foo bar" + "")) + "foo bar") + (("" + "hello") + "\n\t")) + "") {
                            case "":
                                var v38 int32 = -1
                                _ = v38
                            case "hello":
                                v33 += 2
                            case "foo bar":
                                var v39 string = (("" + "foo bar") + "")
                                _ = v39
                            default:
                                var v40 int8 = -1
                                _ = v40
                            }
                        }
                        _ = fi34
                        _ = v35
                    }
                    _ = v33
                }
            }
            v2 = ((-2.25 < (3.14159 + 3.14159)) || (-1 == 0))
        }
        v2 = ((^(v1)) == v1)
    } else {
        var v41 int8 = 42
        v42 := GenS4{F0: 1, F1: (v1 % (((v1 * (v1 / ((189) | 1)))) | 1))}
        v1 += 1
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 1; fi43++ {
        switch "foo bar" {
        case "foo bar":
            v1 += (v1 & 1)
        case "\n\t":
            var v44 uint16 = 65535
            _ = v44
        }
        v1 -= (v1 ^ 65)
        var v45 int16 = -71
        _ = fi43
        _ = v45
    }
    v1 = v1
    v1 = v1
    if true {
        if (65535 <= 18446744073709551615) {
            v1 -= v1
            switch -9223372036854775808 {
            case 1:
                v46 := GenS5{F0: 234, F1: 42}
                _ = v46
            default:
                v47 := GenS5{F0: -1, F1: 13}
                _ = v47
            }
        } else {
            var v48 uint32 = 22
            var v49 bool = ((!(v2 && (127 != -1))) || (!v2))
            _ = v48
            _ = v49
        }
    } else {
        switch -801 {
        case 510:
            for fi50 := 0; fi50 < 18; fi50++ {
                var v51 uint8 = 42
                _ = fi50
                _ = v51
            }
        default:
            if true {
                if false {
                    for fi52 := 0; fi52 < 12; fi52++ {
                        if false {
                            v2 = v2
                        }
                        _ = fi52
                    }
                    fmt.Println(v2, v1)
                    v1 -= (v1 >> ((540) & 63))
                }
                for fi53 := 0; fi53 < 18; fi53++ {
                    fmt.Println(v1)
                    _ = fi53
                }
            }
        }
        v54 := GenS6{F0: 1000}
        _ = v54
    }
    fmt.Println(v1, v1)
    fmt.Println(v2, v1)
    var v55 float32 = -2.25
    v1 += v1
    v55 = float32(0.1)
    for fi56 := 0; fi56 < 6; fi56++ {
        var v57 bool = v2
        for fi58 := 0; fi58 < 15; fi58++ {
            var v59 int32 = -2147483648
            if (!v57) {
                switch fi56 {
                case 32768:
                    v57 = v2
                case 187:
                    fmt.Println(v2)
                }
                v2 = (54 <= 630)
            } else {
                var v60 bool = ((!(-888 <= 1000)) && ((^(fi58)) != (fi58 >> ((0) & 63))))
                switch ("" + ("hello" + (("foo bar" + "foo bar") + "a"))) {
                case "":
                    if true {
                        var v61 float32 = v55
                        var v62 float32 = v61
                        _ = v61
                        _ = v62
                    } else {
                        if (32767 > -565) {
                            if v57 {
                                var v63 int = (fi56 >> ((genFn1(4294967295, 547, ("\n\t" + "hello"))) & 63))
                                v64 := GenS2{F0: 0, F1: -9223372036854775808}
                                var v65 int64 = -9223372036854775808
                                _ = v63
                                _ = v64
                                _ = v65
                            }
                        } else {
                            v3 = append(v3, 2147483648)
                            v60 = true
                        }
                        if v2 {
                            if ((445 <= 725) || true) {
                                v59 = v59
                                var v66 bool = v60
                                _ = v66
                            } else {
                                if (v2 && ((-711 >= -9223372036854775808) && (!v2))) {
                                    if v57 {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                } else {
                                    if false {
                                        var v67 bool = false
                                        v68 := GenS1{F0: (-(fi56))}
                                        var v69 uint32 = 55
                                        _ = v67
                                        _ = v68
                                        _ = v69
                                    }
                                    var v70 int16 = 708
                                    _ = v70
                                }
                                var v71 uint16 = 255
                                _ = v71
                            }
                            fmt.Println(v59, v57)
                            switch "hello" {
                            case "\n\t":
                                var v72 uint = 65535
                                _ = v72
                            case "a":
                                var v73 int = fi56
                                _ = v73
                            case "foo bar":
                                if (1.5 < 3.14159) {
                                    fmt.Println(fi56)
                                    if true {
                                        var v74 float64 = (0.0 / 3.14159)
                                        fmt.Println(v60)
                                        if len(v3) > 0 {
                                            v3[0] = 2
                                        }
                                        _ = v74
                                    }
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            default:
                                if v60 {
                                    v55 -= v55
                                    if v2 {
                                        if len(v3) > 0 {
                                            v3[0] = 776
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v57 = v2
                                        var v75 float64 = 1e10
                                        v76 := GenS3{F0: fi56, F1: ("" + ((("\n\t" + "hello") + "hello") + "hello")), F2: 1}
                                        _ = v75
                                        _ = v76
                                    }
                                    var v77 float64 = -2.25
                                    _ = v77
                                } else {
                                    v1 += v1
                                    fmt.Println(v2, fi56)
                                    v60 = true
                                }
                            }
                        } else {
                            if (1 == 1) {
                                v55 = -2.25
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                fmt.Println(v1, v1, v57)
                            }
                            v78 := GenS7{F0: 1, F1: (fi56 << ((255) & 63)), F2: ("hello" + "")}
                            _ = len(v3)
                            _ = v78
                        }
                        var v79 float64 = (-2.25 - float64(float32((0.1 + 0.0))))
                        if len(v3) > 0 {
                            v3[0] = 1
                        }
                        v3 = append(v3, 2147483648)
                        var v80 float32 = (v55 / (3.14159 + v55))
                        _ = v79
                        _ = v80
                    }
                default:
                    var v81 uint8 = 127
                    _ = v81
                }
                _ = v60
            }
            if len(v3) > 0 {
                v3[0] = 65535
            }
            _ = fi58
            _ = v59
        }
        _ = fi56
        _ = v57
    }
    v55 += (((float32(3.14159) / v55) + 3.14159) + (v55 * v55))
    v55 += v55
    for fi82 := 0; fi82 < 13; fi82++ {
        v55 = float32(float64(((v55 / v55) - -2.25)))
        _ = fi82
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v55
}
