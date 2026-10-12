package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint; F1 int16; F2 uint16 }
type GenS3 struct { F0 string; F1 int16; F2 int32 }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int8; F1 bool }
func genFn0() uint64 {
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        v6 := []int{-9223372036854775808, 967}
        _ = v6
    case "a":
        v7 := GenS1{F0: 381, F1: 855, F2: 127}
        _ = v7
    case "":
        var v8 bool = false
        _ = v8
    default:
        var v9 int = -9223372036854775808
        _ = v9
    }
    var v10 float64 = (0.1 - -2.25)
    _ = v10
    return 18446744073709551615
}
func genFn1() uint32 {
    v11 := GenS2{F0: 4294967295, F1: 0, F2: 65535}
    var v12 float64 = 0.1
    for fi13 := 0; fi13 < 1; fi13++ {
        switch ("foo bar" + "\n\t") {
        case "a":
            v14 := GenS2{F0: 373, F1: 811, F2: 344}
            _ = v14
        case "\n\t":
            v12 = ((v12 + (((v12 - v12) / ((1.5 + (float64(-2.25) * 0.1)) - float64(1.5))) * ((v12 - (v12 - ((v12 - 0.0) / ((((v12 - (float64(1e10) * -2.25)) / (v12 - float64((1.5 / 1e10)))) + 0.1) + 1.5)))) * 1.5))) + v12)
        case "":
            if (32767 > 65535) {
                for fi15 := 0; fi15 < 5; fi15++ {
                    if (!true) {
                        fmt.Println(fi15, v12, fi13)
                        if false {
                            if (-719 >= -781) {
                                v12 += 1.5
                            } else {
                                v12 += float64((1.5 / (1e10 + 1e10)))
                                if (!(0 > 32767)) {
                                    v16 := []int{}
                                    _ = v16
                                }
                            }
                            fmt.Println(v12)
                            if false {
                                var v17 int32 = 42
                                v17 = v17
                                _ = v17
                            }
                        }
                        v12 = (3.14159 / (v12 - v12))
                    }
                    fmt.Println(fi13)
                    var v18 uint8 = 42
                    _ = fi15
                    _ = v18
                }
                v19 := []int{fi13, fi13, fi13}
                _ = v19
            }
        default:
            v12 = -2.25
        }
        v12 -= float64(0.0)
        v20 := []int{fi13, (fi13 & fi13)}
        _ = fi13
        _ = v20
    }
    v12 = v12
    _ = v11
    _ = v12
    return 595
}

func main() {
    v1 := []int{1, -781, 0}
    v2 := make(map[uint32]uint32)
    var v3 uint = 65535
    v4 := []uint32{775, 42, 1}
    var v5 int16 = 523
    if len(v4) > 0 {
        v4[0] = 976
    }
    switch -399 {
    case 32767:
        var v21 uint32 = 834
        _ = v21
    case 1:
        v5 -= v5
    case 9223372036854775807:
        for fi22 := 0; fi22 < 10; fi22++ {
            if (!("" < "foo bar")) {
                fmt.Println(v3)
                var v23 int32 = 0
                _ = v23
            }
            _ = fi22
        }
    default:
        var v24 float32 = 0.0
        _ = v24
    }
    v25 := GenS3{F0: (("hello" + "hello") + "hello"), F1: v5, F2: 0}
    v2[2147483648] = 65535
    var v26 float64 = 0.0
    var v27 uint32 = genFn1()
    for fi28 := 0; fi28 < 7; fi28++ {
        v29 := GenS2{F0: 65535, F1: v5, F2: 255}
        if false {
            switch "a" {
            case "a":
                fmt.Println(v3, v26, v27)
            }
            var v30 uint32 = (v27 + v27)
            fmt.Println(v27, v30, v26)
            _ = v30
        } else {
            v31 := GenS4{F0: ("foo bar" + ""), F1: fi28}
            if true {
                v32 := GenS5{F0: v5}
                v5 -= (v5 + v5)
                _ = v32
            }
            var v33 uint = v3
            _ = v31
            _ = v33
        }
        _ = fi28
        _ = v29
    }
    for fi34 := 0; fi34 < 3; fi34++ {
        var v35 uint64 = 42
        for fi36 := 0; fi36 < 1; fi36++ {
            v37 := GenS4{F0: ("hello" + "foo bar"), F1: 4294967295}
            _ = len(v1)
            if (!true) {
                var v38 int16 = v5
                _ = v38
            } else {
                v26 = (v26 + v26)
                v3 = v3
                v39 := GenS5{F0: 1000}
                _ = v39
            }
            var v40 float32 = (-2.25 - 1e10)
            switch (fi36 << ((v3) & 63)) {
            case 127:
                delete(v2, (v27 % ((4294967295) | 1)))
            case -2147483648:
                v2[779] = v27
            case 0:
                var v41 int32 = 1
                _ = v41
            default:
                switch "a" {
                case "hello":
                    if (true || (0 > 1)) {
                        v26 -= float64((v40 + v40))
                    }
                default:
                    _ = len(v4)
                }
            }
            var v42 bool = false
            v42 = v42
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch (fi34 + 127) {
            case 456:
                fmt.Println(fi34)
            case 9223372036854775807:
                v43 := GenS6{F0: v42, F1: -1}
                _ = v43
            case 4294967295:
                v44 := GenS7{F0: 1, F1: 1, F2: v27}
                _ = v44
            default:
                if ((v35 >= 358) && (255 < 0)) {
                    v27 += 2
                } else {
                    v1 = append(v1, 42)
                    switch (fi34 << ((v3) & 63)) {
                    case -587:
                        var v45 float32 = float32(3.14159)
                        _ = v45
                    case 2147483647:
                        v40 -= v40
                    case 756:
                        switch ("" + "\n\t") {
                        case "foo bar":
                            switch (("foo bar" + "\n\t") + "") {
                            case "foo bar":
                                v5 += (v5 << ((42) & 63))
                            case "hello":
                                fmt.Println(v40, v42)
                            case "":
                                switch fi34 {
                                case -1:
                                    v46 := GenS4{F0: ((("foo bar" + "a") + ("a" + "\n\t")) + "hello"), F1: fi34}
                                    _ = v46
                                case 2147483647:
                                    var v47 uint = v3
                                    _ = v47
                                default:
                                    var v48 int16 = v5
                                    _ = v48
                                }
                            default:
                                var v49 uint64 = (v35 % ((v35) | 1))
                                _ = v49
                            }
                        default:
                            fmt.Println(v27)
                            var v50 uint16 = 65535
                            _ = v50
                        }
                    }
                    if true {
                        switch 32767 {
                        case 32768:
                            var v51 uint8 = 127
                            _ = v51
                        case 873:
                            v42 = v42
                        case 697:
                            var v52 int64 = -476
                            _ = v52
                        default:
                            v3 = v3
                        }
                        v53 := v2[v27]
                        _ = v53
                    }
                    var v54 float64 = -2.25
                    v26 = (v26 * v54)
                    _ = v54
                }
            }
            _ = fi36
            _ = v37
            _ = v40
            _ = v42
        }
        _ = fi34
        _ = v35
    }
    v26 += float64(1.5)
    for fi55 := 0; fi55 < 15; fi55++ {
        var v56 uint16 = 32767
        _ = fi55
        _ = v56
    }
    v57 := GenS8{F0: 32767}
    switch "hello" {
    case "hello":
        v58 := GenS4{F0: "foo bar", F1: 0}
        _ = v58
    case "a":
        for fi59 := 0; fi59 < 8; fi59++ {
            switch fi59 {
            case 32768:
                switch "hello" {
                case "\n\t":
                    var v60 int16 = v5
                    _ = v60
                case "hello":
                    if false {
                        var v61 uint8 = 255
                        var v62 float32 = (1e10 - 1.5)
                        for fi63 := 0; fi63 < 16; fi63++ {
                            var v64 bool = ((v3 - v3) == v3)
                            _ = fi63
                            _ = v64
                        }
                        _ = v61
                        _ = v62
                    }
                }
            }
            v5 = v5
            _ = fi59
        }
    case "\n\t":
        if false {
            for fi65 := 0; fi65 < 11; fi65++ {
                v2[(v27 >> ((65535) & 63))] = 4294967295
                fmt.Println(fi65)
                for fi66 := 0; fi66 < 19; fi66++ {
                    fmt.Println(v26)
                    _ = fi66
                }
                _ = fi65
            }
        } else {
            v27 = uint32(v3)
            fmt.Println(v3, v26, v27)
            v67 := GenS7{F0: -1, F1: 220, F2: (v27 >> (((v3 ^ 255)) & 63))}
            _ = v67
        }
    default:
        if (!false) {
            v68 := GenS9{F0: 127, F1: (!false)}
            _ = v68
        }
    }
    var v69 int16 = 0
    var v70 uint64 = genFn0()
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v69
    _ = v70
}
