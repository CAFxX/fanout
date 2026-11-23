package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 bool, p1 uint32, p2 uint64) uint64 {
    v6 := GenS1{F0: -355, F1: -9223372036854775808}
    switch 127 {
    case 1:
        if p0 {
            var v7 uint8 = 2
            if p0 {
                p1 += (p1 >> ((4294967295) & 63))
                if (true && ((^(p2)) > (p2 / ((p2) | 1)))) {
                    for fi8 := 0; fi8 < 17; fi8++ {
                        v9 := []int{fi8, int(v7), (fi8 | fi8)}
                        _ = fi8
                        _ = v9
                    }
                } else {
                    p2 -= 18446744073709551615
                    p2 += (p2 % (((p2 - 4294967295)) | 1))
                    v10 := GenS1{F0: 1, F1: 9223372036854775807}
                    _ = v10
                }
                var v11 uint16 = 786
                _ = v11
            }
            _ = v7
        } else {
            var v12 int8 = 127
            _ = v12
        }
    case -2147483648:
        var v13 float64 = (float64(1.5) + 0.1)
        _ = v13
    case 32768:
        switch "foo bar" {
        case "hello":
            v14 := []int{32767, 128}
            _ = v14
        }
    default:
        var v15 int8 = -1
        _ = v15
    }
    p2 -= (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 4294967295
}
func genFn1() bool {
    var v16 int = -2147483648
    var v17 uint64 = 0
    v16 = 127
    var v18 uint32 = 11
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            v18 = v18
            v16 = (-(v16))
            v21 := GenS1{F0: -1, F1: -87}
            v18 = v18
            if (true || (127 == 42)) {
                var v22 uint64 = (v17 & v17)
                var v23 int16 = 42
                switch (fi20 / (((fi19 + 42)) | 1)) {
                case -838:
                    var v24 string = "hello"
                    _ = v24
                case -62:
                    var v25 uint8 = 127
                    _ = v25
                }
                var v26 float32 = ((1e10 + (0.1 - 3.14159)) - (1e10 / float32(-2.25)))
                _ = v22
                _ = v23
                _ = v26
            } else {
                v27 := []int{fi19, (^(fi20)), (fi19 ^ (v16 | -67))}
                _ = v27
            }
            _ = fi20
            _ = v21
        }
        for fi28 := 0; fi28 < 20; fi28++ {
            v18 -= 4294967295
            v18 += v18
            if true {
                switch "hello" {
                case "hello":
                    var v29 uint32 = v18
                    var v30 uint32 = v18
                    switch "\n\t" {
                    case "foo bar":
                        switch (v16 % ((fi28) | 1)) {
                        case -1:
                            v29 = (v29 % (((v30 & (^(v29)))) | 1))
                            var v31 uint64 = genFn0(false, v18, v17)
                            _ = v31
                        case 32768:
                            v32 := GenS1{F0: 653, F1: -515}
                            _ = v32
                        default:
                            var v33 float32 = (3.14159 + 3.14159)
                            _ = v33
                        }
                    default:
                        var v34 int8 = 0
                        _ = v34
                    }
                    _ = v29
                    _ = v30
                }
            }
            fmt.Println(v18, v18, v18)
            v35 := []int{(v16 & 32768), 9223372036854775807}
            _ = len(v35)
            v36 := GenS1{F0: -979, F1: -9223372036854775808}
            _ = fi28
            _ = v35
            _ = v36
        }
        for fi37 := 0; fi37 < 9; fi37++ {
            switch "" {
            case "\n\t":
                v38 := GenS1{F0: 482, F1: 720}
                _ = v38
            default:
                var v39 uint64 = (^(v17))
                _ = v39
            }
            _ = fi37
        }
        _ = fi19
    }
    _ = v16
    _ = v17
    _ = v18
    return ((v16 - int(v18)) >= (v16 - 128))
}

func main() {
    v1 := []uint32{}
    var v2 uint16 = 0
    v3 := []int{128, 9223372036854775807}
    var v4 bool = false
    var v5 int16 = -32768
    if (!((32768 < -2147483648) || v4)) {
        switch 26 {
        case 679:
            v40 := GenS2{F0: v2, F1: v2}
            _ = v40
        default:
            if (v4 || true) {
                v41 := GenS2{F0: 65535, F1: (v2 * (v2 % (((v2 ^ (^(v2)))) | 1)))}
                var v42 bool = ((42 != (v2 * 0)) || (2 >= 2))
                var v43 float32 = -2.25
                _ = v41
                _ = v42
                _ = v43
            }
        }
        v5 += (v5 % ((-1) | 1))
    } else {
        if v4 {
            v3 = append(v3, 266)
            var v44 int64 = -887
            if (true || (false || v4)) {
                v44 -= (v44 - v44)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi45 := 0; fi45 < 9; fi45++ {
                    var v46 int64 = (v44 << ((0) & 63))
                    _ = fi45
                    _ = v46
                }
            }
            _ = v44
        } else {
            fmt.Println(v4, v2)
            v47 := GenS3{F0: 2, F1: v2}
            _ = v47
        }
        var v48 uint16 = v2
        _ = v48
    }
    switch 418 {
    case -467:
        var v49 uint64 = 18446744073709551615
        _ = v49
    default:
        switch "foo bar" {
        case "a":
            v5 += (v5 << ((255) & 63))
        case "\n\t":
            switch "a" {
            case "foo bar":
                if (355 >= -963) {
                    v3 = append(v3, 570)
                    for fi50 := 0; fi50 < 13; fi50++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v51 bool = false
                        _ = fi50
                        _ = v51
                    }
                    _ = len(v1)
                } else {
                    fmt.Println(v5, v5, v5)
                }
            case "hello":
                switch ("a" + "a") {
                case "foo bar":
                    for fi52 := 0; fi52 < 12; fi52++ {
                        v5 -= (-(v5))
                        var v53 bool = (1 > 42)
                        _ = fi52
                        _ = v53
                    }
                case "\n\t":
                    for fi54 := 0; fi54 < 2; fi54++ {
                        for fi55 := 0; fi55 < 17; fi55++ {
                            if true {
                                v3 = append(v3, (fi55 - fi54))
                            } else {
                                fmt.Println(fi55, v2, v5)
                            }
                            _ = fi55
                        }
                        var v56 uint16 = v2
                        v56 -= (v2 * 733)
                        _ = fi54
                        _ = v56
                    }
                }
            case "a":
                for fi57 := 0; fi57 < 7; fi57++ {
                    fmt.Println(v4, v2, v4)
                    _ = fi57
                }
            }
        default:
            v5 -= -32768
        }
    }
    for fi58 := 0; fi58 < 12; fi58++ {
        switch "hello" {
        case "a":
            var v59 int32 = -1
            _ = v59
        }
        _ = fi58
    }
    v60 := GenS1{F0: 0, F1: 42}
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    var v61 int64 = 9223372036854775807
    v62 := GenS1{F0: -617, F1: -9223372036854775808}
    for fi63 := 0; fi63 < 3; fi63++ {
        for fi64 := 0; fi64 < 7; fi64++ {
            if (1 <= 646) {
                v65 := GenS2{F0: (v2 >> ((255) & 63)), F1: v2}
                fmt.Println(v4, fi63)
                _ = len(v3)
                _ = v65
            } else {
                var v66 int32 = 76
                v5 += (v5 | v5)
                _ = v66
            }
            _ = fi64
        }
        v67 := GenS4{F0: v2}
        var v68 uint64 = genFn0((443 == 4294967295), 42, 42)
        _ = fi63
        _ = v67
        _ = v68
    }
    var v69 uint32 = 566
    for fi70 := 0; fi70 < 3; fi70++ {
        v5 -= (v5 << ((250) & 63))
        _ = fi70
    }
    if ((v61 ^ (v61 * (v61 & v61))) < (v61 & (v61 - (v61 + v61)))) {
        _ = len(v3)
        if true {
            if false {
                var v71 int = 32768
                _ = v71
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            v72 := GenS3{F0: 127, F1: (v2 & v2)}
            _ = v72
        }
    } else {
        var v73 string = "\n\t"
        fmt.Println(v4, v73, v73)
        _ = v73
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v61
    _ = v62
    _ = v69
}
