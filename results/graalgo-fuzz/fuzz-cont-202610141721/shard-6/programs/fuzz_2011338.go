package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int; F1 uint8 }
func genFn0(p0 uint16) uint8 {
    v7 := []int{4294967295, 127, 32767}
    v8 := GenS1{F0: 2, F1: 1, F2: -128}
    for fi9 := 0; fi9 < 3; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            fmt.Println(p0, fi9, fi9)
            switch "" {
            case "foo bar":
                p0 = p0
                fmt.Println(p0, fi10)
                var v11 int8 = 100
                _ = v11
            case "a":
                var v12 int = 128
                _ = v12
            }
            fmt.Println(p0)
            _ = fi10
        }
        var v13 int64 = 475
        _ = fi9
        _ = v13
    }
    var v14 uint8 = 0
    for fi15 := 0; fi15 < 2; fi15++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi15
    }
    _ = p0
    _ = v14
    _ = v7
    _ = v8
    return 49
}
func genFn1(p0 int8) int64 {
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 uint32 = 2147483648
        _ = fi16
        _ = v17
    }
    v18 := []int{0, -488}
    p0 -= (p0 & (p0 & p0))
    var v19 string = ""
    _ = p0
    _ = v18
    _ = v19
    return -180
}

func main() {
    var v1 uint64 = 711
    v2 := []uint32{65535}
    var v3 float64 = 3.14159
    v4 := make(map[string]bool)
    var v5 int32 = 1
    var v6 uint32 = uint32(v1)
    for fi20 := 0; fi20 < 6; fi20++ {
        var v21 uint8 = 57
        v3 = (v3 / ((1.5 / 1.0) + v3))
        _ = fi20
        _ = v21
    }
    v22 := GenS2{F0: 0, F1: (v5 & -1)}
    var v23 int32 = v5
    var v24 int = -716
    v1 -= 630
    if (1 < 812) {
        v5 -= 42
        if len(v2) > 0 {
            v2[0] = (v6 ^ v6)
        }
    } else {
        for fi25 := 0; fi25 < 14; fi25++ {
            var v26 string = ("a" + "a")
            if (1 <= 255) {
                v3 = float64(0.0)
                switch (v24 & -301) {
                case 9223372036854775807:
                    v27 := v4["\n\t"]
                    _ = v27
                case 32768:
                    v28 := GenS3{F0: v5, F1: 65535}
                    _ = v28
                case -1:
                    switch "hello" {
                    case "\n\t":
                        switch v26 {
                        case "a":
                            for fi29 := 0; fi29 < 8; fi29++ {
                                switch fi29 {
                                case -243:
                                    if len(v2) > 0 {
                                        v2[0] = v6
                                    }
                                case 4294967295:
                                    fmt.Println(v1)
                                }
                                delete(v4, "foo bar")
                                _ = fi29
                            }
                        default:
                            switch v24 {
                            case 127:
                                v26 = "hello"
                            case -1:
                                v23 += v23
                            case 0:
                                for fi30 := 0; fi30 < 8; fi30++ {
                                    if (42 < 833) {
                                        if (!(!(!true))) {
                                            v31 := GenS2{F0: 74, F1: 1}
                                            if true {
                                                v3 += 0.1
                                                fmt.Println(fi25)
                                                switch fi30 {
                                                case 32768:
                                                    v1 = 399
                                                case 0:
                                                    v32 := GenS1{F0: v1, F1: 42, F2: 1}
                                                    _ = v32
                                                }
                                                v33 := GenS2{F0: 16, F1: (v5 / ((v23) | 1))}
                                                _ = v33
                                            }
                                            if true {
                                                switch 682 {
                                                case 1:
                                                    if false {
                                                        fmt.Println(v23, v24)
                                                    }
                                                case -598:
                                                    var v34 int32 = (v23 % ((v23) | 1))
                                                    _ = v34
                                                case 185:
                                                    if ((0.1 - 1.5) >= 3.14159) {
                                                        v23 -= (v5 >> ((996) & 63))
                                                    }
                                                default:
                                                    if false {
                                                        v4[v26] = true
                                                    } else {
                                                        v23 -= v5
                                                        v35 := GenS4{F0: 32768, F1: genFn0(2)}
                                                        switch "a" {
                                                        case "foo bar":
                                                            fmt.Println(fi30)
                                                        case "hello":
                                                            v1 += (v1 - 440)
                                                        case "a":
                                                            fmt.Println(v1)
                                                        }
                                                        _ = v35
                                                    }
                                                }
                                            }
                                            _ = v31
                                        } else {
                                            var v36 float32 = 1.5
                                            _ = v36
                                        }
                                    } else {
                                        var v37 bool = (1 < genFn0(0))
                                        _ = v37
                                    }
                                    var v38 int64 = 1
                                    var v39 string = v26
                                    var v40 uint16 = 984
                                    _ = fi30
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                }
                            default:
                                var v41 float32 = (((0.0 + -2.25) / 1.0) * (1e10 - (1.5 * -2.25)))
                                _ = v41
                            }
                        }
                    case "":
                        if true {
                            v42 := GenS2{F0: -96, F1: (^(v23))}
                            var v43 uint16 = 0
                            _ = v42
                            _ = v43
                        }
                    case "a":
                        var v44 uint16 = 42
                        _ = v44
                    default:
                        if true {
                            for fi45 := 0; fi45 < 17; fi45++ {
                                var v46 int32 = (v5 + v23)
                                _ = fi45
                                _ = v46
                            }
                        }
                    }
                }
            }
            for fi47 := 0; fi47 < 19; fi47++ {
                v4[v26] = true
                var v48 string = "\n\t"
                _ = fi47
                _ = v48
            }
            _ = fi25
            _ = v26
        }
    }
    var v49 float32 = float32(v3)
    if true {
        var v50 uint64 = (v1 << ((348) & 63))
        _ = v50
    }
    switch ("" + ("foo bar" + "a")) {
    case "\n\t":
        v51 := GenS3{F0: v23, F1: 2}
        _ = v51
    case "foo bar":
        var v52 uint64 = (v1 % (((v1 + 1)) | 1))
        _ = v52
    case "hello":
        if ((!true) && true) {
            if false {
                for fi53 := 0; fi53 < 7; fi53++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi53
                }
                v5 = -1
            }
        }
    }
    if (127 >= -53) {
        for fi54 := 0; fi54 < 12; fi54++ {
            var v55 int64 = genFn1(-1)
            var v56 int64 = -9223372036854775808
            if (v1 <= 18446744073709551615) {
                if len(v2) > 0 {
                    v2[0] = 1
                }
                v56 -= v55
                if false {
                    if true {
                        v57 := GenS1{F0: v1, F1: 32767, F2: -96}
                        _ = v57
                    }
                    if true {
                        switch ("\n\t" + "a") {
                        case "a":
                            v58 := v4[("hello" + "a")]
                            _ = v58
                        case "hello":
                            v6 = (v6 % (((^(v6))) | 1))
                        }
                        for fi59 := 0; fi59 < 10; fi59++ {
                            var v60 float32 = -2.25
                            fmt.Println(v60)
                            _ = fi59
                            _ = v60
                        }
                        v6 += (^(v6))
                    }
                    v3 += (v3 + (v3 + (((v3 + v3) - v3) - float64(float32(1.5)))))
                }
            }
            _ = fi54
            _ = v55
            _ = v56
        }
        var v61 int = v24
        var v62 float32 = (v49 * 1.5)
        _ = v61
        _ = v62
    }
    switch "hello" {
    case "":
        var v63 string = ("hello" + ("foo bar" + "a"))
        _ = v63
    }
    switch v24 {
    case 474:
        v5 = v5
    case 2147483647:
        v3 += float64(1.5)
    }
    fmt.Println(v5, v23, v23)
    var v64 int = int(v23)
    if (0.1 == v49) {
        if true {
            v24 -= (v64 + (v24 >> ((2) & 63)))
            for fi65 := 0; fi65 < 3; fi65++ {
                v6 = 218
                _ = fi65
            }
            var v66 string = ("\n\t" + "")
            _ = v66
        } else {
            v67 := GenS3{F0: v5, F1: 684}
            v49 += (v49 - 3.14159)
            var v68 int64 = genFn1(int8(v24))
            _ = v67
            _ = v68
        }
        var v69 int64 = 42
        _ = v69
    } else {
        v5 -= -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v64
}
