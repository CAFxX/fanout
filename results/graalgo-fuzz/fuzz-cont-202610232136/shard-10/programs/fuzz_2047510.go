package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 uint16; F1 uint; F2 int64 }
type GenS5 struct { F0 bool; F1 int; F2 bool }
type GenS6 struct { F0 int64; F1 bool; F2 uint16 }
type GenS7 struct { F0 uint32; F1 bool }

func main() {
    var v1 uint64 = 2
    var v2 string = (("" + "hello") + ("" + "hello"))
    v3 := []int{0}
    var v4 uint = 699
    fmt.Println(v4, v1, v1)
    fmt.Println(v2, v2)
    for fi5 := 0; fi5 < 8; fi5++ {
        _ = len(v3)
        _ = fi5
    }
    v6 := GenS1{F0: 199, F1: 42}
    for fi7 := 0; fi7 < 2; fi7++ {
        if ((0 < (v4 & v4)) || true) {
            if (v2 == v2) {
                v8 := GenS2{F0: (1.5 != ((float32(float64((3.14159 + 1.5))) * 1e10) / 3.14159))}
                var v9 int16 = 0
                for fi10 := 0; fi10 < 8; fi10++ {
                    switch fi10 {
                    case 4294967295:
                        if (false && false) {
                            var v11 int64 = -633
                            _ = v11
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                var v12 uint8 = 127
                                _ = v12
                            case "a":
                                v2 = v2
                                v4 = v4
                            default:
                                fmt.Println(v9, v9)
                            }
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                switch fi7 {
                                case 128:
                                    switch 0 {
                                    case -9223372036854775808:
                                        if len(v3) > 0 {
                                            v3[0] = 93
                                        }
                                    }
                                case 576:
                                    var v13 bool = (420 != -2147483648)
                                    _ = v13
                                case -1:
                                    v4 += v4
                                }
                                v4 += (v4 << (((v4 + (v4 + 4294967295))) & 63))
                                if true {
                                    v2 = (v2 + "\n\t")
                                    switch 128 {
                                    case 1:
                                        fmt.Println(v1)
                                    default:
                                        var v14 uint32 = 111
                                        var v15 int32 = -2147483648
                                        _ = v14
                                        _ = v15
                                    }
                                } else {
                                    var v16 bool = (false && (!false))
                                    v9 = (v9 / (((v9 << ((v4) & 63))) | 1))
                                    v9 = (v9 * 1000)
                                    _ = v16
                                }
                                switch -796 {
                                case 681:
                                    v9 = (v9 + 1)
                                }
                                v17 := GenS3{F0: v2, F1: 975}
                                _ = v17
                            }
                            switch (("a" + (v2 + (("a" + v2) + "hello"))) + (v2 + v2)) {
                            case "":
                                v2 = ""
                            case "foo bar":
                                var v18 uint64 = (v1 >> (((^(v4))) & 63))
                                _ = v18
                            default:
                                var v19 string = v2
                                _ = v19
                            }
                        }
                    case -700:
                        if ((fi10 - (^(fi7))) < (fi7 & fi10)) {
                            v4 = uint(v1)
                            if (4294967295 != 618) {
                                var v20 uint = v4
                                _ = v20
                            } else {
                                var v21 uint8 = 127
                                var v22 uint32 = 2147483648
                                _ = v21
                                _ = v22
                            }
                            switch (fi7 * fi7) {
                            case 1:
                                v23 := GenS4{F0: 2, F1: v4, F2: 42}
                                _ = v23
                            case 2147483647:
                                v4 = (v4 & (v4 << ((v4) & 63)))
                            }
                        } else {
                            var v24 int8 = 0
                            v25 := GenS1{F0: 2147483647, F1: -301}
                            _ = v24
                            _ = v25
                        }
                    case -9223372036854775808:
                        v1 += v1
                    default:
                        if len(v3) > 0 {
                            v3[0] = fi10
                        }
                        switch (v2 + ("foo bar" + "")) {
                        case "a":
                            if (!(-1 != -1)) {
                                var v26 int16 = v9
                                v3 = append(v3, (fi10 << (((v4 >> (((v4 | (v4 << ((v4) & 63)))) & 63))) & 63)))
                                var v27 uint = v4
                                _ = v26
                                _ = v27
                            }
                        case "":
                            var v28 int32 = 1
                            _ = v28
                        default:
                            if (v1 < (v1 ^ v1)) {
                                v2 = v2
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = (fi10 ^ fi7)
                                }
                                switch 573 {
                                case 411:
                                    var v29 int64 = 9223372036854775807
                                    _ = v29
                                }
                            }
                        }
                    }
                    v30 := GenS5{F0: ((!false) || ((1.5 + 1.5) < (0.1 / 1.5))), F1: (^(fi10)), F2: false}
                    _ = fi10
                    _ = v30
                }
                _ = v8
                _ = v9
            } else {
                if (!true) {
                    var v31 uint8 = 242
                    _ = v31
                } else {
                    var v32 uint32 = 4294967295
                    _ = len(v3)
                    v4 = v4
                    _ = v32
                }
                v33 := GenS1{F0: 2, F1: -46}
                v1 += v1
                _ = v33
            }
        }
        v1 += 4294967295
        v2 = v2
        _ = fi7
    }
    var v34 float32 = (-2.25 + -2.25)
    for fi35 := 0; fi35 < 9; fi35++ {
        switch v2 {
        case "hello":
            if (!(32767 < 752)) {
                if (float64(v34) < (((0.0 * float64((v34 - v34))) / 1e10) * float64((3.14159 / (float32(((0.0 / 1.5) * float64((v34 - float32(3.14159))))) + 1.5))))) {
                    for fi36 := 0; fi36 < 13; fi36++ {
                        var v37 string = v2
                        v3 = append(v3, (fi36 % ((fi36) | 1)))
                        fmt.Println(v4, v4)
                        _ = fi36
                        _ = v37
                    }
                    switch fi35 {
                    case 42:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case 4294967295:
                        if true {
                            if len(v3) > 0 {
                                v3[0] = fi35
                            }
                        } else {
                            v38 := GenS6{F0: 1, F1: (42 == 327), F2: 255}
                            v39 := GenS7{F0: 653, F1: (!false)}
                            _ = v38
                            _ = v39
                        }
                    case 32768:
                        var v40 float64 = ((3.14159 * -2.25) * 1e10)
                        _ = v40
                    }
                    for fi41 := 0; fi41 < 17; fi41++ {
                        if len(v3) > 0 {
                            v3[0] = fi35
                        }
                        var v42 int64 = -9223372036854775808
                        _ = len(v3)
                        _ = fi41
                        _ = v42
                    }
                } else {
                    if true {
                        v43 := GenS7{F0: 4294967295, F1: (42 <= 127)}
                        var v44 uint = 4294967295
                        _ = v43
                        _ = v44
                    } else {
                        _ = len(v3)
                    }
                    v1 = 1
                    for fi45 := 0; fi45 < 6; fi45++ {
                        v3 = append(v3, (fi35 + fi35))
                        v4 = uint(v1)
                        _ = fi45
                    }
                }
                for fi46 := 0; fi46 < 1; fi46++ {
                    if len(v3) > 0 {
                        v3[0] = fi46
                    }
                    _ = fi46
                }
                switch fi35 {
                case 127:
                    v2 = ""
                case 128:
                    if false {
                        fmt.Println(v34)
                        if (!false) {
                            v47 := GenS7{F0: 739, F1: (!(!false))}
                            _ = v47
                        } else {
                            v34 = v34
                        }
                        v48 := GenS3{F0: ("hello" + "\n\t"), F1: 255}
                        _ = v48
                    }
                case 459:
                    var v49 uint64 = 42
                    _ = v49
                }
            } else {
                fmt.Println(v4, v4, v34)
                for fi50 := 0; fi50 < 3; fi50++ {
                    var v51 int16 = 0
                    v52 := GenS6{F0: 1, F1: ((((0.1 - (1.5 * 0.1)) / 1e10) / -2.25) <= ((1e10 - (0.1 + 3.14159)) + float64((float32(((0.1 * float64(0.1)) / 0.1)) - v34)))), F2: 32767}
                    if (1 > 0) {
                        var v53 int8 = int8(v51)
                        fmt.Println(fi35, v4)
                        var v54 uint16 = 1
                        v4 = 977
                        _ = v53
                        _ = v54
                    } else {
                        var v55 float32 = 1.5
                        if true {
                            v56 := GenS2{F0: true}
                            fmt.Println(v51, v34, fi35)
                            if false {
                                if true {
                                    fmt.Println(v51, v51)
                                    var v57 float32 = v34
                                    v58 := GenS5{F0: true, F1: (fi35 ^ (fi35 % ((0) | 1))), F2: (538 != fi35)}
                                    _ = v57
                                    _ = v58
                                }
                                if (0 == 995) {
                                    if false {
                                        var v59 uint = (v4 ^ v4)
                                        _ = len(v3)
                                        var v60 uint64 = v1
                                        _ = v59
                                        _ = v60
                                    }
                                }
                                v1 = uint64(v4)
                            }
                            _ = v56
                        }
                        _ = v55
                    }
                    _ = fi50
                    _ = v51
                    _ = v52
                }
            }
        case "\n\t":
            v4 = v4
        }
        fmt.Println(v2, v34)
        _ = fi35
    }
    var v61 bool = (((!(false || (!false))) || true) || false)
    for fi62 := 0; fi62 < 14; fi62++ {
        if (((1.5 - (v34 * ((v34 * (v34 - (-2.25 / (float32(1.5) / (0.1 + float32(1e10)))))) * 3.14159))) + 1.5) < (v34 / v34)) {
            var v63 uint16 = 568
            _ = v63
        } else {
            for fi64 := 0; fi64 < 18; fi64++ {
                if false {
                    fmt.Println(v1, v1)
                    var v65 int32 = 1
                    _ = v65
                } else {
                    v61 = false
                }
                _ = fi64
            }
        }
        _ = fi62
    }
    if ((!(0.0 == -2.25)) || (!(((^(v4)) <= 2) || ((^(v4)) <= (v4 >> ((2) & 63)))))) {
        var v66 string = ("\n\t" + v2)
        if (true || false) {
            var v67 int16 = 823
            for fi68 := 0; fi68 < 3; fi68++ {
                for fi69 := 0; fi69 < 6; fi69++ {
                    if false {
                        var v70 int32 = 309
                        _ = v70
                    }
                    _ = fi69
                }
                _ = fi68
            }
            switch "foo bar" {
            case "\n\t":
                if (!v61) {
                    v1 = 42
                    var v71 int64 = -437
                    _ = v71
                }
            case "a":
                v1 -= 1
            }
            _ = v67
        } else {
            var v72 int64 = 0
            _ = v72
        }
        _ = v66
    } else {
        fmt.Println(v2, v61)
    }
    fmt.Println(v1, v1)
    switch v2 {
    case "hello":
        fmt.Println(v4, v2, v34)
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "\n\t":
        switch ((v2 + "hello") + "a") {
        case "hello":
            v73 := GenS7{F0: uint32(v1), F1: true}
            _ = v73
        case "":
            var v74 float64 = (1e10 - 3.14159)
            _ = v74
        case "\n\t":
            for fi75 := 0; fi75 < 18; fi75++ {
                fmt.Println(v4, v4, v4)
                v61 = false
                fmt.Println(fi75, v61, v34)
                _ = fi75
            }
        }
    default:
        v2 = ((v2 + v2) + "foo bar")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v6
    _ = v61
}
