package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int32; F1 int64; F2 int8 }
type GenS5 struct { F0 int32; F1 bool; F2 int }
type GenS6 struct { F0 uint; F1 uint32; F2 int32 }
type GenS7 struct { F0 int32; F1 uint8; F2 bool }
type GenS8 struct { F0 uint8 }
func genFn0(p0 uint64, p1 string) uint8 {
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := []int{fi5}
        if true {
            fmt.Println(fi5)
        } else {
            if true {
                if len(v6) > 0 {
                    v6[0] = fi5
                }
                for fi7 := 0; fi7 < 7; fi7++ {
                    fmt.Println(fi7, fi5, fi5)
                    switch fi5 {
                    case 1:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    case 735:
                        v8 := GenS1{F0: 1}
                        _ = v8
                    case 0:
                        var v9 uint32 = 42
                        _ = v9
                    default:
                        var v10 int32 = -845
                        _ = v10
                    }
                    v11 := GenS1{F0: 669}
                    _ = fi7
                    _ = v11
                }
            } else {
                v12 := GenS1{F0: 42}
                v13 := GenS2{F0: p0, F1: 385}
                var v14 int32 = 1
                _ = v12
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 7; fi15++ {
                v16 := GenS3{F0: 0, F1: 4294967295}
                _ = fi15
                _ = v16
            }
        }
        _ = len(v6)
        _ = fi5
        _ = v6
    }
    var v17 int = -79
    if false {
        var v18 string = ("foo bar" + "\n\t")
        v19 := GenS1{F0: 503}
        v20 := []int{}
        _ = v18
        _ = v19
        _ = v20
    } else {
        switch p1 {
        case "":
            v17 -= v17
        case "foo bar":
            var v21 int = (^(v17))
            _ = v21
        case "hello":
            v22 := []int{(v17 / ((0) | 1))}
            _ = v22
        default:
            var v23 int64 = 0
            _ = v23
        }
        switch (v17 & v17) {
        case 1:
            var v24 float64 = 3.14159
            _ = v24
        case -530:
            v25 := []int{967, (v17 | (-(v17))), 128, v17}
            _ = v25
        case 248:
            if false {
                if (0.0 <= 0.1) {
                    v26 := []int{v17, 4294967295, v17}
                    v27 := GenS4{F0: 0, F1: 0, F2: 100}
                    _ = v26
                    _ = v27
                }
                var v28 float32 = 3.14159
                _ = v28
            }
        }
        v29 := GenS5{F0: -837, F1: false, F2: v17}
        _ = v29
    }
    var v30 int8 = 26
    _ = p0
    _ = p1
    _ = v17
    _ = v30
    return 237
}
func genFn1() int64 {
    var v31 uint = 949
    switch "\n\t" {
    case "a":
        for fi32 := 0; fi32 < 20; fi32++ {
            v31 += 18446744073709551615
            v31 = v31
            v33 := GenS2{F0: 599, F1: 203}
            _ = fi32
            _ = v33
        }
    case "":
        var v34 float64 = (-2.25 * 0.1)
        _ = v34
    default:
        switch -44 {
        case 32768:
            fmt.Println(v31)
        case 127:
            v31 -= (^(v31))
        }
    }
    v31 = (^(v31))
    _ = v31
    return 1
}

func main() {
    var v1 float64 = (3.14159 * 3.14159)
    var v2 uint64 = 390
    var v3 int16 = -32768
    v4 := []uint32{65535, 2, 2, 1}
    v3 = (v3 * (v3 << ((2) & 63)))
    fmt.Println(v3, v2, v1)
    for fi35 := 0; fi35 < 17; fi35++ {
        switch fi35 {
        case 32768:
            fmt.Println(v2, v3, v1)
        }
        if false {
            v3 -= (v3 + (v3 | (v3 + v3)))
        } else {
            var v36 float64 = v1
            _ = v36
        }
        _ = fi35
    }
    switch -2147483648 {
    case 2147483647:
        for fi37 := 0; fi37 < 17; fi37++ {
            if ((v3 & v3) < v3) {
                v2 = v2
                for fi38 := 0; fi38 < 3; fi38++ {
                    if (935 != 2147483647) {
                        fmt.Println(v3)
                        if (true || (42 >= 1)) {
                            v39 := GenS4{F0: 2147483647, F1: 9223372036854775807, F2: -75}
                            var v40 uint32 = 0
                            v41 := GenS1{F0: 84}
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                    }
                    switch 518 {
                    case 127:
                        var v42 uint = 42
                        _ = v42
                    }
                    v1 += (((float64(1.5) + 3.14159) + v1) + 1.5)
                    v2 = 0
                    fmt.Println(v2, v3)
                    _ = fi38
                }
                for fi43 := 0; fi43 < 14; fi43++ {
                    switch "hello" {
                    case "":
                        if true {
                            var v44 uint8 = 110
                            switch -617 {
                            case 4294967295:
                                v44 = 65
                                v44 += 127
                                if false {
                                    fmt.Println(v3, fi37)
                                }
                                v1 += ((((((float64(float32((float64(0.1) / v1))) * v1) * float64(0.1)) * (v1 + v1)) / v1) / v1) + (3.14159 / (v1 + -2.25)))
                            default:
                                switch "a" {
                                case "foo bar":
                                    var v45 float32 = (0.1 * ((float32((float64((0.0 * (0.1 / ((1.5 * 3.14159) + 3.14159)))) * float64(3.14159))) - ((3.14159 - 1.5) / (((-2.25 / 1.0) * ((-2.25 * 3.14159) / 1.0)) * ((1e10 - 0.0) + (3.14159 - ((-2.25 + (float32(float64(0.1)) / 1.5)) - 0.1)))))) / 1.0))
                                    _ = v45
                                default:
                                    v2 = v2
                                }
                            }
                            fmt.Println(v2, v44, v44)
                            _ = v44
                        }
                    }
                    var v46 uint16 = 65535
                    _ = fi43
                    _ = v46
                }
            }
            v47 := GenS6{F0: 4294967295, F1: 432, F2: -121}
            _ = fi37
            _ = v47
        }
    case -2147483648:
        fmt.Println(v2, v3, v1)
    default:
        if (v3 < (v3 ^ v3)) {
            if (-313 == -994) {
                fmt.Println(v1, v2)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        } else {
            v48 := GenS4{F0: 42, F1: 1, F2: -1}
            if true {
                v2 = 344
                if (false || true) {
                    v49 := GenS6{F0: 42, F1: 320, F2: 918}
                    var v50 int32 = 42
                    switch 863 {
                    case 127:
                        switch "\n\t" {
                        case "":
                            v3 = (v3 + v3)
                        case "a":
                            var v51 uint = 42
                            _ = v51
                        case "\n\t":
                            var v52 uint16 = 1
                            _ = v52
                        default:
                            if (2 != 127) {
                                fmt.Println(v2, v3)
                                var v53 uint64 = v2
                                _ = v53
                            }
                        }
                    }
                    _ = v49
                    _ = v50
                } else {
                    fmt.Println(v3)
                    switch "\n\t" {
                    case "\n\t":
                        var v54 int16 = (v3 * (v3 | v3))
                        _ = v54
                    case "a":
                        for fi55 := 0; fi55 < 9; fi55++ {
                            fmt.Println(v3, fi55, fi55)
                            v4 = append(v4, 1)
                            _ = fi55
                        }
                    case "foo bar":
                        if (("a" + "\n\t") < "\n\t") {
                            v3 = 32767
                            switch 32768 {
                            case -272:
                                if (!(!true)) {
                                    var v56 uint64 = (^(v2))
                                    fmt.Println(v1, v2, v2)
                                    _ = v56
                                } else {
                                    if (42 >= 551) {
                                        switch ("foo bar" + (("a" + "hello") + ((("a" + "") + "a") + "\n\t"))) {
                                        case "hello":
                                            if (42 > 9223372036854775807) {
                                                switch -2147483648 {
                                                case 2147483647:
                                                    if (127 > 127) {
                                                        for fi57 := 0; fi57 < 3; fi57++ {
                                                            v1 += ((v1 - 3.14159) - ((((((v1 - (v1 - 0.1)) - v1) + (-2.25 * (3.14159 * v1))) * ((v1 / -2.25) + (float64(-2.25) - 1.5))) - (v1 + ((((v1 * v1) - (((((1.5 - v1) * (v1 / v1)) - 1e10) + v1) - (0.0 - (((float64(3.14159) - v1) * float64(float32((float64((1.5 * 1.5)) + 3.14159)))) - (float64(float32(3.14159)) * v1))))) * (((v1 - (((v1 / v1) - v1) - 1.5)) + v1) * v1)) + 1e10))) * (((v1 / v1) / v1) + 1.5)))
                                                            var v58 float32 = 0.1
                                                            _ = fi57
                                                            _ = v58
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        if (false || true) {
                                                            if len(v4) > 0 {
                                                                v4[0] = uint32(v2)
                                                            }
                                                            v1 = (v1 * v1)
                                                            if len(v4) > 0 {
                                                                v4[0] = 2147483648
                                                            }
                                                        } else {
                                                            _ = len(v4)
                                                            var v59 int = 127
                                                            v60 := GenS5{F0: int32(v3), F1: false, F2: v59}
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                    }
                                                }
                                                v61 := GenS7{F0: 168, F1: 0, F2: false}
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v61
                                            }
                                        default:
                                            _ = len(v4)
                                        }
                                        var v62 float32 = 3.14159
                                        _ = v62
                                    }
                                }
                            case -9223372036854775808:
                                var v63 uint = 65535
                                _ = v63
                            case 1:
                                switch "a" {
                                case "":
                                    for fi64 := 0; fi64 < 11; fi64++ {
                                        for fi65 := 0; fi65 < 2; fi65++ {
                                            v3 += 1000
                                            var v66 bool = (0 < 886)
                                            _ = fi65
                                            _ = v66
                                        }
                                        var v67 uint = 42
                                        _ = fi64
                                        _ = v67
                                    }
                                }
                            }
                            fmt.Println(v2, v1, v2)
                        } else {
                            v3 = v3
                        }
                    default:
                        fmt.Println(v1, v2, v1)
                    }
                    if (!(42 >= 2)) {
                        var v68 uint32 = 4294967295
                        _ = v68
                    }
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            var v69 uint = 874
            _ = v48
            _ = v69
        }
    }
    switch ("" + ("foo bar" + "\n\t")) {
    case "foo bar":
        v70 := GenS8{F0: 36}
        _ = v70
    default:
        _ = len(v4)
    }
    fmt.Println(v2, v3, v1)
    _ = len(v4)
    switch ("\n\t" + (("hello" + "hello") + "foo bar")) {
    case "foo bar":
        v2 -= 299
    }
    v2 = v2
    var v71 int8 = -109
    v72 := GenS4{F0: 2147483647, F1: 0, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v71
    _ = v72
}
