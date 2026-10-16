package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 uint; F2 int64 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 uint32 = 0
    var v2 uint32 = v1
    var v3 bool = (false && (true || (0 != -2147483648)))
    v4 := make(map[uint32]int)
    v5 := []int{4294967295, -9223372036854775808}
    var v6 uint16 = 860
    v7 := GenS1{F0: "foo bar"}
    var v8 uint = 558
    var v9 uint64 = 952
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 uint16 = v6
        var v12 uint32 = (^(v1))
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: ("a" + "hello")}
    v4[uint32(v9)] = -9223372036854775808
    v9 -= (v9 >> (((v8 - (^(v8)))) & 63))
    fmt.Println(v9)
    v14 := GenS2{F0: 1, F1: (v6 - (^(v6)))}
    v15 := v4[(v2 - (v1 ^ (v1 << ((v8) & 63))))]
    switch 0 {
    case 275:
        fmt.Println(v3)
    case 128:
        _ = len(v5)
    default:
        if (v6 >= v6) {
            switch ("" + ("a" + "hello")) {
            case "a":
                var v16 int32 = 1
                _ = v16
            case "foo bar":
                var v17 int32 = 1
                _ = v17
            case "hello":
                for fi18 := 0; fi18 < 5; fi18++ {
                    v19 := GenS3{F0: v9}
                    var v20 uint64 = v9
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
            if true {
                switch v15 {
                case 42:
                    switch (v15 % (((-(v15))) | 1)) {
                    case 9223372036854775807:
                        fmt.Println(v1, v1, v8)
                    }
                case 9223372036854775807:
                    v9 += (v9 / ((689) | 1))
                }
                var v21 float64 = (1e10 + (0.1 + -2.25))
                fmt.Println(v9, v1)
                _ = v21
            }
            var v22 int32 = 363
            _ = v22
        } else {
            _ = len(v5)
            fmt.Println(v9)
        }
    }
    v1 += (v2 + v2)
    var v23 uint64 = (^(v9))
    if v3 {
        for fi24 := 0; fi24 < 1; fi24++ {
            for fi25 := 0; fi25 < 9; fi25++ {
                var v26 int32 = 971
                v23 += (v9 + v9)
                if v3 {
                    v27 := v4[v1]
                    switch v27 {
                    case 4294967295:
                        switch ("hello" + "foo bar") {
                        case "a":
                            fmt.Println(v6)
                        case "foo bar":
                            v28 := GenS3{F0: uint64(v1)}
                            _ = v28
                        }
                    case 864:
                        v1 -= (^(v2))
                    case -1:
                        v29 := GenS2{F0: 1, F1: (^(v6))}
                        _ = v29
                    default:
                        switch "\n\t" {
                        case "hello":
                            if false {
                                fmt.Println(v2, v27)
                                v4[2147483647] = 4294967295
                                if v3 {
                                    var v30 int16 = -32768
                                    if v3 {
                                        if ((-1 > 110) || false) {
                                            v31 := GenS3{F0: v23}
                                            switch "hello" {
                                            case "\n\t":
                                                fmt.Println(v6)
                                            case "hello":
                                                var v32 int = 949
                                                _ = v32
                                            case "a":
                                                var v33 uint8 = 42
                                                _ = v33
                                            default:
                                                switch "foo bar" {
                                                case "hello":
                                                    v4[65535] = v27
                                                case "foo bar":
                                                    v9 += v23
                                                }
                                            }
                                            fmt.Println(v30, fi24)
                                            _ = v31
                                        }
                                    }
                                    _ = v30
                                } else {
                                    v34 := GenS2{F0: 100, F1: (v6 + 255)}
                                    var v35 int64 = 190
                                    v15 += v15
                                    _ = v34
                                    _ = v35
                                }
                                v36 := GenS1{F0: (("hello" + "") + ("hello" + ("" + ("\n\t" + "hello"))))}
                                _ = v36
                            } else {
                                if (95 <= 1) {
                                    v4[v1] = 525
                                    if true {
                                        switch v27 {
                                        case 639:
                                            v37 := GenS4{F0: v8, F1: (v8 | v8), F2: -159}
                                            _ = v37
                                        case -503:
                                            var v38 uint64 = v23
                                            _ = v38
                                        }
                                        var v39 float32 = -2.25
                                        var v40 bool = ((568 <= 1000) && true)
                                        _ = v39
                                        _ = v40
                                    } else {
                                        var v41 int = 127
                                        v41 += fi25
                                        v8 += (v8 / ((uint(v2)) | 1))
                                        _ = v41
                                    }
                                    v42 := v4[(v2 + v2)]
                                    _ = v42
                                }
                                v1 = v2
                                switch fi25 {
                                case 4294967295:
                                    v43 := GenS2{F0: -128, F1: (^(v6))}
                                    _ = v43
                                case 32768:
                                    fmt.Println(v2)
                                default:
                                    var v44 int64 = -9223372036854775808
                                    _ = v44
                                }
                            }
                            v45 := GenS5{F0: (v26 + -2147483648)}
                            v2 += uint32(v23)
                            _ = v45
                        case "a":
                            v4[v2] = v27
                        default:
                            if len(v5) > 0 {
                                v5[0] = 721
                            }
                        }
                    }
                    _ = v27
                }
                _ = fi25
                _ = v26
            }
            _ = fi24
        }
        switch "" {
        case "foo bar":
            var v46 int8 = -121
            _ = v46
        }
    } else {
        if v3 {
            for fi47 := 0; fi47 < 11; fi47++ {
                delete(v4, 4294967295)
                var v48 int64 = 0
                _ = fi47
                _ = v48
            }
            var v49 uint32 = 0
            v3 = v3
            _ = v49
        }
        if (false || false) {
            if v3 {
                fmt.Println(v2, v23, v3)
            } else {
                _ = len(v5)
                if v3 {
                    fmt.Println(v8, v15, v9)
                    v50 := GenS3{F0: v9}
                    for fi51 := 0; fi51 < 4; fi51++ {
                        if (!v3) {
                            for fi52 := 0; fi52 < 1; fi52++ {
                                var v53 int = v15
                                var v54 int64 = -9223372036854775808
                                switch fi52 {
                                case 9223372036854775807:
                                    var v55 bool = true
                                    _ = v55
                                default:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                _ = fi52
                                _ = v53
                                _ = v54
                            }
                            var v56 uint32 = 7
                            fmt.Println(v56, v3)
                            _ = v56
                        } else {
                            v1 = 221
                            for fi57 := 0; fi57 < 7; fi57++ {
                                v3 = v3
                                v15 = 42
                                _ = fi57
                            }
                            for fi58 := 0; fi58 < 14; fi58++ {
                                var v59 uint = (v8 | 4294967295)
                                if true {
                                    if false {
                                        fmt.Println(v6, v15)
                                    } else {
                                        v60 := GenS2{F0: 100, F1: (v6 >> (((v8 * (v59 >> ((255) & 63)))) & 63))}
                                        _ = v60
                                    }
                                } else {
                                    var v61 uint = (v8 / ((v8) | 1))
                                    var v62 int16 = 42
                                    _ = v61
                                    _ = v62
                                }
                                _ = fi58
                                _ = v59
                            }
                        }
                        _ = fi51
                    }
                    _ = v50
                } else {
                    v9 += v9
                    var v63 int16 = 32767
                    v64 := GenS2{F0: -1, F1: 65535}
                    _ = v63
                    _ = v64
                }
                switch ("\n\t" + "") {
                case "foo bar":
                    fmt.Println(v1, v9)
                }
            }
            var v65 int8 = 42
            var v66 int8 = 100
            _ = v65
            _ = v66
        } else {
            v1 = (v2 / ((1) | 1))
        }
    }
    for fi67 := 0; fi67 < 10; fi67++ {
        var v68 int8 = -1
        _ = fi67
        _ = v68
    }
    var v69 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v7
    _ = v8
    _ = v9
}
