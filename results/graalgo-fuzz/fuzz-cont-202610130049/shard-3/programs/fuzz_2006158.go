package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int8; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS5 struct { F0 int8; F1 bool; F2 int8 }

func main() {
    v1 := []uint32{}
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    var v4 uint32 = 2147483648
    if false {
        var v5 uint64 = 31
        _ = v5
    } else {
        v1 = append(v1, (v4 >> ((255) & 63)))
        v4 -= 65535
        if false {
            if (42 < -383) {
                v4 += 226
            }
            var v6 int8 = 127
            _ = v6
        }
    }
    var v7 int = 127
    _ = len(v2)
    if false {
        var v8 float32 = float32(1e10)
        _ = v8
    }
    v4 -= (^(v4))
    v9 := GenS1{F0: 33}
    if false {
        var v10 int = (v7 + (v7 + -303))
        v4 -= (v4 ^ (^(v4)))
        v11 := GenS1{F0: 127}
        _ = v10
        _ = v11
    }
    v7 = v7
    v12 := GenS1{F0: 1}
    v7 -= 511
    for fi13 := 0; fi13 < 14; fi13++ {
        if (42 > 172) {
            for fi14 := 0; fi14 < 16; fi14++ {
                v7 += (v7 - 149)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v4 = v4
                var v15 uint8 = 23
                _ = fi14
                _ = v15
            }
        } else {
            v4 = v4
            if (276 != -1) {
                for fi16 := 0; fi16 < 10; fi16++ {
                    if (1.5 != 1.5) {
                        v17 := GenS1{F0: -58}
                        switch "foo bar" {
                        case "foo bar":
                            if (4294967295 == 227) {
                                if (-2.25 <= ((((((((((0.1 / (-2.25 * 3.14159)) / 1.5) - 0.0) + 0.0) * -2.25) + 1e10) / 1.0) + ((3.14159 * (((1.5 / (0.0 - 3.14159)) / float64(1e10)) / 0.1)) - 3.14159)) / 1.5) * 1e10)) {
                                    delete(v3, (fi16 ^ -2147483648))
                                }
                                v18 := GenS2{F0: 1, F1: 90, F2: (false || true)}
                                _ = v18
                            }
                        case "hello":
                            var v19 bool = true
                            _ = v19
                        case "\n\t":
                            v20 := GenS1{F0: 6}
                            _ = v20
                        default:
                            fmt.Println(v7, v4, fi13)
                        }
                        _ = v17
                    } else {
                        v4 = (v4 % ((v4) | 1))
                        if false {
                            v21 := GenS2{F0: (fi16 * fi13), F1: 75, F2: true}
                            _ = v21
                        }
                        v22 := GenS1{F0: -1}
                        _ = v22
                    }
                    _ = fi16
                }
            } else {
                v7 = (^(fi13))
            }
        }
        _ = fi13
    }
    if (true || (false && ("foo bar" < ("a" + "foo bar")))) {
        if false {
            for fi23 := 0; fi23 < 18; fi23++ {
                for fi24 := 0; fi24 < 12; fi24++ {
                    v7 -= (fi23 >> ((930) & 63))
                    switch -9223372036854775808 {
                    case -724:
                        var v25 int64 = 0
                        _ = v25
                    case 4294967295:
                        v7 = v7
                    }
                    _ = fi24
                }
                v4 -= 612
                _ = fi23
            }
        } else {
            if (false && ((0 != -32768) || true)) {
                switch "\n\t" {
                case "hello":
                    _ = len(v3)
                case "a":
                    if (true && false) {
                        if ((false || (((3.14159 < 1.5) || ((1 <= (v4 * (v4 + 65535))) && false)) || (!true))) && true) {
                            v4 = (v4 | v4)
                            v26 := v2["\n\t"]
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 18; fi27++ {
                            if true {
                                var v28 uint32 = v4
                                v29 := v3[(fi27 >> ((0) & 63))]
                                if (2 >= 2) {
                                    v28 -= (v28 % ((97) | 1))
                                    switch v29 {
                                    case 0:
                                        for fi30 := 0; fi30 < 20; fi30++ {
                                            v4 -= 557
                                            _ = fi30
                                        }
                                    case 175:
                                        var v31 uint32 = (^(v4))
                                        _ = v31
                                    }
                                } else {
                                    v7 = v7
                                    var v32 uint64 = 2
                                    _ = v32
                                }
                                _ = v28
                                _ = v29
                            } else {
                                v7 = 9223372036854775807
                            }
                            _ = fi27
                        }
                        var v33 int = 1
                        _ = v33
                    } else {
                        if (!(!true)) {
                            if (-137 >= -1) {
                                v7 -= (v7 & 127)
                                v34 := GenS3{F0: 88}
                                _ = v34
                            }
                            v35 := GenS3{F0: 255}
                            _ = v35
                        }
                    }
                default:
                    v36 := GenS1{F0: 0}
                    _ = v36
                }
            } else {
                fmt.Println(v7, v7)
                if ((1e10 + 3.14159) >= 0.1) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v37 := GenS4{F0: -644, F1: 42, F2: 255}
                    switch (v7 << ((600) & 63)) {
                    case 42:
                        var v38 int32 = 42
                        _ = v38
                    default:
                        var v39 string = "a"
                        _ = v39
                    }
                    _ = v37
                } else {
                    fmt.Println(v7)
                }
                v40 := GenS1{F0: 100}
                _ = v40
            }
        }
        for fi41 := 0; fi41 < 18; fi41++ {
            for fi42 := 0; fi42 < 16; fi42++ {
                v4 = v4
                v7 -= v7
                _ = fi42
            }
            if (v4 > v4) {
                var v43 int16 = 26
                v44 := GenS5{F0: 127, F1: true, F2: 0}
                _ = v43
                _ = v44
            } else {
                _ = len(v2)
                v4 += v4
                v7 -= fi41
            }
            _ = fi41
        }
        delete(v3, (v7 >> ((42) & 63)))
    } else {
        var v45 uint8 = 108
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
