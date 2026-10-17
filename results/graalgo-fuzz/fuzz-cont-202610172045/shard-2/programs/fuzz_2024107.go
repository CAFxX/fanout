package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint64; F1 string; F2 int }
type GenS4 struct { F0 int32; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint64; F1 int16; F2 string }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int64 }

func main() {
    v1 := []uint32{434, 2147483648, 4294967295}
    v2 := []string{}
    var v3 int32 = 1
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 int = fi4
        _ = fi4
        _ = v5
    }
    v3 += v3
    var v6 int = 2147483647
    for fi7 := 0; fi7 < 5; fi7++ {
        if true {
            v6 += v6
        }
        var v8 float32 = 1.5
        _ = fi7
        _ = v8
    }
    v3 -= (v3 << ((1) & 63))
    v9 := GenS1{F0: 128}
    v10 := GenS1{F0: v6}
    v6 += (v6 >> ((65535) & 63))
    if (((1e10 != -2.25) && (459 > 346)) || false) {
        switch ("foo bar" + "") {
        case "a":
            if false {
                v11 := GenS1{F0: v6}
                v6 = (-(v6))
                _ = v11
            }
        case "hello":
            switch v6 {
            case 4294967295:
                var v12 string = "foo bar"
                _ = v12
            default:
                var v13 int64 = 1
                _ = v13
            }
        case "":
            if false {
                if len(v1) > 0 {
                    v1[0] = 637
                }
            }
        }
    }
    switch "a" {
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "hello":
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 uint = 255
            for fi16 := 0; fi16 < 15; fi16++ {
                v17 := GenS1{F0: 710}
                v6 = fi16
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 19; fi18++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch 4294967295 {
                case 467:
                    if true {
                        var v19 string = "\n\t"
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                        if len(v2) > 0 {
                            v2[0] = v19
                        }
                        if (0.0 < (float64(3.14159) * float64(-2.25))) {
                            v15 -= (v15 % ((255) | 1))
                            fmt.Println(v3)
                            var v20 int8 = 100
                            var v21 uint64 = 4294967295
                            if false {
                                v1 = append(v1, 2)
                                v15 = (v15 << ((v15) & 63))
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            } else {
                                fmt.Println(fi14, v6, fi18)
                                if true {
                                    var v22 int16 = -32768
                                    if true {
                                        v23 := GenS2{F0: -1, F1: (v21 ^ v21), F2: 0}
                                        var v24 int16 = -32768
                                        var v25 int8 = 42
                                        fmt.Println(v24, fi18, v6)
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    } else {
                                        v26 := GenS3{F0: (v21 - v21), F1: "hello", F2: fi14}
                                        _ = v26
                                    }
                                    var v27 string = v19
                                    _ = v22
                                    _ = v27
                                }
                            }
                            _ = v20
                            _ = v21
                        }
                        _ = v19
                    } else {
                        v6 -= fi18
                        v1 = append(v1, 331)
                        v15 = (v15 ^ (v15 << ((18446744073709551615) & 63)))
                        if (2 > 1) {
                            v28 := GenS1{F0: (fi14 / ((fi18) | 1))}
                            v29 := GenS3{F0: 0, F1: "a", F2: (fi18 + -593)}
                            v3 = -649
                            _ = v28
                            _ = v29
                        } else {
                            v30 := GenS2{F0: -9223372036854775808, F1: 486, F2: 933}
                            var v31 string = "\n\t"
                            if ((3.14159 <= (-2.25 + (0.0 / (((3.14159 - 1e10) - (1e10 + -2.25)) / -2.25)))) || (484 <= (fi14 + 0))) {
                                v6 += (v6 / (((^(fi18))) | 1))
                            }
                            _ = v30
                            _ = v31
                        }
                        var v32 int16 = 0
                        switch fi14 {
                        case 0:
                            switch "hello" {
                            case "a":
                                v33 := GenS2{F0: -1, F1: 41, F2: uint16(v15)}
                                _ = v33
                            case "":
                                v34 := GenS4{F0: v3, F1: 0, F2: 618}
                                _ = v34
                            case "foo bar":
                                fmt.Println(v15, fi14)
                            default:
                                fmt.Println(v32, fi14, fi18)
                            }
                        default:
                            fmt.Println(fi14)
                        }
                        fmt.Println(v32)
                        _ = v32
                    }
                case 42:
                    fmt.Println(v6, fi14, fi14)
                case -9223372036854775808:
                    var v35 int64 = -430
                    _ = v35
                default:
                    switch ("foo bar" + "foo bar") {
                    case "hello":
                        v3 -= v3
                    default:
                        var v36 float32 = 1.5
                        _ = v36
                    }
                    fmt.Println(fi18)
                    if false {
                        _ = len(v2)
                        var v37 uint32 = 4294967295
                        v38 := GenS4{F0: 2147483647, F1: 0, F2: 2}
                        v39 := GenS3{F0: 3, F1: "a", F2: fi14}
                        if ((false || ((uint64(v37) < uint64(v37)) && true)) || (0.0 > 0.0)) {
                            v40 := GenS1{F0: v6}
                            v41 := GenS5{F0: 2, F1: 198, F2: "foo bar"}
                            var v42 float32 = 0.0
                            _ = v40
                            _ = v41
                            _ = v42
                        } else {
                            fmt.Println(fi14, v15, fi18)
                            if false {
                                v6 += (fi18 ^ (v6 ^ (fi18 ^ -9223372036854775808)))
                                if true {
                                    fmt.Println(fi18)
                                    if (229 != 128) {
                                        v3 = -2147483648
                                        if false {
                                            switch "a" {
                                            case "a":
                                                if false {
                                                    var v43 float32 = 3.14159
                                                    _ = v43
                                                }
                                            case "hello":
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            default:
                                                v37 -= v37
                                            }
                                            switch ("hello" + "foo bar") {
                                            case "":
                                                if len(v2) > 0 {
                                                    v2[0] = (((("foo bar" + "\n\t") + "a") + "a") + "hello")
                                                }
                                            case "hello":
                                                var v44 uint64 = 18446744073709551615
                                                _ = v44
                                            case "a":
                                                switch fi14 {
                                                case 128:
                                                    var v45 uint16 = 255
                                                    _ = v45
                                                case -9223372036854775808:
                                                    v15 = (v15 | (v15 + v15))
                                                }
                                            default:
                                                var v46 int = fi18
                                                _ = v46
                                            }
                                        }
                                    }
                                }
                            } else {
                                var v47 uint64 = 1
                                v37 -= (v37 ^ (v37 - v37))
                                var v48 uint16 = 65535
                                _ = v47
                                _ = v48
                            }
                        }
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    if true {
                        fmt.Println(v3, v3, v6)
                        switch (v6 << (((v15 + (v15 & (v15 + v15)))) & 63)) {
                        case 466:
                            v3 = (v3 >> ((18446744073709551615) & 63))
                        case -803:
                            switch ("foo bar" + "foo bar") {
                            case "":
                                v3 = (^(v3))
                            default:
                                var v49 int8 = 127
                                var v50 uint = (v15 << ((v15) & 63))
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v49
                                _ = v50
                            }
                        case 127:
                            v51 := GenS6{F0: true, F1: 0}
                            _ = v51
                        default:
                            var v52 float32 = 0.1
                            _ = v52
                        }
                        if (!true) {
                            var v53 uint32 = 730
                            fmt.Println(v6)
                            _ = v53
                        } else {
                            v54 := GenS7{F0: "a"}
                            _ = v54
                        }
                    }
                }
                _ = fi18
            }
            _ = fi14
            _ = v15
        }
    case "foo bar":
        v55 := GenS4{F0: v3, F1: -1, F2: 42}
        _ = v55
    }
    if false {
        if false {
            v3 -= v3
        } else {
            if false {
                v3 -= v3
            }
            v6 -= 0
            v56 := GenS7{F0: "\n\t"}
            _ = v56
        }
        if (105 == 0) {
            v3 -= -2147483648
            v57 := GenS5{F0: 4294967295, F1: 1, F2: ""}
            v3 += v3
            _ = v57
        }
    } else {
        v58 := GenS8{F0: 367}
        var v59 uint = 42
        _ = v58
        _ = v59
    }
    if false {
        v6 += v6
        v3 -= 0
        if false {
            v3 -= v3
            v60 := GenS4{F0: v3, F1: 127, F2: 32767}
            _ = v60
        } else {
            v3 = (v3 << ((65535) & 63))
            v2 = append(v2, ("" + "a"))
        }
    }
    v2 = append(v2, ((("\n\t" + "\n\t") + "\n\t") + "a"))
    fmt.Println(v3, v6)
    var v61 uint16 = 684
    var v62 uint = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v61
    _ = v62
    _ = v9
}
