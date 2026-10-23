package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 int32; F2 int }
type GenS6 struct { F0 int16; F1 uint; F2 int8 }
func genFn0(p0 uint32, p1 int) int {
    p0 = (p0 | 434)
    var v7 int8 = 1
    p0 -= 308
    switch "foo bar" {
    case "foo bar":
        var v8 uint32 = (p0 - (p0 ^ (^(p0))))
        _ = v8
    case "a":
        v7 -= (v7 >> ((255) & 63))
    case "":
        if false {
            v9 := GenS1{F0: (true || (float32(1.5) < 0.1)), F1: 42}
            v10 := GenS1{F0: (2 < (^(p0))), F1: 625}
            _ = v10
            _ = v9
        } else {
            for fi11 := 0; fi11 < 13; fi11++ {
                p0 = 94
                _ = fi11
            }
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return p1
}

func main() {
    var v1 float64 = -2.25
    var v2 uint8 = 30
    v3 := []string{"foo bar", "hello"}
    var v4 float32 = 3.14159
    v5 := []string{("\n\t" + "a"), "", "", ("a" + "a")}
    v6 := []uint32{597, 4294967295, 691, 1}
    fmt.Println(v2, v2, v4)
    fmt.Println(v2, v1)
    var v12 int8 = 0
    if true {
        switch "foo bar" {
        case "":
            switch -1 {
            case -1:
                var v13 int8 = (v12 | 13)
                _ = v13
            default:
                if false {
                    var v14 int64 = 0
                    v15 := GenS1{F0: true, F1: 186}
                    _ = v14
                    _ = v15
                }
            }
        default:
            v16 := GenS1{F0: true, F1: 18446744073709551615}
            _ = v16
        }
    } else {
        var v17 int64 = 270
        _ = v17
    }
    if len(v3) > 0 {
        v3[0] = ((((("foo bar" + ("a" + (("" + ("" + "")) + ("" + ("\n\t" + "\n\t"))))) + "a") + "\n\t") + "foo bar") + "hello")
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        v19 := GenS2{F0: 32767, F1: 2, F2: 933}
        if false {
            for fi20 := 0; fi20 < 17; fi20++ {
                v21 := GenS1{F0: false, F1: 901}
                _ = fi20
                _ = v21
            }
            v2 = v2
        } else {
            v22 := GenS1{F0: (-181 == -1), F1: 42}
            for fi23 := 0; fi23 < 2; fi23++ {
                var v24 int64 = 1
                v2 = v2
                var v25 uint64 = 4294967295
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = v22
        }
        _ = fi18
        _ = v19
    }
    v4 = float32(float64(v4))
    if false {
        v26 := GenS1{F0: (false && (1 > 255)), F1: 607}
        for fi27 := 0; fi27 < 9; fi27++ {
            for fi28 := 0; fi28 < 17; fi28++ {
                v12 += (v12 + v12)
                v4 = 0.1
                var v29 float64 = v1
                _ = fi28
                _ = v29
            }
            if (false || (!("a" <= "a"))) {
                if true {
                    v12 -= v12
                } else {
                    v2 -= v2
                    fmt.Println(fi27, fi27, v12)
                }
                v2 += v2
                fmt.Println(v2, v12, v12)
            }
            _ = fi27
        }
        _ = v26
    }
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    if false {
        for fi30 := 0; fi30 < 16; fi30++ {
            for fi31 := 0; fi31 < 10; fi31++ {
                if ((fi30 / (((fi30 + 2147483647)) | 1)) != (fi31 >> ((18446744073709551615) & 63))) {
                    v32 := GenS1{F0: (-1 != -32768), F1: 4294967295}
                    v33 := GenS3{F0: 0}
                    _ = v32
                    _ = v33
                } else {
                    if (!(-491 == -544)) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if true {
                            var v34 bool = (18446744073709551615 != 42)
                            var v35 int8 = v12
                            _ = v34
                            _ = v35
                        } else {
                            var v36 float64 = 0.0
                            if (-157 <= 0) {
                                switch (^(fi30)) {
                                case 0:
                                    var v37 bool = (-1 == -872)
                                    _ = v37
                                case 32767:
                                    switch fi30 {
                                    case 2147483647:
                                        v38 := GenS4{F0: (v12 ^ (^(v12)))}
                                        _ = v38
                                    case 32767:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    default:
                                        var v39 uint32 = 610
                                        _ = v39
                                    }
                                    if false {
                                        switch 32767 {
                                        case 32767:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        case 127:
                                            var v40 bool = (fi31 >= (fi31 & (fi30 & (fi31 % ((1) | 1)))))
                                            _ = v40
                                        default:
                                            v41 := GenS3{F0: 1000}
                                            _ = v41
                                        }
                                        fmt.Println(fi31, fi30)
                                    }
                                    v1 += v1
                                default:
                                    var v42 bool = (!true)
                                    v43 := GenS2{F0: 32767, F1: 383, F2: 0}
                                    _ = v42
                                    _ = v43
                                }
                                v44 := GenS5{F0: 42, F1: -1, F2: genFn0(688, (fi31 & (fi31 / ((fi30) | 1))))}
                                if len(v5) > 0 {
                                    v5[0] = "hello"
                                }
                                _ = v44
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = v36
                        }
                    } else {
                        switch "hello" {
                        case "foo bar":
                            if ((v2 + v2) != (v2 + v2)) {
                                v12 = v12
                            } else {
                                v45 := GenS6{F0: -672, F1: 4294967295, F2: v12}
                                v4 += 0.0
                                _ = v45
                            }
                            _ = len(v3)
                        case "a":
                            v2 -= 0
                        default:
                            v46 := GenS3{F0: 433}
                            _ = v46
                        }
                        fmt.Println(v12)
                        if true {
                            v3 = append(v3, "")
                            var v47 uint8 = (^(v2))
                            _ = len(v3)
                            _ = v47
                        } else {
                            var v48 uint32 = 42
                            _ = v48
                        }
                        v2 = (^(v2))
                        v49 := GenS6{F0: -686, F1: 18446744073709551615, F2: -1}
                        v1 -= (((((v1 - v1) + -2.25) + ((float64(-2.25) * 0.0) + v1)) * (v1 / v1)) - v1)
                        _ = v49
                    }
                    v1 -= 1e10
                    var v50 int16 = 120
                    _ = v50
                }
                v51 := GenS5{F0: 981, F1: 11, F2: genFn0(2147483647, genFn0(107, (fi30 / ((fi31) | 1))))}
                v4 = v4
                v12 += (v12 & int8(v2))
                var v52 int16 = 1
                _ = fi31
                _ = v51
                _ = v52
            }
            _ = fi30
        }
        v4 += v4
    } else {
        v12 = (v12 % ((24) | 1))
    }
    v1 = v1
    _ = len(v6)
    v53 := GenS4{F0: v12}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
}
