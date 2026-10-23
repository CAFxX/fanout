package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 bool; F1 uint; F2 string }
type GenS4 struct { F0 string; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint32; F1 uint; F2 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16; F1 string; F2 uint64 }
type GenS8 struct { F0 uint64; F1 uint }

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 547
    v3 := make(map[string]uint32)
    var v4 uint = 1
    var v5 uint64 = 2
    v6 := []uint32{}
    v3[("foo bar" + ("foo bar" + ("hello" + "")))] = 2147483647
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 float64 = (1.5 / 0.1)
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := GenS1{F0: ("hello" + "a"), F1: (v2 + 730), F2: v5}
            if (!(false || false)) {
                var v11 int32 = -2147483648
                var v12 uint32 = 867
                fmt.Println(fi9, v12, v8)
                _ = v11
                _ = v12
            } else {
                if ((fi7 >> (((v4 + 1)) & 63)) == (-(fi9))) {
                    v13 := v3["foo bar"]
                    v2 -= (^(v2))
                    switch (("foo bar" + "a") + (("\n\t" + "hello") + "hello")) {
                    case "\n\t":
                        v13 = v13
                    case "a":
                        var v14 uint64 = v5
                        _ = v14
                    case "":
                        fmt.Println(v4, fi7)
                    }
                    var v15 uint8 = 242
                    v16 := GenS1{F0: "hello", F1: (v5 << ((v4) & 63)), F2: (^(v2))}
                    _ = v13
                    _ = v15
                    _ = v16
                } else {
                    if true {
                        var v17 int16 = 42
                        v18 := GenS1{F0: ("hello" + "\n\t"), F1: 236, F2: (v2 % (((^(v5))) | 1))}
                        var v19 int64 = 1
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                }
            }
            _ = fi9
            _ = v10
        }
        var v20 uint8 = 254
        _ = fi7
        _ = v20
        _ = v8
    }
    var v21 bool = (606 < v4)
    for fi22 := 0; fi22 < 1; fi22++ {
        fmt.Println(v2)
        v3[(("hello" + "foo bar") + (("" + ("" + ("hello" + "\n\t"))) + "foo bar"))] = 1
        var v23 int32 = 1
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        if true {
            v4 = (v4 + (v4 % ((v4) | 1)))
        }
        _ = fi24
    }
    switch "foo bar" {
    case "":
        v3["foo bar"] = 1
    default:
        for fi25 := 0; fi25 < 2; fi25++ {
            fmt.Println(v4, v2, v21)
            v26 := GenS1{F0: (("hello" + "a") + "\n\t"), F1: v2, F2: (v2 / ((v2) | 1))}
            _ = fi25
            _ = v26
        }
    }
    switch (("\n\t" + "") + "\n\t") {
    case "":
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 int = (^(fi27))
            _ = fi27
            _ = v28
        }
    case "a":
        v29 := GenS1{F0: "hello", F1: (v5 | v2), F2: (v5 | (v2 + v5))}
        _ = v29
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        for fi31 := 0; fi31 < 4; fi31++ {
            switch (^(fi31)) {
            case 127:
                _ = len(v6)
            case 32767:
                fmt.Println(v2, v5)
            case 128:
                v32 := GenS1{F0: (("a" + "a") + ("\n\t" + ("a" + "\n\t"))), F1: 42, F2: 71}
                _ = v32
            default:
                v33 := GenS2{F0: (fi30 / ((fi31) | 1)), F1: 42}
                _ = v33
            }
            _ = fi31
        }
        _ = fi30
    }
    if v21 {
        v34 := GenS2{F0: -1, F1: 781}
        if v21 {
            for fi35 := 0; fi35 < 13; fi35++ {
                if ((!v21) && (!true)) {
                    if (!(0 < 175)) {
                        for fi36 := 0; fi36 < 12; fi36++ {
                            _ = len(v6)
                            switch (fi35 << ((v4) & 63)) {
                            case -64:
                                v37 := GenS1{F0: "", F1: (v5 - (v2 << ((961) & 63))), F2: 18446744073709551615}
                                _ = v37
                            case 127:
                                if true {
                                    var v38 int8 = -94
                                    v38 = -121
                                    switch (fi35 * -320) {
                                    case 9223372036854775807:
                                        v2 += 0
                                    default:
                                        v39 := GenS2{F0: fi35, F1: 32767}
                                        _ = v39
                                    }
                                    v2 = 373
                                    if (!("foo bar" < "a")) {
                                        v40 := v1[32768]
                                        v41 := GenS3{F0: (0 >= 2), F1: 1, F2: v40}
                                        v42 := GenS4{F0: "foo bar", F1: 22, F2: 33}
                                        v43 := GenS5{F0: 2147483648, F1: (v4 + v4), F2: 2}
                                        v44 := GenS4{F0: "\n\t", F1: 1, F2: 127}
                                        if (535 <= (v2 - v2)) {
                                            v1[(fi36 + fi36)] = v40
                                            fmt.Println(v21)
                                        }
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    }
                                    _ = v38
                                }
                                v45 := GenS2{F0: (fi36 / (((-(fi35))) | 1)), F1: 0}
                                switch 127 {
                                case -6:
                                    fmt.Println(fi36)
                                }
                                _ = v45
                            case 0:
                                switch -226 {
                                case 42:
                                    _ = len(v6)
                                }
                            default:
                                v5 = v2
                            }
                            _ = fi36
                        }
                    } else {
                        for fi46 := 0; fi46 < 20; fi46++ {
                            if (!false) {
                                v2 += 0
                                v2 -= 18446744073709551615
                            }
                            v3[""] = 42
                            _ = fi46
                        }
                        v6 = append(v6, 321)
                    }
                }
                _ = fi35
            }
            v2 -= v2
        } else {
            var v47 uint16 = 255
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if (false && (((v47 % (((v47 & 2)) | 1)) >= (^(v47))) && (0 > 4294967295))) {
                var v48 int32 = -2147483648
                if ((false || (v4 == (v4 / ((v4) | 1)))) || (382 == 27)) {
                    switch "hello" {
                    case "":
                        var v49 int16 = 1000
                        _ = v49
                    default:
                        if (1e10 > (1e10 - (float64(1e10) + 3.14159))) {
                            switch 2147483647 {
                            case 9223372036854775807:
                                for fi50 := 0; fi50 < 19; fi50++ {
                                    v48 = (v48 | (v48 % ((1) | 1)))
                                    var v51 int32 = 698
                                    _ = fi50
                                    _ = v51
                                }
                            default:
                                var v52 bool = false
                                _ = v52
                            }
                        } else {
                            fmt.Println(v48, v47)
                            var v53 uint32 = 65535
                            delete(v1, 42)
                            _ = v53
                        }
                    }
                    var v54 int32 = 2147483647
                    for fi55 := 0; fi55 < 10; fi55++ {
                        if ((1e10 >= 1e10) || ((v48 >> ((v4) & 63)) > (v48 + (^(v48))))) {
                            for fi56 := 0; fi56 < 20; fi56++ {
                                if false {
                                    var v57 float32 = (-2.25 + 0.0)
                                    v4 = (v4 >> ((v4) & 63))
                                    v58 := GenS6{F0: v2}
                                    _ = v57
                                    _ = v58
                                } else {
                                    fmt.Println(fi56, v47, v2)
                                    var v59 uint16 = 65535
                                    v60 := v1[(fi55 & (fi56 * fi56))]
                                    v2 = v5
                                    var v61 uint32 = 531
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                }
                                _ = fi56
                            }
                            v62 := GenS5{F0: 525, F1: 255, F2: (v4 | 1)}
                            var v63 float32 = (0.0 * (float32(((1.5 - ((-2.25 - 0.0) * 1e10)) + (3.14159 - 3.14159))) / 1e10))
                            _ = v62
                            _ = v63
                        }
                        _ = fi55
                    }
                    _ = v54
                }
                var v64 uint8 = 255
                _ = v48
                _ = v64
            }
            _ = v47
        }
        var v65 uint = v4
        _ = v34
        _ = v65
    } else {
        fmt.Println(v5)
    }
    if v21 {
        if v21 {
            fmt.Println(v2, v21, v5)
            if v21 {
                fmt.Println(v2, v5, v4)
                var v66 uint8 = 175
                v4 = v4
                _ = v66
            }
            v4 += 2
        } else {
            switch "hello" {
            case "hello":
                fmt.Println(v21)
            case "foo bar":
                v6 = append(v6, 2147483648)
            case "":
                v67 := GenS7{F0: 32767, F1: (("a" + "foo bar") + "a"), F2: (^(v5))}
                _ = v67
            }
            if len(v6) > 0 {
                v6[0] = 4294967295
            }
        }
    } else {
        v5 = v2
        for fi68 := 0; fi68 < 3; fi68++ {
            v2 = (^(v5))
            switch (fi68 + (^(fi68))) {
            case 32768:
                v5 = v5
            default:
                fmt.Println(v2, fi68)
            }
            v5 += 2
            _ = fi68
        }
    }
    var v69 int16 = 42
    var v70 uint32 = 65535
    v71 := GenS8{F0: v5, F1: uint(v5)}
    v4 += v4
    var v72 uint8 = 1
    delete(v3, "")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v70
    _ = v71
    _ = v72
}
