package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int; F2 uint8 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint64; F1 uint8 }

func main() {
    v1 := []int{}
    var v2 bool = (true || false)
    var v3 float32 = 1.5
    var v4 float64 = 1.5
    v5 := make(map[string]uint32)
    var v6 int32 = 1
    fmt.Println(v3, v3)
    var v7 int16 = 32767
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            fmt.Println(v4, v3, fi9)
            switch fi9 {
            case -2147483648:
                v2 = true
            case -398:
                if true {
                    var v10 int64 = 1
                    switch ("\n\t" + "foo bar") {
                    case "\n\t":
                        v2 = false
                    case "foo bar":
                        fmt.Println(v6, v3, v10)
                    case "":
                        var v11 int64 = v10
                        _ = v11
                    default:
                        switch ("foo bar" + "") {
                        case "foo bar":
                            v5["a"] = 2147483648
                        case "\n\t":
                            switch "hello" {
                            case "foo bar":
                                fmt.Println(v2, v2, v6)
                            case "":
                                switch (^(fi9)) {
                                case 32767:
                                    var v12 int = 127
                                    _ = v12
                                case 172:
                                    v13 := GenS1{F0: 65535}
                                    _ = v13
                                case -726:
                                    fmt.Println(fi8)
                                }
                            case "a":
                                var v14 string = ("foo bar" + (("" + "hello") + "foo bar"))
                                _ = v14
                            }
                        }
                        var v15 int8 = 123
                        var v16 uint8 = 42
                        _ = v15
                        _ = v16
                    }
                    _ = v10
                } else {
                    fmt.Println(v6)
                    v17 := GenS1{F0: 0}
                    if ((v7 / (((v7 << ((4294967295) & 63))) | 1)) >= 1000) {
                        v2 = v2
                        if v2 {
                            var v18 uint16 = 255
                            v3 -= v3
                            if (!(674 < 0)) {
                                var v19 uint64 = 0
                                v20 := GenS1{F0: 303}
                                _ = v19
                                _ = v20
                            }
                            _ = v18
                        }
                        v7 -= v7
                    }
                    _ = v17
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v21 := GenS2{F0: 0, F1: 0}
            var v22 int = (fi8 & fi8)
            _ = fi9
            _ = v21
            _ = v22
        }
        _ = fi8
    }
    if false {
        v23 := GenS1{F0: 407}
        var v24 uint32 = 2147483647
        _ = v23
        _ = v24
    } else {
        for fi25 := 0; fi25 < 12; fi25++ {
            v7 -= (v7 * (v7 << ((0) & 63)))
            delete(v5, "")
            _ = fi25
        }
        var v26 int16 = (^(v7))
        var v27 uint64 = 0
        _ = v26
        _ = v27
    }
    var v28 uint16 = 65535
    switch "" {
    case "foo bar":
        for fi29 := 0; fi29 < 18; fi29++ {
            if (!v2) {
                _ = len(v1)
                if (-1 == -1) {
                    if (!(true || ((1 < v28) && (v7 <= v7)))) {
                        var v30 uint32 = 0
                        _ = v30
                    } else {
                        if ((v7 & (v7 % ((v7) | 1))) <= (v7 >> ((18446744073709551615) & 63))) {
                            fmt.Println(v6, v4, v2)
                            v31 := GenS3{F0: 255, F1: 128, F2: 42}
                            v2 = v2
                            _ = v31
                        } else {
                            v4 += v4
                            v28 = v28
                            switch fi29 {
                            case -9223372036854775808:
                                v1 = append(v1, fi29)
                            default:
                                var v32 float32 = float32(float64(float32(0.1)))
                                _ = v32
                            }
                        }
                        v33 := GenS4{F0: 695, F1: -1}
                        if true {
                            var v34 uint = 18446744073709551615
                            var v35 float32 = v3
                            _ = v34
                            _ = v35
                        } else {
                            fmt.Println(v7, v6)
                            for fi36 := 0; fi36 < 6; fi36++ {
                                v37 := v5[("a" + "foo bar")]
                                v3 += v3
                                var v38 int8 = -128
                                _ = fi36
                                _ = v37
                                _ = v38
                            }
                        }
                        _ = v33
                    }
                    var v39 int64 = 42
                    _ = v39
                }
                v7 += (v7 ^ 32767)
            }
            _ = fi29
        }
    case "":
        _ = len(v5)
    }
    if false {
        if true {
            for fi40 := 0; fi40 < 19; fi40++ {
                if v2 {
                    for fi41 := 0; fi41 < 16; fi41++ {
                        v42 := v5[""]
                        _ = fi41
                        _ = v42
                    }
                }
                var v43 int = fi40
                _ = fi40
                _ = v43
            }
            v3 = float32((v4 + float64(v3)))
        } else {
            var v44 uint16 = 1
            v44 = 65535
            _ = v44
        }
    }
    switch 128 {
    case -9223372036854775808:
        if (-228 != 9223372036854775807) {
            for fi45 := 0; fi45 < 13; fi45++ {
                var v46 uint = 552
                var v47 uint32 = 2147483648
                _ = fi45
                _ = v46
                _ = v47
            }
            for fi48 := 0; fi48 < 9; fi48++ {
                v49 := GenS3{F0: 1, F1: fi48, F2: 1}
                _ = fi48
                _ = v49
            }
        } else {
            v50 := GenS5{F0: 2, F1: 42}
            _ = v50
        }
    case 1:
        var v51 int32 = v6
        _ = v51
    }
    var v52 uint16 = v28
    switch 1 {
    case 440:
        if len(v1) > 0 {
            v1[0] = -9223372036854775808
        }
    case -9223372036854775808:
        for fi53 := 0; fi53 < 5; fi53++ {
            switch int(v52) {
            case 1:
                if v2 {
                    v54 := GenS3{F0: 1, F1: (fi53 % ((fi53) | 1)), F2: 20}
                    _ = v54
                }
            case -1:
                fmt.Println(v6, v52)
            case -9223372036854775808:
                var v55 uint16 = (v52 + 2)
                _ = v55
            default:
                var v56 bool = ((v7 / ((v7) | 1)) > v7)
                _ = v56
            }
            v57 := GenS5{F0: 42, F1: 128}
            v2 = true
            _ = fi53
            _ = v57
        }
    }
    v3 = (v3 - v3)
    if true {
        for fi58 := 0; fi58 < 16; fi58++ {
            v59 := GenS4{F0: (v6 | -2147483648), F1: 9223372036854775807}
            _ = fi58
            _ = v59
        }
    }
    var v60 float32 = float32(float64(v3))
    v61 := GenS3{F0: 190, F1: 128, F2: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v60
    _ = v61
    _ = v7
}
