package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int64; F2 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 string; F2 uint }
type GenS8 struct { F0 string; F1 int }
type GenS9 struct { F0 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{}
    var v3 uint = 4294967295
    var v4 int64 = -9223372036854775808
    if true {
        if (127 < 42) {
            fmt.Println(v3, v3)
            for fi5 := 0; fi5 < 16; fi5++ {
                v6 := GenS1{F0: "\n\t", F1: ("\n\t" + ""), F2: 2147483647}
                _ = fi5
                _ = v6
            }
        } else {
            for fi7 := 0; fi7 < 8; fi7++ {
                var v8 uint16 = 1
                v9 := GenS2{F0: 100, F1: -9223372036854775808, F2: 1000}
                v10 := GenS2{F0: 35, F1: v4, F2: 32767}
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            var v11 uint = (^(v3))
            _ = v11
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            _ = len(v2)
            if len(v2) > 0 {
                v2[0] = -9223372036854775808
            }
            for fi13 := 0; fi13 < 20; fi13++ {
                delete(v1, fi13)
                var v14 uint8 = 127
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    }
    v15 := GenS2{F0: 255, F1: (^(v4)), F2: -32768}
    var v16 string = "\n\t"
    switch -9223372036854775808 {
    case 285:
        for fi17 := 0; fi17 < 18; fi17++ {
            if (!((v16 >= "a") && (v16 >= "a"))) {
                delete(v1, (fi17 & fi17))
            }
            for fi18 := 0; fi18 < 2; fi18++ {
                if (!(32767 != -1)) {
                    v16 = (((v16 + ("foo bar" + "a")) + ("a" + "\n\t")) + "hello")
                    v19 := GenS2{F0: 232, F1: (v4 << ((v3) & 63)), F2: 0}
                    if true {
                        var v20 uint32 = 65535
                        fmt.Println(fi17)
                        _ = v20
                    } else {
                        var v21 uint = v3
                        _ = v21
                    }
                    _ = v19
                }
                switch (v16 + v16) {
                case "\n\t":
                    var v22 int8 = 100
                    _ = v22
                case "hello":
                    switch v16 {
                    case "foo bar":
                        if (!(0 == 702)) {
                            var v23 uint64 = 42
                            v24 := GenS1{F0: v16, F1: ((("hello" + v16) + v16) + v16), F2: 65535}
                            v25 := GenS3{F0: 2}
                            v16 = v16
                            v3 += v3
                            _ = v23
                            _ = v24
                            _ = v25
                        } else {
                            v1[-9223372036854775808] = true
                            v26 := GenS3{F0: 2}
                            _ = len(v1)
                            _ = v26
                        }
                    case "\n\t":
                        var v27 uint32 = uint32(v3)
                        _ = v27
                    default:
                        if (986 >= 4294967295) {
                            if true {
                                v28 := GenS4{F0: v3, F1: 84}
                                _ = v28
                            }
                            var v29 uint16 = 255
                            _ = v29
                        } else {
                            v1[0] = ((0.1 + 1e10) == 3.14159)
                            if ((false || (!((-2.25 - (0.0 + 1e10)) != ((0.1 / 1.5) * 1e10)))) || (false || true)) {
                                if true {
                                    v2 = append(v2, 2147483647)
                                    switch ("" + "a") {
                                    case "\n\t":
                                        v3 -= v3
                                    case "foo bar":
                                        if false {
                                            v4 = v4
                                            v16 = v16
                                            v2 = append(v2, fi17)
                                        } else {
                                            v4 -= (v4 - v4)
                                            var v30 uint32 = uint32(v3)
                                            _ = v30
                                        }
                                    case "hello":
                                        switch ("a" + "\n\t") {
                                        case "a":
                                            if true {
                                                fmt.Println(v3, fi17)
                                                switch v16 {
                                                case "foo bar":
                                                    fmt.Println(v3, fi18, v3)
                                                case "":
                                                    v31 := v1[(fi17 << ((v3) & 63))]
                                                    _ = v31
                                                }
                                                var v32 int16 = -32768
                                                _ = v32
                                            } else {
                                                fmt.Println(v16)
                                            }
                                        case "":
                                            v2 = append(v2, (^(fi17)))
                                        case "hello":
                                            _ = len(v1)
                                        default:
                                            var v33 uint8 = 127
                                            _ = v33
                                        }
                                    default:
                                        v34 := GenS5{F0: 42, F1: v4}
                                        _ = v34
                                    }
                                    var v35 int16 = 42
                                    if false {
                                        var v36 uint32 = 2147483648
                                        var v37 uint = 0
                                        v38 := GenS6{F0: (^(v37))}
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    } else {
                                        var v39 uint = (v3 << ((v3) & 63))
                                        var v40 uint8 = 42
                                        var v41 uint32 = 369
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    }
                                    _ = v35
                                }
                            }
                            v42 := GenS7{F0: 347, F1: v16, F2: v3}
                            _ = v42
                        }
                    }
                }
                _ = fi18
            }
            _ = fi17
        }
    case -9223372036854775808:
        var v43 int8 = 42
        _ = v43
    case 2147483647:
        if true {
            v4 = v4
            if len(v2) > 0 {
                v2[0] = 918
            }
            var v44 uint8 = 1
            _ = v44
        } else {
            if (!(65535 < 13)) {
                fmt.Println(v16, v4)
            }
        }
    default:
        fmt.Println(v16, v4, v3)
    }
    var v45 uint64 = 692
    for fi46 := 0; fi46 < 18; fi46++ {
        var v47 int64 = (v4 ^ (v4 / ((v4) | 1)))
        var v48 float32 = (1.5 / (3.14159 / float32((float64(-2.25) / 1e10))))
        for fi49 := 0; fi49 < 14; fi49++ {
            if len(v2) > 0 {
                v2[0] = fi49
            }
            v48 -= (v48 + v48)
            var v50 uint64 = (^(v45))
            v2 = append(v2, fi46)
            var v51 uint8 = 0
            _ = fi49
            _ = v50
            _ = v51
        }
        _ = fi46
        _ = v47
        _ = v48
    }
    v52 := GenS8{F0: v16, F1: -290}
    var v53 bool = true
    if (v53 && v53) {
        v3 = (v3 & v3)
        switch 42 {
        case -2147483648:
            var v54 int64 = v4
            _ = v54
        default:
            for fi55 := 0; fi55 < 15; fi55++ {
                _ = len(v1)
                var v56 float32 = 1e10
                switch (v16 + v16) {
                case "":
                    v4 += (v4 ^ (v4 + (v4 << ((255) & 63))))
                }
                _ = fi55
                _ = v56
            }
        }
    } else {
        for fi57 := 0; fi57 < 1; fi57++ {
            for fi58 := 0; fi58 < 18; fi58++ {
                switch v16 {
                case "":
                    if v53 {
                        if (!false) {
                            v3 = 0
                            var v59 int32 = -1
                            var v60 int64 = v4
                            var v61 float64 = 1.5
                            _ = v59
                            _ = v60
                            _ = v61
                        }
                    }
                case "a":
                    var v62 bool = v53
                    _ = v62
                case "hello":
                    if false {
                        v3 = 65535
                    }
                default:
                    var v63 int8 = 42
                    _ = v63
                }
                _ = fi58
            }
            _ = fi57
        }
    }
    for fi64 := 0; fi64 < 4; fi64++ {
        v45 = v45
        for fi65 := 0; fi65 < 11; fi65++ {
            var v66 uint32 = 799
            _ = fi65
            _ = v66
        }
        _ = fi64
    }
    switch 2147483647 {
    case 32768:
        switch 1 {
        case -9223372036854775808:
            if v53 {
                if (1.5 == 3.14159) {
                    if ("a" < "foo bar") {
                        var v67 uint16 = 2
                        _ = v67
                    }
                    for fi68 := 0; fi68 < 20; fi68++ {
                        v53 = ((((3.14159 * 3.14159) / 0.1) - 0.1) > -2.25)
                        var v69 uint = (v3 * v3)
                        var v70 int64 = v4
                        _ = fi68
                        _ = v69
                        _ = v70
                    }
                    fmt.Println(v45, v3, v4)
                }
            } else {
                v2 = append(v2, 42)
                v3 = 870
                fmt.Println(v16, v45)
            }
        case -815:
            var v71 uint8 = 0
            _ = v71
        case 23:
            var v72 int16 = 42
            _ = v72
        }
    case -318:
        v2 = append(v2, 2147483647)
    case 753:
        _ = len(v1)
    }
    switch 154 {
    case 1:
        switch -9223372036854775808 {
        case 0:
            v1[42] = v53
        case 2147483647:
            for fi73 := 0; fi73 < 3; fi73++ {
                if (2147483647 < 473) {
                    if len(v2) > 0 {
                        v2[0] = (fi73 - (fi73 + (fi73 % ((fi73) | 1))))
                    }
                } else {
                    v53 = ((v53 && (32767 >= 196)) || (!v53))
                    var v74 int64 = v4
                    _ = v74
                }
                _ = fi73
            }
        case -501:
            v16 = "\n\t"
        default:
            if true {
                if (!(-488 >= 32768)) {
                    fmt.Println(v16, v4)
                }
                for fi75 := 0; fi75 < 12; fi75++ {
                    _ = len(v2)
                    var v76 int32 = -978
                    _ = fi75
                    _ = v76
                }
            } else {
                for fi77 := 0; fi77 < 2; fi77++ {
                    if true {
                        if (-116 <= 0) {
                            var v78 bool = ((!(v4 > (v4 << ((v3) & 63)))) || v53)
                            v78 = (v3 <= v3)
                            _ = v78
                        } else {
                            var v79 uint8 = uint8(v45)
                            var v80 uint32 = 65535
                            _ = v79
                            _ = v80
                        }
                        var v81 uint64 = (v45 - (v45 + v45))
                        fmt.Println(fi77, v45, fi77)
                        _ = v81
                    } else {
                        v45 = (v45 | 42)
                        _ = len(v2)
                        v3 -= v3
                    }
                    _ = fi77
                }
                v16 = v16
                v16 = (v16 + "a")
            }
        }
    case 127:
        var v82 int32 = 1
        _ = v82
    case -2147483648:
        var v83 float64 = (0.0 * float64(0.1))
        _ = v83
    }
    v84 := GenS9{F0: 2147483648}
    for fi85 := 0; fi85 < 10; fi85++ {
        if (v53 || v53) {
            switch v16 {
            case "foo bar":
                var v86 bool = v53
                _ = v86
            default:
                v1[(fi85 << ((v3) & 63))] = v53
            }
            switch "a" {
            case "hello":
                v53 = (1e10 != (float32(float64((((0.0 / 0.1) + ((1.5 / 1e10) + -2.25)) - 1e10))) / 0.1))
            case "foo bar":
                v3 -= (v3 / ((v3) | 1))
            case "a":
                v3 += (v3 | v3)
            default:
                v3 += (^(v3))
            }
        }
        fmt.Println(v45, fi85, v4)
        if true {
            var v87 float32 = 3.14159
            _ = v87
        } else {
            fmt.Println(v4, v4, v3)
            var v88 int16 = 320
            fmt.Println(v88, v4)
            _ = v88
        }
        _ = fi85
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v52
    _ = v53
    _ = v84
}
