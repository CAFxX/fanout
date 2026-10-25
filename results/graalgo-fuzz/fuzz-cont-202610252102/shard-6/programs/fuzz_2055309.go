package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint16; F2 int32 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 int8; F2 int64 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint; F1 bool; F2 int8 }
type GenS9 struct { F0 bool; F1 uint }

func main() {
    var v1 float64 = 1.5
    var v2 float64 = 1.5
    var v3 int16 = 0
    v4 := []int{32768, 9223372036854775807}
    v5 := make(map[uint32]string)
    var v6 string = "hello"
    v2 += 1e10
    if (2 < 18446744073709551615) {
        var v7 string = v6
        _ = v7
    } else {
        fmt.Println(v1)
    }
    v2 -= (v1 + 3.14159)
    v8 := GenS1{F0: 65535}
    for fi9 := 0; fi9 < 17; fi9++ {
        switch v6 {
        case "\n\t":
            switch v6 {
            case "hello":
                for fi10 := 0; fi10 < 8; fi10++ {
                    var v11 string = v6
                    fmt.Println(v1, fi9, v2)
                    switch (-(fi10)) {
                    case 127:
                        if false {
                            v2 = v2
                            v12 := GenS1{F0: 4294967295}
                            v13 := GenS2{F0: ((!false) || (!(142 < 941))), F1: 32767, F2: -2147483648}
                            v2 = 1e10
                            _ = v12
                            _ = v13
                        } else {
                            v3 = (v3 & v3)
                            var v14 float64 = (3.14159 - float64(-2.25))
                            if (948 < 18446744073709551615) {
                                v15 := GenS1{F0: 255}
                                _ = v15
                            }
                            _ = v14
                        }
                    }
                    _ = fi10
                    _ = v11
                }
            case "\n\t":
                v16 := GenS2{F0: true, F1: 1, F2: -28}
                _ = v16
            }
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi9
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            if len(v4) > 0 {
                v4[0] = fi17
            }
            if false {
                v19 := GenS2{F0: (2147483648 == 719), F1: 417, F2: 0}
                var v20 uint32 = 1
                _ = v19
                _ = v20
            }
            v2 += (3.14159 + 3.14159)
            _ = fi18
        }
        _ = len(v5)
        _ = fi17
    }
    if true {
        v6 = v6
        for fi21 := 0; fi21 < 18; fi21++ {
            v3 += v3
            for fi22 := 0; fi22 < 16; fi22++ {
                if (!false) {
                    switch v6 {
                    case "foo bar":
                        if false {
                            if (3.14159 == 1e10) {
                                v2 = v1
                                v23 := GenS2{F0: false, F1: 42, F2: -353}
                                var v24 uint32 = 1
                                v25 := GenS3{F0: fi22, F1: (((-9223372036854775808 != -9223372036854775808) && true) || (fi21 == (fi22 ^ fi21)))}
                                switch fi22 {
                                case 32767:
                                    if false {
                                        v26 := GenS3{F0: fi21, F1: (!false)}
                                        var v27 uint32 = (v24 | 2147483647)
                                        v28 := GenS1{F0: 4294967295}
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    } else {
                                        var v29 bool = (v24 <= (^(v24)))
                                        var v30 uint8 = 128
                                        if ((false && v29) && (((v3 % ((-32768) | 1)) >= (v3 % ((0) | 1))) || (-2.25 >= (float32((v1 / (0.1 * v1))) * 0.1)))) {
                                            v1 -= (3.14159 + float64((3.14159 - 0.1)))
                                            switch (fi22 >> ((uint(v24)) & 63)) {
                                            case 32768:
                                                switch ((v6 + (v6 + "\n\t")) + v6) {
                                                case "":
                                                    switch (v6 + (v6 + "")) {
                                                    case "foo bar":
                                                        if false {
                                                            var v31 uint64 = 1
                                                            if (2 > 255) {
                                                                v29 = (v3 < (v3 - v3))
                                                                v6 = v6
                                                                var v32 uint16 = 32767
                                                                _ = v32
                                                            }
                                                            v30 += v30
                                                            _ = v31
                                                        }
                                                    case "hello":
                                                        if false {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                    default:
                                                        switch fi21 {
                                                        case 32767:
                                                            var v33 uint16 = 702
                                                            _ = v33
                                                        }
                                                    }
                                                case "a":
                                                    var v34 uint64 = 1
                                                    _ = v34
                                                case "hello":
                                                    v24 += (v24 * 2)
                                                default:
                                                    v1 -= float64(1e10)
                                                }
                                            case 734:
                                                v29 = false
                                            case -9223372036854775808:
                                                var v35 uint32 = 4294967295
                                                _ = v35
                                            default:
                                                v36 := GenS4{F0: v6}
                                                _ = v36
                                            }
                                            var v37 bool = (!(4294967295 > uint(v30)))
                                            var v38 int8 = -40
                                            _ = v37
                                            _ = v38
                                        }
                                        _ = v29
                                        _ = v30
                                    }
                                case 1:
                                    if false {
                                        v39 := GenS5{F0: "hello", F1: 127, F2: -1}
                                        var v40 uint8 = 0
                                        _ = v39
                                        _ = v40
                                    } else {
                                        v41 := GenS5{F0: v6, F1: 42, F2: -1}
                                        var v42 int64 = 652
                                        v3 += -485
                                        _ = v41
                                        _ = v42
                                    }
                                case 2147483647:
                                    if false {
                                        v43 := GenS6{F0: ((v6 + v6) + ((((v6 + "") + (v6 + v6)) + ("foo bar" + v6)) + "foo bar")), F1: (^(v24))}
                                        switch fi21 {
                                        case 492:
                                            v44 := GenS3{F0: (fi22 + (fi22 >> ((65535) & 63))), F1: false}
                                            _ = v44
                                        case 2147483647:
                                            v45 := GenS3{F0: (fi22 - (^(fi21))), F1: false}
                                            _ = v45
                                        }
                                        _ = v43
                                    } else {
                                        if (!(!(false || false))) {
                                            v2 = (-2.25 - (v1 * v2))
                                        } else {
                                            v46 := GenS3{F0: 32767, F1: ((fi21 != fi21) && (((false || false) && (true || false)) || false))}
                                            _ = v46
                                        }
                                        var v47 uint8 = 2
                                        _ = v47
                                    }
                                }
                                var v48 bool = true
                                _ = v23
                                _ = v24
                                _ = v25
                                _ = v48
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = fi21
                                }
                            }
                            var v49 uint16 = 2
                            switch 32767 {
                            case 312:
                                var v50 uint16 = v49
                                _ = v50
                            case 9223372036854775807:
                                if (65535 >= 354) {
                                    v49 -= (v49 >> ((831) & 63))
                                    v51 := GenS1{F0: 0}
                                    _ = v51
                                }
                            default:
                                v52 := GenS3{F0: fi21, F1: ((((v6 + v6) + "\n\t") + (("a" + "\n\t") + "\n\t")) < "hello")}
                                _ = v52
                            }
                            _ = v49
                        } else {
                            var v53 int64 = -1
                            var v54 int8 = 63
                            _ = v53
                            _ = v54
                        }
                        var v55 bool = (false && (0 > -1))
                        v55 = v55
                        v56 := GenS7{F0: 627}
                        if (v6 <= "a") {
                            var v57 int32 = 215
                            var v58 uint = 832
                            fmt.Println(v57, v55)
                            switch v6 {
                            case "a":
                                fmt.Println(v3, v1, v6)
                            case "foo bar":
                                if v55 {
                                    v58 = v58
                                    v5[534] = (v6 + "")
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = len(v4)
                                    v6 = ((("a" + ("foo bar" + v6)) + v6) + v6)
                                }
                            default:
                                v58 += v58
                            }
                            var v59 float32 = -2.25
                            var v60 uint8 = 223
                            _ = v57
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                        _ = v55
                        _ = v56
                    }
                } else {
                    v61 := GenS8{F0: 18446744073709551615, F1: (0 >= -9223372036854775808), F2: 127}
                    if false {
                        var v62 uint32 = 1
                        v62 = (v62 & 65535)
                        delete(v5, 2)
                        _ = v62
                    } else {
                        v63 := v5[2147483647]
                        v64 := v5[42]
                        var v65 int = fi21
                        _ = v63
                        _ = v64
                        _ = v65
                    }
                    _ = v61
                }
                _ = fi22
            }
            if (98 == 8) {
                v6 = ("\n\t" + (v6 + (((v6 + (v6 + v6)) + "a") + (v6 + (v6 + "\n\t")))))
                var v66 int8 = 127
                var v67 uint = 1
                _ = v66
                _ = v67
            } else {
                v68 := GenS2{F0: (((1e10 >= v1) || ((v1 / (0.1 / v1)) == v2)) && false), F1: 255, F2: 2147483647}
                _ = v68
            }
            _ = fi21
        }
        if (2 < 24) {
            if len(v4) > 0 {
                v4[0] = 128
            }
        }
    }
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        v69 := GenS9{F0: false, F1: 18446744073709551615}
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
