package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint; F2 string }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 int; F2 uint }
type GenS6 struct { F0 int8; F1 uint8; F2 int }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 uint; F1 uint32; F2 int16 }
type GenS9 struct { F0 uint8 }

func main() {
    var v1 bool = (false || (-198 != 1))
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    var v4 uint16 = 0
    v5 := GenS1{F0: 2}
    var v6 int16 = -603
    v7 := []int{128, -9223372036854775808, 42}
    if (1e10 < 3.14159) {
        for fi8 := 0; fi8 < 8; fi8++ {
            v6 += 32767
            switch fi8 {
            case -864:
                var v9 int16 = (v6 % ((-1) | 1))
                _ = v9
            case 32767:
                var v10 int64 = -1
                _ = v10
            case -813:
                var v11 int8 = 75
                _ = v11
            }
            _ = fi8
        }
        v12 := GenS2{F0: 65535, F1: 243, F2: "a"}
        var v13 int = -349
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        switch "" {
        case "foo bar":
            v7 = append(v7, fi14)
        }
        switch (^(fi14)) {
        case 128:
            v4 = (^(v4))
        case -168:
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 int32 = 192
                _ = fi15
                _ = v16
            }
        }
        _ = fi14
    }
    if (-528 > -2147483648) {
        fmt.Println(v4)
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            switch (fi17 % ((fi17) | 1)) {
            case 237:
                for fi18 := 0; fi18 < 11; fi18++ {
                    delete(v3, "\n\t")
                    _ = len(v2)
                    _ = fi18
                }
            default:
                v4 = (v4 + v4)
            }
            _ = fi17
        }
        var v19 string = ("" + "a")
        _ = v19
    }
    var v20 string = "\n\t"
    fmt.Println(v1)
    v21 := GenS1{F0: 799}
    v22 := GenS2{F0: (^(v4)), F1: 74, F2: v20}
    if v1 {
        var v23 uint16 = (v4 | 42)
        _ = len(v7)
        v6 = 1000
        _ = v23
    } else {
        v20 = ("foo bar" + (v20 + v20))
        switch -903 {
        case 4294967295:
            v6 += v6
        }
        var v24 int64 = 0
        _ = v24
    }
    _ = len(v7)
    if len(v7) > 0 {
        v7[0] = int(v6)
    }
    v4 = v4
    if false {
        if true {
            var v25 bool = (true && true)
            for fi26 := 0; fi26 < 4; fi26++ {
                if v1 {
                    if v1 {
                        switch 128 {
                        case 127:
                            switch "" {
                            case "hello":
                                if (!v1) {
                                    if ((42 <= -504) || v1) {
                                        switch v20 {
                                        case "a":
                                            var v27 bool = true
                                            _ = v27
                                        case "foo bar":
                                            v25 = true
                                        }
                                        for fi28 := 0; fi28 < 7; fi28++ {
                                            switch 4294967295 {
                                            case 1:
                                                if false {
                                                    switch (v20 + v20) {
                                                    case "a":
                                                        v29 := GenS3{F0: 42, F1: 147, F2: 2}
                                                        _ = v29
                                                    case "foo bar":
                                                        if len(v7) > 0 {
                                                            v7[0] = int(v6)
                                                        }
                                                    case "hello":
                                                        v7 = append(v7, (-(fi26)))
                                                    default:
                                                        switch "a" {
                                                        case "":
                                                            v4 = 948
                                                        case "\n\t":
                                                            v7 = append(v7, (fi26 / (((fi26 % (((fi26 & fi28)) | 1))) | 1)))
                                                        case "a":
                                                            switch (("\n\t" + v20) + ("foo bar" + "\n\t")) {
                                                            case "hello":
                                                                fmt.Println(v1)
                                                            case "\n\t":
                                                                v30 := v3[v20]
                                                                _ = v30
                                                            case "a":
                                                                if v1 {
                                                                    v7 = append(v7, (-(fi26)))
                                                                    if false {
                                                                        var v31 uint16 = (v4 / ((v4) | 1))
                                                                        if v25 {
                                                                            var v32 uint16 = v31
                                                                            var v33 float64 = -2.25
                                                                            var v34 int64 = 9223372036854775807
                                                                            _ = v32
                                                                            _ = v33
                                                                            _ = v34
                                                                        }
                                                                        _ = v31
                                                                    }
                                                                } else {
                                                                    if v1 {
                                                                        v20 = (v20 + v20)
                                                                        switch fi28 {
                                                                        case 42:
                                                                            var v35 uint16 = 65535
                                                                            _ = v35
                                                                        case 128:
                                                                            v36 := GenS2{F0: (v4 >> ((42) & 63)), F1: 4294967295, F2: (v20 + "hello")}
                                                                            _ = v36
                                                                        }
                                                                        var v37 int = (fi28 % ((-266) | 1))
                                                                        _ = v37
                                                                    }
                                                                    v1 = (-544 <= 1)
                                                                }
                                                            default:
                                                                var v38 bool = (2 >= 42)
                                                                _ = v38
                                                            }
                                                        default:
                                                            if ((v4 | (v4 << ((292) & 63))) == v4) {
                                                                v39 := GenS3{F0: 4294967295, F1: 128, F2: 1}
                                                                v6 = 277
                                                                _ = v39
                                                            }
                                                        }
                                                    }
                                                    delete(v2, 2147483648)
                                                } else {
                                                    if v1 {
                                                        var v40 int32 = 0
                                                        _ = v40
                                                    }
                                                    v41 := GenS1{F0: 42}
                                                    if (!v1) {
                                                        delete(v3, v20)
                                                        delete(v3, v20)
                                                        fmt.Println(fi28, v4, v20)
                                                    }
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v6 += 42
                                                    case "":
                                                        var v42 int32 = -21
                                                        _ = v42
                                                    default:
                                                        var v43 string = "foo bar"
                                                        _ = v43
                                                    }
                                                    switch v20 {
                                                    case "hello":
                                                        if v25 {
                                                            if true {
                                                                if v1 {
                                                                    var v44 int64 = 9223372036854775807
                                                                    _ = v44
                                                                } else {
                                                                    v6 = (v6 - (^(v6)))
                                                                    v45 := GenS4{F0: 9223372036854775807}
                                                                    _ = v45
                                                                }
                                                                var v46 uint = 770
                                                                v47 := GenS5{F0: 2147483647, F1: (^(fi28)), F2: v46}
                                                                var v48 bool = v1
                                                                var v49 uint16 = (v4 >> ((v46) & 63))
                                                                var v50 uint8 = 40
                                                                var v51 float64 = 3.14159
                                                                _ = v46
                                                                _ = v47
                                                                _ = v48
                                                                _ = v49
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                        } else {
                                                            var v52 int16 = v6
                                                            v25 = (2 == 2)
                                                            v4 = 2
                                                            _ = v52
                                                        }
                                                    case "a":
                                                        var v53 uint16 = v4
                                                        switch v20 {
                                                        case "":
                                                            var v54 uint16 = 262
                                                            _ = v54
                                                        case "a":
                                                            v55 := v2[0]
                                                            _ = v55
                                                        }
                                                        _ = v53
                                                    case "foo bar":
                                                        var v56 int64 = 9223372036854775807
                                                        _ = v56
                                                    }
                                                    _ = v41
                                                }
                                            case 4294967295:
                                                v1 = true
                                            case 633:
                                                fmt.Println(v6, v20)
                                            }
                                            var v57 int32 = 1
                                            _ = fi28
                                            _ = v57
                                        }
                                    }
                                } else {
                                    var v58 uint8 = 128
                                    if (v58 > v58) {
                                        _ = len(v7)
                                        v4 -= v4
                                        var v59 bool = ((118 > 1) && v25)
                                        _ = v59
                                    } else {
                                        for fi60 := 0; fi60 < 10; fi60++ {
                                            var v61 uint32 = 715
                                            _ = fi60
                                            _ = v61
                                        }
                                        for fi62 := 0; fi62 < 18; fi62++ {
                                            v4 += v4
                                            fmt.Println(v25, v6, v6)
                                            _ = fi62
                                        }
                                        var v63 float32 = (3.14159 - -2.25)
                                        _ = v63
                                    }
                                    fmt.Println(v20, v20, v1)
                                    _ = v58
                                }
                            case "\n\t":
                                var v64 uint = 0
                                _ = v64
                            }
                        case -2147483648:
                            v25 = ((float32(-2.25) + 1e10) == 1.5)
                        case 2147483647:
                            switch 42 {
                            case 32768:
                                var v65 bool = v25
                                _ = v65
                            case 32767:
                                if v25 {
                                    fmt.Println(v4, v25)
                                    var v66 uint = 255
                                    if (!(fi26 < (fi26 - (fi26 - fi26)))) {
                                        var v67 float32 = 1.5
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                        _ = v67
                                    } else {
                                        var v68 int16 = 1000
                                        _ = v68
                                    }
                                    _ = v66
                                } else {
                                    v69 := GenS6{F0: -33, F1: 42, F2: (fi26 << ((18446744073709551615) & 63))}
                                    _ = v69
                                }
                            }
                        }
                        for fi70 := 0; fi70 < 8; fi70++ {
                            switch v20 {
                            case "":
                                v4 = v4
                            case "\n\t":
                                v4 = v4
                            case "hello":
                                var v71 float32 = (1.5 + 1.5)
                                _ = v71
                            default:
                                if (((3.14159 - 3.14159) > 1e10) || ((fi26 / ((9223372036854775807) | 1)) < fi70)) {
                                    var v72 int16 = int16(v4)
                                    _ = v72
                                }
                            }
                            _ = fi70
                        }
                    }
                }
                v73 := GenS7{F0: 127, F1: 590}
                var v74 int64 = 1
                _ = fi26
                _ = v73
                _ = v74
            }
            _ = v25
        }
        v4 -= v4
        var v75 uint64 = 260
        _ = v75
    } else {
        v1 = v1
        if (1 >= 65535) {
            switch 434 {
            case 4294967295:
                v7 = append(v7, -9223372036854775808)
            case 386:
                if (!v1) {
                    for fi76 := 0; fi76 < 8; fi76++ {
                        if (!v1) {
                            switch -9223372036854775808 {
                            case 127:
                                var v77 uint16 = 465
                                _ = v77
                            case -901:
                                if false {
                                    var v78 uint8 = 48
                                    v79 := GenS1{F0: 1}
                                    _ = v78
                                    _ = v79
                                } else {
                                    v1 = true
                                    var v80 int = (fi76 * int(v4))
                                    _ = v80
                                }
                            default:
                                v81 := GenS8{F0: 2, F1: 2147483648, F2: 786}
                                _ = v81
                            }
                        } else {
                            if ((v6 * 42) != 1) {
                                var v82 bool = (!(((!v1) || true) && true))
                                delete(v3, v20)
                                _ = v82
                            }
                        }
                        _ = fi76
                    }
                }
            default:
                v83 := GenS9{F0: 0}
                _ = v83
            }
        }
        var v84 uint16 = v4
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
