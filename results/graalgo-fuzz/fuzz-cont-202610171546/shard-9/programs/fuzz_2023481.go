package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 uint32 }

func main() {
    v1 := []int{-190, 888, 128, 32767}
    var v2 int32 = -1
    v3 := make(map[int]uint32)
    v4 := []uint32{1, 2147483648, 65535, 236}
    var v5 int8 = -1
    var v6 uint = 1
    var v7 string = "a"
    if (-77 <= -32768) {
        _ = len(v4)
        for fi8 := 0; fi8 < 10; fi8++ {
            v7 = "foo bar"
            _ = fi8
        }
        if (v5 <= v5) {
            v9 := GenS1{F0: (false || (0 >= 127)), F1: -2147483648}
            for fi10 := 0; fi10 < 6; fi10++ {
                switch (^(fi10)) {
                case 32767:
                    for fi11 := 0; fi11 < 18; fi11++ {
                        v2 += (v2 / ((v2) | 1))
                        _ = fi11
                    }
                case 2147483647:
                    v2 -= 313
                }
                fmt.Println(v2)
                _ = fi10
            }
            _ = v9
        }
    } else {
        var v12 uint32 = 2147483647
        _ = v12
    }
    fmt.Println(v2)
    v6 = v6
    var v13 int = 42
    var v14 uint8 = 0
    v5 += v5
    v5 -= (-(v5))
    v13 = (v13 % (((v13 + v13)) | 1))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch v7 {
    case "foo bar":
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 int8 = v5
            var v17 float32 = -2.25
            for fi18 := 0; fi18 < 19; fi18++ {
                if (992 >= 42) {
                    if (!(true || ((2147483647 != 492) || false))) {
                        v6 -= 4294967295
                        if (true && (32767 == 65535)) {
                            fmt.Println(v6)
                            v19 := GenS1{F0: (v7 != v7), F1: v2}
                            var v20 int8 = 0
                            var v21 int16 = 0
                            var v22 int8 = v5
                            var v23 int = fi18
                            _ = v19
                            _ = v20
                            _ = v21
                            _ = v22
                            _ = v23
                        } else {
                            if (v16 <= (^(v5))) {
                                v5 -= 100
                                fmt.Println(v5, v7)
                            } else {
                                if true {
                                    var v24 uint = (^(v6))
                                    if true {
                                        if (42 >= 42) {
                                            var v25 uint16 = 983
                                            v26 := v3[(fi15 ^ (v13 % ((42) | 1)))]
                                            _ = v25
                                            _ = v26
                                        } else {
                                            var v27 string = "a"
                                            switch ("a" + v27) {
                                            case "hello":
                                                delete(v3, (^(fi18)))
                                            }
                                            fmt.Println(v24)
                                            _ = v27
                                        }
                                    } else {
                                        v28 := GenS2{F0: 0, F1: false}
                                        _ = v28
                                    }
                                    _ = v24
                                }
                                switch (v7 + v7) {
                                case "":
                                    if false {
                                        fmt.Println(v7)
                                        var v29 int16 = 1000
                                        var v30 int16 = 42
                                        _ = v29
                                        _ = v30
                                    } else {
                                        v5 -= (v5 ^ int8(v6))
                                    }
                                case "foo bar":
                                    var v31 uint16 = 648
                                    _ = v31
                                default:
                                    v32 := GenS1{F0: (false && false), F1: v2}
                                    _ = v32
                                }
                            }
                            fmt.Println(v17, fi18, v17)
                            if (!true) {
                                fmt.Println(v17, v2, v7)
                                var v33 uint8 = v14
                                var v34 int = (-(fi18))
                                if (!true) {
                                    switch ((v7 + "\n\t") + "foo bar") {
                                    case "\n\t":
                                        var v35 int64 = 1
                                        _ = v35
                                    default:
                                        v14 -= v14
                                    }
                                    v33 += v33
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = v33
                                _ = v34
                            } else {
                                fmt.Println(fi15)
                                fmt.Println(v6)
                            }
                        }
                        switch (fi15 - -635) {
                        case -44:
                            v36 := GenS1{F0: false, F1: (v2 - v2)}
                            _ = v36
                        default:
                            var v37 float64 = 1.5
                            _ = v37
                        }
                        switch fi15 {
                        case -9223372036854775808:
                            if ((!false) && (0 <= 2)) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                switch (-(v13)) {
                                case 127:
                                    var v38 int32 = 1
                                    v39 := GenS1{F0: true, F1: 0}
                                    v16 = (v5 + 42)
                                    _ = v38
                                    _ = v39
                                default:
                                    fmt.Println(fi18, fi15)
                                    var v40 int8 = (v16 | v16)
                                    switch "" {
                                    case "a":
                                        fmt.Println(v7, v2, v14)
                                    case "":
                                        switch v13 {
                                        case -2147483648:
                                            var v41 int32 = v2
                                            _ = v41
                                        case 127:
                                            v42 := GenS3{F0: v14}
                                            _ = v42
                                        case 9223372036854775807:
                                            v2 = -287
                                        default:
                                            v43 := GenS3{F0: 2}
                                            _ = v43
                                        }
                                    }
                                    v13 = fi15
                                    _ = v40
                                }
                                fmt.Println(v5, fi15, v5)
                            } else {
                                var v44 int8 = -128
                                _ = v44
                            }
                        case 0:
                            switch v7 {
                            case "\n\t":
                                if true {
                                    if (false && false) {
                                        v45 := GenS3{F0: v14}
                                        if (0 > 32767) {
                                            var v46 uint = v6
                                            v47 := GenS1{F0: ((((v5 >> ((v46) & 63)) == v16) && ((((float32((-2.25 * (1.5 / 1.5))) != v17) && (((true || (9223372036854775807 < int64(v14))) || (566 >= 0)) && (!(-2.25 == 3.14159)))) && ((("a" + v7) + ((v7 + "hello") + v7)) < (v7 + v7))) && true)) || true), F1: v2}
                                            _ = v46
                                            _ = v47
                                        } else {
                                            if ("hello" > (v7 + ((v7 + "") + "\n\t"))) {
                                                var v48 float32 = 1e10
                                                v3[fi18] = 0
                                                _ = len(v3)
                                                _ = v48
                                            } else {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                switch v7 {
                                                case "a":
                                                    v16 += 127
                                                case "\n\t":
                                                    fmt.Println(v17, v5, v2)
                                                default:
                                                    _ = len(v1)
                                                }
                                                var v49 float32 = float32(0.1)
                                                _ = v49
                                            }
                                            if true {
                                                var v50 uint32 = 2
                                                _ = v50
                                            } else {
                                                v6 = v6
                                            }
                                            v2 = v2
                                        }
                                        var v51 string = v7
                                        _ = v45
                                        _ = v51
                                    }
                                    if false {
                                        v4 = append(v4, 0)
                                        var v52 int8 = (v5 | (^(v5)))
                                        var v53 uint16 = 886
                                        _ = v52
                                        _ = v53
                                    } else {
                                        switch ("\n\t" + "hello") {
                                        case "":
                                            v54 := GenS4{F0: (v6 * 255), F1: uint8(v6), F2: (v14 ^ v14)}
                                            _ = v54
                                        case "foo bar":
                                            if true {
                                                fmt.Println(v6, v17, fi18)
                                                var v55 string = (v7 + (((v7 + v7) + v7) + (v7 + "hello")))
                                                switch (v7 + v55) {
                                                case "hello":
                                                    if ((32767 < 1) && true) {
                                                        var v56 uint = 18446744073709551615
                                                        _ = len(v3)
                                                        v17 -= (0.1 / 1e10)
                                                        _ = v56
                                                    }
                                                case "\n\t":
                                                    _ = len(v4)
                                                case "a":
                                                    var v57 int = (^(fi18))
                                                    _ = v57
                                                }
                                                switch 2147483647 {
                                                case 32767:
                                                    var v58 uint16 = 255
                                                    _ = v58
                                                case 2147483647:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                                _ = v55
                                            } else {
                                                var v59 int16 = -32768
                                                v60 := GenS3{F0: 42}
                                                var v61 int16 = (v59 + (v59 * v59))
                                                var v62 string = "\n\t"
                                                _ = v59
                                                _ = v60
                                                _ = v61
                                                _ = v62
                                            }
                                        case "a":
                                            fmt.Println(fi15, v13)
                                        default:
                                            switch "hello" {
                                            case "\n\t":
                                                v63 := GenS2{F0: v2, F1: true}
                                                _ = v63
                                            case "a":
                                                v64 := GenS4{F0: v6, F1: (v14 | 211), F2: (v14 / ((v14) | 1))}
                                                v14 += (v14 >> ((uint(v14)) & 63))
                                                _ = v64
                                            }
                                        }
                                    }
                                    v17 = float32((((1e10 - ((1e10 * -2.25) + (1.5 - float64(v17)))) - -2.25) + 3.14159))
                                    v13 = fi15
                                    fmt.Println(v6, fi15)
                                } else {
                                    switch fi15 {
                                    case 9223372036854775807:
                                        v7 = v7
                                    case -103:
                                        var v65 uint32 = 1
                                        _ = v65
                                    case -2147483648:
                                        if (2 == 1) {
                                            v5 += 127
                                            var v66 int = (-(fi18))
                                            _ = v66
                                        }
                                    }
                                    var v67 int = v13
                                    var v68 bool = (v17 < v17)
                                    v69 := GenS1{F0: v68, F1: v2}
                                    _ = v67
                                    _ = v68
                                    _ = v69
                                }
                            case "hello":
                                if true {
                                    switch "hello" {
                                    case "hello":
                                        _ = len(v1)
                                    case "":
                                        v7 = v7
                                    case "\n\t":
                                        v70 := GenS5{F0: 42}
                                        var v71 uint32 = 4294967295
                                        v72 := GenS6{F0: -1, F1: v2}
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                    }
                                    v73 := GenS7{F0: 2}
                                    _ = v73
                                }
                            default:
                                fmt.Println(v7, fi15, v2)
                            }
                        default:
                            var v74 uint = (v6 >> ((v6) & 63))
                            _ = v74
                        }
                    }
                } else {
                    v2 = v2
                    v75 := v3[(fi15 | v13)]
                    _ = v75
                }
                _ = fi18
            }
            _ = fi15
            _ = v16
            _ = v17
        }
    case "hello":
        v2 = (-(v2))
    case "a":
        var v76 uint8 = v14
        _ = v76
    default:
        switch "\n\t" {
        case "\n\t":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    var v77 int = int(v14)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v77
}
