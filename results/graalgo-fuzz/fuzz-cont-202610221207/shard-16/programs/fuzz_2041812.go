package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 int16; F1 int16 }

func main() {
    var v1 string = ("foo bar" + "")
    var v2 int8 = -1
    var v3 int32 = -2147483648
    v4 := make(map[int]uint32)
    v5 := make(map[int]bool)
    if (282 < -9223372036854775808) {
        fmt.Println(v3, v2)
        fmt.Println(v1)
        fmt.Println(v1, v3)
    } else {
        v2 += 100
    }
    if false {
        var v6 uint = 42
        _ = len(v4)
        _ = v6
    } else {
        v1 = "hello"
        v3 -= v3
    }
    var v7 int = -690
    var v8 int = (v7 ^ 128)
    var v9 int32 = (v3 * v3)
    if ((v2 % ((-1) | 1)) < 100) {
        v10 := GenS1{F0: 42, F1: -9223372036854775808, F2: v3}
        switch "" {
        case "hello":
            switch 645 {
            case 128:
                switch "a" {
                case "foo bar":
                    v2 -= v2
                case "a":
                    v3 -= (v3 | 1)
                case "hello":
                    v11 := GenS2{F0: "foo bar"}
                    _ = v11
                }
            }
        case "foo bar":
            switch "a" {
            case "foo bar":
                v3 += v9
            case "\n\t":
                v3 -= v9
            case "a":
                var v12 float32 = 0.0
                _ = v12
            default:
                fmt.Println(v9, v2, v9)
            }
        }
        _ = v10
    } else {
        if (false && false) {
            if (false && true) {
                var v13 uint32 = 2147483647
                _ = v13
            } else {
                v4[(v8 & (v8 & (v8 / (((v8 << ((65535) & 63))) | 1))))] = 0
                for fi14 := 0; fi14 < 15; fi14++ {
                    v15 := GenS1{F0: 65535, F1: 0, F2: -832}
                    switch 128 {
                    case 127:
                        v7 += (v7 - fi14)
                    case 1:
                        fmt.Println(v7, fi14, v8)
                    default:
                        v2 = v2
                    }
                    var v16 uint = 255
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                if (!true) {
                    _ = len(v4)
                } else {
                    if true {
                        delete(v4, v8)
                        v17 := []int{v7, v8}
                        _ = v17
                    } else {
                        var v18 uint64 = 42
                        v2 -= (v2 >> ((18446744073709551615) & 63))
                        switch int(v18) {
                        case -403:
                            fmt.Println(v9, v1)
                        default:
                            v19 := GenS2{F0: v1}
                            _ = v19
                        }
                        _ = v18
                    }
                }
            }
            var v20 float64 = 0.1
            _ = len(v4)
            _ = v20
        } else {
            delete(v5, (v7 & (v7 ^ (v8 % ((v7) | 1)))))
            v1 = "\n\t"
            v1 = v1
        }
    }
    var v21 uint = 144
    if false {
        var v22 int = v7
        var v23 uint16 = 1
        _ = v22
        _ = v23
    } else {
        v3 -= v9
        v24 := v5[128]
        _ = v24
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        for fi26 := 0; fi26 < 15; fi26++ {
            if ((0.0 - -2.25) != (1.5 + (1e10 / float64(float32(-2.25))))) {
                var v27 int64 = 677
                _ = v27
            } else {
                var v28 float32 = (1.5 / (1.5 * 1.5))
                var v29 float32 = (v28 / 1.0)
                v7 = (fi25 - 32768)
                _ = v28
                _ = v29
            }
            v21 = 18446744073709551615
            _ = fi26
        }
        v30 := []int{}
        v31 := v4[v7]
        _ = fi25
        _ = v30
        _ = v31
    }
    if (false || true) {
        switch ((("\n\t" + v1) + v1) + v1) {
        case "hello":
            for fi32 := 0; fi32 < 6; fi32++ {
                v3 = -2147483648
                _ = fi32
            }
        case "foo bar":
            if true {
                for fi33 := 0; fi33 < 14; fi33++ {
                    if ("a" >= v1) {
                        v34 := GenS2{F0: v1}
                        v7 = v7
                        var v35 bool = (true && false)
                        _ = v34
                        _ = v35
                    } else {
                        v2 -= (v2 + -23)
                        for fi36 := 0; fi36 < 5; fi36++ {
                            _ = len(v5)
                            switch v1 {
                            case "":
                                var v37 int16 = int16(v3)
                                v1 = ""
                                _ = v37
                            case "a":
                                delete(v5, (fi36 << (((v21 * 0)) & 63)))
                            }
                            fmt.Println(v8, v8, v8)
                            _ = fi36
                        }
                    }
                    v8 += (v7 >> ((v21) & 63))
                    var v38 int32 = v3
                    _ = fi33
                    _ = v38
                }
                v39 := v5[(^(v7))]
                var v40 string = v1
                _ = v39
                _ = v40
            } else {
                var v41 uint64 = 1
                var v42 int16 = 1000
                _ = v41
                _ = v42
            }
        }
        v43 := []int{(v8 | (v7 * v8)), v7, (v7 - 32767), 42}
        var v44 string = v1
        _ = v43
        _ = v44
    } else {
        v45 := GenS3{F0: 2, F1: 127}
        if (507 == 32767) {
            v46 := GenS3{F0: 42, F1: 255}
            _ = v46
        } else {
            for fi47 := 0; fi47 < 13; fi47++ {
                for fi48 := 0; fi48 < 13; fi48++ {
                    var v49 int16 = -266
                    if (v2 != (v2 & v2)) {
                        v7 = (fi47 | fi48)
                        v49 = -518
                    }
                    _ = fi48
                    _ = v49
                }
                for fi50 := 0; fi50 < 19; fi50++ {
                    if ((false && false) || false) {
                        fmt.Println(v1, v1, v21)
                        if false {
                            v2 = v2
                            if false {
                                var v51 int16 = 32767
                                var v52 int = v8
                                _ = v51
                                _ = v52
                            }
                            var v53 float64 = 0.1
                            v54 := v5[v8]
                            _ = v53
                            _ = v54
                        }
                        var v55 uint32 = 2147483647
                        var v56 float64 = (3.14159 - -2.25)
                        _ = v55
                        _ = v56
                    } else {
                        v2 = v2
                        switch (fi47 & int(v3)) {
                        case 9223372036854775807:
                            v57 := GenS4{F0: -1, F1: -989}
                            _ = v57
                        case -9223372036854775808:
                            var v58 uint8 = 255
                            _ = v58
                        case 32768:
                            var v59 bool = true
                            _ = v59
                        }
                        v1 = "hello"
                    }
                    _ = fi50
                }
                switch v1 {
                case "\n\t":
                    switch fi47 {
                    case 32767:
                        switch ("hello" + (v1 + v1)) {
                        case "hello":
                            var v60 float32 = -2.25
                            _ = v60
                        case "\n\t":
                            fmt.Println(fi47, v8, v7)
                        default:
                            v3 = (-(v9))
                        }
                    case 32768:
                        if false {
                            var v61 bool = true
                            if ((9223372036854775807 != -167) || (v2 == v2)) {
                                for fi62 := 0; fi62 < 1; fi62++ {
                                    v3 -= (v3 | 1)
                                    _ = fi62
                                }
                                v63 := GenS3{F0: uint32(v21), F1: 42}
                                _ = v63
                            }
                            for fi64 := 0; fi64 < 5; fi64++ {
                                fmt.Println(v61, v2)
                                _ = fi64
                            }
                            _ = v61
                        }
                    default:
                        v65 := GenS1{F0: 65535, F1: 9223372036854775807, F2: v3}
                        _ = v65
                    }
                default:
                    var v66 bool = (42 > v2)
                    _ = v66
                }
                _ = fi47
            }
            v8 -= 32767
            v9 += (v9 ^ v3)
        }
        if (!(21 != -114)) {
            for fi67 := 0; fi67 < 8; fi67++ {
                if (true && (100 > v2)) {
                    v8 += v8
                    v2 = v2
                }
                _ = fi67
            }
            var v68 uint8 = 1
            _ = v68
        }
        _ = v45
    }
    switch v7 {
    case 9223372036854775807:
        v7 = (v7 << (((v21 / (((v21 + v21)) | 1))) & 63))
    case -495:
        v69 := []int{}
        _ = v69
    case 128:
        var v70 int64 = -1
        _ = v70
    default:
        v21 = v21
    }
    v71 := []int{-587, (v8 ^ (v8 / ((2147483647) | 1)))}
    v7 -= v7
    v1 = (v1 + v1)
    v4[(v7 | 42)] = 733
    for fi72 := 0; fi72 < 18; fi72++ {
        v73 := GenS1{F0: 2147483648, F1: 0, F2: (v3 >> ((v21) & 63))}
        v71 = append(v71, v7)
        v7 += v8
        _ = fi72
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v71:", len(v71), v71)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v71
    _ = v8
    _ = v9
}
