package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS3 struct { F0 string; F1 uint8 }
func genFn0(p0 int, p1 bool, p2 uint8) int64 {
    if false {
        var v7 int32 = -1
        switch "" {
        case "":
            var v8 uint8 = p2
            _ = v8
        case "foo bar":
            p0 += 404
        case "\n\t":
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 uint16 = 14
                _ = fi9
                _ = v10
            }
        }
        _ = v7
    }
    v11 := []int{(p0 - (p0 % ((int(p2)) | 1)))}
    if p1 {
        p0 = (p0 | (p0 >> ((18446744073709551615) & 63)))
        fmt.Println(p0)
        var v12 int8 = -65
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return -98
}
func genFn1() int8 {
    v13 := GenS1{F0: 120, F1: 86, F2: 65535}
    v14 := []int{-799, 127, 128}
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 int = 32768
            v17 = (-(fi16))
            var v18 uint = 270
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    _ = v13
    _ = v14
    return 127
}

func main() {
    v1 := []uint32{278, 42}
    v2 := []string{"\n\t", "", ""}
    var v3 float64 = 1e10
    var v4 uint8 = 199
    var v5 int32 = 42
    v6 := []int{128, 2147483647, 638, 642}
    v1 = append(v1, 337)
    for fi19 := 0; fi19 < 13; fi19++ {
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        switch fi19 {
        case -762:
            if (2147483647 <= 42) {
                var v20 uint64 = 1
                _ = len(v2)
                _ = v20
            }
        case 206:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -9223372036854775808:
            if len(v6) > 0 {
                v6[0] = 9223372036854775807
            }
        }
        for fi21 := 0; fi21 < 13; fi21++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi21
        }
        _ = fi19
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v4)
    case "\n\t":
        var v22 int16 = 32767
        _ = v22
    case "":
        switch 0 {
        case 128:
            if (genFn1() <= -60) {
                if len(v2) > 0 {
                    v2[0] = ("hello" + "\n\t")
                }
                v3 += v3
            } else {
                v2 = append(v2, "\n\t")
            }
        case 127:
            if ((!false) || false) {
                if ((!false) && (true && ((1000 != 32767) || true))) {
                    for fi23 := 0; fi23 < 11; fi23++ {
                        v3 = (v3 + ((v3 - 0.0) + v3))
                        _ = fi23
                    }
                    var v24 uint64 = 976
                    var v25 string = (("hello" + ("\n\t" + "hello")) + "")
                    _ = v24
                    _ = v25
                } else {
                    v3 = (v3 - v3)
                    var v26 int8 = 0
                    for fi27 := 0; fi27 < 17; fi27++ {
                        var v28 uint32 = 1
                        v26 = v26
                        _ = fi27
                        _ = v28
                    }
                    _ = v26
                }
                for fi29 := 0; fi29 < 2; fi29++ {
                    fmt.Println(fi29)
                    if true {
                        _ = len(v6)
                        var v30 uint64 = 42
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 5; fi31++ {
                        var v32 float64 = v3
                        _ = fi31
                        _ = v32
                    }
                    _ = fi29
                }
                if ((v3 != (v3 * -2.25)) || (false && (749 != 916))) {
                    fmt.Println(v5, v4)
                    if true {
                        var v33 float64 = v3
                        switch 32768 {
                        case -1:
                            v3 -= (0.1 + (float64(0.0) * (v3 * (0.1 / 1.0))))
                        case 131:
                            if (!true) {
                                var v34 uint32 = 0
                                _ = v34
                            } else {
                                var v35 float32 = 1e10
                                if false {
                                    fmt.Println(v33, v35, v3)
                                    for fi36 := 0; fi36 < 3; fi36++ {
                                        switch 2147483647 {
                                        case 1:
                                            if len(v2) > 0 {
                                                v2[0] = "a"
                                            }
                                        case 32768:
                                            _ = len(v1)
                                        case -671:
                                            if false {
                                                v33 = (v3 + v33)
                                                v3 = v33
                                                v37 := GenS2{F0: true, F1: v4, F2: 42}
                                                _ = v37
                                            }
                                        default:
                                            v33 += (0.0 * ((v3 / ((1e10 - v3) * float64(((v35 + (((v35 + (v35 * 1e10)) + 0.1) * v35)) / 1e10)))) - (3.14159 / -2.25)))
                                        }
                                        switch "" {
                                        case "a":
                                            var v38 int64 = -1
                                            _ = v38
                                        default:
                                            for fi39 := 0; fi39 < 14; fi39++ {
                                                v40 := GenS1{F0: 4294967295, F1: (v4 >> ((65535) & 63)), F2: 1}
                                                _ = fi39
                                                _ = v40
                                            }
                                        }
                                        _ = fi36
                                    }
                                    v41 := GenS1{F0: 765, F1: (v4 - 128), F2: 825}
                                    _ = v41
                                }
                                switch ("foo bar" + ("a" + "a")) {
                                case "foo bar":
                                    v42 := GenS1{F0: 2147483647, F1: v4, F2: 787}
                                    _ = v42
                                }
                                _ = v35
                            }
                        }
                        _ = v33
                    } else {
                        v43 := GenS2{F0: true, F1: v4, F2: 352}
                        for fi44 := 0; fi44 < 5; fi44++ {
                            v45 := GenS1{F0: 1, F1: 1, F2: 65535}
                            v5 -= (^(v5))
                            _ = fi44
                            _ = v45
                        }
                        _ = v43
                    }
                }
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        case 0:
            if len(v6) > 0 {
                v6[0] = 795
            }
        default:
            fmt.Println(v4)
        }
    default:
        for fi46 := 0; fi46 < 2; fi46++ {
            switch fi46 {
            case 42:
                var v47 string = "hello"
                _ = v47
            case -482:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 127:
                if true {
                    v48 := GenS2{F0: true, F1: (v4 % ((v4) | 1)), F2: 2147483647}
                    _ = v48
                }
            }
            if (0.1 > v3) {
                v4 -= (v4 + v4)
                switch "" {
                case "foo bar":
                    if false {
                        var v49 float64 = ((((-2.25 - v3) - (0.1 * (v3 / 1.0))) / 0.1) + (0.1 + v3))
                        _ = v49
                    }
                case "hello":
                    v50 := GenS3{F0: "foo bar", F1: 0}
                    _ = v50
                default:
                    if (255 >= 2) {
                        fmt.Println(v3)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if true {
                            var v51 uint8 = (v4 + (v4 / ((42) | 1)))
                            v3 = (v3 - v3)
                            v3 = v3
                            _ = v51
                        }
                    }
                }
                v5 = (v5 | (v5 ^ v5))
            } else {
                var v52 uint32 = 2147483648
                _ = v52
            }
            for fi53 := 0; fi53 < 5; fi53++ {
                v4 = 255
                v5 -= v5
                _ = fi53
            }
            _ = fi46
        }
    }
    v4 = (v4 * v4)
    switch "hello" {
    case "hello":
        var v54 uint = 0
        _ = v54
    case "":
        _ = len(v2)
    }
    v55 := GenS2{F0: (!true), F1: v4, F2: 42}
    v56 := GenS3{F0: "\n\t", F1: v4}
    fmt.Println(v5, v4, v3)
    var v57 uint32 = 1
    var v58 bool = false
    for fi59 := 0; fi59 < 16; fi59++ {
        switch "\n\t" {
        case "":
            v5 -= (^(v5))
        }
        v4 = (v4 | v4)
        _ = fi59
    }
    var v60 uint8 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v6
    _ = v60
}
