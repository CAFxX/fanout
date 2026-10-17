package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 bool }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int8 = -128
    var v2 int32 = 2147483647
    v3 := []string{"a", ("a" + ""), "a", ("hello" + "")}
    var v4 int = 32768
    var v5 uint32 = 0
    var v6 float64 = float64((0.1 * 1e10))
    var v7 int32 = v2
    v8 := GenS1{F0: 2, F1: v4, F2: v1}
    for fi9 := 0; fi9 < 4; fi9++ {
        v6 = float64(1e10)
        v3 = append(v3, "hello")
        switch fi9 {
        case -2147483648:
            v4 = fi9
        }
        _ = fi9
    }
    var v10 uint64 = 42
    v11 := GenS1{F0: 127, F1: (^(v4)), F2: v1}
    fmt.Println(v5, v4)
    for fi12 := 0; fi12 < 20; fi12++ {
        switch v4 {
        case -316:
            v10 -= 320
        default:
            v4 += 2147483647
        }
        _ = fi12
    }
    var v13 int64 = 11
    for fi14 := 0; fi14 < 20; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            v13 -= (v13 + v13)
            fmt.Println(v2, v1, v1)
            _ = fi15
        }
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        _ = fi14
    }
    if (false && (true || ((!true) && false))) {
        v16 := GenS1{F0: 255, F1: v4, F2: -53}
        var v17 uint = 42
        _ = v16
        _ = v17
    } else {
        switch "" {
        case "":
            v1 += (v1 << ((956) & 63))
        case "foo bar":
            switch "hello" {
            case "hello":
                _ = len(v3)
            case "a":
                _ = len(v3)
            }
        default:
            if (((485 <= 65535) || ((11 == 42) || true)) || false) {
                for fi18 := 0; fi18 < 19; fi18++ {
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v5 = (v5 - v5)
                        var v20 uint64 = v10
                        if ("hello" <= "a") {
                            var v21 int16 = 0
                            _ = v21
                        }
                        if len(v3) > 0 {
                            v3[0] = "hello"
                        }
                        fmt.Println(v6)
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
                for fi22 := 0; fi22 < 15; fi22++ {
                    v6 = (v6 - v6)
                    _ = fi22
                }
                var v23 float64 = (1.5 + v6)
                _ = v23
            } else {
                v2 += v7
                var v24 float64 = (1.5 / (v6 * (v6 * 1e10)))
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 16; fi25++ {
            if true {
                v7 -= (v7 >> ((uint(v10)) & 63))
                fmt.Println(v4)
                v2 += 1
            }
            switch ("a" + "hello") {
            case "hello":
                if true {
                    var v26 bool = false
                    _ = v26
                } else {
                    switch (v4 & -113) {
                    case 4294967295:
                        for fi27 := 0; fi27 < 14; fi27++ {
                            if (-2.25 <= -2.25) {
                                v28 := GenS1{F0: 255, F1: 14, F2: (v1 << ((0) & 63))}
                                var v29 int32 = (v2 + v2)
                                _ = v28
                                _ = v29
                            } else {
                                if false {
                                    v30 := GenS1{F0: 2, F1: (fi27 << ((844) & 63)), F2: v1}
                                    fmt.Println(v10)
                                    var v31 uint64 = (v10 ^ 42)
                                    _ = v30
                                    _ = v31
                                }
                            }
                            var v32 int16 = 1000
                            var v33 uint32 = (v5 * 290)
                            _ = fi27
                            _ = v32
                            _ = v33
                        }
                    case 128:
                        fmt.Println(v2)
                    case 42:
                        for fi34 := 0; fi34 < 6; fi34++ {
                            v1 += -128
                            _ = fi34
                        }
                    default:
                        v4 = (fi25 >> ((4294967295) & 63))
                    }
                }
            case "foo bar":
                switch fi25 {
                case 4294967295:
                    for fi35 := 0; fi35 < 4; fi35++ {
                        var v36 uint16 = 0
                        v2 -= (v2 << ((18446744073709551615) & 63))
                        if (((1.5 + float64(3.14159)) - v6) != 3.14159) {
                            if ((false || (95 >= 2)) && ((0.0 * (0.1 + ((float32((((3.14159 * ((v6 - v6) * float64(0.0))) * (float64(-2.25) - (float64(-2.25) * v6))) / (3.14159 + v6))) - (3.14159 + 3.14159)) / (0.0 - -2.25)))) > 0.0)) {
                                var v37 uint16 = (^(v36))
                                if ("a" >= "") {
                                    v6 += -2.25
                                    v7 -= v2
                                }
                                fmt.Println(fi35, v5, v4)
                                v4 -= (fi25 & int(v37))
                                _ = v37
                            }
                            v38 := GenS1{F0: 31, F1: (fi25 << ((0) & 63)), F2: v1}
                            _ = v38
                        }
                        var v39 uint32 = 715
                        v1 = 70
                        _ = fi35
                        _ = v36
                        _ = v39
                    }
                case 1:
                    fmt.Println(v5, v4, v10)
                case 138:
                    v40 := GenS2{F0: uint64(v5), F1: 0, F2: (false || true)}
                    _ = v40
                default:
                    _ = len(v3)
                }
            case "":
                v41 := GenS3{F0: 846}
                _ = v41
            default:
                var v42 float64 = v6
                _ = v42
            }
            var v43 float64 = ((v6 / 0.1) / (((((float64(float32(v6)) * 1e10) / ((v6 * v6) * (v6 - v6))) * 0.0) / ((v6 - v6) - (-2.25 / ((v6 - 3.14159) + 0.1)))) - 3.14159))
            _ = fi25
            _ = v43
        }
        fmt.Println(v4, v13, v7)
    }
    var v44 int32 = v7
    if (false || ((-(v4)) >= (v4 & (v4 / ((v4) | 1))))) {
        v45 := GenS1{F0: 187, F1: v4, F2: (^(v1))}
        fmt.Println(v6)
        fmt.Println(v2, v44)
        _ = v45
    }
    fmt.Println(v1)
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
