package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 bool; F2 int32 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int64, p1 int32, p2 int32) int {
    fmt.Println(p2, p0, p1)
    v7 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 47
}
func genFn1(p0 int16, p1 bool) string {
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS2{F0: "a"}
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: "foo bar"}
    _ = p0
    _ = p1
    _ = v10
    return "\n\t"
}

func main() {
    var v1 int16 = 42
    var v2 uint8 = 127
    v3 := []string{("\n\t" + ""), "hello", ""}
    var v4 bool = false
    var v5 uint16 = 0
    var v6 float32 = 0.1
    v6 = v6
    v11 := GenS1{F0: int(v2)}
    var v12 uint16 = (v5 * v5)
    var v13 uint = 42
    for fi14 := 0; fi14 < 6; fi14++ {
        switch (fi14 / ((fi14) | 1)) {
        case 204:
            var v15 int32 = 464
            _ = v15
        case 2147483647:
            v16 := GenS3{F0: 0, F1: v4, F2: -1000}
            _ = v16
        default:
            var v17 int16 = v1
            _ = v17
        }
        _ = fi14
    }
    fmt.Println(v1, v5)
    if (0 <= 2147483647) {
        if v4 {
            v2 = v2
            if (42 <= 42) {
                var v18 float64 = 1.5
                fmt.Println(v1)
                for fi19 := 0; fi19 < 6; fi19++ {
                    var v20 float32 = -2.25
                    fmt.Println(v5)
                    _ = fi19
                    _ = v20
                }
                _ = v18
            }
            v21 := GenS2{F0: "\n\t"}
            _ = v21
        } else {
            v5 -= v5
        }
        var v22 uint32 = 313
        for fi23 := 0; fi23 < 14; fi23++ {
            v1 -= (v1 - int16(fi23))
            for fi24 := 0; fi24 < 5; fi24++ {
                if v4 {
                    v22 -= (v22 / ((14) | 1))
                    v13 -= v13
                }
                var v25 int = fi24
                var v26 string = ("foo bar" + ("a" + ("\n\t" + "hello")))
                v6 = -2.25
                if (v4 || ((^(v13)) <= v13)) {
                    fmt.Println(fi24)
                    v22 = 2147483647
                    v26 = v26
                    var v27 int32 = 595
                    var v28 string = v26
                    _ = v27
                    _ = v28
                }
                if false {
                    var v29 uint = v13
                    fmt.Println(v22)
                    var v30 int32 = -2147483648
                    _ = v29
                    _ = v30
                }
                if ((fi23 - (fi24 - (fi24 & fi24))) < (v25 + (fi24 ^ fi23))) {
                    var v31 float32 = v6
                    if v4 {
                        var v32 uint32 = v22
                        v2 = uint8(v12)
                        var v33 int32 = -916
                        _ = v32
                        _ = v33
                    } else {
                        switch v26 {
                        case "a":
                            if (("\n\t" + "\n\t") <= v26) {
                                var v34 int64 = -1
                                _ = v34
                            }
                            switch v26 {
                            case "a":
                                fmt.Println(v6)
                            case "hello":
                                fmt.Println(v2, v2, v25)
                            case "\n\t":
                                v25 -= (v25 << ((4294967295) & 63))
                            default:
                                if v4 {
                                    var v35 int = fi23
                                    v36 := GenS3{F0: 187, F1: v4, F2: 1}
                                    var v37 uint64 = 4294967295
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                } else {
                                    var v38 int8 = 100
                                    v1 -= (v1 & v1)
                                    if v4 {
                                        var v39 uint16 = v12
                                        _ = v39
                                    } else {
                                        var v40 float32 = -2.25
                                        var v41 int16 = v1
                                        _ = v40
                                        _ = v41
                                    }
                                    _ = v38
                                }
                            }
                        case "foo bar":
                            var v42 uint16 = v12
                            _ = v42
                        default:
                            var v43 uint16 = (^(v5))
                            _ = v43
                        }
                    }
                    var v44 bool = v4
                    v12 = 228
                    v3 = append(v3, v26)
                    v45 := GenS3{F0: -958, F1: v4, F2: 114}
                    v22 += (v22 ^ (^(v22)))
                    var v46 bool = v44
                    v46 = ((^(v13)) <= 610)
                    var v47 uint = v13
                    v13 += v47
                    _ = v31
                    _ = v44
                    _ = v45
                    _ = v46
                    _ = v47
                } else {
                    v25 = (fi23 & 32767)
                    v48 := GenS3{F0: 1, F1: (((float32(0.0) - 0.1) - ((float32(1.5) * float32((1e10 - 3.14159))) / v6)) >= v6), F2: int32(v12)}
                    v49 := GenS4{F0: int32(v2), F1: v2}
                    switch v25 {
                    case 1:
                        switch (fi24 + (v25 / ((-2147483648) | 1))) {
                        case 380:
                            var v50 int64 = -446
                            _ = v50
                        case 32768:
                            fmt.Println(fi23)
                        case -9223372036854775808:
                            if (v5 < 0) {
                                fmt.Println(v13, v26, fi24)
                                fmt.Println(v6, v25, fi23)
                                v51 := GenS2{F0: v26}
                                _ = v51
                            }
                        default:
                            var v52 uint32 = (^(v22))
                            _ = v52
                        }
                    case 0:
                        v53 := GenS3{F0: 0, F1: (0.0 >= (0.0 + v6)), F2: 638}
                        _ = v53
                    default:
                        if (!false) {
                            switch (fi23 % ((fi23) | 1)) {
                            case 1:
                                var v54 float64 = (1.5 + 3.14159)
                                _ = v54
                            case -2147483648:
                                fmt.Println(v2, fi24, v5)
                            case 9223372036854775807:
                                v55 := GenS2{F0: "foo bar"}
                                _ = v55
                            }
                        } else {
                            var v56 float32 = v6
                            v57 := GenS3{F0: 42, F1: (fi24 <= fi24), F2: -20}
                            var v58 int64 = 773
                            _ = v56
                            _ = v57
                            _ = v58
                        }
                    }
                    _ = v48
                    _ = v49
                }
                _ = fi24
                _ = v25
                _ = v26
            }
            var v59 float32 = v6
            _ = fi23
            _ = v59
        }
        _ = v22
    }
    for fi60 := 0; fi60 < 7; fi60++ {
        for fi61 := 0; fi61 < 15; fi61++ {
            v4 = true
            _ = fi61
        }
        if v4 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v62 int64 = 9223372036854775807
            _ = v62
        }
        v63 := GenS5{F0: -1}
        _ = fi60
        _ = v63
    }
    if (307 < 65535) {
        v3 = append(v3, "a")
        fmt.Println(v2, v2, v6)
        if (v4 && (v4 || false)) {
            fmt.Println(v12)
        } else {
            switch -914 {
            case 32768:
                var v64 int16 = v1
                _ = v64
            case -830:
                fmt.Println(v2)
            case 300:
                var v65 uint8 = (v2 ^ (v2 / ((30) | 1)))
                _ = v65
            }
        }
    }
    for fi66 := 0; fi66 < 10; fi66++ {
        v2 = v2
        _ = fi66
    }
    if (65535 <= (v13 * (^(v13)))) {
        if (!(-2147483648 > -29)) {
            for fi67 := 0; fi67 < 9; fi67++ {
                v68 := GenS3{F0: 897, F1: (v2 > v2), F2: -921}
                _ = fi67
                _ = v68
            }
        }
        for fi69 := 0; fi69 < 6; fi69++ {
            var v70 uint8 = (v2 << (((v13 | (v13 / (((v13 / (((v13 & 1)) | 1))) | 1)))) & 63))
            v12 += v5
            _ = fi69
            _ = v70
        }
    }
    v4 = (("" + "a") >= "")
    v2 += (v2 | (v2 + (v2 >> (((v13 << (((v13 & v13)) & 63))) & 63))))
    v71 := GenS4{F0: 2147483647, F1: v2}
    v72 := GenS4{F0: -226, F1: v2}
    var v73 bool = (1 <= 0)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
    _ = v72
    _ = v73
}
