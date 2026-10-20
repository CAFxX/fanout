package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint64; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool; F1 int32; F2 int8 }
type GenS7 struct { F0 uint64; F1 int64; F2 uint32 }
func genFn0() uint8 {
    if false {
        v5 := GenS1{F0: 4294967295, F1: 1, F2: 0}
        v6 := []int{1}
        _ = v5
        _ = v6
    } else {
        if true {
            for fi7 := 0; fi7 < 16; fi7++ {
                v8 := GenS1{F0: 2, F1: 64, F2: 632}
                for fi9 := 0; fi9 < 15; fi9++ {
                    var v10 int = (-(fi7))
                    switch "" {
                    case "a":
                        fmt.Println(fi9, fi9)
                        var v11 uint8 = 53
                        if (18446744073709551615 > 535) {
                            v10 += 2147483647
                            v12 := GenS1{F0: 288, F1: -1, F2: 371}
                            v10 += fi7
                            _ = v12
                        } else {
                            var v13 uint8 = v11
                            _ = v13
                        }
                        _ = v11
                    case "foo bar":
                        switch fi7 {
                        case 1:
                            v14 := []int{fi9, (fi9 << ((42) & 63)), (^(v10)), fi9}
                            _ = v14
                        case -58:
                            v10 -= v10
                        case 42:
                            v10 = fi9
                        }
                    case "":
                        v15 := GenS1{F0: 255, F1: 1, F2: 255}
                        _ = v15
                    }
                    _ = fi9
                    _ = v10
                }
                _ = fi7
                _ = v8
            }
            for fi16 := 0; fi16 < 12; fi16++ {
                for fi17 := 0; fi17 < 19; fi17++ {
                    var v18 int8 = 1
                    _ = fi17
                    _ = v18
                }
                fmt.Println(fi16)
                v19 := GenS1{F0: 255, F1: 0, F2: 30}
                _ = fi16
                _ = v19
            }
        } else {
            var v20 bool = true
            _ = v20
        }
    }
    var v21 float64 = -2.25
    var v22 int16 = -768
    v23 := []int{127, -1, -379, -364}
    var v24 string = "\n\t"
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    return 0
}
func genFn1(p0 uint64) int32 {
    if false {
        var v25 int64 = -76
        _ = v25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        if true {
            fmt.Println(fi26)
        } else {
            var v27 int8 = -17
            _ = v27
        }
        for fi28 := 0; fi28 < 13; fi28++ {
            p0 += 2
            if ((true || (42 <= 0)) || true) {
                p0 = 2
                switch ("" + ("a" + ("foo bar" + "foo bar"))) {
                case "foo bar":
                    var v29 float32 = 1e10
                    _ = v29
                }
                if (true && true) {
                    var v30 bool = false
                    fmt.Println(fi26)
                    _ = v30
                } else {
                    p0 = 393
                    v31 := GenS2{F0: 0}
                    var v32 int16 = -461
                    var v33 uint64 = (p0 ^ p0)
                    var v34 uint8 = genFn0()
                    v35 := GenS2{F0: 42}
                    _ = v31
                    _ = v32
                    _ = v33
                    _ = v34
                    _ = v35
                }
            }
            if true {
                if true {
                    fmt.Println(fi28, fi28)
                    var v36 float64 = (3.14159 + 3.14159)
                    _ = v36
                }
                if (0.0 <= -2.25) {
                    var v37 uint16 = 0
                    v38 := []int{}
                    v39 := GenS1{F0: 1, F1: 127, F2: (v37 + 255)}
                    _ = v37
                    _ = v38
                    _ = v39
                } else {
                    switch "foo bar" {
                    case "hello":
                        p0 = (p0 - p0)
                    default:
                        v40 := GenS1{F0: 4294967295, F1: -1, F2: 65535}
                        _ = v40
                    }
                    v41 := GenS2{F0: 1}
                    v42 := GenS3{F0: 1000, F1: (p0 - p0), F2: (-(fi28))}
                    var v43 float32 = float32(1.5)
                    var v44 int64 = 399
                    _ = v41
                    _ = v42
                    _ = v43
                    _ = v44
                }
                v45 := GenS2{F0: 18446744073709551615}
                p0 -= 622
                _ = v45
            }
            if true {
                p0 = (p0 << ((255) & 63))
                p0 += p0
                p0 += p0
            }
            p0 += 655
            _ = fi28
        }
        if false {
            p0 += (p0 % ((2) | 1))
            var v46 int32 = 42
            v46 = v46
            _ = v46
        } else {
            fmt.Println(fi26, fi26)
        }
        _ = fi26
    }
    for fi47 := 0; fi47 < 10; fi47++ {
        p0 = p0
        for fi48 := 0; fi48 < 1; fi48++ {
            var v49 int16 = 32767
            v50 := GenS3{F0: (-(v49)), F1: p0, F2: 9223372036854775807}
            _ = fi48
            _ = v49
            _ = v50
        }
        for fi51 := 0; fi51 < 20; fi51++ {
            p0 = (p0 * 42)
            _ = fi51
        }
        _ = fi47
    }
    v52 := []int{}
    for fi53 := 0; fi53 < 5; fi53++ {
        v52 = append(v52, (-(fi53)))
        _ = fi53
    }
    _ = p0
    _ = v52
    return -374
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{("foo bar" + "a"), ""}
    v3 := []uint32{0, 2}
    var v4 uint8 = 127
    var v54 float64 = (-2.25 - 0.1)
    delete(v1, 2147483648)
    v54 += 0.0
    for fi55 := 0; fi55 < 19; fi55++ {
        v54 -= (1e10 - v54)
        if false {
            switch "a" {
            case "":
                _ = len(v1)
            }
        }
        _ = fi55
    }
    v56 := GenS2{F0: 778}
    v2 = append(v2, ("" + ("foo bar" + (("hello" + "foo bar") + "a"))))
    v57 := v1[536]
    v58 := GenS4{F0: 997}
    for fi59 := 0; fi59 < 5; fi59++ {
        switch (fi59 / (((fi59 + (fi59 << ((42) & 63)))) | 1)) {
        case -162:
            switch 437 {
            case 32768:
                var v60 int = (fi59 | -702)
                _ = v60
            case -832:
                v61 := GenS2{F0: 255}
                _ = v61
            case 9223372036854775807:
                if v57 {
                    v57 = (639 <= 42)
                    if (9223372036854775807 <= 42) {
                        for fi62 := 0; fi62 < 4; fi62++ {
                            v63 := GenS3{F0: 32767, F1: 2, F2: (fi62 << ((255) & 63))}
                            v4 = 1
                            v64 := GenS3{F0: -32768, F1: 2, F2: fi62}
                            _ = fi62
                            _ = v63
                            _ = v64
                        }
                    } else {
                        v57 = false
                        v57 = (!v57)
                    }
                    fmt.Println(v4, v57)
                } else {
                    v65 := GenS5{F0: 127}
                    if false {
                        v66 := GenS2{F0: 18446744073709551615}
                        var v67 uint = 18446744073709551615
                        switch "" {
                        case "\n\t":
                            if v57 {
                                var v68 int64 = 1
                                _ = v68
                            }
                        case "hello":
                            for fi69 := 0; fi69 < 16; fi69++ {
                                switch (^(fi69)) {
                                case 14:
                                    switch (("" + "") + ("foo bar" + (("a" + "\n\t") + "hello"))) {
                                    case "hello":
                                        if false {
                                            v4 = genFn0()
                                            v4 -= (v4 / ((1) | 1))
                                        } else {
                                            v70 := GenS6{F0: true, F1: 2147483647, F2: 0}
                                            var v71 int64 = 1
                                            fmt.Println(fi69, v4)
                                            _ = v70
                                            _ = v71
                                        }
                                    case "\n\t":
                                        var v72 int8 = 0
                                        _ = v72
                                    case "foo bar":
                                        var v73 int32 = 42
                                        _ = v73
                                    }
                                }
                                _ = fi69
                            }
                        }
                        _ = v66
                        _ = v67
                    }
                    _ = v65
                }
            default:
                var v74 float32 = (float32(v54) * -2.25)
                _ = v74
            }
        case -9223372036854775808:
            if len(v2) > 0 {
                v2[0] = "a"
            }
        default:
            v75 := GenS7{F0: 1, F1: 9223372036854775807, F2: 42}
            _ = v75
        }
        _ = fi59
    }
    if v57 {
        v54 = float64((0.0 * (1.5 - (3.14159 + -2.25))))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v56
    _ = v57
    _ = v58
}
