package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 uint; F2 uint }
type GenS7 struct { F0 int8; F1 uint; F2 uint16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS10 struct { F0 int16; F1 int16 }

func main() {
    var v1 int16 = -709
    var v2 uint16 = 65535
    var v3 int16 = (v1 + (v1 % (((v1 % (((v1 % (((v1 << ((65535) & 63))) | 1))) | 1))) | 1)))
    var v4 uint16 = v2
    var v5 string = "hello"
    v6 := []int{-1, -969, 9223372036854775807}
    var v7 bool = false
    var v8 string = v5
    v9 := GenS1{F0: 255, F1: (!true)}
    for fi10 := 0; fi10 < 11; fi10++ {
        if (false || false) {
            if ((!v7) && v7) {
                fmt.Println(v2)
                switch v8 {
                case "\n\t":
                    if false {
                        for fi11 := 0; fi11 < 18; fi11++ {
                            if len(v6) > 0 {
                                v6[0] = fi11
                            }
                            v5 = v5
                            v12 := GenS1{F0: 0, F1: v7}
                            _ = fi11
                            _ = v12
                        }
                        _ = len(v6)
                    }
                case "":
                    if v7 {
                        for fi13 := 0; fi13 < 2; fi13++ {
                            _ = len(v6)
                            v14 := GenS2{F0: 127}
                            _ = fi13
                            _ = v14
                        }
                        v1 -= 32767
                    } else {
                        fmt.Println(v2, fi10)
                        if ((!true) || (v3 > v3)) {
                            var v15 int16 = (v3 | int16(v2))
                            for fi16 := 0; fi16 < 9; fi16++ {
                                var v17 float32 = 3.14159
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                var v18 string = v5
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            v19 := GenS2{F0: 128}
                            _ = v15
                            _ = v19
                        }
                        if len(v6) > 0 {
                            v6[0] = -512
                        }
                    }
                case "a":
                    for fi20 := 0; fi20 < 8; fi20++ {
                        switch "\n\t" {
                        case "hello":
                            v21 := GenS3{F0: 1, F1: (^(v2)), F2: 100}
                            _ = v21
                        default:
                            var v22 uint32 = 1
                            _ = v22
                        }
                        v23 := GenS1{F0: 0, F1: (v7 || v7)}
                        v24 := GenS3{F0: 2, F1: v4, F2: 100}
                        _ = fi20
                        _ = v23
                        _ = v24
                    }
                default:
                    var v25 int64 = 437
                    _ = v25
                }
                if ("hello" >= v8) {
                    var v26 uint32 = 1
                    for fi27 := 0; fi27 < 20; fi27++ {
                        if len(v6) > 0 {
                            v6[0] = fi10
                        }
                        switch (fi10 | fi10) {
                        case -9223372036854775808:
                            fmt.Println(v26)
                        case 1:
                            var v28 uint8 = 70
                            _ = v28
                        case 207:
                            var v29 int8 = -128
                            _ = v29
                        }
                        if v7 {
                            if (-32768 > 1000) {
                                v2 = v4
                                switch fi27 {
                                case -9223372036854775808:
                                    switch v5 {
                                    case "a":
                                        v30 := GenS4{F0: 809, F1: 42}
                                        _ = v30
                                    default:
                                        v6 = append(v6, fi27)
                                    }
                                case -557:
                                    if ((true && ((v3 & v3) != (v1 ^ (-(v3))))) || v7) {
                                        var v31 float32 = 0.1
                                        var v32 string = v5
                                        var v33 bool = v7
                                        v32 = v5
                                        if false {
                                            if v7 {
                                                var v34 int64 = -9223372036854775808
                                                fmt.Println(v33, v34, v2)
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                _ = v34
                                            }
                                            v4 = (v4 / (((v2 | 32767)) | 1))
                                            var v35 uint = 833
                                            _ = v35
                                        } else {
                                            var v36 uint16 = (v4 + 1)
                                            _ = v36
                                        }
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                }
                                var v37 uint = 974
                                _ = v37
                            }
                        } else {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v7 = (255 <= (v4 ^ (v2 / ((931) | 1))))
                        }
                        _ = fi27
                    }
                    _ = v26
                }
            }
            _ = len(v6)
            if v7 {
                var v38 uint = 65535
                v1 -= v1
                _ = v38
            } else {
                var v39 int32 = -966
                v40 := GenS5{F0: -643}
                _ = v39
                _ = v40
            }
        } else {
            v41 := GenS6{F0: 18446744073709551615, F1: 565, F2: 275}
            _ = v41
        }
        _ = len(v6)
        _ = fi10
    }
    var v42 uint16 = (v2 * 783)
    v43 := GenS5{F0: 42}
    if v7 {
        if ((v8 > (v5 + (v8 + v8))) || (v5 == v8)) {
            v1 = v3
        } else {
            var v44 uint = 793
            _ = v44
        }
        v5 = v5
    } else {
        for fi45 := 0; fi45 < 10; fi45++ {
            v46 := GenS7{F0: 127, F1: 65535, F2: 255}
            v7 = (((53 != 0) && (v8 <= v8)) && true)
            _ = fi45
            _ = v46
        }
        v47 := GenS7{F0: 42, F1: 223, F2: (v42 | (v2 >> ((65535) & 63)))}
        switch ("hello" + (v8 + ((v8 + v5) + v5))) {
        case "foo bar":
            for fi48 := 0; fi48 < 7; fi48++ {
                if ((!(v4 <= (^(v42)))) && true) {
                    if (v7 || (fi48 <= (fi48 ^ fi48))) {
                        var v49 int8 = 42
                        if v7 {
                            var v50 int32 = 859
                            if len(v6) > 0 {
                                v6[0] = -9223372036854775808
                            }
                            v5 = "foo bar"
                            _ = v50
                        } else {
                            fmt.Println(v1)
                            v7 = (-650 <= -1)
                        }
                        _ = v49
                    } else {
                        if false {
                            v8 = ("" + v5)
                        }
                        if (2 > 2) {
                            var v51 int = (fi48 - fi48)
                            v4 += v2
                            _ = v51
                        } else {
                            v5 = "hello"
                            v52 := GenS5{F0: -258}
                            _ = v52
                        }
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                } else {
                    fmt.Println(v7)
                    var v53 uint64 = 4294967295
                    if true {
                        if (128 == 131) {
                            for fi54 := 0; fi54 < 14; fi54++ {
                                switch fi48 {
                                case -9223372036854775808:
                                    v55 := GenS1{F0: 128, F1: (100 >= -92)}
                                    _ = v55
                                case -1:
                                    v2 = v2
                                }
                                if len(v6) > 0 {
                                    v6[0] = fi48
                                }
                                _ = fi54
                            }
                            switch (("foo bar" + v8) + v5) {
                            case "foo bar":
                                for fi56 := 0; fi56 < 20; fi56++ {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    if len(v6) > 0 {
                                        v6[0] = fi56
                                    }
                                    _ = fi56
                                }
                            }
                        }
                        switch (fi48 % ((-9223372036854775808) | 1)) {
                        case -9223372036854775808:
                            v3 = -32768
                        case 144:
                            v57 := GenS8{F0: (v53 ^ (v53 ^ 89))}
                            _ = v57
                        }
                        if (!v7) {
                            v5 = "foo bar"
                        } else {
                            v58 := GenS2{F0: 2}
                            v7 = v7
                            _ = v58
                        }
                    }
                    _ = v53
                }
                var v59 float64 = float64(0.1)
                var v60 float32 = float32(v59)
                _ = fi48
                _ = v59
                _ = v60
            }
        case "":
            v2 = v2
        case "hello":
            for fi61 := 0; fi61 < 19; fi61++ {
                v62 := GenS9{F0: v42, F1: (v4 != 2), F2: 42}
                _ = fi61
                _ = v62
            }
        }
        _ = v47
    }
    for fi63 := 0; fi63 < 4; fi63++ {
        for fi64 := 0; fi64 < 15; fi64++ {
            v65 := GenS10{F0: 1000, F1: 305}
            _ = fi64
            _ = v65
        }
        for fi66 := 0; fi66 < 11; fi66++ {
            if (42 > 168) {
                var v67 uint16 = v42
                var v68 float64 = 0.0
                _ = len(v6)
                _ = v67
                _ = v68
            } else {
                var v69 string = v5
                fmt.Println(v7, v1, fi66)
                if true {
                    v8 = ""
                    switch 42 {
                    case -295:
                        v70 := GenS8{F0: 729}
                        var v71 uint8 = 128
                        if len(v6) > 0 {
                            v6[0] = -149
                        }
                        _ = v70
                        _ = v71
                    case -610:
                        _ = len(v6)
                    case 2147483647:
                        var v72 int32 = 42
                        _ = v72
                    default:
                        v8 = v69
                    }
                }
                _ = v69
            }
            v2 += (v42 ^ (v4 ^ (^(v42))))
            _ = fi66
        }
        _ = fi63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
