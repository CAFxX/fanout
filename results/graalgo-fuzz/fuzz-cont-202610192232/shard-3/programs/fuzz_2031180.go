package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint }
type GenS3 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int) uint16 {
    v6 := []int{(-(p0)), p0, p0}
    if (127 > -128) {
        for fi7 := 0; fi7 < 7; fi7++ {
            for fi8 := 0; fi8 < 4; fi8++ {
                p0 -= fi7
                _ = len(v6)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi8
            }
            p0 = p0
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS1{F0: 1}
                p0 -= fi9
                p0 = 32767
                _ = fi9
                _ = v10
            }
            _ = fi7
        }
    } else {
        fmt.Println(p0)
        var v11 int = (p0 + 9223372036854775807)
        _ = v11
    }
    _ = p0
    _ = v6
    return 65535
}

func main() {
    var v1 int16 = 19
    var v2 int64 = 1
    var v3 float32 = 0.1
    v4 := make(map[int]int)
    var v5 float32 = (-2.25 - v3)
    if (true || false) {
        var v12 int64 = (v2 % ((0) | 1))
        _ = v12
    } else {
        if (371 >= 18446744073709551615) {
            var v13 int16 = 42
            fmt.Println(v3, v2, v2)
            if (0.0 < 1.5) {
                if false {
                    for fi14 := 0; fi14 < 19; fi14++ {
                        v3 = (v5 - v3)
                        var v15 int = fi14
                        switch (("foo bar" + (((("hello" + "\n\t") + "hello") + ("hello" + "foo bar")) + "hello")) + "a") {
                        case "":
                            v16 := GenS2{F0: (^(v13)), F1: 4294967295, F2: 42}
                            _ = v16
                        }
                        _ = fi14
                        _ = v15
                    }
                } else {
                    var v17 int64 = 720
                    delete(v4, -49)
                    _ = v17
                }
                v1 += (v13 / ((v1) | 1))
            } else {
                switch 32767 {
                case 32767:
                    var v18 uint8 = 127
                    _ = v18
                case 1:
                    v5 = 0.0
                case 0:
                    v19 := GenS1{F0: 42}
                    _ = v19
                default:
                    if true {
                        for fi20 := 0; fi20 < 7; fi20++ {
                            if (float32(0.1) != (1.5 / -2.25)) {
                                for fi21 := 0; fi21 < 2; fi21++ {
                                    switch (fi20 % (((fi20 | fi20)) | 1)) {
                                    case 4294967295:
                                        var v22 int8 = 64
                                        _ = v22
                                    case 0:
                                        v13 += -874
                                    }
                                    v1 = v13
                                    v5 = float32(1e10)
                                    _ = fi21
                                }
                                switch (^(fi20)) {
                                case 32767:
                                    var v23 float32 = float32(3.14159)
                                    _ = v23
                                }
                                v5 -= (float32((((3.14159 + 0.0) + -2.25) - 1e10)) * (v3 / float32((((1.5 * float64((((0.0 - (1.5 - (1.5 * (v3 + (0.0 * (((1e10 / (v3 - float32((1.5 + 3.14159)))) + (float32(float64(float32(-2.25))) * v5)) * v5)))))) / ((float32((1.5 - ((3.14159 - 0.1) - 1e10))) / v5) / ((0.1 + (v5 + v3)) * v5))) / 1.5))) + 1.5) + 1.5))))
                            }
                            v24 := GenS1{F0: 461}
                            switch (("hello" + "\n\t") + ("hello" + "")) {
                            case "":
                                for fi25 := 0; fi25 < 13; fi25++ {
                                    switch 700 {
                                    case 127:
                                        var v26 int32 = -752
                                        _ = v26
                                    case 128:
                                        switch fi20 {
                                        case 4294967295:
                                            v2 = v2
                                        case -9223372036854775808:
                                            v3 = (v5 - (v3 - (v3 / 0.1)))
                                        default:
                                            v27 := []int{fi20, (fi20 | fi25), fi20, -277}
                                            _ = v27
                                        }
                                    case 32767:
                                        if true {
                                            switch fi20 {
                                            case -217:
                                                var v28 int16 = -32768
                                                _ = v28
                                            case 32768:
                                                var v29 uint16 = 0
                                                _ = v29
                                            case 42:
                                                if true {
                                                    v30 := v4[(-(fi20))]
                                                    _ = v30
                                                }
                                            }
                                        }
                                    }
                                    v31 := []int{(-(fi25))}
                                    if len(v31) > 0 {
                                        v31[0] = 127
                                    }
                                    fmt.Println(v5)
                                    _ = fi25
                                    _ = v31
                                }
                            case "hello":
                                v1 += v13
                            default:
                                v4[(fi20 ^ fi20)] = (fi20 % ((fi20) | 1))
                            }
                            _ = fi20
                            _ = v24
                        }
                    } else {
                        v32 := []int{}
                        v4[1] = -9223372036854775808
                        if true {
                            switch 9223372036854775807 {
                            case -258:
                                for fi33 := 0; fi33 < 12; fi33++ {
                                    v5 += -2.25
                                    _ = fi33
                                }
                            }
                            switch -2147483648 {
                            case 1:
                                v34 := GenS2{F0: (v1 + -1), F1: 2, F2: 264}
                                _ = v34
                            case -2147483648:
                                v35 := v4[32767]
                                _ = v35
                            case 9223372036854775807:
                                var v36 uint64 = 18446744073709551615
                                _ = v36
                            }
                            var v37 int = -2147483648
                            _ = v37
                        }
                        _ = v32
                    }
                }
                if (-284 > -660) {
                    switch 42 {
                    case 127:
                        fmt.Println(v2)
                    case 128:
                        if false {
                            v13 += (v1 % (((v1 / ((-32768) | 1))) | 1))
                            var v38 int64 = 942
                            v4[-562] = 907
                            _ = v38
                        } else {
                            var v39 int64 = v2
                            var v40 int8 = 100
                            if true {
                                v1 -= v1
                            } else {
                                if (!(255 != 808)) {
                                    for fi41 := 0; fi41 < 20; fi41++ {
                                        v4[32767] = int(v1)
                                        v42 := v4[(fi41 + fi41)]
                                        _ = fi41
                                        _ = v42
                                    }
                                    switch 32768 {
                                    case -87:
                                        var v43 int = 42
                                        _ = v43
                                    case 0:
                                        var v44 uint32 = 0
                                        _ = v44
                                    }
                                }
                                v39 -= 9223372036854775807
                            }
                            _ = v39
                            _ = v40
                        }
                    default:
                        v45 := []int{-9223372036854775808, 4294967295, -128, -851}
                        _ = v45
                    }
                    switch "hello" {
                    case "a":
                        switch "foo bar" {
                        case "a":
                            for fi46 := 0; fi46 < 18; fi46++ {
                                for fi47 := 0; fi47 < 20; fi47++ {
                                    var v48 float32 = (v3 * (v5 * v5))
                                    _ = fi47
                                    _ = v48
                                }
                                _ = fi46
                            }
                        case "\n\t":
                            for fi49 := 0; fi49 < 3; fi49++ {
                                switch "hello" {
                                case "\n\t":
                                    var v50 bool = true
                                    _ = v50
                                case "a":
                                    v3 += (v3 / v5)
                                }
                                v4[(-(fi49))] = (fi49 % ((fi49) | 1))
                                _ = fi49
                            }
                        case "hello":
                            if (v13 <= -1) {
                                v1 = v1
                                var v51 float64 = ((1e10 / float64(float32(0.1))) - (1.5 * 1e10))
                                var v52 float64 = v51
                                _ = v51
                                _ = v52
                            } else {
                                v53 := GenS2{F0: v13, F1: 2, F2: 255}
                                var v54 uint32 = 4294967295
                                _ = v53
                                _ = v54
                            }
                        }
                    case "":
                        v55 := []int{32767, 42, -647}
                        _ = v55
                    default:
                        var v56 string = "foo bar"
                        _ = v56
                    }
                    v1 = v1
                }
            }
            _ = v13
        }
        var v57 int8 = 100
        switch ("foo bar" + ("\n\t" + ("\n\t" + ((("a" + "hello") + (("" + "\n\t") + ("hello" + "a"))) + "")))) {
        case "hello":
            for fi58 := 0; fi58 < 16; fi58++ {
                var v59 uint8 = 0
                var v60 float32 = float32(1e10)
                for fi61 := 0; fi61 < 13; fi61++ {
                    if ((false || ((v3 + 3.14159) != ((1.5 - v3) - v60))) || (0.1 != float32(float64(3.14159)))) {
                        var v62 uint8 = (v59 % ((2) | 1))
                        fmt.Println(fi61)
                        _ = v62
                    } else {
                        fmt.Println(v59)
                        v63 := GenS3{F0: 0, F1: 505, F2: v57}
                        v64 := v4[fi58]
                        _ = v63
                        _ = v64
                    }
                    _ = fi61
                }
                _ = fi58
                _ = v59
                _ = v60
            }
        default:
            v3 = (v3 - 0.1)
        }
        _ = v57
    }
    v5 -= (float32(-2.25) * v3)
    var v65 int = 127
    var v66 string = "foo bar"
    var v67 uint64 = 429
    if (128 <= 117) {
        var v68 bool = (563 < 2)
        fmt.Println(v68)
        _ = v68
    }
    if false {
        switch "hello" {
        case "a":
            fmt.Println(v1, v66)
        case "foo bar":
            v69 := GenS4{F0: 42}
            _ = v69
        case "\n\t":
            v70 := []int{}
            _ = v70
        default:
            var v71 int8 = -60
            _ = v71
        }
        v4[v65] = (v65 * 128)
        v72 := []int{v65, 920}
        _ = v72
    } else {
        v73 := GenS3{F0: 486, F1: 32767, F2: -128}
        var v74 int32 = 0
        v74 -= v74
        _ = v73
        _ = v74
    }
    v75 := GenS2{F0: (v1 % (((v1 * v1)) | 1)), F1: 42, F2: 65535}
    var v76 string = v66
    if ((v2 > (v2 / (((v2 * 0)) | 1))) && (v2 >= v2)) {
        if (false || false) {
            v3 -= v5
            v5 += (1.5 * ((v3 / float32(-2.25)) + ((3.14159 + 0.1) / v5)))
        } else {
            delete(v4, (v65 | -2147483648))
            var v77 int32 = int32(v65)
            v78 := v4[(^(v65))]
            _ = v77
            _ = v78
        }
    }
    v79 := v4[42]
    fmt.Println(v65)
    v80 := GenS1{F0: -2147483648}
    v2 = 42
    var v81 uint = 42
    v67 = v67
    switch 9223372036854775807 {
    case 127:
        v82 := []int{4294967295, (v65 / (((v79 ^ v79)) | 1))}
        _ = v82
    default:
        fmt.Println(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v76:", v76)
    fmt.Println("v79:", v79)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v67
    _ = v75
    _ = v76
    _ = v79
    _ = v80
    _ = v81
}
