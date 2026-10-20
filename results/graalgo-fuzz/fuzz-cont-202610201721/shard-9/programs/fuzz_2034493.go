package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int; F2 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 int64; F2 bool }
type GenS6 struct { F0 string; F1 uint }

func main() {
    v1 := []string{((("foo bar" + (("a" + "a") + "hello")) + "a") + "hello"), ("" + ("hello" + ("\n\t" + (((("foo bar" + "foo bar") + "\n\t") + "\n\t") + ("\n\t" + "a"))))), ("\n\t" + "a")}
    var v2 uint16 = 2
    var v3 uint8 = 227
    v4 := make(map[uint32]bool)
    v2 -= v2
    var v5 int8 = 127
    if false {
        var v6 uint16 = (v2 << ((0) & 63))
        _ = v6
    } else {
        _ = len(v4)
        var v7 int = 427
        _ = v7
    }
    if false {
        var v8 int8 = v5
        var v9 bool = (!false)
        fmt.Println(v2, v8)
        _ = v8
        _ = v9
    }
    switch "foo bar" {
    case "foo bar":
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (42 > (v2 % (((v2 ^ (v2 >> ((4294967295) & 63)))) | 1))) {
                if false {
                    v2 = 959
                    for fi10 := 0; fi10 < 2; fi10++ {
                        v11 := GenS1{F0: -1}
                        var v12 int8 = -128
                        var v13 string = "hello"
                        _ = fi10
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    switch -560 {
                    case 127:
                        switch "" {
                        case "foo bar":
                            v14 := GenS1{F0: 1}
                            _ = v14
                        }
                    }
                }
                if (0.0 == ((float32(1e10) - 1.5) * (-2.25 / (1.5 * 1e10)))) {
                    var v15 float64 = 1.5
                    _ = v15
                } else {
                    switch int(v3) {
                    case 2147483647:
                        if (18446744073709551615 <= 42) {
                            v16 := GenS1{F0: -366}
                            _ = v16
                        }
                    case -843:
                        for fi17 := 0; fi17 < 9; fi17++ {
                            switch fi17 {
                            case 127:
                                v5 = (-(v5))
                            case 32768:
                                for fi18 := 0; fi18 < 15; fi18++ {
                                    if true {
                                        v19 := GenS2{F0: (-(fi17)), F1: fi18, F2: 2147483648}
                                        var v20 int = -1
                                        var v21 int16 = 1000
                                        _ = v19
                                        _ = v20
                                        _ = v21
                                    } else {
                                        if true {
                                            v3 = 67
                                        }
                                        if true {
                                            var v22 bool = false
                                            var v23 float32 = 1e10
                                            switch "foo bar" {
                                            case "\n\t":
                                                fmt.Println(v3, fi17, v3)
                                            case "hello":
                                                if false {
                                                    fmt.Println(v22)
                                                }
                                            case "a":
                                                v2 -= (v2 | 42)
                                            default:
                                                v2 -= v2
                                            }
                                            _ = v22
                                            _ = v23
                                        }
                                    }
                                    var v24 int16 = -515
                                    _ = fi18
                                    _ = v24
                                }
                            case 590:
                                fmt.Println(v5, v2, v2)
                            default:
                                fmt.Println(fi17)
                            }
                            for fi25 := 0; fi25 < 9; fi25++ {
                                switch "foo bar" {
                                case "hello":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = fi25
                            }
                            _ = fi17
                        }
                    case -872:
                        switch 158 {
                        case 42:
                            v2 = v2
                        case 127:
                            v5 = v5
                        case 615:
                            v2 = (^(v2))
                        }
                    default:
                        if (!true) {
                            var v26 uint64 = uint64(v2)
                            if false {
                                v27 := GenS3{F0: (-(v5))}
                                v28 := GenS2{F0: -9223372036854775808, F1: 9223372036854775807, F2: 2}
                                _ = v27
                                _ = v28
                            } else {
                                v29 := GenS3{F0: v5}
                                var v30 uint = 255
                                v3 += v3
                                _ = v29
                                _ = v30
                            }
                            v26 += 1
                            _ = v26
                        } else {
                            var v31 uint = 368
                            v32 := GenS2{F0: -1, F1: 541, F2: uint32(v2)}
                            for fi33 := 0; fi33 < 14; fi33++ {
                                v34 := GenS4{F0: v31}
                                _ = fi33
                                _ = v34
                            }
                            _ = v31
                            _ = v32
                        }
                    }
                    v3 = v3
                    if false {
                        for fi35 := 0; fi35 < 12; fi35++ {
                            v36 := GenS4{F0: 68}
                            var v37 uint32 = 2147483647
                            delete(v4, (v37 - 65535))
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                        if (v2 <= v2) {
                            if false {
                                var v38 uint16 = 0
                                var v39 int = -474
                                _ = v38
                                _ = v39
                            }
                        } else {
                            var v40 float32 = 1.5
                            if true {
                                for fi41 := 0; fi41 < 1; fi41++ {
                                    v42 := GenS2{F0: (fi41 - fi41), F1: fi41, F2: 490}
                                    fmt.Println(v3, v3, v2)
                                    _ = fi41
                                    _ = v42
                                }
                            }
                            if true {
                                fmt.Println(v40, v5, v3)
                                if true {
                                    fmt.Println(v2)
                                }
                            } else {
                                for fi43 := 0; fi43 < 9; fi43++ {
                                    for fi44 := 0; fi44 < 9; fi44++ {
                                        var v45 uint32 = 2147483647
                                        if len(v1) > 0 {
                                            v1[0] = "hello"
                                        }
                                        _ = fi44
                                        _ = v45
                                    }
                                    _ = fi43
                                }
                                var v46 int32 = 2147483647
                                _ = v46
                            }
                            _ = v40
                        }
                        v2 = 1
                    }
                }
                for fi47 := 0; fi47 < 12; fi47++ {
                    v2 += 32767
                    _ = fi47
                }
            }
            if (-245 > -723) {
                if false {
                    var v48 int64 = 0
                    switch 2147483647 {
                    case -176:
                        var v49 string = "foo bar"
                        _ = v49
                    }
                    v50 := GenS4{F0: 0}
                    _ = v48
                    _ = v50
                }
            } else {
                if (18446744073709551615 < 0) {
                    var v51 uint16 = (v2 + 583)
                    var v52 int32 = 459
                    _ = v51
                    _ = v52
                } else {
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                }
                _ = len(v1)
            }
        } else {
            var v53 int8 = int8(v3)
            switch 0 {
            case 128:
                for fi54 := 0; fi54 < 16; fi54++ {
                    for fi55 := 0; fi55 < 12; fi55++ {
                        fmt.Println(v53, fi55, v3)
                        _ = fi55
                    }
                    fmt.Println(v5, v2, v5)
                    _ = fi54
                }
            case -2147483648:
                if false {
                    fmt.Println(v3, v5, v5)
                } else {
                    v5 = (v5 / ((-128) | 1))
                    var v56 float32 = ((((1.5 * 1.5) / 1e10) / 1.0) / 0.1)
                    var v57 uint64 = 4294967295
                    _ = v56
                    _ = v57
                }
            }
            _ = v53
        }
    case "":
        if false {
            fmt.Println(v2)
        }
    case "a":
        if false {
            fmt.Println(v2, v3, v5)
        }
    default:
        v58 := GenS5{F0: (-9223372036854775808 >= 42), F1: -1, F2: (true && false)}
        _ = v58
    }
    var v59 int = -1
    if (v2 < (v2 - v2)) {
        _ = len(v4)
        v59 = v59
        switch v59 {
        case -861:
            for fi60 := 0; fi60 < 11; fi60++ {
                delete(v4, 0)
                v61 := GenS1{F0: -241}
                for fi62 := 0; fi62 < 8; fi62++ {
                    switch "" {
                    case "":
                        _ = len(v1)
                    case "\n\t":
                        _ = len(v1)
                    case "a":
                        switch "foo bar" {
                        case "a":
                            if false {
                                var v63 uint64 = 4294967295
                                _ = v63
                            }
                        default:
                            v4[186] = true
                        }
                    default:
                        v3 += (v3 + 1)
                    }
                    if (1 <= -467) {
                        var v64 int32 = 0
                        _ = len(v1)
                        fmt.Println(fi62, v2)
                        _ = v64
                    }
                    _ = fi62
                }
                _ = fi60
                _ = v61
            }
        }
    } else {
        var v65 bool = ("a" != "\n\t")
        var v66 uint64 = 614
        for fi67 := 0; fi67 < 15; fi67++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v68 := GenS6{F0: ("foo bar" + "a"), F1: uint(v2)}
            v69 := GenS1{F0: -9223372036854775808}
            _ = fi67
            _ = v68
            _ = v69
        }
        _ = v65
        _ = v66
    }
    v2 = v2
    switch ("hello" + (("hello" + "hello") + "\n\t")) {
    case "":
        switch (^(v59)) {
        case -530:
            var v70 uint32 = 88
            _ = v70
        case 9223372036854775807:
            var v71 int8 = (v5 << ((255) & 63))
            _ = v71
        }
    default:
        v5 = 127
    }
    v72 := GenS1{F0: 0}
    var v73 float32 = 0.1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v72
    _ = v73
}
