package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int64; F1 int8; F2 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 int8; F2 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 uint8 }

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]int)
    var v3 bool = (19 == 1)
    var v4 int64 = 1
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 10; fi6++ {
            if false {
                v2[40] = (fi6 ^ 1)
                _ = len(v2)
            }
            fmt.Println(v1, fi5, v3)
            switch (fi5 % ((fi6) | 1)) {
            case 128:
                v7 := GenS1{F0: 0}
                _ = v7
            }
            _ = fi6
        }
        v8 := GenS2{F0: 369, F1: v3}
        _ = fi5
        _ = v8
    }
    v1 = (v1 % ((93) | 1))
    v4 = -1
    var v9 int16 = 1000
    if false {
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := GenS1{F0: v1}
            v12 := GenS2{F0: 18446744073709551615, F1: ((!false) || v3)}
            _ = fi10
            _ = v11
            _ = v12
        }
        fmt.Println(v9)
        v13 := v2[764]
        _ = v13
    } else {
        if v3 {
            v14 := []int{127, 42, 0, -259}
            for fi15 := 0; fi15 < 19; fi15++ {
                v1 = v1
                var v16 float64 = 1e10
                _ = fi15
                _ = v16
            }
            if (true || (true || ("foo bar" != ("hello" + "foo bar")))) {
                v3 = false
                if false {
                    fmt.Println(v9, v3)
                }
            } else {
                v3 = (!(((183 >= 42) || (("foo bar" + "foo bar") < "hello")) || v3))
                switch 32767 {
                case 186:
                    v3 = v3
                case 42:
                    v1 = (v1 % ((v1) | 1))
                case 984:
                    switch -341 {
                    case -1:
                        var v17 int64 = -319
                        _ = v17
                    case 1:
                        switch ("" + ("\n\t" + ("\n\t" + ("\n\t" + ("" + "hello"))))) {
                        case "a":
                            var v18 float32 = 1e10
                            _ = v18
                        default:
                            for fi19 := 0; fi19 < 14; fi19++ {
                                if (128 != 127) {
                                    v20 := GenS3{F0: (^(v4)), F1: v1, F2: 1}
                                    if (65535 == 2147483648) {
                                        fmt.Println(v3, v3, v3)
                                    }
                                    if true {
                                        var v21 float32 = ((0.1 / -2.25) * 0.1)
                                        switch "foo bar" {
                                        case "a":
                                            v3 = v3
                                        }
                                        var v22 bool = v3
                                        _ = v21
                                        _ = v22
                                    } else {
                                        if ((1.5 - -2.25) != (-2.25 - (1.5 + 0.0))) {
                                            fmt.Println(v9)
                                            for fi23 := 0; fi23 < 16; fi23++ {
                                                v24 := GenS2{F0: 255, F1: false}
                                                v25 := GenS1{F0: v1}
                                                _ = fi23
                                                _ = v24
                                                _ = v25
                                            }
                                        }
                                    }
                                    _ = v20
                                }
                                if (0 == 934) {
                                    var v26 int16 = v9
                                    v4 = -9223372036854775808
                                    if (199 >= 549) {
                                        var v27 float32 = 1.5
                                        if ((v1 == 42) && (!v3)) {
                                            v9 += (v9 | v9)
                                            switch fi19 {
                                            case 0:
                                                var v28 int32 = -2147483648
                                                _ = v28
                                            default:
                                                switch (-(fi19)) {
                                                case -2147483648:
                                                    v2[2147483648] = fi19
                                                case -687:
                                                    v29 := GenS4{F0: (v4 >> ((42) & 63))}
                                                    _ = v29
                                                case -350:
                                                    v1 = v1
                                                }
                                            }
                                            for fi30 := 0; fi30 < 4; fi30++ {
                                                v31 := GenS3{F0: (v4 ^ (-(v4))), F1: (^(v1)), F2: -81}
                                                _ = fi30
                                                _ = v31
                                            }
                                        }
                                        delete(v2, 4294967295)
                                        _ = v27
                                    }
                                    _ = v26
                                }
                                _ = fi19
                            }
                        }
                    case 128:
                        switch -1 {
                        case 958:
                            for fi32 := 0; fi32 < 15; fi32++ {
                                v33 := GenS5{F0: 1, F1: 84, F2: 0}
                                _ = fi32
                                _ = v33
                            }
                        case -2147483648:
                            v3 = (!true)
                        }
                    }
                }
                v9 -= 264
            }
            _ = v14
        } else {
            var v34 float64 = 0.0
            v34 = 0.0
            _ = v34
        }
        switch -9223372036854775808 {
        case 2147483647:
            v4 = v4
        case 0:
            for fi35 := 0; fi35 < 12; fi35++ {
                if v3 {
                    switch (fi35 << ((18446744073709551615) & 63)) {
                    case 32768:
                        fmt.Println(fi35, v3)
                    }
                    var v36 uint = 36
                    _ = v36
                }
                v4 -= (v4 >> ((65535) & 63))
                _ = fi35
            }
        default:
            var v37 string = ("" + "hello")
            _ = v37
        }
    }
    v9 = v9
    var v38 uint16 = 0
    v39 := []int{int(v4), 128}
    if ((v1 | 42) == (v1 + v1)) {
        v2[929] = 32768
        v40 := GenS6{F0: 0}
        for fi41 := 0; fi41 < 14; fi41++ {
            v42 := GenS2{F0: 65535, F1: ((v38 << ((255) & 63)) == (v38 << ((0) & 63)))}
            _ = fi41
            _ = v42
        }
        _ = v40
    }
    switch 4294967295 {
    case -680:
        for fi43 := 0; fi43 < 2; fi43++ {
            v9 += (v9 >> ((0) & 63))
            if v3 {
                var v44 uint = 1
                v2[65535] = (fi43 ^ 2147483647)
                var v45 float32 = float32((1e10 - 0.1))
                _ = v44
                _ = v45
            }
            _ = fi43
        }
    }
    v1 += -1
    var v46 uint8 = 238
    v38 = (^(v38))
    _ = len(v2)
    v9 -= v9
    if (false && false) {
        for fi47 := 0; fi47 < 11; fi47++ {
            v48 := GenS7{F0: (float32(((1e10 + 0.0) * (1.5 / 1.0))) == 3.14159), F1: 255}
            v46 = (v46 * 1)
            var v49 uint = 211
            _ = fi47
            _ = v48
            _ = v49
        }
    } else {
        var v50 uint64 = 4294967295
        _ = v50
    }
    if len(v39) > 0 {
        v39[0] = 128
    }
    fmt.Println(v9, v9)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v46
    _ = v9
}
