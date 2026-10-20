package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 string; F1 uint32; F2 int16 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 bool; F2 uint }
type GenS6 struct { F0 uint }

func main() {
    v1 := []uint32{}
    var v2 uint64 = 0
    v3 := []string{("foo bar" + ("\n\t" + "a")), "foo bar", "hello", ""}
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := GenS1{F0: 0, F1: fi4}
        var v6 uint8 = uint8(v2)
        _ = fi4
        _ = v5
        _ = v6
    }
    v2 -= v2
    fmt.Println(v2)
    var v7 int16 = 32767
    var v8 int16 = v7
    v9 := GenS1{F0: 255, F1: 2147483647}
    v10 := GenS2{F0: "a", F1: 0, F2: v7}
    if true {
        var v11 float32 = 0.0
        _ = v11
    } else {
        for fi12 := 0; fi12 < 13; fi12++ {
            fmt.Println(fi12, v7, v2)
            v13 := GenS2{F0: "a", F1: 780, F2: int16(fi12)}
            _ = fi12
            _ = v13
        }
    }
    v14 := GenS2{F0: "\n\t", F1: 42, F2: -79}
    switch "foo bar" {
    case "foo bar":
        v2 -= 0
    case "":
        v15 := GenS2{F0: "\n\t", F1: 0, F2: (v7 * (v7 << ((975) & 63)))}
        _ = v15
    case "a":
        v8 = (v7 % ((v7) | 1))
    default:
        fmt.Println(v7, v8, v8)
    }
    if (!(65535 > 65535)) {
        v16 := GenS1{F0: 128, F1: 32768}
        v7 = 194
        _ = v16
    }
    v1 = append(v1, 42)
    if (true || (42 >= v7)) {
        if false {
            v17 := GenS3{F0: "foo bar", F1: 1}
            v8 -= (v8 >> ((2) & 63))
            _ = v17
        }
        if (true && ((2147483647 <= 2147483647) && (true && false))) {
            v8 -= (v7 + v8)
            switch 4294967295 {
            case 2147483647:
                v8 += v8
            case -2147483648:
                if false {
                    var v18 int = 821
                    fmt.Println(v8)
                    v19 := GenS4{F0: ""}
                    _ = v18
                    _ = v19
                }
            }
            v2 -= 240
        } else {
            v7 = v8
            if false {
                if (((1e10 + float32(((0.0 - (1.5 / -2.25)) * (0.1 / 1.0)))) / 1e10) < (float32(0.1) + float32(1.5))) {
                    if (0.0 != ((0.1 * 1.5) * 3.14159)) {
                        if ((1.5 * 1e10) <= 0.0) {
                            if true {
                                switch 32767 {
                                case 127:
                                    var v20 string = ("a" + "a")
                                    _ = v20
                                default:
                                    var v21 int16 = (v8 / (((v8 / (((v8 - 0)) | 1))) | 1))
                                    _ = v21
                                }
                                var v22 uint64 = (v2 + 291)
                                if true {
                                    v23 := GenS5{F0: 213, F1: false, F2: 42}
                                    var v24 int32 = 2147483647
                                    var v25 uint = 2
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                } else {
                                    v26 := GenS2{F0: ("" + "\n\t"), F1: 2147483648, F2: v7}
                                    for fi27 := 0; fi27 < 16; fi27++ {
                                        if (42 <= 42) {
                                            var v28 string = "\n\t"
                                            _ = v28
                                        }
                                        _ = fi27
                                    }
                                    _ = v26
                                }
                                _ = v22
                            }
                        }
                        var v29 float32 = (0.0 * 0.1)
                        if len(v1) > 0 {
                            v1[0] = 293
                        }
                        _ = v29
                    } else {
                        if true {
                            for fi30 := 0; fi30 < 6; fi30++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if (2147483648 > 0) {
                                    for fi31 := 0; fi31 < 20; fi31++ {
                                        fmt.Println(v8, v2, fi30)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if false {
                                            if false {
                                                v32 := GenS3{F0: ("foo bar" + "foo bar"), F1: 887}
                                                _ = v32
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if (!false) {
                                                var v33 string = "\n\t"
                                                _ = v33
                                            } else {
                                                v34 := GenS6{F0: 65535}
                                                var v35 int = (-(fi31))
                                                _ = v34
                                                _ = v35
                                            }
                                            switch "" {
                                            case "foo bar":
                                                if len(v3) > 0 {
                                                    v3[0] = "hello"
                                                }
                                            default:
                                                v8 -= (v8 << ((2) & 63))
                                            }
                                        }
                                        _ = fi31
                                    }
                                }
                                _ = fi30
                            }
                        } else {
                            fmt.Println(v8)
                        }
                    }
                    fmt.Println(v8, v8)
                }
            } else {
                _ = len(v3)
                var v36 bool = (("hello" <= "\n\t") && (true && (!(3.14159 == (3.14159 / (1.5 / (1e10 / 0.1)))))))
                _ = len(v1)
                _ = v36
            }
            for fi37 := 0; fi37 < 3; fi37++ {
                fmt.Println(v8)
                v7 -= (v8 * (v7 << ((7) & 63)))
                fmt.Println(v8)
                _ = fi37
            }
        }
        if (!(!false)) {
            v38 := GenS2{F0: "hello", F1: 611, F2: (v8 + (v7 | v7))}
            if true {
                v39 := GenS2{F0: ("" + "a"), F1: 4294967295, F2: (v8 / (((v7 | (-(v7)))) | 1))}
                for fi40 := 0; fi40 < 6; fi40++ {
                    for fi41 := 0; fi41 < 12; fi41++ {
                        var v42 uint8 = 63
                        var v43 int32 = -1
                        var v44 string = "hello"
                        _ = fi41
                        _ = v42
                        _ = v43
                        _ = v44
                    }
                    v3 = append(v3, "a")
                    _ = fi40
                }
                _ = v39
            }
            _ = v38
        }
    }
    fmt.Println(v8, v2)
    var v45 float64 = (float64(1.5) * 0.0)
    switch 2147483647 {
    case 127:
        if len(v1) > 0 {
            v1[0] = 42
        }
    case 755:
        var v46 float32 = 0.0
        _ = v46
    case 592:
        switch "a" {
        case "":
            v45 = v45
        case "\n\t":
            for fi47 := 0; fi47 < 15; fi47++ {
                v45 = ((v45 / 3.14159) / 3.14159)
                if len(v1) > 0 {
                    v1[0] = 842
                }
                var v48 uint64 = (v2 ^ (v2 * v2))
                _ = fi47
                _ = v48
            }
        case "hello":
            v45 = (1e10 + v45)
        }
    default:
        if (false && (!false)) {
            for fi49 := 0; fi49 < 16; fi49++ {
                v8 -= v7
                _ = fi49
            }
            var v50 int16 = (-(v8))
            _ = v50
        }
    }
    var v51 int = 750
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v45:", v45)
    fmt.Println("v51:", v51)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v45
    _ = v51
    _ = v7
    _ = v8
    _ = v9
}
