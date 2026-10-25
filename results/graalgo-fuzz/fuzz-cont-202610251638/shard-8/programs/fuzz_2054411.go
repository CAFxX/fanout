package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint; F1 int; F2 uint }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 int16; F1 uint16 }
type GenS7 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int8; F1 uint; F2 uint64 }
type GenS10 struct { F0 uint64; F1 uint8; F2 int32 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 103
    var v3 int8 = -1
    var v4 int32 = -2147483648
    if false {
        v5 := []int{32768, -892, 2147483647}
        switch "foo bar" {
        case "hello":
            v3 += (v3 * 100)
        case "":
            switch "foo bar" {
            case "hello":
                if (!(0.1 < 1.5)) {
                    v6 := GenS1{F0: "a", F1: 4294967295, F2: (301 != 1)}
                    var v7 int = 128
                    var v8 uint32 = 337
                    _ = v6
                    _ = v7
                    _ = v8
                } else {
                    v9 := GenS2{F0: 148}
                    if len(v5) > 0 {
                        v5[0] = -9223372036854775808
                    }
                    _ = v9
                }
            }
        case "\n\t":
            v10 := GenS2{F0: 65535}
            _ = v10
        }
        _ = v5
    }
    v4 += v4
    v2 -= (^(v2))
    v3 = v3
    var v11 string = (((("hello" + "\n\t") + "a") + "\n\t") + "foo bar")
    var v12 float32 = float32(0.1)
    v13 := GenS1{F0: v11, F1: 0, F2: (42 > 1)}
    switch v11 {
    case "hello":
        if false {
            if ((-(v3)) > -45) {
                if false {
                    for fi14 := 0; fi14 < 2; fi14++ {
                        for fi15 := 0; fi15 < 8; fi15++ {
                            v11 = (v11 + "a")
                            _ = fi15
                        }
                        for fi16 := 0; fi16 < 17; fi16++ {
                            if (fi16 == 2147483647) {
                                if false {
                                    _ = len(v1)
                                    delete(v1, 32768)
                                    var v17 string = v11
                                    _ = v17
                                }
                                var v18 int = (fi16 + int(v4))
                                var v19 uint64 = 4294967295
                                _ = v18
                                _ = v19
                            } else {
                                v12 = v12
                                if true {
                                    var v20 int = fi14
                                    switch 0 {
                                    case 1:
                                        v2 = (v2 << ((220) & 63))
                                    case -2147483648:
                                        var v21 int = (fi14 + (v20 - v20))
                                        _ = v21
                                    default:
                                        v22 := GenS3{F0: v2, F1: (^(v2)), F2: 2}
                                        _ = v22
                                    }
                                    _ = v20
                                }
                            }
                            _ = len(v1)
                            _ = fi16
                        }
                        v23 := []int{fi14, fi14, 2147483647}
                        _ = fi14
                        _ = v23
                    }
                    v2 -= (^(v2))
                    if (!true) {
                        v24 := GenS1{F0: ((("a" + (v11 + (v11 + v11))) + ((((v11 + "hello") + (((v11 + v11) + "\n\t") + v11)) + (v11 + ("\n\t" + ("hello" + "")))) + v11)) + ("a" + v11)), F1: 2147483647, F2: true}
                        _ = v24
                    }
                }
                if (!(((^(v3)) == -128) && false)) {
                    v25 := GenS2{F0: 42}
                    _ = v25
                } else {
                    v26 := GenS2{F0: 880}
                    _ = v26
                }
            } else {
                var v27 float64 = 0.1
                if false {
                    var v28 string = v11
                    var v29 int = -268
                    v30 := []int{42, -1}
                    _ = v28
                    _ = v29
                    _ = v30
                }
                v31 := GenS1{F0: "a", F1: 545, F2: (4294967295 != 1)}
                _ = v27
                _ = v31
            }
            v32 := []int{42, int(v3), 9223372036854775807}
            _ = v32
        } else {
            for fi33 := 0; fi33 < 9; fi33++ {
                _ = len(v1)
                switch -1 {
                case 127:
                    var v34 uint = 1
                    _ = v34
                case 1:
                    switch 127 {
                    case 42:
                        var v35 int64 = 1
                        _ = v35
                    case 32768:
                        v12 -= 3.14159
                    default:
                        for fi36 := 0; fi36 < 4; fi36++ {
                            v12 -= v12
                            _ = fi36
                        }
                    }
                case -2147483648:
                    v11 = ""
                default:
                    for fi37 := 0; fi37 < 20; fi37++ {
                        if true {
                            if true {
                                fmt.Println(v3)
                                v38 := GenS2{F0: 0}
                                _ = v38
                            } else {
                                v39 := []int{(-(fi33)), (fi37 >> ((1) & 63)), fi33, (fi33 - (fi37 ^ int(v4)))}
                                var v40 int = fi37
                                _ = v39
                                _ = v40
                            }
                            v41 := GenS4{F0: 0, F1: (fi33 * (fi33 | (fi37 * (fi37 + (fi33 % (((fi33 - fi37)) | 1)))))), F2: 18446744073709551615}
                            _ = v41
                        } else {
                            v42 := []int{(fi33 | -1), (fi37 << ((1) & 63)), -44, 997}
                            var v43 uint = 42
                            v44 := GenS5{F0: false, F1: (fi33 + fi33)}
                            if true {
                                if (42 != 0) {
                                    if len(v42) >= 2 {
                                        v42 = v42[1:len(v42)]
                                    }
                                } else {
                                    v3 = (v3 & 112)
                                    var v45 string = (("\n\t" + "\n\t") + "foo bar")
                                    v46 := GenS6{F0: -5, F1: uint16(v2)}
                                    _ = v45
                                    _ = v46
                                }
                            } else {
                                v2 -= v2
                                if len(v42) >= 2 {
                                    v42 = v42[1:len(v42)]
                                }
                                switch v11 {
                                case "a":
                                    fmt.Println(v4, fi37)
                                case "\n\t":
                                    fmt.Println(v12, v43, v11)
                                }
                                v47 := GenS5{F0: true, F1: fi33}
                                v2 = 0
                                _ = v47
                            }
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                        var v48 int = -838
                        if (90 == (v2 - v2)) {
                            v49 := []int{fi37, (fi33 * 42), fi37}
                            _ = v49
                        }
                        _ = fi37
                        _ = v48
                    }
                }
                _ = fi33
            }
        }
    case "":
        var v50 uint16 = 65535
        _ = v50
    }
    switch -433 {
    case 9223372036854775807:
        for fi51 := 0; fi51 < 12; fi51++ {
            switch (-(fi51)) {
            case 2147483647:
                switch (fi51 | 127) {
                case -409:
                    v52 := []int{fi51, -2147483648, (fi51 * (fi51 << ((450) & 63)))}
                    _ = v52
                default:
                    fmt.Println(v11, v2, v2)
                }
            case 790:
                var v53 uint8 = (v2 % (((v2 % ((v2) | 1))) | 1))
                _ = v53
            default:
                if (4294967295 <= 638) {
                    var v54 bool = (v4 == 2147483647)
                    if (!((false && ((v4 | (v4 >> ((1) & 63))) >= -1)) || false)) {
                        v55 := GenS7{F0: 65535, F1: 42, F2: 251}
                        _ = v55
                    } else {
                        v1[(fi51 | -9223372036854775808)] = ((77 <= uint16(v2)) || (!(v54 || v54)))
                    }
                    _ = v54
                }
            }
            v56 := []int{(fi51 | (^(fi51))), (fi51 - (fi51 * (-(fi51)))), 32767}
            _ = fi51
            _ = v56
        }
    }
    if ((true || false) || (255 == 42)) {
        var v57 float32 = v12
        switch 42 {
        case 824:
            if true {
                for fi58 := 0; fi58 < 20; fi58++ {
                    for fi59 := 0; fi59 < 14; fi59++ {
                        var v60 int8 = v3
                        _ = fi59
                        _ = v60
                    }
                    v4 = v4
                    fmt.Println(v2, v12)
                    _ = fi58
                }
                fmt.Println(v3, v11, v57)
            } else {
                v61 := GenS8{F0: -658}
                for fi62 := 0; fi62 < 16; fi62++ {
                    _ = len(v1)
                    fmt.Println(v12, v11)
                    _ = fi62
                }
                delete(v1, -1)
                _ = v61
            }
        case 127:
            fmt.Println(v11, v57, v11)
        }
        _ = v57
    }
    var v63 int64 = 1
    for fi64 := 0; fi64 < 8; fi64++ {
        v3 = 42
        var v65 string = "\n\t"
        if (true || (v4 > (v4 & (v4 | (v4 << ((318) & 63)))))) {
            v66 := GenS7{F0: 625, F1: 2, F2: 65535}
            _ = v66
        }
        _ = fi64
        _ = v65
    }
    v67 := GenS9{F0: v3, F1: uint(v2), F2: 18446744073709551615}
    switch 0 {
    case 0:
        v68 := GenS8{F0: 127}
        _ = v68
    case 127:
        fmt.Println(v2)
    }
    v69 := GenS10{F0: 11, F1: (v2 / (((v2 + v2)) | 1)), F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v63
    _ = v67
    _ = v69
}
