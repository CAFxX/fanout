package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int16; F2 int32 }
type GenS4 struct { F0 int8; F1 int32; F2 uint16 }
func genFn0(p0 uint, p1 int) int32 {
    var v6 float32 = (1.5 - 0.0)
    v7 := []int{p1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    var v1 uint = 255
    v2 := []int{-2147483648, -441}
    var v3 int16 = 1000
    var v4 int64 = -1
    v5 := make(map[uint32]int)
    switch "hello" {
    case "foo bar":
        v8 := v5[42]
        _ = v8
    case "hello":
        v1 = (v1 & 316)
    case "\n\t":
        var v9 string = "hello"
        _ = v9
    default:
        v4 += v4
    }
    v3 += 42
    v4 = -9223372036854775808
    _ = len(v5)
    v3 -= 0
    var v10 int = 6
    var v11 int64 = (-(v4))
    switch 902 {
    case 127:
        var v12 int = 0
        _ = v12
    case -1:
        v1 = 216
    case 32767:
        var v13 uint = (v1 * v1)
        _ = v13
    default:
        fmt.Println(v3, v10, v4)
    }
    v14 := GenS1{F0: v11}
    v15 := GenS2{F0: (!false)}
    switch ((((("a" + (("\n\t" + "") + "")) + ("" + "foo bar")) + (("foo bar" + ("hello" + (("hello" + "") + ""))) + "hello")) + "\n\t") + ("\n\t" + ("\n\t" + "a"))) {
    case "foo bar":
        switch "a" {
        case "foo bar":
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 float32 = 0.1
                for fi18 := 0; fi18 < 20; fi18++ {
                    if ((v3 / ((v3) | 1)) >= v3) {
                        var v19 int32 = int32(v1)
                        if false {
                            v20 := GenS3{F0: 2, F1: v3, F2: (^(v19))}
                            switch "" {
                            case "foo bar":
                                switch -642 {
                                case 638:
                                    v3 = v3
                                case -158:
                                    _ = len(v5)
                                }
                            default:
                                if false {
                                    v19 -= v19
                                }
                            }
                            fmt.Println(fi16, v1, fi16)
                            _ = v20
                        }
                        var v21 uint8 = 2
                        _ = v19
                        _ = v21
                    }
                    _ = fi18
                }
                _ = fi16
                _ = v17
            }
        case "a":
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 uint = v1
                _ = fi22
                _ = v23
            }
        default:
            if false {
                if len(v2) > 0 {
                    v2[0] = v10
                }
                var v24 int16 = (^(v3))
                for fi25 := 0; fi25 < 15; fi25++ {
                    if false {
                        v2 = append(v2, (v10 >> ((255) & 63)))
                        v2 = append(v2, 922)
                        for fi26 := 0; fi26 < 7; fi26++ {
                            var v27 uint32 = 4294967295
                            fmt.Println(v10)
                            switch fi26 {
                            case 127:
                                _ = len(v2)
                            case -1:
                                if ((v4 % (((v4 << ((2) & 63))) | 1)) < v4) {
                                    fmt.Println(v3, fi26, v1)
                                    var v28 int16 = (v24 >> ((445) & 63))
                                    _ = v28
                                }
                            case -351:
                                v29 := GenS4{F0: 4, F1: 2147483647, F2: 323}
                                _ = v29
                            }
                            _ = fi26
                            _ = v27
                        }
                    } else {
                        v30 := GenS1{F0: v11}
                        if true {
                            delete(v5, 2)
                            for fi31 := 0; fi31 < 6; fi31++ {
                                v4 += 1
                                _ = fi31
                            }
                        } else {
                            for fi32 := 0; fi32 < 4; fi32++ {
                                fmt.Println(fi25)
                                v4 -= (v11 - 0)
                                v10 = 32768
                                _ = fi32
                            }
                            var v33 float32 = 1e10
                            var v34 uint16 = 2
                            _ = v33
                            _ = v34
                        }
                        v35 := GenS1{F0: (v4 << ((v1) & 63))}
                        _ = v30
                        _ = v35
                    }
                    _ = fi25
                }
                _ = v24
            } else {
                fmt.Println(v1, v10, v4)
                var v36 float32 = 1.5
                _ = v36
            }
        }
    case "a":
        if len(v2) > 0 {
            v2[0] = 0
        }
    case "\n\t":
        v37 := GenS1{F0: (v11 << ((255) & 63))}
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
