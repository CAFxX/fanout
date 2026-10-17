package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint64; F1 uint64; F2 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 uint32 }

func main() {
    v1 := []uint32{795}
    var v2 uint = 306
    var v3 float64 = 3.14159
    var v4 uint = v2
    var v5 int16 = 1
    v6 := make(map[uint32]int)
    fmt.Println(v3, v5)
    var v7 int64 = 42
    v5 = v5
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := GenS1{F0: int16(v7), F1: v4}
        _ = fi8
        _ = v9
    }
    if true {
        if true {
            var v10 uint = v2
            if false {
                var v11 uint16 = 0
                var v12 int64 = (v7 & (^(v7)))
                _ = v11
                _ = v12
            }
            var v13 uint64 = uint64(v4)
            _ = v10
            _ = v13
        } else {
            for fi14 := 0; fi14 < 19; fi14++ {
                if ((42 == -106) || true) {
                    v5 += int16(v7)
                    if true {
                        v6[42] = (-(fi14))
                        v15 := v6[2147483648]
                        _ = v15
                    }
                    var v16 float64 = (3.14159 - 0.1)
                    _ = v16
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v17 float32 = (1.5 + 0.0)
                    if true {
                        var v18 int32 = 0
                        _ = v18
                    } else {
                        for fi19 := 0; fi19 < 9; fi19++ {
                            fmt.Println(v3, v2, v4)
                            var v20 uint32 = 2147483647
                            switch -2147483648 {
                            case 1:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                v21 := GenS2{F0: 42, F1: 835, F2: "hello"}
                                _ = v21
                            }
                            _ = fi19
                            _ = v20
                        }
                        v22 := GenS3{F0: 4294967295}
                        var v23 float32 = v17
                        _ = v22
                        _ = v23
                    }
                    _ = v17
                }
                _ = len(v1)
                var v24 int16 = v5
                _ = fi14
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 19; fi25++ {
            fmt.Println(v2)
            _ = fi25
        }
        for fi26 := 0; fi26 < 15; fi26++ {
            var v27 int32 = -2147483648
            var v28 bool = (v3 >= v3)
            for fi29 := 0; fi29 < 20; fi29++ {
                v5 = (v5 | v5)
                switch "hello" {
                case "\n\t":
                    if (42 < 1) {
                        switch "\n\t" {
                        case "hello":
                            var v30 uint8 = 42
                            _ = v30
                        case "foo bar":
                            if (!v28) {
                                v5 -= 316
                            }
                        case "":
                            if v28 {
                                v28 = (("" + "a") == "hello")
                                var v31 string = ""
                                fmt.Println(v4)
                                switch fi26 {
                                case 128:
                                    v4 = (v4 >> ((508) & 63))
                                default:
                                    _ = len(v6)
                                }
                                _ = v31
                            }
                        default:
                            v32 := GenS1{F0: v5, F1: (v4 - 4294967295)}
                            _ = v32
                        }
                    }
                default:
                    v28 = v28
                }
                var v33 int64 = 900
                _ = fi29
                _ = v33
            }
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    _ = len(v1)
    delete(v6, 2147483648)
    v34 := GenS4{F0: v4}
    v7 -= (-(v7))
    if true {
        v35 := GenS5{F0: v7, F1: 2147483647}
        if (!false) {
            if true {
                v7 = -451
            } else {
                v4 += v2
                v2 = (v2 & 2)
                v2 += 890
            }
        }
        _ = v35
    } else {
        var v36 uint8 = 183
        v1 = append(v1, 4294967295)
        _ = v36
    }
    switch -904 {
    case -654:
        var v37 int32 = -2147483648
        _ = v37
    case 42:
        v2 = v4
    default:
        switch "a" {
        case "foo bar":
            for fi38 := 0; fi38 < 5; fi38++ {
                for fi39 := 0; fi39 < 11; fi39++ {
                    var v40 uint8 = 2
                    _ = fi39
                    _ = v40
                }
                var v41 int64 = 0
                _ = fi38
                _ = v41
            }
        case "":
            fmt.Println(v7, v3)
        }
    }
    switch -2147483648 {
    case -1:
        var v42 int64 = 0
        _ = v42
    case 1:
        var v43 uint8 = 127
        _ = v43
    default:
        if false {
            fmt.Println(v2, v3)
            v2 = (^(v2))
            v7 -= (v7 ^ (v7 / ((v7) | 1)))
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
