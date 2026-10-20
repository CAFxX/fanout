package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 string }
type GenS2 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS3 struct { F0 bool; F1 int64; F2 int16 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint8; F1 int; F2 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = (0.1 * 1.5)
    v3 := []string{("foo bar" + "hello"), ("foo bar" + ""), ("" + "a"), ("hello" + "")}
    fmt.Println(v2)
    var v4 int = 32768
    v2 -= 1e10
    v5 := GenS1{F0: 229, F1: -9223372036854775808, F2: "hello"}
    if (4294967295 < 487) {
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 int8 = 100
            if (true && true) {
                for fi8 := 0; fi8 < 1; fi8++ {
                    fmt.Println(v4, fi6, fi8)
                    fmt.Println(v2, fi6)
                    var v9 uint32 = 2147483648
                    _ = fi8
                    _ = v9
                }
            } else {
                v10 := GenS2{F0: true, F1: 42, F2: 42}
                if (188 > 1) {
                    var v11 uint8 = 229
                    var v12 int = (fi6 << ((0) & 63))
                    _ = v11
                    _ = v12
                } else {
                    var v13 int64 = 0
                    v4 = v4
                    v7 = (v7 ^ v7)
                    _ = v13
                }
                _ = v10
            }
            for fi14 := 0; fi14 < 2; fi14++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if (true && true) {
                    v2 += 3.14159
                    if false {
                        if ("hello" != (("foo bar" + ("foo bar" + "\n\t")) + "")) {
                            switch ("foo bar" + "a") {
                            case "\n\t":
                                if (false && ((1.5 / v2) > float64(3.14159))) {
                                    if false {
                                        v7 = (-(v7))
                                        var v15 int8 = v7
                                        if false {
                                            fmt.Println(v7)
                                            v7 = v15
                                            v16 := GenS2{F0: false, F1: 1, F2: 42}
                                            _ = v16
                                        } else {
                                            switch (fi6 / ((1) | 1)) {
                                            case 2147483647:
                                                var v17 uint16 = 2
                                                _ = v17
                                            case 32768:
                                                v15 = v7
                                            case 923:
                                                v18 := GenS1{F0: -2147483648, F1: 1, F2: "\n\t"}
                                                _ = v18
                                            default:
                                                v19 := GenS3{F0: true, F1: 1, F2: -1}
                                                _ = v19
                                            }
                                            var v20 int32 = 42
                                            switch (fi6 + v4) {
                                            case 384:
                                                v7 = v15
                                            case 127:
                                                if false {
                                                    var v21 uint32 = 0
                                                    var v22 bool = true
                                                    _ = v21
                                                    _ = v22
                                                }
                                            }
                                            var v23 uint32 = 70
                                            _ = v20
                                            _ = v23
                                        }
                                        v24 := GenS4{F0: 56, F1: "a"}
                                        v25 := v1[371]
                                        _ = v15
                                        _ = v24
                                        _ = v25
                                    }
                                } else {
                                    v7 = v7
                                    v7 = v7
                                }
                            case "hello":
                                var v26 float32 = 0.0
                                _ = v26
                            default:
                                if true {
                                    if (true || false) {
                                        if (1 > 55) {
                                            var v27 int8 = (v7 + v7)
                                            fmt.Println(fi14)
                                            _ = v27
                                        }
                                        var v28 uint32 = 2147483647
                                        _ = len(v3)
                                        _ = v28
                                    }
                                    v3 = append(v3, "hello")
                                } else {
                                    if (!(0.1 < ((((((0.1 + 3.14159) - ((1e10 * ((0.0 / -2.25) - float32((float64(-2.25) - 0.0)))) / 1e10)) * (3.14159 + (0.1 / 1.5))) * (((1e10 / (-2.25 / (0.1 * ((1e10 - (3.14159 / 0.1)) + float32(1e10))))) * 0.0) + -2.25)) - 0.1) * (0.1 / 1.5)))) {
                                        var v29 uint32 = 65535
                                        _ = v29
                                    } else {
                                        v2 = v2
                                        if true {
                                            if len(v3) > 0 {
                                                v3[0] = ("\n\t" + "foo bar")
                                            }
                                        }
                                        var v30 float32 = 1e10
                                        _ = v30
                                    }
                                    v4 = (fi6 >> ((348) & 63))
                                    if true {
                                        if true {
                                            switch (^(fi6)) {
                                            case 4294967295:
                                                var v31 int8 = v7
                                                _ = v31
                                            case 2147483647:
                                                v4 -= 2147483647
                                            case -652:
                                                var v32 float64 = v2
                                                _ = v32
                                            default:
                                                var v33 int = (fi14 << ((230) & 63))
                                                _ = v33
                                            }
                                            if (146 != 412) {
                                                v4 += 32768
                                                v7 = v7
                                            } else {
                                                var v34 int64 = 20
                                                _ = v34
                                            }
                                        } else {
                                            var v35 uint8 = 0
                                            v35 = 1
                                            if false {
                                                fmt.Println(v4)
                                            } else {
                                                v2 += float64(1.5)
                                                var v36 float64 = -2.25
                                                _ = v36
                                            }
                                            _ = v35
                                        }
                                    }
                                }
                            }
                            v37 := GenS3{F0: (float64(float32(3.14159)) == float64(0.1)), F1: 42, F2: 0}
                            var v38 int32 = 1
                            _ = v37
                            _ = v38
                        }
                    }
                    var v39 float32 = 1.5
                    _ = v39
                } else {
                    v2 -= 0.1
                    v3 = append(v3, ("foo bar" + "a"))
                }
                v2 -= 0.1
                delete(v1, 65535)
                _ = fi14
            }
            _ = fi6
            _ = v7
        }
        if true {
            v4 -= (^(v4))
        } else {
            for fi40 := 0; fi40 < 6; fi40++ {
                for fi41 := 0; fi41 < 11; fi41++ {
                    v42 := GenS5{F0: 128, F1: 893, F2: 0}
                    var v43 float32 = ((1.5 - -2.25) - 0.1)
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                for fi44 := 0; fi44 < 15; fi44++ {
                    if (!true) {
                        if (32767 < 1000) {
                            v45 := GenS3{F0: true, F1: 1, F2: -20}
                            if true {
                                v2 = float64(((1.5 / (-2.25 + 3.14159)) / ((((0.1 - 1.5) * (1e10 * ((-2.25 + 3.14159) / 1.0))) + 1.5) - ((0.0 + float32(v2)) + 0.0))))
                            }
                            v4 -= fi44
                            _ = v45
                        } else {
                            v46 := GenS3{F0: true, F1: 9223372036854775807, F2: 1}
                            _ = v46
                        }
                    }
                    v47 := GenS2{F0: (false || (true || false)), F1: 273, F2: 144}
                    _ = fi44
                    _ = v47
                }
                _ = fi40
            }
        }
        var v48 int64 = -819
        _ = v48
    } else {
        if false {
            v3 = append(v3, "foo bar")
        }
        for fi49 := 0; fi49 < 4; fi49++ {
            var v50 uint16 = 65535
            for fi51 := 0; fi51 < 19; fi51++ {
                v50 = (v50 * v50)
                _ = fi51
            }
            _ = fi49
            _ = v50
        }
    }
    v4 += v4
    var v52 int32 = 0
    for fi53 := 0; fi53 < 19; fi53++ {
        var v54 float64 = float64((0.1 / ((3.14159 - 1e10) - 3.14159)))
        _ = fi53
        _ = v54
    }
    v4 = v4
    for fi55 := 0; fi55 < 11; fi55++ {
        var v56 uint16 = 237
        _ = fi55
        _ = v56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
}
