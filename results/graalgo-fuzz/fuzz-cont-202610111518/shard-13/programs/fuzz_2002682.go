package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint32 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 string; F1 uint32; F2 uint16 }
type GenS4 struct { F0 string; F1 string; F2 bool }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int8; F1 int32 }
type GenS8 struct { F0 int; F1 uint8 }

func main() {
    var v1 int64 = 0
    var v2 int16 = -1
    var v3 float32 = 1.5
    v4 := []uint32{18, 2147483647, 527}
    var v5 int64 = 42
    if len(v4) > 0 {
        v4[0] = 372
    }
    var v6 uint16 = 0
    switch (("hello" + "a") + "\n\t") {
    case "\n\t":
        if true {
            for fi7 := 0; fi7 < 12; fi7++ {
                v2 += 468
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 int16 = (v2 + 848)
                    var v10 float32 = float32(0.0)
                    v11 := GenS1{F0: (fi8 | fi7), F1: (v5 + -694), F2: 2147483647}
                    v2 = -32768
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
                var v12 int64 = v5
                _ = fi7
                _ = v12
            }
            switch ("\n\t" + "a") {
            case "foo bar":
                v3 = v3
            case "":
                var v13 int16 = v2
                _ = v13
            case "hello":
                for fi14 := 0; fi14 < 3; fi14++ {
                    v15 := GenS2{F0: (^(v2)), F1: fi14}
                    v6 = v6
                    _ = fi14
                    _ = v15
                }
            default:
                var v16 uint = 812
                _ = v16
            }
        }
    case "foo bar":
        var v17 uint8 = 1
        _ = v17
    case "":
        v4 = append(v4, 1)
    default:
        var v18 int8 = 0
        _ = v18
    }
    var v19 string = "a"
    if true {
        v20 := GenS3{F0: (((((v19 + v19) + ("" + v19)) + v19) + "\n\t") + ((v19 + v19) + ((v19 + v19) + v19))), F1: 1, F2: v6}
        _ = v20
    } else {
        if (!(!false)) {
            if false {
                v21 := GenS3{F0: (v19 + v19), F1: 67, F2: 32767}
                _ = v21
            }
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if ((v2 | 140) < (v2 ^ 286)) {
        v22 := GenS4{F0: "\n\t", F1: (v19 + ""), F2: true}
        _ = v22
    }
    v3 = (v3 * v3)
    var v23 int = 42
    v6 = v6
    var v24 bool = (42 < -703)
    if v24 {
        if (v6 != 390) {
            if v24 {
                v24 = (358 != 709)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if (-1 != 334) {
                    var v25 uint64 = 582
                    var v26 int16 = -636
                    var v27 float64 = 3.14159
                    _ = v25
                    _ = v26
                    _ = v27
                } else {
                    if (0 <= 18446744073709551615) {
                        if ((v6 != (v6 - (v6 & v6))) || (false && (280 >= 42))) {
                            var v28 int64 = 1
                            var v29 uint32 = 996
                            _ = v28
                            _ = v29
                        } else {
                            if (!v24) {
                                for fi30 := 0; fi30 < 12; fi30++ {
                                    var v31 uint32 = 42
                                    _ = fi30
                                    _ = v31
                                }
                                fmt.Println(v19)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                        var v32 bool = v24
                        var v33 uint16 = 255
                        _ = v32
                        _ = v33
                    }
                    v34 := GenS5{F0: 42, F1: (v6 << ((1) & 63))}
                    if len(v4) > 0 {
                        v4[0] = 2147483648
                    }
                    _ = v34
                }
            }
            var v35 bool = v24
            var v36 int16 = (v2 / ((-1) | 1))
            _ = v35
            _ = v36
        } else {
            for fi37 := 0; fi37 < 8; fi37++ {
                v5 = -9223372036854775808
                _ = fi37
            }
            if (55 != 179) {
                for fi38 := 0; fi38 < 19; fi38++ {
                    if v24 {
                        v4 = append(v4, 65535)
                    } else {
                        if len(v4) > 0 {
                            v4[0] = 42
                        }
                        if (false || false) {
                            _ = len(v4)
                            v39 := GenS5{F0: 100, F1: (v6 % (((v6 / ((v6) | 1))) | 1))}
                            switch -855 {
                            case 397:
                                _ = len(v4)
                            case 2147483647:
                                v24 = true
                            case 127:
                                var v40 uint64 = 1
                                _ = v40
                            default:
                                v41 := GenS6{F0: (v1 | (-(v1)))}
                                _ = v41
                            }
                            _ = v39
                        } else {
                            for fi42 := 0; fi42 < 9; fi42++ {
                                var v43 int64 = (v1 & v5)
                                if (!v24) {
                                    v24 = ((int64(v2) >= v5) || (false || v24))
                                    v23 = (fi42 >> ((42) & 63))
                                }
                                var v44 float32 = ((1.5 / v3) * float32(1.5))
                                _ = fi42
                                _ = v43
                                _ = v44
                            }
                            var v45 float32 = 0.0
                            _ = len(v4)
                            _ = v45
                        }
                        if (!(65535 > 255)) {
                            for fi46 := 0; fi46 < 12; fi46++ {
                                v24 = v24
                                fmt.Println(fi38, v1)
                                _ = fi46
                            }
                        } else {
                            var v47 uint32 = 686
                            if len(v4) > 0 {
                                v4[0] = v47
                            }
                            var v48 int = 0
                            _ = v47
                            _ = v48
                        }
                    }
                    fmt.Println(v23, v2)
                    _ = fi38
                }
                var v49 int64 = v5
                _ = v49
            } else {
                fmt.Println(v19, v24)
                v50 := GenS7{F0: 42, F1: int32(v5)}
                v6 -= v6
                _ = v50
            }
            for fi51 := 0; fi51 < 11; fi51++ {
                switch (("\n\t" + v19) + (((v19 + "") + v19) + "a")) {
                case "a":
                    var v52 uint32 = 2147483648
                    _ = v52
                case "\n\t":
                    var v53 int8 = -128
                    _ = v53
                }
                _ = fi51
            }
        }
    }
    var v54 int16 = v2
    v55 := GenS8{F0: (v23 >> ((255) & 63)), F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v6
}
