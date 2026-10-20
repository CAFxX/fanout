package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 bool; F2 int8 }
type GenS6 struct { F0 int32; F1 string; F2 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint32; F1 int16 }
type GenS11 struct { F0 int32; F1 int32 }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[uint32]int)
    var v3 int = 127
    var v4 int8 = 1
    v5 := []uint32{12, 42}
    v6 := []int{-302, (v3 / (((-(v3))) | 1)), v3}
    for fi7 := 0; fi7 < 19; fi7++ {
        switch fi7 {
        case 0:
            v8 := GenS1{F0: 4294967295}
            _ = v8
        case 32768:
            v3 = (^(v3))
        }
        var v9 int = (fi7 + (fi7 - (fi7 ^ (v3 % ((-2147483648) | 1)))))
        _ = fi7
        _ = v9
    }
    if false {
        v5 = append(v5, 2)
        v5 = append(v5, 2147483648)
        if (true || false) {
            if len(v5) > 0 {
                v5[0] = 65535
            }
            v10 := GenS2{F0: 4294967295}
            var v11 int16 = 843
            _ = v10
            _ = v11
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
    }
    v12 := GenS3{F0: (v3 * (v3 / (((v3 | (v3 << ((170) & 63)))) | 1)))}
    switch v3 {
    case 128:
        _ = len(v6)
    default:
        var v13 float32 = 3.14159
        _ = v13
    }
    fmt.Println(v4, v4, v1)
    for fi14 := 0; fi14 < 16; fi14++ {
        fmt.Println(fi14)
        _ = len(v5)
        v15 := v2[2]
        _ = fi14
        _ = v15
    }
    v1 -= 0.0
    if true {
        if false {
            if true {
                v4 -= 73
                var v16 uint16 = 32767
                _ = v16
            }
        } else {
            v1 += v1
            v17 := GenS4{F0: 9223372036854775807}
            if false {
                fmt.Println(v3)
                if false {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        for fi19 := 0; fi19 < 2; fi19++ {
                            var v20 uint16 = 32767
                            _ = fi19
                            _ = v20
                        }
                        if (("hello" + "a") < "\n\t") {
                            _ = len(v2)
                        }
                        _ = fi18
                    }
                } else {
                    v3 -= 32767
                    if true {
                        v4 = -62
                        v21 := GenS3{F0: v3}
                        _ = v21
                    } else {
                        if true {
                            switch v3 {
                            case 0:
                                fmt.Println(v1)
                            case 127:
                                if (((0 <= 65) && true) || (1 < 0)) {
                                    if true {
                                        if true {
                                            v22 := GenS5{F0: 409, F1: false, F2: (v4 ^ v4)}
                                            _ = v22
                                        }
                                        var v23 string = ""
                                        if false {
                                            var v24 int32 = -2147483648
                                            v25 := GenS6{F0: v24, F1: v23, F2: 128}
                                            _ = v24
                                            _ = v25
                                        }
                                        _ = v23
                                    }
                                    var v26 bool = false
                                    _ = v26
                                } else {
                                    for fi27 := 0; fi27 < 8; fi27++ {
                                        if true {
                                            var v28 int32 = 645
                                            if (911 != 42) {
                                                var v29 uint32 = 882
                                                _ = v29
                                            } else {
                                                fmt.Println(fi27, v28)
                                                v30 := GenS7{F0: 1}
                                                var v31 uint64 = 50
                                                _ = v30
                                                _ = v31
                                            }
                                            _ = v28
                                        } else {
                                            _ = len(v2)
                                        }
                                        var v32 uint64 = 42
                                        v3 = v3
                                        _ = fi27
                                        _ = v32
                                    }
                                }
                            }
                            v1 += (3.14159 / v1)
                        }
                    }
                    switch "foo bar" {
                    case "\n\t":
                        v4 = (v4 - (v4 * 127))
                    case "a":
                        for fi33 := 0; fi33 < 17; fi33++ {
                            for fi34 := 0; fi34 < 15; fi34++ {
                                if true {
                                    v35 := GenS5{F0: -1, F1: (true || false), F2: v4}
                                    var v36 int8 = (v4 / ((v4) | 1))
                                    _ = v35
                                    _ = v36
                                }
                                _ = fi34
                            }
                            _ = fi33
                        }
                    case "hello":
                        if (!true) {
                            v1 = (float64(-2.25) - 0.1)
                            v37 := GenS2{F0: 403}
                            if (1000 < 0) {
                                delete(v2, 690)
                                for fi38 := 0; fi38 < 9; fi38++ {
                                    v1 = (v1 - v1)
                                    delete(v2, 2147483648)
                                    v39 := v2[2]
                                    _ = fi38
                                    _ = v39
                                }
                                var v40 bool = (42 != -9223372036854775808)
                                _ = v40
                            } else {
                                v41 := v2[1]
                                var v42 float32 = -2.25
                                v2[85] = v3
                                _ = v41
                                _ = v42
                            }
                            _ = v37
                        } else {
                            var v43 int16 = 42
                            v43 = v43
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = v43
                        }
                    }
                }
            } else {
                if true {
                    for fi44 := 0; fi44 < 12; fi44++ {
                        v4 += 42
                        var v45 uint32 = 955
                        v45 = (v45 - 2)
                        _ = fi44
                        _ = v45
                    }
                    fmt.Println(v1, v1)
                    switch (v3 * 32767) {
                    case 663:
                        var v46 int64 = -575
                        _ = v46
                    case 128:
                        for fi47 := 0; fi47 < 9; fi47++ {
                            var v48 int = 0
                            v49 := GenS6{F0: 1, F1: "hello", F2: 41}
                            fmt.Println(fi47, v4, v48)
                            _ = fi47
                            _ = v48
                            _ = v49
                        }
                    case 42:
                        if true {
                            v50 := GenS8{F0: ((2 <= 59) && true), F1: 255}
                            v51 := GenS9{F0: 429}
                            v4 = -128
                            _ = v50
                            _ = v51
                        } else {
                            v5 = append(v5, 2)
                        }
                    default:
                        if false {
                            v1 = v1
                        }
                    }
                } else {
                    v3 = v3
                    v52 := GenS6{F0: 670, F1: "\n\t", F2: 0}
                    v53 := GenS10{F0: 65535, F1: 390}
                    _ = v52
                    _ = v53
                }
                fmt.Println(v4, v1)
            }
            _ = v17
        }
    } else {
        if true {
            v1 = (float64(1e10) - 0.0)
            switch v3 {
            case 9223372036854775807:
                var v54 uint32 = 1
                _ = v54
            case 975:
                if true {
                    fmt.Println(v1)
                    var v55 uint32 = 2147483647
                    var v56 uint = 42
                    _ = v55
                    _ = v56
                }
            default:
                if ("\n\t" > "hello") {
                    v57 := GenS4{F0: -1}
                    _ = v57
                } else {
                    if true {
                        var v58 int = 9223372036854775807
                        _ = v58
                    } else {
                        for fi59 := 0; fi59 < 3; fi59++ {
                            var v60 int8 = 109
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v61 string = (("\n\t" + "\n\t") + "hello")
                            _ = fi59
                            _ = v60
                            _ = v61
                        }
                        v5 = append(v5, 42)
                    }
                    switch "foo bar" {
                    case "foo bar":
                        var v62 float64 = (v1 / 0.1)
                        _ = v62
                    case "a":
                        v3 -= -2147483648
                    case "hello":
                        v1 = v1
                    default:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                }
            }
            v63 := GenS8{F0: (-96 != 0), F1: 4294967295}
            _ = v63
        }
        fmt.Println(v4, v4, v3)
        for fi64 := 0; fi64 < 18; fi64++ {
            var v65 bool = false
            v4 -= 74
            var v66 float64 = (float64((1e10 - 0.1)) + (3.14159 + 1.5))
            _ = fi64
            _ = v65
            _ = v66
        }
    }
    fmt.Println(v3)
    v4 -= (v4 / (((-(v4))) | 1))
    v1 = v1
    fmt.Println(v3, v1, v1)
    var v67 uint16 = 255
    v4 = v4
    var v68 int32 = 1
    var v69 int8 = (v4 / (((v4 * (v4 & (v4 % ((42) | 1))))) | 1))
    v70 := GenS11{F0: v68, F1: v68}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
    _ = v69
    _ = v70
}
