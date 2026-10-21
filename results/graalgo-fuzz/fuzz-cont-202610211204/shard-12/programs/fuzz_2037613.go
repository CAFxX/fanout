package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int16; F1 uint; F2 int }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS7 struct { F0 uint64; F1 int16; F2 uint }
type GenS8 struct { F0 int8; F1 uint16; F2 int8 }
func genFn0() bool {
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := GenS1{F0: 255, F1: "foo bar"}
        var v6 float64 = 1.5
        var v7 int32 = -1
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 bool = true
    _ = v8
    return (float32(3.14159) <= 0.0)
}

func main() {
    var v1 uint32 = 1
    var v2 float32 = 1.5
    var v3 string = ("hello" + "\n\t")
    var v9 int32 = 2147483647
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := []int{32768, fi10}
        var v12 float32 = 3.14159
        _ = fi10
        _ = v11
        _ = v12
    }
    if false {
        fmt.Println(v9, v1, v1)
    } else {
        for fi13 := 0; fi13 < 3; fi13++ {
            if true {
                fmt.Println(v1, v2)
                var v14 int = 4294967295
                if true {
                    v15 := GenS2{F0: uint(v1)}
                    for fi16 := 0; fi16 < 1; fi16++ {
                        v14 = v14
                        _ = fi16
                    }
                    fmt.Println(v1, fi13, v9)
                    _ = v15
                } else {
                    v17 := GenS2{F0: 255}
                    v3 = ""
                    _ = v17
                }
                _ = v14
            } else {
                for fi18 := 0; fi18 < 12; fi18++ {
                    var v19 uint = 447
                    fmt.Println(v1, v1, v9)
                    var v20 string = ("foo bar" + ("a" + "foo bar"))
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
            v21 := GenS2{F0: 65535}
            v22 := GenS2{F0: 2}
            _ = fi13
            _ = v21
            _ = v22
        }
    }
    var v23 uint64 = 24
    var v24 int64 = -548
    v25 := []int{1, -75}
    v9 = v9
    var v26 uint32 = v1
    var v27 int = -760
    v1 -= (v26 << ((469) & 63))
    fmt.Println(v9, v26)
    for fi28 := 0; fi28 < 1; fi28++ {
        v3 = (v3 + v3)
        var v29 bool = (!false)
        switch int(v23) {
        case 42:
            if (18446744073709551615 > 4294967295) {
                v30 := GenS2{F0: 271}
                var v31 float64 = float64((1e10 * float32(float64(3.14159))))
                var v32 float64 = v31
                _ = v30
                _ = v31
                _ = v32
            } else {
                for fi33 := 0; fi33 < 3; fi33++ {
                    v25 = append(v25, (fi33 >> ((255) & 63)))
                    if (!false) {
                        v34 := GenS1{F0: 128, F1: "foo bar"}
                        _ = v34
                    }
                    v24 = (v24 - v24)
                    v26 += 4294967295
                    v24 -= 911
                    switch (fi33 * 32767) {
                    case -458:
                        var v35 uint16 = 0
                        _ = v35
                    case 0:
                        v27 = v27
                    case 492:
                        fmt.Println(v9, v27)
                    default:
                        _ = len(v25)
                    }
                    _ = fi33
                }
                if (("\n\t" + v3) == v3) {
                    switch (v27 << ((uint(v1)) & 63)) {
                    case 1:
                        var v36 float64 = 0.1
                        _ = v36
                    case 32768:
                        switch v27 {
                        case 42:
                            v37 := GenS3{F0: "", F1: 18446744073709551615}
                            _ = v37
                        }
                    default:
                        fmt.Println(fi28, fi28, v29)
                    }
                    if v29 {
                        if false {
                            v3 = v3
                            for fi38 := 0; fi38 < 16; fi38++ {
                                switch "foo bar" {
                                case "hello":
                                    var v39 uint16 = 2
                                    _ = v39
                                }
                                fmt.Println(v3, v2)
                                _ = fi38
                            }
                        } else {
                            for fi40 := 0; fi40 < 16; fi40++ {
                                var v41 float32 = ((v2 * v2) - v2)
                                _ = fi40
                                _ = v41
                            }
                            if (!v29) {
                                for fi42 := 0; fi42 < 6; fi42++ {
                                    var v43 string = ("foo bar" + "")
                                    var v44 int32 = (v9 + (v9 % ((597) | 1)))
                                    v45 := GenS1{F0: 127, F1: ("foo bar" + v43)}
                                    fmt.Println(v2, fi28)
                                    v1 += (v1 / ((311) | 1))
                                    v44 = v9
                                    if (uint32(v23) == (v26 / ((v26) | 1))) {
                                        if v29 {
                                            if v29 {
                                                if (!true) {
                                                    if len(v25) > 0 {
                                                        v25[0] = 1
                                                    }
                                                } else {
                                                    fmt.Println(v26, v44, v43)
                                                }
                                                if len(v25) >= 2 {
                                                    v25 = v25[1:len(v25)]
                                                }
                                            } else {
                                                if len(v25) >= 2 {
                                                    v25 = v25[1:len(v25)]
                                                }
                                            }
                                        } else {
                                            v46 := GenS3{F0: "hello", F1: 382}
                                            v47 := GenS4{F0: 1000, F1: 255, F2: (^(v27))}
                                            var v48 uint8 = 128
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                        }
                                    } else {
                                        var v49 uint16 = 2
                                        switch v3 {
                                        case "foo bar":
                                            fmt.Println(v26)
                                        case "":
                                            v50 := GenS1{F0: 118, F1: ((v3 + "hello") + ("foo bar" + v43))}
                                            _ = v50
                                        default:
                                            if len(v25) > 0 {
                                                v25[0] = 32768
                                            }
                                        }
                                        _ = v49
                                    }
                                    _ = fi42
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                }
                            }
                        }
                        _ = len(v25)
                    }
                } else {
                    if v29 {
                        v2 = 0.0
                        v27 += fi28
                        v26 = v26
                    } else {
                        if true {
                            var v51 int16 = 1000
                            _ = v51
                        } else {
                            var v52 bool = v29
                            _ = v52
                        }
                    }
                    v23 = 4294967295
                }
                switch 2147483647 {
                case 42:
                    var v53 int8 = 127
                    _ = v53
                default:
                    var v54 uint32 = (v1 + 65535)
                    _ = v54
                }
            }
        }
        _ = fi28
        _ = v29
    }
    var v55 uint64 = v23
    switch v3 {
    case "\n\t":
        for fi56 := 0; fi56 < 8; fi56++ {
            fmt.Println(v1)
            var v57 uint32 = (v1 + (v1 << ((42) & 63)))
            _ = fi56
            _ = v57
        }
    case "foo bar":
        if len(v25) > 0 {
            v25[0] = -227
        }
    default:
        v9 = (-(v9))
    }
    fmt.Println(v27, v23)
    for fi58 := 0; fi58 < 6; fi58++ {
        if false {
            v59 := GenS3{F0: v3, F1: 126}
            _ = v59
        } else {
            var v60 float32 = v2
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            _ = v60
        }
        _ = fi58
    }
    switch "foo bar" {
    case "a":
        if false {
            if true {
                v23 += 0
                for fi61 := 0; fi61 < 19; fi61++ {
                    for fi62 := 0; fi62 < 7; fi62++ {
                        fmt.Println(fi62, v26)
                        _ = fi62
                    }
                    _ = fi61
                }
                var v63 int64 = v24
                _ = v63
            }
        }
    case "\n\t":
        for fi64 := 0; fi64 < 2; fi64++ {
            if (true || (false || false)) {
                var v65 float64 = 0.0
                switch "a" {
                case "":
                    v9 -= (v9 & (v9 + (^(v9))))
                }
                var v66 int16 = 1
                _ = v65
                _ = v66
            } else {
                fmt.Println(v24)
                v1 += (v26 / (((^(v26))) | 1))
                switch fi64 {
                case 9223372036854775807:
                    var v67 int32 = v9
                    _ = v67
                case 4294967295:
                    fmt.Println(v9, v27)
                case -799:
                    for fi68 := 0; fi68 < 4; fi68++ {
                        v69 := GenS1{F0: 255, F1: v3}
                        _ = fi68
                        _ = v69
                    }
                default:
                    var v70 uint8 = 255
                    _ = v70
                }
            }
            var v71 float32 = (v2 / (1.5 + v2))
            if (fi64 == (v27 % ((v27) | 1))) {
                for fi72 := 0; fi72 < 3; fi72++ {
                    v1 = (v1 * v1)
                    _ = fi72
                }
                switch fi64 {
                case -737:
                    for fi73 := 0; fi73 < 2; fi73++ {
                        var v74 uint8 = 169
                        if true {
                            if len(v25) > 0 {
                                v25[0] = (fi64 % (((fi64 / ((2147483647) | 1))) | 1))
                            }
                        }
                        var v75 int32 = v9
                        var v76 string = v3
                        switch fi73 {
                        case -242:
                            if (!((32767 == 922) && true)) {
                                v77 := GenS5{F0: 255, F1: v3}
                                var v78 int8 = 0
                                _ = v77
                                _ = v78
                            }
                        default:
                            fmt.Println(v27, v75)
                        }
                        v79 := GenS6{F0: (v26 * 4294967295), F1: 995, F2: 39}
                        _ = fi73
                        _ = v74
                        _ = v75
                        _ = v76
                        _ = v79
                    }
                case 2147483647:
                    for fi80 := 0; fi80 < 16; fi80++ {
                        var v81 bool = ((v23 / (((v23 ^ 712)) | 1)) < v23)
                        _ = fi80
                        _ = v81
                    }
                case 32768:
                    v55 = (^(v23))
                default:
                    v24 = v24
                }
                if true {
                    v82 := GenS7{F0: (v55 - (v23 & (^(v55)))), F1: 1000, F2: 0}
                    v55 = v23
                    _ = v82
                }
            }
            _ = fi64
            _ = v71
        }
    default:
        var v83 float32 = ((v2 / (v2 - (float32((1e10 - -2.25)) / v2))) * ((1e10 - v2) - v2))
        _ = v83
    }
    switch v3 {
    case "hello":
        for fi84 := 0; fi84 < 1; fi84++ {
            fmt.Println(v24, v24, v2)
            _ = fi84
        }
    case "a":
        var v85 float64 = ((1.5 / 3.14159) + (float64(v2) - 3.14159))
        _ = v85
    case "\n\t":
        v86 := GenS8{F0: 127, F1: 42, F2: 42}
        _ = v86
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v55
    _ = v9
}
