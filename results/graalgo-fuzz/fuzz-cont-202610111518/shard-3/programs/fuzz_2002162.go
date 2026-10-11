package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint }
type GenS4 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 bool; F1 uint; F2 int }
type GenS8 struct { F0 int64; F1 uint; F2 int8 }
func genFn0() uint {
    var v5 uint = 905
    fmt.Println(v5)
    switch ("" + ("\n\t" + ("hello" + "foo bar"))) {
    case "\n\t":
        var v6 int16 = 0
        _ = v6
    case "":
        v5 = 4294967295
    default:
        v7 := []int{-1, 32768}
        _ = v7
    }
    v8 := []int{}
    switch "a" {
    case "\n\t":
        switch 42 {
        case 1:
            v9 := GenS1{F0: (v5 >> ((v5) & 63))}
            _ = v9
        case 891:
            for fi10 := 0; fi10 < 17; fi10++ {
                switch -1 {
                case 42:
                    if (!false) {
                        var v11 string = ("" + (("" + "foo bar") + "foo bar"))
                        var v12 int8 = 100
                        _ = v11
                        _ = v12
                    }
                case 32767:
                    if false {
                        for fi13 := 0; fi13 < 16; fi13++ {
                            v5 = (v5 / ((v5) | 1))
                            if ((v5 % ((v5) | 1)) <= v5) {
                                var v14 int32 = 2147483647
                                _ = len(v8)
                                var v15 bool = (1 > 0)
                                var v16 uint64 = 868
                                if false {
                                    fmt.Println(v16, fi13, fi13)
                                    v17 := GenS2{F0: 1}
                                    _ = v17
                                } else {
                                    v18 := GenS2{F0: 1}
                                    v14 -= v14
                                    switch ("foo bar" + ("hello" + "\n\t")) {
                                    case "\n\t":
                                        fmt.Println(fi13, v14)
                                    case "foo bar":
                                        _ = len(v8)
                                    case "hello":
                                        v19 := GenS3{F0: 32767, F1: 127, F2: 18446744073709551615}
                                        _ = v19
                                    default:
                                        if v15 {
                                            var v20 int8 = -96
                                            _ = v20
                                        }
                                    }
                                    v15 = v15
                                    _ = v18
                                }
                                var v21 int16 = -56
                                v21 = (v21 & (v21 - v21))
                                _ = v14
                                _ = v15
                                _ = v16
                                _ = v21
                            }
                            var v22 string = "\n\t"
                            _ = fi13
                            _ = v22
                        }
                        v5 = v5
                    }
                case -655:
                    if (true && true) {
                        v5 += 18446744073709551615
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        var v23 int32 = 2147483647
                        _ = v23
                    } else {
                        var v24 float32 = 1.5
                        v25 := GenS1{F0: (^(v5))}
                        _ = v24
                        _ = v25
                    }
                default:
                    fmt.Println(v5, v5)
                }
                _ = fi10
            }
        case -969:
            fmt.Println(v5)
        }
    case "a":
        v8 = append(v8, 91)
    case "foo bar":
        for fi26 := 0; fi26 < 19; fi26++ {
            v27 := GenS4{F0: 42, F1: 659, F2: 574}
            var v28 string = ("hello" + "foo bar")
            _ = fi26
            _ = v27
            _ = v28
        }
    default:
        v5 += (v5 | (v5 - (v5 / (((v5 ^ v5)) | 1))))
    }
    _ = v5
    _ = v8
    return v5
}

func main() {
    var v1 string = "a"
    var v2 uint8 = 127
    var v3 float32 = float32(-2.25)
    v4 := []int{32767, -879, -137}
    _ = len(v4)
    var v29 float32 = 0.0
    switch v1 {
    case "foo bar":
        var v30 uint32 = 0
        _ = v30
    case "\n\t":
        fmt.Println(v1, v29)
    case "hello":
        _ = len(v4)
    }
    v3 = -2.25
    if ((v2 & 1) != 1) {
        v31 := GenS1{F0: 614}
        _ = v31
    } else {
        if len(v4) > 0 {
            v4[0] = 0
        }
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        v3 = v29
        _ = fi32
    }
    v33 := GenS3{F0: 662, F1: -1, F2: 42}
    switch 32768 {
    case 823:
        for fi34 := 0; fi34 < 16; fi34++ {
            if true {
                var v35 uint64 = 935
                var v36 uint8 = 2
                fmt.Println(v35, v36, v3)
                _ = v35
                _ = v36
            }
            var v37 uint = 4294967295
            if (int32(fi34) == 42) {
                var v38 uint16 = uint16(v2)
                _ = v38
            }
            _ = fi34
            _ = v37
        }
    case 32768:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case -1:
        if len(v4) > 0 {
            v4[0] = 9223372036854775807
        }
    default:
        v39 := GenS5{F0: -2147483648}
        _ = v39
    }
    var v40 uint = 42
    v3 -= ((0.0 - 1.5) + (0.0 / v29))
    v40 = (v40 + (v40 / ((v40) | 1)))
    switch (v1 + "") {
    case "a":
        switch v1 {
        case "foo bar":
            var v41 bool = true
            _ = v41
        }
    case "hello":
        v1 = v1
    case "foo bar":
        var v42 uint32 = 0
        _ = v42
    default:
        v4 = append(v4, 32767)
    }
    if (int8(v2) == -128) {
        for fi43 := 0; fi43 < 2; fi43++ {
            var v44 int = 9223372036854775807
            if true {
                switch (^(fi43)) {
                case -655:
                    var v45 uint8 = 2
                    _ = v45
                case -251:
                    for fi46 := 0; fi46 < 2; fi46++ {
                        var v47 int64 = 9223372036854775807
                        _ = fi46
                        _ = v47
                    }
                case 1:
                    switch (^(fi43)) {
                    case 754:
                        if len(v4) > 0 {
                            v4[0] = v44
                        }
                    case -2147483648:
                        fmt.Println(v2, v44)
                    default:
                        v4 = append(v4, v44)
                    }
                }
                v2 -= 1
                if (!(2147483648 <= 65535)) {
                    v48 := GenS3{F0: uint16(v40), F1: -128, F2: 913}
                    for fi49 := 0; fi49 < 18; fi49++ {
                        switch v1 {
                        case "":
                            fmt.Println(fi43, v1)
                        default:
                            v1 = (v1 + v1)
                        }
                        if len(v4) > 0 {
                            v4[0] = fi49
                        }
                        _ = fi49
                    }
                    if true {
                        v29 -= v29
                    }
                    _ = v48
                } else {
                    v50 := GenS6{F0: 42, F1: v40}
                    switch v1 {
                    case "":
                        fmt.Println(v2)
                    case "a":
                        var v51 float64 = 0.0
                        _ = v51
                    }
                    _ = v50
                }
            } else {
                v40 = v40
                for fi52 := 0; fi52 < 12; fi52++ {
                    var v53 int32 = -1
                    _ = fi52
                    _ = v53
                }
            }
            if false {
                for fi54 := 0; fi54 < 18; fi54++ {
                    v55 := GenS7{F0: (-449 > 32767), F1: (v40 | v40), F2: -602}
                    v3 += 1e10
                    _ = fi54
                    _ = v55
                }
            } else {
                var v56 bool = true
                var v57 bool = (genFn0() > (v40 - 18446744073709551615))
                fmt.Println(v29, v3)
                _ = v56
                _ = v57
            }
            _ = fi43
            _ = v44
        }
        for fi58 := 0; fi58 < 10; fi58++ {
            v40 = (^(v40))
            for fi59 := 0; fi59 < 9; fi59++ {
                switch v1 {
                case "a":
                    v40 -= 951
                }
                var v60 int16 = 1000
                if ((!false) || false) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch v1 {
                    case "hello":
                        var v61 uint8 = (v2 + v2)
                        _ = v61
                    default:
                        var v62 uint64 = 42
                        switch (fi58 % (((fi58 ^ (fi58 - -287))) | 1)) {
                        case -1:
                            v60 = v60
                        case 0:
                            var v63 uint16 = 1
                            if true {
                                fmt.Println(v2)
                                _ = len(v4)
                            }
                            v64 := GenS8{F0: int64(v60), F1: (v40 | v40), F2: -38}
                            _ = len(v4)
                            var v65 float32 = (v29 * (v3 / 1e10))
                            _ = v63
                            _ = v64
                            _ = v65
                        }
                        _ = v62
                    }
                }
                var v66 bool = false
                _ = fi59
                _ = v60
                _ = v66
            }
            if len(v4) > 0 {
                v4[0] = (fi58 ^ -1)
            }
            _ = fi58
        }
    }
    v67 := GenS5{F0: 534}
    var v68 uint32 = 2147483648
    for fi69 := 0; fi69 < 6; fi69++ {
        v70 := GenS8{F0: -1, F1: 0, F2: 0}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch -2147483648 {
        case 414:
            v29 = float32((1.5 - -2.25))
        case 2147483647:
            var v71 uint32 = v68
            _ = v71
        case 32768:
            var v72 uint16 = 1
            _ = v72
        default:
            v3 += (v3 * -2.25)
        }
        _ = fi69
        _ = v70
    }
    v68 += v68
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v40
    _ = v67
    _ = v68
}
