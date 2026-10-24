package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint; F2 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 string; F2 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 int64; F1 uint }

func main() {
    var v1 int8 = 27
    var v2 bool = false
    v3 := []uint32{}
    var v4 int16 = int16(v1)
    var v5 float64 = float64(0.1)
    v6 := make(map[string]bool)
    if v2 {
        v5 -= 0.0
        var v7 string = "hello"
        _ = v7
    }
    if true {
        if (false && v2) {
            v2 = v2
            if (0 == 4294967295) {
                _ = len(v3)
                v8 := GenS1{F0: 423, F1: 0}
                var v9 int64 = 420
                _ = v8
                _ = v9
            } else {
                var v10 int8 = (v1 ^ v1)
                _ = v10
            }
            v5 = v5
        } else {
            v11 := GenS1{F0: 42, F1: 14}
            _ = v11
        }
    }
    var v12 uint16 = 255
    var v13 float64 = 0.1
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 uint64 = 1
        _ = fi14
        _ = v15
    }
    var v16 int32 = 0
    _ = len(v3)
    switch 9223372036854775807 {
    case 2147483647:
        v17 := GenS1{F0: 215, F1: 1}
        _ = v17
    case 9223372036854775807:
        var v18 uint32 = 82
        _ = v18
    case -598:
        if v2 {
            if (!v2) {
                v19 := GenS2{F0: v16, F1: 255, F2: 65535}
                if false {
                    if v2 {
                        for fi20 := 0; fi20 < 9; fi20++ {
                            v5 = v5
                            v3 = append(v3, 471)
                            _ = fi20
                        }
                    } else {
                        for fi21 := 0; fi21 < 5; fi21++ {
                            var v22 uint = 65535
                            if len(v3) > 0 {
                                v3[0] = 855
                            }
                            _ = fi21
                            _ = v22
                        }
                        v13 = v5
                    }
                    var v23 float32 = (3.14159 - (float32(float64(3.14159)) / 1.5))
                    _ = v23
                } else {
                    delete(v6, "\n\t")
                }
                if (!v2) {
                    var v24 int16 = (v4 - v4)
                    _ = v24
                } else {
                    if len(v3) > 0 {
                        v3[0] = 690
                    }
                    var v25 int16 = (v4 % ((1) | 1))
                    _ = v25
                }
                _ = v19
            } else {
                if (0 < 128) {
                    switch -148 {
                    case -9223372036854775808:
                        if true {
                            for fi26 := 0; fi26 < 5; fi26++ {
                                if true {
                                    v27 := GenS3{F0: 2147483647}
                                    for fi28 := 0; fi28 < 18; fi28++ {
                                        v1 = v1
                                        _ = fi28
                                    }
                                    _ = v27
                                } else {
                                    var v29 uint8 = 232
                                    for fi30 := 0; fi30 < 8; fi30++ {
                                        v31 := GenS3{F0: 42}
                                        _ = fi30
                                        _ = v31
                                    }
                                    _ = v29
                                }
                                _ = len(v3)
                                v32 := v6[""]
                                _ = fi26
                                _ = v32
                            }
                        } else {
                            if (false || (((0.0 * 1.5) + 0.1) > 3.14159)) {
                                v33 := GenS2{F0: v16, F1: 122, F2: 65535}
                                _ = v33
                            }
                            v16 = v16
                            var v34 int32 = (v16 & v16)
                            _ = v34
                        }
                    case 4294967295:
                        var v35 float32 = 3.14159
                        _ = v35
                    }
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        switch "a" {
                        case "":
                            v36 := GenS3{F0: 4294967295}
                            _ = v36
                        case "foo bar":
                            if len(v3) > 0 {
                                v3[0] = 1
                            }
                        case "a":
                            for fi37 := 0; fi37 < 5; fi37++ {
                                if ((-(v1)) != (-(v1))) {
                                    if (v2 && false) {
                                        var v38 bool = (!(42 != v1))
                                        if v2 {
                                            switch "a" {
                                            case "a":
                                                v39 := GenS1{F0: 18446744073709551615, F1: 255}
                                                _ = v39
                                            default:
                                                v40 := GenS3{F0: 2147483647}
                                                _ = v40
                                            }
                                            var v41 float64 = float64(float32(((v5 / v5) / v5)))
                                            _ = v41
                                        } else {
                                            v16 += v16
                                        }
                                        fmt.Println(v4, v2)
                                        _ = v38
                                    }
                                } else {
                                    for fi42 := 0; fi42 < 12; fi42++ {
                                        switch (fi42 % (((fi42 - (fi37 + fi42))) | 1)) {
                                        case 42:
                                            var v43 int = 886
                                            fmt.Println(fi37, v12, v5)
                                            _ = v43
                                        case 127:
                                            var v44 int16 = 42
                                            _ = v44
                                        case -2147483648:
                                            v2 = ((!v2) || false)
                                            switch (fi37 | (^(fi37))) {
                                            case 1:
                                                if v2 {
                                                    var v45 string = ("hello" + "")
                                                    if ((0.1 - (((0.1 * (0.1 + 0.1)) + ((1.5 * 0.1) + 3.14159)) / 3.14159)) <= ((1.5 * 0.1) - ((3.14159 / 3.14159) / 3.14159))) {
                                                        v46 := v6[(v45 + ((v45 + v45) + ("foo bar" + "hello")))]
                                                        v47 := GenS4{F0: 423, F1: "foo bar", F2: (v12 & v12)}
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    if v2 {
                                                        var v48 int8 = (^(v1))
                                                        _ = v48
                                                    } else {
                                                        v3 = append(v3, 778)
                                                        v49 := GenS4{F0: 42, F1: v45, F2: (v12 >> ((65535) & 63))}
                                                        if v2 {
                                                            if len(v3) > 0 {
                                                                v3[0] = 65535
                                                            }
                                                            v50 := GenS5{F0: (v4 & (-(v4)))}
                                                            if v2 {
                                                                var v51 uint = 804
                                                                _ = v51
                                                            }
                                                            v16 = (-(v16))
                                                            _ = v50
                                                        } else {
                                                            switch fi37 {
                                                            case 0:
                                                                if true {
                                                                    if v2 {
                                                                        var v52 float32 = 0.0
                                                                        var v53 uint64 = 4294967295
                                                                        _ = v52
                                                                        _ = v53
                                                                    }
                                                                    var v54 uint16 = 255
                                                                    _ = v54
                                                                } else {
                                                                    v55 := GenS4{F0: 2, F1: v45, F2: v12}
                                                                    _ = v55
                                                                }
                                                            }
                                                            var v56 uint16 = v12
                                                            v1 = (v1 * v1)
                                                            _ = v56
                                                        }
                                                        v13 = v5
                                                        _ = v49
                                                    }
                                                    v4 -= v4
                                                    _ = v45
                                                }
                                            case -197:
                                                var v57 float64 = v5
                                                _ = v57
                                            case 339:
                                                v12 -= (v12 ^ (v12 << ((1) & 63)))
                                            }
                                        }
                                        _ = fi42
                                    }
                                }
                                _ = fi37
                            }
                        default:
                            v58 := GenS6{F0: -9223372036854775808, F1: 42}
                            _ = v58
                        }
                    }
                    v13 -= (v5 / (float64(3.14159) - -2.25))
                    var v59 int8 = int8(v4)
                    _ = v59
                }
                var v60 string = "foo bar"
                _ = v60
            }
            v61 := GenS2{F0: (v16 + -2147483648), F1: 2, F2: 0}
            _ = v61
        } else {
            switch "" {
            case "foo bar":
                v62 := GenS7{F0: 9223372036854775807, F1: 754}
                _ = v62
            }
        }
    }
    v12 = (v12 & 490)
    v16 = (v16 ^ -2147483648)
    v2 = v2
    if (v12 != 248) {
        var v63 int = 566
        _ = v63
    }
    fmt.Println(v4)
    var v64 bool = v2
    v12 = 550
    v4 += (v4 << ((205) & 63))
    for fi65 := 0; fi65 < 12; fi65++ {
        var v66 uint8 = 42
        var v67 float32 = 0.1
        var v68 int32 = int32(v12)
        _ = fi65
        _ = v66
        _ = v67
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v64:", v64)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
