package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int16; F1 bool; F2 int }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int64 }

func main() {
    var v1 int32 = -1
    var v2 int = 0
    v3 := []string{"\n\t", "foo bar", "foo bar"}
    var v4 int64 = 42
    var v5 int32 = v1
    if ("\n\t" != "") {
        switch 128 {
        case 32768:
            var v6 uint32 = 4294967295
            _ = v6
        case 695:
            switch ("" + "foo bar") {
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "\n\t":
                v4 += v4
            default:
                var v7 uint8 = 1
                _ = v7
            }
        case 1:
            if false {
                v8 := GenS1{F0: (v2 << ((1) & 63)), F1: 2, F2: 128}
                for fi9 := 0; fi9 < 5; fi9++ {
                    var v10 uint16 = 2
                    var v11 uint16 = 65535
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                fmt.Println(v1)
                _ = v8
            } else {
                if ((18446744073709551615 > 294) && (((!false) || true) && true)) {
                    var v12 int8 = -1
                    var v13 bool = (((!(((((-2.25 * 3.14159) / ((((float64(1.5) - 0.0) + 0.0) / 3.14159) / float64(1.5))) / 1e10) * (0.1 + ((-2.25 / float64((0.1 * (((float32(0.1) - float32(float64(float32(((1e10 - 3.14159) + 1.5))))) - -2.25) + -2.25)))) + -2.25))) > (0.0 * 1.5))) && ((^(v5)) == v5)) && true)
                    _ = v12
                    _ = v13
                }
                var v14 int64 = (v4 << ((42) & 63))
                v5 = (v5 + (v1 & v5))
                _ = v14
            }
        }
    }
    v15 := GenS2{F0: 664, F1: (!(2147483648 != 0)), F2: v2}
    fmt.Println(v1, v2, v4)
    var v16 int64 = (v4 % ((-9223372036854775808) | 1))
    var v17 int16 = 32767
    var v18 uint = 857
    if true {
        v5 = 0
    }
    fmt.Println(v18, v5)
    var v19 int = int(v4)
    var v20 uint64 = 673
    v19 += -9223372036854775808
    switch int(v17) {
    case 4294967295:
        if false {
            if true {
                for fi21 := 0; fi21 < 4; fi21++ {
                    var v22 int32 = (v1 - int32(v19))
                    _ = fi21
                    _ = v22
                }
                var v23 float32 = 0.0
                _ = v23
            }
            switch "hello" {
            case "a":
                var v24 string = "\n\t"
                _ = v24
            case "":
                if (v18 < v18) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    for fi25 := 0; fi25 < 2; fi25++ {
                        for fi26 := 0; fi26 < 13; fi26++ {
                            var v27 uint8 = uint8(v20)
                            _ = fi26
                            _ = v27
                        }
                        if len(v3) > 0 {
                            v3[0] = ("" + "hello")
                        }
                        _ = fi25
                    }
                    switch ("foo bar" + ("hello" + ("hello" + ("a" + "")))) {
                    case "\n\t":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "foo bar":
                        v16 = (v4 ^ (v4 + int64(v20)))
                    case "":
                        for fi28 := 0; fi28 < 6; fi28++ {
                            v4 += -9223372036854775808
                            _ = fi28
                        }
                    default:
                        var v29 uint64 = v20
                        _ = v29
                    }
                    v30 := GenS3{F0: -760, F1: v4}
                    _ = v30
                }
            case "\n\t":
                v16 -= v4
            }
        }
    case -2147483648:
        v4 = v4
    default:
        switch int(v18) {
        case 1:
            var v31 int8 = 84
            _ = v31
        case 4294967295:
            if (true || true) {
                if (2 < 136) {
                    switch v19 {
                    case -9223372036854775808:
                        v32 := GenS4{F0: v20, F1: 255}
                        _ = v32
                    case 2147483647:
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                    case -695:
                        v16 += -559
                    }
                    for fi33 := 0; fi33 < 5; fi33++ {
                        var v34 float64 = -2.25
                        var v35 int64 = -1
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    if false {
                        for fi36 := 0; fi36 < 4; fi36++ {
                            if (!false) {
                                var v37 int16 = 1000
                                _ = v37
                            } else {
                                if true {
                                    v38 := GenS5{F0: 221, F1: v4}
                                    _ = v38
                                } else {
                                    v18 = v18
                                    switch (v2 ^ v2) {
                                    case 810:
                                        v16 = v16
                                    }
                                    var v39 uint32 = 449
                                    _ = v39
                                }
                            }
                            v40 := GenS3{F0: -1, F1: v16}
                            _ = fi36
                            _ = v40
                        }
                        v17 = (v17 % ((int16(v16)) | 1))
                        var v41 uint32 = 1
                        _ = v41
                    } else {
                        for fi42 := 0; fi42 < 12; fi42++ {
                            switch fi42 {
                            case 127:
                                if len(v3) > 0 {
                                    v3[0] = "hello"
                                }
                            }
                            var v43 string = "foo bar"
                            _ = fi42
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 13; fi44++ {
                            fmt.Println(v20)
                            if len(v3) > 0 {
                                v3[0] = "foo bar"
                            }
                            v45 := GenS5{F0: 42, F1: int64(v5)}
                            _ = fi44
                            _ = v45
                        }
                    }
                }
            } else {
                switch (-(v2)) {
                case 0:
                    for fi46 := 0; fi46 < 12; fi46++ {
                        for fi47 := 0; fi47 < 19; fi47++ {
                            if true {
                                var v48 int64 = v4
                                var v49 int64 = (v16 / ((v16) | 1))
                                var v50 int8 = 42
                                _ = v48
                                _ = v49
                                _ = v50
                            } else {
                                v51 := GenS5{F0: 127, F1: (v16 ^ v16)}
                                _ = v51
                            }
                            _ = fi47
                        }
                        var v52 int8 = 42
                        _ = fi46
                        _ = v52
                    }
                case 69:
                    switch (-(v2)) {
                    case 9223372036854775807:
                        if (true || false) {
                            v17 = (v17 | v17)
                            v20 -= v20
                        }
                    case 288:
                        for fi53 := 0; fi53 < 16; fi53++ {
                            if false {
                                v54 := GenS5{F0: 0, F1: v16}
                                fmt.Println(v1)
                                for fi55 := 0; fi55 < 15; fi55++ {
                                    var v56 uint16 = 0
                                    _ = fi55
                                    _ = v56
                                }
                                _ = v54
                            }
                            if (v20 == v20) {
                                v2 = (v2 % ((-9223372036854775808) | 1))
                            } else {
                                if (false || false) {
                                    for fi57 := 0; fi57 < 11; fi57++ {
                                        var v58 int32 = -248
                                        if ((128 < 168) || ("a" < "hello")) {
                                            v59 := GenS2{F0: (v17 % (((v17 % ((32767) | 1))) | 1)), F1: (!((v2 % ((int(v18)) | 1)) != (v2 ^ (v19 * fi53)))), F2: -503}
                                            _ = v59
                                        }
                                        var v60 uint8 = 1
                                        v4 = (-(v16))
                                        _ = fi57
                                        _ = v58
                                        _ = v60
                                    }
                                }
                                fmt.Println(v1)
                                switch ("\n\t" + "") {
                                case "a":
                                    if (false && false) {
                                        fmt.Println(v5, v20)
                                    } else {
                                        var v61 bool = (!(true || true))
                                        v5 += v5
                                        _ = v61
                                    }
                                case "hello":
                                    for fi62 := 0; fi62 < 17; fi62++ {
                                        v16 = -1
                                        fmt.Println(fi62, v18, v19)
                                        v63 := GenS2{F0: 0, F1: (!false), F2: fi62}
                                        _ = fi62
                                        _ = v63
                                    }
                                case "":
                                    var v64 uint8 = 119
                                    _ = v64
                                default:
                                    v65 := GenS1{F0: (fi53 / (((fi53 >> (((v18 >> (((v18 & v18)) & 63))) & 63))) | 1)), F1: 1, F2: 255}
                                    _ = v65
                                }
                            }
                            _ = fi53
                        }
                    default:
                        var v66 int64 = int64(v18)
                        _ = v66
                    }
                case 32768:
                    switch "foo bar" {
                    case "":
                        v67 := GenS2{F0: (v17 * (^(v17))), F1: (231 == (v17 - 32767)), F2: (v19 << (((^(v18))) & 63))}
                        _ = v67
                    case "a":
                        var v68 string = ""
                        _ = v68
                    case "foo bar":
                        v69 := GenS3{F0: (^(v17)), F1: -422}
                        _ = v69
                    default:
                        var v70 uint32 = 382
                        _ = v70
                    }
                }
            }
        default:
            v20 = (v20 / ((v20) | 1))
        }
    }
    var v71 uint64 = 446
    v17 = v17
    if (!(!true)) {
        v72 := GenS3{F0: v17, F1: 328}
        var v73 int16 = 1
        fmt.Println(v73, v19)
        _ = v72
        _ = v73
    }
    fmt.Println(v4, v4)
    v3 = append(v3, ("hello" + "hello"))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v71
}
