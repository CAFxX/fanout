package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 bool; F1 uint32; F2 int }
type GenS3 struct { F0 string; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 uint64; F2 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint32 }
func genFn0() uint32 {
    var v6 uint32 = 2
    var v7 uint8 = 24
    _ = v6
    _ = v7
    return v6
}

func main() {
    var v1 uint16 = 42
    var v2 bool = true
    v3 := []string{""}
    var v4 uint16 = 65535
    v5 := make(map[int]uint32)
    if false {
        switch "" {
        case "":
            for fi8 := 0; fi8 < 15; fi8++ {
                switch fi8 {
                case 921:
                    for fi9 := 0; fi9 < 16; fi9++ {
                        delete(v5, -2147483648)
                        v10 := GenS1{F0: false, F1: 100}
                        var v11 int64 = -9223372036854775808
                        v11 += v11
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                }
                fmt.Println(v1, fi8)
                var v12 float32 = 3.14159
                _ = fi8
                _ = v12
            }
        case "\n\t":
            switch 370 {
            case 9223372036854775807:
                v2 = true
            case 2147483647:
                v4 -= (v1 + 65535)
            default:
                v13 := v5[-16]
                _ = v13
            }
        case "hello":
            switch ("hello" + "hello") {
            case "a":
                var v14 uint16 = v1
                _ = v14
            case "\n\t":
                for fi15 := 0; fi15 < 17; fi15++ {
                    switch fi15 {
                    case -2147483648:
                        for fi16 := 0; fi16 < 18; fi16++ {
                            v2 = (-1 <= 42)
                            v17 := GenS2{F0: v2, F1: genFn0(), F2: fi15}
                            if true {
                                v5[(fi16 / ((fi15) | 1))] = 42
                                var v18 uint8 = 53
                                var v19 uint64 = 293
                                _ = v18
                                _ = v19
                            }
                            _ = fi16
                            _ = v17
                        }
                    case 32768:
                        var v20 int8 = 99
                        _ = v20
                    case -1:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    default:
                        if len(v3) > 0 {
                            v3[0] = ("a" + "")
                        }
                    }
                    v1 = v4
                    v21 := GenS3{F0: ("hello" + ("\n\t" + "\n\t")), F1: 18446744073709551615, F2: uint32(v1)}
                    _ = fi15
                    _ = v21
                }
            case "foo bar":
                for fi22 := 0; fi22 < 8; fi22++ {
                    for fi23 := 0; fi23 < 19; fi23++ {
                        switch int(v1) {
                        case 128:
                            var v24 uint8 = 1
                            _ = v24
                        default:
                            v4 = (^(v1))
                        }
                        v25 := GenS1{F0: (1 != -2147483648), F1: 0}
                        _ = fi23
                        _ = v25
                    }
                    fmt.Println(fi22, v1)
                    v26 := GenS3{F0: "\n\t", F1: 42, F2: 0}
                    _ = fi22
                    _ = v26
                }
            }
        default:
            v4 = 2
        }
        v1 = (v1 ^ v1)
    } else {
        var v27 uint64 = 1
        if len(v3) > 0 {
            v3[0] = ""
        }
        _ = v27
    }
    switch -1 {
    case 127:
        if (32768 != 127) {
            for fi28 := 0; fi28 < 6; fi28++ {
                if true {
                    v1 = (v4 ^ v4)
                    delete(v5, fi28)
                } else {
                    for fi29 := 0; fi29 < 2; fi29++ {
                        switch "\n\t" {
                        case "\n\t":
                            if v2 {
                                if v2 {
                                    var v30 float32 = (((1.5 * 1e10) * (-2.25 + ((1.5 + (float32(-2.25) - 3.14159)) + 0.0))) / 1.5)
                                    var v31 int = fi29
                                    _ = v30
                                    _ = v31
                                } else {
                                    v1 -= v4
                                }
                            } else {
                                fmt.Println(v2, v4)
                            }
                        }
                        var v32 float64 = -2.25
                        _ = fi29
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 4; fi33++ {
                        switch fi28 {
                        case 4294967295:
                            v34 := GenS4{F0: 255, F1: 42, F2: 0}
                            _ = v34
                        case -1:
                            if (!(41 >= 18446744073709551615)) {
                                fmt.Println(v2, fi33)
                            } else {
                                var v35 uint8 = 53
                                var v36 uint8 = v35
                                var v37 int8 = 1
                                if true {
                                    v37 += v37
                                }
                                _ = v35
                                _ = v36
                                _ = v37
                            }
                        }
                        switch "hello" {
                        case "a":
                            v38 := GenS5{F0: (v1 + (v4 * v4))}
                            switch (fi33 ^ (^(fi28))) {
                            case -1:
                                v2 = v2
                            case 174:
                                _ = len(v5)
                            case -988:
                                v4 = (v1 / (((v1 >> ((255) & 63))) | 1))
                            default:
                                fmt.Println(fi28)
                            }
                            delete(v5, 32767)
                            var v39 int16 = -1
                            _ = v38
                            _ = v39
                        case "":
                            fmt.Println(v4, v2)
                        default:
                            v40 := GenS3{F0: "hello", F1: 42, F2: 1}
                            _ = v40
                        }
                        _ = fi33
                    }
                    switch "" {
                    case "a":
                        v41 := GenS5{F0: (v1 * 42)}
                        _ = v41
                    }
                }
                v42 := GenS4{F0: 127, F1: 255, F2: 450}
                v43 := v5[(fi28 << ((255) & 63))]
                _ = fi28
                _ = v42
                _ = v43
            }
            for fi44 := 0; fi44 < 10; fi44++ {
                v4 -= (^(v1))
                _ = fi44
            }
            switch "hello" {
            case "foo bar":
                if v2 {
                    var v45 float32 = 3.14159
                    _ = v45
                } else {
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    v46 := GenS6{F0: genFn0(), F1: 630, F2: (true || ("foo bar" < ("" + "\n\t")))}
                    v47 := GenS7{F0: 0}
                    _ = v46
                    _ = v47
                }
            }
        }
    case 572:
        var v48 uint64 = 489
        _ = v48
    default:
        v2 = true
    }
    v49 := GenS3{F0: ("\n\t" + "\n\t"), F1: 42, F2: 4294967295}
    if v2 {
        for fi50 := 0; fi50 < 9; fi50++ {
            v51 := GenS5{F0: 2}
            var v52 float64 = (((0.1 / 0.1) / 1.5) / 0.1)
            fmt.Println(fi50, fi50)
            _ = fi50
            _ = v51
            _ = v52
        }
        var v53 uint16 = (^(v1))
        delete(v5, -1)
        _ = v53
    }
    v5[4294967295] = genFn0()
    var v54 float32 = 0.0
    v55 := GenS8{F0: 65535}
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v54
    _ = v55
}
