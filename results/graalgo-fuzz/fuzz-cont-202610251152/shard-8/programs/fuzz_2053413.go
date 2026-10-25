package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int16; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 int8; F2 int16 }
func genFn0() uint16 {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint8 = 1
        if (false || true) {
            switch (fi6 >> ((65535) & 63)) {
            case 899:
                fmt.Println(v7)
            case 2147483647:
                fmt.Println(v7, fi6)
            default:
                v8 := GenS1{F0: ("a" + "a")}
                _ = v8
            }
        } else {
            if (fi6 < (fi6 & 127)) {
                v9 := GenS2{F0: "foo bar"}
                fmt.Println(fi6, fi6)
                for fi10 := 0; fi10 < 13; fi10++ {
                    var v11 int16 = 42
                    var v12 float32 = 3.14159
                    switch ("a" + "a") {
                    case "\n\t":
                        if (!true) {
                            v13 := GenS3{F0: -128, F1: -186, F2: 42}
                            v7 += (v7 / ((v7) | 1))
                            v11 -= 1
                            _ = v13
                        } else {
                            var v14 int64 = 1
                            var v15 uint64 = 42
                            if (!(!false)) {
                                v16 := GenS3{F0: -128, F1: (^(v11)), F2: 42}
                                switch "a" {
                                case "\n\t":
                                    v17 := []int{fi10, fi6}
                                    _ = len(v17)
                                    if ((!((!(!(83 > 74))) && false)) || true) {
                                        v12 = v12
                                        v7 += v7
                                    }
                                    _ = v17
                                case "":
                                    var v18 bool = (!(v12 <= (v12 / (v12 + (((v12 * (v12 - float32(-2.25))) / v12) / 3.14159)))))
                                    _ = v18
                                }
                                _ = v16
                            }
                            _ = v14
                            _ = v15
                        }
                    default:
                        v12 -= 0.0
                    }
                    if true {
                        if false {
                            var v19 uint16 = 660
                            _ = v19
                        } else {
                            fmt.Println(fi6, v11, fi10)
                            v20 := GenS4{F0: 1}
                            v21 := GenS3{F0: 0, F1: (v11 ^ v11), F2: -304}
                            _ = v20
                            _ = v21
                        }
                    } else {
                        var v22 uint16 = 365
                        _ = v22
                    }
                    if (0 >= 101) {
                        var v23 uint16 = 191
                        v24 := []int{fi6, (-(fi6)), fi6, fi6}
                        if (3.14159 > float32(0.1)) {
                            fmt.Println(fi10)
                        }
                        _ = v23
                        _ = v24
                    } else {
                        var v25 float64 = 1e10
                        v26 := []int{(fi6 + fi10), fi10, 32767}
                        _ = v25
                        _ = v26
                    }
                    var v27 uint8 = v7
                    var v28 float64 = 3.14159
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v27
                    _ = v28
                }
                _ = v9
            }
        }
        _ = fi6
        _ = v7
    }
    var v29 uint16 = 65535
    var v30 int = 42
    if false {
        switch v30 {
        case -331:
            for fi31 := 0; fi31 < 10; fi31++ {
                v30 = v30
                v32 := []int{(fi31 & 1), v30, 128}
                _ = fi31
                _ = v32
            }
        case 1:
            v30 = v30
        default:
            var v33 uint64 = 938
            _ = v33
        }
    } else {
        v34 := []int{v30}
        var v35 int8 = -58
        _ = v34
        _ = v35
    }
    var v36 int8 = 100
    _ = v29
    _ = v30
    _ = v36
    return (v29 + v29)
}
func genFn1(p0 uint, p1 uint8) uint64 {
    fmt.Println(p0, p0)
    for fi37 := 0; fi37 < 2; fi37++ {
        v38 := GenS1{F0: ("\n\t" + ("\n\t" + ""))}
        _ = fi37
        _ = v38
    }
    var v39 int = 32768
    for fi40 := 0; fi40 < 12; fi40++ {
        fmt.Println(v39)
        v39 = 32767
        p1 -= p1
        _ = fi40
    }
    _ = p0
    _ = p1
    _ = v39
    return 202
}

func main() {
    v1 := []uint32{}
    var v2 uint16 = 145
    var v3 int = -2147483648
    var v4 float64 = -2.25
    var v5 uint16 = v2
    var v41 uint32 = 2147483647
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 42 {
    case 232:
        for fi42 := 0; fi42 < 15; fi42++ {
            if (42 != 1) {
                var v43 int = (-(fi42))
                for fi44 := 0; fi44 < 7; fi44++ {
                    v2 = genFn0()
                    fmt.Println(fi42)
                    v4 = ((float64(float32(v4)) - v4) - (((1.5 * v4) + 0.1) * 1.5))
                    _ = fi44
                }
                _ = v43
            } else {
                if true {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if true {
                        for fi45 := 0; fi45 < 10; fi45++ {
                            v46 := GenS5{F0: 974}
                            v41 += v41
                            _ = fi45
                            _ = v46
                        }
                        if false {
                            for fi47 := 0; fi47 < 3; fi47++ {
                                fmt.Println(fi47, fi47, fi42)
                                var v48 uint = 1
                                var v49 int8 = -128
                                _ = fi47
                                _ = v48
                                _ = v49
                            }
                            if (!(true || false)) {
                                v5 += (v2 - 65535)
                            }
                            if (!true) {
                                for fi50 := 0; fi50 < 13; fi50++ {
                                    var v51 float32 = (1.5 - (0.1 * float32(v4)))
                                    fmt.Println(fi50, v4)
                                    v4 -= (v4 - (v4 - v4))
                                    _ = fi50
                                    _ = v51
                                }
                            }
                        } else {
                            var v52 uint16 = 1
                            if true {
                                v5 = v2
                                var v53 uint16 = (v52 << ((255) & 63))
                                _ = v53
                            } else {
                                for fi54 := 0; fi54 < 10; fi54++ {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v55 int = fi42
                                    var v56 uint8 = 255
                                    _ = fi54
                                    _ = v55
                                    _ = v56
                                }
                                v5 = v52
                                v57 := GenS5{F0: 9223372036854775807}
                                _ = v57
                            }
                            _ = v52
                        }
                    }
                    v3 = v3
                }
                var v58 uint16 = 2
                fmt.Println(v58, v58, v3)
                _ = v58
            }
            var v59 int8 = 33
            for fi60 := 0; fi60 < 2; fi60++ {
                _ = len(v1)
                _ = fi60
            }
            _ = fi42
            _ = v59
        }
    case -740:
        fmt.Println(v2, v5)
    case 2147483647:
        v3 -= (v3 / ((v3) | 1))
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 = (0.0 / v4)
    var v61 float64 = float64(0.1)
    v5 = (v5 % ((v5) | 1))
    fmt.Println(v61, v2, v61)
    var v62 uint8 = 55
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v63 int16 = 610
    v5 = (^(v2))
    v64 := GenS2{F0: ""}
    if false {
        v65 := GenS6{F0: 42, F1: 1, F2: v63}
        if false {
            v2 = (v2 | 649)
            fmt.Println(v2, v4)
        }
        var v66 uint = 1
        _ = v65
        _ = v66
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v61
    _ = v62
    _ = v63
    _ = v64
}
