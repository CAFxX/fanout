package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int8 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int64; F1 uint; F2 string }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int32) uint64 {
    var v4 int16 = 1000
    fmt.Println(v4, v4)
    for fi5 := 0; fi5 < 9; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 uint64 = 4294967295
            var v8 uint = 255
            v7 -= v7
            _ = fi6
            _ = v7
            _ = v8
        }
        switch "" {
        case "foo bar":
            var v9 int64 = -1
            _ = v9
        }
        _ = fi5
    }
    var v10 bool = (4294967295 > 4294967295)
    _ = p0
    _ = v10
    _ = v4
    return 1
}
func genFn1() uint32 {
    switch 32768 {
    case -2147483648:
        var v11 uint = 2
        _ = v11
    }
    switch 4294967295 {
    case 1:
        var v12 float32 = (((((3.14159 * 0.1) + -2.25) / (0.1 + (1.5 * 0.1))) + (0.1 * 1e10)) - 1.5)
        _ = v12
    }
    var v13 uint32 = 2147483647
    v14 := []int{788, 128, -226}
    _ = v13
    _ = v14
    return (v13 / (((v13 / (((v13 & 2147483647)) | 1))) | 1))
}

func main() {
    var v1 uint32 = 371
    v2 := []int{}
    v3 := []int{-9223372036854775808, -2147483648, -9223372036854775808, -739}
    if true {
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := GenS1{F0: 2, F1: 42, F2: -101}
            var v17 float32 = 0.0
            switch ("" + "a") {
            case "\n\t":
                if (100 != 42) {
                    _ = len(v2)
                    for fi18 := 0; fi18 < 15; fi18++ {
                        fmt.Println(fi15)
                        _ = fi18
                    }
                    if (!false) {
                        v1 += (v1 >> ((65535) & 63))
                        var v19 uint = 18446744073709551615
                        _ = v19
                    }
                } else {
                    if (!(60 < -43)) {
                        fmt.Println(v1, v17, v17)
                    } else {
                        if len(v2) > 0 {
                            v2[0] = fi15
                        }
                        fmt.Println(v17, v1)
                    }
                    v20 := GenS1{F0: 0, F1: 232, F2: 93}
                    v1 -= v1
                    _ = v20
                }
            case "foo bar":
                v3 = append(v3, -627)
            default:
                v17 = (v17 + v17)
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        if true {
            if false {
                v1 = (^(v1))
            } else {
                var v21 uint = 736
                _ = v21
            }
        }
    } else {
        if (128 >= 1) {
            v22 := GenS2{F0: 0, F1: 16}
            v23 := GenS2{F0: 42, F1: 100}
            _ = v22
            _ = v23
        }
    }
    switch ("" + "a") {
    case "\n\t":
        for fi24 := 0; fi24 < 10; fi24++ {
            v1 += genFn1()
            v1 = (^(v1))
            _ = fi24
        }
    case "":
        fmt.Println(v1)
    default:
        for fi25 := 0; fi25 < 13; fi25++ {
            switch "hello" {
            case "\n\t":
                var v26 uint = 255
                _ = v26
            case "a":
                if (2 > 75) {
                    v1 += (v1 >> ((2) & 63))
                    for fi27 := 0; fi27 < 8; fi27++ {
                        v1 -= 42
                        switch (fi27 >> ((255) & 63)) {
                        case 32767:
                            var v28 uint = 435
                            _ = v28
                        case 0:
                            var v29 uint16 = 16
                            var v30 uint = 0
                            _ = v29
                            _ = v30
                        case -419:
                            v1 += v1
                        }
                        switch "" {
                        case "a":
                            var v31 int8 = 0
                            _ = v31
                        case "hello":
                            _ = len(v2)
                        case "":
                            if (!(!(!(-117 <= int8(v1))))) {
                                v32 := GenS1{F0: 65535, F1: 1000, F2: 25}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v1 += (v1 * 544)
                                _ = v32
                            }
                        default:
                            _ = len(v3)
                        }
                        _ = fi27
                    }
                } else {
                    if false {
                        for fi33 := 0; fi33 < 9; fi33++ {
                            var v34 float64 = -2.25
                            switch fi25 {
                            case 42:
                                v34 = 1e10
                            case -1:
                                v35 := GenS2{F0: 10, F1: 100}
                                _ = v35
                            case -78:
                                fmt.Println(v1, fi33, fi33)
                            default:
                                v34 -= v34
                            }
                            _ = fi33
                            _ = v34
                        }
                    } else {
                        switch ("\n\t" + ("a" + "foo bar")) {
                        case "hello":
                            v36 := GenS2{F0: -90, F1: 1}
                            _ = v36
                        case "a":
                            var v37 int32 = 191
                            _ = v37
                        case "":
                            fmt.Println(fi25, v1)
                        }
                        var v38 uint64 = genFn0(2147483647)
                        var v39 int64 = 42
                        _ = v38
                        _ = v39
                    }
                    if len(v2) > 0 {
                        v2[0] = fi25
                    }
                    v40 := GenS2{F0: 114, F1: -1}
                    _ = v40
                }
            case "foo bar":
                for fi41 := 0; fi41 < 3; fi41++ {
                    v1 += v1
                    v42 := GenS3{F0: 211, F1: 4294967295}
                    if false {
                        v1 = (v1 * (v1 % ((genFn1()) | 1)))
                        var v43 int16 = 295
                        if false {
                            var v44 uint32 = v1
                            if false {
                                v45 := GenS2{F0: 77, F1: -128}
                                _ = v45
                            }
                            _ = v44
                        } else {
                            var v46 int16 = -1
                            switch (fi25 | (fi41 ^ (^(fi25)))) {
                            case 983:
                                v43 += v43
                            case 128:
                                var v47 uint16 = 0
                                _ = v47
                            }
                            _ = v46
                        }
                        _ = v43
                    } else {
                        if (((((0.0 * 1.5) - 3.14159) * 0.1) * 1e10) > float32(1e10)) {
                            var v48 uint8 = 0
                            var v49 string = "foo bar"
                            if (-1 < 127) {
                                if len(v2) > 0 {
                                    v2[0] = (fi41 ^ (^(fi25)))
                                }
                                v50 := GenS3{F0: uint64(v48), F1: 65535}
                                var v51 bool = (878 > -32768)
                                if (!v51) {
                                    if (1.5 == (1e10 * (1e10 - 1.5))) {
                                        if v51 {
                                            v52 := GenS1{F0: 42, F1: 32767, F2: 94}
                                            _ = v52
                                        } else {
                                            v48 -= 2
                                        }
                                        var v53 uint64 = 806
                                        _ = v53
                                    } else {
                                        v48 -= uint8(v1)
                                        if v51 {
                                            fmt.Println(fi41, v48)
                                            v49 = "hello"
                                            if len(v3) > 0 {
                                                v3[0] = fi41
                                            }
                                        } else {
                                            v49 = (v49 + v49)
                                            v54 := GenS1{F0: 255, F1: 42, F2: 100}
                                            var v55 uint8 = (v48 - 12)
                                            _ = v54
                                            _ = v55
                                        }
                                    }
                                    if v51 {
                                        fmt.Println(fi25)
                                    } else {
                                        var v56 string = "foo bar"
                                        v56 = ("" + v49)
                                        v56 = "a"
                                        _ = v56
                                    }
                                    v51 = (!(!true))
                                    if v51 {
                                        switch 32768 {
                                        case -9223372036854775808:
                                            v1 -= (v1 / ((v1) | 1))
                                        case -1:
                                            var v57 uint64 = 1
                                            _ = v57
                                        }
                                        v58 := GenS1{F0: 32767, F1: 58, F2: -48}
                                        _ = v58
                                    }
                                } else {
                                    v1 = 2147483647
                                    fmt.Println(v49, v51)
                                }
                                v48 += (v48 ^ v48)
                                _ = v50
                                _ = v51
                            }
                            _ = v48
                            _ = v49
                        } else {
                            v1 = genFn1()
                            var v59 uint = 96
                            fmt.Println(v1, v1)
                            _ = v59
                        }
                        var v60 uint64 = 401
                        _ = v60
                    }
                    _ = fi41
                    _ = v42
                }
            default:
                fmt.Println(v1, fi25)
            }
            v1 = v1
            v1 += (v1 | v1)
            _ = fi25
        }
    }
    var v61 int = 127
    v62 := GenS4{F0: 1, F1: 2, F2: ""}
    for fi63 := 0; fi63 < 6; fi63++ {
        fmt.Println(v1, fi63, v61)
        _ = fi63
    }
    if (42 <= 867) {
        _ = len(v2)
        v64 := GenS2{F0: 41, F1: -108}
        v61 += 1
        _ = v64
    } else {
        for fi65 := 0; fi65 < 6; fi65++ {
            v1 -= (v1 | v1)
            _ = fi65
        }
        v61 = -1
        v1 = v1
    }
    var v66 string = "a"
    v67 := GenS1{F0: 32767, F1: 1, F2: -1}
    v66 = "foo bar"
    _ = len(v3)
    if ((!true) || true) {
        for fi68 := 0; fi68 < 6; fi68++ {
            for fi69 := 0; fi69 < 17; fi69++ {
                fmt.Println(v66, v61)
                _ = fi69
            }
            if (54 != 359) {
                var v70 bool = ((-2.25 - (float32(1.5) - 1.5)) < -2.25)
                _ = v70
            } else {
                for fi71 := 0; fi71 < 8; fi71++ {
                    v66 = "foo bar"
                    fmt.Println(v1)
                    _ = len(v2)
                    _ = fi71
                }
            }
            _ = fi68
        }
        var v72 uint = 255
        _ = v72
    } else {
        if true {
            var v73 uint32 = 891
            v74 := GenS3{F0: 42, F1: genFn1()}
            switch v66 {
            case "\n\t":
                v2 = append(v2, 320)
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = v73
            _ = v74
        } else {
            if true {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v1 = genFn1()
        }
        fmt.Println(v1)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v75 := GenS5{F0: 1}
    var v76 bool = (((!true) && false) && ((42 != 1) || false))
    var v77 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v61:", v61)
    fmt.Println("v66:", v66)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v61
    _ = v62
    _ = v66
    _ = v67
    _ = v75
    _ = v76
    _ = v77
}
