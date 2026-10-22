package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS6 struct { F0 uint32; F1 uint; F2 int64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool; F1 uint16; F2 int8 }
func genFn0() int {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 bool = false
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 289
    _ = v9
    return -972
}
func genFn1() uint64 {
    var v10 int8 = 0
    var v11 uint8 = 30
    v12 := []int{128, 128}
    for fi13 := 0; fi13 < 2; fi13++ {
        switch -1 {
        case -897:
            if len(v12) > 0 {
                v12[0] = (fi13 % (((fi13 / (((-(fi13))) | 1))) | 1))
            }
        case 335:
            switch "foo bar" {
            case "a":
                fmt.Println(v11, v10)
            case "hello":
                v10 = v10
            default:
                if len(v12) > 0 {
                    v12[0] = genFn0()
                }
            }
        case 42:
            for fi14 := 0; fi14 < 4; fi14++ {
                v10 += (v10 / (((v10 >> ((1) & 63))) | 1))
                _ = fi14
            }
        default:
            for fi15 := 0; fi15 < 8; fi15++ {
                v16 := GenS1{F0: 65535, F1: -32768}
                _ = fi15
                _ = v16
            }
        }
        var v17 uint32 = 4294967295
        _ = fi13
        _ = v17
    }
    _ = v10
    _ = v11
    _ = v12
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 1
    var v3 int32 = -1
    var v4 uint8 = v2
    var v5 uint8 = v4
    v6 := make(map[uint32]uint32)
    if true {
        switch 783 {
        case 127:
            var v18 uint8 = v5
            _ = v18
        case 2147483647:
            var v19 float64 = (float64(float32(1.5)) - (((0.0 * -2.25) / 0.1) - -2.25))
            _ = v19
        case -9223372036854775808:
            fmt.Println(v2, v2)
        default:
            v20 := GenS1{F0: 1, F1: -1}
            _ = v20
        }
    } else {
        var v21 uint32 = 2147483647
        v22 := GenS1{F0: v21, F1: -32768}
        switch "\n\t" {
        case "":
            switch 128 {
            case 2147483647:
                v23 := GenS1{F0: 42, F1: -395}
                _ = v23
            }
        case "a":
            var v24 uint = 324
            _ = v24
        }
        _ = v21
        _ = v22
    }
    v2 = 255
    var v25 uint16 = 42
    v4 += v2
    var v26 float64 = 1e10
    v27 := GenS1{F0: 548, F1: -32768}
    if (false || (-862 >= 0)) {
        v3 -= -319
        var v28 uint = uint(v25)
        var v29 int = -519
        _ = v28
        _ = v29
    } else {
        v30 := GenS2{F0: 0}
        _ = v30
    }
    v5 = 127
    if true {
        var v31 string = "a"
        switch 32767 {
        case -9223372036854775808:
            fmt.Println(v2, v26, v25)
        default:
            switch v31 {
            case "":
                var v32 float64 = v26
                _ = v32
            case "\n\t":
                v33 := GenS3{F0: 0, F1: 878}
                _ = v33
            case "a":
                if ((0.1 * (-2.25 / 1e10)) >= 0.0) {
                    var v34 int32 = 2147483647
                    _ = v34
                } else {
                    v35 := []int{-2147483648, -1}
                    _ = v35
                }
            default:
                if (!true) {
                    v36 := GenS3{F0: 436, F1: 2147483647}
                    if true {
                        v37 := []int{-184, 0}
                        for fi38 := 0; fi38 < 9; fi38++ {
                            v5 = (v5 >> ((0) & 63))
                            if true {
                                v5 -= v2
                                if false {
                                    switch v31 {
                                    case "":
                                        for fi39 := 0; fi39 < 12; fi39++ {
                                            if (true && (!false)) {
                                                var v40 string = v31
                                                v3 = (-(v3))
                                                var v41 int16 = -1
                                                _ = v40
                                                _ = v41
                                            }
                                            v3 += (^(v3))
                                            var v42 int = (fi38 << ((393) & 63))
                                            _ = fi39
                                            _ = v42
                                        }
                                    default:
                                        for fi43 := 0; fi43 < 1; fi43++ {
                                            switch 128 {
                                            case 4294967295:
                                                var v44 float64 = float64(((3.14159 * 1.5) + (1e10 / 0.1)))
                                                _ = v44
                                            }
                                            v45 := GenS4{F0: -53, F1: -1}
                                            v26 = (0.0 * (v26 + 1.5))
                                            fmt.Println(v26, v4, v2)
                                            _ = fi43
                                            _ = v45
                                        }
                                    }
                                }
                            }
                            v46 := GenS5{F0: -54, F1: 18446744073709551615, F2: v25}
                            _ = fi38
                            _ = v46
                        }
                        _ = v37
                    }
                    _ = v36
                } else {
                    switch 127 {
                    case 4294967295:
                        var v47 float32 = ((0.0 * (0.1 - 3.14159)) + 3.14159)
                        _ = v47
                    case 819:
                        v48 := GenS6{F0: 4294967295, F1: 1, F2: 923}
                        _ = v48
                    default:
                        var v49 uint32 = 2147483648
                        _ = v49
                    }
                    v5 -= 127
                }
            }
        }
        for fi50 := 0; fi50 < 19; fi50++ {
            v51 := GenS7{F0: 1000}
            _ = fi50
            _ = v51
        }
        _ = v31
    }
    var v52 uint = uint(v25)
    for fi53 := 0; fi53 < 11; fi53++ {
        var v54 int64 = 488
        switch "hello" {
        case "\n\t":
            v26 = v26
        }
        v55 := GenS8{F0: false, F1: v25, F2: -124}
        _ = fi53
        _ = v54
        _ = v55
    }
    for fi56 := 0; fi56 < 18; fi56++ {
        if (true && false) {
            fmt.Println(v2)
            var v57 uint8 = v4
            switch ("foo bar" + "") {
            case "hello":
                var v58 uint8 = 60
                _ = v58
            case "a":
                var v59 int8 = 1
                _ = v59
            }
            _ = v57
        }
        _ = fi56
    }
    var v60 string = ("\n\t" + ("\n\t" + "hello"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v60
}
