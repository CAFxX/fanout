package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint8; F2 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 string; F2 uint16 }
type GenS6 struct { F0 uint; F1 uint16; F2 int64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint16 }
func genFn0(p0 uint64) uint {
    v6 := GenS1{F0: (p0 - 113), F1: 4294967295, F2: true}
    var v7 float32 = 1.5
    var v8 float32 = float32((-2.25 + (float64((v7 - float32(1e10))) - (-2.25 * 1e10))))
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 136
}

func main() {
    var v1 float32 = (0.0 + 3.14159)
    var v2 string = "\n\t"
    var v3 uint8 = 128
    var v4 string = v2
    v5 := make(map[int]int)
    for fi9 := 0; fi9 < 6; fi9++ {
        switch ((v4 + (v4 + v2)) + v4) {
        case "\n\t":
            v10 := []int{fi9, (fi9 >> ((2) & 63))}
            _ = v10
        }
        var v11 int8 = -128
        var v12 bool = false
        _ = fi9
        _ = v11
        _ = v12
    }
    var v13 int16 = 438
    v5[int(v13)] = 32767
    if true {
        v14 := v5[-154]
        v15 := GenS1{F0: 0, F1: 4294967295, F2: (4294967295 > 706)}
        _ = v14
        _ = v15
    }
    var v16 int32 = 1
    for fi17 := 0; fi17 < 17; fi17++ {
        v18 := GenS2{F0: (^(v3))}
        _ = fi17
        _ = v18
    }
    var v19 uint64 = 490
    var v20 uint = uint(v3)
    for fi21 := 0; fi21 < 11; fi21++ {
        switch (^(fi21)) {
        case 127:
            v22 := GenS3{F0: 42, F1: (^(v3)), F2: 646}
            _ = v22
        }
        v23 := v5[fi21]
        _ = fi21
        _ = v23
    }
    if (!true) {
        if true {
            delete(v5, 4294967295)
            switch 1 {
            case 4294967295:
                var v24 float64 = 1.5
                _ = v24
            default:
                v25 := []int{2147483647, -1, 32767, 128}
                _ = v25
            }
        } else {
            v26 := []int{0}
            var v27 uint32 = 827
            fmt.Println(v1)
            _ = v26
            _ = v27
        }
        v28 := GenS2{F0: 255}
        for fi29 := 0; fi29 < 17; fi29++ {
            v30 := GenS4{F0: (v3 | 241)}
            v20 = v20
            fmt.Println(v2)
            _ = fi29
            _ = v30
        }
        _ = v28
    } else {
        for fi31 := 0; fi31 < 9; fi31++ {
            for fi32 := 0; fi32 < 15; fi32++ {
                fmt.Println(v19, v19, fi31)
                var v33 uint32 = 4294967295
                _ = fi32
                _ = v33
            }
            fmt.Println(v13, v20)
            _ = fi31
        }
        var v34 string = v4
        _ = v34
    }
    v4 = v2
    var v35 float64 = 0.0
    if (((v1 * v1) - 0.0) > v1) {
        switch v2 {
        case "":
            v3 -= v3
        case "\n\t":
            if (v16 != (-(v16))) {
                var v36 int32 = 42
                v13 -= (v13 & v13)
                _ = v36
            } else {
                switch v2 {
                case "\n\t":
                    switch ((v2 + v4) + (v2 + v2)) {
                    case "a":
                        v19 += v19
                    case "":
                        if true {
                            if ((v20 << ((18446744073709551615) & 63)) >= v20) {
                                switch ((v2 + ("foo bar" + ("a" + v4))) + (v4 + "\n\t")) {
                                case "foo bar":
                                    var v37 int8 = -105
                                    _ = v37
                                case "a":
                                    if ((v19 / (((^(v19))) | 1)) != 42) {
                                        var v38 int = 42
                                        var v39 string = v2
                                        _ = v38
                                        _ = v39
                                    }
                                case "hello":
                                    if ((v13 << ((genFn0((v19 * (v19 - v19)))) & 63)) != 1) {
                                        if (-9223372036854775808 < 9223372036854775807) {
                                            if true {
                                                v40 := GenS4{F0: (v3 + (v3 / ((v3) | 1)))}
                                                for fi41 := 0; fi41 < 18; fi41++ {
                                                    for fi42 := 0; fi42 < 11; fi42++ {
                                                        v4 = "hello"
                                                        fmt.Println(v35, v3, v16)
                                                        if false {
                                                            var v43 float32 = v1
                                                            v3 = v3
                                                            fmt.Println(v43)
                                                            v44 := []int{fi41}
                                                            _ = v43
                                                            _ = v44
                                                        }
                                                        _ = fi42
                                                    }
                                                    switch v4 {
                                                    case "foo bar":
                                                        fmt.Println(v13)
                                                    case "a":
                                                        var v45 uint32 = 537
                                                        _ = v45
                                                    default:
                                                        var v46 bool = (2 > 65535)
                                                        _ = v46
                                                    }
                                                    if false {
                                                        v47 := GenS3{F0: 540, F1: (v3 ^ v3), F2: genFn0(v19)}
                                                        _ = v47
                                                    }
                                                    _ = fi41
                                                }
                                                _ = v40
                                            } else {
                                                v48 := GenS3{F0: 1, F1: v3, F2: 0}
                                                v4 = ("hello" + v4)
                                                _ = v48
                                            }
                                        } else {
                                            if (!((-1 < 100) && true)) {
                                                var v49 int8 = 100
                                                _ = len(v5)
                                                _ = v49
                                            } else {
                                                v2 = v4
                                                v2 = ""
                                            }
                                            for fi50 := 0; fi50 < 13; fi50++ {
                                                v51 := GenS2{F0: 198}
                                                if false {
                                                    v3 += v3
                                                    fmt.Println(v16)
                                                }
                                                _ = fi50
                                                _ = v51
                                            }
                                        }
                                        v4 = v2
                                    } else {
                                        v52 := GenS5{F0: 2147483648, F1: v4, F2: uint16(v19)}
                                        _ = v52
                                    }
                                default:
                                    var v53 int16 = v13
                                    _ = v53
                                }
                            }
                        }
                    default:
                        var v54 string = v4
                        _ = v54
                    }
                case "hello":
                    v55 := []int{0, -562}
                    _ = v55
                case "":
                    v56 := GenS6{F0: genFn0(18446744073709551615), F1: uint16(v3), F2: 1}
                    _ = v56
                default:
                    for fi57 := 0; fi57 < 9; fi57++ {
                        fmt.Println(v2)
                        switch int(v3) {
                        case 4294967295:
                            delete(v5, fi57)
                        case 32768:
                            v58 := []int{(fi57 - -35), fi57, fi57, (fi57 ^ (fi57 - fi57))}
                            _ = v58
                        case 1:
                            v59 := GenS7{F0: 919}
                            _ = v59
                        }
                        v1 = 1.5
                        _ = fi57
                    }
                }
            }
        }
        for fi60 := 0; fi60 < 3; fi60++ {
            var v61 uint32 = 0
            fmt.Println(v61)
            _ = fi60
            _ = v61
        }
    } else {
        for fi62 := 0; fi62 < 18; fi62++ {
            if ((-2147483648 != (-(v16))) && true) {
                var v63 uint8 = 255
                _ = v63
            }
            for fi64 := 0; fi64 < 3; fi64++ {
                var v65 uint8 = v3
                _ = fi64
                _ = v65
            }
            v66 := GenS7{F0: v16}
            _ = fi62
            _ = v66
        }
        if (-1 <= 1) {
            delete(v5, 612)
            var v67 int32 = (^(v16))
            _ = v67
        }
    }
    for fi68 := 0; fi68 < 3; fi68++ {
        var v69 int = (-(fi68))
        for fi70 := 0; fi70 < 20; fi70++ {
            v2 = v2
            v1 += 0.0
            _ = fi70
        }
        var v71 float32 = 3.14159
        _ = fi68
        _ = v69
        _ = v71
    }
    switch 128 {
    case -2147483648:
        for fi72 := 0; fi72 < 10; fi72++ {
            for fi73 := 0; fi73 < 2; fi73++ {
                fmt.Println(v19)
                _ = fi73
            }
            v74 := v5[(fi72 >> ((uint(v3)) & 63))]
            _ = fi72
            _ = v74
        }
    default:
        v75 := v5[851]
        _ = v75
    }
    switch 1 {
    case -9223372036854775808:
        for fi76 := 0; fi76 < 20; fi76++ {
            fmt.Println(v4)
            v35 = v35
            v2 = "\n\t"
            _ = fi76
        }
    case 4294967295:
        if true {
            v77 := GenS8{F0: 32767}
            _ = v77
        }
    case 682:
        var v78 int8 = 127
        _ = v78
    default:
        v79 := GenS2{F0: v3}
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
