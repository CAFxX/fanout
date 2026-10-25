package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 int32; F2 uint32 }
type GenS5 struct { F0 uint8; F1 int16; F2 uint }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 int8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint16; F1 string }

func main() {
    var v1 float32 = 3.14159
    v2 := []uint32{}
    v3 := make(map[int]string)
    var v4 int32 = 1
    var v5 uint32 = 4294967295
    var v6 uint32 = (v5 * (v5 >> ((0) & 63)))
    if (65535 != 2) {
        var v7 int = 32768
        switch (v7 / (((v7 >> ((929) & 63))) | 1)) {
        case 32768:
            if true {
                v5 -= v5
                fmt.Println(v1)
                if len(v2) > 0 {
                    v2[0] = (v6 + v5)
                }
            } else {
                for fi8 := 0; fi8 < 1; fi8++ {
                    v1 = (v1 * v1)
                    _ = fi8
                }
                if false {
                    var v9 float64 = (1e10 + 1.5)
                    fmt.Println(v4, v4)
                    v10 := GenS1{F0: 42, F1: v7}
                    _ = v10
                    _ = v9
                }
                if (!(true && (false || (1 == 0)))) {
                    fmt.Println(v4, v4)
                    if true {
                        for fi11 := 0; fi11 < 1; fi11++ {
                            v12 := GenS1{F0: -1, F1: (v7 * (v7 & 554))}
                            v13 := GenS2{F0: 0, F1: v6, F2: 426}
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    }
                    v5 -= (v6 / ((v6) | 1))
                }
            }
        case 0:
            v14 := GenS3{F0: 0}
            _ = v14
        default:
            v15 := GenS3{F0: 42}
            _ = v15
        }
        _ = v7
    }
    v16 := GenS4{F0: -837, F1: v4, F2: (v6 & 411)}
    v1 = ((v1 * ((((v1 + v1) + ((3.14159 / v1) * v1)) / float32(float64(v1))) * 1.5)) * v1)
    var v17 int32 = 1
    v4 = v17
    var v18 float32 = (v1 / 1.5)
    var v19 int32 = (-(v17))
    v20 := GenS5{F0: 207, F1: 649, F2: 255}
    fmt.Println(v19, v1, v17)
    if (343 != -1) {
        fmt.Println(v18, v1)
        v18 -= (float32(1.5) * 0.0)
    } else {
        var v21 int64 = -1
        _ = v21
    }
    v19 -= v4
    fmt.Println(v1, v6, v1)
    if ((!false) || (("" + "hello") < (("hello" + "") + ("hello" + "a")))) {
        v6 += 65535
        var v22 uint8 = 137
        _ = v22
    } else {
        v23 := GenS6{F0: 32767, F1: (v19 - (v19 / ((v4) | 1)))}
        if (true && (!(0.1 > (-2.25 + (0.0 * (((0.0 / 3.14159) * 0.0) + 3.14159)))))) {
            v17 = (v4 + v17)
            v24 := GenS7{F0: 42}
            for fi25 := 0; fi25 < 9; fi25++ {
                if true {
                    v2 = append(v2, 4294967295)
                    var v26 int64 = -1
                    _ = v26
                }
                _ = fi25
            }
            _ = v24
        } else {
            fmt.Println(v6)
        }
        var v27 uint8 = 42
        _ = v23
        _ = v27
    }
    v28 := GenS7{F0: 4294967295}
    switch "a" {
    case "":
        v29 := GenS8{F0: 2147483647, F1: 127}
        _ = v29
    default:
        if false {
            if false {
                if (!(42 < 32767)) {
                    v3[9223372036854775807] = ("hello" + "")
                }
                if (true || true) {
                    v3[4294967295] = "\n\t"
                    for fi30 := 0; fi30 < 15; fi30++ {
                        for fi31 := 0; fi31 < 15; fi31++ {
                            var v32 int8 = -103
                            if (18446744073709551615 > 42) {
                                if (v32 == 42) {
                                    v2 = append(v2, (v6 << ((587) & 63)))
                                    fmt.Println(v4, v18)
                                }
                            } else {
                                switch (fi30 * 32767) {
                                case 42:
                                    var v33 string = "foo bar"
                                    _ = v33
                                case 2147483647:
                                    if (-891 < 0) {
                                        fmt.Println(fi30)
                                    }
                                case -9223372036854775808:
                                    v34 := GenS3{F0: 128}
                                    _ = v34
                                }
                            }
                            _ = fi31
                            _ = v32
                        }
                        if (!false) {
                            for fi35 := 0; fi35 < 8; fi35++ {
                                var v36 int = (fi35 & fi30)
                                var v37 uint16 = 0
                                v5 += v6
                                v5 = (v6 - v6)
                                var v38 float32 = v18
                                _ = fi35
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                        }
                        var v39 uint16 = 42
                        _ = fi30
                        _ = v39
                    }
                } else {
                    var v40 uint16 = 262
                    switch "" {
                    case "hello":
                        v17 = v19
                    case "foo bar":
                        var v41 int64 = -101
                        _ = v41
                    case "":
                        v42 := GenS4{F0: 1, F1: (v19 % ((v19) | 1)), F2: 4294967295}
                        _ = v42
                    default:
                        for fi43 := 0; fi43 < 9; fi43++ {
                            if true {
                                var v44 uint8 = 193
                                _ = v44
                            } else {
                                var v45 int8 = 42
                                var v46 bool = ("hello" < "\n\t")
                                v4 = 1
                                _ = v45
                                _ = v46
                            }
                            _ = len(v3)
                            fmt.Println(v17)
                            _ = fi43
                        }
                    }
                    _ = v40
                }
                switch 4294967295 {
                case -121:
                    for fi47 := 0; fi47 < 5; fi47++ {
                        var v48 uint = 255
                        switch "a" {
                        case "hello":
                            var v49 int16 = 0
                            _ = v49
                        case "":
                            for fi50 := 0; fi50 < 3; fi50++ {
                                var v51 uint16 = 641
                                switch fi50 {
                                case 0:
                                    var v52 uint16 = v51
                                    _ = v52
                                case 32767:
                                    v53 := GenS9{F0: false}
                                    _ = v53
                                default:
                                    _ = len(v2)
                                }
                                if true {
                                    v54 := GenS10{F0: -1}
                                    var v55 int64 = 492
                                    _ = v54
                                    _ = v55
                                } else {
                                    var v56 uint32 = 1
                                    v5 = (v5 >> (((v48 * (v48 - 255))) & 63))
                                    switch "hello" {
                                    case "foo bar":
                                        var v57 uint = 255
                                        _ = v57
                                    case "hello":
                                        var v58 uint8 = 255
                                        _ = v58
                                    case "a":
                                        var v59 string = "a"
                                        _ = v59
                                    default:
                                        v60 := GenS11{F0: v51, F1: (((("hello" + ("" + "\n\t")) + "\n\t") + "hello") + "foo bar")}
                                        _ = v60
                                    }
                                    _ = v56
                                }
                                v17 = 0
                                _ = fi50
                                _ = v51
                            }
                        }
                        switch 128 {
                        case -1:
                            v2 = append(v2, 42)
                        }
                        _ = fi47
                        _ = v48
                    }
                }
            }
            var v61 int = 128
            _ = v61
        }
    }
    v18 += v18
    var v62 int = 1
    var v63 bool = (!(true && ((-110 < 12) && ((1 > 2) || false))))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v63
}
