package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int }
type GenS2 struct { F0 int16; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int32; F1 uint8 }
func genFn0(p0 int64) int32 {
    for fi5 := 0; fi5 < 14; fi5++ {
        switch (-(fi5)) {
        case 4294967295:
            p0 += 505
        case 32767:
            p0 += -1
        case 1:
            v6 := GenS1{F0: fi5, F1: 0, F2: (fi5 << ((439) & 63))}
            _ = v6
        }
        switch "" {
        case "":
            v7 := []int{}
            _ = v7
        case "hello":
            v8 := GenS1{F0: (fi5 >> ((4294967295) & 63)), F1: 0, F2: fi5}
            _ = v8
        case "foo bar":
            var v9 float64 = 1e10
            _ = v9
        default:
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 float64 = 1.5
                v12 := GenS2{F0: 0, F1: 424, F2: 304}
                p0 += 609
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        _ = fi5
    }
    p0 -= p0
    var v13 int32 = 2147483647
    var v14 uint = 255
    var v15 int64 = (p0 - 1)
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    return 449
}
func genFn1(p0 int32, p1 bool) int16 {
    var v16 string = "hello"
    v16 = v16
    p0 = (p0 + (p0 & p0))
    _ = p0
    _ = p1
    _ = v16
    return 0
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 42
    var v3 uint64 = 2
    var v4 int16 = -32768
    for fi17 := 0; fi17 < 7; fi17++ {
        if len(v1) > 0 {
            v1[0] = (^(v2))
        }
        _ = fi17
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        if false {
            switch "" {
            case "\n\t":
                for fi19 := 0; fi19 < 20; fi19++ {
                    if (1e10 < 1e10) {
                        var v20 int64 = 211
                        if (!true) {
                            var v21 uint8 = 2
                            if (!(-1 < int32(v21))) {
                                var v22 string = "a"
                                _ = v22
                            }
                            _ = v21
                        } else {
                            switch ("" + "\n\t") {
                            case "":
                                v23 := GenS3{F0: 1}
                                _ = v23
                            }
                            var v24 float64 = 0.1
                            fmt.Println(fi18, v3)
                            if true {
                                v25 := GenS2{F0: 0, F1: 558, F2: 32767}
                                _ = v25
                            } else {
                                switch ("" + ((("\n\t" + "hello") + "\n\t") + (("\n\t" + "\n\t") + (("a" + "\n\t") + "hello")))) {
                                case "\n\t":
                                    if true {
                                        var v26 uint32 = 490
                                        _ = v26
                                    } else {
                                        v27 := GenS1{F0: (^(fi18)), F1: (^(v3)), F2: fi19}
                                        var v28 uint8 = 255
                                        _ = v27
                                        _ = v28
                                    }
                                case "hello":
                                    var v29 uint = 503
                                    _ = v29
                                case "foo bar":
                                    switch (fi18 / ((fi19) | 1)) {
                                    case 127:
                                        if false {
                                            v30 := GenS1{F0: -9223372036854775808, F1: 363, F2: -429}
                                            var v31 uint16 = 65535
                                            v3 = uint64(v31)
                                            _ = v30
                                            _ = v31
                                        } else {
                                            var v32 uint8 = 255
                                            var v33 int64 = (v20 >> ((255) & 63))
                                            _ = v32
                                            _ = v33
                                        }
                                    }
                                }
                                var v34 bool = ("\n\t" == "foo bar")
                                _ = v34
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v24
                        }
                        _ = v20
                    }
                    _ = fi19
                }
            case "a":
                fmt.Println(fi18, v4, v3)
            case "foo bar":
                for fi35 := 0; fi35 < 1; fi35++ {
                    fmt.Println(v4, v3, fi18)
                    switch ("" + "a") {
                    case "a":
                        v1 = append(v1, v2)
                    case "foo bar":
                        v3 = v3
                    }
                    _ = fi35
                }
            }
        } else {
            var v36 uint = 2
            var v37 uint16 = 42
            _ = v36
            _ = v37
        }
        v2 = (v2 | 42)
        if false {
            v38 := GenS1{F0: -957, F1: (v3 % ((v3) | 1)), F2: (fi18 * 2147483647)}
            if false {
                v39 := GenS2{F0: (^(v4)), F1: genFn0(0), F2: 2}
                var v40 int64 = 9223372036854775807
                for fi41 := 0; fi41 < 13; fi41++ {
                    var v42 uint = 184
                    var v43 uint32 = 331
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                _ = v39
                _ = v40
            }
            var v44 int32 = 625
            _ = v38
            _ = v44
        }
        _ = fi18
    }
    switch "" {
    case "hello":
        switch ("\n\t" + "") {
        case "hello":
            v2 = 45
        }
    }
    v1 = append(v1, 2147483647)
    _ = len(v1)
    switch "hello" {
    case "foo bar":
        v45 := GenS4{F0: 156, F1: 1}
        _ = v45
    case "":
        var v46 int8 = 100
        _ = v46
    case "\n\t":
        v4 -= genFn1(genFn0(42), ((v3 & v3) < (v3 ^ v3)))
    }
    if false {
        v3 = 0
        var v47 float64 = ((1.5 * 1e10) - (1.5 - 1.5))
        v48 := GenS1{F0: 307, F1: v3, F2: 0}
        _ = v47
        _ = v48
    }
    fmt.Println(v2, v4, v4)
    if false {
        var v49 uint16 = 0
        var v50 int32 = 302
        _ = v49
        _ = v50
    } else {
        var v51 string = "a"
        if false {
            switch 127 {
            case -9223372036854775808:
                for fi52 := 0; fi52 < 7; fi52++ {
                    var v53 uint64 = v3
                    for fi54 := 0; fi54 < 5; fi54++ {
                        if (42 > 42) {
                            var v55 uint64 = v53
                            v2 = v2
                            if len(v1) > 0 {
                                v1[0] = 891
                            }
                            _ = len(v1)
                            _ = v55
                        }
                        v1 = append(v1, 65535)
                        var v56 int16 = (v4 / ((v4) | 1))
                        _ = fi54
                        _ = v56
                    }
                    _ = fi52
                    _ = v53
                }
            case 4294967295:
                var v57 uint32 = v2
                _ = v57
            }
        } else {
            if (!(0 != 255)) {
                v51 = v51
                if len(v1) > 0 {
                    v1[0] = (v2 << ((18446744073709551615) & 63))
                }
            } else {
                for fi58 := 0; fi58 < 9; fi58++ {
                    if false {
                        switch ("" + v51) {
                        case "":
                            for fi59 := 0; fi59 < 6; fi59++ {
                                v3 = v3
                                if (genFn1(genFn0(-9223372036854775808), true) == genFn1(340, false)) {
                                    switch (fi58 & fi59) {
                                    case 2147483647:
                                        v51 = (v51 + v51)
                                    case 902:
                                        var v60 float32 = 1.5
                                        _ = v60
                                    default:
                                        var v61 string = (v51 + v51)
                                        _ = v61
                                    }
                                }
                                _ = fi59
                            }
                        case "hello":
                            switch fi58 {
                            case 9223372036854775807:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case 911:
                                var v62 uint32 = uint32(v3)
                                _ = v62
                            case -562:
                                v2 = (v2 ^ 511)
                            }
                        }
                    }
                    _ = fi58
                }
                _ = len(v1)
            }
            var v63 uint64 = 380
            var v64 int32 = -477
            _ = v63
            _ = v64
        }
        switch 1 {
        case -550:
            for fi65 := 0; fi65 < 7; fi65++ {
                v66 := GenS5{F0: -431, F1: 127}
                v4 = v4
                _ = fi65
                _ = v66
            }
        }
        _ = v51
    }
    var v67 uint64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v67
}
