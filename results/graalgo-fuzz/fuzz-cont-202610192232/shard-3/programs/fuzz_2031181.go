package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 uint64; F1 string; F2 int32 }
type GenS7 struct { F0 int8 }

func main() {
    var v1 string = "\n\t"
    var v2 float32 = 3.14159
    var v3 string = v1
    v4 := []int{32768}
    var v5 float32 = v2
    var v6 float32 = float32((0.0 - (0.1 / 3.14159)))
    switch 682 {
    case 0:
        if true {
            if (!(true || false)) {
                var v7 float64 = (-2.25 * 0.1)
                _ = v7
            }
            v8 := GenS1{F0: 2, F1: 0, F2: 0}
            _ = len(v4)
            _ = v8
        } else {
            _ = len(v4)
            v9 := GenS2{F0: true}
            for fi10 := 0; fi10 < 3; fi10++ {
                v11 := GenS1{F0: 373, F1: 42, F2: (fi10 - (fi10 * -222))}
                v4 = append(v4, fi10)
                if (4294967295 != 1) {
                    var v12 int16 = 42
                    v4 = append(v4, fi10)
                    switch fi10 {
                    case 2147483647:
                        if ((!(((0 > 969) || (!true)) && (false && (((v6 * v6) == (v5 * v2)) || true)))) && true) {
                            if len(v4) > 0 {
                                v4[0] = (fi10 % (((fi10 + fi10)) | 1))
                            }
                            var v13 int32 = 563
                            _ = v13
                        } else {
                            var v14 uint8 = 255
                            v5 = (v2 * ((0.1 + v2) / -2.25))
                            _ = v14
                        }
                    case -2147483648:
                        v15 := GenS1{F0: 401, F1: 519, F2: fi10}
                        _ = v15
                    default:
                        v16 := GenS1{F0: 764, F1: 0, F2: 1}
                        _ = v16
                    }
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
            _ = v9
        }
    default:
        fmt.Println(v3, v6, v6)
    }
    var v17 int32 = -2147483648
    if (0 >= 745) {
        if (!true) {
            fmt.Println(v6, v1)
            var v18 uint = 460
            v19 := GenS1{F0: 0, F1: 0, F2: 908}
            _ = v18
            _ = v19
        } else {
            v6 += (v6 / (float32(0.1) * (((((v5 / 0.1) / v2) * v2) / 1.5) * (0.1 - v2))))
        }
        for fi20 := 0; fi20 < 8; fi20++ {
            v21 := GenS3{F0: 127}
            _ = fi20
            _ = v21
        }
        v22 := GenS4{F0: -1}
        _ = v22
    } else {
        fmt.Println(v3, v6, v6)
    }
    v5 -= v2
    v23 := GenS1{F0: 866, F1: 2, F2: 394}
    v5 = (1.5 + (0.1 / 1e10))
    if (false && (123 <= 1000)) {
        v4 = append(v4, -9223372036854775808)
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi25 := 0; fi25 < 18; fi25++ {
            v6 += (((v5 * float32(3.14159)) * v6) - v5)
            v17 = (v17 >> ((405) & 63))
            _ = fi25
        }
        _ = fi24
    }
    _ = len(v4)
    for fi26 := 0; fi26 < 17; fi26++ {
        fmt.Println(v3, v6, v1)
        _ = fi26
    }
    if (true || (true && (2 != 1))) {
        var v27 string = v3
        for fi28 := 0; fi28 < 20; fi28++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if false {
                v4 = append(v4, (fi28 << ((4294967295) & 63)))
                v4 = append(v4, (-(fi28)))
                switch v3 {
                case "\n\t":
                    var v29 int32 = 42
                    _ = v29
                case "hello":
                    var v30 int = fi28
                    _ = v30
                }
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi28
        }
        if (true || (!false)) {
            var v31 float32 = ((v5 * v5) + v6)
            v32 := GenS5{F0: 4294967295, F1: 42}
            if (-128 > 52) {
                var v33 int32 = -280
                var v34 int64 = 0
                v17 = v33
                _ = v33
                _ = v34
            }
            _ = v31
            _ = v32
        } else {
            _ = len(v4)
            switch v1 {
            case "hello":
                switch v1 {
                case "foo bar":
                    v35 := GenS6{F0: 18446744073709551615, F1: ("" + v3), F2: v17}
                    _ = v35
                case "a":
                    fmt.Println(v17, v3, v17)
                case "hello":
                    switch -953 {
                    case -904:
                        var v36 int8 = 0
                        _ = v36
                    case 43:
                        if true {
                            v5 += (v6 / ((((v2 + ((0.0 - (-2.25 - 1e10)) * v2)) * 3.14159) * v6) + float32(0.1)))
                            if (-112 != 100) {
                                v6 = 3.14159
                            } else {
                                v37 := GenS7{F0: 86}
                                fmt.Println(v6)
                                v38 := GenS1{F0: 139, F1: 2, F2: 0}
                                _ = v37
                                _ = v38
                            }
                        }
                    case -9223372036854775808:
                        v17 = (v17 + 2147483647)
                    }
                default:
                    for fi39 := 0; fi39 < 8; fi39++ {
                        v4 = append(v4, (fi39 & 894))
                        _ = fi39
                    }
                }
            }
            v40 := GenS5{F0: 379, F1: 1}
            _ = v40
        }
        _ = v27
    }
    var v41 uint8 = 217
    v17 -= (-(v17))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
