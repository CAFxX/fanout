package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 bool; F1 int16; F2 int64 }
type GenS3 struct { F0 int64; F1 int8; F2 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64 }

func main() {
    v1 := []string{("\n\t" + "foo bar")}
    v2 := make(map[int]string)
    v3 := []string{}
    var v4 uint64 = 993
    var v5 int8 = 100
    var v6 uint = 4294967295
    v2[128] = ("foo bar" + "\n\t")
    if false {
        fmt.Println(v6, v5, v6)
    }
    if (!(306 <= 32767)) {
        v7 := GenS1{F0: 0, F1: int32(v4)}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint = (^(v6))
            if false {
                v10 := GenS2{F0: (false || true), F1: 224, F2: -1}
                if false {
                    v11 := v2[(fi8 - (fi8 >> ((v6) & 63)))]
                    v4 += (v4 & (^(v4)))
                    var v12 int64 = -9223372036854775808
                    _ = v11
                    _ = v12
                } else {
                    switch (fi8 ^ (fi8 - (fi8 - fi8))) {
                    case 382:
                        if (-75 > 32767) {
                            if false {
                                if true {
                                    v4 -= v4
                                } else {
                                    for fi13 := 0; fi13 < 15; fi13++ {
                                        v9 += v6
                                        var v14 uint16 = 42
                                        v15 := v2[-1]
                                        v16 := GenS3{F0: -242, F1: v5, F2: (-(fi8))}
                                        _ = fi13
                                        _ = v14
                                        _ = v15
                                        _ = v16
                                    }
                                    v6 -= 0
                                }
                                fmt.Println(v9)
                            }
                            v5 = (v5 >> (((v9 - (^(v9)))) & 63))
                            var v17 int = 32768
                            _ = v17
                        }
                    case 127:
                        var v18 int16 = -32768
                        _ = v18
                    case 72:
                        v4 -= (v4 / ((13) | 1))
                    }
                }
                v19 := GenS2{F0: true, F1: -596, F2: -9223372036854775808}
                _ = v10
                _ = v19
            } else {
                v20 := GenS4{F0: 0}
                _ = v20
            }
            _ = fi8
            _ = v9
        }
        fmt.Println(v4)
    }
    var v21 uint = 42
    var v22 float64 = 1e10
    var v23 uint16 = 845
    v2[-283] = "foo bar"
    var v24 uint8 = 1
    if (!(9223372036854775807 > 42)) {
        v22 -= ((float64(3.14159) - v22) / v22)
    }
    fmt.Println(v4, v22, v23)
    if false {
        _ = len(v3)
        v21 -= uint(v23)
    } else {
        for fi25 := 0; fi25 < 17; fi25++ {
            v3 = append(v3, "a")
            for fi26 := 0; fi26 < 15; fi26++ {
                var v27 uint64 = (v4 / ((v4) | 1))
                _ = fi26
                _ = v27
            }
            v23 += v23
            _ = fi25
        }
        var v28 uint = 0
        _ = v28
    }
    v22 = 1e10
    for fi29 := 0; fi29 < 4; fi29++ {
        if false {
            for fi30 := 0; fi30 < 15; fi30++ {
                v31 := v2[(fi30 >> ((v6) & 63))]
                _ = fi30
                _ = v31
            }
            switch "" {
            case "a":
                v21 = uint(v24)
            default:
                _ = len(v2)
            }
            fmt.Println(v22)
        }
        switch 4294967295 {
        case -215:
            v6 += v6
        }
        _ = fi29
    }
    v32 := GenS5{F0: -128}
    for fi33 := 0; fi33 < 2; fi33++ {
        switch fi33 {
        case 4294967295:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -81:
            if true {
                switch ("hello" + (("" + "hello") + "foo bar")) {
                case "hello":
                    if (!(!true)) {
                        v6 = (v6 * v21)
                        if true {
                            switch fi33 {
                            case 130:
                                for fi34 := 0; fi34 < 11; fi34++ {
                                    if true {
                                        v2[(fi33 & int(v24))] = ""
                                        v35 := GenS5{F0: -114}
                                        if (((v22 - (v22 / -2.25)) / (1.5 * v22)) != 0.1) {
                                            v36 := GenS1{F0: uint(v24), F1: 0}
                                            _ = v36
                                        } else {
                                            fmt.Println(v21, v6)
                                            var v37 string = "hello"
                                            _ = v37
                                        }
                                        _ = v35
                                    }
                                    _ = fi34
                                }
                            case 32768:
                                var v38 uint8 = (v24 >> ((v6) & 63))
                                _ = v38
                            default:
                                switch ("" + "\n\t") {
                                case "hello":
                                    _ = len(v1)
                                default:
                                    v24 = 0
                                }
                            }
                        } else {
                            for fi39 := 0; fi39 < 3; fi39++ {
                                v40 := GenS2{F0: (0.1 <= 3.14159), F1: -32768, F2: -1}
                                var v41 int32 = 611
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                            _ = len(v2)
                        }
                    }
                case "":
                    fmt.Println(v23, v4)
                case "foo bar":
                    var v42 string = "hello"
                    _ = v42
                default:
                    var v43 int32 = 0
                    _ = v43
                }
            } else {
                fmt.Println(v22, v24, v4)
                var v44 uint8 = v24
                var v45 uint16 = 255
                _ = v44
                _ = v45
            }
        case -752:
            v46 := GenS6{F0: 42}
            _ = v46
        default:
            v47 := GenS6{F0: 45}
            _ = v47
        }
        var v48 float64 = v22
        _ = fi33
        _ = v48
    }
    for fi49 := 0; fi49 < 9; fi49++ {
        if (2 >= v23) {
            fmt.Println(v6)
        } else {
            fmt.Println(v21, fi49, v4)
            delete(v2, (fi49 ^ (fi49 << ((v21) & 63))))
            v50 := GenS1{F0: 65535, F1: 0}
            _ = v50
        }
        _ = fi49
    }
    if false {
        var v51 uint64 = v4
        if (false || (v5 >= v5)) {
            switch ("a" + "foo bar") {
            case "\n\t":
                var v52 int8 = (^(v5))
                _ = v52
            }
            v24 -= (v24 - uint8(v23))
        } else {
            v51 = 250
        }
        _ = v51
    } else {
        v53 := GenS6{F0: 0}
        _ = v53
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
