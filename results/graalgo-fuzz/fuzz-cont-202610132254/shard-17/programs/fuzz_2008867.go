package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int16 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS8 struct { F0 int32; F1 bool; F2 int32 }
type GenS9 struct { F0 uint64; F1 int }
type GenS10 struct { F0 bool }
func genFn0() int32 {
    v5 := GenS1{F0: 44, F1: -128, F2: -4}
    var v6 int32 = -303
    for fi7 := 0; fi7 < 5; fi7++ {
        v6 -= 42
        fmt.Println(fi7, v6)
        v8 := []int{fi7, fi7, (fi7 * (fi7 & (fi7 | (fi7 ^ (fi7 | (fi7 * 32767))))))}
        _ = fi7
        _ = v8
    }
    v6 += 0
    if false {
        if true {
            if (18446744073709551615 < 562) {
                v6 += (v6 ^ (-(v6)))
                fmt.Println(v6)
            } else {
                v9 := GenS2{F0: -402, F1: -507}
                for fi10 := 0; fi10 < 7; fi10++ {
                    v11 := GenS2{F0: 581, F1: (-(v6))}
                    v6 -= v6
                    v12 := []int{32768}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = v9
            }
            v6 -= v6
        } else {
            v13 := []int{689}
            var v14 int8 = 75
            _ = v13
            _ = v14
        }
        v6 += (v6 >> ((65535) & 63))
    }
    _ = v5
    _ = v6
    return v6
}
func genFn1(p0 string) uint16 {
    var v15 uint64 = 2
    if (true && ("" <= (p0 + p0))) {
        p0 = ("a" + "a")
        v16 := GenS3{F0: 4294967295, F1: 972}
        _ = v16
    }
    var v17 uint8 = 128
    v18 := []int{}
    _ = p0
    _ = v15
    _ = v17
    _ = v18
    return 0
}

func main() {
    var v1 uint8 = 42
    var v2 int = -2147483648
    var v3 int8 = 1
    v4 := []string{"\n\t"}
    _ = len(v4)
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(v2, fi19)
        v20 := GenS3{F0: 65535, F1: 886}
        _ = fi19
        _ = v20
    }
    v1 = 2
    v2 += v2
    fmt.Println(v3, v2, v1)
    var v21 uint = 1
    if ("a" <= "foo bar") {
        for fi22 := 0; fi22 < 18; fi22++ {
            v1 = (^(v1))
            if (255 >= v21) {
                var v23 float32 = 3.14159
                var v24 float32 = v23
                _ = v23
                _ = v24
            } else {
                for fi25 := 0; fi25 < 4; fi25++ {
                    var v26 bool = true
                    _ = fi25
                    _ = v26
                }
                var v27 float32 = 3.14159
                var v28 int32 = 0
                _ = v27
                _ = v28
            }
            fmt.Println(fi22, v2)
            _ = fi22
        }
        fmt.Println(v1)
        fmt.Println(v21)
    } else {
        v29 := GenS3{F0: (v21 * (^(v21))), F1: 1}
        if false {
            v3 -= (-(v3))
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            switch ("foo bar" + ("foo bar" + (("foo bar" + "hello") + "foo bar"))) {
            case "hello":
                for fi30 := 0; fi30 < 9; fi30++ {
                    v1 += (v1 * (v1 - (^(v1))))
                    var v31 uint32 = 2
                    _ = fi30
                    _ = v31
                }
            case "a":
                _ = len(v4)
            case "foo bar":
                switch v2 {
                case 42:
                    if ("hello" > "") {
                        var v32 int16 = 259
                        if false {
                            v21 -= 1
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                        } else {
                            if false {
                                v2 += (v2 & 0)
                            }
                            if false {
                                if true {
                                    v32 -= -838
                                    if false {
                                        for fi33 := 0; fi33 < 10; fi33++ {
                                            if ((v3 << ((18446744073709551615) & 63)) < (v3 / ((v3) | 1))) {
                                                v34 := GenS4{F0: (v21 >> (((v21 / ((v21) | 1))) & 63)), F1: ((v1 - v1) < (^(v1)))}
                                                v35 := GenS4{F0: (v21 * 33), F1: (v3 != v3)}
                                                _ = v34
                                                _ = v35
                                            }
                                            if (2 < 482) {
                                                if len(v4) > 0 {
                                                    v4[0] = "a"
                                                }
                                            } else {
                                                for fi36 := 0; fi36 < 13; fi36++ {
                                                    var v37 uint32 = uint32(v1)
                                                    var v38 int8 = (v3 | (v3 & v3))
                                                    _ = fi36
                                                    _ = v37
                                                    _ = v38
                                                }
                                                var v39 int32 = genFn0()
                                                _ = v39
                                            }
                                            _ = fi33
                                        }
                                    } else {
                                        var v40 int8 = (v3 + v3)
                                        if (!false) {
                                            var v41 int16 = v32
                                            _ = v41
                                        }
                                        _ = v40
                                    }
                                    var v42 uint32 = 2147483648
                                    _ = v42
                                } else {
                                    if false {
                                        var v43 int32 = -2147483648
                                        _ = v43
                                    }
                                    var v44 string = "foo bar"
                                    _ = v44
                                }
                                var v45 uint16 = 32767
                                if true {
                                    var v46 uint = (v21 % ((uint(v45)) | 1))
                                    switch (v2 / (((v2 << (((v21 ^ v21)) & 63))) | 1)) {
                                    case 32767:
                                        if (296 >= 4294967295) {
                                            var v47 int32 = 148
                                            _ = v47
                                        } else {
                                            v45 = 65535
                                            v48 := GenS5{F0: v32, F1: "\n\t"}
                                            v2 += 2147483647
                                            _ = v48
                                        }
                                    default:
                                        v46 += v46
                                    }
                                    _ = v46
                                }
                                _ = v45
                            } else {
                                v49 := GenS4{F0: v21, F1: (((float64(((1.5 + 1e10) + 1e10)) * 0.1) >= 0.0) && true)}
                                var v50 uint32 = 65535
                                _ = v49
                                _ = v50
                            }
                            for fi51 := 0; fi51 < 2; fi51++ {
                                v32 = (v32 + v32)
                                _ = fi51
                            }
                        }
                        switch (^(v2)) {
                        case 9223372036854775807:
                            fmt.Println(v21, v3, v2)
                        case 42:
                            if false {
                                switch -965 {
                                case -2147483648:
                                    switch v2 {
                                    case 420:
                                        _ = len(v4)
                                    case 2147483647:
                                        if false {
                                            fmt.Println(v2, v21)
                                            if false {
                                                switch v2 {
                                                case 127:
                                                    for fi52 := 0; fi52 < 17; fi52++ {
                                                        v21 = v21
                                                        if true {
                                                            if (false && false) {
                                                                fmt.Println(v21, v32)
                                                                fmt.Println(v32, v1)
                                                                v32 += -1
                                                            } else {
                                                                v53 := GenS2{F0: 0, F1: -2147483648}
                                                                _ = v53
                                                            }
                                                            var v54 uint64 = 1
                                                            switch "\n\t" {
                                                            case "a":
                                                                v54 = 238
                                                            default:
                                                                v32 = 0
                                                            }
                                                            _ = v54
                                                        }
                                                        switch v2 {
                                                        case 726:
                                                            var v55 float64 = (1.5 - 1.5)
                                                            _ = v55
                                                        default:
                                                            for fi56 := 0; fi56 < 13; fi56++ {
                                                                if false {
                                                                    fmt.Println(v3, fi52)
                                                                } else {
                                                                    v2 += (v2 | (^(fi52)))
                                                                    v21 = 290
                                                                    v57 := GenS6{F0: (fi52 % ((fi52) | 1)), F1: (v2 << (((v21 * v21)) & 63))}
                                                                    _ = v57
                                                                }
                                                                var v58 float32 = ((0.1 - (((-2.25 + -2.25) / (((1.5 * (((((0.1 * 1.5) * 3.14159) - 0.0) - ((((1.5 - ((1.5 + 1.5) / 0.1)) - ((3.14159 - 1e10) * (3.14159 / -2.25))) + 0.1) * float32(-2.25))) + 0.1)) * float32((1e10 * -2.25))) * 1e10)) - 3.14159)) + -2.25)
                                                                _ = fi56
                                                                _ = v58
                                                            }
                                                        }
                                                        _ = fi52
                                                    }
                                                case 635:
                                                    if (((-2.25 * 0.1) > float32(0.0)) && true) {
                                                        v21 = (^(v21))
                                                        v1 = 0
                                                        for fi59 := 0; fi59 < 16; fi59++ {
                                                            var v60 uint = v21
                                                            _ = fi59
                                                            _ = v60
                                                        }
                                                    } else {
                                                        v4 = append(v4, ("hello" + ("" + (("foo bar" + "a") + "hello"))))
                                                        fmt.Println(v32)
                                                    }
                                                case 169:
                                                    var v61 uint8 = v1
                                                    _ = v61
                                                default:
                                                    var v62 string = (("a" + "") + ("hello" + "\n\t"))
                                                    _ = v62
                                                }
                                            } else {
                                                var v63 uint32 = 42
                                                _ = v63
                                            }
                                        } else {
                                            for fi64 := 0; fi64 < 7; fi64++ {
                                                var v65 uint = (v21 / ((310) | 1))
                                                var v66 uint8 = v1
                                                _ = fi64
                                                _ = v65
                                                _ = v66
                                            }
                                            fmt.Println(v21, v1)
                                            v67 := GenS7{F0: 65535, F1: 2, F2: v3}
                                            _ = v67
                                        }
                                    }
                                case 926:
                                    v32 -= -258
                                default:
                                    v1 = 42
                                }
                            }
                        case 32767:
                            _ = len(v4)
                        default:
                            switch "foo bar" {
                            case "a":
                                _ = len(v4)
                            }
                        }
                        _ = v32
                    } else {
                        fmt.Println(v3, v1)
                        var v68 int64 = 17
                        v21 += v21
                        _ = v68
                    }
                case 25:
                    v69 := GenS8{F0: int32(v3), F1: false, F2: 1}
                    _ = v69
                case -826:
                    switch v2 {
                    case -9223372036854775808:
                        var v70 int = v2
                        _ = v70
                    default:
                        if true {
                            var v71 string = ""
                            _ = v71
                        } else {
                            for fi72 := 0; fi72 < 9; fi72++ {
                                v73 := GenS9{F0: 1, F1: v2}
                                v21 = v21
                                _ = fi72
                                _ = v73
                            }
                            var v74 int8 = 100
                            var v75 string = "hello"
                            _ = v74
                            _ = v75
                        }
                    }
                }
            }
        }
        v1 = 1
        _ = v29
    }
    for fi76 := 0; fi76 < 13; fi76++ {
        if false {
            v77 := GenS10{F0: (v1 >= (v1 / (((v1 << ((uint(v1)) & 63))) | 1)))}
            _ = len(v4)
            var v78 uint32 = 1
            _ = v77
            _ = v78
        } else {
            var v79 int32 = genFn0()
            var v80 bool = (v21 <= v21)
            _ = v79
            _ = v80
        }
        fmt.Println(v21, fi76, v21)
        v81 := GenS9{F0: 18446744073709551615, F1: fi76}
        _ = fi76
        _ = v81
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
