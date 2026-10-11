package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 uint32; F2 uint }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 uint32 }

func main() {
    var v1 int32 = 0
    var v2 int16 = 1000
    v3 := make(map[uint32]bool)
    v4 := []uint32{}
    var v5 bool = false
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 float64 = 1e10
        _ = fi6
        _ = v7
    }
    v5 = (4294967295 < 42)
    switch 32768 {
    case 127:
        var v8 string = "\n\t"
        _ = v8
    case 195:
        for fi9 := 0; fi9 < 17; fi9++ {
            v5 = true
            v5 = (!false)
            for fi10 := 0; fi10 < 1; fi10++ {
                if (("\n\t" >= "a") && false) {
                    var v11 uint = 973
                    _ = v11
                } else {
                    v12 := GenS1{F0: 9223372036854775807}
                    var v13 bool = v5
                    v14 := GenS2{F0: 881}
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = fi10
            }
            _ = fi9
        }
    default:
        for fi15 := 0; fi15 < 16; fi15++ {
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 string = "a"
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 8; fi18++ {
                switch "a" {
                case "\n\t":
                    switch (("foo bar" + ("foo bar" + "hello")) + "foo bar") {
                    case "":
                        v5 = (0 == 57)
                    case "a":
                        v1 = (v1 % ((2147483647) | 1))
                    }
                case "hello":
                    v5 = (false && (!(497 == 2147483648)))
                case "":
                    if false {
                        if (0.1 == ((1.5 + 1.5) / -2.25)) {
                            v19 := GenS3{F0: v1, F1: 2}
                            if false {
                                var v20 int = fi15
                                switch "a" {
                                case "\n\t":
                                    _ = len(v3)
                                case "a":
                                    if (v1 > v1) {
                                        var v21 int32 = v1
                                        if true {
                                            if (-2.25 != float64(3.14159)) {
                                                var v22 bool = false
                                                _ = v22
                                            } else {
                                                var v23 uint32 = 1
                                                v1 += (v21 % (((v21 * -1)) | 1))
                                                _ = v23
                                            }
                                        } else {
                                            v2 = (v2 << ((97) & 63))
                                            if (305 < 1) {
                                                if (!v5) {
                                                    v1 = 403
                                                    v24 := GenS2{F0: 97}
                                                    v5 = true
                                                    _ = v24
                                                }
                                                v25 := GenS4{F0: int32(v2), F1: 2}
                                                _ = v25
                                            } else {
                                                v2 += 0
                                            }
                                            var v26 int64 = 9223372036854775807
                                            _ = v26
                                        }
                                        var v27 uint64 = 590
                                        _ = v21
                                        _ = v27
                                    } else {
                                        var v28 float32 = 0.1
                                        _ = v28
                                    }
                                    v20 += -9223372036854775808
                                }
                                v4 = append(v4, 2)
                                _ = v20
                            }
                            v1 = v1
                            v2 = (v2 >> ((343) & 63))
                            _ = v19
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v29 uint64 = 2
                            _ = v29
                        }
                        delete(v3, 4294967295)
                        var v30 uint = 852
                        _ = v30
                    } else {
                        v31 := GenS4{F0: 741, F1: 610}
                        _ = v31
                    }
                }
                switch (fi18 & (fi15 | (^(fi15)))) {
                case 32767:
                    fmt.Println(v2)
                }
                fmt.Println(v2)
                var v32 bool = ((^(v2)) == (v2 - (v2 % ((v2) | 1))))
                if (!(!(((float32((float64((1e10 / 1.0)) / 0.1)) - 3.14159) - (3.14159 - (1.5 / 1e10))) < 1.5))) {
                    _ = len(v3)
                } else {
                    switch (^(fi18)) {
                    case 32767:
                        v2 = (v2 + 1)
                    default:
                        var v33 uint8 = 237
                        _ = v33
                    }
                }
                v5 = (!(v32 || true))
                v5 = true
                _ = fi18
                _ = v32
            }
            switch fi15 {
            case 259:
                v2 += v2
            }
            _ = fi15
        }
    }
    _ = len(v3)
    var v34 bool = v5
    for fi35 := 0; fi35 < 10; fi35++ {
        var v36 float32 = (0.1 * 1e10)
        _ = fi35
        _ = v36
    }
    var v37 float64 = (-2.25 * 0.1)
    v37 = ((v37 + float64((1.5 / 1e10))) * v37)
    v34 = v5
    v38 := GenS5{F0: 0}
    for fi39 := 0; fi39 < 14; fi39++ {
        for fi40 := 0; fi40 < 13; fi40++ {
            var v41 int16 = (v2 / ((-1) | 1))
            var v42 float32 = 0.0
            delete(v3, 994)
            _ = fi40
            _ = v41
            _ = v42
        }
        v5 = (0 > 2)
        fmt.Println(v1)
        _ = fi39
    }
    v43 := GenS3{F0: v1, F1: 18446744073709551615}
    v4 = append(v4, 743)
    if (false || true) {
        for fi44 := 0; fi44 < 6; fi44++ {
            if v5 {
                var v45 float32 = ((float32(v37) - 1.5) / ((((0.1 * float32(v37)) / 1.5) * ((1.5 * (1.5 - -2.25)) * (0.1 - (1.5 * 1e10)))) / (1e10 + 1e10)))
                var v46 uint8 = 127
                _ = v45
                _ = v46
            }
            fmt.Println(fi44, v5, v1)
            var v47 uint8 = 99
            _ = fi44
            _ = v47
        }
        switch 4294967295 {
        case 349:
            for fi48 := 0; fi48 < 11; fi48++ {
                v3[901] = (float32((-2.25 / 1.0)) == (float32(v37) / 3.14159))
                _ = fi48
            }
        case 128:
            fmt.Println(v1)
        default:
            v4 = append(v4, 2147483647)
        }
        switch "a" {
        case "a":
            for fi49 := 0; fi49 < 7; fi49++ {
                fmt.Println(v5)
                var v50 int8 = -16
                if true {
                    if ((65535 > 18446744073709551615) || v34) {
                        var v51 uint16 = 269
                        _ = v51
                    }
                }
                _ = fi49
                _ = v50
            }
        case "\n\t":
            for fi52 := 0; fi52 < 14; fi52++ {
                switch ("\n\t" + "foo bar") {
                case "":
                    v5 = (!v5)
                case "a":
                    if v34 {
                        if (179 == -1) {
                            v1 = -74
                            for fi53 := 0; fi53 < 17; fi53++ {
                                if v34 {
                                    fmt.Println(v2, fi52)
                                } else {
                                    v2 = -1
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if (42 >= 127) {
                                        var v54 int8 = 70
                                        _ = v54
                                    }
                                }
                                _ = fi53
                            }
                        } else {
                            var v55 bool = (1 != (v2 + v2))
                            if v55 {
                                if (!v34) {
                                    if v55 {
                                        switch 32767 {
                                        case 1:
                                            if v5 {
                                                if v55 {
                                                    v55 = (-450 >= 0)
                                                    v5 = v34
                                                }
                                            }
                                        }
                                    } else {
                                        var v56 int32 = v1
                                        _ = v56
                                    }
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    switch "" {
                                    case "a":
                                        fmt.Println(v37)
                                    case "hello":
                                        v57 := GenS6{F0: 2147483647, F1: 2, F2: 867}
                                        _ = v57
                                    }
                                    v58 := GenS7{F0: 2147483647, F1: fi52}
                                    _ = v58
                                }
                                _ = len(v4)
                                v59 := GenS2{F0: 42}
                                _ = v59
                            }
                            _ = v55
                        }
                        var v60 int32 = (^(v1))
                        v34 = v34
                        _ = v60
                    } else {
                        v61 := v3[2]
                        v37 -= float64(3.14159)
                        fmt.Println(v61, v34, v5)
                        _ = v61
                    }
                }
                v2 = 0
                if (2 > 42) {
                    var v62 uint8 = 229
                    _ = v62
                } else {
                    fmt.Println(v1)
                    for fi63 := 0; fi63 < 6; fi63++ {
                        v64 := GenS8{F0: 65535}
                        var v65 int8 = 127
                        v5 = v5
                        _ = fi63
                        _ = v64
                        _ = v65
                    }
                    v37 = (v37 - (1.5 - v37))
                }
                _ = fi52
            }
        }
    }
    var v66 int32 = 1
    if len(v4) > 0 {
        v4[0] = 208
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v5
    _ = v66
}
