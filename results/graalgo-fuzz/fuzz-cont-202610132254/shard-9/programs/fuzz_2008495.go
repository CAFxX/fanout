package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint8; F1 uint32; F2 int16 }

func main() {
    var v1 uint64 = 926
    var v2 int32 = -389
    v3 := []uint32{4294967295, 2, 42}
    var v4 uint = 18446744073709551615
    if len(v3) > 0 {
        v3[0] = 42
    }
    v5 := GenS1{F0: 2147483648}
    for fi6 := 0; fi6 < 11; fi6++ {
        if ((uint32(v4) > 1) && (!true)) {
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 uint32 = 2
                var v9 string = ("\n\t" + "hello")
                v1 = v1
                _ = len(v3)
                _ = fi7
                _ = v8
                _ = v9
            }
        } else {
            switch "" {
            case "foo bar":
                fmt.Println(v1)
            case "\n\t":
                switch "hello" {
                case "foo bar":
                    v2 = v2
                case "\n\t":
                    var v10 bool = false
                    _ = v10
                case "hello":
                    v11 := GenS2{F0: 65535}
                    _ = v11
                }
            default:
                v12 := GenS1{F0: 236}
                _ = v12
            }
            v13 := GenS3{F0: v2, F1: 128, F2: (v1 >> ((v4) & 63))}
            _ = v13
        }
        v2 = (v2 * 42)
        _ = fi6
    }
    fmt.Println(v4, v2, v1)
    fmt.Println(v2)
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 int8 = 0
        _ = fi14
        _ = v15
    }
    var v16 int8 = 100
    var v17 bool = false
    for fi18 := 0; fi18 < 6; fi18++ {
        v19 := GenS4{F0: 602}
        _ = fi18
        _ = v19
    }
    v17 = (!(-817 >= 1000))
    fmt.Println(v17)
    switch "foo bar" {
    case "a":
        if ((^(v4)) == 279) {
            fmt.Println(v16, v17)
            if (244 <= 127) {
                for fi20 := 0; fi20 < 20; fi20++ {
                    var v21 int16 = 42
                    var v22 bool = v17
                    fmt.Println(fi20, v4)
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                v2 = v2
                v17 = ("foo bar" > ((("hello" + ("a" + "foo bar")) + "hello") + "a"))
            }
            fmt.Println(v16, v4, v1)
        } else {
            var v23 uint32 = 958
            for fi24 := 0; fi24 < 19; fi24++ {
                switch "" {
                case "foo bar":
                    if v17 {
                        for fi25 := 0; fi25 < 20; fi25++ {
                            v26 := GenS1{F0: (v23 / (((v23 % ((2147483648) | 1))) | 1))}
                            var v27 int32 = 0
                            var v28 int16 = 0
                            var v29 uint32 = v23
                            _ = len(v3)
                            _ = fi25
                            _ = v26
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        if (fi24 <= int(v2)) {
                            for fi30 := 0; fi30 < 17; fi30++ {
                                v2 -= (^(v2))
                                _ = fi30
                            }
                            for fi31 := 0; fi31 < 7; fi31++ {
                                v17 = v17
                                switch "" {
                                case "hello":
                                    var v32 uint32 = v23
                                    _ = v32
                                case "a":
                                    v33 := GenS5{F0: ("hello" + ""), F1: v4}
                                    _ = v33
                                case "\n\t":
                                    fmt.Println(v17, v16, v1)
                                }
                                fmt.Println(v16)
                                _ = fi31
                            }
                        }
                        var v34 float32 = -2.25
                        _ = v34
                    } else {
                        if false {
                            fmt.Println(v1, v1, fi24)
                            if (((1e10 / 1.0) / 1e10) < 1e10) {
                                if (v16 < 0) {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    fmt.Println(v16)
                                }
                            }
                        } else {
                            v35 := GenS4{F0: 32767}
                            for fi36 := 0; fi36 < 20; fi36++ {
                                fmt.Println(v23, v23, v2)
                                var v37 int32 = 0
                                _ = fi36
                                _ = v37
                            }
                            v16 = (v16 - v16)
                            _ = v35
                        }
                        if ((0 < v1) || ((fi24 * (fi24 / ((127) | 1))) <= (fi24 >> ((v4) & 63)))) {
                            var v38 uint32 = 398
                            v17 = ((v2 << ((659) & 63)) > -2147483648)
                            _ = v38
                        } else {
                            var v39 int64 = 42
                            var v40 uint = v4
                            _ = v39
                            _ = v40
                        }
                        if ((1 > -1) || ("hello" != "hello")) {
                            switch fi24 {
                            case 128:
                                if false {
                                    var v41 int32 = 0
                                    v3 = append(v3, 42)
                                    var v42 int64 = -1
                                    _ = v41
                                    _ = v42
                                } else {
                                    var v43 uint = 42
                                    var v44 uint32 = v23
                                    _ = v43
                                    _ = v44
                                }
                            }
                        } else {
                            v45 := GenS2{F0: v23}
                            v16 = (v16 >> ((v4) & 63))
                            for fi46 := 0; fi46 < 9; fi46++ {
                                switch (fi46 - int(v2)) {
                                case 0:
                                    if (v17 && (42 == uint8(v1))) {
                                        var v47 bool = v17
                                        switch (-(fi46)) {
                                        case 505:
                                            v23 += (v23 >> (((^(v4))) & 63))
                                        case 47:
                                            v23 += v23
                                        case 0:
                                            v4 = (v4 << (((v4 - (v4 & 2))) & 63))
                                        }
                                        _ = v47
                                    }
                                default:
                                    v1 -= 4294967295
                                }
                                _ = fi46
                            }
                            _ = v45
                        }
                    }
                case "hello":
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                default:
                    var v48 int8 = 42
                    _ = v48
                }
                for fi49 := 0; fi49 < 18; fi49++ {
                    v4 = 546
                    v50 := GenS4{F0: -1}
                    _ = fi49
                    _ = v50
                }
                _ = fi24
            }
            switch "foo bar" {
            case "":
                if false {
                    if v17 {
                        if v17 {
                            fmt.Println(v17, v1, v2)
                        }
                        switch (("" + "hello") + "a") {
                        case "":
                            for fi51 := 0; fi51 < 4; fi51++ {
                                fmt.Println(v16, v1, v17)
                                _ = fi51
                            }
                        }
                    }
                } else {
                    fmt.Println(v2)
                    var v52 uint = (^(v4))
                    if (18446744073709551615 < (^(v1))) {
                        for fi53 := 0; fi53 < 8; fi53++ {
                            if (!false) {
                                for fi54 := 0; fi54 < 10; fi54++ {
                                    v52 = 2
                                    switch ("hello" + "\n\t") {
                                    case "a":
                                        v55 := GenS3{F0: -207, F1: 198, F2: 4294967295}
                                        _ = v55
                                    }
                                    v56 := GenS6{F0: (v1 << (((^(v52))) & 63)), F1: ("hello" + "hello")}
                                    _ = fi54
                                    _ = v56
                                }
                                switch int(v4) {
                                case -1:
                                    for fi57 := 0; fi57 < 16; fi57++ {
                                        v1 += 2
                                        _ = fi57
                                    }
                                case -2147483648:
                                    _ = len(v3)
                                }
                                v58 := GenS7{F0: v1}
                                _ = v58
                            } else {
                                var v59 uint = 1
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v60 int = fi53
                                _ = v59
                                _ = v60
                            }
                            v1 += 0
                            _ = fi53
                        }
                        v52 -= 42
                    }
                    _ = v52
                }
            }
            _ = v23
        }
    default:
        var v61 uint32 = 0
        _ = v61
    }
    switch -1 {
    case 2147483647:
        var v62 float32 = 1e10
        _ = v62
    case 128:
        for fi63 := 0; fi63 < 18; fi63++ {
            var v64 uint16 = 0
            if (!v17) {
                v1 = (v1 << ((1) & 63))
                v65 := GenS5{F0: ("hello" + (("foo bar" + "hello") + "\n\t")), F1: 726}
                _ = v65
            }
            fmt.Println(v16, fi63, v4)
            _ = fi63
            _ = v64
        }
    default:
        v2 = (^(v2))
    }
    if (!(-2147483648 >= (v2 << (((v4 & (v4 >> ((1) & 63)))) & 63)))) {
        for fi66 := 0; fi66 < 1; fi66++ {
            switch 0 {
            case 0:
                if (1e10 == 1e10) {
                    v4 = v4
                    fmt.Println(fi66)
                    var v67 uint16 = uint16(v1)
                    _ = v67
                } else {
                    if (!v17) {
                        fmt.Println(v4, v17, v4)
                        v1 = 0
                    } else {
                        v68 := GenS2{F0: 2}
                        for fi69 := 0; fi69 < 18; fi69++ {
                            v1 -= uint64(v4)
                            if len(v3) > 0 {
                                v3[0] = 816
                            }
                            _ = fi69
                        }
                        var v70 float32 = (3.14159 + 0.1)
                        _ = v68
                        _ = v70
                    }
                }
            default:
                if v17 {
                    var v71 string = ""
                    var v72 int64 = -418
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                    _ = v71
                    _ = v72
                } else {
                    v1 = v1
                    if (false && (-2147483648 != (v2 | (^(v2))))) {
                        v73 := GenS5{F0: "foo bar", F1: v4}
                        _ = v73
                    }
                    fmt.Println(fi66, v4)
                }
            }
            var v74 int32 = (v2 | v2)
            v75 := GenS8{F0: 255, F1: 42, F2: 939}
            _ = fi66
            _ = v74
            _ = v75
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
