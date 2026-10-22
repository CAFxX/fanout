package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 uint8; F1 int32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int8 }
func genFn0(p0 uint64, p1 int, p2 uint8) uint {
    for fi4 := 0; fi4 < 7; fi4++ {
        switch "foo bar" {
        case "":
            switch p1 {
            case 42:
                v5 := GenS1{F0: 1}
                _ = v5
            default:
                fmt.Println(p0, p2, p0)
            }
        case "hello":
            for fi6 := 0; fi6 < 20; fi6++ {
                var v7 float64 = (1e10 - (3.14159 + (3.14159 - (1e10 - float64((1.5 / 1e10))))))
                _ = fi6
                _ = v7
            }
        default:
            v8 := []int{p1, 127, (p1 / (((^(fi4))) | 1))}
            _ = v8
        }
        _ = fi4
    }
    switch p1 {
    case 9223372036854775807:
        var v9 int8 = 42
        _ = v9
    default:
        if true {
            switch -2147483648 {
            case -531:
                switch "a" {
                case "":
                    if true {
                        for fi10 := 0; fi10 < 12; fi10++ {
                            v11 := GenS1{F0: 2}
                            v12 := []int{int(p0)}
                            var v13 uint16 = 65535
                            _ = fi10
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                        v14 := GenS2{F0: 32767, F1: 1}
                        p1 += p1
                        _ = v14
                    }
                case "\n\t":
                    if true {
                        p2 -= p2
                        fmt.Println(p0)
                    } else {
                        if true {
                            for fi15 := 0; fi15 < 14; fi15++ {
                                var v16 uint32 = 2
                                _ = fi15
                                _ = v16
                            }
                        } else {
                            v17 := GenS3{F0: 189, F1: 42}
                            var v18 uint32 = 140
                            _ = v17
                            _ = v18
                        }
                        v19 := GenS4{F0: true, F1: (127 >= 47)}
                        _ = v19
                    }
                case "foo bar":
                    v20 := GenS5{F0: 1, F1: "foo bar"}
                    _ = v20
                default:
                    var v21 bool = (false || (1 <= 1))
                    _ = v21
                }
            case 128:
                v22 := []int{p1}
                _ = v22
            default:
                for fi23 := 0; fi23 < 20; fi23++ {
                    v24 := GenS6{F0: 42, F1: "foo bar"}
                    _ = fi23
                    _ = v24
                }
            }
            if (127 == 100) {
                var v25 int = -2147483648
                var v26 float32 = -2.25
                for fi27 := 0; fi27 < 14; fi27++ {
                    for fi28 := 0; fi28 < 2; fi28++ {
                        var v29 int = 1
                        _ = fi28
                        _ = v29
                    }
                    v30 := []int{128, (fi27 * p1), fi27, (v25 >> ((1) & 63))}
                    _ = fi27
                    _ = v30
                }
                _ = v25
                _ = v26
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 int32 = 0
    v2 := []uint32{55, 65535, 0, 860}
    var v3 uint = 317
    var v31 float64 = (1.5 + (3.14159 - float64((1e10 + 0.1))))
    if (!true) {
        switch "" {
        case "a":
            switch 32768 {
            case 32767:
                v3 = (^(v3))
            case -9223372036854775808:
                v32 := GenS6{F0: v3, F1: "a"}
                _ = v32
            case -848:
                v3 -= (v3 + v3)
            }
        default:
            if (103 <= 42) {
                var v33 string = "hello"
                if false {
                    var v34 float64 = -2.25
                    if false {
                        fmt.Println(v3)
                    } else {
                        if false {
                            switch v33 {
                            case "a":
                                for fi35 := 0; fi35 < 2; fi35++ {
                                    var v36 float32 = 3.14159
                                    v2 = append(v2, 602)
                                    if false {
                                        v37 := GenS7{F0: "a", F1: (18446744073709551615 > 954)}
                                        _ = v37
                                    } else {
                                        for fi38 := 0; fi38 < 9; fi38++ {
                                            switch 32767 {
                                            case 1:
                                                var v39 float32 = ((v36 / v36) + ((v36 + v36) / (3.14159 + (float32(v31) * v36))))
                                                v39 -= v39
                                                _ = v39
                                            case -1:
                                                switch fi35 {
                                                case 128:
                                                    if (!(false || false)) {
                                                        v31 += (v34 * 1.5)
                                                        fmt.Println(v3)
                                                        v40 := GenS8{F0: 127, F1: (v1 & -367)}
                                                        _ = v40
                                                    } else {
                                                        var v41 uint8 = 79
                                                        _ = v41
                                                    }
                                                default:
                                                    v42 := GenS9{F0: ""}
                                                    _ = v42
                                                }
                                            default:
                                                var v43 uint16 = 255
                                                _ = v43
                                            }
                                            v31 = (3.14159 * ((v31 + ((float64(v36) - v34) - float64(1e10))) + (float64(float32(v31)) * (-2.25 - (3.14159 - (-2.25 - (0.0 + v34)))))))
                                            var v44 uint16 = 402
                                            _ = fi38
                                            _ = v44
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = 747
                                        }
                                        _ = len(v2)
                                    }
                                    _ = fi35
                                    _ = v36
                                }
                            case "hello":
                                v45 := GenS5{F0: 4294967295, F1: ((((((v33 + ("\n\t" + v33)) + "") + "\n\t") + v33) + (v33 + (v33 + ((("" + "") + "foo bar") + (v33 + "hello"))))) + v33)}
                                _ = v45
                            default:
                                var v46 uint16 = uint16(v3)
                                _ = v46
                            }
                        } else {
                            v31 += v31
                        }
                        var v47 bool = true
                        _ = v47
                    }
                    _ = v34
                } else {
                    if (false || (606 == -9223372036854775808)) {
                        for fi48 := 0; fi48 < 12; fi48++ {
                            var v49 uint = (v3 | (v3 & v3))
                            switch v33 {
                            case "":
                                switch "\n\t" {
                                case "":
                                    for fi50 := 0; fi50 < 4; fi50++ {
                                        if true {
                                            var v51 int = (fi48 + fi48)
                                            _ = v51
                                        }
                                        _ = fi50
                                    }
                                case "hello":
                                    fmt.Println(v3, v31)
                                case "a":
                                    v3 = 65535
                                }
                            case "hello":
                                v31 = -2.25
                            case "foo bar":
                                for fi52 := 0; fi52 < 20; fi52++ {
                                    v3 -= genFn0(1, (fi48 * (fi52 >> ((v49) & 63))), 127)
                                    _ = fi52
                                }
                            default:
                                v53 := GenS10{F0: 90}
                                _ = v53
                            }
                            _ = fi48
                            _ = v49
                        }
                        var v54 uint16 = 753
                        fmt.Println(v1)
                        _ = v54
                    }
                    v55 := GenS1{F0: 255}
                    _ = v55
                }
                _ = v33
            }
        }
    }
    fmt.Println(v3, v1)
    var v56 int32 = (v1 & v1)
    var v57 uint = genFn0(0, 668, 127)
    switch "a" {
    case "foo bar":
        v57 -= v57
    case "":
        var v58 uint16 = 206
        _ = v58
    }
    fmt.Println(v31)
    var v59 int32 = v56
    if false {
        v1 += 0
    } else {
        var v60 int64 = 0
        fmt.Println(v31, v31, v57)
        _ = v60
    }
    v57 = v57
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v56
    _ = v57
    _ = v59
}
