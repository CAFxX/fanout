package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 int; F1 uint32; F2 int32 }
func genFn0(p0 string, p1 uint8, p2 string) int8 {
    v5 := []int{}
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := GenS1{F0: 2, F1: (p0 + p2), F2: (fi6 ^ (fi7 / ((fi6) | 1)))}
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -1
}
func genFn1(p0 int, p1 uint64) uint32 {
    p0 = p0
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := []int{p0, (p0 & 32767)}
        var v11 float64 = ((0.1 - 1e10) * -2.25)
        var v12 int = fi9
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    if (-1 < -1) {
        switch (p0 * p0) {
        case 1:
            for fi13 := 0; fi13 < 2; fi13++ {
                var v14 int = (fi13 % ((int(p1)) | 1))
                _ = fi13
                _ = v14
            }
        case 4294967295:
            switch -1 {
            case 1:
                var v15 int8 = -1
                _ = v15
            case 9223372036854775807:
                if false {
                    p0 += p0
                }
            default:
                if ("a" == ("a" + "a")) {
                    v16 := GenS2{F0: 1, F1: 0}
                    switch (p0 + 2147483647) {
                    case 32767:
                        fmt.Println(p0)
                    case 1:
                        v17 := GenS2{F0: 298, F1: 2}
                        _ = v17
                    }
                    _ = v16
                } else {
                    var v18 float64 = 1e10
                    p0 = (p0 % ((598) | 1))
                    _ = v18
                }
            }
        case 127:
            if (1000 > -32768) {
                if (!true) {
                    if (true || (0 == 42)) {
                        switch (p0 + p0) {
                        case 878:
                            for fi19 := 0; fi19 < 13; fi19++ {
                                switch fi19 {
                                case 32767:
                                    v20 := []int{p0, (fi19 << ((4294967295) & 63)), (fi19 - fi19), -2147483648}
                                    _ = v20
                                }
                                v21 := []int{}
                                for fi22 := 0; fi22 < 10; fi22++ {
                                    var v23 uint = 4294967295
                                    _ = fi22
                                    _ = v23
                                }
                                _ = fi19
                                _ = v21
                            }
                        case 42:
                            p0 -= 0
                        }
                    }
                    var v24 int32 = 42
                    _ = v24
                }
                p1 += (p1 - (p1 >> ((65535) & 63)))
                if (false && (2 > 731)) {
                    v25 := GenS3{F0: p0, F1: 42}
                    p1 = p1
                    _ = v25
                } else {
                    p0 = p0
                    p0 = int(p1)
                    var v26 int = 1
                    _ = v26
                }
            } else {
                v27 := []int{-537, 1, (p0 - p0)}
                _ = v27
            }
        }
        v28 := []int{p0, 9223372036854775807}
        _ = v28
    } else {
        v29 := GenS3{F0: 42, F1: 18446744073709551615}
        _ = v29
    }
    _ = p0
    _ = p1
    return 2147483648
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 58
    v3 := make(map[string]bool)
    v4 := []int{127, 754, 2147483647}
    if true {
        switch "a" {
        case "a":
            var v30 int16 = -32768
            _ = v30
        case "\n\t":
            if (true || true) {
                if false {
                    if (!true) {
                        v1[2147483647] = 2
                        switch -9223372036854775808 {
                        case -419:
                            fmt.Println(v2)
                        default:
                            if (!false) {
                                _ = len(v1)
                                v2 -= (^(v2))
                                var v31 uint32 = 65535
                                _ = v31
                            }
                        }
                    } else {
                        if true {
                            var v32 string = "a"
                            v33 := GenS3{F0: 0, F1: uint(v2)}
                            _ = v32
                            _ = v33
                        }
                    }
                    if len(v4) > 0 {
                        v4[0] = 128
                    }
                    for fi34 := 0; fi34 < 11; fi34++ {
                        var v35 float64 = (-2.25 - -2.25)
                        _ = fi34
                        _ = v35
                    }
                } else {
                    var v36 float64 = 3.14159
                    _ = v36
                }
                if (-1 != 1) {
                    v37 := GenS4{F0: 255, F1: 0}
                    v38 := v3["a"]
                    v1[558] = -465
                    _ = v37
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 7; fi39++ {
                        if (((-9223372036854775808 <= -374) && (-1 > (fi39 & fi39))) || true) {
                            fmt.Println(fi39)
                            if len(v4) > 0 {
                                v4[0] = (fi39 & fi39)
                            }
                            for fi40 := 0; fi40 < 13; fi40++ {
                                v2 = (^(v2))
                                _ = fi40
                            }
                        } else {
                            v41 := GenS5{F0: 35, F1: 42}
                            for fi42 := 0; fi42 < 1; fi42++ {
                                var v43 float32 = float32(0.0)
                                if true {
                                    if (18446744073709551615 != 2) {
                                        if ((!(false && true)) || ((-(fi42)) < fi42)) {
                                            var v44 int64 = 1
                                            _ = v44
                                        }
                                        var v45 uint64 = 18446744073709551615
                                        if len(v4) > 0 {
                                            v4[0] = 42
                                        }
                                        _ = v45
                                    }
                                } else {
                                    _ = len(v4)
                                    v46 := GenS1{F0: 1, F1: "\n\t", F2: fi42}
                                    _ = v46
                                }
                                _ = fi42
                                _ = v43
                            }
                            _ = v41
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        delete(v3, "\n\t")
                        _ = fi39
                    }
                }
            } else {
                if false {
                    _ = len(v4)
                    var v47 float32 = float32((1.5 - -2.25))
                    for fi48 := 0; fi48 < 20; fi48++ {
                        for fi49 := 0; fi49 < 4; fi49++ {
                            fmt.Println(fi49, fi49)
                            fmt.Println(v2, fi48, v47)
                            _ = fi49
                        }
                        _ = fi48
                    }
                    _ = v47
                } else {
                    v2 += v2
                }
            }
        case "foo bar":
            if len(v4) > 0 {
                v4[0] = 0
            }
        default:
            fmt.Println(v2)
        }
    }
    v50 := v3[(("" + "hello") + "hello")]
    for fi51 := 0; fi51 < 14; fi51++ {
        if (0 <= 2) {
            for fi52 := 0; fi52 < 20; fi52++ {
                v50 = (1 <= 1)
                _ = fi52
            }
        } else {
            v4 = append(v4, (fi51 - fi51))
            var v53 string = "hello"
            _ = v53
        }
        v50 = ((127 >= -44) && (v50 && v50))
        _ = fi51
    }
    if (0 > 42) {
        _ = len(v4)
        v50 = (true && v50)
        var v54 uint = 42
        _ = v54
    } else {
        var v55 int = 590
        fmt.Println(v50, v2)
        var v56 bool = true
        _ = v55
        _ = v56
    }
    if v50 {
        v57 := GenS6{F0: (42 < 255), F1: 42}
        _ = v57
    }
    v58 := GenS3{F0: -545, F1: 255}
    var v59 int = -912
    v50 = (-1 < -957)
    for fi60 := 0; fi60 < 20; fi60++ {
        v61 := GenS7{F0: v59, F1: 2, F2: 42}
        _ = fi60
        _ = v61
    }
    for fi62 := 0; fi62 < 16; fi62++ {
        var v63 uint8 = (v2 & v2)
        v2 = (v2 + v63)
        v64 := GenS1{F0: 1, F1: "a", F2: (v59 << ((255) & 63))}
        _ = fi62
        _ = v63
        _ = v64
    }
    var v65 uint = 18446744073709551615
    var v66 uint16 = 255
    v65 = (v65 << (((v65 + v65)) & 63))
    for fi67 := 0; fi67 < 5; fi67++ {
        var v68 uint64 = uint64(v65)
        if (("hello" + ("" + "a")) > ("a" + "\n\t")) {
            v68 = (v68 / ((v68) | 1))
        } else {
            var v69 int = fi67
            _ = v69
        }
        v68 = (v68 * 42)
        _ = fi67
        _ = v68
    }
    var v70 uint8 = uint8(v65)
    if v50 {
        var v71 int = (v59 % (((v59 & (v59 ^ -830))) | 1))
        v59 = (v71 % ((v71) | 1))
        _ = v71
    }
    switch 127 {
    case 142:
        var v72 int16 = -32768
        _ = v72
    case 127:
        if (v50 && v50) {
            switch (v59 | (^(v59))) {
            case 42:
                v66 += (^(v66))
            case 32768:
                var v73 int32 = 899
                _ = v73
            }
        }
    case -1:
        v66 -= 42
    default:
        var v74 int8 = genFn0(("hello" + "foo bar"), v2, (("hello" + "a") + "\n\t"))
        _ = v74
    }
    switch "hello" {
    case "foo bar":
        var v75 float64 = (float64(1.5) + (float64(1e10) - 0.0))
        _ = v75
    case "":
        v4 = append(v4, v59)
    case "a":
        v59 -= int(v66)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v50:", v50)
    fmt.Println("v59:", v59)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v58
    _ = v59
    _ = v65
    _ = v66
    _ = v70
}
