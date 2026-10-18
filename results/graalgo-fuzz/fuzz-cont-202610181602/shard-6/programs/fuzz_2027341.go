package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint16; F1 string; F2 string }
type GenS3 struct { F0 bool; F1 int16; F2 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 bool; F1 string }
type GenS8 struct { F0 uint; F1 int64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint64; F1 bool }
type GenS11 struct { F0 uint16; F1 int16; F2 int }
type GenS12 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 uint32 = 377
    var v2 bool = (true || ((true || ((v1 ^ (v1 << ((419) & 63))) > (v1 & 2147483647))) || false))
    var v3 uint32 = 4294967295
    var v4 int8 = 127
    v5 := GenS1{F0: "hello", F1: (v4 % (((v4 % (((v4 ^ v4)) | 1))) | 1))}
    if (true || (false && ((!(v2 || (3.14159 > ((1e10 * float64(0.0)) - 1e10)))) && (!(!v2))))) {
        v6 := GenS2{F0: 65535, F1: "a", F2: ""}
        var v7 int = 1
        v8 := []int{v7, 1, (v7 << ((1) & 63))}
        _ = v6
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := []int{128, fi9}
            _ = fi9
            _ = v10
        }
        v2 = (255 > 4)
        var v11 int64 = 882
        _ = v11
    }
    v12 := GenS1{F0: "", F1: (^(v4))}
    var v13 uint16 = 42
    fmt.Println(v4)
    v14 := []int{-2147483648, 55}
    if (((v13 * (v13 ^ v13)) >= (v13 >> ((2) & 63))) || (32767 <= int16(v3))) {
        switch 1 {
        case -621:
            if len(v14) > 0 {
                v14[0] = 9223372036854775807
            }
        case -2147483648:
            var v15 int64 = -655
            _ = v15
        case -814:
            v13 = (v13 | v13)
        default:
            if true {
                if (1 > 6) {
                    if (!v2) {
                        if ((v1 / ((v1) | 1)) == 4294967295) {
                            switch 798 {
                            case -2147483648:
                                for fi16 := 0; fi16 < 3; fi16++ {
                                    var v17 bool = v2
                                    fmt.Println(v17, fi16)
                                    _ = fi16
                                    _ = v17
                                }
                            default:
                                v18 := GenS3{F0: v2, F1: 32767, F2: -544}
                                _ = v18
                            }
                            var v19 int16 = 0
                            _ = v19
                        } else {
                            v4 = (v4 << ((uint(v3)) & 63))
                            v20 := GenS2{F0: 255, F1: "", F2: ((("\n\t" + ("hello" + "")) + "hello") + "\n\t")}
                            _ = v20
                        }
                    } else {
                        var v21 int64 = -9223372036854775808
                        _ = v21
                    }
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    v22 := GenS4{F0: -2147483648}
                    _ = v22
                } else {
                    v23 := GenS5{F0: 18446744073709551615, F1: 0, F2: 1}
                    switch 128 {
                    case 980:
                        if false {
                            switch "" {
                            case "\n\t":
                                if len(v14) > 0 {
                                    v14[0] = 764
                                }
                            case "hello":
                                if ((!v2) || v2) {
                                    for fi24 := 0; fi24 < 17; fi24++ {
                                        var v25 uint = 184
                                        var v26 int32 = 2147483647
                                        for fi27 := 0; fi27 < 18; fi27++ {
                                            if len(v14) > 0 {
                                                v14[0] = fi24
                                            }
                                            v28 := GenS3{F0: (42 != 0), F1: -32768, F2: (fi24 - fi24)}
                                            _ = fi27
                                            _ = v28
                                        }
                                        _ = fi24
                                        _ = v25
                                        _ = v26
                                    }
                                    v14 = append(v14, 127)
                                } else {
                                    v29 := GenS1{F0: ("hello" + "a"), F1: v4}
                                    if v2 {
                                        for fi30 := 0; fi30 < 9; fi30++ {
                                            v3 = v3
                                            v1 -= (v1 / (((v3 - uint32(v13))) | 1))
                                            _ = fi30
                                        }
                                        var v31 string = "foo bar"
                                        _ = v31
                                    } else {
                                        var v32 string = ""
                                        if v2 {
                                            v33 := GenS5{F0: 2, F1: 196, F2: 255}
                                            var v34 string = (v32 + v32)
                                            _ = v33
                                            _ = v34
                                        } else {
                                            var v35 float64 = (0.0 * (1e10 + -2.25))
                                            _ = v35
                                        }
                                        _ = v32
                                    }
                                    _ = v29
                                }
                            case "a":
                                var v36 string = "hello"
                                _ = v36
                            }
                            v2 = v2
                            for fi37 := 0; fi37 < 6; fi37++ {
                                v38 := GenS6{F0: 304, F1: (fi37 | fi37)}
                                for fi39 := 0; fi39 < 20; fi39++ {
                                    fmt.Println(v3, v4)
                                    v4 = (v4 + (^(v4)))
                                    v3 -= v1
                                    v13 += v13
                                    _ = fi39
                                }
                                _ = fi37
                                _ = v38
                            }
                        } else {
                            switch -887 {
                            case 0:
                                v2 = v2
                            default:
                                v40 := GenS7{F0: (!((!true) && (v2 && (!(false || v2))))), F1: ("" + "foo bar")}
                                _ = v40
                            }
                            fmt.Println(v1, v4)
                            v41 := GenS8{F0: 2, F1: -1}
                            _ = v41
                        }
                    default:
                        for fi42 := 0; fi42 < 13; fi42++ {
                            for fi43 := 0; fi43 < 18; fi43++ {
                                v4 -= int8(v1)
                                _ = fi43
                            }
                            if true {
                                if v2 {
                                    if len(v14) > 0 {
                                        v14[0] = (fi42 - (^(fi42)))
                                    }
                                    v44 := GenS9{F0: v2}
                                    var v45 uint = 42
                                    _ = v44
                                    _ = v45
                                }
                                v2 = (v4 <= v4)
                            }
                            _ = fi42
                        }
                    }
                    _ = v23
                }
                switch (("\n\t" + ("hello" + "foo bar")) + "foo bar") {
                case "foo bar":
                    var v46 int32 = 208
                    _ = v46
                case "":
                    if v2 {
                        if v2 {
                            v1 += v1
                            var v47 int = 368
                            if v2 {
                                fmt.Println(v13)
                                v48 := GenS10{F0: 816, F1: (true && true)}
                                _ = v48
                            }
                            _ = v47
                        } else {
                            v49 := GenS11{F0: 32767, F1: 1, F2: -524}
                            _ = v49
                        }
                    } else {
                        if (3.14159 <= ((1.5 - 3.14159) - float32((1e10 - (3.14159 - (float64((1e10 / (0.1 * (float32(1e10) / 1.5)))) + 3.14159)))))) {
                            for fi50 := 0; fi50 < 1; fi50++ {
                                var v51 string = "a"
                                _ = fi50
                                _ = v51
                            }
                            if (!(!(-302 == 0))) {
                                var v52 bool = (1.5 >= -2.25)
                                _ = v52
                            }
                        } else {
                            v14 = append(v14, -523)
                        }
                        for fi53 := 0; fi53 < 17; fi53++ {
                            v54 := GenS6{F0: -2147483648, F1: fi53}
                            _ = fi53
                            _ = v54
                        }
                    }
                default:
                    _ = len(v14)
                }
                var v55 uint8 = 2
                _ = v55
            } else {
                for fi56 := 0; fi56 < 17; fi56++ {
                    v14 = append(v14, fi56)
                    v57 := GenS3{F0: v2, F1: 1000, F2: 0}
                    var v58 int = fi56
                    _ = fi56
                    _ = v57
                    _ = v58
                }
            }
        }
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        v59 := GenS12{F0: 227, F1: 255}
        _ = v59
    } else {
        var v60 uint32 = v3
        v61 := GenS7{F0: (!(("a" + "foo bar") != ("hello" + ("foo bar" + "\n\t")))), F1: "\n\t"}
        _ = v60
        _ = v61
    }
    v4 = v4
    v3 += v1
    var v62 bool = (v4 != (v4 >> ((154) & 63)))
    for fi63 := 0; fi63 < 16; fi63++ {
        for fi64 := 0; fi64 < 5; fi64++ {
            switch -1 {
            case -769:
                var v65 int8 = v4
                _ = v65
            case 972:
                var v66 uint = 0
                _ = v66
            }
            v67 := GenS8{F0: 18446744073709551615, F1: 0}
            switch fi64 {
            case 9223372036854775807:
                var v68 uint = 4294967295
                _ = v68
            case 2147483647:
                v13 = (v13 % ((185) | 1))
            }
            _ = fi64
            _ = v67
        }
        _ = fi63
    }
    var v69 uint64 = 42
    fmt.Println(v69, v4, v62)
    var v70 int16 = 44
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v62:", v62)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v69
    _ = v70
}
