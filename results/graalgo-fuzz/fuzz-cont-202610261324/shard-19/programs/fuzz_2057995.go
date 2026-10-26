package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 int64; F1 int16; F2 int64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 string; F1 uint; F2 int }
func genFn0(p0 int8, p1 string, p2 int16) uint {
    var v4 uint = 42
    switch -244 {
    case 0:
        if true {
            var v5 uint8 = 255
            _ = v5
        }
    case -2147483648:
        var v6 uint8 = 2
        _ = v6
    default:
        v7 := []int{2147483647, 0, -9223372036854775808, 1}
        _ = v7
    }
    var v8 uint8 = 127
    switch ((p1 + p1) + (("" + p1) + p1)) {
    case "hello":
        if (1 <= 0) {
            if false {
                var v9 int32 = 442
                p1 = ("foo bar" + p1)
                _ = v9
            }
            var v10 bool = false
            switch p1 {
            case "hello":
                var v11 string = p1
                _ = v11
            case "":
                var v12 uint64 = 87
                _ = v12
            case "a":
                p2 = p2
            }
            _ = v10
        } else {
            v13 := GenS1{F0: (v4 >> (((^(v4))) & 63))}
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := []int{fi14, fi14, 42}
                var v16 int8 = p0
                v17 := GenS2{F0: 2147483647, F1: 175}
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
            _ = v13
        }
    case "foo bar":
        switch ("\n\t" + "hello") {
        case "foo bar":
            p1 = p1
        case "":
            var v18 int16 = 32767
            _ = v18
        default:
            p1 = p1
        }
    }
    fmt.Println(p1, v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return (^(v4))
}
func genFn1(p0 bool, p1 uint32) int32 {
    if (128 > 255) {
        for fi19 := 0; fi19 < 20; fi19++ {
            p1 = (p1 << ((23) & 63))
            var v20 int32 = 265
            _ = fi19
            _ = v20
        }
        var v21 int = -376
        _ = v21
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        if true {
            if (-243 >= 1) {
                p1 += 65535
                var v23 float32 = 0.0
                _ = v23
            }
            v24 := GenS2{F0: (^(p1)), F1: -1}
            _ = v24
        } else {
            p1 += (p1 / ((p1) | 1))
            v25 := GenS3{F0: p0}
            _ = v25
        }
        if (42 >= 16) {
            for fi26 := 0; fi26 < 17; fi26++ {
                p1 -= p1
                _ = fi26
            }
        } else {
            v27 := GenS4{F0: 293, F1: 32767}
            _ = v27
        }
        if (253 < 2) {
            if p0 {
                var v28 uint64 = 0
                fmt.Println(p1, v28, fi22)
                _ = v28
            } else {
                p0 = false
                if (false || p0) {
                    var v29 uint16 = 107
                    v30 := GenS5{F0: 318, F1: 42}
                    _ = v29
                    _ = v30
                } else {
                    if (true || (1 != 0)) {
                        v31 := GenS1{F0: 65535}
                        v32 := []int{(^(fi22)), 2147483647}
                        v33 := GenS6{F0: -870, F1: 288, F2: 42}
                        _ = v31
                        _ = v32
                        _ = v33
                    } else {
                        if true {
                            if p0 {
                                v34 := []int{fi22, (fi22 * fi22), (fi22 + fi22)}
                                for fi35 := 0; fi35 < 7; fi35++ {
                                    var v36 uint8 = 62
                                    var v37 uint8 = 255
                                    _ = fi35
                                    _ = v36
                                    _ = v37
                                }
                                _ = v34
                            } else {
                                v38 := GenS1{F0: 493}
                                v39 := []int{fi22, 1}
                                switch "" {
                                case "\n\t":
                                    p1 = 366
                                case "foo bar":
                                    fmt.Println(fi22)
                                case "":
                                    if p0 {
                                        var v40 int64 = 42
                                        var v41 float32 = 3.14159
                                        _ = v40
                                        _ = v41
                                    } else {
                                        if p0 {
                                            for fi42 := 0; fi42 < 7; fi42++ {
                                                v43 := GenS4{F0: 1, F1: 387}
                                                var v44 uint = 1
                                                fmt.Println(p0, p1)
                                                p1 -= p1
                                                _ = fi42
                                                _ = v43
                                                _ = v44
                                            }
                                            if len(v39) >= 2 {
                                                v39 = v39[1:len(v39)]
                                            }
                                        } else {
                                            p1 += (p1 + p1)
                                            p1 = 884
                                            var v45 uint16 = 255
                                            _ = v45
                                        }
                                    }
                                default:
                                    p1 = (^(p1))
                                }
                                _ = v38
                                _ = v39
                            }
                            v46 := GenS3{F0: p0}
                            for fi47 := 0; fi47 < 15; fi47++ {
                                p1 += (^(p1))
                                var v48 int32 = -414
                                p0 = (p0 && true)
                                _ = fi47
                                _ = v48
                            }
                            _ = v46
                        }
                        fmt.Println(p0, fi22)
                        var v49 uint8 = 255
                        _ = v49
                    }
                    var v50 uint8 = 0
                    p1 = 0
                    _ = v50
                }
                v51 := GenS1{F0: 42}
                _ = v51
            }
        }
        _ = fi22
    }
    switch -2147483648 {
    case 9223372036854775807:
        var v52 uint16 = 42
        _ = v52
    case 32768:
        p1 += 648
    }
    v53 := GenS7{F0: 0}
    _ = p0
    _ = p1
    _ = v53
    return 2147483647
}

func main() {
    v1 := []uint32{}
    var v2 int64 = 0
    v3 := []string{(("\n\t" + "foo bar") + "hello"), ("a" + "\n\t"), "a"}
    fmt.Println(v2)
    var v54 float32 = -2.25
    v54 -= float32(0.1)
    if (v2 >= (-(v2))) {
        v3 = append(v3, "hello")
    }
    switch 453 {
    case -602:
        for fi55 := 0; fi55 < 1; fi55++ {
            fmt.Println(fi55, fi55)
            _ = fi55
        }
    case 32767:
        if len(v1) > 0 {
            v1[0] = 65535
        }
    case 398:
        v2 -= (v2 | v2)
    default:
        for fi56 := 0; fi56 < 14; fi56++ {
            var v57 float64 = 0.0
            _ = fi56
            _ = v57
        }
    }
    v58 := GenS8{F0: "hello", F1: 632, F2: 0}
    var v59 uint32 = 0
    for fi60 := 0; fi60 < 9; fi60++ {
        _ = len(v3)
        _ = fi60
    }
    if false {
        var v61 int8 = 42
        var v62 uint = 809
        _ = v61
        _ = v62
    } else {
        fmt.Println(v59)
        v63 := GenS2{F0: (v59 / (((^(v59))) | 1)), F1: v2}
        for fi64 := 0; fi64 < 2; fi64++ {
            var v65 uint8 = 1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi66 := 0; fi66 < 9; fi66++ {
                switch ("a" + "hello") {
                case "hello":
                    if false {
                        var v67 uint64 = 0
                        _ = v67
                    } else {
                        v65 = v65
                    }
                }
                var v68 bool = false
                var v69 int32 = 2147483647
                _ = fi66
                _ = v68
                _ = v69
            }
            _ = fi64
            _ = v65
        }
        _ = v63
    }
    var v70 int8 = -1
    var v71 float64 = -2.25
    switch "hello" {
    case "hello":
        if len(v3) > 0 {
            v3[0] = ("a" + "foo bar")
        }
    }
    var v72 int32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v54:", v54)
    fmt.Println("v59:", v59)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v58
    _ = v59
    _ = v70
    _ = v71
    _ = v72
}
