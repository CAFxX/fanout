package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint16 }
func genFn0() int64 {
    var v7 int = 2147483647
    if (!true) {
        if true {
            for fi8 := 0; fi8 < 4; fi8++ {
                v9 := []int{(fi8 % ((-9223372036854775808) | 1)), 4294967295, (fi8 - v7), -693}
                if len(v9) > 0 {
                    v9[0] = (fi8 % (((^(v7))) | 1))
                }
                _ = fi8
                _ = v9
            }
            switch (^(v7)) {
            case 32768:
                var v10 float64 = (0.1 / 3.14159)
                _ = v10
            case 423:
                for fi11 := 0; fi11 < 14; fi11++ {
                    v12 := []int{127, (fi11 % (((fi11 & 32768)) | 1))}
                    if true {
                        fmt.Println(v7, fi11)
                    }
                    _ = fi11
                    _ = v12
                }
            }
            v7 = (v7 ^ (v7 << ((1) & 63)))
        }
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 uint = 4294967295
        fmt.Println(fi13)
        _ = fi13
        _ = v14
    }
    fmt.Println(v7)
    v15 := []int{(v7 / ((42) | 1)), (v7 * (v7 & (v7 * 1)))}
    _ = v15
    _ = v7
    return int64(v7)
}

func main() {
    v1 := []string{}
    v2 := make(map[int]bool)
    var v3 float32 = -2.25
    var v4 string = "foo bar"
    v5 := []int{-790, 810}
    var v6 int8 = 42
    var v16 uint8 = 168
    var v17 uint32 = 2147483648
    if false {
        v18 := GenS1{F0: v4, F1: 20}
        _ = v18
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        if true {
            var v20 float64 = (0.0 * 1e10)
            _ = v20
        } else {
            switch fi19 {
            case 1:
                fmt.Println(v16)
            }
            v1 = append(v1, v4)
            var v21 int16 = -158
            _ = v21
        }
        _ = fi19
    }
    var v22 int8 = (v6 ^ v6)
    switch v4 {
    case "hello":
        v3 += v3
    case "foo bar":
        v2[-726] = (42 == 4294967295)
    case "\n\t":
        for fi23 := 0; fi23 < 1; fi23++ {
            switch v4 {
            case "foo bar":
                for fi24 := 0; fi24 < 10; fi24++ {
                    v5 = append(v5, fi23)
                    if true {
                        if false {
                            switch (v4 + "\n\t") {
                            case "hello":
                                var v25 int16 = 69
                                _ = v25
                            case "a":
                                switch (fi24 * 32767) {
                                case -2147483648:
                                    if false {
                                        if true {
                                            _ = len(v2)
                                        } else {
                                            var v26 bool = true
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            if v26 {
                                                v27 := GenS1{F0: v4, F1: 65535}
                                                _ = v27
                                            }
                                            _ = v26
                                        }
                                        var v28 int64 = 1
                                        fmt.Println(v3)
                                        _ = v28
                                    }
                                case 127:
                                    if true {
                                        if ((v17 ^ v17) != (v17 % ((4294967295) | 1))) {
                                            if true {
                                                if true {
                                                    var v29 float64 = 0.1
                                                    if false {
                                                        v30 := GenS1{F0: v4, F1: 547}
                                                        if false {
                                                            v31 := GenS1{F0: v4, F1: 32767}
                                                            _ = v31
                                                        }
                                                        _ = v30
                                                    } else {
                                                        v3 = 1e10
                                                        var v32 int32 = 511
                                                        _ = v32
                                                    }
                                                    var v33 float64 = (0.0 * ((float64(v3) * float64(float32((v29 - v29)))) - -2.25))
                                                    _ = v29
                                                    _ = v33
                                                } else {
                                                    var v34 int8 = 127
                                                    fmt.Println(fi23, v22, fi24)
                                                    var v35 uint16 = 32767
                                                    v36 := v2[(fi23 & -1)]
                                                    v3 -= v3
                                                    _ = v34
                                                    _ = v35
                                                    _ = v36
                                                }
                                                var v37 uint32 = (^(v17))
                                                var v38 float64 = 3.14159
                                                _ = v37
                                                _ = v38
                                            }
                                            switch v4 {
                                            case "\n\t":
                                                var v39 float32 = 0.0
                                                _ = v39
                                            case "a":
                                                var v40 uint8 = (v16 & (v16 + 127))
                                                _ = v40
                                            case "hello":
                                                v41 := GenS1{F0: "foo bar", F1: 228}
                                                _ = v41
                                            }
                                        }
                                    }
                                case -1:
                                    v42 := GenS2{F0: 2}
                                    _ = v42
                                default:
                                    v43 := GenS1{F0: v4, F1: 1}
                                    _ = v43
                                }
                            case "foo bar":
                                fmt.Println(v17, fi23)
                            default:
                                fmt.Println(v6, v3, v22)
                            }
                            var v44 uint = uint(v16)
                            _ = v44
                        }
                        v6 -= -54
                        switch 1 {
                        case 528:
                            v45 := GenS2{F0: 42}
                            _ = v45
                        }
                    }
                    _ = fi24
                }
            default:
                v22 = v6
            }
            var v46 int64 = -1
            _ = fi23
            _ = v46
        }
    default:
        var v47 uint64 = 2
        _ = v47
    }
    if ((!((v17 * v17) <= (v17 % (((v17 % (((v17 & 616)) | 1))) | 1)))) || true) {
        v17 -= 42
        v6 -= v22
        switch (v4 + "\n\t") {
        case "hello":
            var v48 bool = true
            _ = v48
        default:
            var v49 bool = (true && (2 == 255))
            _ = v49
        }
    }
    switch 9223372036854775807 {
    case 306:
        if (!(!((v16 * (v16 << ((0) & 63))) > (^(v16))))) {
            if (!true) {
                v50 := GenS2{F0: 138}
                v51 := v2[32767]
                v51 = true
                _ = v50
                _ = v51
            } else {
                var v52 int32 = 2147483647
                fmt.Println(v3)
                _ = v52
            }
            for fi53 := 0; fi53 < 11; fi53++ {
                v54 := GenS2{F0: 65535}
                if true {
                    v17 += (v17 - (v17 << ((uint(v17)) & 63)))
                    if false {
                        v16 = (v16 | 2)
                        _ = len(v5)
                    }
                    if false {
                        if len(v1) > 0 {
                            v1[0] = v4
                        }
                    } else {
                        var v55 uint16 = 2
                        v22 -= (v22 - 42)
                        if (false || true) {
                            if len(v5) > 0 {
                                v5[0] = (fi53 * 32768)
                            }
                        }
                        _ = v55
                    }
                }
                v17 -= (v17 | (v17 / (((v17 / ((v17) | 1))) | 1)))
                _ = fi53
                _ = v54
            }
            switch v4 {
            case "":
                v4 = v4
            case "a":
                fmt.Println(v22, v16)
            case "foo bar":
                var v56 bool = true
                _ = v56
            default:
                if false {
                    switch ("hello" + v4) {
                    case "foo bar":
                        fmt.Println(v6)
                    default:
                        for fi57 := 0; fi57 < 18; fi57++ {
                            fmt.Println(v16)
                            var v58 int32 = 0
                            if false {
                                var v59 int8 = -1
                                var v60 uint8 = 42
                                _ = v59
                                _ = v60
                            } else {
                                switch v4 {
                                case "hello":
                                    fmt.Println(fi57)
                                case "a":
                                    if ((v6 >= (-(v6))) || (!(1 == -1))) {
                                        switch -360 {
                                        case -87:
                                            switch fi57 {
                                            case -9223372036854775808:
                                                v2[(fi57 * 2147483647)] = (1 > v58)
                                            case -357:
                                                v3 = v3
                                            }
                                        case 9223372036854775807:
                                            var v61 string = "\n\t"
                                            _ = v61
                                        case 2147483647:
                                            v1 = append(v1, v4)
                                        default:
                                            fmt.Println(v4, v3)
                                        }
                                        if (4294967295 == (v17 | 662)) {
                                            v62 := GenS2{F0: 32767}
                                            _ = v62
                                        } else {
                                            v2[(-(fi57))] = ((v58 * v58) == (-(v58)))
                                        }
                                    } else {
                                        delete(v2, fi57)
                                    }
                                case "foo bar":
                                    v2[32767] = false
                                }
                                for fi63 := 0; fi63 < 18; fi63++ {
                                    v64 := GenS1{F0: v4, F1: 446}
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = fi63
                                    _ = v64
                                }
                                for fi65 := 0; fi65 < 3; fi65++ {
                                    switch 0 {
                                    case -901:
                                        v66 := GenS2{F0: 42}
                                        _ = v66
                                    }
                                    _ = fi65
                                }
                            }
                            _ = fi57
                            _ = v58
                        }
                    }
                    v67 := GenS1{F0: "foo bar", F1: 337}
                    _ = v67
                } else {
                    var v68 int16 = 42
                    v68 -= 0
                    v68 = v68
                    _ = v68
                }
            }
        }
    case 32767:
        var v69 float64 = (0.0 * 0.1)
        _ = v69
    case -926:
        for fi70 := 0; fi70 < 2; fi70++ {
            switch "a" {
            case "":
                v4 = ("" + (((v4 + v4) + (v4 + "")) + "foo bar"))
            default:
                v17 -= v17
            }
            v2[32768] = ((false || false) && true)
            var v71 uint32 = v17
            _ = fi70
            _ = v71
        }
    default:
        v4 = ("\n\t" + (((v4 + (v4 + (v4 + (("hello" + v4) + ((("\n\t" + v4) + (v4 + v4)) + ((v4 + ("\n\t" + "a")) + ("\n\t" + ((v4 + (("foo bar" + (v4 + "")) + "hello")) + (("foo bar" + v4) + v4))))))))) + v4) + (v4 + v4)))
    }
    v16 = (v16 | (v16 ^ (v16 % ((v16) | 1))))
    v3 = 0.0
    var v72 int = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
}
