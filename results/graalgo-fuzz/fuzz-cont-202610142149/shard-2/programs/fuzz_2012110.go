package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 int }

func main() {
    v1 := []int{-2147483648}
    var v2 float32 = -2.25
    var v3 int = 40
    var v4 int = v3
    var v5 float32 = (v2 / v2)
    var v6 string = ("foo bar" + ("foo bar" + ("\n\t" + "foo bar")))
    var v7 uint8 = 246
    switch "a" {
    case "\n\t":
        fmt.Println(v5, v2, v4)
    default:
        for fi8 := 0; fi8 < 11; fi8++ {
            v6 = v6
            var v9 int = 261
            _ = fi8
            _ = v9
        }
    }
    if (0.0 != (v5 - v2)) {
        switch (v3 | (v3 ^ -341)) {
        case 127:
            v10 := GenS1{F0: int8(v4), F1: 0, F2: 4294967295}
            _ = v10
        case -547:
            switch 470 {
            case 804:
                for fi11 := 0; fi11 < 12; fi11++ {
                    v12 := GenS1{F0: 42, F1: 255, F2: 18446744073709551615}
                    if ((v7 == (v7 | v7)) && false) {
                        switch v6 {
                        case "":
                            v1 = append(v1, v4)
                        case "a":
                            if false {
                                var v13 int16 = -565
                                _ = v13
                            }
                        case "\n\t":
                            v4 -= (v4 ^ (fi11 + v4))
                        default:
                            if (3.14159 != 3.14159) {
                                var v14 uint8 = 223
                                if (!(18446744073709551615 >= 42)) {
                                    switch v6 {
                                    case "hello":
                                        for fi15 := 0; fi15 < 20; fi15++ {
                                            var v16 int = (^(v4))
                                            _ = fi15
                                            _ = v16
                                        }
                                    }
                                    for fi17 := 0; fi17 < 5; fi17++ {
                                        v2 = 0.0
                                        if true {
                                            if true {
                                                var v18 uint = 982
                                                v1 = append(v1, v3)
                                                if len(v1) > 0 {
                                                    v1[0] = v4
                                                }
                                                _ = len(v1)
                                                var v19 int32 = -2147483648
                                                var v20 float32 = 1.5
                                                var v21 uint32 = 1
                                                _ = v18
                                                _ = v19
                                                _ = v20
                                                _ = v21
                                            } else {
                                                var v22 int8 = 120
                                                var v23 bool = (128 <= (v14 | 229))
                                                v24 := GenS1{F0: v22, F1: 0, F2: 0}
                                                _ = v22
                                                _ = v23
                                                _ = v24
                                            }
                                            switch 32767 {
                                            case 1:
                                                v25 := GenS2{F0: 42}
                                                _ = v25
                                            }
                                            fmt.Println(fi11, fi17, v5)
                                        } else {
                                            if (-1 < -140) {
                                                if (false && false) {
                                                    if (862 > 0) {
                                                        v26 := GenS3{F0: ("a" + v6), F1: 32767}
                                                        _ = v26
                                                    } else {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        var v27 int16 = 1000
                                                        _ = v27
                                                    }
                                                    v14 -= 172
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v28 uint64 = 42
                                                    fmt.Println(v14, v14)
                                                    var v29 float64 = 1e10
                                                    if false {
                                                        var v30 uint = 0
                                                        _ = v30
                                                    } else {
                                                        switch (v3 << ((2) & 63)) {
                                                        case 127:
                                                            v3 = v3
                                                        case -633:
                                                            var v31 float32 = float32(v29)
                                                            _ = v31
                                                        case 128:
                                                            fmt.Println(v2, v6)
                                                        }
                                                        var v32 int = (v4 & fi17)
                                                        var v33 int32 = -1
                                                        if len(v1) > 0 {
                                                            v1[0] = v3
                                                        }
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                    switch 9223372036854775807 {
                                                    case -825:
                                                        v5 = (1e10 * v2)
                                                    }
                                                    v5 -= 1.5
                                                    _ = v28
                                                    _ = v29
                                                } else {
                                                    v5 += v5
                                                    v3 = fi11
                                                }
                                            } else {
                                                if (1 < 0) {
                                                    if false {
                                                        v2 = v5
                                                        var v34 uint16 = 65535
                                                        _ = v34
                                                    } else {
                                                        v35 := GenS1{F0: 42, F1: 914, F2: 255}
                                                        v6 = "a"
                                                        _ = v35
                                                    }
                                                    v36 := GenS4{F0: 663}
                                                    v37 := GenS5{F0: 2147483647, F1: v3}
                                                    _ = v36
                                                    _ = v37
                                                } else {
                                                    v3 = -1
                                                }
                                            }
                                        }
                                        _ = fi17
                                    }
                                    v38 := GenS3{F0: v6, F1: 1}
                                    _ = v38
                                } else {
                                    _ = len(v1)
                                    if (1 == 32767) {
                                        if ((0.0 * float64(3.14159)) > ((0.1 / 1.5) * (0.0 * (1.5 * 1.5)))) {
                                            _ = len(v1)
                                        }
                                        if (fi11 <= v3) {
                                            var v39 uint16 = 2
                                            _ = v39
                                        } else {
                                            fmt.Println(v6, v2, v2)
                                        }
                                    }
                                }
                                if true {
                                    for fi40 := 0; fi40 < 6; fi40++ {
                                        _ = len(v1)
                                        if len(v1) > 0 {
                                            v1[0] = fi11
                                        }
                                        _ = fi40
                                    }
                                    for fi41 := 0; fi41 < 17; fi41++ {
                                        v4 += fi41
                                        _ = fi41
                                    }
                                    v2 = v5
                                } else {
                                    v42 := GenS2{F0: 59}
                                    _ = v42
                                }
                                _ = v14
                            }
                        }
                        v43 := GenS1{F0: 100, F1: 65535, F2: 2}
                        _ = v43
                    }
                    v5 += 3.14159
                    _ = fi11
                    _ = v12
                }
            case 32767:
                v7 += v7
            }
        default:
            switch v6 {
            case "hello":
                v44 := GenS1{F0: 1, F1: 1, F2: 0}
                _ = v44
            case "a":
                var v45 string = "foo bar"
                _ = v45
            case "":
                for fi46 := 0; fi46 < 9; fi46++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if true {
                        v2 += (v2 - ((v2 + 1.5) + v2))
                    }
                    _ = fi46
                }
            }
        }
        v4 -= (v4 | 127)
        switch "a" {
        case "foo bar":
            for fi47 := 0; fi47 < 13; fi47++ {
                var v48 int = -166
                v2 += v2
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi47
                _ = v48
            }
        case "":
            _ = len(v1)
        case "\n\t":
            _ = len(v1)
        default:
            for fi49 := 0; fi49 < 2; fi49++ {
                fmt.Println(v5)
                _ = fi49
            }
        }
    } else {
        var v50 uint8 = 255
        if len(v1) > 0 {
            v1[0] = (v3 * 9223372036854775807)
        }
        _ = v50
    }
    if true {
        if (!true) {
            if (0.1 != (1.5 / v5)) {
                if (("\n\t" + (v6 + (v6 + v6))) < ((v6 + v6) + (("a" + v6) + "hello"))) {
                    if len(v1) > 0 {
                        v1[0] = (v3 % (((v4 + (^(v3)))) | 1))
                    }
                    if len(v1) > 0 {
                        v1[0] = (v4 >> ((454) & 63))
                    }
                    v6 = v6
                } else {
                    v51 := GenS2{F0: 42}
                    var v52 int64 = 724
                    var v53 int64 = (v52 >> ((4294967295) & 63))
                    _ = v51
                    _ = v52
                    _ = v53
                }
                for fi54 := 0; fi54 < 1; fi54++ {
                    v2 = (0.1 + float32(3.14159))
                    _ = fi54
                }
                switch (^(v3)) {
                case 128:
                    var v55 int = -2147483648
                    _ = v55
                default:
                    var v56 uint = 42
                    _ = v56
                }
            }
            var v57 bool = true
            var v58 uint32 = 1
            _ = v57
            _ = v58
        }
        v7 -= v7
    } else {
        for fi59 := 0; fi59 < 11; fi59++ {
            v5 += 0.0
            fmt.Println(fi59, v2)
            _ = fi59
        }
        var v60 uint8 = (v7 - v7)
        _ = v60
    }
    var v61 uint = 2
    switch -961 {
    case 4294967295:
        var v62 int = int(v7)
        _ = v62
    case 127:
        var v63 uint32 = 4294967295
        _ = v63
    case 9223372036854775807:
        v3 -= v4
    }
    switch v3 {
    case -1:
        for fi64 := 0; fi64 < 2; fi64++ {
            var v65 string = (v6 + "hello")
            _ = len(v1)
            var v66 int8 = 0
            _ = fi64
            _ = v65
            _ = v66
        }
    case 128:
        v7 = v7
    case 2147483647:
        v61 -= 154
    default:
        var v67 int64 = -9223372036854775808
        _ = v67
    }
    var v68 uint16 = 757
    fmt.Println(v68, v5)
    if true {
        v69 := GenS2{F0: uint16(v7)}
        if len(v1) > 0 {
            v1[0] = (v4 / (((v4 + v3)) | 1))
        }
        _ = v69
    } else {
        var v70 float32 = v5
        _ = v70
    }
    fmt.Println(v6, v7)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v68:", v68)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v68
    _ = v7
}
