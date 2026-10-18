package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 int8; F2 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint }
func genFn0(p0 uint16, p1 int64) uint8 {
    switch "\n\t" {
    case "hello":
        p1 += p1
    case "\n\t":
        v6 := []int{506, 42, 32768}
        _ = v6
    case "a":
        v7 := []int{32767, -9223372036854775808, int(p1), 0}
        _ = v7
    default:
        p0 = p0
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int8 = 100
        p0 -= (p0 >> ((1) & 63))
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int = (fi10 + 127)
            v12 := GenS1{F0: (("a" + "foo bar") + "a")}
            v11 = (^(fi8))
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 int8 = 0
    var v2 int64 = 0
    var v3 bool = false
    v4 := []uint32{0}
    v5 := []int{-960, 243}
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int64 = (-(v2))
        for fi15 := 0; fi15 < 14; fi15++ {
            v3 = v3
            _ = fi15
        }
        var v16 int8 = -128
        _ = fi13
        _ = v14
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        v2 += (v2 << ((255) & 63))
        v3 = v3
        _ = fi17
    }
    if false {
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS1{F0: "foo bar"}
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 uint16 = 2
                v1 = (v1 << ((42) & 63))
                _ = fi20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
    }
    v2 = v2
    v1 -= v1
    var v22 int16 = 32767
    var v23 uint64 = 4294967295
    for fi24 := 0; fi24 < 2; fi24++ {
        fmt.Println(v1)
        v22 -= v22
        _ = fi24
    }
    if (42 != 1) {
        for fi25 := 0; fi25 < 7; fi25++ {
            v1 = (v1 / ((v1) | 1))
            if (2 != 958) {
                fmt.Println(v22, v1)
                v26 := GenS2{F0: 357}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v26
            } else {
                switch (fi25 % ((-526) | 1)) {
                case 128:
                    v4 = append(v4, 4294967295)
                case -442:
                    var v27 uint32 = 4294967295
                    _ = v27
                default:
                    for fi28 := 0; fi28 < 5; fi28++ {
                        v1 = v1
                        v5 = append(v5, -689)
                        switch (("a" + "foo bar") + "") {
                        case "\n\t":
                            var v29 uint32 = 42
                            _ = v29
                        case "hello":
                            if (2147483648 < 2147483647) {
                                v1 -= (v1 % ((-51) | 1))
                                _ = len(v4)
                                var v30 int64 = v2
                                _ = v30
                            } else {
                                v23 = 0
                                v22 = -265
                                _ = len(v5)
                            }
                        case "foo bar":
                            if (("\n\t" + (("" + "") + "\n\t")) < ("foo bar" + "hello")) {
                                if (42 == 628) {
                                    v31 := GenS3{F0: (v22 & 0), F1: (2 == 255)}
                                    if false {
                                        _ = len(v5)
                                    } else {
                                        var v32 int16 = (v22 | (v22 >> ((255) & 63)))
                                        v33 := GenS4{F0: (1 == 127)}
                                        _ = v32
                                        _ = v33
                                    }
                                    switch "\n\t" {
                                    case "":
                                        var v34 float64 = float64(1.5)
                                        _ = v34
                                    }
                                    fmt.Println(v1)
                                    var v35 uint16 = 65535
                                    _ = v31
                                    _ = v35
                                } else {
                                    switch fi28 {
                                    case 4294967295:
                                        v3 = v3
                                    case 445:
                                        var v36 bool = (42 >= 247)
                                        _ = v36
                                    case -430:
                                        var v37 int = fi28
                                        _ = v37
                                    default:
                                        if false {
                                            v1 += v1
                                            if (!v3) {
                                                if v3 {
                                                    fmt.Println(v23)
                                                    v1 += (v1 + v1)
                                                    if v3 {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        if true {
                                                            v23 -= (v23 & (v23 >> ((65535) & 63)))
                                                            v3 = (32767 < 42)
                                                            v2 = v2
                                                        } else {
                                                            var v38 bool = ((v22 - (v22 / ((-332) | 1))) != -32768)
                                                            v3 = (!v38)
                                                            _ = v38
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v2)
                                                }
                                                switch "a" {
                                                case "a":
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v39 := GenS5{F0: 209, F1: v1, F2: (v22 & 42)}
                                                        _ = v39
                                                    default:
                                                        if ((!(v3 && true)) || (!v3)) {
                                                            if v3 {
                                                                var v40 float32 = -2.25
                                                                _ = v40
                                                            }
                                                            var v41 uint16 = 888
                                                            var v42 float32 = 0.0
                                                            v3 = (1 < (v1 + -54))
                                                            _ = v41
                                                            _ = v42
                                                        } else {
                                                            var v43 bool = (v3 || (fi25 > (fi28 & (fi28 | 2147483647))))
                                                            _ = v43
                                                        }
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v1, v2)
                                                switch fi25 {
                                                case 604:
                                                    v44 := GenS3{F0: (v22 << ((908) & 63)), F1: false}
                                                    _ = v44
                                                default:
                                                    switch (-(fi25)) {
                                                    case 242:
                                                        if v3 {
                                                            v45 := GenS5{F0: 42, F1: (v1 + -128), F2: v22}
                                                            _ = v45
                                                        } else {
                                                            v46 := GenS3{F0: -431, F1: (v22 <= v22)}
                                                            var v47 uint32 = 2147483648
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                    default:
                                                        var v48 uint16 = 0
                                                        _ = v48
                                                    }
                                                }
                                                v2 = v2
                                            }
                                            fmt.Println(v2, v22)
                                        }
                                    }
                                    v23 += v23
                                    v4 = append(v4, 2)
                                }
                                v2 = 1
                                fmt.Println(v2, fi28, v22)
                            }
                        }
                        _ = fi28
                    }
                }
                if len(v5) > 0 {
                    v5[0] = fi25
                }
                v49 := GenS6{F0: 42}
                _ = v49
            }
            _ = fi25
        }
        switch "a" {
        case "\n\t":
            for fi50 := 0; fi50 < 18; fi50++ {
                v3 = (("hello" + "a") > "foo bar")
                for fi51 := 0; fi51 < 16; fi51++ {
                    switch fi50 {
                    case 2147483647:
                        var v52 uint = 0
                        _ = v52
                    case 9223372036854775807:
                        v53 := GenS7{F0: 2}
                        _ = v53
                    case -796:
                        switch "a" {
                        case "foo bar":
                            var v54 int32 = int32(v23)
                            _ = v54
                        }
                    default:
                        if v3 {
                            v23 = (v23 << ((13) & 63))
                            fmt.Println(v22, v23)
                            var v55 uint16 = 2
                            switch -321 {
                            case 84:
                                v55 -= v55
                            case 415:
                                v56 := GenS6{F0: 2147483647}
                                _ = v56
                            case -9223372036854775808:
                                if len(v4) > 0 {
                                    v4[0] = 2147483647
                                }
                            default:
                                var v57 float32 = 1.5
                                _ = v57
                            }
                            v23 = v23
                            _ = v55
                        } else {
                            var v58 uint32 = 2147483647
                            var v59 float64 = (float64((((3.14159 / (1.5 - 3.14159)) / (-2.25 * 0.1)) - 1e10)) - 0.0)
                            _ = len(v5)
                            v4 = append(v4, v58)
                            _ = v58
                            _ = v59
                        }
                    }
                    _ = fi51
                }
                if len(v5) > 0 {
                    v5[0] = fi50
                }
                _ = fi50
            }
        case "hello":
            v60 := GenS3{F0: v22, F1: (!(1 != 2147483648))}
            _ = v60
        default:
            fmt.Println(v2)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
