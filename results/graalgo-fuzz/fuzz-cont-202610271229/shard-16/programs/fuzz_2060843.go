package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 int; F1 uint64; F2 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint16; F1 bool; F2 int }
func genFn0(p0 int, p1 uint, p2 string) uint16 {
    v5 := GenS1{F0: 0, F1: 144, F2: (^(p0))}
    v6 := []int{}
    v7 := GenS2{F0: (^(p1))}
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 32767
}

func main() {
    var v1 bool = ((219 != 255) || (244 > 128))
    v2 := []uint32{65535, 348, 42, 1}
    v3 := []int{9223372036854775807, 127, 32768}
    var v4 bool = v1
    var v8 uint8 = 124
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS3{F0: 1}
        _ = fi9
        _ = v10
    }
    switch "a" {
    case "\n\t":
        var v11 uint32 = 1
        _ = v11
    default:
        if true {
            var v12 float64 = 1.5
            _ = v12
        } else {
            v1 = v4
        }
    }
    var v13 int8 = -1
    v1 = (genFn0(-656, 2, ("hello" + "foo bar")) >= 11)
    switch ("foo bar" + ("a" + "")) {
    case "":
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := GenS3{F0: -32768}
            _ = fi14
            _ = v15
        }
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v4 = v4
    _ = len(v3)
    if (-2.25 <= 1e10) {
        v2 = append(v2, 1)
    } else {
        v4 = (v13 != (v13 ^ v13))
        if ((v8 < (v8 / (((v8 | v8)) | 1))) || (v4 || v1)) {
            switch 326 {
            case -19:
                for fi16 := 0; fi16 < 15; fi16++ {
                    v4 = true
                    _ = fi16
                }
            case 42:
                v17 := GenS4{F0: v1, F1: genFn0(9223372036854775807, 42, "foo bar")}
                _ = v17
            case 32768:
                switch "" {
                case "hello":
                    fmt.Println(v1)
                case "":
                    var v18 bool = (1e10 != 0.1)
                    _ = v18
                case "\n\t":
                    switch "foo bar" {
                    case "a":
                        if v4 {
                            fmt.Println(v1, v8, v13)
                            for fi19 := 0; fi19 < 12; fi19++ {
                                v13 = 1
                                for fi20 := 0; fi20 < 9; fi20++ {
                                    v8 -= (^(v8))
                                    _ = len(v3)
                                    var v21 int64 = 1
                                    _ = fi20
                                    _ = v21
                                }
                                _ = fi19
                            }
                        }
                    default:
                        var v22 int32 = 2147483647
                        _ = v22
                    }
                }
            default:
                var v23 uint64 = 1
                _ = v23
            }
            if (65535 == 758) {
                fmt.Println(v4, v8, v8)
                v4 = true
            } else {
                for fi24 := 0; fi24 < 9; fi24++ {
                    _ = len(v3)
                    _ = fi24
                }
                fmt.Println(v4, v1)
                v13 = -128
            }
        }
    }
    switch 32767 {
    case 313:
        switch "" {
        case "\n\t":
            _ = len(v2)
        case "foo bar":
            switch "\n\t" {
            case "a":
                fmt.Println(v4)
            }
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    default:
        switch "\n\t" {
        case "":
            _ = len(v3)
        case "a":
            fmt.Println(v13, v13, v13)
        case "foo bar":
            if false {
                v13 -= v13
            } else {
                v4 = (!v1)
                if v1 {
                    switch "foo bar" {
                    case "\n\t":
                        if v4 {
                            if (-902 >= 861) {
                                for fi25 := 0; fi25 < 11; fi25++ {
                                    v13 = v13
                                    for fi26 := 0; fi26 < 8; fi26++ {
                                        v8 -= v8
                                        var v27 int16 = -889
                                        var v28 bool = (1e10 < 1.5)
                                        _ = fi26
                                        _ = v27
                                        _ = v28
                                    }
                                    _ = fi25
                                }
                                var v29 float64 = 0.1
                                _ = v29
                            }
                            v30 := GenS5{F0: 0, F1: 0, F2: -1}
                            _ = v30
                        }
                    case "hello":
                        _ = len(v2)
                    }
                    v1 = (42 != -437)
                    var v31 string = "hello"
                    _ = v31
                }
            }
        default:
            _ = len(v3)
        }
    }
    switch 32768 {
    case 2147483647:
        if true {
            var v32 uint8 = (v8 << ((255) & 63))
            for fi33 := 0; fi33 < 10; fi33++ {
                v34 := GenS6{F0: 847}
                _ = fi33
                _ = v34
            }
            _ = v32
        } else {
            var v35 uint8 = 255
            _ = v35
        }
    case -13:
        var v36 int16 = 1
        _ = v36
    }
    if len(v2) > 0 {
        v2[0] = 65535
    }
    for fi37 := 0; fi37 < 2; fi37++ {
        var v38 uint = 255
        var v39 uint32 = 0
        _ = fi37
        _ = v38
        _ = v39
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v4, v4, v1)
    case "hello":
        var v40 float32 = 1.5
        _ = v40
    case "a":
        if len(v3) > 0 {
            v3[0] = -242
        }
    default:
        if v1 {
            if true {
                v4 = false
                _ = len(v3)
            } else {
                v41 := GenS7{F0: 970, F1: (0.0 < 3.14159), F2: 128}
                var v42 uint32 = uint32(v8)
                v13 -= (v13 >> ((422) & 63))
                _ = v41
                _ = v42
            }
            v4 = ("a" > "foo bar")
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    for fi43 := 0; fi43 < 20; fi43++ {
        if (false || v4) {
            for fi44 := 0; fi44 < 13; fi44++ {
                switch -894 {
                case 2147483647:
                    var v45 int64 = 687
                    _ = v45
                case 9223372036854775807:
                    var v46 int8 = -43
                    _ = v46
                case -326:
                    if (-7 != -2147483648) {
                        v1 = true
                        var v47 int64 = 0
                        _ = v47
                    } else {
                        v1 = (v4 || false)
                        if (2147483647 >= 42) {
                            var v48 uint32 = 65535
                            _ = v48
                        }
                        var v49 int16 = -595
                        _ = v49
                    }
                }
                _ = fi44
            }
        } else {
            var v50 string = "a"
            _ = v50
        }
        _ = fi43
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
