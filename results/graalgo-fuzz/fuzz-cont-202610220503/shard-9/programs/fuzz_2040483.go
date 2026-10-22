package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint16; F2 bool }
type GenS3 struct { F0 uint8; F1 int64; F2 string }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64 }
func genFn0(p0 uint16, p1 int32, p2 int) int8 {
    switch "hello" {
    case "hello":
        v7 := GenS1{F0: -128}
        _ = v7
    case "":
        switch "hello" {
        case "":
            v8 := []int{(^(p2)), p2, (p2 << ((42) & 63)), (p2 | (p2 % (((p2 % ((p2) | 1))) | 1)))}
            _ = v8
        case "hello":
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 uint8 = 127
                _ = fi9
                _ = v10
            }
        case "foo bar":
            var v11 bool = (true || true)
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := []int{fi12}
            for fi14 := 0; fi14 < 7; fi14++ {
                _ = len(v13)
                if (42 < 1) {
                    var v15 int16 = 0
                    _ = v15
                }
                p0 = (p0 % (((p0 + 358)) | 1))
                if false {
                    v16 := GenS2{F0: -128, F1: (p0 % ((2) | 1)), F2: (p0 < p0)}
                    _ = v16
                }
                _ = fi14
            }
            switch "a" {
            case "":
                var v17 uint = 18446744073709551615
                _ = v17
            }
            _ = fi12
            _ = v13
        }
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 bool = true
            switch "\n\t" {
            case "":
                var v21 uint32 = 2147483648
                _ = v21
            default:
                p1 = (p1 << ((194) & 63))
            }
            p2 += (fi19 >> ((4294967295) & 63))
            _ = fi19
            _ = v20
        }
        v22 := GenS2{F0: -21, F1: p0, F2: true}
        var v23 string = "a"
        _ = fi18
        _ = v22
        _ = v23
    }
    v24 := GenS3{F0: 1, F1: 9223372036854775807, F2: ("hello" + ("hello" + "\n\t"))}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return -1
}
func genFn1(p0 string) uint8 {
    for fi25 := 0; fi25 < 1; fi25++ {
        fmt.Println(fi25, p0)
        var v26 int = (fi25 | fi25)
        for fi27 := 0; fi27 < 15; fi27++ {
            v26 -= (v26 >> ((65535) & 63))
            var v28 int16 = 0
            _ = fi27
            _ = v28
        }
        _ = fi25
        _ = v26
    }
    p0 = p0
    var v29 float32 = (1.5 * ((-2.25 * 1.5) - -2.25))
    var v30 float32 = v29
    _ = p0
    _ = v29
    _ = v30
    return 42
}

func main() {
    var v1 uint16 = 376
    v2 := []string{("" + ("" + ("\n\t" + "foo bar")))}
    v3 := make(map[uint32]int)
    var v4 int64 = -9223372036854775808
    var v5 int64 = v4
    var v6 uint = 4294967295
    var v31 float64 = 1e10
    if false {
        v5 = -9223372036854775808
        var v32 uint16 = (v1 % (((v1 & uint16(v6))) | 1))
        _ = v32
    } else {
        switch 127 {
        case 4294967295:
            switch 1 {
            case 230:
                var v33 float32 = 1.5
                _ = v33
            case -9223372036854775808:
                var v34 uint = 255
                _ = v34
            }
        case 419:
            switch "a" {
            case "":
                switch 1 {
                case -680:
                    v4 = (v5 ^ 157)
                case -204:
                    v35 := v3[2147483648]
                    _ = v35
                case 664:
                    var v36 uint8 = 54
                    _ = v36
                }
            case "\n\t":
                if ((42 <= 225) || false) {
                    switch "a" {
                    case "a":
                        var v37 int8 = -128
                        _ = v37
                    }
                }
            default:
                if true {
                    if (((("\n\t" + ("hello" + ("" + ("" + "hello")))) + "foo bar") <= "\n\t") || false) {
                        v38 := GenS3{F0: 241, F1: (v4 & 1), F2: "hello"}
                        _ = v38
                    } else {
                        v39 := GenS4{F0: 127, F1: 65535}
                        if true {
                            var v40 int64 = int64(v1)
                            switch ("\n\t" + "") {
                            case "":
                                v41 := GenS2{F0: 0, F1: v1, F2: ((-2.25 == ((float32((0.0 - ((-2.25 - v31) - 0.0))) / (-2.25 / 0.1)) / (1e10 - 0.1))) || (true || ("\n\t" <= "\n\t")))}
                                _ = v41
                            case "\n\t":
                                fmt.Println(v6, v1, v5)
                            case "hello":
                                fmt.Println(v40, v40)
                            default:
                                if (18446744073709551615 > 1) {
                                    var v42 uint = (v6 | (^(v6)))
                                    var v43 float64 = v31
                                    if (!(840 >= -928)) {
                                        switch 32767 {
                                        case -1:
                                            var v44 float32 = 3.14159
                                            _ = v44
                                        default:
                                            if ((1.5 - ((1e10 + (3.14159 / -2.25)) + 0.0)) == 3.14159) {
                                                v45 := GenS2{F0: -65, F1: 949, F2: true}
                                                fmt.Println(v6)
                                                _ = v45
                                            } else {
                                                v5 += 42
                                                if true {
                                                    if true {
                                                        fmt.Println(v40)
                                                        if true {
                                                            for fi46 := 0; fi46 < 15; fi46++ {
                                                                v40 = v40
                                                                var v47 int8 = -29
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    if (439 <= -131) {
                                                                        var v48 bool = (((2 > 2) || true) || (706 >= 2147483648))
                                                                        _ = v48
                                                                    } else {
                                                                        var v49 uint16 = (v1 + 1)
                                                                        _ = v49
                                                                    }
                                                                }
                                                                _ = fi46
                                                                _ = v47
                                                            }
                                                            var v50 string = "hello"
                                                            switch 1 {
                                                            case -2147483648:
                                                                var v51 string = v50
                                                                _ = v51
                                                            case -999:
                                                                v2 = append(v2, v50)
                                                            case 719:
                                                                var v52 uint16 = v1
                                                                _ = v52
                                                            }
                                                            _ = v50
                                                        } else {
                                                            if true {
                                                                if (0 < 0) {
                                                                    fmt.Println(v43)
                                                                }
                                                            } else {
                                                                v53 := v3[2]
                                                                v2 = append(v2, (((((((("a" + "\n\t") + ("\n\t" + "foo bar")) + "hello") + ("foo bar" + "")) + "foo bar") + "foo bar") + "foo bar") + ("foo bar" + "hello")))
                                                                var v54 uint16 = (v1 >> ((v6) & 63))
                                                                _ = v53
                                                                _ = v54
                                                            }
                                                        }
                                                        var v55 int32 = -2147483648
                                                        _ = v55
                                                    }
                                                    var v56 int16 = 1
                                                    _ = v56
                                                }
                                            }
                                        }
                                        var v57 float64 = 0.0
                                        v58 := GenS4{F0: 0, F1: 2}
                                        _ = v57
                                        _ = v58
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        for fi59 := 0; fi59 < 17; fi59++ {
                                            var v60 int8 = -1
                                            fmt.Println(v40, v6, v31)
                                            v43 -= (v31 / (v43 + v43))
                                            _ = fi59
                                            _ = v60
                                        }
                                        fmt.Println(v42, v31, v4)
                                    }
                                    _ = v42
                                    _ = v43
                                } else {
                                    fmt.Println(v5, v31, v31)
                                }
                            }
                            var v61 int64 = v5
                            _ = v40
                            _ = v61
                        } else {
                            v62 := GenS5{F0: 32767}
                            _ = v62
                        }
                        if false {
                            var v63 string = "foo bar"
                            if (0.0 < (v31 + ((float64((float32((v31 - 0.1)) * ((0.0 - (1e10 / 0.1)) / (float32(1e10) + 1.5)))) - -2.25) - (v31 * 1e10)))) {
                                v31 -= 0.0
                                var v64 bool = true
                                _ = v64
                            }
                            _ = v63
                        } else {
                            v65 := GenS6{F0: (v4 >> ((v6) & 63))}
                            delete(v3, 2)
                            _ = v65
                        }
                        _ = v39
                    }
                    var v66 int64 = (v5 & (-(v4)))
                    for fi67 := 0; fi67 < 10; fi67++ {
                        fmt.Println(v6, fi67, v1)
                        switch 32768 {
                        case 9223372036854775807:
                            if (!true) {
                                v2 = append(v2, "a")
                            } else {
                                var v68 uint32 = uint32(v1)
                                _ = v68
                            }
                        default:
                            var v69 int16 = -32768
                            _ = v69
                        }
                        _ = fi67
                    }
                    _ = v66
                }
            }
        case 42:
            for fi70 := 0; fi70 < 1; fi70++ {
                v1 = 32767
                _ = fi70
            }
        default:
            if false {
                v6 -= uint(v1)
                switch 42 {
                case -2147483648:
                    var v71 int64 = v5
                    _ = v71
                case -9223372036854775808:
                    var v72 uint32 = 2147483647
                    _ = v72
                case 0:
                    fmt.Println(v4, v31, v31)
                default:
                    switch 728 {
                    case -572:
                        v31 = float64(1.5)
                    case 4294967295:
                        if ((true || ((!true) || (2 > 4294967295))) || (true && (85 == 255))) {
                            var v73 int32 = 1
                            var v74 int64 = (v5 & v5)
                            _ = v73
                            _ = v74
                        } else {
                            v2 = append(v2, (("hello" + ("foo bar" + "foo bar")) + "foo bar"))
                        }
                    }
                }
            } else {
                for fi75 := 0; fi75 < 13; fi75++ {
                    v6 = v6
                    for fi76 := 0; fi76 < 2; fi76++ {
                        v5 -= 742
                        v5 = v5
                        fmt.Println(v5)
                        _ = fi76
                    }
                    _ = fi75
                }
                fmt.Println(v6, v4)
                switch ((("a" + "hello") + "hello") + "hello") {
                case "hello":
                    v77 := GenS5{F0: 128}
                    _ = v77
                case "foo bar":
                    v78 := GenS3{F0: 127, F1: v4, F2: ("\n\t" + "hello")}
                    _ = v78
                case "":
                    for fi79 := 0; fi79 < 11; fi79++ {
                        var v80 int32 = int32(v4)
                        v81 := GenS7{F0: v5}
                        _ = fi79
                        _ = v80
                        _ = v81
                    }
                }
            }
        }
        var v82 uint32 = 1
        _ = len(v2)
        _ = v82
    }
    _ = len(v2)
    var v83 bool = (genFn0(v1, -1, 32768) <= 0)
    fmt.Println(v6)
    var v84 int = 351
    v6 = v6
    v85 := GenS1{F0: 100}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v83
    _ = v84
    _ = v85
}
