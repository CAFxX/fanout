package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 int8; F1 int64 }

func main() {
    var v1 bool = (42 <= -9223372036854775808)
    v2 := make(map[uint32]uint32)
    var v3 int32 = 2147483647
    v4 := make(map[uint32]bool)
    v5 := []uint32{42, 127, 107}
    switch "foo bar" {
    case "foo bar":
        _ = len(v5)
    case "a":
        if (v1 && (65535 < 4294967295)) {
            v4[0] = (!v1)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if (1 < 1000) {
                for fi6 := 0; fi6 < 9; fi6++ {
                    var v7 uint32 = 0
                    _ = fi6
                    _ = v7
                }
                if true {
                    var v8 float64 = (1e10 / 1.0)
                    for fi9 := 0; fi9 < 4; fi9++ {
                        delete(v4, 42)
                        for fi10 := 0; fi10 < 12; fi10++ {
                            fmt.Println(v3, fi9, v1)
                            v11 := GenS1{F0: 204}
                            _ = fi10
                            _ = v11
                        }
                        var v12 float32 = (0.1 - (1e10 / 1.0))
                        _ = fi9
                        _ = v12
                    }
                    if ((42 >= 4294967295) && v1) {
                        var v13 float64 = 1e10
                        v8 += (((((v8 / v8) / (0.1 + 1.5)) - 0.0) - v13) - (v13 - (((((float64(3.14159) - (0.1 + v13)) * v8) - (-2.25 / v13)) * 1e10) - v13)))
                        _ = v13
                    }
                    _ = v8
                }
                var v14 uint64 = 344
                _ = v14
            }
        }
    case "hello":
        if (65535 > 1) {
            v3 = v3
            if false {
                if v1 {
                    v5 = append(v5, 347)
                    for fi15 := 0; fi15 < 5; fi15++ {
                        if v1 {
                            if (true && (465 != -1)) {
                                v1 = true
                                var v16 uint16 = 378
                                for fi17 := 0; fi17 < 1; fi17++ {
                                    v1 = true
                                    _ = fi17
                                }
                                _ = v16
                            }
                            fmt.Println(v1, v3, v3)
                            for fi18 := 0; fi18 < 11; fi18++ {
                                v1 = (!false)
                                _ = fi18
                            }
                        }
                        fmt.Println(v1)
                        v19 := GenS1{F0: 0}
                        _ = fi15
                        _ = v19
                    }
                }
            }
        } else {
            for fi20 := 0; fi20 < 2; fi20++ {
                fmt.Println(fi20, fi20)
                _ = fi20
            }
            v21 := GenS2{F0: 255, F1: v1}
            _ = v21
        }
    default:
        fmt.Println(v3, v1)
    }
    if len(v5) > 0 {
        v5[0] = 65535
    }
    if (v1 || (!true)) {
        var v22 int16 = -1
        _ = v22
    } else {
        switch ("a" + "\n\t") {
        case "\n\t":
            if v1 {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                fmt.Println(v3)
                v1 = (true && false)
            } else {
                v23 := GenS3{F0: 1, F1: 1000}
                _ = v23
            }
        case "a":
            var v24 uint64 = 4294967295
            _ = v24
        case "":
            if (42 == (v3 - 511)) {
                delete(v4, 467)
            } else {
                for fi25 := 0; fi25 < 19; fi25++ {
                    var v26 uint16 = 292
                    v5 = append(v5, 1)
                    v27 := GenS4{F0: (v26 < 32767), F1: -77}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
        default:
            if (float64((1.5 + 1.5)) >= 1e10) {
                for fi28 := 0; fi28 < 1; fi28++ {
                    for fi29 := 0; fi29 < 14; fi29++ {
                        v30 := GenS1{F0: 1}
                        v3 -= (v3 << ((42) & 63))
                        v3 = (v3 & (v3 << ((401) & 63)))
                        if (0 >= -1) {
                            var v31 int16 = 179
                            _ = v31
                        }
                        _ = fi29
                        _ = v30
                    }
                    _ = fi28
                }
                if v1 {
                    if v1 {
                        switch ("foo bar" + "") {
                        case "hello":
                            if (!((^(v3)) <= -2147483648)) {
                                if (-670 >= 194) {
                                    var v32 int32 = v3
                                    var v33 string = "\n\t"
                                    _ = v32
                                    _ = v33
                                }
                                if v1 {
                                    var v34 string = "hello"
                                    fmt.Println(v3, v1, v34)
                                    _ = v34
                                } else {
                                    fmt.Println(v3, v1)
                                }
                            }
                        case "foo bar":
                            v1 = (4294967295 != 42)
                        case "\n\t":
                            fmt.Println(v1)
                        }
                    }
                    v35 := v4[2147483647]
                    v36 := GenS1{F0: 0}
                    _ = v35
                    _ = v36
                }
            } else {
                v37 := GenS5{F0: -73, F1: -1}
                fmt.Println(v1)
                fmt.Println(v3, v3)
                _ = v37
            }
        }
        var v38 int = -9223372036854775808
        fmt.Println(v1, v38)
        _ = v38
    }
    if (false && (-9223372036854775808 >= 9223372036854775807)) {
        v39 := GenS1{F0: 42}
        fmt.Println(v3, v1)
        _ = v39
    } else {
        v1 = v1
        var v40 float32 = 1.5
        _ = v40
    }
    if ((!false) || v1) {
        var v41 bool = false
        var v42 int64 = 1
        _ = v41
        _ = v42
    } else {
        var v43 uint = 0
        _ = v43
    }
    fmt.Println(v1, v1)
    var v44 uint8 = 0
    var v45 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
}
