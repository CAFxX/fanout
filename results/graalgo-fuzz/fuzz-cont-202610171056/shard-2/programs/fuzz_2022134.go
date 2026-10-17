package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int32; F2 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 bool = false
    v3 := make(map[uint32]uint32)
    var v4 int16 = 32767
    var v5 int8 = -30
    v6 := make(map[int]uint32)
    fmt.Println(v2, v4, v2)
    var v7 int32 = 2147483647
    switch "" {
    case "":
        var v8 uint8 = 0
        _ = v8
    case "\n\t":
        switch -1 {
        case 127:
            var v9 uint16 = 65535
            _ = v9
        case -691:
            if v2 {
                v10 := GenS1{F0: 464}
                var v11 uint8 = 0
                for fi12 := 0; fi12 < 13; fi12++ {
                    if v2 {
                        switch (fi12 & fi12) {
                        case -9223372036854775808:
                            v2 = v2
                        default:
                            fmt.Println(v2, v11)
                        }
                        for fi13 := 0; fi13 < 14; fi13++ {
                            v4 = (v4 >> ((404) & 63))
                            fmt.Println(v5, v4)
                            v5 += (v5 - (v5 - int8(v4)))
                            _ = fi13
                        }
                    } else {
                        v14 := GenS2{F0: fi12, F1: (v7 / ((v7) | 1)), F2: 2}
                        if ("" != ("" + ("a" + (("foo bar" + "foo bar") + ("\n\t" + ("foo bar" + "\n\t")))))) {
                            var v15 int64 = 299
                            _ = v15
                        } else {
                            for fi16 := 0; fi16 < 7; fi16++ {
                                v17 := GenS2{F0: 1, F1: 60, F2: (v11 + v11)}
                                if true {
                                    v5 -= v5
                                    var v18 float32 = 0.1
                                    v18 += v18
                                    _ = v18
                                }
                                fmt.Println(v7)
                                _ = fi16
                                _ = v17
                            }
                        }
                        v5 += v5
                        _ = v14
                    }
                    v19 := GenS2{F0: int(v5), F1: v7, F2: (v11 | (v11 / (((^(v11))) | 1)))}
                    _ = fi12
                    _ = v19
                }
                _ = v10
                _ = v11
            }
        default:
            v20 := GenS3{F0: 42}
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 uint64 = 2
        fmt.Println(v5, v22, fi21)
        var v23 float64 = (1.5 * (-2.25 * 3.14159))
        _ = fi21
        _ = v22
        _ = v23
    }
    if false {
        var v24 string = "foo bar"
        for fi25 := 0; fi25 < 19; fi25++ {
            if ((0.1 / ((float32((1.5 + -2.25)) + (float32(1.5) - 0.1)) - (float32(((float64(float32(1e10)) / (-2.25 / ((1e10 - (0.0 * float64((1.5 - (float32(1.5) + 3.14159))))) + 0.0))) - (-2.25 + (1.5 - (1.5 + -2.25))))) + float32(1e10)))) != 1.5) {
                v26 := GenS3{F0: v4}
                v27 := GenS4{F0: 18446744073709551615}
                if ((v5 * (v5 << ((2) & 63))) >= v5) {
                    for fi28 := 0; fi28 < 18; fi28++ {
                        v24 = v24
                        _ = fi28
                    }
                    v29 := GenS5{F0: 68}
                    var v30 uint32 = 833
                    _ = v29
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 11; fi31++ {
                        var v32 int64 = 453
                        _ = fi31
                        _ = v32
                    }
                    v33 := GenS5{F0: 0}
                    _ = v33
                }
                _ = v26
                _ = v27
            } else {
                var v34 int8 = (v5 >> ((18446744073709551615) & 63))
                for fi35 := 0; fi35 < 10; fi35++ {
                    if false {
                        if (!v2) {
                            var v36 uint16 = 255
                            v34 = v5
                            _ = v36
                        } else {
                            v37 := []int{fi25, (fi25 ^ fi25), fi35}
                            var v38 uint64 = 621
                            if len(v37) > 0 {
                                v37[0] = 42
                            }
                            _ = len(v37)
                            _ = v37
                            _ = v38
                        }
                        if true {
                            var v39 int16 = v4
                            _ = v39
                        }
                    }
                    var v40 int32 = (-(v7))
                    _ = fi35
                    _ = v40
                }
                if false {
                    v41 := GenS4{F0: 65535}
                    v24 = ("foo bar" + ("" + (v24 + v24)))
                    v42 := GenS6{F0: 255, F1: 42}
                    _ = v41
                    _ = v42
                }
                _ = v34
            }
            fmt.Println(v4, v24, v4)
            v43 := v3[0]
            _ = fi25
            _ = v43
        }
        for fi44 := 0; fi44 < 12; fi44++ {
            for fi45 := 0; fi45 < 18; fi45++ {
                var v46 int64 = -597
                switch (("hello" + v24) + v24) {
                case "hello":
                    v4 += (v4 + (v4 ^ (v4 * 1)))
                case "":
                    v47 := GenS1{F0: 1}
                    _ = v47
                default:
                    _ = len(v6)
                }
                _ = fi45
                _ = v46
            }
            switch ("" + v24) {
            case "":
                v24 = v24
            case "hello":
                for fi48 := 0; fi48 < 18; fi48++ {
                    v49 := GenS2{F0: fi48, F1: (v7 * int32(v5)), F2: 0}
                    if v2 {
                        fmt.Println(v2)
                    }
                    _ = fi48
                    _ = v49
                }
            }
            _ = fi44
        }
        _ = v24
    }
    fmt.Println(v5, v5)
    fmt.Println(v5, v2, v2)
    if (false && (!v2)) {
        fmt.Println(v2, v4)
        switch "foo bar" {
        case "hello":
            if v2 {
                delete(v1, (("\n\t" + "a") + ""))
                if v2 {
                    if ((v7 << ((255) & 63)) != (v7 & v7)) {
                        if true {
                            v2 = (!v2)
                            for fi50 := 0; fi50 < 16; fi50++ {
                                v4 -= v4
                                fmt.Println(v4)
                                _ = fi50
                            }
                            var v51 int8 = (-(v5))
                            _ = v51
                        }
                        switch "" {
                        case "\n\t":
                            for fi52 := 0; fi52 < 3; fi52++ {
                                v53 := GenS4{F0: 555}
                                v4 += (^(v4))
                                _ = fi52
                                _ = v53
                            }
                        }
                        v4 -= 42
                    }
                    switch 32767 {
                    case -1:
                        var v54 int = -1
                        _ = v54
                    case -9223372036854775808:
                        v2 = true
                    default:
                        v55 := GenS5{F0: 220}
                        _ = v55
                    }
                }
            }
        case "foo bar":
            fmt.Println(v7, v4)
        case "\n\t":
            fmt.Println(v5, v2)
        }
        for fi56 := 0; fi56 < 14; fi56++ {
            var v57 uint32 = 637
            _ = fi56
            _ = v57
        }
    }
    v4 -= (-(v4))
    v6[0] = 1
    v7 += (v7 ^ (v7 & 42))
    v58 := []int{-1, 0, 142, -415}
    fmt.Println(v4, v2)
    v7 = v7
    fmt.Println(v2, v2)
    v2 = (!false)
    fmt.Println(v7, v5)
    _ = len(v58)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v58:", len(v58), v58)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v7
}
