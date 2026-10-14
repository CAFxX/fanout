package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint64; F1 string; F2 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 int16; F1 int16; F2 int }
type GenS8 struct { F0 int16 }
func genFn0() int8 {
    var v5 uint64 = 2
    if (784 <= -32768) {
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 int = -26
            _ = fi6
            _ = v7
        }
        v8 := []int{9223372036854775807}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = v8
    } else {
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint = 1
            _ = fi9
            _ = v10
        }
        v5 -= (v5 & 42)
    }
    switch (("" + ("\n\t" + ("a" + "hello"))) + "a") {
    case "hello":
        v5 -= (v5 << ((4294967295) & 63))
    case "\n\t":
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := []int{fi11, fi11, (fi11 + 39)}
            _ = fi11
            _ = v12
        }
    case "":
        var v13 float32 = ((1e10 + 3.14159) + (float32(-2.25) - float32(0.1)))
        _ = v13
    default:
        v14 := GenS1{F0: -1}
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        v5 = v5
        v16 := []int{(fi15 >> ((255) & 63)), (fi15 | fi15), (fi15 / ((fi15) | 1))}
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        v5 += (v5 * (^(v5)))
        _ = fi17
    }
    _ = v5
    return 0
}
func genFn1(p0 uint, p1 string) string {
    var v18 int16 = 32767
    var v19 int = 204
    switch p1 {
    case "a":
        p0 += (p0 % (((^(p0))) | 1))
    case "hello":
        for fi20 := 0; fi20 < 6; fi20++ {
            fmt.Println(fi20)
            v21 := GenS1{F0: (v18 & (^(v18)))}
            _ = fi20
            _ = v21
        }
    case "":
        v22 := GenS1{F0: 32767}
        _ = v22
    }
    v23 := GenS1{F0: (v18 % (((v18 >> ((18446744073709551615) & 63))) | 1))}
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    _ = v23
    return (p1 + "")
}

func main() {
    var v1 uint16 = 2
    var v2 float32 = 0.0
    v3 := make(map[uint32]uint32)
    v4 := []uint32{811, 486, 42, 578}
    switch ("\n\t" + "hello") {
    case "hello":
        switch 9223372036854775807 {
        case 512:
            _ = len(v4)
        case 720:
            if false {
                var v24 int16 = 0
                if false {
                    var v25 int16 = v24
                    for fi26 := 0; fi26 < 3; fi26++ {
                        delete(v3, 42)
                        _ = fi26
                    }
                    _ = v25
                }
                _ = v24
            } else {
                var v27 float64 = (3.14159 * (1e10 / -2.25))
                _ = v27
            }
        case 384:
            fmt.Println(v2, v2)
        default:
            var v28 uint = 1
            _ = v28
        }
    case "\n\t":
        v2 += v2
    case "foo bar":
        switch -430 {
        case 1:
            fmt.Println(v1, v1)
        case 32767:
            v1 = v1
        case -296:
            var v29 int = 127
            _ = v29
        }
    default:
        if false {
            for fi30 := 0; fi30 < 15; fi30++ {
                if (false || false) {
                    if len(v4) > 0 {
                        v4[0] = 0
                    }
                    _ = len(v3)
                    v2 += (v2 / -2.25)
                } else {
                    for fi31 := 0; fi31 < 17; fi31++ {
                        v32 := v3[2]
                        v1 -= v1
                        v2 = (v2 + v2)
                        _ = fi31
                        _ = v32
                    }
                    v33 := GenS2{F0: 1000, F1: 1}
                    if false {
                        v2 = v2
                    } else {
                        fmt.Println(v2, v1, fi30)
                        v2 = v2
                    }
                    _ = v33
                }
                v2 = (v2 / ((float32(0.0) / 1e10) / v2))
                _ = fi30
            }
        }
    }
    _ = len(v3)
    for fi34 := 0; fi34 < 15; fi34++ {
        v35 := GenS3{F0: (v1 + (v1 | 1)), F1: (^(v1)), F2: 186}
        _ = fi34
        _ = v35
    }
    var v36 int32 = 42
    v37 := GenS2{F0: -1, F1: 0}
    _ = len(v4)
    var v38 uint = 2
    v39 := v3[4294967295]
    fmt.Println(v38, v1, v1)
    switch 32768 {
    case -453:
        v40 := GenS1{F0: -1}
        _ = v40
    default:
        v41 := GenS4{F0: 1, F1: "a", F2: 42}
        _ = v41
    }
    if true {
        if ((^(v38)) != 65535) {
            fmt.Println(v38, v36, v2)
            if ((1.5 - (v2 - v2)) < v2) {
                if true {
                    switch "a" {
                    case "\n\t":
                        var v42 uint = (v38 | v38)
                        _ = v42
                    case "a":
                        switch ("foo bar" + ("foo bar" + "hello")) {
                        case "hello":
                            fmt.Println(v1, v38, v39)
                        case "foo bar":
                            fmt.Println(v36)
                        default:
                            if false {
                                _ = len(v4)
                                if (!false) {
                                    var v43 float64 = 1e10
                                    _ = v43
                                } else {
                                    v36 = (-(v36))
                                    fmt.Println(v2)
                                    if (1 <= 2147483647) {
                                        v39 = 65535
                                        v2 -= v2
                                    }
                                }
                                if false {
                                    var v44 int16 = -1
                                    _ = v44
                                } else {
                                    for fi45 := 0; fi45 < 12; fi45++ {
                                        fmt.Println(v2, fi45)
                                        _ = fi45
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = v39
                                    }
                                }
                            } else {
                                switch ("\n\t" + "\n\t") {
                                case "":
                                    var v46 uint64 = 973
                                    _ = v46
                                case "\n\t":
                                    if true {
                                        v3[v39] = 2147483647
                                        v4 = append(v4, (v39 * uint32(v38)))
                                    }
                                }
                            }
                        }
                    case "":
                        v39 = v39
                    }
                    fmt.Println(v2, v36)
                } else {
                    if true {
                        for fi47 := 0; fi47 < 2; fi47++ {
                            v39 += (v39 << ((v38) & 63))
                            var v48 uint = (v38 >> ((4294967295) & 63))
                            _ = fi47
                            _ = v48
                        }
                        v49 := GenS1{F0: 352}
                        _ = v49
                    } else {
                        v39 = 2147483647
                        if false {
                            switch (("\n\t" + "a") + "foo bar") {
                            case "\n\t":
                                switch "foo bar" {
                                case "foo bar":
                                    fmt.Println(v38, v2)
                                case "hello":
                                    v50 := GenS2{F0: 1000, F1: (v39 % (((v39 << ((1) & 63))) | 1))}
                                    _ = v50
                                case "\n\t":
                                    if true {
                                        var v51 float32 = 0.0
                                        fmt.Println(v1)
                                        if (921 <= -1) {
                                            v52 := GenS5{F0: false}
                                            var v53 int = -1
                                            _ = v52
                                            _ = v53
                                        } else {
                                            v54 := GenS1{F0: 0}
                                            if true {
                                                v2 = float32(3.14159)
                                                for fi55 := 0; fi55 < 20; fi55++ {
                                                    fmt.Println(v2, fi55)
                                                    var v56 bool = (v38 < (v38 / ((v38) | 1)))
                                                    fmt.Println(v51)
                                                    _ = fi55
                                                    _ = v56
                                                }
                                                v36 -= v36
                                            }
                                            _ = v54
                                        }
                                        _ = v51
                                    } else {
                                        var v57 int64 = 42
                                        for fi58 := 0; fi58 < 3; fi58++ {
                                            var v59 string = ("foo bar" + ("a" + ((("hello" + "\n\t") + "foo bar") + "\n\t")))
                                            v36 = 45
                                            var v60 int16 = 32767
                                            _ = fi58
                                            _ = v59
                                            _ = v60
                                        }
                                        _ = v57
                                    }
                                default:
                                    if false {
                                        var v61 float64 = ((-2.25 - -2.25) / 0.1)
                                        _ = len(v3)
                                        v62 := GenS6{F0: v39, F1: 787}
                                        _ = v61
                                        _ = v62
                                    } else {
                                        delete(v3, v39)
                                    }
                                }
                            case "foo bar":
                                v63 := GenS7{F0: 0, F1: 1000, F2: 128}
                                _ = v63
                            default:
                                if (true && false) {
                                    for fi64 := 0; fi64 < 17; fi64++ {
                                        for fi65 := 0; fi65 < 10; fi65++ {
                                            var v66 int64 = -1
                                            _ = fi65
                                            _ = v66
                                        }
                                        _ = fi64
                                    }
                                }
                            }
                            fmt.Println(v2, v36)
                        }
                    }
                    for fi67 := 0; fi67 < 1; fi67++ {
                        var v68 uint32 = (v39 ^ v39)
                        if true {
                            v69 := GenS3{F0: uint16(v39), F1: v1, F2: uint8(v1)}
                            v70 := GenS7{F0: 1, F1: 1, F2: -174}
                            _ = v69
                            _ = v70
                        }
                        _ = fi67
                        _ = v68
                    }
                }
            }
            v1 -= v1
        }
    }
    switch "foo bar" {
    case "a":
        switch 769 {
        case 128:
            fmt.Println(v38, v36)
        default:
            for fi71 := 0; fi71 < 9; fi71++ {
                _ = len(v3)
                _ = fi71
            }
        }
    default:
        switch "" {
        case "hello":
            fmt.Println(v38, v38, v38)
        case "\n\t":
            var v72 string = "a"
            _ = v72
        case "a":
            if true {
                var v73 float64 = (float64(3.14159) - 1.5)
                _ = v73
            }
        default:
            fmt.Println(v2, v1)
        }
    }
    v74 := GenS4{F0: 1, F1: ((("hello" + "") + "foo bar") + ""), F2: 0}
    v4 = append(v4, v39)
    var v75 uint32 = 979
    for fi76 := 0; fi76 < 19; fi76++ {
        fmt.Println(v38)
        for fi77 := 0; fi77 < 11; fi77++ {
            var v78 int8 = 0
            _ = fi77
            _ = v78
        }
        switch "\n\t" {
        case "":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "a":
            if (false || true) {
                var v79 string = "foo bar"
                v80 := GenS3{F0: v1, F1: v1, F2: 128}
                _ = v79
                _ = v80
            }
        default:
            fmt.Println(v2)
        }
        _ = fi76
    }
    v39 -= v39
    v81 := GenS8{F0: -727}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v74
    _ = v75
    _ = v81
}
