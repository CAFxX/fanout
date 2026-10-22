package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint; F1 int32; F2 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 int16; F2 int16 }
type GenS5 struct { F0 int16 }
func genFn0(p0 bool, p1 uint8, p2 uint8) uint {
    for fi7 := 0; fi7 < 10; fi7++ {
        switch fi7 {
        case 42:
            for fi8 := 0; fi8 < 10; fi8++ {
                v9 := []int{}
                v10 := GenS1{F0: 172, F1: 0}
                v11 := GenS1{F0: -1, F1: 231}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        case 32767:
            v12 := []int{}
            _ = v12
        default:
            var v13 string = "\n\t"
            _ = v13
        }
        p1 = 102
        p0 = p0
        _ = fi7
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        p0 = true
        var v15 int16 = 1
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 uint64) int16 {
    v16 := []int{4294967295, -1, 32768, 1}
    if false {
        var v17 int16 = 1
        if len(v16) > 0 {
            v16[0] = -536
        }
        _ = v17
    }
    var v18 int64 = 1
    for fi19 := 0; fi19 < 6; fi19++ {
        v20 := GenS2{F0: genFn0(true, 255, 2), F1: -1, F2: 0}
        p0 = 18446744073709551615
        _ = fi19
        _ = v20
    }
    v21 := GenS3{F0: 100}
    _ = p0
    _ = v16
    _ = v18
    _ = v21
    return 1
}

func main() {
    var v1 int = 127
    var v2 string = "a"
    var v3 int8 = 1
    var v4 int32 = 0
    var v5 float32 = 0.0
    var v6 uint16 = 32767
    v22 := []int{v1, (v1 >> ((319) & 63))}
    if (1.5 < 0.0) {
        v3 = v3
        if true {
            fmt.Println(v6)
            v4 += (v4 << ((255) & 63))
            if len(v22) > 0 {
                v22[0] = v1
            }
        }
        if len(v22) > 0 {
            v22[0] = v1
        }
    }
    switch ("" + "\n\t") {
    case "\n\t":
        for fi23 := 0; fi23 < 12; fi23++ {
            for fi24 := 0; fi24 < 5; fi24++ {
                fmt.Println(v3, v1, v4)
                var v25 int = 530
                v3 = -80
                _ = fi24
                _ = v25
            }
            if false {
                switch 127 {
                case 9223372036854775807:
                    switch 9223372036854775807 {
                    case -1:
                        for fi26 := 0; fi26 < 16; fi26++ {
                            var v27 int8 = 127
                            var v28 bool = false
                            v1 = v1
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                    case 32767:
                        for fi29 := 0; fi29 < 20; fi29++ {
                            if false {
                                fmt.Println(v3)
                                var v30 int64 = -1
                                _ = v30
                            }
                            if true {
                                var v31 int64 = -1
                                _ = len(v22)
                                _ = v31
                            } else {
                                v32 := GenS2{F0: 0, F1: v4, F2: 255}
                                var v33 int64 = -399
                                switch v2 {
                                case "foo bar":
                                    fmt.Println(fi29)
                                case "":
                                    var v34 int8 = -1
                                    _ = v34
                                default:
                                    switch "hello" {
                                    case "\n\t":
                                        var v35 float64 = float64(1.5)
                                        _ = v35
                                    default:
                                        v3 = v3
                                    }
                                }
                                _ = v32
                                _ = v33
                            }
                            var v36 int16 = genFn1(42)
                            _ = fi29
                            _ = v36
                        }
                    case -612:
                        var v37 int = (v1 % ((2147483647) | 1))
                        _ = v37
                    }
                case -243:
                    var v38 uint64 = 0
                    _ = v38
                }
                for fi39 := 0; fi39 < 20; fi39++ {
                    switch fi23 {
                    case -1:
                        var v40 string = (v2 + (v2 + ("a" + v2)))
                        v4 = 295
                        _ = v40
                    case 1:
                        v41 := GenS2{F0: 18446744073709551615, F1: (v4 << ((65535) & 63)), F2: 2}
                        _ = v41
                    }
                    _ = fi39
                }
            }
            _ = fi23
        }
    case "":
        if ((!false) && true) {
            var v42 uint16 = 73
            switch ((v2 + "\n\t") + ("a" + "")) {
            case "foo bar":
                fmt.Println(v2, v6, v4)
            case "\n\t":
                for fi43 := 0; fi43 < 10; fi43++ {
                    var v44 string = ("hello" + "foo bar")
                    switch (-(v1)) {
                    case 2147483647:
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                    case -238:
                        v45 := GenS2{F0: 857, F1: 532, F2: uint8(v42)}
                        _ = v45
                    default:
                        var v46 float64 = float64(-2.25)
                        _ = v46
                    }
                    _ = fi43
                    _ = v44
                }
            }
            for fi47 := 0; fi47 < 4; fi47++ {
                for fi48 := 0; fi48 < 3; fi48++ {
                    v42 -= (v42 >> ((255) & 63))
                    _ = fi48
                }
                var v49 int32 = v4
                _ = fi47
                _ = v49
            }
            _ = v42
        }
    case "foo bar":
        var v50 int64 = -9223372036854775808
        _ = v50
    }
    v1 = -2147483648
    if false {
        v2 = v2
    } else {
        v51 := GenS4{F0: 65535, F1: genFn1(42), F2: -418}
        switch "foo bar" {
        case "a":
            var v52 uint16 = v6
            _ = v52
        case "foo bar":
            v5 = 3.14159
        }
        _ = v51
    }
    var v53 int8 = v3
    switch (v1 + 9223372036854775807) {
    case 307:
        var v54 bool = (!true)
        _ = v54
    case 0:
        v55 := GenS5{F0: 1000}
        _ = v55
    default:
        for fi56 := 0; fi56 < 13; fi56++ {
            var v57 uint32 = 1
            if true {
                v22 = append(v22, (^(fi56)))
                var v58 bool = ((v53 >> ((uint(v6)) & 63)) != (^(v3)))
                _ = v58
            } else {
                v5 -= (v5 + v5)
                v57 += v57
            }
            v59 := GenS2{F0: 167, F1: (^(v4)), F2: 128}
            _ = fi56
            _ = v57
            _ = v59
        }
    }
    _ = len(v22)
    v3 += (v3 + (v3 & (^(v3))))
    if (((float32(float64((v5 - v5))) - 1.5) <= (-2.25 * 1.5)) || ((42 >= 855) && true)) {
        v60 := GenS1{F0: (^(v4)), F1: v6}
        for fi61 := 0; fi61 < 16; fi61++ {
            v6 -= (v6 >> ((927) & 63))
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            _ = fi61
        }
        _ = v60
    } else {
        if true {
            switch v2 {
            case "":
                if (!(false || true)) {
                    switch v1 {
                    case 459:
                        v3 -= 40
                    case 1:
                        if false {
                            v4 = v4
                            for fi62 := 0; fi62 < 17; fi62++ {
                                v63 := GenS5{F0: int16(v3)}
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                                _ = fi62
                                _ = v63
                            }
                            if len(v22) > 0 {
                                v22[0] = 614
                            }
                        }
                    default:
                        var v64 uint16 = (v6 % ((0) | 1))
                        _ = v64
                    }
                } else {
                    v5 = v5
                    v4 -= (v4 & (-(v4)))
                    var v65 uint16 = v6
                    _ = v65
                }
            case "\n\t":
                var v66 uint8 = 0
                _ = v66
            case "a":
                for fi67 := 0; fi67 < 1; fi67++ {
                    for fi68 := 0; fi68 < 12; fi68++ {
                        var v69 uint = 1
                        _ = fi68
                        _ = v69
                    }
                    var v70 uint64 = 1
                    var v71 bool = false
                    _ = fi67
                    _ = v70
                    _ = v71
                }
            }
            var v72 bool = (393 != v6)
            v3 = (v3 / ((-1) | 1))
            _ = v72
        } else {
            v73 := GenS4{F0: 709, F1: 320, F2: 0}
            _ = v73
        }
        v6 = (v6 + (v6 >> ((302) & 63)))
        v4 += v4
    }
    fmt.Println(v2)
    if len(v22) > 0 {
        v22[0] = (v1 ^ (v1 << ((42) & 63)))
    }
    v74 := GenS2{F0: 2, F1: (-(v4)), F2: 2}
    v3 += (v3 / ((0) | 1))
    v53 = (-(v3))
    _ = len(v22)
    v2 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v74
}
