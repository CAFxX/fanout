package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS6 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS7 struct { F0 int16; F1 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -1
    v3 := []uint32{694, 2147483648, 2}
    var v4 int = -1
    v5 := make(map[string]int)
    var v6 bool = true
    for fi7 := 0; fi7 < 12; fi7++ {
        if false {
            fmt.Println(v2, v2, fi7)
        } else {
            v6 = v6
        }
        _ = fi7
    }
    switch "a" {
    case "foo bar":
        var v8 bool = v6
        _ = v8
    case "a":
        v9 := v1[612]
        _ = v9
    case "hello":
        var v10 uint64 = 42
        _ = v10
    default:
        v11 := GenS1{F0: 42, F1: -661}
        _ = v11
    }
    var v12 int64 = 0
    v12 = (v12 >> ((4294967295) & 63))
    switch "a" {
    case "hello":
        var v13 uint64 = 0
        _ = v13
    case "a":
        if (0 <= 4294967295) {
            for fi14 := 0; fi14 < 9; fi14++ {
                v4 = v4
                fmt.Println(v2, v2)
                _ = fi14
            }
            var v15 int8 = 34
            if ((-(v4)) < 127) {
                var v16 uint8 = 2
                _ = v16
            }
            _ = v15
        } else {
            if (62 == -128) {
                v2 = (v2 / ((v2) | 1))
                if false {
                    if (v6 && (!(false || v6))) {
                        for fi17 := 0; fi17 < 5; fi17++ {
                            switch "" {
                            case "hello":
                                var v18 bool = (v6 || (658 == 2))
                                _ = v18
                            case "foo bar":
                                v19 := GenS2{F0: 26}
                                _ = v19
                            case "a":
                                fmt.Println(v6)
                            default:
                                if ((v2 * (v2 % (((v2 << ((1) & 63))) | 1))) != (v2 % ((v2) | 1))) {
                                    if v6 {
                                        switch "\n\t" {
                                        case "hello":
                                            v5["foo bar"] = fi17
                                        case "a":
                                            var v20 bool = v6
                                            _ = v20
                                        }
                                        v21 := GenS3{F0: v6, F1: ""}
                                        v2 = 2147483647
                                        _ = v21
                                    } else {
                                        v2 = int32(v12)
                                        v12 -= (v12 | v12)
                                    }
                                }
                            }
                            _ = fi17
                        }
                    } else {
                        fmt.Println(v4, v4)
                        if v6 {
                            switch v4 {
                            case -9223372036854775808:
                                var v22 int32 = 42
                                _ = v22
                            case 0:
                                if v6 {
                                    var v23 bool = ((v2 == v2) && (851 <= 724))
                                    _ = v23
                                } else {
                                    if (v6 || (-1 < (v2 | 1))) {
                                        v24 := GenS3{F0: (0 < 34), F1: "foo bar"}
                                        for fi25 := 0; fi25 < 18; fi25++ {
                                            if (!((fi25 & -2147483648) >= (fi25 / ((32767) | 1)))) {
                                                if (1 > 555) {
                                                    v26 := GenS4{F0: 127}
                                                    _ = v26
                                                } else {
                                                    v27 := GenS2{F0: 49}
                                                    _ = v27
                                                }
                                            }
                                            _ = fi25
                                        }
                                        _ = v24
                                    } else {
                                        v2 = v2
                                        fmt.Println(v2, v4, v2)
                                    }
                                }
                            case 9223372036854775807:
                                if ((v12 % ((v12) | 1)) > 9223372036854775807) {
                                    switch -84 {
                                    case -301:
                                        v28 := v1[2147483648]
                                        _ = v28
                                    }
                                    var v29 uint64 = 42
                                    var v30 float32 = float32((-2.25 * 0.1))
                                    _ = v29
                                    _ = v30
                                } else {
                                    if false {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v31 uint8 = 255
                                        _ = v31
                                    }
                                }
                            default:
                                switch (-(v4)) {
                                case 32768:
                                    var v32 uint64 = 18446744073709551615
                                    _ = v32
                                default:
                                    var v33 uint64 = 1
                                    _ = v33
                                }
                            }
                            v34 := GenS5{F0: 42, F1: 255, F2: 0}
                            v6 = v6
                            _ = v34
                        } else {
                            fmt.Println(v2, v6)
                            for fi35 := 0; fi35 < 14; fi35++ {
                                for fi36 := 0; fi36 < 17; fi36++ {
                                    switch ("" + ("" + "hello")) {
                                    case "hello":
                                        if ("foo bar" > "hello") {
                                            switch "hello" {
                                            case "":
                                                if (3.14159 > 0.0) {
                                                    var v37 int = (-(fi35))
                                                    _ = v37
                                                }
                                            default:
                                                if true {
                                                    if ((1 == (v2 ^ 2147483647)) || v6) {
                                                        var v38 int8 = 1
                                                        v2 = (v2 << ((0) & 63))
                                                        v6 = (1 > 0)
                                                        _ = v38
                                                    } else {
                                                        v39 := GenS6{F0: 65535, F1: 255, F2: 296}
                                                        _ = v39
                                                    }
                                                }
                                            }
                                        } else {
                                            v4 = (fi36 + (^(fi35)))
                                            var v40 uint8 = 128
                                            var v41 string = "hello"
                                            _ = v40
                                            _ = v41
                                        }
                                    case "a":
                                        switch fi36 {
                                        case 902:
                                            v1[762] = (v12 <= (v12 / (((^(v12))) | 1)))
                                        default:
                                            v2 = 2147483647
                                        }
                                    default:
                                        var v42 uint = 110
                                        _ = v42
                                    }
                                    fmt.Println(fi35, v6, v4)
                                    _ = fi36
                                }
                                if (("hello" + "") == ("\n\t" + "")) {
                                    fmt.Println(v6, v12)
                                    for fi43 := 0; fi43 < 20; fi43++ {
                                        var v44 uint64 = 4294967295
                                        _ = fi43
                                        _ = v44
                                    }
                                }
                                v45 := v5[("foo bar" + ("" + "\n\t"))]
                                _ = fi35
                                _ = v45
                            }
                        }
                        v12 += int64(v4)
                    }
                    var v46 bool = v6
                    var v47 int64 = v12
                    _ = v46
                    _ = v47
                }
            } else {
                v6 = (!v6)
                var v48 uint32 = 865
                var v49 bool = v6
                _ = v48
                _ = v49
            }
        }
    default:
        if v6 {
            v50 := GenS3{F0: false, F1: ("a" + ("foo bar" + "a"))}
            switch (v4 % (((v4 | (v4 - (v4 - 42)))) | 1)) {
            case 207:
                var v51 int32 = -2147483648
                _ = v51
            }
            _ = v50
        } else {
            v4 = v4
        }
    }
    var v52 int32 = v2
    switch (v4 / (((v4 - v4)) | 1)) {
    case 127:
        var v53 int8 = -128
        _ = v53
    default:
        v54 := GenS4{F0: (v4 + 0)}
        _ = v54
    }
    v55 := GenS7{F0: 1, F1: 0}
    fmt.Println(v12, v12)
    v2 = v52
    for fi56 := 0; fi56 < 6; fi56++ {
        var v57 uint = 42
        _ = fi56
        _ = v57
    }
    if (1 <= 32767) {
        v58 := GenS3{F0: v6, F1: ("" + "foo bar")}
        var v59 uint64 = 979
        _ = v58
        _ = v59
    } else {
        var v60 uint = 1
        v4 += (v4 << (((v60 | v60)) & 63))
        v61 := GenS7{F0: 32767, F1: 14}
        _ = v60
        _ = v61
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v55
    _ = v6
}
