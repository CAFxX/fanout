package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 bool }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 int64; F2 string }
type GenS6 struct { F0 int8; F1 uint16; F2 string }

func main() {
    var v1 bool = false
    var v2 uint32 = 2
    var v3 uint16 = 886
    v4 := []int{}
    var v5 float32 = 0.1
    fmt.Println(v3)
    var v6 string = ""
    var v7 float32 = float32(float64(v5))
    fmt.Println(v7, v5)
    if (false && false) {
        switch -2147483648 {
        case 127:
            v8 := GenS1{F0: 557, F1: 0, F2: v1}
            _ = v8
        case 469:
            for fi9 := 0; fi9 < 10; fi9++ {
                v1 = (1e10 >= v5)
                _ = fi9
            }
        case 4294967295:
            switch 780 {
            case 42:
                var v10 uint32 = 1
                _ = v10
            case 32768:
                _ = len(v4)
            default:
                _ = len(v4)
            }
        }
        v4 = append(v4, 32767)
        v1 = (32767 >= -419)
    }
    fmt.Println(v2, v6, v3)
    if v1 {
        v6 = "hello"
        if v1 {
            for fi11 := 0; fi11 < 19; fi11++ {
                v1 = (0 < 1)
                _ = fi11
            }
            var v12 uint16 = 1
            _ = v12
        }
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 uint8 = 1
            if len(v4) > 0 {
                v4[0] = 179
            }
            switch 42 {
            case 32767:
                v1 = v1
            case -2147483648:
                v15 := GenS1{F0: 255, F1: 612, F2: (v1 || v1)}
                _ = v15
            }
            _ = fi13
            _ = v14
        }
    } else {
        if v1 {
            if (!((v2 % (((^(v2))) | 1)) != 1)) {
                v16 := GenS1{F0: 4294967295, F1: 42, F2: (!(42 >= 0))}
                var v17 uint = 1
                _ = v16
                _ = v17
            } else {
                var v18 uint16 = v3
                if ((3.14159 - ((1e10 * ((0.0 + float64(v7)) / (0.1 + (0.1 / -2.25)))) * 0.0)) >= (3.14159 / 3.14159)) {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v1 = (0.0 > 0.1)
                        v18 = (v3 >> ((992) & 63))
                        _ = fi19
                    }
                    var v20 int16 = 0
                    if ((!v1) && (128 != -760)) {
                        switch "" {
                        case "a":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            if len(v4) > 0 {
                                v4[0] = 9223372036854775807
                            }
                        }
                    } else {
                        var v21 uint32 = (^(v2))
                        switch 128 {
                        case 1:
                            switch 313 {
                            case 9223372036854775807:
                                _ = len(v4)
                            case 672:
                                if len(v4) > 0 {
                                    v4[0] = -179
                                }
                            default:
                                var v22 float32 = 0.0
                                _ = v22
                            }
                        case 32768:
                            if true {
                                if false {
                                    switch v6 {
                                    case "a":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case "":
                                        v23 := GenS1{F0: 255, F1: 804, F2: true}
                                        _ = v23
                                    default:
                                        for fi24 := 0; fi24 < 13; fi24++ {
                                            for fi25 := 0; fi25 < 13; fi25++ {
                                                fmt.Println(v2, v7, v2)
                                                v26 := GenS1{F0: 304, F1: 971, F2: v1}
                                                var v27 uint = 42
                                                _ = fi25
                                                _ = v26
                                                _ = v27
                                            }
                                            for fi28 := 0; fi28 < 2; fi28++ {
                                                v4 = append(v4, fi28)
                                                _ = fi28
                                            }
                                            _ = fi24
                                        }
                                    }
                                }
                                var v29 int8 = 0
                                if v1 {
                                    switch (v6 + (((v6 + v6) + v6) + "hello")) {
                                    case "a":
                                        _ = len(v4)
                                    case "":
                                        if v1 {
                                            if v1 {
                                                for fi30 := 0; fi30 < 12; fi30++ {
                                                    if ((!v1) && (v1 || v1)) {
                                                        if true {
                                                            for fi31 := 0; fi31 < 12; fi31++ {
                                                                var v32 uint = 0
                                                                _ = fi31
                                                                _ = v32
                                                            }
                                                            v33 := GenS2{F0: v3, F1: v29}
                                                            v34 := GenS2{F0: v18, F1: (v29 << ((65535) & 63))}
                                                            _ = v33
                                                            _ = v34
                                                        } else {
                                                            v4 = append(v4, fi30)
                                                            v1 = v1
                                                        }
                                                        v6 = v6
                                                    } else {
                                                        v35 := GenS1{F0: 476, F1: 952, F2: ((((((0.1 * float64(v7)) * (3.14159 - 1.5)) / (1e10 - 3.14159)) * 1e10) / float64(v5)) >= ((-2.25 + 1e10) + (((1e10 + ((0.1 + float64((1e10 - float32(1.5)))) * 0.1)) - (-2.25 / 0.1)) + -2.25)))}
                                                        var v36 string = v6
                                                        _ = len(v4)
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                    _ = fi30
                                                }
                                            }
                                            for fi37 := 0; fi37 < 20; fi37++ {
                                                v21 = (v21 + (^(v2)))
                                                if ((!v1) || (v29 < (^(v29)))) {
                                                    switch fi37 {
                                                    case 0:
                                                        _ = len(v4)
                                                    default:
                                                        if (true || (0 > 2147483647)) {
                                                            var v38 uint64 = 2
                                                            v6 = (("\n\t" + ("a" + "a")) + "foo bar")
                                                            _ = v38
                                                        }
                                                    }
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        switch v6 {
                                                        case "":
                                                            var v39 uint8 = 42
                                                            _ = v39
                                                        }
                                                    case "\n\t":
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                    var v40 float64 = -2.25
                                                    _ = v40
                                                }
                                                _ = fi37
                                            }
                                        } else {
                                            switch -2147483648 {
                                            case 97:
                                                for fi41 := 0; fi41 < 15; fi41++ {
                                                    v4 = append(v4, -9223372036854775808)
                                                    _ = fi41
                                                }
                                            case 0:
                                                if len(v4) > 0 {
                                                    v4[0] = -9223372036854775808
                                                }
                                            }
                                            v42 := GenS3{F0: 255, F1: ((v6 >= (v6 + "")) && false)}
                                            for fi43 := 0; fi43 < 4; fi43++ {
                                                var v44 int16 = (v20 % ((0) | 1))
                                                var v45 uint16 = 32767
                                                var v46 string = ((v6 + v6) + v6)
                                                _ = fi43
                                                _ = v44
                                                _ = v45
                                                _ = v46
                                            }
                                            _ = v42
                                        }
                                    case "foo bar":
                                        v3 = 279
                                    }
                                } else {
                                    fmt.Println(v21)
                                    if (255 >= 65535) {
                                        if v1 {
                                            v5 += v5
                                        }
                                        var v47 uint = 852
                                        _ = v47
                                    }
                                }
                                _ = v29
                            } else {
                                switch "a" {
                                case "foo bar":
                                    var v48 uint32 = (v21 + v2)
                                    _ = v48
                                case "a":
                                    if len(v4) > 0 {
                                        v4[0] = 128
                                    }
                                case "hello":
                                    v4 = append(v4, 32767)
                                }
                            }
                        case -9223372036854775808:
                            fmt.Println(v21, v18, v1)
                        default:
                            v49 := GenS4{F0: ((float32(float64(1.5)) == (0.1 / v7)) && ((!v1) || v1))}
                            _ = v49
                        }
                        _ = v21
                    }
                    _ = v20
                } else {
                    var v50 string = "\n\t"
                    fmt.Println(v50, v5, v6)
                    v51 := GenS3{F0: 43, F1: ((v5 < 1.5) && v1)}
                    _ = v50
                    _ = v51
                }
                _ = v18
            }
            for fi52 := 0; fi52 < 4; fi52++ {
                var v53 int32 = 1
                v5 = v7
                _ = fi52
                _ = v53
            }
        }
        v7 = v7
        switch ((v6 + (v6 + ("foo bar" + "a"))) + "a") {
        case "\n\t":
            var v54 int16 = int16(v2)
            _ = v54
        case "":
            if (0.1 != ((3.14159 * (float64((float32((0.1 * 1.5)) - 0.0)) * 0.1)) + (-2.25 - 1.5))) {
                v7 = v7
                if ((-1 <= 470) && ((1e10 <= 0.0) || v1)) {
                    for fi55 := 0; fi55 < 1; fi55++ {
                        var v56 string = v6
                        _ = fi55
                        _ = v56
                    }
                    switch "foo bar" {
                    case "hello":
                        var v57 uint8 = 127
                        _ = v57
                    case "\n\t":
                        var v58 float32 = ((float32(3.14159) - (v5 * v5)) + v7)
                        _ = v58
                    case "a":
                        var v59 uint8 = 46
                        _ = v59
                    default:
                        var v60 string = "a"
                        _ = v60
                    }
                }
                v5 -= v7
            }
        }
    }
    var v61 int8 = int8(v2)
    v61 += 100
    if (!true) {
        var v62 string = v6
        for fi63 := 0; fi63 < 1; fi63++ {
            if v1 {
                v1 = v1
                if (!v1) {
                    v6 = v62
                } else {
                    v5 += v7
                }
                if (128 == 127) {
                    if v1 {
                        v4 = append(v4, fi63)
                        switch "a" {
                        case "hello":
                            _ = len(v4)
                        case "":
                            var v64 uint32 = v2
                            _ = v64
                        }
                        var v65 float64 = (1e10 * -2.25)
                        _ = v65
                    }
                    if v1 {
                        switch "a" {
                        case "a":
                            for fi66 := 0; fi66 < 17; fi66++ {
                                if v1 {
                                    v67 := GenS3{F0: 128, F1: ((!false) && (578 == 0))}
                                    _ = v67
                                }
                                var v68 int = (fi63 + 2147483647)
                                _ = fi66
                                _ = v68
                            }
                        case "\n\t":
                            fmt.Println(v62)
                        }
                        if true {
                            var v69 int32 = 0
                            _ = v69
                        }
                        if v1 {
                            v70 := GenS3{F0: uint8(v2), F1: true}
                            _ = v70
                        }
                    }
                }
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v71 := GenS5{F0: 541, F1: 42, F2: v6}
                _ = v71
            }
            _ = fi63
        }
        _ = v62
    } else {
        v72 := GenS6{F0: -12, F1: (^(v3)), F2: (((v6 + v6) + "a") + ((v6 + "\n\t") + "hello"))}
        var v73 uint8 = 42
        _ = v72
        _ = v73
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi74 := 0; fi74 < 4; fi74++ {
        v4 = append(v4, 4294967295)
        _ = fi74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v7
}
