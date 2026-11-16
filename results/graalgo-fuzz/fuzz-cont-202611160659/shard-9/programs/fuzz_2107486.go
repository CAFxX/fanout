package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }
type GenS2 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool; F1 bool; F2 int16 }
type GenS6 struct { F0 uint16 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    var v4 bool = (42 != -128)
    var v5 string = "foo bar"
    var v6 uint = 0
    fmt.Println(v6, v4)
    var v7 float64 = float64(((1.5 * (0.0 - ((float32(1e10) * 3.14159) / 1e10))) + float32((0.1 / 3.14159))))
    v7 = v7
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 uint32 = 199
        var v10 int8 = -128
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 0, F1: 2147483647, F2: ("\n\t" + "\n\t")}
    if (255 == 168) {
        fmt.Println(v6, v4, v6)
    } else {
        fmt.Println(v4, v5, v7)
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 int8 = -21
        for fi14 := 0; fi14 < 8; fi14++ {
            v15 := GenS1{F0: 609, F1: 1, F2: "\n\t"}
            if (127 < 42) {
                var v16 int = (-(fi12))
                if true {
                    v6 -= (v6 * (v6 - v6))
                    switch fi12 {
                    case -9223372036854775808:
                        v17 := GenS1{F0: 651, F1: 42, F2: (v5 + v5)}
                        _ = v17
                    case 9223372036854775807:
                        if true {
                            v5 = v5
                            v7 += v7
                            v18 := GenS2{F0: 32767, F1: 1, F2: 1}
                            _ = v18
                        } else {
                            var v19 int32 = 1
                            if v4 {
                                delete(v3, (("\n\t" + v5) + ""))
                            }
                            switch -376 {
                            case -699:
                                v20 := GenS3{F0: 1, F1: 42}
                                var v21 string = ((v5 + "a") + v5)
                                _ = v20
                                _ = v21
                            case 9223372036854775807:
                                delete(v3, v5)
                                v22 := GenS1{F0: 4294967295, F1: v19, F2: (v5 + (v5 + (("a" + (("\n\t" + v5) + v5)) + (("hello" + (v5 + v5)) + "hello"))))}
                                var v23 uint8 = 33
                                v24 := GenS2{F0: 1, F1: uint64(v23), F2: 65535}
                                v25 := []int{}
                                _ = v22
                                _ = v23
                                _ = v24
                                _ = v25
                            default:
                                switch "hello" {
                                case "foo bar":
                                    v26 := []int{}
                                    _ = v26
                                default:
                                    v19 -= (v19 >> ((v6) & 63))
                                }
                            }
                            v5 = "hello"
                            fmt.Println(v6, v6)
                            _ = v19
                        }
                    case 440:
                        v5 = (v5 + (v5 + v5))
                    }
                    fmt.Println(v6)
                } else {
                    var v27 uint64 = 891
                    if v4 {
                        v4 = (12 < 251)
                        fmt.Println(v5, v16)
                        var v28 int = fi14
                        _ = v28
                    }
                    _ = v27
                }
                _ = v16
            }
            if v4 {
                v29 := GenS3{F0: 222, F1: 128}
                v30 := GenS4{F0: 42}
                _ = v29
                _ = v30
            } else {
                var v31 int8 = (v13 << ((42) & 63))
                switch -954 {
                case 2147483647:
                    v32 := GenS5{F0: (0 != 522), F1: (!false), F2: 32767}
                    _ = v32
                case 0:
                    if true {
                        v33 := GenS5{F0: v4, F1: (!true), F2: 32767}
                        if (!(v4 && false)) {
                            v7 -= v7
                            v34 := GenS6{F0: 65535}
                            _ = v34
                        } else {
                            fmt.Println(v6)
                        }
                        _ = v33
                    }
                }
                _ = v31
            }
            _ = fi14
            _ = v15
        }
        for fi35 := 0; fi35 < 13; fi35++ {
            var v36 int8 = (v13 >> ((0) & 63))
            _ = fi35
            _ = v36
        }
        _ = fi12
        _ = v13
    }
    v37 := []int{-1, 743, 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
