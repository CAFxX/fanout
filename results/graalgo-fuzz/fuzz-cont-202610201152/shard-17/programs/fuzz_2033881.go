package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 string; F1 int; F2 string }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int32; F1 int16; F2 int }

func main() {
    v1 := []uint32{}
    var v2 bool = true
    v3 := []uint32{541, 1, 42, 42}
    var v4 float64 = -2.25
    var v5 int64 = -1
    _ = len(v3)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = v4
    if (-9223372036854775808 != -429) {
        var v6 uint16 = 1
        _ = v6
    } else {
        v5 -= 359
        for fi7 := 0; fi7 < 10; fi7++ {
            v2 = (("" <= "") || v2)
            fmt.Println(v2, v5)
            var v8 int32 = 282
            _ = fi7
            _ = v8
        }
    }
    var v9 int16 = -762
    if len(v1) > 0 {
        v1[0] = 42
    }
    var v10 int32 = 0
    switch 0 {
    case -447:
        fmt.Println(v9, v9)
    case -9223372036854775808:
        switch ("\n\t" + "a") {
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "":
            v5 = 42
        }
    default:
        v2 = true
    }
    if ((290 != 42) || v2) {
        switch 4294967295 {
        case 32768:
            var v11 float64 = 1.5
            _ = v11
        case 2147483647:
            fmt.Println(v10, v2)
        case 0:
            if v2 {
                var v12 string = ""
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v13 := GenS1{F0: v10, F1: 32768}
                _ = v12
                _ = v13
            }
        }
        v4 = ((v4 + 1.5) / -2.25)
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        if v2 {
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 float32 = 3.14159
                if len(v3) > 0 {
                    v3[0] = 1
                }
                fmt.Println(fi14, fi14)
                if true {
                    if v2 {
                        var v17 float32 = ((1e10 / (((v16 / v16) + (v16 * (0.0 - (1.5 / (v16 + v16))))) - (v16 + ((v16 + v16) / (0.1 / v16))))) * float32((v4 * ((((0.0 / 1.0) / float64(v16)) + float64(v16)) - v4))))
                        v18 := GenS2{F0: "a", F1: (fi14 & (fi15 - (^(fi15)))), F2: ""}
                        var v19 bool = true
                        var v20 float64 = v4
                        if v19 {
                            if v19 {
                                v2 = v19
                                var v21 int64 = (v5 / (((^(v5))) | 1))
                                v22 := GenS2{F0: "", F1: (fi15 ^ (fi15 % ((fi15) | 1))), F2: "\n\t"}
                                if true {
                                    var v23 int = (^(fi14))
                                    _ = v23
                                }
                                if v19 {
                                    switch "" {
                                    case "":
                                        if len(v3) > 0 {
                                            v3[0] = 923
                                        }
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 65535
                                    }
                                    v1 = append(v1, 448)
                                    switch fi14 {
                                    case 905:
                                        var v24 int32 = v10
                                        _ = v24
                                    case 127:
                                        v25 := GenS2{F0: "hello", F1: fi15, F2: "hello"}
                                        _ = v25
                                    case -516:
                                        if (2147483647 != 531) {
                                            var v26 int64 = v21
                                            var v27 uint16 = 126
                                            _ = v26
                                            _ = v27
                                        }
                                    default:
                                        v16 += 0.1
                                    }
                                    var v28 uint64 = 42
                                    v3 = append(v3, 0)
                                    _ = v28
                                }
                                _ = v21
                                _ = v22
                            }
                        }
                        _ = v17
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        fmt.Println(fi14, fi15)
                        v29 := GenS3{F0: int16(v10), F1: (v5 & -88)}
                        if v2 {
                            if len(v3) > 0 {
                                v3[0] = 65535
                            }
                        } else {
                            v5 -= (v5 / ((v5) | 1))
                        }
                        _ = v29
                    }
                }
                _ = len(v1)
                _ = fi15
                _ = v16
            }
            v30 := GenS2{F0: "\n\t", F1: fi14, F2: ("a" + ("hello" + "a"))}
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            _ = v30
        } else {
            var v31 uint64 = 1
            var v32 uint8 = 2
            _ = v31
            _ = v32
        }
        _ = fi14
    }
    if v2 {
        v4 -= ((float64(3.14159) / 1.0) / v4)
        var v33 int8 = 93
        v34 := GenS4{F0: v33, F1: v33}
        _ = v33
        _ = v34
    }
    _ = len(v3)
    for fi35 := 0; fi35 < 17; fi35++ {
        v36 := GenS4{F0: int8(fi35), F1: 1}
        for fi37 := 0; fi37 < 19; fi37++ {
            if (0.0 > -2.25) {
                var v38 int8 = -1
                v9 = (v9 - (v9 % ((-32768) | 1)))
                fmt.Println(v5, v5)
                _ = v38
            } else {
                if (1 <= 278) {
                    var v39 int32 = (v10 % ((-2147483648) | 1))
                    _ = v39
                }
                if (0 <= 65535) {
                    switch (fi35 >> ((255) & 63)) {
                    case 9223372036854775807:
                        v40 := GenS1{F0: -2147483648, F1: (fi37 & fi37)}
                        _ = v40
                    case 128:
                        switch fi35 {
                        case -446:
                            var v41 int = (fi37 << ((42) & 63))
                            _ = v41
                        case 127:
                            switch (fi35 & (fi35 << ((563) & 63))) {
                            case -36:
                                v42 := GenS1{F0: 1, F1: fi37}
                                _ = v42
                            }
                        }
                    }
                    switch (fi35 & 32768) {
                    case 513:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v43 uint32 = 2
                        var v44 bool = ((fi35 ^ fi35) != (^(fi35)))
                        fmt.Println(v4, fi35)
                        _ = v43
                        _ = v44
                    case 0:
                        switch ("foo bar" + "a") {
                        case "\n\t":
                            v9 -= 42
                        case "foo bar":
                            v45 := GenS5{F0: 42}
                            _ = v45
                        case "":
                            if v2 {
                                switch fi35 {
                                case 127:
                                    var v46 int8 = 42
                                    _ = v46
                                case 1:
                                    v47 := GenS3{F0: (v9 % ((v9) | 1)), F1: v5}
                                    _ = v47
                                case -1:
                                    _ = len(v1)
                                }
                                if (42 > 0) {
                                    var v48 uint16 = 65535
                                    var v49 uint16 = 1
                                    var v50 int16 = 0
                                    v2 = (2 == 127)
                                    _ = v48
                                    _ = v49
                                    _ = v50
                                } else {
                                    var v51 uint = 65535
                                    var v52 uint16 = 844
                                    fmt.Println(v51)
                                    _ = v51
                                    _ = v52
                                }
                                if v2 {
                                    fmt.Println(v5, fi37, v2)
                                }
                            } else {
                                _ = len(v3)
                            }
                        }
                    case 268:
                        _ = len(v1)
                    default:
                        switch 4294967295 {
                        case 754:
                            switch 42 {
                            case 9223372036854775807:
                                fmt.Println(v5, v9)
                            case -946:
                                v5 = 0
                            case 32768:
                                var v53 int32 = v10
                                _ = v53
                            default:
                                var v54 uint64 = 761
                                _ = v54
                            }
                        case 32767:
                            v55 := GenS3{F0: v9, F1: (v5 * v5)}
                            _ = v55
                        case 1:
                            v10 += 0
                        default:
                            v4 = (float64(3.14159) / v4)
                        }
                    }
                } else {
                    if v2 {
                        switch (fi35 * fi37) {
                        case 0:
                            v56 := GenS6{F0: v2}
                            _ = v56
                        case -1:
                            var v57 int64 = v5
                            _ = v57
                        case 4294967295:
                            v58 := GenS7{F0: 42, F1: -1}
                            var v59 int32 = 599
                            var v60 uint64 = 2
                            _ = v58
                            _ = v59
                            _ = v60
                        default:
                            switch (fi35 + 1) {
                            case 632:
                                v10 = (v10 & v10)
                            case -9223372036854775808:
                                var v61 int64 = v5
                                _ = v61
                            case 4294967295:
                                if (((float32((((v4 - v4) + 3.14159) + (3.14159 * (v4 + v4)))) + float32((float64(0.0) - 0.0))) <= (0.0 * 1e10)) && ((v9 * (v9 * (-(v9)))) > v9)) {
                                    var v62 int32 = v10
                                    switch ("" + "\n\t") {
                                    case "":
                                        v63 := GenS1{F0: (v10 - int32(v9)), F1: fi37}
                                        _ = v63
                                    case "hello":
                                        var v64 int32 = v10
                                        _ = v64
                                    case "\n\t":
                                        if v2 {
                                            v65 := GenS8{F0: 42, F1: 4294967295, F2: (v9 ^ v9)}
                                            _ = v65
                                        }
                                    default:
                                        var v66 string = ("foo bar" + "")
                                        _ = v66
                                    }
                                    v67 := GenS9{F0: ("a" + "")}
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v62
                                    _ = v67
                                }
                            default:
                                var v68 uint64 = 0
                                _ = v68
                            }
                        }
                        fmt.Println(v9, v9)
                        var v69 string = (("hello" + "hello") + (("a" + "hello") + ""))
                        _ = v69
                    } else {
                        v4 = v4
                        v3 = append(v3, 2)
                        switch "a" {
                        case "a":
                            if true {
                                fmt.Println(v5, v5)
                                switch "" {
                                case "":
                                    v10 -= -2147483648
                                }
                                fmt.Println(v9, v9)
                                v9 = v9
                                switch (("a" + "foo bar") + "a") {
                                case "a":
                                    fmt.Println(fi35, v4)
                                case "":
                                    fmt.Println(v5, fi37)
                                }
                                fmt.Println(fi37, fi37, v4)
                            }
                            v70 := GenS1{F0: v10, F1: fi35}
                            _ = v70
                        case "hello":
                            v5 += (v5 << ((255) & 63))
                        default:
                            v71 := GenS10{F0: v10, F1: -1, F2: fi35}
                            _ = v71
                        }
                    }
                    var v72 float32 = 1e10
                    v3 = append(v3, 756)
                    if v2 {
                        v9 = v9
                        v73 := GenS10{F0: 2147483647, F1: -32768, F2: (fi35 | (fi37 - fi35))}
                        v9 += v9
                        fmt.Println(v9)
                        _ = v73
                    }
                    _ = v72
                }
            }
            _ = fi37
        }
        _ = fi35
        _ = v36
    }
    var v74 float32 = ((1.5 - 1e10) / 3.14159)
    for fi75 := 0; fi75 < 2; fi75++ {
        var v76 int32 = v10
        fmt.Println(v74)
        _ = fi75
        _ = v76
    }
    var v77 bool = (255 > 42)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v74:", v74)
    fmt.Println("v77:", v77)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v74
    _ = v77
    _ = v9
}
