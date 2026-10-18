package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 uint8; F1 int; F2 uint16 }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 int16 }
func genFn0(p0 int32) uint32 {
    fmt.Println(p0)
    p0 = (p0 % (((p0 ^ p0)) | 1))
    for fi4 := 0; fi4 < 9; fi4++ {
        p0 += p0
        for fi5 := 0; fi5 < 9; fi5++ {
            p0 = p0
            var v6 float64 = (1.5 + (3.14159 * -2.25))
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    p0 = (-(p0))
    _ = p0
    return 65535
}
func genFn1() int64 {
    var v7 uint = 0
    for fi8 := 0; fi8 < 7; fi8++ {
        v7 = (v7 % ((0) | 1))
        v7 -= v7
        v9 := []int{fi8, fi8, 1}
        _ = fi8
        _ = v9
    }
    _ = v7
    return -83
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 9223372036854775807
    var v3 int8 = -128
    var v10 int64 = genFn1()
    if (2147483647 == -335) {
        for fi11 := 0; fi11 < 16; fi11++ {
            fmt.Println(v2, v10)
            var v12 int32 = int32(v3)
            v2 -= v10
            _ = fi11
            _ = v12
        }
        fmt.Println(v2, v10, v3)
    }
    if true {
        v13 := GenS1{F0: 1, F1: 202}
        var v14 int64 = v10
        _ = v13
        _ = v14
    } else {
        var v15 int8 = v3
        fmt.Println(v3, v10)
        var v16 float64 = float64(-2.25)
        _ = v15
        _ = v16
    }
    fmt.Println(v2, v10, v2)
    switch -2147483648 {
    case 127:
        switch 0 {
        case 0:
            var v17 int64 = (v2 >> ((2) & 63))
            _ = v17
        case -2147483648:
            fmt.Println(v2, v10)
        default:
            v10 = genFn1()
        }
    case -36:
        v18 := GenS1{F0: -1, F1: 123}
        _ = v18
    case -1:
        if false {
            v10 -= (v10 & 604)
            v19 := GenS2{F0: -827}
            _ = v19
        } else {
            v20 := GenS2{F0: 4294967295}
            _ = v20
        }
    default:
        var v21 uint = 65535
        _ = v21
    }
    v22 := GenS1{F0: -2147483648, F1: 2}
    for fi23 := 0; fi23 < 17; fi23++ {
        fmt.Println(v3)
        v3 += (v3 - -87)
        v24 := v1["a"]
        _ = fi23
        _ = v24
    }
    fmt.Println(v3, v10, v2)
    var v25 string = "foo bar"
    var v26 uint8 = 58
    v3 = (-(v3))
    switch v25 {
    case "":
        v27 := GenS1{F0: -949, F1: v26}
        _ = v27
    case "a":
        switch 42 {
        case -1:
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 uint8 = 227
                _ = fi28
                _ = v29
            }
        }
    case "foo bar":
        v25 = v25
    }
    if (!false) {
        var v30 bool = true
        for fi31 := 0; fi31 < 12; fi31++ {
            for fi32 := 0; fi32 < 1; fi32++ {
                v2 = v10
                _ = fi32
            }
            _ = fi31
        }
        _ = v30
    }
    switch 127 {
    case -2147483648:
        switch 0 {
        case 128:
            switch 32767 {
            case 281:
                v33 := GenS2{F0: -9223372036854775808}
                _ = v33
            case -466:
                switch -967 {
                case -260:
                    var v34 int32 = 2147483647
                    _ = v34
                case -779:
                    fmt.Println(v25)
                }
            case 0:
                v26 = (v26 | (v26 % ((v26) | 1)))
            }
        case 0:
            switch 32767 {
            case 9223372036854775807:
                v25 = "hello"
            case 134:
                v25 = ((v25 + (v25 + ("\n\t" + v25))) + "a")
            case 510:
                v35 := GenS1{F0: -47, F1: v26}
                _ = v35
            default:
                if (!(42 != 0)) {
                    if true {
                        var v36 int = 32767
                        if true {
                            v37 := []int{int(v3)}
                            _ = v37
                        }
                        for fi38 := 0; fi38 < 11; fi38++ {
                            _ = len(v1)
                            _ = fi38
                        }
                        _ = v36
                    }
                }
            }
        case 1:
            switch "hello" {
            case "\n\t":
                fmt.Println(v3)
            case "":
                fmt.Println(v25)
            case "a":
                v3 = v3
            default:
                v39 := GenS2{F0: 32767}
                _ = v39
            }
        }
    }
    v40 := GenS2{F0: 4294967295}
    if (!false) {
        if true {
            for fi41 := 0; fi41 < 15; fi41++ {
                for fi42 := 0; fi42 < 19; fi42++ {
                    if (-32768 != 1000) {
                        var v43 uint8 = 128
                        _ = v43
                    }
                    _ = fi42
                }
                _ = fi41
            }
        } else {
            v2 += (-(v2))
            switch v25 {
            case "":
                v10 -= (^(v2))
            case "foo bar":
                v10 -= genFn1()
            case "\n\t":
                switch -1 {
                case 0:
                    v44 := GenS3{F0: v26}
                    _ = v44
                }
            default:
                var v45 int16 = 1000
                _ = v45
            }
            var v46 uint64 = 508
            _ = v46
        }
        v26 = (v26 + (v26 / ((v26) | 1)))
    } else {
        switch v25 {
        case "hello":
            v47 := GenS1{F0: -195, F1: 0}
            _ = v47
        case "\n\t":
            if false {
                fmt.Println(v10)
                v48 := GenS4{F0: (v25 + v25)}
                fmt.Println(v25, v10, v26)
                _ = v48
            } else {
                if (!(((v26 & v26) < v26) || true)) {
                    v49 := GenS5{F0: 42, F1: -2147483648}
                    v50 := GenS6{F0: (v26 * v26), F1: 127, F2: 42}
                    _ = v49
                    _ = v50
                } else {
                    var v51 int = int(v3)
                    if true {
                        for fi52 := 0; fi52 < 15; fi52++ {
                            var v53 int32 = 2147483647
                            fmt.Println(v53, v2, v26)
                            _ = fi52
                            _ = v53
                        }
                    } else {
                        for fi54 := 0; fi54 < 12; fi54++ {
                            for fi55 := 0; fi55 < 18; fi55++ {
                                var v56 uint16 = 1
                                switch (("foo bar" + ((v25 + v25) + v25)) + v25) {
                                case "":
                                    v57 := v1[(v25 + "foo bar")]
                                    if false {
                                        var v58 int32 = 1
                                        var v59 string = "a"
                                        _ = v58
                                        _ = v59
                                    }
                                    v25 = (v25 + (v25 + "hello"))
                                    _ = v57
                                case "a":
                                    switch fi55 {
                                    case 9223372036854775807:
                                        if true {
                                            v60 := GenS7{F0: v25, F1: 4294967295}
                                            v61 := []int{fi55}
                                            switch (v51 + -297) {
                                            case 636:
                                                if false {
                                                    var v62 int = (fi55 * fi55)
                                                    if false {
                                                        v2 += genFn1()
                                                    }
                                                    v63 := GenS8{F0: 0}
                                                    _ = v62
                                                    _ = v63
                                                } else {
                                                    if false {
                                                        v64 := GenS4{F0: ("hello" + v25)}
                                                        v56 += v56
                                                        _ = v64
                                                    } else {
                                                        fmt.Println(v26, v2)
                                                        v65 := GenS7{F0: ("" + v25), F1: 0}
                                                        v56 = 42
                                                        _ = v65
                                                    }
                                                    v56 -= v56
                                                    if (v56 == uint16(v26)) {
                                                        v66 := GenS1{F0: 1, F1: v26}
                                                        var v67 int32 = -99
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                }
                                            case 128:
                                                if (4294967295 <= 0) {
                                                    var v68 uint32 = uint32(v56)
                                                    if true {
                                                        if true {
                                                            if false {
                                                                var v69 bool = (v56 >= (v56 >> ((441) & 63)))
                                                                var v70 int = fi55
                                                                if ((2 != 2) || true) {
                                                                    v26 = (v26 + v26)
                                                                    var v71 int32 = -183
                                                                    _ = v71
                                                                }
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                            v3 -= (v3 % (((v3 ^ -128)) | 1))
                                                        }
                                                    }
                                                    _ = v68
                                                }
                                            }
                                            _ = v60
                                            _ = v61
                                        } else {
                                            v72 := GenS6{F0: (v26 | v26), F1: fi55, F2: uint16(v26)}
                                            _ = v72
                                        }
                                    case 42:
                                        v51 = 0
                                    }
                                case "\n\t":
                                    delete(v1, v25)
                                default:
                                    fmt.Println(fi55, v26)
                                }
                                fmt.Println(fi55, v25, fi55)
                                var v73 float64 = 0.1
                                v26 = v26
                                _ = fi55
                                _ = v56
                                _ = v73
                            }
                            _ = fi54
                        }
                        switch (v51 << ((900) & 63)) {
                        case 32768:
                            if true {
                                v51 -= v51
                                if (0 > 42) {
                                    if true {
                                        fmt.Println(v25, v2)
                                    }
                                }
                            } else {
                                v1[(v25 + "a")] = v51
                            }
                        case 32767:
                            if true {
                                v10 -= (v2 + v2)
                                var v74 string = (v25 + "foo bar")
                                _ = v74
                            } else {
                                var v75 int16 = 32767
                                var v76 bool = (int32(v3) <= -1)
                                _ = v75
                                _ = v76
                            }
                        case 4294967295:
                            fmt.Println(v10, v2)
                        }
                    }
                    _ = v51
                }
            }
        case "foo bar":
            delete(v1, ((v25 + "a") + v25))
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v40
}
