package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int64; F1 int16 }
func genFn0(p0 uint16, p1 string, p2 uint) int16 {
    v5 := GenS1{F0: 0}
    fmt.Println(p1)
    if (9223372036854775807 >= 0) {
        p1 = p1
        p2 -= (p2 >> ((p2) & 63))
    } else {
        var v6 uint32 = 42
        _ = v6
    }
    fmt.Println(p0, p1)
    var v7 uint8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return 42
}
func genFn1(p0 bool) uint16 {
    var v8 int16 = -32768
    var v9 uint32 = 1
    var v10 float64 = -2.25
    var v11 int64 = 610
    _ = p0
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 string = "foo bar"
    v2 := []uint32{2147483648, 471, 42}
    var v3 int64 = -1
    var v4 bool = (42 >= 18446744073709551615)
    for fi12 := 0; fi12 < 16; fi12++ {
        if true {
            switch (-(fi12)) {
            case -826:
                fmt.Println(v3, v4)
            case 9223372036854775807:
                switch v1 {
                case "\n\t":
                    fmt.Println(v4, v3)
                case "":
                    var v13 float32 = 0.1
                    _ = v13
                case "a":
                    fmt.Println(fi12)
                default:
                    var v14 uint16 = 42
                    _ = v14
                }
            default:
                var v15 int = 9223372036854775807
                _ = v15
            }
            for fi16 := 0; fi16 < 9; fi16++ {
                switch (fi16 * (fi12 % ((fi16) | 1))) {
                case -1:
                    var v17 int32 = 0
                    _ = v17
                case 32767:
                    var v18 float64 = -2.25
                    _ = v18
                case -2147483648:
                    if v4 {
                        var v19 int16 = -1
                        _ = v19
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                default:
                    if (float32(1.5) != 0.1) {
                        if (!(-128 <= -80)) {
                            if false {
                                v20 := GenS1{F0: 1}
                                if (266 <= 2147483648) {
                                    var v21 float64 = 0.0
                                    v21 += v21
                                    fmt.Println(v21, v4)
                                    v22 := GenS2{F0: 1, F1: 0}
                                    fmt.Println(v1, fi16, fi16)
                                    _ = v21
                                    _ = v22
                                } else {
                                    var v23 uint64 = 4294967295
                                    _ = v23
                                }
                                var v24 uint64 = 18446744073709551615
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                switch v1 {
                                case "a":
                                    if (!true) {
                                        v24 = v24
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    } else {
                                        if v4 {
                                            v24 = (v24 ^ (v24 & v24))
                                            var v25 uint32 = 2147483647
                                            if ((v3 ^ v3) != (v3 << ((65) & 63))) {
                                                v26 := GenS3{F0: 232, F1: v24}
                                                v1 = v1
                                                _ = len(v2)
                                                _ = v26
                                            }
                                            var v27 uint = 42
                                            fmt.Println(fi16, v27)
                                            fmt.Println(v25, fi12)
                                            _ = v25
                                            _ = v27
                                        } else {
                                            if false {
                                                var v28 int16 = genFn0(74, "foo bar", 255)
                                                v1 = (v1 + "a")
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v28
                                            } else {
                                                v1 = "foo bar"
                                                var v29 int = (-(fi12))
                                                _ = v29
                                            }
                                            switch (fi16 | fi12) {
                                            case -1:
                                                v30 := GenS4{F0: fi12, F1: 711}
                                                _ = v30
                                            case -2147483648:
                                                v24 = (^(v24))
                                            default:
                                                fmt.Println(v3, v24)
                                            }
                                            var v31 uint8 = 127
                                            _ = v31
                                        }
                                        if (!v4) {
                                            var v32 int16 = -265
                                            v33 := GenS5{F0: 2, F1: (v1 + (("" + v1) + v1))}
                                            _ = v32
                                            _ = v33
                                        } else {
                                            var v34 int = (fi12 << ((301) & 63))
                                            _ = len(v2)
                                            v35 := GenS1{F0: 65535}
                                            _ = v34
                                            _ = v35
                                        }
                                        var v36 int64 = 9223372036854775807
                                        _ = v36
                                    }
                                    fmt.Println(fi12)
                                default:
                                    v1 = ""
                                }
                                _ = v20
                                _ = v24
                            } else {
                                switch "\n\t" {
                                case "a":
                                    switch (fi16 / (((^(fi12))) | 1)) {
                                    case -2147483648:
                                        if (!(-30 > 92)) {
                                            var v37 int16 = 42
                                            _ = v37
                                        }
                                    case 32768:
                                        v4 = v4
                                    }
                                default:
                                    if len(v2) > 0 {
                                        v2[0] = 44
                                    }
                                }
                                v4 = v4
                            }
                            var v38 uint64 = 4294967295
                            _ = v38
                        } else {
                            v39 := GenS6{F0: "foo bar"}
                            _ = v39
                        }
                    } else {
                        v2 = append(v2, 4294967295)
                        v4 = v4
                        var v40 uint8 = 1
                        _ = v40
                    }
                }
                if (16 == 1) {
                    var v41 int32 = -634
                    if v4 {
                        fmt.Println(fi16, v3, v41)
                    } else {
                        var v42 uint16 = 2
                        v42 = v42
                        _ = v42
                    }
                    if false {
                        var v43 uint64 = 204
                        _ = v43
                    } else {
                        _ = len(v2)
                    }
                    _ = v41
                } else {
                    if (!false) {
                        v44 := GenS7{F0: -951}
                        _ = v44
                    }
                    v45 := GenS1{F0: 0}
                    _ = len(v2)
                    _ = v45
                }
                _ = fi16
            }
        }
        if true {
            fmt.Println(v1, v1, v3)
            var v46 uint = 405
            var v47 bool = true
            _ = v46
            _ = v47
        } else {
            var v48 uint16 = 0
            v1 = (v1 + (v1 + v1))
            _ = v48
        }
        _ = fi12
    }
    fmt.Println(v1)
    var v49 uint64 = 2
    if v4 {
        if len(v2) > 0 {
            v2[0] = 728
        }
        if false {
            v3 += (v3 << ((4294967295) & 63))
            if v4 {
                var v50 int8 = 0
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
                _ = v50
            } else {
                var v51 int = 127
                fmt.Println(v1)
                for fi52 := 0; fi52 < 19; fi52++ {
                    v49 += v49
                    _ = fi52
                }
                _ = v51
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v4 = (1 > 65535)
        }
    } else {
        for fi53 := 0; fi53 < 6; fi53++ {
            for fi54 := 0; fi54 < 15; fi54++ {
                v55 := GenS6{F0: "a"}
                v49 -= 1
                switch v1 {
                case "foo bar":
                    switch fi53 {
                    case -658:
                        v1 = ("hello" + (v1 + "hello"))
                    case 9223372036854775807:
                        var v56 int32 = 243
                        _ = v56
                    case -429:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                case "":
                    switch "foo bar" {
                    case "a":
                        var v57 string = (v1 + ("" + v1))
                        _ = v57
                    case "hello":
                        var v58 float32 = 0.0
                        _ = v58
                    default:
                        fmt.Println(v49)
                    }
                }
                _ = fi54
                _ = v55
            }
            var v59 bool = v4
            switch "" {
            case "hello":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "\n\t":
                v4 = ((!(((v49 - v49) <= (v49 * (v49 % ((v49) | 1)))) && v4)) && (1 < 42))
            case "":
                _ = len(v2)
            default:
                v60 := GenS4{F0: (fi53 % (((fi53 & (-(fi53)))) | 1)), F1: 0}
                _ = v60
            }
            _ = fi53
            _ = v59
        }
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    }
    for fi61 := 0; fi61 < 3; fi61++ {
        switch "" {
        case "a":
            for fi62 := 0; fi62 < 7; fi62++ {
                v1 = (((v1 + ("foo bar" + v1)) + v1) + ((v1 + v1) + v1))
                _ = fi62
            }
        case "":
            var v63 uint = 65535
            _ = v63
        default:
            switch "a" {
            case "foo bar":
                fmt.Println(v4, v1, v49)
            case "hello":
                for fi64 := 0; fi64 < 16; fi64++ {
                    fmt.Println(v3, v1)
                    v1 = v1
                    var v65 uint32 = 2147483648
                    _ = fi64
                    _ = v65
                }
            default:
                v66 := GenS8{F0: v3, F1: -874}
                _ = v66
            }
        }
        v4 = (1 >= 28)
        _ = fi61
    }
    v1 = v1
    fmt.Println(v49, v3)
    if v4 {
        v1 = (v1 + "")
        if v4 {
            v3 = 990
        } else {
            switch 329 {
            case 2147483647:
                v67 := GenS1{F0: 2}
                _ = v67
            default:
                var v68 bool = (!(!((-2.25 > -2.25) || (v49 >= 42))))
                _ = v68
            }
            v4 = v4
            if v4 {
                switch 0 {
                case -2147483648:
                    switch (v1 + v1) {
                    case "foo bar":
                        var v69 uint = 42
                        _ = v69
                    case "\n\t":
                        var v70 uint = 106
                        _ = v70
                    case "":
                        switch 909 {
                        case 128:
                            v1 = v1
                        }
                    default:
                        var v71 int64 = (v3 - v3)
                        _ = v71
                    }
                case 2147483647:
                    var v72 string = v1
                    _ = v72
                case 1:
                    v1 = v1
                }
            }
        }
        for fi73 := 0; fi73 < 5; fi73++ {
            var v74 int64 = (v3 + 0)
            if v4 {
                fmt.Println(v74, v4)
                v3 -= 42
                v74 += (v74 << ((1) & 63))
            } else {
                for fi75 := 0; fi75 < 2; fi75++ {
                    fmt.Println(fi75, v4, fi75)
                    var v76 int = fi75
                    var v77 bool = (v4 || v4)
                    _ = fi75
                    _ = v76
                    _ = v77
                }
                v4 = false
            }
            var v78 float32 = 3.14159
            _ = fi73
            _ = v74
            _ = v78
        }
    } else {
        var v79 int32 = -28
        _ = v79
    }
    for fi80 := 0; fi80 < 16; fi80++ {
        var v81 uint = 1
        v3 += v3
        _ = fi80
        _ = v81
    }
    var v82 int = 4294967295
    v49 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v82
}
