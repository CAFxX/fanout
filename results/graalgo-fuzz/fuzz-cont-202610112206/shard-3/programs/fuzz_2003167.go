package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 uint }

func main() {
    v1 := []int{2147483647}
    var v2 int32 = -2147483648
    v3 := make(map[int]uint32)
    if false {
        for fi4 := 0; fi4 < 18; fi4++ {
            if (955 != 330) {
                v5 := GenS1{F0: 9223372036854775807}
                _ = v5
            }
            _ = fi4
        }
        if false {
            var v6 int = -1
            v7 := GenS1{F0: 9223372036854775807}
            _ = v6
            _ = v7
        } else {
            fmt.Println(v2)
            v8 := GenS1{F0: 9223372036854775807}
            _ = v8
        }
        _ = len(v1)
    } else {
        if false {
            v2 = 0
            var v9 string = ("a" + "\n\t")
            _ = v9
        }
    }
    switch "" {
    case "foo bar":
        var v10 int64 = -875
        _ = v10
    case "hello":
        var v11 uint64 = 573
        _ = v11
    case "a":
        fmt.Println(v2)
    default:
        v3[2147483647] = 4294967295
    }
    v12 := GenS1{F0: -1}
    switch ("hello" + (("hello" + "hello") + "a")) {
    case "foo bar":
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 int = (fi13 << ((4294967295) & 63))
            v14 = (^(v14))
            v2 -= 800
            _ = fi13
            _ = v14
        }
    case "":
        var v15 float64 = 0.0
        _ = v15
    }
    if (1.5 < 0.0) {
        v2 = v2
        if false {
            v16 := GenS2{F0: 0, F1: 1, F2: 42}
            switch ("\n\t" + "hello") {
            case "\n\t":
                v2 = -404
            case "":
                switch "\n\t" {
                case "\n\t":
                    v17 := GenS3{F0: "a", F1: 65535}
                    _ = v17
                case "":
                    switch -1 {
                    case 42:
                        v2 -= -924
                    case 127:
                        switch 128 {
                        case 0:
                            if false {
                                if (!(42 > 1)) {
                                    fmt.Println(v2)
                                    for fi18 := 0; fi18 < 18; fi18++ {
                                        fmt.Println(fi18, v2)
                                        _ = fi18
                                    }
                                }
                                if false {
                                    v19 := GenS2{F0: 65535, F1: 696, F2: 2}
                                    fmt.Println(v2)
                                    _ = v19
                                } else {
                                    fmt.Println(v2)
                                    var v20 float32 = (1e10 * -2.25)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v20
                                }
                                for fi21 := 0; fi21 < 7; fi21++ {
                                    if false {
                                        v22 := v3[(fi21 * fi21)]
                                        v23 := v3[fi21]
                                        var v24 uint64 = 4294967295
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                    }
                                    var v25 int = (fi21 | (fi21 - (fi21 ^ fi21)))
                                    v2 += (v2 | v2)
                                    _ = fi21
                                    _ = v25
                                }
                            }
                        case 511:
                            v2 += v2
                        case 4294967295:
                            var v26 int32 = 42
                            _ = v26
                        default:
                            for fi27 := 0; fi27 < 16; fi27++ {
                                if (((-2.25 * -2.25) <= 1e10) && false) {
                                    var v28 float32 = -2.25
                                    _ = v28
                                } else {
                                    if false {
                                        var v29 uint16 = 42
                                        var v30 string = ("" + ("\n\t" + "\n\t"))
                                        _ = v29
                                        _ = v30
                                    } else {
                                        var v31 uint32 = 704
                                        var v32 uint8 = 64
                                        _ = v31
                                        _ = v32
                                    }
                                    fmt.Println(fi27, v2)
                                    v33 := GenS4{F0: (32767 != 1), F1: 699}
                                    _ = v33
                                }
                                fmt.Println(fi27)
                                for fi34 := 0; fi34 < 3; fi34++ {
                                    var v35 int16 = -681
                                    v36 := GenS1{F0: 1}
                                    _ = fi34
                                    _ = v35
                                    _ = v36
                                }
                                _ = fi27
                            }
                        }
                    case 2147483647:
                        var v37 string = (("\n\t" + (("a" + "") + "")) + "a")
                        _ = v37
                    }
                }
            }
            _ = v16
        } else {
            for fi38 := 0; fi38 < 18; fi38++ {
                v2 = (v2 - v2)
                if (false && (!false)) {
                    for fi39 := 0; fi39 < 9; fi39++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v2)
                        var v40 float64 = (float64(3.14159) - 0.1)
                        var v41 float32 = (1.5 - (1e10 * 1.5))
                        var v42 int8 = 7
                        _ = fi39
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                    var v43 float64 = 1.5
                    _ = v43
                } else {
                    var v44 uint16 = 1
                    v2 -= (v2 - (-(v2)))
                    _ = v44
                }
                var v45 int = fi38
                _ = fi38
                _ = v45
            }
            delete(v3, 4294967295)
        }
    } else {
        v2 = (-(v2))
        v2 = 716
    }
    switch 166 {
    case 9223372036854775807:
        v2 += (v2 / ((1) | 1))
    case 193:
        v2 -= v2
    default:
        fmt.Println(v2)
    }
    if (!(!(0.0 != 3.14159))) {
        for fi46 := 0; fi46 < 12; fi46++ {
            if ("" >= "\n\t") {
                v2 = v2
            }
            v2 = 165
            _ = fi46
        }
        v1 = append(v1, -9223372036854775808)
    }
    v47 := GenS5{F0: 4294967295}
    var v48 uint = 255
    v48 = (v48 & v48)
    v48 += 255
    v49 := GenS6{F0: 106, F1: v48}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v48:", v48)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v47
    _ = v48
    _ = v49
}
