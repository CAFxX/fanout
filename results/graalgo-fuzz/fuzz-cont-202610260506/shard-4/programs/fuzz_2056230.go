package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 uint; F2 bool }
type GenS6 struct { F0 int64; F1 string }
func genFn0(p0 int32) int {
    var v6 float32 = float32(1.5)
    var v7 int16 = -573
    p0 -= 506
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS1{F0: 680, F1: -291, F2: 416}
        if true {
            if true {
                for fi10 := 0; fi10 < 9; fi10++ {
                    if true {
                        var v11 float32 = ((1.5 - 1.5) / (v6 * 3.14159))
                        var v12 string = "foo bar"
                        _ = v11
                        _ = v12
                    }
                    _ = fi10
                }
                var v13 bool = (!true)
                var v14 int64 = -9223372036854775808
                _ = v13
                _ = v14
            }
        }
        _ = fi8
        _ = v9
    }
    v15 := GenS2{F0: 1000}
    _ = p0
    _ = v15
    _ = v6
    _ = v7
    return 128
}
func genFn1(p0 int16, p1 uint, p2 uint) uint16 {
    v16 := GenS1{F0: 0, F1: p0, F2: p1}
    v17 := GenS2{F0: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 1
}

func main() {
    var v1 int16 = 405
    v2 := make(map[uint32]uint32)
    v3 := []uint32{65535, 2147483647, 65535}
    v4 := make(map[uint32]bool)
    var v5 float32 = -2.25
    if (65535 != 583) {
        v18 := GenS3{F0: v1}
        _ = v18
    } else {
        if (!("foo bar" < "\n\t")) {
            if true {
                v19 := GenS3{F0: v1}
                _ = v19
            } else {
                for fi20 := 0; fi20 < 8; fi20++ {
                    for fi21 := 0; fi21 < 20; fi21++ {
                        _ = len(v3)
                        _ = fi21
                    }
                    _ = fi20
                }
                for fi22 := 0; fi22 < 13; fi22++ {
                    var v23 float32 = 3.14159
                    _ = fi22
                    _ = v23
                }
                v24 := GenS2{F0: v1}
                _ = v24
            }
        } else {
            v1 -= v1
            v1 += v1
            var v25 int8 = -66
            _ = v25
        }
        switch ("" + "foo bar") {
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "a":
            for fi26 := 0; fi26 < 15; fi26++ {
                switch "hello" {
                case "a":
                    v27 := GenS4{F0: 168}
                    _ = v27
                case "foo bar":
                    v5 -= v5
                default:
                    v1 += v1
                }
                var v28 uint8 = 93
                var v29 float32 = (v5 - v5)
                _ = fi26
                _ = v28
                _ = v29
            }
        }
    }
    var v30 uint64 = 986
    var v31 float32 = 1.5
    if false {
        var v32 float32 = float32(float64(v31))
        switch "foo bar" {
        case "foo bar":
            var v33 float32 = v31
            _ = v33
        case "hello":
            switch genFn0(0) {
            case 127:
                if len(v3) > 0 {
                    v3[0] = 1
                }
            case 224:
                var v34 uint = 65535
                _ = v34
            }
        default:
            var v35 uint8 = 2
            _ = v35
        }
        _ = v32
    } else {
        v31 += v5
        delete(v4, 2147483647)
        for fi36 := 0; fi36 < 19; fi36++ {
            fmt.Println(v31, v1)
            v37 := GenS1{F0: 42, F1: (v1 & (v1 | (v1 >> ((0) & 63)))), F2: 18446744073709551615}
            if false {
                for fi38 := 0; fi38 < 9; fi38++ {
                    if (1 < -2147483648) {
                        v30 = v30
                    } else {
                        switch fi36 {
                        case 4294967295:
                            v31 = 0.0
                        case 127:
                            var v39 int8 = 0
                            _ = v39
                        case 642:
                            fmt.Println(fi38)
                        }
                    }
                    _ = fi38
                }
                delete(v2, 2147483647)
            } else {
                if false {
                    for fi40 := 0; fi40 < 15; fi40++ {
                        v3 = append(v3, 65535)
                        v1 = v1
                        if true {
                            if false {
                                v30 += (v30 >> ((2) & 63))
                                v31 = (float32((0.0 / (1.5 / float64((0.1 - -2.25))))) / (v31 / v31))
                                v31 = (v5 + v5)
                                var v41 uint32 = 1
                                _ = v41
                            } else {
                                switch ("hello" + "hello") {
                                case "\n\t":
                                    v42 := GenS1{F0: 832, F1: v1, F2: 258}
                                    _ = v42
                                case "":
                                    v5 += ((((3.14159 * v5) + v5) / v31) - -2.25)
                                case "foo bar":
                                    var v43 int = (fi40 << ((91) & 63))
                                    _ = v43
                                }
                                if false {
                                    v3 = append(v3, 229)
                                    _ = len(v2)
                                }
                            }
                            switch fi36 {
                            case 128:
                                if len(v3) > 0 {
                                    v3[0] = 65535
                                }
                                var v44 int = (fi40 & (fi40 * (fi36 >> ((42) & 63))))
                                _ = v44
                            case 1:
                                var v45 uint64 = 42
                                switch ("a" + "foo bar") {
                                case "foo bar":
                                    v5 -= 0.1
                                case "a":
                                    fmt.Println(v30, fi40, v31)
                                case "":
                                    var v46 bool = true
                                    _ = v46
                                }
                                var v47 uint64 = (v45 % ((v45) | 1))
                                _ = v45
                                _ = v47
                            case 127:
                                v1 = v1
                            }
                            v1 += v1
                        } else {
                            var v48 uint64 = (v30 | v30)
                            switch "" {
                            case "":
                                if true {
                                    switch "a" {
                                    case "":
                                        var v49 bool = true
                                        _ = v49
                                    case "hello":
                                        var v50 bool = (670 <= 0)
                                        _ = v50
                                    }
                                }
                            case "foo bar":
                                if len(v3) > 0 {
                                    v3[0] = 65535
                                }
                            default:
                                v5 = ((0.0 * v5) * (0.0 / v31))
                            }
                            _ = v48
                        }
                        _ = fi40
                    }
                    for fi51 := 0; fi51 < 3; fi51++ {
                        fmt.Println(fi36, v1, v1)
                        if ((0 != -1) && (841 != 490)) {
                            if false {
                                var v52 int = fi51
                                _ = v52
                            } else {
                                if false {
                                    fmt.Println(v1)
                                    switch (-(fi36)) {
                                    case -9223372036854775808:
                                        _ = len(v2)
                                    case 4294967295:
                                        v31 -= 1e10
                                    }
                                    if (0 < v1) {
                                        v53 := GenS5{F0: true, F1: 65535, F2: false}
                                        _ = v53
                                    } else {
                                        v3 = append(v3, 2147483648)
                                    }
                                } else {
                                    v5 = ((float32(0.0) - 0.0) + v31)
                                    fmt.Println(v30, v30, v1)
                                    delete(v2, 2147483648)
                                }
                            }
                            v1 = int16(fi51)
                            var v54 uint16 = 65535
                            _ = v54
                        }
                        delete(v2, 4294967295)
                        _ = fi51
                    }
                } else {
                    v55 := GenS2{F0: (v1 >> ((0) & 63))}
                    var v56 uint = 0
                    _ = v55
                    _ = v56
                }
                if false {
                    fmt.Println(fi36, v5, v31)
                    v30 = (v30 << ((2) & 63))
                    switch fi36 {
                    case 42:
                        v57 := v2[573]
                        _ = v57
                    case 9223372036854775807:
                        for fi58 := 0; fi58 < 16; fi58++ {
                            var v59 int8 = 42
                            _ = fi58
                            _ = v59
                        }
                    }
                } else {
                    v30 -= 630
                    var v60 uint = 65535
                    v3 = append(v3, 2)
                    _ = v60
                }
            }
            _ = fi36
            _ = v37
        }
    }
    v2[42] = 4294967295
    v5 = v31
    var v61 int16 = (v1 % ((-32768) | 1))
    for fi62 := 0; fi62 < 7; fi62++ {
        var v63 int = 32767
        v31 = 1e10
        _ = fi62
        _ = v63
    }
    delete(v2, 42)
    switch "hello" {
    case "a":
        v64 := GenS6{F0: -9223372036854775808, F1: ""}
        _ = v64
    case "":
        for fi65 := 0; fi65 < 11; fi65++ {
            for fi66 := 0; fi66 < 1; fi66++ {
                v61 = (v61 & (v61 << ((255) & 63)))
                var v67 uint8 = 42
                var v68 float32 = (-2.25 - 0.1)
                _ = fi66
                _ = v67
                _ = v68
            }
            _ = fi65
        }
    }
    if (true || (!((v30 > (v30 * 42)) && true))) {
        v5 += (-2.25 * 3.14159)
    } else {
        var v69 int32 = 0
        var v70 uint8 = 2
        _ = v69
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v61
}
