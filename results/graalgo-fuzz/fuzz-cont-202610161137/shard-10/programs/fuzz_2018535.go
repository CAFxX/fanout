package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool }
func genFn0(p0 uint, p1 uint64) uint {
    v7 := []int{-862, 1, 1, 9223372036854775807}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v8 uint = p0
    _ = len(v7)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (^(v8))
}

func main() {
    v1 := []int{127, -1, 4294967295, 128}
    var v2 uint8 = 127
    var v3 uint16 = 255
    var v4 int = 961
    v5 := make(map[uint32]bool)
    var v6 uint16 = v3
    if false {
        if true {
            if false {
                switch "hello" {
                case "foo bar":
                    if (true || (42 == 4294967295)) {
                        delete(v5, 42)
                    }
                }
            }
        } else {
            _ = len(v1)
        }
    } else {
        var v9 int = v4
        if (!((v3 << ((518) & 63)) != 690)) {
            if false {
                if ((32767 > 42) && (0 <= 705)) {
                    v10 := GenS1{F0: 42}
                    _ = v10
                } else {
                    if ("foo bar" >= ("hello" + ("a" + ""))) {
                        v9 -= v4
                    } else {
                        for fi11 := 0; fi11 < 5; fi11++ {
                            v12 := GenS1{F0: 854}
                            v3 = 2
                            v13 := GenS1{F0: 74}
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        for fi14 := 0; fi14 < 6; fi14++ {
                            v1 = append(v1, fi14)
                            _ = fi14
                        }
                    }
                }
            } else {
                var v15 bool = ((v3 / ((1) | 1)) <= 0)
                _ = v15
            }
            for fi16 := 0; fi16 < 2; fi16++ {
                fmt.Println(v9, v6, fi16)
                v3 = (^(v6))
                _ = fi16
            }
        } else {
            for fi17 := 0; fi17 < 7; fi17++ {
                v2 += v2
                if ((1e10 * 3.14159) > -2.25) {
                    var v18 int64 = -536
                    var v19 int8 = 127
                    if true {
                        _ = len(v5)
                        v3 -= v6
                    } else {
                        if len(v1) > 0 {
                            v1[0] = fi17
                        }
                        v20 := GenS2{F0: 2, F1: (v19 - (v19 + 100))}
                        _ = v20
                    }
                    _ = v18
                    _ = v19
                }
                switch 42 {
                case 9223372036854775807:
                    switch "\n\t" {
                    case "foo bar":
                        delete(v5, 482)
                    case "":
                        var v21 uint = 2
                        _ = v21
                    default:
                        v5[42] = false
                    }
                case -1:
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v23 := GenS1{F0: 872}
                        var v24 float64 = (float64(((0.1 - float32(((float64(0.1) * 1e10) - (0.0 / 1e10)))) - -2.25)) - (1.5 * -2.25))
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                default:
                    for fi25 := 0; fi25 < 13; fi25++ {
                        fmt.Println(v4, fi17, v3)
                        _ = fi25
                    }
                }
                _ = fi17
            }
        }
        _ = v9
    }
    fmt.Println(v3, v2)
    var v26 int64 = -9223372036854775808
    if true {
        for fi27 := 0; fi27 < 20; fi27++ {
            fmt.Println(v26, v6)
            v28 := GenS3{F0: true, F1: 4294967295}
            _ = fi27
            _ = v28
        }
        var v29 float32 = (0.0 * (((0.0 - ((1e10 * 0.1) * 0.1)) / float32((0.1 + -2.25))) + 0.1))
        if false {
            for fi30 := 0; fi30 < 1; fi30++ {
                v2 -= 254
                _ = fi30
            }
            v31 := GenS4{F0: (v6 | (v3 % ((v6) | 1)))}
            _ = v31
        } else {
            switch "foo bar" {
            case "\n\t":
                v32 := GenS5{F0: 1}
                _ = v32
            case "hello":
                fmt.Println(v26, v2, v4)
            case "a":
                v33 := GenS4{F0: v6}
                _ = v33
            }
            _ = len(v5)
        }
        _ = v29
    } else {
        var v34 uint16 = (v6 + 2)
        _ = len(v5)
        if ("foo bar" != "a") {
            switch (v4 << ((1) & 63)) {
            case -273:
                if (!false) {
                    switch "\n\t" {
                    case "\n\t":
                        if (true && true) {
                            var v35 int64 = int64(v3)
                            for fi36 := 0; fi36 < 2; fi36++ {
                                v34 = (v6 | (v6 >> ((genFn0(uint(v6), 1)) & 63)))
                                _ = fi36
                            }
                            _ = v35
                        }
                    case "hello":
                        for fi37 := 0; fi37 < 5; fi37++ {
                            v38 := GenS2{F0: 43, F1: 127}
                            _ = fi37
                            _ = v38
                        }
                    }
                    for fi39 := 0; fi39 < 4; fi39++ {
                        v40 := v5[uint32(v2)]
                        if (!false) {
                            switch ("foo bar" + "a") {
                            case "\n\t":
                                switch fi39 {
                                case -1:
                                    var v41 int16 = 42
                                    _ = v41
                                case 0:
                                    v42 := v5[65535]
                                    _ = v42
                                case -2147483648:
                                    v34 -= v34
                                default:
                                    if (v40 || v40) {
                                        if len(v1) > 0 {
                                            v1[0] = fi39
                                        }
                                        fmt.Println(v3, v26)
                                    }
                                }
                            case "":
                                for fi43 := 0; fi43 < 19; fi43++ {
                                    v44 := GenS5{F0: -2147483648}
                                    _ = fi43
                                    _ = v44
                                }
                            case "hello":
                                _ = len(v1)
                            }
                            if (!v40) {
                                var v45 int8 = 100
                                if true {
                                    var v46 float64 = ((1e10 / 1e10) + (1.5 + 0.1))
                                    _ = v46
                                } else {
                                    for fi47 := 0; fi47 < 14; fi47++ {
                                        v6 -= (v6 / ((649) | 1))
                                        var v48 int16 = -830
                                        v49 := GenS5{F0: 1}
                                        _ = fi47
                                        _ = v48
                                        _ = v49
                                    }
                                }
                                _ = v45
                            }
                            v40 = (506 > -1)
                        }
                        _ = fi39
                        _ = v40
                    }
                    if true {
                        _ = len(v1)
                    }
                }
            case 32768:
                if true {
                    for fi50 := 0; fi50 < 18; fi50++ {
                        var v51 uint32 = 2147483648
                        if true {
                            v2 = 127
                            fmt.Println(v34, v4, v4)
                        } else {
                            v51 = (^(v51))
                        }
                        _ = fi50
                        _ = v51
                    }
                }
            }
            v26 = (^(v26))
            switch "" {
            case "":
                var v52 float32 = 0.0
                _ = v52
            case "hello":
                if false {
                    var v53 uint8 = (v2 - 127)
                    var v54 float64 = -2.25
                    for fi55 := 0; fi55 < 2; fi55++ {
                        for fi56 := 0; fi56 < 17; fi56++ {
                            v54 = float64(3.14159)
                            var v57 int16 = 32767
                            _ = fi56
                            _ = v57
                        }
                        v2 -= v2
                        v4 -= (v4 % ((fi55) | 1))
                        _ = fi55
                    }
                    _ = v53
                    _ = v54
                }
            }
        } else {
            v2 = v2
            for fi58 := 0; fi58 < 19; fi58++ {
                var v59 int32 = 2147483647
                _ = fi58
                _ = v59
            }
            for fi60 := 0; fi60 < 11; fi60++ {
                var v61 uint16 = (v6 % (((v6 - 0)) | 1))
                var v62 int32 = -997
                for fi63 := 0; fi63 < 2; fi63++ {
                    v64 := GenS6{F0: (v62 < v62)}
                    var v65 bool = (792 == 629)
                    v26 = (v26 * (-(v26)))
                    _ = fi63
                    _ = v64
                    _ = v65
                }
                _ = fi60
                _ = v61
                _ = v62
            }
        }
        _ = v34
    }
    v3 = v6
    var v66 float64 = 1.5
    switch 127 {
    case -9:
        fmt.Println(v2, v6)
    case -2147483648:
        var v67 int16 = 32767
        _ = v67
    case 32768:
        var v68 float64 = ((v66 + v66) - (v66 * 1e10))
        _ = v68
    }
    var v69 int32 = 42
    for fi70 := 0; fi70 < 19; fi70++ {
        v3 += v6
        _ = fi70
    }
    v66 = v66
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v69
}
