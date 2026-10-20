package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64 }

func main() {
    var v1 uint32 = 42
    v2 := make(map[uint32]uint32)
    var v3 int8 = -128
    var v4 int64 = -9223372036854775808
    fmt.Println(v1, v1)
    v3 = (v3 >> ((2) & 63))
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 uint8 = 127
            if (1 <= 4294967295) {
                fmt.Println(v7, fi6, v1)
            } else {
                var v8 int32 = -226
                v2[v1] = v1
                _ = v8
            }
            _ = fi6
            _ = v7
        }
        var v9 uint64 = 2
        var v10 uint = 255
        _ = fi5
        _ = v10
        _ = v9
    }
    v4 -= v4
    v4 = -197
    var v11 int = 128
    v3 = v3
    v12 := GenS1{F0: (v1 + (v1 | (v1 % (((v1 >> ((661) & 63))) | 1)))), F1: v11, F2: (v11 - 1)}
    for fi13 := 0; fi13 < 19; fi13++ {
        v1 -= 0
        _ = fi13
    }
    v14 := GenS2{F0: 42, F1: 1}
    var v15 int16 = 32767
    v1 = (^(v1))
    if true {
        switch ("\n\t" + ("a" + "hello")) {
        case "foo bar":
            if (!true) {
                v16 := GenS3{F0: 127, F1: v1}
                for fi17 := 0; fi17 < 9; fi17++ {
                    v18 := GenS3{F0: (v3 | (v3 ^ int8(v4))), F1: (^(v1))}
                    for fi19 := 0; fi19 < 14; fi19++ {
                        var v20 int = 9223372036854775807
                        if true {
                            delete(v2, (v1 >> ((2) & 63)))
                            v21 := GenS4{F0: (-2147483648 <= 1), F1: -1}
                            _ = v21
                        } else {
                            v22 := []int{}
                            var v23 uint32 = (v1 + (v1 + v1))
                            v24 := GenS4{F0: true, F1: -791}
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        _ = fi19
                        _ = v20
                    }
                    v11 = v11
                    _ = fi17
                    _ = v18
                }
                _ = v16
            } else {
                if false {
                    switch v11 {
                    case 9223372036854775807:
                        switch (-(v11)) {
                        case 32768:
                            var v25 float32 = ((1.5 - (float32(1e10) * 3.14159)) + float32((0.0 + 0.1)))
                            _ = v25
                        }
                    case 819:
                        v11 -= v11
                    default:
                        v26 := []int{}
                        _ = v26
                    }
                    v11 += v11
                } else {
                    fmt.Println(v1, v15, v15)
                }
                if ((v4 == (^(v4))) || (-128 <= v3)) {
                    v1 = 199
                    switch ("\n\t" + "foo bar") {
                    case "a":
                        fmt.Println(v15)
                    case "":
                        for fi27 := 0; fi27 < 20; fi27++ {
                            for fi28 := 0; fi28 < 13; fi28++ {
                                fmt.Println(v11, v1, v1)
                                _ = fi28
                            }
                            _ = fi27
                        }
                    default:
                        if ((!(float64(0.1) < (1.5 + ((1.5 * (1.5 / 1.5)) + 1e10)))) || true) {
                            v29 := GenS2{F0: 44, F1: (v15 >> ((1) & 63))}
                            v11 = (v11 * (v11 * (v11 >> ((450) & 63))))
                            _ = v29
                        }
                    }
                    fmt.Println(v11, v1, v11)
                } else {
                    v30 := []int{0, 0, (v11 | (v11 * v11))}
                    if false {
                        for fi31 := 0; fi31 < 12; fi31++ {
                            var v32 int8 = v3
                            _ = fi31
                            _ = v32
                        }
                        for fi33 := 0; fi33 < 6; fi33++ {
                            if (0 < 1) {
                                if (false && (v15 < -32768)) {
                                    var v34 float64 = 1.5
                                    _ = v34
                                } else {
                                    v35 := v2[(v1 & (v1 % (((v1 / (((v1 | 2)) | 1))) | 1)))]
                                    switch "" {
                                    case "foo bar":
                                        v30 = append(v30, (^(v11)))
                                    case "\n\t":
                                        _ = len(v2)
                                    default:
                                        fmt.Println(fi33, fi33, v4)
                                    }
                                    _ = v35
                                }
                                for fi36 := 0; fi36 < 3; fi36++ {
                                    var v37 int64 = v4
                                    var v38 uint = 65535
                                    _ = fi36
                                    _ = v37
                                    _ = v38
                                }
                                for fi39 := 0; fi39 < 17; fi39++ {
                                    if false {
                                        fmt.Println(v3, fi39, v3)
                                        var v40 uint = 200
                                        switch fi39 {
                                        case -479:
                                            v41 := GenS2{F0: 51, F1: v15}
                                            _ = v41
                                        case 93:
                                            var v42 int8 = v3
                                            _ = v42
                                        default:
                                            v43 := GenS4{F0: true, F1: -1}
                                            _ = v43
                                        }
                                        _ = v40
                                    } else {
                                        delete(v2, v1)
                                        v15 += 42
                                    }
                                    _ = fi39
                                }
                            }
                            v4 += -9223372036854775808
                            _ = fi33
                        }
                        v30 = append(v30, v11)
                    } else {
                        switch int(v1) {
                        case 1:
                            if false {
                                _ = len(v2)
                            }
                        case -2147483648:
                            var v44 int8 = 127
                            _ = v44
                        }
                    }
                    _ = v30
                }
                if false {
                    if false {
                        fmt.Println(v15, v15, v11)
                    }
                    for fi45 := 0; fi45 < 5; fi45++ {
                        var v46 int32 = 42
                        _ = fi45
                        _ = v46
                    }
                } else {
                    for fi47 := 0; fi47 < 11; fi47++ {
                        fmt.Println(v4, v11, v1)
                        var v48 string = ""
                        _ = fi47
                        _ = v48
                    }
                    v49 := []int{}
                    if len(v49) >= 2 {
                        v49 = v49[1:len(v49)]
                    }
                    _ = v49
                }
            }
        }
        fmt.Println(v11, v11, v3)
    } else {
        v50 := GenS5{F0: (v4 + -1)}
        fmt.Println(v15, v11, v4)
        _ = v50
    }
    for fi51 := 0; fi51 < 1; fi51++ {
        _ = len(v2)
        if (!true) {
            for fi52 := 0; fi52 < 14; fi52++ {
                switch ("hello" + "foo bar") {
                case "a":
                    v15 += (-(v15))
                }
                if true {
                    v53 := GenS3{F0: 100, F1: v1}
                    var v54 uint64 = 519
                    v4 = 1
                    if (1 == 18446744073709551615) {
                        v55 := GenS2{F0: 128, F1: (v15 << ((0) & 63))}
                        var v56 uint32 = v1
                        switch v11 {
                        case 32767:
                            fmt.Println(v54, v4, v15)
                        case -774:
                            if false {
                                switch "" {
                                case "\n\t":
                                    switch -526 {
                                    case 190:
                                        var v57 uint32 = (v56 % (((v56 << ((2) & 63))) | 1))
                                        _ = v57
                                    case 2147483647:
                                        if false {
                                            fmt.Println(v11, fi52)
                                            v58 := GenS4{F0: true, F1: -961}
                                            _ = v58
                                        }
                                    case -9223372036854775808:
                                        var v59 int16 = v15
                                        _ = v59
                                    default:
                                        v60 := []int{4294967295, int(v1)}
                                        _ = v60
                                    }
                                case "":
                                    v61 := []int{fi52, 128}
                                    _ = v61
                                case "a":
                                    var v62 float32 = -2.25
                                    _ = v62
                                }
                            } else {
                                if true {
                                    if true {
                                        var v63 uint32 = (v56 | (v56 * v56))
                                        _ = v63
                                    }
                                    v64 := GenS3{F0: (v3 % ((v3) | 1)), F1: uint32(v54)}
                                    var v65 uint16 = 65535
                                    _ = v64
                                    _ = v65
                                } else {
                                    v66 := []int{(v11 + int(v56)), fi51}
                                    switch fi51 {
                                    case 42:
                                        v67 := GenS1{F0: (v1 - (v1 * v1)), F1: (v11 - 2147483647), F2: fi52}
                                        _ = v67
                                    case 1:
                                        if false {
                                            v2[(v1 * (v56 / ((2147483648) | 1)))] = (v1 / ((v1) | 1))
                                        } else {
                                            fmt.Println(v11, v1)
                                            v68 := GenS2{F0: 128, F1: (v15 >> ((930) & 63))}
                                            _ = v68
                                        }
                                    default:
                                        v3 = v3
                                    }
                                    v4 += v4
                                    v69 := GenS2{F0: 128, F1: 42}
                                    _ = v66
                                    _ = v69
                                }
                            }
                        case 2147483647:
                            v70 := GenS6{F0: -107}
                            _ = v70
                        }
                        _ = v55
                        _ = v56
                    } else {
                        var v71 float64 = 0.0
                        var v72 string = "\n\t"
                        _ = v71
                        _ = v72
                    }
                    _ = v53
                    _ = v54
                }
                v73 := GenS2{F0: 128, F1: v15}
                _ = fi52
                _ = v73
            }
            v74 := []int{fi51, v11, (v11 / (((v11 | v11)) | 1))}
            var v75 bool = (!false)
            _ = v74
            _ = v75
        } else {
            var v76 bool = ((((1e10 * 1e10) - 0.1) * 3.14159) < 0.1)
            switch "hello" {
            case "\n\t":
                if (!false) {
                    if (!false) {
                        var v77 int8 = 100
                        v3 = (-(v77))
                        _ = v77
                    } else {
                        fmt.Println(v3)
                    }
                    v3 -= 127
                }
            case "":
                for fi78 := 0; fi78 < 19; fi78++ {
                    v2[(^(v1))] = v1
                    _ = fi78
                }
            }
            _ = v76
        }
        var v79 int = fi51
        _ = fi51
        _ = v79
    }
    if (v1 < v1) {
        fmt.Println(v1, v11)
    } else {
        var v80 uint8 = 1
        _ = len(v2)
        _ = len(v2)
        _ = v80
    }
    var v81 int8 = (v3 >> ((uint(v1)) & 63))
    var v82 bool = (102 != 54)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v81
    _ = v82
}
