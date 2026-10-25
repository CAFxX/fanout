package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 string; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int64; F2 uint }
type GenS6 struct { F0 int32; F1 uint32 }
type GenS7 struct { F0 int64; F1 int32; F2 uint }
func genFn0() uint16 {
    v4 := GenS1{F0: false}
    var v5 uint = 42
    _ = v4
    _ = v5
    return 938
}

func main() {
    var v1 uint16 = 702
    var v2 string = ""
    var v3 string = (v2 + v2)
    var v6 uint64 = 4294967295
    switch "a" {
    case "hello":
        v1 += v1
    case "a":
        v7 := []int{-232, -9223372036854775808, 127}
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 uint16 = v1
            v9 = v9
            var v10 string = v2
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        v11 := GenS1{F0: false}
        _ = v11
    }
    v12 := []int{4294967295, 9223372036854775807, 32767}
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 uint = 42
        _ = fi13
        _ = v14
    }
    var v15 bool = (true || (1.5 >= 0.0))
    if ((1.5 - 1.5) < 0.0) {
        var v16 string = "foo bar"
        var v17 bool = v15
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 18; fi18++ {
            v12 = append(v12, (fi18 % ((fi18) | 1)))
            if v15 {
                v1 = genFn0()
                v6 = (v6 | v6)
            } else {
                if (493 > 2147483648) {
                    v2 = ((((v3 + "hello") + "hello") + v3) + v2)
                    v12 = append(v12, 32767)
                    var v19 float32 = (-2.25 * float32(((-2.25 * (0.0 / 1.5)) / (1e10 * 3.14159))))
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 9; fi20++ {
                        fmt.Println(v3, fi20)
                        v21 := GenS2{F0: fi18, F1: "hello", F2: 2}
                        v6 = v6
                        _ = fi20
                        _ = v21
                    }
                }
                if false {
                    v3 = v3
                    v12 = append(v12, (fi18 + (fi18 & 70)))
                    var v22 int = int(v1)
                    _ = v22
                } else {
                    v23 := GenS2{F0: fi18, F1: v2, F2: 2}
                    _ = v23
                }
                fmt.Println(v3, fi18, v1)
            }
            if ((-2.25 <= float64(1e10)) && v15) {
                if v15 {
                    if (!(1 > 42)) {
                        var v24 float64 = (-2.25 - 0.0)
                        for fi25 := 0; fi25 < 5; fi25++ {
                            v12 = append(v12, fi25)
                            v26 := GenS3{F0: (!(v2 < v3)), F1: (v1 & v1)}
                            v27 := GenS4{F0: ("\n\t" == v3)}
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        _ = v24
                    }
                    v28 := GenS3{F0: (42 == 1), F1: (v1 ^ v1)}
                    _ = v28
                } else {
                    v29 := GenS4{F0: (!v15)}
                    _ = v29
                }
                v30 := GenS1{F0: (!true)}
                var v31 uint = 968
                _ = v30
                _ = v31
            }
            _ = fi18
        }
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        for fi33 := 0; fi33 < 17; fi33++ {
            v34 := GenS2{F0: (fi32 / (((fi33 & fi33)) | 1)), F1: v2, F2: 68}
            _ = fi33
            _ = v34
        }
        var v35 bool = v15
        v6 = (v6 << ((0) & 63))
        _ = fi32
        _ = v35
    }
    switch 32767 {
    case -8:
        v36 := GenS5{F0: -375, F1: 0, F2: 674}
        _ = v36
    case 4294967295:
        if (23 <= 1) {
            switch -1 {
            case -1:
                v15 = (1 <= 1)
            case -2147483648:
                switch 4294967295 {
                case 9223372036854775807:
                    for fi37 := 0; fi37 < 13; fi37++ {
                        switch v2 {
                        case "a":
                            fmt.Println(v15)
                        default:
                            switch v2 {
                            case "hello":
                                for fi38 := 0; fi38 < 4; fi38++ {
                                    if (1.5 == 3.14159) {
                                        if len(v12) > 0 {
                                            v12[0] = (^(fi37))
                                        }
                                        v2 = "foo bar"
                                        v39 := GenS1{F0: v15}
                                        _ = v39
                                    }
                                    var v40 uint8 = 220
                                    v6 = (v6 | v6)
                                    _ = fi38
                                    _ = v40
                                }
                            }
                        }
                        _ = fi37
                    }
                case 574:
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                default:
                    var v41 uint16 = (v1 * v1)
                    _ = v41
                }
            case -784:
                switch 128 {
                case -9223372036854775808:
                    if len(v12) > 0 {
                        v12[0] = 128
                    }
                case 432:
                    if (!(v15 || ((1.5 != (float64((-2.25 * float32(1e10))) / float64((((3.14159 / 3.14159) / float32(-2.25)) / float32(float64(1e10)))))) || (v15 || ((-727 <= -1) && false))))) {
                        var v42 float64 = ((1e10 - float64(0.0)) * -2.25)
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        var v43 int = 32768
                        _ = v42
                        _ = v43
                    } else {
                        var v44 uint8 = 255
                        v2 = v2
                        var v45 uint8 = v44
                        _ = v44
                        _ = v45
                    }
                case 32767:
                    v15 = (!(v15 && v15))
                }
            }
            v2 = "a"
        } else {
            v6 = v6
            fmt.Println(v6, v15, v1)
        }
    default:
        v3 = "a"
    }
    v46 := GenS6{F0: 560, F1: 1}
    var v47 string = (("a" + ("foo bar" + v2)) + ((v2 + (("foo bar" + ("foo bar" + v3)) + "foo bar")) + v3))
    if len(v12) > 0 {
        v12[0] = -40
    }
    v12 = append(v12, -9223372036854775808)
    var v48 uint32 = 2
    v49 := GenS7{F0: 42, F1: 437, F2: 451}
    fmt.Println(v3)
    if len(v12) > 0 {
        v12[0] = 665
    }
    var v50 float32 = ((3.14159 / 1.5) / (float32((1.5 - 0.0)) / 1.5))
    for fi51 := 0; fi51 < 4; fi51++ {
        var v52 string = v3
        _ = fi51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v6
}
