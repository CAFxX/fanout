package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 bool; F2 int8 }
type GenS3 struct { F0 int16; F1 string; F2 uint32 }
type GenS4 struct { F0 int32; F1 int16; F2 int16 }
type GenS5 struct { F0 int32; F1 uint32; F2 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32 }

func main() {
    v1 := []int{}
    v2 := make(map[int]string)
    var v3 uint = 417
    v4 := make(map[int]bool)
    var v5 string = ""
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := GenS1{F0: (v3 - (^(v3)))}
        v3 = (v3 | 849)
        for fi8 := 0; fi8 < 1; fi8++ {
            v5 = v5
            if len(v1) > 0 {
                v1[0] = fi6
            }
            if (-2147483648 > -931) {
                v9 := GenS2{F0: 0, F1: true, F2: 0}
                if true {
                    var v10 uint32 = 42
                    v11 := GenS3{F0: 0, F1: v5, F2: (v10 - v10)}
                    if true {
                        var v12 uint16 = 255
                        _ = v12
                    } else {
                        v13 := GenS2{F0: 871, F1: true, F2: 0}
                        _ = v13
                    }
                    _ = v10
                    _ = v11
                } else {
                    var v14 uint = v3
                    _ = v14
                }
                var v15 uint16 = 32767
                _ = v15
                _ = v9
            }
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    if false {
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 int16 = 42
            v5 = (("foo bar" + (v5 + "")) + v5)
            for fi18 := 0; fi18 < 1; fi18++ {
                if len(v1) > 0 {
                    v1[0] = fi18
                }
                if true {
                    if (2147483647 != 814) {
                        var v19 float32 = ((1e10 * (-2.25 * -2.25)) - 0.1)
                        if (true || (-2147483648 >= 0)) {
                            switch "hello" {
                            case "":
                                v17 -= (-(v17))
                            default:
                                v3 += v3
                            }
                            v3 += v3
                            if (0 >= 255) {
                                v19 += ((v19 - v19) - (((v19 + 0.1) - v19) * float32(1e10)))
                                _ = len(v2)
                                var v20 uint16 = 0
                                var v21 int32 = -2147483648
                                v3 += (^(v3))
                                _ = v20
                                _ = v21
                            }
                        }
                        if false {
                            var v22 uint64 = 385
                            _ = v22
                        } else {
                            if (2 == uint8(v3)) {
                                if (((v5 + v5) + (v5 + (v5 + v5))) >= (((v5 + "hello") + "") + v5)) {
                                    var v23 int64 = -380
                                    _ = v23
                                }
                                v5 = v5
                            }
                            v1 = append(v1, fi16)
                        }
                        _ = v19
                    } else {
                        if true {
                            v5 = (((((v5 + v5) + v5) + v5) + (v5 + v5)) + "foo bar")
                        }
                        _ = len(v1)
                        var v24 uint32 = 797
                        _ = v24
                    }
                } else {
                    fmt.Println(fi18, v17)
                    _ = len(v4)
                }
                _ = fi18
            }
            _ = fi16
            _ = v17
        }
        fmt.Println(v5)
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        var v26 uint16 = 142
        var v27 uint16 = v26
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint8 = 1
    v28 = 1
    for fi29 := 0; fi29 < 13; fi29++ {
        if (fi29 != (fi29 << ((v3) & 63))) {
            switch v5 {
            case "hello":
                if (!true) {
                    v28 -= (v28 % ((114) | 1))
                    _ = len(v1)
                    if true {
                        v30 := GenS4{F0: -2147483648, F1: -32768, F2: 1000}
                        _ = v30
                    } else {
                        delete(v4, (fi29 | (fi29 >> ((v3) & 63))))
                        v3 += (v3 << (((v3 * v3)) & 63))
                    }
                } else {
                    v28 = (v28 % ((v28) | 1))
                }
            case "foo bar":
                var v31 uint32 = 42
                _ = v31
            case "a":
                switch v5 {
                case "foo bar":
                    switch (v5 + (v5 + v5)) {
                    case "a":
                        var v32 uint8 = v28
                        _ = v32
                    case "":
                        for fi33 := 0; fi33 < 19; fi33++ {
                            if true {
                                fmt.Println(fi33, fi29)
                            }
                            _ = fi33
                        }
                    }
                case "":
                    var v34 int32 = 0
                    _ = v34
                }
            default:
                switch v5 {
                case "foo bar":
                    v35 := GenS5{F0: 1, F1: 2147483648, F2: 0}
                    _ = v35
                case "\n\t":
                    switch fi29 {
                    case 0:
                        v4[fi29] = false
                    }
                case "a":
                    for fi36 := 0; fi36 < 18; fi36++ {
                        fmt.Println(fi36)
                        _ = fi36
                    }
                default:
                    v3 -= (v3 + (v3 | v3))
                }
            }
        }
        v37 := GenS6{F0: (v3 & v3)}
        v38 := GenS1{F0: uint(v28)}
        _ = fi29
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        var v40 int16 = 1
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi41 := 0; fi41 < 20; fi41++ {
            v42 := GenS5{F0: 42, F1: uint32(v3), F2: (v40 & 312)}
            _ = fi41
            _ = v42
        }
        _ = fi39
        _ = v40
    }
    if (!true) {
        v43 := GenS3{F0: 1, F1: v5, F2: 65535}
        for fi44 := 0; fi44 < 14; fi44++ {
            _ = len(v4)
            v28 = (v28 * (v28 + (v28 / (((v28 / ((v28) | 1))) | 1))))
            _ = fi44
        }
        if (!(false && (-9223372036854775808 == 938))) {
            for fi45 := 0; fi45 < 14; fi45++ {
                v28 -= uint8(v3)
                v28 = 128
                for fi46 := 0; fi46 < 11; fi46++ {
                    v3 += 909
                    _ = fi46
                }
                _ = fi45
            }
            var v47 int32 = -786
            _ = v47
        }
        _ = v43
    }
    for fi48 := 0; fi48 < 9; fi48++ {
        if true {
            for fi49 := 0; fi49 < 6; fi49++ {
                v3 = 422
                if len(v1) > 0 {
                    v1[0] = 32768
                }
                var v50 uint8 = v28
                _ = fi49
                _ = v50
            }
        } else {
            if (42 != -305) {
                for fi51 := 0; fi51 < 18; fi51++ {
                    fmt.Println(fi51, v28)
                    if len(v1) > 0 {
                        v1[0] = (fi51 * fi51)
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v52 string = (v5 + "hello")
                    _ = fi51
                    _ = v52
                }
            }
            v3 = v3
            var v53 int64 = 664
            _ = v53
        }
        v54 := GenS7{F0: 255}
        _ = fi48
        _ = v54
    }
    fmt.Println(v5, v5, v3)
    if true {
        for fi55 := 0; fi55 < 10; fi55++ {
            var v56 int16 = -32768
            _ = fi55
            _ = v56
        }
        v28 -= 1
        v28 = (v28 / ((v28) | 1))
    }
    for fi57 := 0; fi57 < 15; fi57++ {
        var v58 uint8 = uint8(v3)
        _ = fi57
        _ = v58
    }
    v5 = "hello"
    v59 := GenS8{F0: -1}
    for fi60 := 0; fi60 < 7; fi60++ {
        fmt.Println(v3, v28, fi60)
        _ = fi60
    }
    var v61 float64 = 1e10
    v5 = (v5 + v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v61
}
