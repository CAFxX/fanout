package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int8; F1 int16; F2 int64 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 bool; F1 int16 }

func main() {
    var v1 int64 = -1
    v2 := make(map[int]int)
    var v3 int8 = 127
    var v4 float64 = 1.5
    var v5 bool = (true && false)
    var v6 uint64 = 2
    fmt.Println(v1, v5, v3)
    if (!(2147483648 < 42)) {
        if false {
            v6 = 0
        } else {
            if false {
                switch 150 {
                case 0:
                    v4 += (0.1 - v4)
                default:
                    v1 = v1
                }
                v7 := []int{}
                _ = v7
            } else {
                v2[-751] = 42
                switch "" {
                case "\n\t":
                    v8 := GenS1{F0: 127, F1: 0}
                    _ = v8
                case "":
                    v9 := v2[32767]
                    _ = v9
                case "hello":
                    var v10 int16 = 1000
                    _ = v10
                default:
                    var v11 uint8 = 150
                    _ = v11
                }
            }
            v12 := []int{}
            for fi13 := 0; fi13 < 7; fi13++ {
                v12 = append(v12, fi13)
                var v14 uint16 = 979
                _ = fi13
                _ = v14
            }
            _ = v12
        }
        switch ("\n\t" + ("hello" + "\n\t")) {
        case "\n\t":
            if (329 <= 255) {
                for fi15 := 0; fi15 < 10; fi15++ {
                    if (!(float32(float64(0.0)) > (-2.25 / ((1.5 * 1.5) * 3.14159)))) {
                        v6 = (v6 >> ((255) & 63))
                    } else {
                        v2[fi15] = (fi15 - (fi15 % ((fi15) | 1)))
                    }
                    if ("hello" < "\n\t") {
                        v3 += v3
                        v16 := GenS2{F0: 20, F1: -538, F2: v1}
                        _ = v16
                    } else {
                        v17 := v2[fi15]
                        v18 := GenS2{F0: 100, F1: -1, F2: (^(v1))}
                        v19 := []int{0, (v17 >> ((4294967295) & 63)), -1, 42}
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = fi15
                }
                if (!v5) {
                    for fi20 := 0; fi20 < 13; fi20++ {
                        var v21 uint8 = 233
                        var v22 float64 = (1.5 + 1e10)
                        for fi23 := 0; fi23 < 10; fi23++ {
                            var v24 int64 = -582
                            switch 9223372036854775807 {
                            case 32768:
                                delete(v2, int(v3))
                            }
                            switch -776 {
                            case 42:
                                if (((!(((v1 % (((v24 | v1)) | 1)) == -1) || ((-460 >= -75) && v5))) && ((844 >= 1) && false)) || (!v5)) {
                                    var v25 uint16 = 42
                                    fmt.Println(v5, fi23, v3)
                                    v25 += (v25 | v25)
                                    _ = v25
                                } else {
                                    v22 -= ((v22 + v22) * -2.25)
                                    v26 := []int{}
                                    _ = v26
                                }
                            default:
                                if v5 {
                                    v2[fi23] = fi23
                                    v27 := []int{-2147483648}
                                    if len(v27) > 0 {
                                        v27[0] = (-(fi20))
                                    }
                                    _ = v27
                                } else {
                                    v22 = v22
                                    v28 := []int{fi20}
                                    if ((!(-766 < 0)) || ((1 <= (v21 - 128)) || true)) {
                                        v29 := GenS2{F0: -1, F1: -32768, F2: v24}
                                        _ = v29
                                    } else {
                                        if false {
                                            if v5 {
                                                fmt.Println(v24, v4, v1)
                                                var v30 float64 = (1e10 / (v22 * (v22 + v22)))
                                                v31 := GenS1{F0: fi23, F1: 2147483647}
                                                _ = v30
                                                _ = v31
                                            }
                                            var v32 int16 = -796
                                            _ = v32
                                        } else {
                                            v33 := v2[(fi23 ^ (^(fi23)))]
                                            v34 := GenS2{F0: -1, F1: 1, F2: (^(v24))}
                                            var v35 float32 = 0.0
                                            v22 += (0.1 / v4)
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                        }
                                        var v36 uint = 0
                                        var v37 uint16 = 0
                                        _ = v36
                                        _ = v37
                                    }
                                    switch "\n\t" {
                                    case "a":
                                        fmt.Println(v21, v22)
                                    case "foo bar":
                                        var v38 int64 = (-(v24))
                                        _ = v38
                                    case "hello":
                                        v21 = (v21 | (v21 >> ((852) & 63)))
                                    default:
                                        if v5 {
                                            var v39 uint64 = 924
                                            v2[fi23] = (^(fi20))
                                            _ = v39
                                        }
                                    }
                                    var v40 int = fi23
                                    if (32767 >= -850) {
                                        var v41 int = 32767
                                        fmt.Println(fi23, v6)
                                        _ = v41
                                    }
                                    v42 := GenS2{F0: (v3 << ((18446744073709551615) & 63)), F1: 998, F2: (v1 & v1)}
                                    v43 := GenS1{F0: fi23, F1: 2147483648}
                                    _ = len(v28)
                                    if v5 {
                                        v24 = -1
                                    }
                                    _ = v28
                                    _ = v40
                                    _ = v42
                                    _ = v43
                                }
                            }
                            _ = fi23
                            _ = v24
                        }
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    v6 += 4294967295
                }
            } else {
                for fi44 := 0; fi44 < 19; fi44++ {
                    v45 := v2[fi44]
                    _ = fi44
                    _ = v45
                }
                fmt.Println(v6, v3, v4)
            }
        case "":
            v1 = v1
        }
        delete(v2, 0)
    } else {
        switch "" {
        case "foo bar":
            if (v5 && (!((!v5) && ((!(343 < 554)) && (!v5))))) {
                v46 := GenS2{F0: v3, F1: 1, F2: (v1 >> ((4294967295) & 63))}
                _ = v46
            }
        case "\n\t":
            v3 = v3
        default:
            v47 := GenS2{F0: 0, F1: -1, F2: v1}
            _ = v47
        }
        v2[9223372036854775807] = 4294967295
    }
    v48 := GenS1{F0: 127, F1: 42}
    v1 -= (-(v1))
    switch (("\n\t" + (((("" + "") + "\n\t") + "hello") + "a")) + "foo bar") {
    case "a":
        fmt.Println(v3, v4, v4)
    }
    switch 4294967295 {
    case 32768:
        var v49 string = "hello"
        _ = v49
    default:
        var v50 string = "a"
        _ = v50
    }
    for fi51 := 0; fi51 < 16; fi51++ {
        fmt.Println(v1, fi51)
        _ = fi51
    }
    v3 = (-(v3))
    var v52 uint8 = 128
    v1 += v1
    if ((-32 != -128) && (265 != uint32(v6))) {
        if (32768 == 32767) {
            if true {
                if true {
                    if (v3 <= v3) {
                        v2[719] = 0
                    } else {
                        var v53 bool = (681 < 1)
                        var v54 string = "hello"
                        v55 := GenS3{F0: 158, F1: 255}
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                } else {
                    v5 = ((525 >= -9223372036854775808) && true)
                }
                for fi56 := 0; fi56 < 17; fi56++ {
                    var v57 uint8 = 2
                    v57 = (v52 << ((18446744073709551615) & 63))
                    v52 = 2
                    _ = fi56
                    _ = v57
                }
            }
        }
        if (v5 && v5) {
            switch -812 {
            case -242:
                var v58 int8 = (v3 >> ((4294967295) & 63))
                _ = v58
            }
            if false {
                v3 = v3
            }
        } else {
            switch "" {
            case "foo bar":
                if v5 {
                    var v59 int32 = -504
                    if v5 {
                        v59 += (-(v59))
                        for fi60 := 0; fi60 < 3; fi60++ {
                            var v61 float64 = 1.5
                            v59 += v59
                            _ = fi60
                            _ = v61
                        }
                    } else {
                        if (1e10 == float32(v4)) {
                            for fi62 := 0; fi62 < 8; fi62++ {
                                var v63 uint64 = (^(v6))
                                var v64 uint32 = 4294967295
                                if (0 >= 1000) {
                                    fmt.Println(v5)
                                    v1 += (v1 | (v1 ^ v1))
                                } else {
                                    v65 := []int{(^(fi62)), fi62, fi62, fi62}
                                    if v5 {
                                        switch (^(fi62)) {
                                        case 2147483647:
                                            v65 = append(v65, fi62)
                                        case 42:
                                            if false {
                                                v66 := GenS3{F0: 1, F1: 1}
                                                for fi67 := 0; fi67 < 6; fi67++ {
                                                    v68 := GenS2{F0: (v3 * (v3 % ((v3) | 1))), F1: -647, F2: (v1 | (^(v1)))}
                                                    _ = fi67
                                                    _ = v68
                                                }
                                                v59 = 82
                                                _ = v66
                                            }
                                        default:
                                            v69 := GenS4{F0: (v59 + -1), F1: 127}
                                            _ = v69
                                        }
                                    }
                                    _ = v65
                                }
                                _ = fi62
                                _ = v63
                                _ = v64
                            }
                        }
                        var v70 uint64 = v6
                        v70 += v70
                        _ = v70
                    }
                    _ = v59
                } else {
                    delete(v2, 42)
                    v71 := GenS5{F0: false, F1: -1}
                    _ = v71
                }
            }
        }
        var v72 int16 = 0
        _ = v72
    }
    if (-1 < 0) {
        if (!false) {
            var v73 bool = (!true)
            _ = v73
        }
    }
    var v74 int16 = 42
    v4 = v4
    if ((1 != v6) && (360 <= 42)) {
        v5 = (((1e10 - v4) - (((0.0 / 1.0) - -2.25) + v4)) > 0.0)
        var v75 uint8 = v52
        _ = v75
    } else {
        v3 += int8(v1)
        fmt.Println(v52, v4, v6)
        fmt.Println(v6, v4, v1)
    }
    v76 := []int{128}
    v77 := GenS5{F0: v5, F1: (v74 >> ((419) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    fmt.Println("slicelen v76:", len(v76), v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v52
    _ = v6
    _ = v74
    _ = v76
    _ = v77
}
