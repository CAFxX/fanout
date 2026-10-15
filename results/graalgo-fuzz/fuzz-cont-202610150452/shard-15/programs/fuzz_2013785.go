package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 bool; F1 uint32; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 string; F1 int32; F2 int16 }
func genFn0(p0 uint64, p1 int8) int8 {
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 int16 = 1
        var v6 float64 = float64(((3.14159 / (-2.25 * (3.14159 - 0.0))) / 1.5))
        if (false && ("a" == "foo bar")) {
            for fi7 := 0; fi7 < 18; fi7++ {
                v8 := GenS1{F0: 94, F1: 0}
                _ = fi7
                _ = v8
            }
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    var v9 float32 = (0.0 / float32(0.1))
    if false {
        v10 := GenS2{F0: true, F1: 1, F2: 42}
        _ = v10
    } else {
        p1 -= -128
        v11 := GenS3{F0: (p1 | (p1 - (p1 << ((4294967295) & 63))))}
        if true {
            switch 32767 {
            case -9223372036854775808:
                var v12 uint64 = p0
                _ = v12
            case 896:
                for fi13 := 0; fi13 < 8; fi13++ {
                    fmt.Println(v9, fi13)
                    var v14 uint = 2
                    _ = fi13
                    _ = v14
                }
            }
            var v15 uint64 = (^(p0))
            _ = v15
        } else {
            for fi16 := 0; fi16 < 19; fi16++ {
                v17 := GenS2{F0: ((v9 * v9) >= -2.25), F1: 0, F2: (fi16 + (fi16 / ((fi16) | 1)))}
                switch "\n\t" {
                case "hello":
                    var v18 uint8 = 127
                    _ = v18
                }
                var v19 int = 32768
                _ = fi16
                _ = v17
                _ = v19
            }
            v20 := []int{32767, 2147483647}
            _ = v20
        }
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return p1
}

func main() {
    var v1 int8 = -1
    v2 := make(map[int]uint32)
    var v3 uint64 = 964
    for fi21 := 0; fi21 < 14; fi21++ {
        v3 = v3
        if ((3.14159 == 0.1) || (false || false)) {
            for fi22 := 0; fi22 < 11; fi22++ {
                if false {
                    v23 := GenS2{F0: (uint16(v3) != 65535), F1: 2, F2: (fi22 / ((fi21) | 1))}
                    _ = v23
                } else {
                    v3 = v3
                    if (1 < 65535) {
                        v24 := []int{(fi22 << ((18446744073709551615) & 63)), 127, int(v1), (-(fi22))}
                        var v25 bool = false
                        _ = v24
                        _ = v25
                    }
                }
                v26 := GenS3{F0: v1}
                v27 := []int{fi21}
                _ = fi22
                _ = v26
                _ = v27
            }
            v2[(fi21 & (fi21 | fi21))] = 2147483647
            v3 += (v3 * v3)
        } else {
            v28 := []int{(-(fi21)), -2147483648, (fi21 & (fi21 | fi21)), 32767}
            _ = v28
        }
        switch (^(fi21)) {
        case -689:
            for fi29 := 0; fi29 < 18; fi29++ {
                delete(v2, 4294967295)
                var v30 uint = 492
                _ = fi29
                _ = v30
            }
        default:
            fmt.Println(v1, v3)
        }
        _ = fi21
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        for fi32 := 0; fi32 < 13; fi32++ {
            v33 := GenS4{F0: -607}
            var v34 float32 = (1e10 * -2.25)
            v1 -= v1
            _ = fi32
            _ = v33
            _ = v34
        }
        v2[(fi31 + (fi31 / (((fi31 / ((fi31) | 1))) | 1)))] = 2147483648
        _ = fi31
    }
    var v35 uint = 0
    for fi36 := 0; fi36 < 12; fi36++ {
        var v37 float32 = ((0.0 + 0.1) - (1e10 + 1.5))
        var v38 uint32 = 1
        var v39 bool = true
        _ = fi36
        _ = v37
        _ = v38
        _ = v39
    }
    v3 = (^(v3))
    v40 := []int{32767, 32768, -675, 9223372036854775807}
    switch 42 {
    case 0:
        delete(v2, 32767)
    }
    v1 = (v1 + genFn0(v3, v1))
    if (-593 < -1) {
        fmt.Println(v35, v3)
        v1 = genFn0(v3, genFn0(910, 127))
        if (65535 == 4294967295) {
            v35 -= 42
            var v41 uint16 = 1
            _ = v41
        }
    } else {
        v35 = (v35 << ((18446744073709551615) & 63))
    }
    _ = len(v40)
    v1 = (^(v1))
    for fi42 := 0; fi42 < 16; fi42++ {
        v43 := GenS2{F0: (0 <= v3), F1: 721, F2: (fi42 * (fi42 * 127))}
        for fi44 := 0; fi44 < 18; fi44++ {
            if false {
                fmt.Println(v1, fi44)
            } else {
                v2[fi44] = 424
                delete(v2, fi44)
            }
            _ = fi44
        }
        _ = fi42
        _ = v43
    }
    v3 = (v3 % ((v3) | 1))
    v45 := GenS4{F0: -1}
    if true {
        fmt.Println(v3)
        var v46 float32 = 0.1
        if true {
            v35 = v35
            for fi47 := 0; fi47 < 20; fi47++ {
                for fi48 := 0; fi48 < 7; fi48++ {
                    if false {
                        fmt.Println(v35, v1)
                        if false {
                            var v49 uint32 = 2147483648
                            var v50 int64 = -9223372036854775808
                            var v51 int32 = 42
                            _ = v49
                            _ = v50
                            _ = v51
                        }
                        switch fi47 {
                        case 128:
                            v2[(fi47 ^ (fi48 - fi48))] = 385
                        case 48:
                            v1 = (v1 << (((v35 * (v35 - (v35 / ((v35) | 1))))) & 63))
                            if false {
                                var v52 uint32 = 32
                                v3 = v3
                                if len(v40) >= 2 {
                                    v40 = v40[1:len(v40)]
                                }
                                _ = v52
                            } else {
                                v53 := GenS5{F0: 1, F1: v1}
                                v1 = v1
                                if ((((((("foo bar" + "") + "") + (("a" + ((("hello" + "a") + ("\n\t" + "a")) + ("" + "foo bar"))) + (("\n\t" + (("hello" + "\n\t") + ("\n\t" + ("\n\t" + "a")))) + "foo bar"))) + "hello") + "") <= "foo bar") && (429 >= v35)) {
                                    if len(v40) > 0 {
                                        v40[0] = 913
                                    }
                                } else {
                                    v54 := GenS6{F0: "\n\t", F1: 938, F2: 1}
                                    v55 := GenS6{F0: "a", F1: 2147483647, F2: -1}
                                    var v56 int16 = 32767
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                }
                                _ = v53
                            }
                        case 32768:
                            v3 += (v3 >> ((v35) & 63))
                        default:
                            var v57 uint32 = 98
                            _ = v57
                        }
                    }
                    _ = fi48
                }
                _ = fi47
            }
            _ = len(v2)
        }
        _ = v46
    }
    fmt.Println(v1, v35)
    for fi58 := 0; fi58 < 2; fi58++ {
        for fi59 := 0; fi59 < 6; fi59++ {
            if (-2.25 <= (float32(float64(0.0)) / 1.5)) {
                if len(v40) >= 2 {
                    v40 = v40[1:len(v40)]
                }
            }
            v1 += v1
            switch -1 {
            case -828:
                var v60 int8 = (-(v1))
                _ = v60
            case -932:
                var v61 string = "a"
                _ = v61
            case -956:
                var v62 int8 = v1
                _ = v62
            }
            _ = fi59
        }
        _ = len(v40)
        _ = fi58
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v40
    _ = v45
}
