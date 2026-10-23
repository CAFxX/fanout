package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS6 struct { F0 int8 }
func genFn0(p0 int8) uint16 {
    p0 += (p0 % (((p0 >> ((1) & 63))) | 1))
    p0 += 100
    p0 = (p0 >> ((42) & 63))
    v7 := GenS1{F0: 704, F1: ""}
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 uint8, p1 int32) uint32 {
    p0 = p0
    fmt.Println(p0, p0)
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := []int{fi9}
            var v11 uint = 507
            _ = fi9
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 7; fi12++ {
            p1 = p1
            v13 := []int{fi12, fi8}
            p1 = p1
            _ = fi12
            _ = v13
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    return 958
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{}
    var v3 uint16 = 0
    var v4 string = "\n\t"
    var v5 int = 538
    v6 := make(map[uint32]uint32)
    v5 = 127
    var v14 int8 = -66
    v14 = (^(v14))
    if false {
        switch v5 {
        case 32767:
            for fi15 := 0; fi15 < 5; fi15++ {
                for fi16 := 0; fi16 < 7; fi16++ {
                    fmt.Println(fi16)
                    _ = fi16
                }
                _ = fi15
            }
        }
    }
    fmt.Println(v5, v4)
    if (true || ((v14 >= v14) || false)) {
        for fi17 := 0; fi17 < 19; fi17++ {
            if true {
                switch v4 {
                case "hello":
                    switch v5 {
                    case 0:
                        switch 4294967295 {
                        case -422:
                            if ((v14 % (((v14 / ((v14) | 1))) | 1)) != (^(v14))) {
                                switch "hello" {
                                case "":
                                    var v18 int8 = v14
                                    _ = v18
                                case "foo bar":
                                    if false {
                                        v5 = (v5 + fi17)
                                    }
                                default:
                                    switch 981 {
                                    case 4294967295:
                                        var v19 int = (fi17 + fi17)
                                        _ = v19
                                    case 0:
                                        var v20 uint8 = 2
                                        _ = v20
                                    case 139:
                                        if false {
                                            v4 = "\n\t"
                                            v3 = genFn0(v14)
                                        }
                                    default:
                                        v21 := GenS1{F0: 101, F1: (v4 + v4)}
                                        _ = v21
                                    }
                                }
                                v5 += v5
                                v3 -= (v3 - 2)
                            } else {
                                for fi22 := 0; fi22 < 5; fi22++ {
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v23 := v6[209]
                                        _ = v23
                                    } else {
                                        var v24 int16 = 32767
                                        v25 := v1[v4]
                                        v24 += v24
                                        _ = v24
                                        _ = v25
                                    }
                                    switch "\n\t" {
                                    case "":
                                        switch v4 {
                                        case "\n\t":
                                            v14 -= (v14 | 1)
                                        default:
                                            v14 += 100
                                        }
                                    }
                                    v14 = 42
                                    _ = fi22
                                }
                                if (!(723 == 526)) {
                                    if true {
                                        var v26 int8 = 0
                                        for fi27 := 0; fi27 < 2; fi27++ {
                                            _ = len(v2)
                                            v14 -= (v14 % (((v26 | v26)) | 1))
                                            var v28 float32 = -2.25
                                            _ = fi27
                                            _ = v28
                                        }
                                        _ = v26
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v29 uint32 = 2147483648
                                        v3 += v3
                                        _ = v29
                                    }
                                    v4 = v4
                                }
                                v30 := GenS2{F0: true, F1: 4294967295}
                                _ = v30
                            }
                        case 64:
                            switch (fi17 - v5) {
                            case 203:
                                var v31 uint8 = 1
                                _ = v31
                            case -1:
                                if false {
                                    if len(v2) > 0 {
                                        v2[0] = (fi17 & (fi17 << ((82) & 63)))
                                    }
                                    v32 := GenS3{F0: 2}
                                    fmt.Println(v14, v3)
                                    _ = v32
                                } else {
                                    v4 = "a"
                                    var v33 bool = true
                                    _ = v33
                                }
                            case 128:
                                if (65535 < (^(v3))) {
                                    if false {
                                        var v34 int16 = -362
                                        var v35 uint = 18446744073709551615
                                        _ = v34
                                        _ = v35
                                    }
                                    _ = len(v2)
                                    var v36 uint16 = genFn0((^(v14)))
                                    _ = v36
                                }
                            }
                        default:
                            var v37 uint8 = 128
                            _ = v37
                        }
                    case -2147483648:
                        for fi38 := 0; fi38 < 7; fi38++ {
                            var v39 uint8 = 255
                            _ = fi38
                            _ = v39
                        }
                    default:
                        if true {
                            var v40 uint8 = 2
                            for fi41 := 0; fi41 < 8; fi41++ {
                                var v42 uint = 2
                                v43 := GenS2{F0: (42 > 0), F1: 1}
                                var v44 uint = v42
                                _ = fi41
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                            for fi45 := 0; fi45 < 15; fi45++ {
                                var v46 uint64 = 2
                                if ((v14 & v14) < v14) {
                                    v4 = v4
                                    switch (v5 & v5) {
                                    case -9223372036854775808:
                                        v47 := GenS4{F0: v46, F1: v46, F2: 0}
                                        _ = v47
                                    case -707:
                                        var v48 int32 = 250
                                        _ = v48
                                    default:
                                        switch (v4 + ("a" + v4)) {
                                        case "":
                                            fmt.Println(v5, v14)
                                        }
                                    }
                                    if false {
                                        v3 = v3
                                        var v49 uint16 = v3
                                        _ = v49
                                    }
                                }
                                var v50 int64 = 789
                                _ = fi45
                                _ = v46
                                _ = v50
                            }
                            _ = v40
                        } else {
                            for fi51 := 0; fi51 < 19; fi51++ {
                                v52 := GenS5{F0: 2, F1: 1, F2: 2}
                                v2 = append(v2, int(v14))
                                _ = len(v6)
                                v5 = (^(fi51))
                                _ = fi51
                                _ = v52
                            }
                            var v53 uint = 4294967295
                            switch v5 {
                            case 1:
                                for fi54 := 0; fi54 < 8; fi54++ {
                                    v2 = append(v2, fi54)
                                    var v55 int = 2147483647
                                    v55 = 32767
                                    _ = fi54
                                    _ = v55
                                }
                            default:
                                if (42 != v53) {
                                    if false {
                                        fmt.Println(v4, v53)
                                        switch v4 {
                                        case "\n\t":
                                            switch (v5 % ((2147483647) | 1)) {
                                            case -326:
                                                v56 := GenS6{F0: int8(v53)}
                                                _ = v56
                                            case -313:
                                                for fi57 := 0; fi57 < 15; fi57++ {
                                                    var v58 float32 = 3.14159
                                                    _ = fi57
                                                    _ = v58
                                                }
                                            }
                                        case "a":
                                            v14 -= v14
                                        }
                                        v53 = v53
                                    } else {
                                        v59 := GenS1{F0: 9223372036854775807, F1: (v4 + v4)}
                                        _ = v59
                                    }
                                    for fi60 := 0; fi60 < 18; fi60++ {
                                        v2 = append(v2, fi60)
                                        _ = fi60
                                    }
                                    v14 = (v14 / ((-119) | 1))
                                }
                            }
                            _ = v53
                        }
                    }
                case "\n\t":
                    var v61 float32 = 1.5
                    _ = v61
                default:
                    if false {
                        if (!false) {
                            for fi62 := 0; fi62 < 7; fi62++ {
                                v5 -= fi62
                                _ = fi62
                            }
                            _ = len(v1)
                        } else {
                            v5 += (fi17 * (v5 * 9223372036854775807))
                            var v63 uint64 = 42
                            _ = len(v1)
                            _ = v63
                        }
                        var v64 uint64 = 565
                        for fi65 := 0; fi65 < 20; fi65++ {
                            if true {
                                fmt.Println(v4)
                                v5 = v5
                                var v66 int64 = -185
                                _ = v66
                            } else {
                                v14 += 127
                                var v67 int8 = -128
                                v68 := GenS5{F0: (v64 & 4294967295), F1: 255, F2: (^(v64))}
                                _ = v67
                                _ = v68
                            }
                            var v69 string = v4
                            switch 42 {
                            case 944:
                                if (!false) {
                                    v4 = ("\n\t" + (v4 + (v4 + v69)))
                                }
                            case 905:
                                v69 = v69
                            }
                            _ = fi65
                            _ = v69
                        }
                        _ = v64
                    }
                }
            } else {
                for fi70 := 0; fi70 < 17; fi70++ {
                    var v71 bool = (0 != 913)
                    _ = len(v2)
                    _ = fi70
                    _ = v71
                }
                v4 = ""
            }
            v72 := GenS6{F0: 1}
            _ = fi17
            _ = v72
        }
        v5 += v5
        var v73 string = (v4 + ("a" + v4))
        _ = v73
    } else {
        if false {
            fmt.Println(v14, v3)
            v74 := GenS2{F0: (true && (1 != -261)), F1: 160}
            _ = v74
        } else {
            var v75 float32 = 3.14159
            for fi76 := 0; fi76 < 6; fi76++ {
                v5 += (-(v5))
                v5 -= (v5 / (((fi76 << ((0) & 63))) | 1))
                v2 = append(v2, (fi76 / (((fi76 << ((255) & 63))) | 1)))
                _ = fi76
            }
            _ = v75
        }
        var v77 uint = 42
        _ = v77
    }
    fmt.Println(v5, v14, v5)
    fmt.Println(v3, v3)
    v14 += v14
    var v78 int32 = 0
    v14 -= 1
    for fi79 := 0; fi79 < 2; fi79++ {
        if true {
            v80 := GenS3{F0: 822}
            v81 := GenS3{F0: 651}
            if (!(!((32767 < -634) || (-32768 != 32767)))) {
                v5 += (v5 % ((v5) | 1))
                v14 += -32
            }
            _ = v80
            _ = v81
        }
        v5 -= 127
        fmt.Println(fi79, v5)
        _ = fi79
    }
    v82 := GenS2{F0: true, F1: 597}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v78:", v78)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
    _ = v82
}
