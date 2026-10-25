package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 bool }
type GenS2 struct { F0 uint; F1 uint16; F2 int8 }
type GenS3 struct { F0 int64; F1 int; F2 uint }
type GenS4 struct { F0 int; F1 uint; F2 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int16 }

func main() {
    var v1 int32 = 42
    var v2 int8 = 1
    var v3 uint8 = 42
    var v4 uint8 = 128
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: 42, F1: v1, F2: false}
        fmt.Println(v2, v3, v4)
        _ = fi5
        _ = v6
    }
    v7 := GenS2{F0: 18446744073709551615, F1: 237, F2: (v2 >> ((73) & 63))}
    fmt.Println(v4)
    if false {
        fmt.Println(v1, v1, v4)
        fmt.Println(v2, v1, v3)
    } else {
        v1 -= -2147483648
        v2 = v2
    }
    v8 := []int{42, 118}
    if (33 < 65535) {
        switch ("hello" + "\n\t") {
        case "":
            switch "hello" {
            case "foo bar":
                fmt.Println(v4, v2, v1)
            default:
                var v9 int32 = v1
                _ = v9
            }
        case "a":
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        case "foo bar":
            var v10 float32 = float32((0.1 / 3.14159))
            _ = v10
        }
        switch "a" {
        case "\n\t":
            switch (("\n\t" + ("" + (("foo bar" + ("foo bar" + "\n\t")) + "a"))) + ("foo bar" + ("hello" + ("foo bar" + ("\n\t" + "foo bar"))))) {
            case "":
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            default:
                switch 2147483647 {
                case 665:
                    v11 := GenS3{F0: 9223372036854775807, F1: -1, F2: 325}
                    _ = v11
                case 1:
                    if len(v8) > 0 {
                        v8[0] = 32768
                    }
                }
            }
        default:
            var v12 int64 = 1
            _ = v12
        }
        var v13 int16 = -32768
        _ = v13
    }
    var v14 uint = uint(v4)
    v15 := GenS2{F0: (v14 | v14), F1: 307, F2: v2}
    switch 916 {
    case -1:
        var v16 float64 = ((-2.25 / 1.5) - float64(1e10))
        _ = v16
    case 32767:
        var v17 uint8 = uint8(v14)
        _ = v17
    default:
        v1 -= (v1 % ((494) | 1))
    }
    if true {
        v14 += 2
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        switch fi18 {
        case 42:
            switch "foo bar" {
            case "hello":
                v19 := GenS1{F0: 0, F1: v1, F2: (489 > 407)}
                _ = v19
            }
        case -122:
            v4 += v4
        case 9223372036854775807:
            var v20 uint64 = 167
            _ = v20
        default:
            var v21 uint32 = 1
            _ = v21
        }
        switch -1 {
        case -916:
            if true {
                if false {
                    v4 += (v3 * v4)
                    var v22 int8 = v2
                    var v23 uint = 0
                    _ = v22
                    _ = v23
                }
                v24 := GenS2{F0: 18446744073709551615, F1: 0, F2: (v2 % ((int8(fi18)) | 1))}
                _ = v24
            } else {
                var v25 uint32 = 549
                _ = v25
            }
        case 588:
            var v26 int32 = (v1 ^ -886)
            _ = v26
        }
        _ = fi18
    }
    if (!true) {
        var v27 int8 = v2
        _ = len(v8)
        switch -1 {
        case 895:
            v3 = (^(v3))
        case 42:
            for fi28 := 0; fi28 < 9; fi28++ {
                var v29 uint16 = 604
                fmt.Println(v1)
                _ = fi28
                _ = v29
            }
        default:
            if false {
                v27 -= (v2 ^ (v27 / ((v2) | 1)))
                for fi30 := 0; fi30 < 17; fi30++ {
                    var v31 int = (fi30 / ((fi30) | 1))
                    _ = len(v8)
                    v14 = (v14 & v14)
                    _ = fi30
                    _ = v31
                }
            } else {
                if false {
                    var v32 uint = 255
                    v3 = (v4 - v3)
                    _ = v32
                }
            }
        }
        _ = v27
    }
    switch "" {
    case "\n\t":
        fmt.Println(v1, v2, v4)
    case "hello":
        for fi33 := 0; fi33 < 9; fi33++ {
            for fi34 := 0; fi34 < 13; fi34++ {
                if true {
                    if false {
                        v3 = (^(v4))
                    } else {
                        if (-1 != 42) {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        } else {
                            var v35 float64 = 1.5
                            _ = v35
                        }
                        v4 += v4
                    }
                }
                v4 -= v4
                _ = len(v8)
                v36 := GenS4{F0: fi33, F1: (v14 % (((^(v14))) | 1)), F2: v3}
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi34
                _ = v36
            }
            for fi37 := 0; fi37 < 16; fi37++ {
                v8 = append(v8, -9223372036854775808)
                if len(v8) > 0 {
                    v8[0] = fi33
                }
                fmt.Println(v14, v2)
                _ = fi37
            }
            var v38 bool = false
            _ = fi33
            _ = v38
        }
    default:
        for fi39 := 0; fi39 < 8; fi39++ {
            var v40 int16 = 0
            if true {
                if (((!false) && false) && (255 > 42)) {
                    var v41 int8 = (v2 - (v2 & v2))
                    _ = v41
                } else {
                    if ((fi39 ^ (^(fi39))) > (fi39 >> ((v14) & 63))) {
                        v42 := GenS5{F0: 4294967295}
                        v3 = (v3 >> (((v14 | (v14 / (((v14 | 92)) | 1)))) & 63))
                        v1 -= 621
                        _ = v42
                    } else {
                        var v43 uint = v14
                        var v44 uint64 = 412
                        v2 = v2
                        _ = v43
                        _ = v44
                    }
                }
                switch 718 {
                case -482:
                    if true {
                        fmt.Println(v2, v14, v40)
                    }
                case 9223372036854775807:
                    var v45 int64 = int64(v3)
                    _ = v45
                default:
                    if (!(7 <= 649)) {
                        if len(v8) > 0 {
                            v8[0] = fi39
                        }
                        if false {
                            fmt.Println(fi39)
                            v4 = 1
                        } else {
                            v46 := GenS6{F0: (^(v40))}
                            _ = v46
                        }
                        for fi47 := 0; fi47 < 20; fi47++ {
                            var v48 float64 = 1.5
                            v2 -= v2
                            switch "a" {
                            case "foo bar":
                                switch ("hello" + "foo bar") {
                                case "hello":
                                    var v49 uint = (^(v14))
                                    _ = v49
                                case "foo bar":
                                    if (!(!false)) {
                                        v1 += (v1 % ((v1) | 1))
                                        v2 = v2
                                    } else {
                                        v4 -= (v4 / ((v4) | 1))
                                        v4 -= v4
                                    }
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                case "\n\t":
                                    v14 = (v14 - v14)
                                default:
                                    v50 := GenS1{F0: 18446744073709551615, F1: (v1 ^ v1), F2: true}
                                    _ = v50
                                }
                            default:
                                var v51 uint = (^(v14))
                                _ = v51
                            }
                            _ = fi47
                            _ = v48
                        }
                    }
                }
            }
            var v52 int32 = 604
            _ = fi39
            _ = v40
            _ = v52
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
