package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 bool; F2 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 uint8; F2 uint }

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 9223372036854775807
    var v3 uint32 = 0
    var v4 uint8 = 0
    var v5 uint16 = 2
    var v6 float64 = -2.25
    var v7 string = "hello"
    if false {
        if true {
            _ = len(v1)
            v8 := GenS1{F0: -430, F1: 65535}
            _ = v8
        }
        var v9 float64 = v6
        _ = v9
    }
    v10 := []int{128, 128}
    v11 := GenS1{F0: 2147483647, F1: 967}
    var v12 int32 = -2147483648
    switch -849 {
    case -9223372036854775808:
        var v13 int = 1
        _ = v13
    case -389:
        if (42 <= 1) {
            for fi14 := 0; fi14 < 4; fi14++ {
                v15 := GenS2{F0: (v3 << ((255) & 63))}
                v16 := GenS1{F0: 633, F1: 2}
                _ = fi14
                _ = v15
                _ = v16
            }
            v5 -= v5
        }
    case 42:
        fmt.Println(v5)
    default:
        switch 9223372036854775807 {
        case -524:
            for fi17 := 0; fi17 < 15; fi17++ {
                v1[-533] = (!(127 < v4))
                _ = fi17
            }
        case 1:
            v18 := GenS2{F0: (v3 >> ((255) & 63))}
            _ = v18
        }
    }
    v5 = (^(v5))
    fmt.Println(v5, v12, v4)
    if (-2.25 == -2.25) {
        var v19 int32 = (v12 - v12)
        v20 := v1[9223372036854775807]
        _ = v19
        _ = v20
    }
    var v21 int = 4294967295
    var v22 float64 = 1e10
    fmt.Println(v5)
    switch int(v4) {
    case 468:
        var v23 int8 = 100
        _ = v23
    case -1:
        var v24 uint8 = (v4 >> ((42) & 63))
        _ = v24
    default:
        if (true && false) {
            v6 -= (v22 + 3.14159)
            var v25 int16 = -1
            _ = v25
        }
    }
    v4 = uint8(v3)
    if (3 > 1000) {
        v5 = 65535
        for fi26 := 0; fi26 < 12; fi26++ {
            fmt.Println(v5)
            _ = fi26
        }
    } else {
        if true {
            v27 := GenS1{F0: v12, F1: 2}
            _ = v27
        }
    }
    _ = len(v1)
    if true {
        for fi28 := 0; fi28 < 2; fi28++ {
            switch int(v3) {
            case 1:
                v10 = append(v10, fi28)
            case -2147483648:
                v5 += 32767
            case -856:
                v3 -= uint32(v4)
            default:
                if (1e10 <= (v6 - v22)) {
                    if false {
                        switch (v7 + v7) {
                        case "\n\t":
                            if false {
                                v29 := GenS3{F0: 0}
                                _ = v29
                            } else {
                                var v30 string = (v7 + v7)
                                _ = v30
                            }
                        case "":
                            var v31 uint32 = uint32(v4)
                            _ = v31
                        case "hello":
                            for fi32 := 0; fi32 < 17; fi32++ {
                                if true {
                                    if ((v3 / ((65535) | 1)) == 2147483648) {
                                        if false {
                                            var v33 float32 = 1e10
                                            _ = v33
                                        } else {
                                            v3 -= v3
                                        }
                                        var v34 int8 = -107
                                        v35 := v1[(^(fi32))]
                                        _ = v34
                                        _ = v35
                                    } else {
                                        var v36 int16 = 42
                                        var v37 int = fi32
                                        _ = v36
                                        _ = v37
                                    }
                                    var v38 bool = ("" < ("a" + ("" + ((((v7 + v7) + "foo bar") + ((v7 + ("a" + (v7 + ((((v7 + "a") + v7) + (("a" + "hello") + (v7 + ""))) + v7)))) + v7)) + v7))))
                                    _ = v38
                                }
                                v39 := GenS4{F0: false, F1: (((v7 + (v7 + v7)) + "\n\t") < v7), F2: ""}
                                var v40 int8 = 0
                                _ = fi32
                                _ = v39
                                _ = v40
                            }
                        default:
                            v1[fi28] = ((v12 / ((v12) | 1)) > v12)
                        }
                    } else {
                        v2 += (v2 ^ -1)
                        for fi41 := 0; fi41 < 14; fi41++ {
                            v6 += float64(0.1)
                            fmt.Println(v2, v6)
                            var v42 int64 = 998
                            _ = fi41
                            _ = v42
                        }
                        v1[(-(fi28))] = (!true)
                    }
                    for fi43 := 0; fi43 < 4; fi43++ {
                        if (129 == v21) {
                            var v44 float64 = v6
                            v45 := GenS2{F0: v3}
                            var v46 uint16 = v5
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                        _ = fi43
                    }
                } else {
                    if true {
                        for fi47 := 0; fi47 < 9; fi47++ {
                            if false {
                                if (18446744073709551615 >= 4294967295) {
                                    var v48 uint16 = (v5 << ((82) & 63))
                                    fmt.Println(v5, v6, fi28)
                                    switch 32767 {
                                    case -1:
                                        delete(v1, fi47)
                                    case 128:
                                        v6 -= v22
                                        v21 = int(v4)
                                        v49 := GenS1{F0: v12, F1: 18446744073709551615}
                                        _ = v49
                                    case -72:
                                        if true {
                                            var v50 int64 = v2
                                            delete(v1, 134)
                                            var v51 uint64 = uint64(v3)
                                            _ = v50
                                            _ = v51
                                        }
                                    }
                                    switch fi47 {
                                    case -261:
                                        switch (fi28 & fi28) {
                                        case 32767:
                                            if ((v5 - (v5 / (((^(v5))) | 1))) != v48) {
                                                v22 = float64(-2.25)
                                            }
                                        case 1:
                                            delete(v1, (fi47 ^ 32767))
                                        }
                                    case -2147483648:
                                        if true {
                                            v1[(v21 % (((-(fi47))) | 1))] = true
                                            var v52 bool = false
                                            var v53 int32 = int32(v2)
                                            _ = v52
                                            _ = v53
                                        }
                                    case 414:
                                        var v54 string = v7
                                        v54 = v7
                                        _ = v54
                                    }
                                    switch int(v5) {
                                    case 42:
                                        var v55 int32 = v12
                                        _ = v55
                                    case -587:
                                        v3 = v3
                                    case -365:
                                        if (v4 >= v4) {
                                            switch fi47 {
                                            case -2147483648:
                                                var v56 uint64 = 0
                                                _ = v56
                                            case 788:
                                                v57 := GenS4{F0: false, F1: (-1 < 208), F2: (v7 + v7)}
                                                _ = v57
                                            case -9223372036854775808:
                                                v12 -= (-(v12))
                                                if true {
                                                    _ = len(v1)
                                                    v22 = v6
                                                    v4 = (v4 + v4)
                                                }
                                            default:
                                                var v58 int32 = v12
                                                _ = v58
                                            }
                                            var v59 uint32 = v3
                                            var v60 uint8 = 1
                                            _ = v59
                                            _ = v60
                                        } else {
                                            switch v7 {
                                            case "hello":
                                                v5 = 726
                                            case "":
                                                v21 += (fi47 | fi47)
                                            }
                                            v5 -= 32767
                                        }
                                    }
                                    v61 := GenS5{F0: 4294967295}
                                    _ = v48
                                    _ = v61
                                }
                                var v62 string = (v7 + v7)
                                _ = v62
                            } else {
                                _ = len(v10)
                            }
                            v12 -= v12
                            switch v7 {
                            case "":
                                if (((^(v3)) == v3) || true) {
                                    v22 += 0.0
                                    fmt.Println(v22, fi47, v4)
                                } else {
                                    var v63 int8 = 107
                                    fmt.Println(fi28)
                                    _ = v63
                                }
                            case "\n\t":
                                v64 := GenS6{F0: v2, F1: uint8(v3), F2: 2}
                                _ = v64
                            default:
                                fmt.Println(fi47, v12)
                            }
                            _ = fi47
                        }
                    } else {
                        v65 := GenS2{F0: (v3 * (v3 % (((v3 * (v3 >> ((4294967295) & 63)))) | 1)))}
                        v66 := v1[(v21 >> ((uint(v5)) & 63))]
                        var v67 uint16 = (v5 / ((v5) | 1))
                        _ = v65
                        _ = v66
                        _ = v67
                    }
                    v1[fi28] = (v7 < (((((v7 + v7) + v7) + v7) + "") + "foo bar"))
                    if (758 > 0) {
                        switch (fi28 & (^(fi28))) {
                        case 128:
                            for fi68 := 0; fi68 < 10; fi68++ {
                                v22 -= ((0.1 - v6) + 0.1)
                                v22 = (v22 * v6)
                                _ = fi68
                            }
                        case 646:
                            v12 -= (^(v12))
                        case -575:
                            switch (-(fi28)) {
                            case 29:
                                v69 := GenS6{F0: 0, F1: (v4 - v4), F2: 932}
                                _ = v69
                            }
                        default:
                            var v70 bool = (!(v2 <= (v2 >> ((65535) & 63))))
                            _ = v70
                        }
                        v7 = (v7 + "a")
                    }
                }
            }
            if false {
                v3 = (v3 ^ v3)
            } else {
                if true {
                    v71 := GenS4{F0: false, F1: false, F2: v7}
                    v12 = (-(v12))
                    _ = v71
                }
                fmt.Println(v2, v21, v12)
            }
            fmt.Println(v22, v21, v5)
            _ = fi28
        }
    }
    var v72 int64 = (-(v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v72
}
