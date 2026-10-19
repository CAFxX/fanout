package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int32; F1 int32; F2 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 0
    var v3 float64 = -2.25
    var v4 int = -2147483648
    var v5 int8 = 127
    if (v4 > (v4 | v4)) {
        var v6 uint = 42
        v7 := []int{v4, 1, v4}
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 uint = 4294967295
            _ = fi8
            _ = v9
        }
        _ = v6
        _ = v7
    }
    var v10 uint16 = 366
    var v11 uint32 = 1
    var v12 string = "a"
    var v13 float64 = 3.14159
    v14 := GenS1{F0: 65535, F1: -66}
    v10 -= v10
    v15 := []int{(v4 - 32768)}
    if false {
        switch v12 {
        case "":
            switch v12 {
            case "":
                v15 = append(v15, (v4 | 1))
            case "\n\t":
                for fi16 := 0; fi16 < 11; fi16++ {
                    v2 += v2
                    fmt.Println(v12)
                    _ = fi16
                }
            default:
                for fi17 := 0; fi17 < 10; fi17++ {
                    v10 += (v10 / ((65535) | 1))
                    var v18 string = v12
                    _ = fi17
                    _ = v18
                }
            }
        case "\n\t":
            v15 = append(v15, (v4 & v4))
        case "hello":
            for fi19 := 0; fi19 < 15; fi19++ {
                v20 := GenS2{F0: (v11 - (v11 / ((65535) | 1))), F1: 0}
                var v21 uint32 = (^(v11))
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        v15 = append(v15, (v4 / (((v4 - -727)) | 1)))
        var v22 float64 = 3.14159
        _ = v22
    } else {
        fmt.Println(v2, v4, v2)
    }
    fmt.Println(v10, v11, v5)
    for fi23 := 0; fi23 < 16; fi23++ {
        v2 = uint8(v11)
        if (4294967295 <= 42) {
            var v24 int64 = -1
            if len(v15) > 0 {
                v15[0] = int(v5)
            }
            _ = v24
        } else {
            var v25 uint64 = 1
            _ = v25
        }
        switch v12 {
        case "\n\t":
            for fi26 := 0; fi26 < 3; fi26++ {
                if ((^(v2)) <= 52) {
                    fmt.Println(v13, v13)
                    if false {
                        if (!true) {
                            var v27 int8 = (^(v5))
                            _ = v27
                        } else {
                            var v28 float64 = 0.0
                            _ = v28
                        }
                    }
                    var v29 int64 = 1
                    _ = v29
                }
                _ = fi26
            }
        case "a":
            _ = len(v15)
        case "hello":
            var v30 int64 = 1
            _ = v30
        default:
            for fi31 := 0; fi31 < 20; fi31++ {
                if false {
                    if (false && (-1 < -209)) {
                        v32 := GenS3{F0: 42, F1: 40, F2: (((("hello" + (v12 + "a")) + (v12 + (v12 + "hello"))) + (v12 + v12)) + ("hello" + v12))}
                        _ = v32
                    } else {
                        if (-287 != -32768) {
                            if (1.5 != v13) {
                                v12 = v12
                                var v33 int64 = 210
                                if (!("a" == v12)) {
                                    fmt.Println(v4, v3, v12)
                                    if true {
                                        if (-495 > -453) {
                                            v34 := GenS4{F0: v11}
                                            if false {
                                                var v35 int16 = -333
                                                var v36 float64 = 1.5
                                                _ = v35
                                                _ = v36
                                            }
                                            if (true && false) {
                                                var v37 int = int(v10)
                                                fmt.Println(v37, v5)
                                                _ = v37
                                            }
                                            _ = v34
                                        } else {
                                            v1[v12] = (((!(2147483647 == 0)) && ((((0.0 * 1e10) - 0.1) + 3.14159) >= float32(v3))) && ((v11 % ((0) | 1)) == v11))
                                            if false {
                                                v12 = ("a" + "\n\t")
                                                switch 42 {
                                                case 1:
                                                    v38 := GenS4{F0: (v11 / ((0) | 1))}
                                                    _ = v38
                                                }
                                            }
                                        }
                                        if (-1 != 0) {
                                            var v39 uint8 = (v2 << ((4294967295) & 63))
                                            switch v12 {
                                            case "foo bar":
                                                var v40 uint = 255
                                                _ = v40
                                            }
                                            if (false && true) {
                                                v10 += v10
                                            } else {
                                                var v41 uint = 848
                                                v42 := GenS5{F0: v33, F1: 18446744073709551615}
                                                _ = v41
                                                _ = v42
                                            }
                                            _ = v39
                                        }
                                    }
                                    if (42 > 65535) {
                                        if true {
                                            var v43 int64 = v33
                                            if ((^(v33)) <= int64(v11)) {
                                                fmt.Println(v11)
                                            } else {
                                                v44 := GenS6{F0: 42, F1: 42}
                                                _ = len(v1)
                                                _ = v44
                                            }
                                            if (!false) {
                                                _ = len(v15)
                                                v45 := GenS5{F0: (v43 | 0), F1: 1}
                                                var v46 int32 = 42
                                                _ = v45
                                                _ = v46
                                            } else {
                                                fmt.Println(v11, fi31, v12)
                                                var v47 int = (fi23 >> ((1) & 63))
                                                var v48 int = 128
                                                _ = v47
                                                _ = v48
                                            }
                                            if ((!true) || false) {
                                                var v49 float64 = (v13 * (v13 * 0.0))
                                                var v50 string = "\n\t"
                                                fmt.Println(v3, v4)
                                                _ = v49
                                                _ = v50
                                            }
                                            _ = v43
                                        }
                                        v51 := GenS7{F0: "hello"}
                                        v10 -= 65535
                                        _ = v51
                                    }
                                } else {
                                    if false {
                                        v5 = v5
                                    } else {
                                        var v52 uint64 = 2
                                        var v53 bool = true
                                        if true {
                                            fmt.Println(v2, v11, fi23)
                                            if (false || (65535 <= 4294967295)) {
                                                switch fi23 {
                                                case -891:
                                                    fmt.Println(v3)
                                                case 97:
                                                    fmt.Println(v11, v2)
                                                default:
                                                    switch 43 {
                                                    case -296:
                                                        v4 += -653
                                                    case -2147483648:
                                                        _ = len(v1)
                                                    case 868:
                                                        switch (fi31 | (^(v4))) {
                                                        case 127:
                                                            var v54 string = v12
                                                            _ = v54
                                                        }
                                                        fmt.Println(v53)
                                                        v11 = 2
                                                    }
                                                }
                                            } else {
                                                v55 := GenS8{F0: v10}
                                                if (v13 > v13) {
                                                    if len(v15) >= 2 {
                                                        v15 = v15[1:len(v15)]
                                                    }
                                                    var v56 bool = (((v11 << ((4294967295) & 63)) > 1) || v53)
                                                    _ = v56
                                                } else {
                                                    if v53 {
                                                        fmt.Println(fi23, v11)
                                                        if v53 {
                                                            v12 = "foo bar"
                                                        } else {
                                                            v57 := GenS4{F0: v11}
                                                            v10 -= v10
                                                            v13 = v13
                                                            _ = v57
                                                        }
                                                    }
                                                    v15 = append(v15, v4)
                                                    v58 := GenS9{F0: 1}
                                                    var v59 float32 = (3.14159 - 0.0)
                                                    if v53 {
                                                        if len(v15) >= 2 {
                                                            v15 = v15[1:len(v15)]
                                                        }
                                                        v60 := GenS10{F0: v10}
                                                        _ = v60
                                                    }
                                                    v61 := GenS7{F0: (v12 + "a")}
                                                    var v62 float32 = (((v59 - (v59 * float32(v3))) + v59) + -2.25)
                                                    _ = v58
                                                    _ = v59
                                                    _ = v61
                                                    _ = v62
                                                }
                                                var v63 uint16 = 157
                                                var v64 uint16 = v63
                                                v15 = append(v15, (fi31 & 2147483647))
                                                _ = v55
                                                _ = v63
                                                _ = v64
                                            }
                                        }
                                        var v65 int = v4
                                        _ = v52
                                        _ = v53
                                        _ = v65
                                    }
                                    v10 = 32767
                                    fmt.Println(v13, v33)
                                }
                                _ = v33
                            } else {
                                v2 = v2
                                v10 = 0
                            }
                        }
                        var v66 float32 = (1e10 - (((3.14159 * 0.0) * 3.14159) / 1.5))
                        switch ("foo bar" + v12) {
                        case "":
                            var v67 float64 = (-2.25 / -2.25)
                            _ = v67
                        case "hello":
                            v68 := GenS6{F0: -196, F1: 162}
                            _ = v68
                        case "a":
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                        default:
                            if (false && false) {
                                v69 := GenS2{F0: v11, F1: uint64(v10)}
                                fmt.Println(v66)
                                _ = v69
                            } else {
                                fmt.Println(v11)
                            }
                        }
                        _ = v66
                    }
                } else {
                    v12 = v12
                }
                var v70 float64 = float64((1e10 * 3.14159))
                _ = fi31
                _ = v70
            }
        }
        _ = fi23
    }
    fmt.Println(v4, v10)
    if (!(false && false)) {
        var v71 int16 = 0
        if false {
            var v72 float32 = 0.1
            var v73 uint32 = v11
            _ = v72
            _ = v73
        } else {
            var v74 uint64 = 42
            _ = len(v1)
            _ = v74
        }
        var v75 int = 32767
        _ = v71
        _ = v75
    }
    var v76 bool = (3.14159 >= -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v76
}
