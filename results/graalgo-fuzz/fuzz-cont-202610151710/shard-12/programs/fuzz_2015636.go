package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint16; F2 int32 }
type GenS4 struct { F0 bool; F1 bool; F2 string }
type GenS5 struct { F0 uint8; F1 bool; F2 int8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64; F1 int8; F2 int8 }

func main() {
    v1 := []string{}
    var v2 int = -9223372036854775808
    var v3 float32 = -2.25
    for fi4 := 0; fi4 < 6; fi4++ {
        v5 := GenS1{F0: 175}
        _ = fi4
        _ = v5
    }
    var v6 string = "a"
    v6 = v6
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 string = v6
            if (false || false) {
                if (0 < 1) {
                    switch v9 {
                    case "hello":
                        var v10 uint16 = 1
                        _ = v10
                    case "foo bar":
                        v11 := GenS1{F0: 64}
                        _ = v11
                    default:
                        if false {
                            fmt.Println(v6, fi7)
                            v2 += (fi7 % (((v2 ^ -1)) | 1))
                        } else {
                            v6 = v6
                            var v12 bool = true
                            _ = v12
                        }
                    }
                    switch 2147483647 {
                    case -9223372036854775808:
                        if false {
                            v9 = (v9 + (((v9 + "a") + v6) + v6))
                        }
                    case 127:
                        if len(v1) > 0 {
                            v1[0] = v6
                        }
                    }
                    fmt.Println(v6, fi7)
                } else {
                    v13 := GenS2{F0: 78}
                    v14 := GenS2{F0: 42}
                    _ = v13
                    _ = v14
                }
            }
            _ = fi8
            _ = v9
        }
        switch "\n\t" {
        case "hello":
            v15 := GenS2{F0: -1}
            _ = v15
        }
        _ = fi7
    }
    if false {
        var v16 uint64 = 0
        v17 := GenS2{F0: -128}
        var v18 int16 = 1
        _ = v16
        _ = v17
        _ = v18
    } else {
        switch (v2 ^ (v2 - v2)) {
        case -303:
            var v19 int32 = 0
            _ = v19
        case 9223372036854775807:
            if (155 == 2147483648) {
                fmt.Println(v2)
                var v20 int32 = 656
                switch v6 {
                case "foo bar":
                    v2 = (v2 | v2)
                default:
                    v21 := GenS3{F0: -1, F1: 759, F2: 42}
                    _ = v21
                }
                _ = v20
            }
        case 1:
            var v22 int32 = 42
            _ = v22
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 uint64 = 932
            _ = len(v1)
            _ = fi23
            _ = v24
        }
        v6 = (v6 + v6)
    }
    var v25 string = v6
    for fi26 := 0; fi26 < 8; fi26++ {
        if false {
            v27 := GenS3{F0: 42, F1: 574, F2: 2147483647}
            v6 = ((v25 + v6) + (v25 + "\n\t"))
            _ = v27
        } else {
            var v28 string = v25
            _ = v28
        }
        v29 := GenS1{F0: 33}
        _ = fi26
        _ = v29
    }
    var v30 uint16 = 2
    v31 := GenS4{F0: ((4294967295 > 0) && (!(!false))), F1: false, F2: v6}
    if (v30 == v30) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi32 := 0; fi32 < 4; fi32++ {
            var v33 uint16 = v30
            _ = fi32
            _ = v33
        }
    } else {
        v34 := GenS5{F0: 0, F1: true, F2: -34}
        _ = v34
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        v25 = (v6 + "foo bar")
        v1 = append(v1, "")
        _ = fi35
    }
    if false {
        switch v25 {
        case "":
            v30 = v30
        }
        for fi36 := 0; fi36 < 8; fi36++ {
            var v37 uint16 = v30
            var v38 int = v2
            for fi39 := 0; fi39 < 10; fi39++ {
                fmt.Println(fi36, v37, v6)
                _ = fi39
            }
            _ = fi36
            _ = v37
            _ = v38
        }
    } else {
        if true {
            v1 = append(v1, (("foo bar" + (v6 + (("a" + ("" + v6)) + (((("foo bar" + (("a" + v6) + v25)) + (v25 + "\n\t")) + v25) + "foo bar")))) + (v6 + "")))
        } else {
            var v40 float32 = (v3 * (float32((0.0 + float64(v3))) / ((float32(0.0) + (((v3 - 0.0) - v3) / 3.14159)) + 1.5)))
            for fi41 := 0; fi41 < 7; fi41++ {
                for fi42 := 0; fi42 < 7; fi42++ {
                    if false {
                        var v43 float64 = 1e10
                        v43 = v43
                        _ = v43
                    } else {
                        if false {
                            var v44 uint = 134
                            _ = v44
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    if (false && (((0.0 / float64(v40)) + 3.14159) != 3.14159)) {
                        var v45 int32 = -630
                        if true {
                            var v46 int8 = -37
                            if (v2 == v2) {
                                var v47 int = (^(fi41))
                                _ = v47
                            }
                            _ = v46
                        }
                        _ = v45
                    } else {
                        _ = len(v1)
                        v2 += (fi41 + (fi42 / (((fi42 + (fi42 - 4294967295))) | 1)))
                        _ = len(v1)
                    }
                    _ = fi42
                }
                for fi48 := 0; fi48 < 16; fi48++ {
                    if (!(!true)) {
                        if (!true) {
                            if len(v1) > 0 {
                                v1[0] = ((v25 + ("" + "a")) + "foo bar")
                            }
                            switch (fi41 >> ((0) & 63)) {
                            case 254:
                                v6 = v6
                            case 42:
                                fmt.Println(v25, v25, v30)
                            case -9223372036854775808:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        } else {
                            fmt.Println(fi41, v2, v2)
                        }
                        v25 = ((("a" + v6) + "") + "foo bar")
                        if (1 > 65535) {
                            v49 := GenS2{F0: -88}
                            _ = v49
                        }
                    }
                    if (!(false || (-1 > -181))) {
                        v50 := GenS6{F0: 311}
                        if (((v40 + v3) >= (v40 / ((-2.25 * (v40 / v3)) * (v40 * 1e10)))) && (127 < 0)) {
                            v6 = v25
                        } else {
                            if len(v1) > 0 {
                                v1[0] = v25
                            }
                            switch 128 {
                            case -1:
                                var v51 int8 = -1
                                _ = v51
                            default:
                                var v52 int = int(v30)
                                _ = v52
                            }
                            var v53 int8 = 42
                            v54 := GenS7{F0: 42, F1: v53, F2: v53}
                            _ = v53
                            _ = v54
                        }
                        _ = v50
                    }
                    _ = fi48
                }
                _ = fi41
            }
            for fi55 := 0; fi55 < 7; fi55++ {
                switch v25 {
                case "foo bar":
                    for fi56 := 0; fi56 < 8; fi56++ {
                        var v57 int16 = 42
                        _ = fi56
                        _ = v57
                    }
                default:
                    fmt.Println(v2)
                }
                _ = fi55
            }
            _ = v40
        }
    }
    for fi58 := 0; fi58 < 19; fi58++ {
        v59 := GenS1{F0: 186}
        switch (v2 - (fi58 << ((18446744073709551615) & 63))) {
        case 42:
            switch "foo bar" {
            case "foo bar":
                v6 = v6
            }
        case -430:
            if (!false) {
                v30 = v30
                fmt.Println(v2, fi58)
                var v60 float32 = v3
                _ = v60
            }
        default:
            if true {
                var v61 uint64 = 1
                _ = v61
            }
        }
        _ = fi58
        _ = v59
    }
    fmt.Println(v6, v30)
    v30 += v30
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v6
}
