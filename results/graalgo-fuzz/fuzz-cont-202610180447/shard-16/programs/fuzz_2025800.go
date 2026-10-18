package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS3 struct { F0 string; F1 uint32; F2 uint }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 uint8; F1 uint16; F2 uint16 }

func main() {
    var v1 bool = false
    var v2 int64 = 9223372036854775807
    var v3 int16 = 42
    fmt.Println(v1, v2)
    if (true && (100 < -23)) {
        v4 := []int{32767, 43}
        v5 := GenS1{F0: 910}
        switch "hello" {
        case "hello":
            switch "\n\t" {
            case "hello":
                var v6 int64 = (v2 * v2)
                _ = v6
            }
        case "":
            v7 := GenS2{F0: 178, F1: 0, F2: v2}
            _ = v7
        case "a":
            if (-128 > -104) {
                if (127 > 88) {
                    for fi8 := 0; fi8 < 5; fi8++ {
                        for fi9 := 0; fi9 < 13; fi9++ {
                            v2 = v2
                            var v10 int64 = v2
                            _ = fi9
                            _ = v10
                        }
                        var v11 uint64 = 2
                        v11 += (v11 + (^(v11)))
                        _ = fi8
                        _ = v11
                    }
                    v12 := GenS3{F0: "a", F1: 1, F2: 4294967295}
                    _ = v12
                } else {
                    v3 -= (v3 >> ((799) & 63))
                    for fi13 := 0; fi13 < 17; fi13++ {
                        for fi14 := 0; fi14 < 11; fi14++ {
                            fmt.Println(fi13, v1)
                            _ = fi14
                        }
                        fmt.Println(fi13, v1)
                        _ = fi13
                    }
                    v4 = append(v4, 954)
                }
            } else {
                v2 -= (v2 << ((1) & 63))
            }
        }
        _ = v4
        _ = v5
    } else {
        v15 := GenS4{F0: 127, F1: 101}
        v16 := GenS1{F0: 4294967295}
        _ = v15
        _ = v16
    }
    v2 = (-(v2))
    var v17 uint32 = 663
    v1 = v1
    switch "\n\t" {
    case "hello":
        fmt.Println(v17)
    case "":
        v3 = (v3 ^ 32767)
    case "\n\t":
        for fi18 := 0; fi18 < 5; fi18++ {
            if (false && (127 > 1)) {
                v3 = (-(v3))
            } else {
                if v1 {
                    switch "hello" {
                    case "a":
                        v19 := []int{}
                        _ = v19
                    case "hello":
                        v20 := []int{(fi18 << ((42) & 63)), (fi18 ^ (fi18 % (((fi18 << ((378) & 63))) | 1))), (fi18 / ((fi18) | 1))}
                        _ = v20
                    case "":
                        v21 := GenS1{F0: 790}
                        _ = v21
                    }
                    v1 = false
                } else {
                    if (!v1) {
                        v22 := GenS3{F0: "hello", F1: (v17 % ((v17) | 1)), F2: 255}
                        _ = v22
                    } else {
                        if (v1 || (0.1 <= float64((((0.1 - 3.14159) + 1.5) / 1.5)))) {
                            v23 := []int{(fi18 >> ((819) & 63)), 127}
                            v24 := GenS5{F0: -135, F1: 4294967295, F2: 127}
                            _ = v23
                            _ = v24
                        }
                        var v25 uint8 = 212
                        _ = v25
                    }
                }
            }
            v2 -= 1
            _ = fi18
        }
    }
    if true {
        v17 -= (^(v17))
        for fi26 := 0; fi26 < 7; fi26++ {
            for fi27 := 0; fi27 < 8; fi27++ {
                var v28 int = fi27
                if (!((!v1) || true)) {
                    v17 -= v17
                    var v29 float32 = 0.1
                    _ = v29
                } else {
                    switch (v28 / ((32767) | 1)) {
                    case 32768:
                        fmt.Println(v3, v28)
                    default:
                        var v30 int = v28
                        _ = v30
                    }
                    var v31 int64 = v2
                    if v1 {
                        if (993 != 1) {
                            v31 -= int64(v17)
                        } else {
                            v32 := []int{fi27, (v28 / (((-(v28))) | 1)), (fi26 >> ((18446744073709551615) & 63)), (fi26 << ((255) & 63))}
                            var v33 int64 = v2
                            var v34 string = "hello"
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                    }
                    _ = v31
                }
                _ = fi27
                _ = v28
            }
            switch int(v17) {
            case 898:
                v1 = (127 == -128)
            }
            if v1 {
                v35 := GenS6{F0: 2147483647, F1: 65535}
                if (!(v1 && (1 > 42))) {
                    fmt.Println(v3, v3)
                    if v1 {
                        var v36 uint8 = 1
                        for fi37 := 0; fi37 < 4; fi37++ {
                            var v38 float32 = -2.25
                            if v1 {
                                v39 := GenS4{F0: fi37, F1: fi26}
                                _ = v39
                            }
                            _ = fi37
                            _ = v38
                        }
                        _ = v36
                    } else {
                        v40 := []int{-626, fi26, (^(fi26)), fi26}
                        _ = v40
                    }
                    var v41 uint16 = 2
                    _ = v41
                } else {
                    fmt.Println(v3, fi26, v1)
                    if v1 {
                        v42 := GenS7{F0: 0, F1: 32767, F2: 42}
                        _ = v42
                    } else {
                        for fi43 := 0; fi43 < 17; fi43++ {
                            switch -566 {
                            case -351:
                                v44 := GenS5{F0: 1, F1: 780, F2: 128}
                                _ = v44
                            default:
                                switch "\n\t" {
                                case "foo bar":
                                    switch "" {
                                    case "hello":
                                        switch (fi43 + (fi43 ^ (fi43 | -378))) {
                                        case -825:
                                            v45 := []int{2147483647, 307, (fi26 ^ 127)}
                                            _ = v45
                                        case -751:
                                            if true {
                                                v2 += (v2 & v2)
                                                v46 := GenS7{F0: 242, F1: 902, F2: 0}
                                                v47 := []int{}
                                                _ = v46
                                                _ = v47
                                            } else {
                                                v48 := []int{fi26, (fi43 | (fi26 - int(v17)))}
                                                v49 := GenS2{F0: 166, F1: 187, F2: v2}
                                                var v50 int64 = (v2 + 1)
                                                _ = v48
                                                _ = v49
                                                _ = v50
                                            }
                                        default:
                                            v17 = v17
                                        }
                                    case "a":
                                        if true {
                                            v1 = v1
                                            if (!false) {
                                                fmt.Println(fi43, v17, v17)
                                                fmt.Println(v17, v17, v3)
                                                v17 += v17
                                                v3 = v3
                                                v3 += (-(v3))
                                                v51 := []int{fi26, 370, (fi43 | fi26)}
                                                _ = v51
                                            } else {
                                                if (62 < 127) {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        fmt.Println(v1)
                                                    }
                                                }
                                            }
                                        } else {
                                            var v52 float32 = float32(float64((3.14159 / -2.25)))
                                            _ = v52
                                        }
                                    case "foo bar":
                                        v17 += (v17 ^ (^(v17)))
                                    default:
                                        v17 += 0
                                    }
                                case "hello":
                                    v53 := []int{}
                                    _ = v53
                                }
                            }
                            _ = fi43
                        }
                    }
                    switch (fi26 / ((1) | 1)) {
                    case 4294967295:
                        for fi54 := 0; fi54 < 10; fi54++ {
                            v55 := []int{(fi54 % ((-2147483648) | 1)), (fi26 + fi26), 2147483647}
                            v3 += v3
                            _ = fi54
                            _ = v55
                        }
                    default:
                        for fi56 := 0; fi56 < 8; fi56++ {
                            if v1 {
                                v2 -= v2
                                v57 := GenS1{F0: 332}
                                _ = v57
                            }
                            _ = fi56
                        }
                    }
                }
                v3 -= int16(v2)
                _ = v35
            }
            _ = fi26
        }
    } else {
        if v1 {
            for fi58 := 0; fi58 < 12; fi58++ {
                var v59 bool = (!false)
                _ = fi58
                _ = v59
            }
        } else {
            v2 = v2
            v60 := GenS5{F0: 43, F1: 913, F2: 107}
            _ = v60
        }
        for fi61 := 0; fi61 < 17; fi61++ {
            for fi62 := 0; fi62 < 7; fi62++ {
                var v63 uint16 = 723
                var v64 uint64 = 1
                var v65 string = ("a" + "a")
                _ = fi62
                _ = v63
                _ = v64
                _ = v65
            }
            var v66 int8 = -96
            _ = fi61
            _ = v66
        }
    }
    var v67 int16 = v3
    v68 := []int{}
    v68 = append(v68, -2147483648)
    var v69 uint16 = 2
    var v70 bool = (("\n\t" + "hello") < ("a" + (("a" + "\n\t") + "a")))
    v67 = v3
    var v71 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v67:", v67)
    fmt.Println("slicelen v68:", len(v68), v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v67
    _ = v68
    _ = v69
    _ = v70
    _ = v71
}
