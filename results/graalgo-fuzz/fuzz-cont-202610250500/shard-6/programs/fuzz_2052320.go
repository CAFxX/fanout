package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 int }
type GenS8 struct { F0 int32 }
func genFn0(p0 int32) int8 {
    v7 := GenS1{F0: -277, F1: 100}
    p0 -= (p0 / ((p0) | 1))
    v8 := []int{1, 9223372036854775807, -9223372036854775808}
    switch "a" {
    case "":
        v8 = append(v8, 32768)
    case "foo bar":
        v8 = append(v8, 128)
    default:
        if false {
            if false {
                p0 += p0
                switch int(p0) {
                case -2147483648:
                    p0 = (p0 ^ (^(p0)))
                case 714:
                    p0 = (-(p0))
                }
            } else {
                var v9 int8 = 100
                var v10 int8 = (v9 + 100)
                v11 := GenS2{F0: p0}
                _ = v10
                _ = v11
                _ = v9
            }
            var v12 uint = 855
            var v13 float64 = -2.25
            _ = v12
            _ = v13
        } else {
            if (2 >= 2) {
                if (1.5 <= float32(0.1)) {
                    if (!((-1 <= 1) || (4294967295 > -1))) {
                        p0 = -1
                        for fi14 := 0; fi14 < 15; fi14++ {
                            if (32767 != 0) {
                                v15 := GenS3{F0: true, F1: p0}
                                var v16 bool = (int64(fi14) > -9223372036854775808)
                                var v17 bool = (42 != -1)
                                _ = v15
                                _ = v16
                                _ = v17
                            } else {
                                for fi18 := 0; fi18 < 16; fi18++ {
                                    var v19 uint8 = 128
                                    _ = fi18
                                    _ = v19
                                }
                                for fi20 := 0; fi20 < 11; fi20++ {
                                    v21 := GenS4{F0: fi14, F1: "hello"}
                                    v22 := GenS4{F0: -468, F1: "a"}
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                            }
                            p0 = (p0 % (((-(p0))) | 1))
                            _ = fi14
                        }
                    }
                } else {
                    if ((!(56 != -44)) || false) {
                        _ = len(v8)
                        if len(v8) > 0 {
                            v8[0] = -688
                        }
                        v23 := GenS1{F0: (p0 / (((p0 & p0)) | 1)), F1: 42}
                        _ = v23
                    } else {
                        p0 += p0
                        p0 = -1
                    }
                    var v24 float32 = 1.5
                    _ = v24
                }
                var v25 int32 = p0
                v25 = (v25 | -1)
                _ = v25
            }
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return 89
}

func main() {
    var v1 int32 = -2147483648
    v2 := []uint32{42, 2147483647}
    var v3 uint16 = 2
    v4 := []int{}
    var v5 int32 = (v1 | v1)
    v6 := make(map[string]uint32)
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    v1 = 1
    for fi26 := 0; fi26 < 15; fi26++ {
        v27 := GenS3{F0: false, F1: v5}
        var v28 float32 = 1.5
        _ = fi26
        _ = v27
        _ = v28
    }
    v1 = v1
    var v29 uint8 = 127
    if (false || true) {
        v30 := GenS4{F0: 990, F1: "hello"}
        _ = v30
    } else {
        v31 := GenS4{F0: 69, F1: ("" + "\n\t")}
        v32 := GenS4{F0: 32767, F1: "hello"}
        v5 -= v1
        _ = v31
        _ = v32
    }
    switch 32768 {
    case 256:
        if false {
            v5 -= (^(v5))
        } else {
            v33 := GenS2{F0: 42}
            var v34 int8 = -128
            _ = v33
            _ = v34
        }
    case 1:
        switch "hello" {
        case "foo bar":
            for fi35 := 0; fi35 < 15; fi35++ {
                v5 = (v5 & v5)
                switch (fi35 % (((-(fi35))) | 1)) {
                case 127:
                    v29 = v29
                case -1:
                    switch fi35 {
                    case -300:
                        switch "a" {
                        case "hello":
                            var v36 int64 = 0
                            _ = v36
                        case "":
                            v3 += 32767
                        default:
                            if (v29 == (^(v29))) {
                                v2 = append(v2, 1)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                    case -9223372036854775808:
                        if false {
                            var v37 int64 = 42
                            v38 := GenS3{F0: (fi35 < int(v37)), F1: -2147483648}
                            for fi39 := 0; fi39 < 6; fi39++ {
                                v4 = append(v4, (fi35 + (-(fi39))))
                                v3 += v3
                                _ = fi39
                            }
                            _ = v37
                            _ = v38
                        }
                    case 128:
                        fmt.Println(v29, v1)
                    default:
                        v5 = -326
                    }
                default:
                    v40 := GenS5{F0: 42}
                    _ = v40
                }
                if (42 == 42) {
                    var v41 uint32 = 959
                    var v42 int64 = -9223372036854775808
                    _ = v41
                    _ = v42
                }
                _ = fi35
            }
        case "hello":
            for fi43 := 0; fi43 < 14; fi43++ {
                fmt.Println(v1, v1)
                for fi44 := 0; fi44 < 20; fi44++ {
                    v45 := GenS6{F0: "a"}
                    _ = fi44
                    _ = v45
                }
                _ = fi43
            }
        case "\n\t":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    case -958:
        if false {
            if ((!(false || (!((!((!(!(((v3 << ((924) & 63)) != v3) || true))) && true)) && true)))) && (50 > 100)) {
                for fi46 := 0; fi46 < 3; fi46++ {
                    if false {
                        for fi47 := 0; fi47 < 3; fi47++ {
                            v6["foo bar"] = 2147483648
                            v48 := GenS7{F0: 1, F1: fi46}
                            if (!(false || false)) {
                                if true {
                                    v3 = v3
                                }
                                fmt.Println(v1, fi46)
                                _ = len(v6)
                            }
                            _ = fi47
                            _ = v48
                        }
                    }
                    _ = fi46
                }
                v3 = (v3 << ((276) & 63))
            }
            v3 = (v3 / ((v3) | 1))
        }
    }
    for fi49 := 0; fi49 < 20; fi49++ {
        switch 0 {
        case 795:
            var v50 uint64 = 365
            _ = v50
        }
        v6["a"] = 1
        var v51 float32 = (float32(float64(float32((1e10 / 3.14159)))) * 0.1)
        _ = fi49
        _ = v51
    }
    if (!true) {
        switch -9223372036854775808 {
        case 127:
            v5 = v1
        case 1:
            switch "a" {
            case "hello":
                fmt.Println(v1)
            case "a":
                if true {
                    if (323 <= 2) {
                        v29 -= v29
                        for fi52 := 0; fi52 < 12; fi52++ {
                            var v53 bool = true
                            v1 = (v1 & (v1 ^ (^(v5))))
                            var v54 float32 = (-2.25 + ((0.1 * float32((float64(1.5) * (-2.25 / 3.14159)))) - (float32(-2.25) / 3.14159)))
                            _ = fi52
                            _ = v53
                            _ = v54
                        }
                    } else {
                        v1 = v1
                        switch ("hello" + "a") {
                        case "foo bar":
                            var v55 int16 = 76
                            _ = v55
                        case "hello":
                            var v56 int = -2147483648
                            _ = v56
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    var v57 uint8 = v29
                    v58 := GenS3{F0: true, F1: v5}
                    _ = v57
                    _ = v58
                }
            case "foo bar":
                v29 += v29
            default:
                v59 := GenS8{F0: (v5 + (v1 / ((v5) | 1)))}
                _ = v59
            }
        case -383:
            v3 = v3
        default:
            if false {
                v60 := GenS8{F0: v1}
                _ = v60
            }
        }
        for fi61 := 0; fi61 < 14; fi61++ {
            v3 = v3
            _ = fi61
        }
        if (2 != 914) {
            var v62 int16 = 1
            _ = v62
        }
    }
    for fi63 := 0; fi63 < 14; fi63++ {
        v29 = (v29 + v29)
        var v64 string = ("" + ("hello" + "foo bar"))
        _ = fi63
        _ = v64
    }
    switch ((("" + "hello") + "") + "\n\t") {
    case "foo bar":
        var v65 int = 1
        _ = v65
    case "":
        var v66 int32 = -969
        _ = v66
    case "\n\t":
        v67 := GenS6{F0: ("a" + "\n\t")}
        _ = v67
    default:
        var v68 float32 = 1.5
        _ = v68
    }
    if (!((1e10 - ((float64(-2.25) + float64((1e10 / ((1.5 - 3.14159) + 1.5)))) * 1e10)) <= 0.1)) {
        for fi69 := 0; fi69 < 12; fi69++ {
            if (true || (4294967295 > 69)) {
                _ = len(v4)
            } else {
                if (!true) {
                    v70 := GenS2{F0: (v5 ^ 42)}
                    _ = v70
                }
                v1 += (v5 % ((v1) | 1))
            }
            _ = fi69
        }
    }
    if true {
        var v71 uint64 = 156
        for fi72 := 0; fi72 < 20; fi72++ {
            v1 = int32(fi72)
            _ = fi72
        }
        var v73 uint64 = v71
        _ = v71
        _ = v73
    } else {
        fmt.Println(v1, v1)
        fmt.Println(v29, v1, v3)
    }
    v1 -= v1
    if true {
        var v74 bool = true
        v1 = 42
        _ = v74
    } else {
        v75 := GenS6{F0: (("foo bar" + "a") + "\n\t")}
        var v76 uint16 = v3
        v3 = 32767
        _ = v75
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
