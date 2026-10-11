package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint; F2 int }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 int16; F2 uint8 }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 int; F1 uint32; F2 uint }
type GenS8 struct { F0 uint; F1 bool; F2 string }
type GenS9 struct { F0 int; F1 uint; F2 int16 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int64; F1 int16; F2 int32 }

func main() {
    var v1 float64 = (3.14159 - 0.0)
    v2 := make(map[string]bool)
    var v3 uint64 = 1
    v4 := []int{9223372036854775807, 0}
    v5 := GenS1{F0: 1}
    var v6 uint32 = 2147483648
    switch "a" {
    case "":
        fmt.Println(v3, v6, v3)
    case "\n\t":
        v7 := GenS1{F0: -1}
        _ = v7
    default:
        if (2147483647 == 42) {
            v1 = v1
            if (32767 <= 65535) {
                v1 = v1
                var v8 uint16 = uint16(v6)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v8
            } else {
                if len(v4) > 0 {
                    v4[0] = 574
                }
                var v9 uint = 4294967295
                _ = v9
            }
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 uint8 = 0
                _ = fi10
                _ = v11
            }
        }
    }
    v3 -= (v3 * (^(v3)))
    if true {
        if (4294967295 >= 90) {
            v12 := GenS1{F0: 124}
            for fi13 := 0; fi13 < 10; fi13++ {
                var v14 string = ""
                var v15 bool = false
                _ = fi13
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 6; fi16++ {
                v3 = (v3 * v3)
                _ = fi16
            }
            _ = v12
        } else {
            v1 = ((float64(float32(((v1 - float64(0.1)) / float64(1.5)))) / v1) + float64(0.1))
            var v17 uint32 = 441
            _ = v17
        }
        v18 := GenS2{F0: 301, F1: 42, F2: 42}
        _ = v18
    } else {
        var v19 uint64 = v3
        _ = v19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 float64 = (v1 * v1)
        for fi22 := 0; fi22 < 15; fi22++ {
            var v23 bool = ((!true) || (151 >= 42))
            fmt.Println(v6)
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    v24 := GenS3{F0: 1000, F1: 317}
    _ = len(v4)
    delete(v2, (("" + "foo bar") + "foo bar"))
    v1 = (v1 - (((v1 + 1e10) + v1) + (float64(0.1) - ((v1 - ((float64(0.0) * ((v1 / 1e10) - float64(3.14159))) / -2.25)) - (0.1 - -2.25)))))
    v25 := v2[""]
    v6 -= v6
    switch 4294967295 {
    case 2147483647:
        for fi26 := 0; fi26 < 5; fi26++ {
            switch 4294967295 {
            case 9223372036854775807:
                if false {
                    if len(v4) > 0 {
                        v4[0] = (fi26 << ((38) & 63))
                    }
                } else {
                    v27 := GenS4{F0: 0}
                    _ = len(v2)
                    var v28 float64 = v1
                    _ = v27
                    _ = v28
                }
            case -77:
                switch "" {
                case "a":
                    if v25 {
                        fmt.Println(v1, v3, v6)
                        v3 = (^(v3))
                    }
                case "":
                    v6 = 107
                case "\n\t":
                    var v29 int64 = -530
                    _ = v29
                }
            case 348:
                var v30 int8 = 1
                _ = v30
            }
            var v31 int32 = 0
            _ = fi26
            _ = v31
        }
    case 4294967295:
        if (v25 && false) {
            if v25 {
                if (377 <= 0) {
                    var v32 int8 = 54
                    var v33 uint8 = 42
                    if (true || (42 != 0)) {
                        v25 = v25
                        v34 := GenS2{F0: 1000, F1: 0, F2: -1}
                        for fi35 := 0; fi35 < 18; fi35++ {
                            if (v25 && (v33 <= v33)) {
                                v1 += 0.0
                            }
                            v32 += (v32 / ((v32) | 1))
                            if (v25 || (true && false)) {
                                v36 := v2["foo bar"]
                                v6 = (v6 & (v6 << ((2) & 63)))
                                _ = v36
                            }
                            _ = fi35
                        }
                        _ = v34
                    }
                    _ = v32
                    _ = v33
                } else {
                    v3 = (v3 | (v3 * v3))
                }
                var v37 int64 = 9223372036854775807
                v37 -= v37
                _ = v37
            } else {
                if ((false || v25) || v25) {
                    if ((v25 && false) && v25) {
                        var v38 uint32 = 42
                        var v39 uint8 = 42
                        v1 = (v1 - v1)
                        _ = v38
                        _ = v39
                    } else {
                        if (1 < -128) {
                            var v40 uint32 = v6
                            _ = v40
                        }
                    }
                }
                v3 = v3
                v6 += (^(v6))
            }
        } else {
            var v41 bool = false
            _ = v41
        }
    }
    var v42 string = "foo bar"
    v43 := GenS2{F0: 42, F1: 1, F2: 42}
    if v25 {
        v6 += 2147483647
    } else {
        if (!(65535 < 32767)) {
            v44 := v2[("\n\t" + v42)]
            v6 = (v6 ^ (v6 + v6))
            v45 := GenS2{F0: -274, F1: 255, F2: 127}
            _ = v44
            _ = v45
        } else {
            v46 := GenS1{F0: -39}
            v42 = "hello"
            _ = v46
        }
        if true {
            var v47 int64 = -534
            var v48 int32 = 252
            if v25 {
                v49 := GenS5{F0: v48, F1: 32767, F2: 102}
                v50 := GenS4{F0: 127}
                _ = v49
                _ = v50
            } else {
                if v25 {
                    v2[v42] = (true && true)
                }
            }
            _ = v47
            _ = v48
        } else {
            v51 := GenS6{F0: (0 == 0), F1: 1}
            fmt.Println(v3, v42, v1)
            for fi52 := 0; fi52 < 20; fi52++ {
                switch v42 {
                case "a":
                    v53 := v2[v42]
                    _ = v53
                case "foo bar":
                    var v54 string = v42
                    _ = v54
                case "hello":
                    for fi55 := 0; fi55 < 16; fi55++ {
                        switch (-(fi55)) {
                        case 0:
                            v56 := GenS7{F0: fi55, F1: v6, F2: 41}
                            _ = v56
                        case -904:
                            var v57 int8 = -124
                            _ = v57
                        case -1:
                            var v58 int = -1
                            _ = v58
                        }
                        v59 := GenS8{F0: 429, F1: v25, F2: ("\n\t" + v42)}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi55
                        _ = v59
                    }
                }
                fmt.Println(v1, v3)
                _ = fi52
            }
            _ = v51
        }
        var v60 int8 = 1
        _ = v60
    }
    switch 127 {
    case 2147483647:
        v61 := GenS2{F0: 1000, F1: 18446744073709551615, F2: 32767}
        _ = v61
    default:
        fmt.Println(v3, v6, v6)
    }
    for fi62 := 0; fi62 < 13; fi62++ {
        var v63 float32 = 1e10
        if v25 {
            v42 = "foo bar"
            for fi64 := 0; fi64 < 3; fi64++ {
                var v65 int = (fi64 / (((fi64 ^ (fi62 & fi62))) | 1))
                _ = fi64
                _ = v65
            }
        } else {
            var v66 bool = (fi62 == -2147483648)
            v4 = append(v4, fi62)
            if ((0 >= -128) && ((v42 == v42) && (!(!(42 == 127))))) {
                for fi67 := 0; fi67 < 1; fi67++ {
                    fmt.Println(v3, v25)
                    v63 -= -2.25
                    v68 := GenS9{F0: fi62, F1: 0, F2: -1}
                    _ = fi67
                    _ = v68
                }
                fmt.Println(v25)
                if true {
                    v3 = v3
                    for fi69 := 0; fi69 < 12; fi69++ {
                        _ = len(v4)
                        fmt.Println(fi69, v25, v63)
                        v42 = v42
                        _ = fi69
                    }
                    var v70 int = 128
                    _ = v70
                } else {
                    v71 := GenS10{F0: 42}
                    fmt.Println(v42, v1, v63)
                    if v25 {
                        var v72 uint = 18446744073709551615
                        _ = v72
                    }
                    _ = v71
                }
            } else {
                v73 := GenS11{F0: 9223372036854775807, F1: -997, F2: 0}
                v63 = 3.14159
                if true {
                    if (v63 > float32((3.14159 * ((v1 * v1) - v1)))) {
                        var v74 float32 = 0.1
                        fmt.Println(v25, v74, v74)
                        fmt.Println(v63, v6, v25)
                        _ = v74
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            var v75 uint = 434
                            _ = v75
                        }
                    }
                    var v76 uint32 = v6
                    _ = v76
                }
                _ = v73
            }
            _ = v66
        }
        _ = fi62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
