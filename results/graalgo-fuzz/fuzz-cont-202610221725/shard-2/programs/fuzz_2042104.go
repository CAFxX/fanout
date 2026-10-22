package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint8; F1 int; F2 uint64 }
func genFn0() uint8 {
    v4 := []int{32768, 128, -258}
    var v5 float32 = ((1.5 / 1.0) * 3.14159)
    _ = v4
    _ = v5
    return 19
}
func genFn1(p0 int8) int32 {
    p0 = (p0 & p0)
    v6 := []int{413, -575}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 int16 = -334
    _ = p0
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := []int{-47, 32767, 32767, 260}
    v2 := make(map[uint32]string)
    var v3 int16 = -32768
    var v8 int32 = -2147483648
    var v9 int = 127
    delete(v2, 2147483647)
    if false {
        v9 = v9
        v10 := GenS1{F0: 42, F1: -720}
        if false {
            var v11 uint32 = 184
            v9 = v9
            _ = v11
        } else {
            v3 += 42
            v8 += 238
        }
        _ = v10
    }
    var v12 uint16 = 171
    switch ("hello" + "hello") {
    case "\n\t":
        if true {
            _ = len(v1)
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := GenS1{F0: 138, F1: 42}
                v15 := GenS1{F0: 2, F1: -1}
                _ = fi13
                _ = v14
                _ = v15
            }
            var v16 int8 = 101
            _ = v16
        } else {
            var v17 uint16 = v12
            if false {
                for fi18 := 0; fi18 < 18; fi18++ {
                    var v19 uint = 691
                    _ = len(v1)
                    _ = fi18
                    _ = v19
                }
            } else {
                v17 = 1
            }
            _ = v17
        }
    case "a":
        switch (v9 - 32768) {
        case -160:
            var v20 int32 = (v8 | v8)
            _ = v20
        case -654:
            switch "a" {
            case "":
                var v21 uint16 = v12
                _ = v21
            default:
                if true {
                    v22 := GenS1{F0: 47, F1: 42}
                    if (1e10 < -2.25) {
                        if true {
                            v8 = v8
                            if len(v1) > 0 {
                                v1[0] = v9
                            }
                        } else {
                            v23 := GenS1{F0: 255, F1: 9223372036854775807}
                            _ = v23
                        }
                        switch v9 {
                        case 2147483647:
                            v2[2147483648] = ("foo bar" + ("foo bar" + "a"))
                        case -886:
                            var v24 uint32 = 0
                            _ = v24
                        case 42:
                            var v25 float64 = 1.5
                            _ = v25
                        }
                    }
                    _ = v22
                }
            }
        case 1:
            v3 = 1
        default:
            if (!false) {
                var v26 int = (v9 / (((v9 + 32768)) | 1))
                var v27 uint16 = 2
                v28 := GenS1{F0: 1, F1: 1}
                _ = v26
                _ = v27
                _ = v28
            } else {
                v9 += 130
                var v29 int16 = -406
                _ = v29
            }
        }
    case "hello":
        var v30 uint = 291
        _ = v30
    }
    if true {
        var v31 string = "hello"
        if (4294967295 == 4294967295) {
            fmt.Println(v3, v12)
        }
        if true {
            switch v9 {
            case 962:
                v9 -= 0
            case -772:
                if (int8(v8) <= -128) {
                    if true {
                        var v32 string = v31
                        _ = v32
                    } else {
                        if true {
                            var v33 uint16 = (v12 + (v12 / (((v12 >> ((1) & 63))) | 1)))
                            switch (^(v9)) {
                            case -906:
                                v34 := GenS1{F0: 0, F1: 1}
                                _ = v34
                            case 32767:
                                v35 := GenS1{F0: 42, F1: 42}
                                _ = v35
                            case 42:
                                v1 = append(v1, v9)
                            }
                            _ = v33
                        }
                        if false {
                            fmt.Println(v9)
                            v12 = 777
                        } else {
                            fmt.Println(v3, v12, v3)
                            var v36 int8 = 1
                            for fi37 := 0; fi37 < 16; fi37++ {
                                v1 = append(v1, v9)
                                for fi38 := 0; fi38 < 2; fi38++ {
                                    v39 := GenS1{F0: 140, F1: -9223372036854775808}
                                    _ = fi38
                                    _ = v39
                                }
                                switch v31 {
                                case "foo bar":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case "":
                                    v8 -= (v8 << ((662) & 63))
                                default:
                                    fmt.Println(v12, v12, v8)
                                }
                                _ = fi37
                            }
                            _ = v36
                        }
                    }
                    var v40 int16 = v3
                    _ = v40
                }
            }
            for fi41 := 0; fi41 < 5; fi41++ {
                switch (v9 - fi41) {
                case 122:
                    fmt.Println(v9, v31, v9)
                case 245:
                    switch (fi41 % (((v9 ^ 885)) | 1)) {
                    case 42:
                        v12 += 46
                    case -268:
                        v1 = append(v1, v9)
                    }
                }
                switch (v31 + v31) {
                case "hello":
                    fmt.Println(v3, v12)
                }
                _ = fi41
            }
        } else {
            v12 = (v12 % (((v12 << ((4294967295) & 63))) | 1))
            v8 -= v8
            for fi42 := 0; fi42 < 1; fi42++ {
                var v43 uint = 2
                switch "hello" {
                case "":
                    v44 := GenS2{F0: (v31 + v31), F1: v31}
                    _ = v44
                }
                for fi45 := 0; fi45 < 14; fi45++ {
                    v31 = "a"
                    if true {
                        var v46 uint8 = 2
                        var v47 uint64 = 427
                        switch "" {
                        case "a":
                            fmt.Println(v46, v12)
                        case "":
                            switch v31 {
                            case "a":
                                v48 := GenS2{F0: (v31 + ("foo bar" + "a")), F1: "hello"}
                                var v49 int32 = v8
                                v8 = genFn1(42)
                                _ = v48
                                _ = v49
                            default:
                                _ = len(v2)
                            }
                        case "foo bar":
                            v50 := GenS3{F0: (v46 / (((^(v46))) | 1)), F1: -1, F2: v47}
                            _ = v50
                        default:
                            if true {
                                switch (v31 + "") {
                                case "\n\t":
                                    var v51 uint32 = 4294967295
                                    _ = v51
                                case "hello":
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v52 uint8 = (v46 * genFn0())
                                        _ = v52
                                    case "hello":
                                        if (2147483647 > 42) {
                                            var v53 int64 = int64(v12)
                                            if false {
                                                v53 += (v53 / ((int64(v46)) | 1))
                                                var v54 uint = (v43 | 859)
                                                _ = v54
                                            } else {
                                                if (false || (((0.0 == 0.0) && false) || (float32((((0.1 / -2.25) - 1.5) + -2.25)) < 0.0))) {
                                                    var v55 float32 = (float32((1.5 + 1.5)) * (1.5 + 0.1))
                                                    _ = v55
                                                }
                                            }
                                            _ = v53
                                        }
                                    }
                                case "":
                                    if (32767 > 1) {
                                        var v56 float64 = ((3.14159 + 1e10) / -2.25)
                                        _ = v56
                                    } else {
                                        var v57 int8 = -1
                                        if false {
                                            switch (v9 << ((uint(v46)) & 63)) {
                                            case -39:
                                                v58 := GenS1{F0: 2, F1: -355}
                                                _ = v58
                                            case 9223372036854775807:
                                                var v59 uint64 = 903
                                                _ = v59
                                            case 1:
                                                var v60 uint = 255
                                                _ = v60
                                            }
                                            switch 32768 {
                                            case 32768:
                                                delete(v2, 0)
                                            }
                                            var v61 int64 = int64(v43)
                                            _ = v61
                                        }
                                        fmt.Println(v57)
                                        _ = v57
                                    }
                                }
                                v2[65535] = (("hello" + (v31 + ((v31 + v31) + "hello"))) + ((("" + ("a" + "foo bar")) + ((v31 + ("foo bar" + v31)) + "\n\t")) + (v31 + v31)))
                            } else {
                                v43 -= (v43 / ((v43) | 1))
                            }
                        }
                        _ = v46
                        _ = v47
                    }
                    _ = fi45
                }
                _ = fi42
                _ = v43
            }
        }
        _ = v31
    }
    v62 := v2[1]
    v63 := GenS2{F0: "foo bar", F1: v62}
    for fi64 := 0; fi64 < 12; fi64++ {
        var v65 uint8 = 224
        if (v8 < (v8 % ((-174) | 1))) {
            var v66 float64 = ((((float64(0.1) - 1.5) + 3.14159) / (-2.25 + 1e10)) * 1e10)
            fmt.Println(v65, v12, v65)
            _ = v66
        }
        if (true && false) {
            var v67 uint = 475
            _ = v67
        }
        _ = fi64
        _ = v65
    }
    var v68 uint16 = (v12 & (^(v12)))
    v9 -= v9
    if ((^(v12)) > (^(v12))) {
        v3 -= (^(v3))
        var v69 int64 = 722
        _ = v69
    } else {
        v68 = (v68 ^ 255)
        for fi70 := 0; fi70 < 10; fi70++ {
            _ = len(v1)
            v8 -= 279
            _ = fi70
        }
        fmt.Println(v3, v3, v9)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v62:", v62)
    fmt.Println("v68:", v68)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v62
    _ = v63
    _ = v68
    _ = v8
    _ = v9
}
