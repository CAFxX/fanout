package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 uint8; F2 uint32 }

func main() {
    var v1 int8 = -1
    v2 := []string{}
    v3 := make(map[int]uint32)
    var v4 int = -460
    var v5 int8 = int8(v4)
    v6 := make(map[string]bool)
    fmt.Println(v4)
    switch v4 {
    case 0:
        var v7 float32 = ((float32(float64(1e10)) + float32(1e10)) + -2.25)
        _ = v7
    case 32768:
        v8 := v3[1]
        _ = v8
    case 127:
        fmt.Println(v1, v5)
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            v5 += (v5 << ((18446744073709551615) & 63))
            switch (fi10 / ((v4) | 1)) {
            case 9223372036854775807:
                if true {
                    if false {
                        fmt.Println(fi9, v4)
                        v11 := v6[""]
                        _ = v11
                    }
                    v1 -= 0
                    v1 = (^(v5))
                } else {
                    v12 := GenS1{F0: 2, F1: -809, F2: 65535}
                    _ = v12
                }
            }
            _ = fi10
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi13 := 0; fi13 < 15; fi13++ {
            switch "foo bar" {
            case "hello":
                var v14 float32 = 0.0
                _ = v14
            }
            switch (fi9 - 128) {
            case -2147483648:
                fmt.Println(v4, fi13, v1)
            case 589:
                var v15 int8 = 127
                _ = v15
            case 127:
                v1 = v1
            default:
                v16 := GenS1{F0: 42, F1: -504, F2: 18446744073709551615}
                _ = v16
            }
            _ = fi13
        }
        _ = fi9
    }
    v4 = v4
    var v17 bool = true
    for fi18 := 0; fi18 < 11; fi18++ {
        v19 := GenS1{F0: 2147483648, F1: 333, F2: 829}
        if (537 <= -2147483648) {
            for fi20 := 0; fi20 < 2; fi20++ {
                if false {
                    v17 = true
                    v1 -= (v5 & (v1 * 42))
                    v21 := GenS1{F0: 2147483647, F1: 0, F2: 2}
                    var v22 bool = v17
                    v23 := v6[""]
                    var v24 uint32 = 2147483647
                    fmt.Println(v23)
                    _ = v21
                    _ = v22
                    _ = v23
                    _ = v24
                }
                _ = fi20
            }
        } else {
            fmt.Println(v5, v17, v5)
            for fi25 := 0; fi25 < 16; fi25++ {
                if (255 != 5) {
                    v1 -= (v5 << ((42) & 63))
                    v5 += (v1 % ((v5) | 1))
                } else {
                    v1 += 127
                    var v26 uint32 = 596
                    _ = v26
                }
                if (!(!((v1 <= v1) && (76 < 160)))) {
                    var v27 bool = (-594 == 42)
                    v28 := GenS1{F0: 2, F1: 322, F2: 786}
                    if (v17 && v17) {
                        v1 -= v1
                    }
                    fmt.Println(fi18)
                    v29 := GenS1{F0: 359, F1: 0, F2: 4294967295}
                    var v30 bool = v27
                    _ = v27
                    _ = v28
                    _ = v29
                    _ = v30
                } else {
                    switch (v4 * (fi25 & v4)) {
                    case -442:
                        if (1.5 == -2.25) {
                            v31 := GenS2{F0: 65535, F1: 1}
                            fmt.Println(fi18, fi25, v1)
                            if (2 > 0) {
                                fmt.Println(v4)
                            } else {
                                v32 := GenS2{F0: 2, F1: 1}
                                v5 += (v1 - (v5 >> ((65535) & 63)))
                                _ = v32
                            }
                            if (!false) {
                                switch "\n\t" {
                                case "\n\t":
                                    var v33 bool = (!true)
                                    _ = v33
                                case "a":
                                    v1 -= 113
                                default:
                                    var v34 float64 = 1e10
                                    _ = v34
                                }
                                v35 := GenS2{F0: 1, F1: 891}
                                v36 := GenS2{F0: 2, F1: 600}
                                _ = v35
                                _ = v36
                            } else {
                                var v37 int = fi18
                                _ = v37
                            }
                            v17 = (75 >= 42)
                            _ = v31
                        } else {
                            v6["foo bar"] = v17
                            fmt.Println(v17, v4, fi25)
                            fmt.Println(v1)
                        }
                    case -2147483648:
                        v38 := GenS2{F0: 1, F1: 952}
                        _ = v38
                    case -9223372036854775808:
                        var v39 string = ("\n\t" + "a")
                        _ = v39
                    default:
                        if (-1 > 1) {
                            v40 := GenS2{F0: 32767, F1: 42}
                            var v41 uint8 = 255
                            var v42 uint8 = 255
                            fmt.Println(fi25)
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                    }
                }
                _ = fi25
            }
        }
        fmt.Println(v5, v4, v4)
        _ = fi18
        _ = v19
    }
    var v43 uint32 = 42
    fmt.Println(v4)
    v44 := GenS1{F0: 231, F1: 222, F2: 255}
    switch "" {
    case "foo bar":
        v1 -= v1
    default:
        for fi45 := 0; fi45 < 12; fi45++ {
            v3[(v4 & 473)] = v43
            _ = fi45
        }
    }
    v46 := v6[""]
    if v17 {
        if true {
            if (v46 || (float32((3.14159 - 1e10)) < 0.1)) {
                fmt.Println(v17, v4, v17)
                if (2 == 0) {
                    for fi47 := 0; fi47 < 19; fi47++ {
                        var v48 uint32 = (^(v43))
                        _ = fi47
                        _ = v48
                    }
                    if true {
                        v49 := v6["foo bar"]
                        v17 = (!(-404 >= -539))
                        v4 -= v4
                        _ = v49
                    }
                } else {
                    switch 32768 {
                    case 42:
                        var v50 string = ""
                        _ = v50
                    case 128:
                        if (v17 && (!true)) {
                            var v51 float64 = 0.0
                            v4 -= v4
                            _ = v51
                        } else {
                            v46 = v17
                        }
                    }
                    v1 += (v5 + (v5 + 100))
                    var v52 int16 = 538
                    _ = v52
                }
            }
        } else {
            if ((v46 && (v43 == (v43 >> ((326) & 63)))) && false) {
                if v46 {
                    fmt.Println(v43, v4, v17)
                } else {
                    v5 = (v5 * v5)
                    v53 := v3[9223372036854775807]
                    _ = v53
                }
            } else {
                if true {
                    v43 = v43
                    if (("" + "a") >= "foo bar") {
                        var v54 uint8 = 127
                        fmt.Println(v46, v54, v1)
                        _ = v54
                    }
                    v4 = v4
                } else {
                    if (112 != 70) {
                        fmt.Println(v4, v46)
                        if v17 {
                            var v55 bool = v17
                            delete(v3, 32767)
                            _ = v55
                        }
                    }
                    if (v46 && true) {
                        v3[v4] = v43
                    } else {
                        var v56 uint32 = v43
                        if (2147483647 < -614) {
                            var v57 int = v4
                            var v58 int = (v4 % ((42) | 1))
                            v59 := GenS3{F0: 2147483647}
                            _ = v57
                            _ = v58
                            _ = v59
                        } else {
                            for fi60 := 0; fi60 < 10; fi60++ {
                                fmt.Println(v4)
                                _ = fi60
                            }
                        }
                        _ = v56
                    }
                }
                v61 := GenS3{F0: 313}
                _ = v61
            }
        }
        var v62 int32 = 911
        v4 += (v4 ^ v4)
        _ = v62
    } else {
        for fi63 := 0; fi63 < 8; fi63++ {
            _ = len(v6)
            v64 := GenS4{F0: 860}
            v46 = (int(v43) == (v4 >> ((0) & 63)))
            _ = fi63
            _ = v64
        }
        switch "\n\t" {
        case "":
            v5 = (v5 & 72)
        case "a":
            if v46 {
                v65 := GenS5{F0: 738, F1: 128, F2: v43}
                var v66 float64 = ((float64(1e10) + (1e10 + -2.25)) + -2.25)
                for fi67 := 0; fi67 < 13; fi67++ {
                    v17 = v17
                    switch v4 {
                    case 0:
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                    default:
                        switch "" {
                        case "hello":
                            fmt.Println(v5)
                        case "a":
                            fmt.Println(v43, v66, v4)
                        }
                    }
                    if (42 > 730) {
                        if len(v2) > 0 {
                            v2[0] = ("" + "hello")
                        }
                        var v68 string = ("hello" + "a")
                        fmt.Println(v4, v4)
                        _ = v68
                    }
                    _ = fi67
                }
                _ = v65
                _ = v66
            } else {
                fmt.Println(v46, v1)
                fmt.Println(v46)
            }
        case "\n\t":
            for fi69 := 0; fi69 < 2; fi69++ {
                if true {
                    v1 = v5
                }
                for fi70 := 0; fi70 < 18; fi70++ {
                    fmt.Println(fi69)
                    var v71 int32 = -1
                    v46 = v17
                    v72 := GenS5{F0: 946, F1: 127, F2: v43}
                    v73 := GenS5{F0: 1, F1: 0, F2: v43}
                    v6["foo bar"] = (0.1 != 0.1)
                    switch fi70 {
                    case 0:
                        switch ("foo bar" + "foo bar") {
                        case "a":
                            var v74 bool = (3.14159 < 0.0)
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                            _ = v74
                        case "foo bar":
                            v5 = (v5 - (v1 % ((1) | 1)))
                        }
                    case 1:
                        v46 = v46
                    }
                    _ = fi70
                    _ = v71
                    _ = v72
                    _ = v73
                }
                _ = fi69
            }
        default:
            if true {
                v1 = v1
                var v75 float64 = 0.1
                _ = v75
            }
        }
    }
    var v76 uint = 129
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v76:", v76)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v46
    _ = v5
    _ = v6
    _ = v76
}
