package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int8; F2 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 bool; F2 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint }

func main() {
    v1 := make(map[string]string)
    v2 := []int{-9223372036854775808, 9223372036854775807, 1, 2147483647}
    v3 := []uint32{676}
    var v4 int64 = 42
    v5 := GenS1{F0: (v4 - (v4 * v4))}
    if (0 == 65535) {
        if (42 != 1) {
            fmt.Println(v4)
            var v6 uint16 = 65535
            _ = v6
        } else {
            fmt.Println(v4)
        }
        if (false && true) {
            v4 += (v4 * (v4 * (v4 | v4)))
            switch 921 {
            case 0:
                v7 := GenS1{F0: v4}
                _ = v7
            case 32768:
                for fi8 := 0; fi8 < 20; fi8++ {
                    v2 = append(v2, 397)
                    _ = fi8
                }
            case 1:
                fmt.Println(v4)
            }
            fmt.Println(v4)
        } else {
            v2 = append(v2, 89)
            v4 += (v4 | v4)
            fmt.Println(v4)
        }
        if (!true) {
            if len(v3) > 0 {
                v3[0] = 163
            }
        }
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        switch ("hello" + ("foo bar" + "")) {
        case "hello":
            for fi10 := 0; fi10 < 14; fi10++ {
                fmt.Println(v4)
                v4 -= v4
                v11 := GenS2{F0: 42, F1: 0, F2: (1000 <= -1)}
                _ = fi10
                _ = v11
            }
        case "":
            if (0.0 != float32((1.5 - (1e10 - (1e10 * (-2.25 + (((3.14159 / (1e10 * 3.14159)) * -2.25) / 1.0))))))) {
                var v12 uint32 = 2
                _ = v12
            }
        case "\n\t":
            if true {
                if (1.5 > 3.14159) {
                    if true {
                        for fi13 := 0; fi13 < 20; fi13++ {
                            if false {
                                if (42 > 918) {
                                    if true {
                                        if (-76 > -1) {
                                            var v14 uint16 = 2
                                            if (18446744073709551615 == 18446744073709551615) {
                                                v4 += v4
                                                var v15 float64 = -2.25
                                                var v16 uint64 = 4294967295
                                                v15 = ((1e10 * 1e10) * v15)
                                                v1[""] = "hello"
                                                _ = v15
                                                _ = v16
                                            } else {
                                                if true {
                                                    fmt.Println(v14, v4, v4)
                                                    v4 += v4
                                                }
                                            }
                                            _ = v14
                                        } else {
                                            v4 -= v4
                                            v17 := GenS3{F0: 42}
                                            v18 := v1[""]
                                            _ = v17
                                            _ = v18
                                        }
                                        var v19 int8 = 0
                                        _ = v19
                                    }
                                    v4 += (v4 | v4)
                                } else {
                                    var v20 int = fi9
                                    v3 = append(v3, 17)
                                    v21 := GenS4{F0: 42, F1: ((!(42 > 651)) || false), F2: 128}
                                    _ = v20
                                    _ = v21
                                }
                                if false {
                                    fmt.Println(fi9)
                                    v22 := GenS3{F0: 255}
                                    _ = v22
                                } else {
                                    v4 += (v4 >> ((255) & 63))
                                }
                            }
                            switch "a" {
                            case "\n\t":
                                v23 := v1[("a" + "foo bar")]
                                _ = v23
                            case "a":
                                var v24 uint8 = 206
                                _ = v24
                            case "foo bar":
                                var v25 int16 = -362
                                _ = v25
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            if false {
                                var v26 int16 = 395
                                _ = v26
                            } else {
                                if (v4 <= v4) {
                                    var v27 uint16 = 42
                                    if (v27 > v27) {
                                        v27 = 32767
                                        if (462 <= 42) {
                                            v28 := GenS1{F0: (v4 / ((17) | 1))}
                                            _ = v28
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v29 int8 = 100
                                            v29 += (v29 + 1)
                                            _ = v29
                                        }
                                        var v30 uint32 = 4294967295
                                        var v31 uint = 2
                                        var v32 uint16 = v27
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    }
                                    v33 := v1["a"]
                                    var v34 uint8 = 255
                                    var v35 float64 = ((1.5 / float64((((((-2.25 - 1e10) + (1e10 / ((3.14159 + -2.25) * float32((float64(1e10) * (-2.25 + float64((0.0 - 0.0)))))))) / float32(float64((1.5 + 3.14159)))) * -2.25) / (0.1 / 3.14159)))) * 0.0)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v27
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                } else {
                                    v36 := GenS5{F0: -1}
                                    _ = v36
                                }
                                v37 := GenS6{F0: 0}
                                _ = v37
                            }
                            _ = fi13
                        }
                        var v38 bool = false
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v38
                    }
                }
            }
        default:
            if (-1 < -392) {
                v4 = (v4 & -9223372036854775808)
                for fi39 := 0; fi39 < 3; fi39++ {
                    v3 = append(v3, 42)
                    _ = fi39
                }
            }
        }
        v4 = 148
        _ = fi9
    }
    if true {
        var v40 int32 = -1
        fmt.Println(v4, v4)
        v41 := GenS6{F0: 4294967295}
        _ = v40
        _ = v41
    }
    switch 140 {
    case -702:
        v4 = -1
    case 127:
        for fi42 := 0; fi42 < 14; fi42++ {
            v43 := v1["\n\t"]
            v44 := GenS4{F0: 100, F1: ((2 == 386) && (572 <= 65535)), F2: 2}
            _ = fi42
            _ = v43
            _ = v44
        }
    default:
        var v45 int8 = 100
        _ = v45
    }
    var v46 float64 = 3.14159
    v47 := v1["hello"]
    fmt.Println(v4, v4, v46)
    fmt.Println(v47, v4)
    var v48 uint64 = 630
    for fi49 := 0; fi49 < 9; fi49++ {
        switch ("" + ((v47 + v47) + ("\n\t" + ((v47 + (((v47 + "hello") + "hello") + ((((v47 + v47) + v47) + (v47 + (v47 + v47))) + "foo bar"))) + v47)))) {
        case "hello":
            if (false && true) {
                v48 = 725
                v50 := GenS1{F0: (v4 | v4)}
                _ = v50
            } else {
                v2 = append(v2, fi49)
                v48 += v48
            }
        default:
            if false {
                var v51 int64 = v4
                _ = v51
            }
        }
        var v52 bool = ((-2.25 + v46) > v46)
        _ = fi49
        _ = v52
    }
    var v53 int32 = -1
    v53 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v53
}
