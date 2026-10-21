package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 int8; F2 uint64 }
type GenS7 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS10 struct { F0 uint64; F1 int32; F2 uint64 }
func genFn0(p0 int8) int16 {
    p0 -= p0
    var v6 int = 42
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 0}
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := []int{v6, fi9, fi7, 0}
            var v11 int8 = (^(p0))
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    var v12 uint8 = 99
    var v13 uint16 = 2
    _ = p0
    _ = v12
    _ = v13
    _ = v6
    return 125
}
func genFn1(p0 int8) uint {
    p0 = p0
    if true {
        for fi14 := 0; fi14 < 4; fi14++ {
            p0 = (-(p0))
            switch (fi14 & fi14) {
            case 2147483647:
                var v15 int64 = 887
                _ = v15
            default:
                for fi16 := 0; fi16 < 8; fi16++ {
                    v17 := GenS2{F0: -1, F1: false, F2: 4294967295}
                    p0 += 127
                    if true {
                        p0 = p0
                    }
                    var v18 int32 = 916
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            var v19 uint16 = 540
            _ = fi14
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        if ((112 <= 42) && (18446744073709551615 < 42)) {
            switch (fi20 ^ fi20) {
            case 4294967295:
                v21 := GenS3{F0: 42, F1: -1, F2: 4294967295}
                _ = v21
            }
            var v22 uint64 = 220
            _ = v22
        } else {
            v23 := []int{}
            p0 = (p0 & (p0 | p0))
            for fi24 := 0; fi24 < 1; fi24++ {
                if (false && (65535 <= 619)) {
                    var v25 uint = 65535
                    _ = v25
                }
                if true {
                    var v26 uint64 = 895
                    if true {
                        switch (-(fi20)) {
                        case 42:
                            v27 := GenS1{F0: 0}
                            p0 = 127
                            v26 = 0
                            p0 = p0
                            v28 := GenS3{F0: 810, F1: 545, F2: 2}
                            if (v26 < v26) {
                                var v29 uint16 = 255
                                var v30 float32 = float32((3.14159 / -2.25))
                                _ = v29
                                _ = v30
                            } else {
                                var v31 int16 = 32767
                                _ = v31
                            }
                            _ = v27
                            _ = v28
                        default:
                            var v32 uint64 = v26
                            _ = v32
                        }
                        fmt.Println(v26, fi20, v26)
                        v33 := GenS4{F0: true}
                        _ = v33
                    }
                    if false {
                        v34 := GenS3{F0: uint16(v26), F1: 9223372036854775807, F2: 65535}
                        _ = v34
                    }
                    _ = v26
                } else {
                    var v35 uint8 = 42
                    v36 := GenS3{F0: 32767, F1: 9223372036854775807, F2: 1}
                    if (((3.14159 / -2.25) < 3.14159) && true) {
                        var v37 int = fi24
                        v38 := GenS5{F0: 1}
                        _ = v37
                        _ = v38
                    }
                    if ("foo bar" <= ((("" + "\n\t") + "\n\t") + "foo bar")) {
                        var v39 int16 = 0
                        _ = v39
                    }
                    p0 = -128
                    v23 = append(v23, -435)
                    if true {
                        fmt.Println(v35)
                        v35 = v35
                    } else {
                        var v40 uint = 2
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(fi24, p0)
                        case "\n\t":
                            var v41 bool = (fi20 == (fi20 | (^(fi20))))
                            _ = v41
                        case "a":
                            fmt.Println(p0, v35)
                        }
                        _ = v40
                    }
                    switch fi20 {
                    case 127:
                        p0 = p0
                    default:
                        v42 := GenS2{F0: (p0 & (p0 - p0)), F1: (2 >= 0), F2: 0}
                        _ = v42
                    }
                    switch ("\n\t" + ("\n\t" + ("" + "\n\t"))) {
                    case "":
                        p0 = (p0 | p0)
                    case "\n\t":
                        var v43 uint32 = 12
                        _ = v43
                    }
                    p0 += (^(p0))
                    _ = v35
                    _ = v36
                }
                var v44 bool = true
                if true {
                    v45 := GenS5{F0: 42}
                    _ = v45
                }
                var v46 uint8 = 0
                if false {
                    p0 += -16
                } else {
                    v44 = (!((!v44) || true))
                    fmt.Println(v46, fi20)
                    v23 = append(v23, (^(fi20)))
                    var v47 float32 = -2.25
                    _ = v47
                }
                _ = fi24
                _ = v44
                _ = v46
            }
            _ = v23
        }
        _ = fi20
    }
    if true {
        var v48 string = (("\n\t" + "hello") + "\n\t")
        if false {
            if false {
                var v49 uint8 = 2
                v50 := []int{4294967295, 237}
                var v51 int64 = -1
                _ = v49
                _ = v50
                _ = v51
            } else {
                v48 = v48
                var v52 int = 32767
                v48 = v48
                _ = v52
            }
            v53 := GenS6{F0: false, F1: (p0 + p0), F2: 2}
            for fi54 := 0; fi54 < 8; fi54++ {
                if (!(128 != 1)) {
                    for fi55 := 0; fi55 < 12; fi55++ {
                        var v56 float32 = 1.5
                        _ = fi55
                        _ = v56
                    }
                    switch fi54 {
                    case -843:
                        v48 = "a"
                    }
                    if false {
                        v57 := []int{(fi54 | (fi54 - 862)), fi54}
                        for fi58 := 0; fi58 < 2; fi58++ {
                            if (false && ((true || false) || true)) {
                                var v59 uint32 = 0
                                var v60 int8 = p0
                                _ = v59
                                _ = v60
                            }
                            _ = fi58
                        }
                        _ = v57
                    }
                }
                v48 = ""
                _ = fi54
            }
            _ = v53
        }
        v61 := []int{127, 225, -2147483648, 32767}
        _ = v48
        _ = v61
    }
    var v62 float32 = (0.1 * 1.5)
    _ = p0
    _ = v62
    return 2
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]uint32)
    v3 := []string{"foo bar"}
    v4 := make(map[string]bool)
    var v5 int16 = (^(v1))
    v5 -= v5
    v1 += genFn0(47)
    switch -1 {
    case -1:
        if true {
            v63 := GenS7{F0: 128, F1: 0, F2: 42}
            if (-2.25 > (-2.25 + 1e10)) {
                v1 = genFn0(42)
            } else {
                for fi64 := 0; fi64 < 4; fi64++ {
                    v1 -= -118
                    v5 -= (v5 - (v1 * v1))
                    _ = fi64
                }
            }
            _ = v63
        }
    }
    v65 := GenS6{F0: (2147483648 > 65535), F1: -61, F2: 903}
    var v66 int = -400
    if true {
        v67 := GenS8{F0: -668}
        v66 = v66
        _ = v67
    } else {
        var v68 uint32 = 0
        fmt.Println(v66, v5, v68)
        switch "foo bar" {
        case "\n\t":
            _ = len(v3)
        }
        _ = v68
    }
    if true {
        if false {
            for fi69 := 0; fi69 < 18; fi69++ {
                v1 += (v1 << ((genFn1(42)) & 63))
                _ = fi69
            }
        }
    } else {
        for fi70 := 0; fi70 < 10; fi70++ {
            v71 := GenS9{F0: 32767, F1: 42, F2: 255}
            v1 = (^(v1))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi70
            _ = v71
        }
    }
    if (!(619 < 42)) {
        if (float32((0.0 + (3.14159 / (0.1 - float64((0.1 / float32((3.14159 / (-2.25 / 1.5))))))))) == 0.0) {
            var v72 int = (v66 << ((2) & 63))
            _ = v72
        }
    }
    v73 := GenS10{F0: 42, F1: -1, F2: 18446744073709551615}
    var v74 bool = true
    fmt.Println(v74)
    switch v66 {
    case 0:
        v3 = append(v3, ("" + "a"))
    }
    var v75 int16 = 1000
    v75 += v5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v73
    _ = v74
    _ = v75
}
