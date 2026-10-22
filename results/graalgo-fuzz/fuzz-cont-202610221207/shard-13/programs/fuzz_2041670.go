package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint32; F1 uint16 }
type GenS10 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint8; F1 bool; F2 uint }

func main() {
    v1 := []string{(("hello" + ("\n\t" + "\n\t")) + "\n\t"), "hello", "foo bar"}
    var v2 uint8 = 42
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    v5 := []string{"foo bar"}
    var v6 bool = (0 == 42)
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            if len(v1) > 0 {
                v1[0] = (("a" + "a") + ("\n\t" + "a"))
            }
            v6 = (100 != -1)
            _ = fi8
        }
        fmt.Println(v2)
        _ = fi7
    }
    v9 := v3[-797]
    if (!v6) {
        _ = len(v4)
    } else {
        fmt.Println(v9, v2, v2)
    }
    if v6 {
        if (-9223372036854775808 != -9223372036854775808) {
            var v10 string = "a"
            v11 := GenS1{F0: 429, F1: 2}
            _ = v10
            _ = v11
        }
        fmt.Println(v9)
        for fi12 := 0; fi12 < 5; fi12++ {
            if v6 {
                if v6 {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        v4[1] = (fi12 % ((644) | 1))
                        fmt.Println(v2)
                        _ = fi13
                    }
                    for fi14 := 0; fi14 < 19; fi14++ {
                        if false {
                            v2 = (v2 / (((^(v2))) | 1))
                        } else {
                            var v15 uint32 = 2147483648
                            _ = v15
                        }
                        v2 = 242
                        _ = fi14
                    }
                } else {
                    v6 = (1 == 2)
                    fmt.Println(fi12, v6, v9)
                    var v16 int = (-(fi12))
                    _ = v16
                }
                if v6 {
                    v2 = 0
                    var v17 float32 = 1e10
                    _ = v17
                } else {
                    if ((fi12 | fi12) < v9) {
                        v18 := GenS1{F0: 255, F1: 42}
                        for fi19 := 0; fi19 < 6; fi19++ {
                            v20 := GenS2{F0: (v2 / ((v2) | 1))}
                            if (false && v6) {
                                v21 := GenS2{F0: (^(v2))}
                                _ = len(v3)
                                _ = v21
                            } else {
                                if true {
                                    if (-2.25 <= 3.14159) {
                                        v2 = 0
                                        v6 = v6
                                        if (1 >= 1) {
                                            fmt.Println(v2)
                                            v1 = append(v1, "")
                                            fmt.Println(v9, fi12)
                                        } else {
                                            var v22 uint = 65535
                                            _ = v22
                                        }
                                        if v6 {
                                            if v6 {
                                                switch 128 {
                                                case 285:
                                                    v23 := GenS1{F0: 0, F1: 540}
                                                    _ = v23
                                                default:
                                                    var v24 string = "foo bar"
                                                    _ = v24
                                                }
                                                v25 := GenS3{F0: -89}
                                                var v26 uint16 = 621
                                                _ = v25
                                                _ = v26
                                            } else {
                                                _ = len(v1)
                                            }
                                            v4[(^(fi19))] = fi12
                                            v9 -= v9
                                        } else {
                                            v27 := GenS2{F0: 42}
                                            if len(v5) > 0 {
                                                v5[0] = ("hello" + ("a" + ("hello" + "hello")))
                                            }
                                            v28 := v4[v9]
                                            _ = v27
                                            _ = v28
                                        }
                                        v6 = (839 < 606)
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        switch fi12 {
                                        case 127:
                                            v6 = (4294967295 >= 42)
                                        }
                                        v2 = 0
                                    }
                                    delete(v4, (^(fi12)))
                                    switch "hello" {
                                    case "hello":
                                        fmt.Println(fi19, v2)
                                    case "\n\t":
                                        delete(v3, (-(fi19)))
                                    default:
                                        if v6 {
                                            v9 = fi19
                                            v29 := GenS4{F0: 1}
                                            _ = v29
                                        }
                                    }
                                } else {
                                    if (255 <= 882) {
                                        switch "foo bar" {
                                        case "foo bar":
                                            if ((^(v9)) >= (fi19 << ((0) & 63))) {
                                                fmt.Println(v6, v9, fi12)
                                                v2 = (v2 / (((v2 >> ((255) & 63))) | 1))
                                            } else {
                                                var v30 uint64 = 983
                                                var v31 float32 = 1e10
                                                var v32 int = -269
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            }
                                        default:
                                            var v33 uint64 = 288
                                            _ = v33
                                        }
                                        var v34 int32 = -587
                                        var v35 float32 = 1.5
                                        var v36 int = 127
                                        if len(v1) > 0 {
                                            v1[0] = "\n\t"
                                        }
                                        var v37 uint16 = 747
                                        if v6 {
                                            switch fi19 {
                                            case 4294967295:
                                                v34 -= -815
                                            case 128:
                                                var v38 uint32 = 4294967295
                                                _ = v38
                                            case 42:
                                                if (!v6) {
                                                    v3[(v9 / ((v36) | 1))] = v9
                                                } else {
                                                    if v6 {
                                                        v34 = (v34 & -930)
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        v39 := GenS5{F0: 2147483647, F1: 681}
                                                        _ = v39
                                                    } else {
                                                        v40 := GenS6{F0: -1}
                                                        _ = v40
                                                    }
                                                    var v41 int32 = v34
                                                    _ = v41
                                                }
                                            }
                                            v42 := GenS7{F0: "a"}
                                            _ = v42
                                        }
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                }
                            }
                            v43 := GenS6{F0: 1000}
                            _ = fi19
                            _ = v20
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 4; fi44++ {
                            if (834 > 1) {
                                if v6 {
                                    if (v6 || (v2 >= v2)) {
                                        if v6 {
                                            fmt.Println(fi44, v2, v6)
                                            v9 = -472
                                        }
                                        v3[fi44] = (v9 ^ (fi44 << ((4294967295) & 63)))
                                    } else {
                                        var v45 uint8 = (v2 & v2)
                                        var v46 float32 = 1e10
                                        v6 = (1000 > -775)
                                        var v47 string = "foo bar"
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                    }
                                }
                                v6 = v6
                                v6 = (!true)
                                if ("" != "") {
                                    v2 = v2
                                    fmt.Println(fi12)
                                }
                                fmt.Println(fi44, v6, fi44)
                            } else {
                                if v6 {
                                    v9 = fi44
                                    fmt.Println(fi12)
                                }
                                v6 = ((!false) || (v6 && (1.5 > -2.25)))
                            }
                            _ = fi44
                        }
                        _ = v18
                    }
                }
                v48 := GenS8{F0: 65535}
                _ = v48
            } else {
                var v49 int64 = 574
                _ = v49
            }
            var v50 int16 = -603
            var v51 float32 = 0.1
            _ = fi12
            _ = v50
            _ = v51
        }
    }
    fmt.Println(v6, v6)
    v52 := GenS9{F0: 2147483647, F1: 32767}
    v53 := GenS4{F0: 99}
    v54 := GenS10{F0: (^(v2)), F1: 42, F2: 18446744073709551615}
    v55 := GenS11{F0: 2}
    var v56 int16 = -1
    v57 := GenS12{F0: 255, F1: true, F2: 4294967295}
    v2 = 0
    fmt.Println(v2, v6)
    v58 := GenS5{F0: 1, F1: 964}
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v6
    _ = v9
}
