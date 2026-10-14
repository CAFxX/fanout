package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint; F2 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 int16; F2 int64 }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 uint64; F1 string; F2 int }
type GenS9 struct { F0 string }

func main() {
    var v1 string = ("hello" + "\n\t")
    var v2 int64 = 621
    var v3 int8 = 64
    v3 -= v3
    for fi4 := 0; fi4 < 7; fi4++ {
        v2 += (v2 ^ v2)
        for fi5 := 0; fi5 < 4; fi5++ {
            v6 := GenS1{F0: v1}
            v7 := GenS1{F0: "hello"}
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := GenS1{F0: v1}
        _ = fi4
        _ = v8
    }
    v9 := GenS2{F0: 47, F1: 32768, F2: 721}
    var v10 bool = (float64(3.14159) <= 1e10)
    var v11 float64 = (1.5 + 1.5)
    v3 += 122
    v12 := GenS1{F0: ((v1 + ((v1 + (((v1 + v1) + ("foo bar" + "foo bar")) + "")) + (v1 + ""))) + v1)}
    switch (v1 + "foo bar") {
    case "foo bar":
        v13 := GenS1{F0: v1}
        _ = v13
    }
    var v14 uint64 = 549
    if (v10 && true) {
        if ((!(128 == 1)) || v10) {
            v15 := []int{-9223372036854775808}
            if v10 {
                for fi16 := 0; fi16 < 13; fi16++ {
                    if true {
                        fmt.Println(v10)
                        v17 := GenS3{F0: 818}
                        v15 = append(v15, (fi16 + 4294967295))
                        _ = v17
                    }
                    v1 = v1
                    switch "" {
                    case "\n\t":
                        fmt.Println(v3, v1)
                    case "":
                        switch "a" {
                        case "\n\t":
                            switch 128 {
                            case -816:
                                v18 := GenS1{F0: v1}
                                _ = v18
                            case 32767:
                                for fi19 := 0; fi19 < 17; fi19++ {
                                    if len(v15) > 0 {
                                        v15[0] = (fi19 & -318)
                                    }
                                    v20 := GenS1{F0: (((v1 + v1) + "a") + v1)}
                                    fmt.Println(v11)
                                    v14 = v14
                                    var v21 float64 = (float64((1.5 / 1.0)) - ((((((v11 * v11) - v11) + float64((1.5 - 0.0))) - (v11 + float64(-2.25))) - v11) / v11))
                                    var v22 uint16 = 98
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                }
                            default:
                                if (v3 < v3) {
                                    var v23 string = ""
                                    _ = v23
                                } else {
                                    var v24 float32 = -2.25
                                    _ = v24
                                }
                            }
                        }
                    case "foo bar":
                        var v25 uint16 = 255
                        _ = v25
                    }
                    _ = fi16
                }
                v1 = ""
            } else {
                for fi26 := 0; fi26 < 1; fi26++ {
                    var v27 float64 = (float64((-2.25 * float32(-2.25))) + 1.5)
                    _ = fi26
                    _ = v27
                }
                v2 += 974
                switch v1 {
                case "hello":
                    for fi28 := 0; fi28 < 9; fi28++ {
                        for fi29 := 0; fi29 < 14; fi29++ {
                            if v10 {
                                var v30 uint8 = 0
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                                v3 = (v3 & 127)
                                switch "" {
                                case "":
                                    var v31 float64 = 1e10
                                    fmt.Println(fi29)
                                    _ = v31
                                case "foo bar":
                                    if (false || v10) {
                                        v30 = 2
                                        fmt.Println(v1)
                                        v32 := GenS2{F0: v30, F1: (fi29 / ((42) | 1)), F2: 484}
                                        if v10 {
                                            if (v2 == v2) {
                                                if (4294967295 < 4294967295) {
                                                    if len(v15) > 0 {
                                                        v15[0] = 0
                                                    }
                                                    var v33 string = v1
                                                    _ = v33
                                                } else {
                                                    if len(v15) > 0 {
                                                        v15[0] = (fi29 & fi29)
                                                    }
                                                }
                                                var v34 uint = 258
                                                var v35 float64 = v11
                                                _ = v34
                                                _ = v35
                                            }
                                            v3 = v3
                                            v36 := GenS4{F0: 4294967295, F1: uint(v30), F2: 2147483647}
                                            _ = v36
                                        }
                                        switch "a" {
                                        case "hello":
                                            var v37 int8 = v3
                                            _ = v37
                                        case "foo bar":
                                            var v38 int64 = (v2 & int64(v30))
                                            _ = v38
                                        case "\n\t":
                                            switch "a" {
                                            case "foo bar":
                                                var v39 uint32 = 313
                                                _ = v39
                                            }
                                        default:
                                            var v40 int64 = (v2 & (v2 | -9223372036854775808))
                                            _ = v40
                                        }
                                        _ = v32
                                    } else {
                                        var v41 int8 = v3
                                        if v10 {
                                            v11 = v11
                                            v42 := GenS2{F0: 255, F1: fi29, F2: 240}
                                            _ = v42
                                        } else {
                                            if true {
                                                v43 := GenS5{F0: 0}
                                                var v44 uint32 = 449
                                                switch v1 {
                                                case "hello":
                                                    v15 = append(v15, (fi29 + fi28))
                                                case "foo bar":
                                                    var v45 string = (v1 + ((v1 + ((v1 + v1) + ("\n\t" + "\n\t"))) + "foo bar"))
                                                    _ = v45
                                                }
                                                _ = v43
                                                _ = v44
                                            }
                                            v30 = (v30 + v30)
                                            v15 = append(v15, (^(fi29)))
                                            if (2147483648 > 0) {
                                                v14 -= (v14 ^ (v14 % ((v14) | 1)))
                                            }
                                        }
                                        fmt.Println(v1, fi29, v30)
                                        _ = v41
                                    }
                                case "a":
                                    var v46 string = (v1 + "a")
                                    _ = v46
                                }
                                _ = v30
                            } else {
                                if (v10 || (v3 >= v3)) {
                                    fmt.Println(v11, v3, v1)
                                    var v47 uint8 = 99
                                    fmt.Println(v1)
                                    v11 += v11
                                    _ = v47
                                }
                                if ((482 > 0) || v10) {
                                    v48 := GenS6{F0: v1, F1: -1, F2: (v2 | 9223372036854775807)}
                                    switch v1 {
                                    case "hello":
                                        fmt.Println(v1)
                                    case "":
                                        v14 = v14
                                    default:
                                        if len(v15) >= 2 {
                                            v15 = v15[1:len(v15)]
                                        }
                                    }
                                    fmt.Println(fi28)
                                    _ = v48
                                } else {
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                }
                            }
                            var v49 uint16 = 0
                            fmt.Println(v11, fi28)
                            if v10 {
                                if (0 == 569) {
                                    if len(v15) > 0 {
                                        v15[0] = fi28
                                    }
                                    v50 := GenS7{F0: (fi29 * (fi29 * 9223372036854775807)), F1: 1}
                                    _ = v50
                                }
                            }
                            var v51 uint32 = 42
                            _ = fi29
                            _ = v49
                            _ = v51
                        }
                        v11 = v11
                        if (!(v1 != (v1 + v1))) {
                            v52 := GenS8{F0: (v14 - v14), F1: (v1 + "hello"), F2: fi28}
                            var v53 int16 = 0
                            if v10 {
                                switch v1 {
                                case "hello":
                                    var v54 uint = 0
                                    _ = v54
                                case "\n\t":
                                    v55 := GenS7{F0: 1, F1: 2}
                                    _ = v55
                                case "":
                                    var v56 uint64 = (v14 * 4294967295)
                                    _ = v56
                                default:
                                    var v57 uint = 0
                                    _ = v57
                                }
                                for fi58 := 0; fi58 < 20; fi58++ {
                                    v59 := GenS7{F0: -1, F1: 32767}
                                    var v60 string = v1
                                    _ = fi58
                                    _ = v59
                                    _ = v60
                                }
                                v11 = 0.0
                            }
                            _ = v52
                            _ = v53
                        } else {
                            v15 = append(v15, fi28)
                            if true {
                                v2 += (v2 << ((471) & 63))
                            } else {
                                var v61 uint8 = 0
                                for fi62 := 0; fi62 < 8; fi62++ {
                                    var v63 int64 = (-(v2))
                                    var v64 uint = 604
                                    v3 = (v3 >> (((v64 << (((v64 * 551)) & 63))) & 63))
                                    _ = fi62
                                    _ = v63
                                    _ = v64
                                }
                                _ = v61
                            }
                        }
                        _ = fi28
                    }
                case "":
                    _ = len(v15)
                }
            }
            for fi65 := 0; fi65 < 3; fi65++ {
                switch fi65 {
                case -93:
                    var v66 string = (((("foo bar" + ("a" + v1)) + v1) + (v1 + v1)) + (v1 + v1))
                    _ = v66
                case -9223372036854775808:
                    var v67 string = "a"
                    _ = v67
                default:
                    var v68 uint = 65535
                    _ = v68
                }
                _ = fi65
            }
            _ = v15
        }
        v69 := GenS9{F0: v1}
        _ = v69
    } else {
        var v70 int32 = 2147483647
        var v71 uint16 = 0
        _ = v70
        _ = v71
    }
    if ((-506 >= 298) && ((972 < -193) || v10)) {
        var v72 uint8 = 78
        _ = v72
    } else {
        for fi73 := 0; fi73 < 7; fi73++ {
            v74 := []int{fi73, -240, 32768}
            fmt.Println(v10, v3)
            v75 := GenS1{F0: ((v1 + (v1 + v1)) + v1)}
            _ = fi73
            _ = v74
            _ = v75
        }
        var v76 uint32 = 42
        _ = v76
    }
    var v77 int32 = 2147483647
    var v78 string = ((((v1 + v1) + v1) + "hello") + v1)
    var v79 float64 = (-2.25 + v11)
    switch v1 {
    case "hello":
        v2 -= v2
    case "a":
        v79 += 0.1
    case "foo bar":
        if v10 {
            v80 := []int{}
            _ = v80
        }
    default:
        v77 -= v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v77
    _ = v78
    _ = v79
    _ = v9
}
