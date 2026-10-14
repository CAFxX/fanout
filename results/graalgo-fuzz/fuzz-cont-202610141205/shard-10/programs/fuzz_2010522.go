package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 bool; F1 int16; F2 int8 }
type GenS5 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS6 struct { F0 int64; F1 int32 }
func genFn0() uint {
    var v4 int16 = -1
    if false {
        v5 := []int{128, -1, 127}
        var v6 int = -9223372036854775808
        var v7 uint8 = 180
        _ = v5
        _ = v6
        _ = v7
    } else {
        v4 -= -1
    }
    if false {
        v8 := []int{}
        v9 := GenS1{F0: -625, F1: 1, F2: 0}
        v4 = 896
        _ = v8
        _ = v9
    }
    v4 = 1
    v4 = (v4 * (^(v4)))
    _ = v4
    return 137
}

func main() {
    v1 := []uint32{42, 628, 198, 805}
    var v2 uint64 = 4294967295
    var v3 uint8 = 0
    if (329 >= 65535) {
        var v10 uint32 = 65535
        if false {
            var v11 int16 = 1
            v12 := GenS1{F0: -2147483648, F1: 42, F2: 2}
            if true {
                for fi13 := 0; fi13 < 16; fi13++ {
                    switch "hello" {
                    case "foo bar":
                        v10 = (^(v10))
                    case "hello":
                        for fi14 := 0; fi14 < 12; fi14++ {
                            if true {
                                v11 -= (v11 & v11)
                                var v15 int32 = -2147483648
                                _ = v15
                            } else {
                                fmt.Println(v10, v11, fi14)
                                v10 -= (v10 - v10)
                            }
                            var v16 int32 = 1
                            _ = fi14
                            _ = v16
                        }
                    case "a":
                        var v17 int16 = v11
                        _ = v17
                    }
                    switch "\n\t" {
                    case "hello":
                        v1 = append(v1, 42)
                    }
                    _ = fi13
                }
                var v18 uint16 = 2
                _ = v18
            } else {
                switch "" {
                case "a":
                    v19 := GenS2{F0: (^(v10))}
                    _ = v19
                default:
                    v1 = append(v1, (^(v10)))
                }
                if (false && (!false)) {
                    v2 -= v2
                    switch (("foo bar" + "a") + ("" + ("foo bar" + "a"))) {
                    case "":
                        var v20 int8 = 1
                        _ = v20
                    default:
                        if false {
                            v3 = (^(v3))
                            var v21 int64 = 9223372036854775807
                            for fi22 := 0; fi22 < 5; fi22++ {
                                for fi23 := 0; fi23 < 19; fi23++ {
                                    v1 = append(v1, (v10 << ((255) & 63)))
                                    v11 = (v11 >> ((508) & 63))
                                    _ = fi23
                                }
                                v24 := GenS2{F0: v10}
                                _ = fi22
                                _ = v24
                            }
                            _ = v21
                        }
                    }
                } else {
                    _ = len(v1)
                }
                for fi25 := 0; fi25 < 2; fi25++ {
                    v26 := GenS2{F0: (v10 % (((v10 % ((42) | 1))) | 1))}
                    for fi27 := 0; fi27 < 14; fi27++ {
                        var v28 int = (fi25 & (fi25 << ((255) & 63)))
                        _ = fi27
                        _ = v28
                    }
                    v11 += (v11 >> ((2) & 63))
                    _ = fi25
                    _ = v26
                }
            }
            _ = v11
            _ = v12
        } else {
            fmt.Println(v2, v3)
            v29 := GenS2{F0: 133}
            var v30 uint64 = 42
            _ = v29
            _ = v30
        }
        _ = v10
    }
    v2 = v2
    v3 -= uint8(v2)
    v2 = (v2 >> ((4294967295) & 63))
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 uint8 = v3
        for fi33 := 0; fi33 < 19; fi33++ {
            switch (fi33 | (fi31 % ((fi33) | 1))) {
            case 1:
                if false {
                    fmt.Println(fi33)
                }
            case -2147483648:
                v34 := GenS1{F0: 1, F1: 127, F2: 1}
                _ = v34
            default:
                v32 = (v32 + 2)
            }
            var v35 bool = (!(!(2147483647 <= 4294967295)))
            v3 = 128
            _ = fi33
            _ = v35
        }
        v3 -= v32
        _ = fi31
        _ = v32
    }
    v3 -= (v3 ^ 127)
    v36 := GenS2{F0: 1}
    fmt.Println(v2, v3)
    v37 := GenS3{F0: 4294967295, F1: 1000}
    v3 = v3
    if true {
        for fi38 := 0; fi38 < 2; fi38++ {
            switch "\n\t" {
            case "":
                fmt.Println(v3)
            case "foo bar":
                if (!(0.1 != ((1.5 + 1.5) / 3.14159))) {
                    v3 -= v3
                    var v39 float32 = -2.25
                    for fi40 := 0; fi40 < 9; fi40++ {
                        switch (fi40 | -2147483648) {
                        case 42:
                            var v41 uint8 = v3
                            _ = v41
                        case 0:
                            var v42 int16 = -1
                            _ = v42
                        case 127:
                            switch (fi40 + 468) {
                            case 128:
                                v3 -= uint8(v2)
                            case 127:
                                fmt.Println(fi40)
                                var v43 int32 = 163
                                _ = v43
                            case 0:
                                v44 := GenS4{F0: (!(0.0 >= 0.0)), F1: 1000, F2: 100}
                                _ = v44
                            default:
                                var v45 int8 = 70
                                _ = v45
                            }
                        default:
                            var v46 float64 = 1e10
                            _ = v46
                        }
                        v39 -= v39
                        v47 := GenS1{F0: -627, F1: 0, F2: 453}
                        _ = fi40
                        _ = v47
                    }
                    _ = v39
                } else {
                    var v48 uint8 = v3
                    if (1 != 0) {
                        if true {
                            v49 := GenS5{F0: (^(v2)), F1: 4294967295, F2: 255}
                            var v50 int = fi38
                            _ = v49
                            _ = v50
                        } else {
                            v48 = v48
                        }
                    }
                    for fi51 := 0; fi51 < 10; fi51++ {
                        v3 = v3
                        var v52 float64 = 0.0
                        if false {
                            v48 = v3
                            v53 := GenS6{F0: 42, F1: 596}
                            _ = v53
                        } else {
                            fmt.Println(v48, fi51, v3)
                            if false {
                                fmt.Println(fi38, fi51, v52)
                                var v54 int16 = 1000
                                _ = v54
                            }
                            switch "foo bar" {
                            case "foo bar":
                                if false {
                                    v1 = append(v1, 568)
                                    if true {
                                        v48 = (v3 * v3)
                                        if (false || (42 >= 127)) {
                                            var v55 bool = true
                                            _ = v55
                                        }
                                    }
                                } else {
                                    if (3.14159 != (0.1 + ((3.14159 - float32(0.0)) / 1e10))) {
                                        if (4294967295 >= (v2 & 0)) {
                                            v52 = v52
                                            var v56 string = ("foo bar" + "\n\t")
                                            _ = v56
                                        } else {
                                            var v57 uint8 = (v48 * (v48 * v3))
                                            _ = v57
                                        }
                                        fmt.Println(v48, v48, v48)
                                        fmt.Println(v52, fi38, fi38)
                                        switch "\n\t" {
                                        case "":
                                            v58 := GenS1{F0: 1, F1: 0, F2: 0}
                                            _ = v58
                                        case "a":
                                            v48 = (v3 << ((0) & 63))
                                        case "\n\t":
                                            fmt.Println(v2, v48, v52)
                                        }
                                    }
                                    var v59 uint = 4294967295
                                    v3 = 93
                                    _ = v59
                                }
                            case "\n\t":
                                fmt.Println(v48)
                            }
                            var v60 int = 4294967295
                            v3 = (^(v48))
                            _ = v60
                        }
                        _ = fi51
                        _ = v52
                    }
                    _ = v48
                }
            }
            _ = fi38
        }
        if false {
            for fi61 := 0; fi61 < 20; fi61++ {
                if (uint(v2) == 2) {
                    var v62 bool = true
                    v62 = (v3 == 128)
                    _ = v62
                } else {
                    if (!((42 != 0) && (!(!true)))) {
                        for fi63 := 0; fi63 < 8; fi63++ {
                            var v64 uint = genFn0()
                            v1 = append(v1, 171)
                            v2 -= (v2 + (v2 >> ((1) & 63)))
                            var v65 int64 = 1
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                    }
                }
                _ = fi61
            }
            v3 -= (v3 % (((v3 << ((genFn0()) & 63))) | 1))
        }
    } else {
        var v66 float32 = 0.1
        if (!(424 < 2147483648)) {
            switch -981 {
            case 128:
                v2 = v2
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = len(v1)
        } else {
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
        }
        _ = v66
    }
    for fi67 := 0; fi67 < 18; fi67++ {
        switch "" {
        case "a":
            for fi68 := 0; fi68 < 20; fi68++ {
                var v69 int16 = -654
                _ = fi68
                _ = v69
            }
        default:
            var v70 int8 = int8(v2)
            _ = v70
        }
        _ = len(v1)
        _ = fi67
    }
    if false {
        fmt.Println(v2)
        var v71 bool = true
        _ = v71
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
}
