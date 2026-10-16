package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 bool; F1 int32; F2 int32 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string; F1 int; F2 uint64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 string; F1 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = -938
    var v3 int64 = 0
    var v4 uint64 = 209
    v5 := make(map[uint32]int)
    var v6 uint = 18446744073709551615
    switch "foo bar" {
    case "a":
        var v7 uint8 = 127
        _ = v7
    default:
        v6 = v6
    }
    v8 := GenS1{F0: "a", F1: 1}
    v9 := []int{4294967295, 127, -1}
    var v10 int8 = 1
    for fi11 := 0; fi11 < 9; fi11++ {
        v10 = int8(v6)
        _ = fi11
    }
    delete(v5, 42)
    if true {
        v2 -= (v2 & (v3 + v3))
        v4 += (v4 * (v4 | (^(v4))))
        if false {
            if true {
                fmt.Println(v3)
                var v12 uint = (v6 / (((v6 >> ((v6) & 63))) | 1))
                switch int(v12) {
                case -187:
                    v13 := GenS2{F0: true, F1: int32(v6), F2: -1}
                    _ = v13
                case 997:
                    v12 = (v12 & v6)
                case -700:
                    var v14 int64 = v2
                    _ = v14
                }
                _ = v12
            }
            v9 = append(v9, 118)
            if (0.0 < (0.1 + 3.14159)) {
                fmt.Println(v10, v6)
                v15 := GenS1{F0: "a", F1: v6}
                for fi16 := 0; fi16 < 8; fi16++ {
                    if false {
                        if (v4 >= (v4 >> ((v6) & 63))) {
                            var v17 int8 = 127
                            if false {
                                for fi18 := 0; fi18 < 4; fi18++ {
                                    if true {
                                        switch fi18 {
                                        case 42:
                                            v2 = v3
                                        default:
                                            v19 := GenS3{F0: 944, F1: 32767}
                                            v10 = (v10 + v17)
                                            var v20 string = "a"
                                            _ = v19
                                            _ = v20
                                        }
                                    } else {
                                        v10 += (v17 % ((-73) | 1))
                                    }
                                    v2 = v3
                                    switch ("foo bar" + "") {
                                    case "\n\t":
                                        if false {
                                            v17 -= v17
                                        }
                                    case "":
                                        if (0 != -283) {
                                            if (v4 != v4) {
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                                var v21 uint64 = 4294967295
                                                _ = v21
                                            } else {
                                                if (false && (!true)) {
                                                    delete(v5, 364)
                                                } else {
                                                    v10 -= v10
                                                    v4 = 0
                                                }
                                            }
                                            if ((^(fi18)) > fi18) {
                                                var v22 int64 = (v3 - -9223372036854775808)
                                                if (18446744073709551615 != (v6 - (v6 ^ v6))) {
                                                    v23 := v1[int(v10)]
                                                    if (!(v23 || (true || v23))) {
                                                        fmt.Println(fi16, fi16, v22)
                                                        var v24 float64 = 0.0
                                                        _ = v24
                                                    }
                                                    v25 := GenS1{F0: "foo bar", F1: 0}
                                                    v9 = append(v9, (fi18 >> ((v6) & 63)))
                                                    _ = v23
                                                    _ = v25
                                                } else {
                                                    var v26 int16 = -1
                                                    v22 += v2
                                                    _ = v26
                                                }
                                                _ = v22
                                            } else {
                                                if false {
                                                    fmt.Println(v10, v2, fi18)
                                                    v27 := GenS2{F0: false, F1: 969, F2: 42}
                                                    if (0.1 == (-2.25 / 1.0)) {
                                                        var v28 float64 = -2.25
                                                        _ = v28
                                                    } else {
                                                        switch "a" {
                                                        case "a":
                                                            v2 = 42
                                                        case "":
                                                            var v29 int64 = int64(v4)
                                                            _ = v29
                                                        case "\n\t":
                                                            v30 := GenS4{F0: false}
                                                            var v31 string = "hello"
                                                            delete(v1, fi18)
                                                            _ = v30
                                                            _ = v31
                                                        default:
                                                            var v32 uint16 = 32767
                                                            _ = v32
                                                        }
                                                        fmt.Println(v3)
                                                        v4 += (v4 & (^(v4)))
                                                        delete(v5, 42)
                                                        var v33 string = ("a" + ("hello" + ("" + "\n\t")))
                                                        v6 = (v6 % ((v6) | 1))
                                                        _ = v33
                                                    }
                                                    if true {
                                                        v9 = append(v9, 128)
                                                        v10 -= (^(v17))
                                                    }
                                                    if len(v9) > 0 {
                                                        v9[0] = fi16
                                                    }
                                                    _ = v27
                                                }
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                                var v34 uint8 = 127
                                                _ = v34
                                            }
                                            fmt.Println(v17)
                                        }
                                        if (186 >= 107) {
                                            switch fi16 {
                                            case 4294967295:
                                                var v35 uint16 = 660
                                                _ = v35
                                            case 1:
                                                if ((112 == 1) && (207 == 164)) {
                                                    if true {
                                                        var v36 int16 = 1000
                                                        _ = v36
                                                    }
                                                    if ("a" <= ("foo bar" + "foo bar")) {
                                                        if (v3 != 9223372036854775807) {
                                                            v17 -= 127
                                                            v17 = v10
                                                        } else {
                                                            v3 = v2
                                                            var v37 int16 = 0
                                                            _ = v37
                                                        }
                                                    } else {
                                                        v38 := GenS5{F0: (v6 << ((v6) & 63))}
                                                        v39 := GenS1{F0: "a", F1: v6}
                                                        _ = v38
                                                        _ = v39
                                                    }
                                                } else {
                                                    v40 := GenS3{F0: 2, F1: 364}
                                                    _ = v40
                                                }
                                            case -9223372036854775808:
                                                v41 := GenS6{F0: (1.5 < ((0.1 - -2.25) / (1e10 / 3.14159))), F1: 1}
                                                _ = v41
                                            default:
                                                _ = len(v9)
                                            }
                                            v6 = 0
                                        } else {
                                            var v42 int = (fi16 / ((fi18) | 1))
                                            _ = v42
                                        }
                                        if (-750 < int32(v4)) {
                                            var v43 bool = (!false)
                                            _ = v43
                                        }
                                    }
                                    v44 := GenS7{F0: 18446744073709551615}
                                    var v45 string = "foo bar"
                                    _ = fi18
                                    _ = v44
                                    _ = v45
                                }
                                v4 = 4294967295
                                switch "hello" {
                                case "":
                                    if false {
                                        v46 := GenS8{F0: "", F1: 943, F2: (v4 << ((v6) & 63))}
                                        switch (fi16 + fi16) {
                                        case -2147483648:
                                            for fi47 := 0; fi47 < 17; fi47++ {
                                                if false {
                                                    var v48 float32 = 1.5
                                                    v49 := GenS1{F0: "", F1: (^(v6))}
                                                    var v50 string = ("\n\t" + ("hello" + "\n\t"))
                                                    _ = v48
                                                    _ = v49
                                                    _ = v50
                                                } else {
                                                    var v51 int8 = -1
                                                    _ = v51
                                                }
                                                switch ("" + "hello") {
                                                case "a":
                                                    v52 := GenS7{F0: 1}
                                                    _ = v52
                                                default:
                                                    if (true || true) {
                                                        v2 = v2
                                                        v10 = v10
                                                    }
                                                }
                                                var v53 float64 = 0.1
                                                _ = fi47
                                                _ = v53
                                            }
                                        case -672:
                                            switch "\n\t" {
                                            case "hello":
                                                for fi54 := 0; fi54 < 19; fi54++ {
                                                    _ = len(v9)
                                                    _ = fi54
                                                }
                                            case "a":
                                                _ = len(v9)
                                            }
                                        default:
                                            for fi55 := 0; fi55 < 12; fi55++ {
                                                v3 += v3
                                                var v56 uint16 = uint16(v4)
                                                var v57 int = fi16
                                                _ = fi55
                                                _ = v56
                                                _ = v57
                                            }
                                        }
                                        if ((!false) && false) {
                                            v10 = v10
                                            for fi58 := 0; fi58 < 3; fi58++ {
                                                var v59 uint = v6
                                                v60 := GenS3{F0: 1, F1: 830}
                                                if (v3 == 9223372036854775807) {
                                                    switch 32768 {
                                                    case 1:
                                                        var v61 uint64 = (v4 >> (((v59 * v6)) & 63))
                                                        _ = v61
                                                    case 4294967295:
                                                        fmt.Println(v2)
                                                    default:
                                                        v62 := GenS6{F0: ((((fi16 & fi16) > (fi58 | fi58)) || (!false)) || true), F1: 973}
                                                        _ = v62
                                                    }
                                                }
                                                _ = fi58
                                                _ = v59
                                                _ = v60
                                            }
                                        } else {
                                            switch int(v3) {
                                            case -1:
                                                if false {
                                                    var v63 int8 = (v10 * (v10 | 0))
                                                    _ = v63
                                                }
                                            }
                                        }
                                        _ = v46
                                    }
                                case "hello":
                                    delete(v1, -1)
                                }
                            } else {
                                if false {
                                    var v64 uint64 = v4
                                    var v65 uint = 0
                                    var v66 int = -2147483648
                                    _ = v64
                                    _ = v65
                                    _ = v66
                                }
                            }
                            for fi67 := 0; fi67 < 9; fi67++ {
                                v68 := GenS3{F0: 42, F1: 65535}
                                var v69 uint = v6
                                if ("\n\t" < "a") {
                                    v70 := GenS1{F0: "a", F1: 18446744073709551615}
                                    v71 := v1[(fi16 | fi67)]
                                    var v72 float32 = 1.5
                                    v5[2147483648] = 2147483647
                                    _ = v70
                                    _ = v71
                                    _ = v72
                                } else {
                                    if (-911 != 765) {
                                        var v73 bool = false
                                        _ = v73
                                    }
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    if true {
                                        if true {
                                            fmt.Println(v4, v69, v2)
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                        } else {
                                            var v74 int64 = v2
                                            var v75 float32 = ((0.1 * (0.1 / 1.0)) - 0.0)
                                            var v76 int8 = v10
                                            _ = v74
                                            _ = v75
                                            _ = v76
                                        }
                                        v77 := v5[2147483648]
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                        _ = v77
                                    }
                                }
                                _ = fi67
                                _ = v68
                                _ = v69
                            }
                            _ = v17
                        }
                    }
                    for fi78 := 0; fi78 < 15; fi78++ {
                        v79 := GenS9{F0: 32767}
                        _ = fi78
                        _ = v79
                    }
                    _ = fi16
                }
                _ = v15
            }
        }
    }
    v80 := GenS10{F0: "hello", F1: v2}
    var v81 int8 = -6
    v3 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v81:", v81)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v80
    _ = v81
    _ = v9
}
