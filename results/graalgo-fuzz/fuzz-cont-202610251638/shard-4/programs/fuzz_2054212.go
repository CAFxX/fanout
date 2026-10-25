package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 uint16; F2 int }
type GenS5 struct { F0 uint; F1 string; F2 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 uint; F2 uint64 }
type GenS8 struct { F0 bool; F1 int8 }

func main() {
    var v1 int8 = -1
    var v2 bool = false
    var v3 uint64 = 18446744073709551615
    var v4 uint32 = uint32(v3)
    fmt.Println(v4, v1)
    if v2 {
        fmt.Println(v1)
        fmt.Println(v2, v2)
        v5 := GenS1{F0: 32767}
        _ = v5
    } else {
        v6 := GenS1{F0: 32767}
        _ = v6
    }
    if (226 <= 255) {
        for fi7 := 0; fi7 < 10; fi7++ {
            v2 = true
            v4 = (v4 + v4)
            if v2 {
                v1 = (v1 * v1)
            }
            _ = fi7
        }
        v3 -= (v3 % ((v3) | 1))
    }
    var v8 int64 = 0
    switch ("hello" + "\n\t") {
    case "foo bar":
        v9 := GenS1{F0: 320}
        _ = v9
    case "a":
        if true {
            v8 = (^(v8))
            var v10 uint64 = (v3 >> ((660) & 63))
            v11 := GenS1{F0: -32768}
            _ = v10
            _ = v11
        } else {
            var v12 bool = true
            var v13 float32 = ((3.14159 - (1e10 + float32(0.0))) - 0.1)
            v14 := []int{206, 32768, 528}
            _ = v12
            _ = v13
            _ = v14
        }
    case "":
        var v15 bool = ((455 <= 328) || (!(false || (1 <= 65535))))
        _ = v15
    }
    v16 := GenS1{F0: -391}
    if true {
        v17 := []int{}
        var v18 int8 = v1
        _ = v17
        _ = v18
    } else {
        var v19 string = "a"
        v20 := []int{}
        _ = v19
        _ = v20
    }
    switch "" {
    case "foo bar":
        v21 := GenS2{F0: -171, F1: 42, F2: v3}
        _ = v21
    }
    switch 434 {
    case -757:
        v22 := GenS3{F0: v2}
        _ = v22
    case 705:
        var v23 float32 = 1e10
        _ = v23
    }
    v24 := GenS4{F0: 381, F1: 270, F2: 2147483647}
    v3 = 4294967295
    if (18446744073709551615 > 1) {
        if v2 {
            for fi25 := 0; fi25 < 6; fi25++ {
                for fi26 := 0; fi26 < 16; fi26++ {
                    v27 := []int{fi25, fi26}
                    var v28 bool = true
                    v4 += v4
                    v28 = false
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = fi25
            }
            v29 := []int{142, int(v1), -828, 0}
            _ = v29
        }
    }
    var v30 int16 = 42
    v31 := GenS1{F0: 1}
    var v32 float64 = (float64(float32(1e10)) - 0.0)
    switch "\n\t" {
    case "":
        fmt.Println(v32)
    case "hello":
        fmt.Println(v1)
    case "\n\t":
        if v2 {
            switch -125 {
            case 4294967295:
                v33 := GenS3{F0: ((2147483647 == 376) && v2)}
                _ = v33
            case 32768:
                v34 := GenS5{F0: 65535, F1: ("foo bar" + "\n\t"), F2: 127}
                _ = v34
            case -910:
                v35 := []int{0, 32767, -9223372036854775808, -936}
                _ = v35
            default:
                var v36 string = "\n\t"
                _ = v36
            }
        } else {
            v3 += 22
            if true {
                if v2 {
                    switch 42 {
                    case 4294967295:
                        v4 = v4
                    case -100:
                        if ("\n\t" != "\n\t") {
                            fmt.Println(v4, v30, v2)
                        } else {
                            if false {
                                var v37 int8 = 72
                                switch 29 {
                                case 32768:
                                    var v38 int64 = 0
                                    _ = v38
                                case 32767:
                                    if v2 {
                                        var v39 string = "a"
                                        _ = v39
                                    } else {
                                        fmt.Println(v1, v8)
                                        for fi40 := 0; fi40 < 7; fi40++ {
                                            fmt.Println(v32)
                                            fmt.Println(v1, v30, v3)
                                            _ = fi40
                                        }
                                    }
                                default:
                                    v41 := []int{757, 32767, -706}
                                    _ = v41
                                }
                                for fi42 := 0; fi42 < 10; fi42++ {
                                    for fi43 := 0; fi43 < 16; fi43++ {
                                        if v2 {
                                            if false {
                                                var v44 int = fi43
                                                v45 := []int{}
                                                if (!(3.14159 >= 0.1)) {
                                                    fmt.Println(fi42)
                                                }
                                                _ = v44
                                                _ = v45
                                            } else {
                                                v3 = 608
                                            }
                                        } else {
                                            v30 += v30
                                            v46 := []int{(fi42 >> ((621) & 63)), (fi43 & 127)}
                                            if (true && v2) {
                                                var v47 float32 = 3.14159
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v47 -= (0.1 - (v47 / ((float32(1e10) - float32(v32)) / ((((v47 - 1e10) - (float32(((v32 + v32) * (float64(3.14159) - 0.1))) - 0.0)) + (float32((0.0 - (1e10 * (v32 + v32)))) + 3.14159)) - (float32(((-2.25 + (3.14159 + 1e10)) / v32)) - 1e10)))))
                                                case "a":
                                                    var v48 bool = (42 == v3)
                                                    _ = v48
                                                default:
                                                    fmt.Println(v4)
                                                }
                                                _ = v47
                                            }
                                            _ = v46
                                        }
                                        var v49 int = 942
                                        switch 624 {
                                        case 1:
                                            v50 := []int{(fi43 / (((-(v49))) | 1)), 201, 2147483647}
                                            _ = v50
                                        case 4294967295:
                                            v30 = (v30 * v30)
                                        case 0:
                                            v51 := []int{fi42, 90}
                                            _ = v51
                                        default:
                                            fmt.Println(v32, v8)
                                        }
                                        _ = fi43
                                        _ = v49
                                    }
                                    if (("\n\t" <= ("hello" + (((("" + "foo bar") + ("a" + "foo bar")) + "") + "hello"))) || v2) {
                                        var v52 uint = 42
                                        var v53 int32 = 482
                                        var v54 bool = v2
                                        _ = v52
                                        _ = v53
                                        _ = v54
                                    } else {
                                        v8 += (v8 | 1)
                                        switch -9223372036854775808 {
                                        case 32768:
                                            v55 := GenS6{F0: "a"}
                                            _ = v55
                                        case 2147483647:
                                            v56 := GenS6{F0: ("" + (("" + "\n\t") + ""))}
                                            _ = v56
                                        case 506:
                                            if v2 {
                                                var v57 uint64 = v3
                                                v30 += 1
                                                _ = v57
                                            }
                                        }
                                        for fi58 := 0; fi58 < 18; fi58++ {
                                            v59 := []int{}
                                            var v60 int = -226
                                            _ = fi58
                                            _ = v59
                                            _ = v60
                                        }
                                    }
                                    for fi61 := 0; fi61 < 8; fi61++ {
                                        fmt.Println(v32, v32)
                                        v32 = v32
                                        v32 = (((0.0 / 3.14159) + v32) + v32)
                                        v62 := GenS3{F0: true}
                                        switch (fi61 * 32767) {
                                        case 42:
                                            if false {
                                                v63 := GenS7{F0: (-(v30)), F1: 18446744073709551615, F2: 18446744073709551615}
                                                if false {
                                                    v64 := GenS1{F0: (v30 ^ int16(v8))}
                                                    _ = v64
                                                }
                                                _ = v63
                                            }
                                        default:
                                            v2 = v2
                                        }
                                        _ = fi61
                                        _ = v62
                                    }
                                    _ = fi42
                                }
                                _ = v37
                            } else {
                                var v65 int64 = (v8 - -1)
                                for fi66 := 0; fi66 < 13; fi66++ {
                                    fmt.Println(v65, v8)
                                    fmt.Println(v4, fi66, v3)
                                    _ = fi66
                                }
                                _ = v65
                            }
                            for fi67 := 0; fi67 < 6; fi67++ {
                                var v68 float64 = v32
                                switch -328 {
                                case 32767:
                                    var v69 string = ("foo bar" + "a")
                                    _ = v69
                                case 917:
                                    v70 := GenS8{F0: v2, F1: 1}
                                    _ = v70
                                case 128:
                                    var v71 int64 = (v8 << ((27) & 63))
                                    _ = v71
                                default:
                                    v4 = v4
                                }
                                _ = fi67
                                _ = v68
                            }
                        }
                    }
                }
            } else {
                v72 := []int{-987, -2147483648, 32767}
                var v73 int64 = (v8 / ((v8) | 1))
                _ = v72
                _ = v73
            }
        }
    default:
        if v2 {
            for fi74 := 0; fi74 < 8; fi74++ {
                switch "foo bar" {
                case "hello":
                    fmt.Println(v4)
                case "\n\t":
                    fmt.Println(v2, fi74)
                case "a":
                    v3 -= v3
                default:
                    v75 := GenS5{F0: 890, F1: "\n\t", F2: 158}
                    _ = v75
                }
                _ = fi74
            }
            v1 += v1
        }
    }
    v32 = -2.25
    for fi76 := 0; fi76 < 11; fi76++ {
        var v77 int64 = 42
        _ = fi76
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v8
}
