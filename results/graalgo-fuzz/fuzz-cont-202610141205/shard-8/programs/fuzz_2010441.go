package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 string; F2 uint8 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS8 struct { F0 uint64; F1 int }
type GenS9 struct { F0 uint16 }

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]bool)
    var v3 uint32 = 0
    var v4 uint = 939
    var v5 string = ("\n\t" + "")
    var v6 uint = v4
    var v7 uint8 = 195
    v8 := GenS1{F0: 979, F1: (v6 & v4), F2: -159}
    switch v5 {
    case "hello":
        fmt.Println(v5, v7)
    case "a":
        switch 2147483647 {
        case 494:
            if (v3 != v3) {
                switch v5 {
                case "":
                    if false {
                        if ((18446744073709551615 == 1) && (9223372036854775807 >= -1)) {
                            var v9 int8 = 48
                            switch (v5 + "hello") {
                            case "a":
                                v10 := []int{32768, -2147483648, -9223372036854775808, 470}
                                _ = v10
                            default:
                                v3 = v3
                            }
                            _ = v9
                        }
                        if false {
                            fmt.Println(v6)
                        }
                    } else {
                        if (32767 <= 32767) {
                            v4 = v6
                            var v11 uint16 = uint16(v3)
                            var v12 int = 9223372036854775807
                            _ = v11
                            _ = v12
                        } else {
                            var v13 uint64 = uint64(v7)
                            v7 = 39
                            _ = v13
                        }
                        for fi14 := 0; fi14 < 3; fi14++ {
                            var v15 float32 = -2.25
                            v16 := GenS2{F0: false}
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        v17 := []int{}
                        _ = v17
                    }
                case "\n\t":
                    v18 := GenS3{F0: 42, F1: "", F2: v7}
                    _ = v18
                case "a":
                    var v19 int32 = 2147483647
                    _ = v19
                default:
                    var v20 bool = true
                    _ = v20
                }
            } else {
                var v21 int16 = v1
                v22 := GenS4{F0: 127, F1: 497}
                _ = v21
                _ = v22
            }
        case 32767:
            switch v5 {
            case "hello":
                var v23 uint16 = 65535
                _ = v23
            case "a":
                v24 := []int{32767, int(v7)}
                _ = v24
            case "":
                v25 := GenS3{F0: 557, F1: v5, F2: v7}
                _ = v25
            default:
                v6 = (v4 + (^(v6)))
            }
        case -691:
            v26 := []int{-9223372036854775808, -9223372036854775808, -121}
            _ = v26
        }
    case "foo bar":
        switch ((v5 + (v5 + (((v5 + ("a" + v5)) + v5) + (v5 + (v5 + v5))))) + v5) {
        case "":
            fmt.Println(v7, v3)
        case "\n\t":
            var v27 float64 = -2.25
            _ = v27
        case "hello":
            v3 -= v3
        }
    }
    var v28 int8 = 42
    for fi29 := 0; fi29 < 7; fi29++ {
        v30 := GenS5{F0: 1, F1: 578}
        _ = len(v2)
        for fi31 := 0; fi31 < 5; fi31++ {
            if (true && true) {
                v32 := GenS1{F0: 9223372036854775807, F1: v4, F2: 42}
                switch (v5 + ("" + ("a" + v5))) {
                case "hello":
                    if ((v28 * int8(v6)) >= -1) {
                        fmt.Println(v3, v5)
                        var v33 int32 = -1
                        var v34 uint8 = 42
                        _ = v33
                        _ = v34
                    } else {
                        var v35 int32 = -2147483648
                        var v36 uint64 = 4294967295
                        _ = v35
                        _ = v36
                    }
                case "a":
                    var v37 string = (v5 + ("" + (((v5 + v5) + (v5 + v5)) + v5)))
                    fmt.Println(v5, v5)
                    var v38 float32 = 1.5
                    _ = v37
                    _ = v38
                default:
                    fmt.Println(v28, v5, fi31)
                }
                fmt.Println(fi29, v5)
                _ = v32
            }
            _ = fi31
        }
        _ = fi29
        _ = v30
    }
    v28 = v28
    fmt.Println(v5)
    var v39 uint16 = 1
    delete(v2, 128)
    for fi40 := 0; fi40 < 16; fi40++ {
        v41 := GenS6{F0: 1, F1: 2}
        switch v5 {
        case "\n\t":
            fmt.Println(v39, v1)
        }
        _ = fi40
        _ = v41
    }
    v42 := []int{}
    if false {
        v39 -= (v39 - v39)
        v4 += 2
        for fi43 := 0; fi43 < 11; fi43++ {
            if false {
                switch fi43 {
                case 290:
                    v3 = uint32(v39)
                case 32768:
                    if (true && false) {
                        for fi44 := 0; fi44 < 14; fi44++ {
                            var v45 int32 = 2147483647
                            _ = fi44
                            _ = v45
                        }
                    } else {
                        for fi46 := 0; fi46 < 19; fi46++ {
                            var v47 int64 = 0
                            _ = len(v2)
                            v48 := GenS7{F0: 576, F1: v39, F2: v3}
                            _ = fi46
                            _ = v47
                            _ = v48
                        }
                        v3 -= v3
                    }
                case 42:
                    if false {
                        var v49 int8 = -128
                        var v50 uint8 = 2
                        _ = v49
                        _ = v50
                    }
                }
                if (!true) {
                    if false {
                        if true {
                            v4 = 255
                            v51 := GenS8{F0: 771, F1: 42}
                            if false {
                                v42 = append(v42, fi43)
                                v3 = uint32(v7)
                            } else {
                                var v52 uint8 = v7
                                var v53 uint32 = uint32(v52)
                                v3 = 2
                                _ = v52
                                _ = v53
                            }
                            _ = v51
                        } else {
                            v54 := GenS4{F0: int8(v7), F1: 9223372036854775807}
                            _ = v54
                        }
                    }
                    for fi55 := 0; fi55 < 3; fi55++ {
                        var v56 int32 = int32(fi55)
                        var v57 uint = 1
                        if false {
                            if false {
                                v1 = v1
                                v57 -= (v57 - 2)
                            }
                        } else {
                            if true {
                                fmt.Println(v3)
                                if (-955 >= 0) {
                                    fmt.Println(v4)
                                    switch "\n\t" {
                                    case "":
                                        var v58 float64 = (1.5 * float64((3.14159 + (1.5 + float32((0.0 - (float64(float32(((3.14159 - float64((0.0 / ((-2.25 * 3.14159) - 0.0)))) / 0.1))) + 0.0)))))))
                                        _ = v58
                                    case "hello":
                                        v39 -= v39
                                    case "a":
                                        if (-1 == 1) {
                                            v42 = append(v42, (fi43 / ((fi55) | 1)))
                                            if len(v42) > 0 {
                                                v42[0] = fi55
                                            }
                                        }
                                    default:
                                        v5 = ("a" + ("hello" + ((v5 + v5) + (((v5 + v5) + ("foo bar" + (v5 + "foo bar"))) + "foo bar"))))
                                        var v59 uint32 = 295
                                        _ = v59
                                    }
                                    if ((v4 ^ (v57 * (v6 - v57))) != 42) {
                                        v60 := GenS9{F0: 0}
                                        v2[(^(fi55))] = true
                                        if len(v42) > 0 {
                                            v42[0] = (fi55 - (fi55 ^ fi55))
                                        }
                                        _ = v60
                                    }
                                } else {
                                    if len(v42) > 0 {
                                        v42[0] = (fi43 << ((v6) & 63))
                                    }
                                    v39 = uint16(v7)
                                }
                                var v61 int16 = v1
                                _ = len(v2)
                                _ = v61
                            } else {
                                if (465 > 2) {
                                    var v62 int32 = int32(fi43)
                                    _ = v62
                                }
                                if false {
                                    switch -842 {
                                    case -352:
                                        fmt.Println(v57, v28, v3)
                                    case 32768:
                                        v56 += int32(v3)
                                    case 32767:
                                        fmt.Println(v4, v3)
                                    }
                                    if true {
                                        _ = len(v2)
                                    }
                                } else {
                                    if false {
                                        switch v5 {
                                        case "a":
                                            if false {
                                                v57 = uint(v7)
                                                delete(v2, (fi55 & fi55))
                                                var v63 uint16 = 65535
                                                _ = v63
                                            }
                                        default:
                                            var v64 uint16 = 2
                                            _ = v64
                                        }
                                        fmt.Println(v28)
                                        v57 = (^(v57))
                                        switch fi43 {
                                        case 0:
                                            var v65 int8 = v28
                                            _ = v65
                                        case 14:
                                            var v66 uint32 = (^(v3))
                                            _ = v66
                                        default:
                                            v39 = 2
                                        }
                                        var v67 uint = (v57 + v6)
                                        _ = v67
                                    }
                                }
                                v56 -= v56
                            }
                            var v68 uint16 = v39
                            v1 = 0
                            _ = v68
                        }
                        _ = fi55
                        _ = v56
                        _ = v57
                    }
                    v39 = 42
                } else {
                    v4 = (^(v6))
                    v69 := GenS6{F0: 0, F1: 0}
                    _ = v69
                }
                var v70 uint = v6
                _ = v70
            }
            v4 = v6
            var v71 uint8 = (v7 / (((^(v7))) | 1))
            _ = fi43
            _ = v71
        }
    } else {
        for fi72 := 0; fi72 < 3; fi72++ {
            if true {
                var v73 int64 = -1
                var v74 float32 = float32(0.0)
                _ = v73
                _ = v74
            }
            v1 -= 32767
            if ((v3 + v3) <= v3) {
                var v75 uint32 = v3
                _ = v75
            } else {
                var v76 int8 = (v28 + -50)
                fmt.Println(v3, v28)
                _ = v76
            }
            _ = fi72
        }
    }
    v3 = (^(v3))
    v39 -= (v39 + 2)
    v3 -= (v3 << ((uint(v3)) & 63))
    v3 -= (v3 - (v3 & 2147483647))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
