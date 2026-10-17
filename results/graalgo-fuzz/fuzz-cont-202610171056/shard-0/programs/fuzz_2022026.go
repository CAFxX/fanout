package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int64 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int64; F2 int }
type GenS4 struct { F0 int8; F1 int8; F2 int32 }

func main() {
    var v1 float32 = ((0.1 - float32(3.14159)) * 1.5)
    var v2 float64 = float64(((((((0.0 / v1) * -2.25) - ((1e10 * v1) * (v1 / float32((0.1 - -2.25))))) + ((1e10 / float32(((0.0 + 1.5) / 3.14159))) * 1e10)) - 1e10) * ((v1 * 0.1) + ((1e10 / v1) / v1))))
    var v3 int32 = -2147483648
    var v4 int8 = 31
    v5 := []uint32{1, 2147483647, 892}
    var v6 bool = ((-32768 == -1) && false)
    switch "foo bar" {
    case "":
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 uint = 4294967295
            v2 -= (float64(-2.25) + v2)
            _ = fi7
            _ = v8
        }
    case "foo bar":
        v6 = true
    case "\n\t":
        if false {
            v9 := GenS1{F0: 1, F1: -1, F2: -673}
            fmt.Println(v2, v1)
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 int16 = -20
                var v12 int8 = v4
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v9
        }
    }
    if (!(42 < -9223372036854775808)) {
        if v6 {
            v6 = (42 < 4294967295)
        } else {
            var v13 string = ""
            if (-587 >= 9223372036854775807) {
                switch v13 {
                case "foo bar":
                    switch ("foo bar" + v13) {
                    case "hello":
                        var v14 int = -107
                        _ = v14
                    case "\n\t":
                        var v15 bool = (!false)
                        _ = v15
                    case "a":
                        if v6 {
                            v16 := GenS2{F0: 42, F1: 1, F2: 551}
                            v17 := GenS2{F0: -32768, F1: 255, F2: 233}
                            fmt.Println(v4)
                            _ = v16
                            _ = v17
                        }
                    default:
                        v2 += v2
                    }
                default:
                    var v18 uint32 = 2
                    _ = v18
                }
                var v19 uint = 0
                if v6 {
                    if true {
                        v2 -= (float64(float32(v2)) + float64(v1))
                    }
                    for fi20 := 0; fi20 < 6; fi20++ {
                        v21 := GenS2{F0: -1, F1: 1, F2: 0}
                        var v22 uint = (^(v19))
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                }
                _ = v19
            }
            _ = v13
        }
        if v6 {
            _ = len(v5)
            if v6 {
                switch ((("" + "hello") + "foo bar") + "") {
                case "hello":
                    v6 = ((true || (601 >= 32767)) || v6)
                case "":
                    var v23 int8 = v4
                    _ = v23
                case "foo bar":
                    switch 32768 {
                    case 2147483647:
                        v24 := GenS3{F0: 255, F1: -8, F2: 950}
                        _ = v24
                    }
                default:
                    v3 = (v3 | v3)
                }
                if v6 {
                    if false {
                        for fi25 := 0; fi25 < 4; fi25++ {
                            switch "a" {
                            case "":
                                var v26 uint16 = 1
                                _ = v26
                            case "\n\t":
                                if (936 != 255) {
                                    v3 -= (^(v3))
                                    var v27 bool = v6
                                    _ = v27
                                }
                            case "foo bar":
                                var v28 uint8 = 0
                                _ = v28
                            }
                            for fi29 := 0; fi29 < 9; fi29++ {
                                v3 = v3
                                v30 := GenS4{F0: (v4 + (v4 ^ v4)), F1: 1, F2: (v3 * -989)}
                                _ = fi29
                                _ = v30
                            }
                            _ = fi25
                        }
                    }
                }
            } else {
                v4 = (-(v4))
            }
        }
    } else {
        v3 = (v3 << ((4294967295) & 63))
    }
    var v31 uint16 = 2
    v1 += v1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v31 = v31
    v32 := GenS3{F0: 127, F1: -278, F2: 32768}
    for fi33 := 0; fi33 < 11; fi33++ {
        v2 += (v2 - -2.25)
        _ = fi33
    }
    if (false && (!(true && (-2147483648 > (v3 * (^(v3))))))) {
        var v34 uint64 = 2
        _ = v34
    } else {
        if len(v5) > 0 {
            v5[0] = 2147483648
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(v2)
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        fmt.Println(v4, v6, v3)
        for fi36 := 0; fi36 < 2; fi36++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if (!("a" <= ("foo bar" + ("" + ("foo bar" + "a"))))) {
                switch 32768 {
                case -1:
                    var v37 string = "\n\t"
                    _ = v37
                case -2147483648:
                    if v6 {
                        v1 -= v1
                    } else {
                        var v38 int = fi36
                        v31 -= v31
                        _ = v38
                    }
                case 4294967295:
                    v2 = 0.0
                default:
                    v4 -= v4
                }
                var v39 int16 = 801
                _ = v39
            }
            _ = fi36
        }
        _ = fi35
    }
    var v40 int8 = v4
    v40 = 42
    fmt.Println(v2, v40, v3)
    v40 -= (v4 % (((v4 / (((v4 / (((v4 | (-(v40)))) | 1))) | 1))) | 1))
    var v41 int = 1
    v41 -= int(v40)
    v6 = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
