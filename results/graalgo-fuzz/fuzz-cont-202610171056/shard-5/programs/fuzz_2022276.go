package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 uint8; F1 bool; F2 int8 }
type GenS7 struct { F0 string; F1 int8; F2 int16 }
type GenS8 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 18446744073709551615
    var v3 int32 = -2147483648
    switch 128 {
    case -1:
        v3 = -376
    case 1:
        v3 = (-(v3))
    }
    for fi4 := 0; fi4 < 16; fi4++ {
        if ((3.14159 * (1e10 - 3.14159)) != float64(3.14159)) {
            v5 := []int{fi4, fi4, fi4}
            _ = v5
        }
        v6 := []int{(fi4 << ((v2) & 63))}
        _ = fi4
        _ = v6
    }
    v3 += (v3 ^ int32(v2))
    var v7 string = "hello"
    var v8 int64 = 983
    v9 := []int{32768, 1, 128}
    fmt.Println(v7)
    var v10 int = 9223372036854775807
    switch (v7 + v7) {
    case "a":
        var v11 int64 = (v8 & (v8 / (((v8 ^ int64(v2))) | 1)))
        _ = v11
    case "hello":
        var v12 uint8 = 127
        _ = v12
    default:
        for fi13 := 0; fi13 < 17; fi13++ {
            if (10 > 67) {
                for fi14 := 0; fi14 < 5; fi14++ {
                    switch (v7 + "hello") {
                    case "foo bar":
                        v15 := GenS1{F0: 0}
                        _ = v15
                    case "":
                        v7 = "a"
                    default:
                        if ((v8 ^ v8) <= (v8 | (v8 * (v8 / (((v8 / ((0) | 1))) | 1))))) {
                            if (float64(-2.25) >= (0.0 + ((-2.25 / 1.0) / 1e10))) {
                                var v16 float64 = 3.14159
                                _ = v16
                            }
                            switch (v10 & (fi13 % (((v10 * fi13)) | 1))) {
                            case -1:
                                delete(v1, 396)
                            case 993:
                                var v17 int8 = 1
                                _ = v17
                            default:
                                v18 := GenS2{F0: v8, F1: 0}
                                _ = v18
                            }
                            var v19 uint8 = 1
                            if true {
                                v20 := GenS2{F0: (v8 | v8), F1: 32767}
                                v21 := GenS3{F0: uint8(v2), F1: (v19 << ((v2) & 63)), F2: -607}
                                v19 += (v19 >> (((v2 << (((v2 % ((0) | 1))) & 63))) & 63))
                                _ = v20
                                _ = v21
                            }
                            _ = v19
                        } else {
                            switch fi14 {
                            case -451:
                                v22 := GenS2{F0: (v8 | v8), F1: 736}
                                _ = v22
                            default:
                                if len(v9) > 0 {
                                    v9[0] = fi14
                                }
                                var v23 int16 = -32768
                                _ = v23
                            }
                        }
                    }
                    _ = fi14
                }
            } else {
                v1[42] = 400
                switch (((v7 + (v7 + ((v7 + "a") + "\n\t"))) + "") + "") {
                case "hello":
                    if false {
                        for fi24 := 0; fi24 < 19; fi24++ {
                            if true {
                                fmt.Println(fi24, v10)
                                var v25 int64 = (v8 ^ 0)
                                if (v8 != (v8 - 9223372036854775807)) {
                                    var v26 int = (^(fi24))
                                    v9 = append(v9, (v26 / (((v26 | v10)) | 1)))
                                    var v27 float64 = (((0.1 / -2.25) + 0.1) * 1e10)
                                    _ = v26
                                    _ = v27
                                } else {
                                    v25 = (v8 << ((4294967295) & 63))
                                    var v28 string = ("" + ((v7 + (v7 + "")) + (v7 + v7)))
                                    var v29 int8 = -128
                                    _ = v28
                                    _ = v29
                                }
                                _ = v25
                            } else {
                                switch (((((v7 + ("" + v7)) + ("\n\t" + v7)) + (v7 + "\n\t")) + "foo bar") + (((v7 + v7) + v7) + ("foo bar" + (v7 + "hello")))) {
                                case "\n\t":
                                    v7 = "a"
                                case "a":
                                    fmt.Println(fi13, v8)
                                case "foo bar":
                                    if false {
                                        var v30 uint64 = 893
                                        v31 := GenS4{F0: 537, F1: v2}
                                        var v32 bool = (!false)
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    } else {
                                        var v33 string = v7
                                        v34 := GenS5{F0: 0, F1: v33}
                                        v33 = v33
                                        _ = v33
                                        _ = v34
                                    }
                                }
                                v35 := v1[45]
                                _ = v35
                            }
                            var v36 uint32 = 4294967295
                            switch v7 {
                            case "hello":
                                if true {
                                    fmt.Println(v10, v3, v2)
                                    v37 := GenS6{F0: 2, F1: true, F2: 42}
                                    switch 4294967295 {
                                    case 2147483647:
                                        v36 = (v36 & 4294967295)
                                        v2 = v2
                                        if ((fi24 | int(v8)) < (fi13 >> ((65535) & 63))) {
                                            v2 = (^(v2))
                                        }
                                        var v38 uint64 = 30
                                        _ = v38
                                    case -9223372036854775808:
                                        v3 += v3
                                    case 713:
                                        v36 -= (^(v36))
                                    }
                                    _ = v37
                                } else {
                                    v2 = (v2 - v2)
                                }
                            default:
                                v8 -= (v8 % (((v8 % ((1) | 1))) | 1))
                            }
                            v39 := GenS5{F0: 1000, F1: "\n\t"}
                            var v40 int16 = 0
                            var v41 string = v7
                            _ = fi24
                            _ = v36
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        delete(v1, 475)
                        v42 := GenS5{F0: 32767, F1: (v7 + "foo bar")}
                        _ = v42
                    } else {
                        v3 += v3
                        var v43 float32 = -2.25
                        if (18446744073709551615 <= 42) {
                            var v44 int = (v10 ^ v10)
                            v9 = append(v9, 32767)
                            _ = v44
                        } else {
                            v3 -= 551
                            v2 += v2
                            v45 := GenS7{F0: "hello", F1: 100, F2: -590}
                            _ = v45
                        }
                        _ = v43
                    }
                case "":
                    var v46 string = v7
                    _ = v46
                }
            }
            _ = fi13
        }
    }
    for fi47 := 0; fi47 < 1; fi47++ {
        v8 = (v8 ^ v8)
        if (v7 <= (v7 + "\n\t")) {
            v10 += fi47
            var v48 uint64 = 2
            for fi49 := 0; fi49 < 6; fi49++ {
                if false {
                    v10 = (v10 << ((v2) & 63))
                }
                v3 = 2147483647
                var v50 uint16 = 449
                _ = fi49
                _ = v50
            }
            _ = v48
        } else {
            for fi51 := 0; fi51 < 20; fi51++ {
                switch (-(v10)) {
                case 0:
                    switch 2147483647 {
                    case -236:
                        delete(v1, 2147483647)
                    case 32767:
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        switch ((v7 + "foo bar") + v7) {
                        case "hello":
                            var v52 int = 32768
                            _ = v52
                        default:
                            var v53 uint = (v2 & (v2 | 786))
                            _ = v53
                        }
                        v3 += -342
                    case 46:
                        v8 -= v8
                    }
                default:
                    if ((v2 << ((v2) & 63)) < 121) {
                        fmt.Println(v2)
                        fmt.Println(v2, v2)
                    } else {
                        v8 = v8
                        v8 = (-(v8))
                    }
                }
                v8 = 0
                v10 = 143
                _ = fi51
            }
            var v54 uint16 = 1
            _ = v54
        }
        for fi55 := 0; fi55 < 2; fi55++ {
            if len(v9) > 0 {
                v9[0] = (fi47 >> (((v2 & v2)) & 63))
            }
            fmt.Println(fi55, fi47, v10)
            _ = fi55
        }
        _ = fi47
    }
    v56 := GenS8{F0: v8}
    for fi57 := 0; fi57 < 2; fi57++ {
        switch 10 {
        case 128:
            var v58 int64 = v8
            _ = v58
        case -9223372036854775808:
            var v59 float64 = -2.25
            _ = v59
        }
        v2 -= (v2 & v2)
        _ = fi57
    }
    for fi60 := 0; fi60 < 14; fi60++ {
        if len(v9) > 0 {
            v9[0] = v10
        }
        _ = fi60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v56
    _ = v7
    _ = v8
    _ = v9
}
