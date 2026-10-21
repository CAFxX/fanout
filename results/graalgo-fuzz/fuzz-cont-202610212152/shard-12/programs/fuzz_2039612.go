package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 int64; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 int16; F1 int; F2 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint8; F1 int32; F2 int8 }
func genFn0(p0 int64) uint {
    v6 := GenS1{F0: 1}
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := []int{(fi8 ^ 1), 521, fi7, 4294967295}
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 100}
        _ = fi7
        _ = v10
    }
    var v11 uint = 568
    for fi12 := 0; fi12 < 10; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 uint16 = 266
            if true {
                v15 := []int{fi13, 0, (fi12 | fi12)}
                _ = len(v15)
                var v16 uint = v11
                v14 += (v14 >> ((v16) & 63))
                var v17 string = ""
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    v11 += 913
    _ = p0
    _ = v11
    _ = v6
    return (v11 / (((v11 & v11)) | 1))
}
func genFn1(p0 int32) uint16 {
    p0 = -2147483648
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 string = "\n\t"
        _ = fi18
        _ = v19
    }
    _ = p0
    return 42
}

func main() {
    v1 := []string{}
    var v2 int32 = -2147483648
    var v3 uint32 = 1
    v4 := make(map[string]int)
    var v5 uint32 = v3
    for fi20 := 0; fi20 < 6; fi20++ {
        var v21 int8 = 127
        _ = fi20
        _ = v21
    }
    var v22 float32 = float32(float64(float32(1.5)))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = (v2 & v2)
    var v23 uint64 = 637
    if (!((false && (42 > 0)) && false)) {
        switch "hello" {
        case "":
            v2 = (v2 + v2)
        case "a":
            var v24 int8 = -128
            _ = v24
        case "foo bar":
            if (-10 != -70) {
                fmt.Println(v5, v22, v5)
            } else {
                v2 -= -2147483648
                var v25 float32 = ((v22 - v22) / v22)
                delete(v4, "")
                _ = v25
            }
        default:
            var v26 bool = true
            _ = v26
        }
    } else {
        v23 = v23
    }
    if false {
        v27 := GenS1{F0: 42}
        _ = v27
    } else {
        var v28 uint16 = 2
        v29 := GenS2{F0: (v28 % ((genFn1(int32(v28))) | 1))}
        if (!true) {
            switch 1 {
            case -55:
                if true {
                    fmt.Println(v23, v5, v2)
                } else {
                    if len(v1) > 0 {
                        v1[0] = ("\n\t" + ("hello" + ""))
                    }
                    v28 = uint16(v5)
                }
            case 32767:
                v2 = v2
            }
            for fi30 := 0; fi30 < 14; fi30++ {
                var v31 int8 = 1
                _ = fi30
                _ = v31
            }
            for fi32 := 0; fi32 < 19; fi32++ {
                switch "foo bar" {
                case "hello":
                    var v33 int8 = 7
                    _ = v33
                case "":
                    _ = len(v4)
                case "a":
                    var v34 float32 = v22
                    _ = v34
                }
                _ = len(v1)
                var v35 uint = genFn0(-1)
                _ = fi32
                _ = v35
            }
        } else {
            v36 := GenS1{F0: 1}
            for fi37 := 0; fi37 < 3; fi37++ {
                switch (fi37 / (((fi37 | -965)) | 1)) {
                case 2147483647:
                    v38 := GenS2{F0: 42}
                    _ = v38
                case 862:
                    v39 := GenS2{F0: (v28 % ((v28) | 1))}
                    _ = v39
                case 32767:
                    for fi40 := 0; fi40 < 11; fi40++ {
                        var v41 int = (fi40 / (((^(fi37))) | 1))
                        v2 += (v2 << ((2) & 63))
                        var v42 int8 = -1
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                }
                var v43 uint16 = 2
                if false {
                    for fi44 := 0; fi44 < 15; fi44++ {
                        if (-56 != 42) {
                            switch "a" {
                            case "foo bar":
                                v4[((("\n\t" + ("foo bar" + ("hello" + ((("hello" + "foo bar") + "hello") + (("" + "") + ""))))) + "") + ("hello" + ""))] = fi37
                            case "a":
                                v5 = v3
                            case "":
                                v3 = 1
                            }
                            var v45 uint = 65535
                            _ = v45
                        }
                        v46 := GenS3{F0: 9223372036854775807, F1: 42, F2: 1}
                        v1 = append(v1, ("foo bar" + "a"))
                        _ = fi44
                        _ = v46
                    }
                    var v47 uint16 = v43
                    if (1 < genFn0(42)) {
                        var v48 int = (fi37 / ((fi37) | 1))
                        if (!true) {
                            var v49 uint32 = v3
                            for fi50 := 0; fi50 < 11; fi50++ {
                                if len(v1) > 0 {
                                    v1[0] = "hello"
                                }
                                if (int8(v49) >= 124) {
                                    _ = len(v1)
                                    var v51 uint8 = uint8(v49)
                                    _ = v51
                                } else {
                                    v43 -= genFn1(v2)
                                    switch fi37 {
                                    case 399:
                                        v52 := GenS4{F0: (v28 + (v47 / ((genFn1(701)) | 1))), F1: 255}
                                        _ = v52
                                    default:
                                        var v53 uint = 4294967295
                                        _ = v53
                                    }
                                    v54 := GenS5{F0: 42, F1: fi37, F2: 1}
                                    _ = v54
                                }
                                _ = fi50
                            }
                            v22 += v22
                            _ = v49
                        } else {
                            v2 = (v2 - v2)
                        }
                        v55 := GenS6{F0: (v48 - v48)}
                        _ = v48
                        _ = v55
                    }
                    _ = v47
                }
                _ = fi37
                _ = v43
            }
            _ = v36
        }
        _ = v28
        _ = v29
    }
    fmt.Println(v5, v3, v23)
    v5 -= 0
    if (genFn1((v2 | (v2 | 1))) >= genFn1(v2)) {
        switch 1 {
        case 32768:
            for fi56 := 0; fi56 < 1; fi56++ {
                if len(v1) > 0 {
                    v1[0] = ("" + "a")
                }
                v5 -= (^(v3))
                _ = fi56
            }
        case 540:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 128:
            v23 = 514
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        var v57 float64 = -2.25
        for fi58 := 0; fi58 < 3; fi58++ {
            fmt.Println(fi58)
            _ = fi58
        }
        if (v57 == -2.25) {
            v23 = 652
            v59 := GenS7{F0: ("" + "a")}
            _ = v59
        } else {
            for fi60 := 0; fi60 < 3; fi60++ {
                for fi61 := 0; fi61 < 20; fi61++ {
                    var v62 int32 = (^(v2))
                    fmt.Println(v22, v5, fi60)
                    var v63 int8 = 127
                    _ = fi61
                    _ = v62
                    _ = v63
                }
                switch (fi60 << ((18446744073709551615) & 63)) {
                case 128:
                    fmt.Println(fi60)
                case 42:
                    v23 = v23
                default:
                    v64 := v4[("\n\t" + "foo bar")]
                    _ = v64
                }
                v65 := GenS5{F0: 42, F1: (fi60 / (((fi60 & fi60)) | 1)), F2: 304}
                _ = fi60
                _ = v65
            }
            var v66 int32 = -151
            if (("hello" + "") == "\n\t") {
                v67 := GenS8{F0: 2, F1: v2, F2: 69}
                _ = v67
            }
            _ = v66
        }
        _ = v57
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
