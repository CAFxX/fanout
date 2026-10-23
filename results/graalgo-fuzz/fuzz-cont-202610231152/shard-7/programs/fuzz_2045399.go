package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 int8; F1 int8; F2 int8 }
type GenS6 struct { F0 int16; F1 uint; F2 int64 }
type GenS7 struct { F0 uint16; F1 bool }

func main() {
    v1 := []string{"", (("\n\t" + "hello") + "a"), "a", ("" + "foo bar")}
    v2 := []uint32{494, 892, 4294967295, 88}
    v3 := []int{2147483647, -208, -1}
    v4 := make(map[int]int)
    var v5 string = ("" + (("a" + "a") + "foo bar"))
    v6 := make(map[string]string)
    if (((1.5 * 0.0) * 1.5) != 3.14159) {
        var v7 uint16 = 65535
        if true {
            v5 = "hello"
            var v8 float32 = 0.1
            v9 := GenS1{F0: 106, F1: -32768, F2: 2147483647}
            _ = v8
            _ = v9
        } else {
            delete(v6, v5)
            v7 = (v7 >> ((759) & 63))
        }
        var v10 int8 = -1
        _ = v10
        _ = v7
    }
    var v11 int32 = 0
    for fi12 := 0; fi12 < 18; fi12++ {
        if (286 <= 9223372036854775807) {
            v13 := GenS1{F0: 2, F1: 1000, F2: v11}
            var v14 uint16 = 255
            _ = v13
            _ = v14
        }
        if true {
            v15 := GenS2{F0: 42}
            var v16 uint8 = 42
            _ = v15
            _ = v16
        } else {
            var v17 int32 = (-(v11))
            if (1.5 >= float64(float32((0.0 - (0.0 / ((-2.25 + 1e10) + 0.1)))))) {
                if ((fi12 << ((2) & 63)) != fi12) {
                    fmt.Println(v5, v17)
                }
            }
            _ = v17
        }
        _ = fi12
    }
    v11 += v11
    var v18 uint = 0
    v19 := GenS2{F0: 42}
    v18 = (v18 / ((v18) | 1))
    switch 127 {
    case 42:
        fmt.Println(v5, v11)
    case 127:
        switch 128 {
        case 32767:
            v5 = (v5 + "hello")
        case -319:
            if (313 <= 18446744073709551615) {
                v11 += (v11 | (v11 | (v11 & (v11 << ((v18) & 63)))))
                fmt.Println(v11, v18)
                if true {
                    var v20 int16 = -425
                    v18 -= 18446744073709551615
                    var v21 uint8 = 2
                    _ = v20
                    _ = v21
                }
            }
        case 127:
            v22 := GenS2{F0: 0}
            _ = v22
        default:
            for fi23 := 0; fi23 < 8; fi23++ {
                if (((3.14159 != 0.0) && false) && true) {
                    if (!false) {
                        v24 := v4[0]
                        if (v5 <= v5) {
                            if false {
                                for fi25 := 0; fi25 < 5; fi25++ {
                                    var v26 int32 = v11
                                    v27 := GenS1{F0: 798, F1: -32768, F2: v26}
                                    var v28 int16 = 223
                                    v6[v5] = v5
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            } else {
                                for fi29 := 0; fi29 < 12; fi29++ {
                                    v5 = v5
                                    var v30 uint8 = 216
                                    var v31 int64 = 1
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                                for fi32 := 0; fi32 < 20; fi32++ {
                                    delete(v4, (fi23 ^ (fi32 << ((42) & 63))))
                                    _ = fi32
                                }
                                v5 = "a"
                            }
                            v33 := GenS3{F0: 42, F1: 42}
                            v34 := GenS3{F0: 2147483647, F1: 32767}
                            _ = v33
                            _ = v34
                        } else {
                            fmt.Println(v24, v18)
                        }
                        _ = v24
                    }
                }
                if len(v3) > 0 {
                    v3[0] = (fi23 ^ fi23)
                }
                _ = fi23
            }
        }
    default:
        for fi35 := 0; fi35 < 6; fi35++ {
            var v36 bool = false
            _ = len(v3)
            if v36 {
                if true {
                    if v36 {
                        v3 = append(v3, (fi35 + fi35))
                        var v37 int16 = 1
                        _ = v37
                    } else {
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                    }
                    var v38 float64 = 0.1
                    if ((v11 << ((v18) & 63)) != v11) {
                        for fi39 := 0; fi39 < 15; fi39++ {
                            if v36 {
                                if v36 {
                                    v38 += v38
                                }
                            }
                            _ = fi39
                        }
                    } else {
                        switch fi35 {
                        case 9223372036854775807:
                            v6[((v5 + "hello") + "")] = (v5 + v5)
                        case 42:
                            v18 += v18
                        }
                        if v36 {
                            v40 := GenS2{F0: 4294967295}
                            _ = v40
                        } else {
                            fmt.Println(v5, v36, v36)
                        }
                    }
                    _ = v38
                } else {
                    switch (fi35 / ((fi35) | 1)) {
                    case -9223372036854775808:
                        if v36 {
                            if (!v36) {
                                if (("hello" != v5) && true) {
                                    if true {
                                        switch v5 {
                                        case "\n\t":
                                            v18 += v18
                                        default:
                                            switch fi35 {
                                            case 446:
                                                fmt.Println(fi35)
                                            case -9223372036854775808:
                                                if ((!true) || false) {
                                                    var v41 int32 = 0
                                                    for fi42 := 0; fi42 < 14; fi42++ {
                                                        fmt.Println(v36)
                                                        _ = fi42
                                                    }
                                                    delete(v6, (v5 + v5))
                                                    _ = v41
                                                }
                                            case 0:
                                                for fi43 := 0; fi43 < 20; fi43++ {
                                                    fmt.Println(v5)
                                                    var v44 int32 = v11
                                                    switch (((v5 + v5) + "\n\t") + "") {
                                                    case "":
                                                        var v45 uint32 = 0
                                                        _ = v45
                                                    case "hello":
                                                        var v46 uint = v18
                                                        _ = v46
                                                    case "a":
                                                        v36 = v36
                                                    default:
                                                        v36 = false
                                                    }
                                                    _ = fi43
                                                    _ = v44
                                                }
                                            default:
                                                for fi47 := 0; fi47 < 16; fi47++ {
                                                    if v36 {
                                                        fmt.Println(v36, v5, v18)
                                                        fmt.Println(v11)
                                                    }
                                                    _ = fi47
                                                }
                                            }
                                        }
                                    } else {
                                        _ = len(v3)
                                    }
                                    var v48 uint = (v18 % ((v18) | 1))
                                    _ = v48
                                } else {
                                    fmt.Println(v18)
                                    if (127 >= 128) {
                                        v11 = 42
                                        switch (fi35 + fi35) {
                                        case -2147483648:
                                            if v36 {
                                                for fi49 := 0; fi49 < 9; fi49++ {
                                                    var v50 int16 = 1000
                                                    fmt.Println(fi35, v18)
                                                    _ = fi49
                                                    _ = v50
                                                }
                                                var v51 int64 = -394
                                                _ = v51
                                            }
                                        }
                                    } else {
                                        for fi52 := 0; fi52 < 1; fi52++ {
                                            var v53 uint8 = 1
                                            _ = fi52
                                            _ = v53
                                        }
                                    }
                                }
                            }
                        }
                    case -1:
                        _ = len(v2)
                    default:
                        v54 := GenS2{F0: 0}
                        _ = v54
                    }
                    var v55 int = fi35
                    _ = v55
                }
                _ = len(v4)
            } else {
                if (!true) {
                    v11 = 1
                    if ((1000 != -32768) || (251 <= 4294967295)) {
                        var v56 int8 = 0
                        if (42 < 264) {
                            if (v36 || ((v18 / ((v18) | 1)) <= 413)) {
                                fmt.Println(v18, v11)
                                fmt.Println(v18)
                            }
                            v57 := GenS4{F0: -245, F1: v56}
                            var v58 uint = (v18 + (v18 * v18))
                            _ = v57
                            _ = v58
                        } else {
                            v5 = "hello"
                        }
                        var v59 bool = true
                        _ = v56
                        _ = v59
                    }
                    var v60 int16 = 4
                    _ = v60
                } else {
                    var v61 uint16 = 644
                    v61 += v61
                    _ = v61
                }
                for fi62 := 0; fi62 < 14; fi62++ {
                    fmt.Println(v11, fi62)
                    _ = fi62
                }
            }
            _ = fi35
            _ = v36
        }
    }
    v11 -= v11
    if (798 > 32767) {
        var v63 int16 = -1
        _ = v63
    } else {
        _ = len(v4)
    }
    if true {
        var v64 uint32 = 2
        if false {
            v11 += (v11 ^ v11)
        } else {
            delete(v6, v5)
            v65 := GenS3{F0: v64, F1: 65535}
            v2 = append(v2, (v64 ^ (^(v64))))
            _ = v65
        }
        _ = v64
    }
    if true {
        v11 += (^(v11))
        var v66 float64 = 0.0
        _ = v66
    }
    for fi67 := 0; fi67 < 14; fi67++ {
        switch "foo bar" {
        case "hello":
            v5 = v5
        case "a":
            fmt.Println(v5)
        }
        _ = fi67
    }
    if true {
        fmt.Println(v11, v5)
    } else {
        if false {
            v11 -= -1
            if true {
                if false {
                    var v68 bool = false
                    _ = v68
                } else {
                    for fi69 := 0; fi69 < 4; fi69++ {
                        v18 -= (v18 >> ((v18) & 63))
                        _ = fi69
                    }
                    if (uint16(v18) < 0) {
                        var v70 int32 = v11
                        var v71 uint8 = 127
                        _ = v70
                        _ = v71
                    }
                }
            }
        }
        v11 -= v11
        switch 128 {
        case -28:
            fmt.Println(v11)
        case -537:
            if true {
                if (true || false) {
                    var v72 uint8 = 50
                    var v73 uint64 = 4294967295
                    _ = v72
                    _ = v73
                } else {
                    v74 := GenS4{F0: 0, F1: 30}
                    _ = len(v6)
                    _ = v74
                }
            } else {
                for fi75 := 0; fi75 < 2; fi75++ {
                    switch fi75 {
                    case 2147483647:
                        var v76 uint64 = 42
                        _ = v76
                    case 4294967295:
                        v77 := GenS5{F0: -128, F1: -128, F2: 1}
                        _ = v77
                    default:
                        v11 += int32(fi75)
                    }
                    for fi78 := 0; fi78 < 8; fi78++ {
                        switch v5 {
                        case "":
                            switch "foo bar" {
                            case "a":
                                v4[(^(fi75))] = fi75
                            default:
                                var v79 string = v5
                                _ = v79
                            }
                            fmt.Println(fi75, v5, v11)
                            switch -144 {
                            case 2147483647:
                                if (false || ((fi75 - -9223372036854775808) >= fi75)) {
                                    v5 = ((v5 + v5) + ((v5 + ("hello" + "hello")) + ((v5 + "\n\t") + (((v5 + (("" + v5) + v5)) + (("hello" + ("\n\t" + (v5 + "hello"))) + v5)) + ("" + "\n\t")))))
                                    var v80 uint64 = 4294967295
                                    _ = v80
                                }
                            }
                        case "foo bar":
                            fmt.Println(fi75, v11)
                        case "a":
                            var v81 int32 = (v11 / ((v11) | 1))
                            _ = v81
                        default:
                            fmt.Println(fi78)
                        }
                        if (2147483647 >= 373) {
                            v2 = append(v2, 844)
                        }
                        v18 += 4294967295
                        _ = fi78
                    }
                    _ = fi75
                }
                for fi82 := 0; fi82 < 6; fi82++ {
                    fmt.Println(fi82)
                    switch (fi82 * (-(fi82))) {
                    case 127:
                        var v83 int64 = 1
                        _ = v83
                    case -2147483648:
                        var v84 int32 = (^(v11))
                        _ = v84
                    case 2147483647:
                        v85 := GenS6{F0: 0, F1: (v18 & (^(v18))), F2: -4}
                        _ = v85
                    default:
                        if false {
                            v86 := GenS7{F0: 255, F1: false}
                            _ = v86
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v18 = v18
                        }
                    }
                    _ = fi82
                }
            }
        default:
            var v87 bool = (162 > 1)
            _ = v87
        }
    }
    fmt.Println(v11, v5)
    var v88 float32 = (0.0 / 0.1)
    delete(v6, ("" + (v5 + "")))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v88:", v88)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v88
}
