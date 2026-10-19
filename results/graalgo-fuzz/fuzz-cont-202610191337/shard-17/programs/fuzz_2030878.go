package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 uint; F2 int8 }
type GenS4 struct { F0 int8; F1 int; F2 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 int }

func main() {
    var v1 int16 = -32768
    v2 := []int{114, -1, 1, 163}
    var v3 int32 = 1
    var v4 uint32 = 249
    v5 := []string{(("foo bar" + "") + ("" + "hello")), ("a" + "\n\t")}
    var v6 bool = (702 > 65535)
    fmt.Println(v6, v1, v3)
    v4 = (v4 | v4)
    v2 = append(v2, 32767)
    v3 += 42
    for fi7 := 0; fi7 < 18; fi7++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi7
    }
    switch 2147483647 {
    case -9223372036854775808:
        fmt.Println(v1, v3, v4)
    case 32767:
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int32 = v3
            v1 = (v1 ^ (v1 & v1))
            if (-1 == 9223372036854775807) {
                v1 -= (v1 % ((v1) | 1))
                if v6 {
                    if (42 <= 9223372036854775807) {
                        for fi10 := 0; fi10 < 3; fi10++ {
                            var v11 uint = 65535
                            _ = fi10
                            _ = v11
                        }
                        v3 += v9
                        fmt.Println(fi8, v9, v1)
                    } else {
                        if (156 <= 145) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        } else {
                            if len(v5) > 0 {
                                v5[0] = "a"
                            }
                            fmt.Println(fi8)
                            var v12 int16 = (v1 >> ((255) & 63))
                            _ = v12
                        }
                    }
                    v9 = (v3 ^ (v3 >> ((452) & 63)))
                }
            } else {
                switch (fi8 & fi8) {
                case 1:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case -470:
                    switch fi8 {
                    case 712:
                        v13 := GenS1{F0: 7}
                        _ = v13
                    }
                }
            }
            _ = fi8
            _ = v9
        }
    case 512:
        v6 = v6
    default:
        switch "\n\t" {
        case "foo bar":
            v1 = 32767
        case "a":
            v1 -= (v1 * v1)
        case "hello":
            if (1.5 != 1e10) {
                for fi14 := 0; fi14 < 15; fi14++ {
                    v15 := GenS1{F0: v4}
                    var v16 float32 = (1e10 / 1.5)
                    var v17 uint = 2
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
                fmt.Println(v6, v1, v1)
            }
        }
    }
    v18 := GenS1{F0: (v4 * (^(v4)))}
    v1 = v1
    fmt.Println(v3, v1, v3)
    for fi19 := 0; fi19 < 3; fi19++ {
        for fi20 := 0; fi20 < 12; fi20++ {
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
            _ = fi20
        }
        _ = fi19
    }
    v6 = (v3 <= (-(v3)))
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 5; fi22++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v23 uint = 4294967295
            v5 = append(v5, ("hello" + "\n\t"))
            _ = fi22
            _ = v23
        }
        _ = len(v5)
        _ = fi21
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v25 := GenS2{F0: v4}
        for fi26 := 0; fi26 < 16; fi26++ {
            if true {
                var v27 bool = ((v4 + 1) >= v4)
                _ = v27
            }
            _ = fi26
        }
        if (-1 >= 46) {
            fmt.Println(v1)
            for fi28 := 0; fi28 < 6; fi28++ {
                switch (fi24 ^ (-(fi24))) {
                case -1:
                    switch "\n\t" {
                    case "":
                        v29 := GenS1{F0: (v4 ^ 485)}
                        _ = v29
                    case "hello":
                        v30 := GenS3{F0: 42, F1: 255, F2: 42}
                        _ = v30
                    case "a":
                        v31 := GenS2{F0: (v4 << ((65535) & 63))}
                        _ = v31
                    }
                case -928:
                    fmt.Println(v1)
                case 32767:
                    if (v6 && ((fi28 << ((65535) & 63)) != fi28)) {
                        switch "\n\t" {
                        case "a":
                            fmt.Println(v1, v4, v4)
                        }
                    } else {
                        v2 = append(v2, fi24)
                    }
                default:
                    var v32 int64 = 241
                    _ = v32
                }
                fmt.Println(v4, fi24)
                _ = fi28
            }
            var v33 uint16 = 495
            _ = v33
        } else {
            switch fi24 {
            case 4294967295:
                switch ("foo bar" + ("a" + "hello")) {
                case "\n\t":
                    var v34 uint16 = 627
                    _ = v34
                case "foo bar":
                    for fi35 := 0; fi35 < 4; fi35++ {
                        v6 = (771 == 18446744073709551615)
                        fmt.Println(v3, fi24, v3)
                        var v36 uint8 = 59
                        _ = fi35
                        _ = v36
                    }
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            case 949:
                v37 := GenS4{F0: 1, F1: fi24, F2: (v4 >> ((255) & 63))}
                _ = v37
            case -9223372036854775808:
                var v38 int32 = 184
                _ = v38
            default:
                if v6 {
                    switch "" {
                    case "a":
                        if (("a" + "\n\t") < "hello") {
                            if true {
                                if (float32(float64((-2.25 / 1.5))) < 1.5) {
                                    for fi39 := 0; fi39 < 18; fi39++ {
                                        if len(v2) > 0 {
                                            v2[0] = fi24
                                        }
                                        _ = fi39
                                    }
                                }
                            } else {
                                var v40 uint8 = 1
                                switch (fi24 / ((-469) | 1)) {
                                case -2147483648:
                                    var v41 int = fi24
                                    _ = v41
                                case 427:
                                    v42 := GenS5{F0: 116}
                                    _ = v42
                                default:
                                    v40 = v40
                                }
                                _ = v40
                            }
                            v6 = v6
                            var v43 uint8 = 2
                            _ = v43
                        } else {
                            v5 = append(v5, "foo bar")
                            v4 -= (v4 / (((^(v4))) | 1))
                            v44 := GenS2{F0: (v4 & (v4 << ((320) & 63)))}
                            _ = v44
                        }
                    case "":
                        switch "foo bar" {
                        case "a":
                            for fi45 := 0; fi45 < 11; fi45++ {
                                var v46 uint32 = (^(v4))
                                switch "hello" {
                                case "foo bar":
                                    if false {
                                        if len(v5) > 0 {
                                            v5[0] = ""
                                        }
                                        v47 := GenS4{F0: int8(v3), F1: fi45, F2: (v46 % (((^(v4))) | 1))}
                                        v48 := GenS3{F0: 0, F1: 18446744073709551615, F2: 40}
                                        _ = v47
                                        _ = v48
                                    }
                                }
                                v1 -= v1
                                _ = fi45
                                _ = v46
                            }
                        case "foo bar":
                            var v49 int = fi24
                            _ = v49
                        }
                    default:
                        var v50 int8 = -79
                        _ = v50
                    }
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if (551 == 2) {
                        v4 -= (v4 >> ((255) & 63))
                        v6 = (false && v6)
                    }
                    v4 = 4294967295
                }
            }
            v51 := GenS6{F0: (v4 / (((^(v4))) | 1)), F1: fi24}
            var v52 uint = 2
            _ = v51
            _ = v52
        }
        _ = fi24
        _ = v25
    }
    var v53 int = 127
    var v54 uint16 = 441
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
