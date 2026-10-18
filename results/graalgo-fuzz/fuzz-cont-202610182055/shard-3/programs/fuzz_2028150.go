package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int8; F2 int8 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 uint8; F2 int8 }
type GenS6 struct { F0 int16; F1 int }

func main() {
    var v1 int32 = -455
    var v2 float64 = 0.1
    var v3 int64 = -986
    var v4 int = 42
    var v5 bool = true
    var v6 float32 = (1e10 + -2.25)
    switch ("\n\t" + "") {
    case "hello":
        if v5 {
            var v7 int32 = v1
            switch "\n\t" {
            case "a":
                v7 = 2147483647
            case "\n\t":
                v8 := []int{int(v3)}
                _ = v8
            }
            _ = v7
        } else {
            if (242 > -32768) {
                if true {
                    v9 := GenS1{F0: (-(v4))}
                    v3 += v3
                    var v10 string = ("hello" + "\n\t")
                    _ = v10
                    _ = v9
                } else {
                    var v11 uint64 = 18446744073709551615
                    _ = v11
                }
                v1 = (v1 | 0)
                fmt.Println(v5, v5)
            }
        }
    case "\n\t":
        v12 := []int{-1, v4, (v4 ^ (v4 * v4)), (v4 - v4)}
        _ = v12
    case "":
        fmt.Println(v1, v4)
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := []int{32767, (v4 % ((v4) | 1)), v4}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 bool = v5
        _ = fi15
        _ = v16
    }
    if true {
        if (v5 && v5) {
            fmt.Println(v6)
            var v17 int = (^(v4))
            fmt.Println(v3, v17)
            _ = v17
        } else {
            for fi18 := 0; fi18 < 7; fi18++ {
                v19 := []int{fi18}
                if true {
                    for fi20 := 0; fi20 < 17; fi20++ {
                        v1 += v1
                        v21 := GenS1{F0: (fi20 / (((fi20 % ((v4) | 1))) | 1))}
                        _ = fi20
                        _ = v21
                    }
                } else {
                    v5 = v5
                    v2 = 1.5
                }
                v22 := GenS2{F0: 42, F1: 82, F2: -128}
                _ = fi18
                _ = v19
                _ = v22
            }
            v23 := []int{(v4 * v4), (v4 * (v4 << ((42) & 63)))}
            fmt.Println(v5, v4)
            _ = v23
        }
    }
    v24 := []int{}
    if (0.0 != 1e10) {
        if v5 {
            fmt.Println(v3)
            var v25 int32 = 42
            _ = v25
        }
        var v26 float32 = v6
        v27 := GenS1{F0: -2147483648}
        _ = v26
        _ = v27
    } else {
        for fi28 := 0; fi28 < 13; fi28++ {
            for fi29 := 0; fi29 < 10; fi29++ {
                switch int(v1) {
                case 4294967295:
                    v24 = append(v24, 1)
                default:
                    v30 := GenS2{F0: 1, F1: 42, F2: 3}
                    _ = v30
                }
                v6 = v6
                v3 += 0
                var v31 float32 = (v6 + (v6 + v6))
                _ = fi29
                _ = v31
            }
            _ = fi28
        }
    }
    if v5 {
        for fi32 := 0; fi32 < 3; fi32++ {
            if v5 {
                v3 += v3
                for fi33 := 0; fi33 < 2; fi33++ {
                    v3 = v3
                    _ = fi33
                }
                if v5 {
                    v2 = v2
                    v2 += v2
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                } else {
                    if false {
                        _ = len(v24)
                        switch "\n\t" {
                        case "hello":
                            v34 := GenS1{F0: (fi32 ^ (-(v4)))}
                            _ = v34
                        }
                    } else {
                        switch v4 {
                        case -1:
                            v6 += 3.14159
                        case 2147483647:
                            switch "\n\t" {
                            case "hello":
                                for fi35 := 0; fi35 < 20; fi35++ {
                                    var v36 uint32 = 42
                                    _ = fi35
                                    _ = v36
                                }
                            case "":
                                var v37 uint16 = 42
                                _ = v37
                            }
                        case 9223372036854775807:
                            if (!true) {
                                if len(v24) > 0 {
                                    v24[0] = (fi32 << ((790) & 63))
                                }
                            }
                        }
                    }
                    var v38 uint8 = 128
                    v39 := GenS3{F0: 640, F1: uint32(v38)}
                    _ = v38
                    _ = v39
                }
            } else {
                fmt.Println(v6, v5, v1)
                var v40 int16 = 1
                _ = v40
            }
            for fi41 := 0; fi41 < 6; fi41++ {
                v42 := GenS4{F0: fi41}
                var v43 string = "a"
                if v5 {
                    switch fi41 {
                    case -1:
                        switch -45 {
                        case 9223372036854775807:
                            v44 := GenS5{F0: -127, F1: 0, F2: -128}
                            _ = v44
                        default:
                            if len(v24) > 0 {
                                v24[0] = -1
                            }
                        }
                    case 32768:
                        switch v43 {
                        case "foo bar":
                            var v45 uint16 = 435
                            _ = v45
                        case "":
                            if ((fi32 ^ (^(v4))) < -1) {
                                var v46 uint16 = 2
                                _ = v46
                            }
                        }
                        fmt.Println(v5, v6)
                        v47 := GenS4{F0: 9223372036854775807}
                        var v48 bool = v5
                        _ = v47
                        _ = v48
                    case -2147483648:
                        v49 := GenS1{F0: fi41}
                        _ = v49
                    default:
                        v4 = (v4 & fi32)
                    }
                } else {
                    v6 -= v6
                }
                if (-1 < v3) {
                    if (v5 && v5) {
                        v5 = (523 < 42)
                        if len(v24) > 0 {
                            v24[0] = 32767
                        }
                    }
                }
                var v50 uint16 = 42
                _ = fi41
                _ = v42
                _ = v43
                _ = v50
            }
            var v51 float32 = 1e10
            _ = fi32
            _ = v51
        }
        if len(v24) > 0 {
            v24[0] = (v4 + 611)
        }
        for fi52 := 0; fi52 < 1; fi52++ {
            fmt.Println(v6, v4, v1)
            var v53 uint32 = 4294967295
            var v54 int8 = -49
            _ = fi52
            _ = v53
            _ = v54
        }
    }
    for fi55 := 0; fi55 < 8; fi55++ {
        v3 += (v3 << ((139) & 63))
        _ = fi55
    }
    v2 = 0.0
    fmt.Println(v5, v6)
    if (v5 || (!false)) {
        _ = len(v24)
        switch "hello" {
        case "\n\t":
            if false {
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
                v3 -= v3
                fmt.Println(v3, v6)
            }
        case "hello":
            for fi56 := 0; fi56 < 2; fi56++ {
                for fi57 := 0; fi57 < 2; fi57++ {
                    if len(v24) > 0 {
                        v24[0] = (-(v4))
                    }
                    _ = fi57
                }
                _ = fi56
            }
        }
    } else {
        switch v4 {
        case 710:
            if (-128 != 1) {
                v5 = v5
                v58 := GenS4{F0: (v4 + 32767)}
                _ = v58
            }
        case 32768:
            if v5 {
                v6 = 1e10
                for fi59 := 0; fi59 < 13; fi59++ {
                    v5 = (0 == 2147483647)
                    var v60 uint64 = 2
                    _ = fi59
                    _ = v60
                }
            }
        case 2147483647:
            v5 = (32767 <= -1)
        default:
            _ = len(v24)
        }
        if (100 != int8(v3)) {
            var v61 int32 = v1
            var v62 uint16 = 73
            _ = v61
            _ = v62
        }
        for fi63 := 0; fi63 < 5; fi63++ {
            for fi64 := 0; fi64 < 13; fi64++ {
                fmt.Println(v2, v4)
                _ = len(v24)
                if len(v24) > 0 {
                    v24[0] = fi63
                }
                _ = fi64
            }
            _ = fi63
        }
    }
    var v65 int16 = -904
    v24 = append(v24, (v4 >> ((143) & 63)))
    var v66 uint64 = 18446744073709551615
    switch v4 {
    case -9223372036854775808:
        v67 := GenS6{F0: (v65 % ((v65) | 1)), F1: (v4 * (v4 & v4))}
        _ = v67
    case 2147483647:
        var v68 int32 = (-(v1))
        _ = v68
    case 1:
        for fi69 := 0; fi69 < 13; fi69++ {
            v65 -= v65
            v1 -= -2
            _ = fi69
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
}
