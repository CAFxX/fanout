package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS6 struct { F0 int32; F1 int16 }

func main() {
    var v1 uint64 = 0
    var v2 uint64 = (v1 * (v1 + 109))
    var v3 int64 = 0
    v4 := []string{"", "hello", "hello", "\n\t"}
    v5 := make(map[uint32]bool)
    var v6 float64 = 3.14159
    var v7 uint8 = 133
    fmt.Println(v1, v2, v1)
    var v8 float64 = (v6 / (v6 / (float64((1.5 / 1.0)) + v6)))
    v4 = append(v4, (("hello" + "foo bar") + "foo bar"))
    if (255 <= 18446744073709551615) {
        switch "hello" {
        case "\n\t":
            var v9 int32 = 517
            _ = v9
        default:
            if len(v4) > 0 {
                v4[0] = "hello"
            }
        }
        var v10 uint16 = 65535
        switch 32768 {
        case 32767:
            fmt.Println(v2)
        case 1:
            for fi11 := 0; fi11 < 9; fi11++ {
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
                _ = fi11
            }
        case 618:
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 string = "a"
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi12
                _ = v13
            }
        default:
            switch -2147483648 {
            case 9223372036854775807:
                if true {
                    var v14 string = "a"
                    _ = v14
                } else {
                    var v15 uint32 = 810
                    if false {
                        var v16 uint64 = 873
                        v17 := GenS1{F0: 32768}
                        _ = v16
                        _ = v17
                    } else {
                        if (9223372036854775807 >= 32768) {
                            var v18 uint16 = 255
                            v8 += v8
                            _ = v18
                        }
                        if len(v4) > 0 {
                            v4[0] = ("" + ("" + ("foo bar" + ("" + "a"))))
                        }
                        if (uint16(v1) != (^(v10))) {
                            switch ("\n\t" + (("\n\t" + "") + (("hello" + (("\n\t" + ("a" + ("hello" + ""))) + ("hello" + "a"))) + "\n\t"))) {
                            case "a":
                                for fi19 := 0; fi19 < 14; fi19++ {
                                    var v20 uint8 = (v7 / ((v7) | 1))
                                    _ = fi19
                                    _ = v20
                                }
                            case "\n\t":
                                if false {
                                    var v21 int16 = -1
                                    _ = v21
                                } else {
                                    fmt.Println(v1)
                                    delete(v5, 1)
                                }
                            default:
                                v22 := GenS1{F0: 32767}
                                _ = v22
                            }
                        }
                    }
                    for fi23 := 0; fi23 < 11; fi23++ {
                        var v24 bool = false
                        _ = len(v4)
                        fmt.Println(v6)
                        _ = fi23
                        _ = v24
                    }
                    _ = v15
                }
            case -2147483648:
                var v25 int8 = 127
                _ = v25
            default:
                if true {
                    if false {
                        var v26 int = 32768
                        v26 += (v26 * v26)
                        for fi27 := 0; fi27 < 19; fi27++ {
                            if (v1 == v2) {
                                v2 = v2
                                fmt.Println(v7)
                                var v28 string = "foo bar"
                                _ = v28
                            } else {
                                v29 := GenS1{F0: (^(v26))}
                                for fi30 := 0; fi30 < 15; fi30++ {
                                    switch v26 {
                                    case 4294967295:
                                        var v31 bool = true
                                        _ = v31
                                    case -1:
                                        v32 := GenS1{F0: (v26 * (v26 >> ((346) & 63)))}
                                        _ = v32
                                    }
                                    var v33 float64 = float64(((-2.25 / 1.0) * 0.0))
                                    _ = fi30
                                    _ = v33
                                }
                                _ = v29
                            }
                            if true {
                                v34 := GenS1{F0: -1}
                                if (("" <= "a") || true) {
                                    var v35 bool = false
                                    v36 := GenS2{F0: -802, F1: (v3 + -482), F2: 42}
                                    switch ("foo bar" + "") {
                                    case "a":
                                        fmt.Println(v10, v26, v8)
                                    case "":
                                        var v37 uint16 = (v10 | 65535)
                                        _ = v37
                                    case "\n\t":
                                        fmt.Println(v26)
                                    }
                                    _ = v35
                                    _ = v36
                                } else {
                                    v1 += v1
                                }
                                _ = v34
                            }
                            _ = len(v4)
                            _ = fi27
                        }
                        _ = v26
                    }
                    for fi38 := 0; fi38 < 13; fi38++ {
                        _ = len(v5)
                        v10 += 42
                        fmt.Println(v10, v3, v6)
                        _ = fi38
                    }
                    v39 := GenS3{F0: v10, F1: -284, F2: (v10 >> ((384) & 63))}
                    _ = v39
                }
            }
        }
        _ = v10
    }
    v2 = 840
    var v40 float64 = 0.0
    if (true && true) {
        var v41 int = 4294967295
        var v42 int = (v41 * v41)
        v43 := GenS1{F0: v42}
        _ = v41
        _ = v42
        _ = v43
    } else {
        for fi44 := 0; fi44 < 19; fi44++ {
            var v45 float32 = 1.5
            v8 = (v40 + ((v6 + float64((0.1 - 0.1))) / ((0.0 / ((v6 + v8) * (((-2.25 / v8) + (v8 - v40)) - v8))) * v8)))
            v8 = float64(((v45 / (v45 - v45)) * v45))
            _ = fi44
            _ = v45
        }
        fmt.Println(v1, v7, v8)
        for fi46 := 0; fi46 < 8; fi46++ {
            var v47 int = fi46
            v8 = v40
            if (1 < uint32(v7)) {
                var v48 int8 = 100
                if ((v1 % ((4294967295) | 1)) < (v2 << ((255) & 63))) {
                    var v49 int64 = (v3 ^ (-(v3)))
                    var v50 float64 = v6
                    _ = v49
                    _ = v50
                } else {
                    v7 = v7
                    v8 = float64(((0.1 + (1e10 + 0.0)) + (((1e10 / 3.14159) - 0.0) * 3.14159)))
                }
                v51 := GenS1{F0: -2147483648}
                _ = v48
                _ = v51
            }
            _ = fi46
            _ = v47
        }
    }
    var v52 int8 = 8
    v53 := GenS4{F0: 32767, F1: v7}
    switch int(v1) {
    case 0:
        var v54 int16 = -572
        _ = v54
    case 139:
        _ = len(v4)
    case 2147483647:
        for fi55 := 0; fi55 < 6; fi55++ {
            v7 = (v7 ^ (v7 + v7))
            if true {
                for fi56 := 0; fi56 < 7; fi56++ {
                    v57 := GenS5{F0: 42, F1: 18446744073709551615, F2: (v7 & 0)}
                    _ = fi56
                    _ = v57
                }
                v58 := GenS6{F0: 1, F1: 32767}
                var v59 uint8 = (v7 ^ (v7 * (v7 + v7)))
                _ = v58
                _ = v59
            } else {
                switch "a" {
                case "a":
                    v6 += float64(float32((-2.25 / 1.0)))
                case "hello":
                    v4 = append(v4, (("hello" + "hello") + ""))
                case "\n\t":
                    var v60 float64 = 3.14159
                    _ = v60
                }
                fmt.Println(v3, fi55, v52)
            }
            _ = fi55
        }
    }
    var v61 uint16 = 86
    v1 += (v1 * (^(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v52
    _ = v53
    _ = v6
    _ = v61
    _ = v7
    _ = v8
}
