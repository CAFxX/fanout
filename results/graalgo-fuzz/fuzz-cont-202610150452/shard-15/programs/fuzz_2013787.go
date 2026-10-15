package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 bool; F1 bool; F2 int16 }
type GenS4 struct { F0 uint8; F1 int16; F2 uint16 }

func main() {
    var v1 string = ("\n\t" + ("a" + ("" + "")))
    v2 := []int{}
    v3 := []int{32768, -1, -298, 9223372036854775807}
    var v4 int32 = -53
    if ((!true) || false) {
        switch -76 {
        case 128:
            var v5 float64 = (0.1 + 0.0)
            _ = v5
        case -469:
            for fi6 := 0; fi6 < 5; fi6++ {
                var v7 uint64 = 964
                fmt.Println(fi6, v7)
                _ = fi6
                _ = v7
            }
        case -706:
            var v8 uint = 4294967295
            _ = v8
        }
        switch v1 {
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "\n\t":
            switch v1 {
            case "hello":
                if (65535 <= 105) {
                    var v9 uint = 1
                    fmt.Println(v9, v9)
                    _ = v9
                }
            case "\n\t":
                var v10 int64 = 28
                _ = v10
            case "":
                for fi11 := 0; fi11 < 18; fi11++ {
                    if false {
                        if true {
                            for fi12 := 0; fi12 < 3; fi12++ {
                                fmt.Println(v4, v4)
                                if true {
                                    if len(v3) > 0 {
                                        v3[0] = (fi11 + 1)
                                    }
                                    v1 = "a"
                                    if true {
                                        if len(v2) > 0 {
                                            v2[0] = fi12
                                        }
                                        v1 = v1
                                    }
                                    v1 = v1
                                    if len(v2) > 0 {
                                        v2[0] = fi12
                                    }
                                } else {
                                    v3 = append(v3, 722)
                                    var v13 int8 = 100
                                    _ = v13
                                }
                                v1 = "a"
                                _ = fi12
                            }
                        } else {
                            v14 := GenS1{F0: 100, F1: 255}
                            _ = v14
                        }
                    } else {
                        if (-94 > 127) {
                            var v15 float64 = float64(0.1)
                            for fi16 := 0; fi16 < 16; fi16++ {
                                var v17 uint8 = 2
                                _ = len(v2)
                                _ = fi16
                                _ = v17
                            }
                            _ = v15
                        } else {
                            var v18 uint16 = 2
                            _ = v18
                        }
                        v1 = v1
                        if true {
                            v1 = ("hello" + ((((v1 + "foo bar") + ("a" + ("\n\t" + v1))) + v1) + ("a" + "")))
                            fmt.Println(fi11)
                        }
                    }
                    for fi19 := 0; fi19 < 15; fi19++ {
                        v20 := GenS2{F0: -32768, F1: 128}
                        if (130 > 127) {
                            if (!(false && false)) {
                                v2 = append(v2, 353)
                                var v21 uint64 = 42
                                fmt.Println(fi19)
                                _ = v21
                            }
                            if (-2.25 < float32(1e10)) {
                                v4 = (v4 * (v4 + -1))
                                var v22 string = "hello"
                                _ = v22
                            } else {
                                fmt.Println(fi11, v1)
                                _ = len(v3)
                                v23 := GenS1{F0: 127, F1: 1}
                                _ = v23
                            }
                        } else {
                            v4 = (v4 >> ((1) & 63))
                            v1 = v1
                        }
                        var v24 int8 = 0
                        _ = fi19
                        _ = v20
                        _ = v24
                    }
                    _ = fi11
                }
            default:
                for fi25 := 0; fi25 < 11; fi25++ {
                    v1 = v1
                    if (v1 >= v1) {
                        var v26 string = "a"
                        _ = v26
                    }
                    if (true && ((1.5 * ((float64((1.5 / (1e10 - 3.14159))) * (-2.25 * 1.5)) - 0.1)) <= 0.0)) {
                        v4 += 2147483647
                    }
                    _ = fi25
                }
            }
        case "hello":
            for fi27 := 0; fi27 < 3; fi27++ {
                var v28 int8 = 100
                _ = fi27
                _ = v28
            }
        default:
            v4 -= -34
        }
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        v4 += (v4 | (v4 % (((v4 - (v4 - v4))) | 1)))
        v4 = (v4 | 2147483647)
        v3 = append(v3, (fi29 ^ fi29))
        _ = fi29
    }
    _ = len(v3)
    v1 = v1
    fmt.Println(v4)
    for fi30 := 0; fi30 < 13; fi30++ {
        var v31 float32 = 3.14159
        var v32 uint32 = 2
        _ = fi30
        _ = v31
        _ = v32
    }
    v4 = (v4 % ((1) | 1))
    switch (v1 + "\n\t") {
    case "a":
        var v33 int64 = 42
        _ = v33
    default:
        v4 = (v4 - v4)
    }
    if false {
        for fi34 := 0; fi34 < 6; fi34++ {
            var v35 uint16 = 1
            v36 := GenS3{F0: true, F1: false, F2: 951}
            v35 = v35
            _ = fi34
            _ = v35
            _ = v36
        }
        fmt.Println(v4)
        if false {
            for fi37 := 0; fi37 < 12; fi37++ {
                _ = len(v2)
                var v38 uint8 = 1
                fmt.Println(fi37, v1)
                _ = fi37
                _ = v38
            }
            v2 = append(v2, 127)
        }
    }
    v1 = v1
    v4 = (v4 & v4)
    if (42 != 1) {
        var v39 int16 = 1
        _ = v39
    } else {
        for fi40 := 0; fi40 < 20; fi40++ {
            var v41 float32 = ((0.1 - (1.5 + 3.14159)) * -2.25)
            _ = fi40
            _ = v41
        }
        var v42 uint = 65535
        var v43 int32 = v4
        _ = v42
        _ = v43
    }
    for fi44 := 0; fi44 < 8; fi44++ {
        for fi45 := 0; fi45 < 7; fi45++ {
            v3 = append(v3, fi45)
            _ = fi45
        }
        for fi46 := 0; fi46 < 5; fi46++ {
            if true {
                switch fi44 {
                case -388:
                    _ = len(v3)
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v1 = (v1 + ((v1 + v1) + v1))
                var v47 bool = (v1 >= ((("a" + v1) + v1) + v1))
                _ = v47
            }
            fmt.Println(v1)
            switch v1 {
            case "a":
                v1 = v1
            }
            _ = fi46
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi44
    }
    var v48 int64 = -24
    v48 += v48
    if true {
        switch v1 {
        case "":
            var v49 float32 = (3.14159 - -2.25)
            _ = v49
        case "\n\t":
            if true {
                v1 = (("foo bar" + (v1 + v1)) + v1)
                fmt.Println(v4, v4, v48)
                _ = len(v3)
            } else {
                if (-1 == 9223372036854775807) {
                    for fi50 := 0; fi50 < 8; fi50++ {
                        var v51 float64 = float64(1e10)
                        _ = fi50
                        _ = v51
                    }
                    v1 = v1
                    var v52 uint16 = 18
                    _ = v52
                } else {
                    fmt.Println(v1, v48, v4)
                    v1 = ("" + v1)
                }
                for fi53 := 0; fi53 < 16; fi53++ {
                    var v54 string = "\n\t"
                    v1 = v1
                    if (!true) {
                        for fi55 := 0; fi55 < 15; fi55++ {
                            var v56 uint16 = 270
                            v48 += v48
                            _ = fi55
                            _ = v56
                        }
                        fmt.Println(fi53, fi53, fi53)
                        switch v1 {
                        case "hello":
                            if (1 == 2) {
                                fmt.Println(v54)
                                v57 := GenS1{F0: -1, F1: 746}
                                v58 := GenS2{F0: 1000, F1: 127}
                                _ = v57
                                _ = v58
                            }
                        case "a":
                            v59 := GenS4{F0: 255, F1: -32768, F2: 2}
                            _ = v59
                        case "foo bar":
                            v60 := GenS3{F0: (fi53 <= -394), F1: ((float64(1e10) + ((3.14159 * 0.0) * (3.14159 * 1.5))) < -2.25), F2: 557}
                            _ = v60
                        default:
                            v3 = append(v3, 699)
                        }
                    } else {
                        var v61 uint8 = 128
                        v1 = ("a" + "\n\t")
                        v2 = append(v2, (fi53 >> ((4294967295) & 63)))
                        _ = v61
                    }
                    _ = fi53
                    _ = v54
                }
                v48 -= (v48 / (((v48 | 0)) | 1))
            }
        case "a":
            switch v1 {
            case "foo bar":
                v62 := GenS2{F0: 1, F1: 127}
                _ = v62
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
    }
    v4 -= v4
    if false {
        if true {
            fmt.Println(v48)
        } else {
            var v63 int16 = -124
            _ = v63
        }
        fmt.Println(v1, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
}
