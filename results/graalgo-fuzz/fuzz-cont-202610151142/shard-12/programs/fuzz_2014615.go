package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 int8; F1 string; F2 bool }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 uint64; F1 int16; F2 int64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int64; F1 string; F2 int32 }
type GenS10 struct { F0 int64; F1 int; F2 string }
type GenS11 struct { F0 uint8 }
func genFn0(p0 uint32) uint {
    v6 := []int{0, 127}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 string = ("foo bar" + (("\n\t" + "a") + ("" + "\n\t")))
    var v8 int = 4294967295
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 152
}

func main() {
    var v1 int32 = -2147483648
    var v2 int8 = -1
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    v5 := make(map[string]uint32)
    var v9 bool = (9223372036854775807 > 9223372036854775807)
    if (56 < -1) {
        v1 = int32(v2)
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 uint64 = 539
            fmt.Println(v2, v11)
            _ = fi10
            _ = v11
        }
    } else {
        fmt.Println(v1)
    }
    v12 := GenS1{F0: 921, F1: 2}
    v13 := []int{}
    var v14 uint = 4294967295
    for fi15 := 0; fi15 < 16; fi15++ {
        if v9 {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = len(v5)
        }
        _ = fi15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        switch "" {
        case "\n\t":
            if v9 {
                v17 := GenS1{F0: (fi16 % ((fi16) | 1)), F1: 42}
                v18 := GenS2{F0: -636, F1: 255}
                _ = v17
                _ = v18
            }
        default:
            v2 = v2
        }
        var v19 float64 = 0.0
        if (!(!v9)) {
            for fi20 := 0; fi20 < 4; fi20++ {
                v14 += v14
                var v21 int8 = v2
                _ = fi20
                _ = v21
            }
            var v22 uint32 = 0
            var v23 uint = (v14 ^ genFn0(2))
            _ = v22
            _ = v23
        } else {
            if (896 <= 674) {
                switch "" {
                case "a":
                    for fi24 := 0; fi24 < 10; fi24++ {
                        if v9 {
                            v2 = v2
                            if (-1 == 847) {
                                v25 := GenS3{F0: 47, F1: 1}
                                var v26 int8 = (v2 - 0)
                                _ = v25
                                _ = v26
                            } else {
                                v2 += 127
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(v2)
                                case "":
                                    if v9 {
                                        var v27 int64 = 0
                                        _ = v27
                                    } else {
                                        fmt.Println(fi24)
                                        switch "foo bar" {
                                        case "":
                                            if len(v13) > 0 {
                                                v13[0] = (fi16 >> ((v14) & 63))
                                            }
                                        case "a":
                                            v28 := GenS4{F0: int(v2), F1: -1}
                                            _ = v28
                                        case "\n\t":
                                            var v29 uint64 = 18446744073709551615
                                            _ = v29
                                        }
                                        v9 = false
                                    }
                                case "foo bar":
                                    v30 := GenS5{F0: v2, F1: "foo bar", F2: (v9 && (v2 == (-(v2))))}
                                    _ = v30
                                default:
                                    delete(v3, "foo bar")
                                }
                                v31 := GenS6{F0: 79, F1: 65535}
                                _ = v31
                            }
                            switch "foo bar" {
                            case "\n\t":
                                v32 := v4[1]
                                _ = v32
                            case "foo bar":
                                v14 = 4294967295
                            case "a":
                                var v33 int64 = -681
                                _ = v33
                            }
                        } else {
                            v34 := GenS1{F0: (fi24 * fi24), F1: 42}
                            _ = len(v4)
                            v35 := GenS7{F0: 1, F1: -32768, F2: 42}
                            _ = v34
                            _ = v35
                        }
                        v36 := GenS7{F0: 2, F1: 0, F2: 1}
                        _ = fi24
                        _ = v36
                    }
                case "":
                    if (2 >= 42) {
                        if (((v9 || (v2 == v2)) || v9) || v9) {
                            v13 = append(v13, (fi16 >> (((v14 * (v14 ^ 157))) & 63)))
                        }
                    } else {
                        v37 := GenS3{F0: (v2 - v2), F1: -8}
                        v13 = append(v13, fi16)
                        _ = v37
                    }
                }
                if (-368 == (v1 >> ((v14) & 63))) {
                    v2 += v2
                }
                var v38 int = 127
                _ = v38
            } else {
                v14 = v14
            }
        }
        _ = fi16
        _ = v19
    }
    if (v14 != (v14 / ((genFn0(1)) | 1))) {
        for fi39 := 0; fi39 < 17; fi39++ {
            if (int(v14) > (fi39 >> ((genFn0(1)) & 63))) {
                v1 = -2147483648
            }
            v40 := GenS8{F0: ("foo bar" + "\n\t")}
            for fi41 := 0; fi41 < 2; fi41++ {
                if (!(!(!(fi41 == -296)))) {
                    var v42 uint = (^(v14))
                    v42 += (v14 % ((v14) | 1))
                    v13 = append(v13, 32768)
                    _ = v42
                } else {
                    var v43 uint32 = 2147483647
                    v1 -= v1
                    fmt.Println(v43, v14)
                    _ = v43
                }
                v44 := GenS4{F0: fi39, F1: -1}
                if v9 {
                    v2 = v2
                    v14 = genFn0(2)
                }
                _ = fi41
                _ = v44
            }
            _ = fi39
            _ = v40
        }
        switch -9223372036854775808 {
        case 382:
            v2 = v2
        case 42:
            fmt.Println(v9)
        case -2147483648:
            v1 = v1
        }
    } else {
        if true {
            var v45 int8 = (v2 / ((126) | 1))
            switch -2147483648 {
            case -994:
                for fi46 := 0; fi46 < 11; fi46++ {
                    for fi47 := 0; fi47 < 2; fi47++ {
                        if (!(((-1 != 1) || v9) && v9)) {
                            var v48 string = "hello"
                            _ = v48
                        } else {
                            if (273 >= 2147483647) {
                                var v49 float64 = (1.5 / 1.5)
                                if len(v13) > 0 {
                                    v13[0] = (fi47 << (((v14 * (v14 ^ 224))) & 63))
                                }
                                _ = v49
                            } else {
                                fmt.Println(v2, v45)
                            }
                        }
                        _ = fi47
                    }
                    v50 := GenS9{F0: -9223372036854775808, F1: "a", F2: v1}
                    v13 = append(v13, int(v2))
                    _ = fi46
                    _ = v50
                }
            case 9223372036854775807:
                v51 := GenS10{F0: 216, F1: -2147483648, F2: ("hello" + ("hello" + ((("\n\t" + "hello") + "foo bar") + "")))}
                _ = v51
            case 0:
                for fi52 := 0; fi52 < 13; fi52++ {
                    switch "foo bar" {
                    case "foo bar":
                        v14 += 2
                    default:
                        for fi53 := 0; fi53 < 12; fi53++ {
                            switch fi52 {
                            case -9223372036854775808:
                                v45 += v45
                            case 1:
                                fmt.Println(v1)
                            case 0:
                                if v9 {
                                    if len(v13) > 0 {
                                        v13[0] = (-(fi53))
                                    }
                                    var v54 float32 = 0.0
                                    _ = v54
                                } else {
                                    v55 := GenS11{F0: 128}
                                    var v56 uint8 = 214
                                    v1 -= 190
                                    _ = v55
                                    _ = v56
                                }
                            }
                            v2 = (v2 & v2)
                            v57 := GenS4{F0: fi53, F1: -1}
                            _ = fi53
                            _ = v57
                        }
                    }
                    v58 := v4[0]
                    if true {
                        switch int(v2) {
                        case 4294967295:
                            if (v2 <= 0) {
                                fmt.Println(v58, v9, v45)
                                if (9223372036854775807 < 42) {
                                    _ = len(v13)
                                } else {
                                    delete(v3, "\n\t")
                                }
                                v9 = v9
                            } else {
                                for fi59 := 0; fi59 < 16; fi59++ {
                                    var v60 uint32 = 729
                                    v61 := v3[("\n\t" + (("\n\t" + "hello") + ("foo bar" + (("" + "\n\t") + "foo bar"))))]
                                    var v62 bool = v9
                                    _ = fi59
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                }
                            }
                        }
                        var v63 int64 = 0
                        _ = v63
                    }
                    _ = fi52
                    _ = v58
                }
            default:
                var v64 int32 = -528
                _ = v64
            }
            v14 = 548
            _ = v45
        }
    }
    var v65 int64 = 1
    v2 = int8(v14)
    for fi66 := 0; fi66 < 1; fi66++ {
        for fi67 := 0; fi67 < 9; fi67++ {
            var v68 string = (("\n\t" + ("\n\t" + ("" + "foo bar"))) + (("" + "a") + ""))
            var v69 float32 = 1e10
            v68 = (v68 + (("a" + v68) + v68))
            _ = fi67
            _ = v68
            _ = v69
        }
        var v70 uint8 = uint8(v14)
        v65 += v65
        _ = fi66
        _ = v70
    }
    var v71 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v65:", v65)
    fmt.Println("v71:", v71)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v71
    _ = v9
}
