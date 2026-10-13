package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint }
type GenS2 struct { F0 bool; F1 int8; F2 bool }
type GenS3 struct { F0 string; F1 int32; F2 string }
func genFn0(p0 int64, p1 uint16, p2 int) int64 {
    fmt.Println(p0, p1, p2)
    p2 = p2
    p2 = (p2 & 128)
    var v5 uint8 = 42
    p1 -= (p1 + 970)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = -1
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    if (1 < 133) {
        fmt.Println(v2)
        v6 := v4[2147483648]
        v7 := GenS1{F0: 9223372036854775807, F1: genFn0(1, 226, 32767), F2: 278}
        _ = v6
        _ = v7
    }
    if true {
        v2 -= (v2 >> ((524) & 63))
        switch 270 {
        case -1:
            if false {
                var v8 int8 = 22
                var v9 string = "\n\t"
                _ = v8
                _ = v9
            }
        case 128:
            for fi10 := 0; fi10 < 20; fi10++ {
                switch fi10 {
                case 2147483647:
                    if (42 == 709) {
                        var v11 float64 = (0.0 + 1.5)
                        for fi12 := 0; fi12 < 1; fi12++ {
                            switch fi12 {
                            case -9223372036854775808:
                                var v13 string = ("foo bar" + "foo bar")
                                _ = v13
                            case 708:
                                v2 = (v2 >> ((0) & 63))
                            case 42:
                                v11 += 0.1
                            default:
                                switch "a" {
                                case "":
                                    var v14 int16 = 1
                                    _ = v14
                                }
                                v15 := GenS1{F0: fi10, F1: 881, F2: 255}
                                _ = v15
                            }
                            _ = fi12
                        }
                        _ = v11
                    } else {
                        for fi16 := 0; fi16 < 2; fi16++ {
                            var v17 uint = 0
                            v18 := GenS2{F0: true, F1: -1, F2: (42 <= -9223372036854775808)}
                            if true {
                                v17 = (v17 ^ (^(v17)))
                                v2 = v2
                                var v19 int = fi10
                                _ = v19
                            }
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                case 4294967295:
                    v2 += v2
                case 752:
                    v20 := GenS1{F0: -2147483648, F1: genFn0(42, 2, -2147483648), F2: 18446744073709551615}
                    _ = v20
                default:
                    v2 -= v2
                }
                _ = fi10
            }
        case 2147483647:
            fmt.Println(v2)
        default:
            var v21 string = "hello"
            _ = v21
        }
    }
    if (((1.5 >= (3.14159 - 0.0)) && false) && false) {
        var v22 int32 = -1
        _ = v22
    }
    v23 := []int{834}
    v3[128] = 1
    if true {
        if false {
            switch 1 {
            case 32767:
                for fi24 := 0; fi24 < 16; fi24++ {
                    if true {
                        v2 -= v2
                        switch (fi24 - fi24) {
                        case 9223372036854775807:
                            for fi25 := 0; fi25 < 19; fi25++ {
                                var v26 int64 = 767
                                _ = fi25
                                _ = v26
                            }
                        case 2147483647:
                            var v27 string = "hello"
                            _ = v27
                        case 1:
                            v2 = (-(v2))
                        default:
                            v28 := GenS2{F0: false, F1: 1, F2: (!true)}
                            _ = v28
                        }
                        fmt.Println(v2)
                    } else {
                        _ = len(v3)
                        v29 := GenS2{F0: (42 < 42), F1: -61, F2: (42 == 897)}
                        for fi30 := 0; fi30 < 3; fi30++ {
                            var v31 float64 = 3.14159
                            var v32 int32 = 1
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        _ = v29
                    }
                    v4[29] = ""
                    var v33 int32 = 42
                    _ = fi24
                    _ = v33
                }
            }
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            _ = len(v23)
        }
    } else {
        var v34 string = ("hello" + "a")
        switch 817 {
        case -479:
            var v35 float32 = 0.0
            _ = v35
        case -964:
            fmt.Println(v34, v2)
        case 42:
            switch -1 {
            case 32767:
                fmt.Println(v34)
            case 718:
                switch 32768 {
                case -1:
                    if len(v23) > 0 {
                        v23[0] = 305
                    }
                default:
                    if (178 >= 4294967295) {
                        fmt.Println(v34, v2)
                        for fi36 := 0; fi36 < 17; fi36++ {
                            for fi37 := 0; fi37 < 20; fi37++ {
                                v2 = v2
                                v2 = -1
                                _ = fi37
                            }
                            _ = fi36
                        }
                    } else {
                        _ = len(v23)
                        if len(v23) > 0 {
                            v23[0] = 0
                        }
                    }
                }
            default:
                v2 = v2
            }
        }
        for fi38 := 0; fi38 < 1; fi38++ {
            fmt.Println(fi38, v2, v34)
            _ = fi38
        }
        _ = v34
    }
    var v39 uint = 42
    var v40 int = 56
    v41 := GenS1{F0: 32768, F1: 9223372036854775807, F2: (v39 | 0)}
    v42 := GenS3{F0: (("a" + "a") + (("\n\t" + ("a" + "hello")) + ("a" + ""))), F1: 2147483647, F2: ("foo bar" + ("\n\t" + "a"))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
