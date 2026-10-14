package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8; F1 string }
type GenS9 struct { F0 uint64; F1 uint8 }

func main() {
    v1 := []uint32{}
    var v2 int16 = 32767
    v3 := []int{360, 32767}
    var v4 bool = false
    var v5 uint = 2
    if len(v1) > 0 {
        v1[0] = 2
    }
    if len(v1) > 0 {
        v1[0] = 2
    }
    if (2147483647 >= 128) {
        if v4 {
            v6 := GenS1{F0: 224}
            v7 := GenS2{F0: 127, F1: 128}
            if (32767 != (v2 & -432)) {
                if v4 {
                    v2 += (v2 - 0)
                    for fi8 := 0; fi8 < 8; fi8++ {
                        if ((fi8 / (((fi8 << (((v5 - v5)) & 63))) | 1)) >= (fi8 | (^(fi8)))) {
                            fmt.Println(fi8)
                        } else {
                            v4 = (0.0 > (((0.0 * 3.14159) - 0.1) + 3.14159))
                            v9 := GenS2{F0: 32768, F1: -324}
                            var v10 uint8 = 42
                            _ = v10
                            _ = v9
                        }
                        if v4 {
                            v5 = (v5 & 916)
                            _ = len(v1)
                            if true {
                                var v11 float32 = float32(((0.1 + 1.5) + 1e10))
                                _ = v11
                            }
                        } else {
                            if false {
                                if v4 {
                                    v4 = v4
                                    v1 = append(v1, 42)
                                    v5 = 919
                                }
                                v4 = (!v4)
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = (-(fi8))
                                }
                                v12 := GenS3{F0: -801, F1: v4}
                                _ = v12
                            }
                        }
                        _ = fi8
                    }
                    if v4 {
                        var v13 float32 = 0.1
                        _ = v13
                    }
                } else {
                    fmt.Println(v4)
                    for fi14 := 0; fi14 < 1; fi14++ {
                        v4 = (!v4)
                        for fi15 := 0; fi15 < 1; fi15++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if false {
                                var v16 uint32 = 1
                                _ = v16
                            }
                            v1 = append(v1, 0)
                            _ = fi15
                        }
                        var v17 float32 = (0.1 / 0.1)
                        _ = fi14
                        _ = v17
                    }
                }
            }
            _ = v6
            _ = v7
        } else {
            for fi18 := 0; fi18 < 3; fi18++ {
                v5 = (v5 << ((v5) & 63))
                var v19 int32 = 234
                var v20 int = (fi18 | (^(fi18)))
                _ = fi18
                _ = v19
                _ = v20
            }
            if ((-2.25 == 0.0) && true) {
                switch -1 {
                case -641:
                    var v21 int32 = 42
                    _ = v21
                }
            } else {
                for fi22 := 0; fi22 < 8; fi22++ {
                    var v23 uint16 = 42
                    _ = fi22
                    _ = v23
                }
                for fi24 := 0; fi24 < 2; fi24++ {
                    var v25 int8 = 127
                    if (!(255 < 255)) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v26 int32 = 2147483647
                        _ = v26
                    } else {
                        v5 = v5
                        fmt.Println(v2, v2)
                    }
                    v27 := GenS3{F0: v2, F1: v4}
                    _ = fi24
                    _ = v25
                    _ = v27
                }
                if len(v1) > 0 {
                    v1[0] = 2
                }
            }
            var v28 uint64 = 179
            _ = v28
        }
    } else {
        if (!(v2 > 32767)) {
            v29 := GenS4{F0: -1, F1: 136}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v30 := GenS5{F0: ((((-2.25 / 3.14159) * 1e10) <= (3.14159 / (1e10 - (1.5 / 0.1)))) || v4)}
            _ = v29
            _ = v30
        }
    }
    v31 := GenS1{F0: 207}
    v2 = v2
    switch 42 {
    case -9223372036854775808:
        for fi32 := 0; fi32 < 15; fi32++ {
            v33 := GenS6{F0: "hello", F1: 18446744073709551615}
            _ = fi32
            _ = v33
        }
    default:
        if v4 {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if true {
                v3 = append(v3, 32768)
                v4 = v4
                var v34 float32 = 0.0
                _ = v34
            }
            fmt.Println(v5, v5, v2)
        }
    }
    switch 42 {
    case 32768:
        fmt.Println(v2, v4, v5)
    }
    fmt.Println(v2, v5)
    for fi35 := 0; fi35 < 17; fi35++ {
        for fi36 := 0; fi36 < 13; fi36++ {
            var v37 int16 = v2
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    for fi38 := 0; fi38 < 11; fi38++ {
        v4 = v4
        v39 := GenS5{F0: (((float64(0.0) / float64((0.1 + 1e10))) * float64(1e10)) < (1.5 / 1.0))}
        _ = fi38
        _ = v39
    }
    if true {
        for fi40 := 0; fi40 < 13; fi40++ {
            for fi41 := 0; fi41 < 5; fi41++ {
                fmt.Println(v5)
                v5 += 65535
                var v42 float32 = 0.1
                _ = fi41
                _ = v42
            }
            switch "foo bar" {
            case "foo bar":
                fmt.Println(fi40)
            case "hello":
                if (42 > 42) {
                    var v43 float32 = (1e10 * -2.25)
                    switch ("foo bar" + (("" + "foo bar") + "")) {
                    case "\n\t":
                        if len(v1) > 0 {
                            v1[0] = 291
                        }
                    }
                    fmt.Println(v2, fi40)
                    _ = v43
                }
            case "a":
                v4 = (-1 < 42)
            }
            if true {
                _ = len(v1)
                v2 = (v2 / ((v2) | 1))
            } else {
                v4 = (v4 && false)
                v44 := GenS7{F0: 65535}
                _ = v44
            }
            _ = fi40
        }
    }
    if false {
        v2 = v2
        switch -223 {
        case -787:
            if (2 > 127) {
                fmt.Println(v5, v5, v5)
            }
        case 4294967295:
            switch ("" + "") {
            case "a":
                if (!v4) {
                    v45 := GenS5{F0: (127 != -844)}
                    var v46 uint16 = 255
                    switch 0 {
                    case 4294967295:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 0:
                        var v47 int8 = 0
                        _ = v47
                    default:
                        switch 108 {
                        case 9223372036854775807:
                            if v4 {
                                var v48 uint = (v5 >> ((v5) & 63))
                                _ = v48
                            } else {
                                v49 := GenS6{F0: (("" + ("\n\t" + "")) + "hello"), F1: v5}
                                var v50 uint16 = v46
                                _ = v49
                                _ = v50
                            }
                        case 128:
                            if v4 {
                                v2 = (v2 >> ((v5) & 63))
                            } else {
                                v4 = true
                                v5 -= v5
                            }
                        case 127:
                            var v51 float64 = 0.1
                            _ = v51
                        default:
                            v52 := GenS7{F0: (v46 << ((v5) & 63))}
                            _ = v52
                        }
                    }
                    _ = v45
                    _ = v46
                }
            case "":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "foo bar":
                fmt.Println(v4)
            }
        }
        if (v5 >= v5) {
            v2 -= v2
            if (false && true) {
                v5 = v5
                v53 := GenS8{F0: 127, F1: "\n\t"}
                _ = v53
            }
            v5 -= v5
        }
    } else {
        v2 += -269
        v54 := GenS9{F0: 4294967295, F1: 255}
        _ = v54
    }
    for fi55 := 0; fi55 < 8; fi55++ {
        v56 := GenS2{F0: fi55, F1: 2147483647}
        _ = fi55
        _ = v56
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
