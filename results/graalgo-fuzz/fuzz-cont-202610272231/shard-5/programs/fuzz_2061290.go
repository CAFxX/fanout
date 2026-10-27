package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 bool; F2 bool }
func genFn0() uint8 {
    switch 1 {
    case 2147483647:
        if ((-267 == -9223372036854775808) || (true || false)) {
            var v7 int16 = 42
            v8 := []int{4294967295, -2147483648, 1, 684}
            var v9 bool = (!("a" == "a"))
            _ = v7
            _ = v8
            _ = v9
        }
    case -9223372036854775808:
        if ((32768 != 128) && (65535 <= 42)) {
            v10 := GenS1{F0: 772, F1: 0, F2: 255}
            _ = v10
        }
    }
    var v11 bool = (9223372036854775807 > 4294967295)
    _ = v11
    return 0
}

func main() {
    v1 := []string{"foo bar"}
    var v2 uint64 = 721
    var v3 float32 = (3.14159 - (0.1 + 3.14159))
    var v4 float64 = 1.5
    var v5 uint = 42
    v6 := []uint32{}
    v12 := GenS1{F0: 1, F1: -295, F2: uint16(v2)}
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS2{F0: -128}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 uint = 255
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        fmt.Println(v4, v4)
        v5 = v5
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi17
    }
    var v18 float32 = 1.5
    v19 := GenS1{F0: -32768, F1: -483, F2: 651}
    fmt.Println(v2, v4)
    fmt.Println(v4, v18)
    if false {
        _ = len(v6)
        for fi20 := 0; fi20 < 1; fi20++ {
            var v21 uint8 = 154
            _ = fi20
            _ = v21
        }
    } else {
        v3 -= v3
        v22 := GenS3{F0: "a", F1: (0 > 562), F2: true}
        fmt.Println(v5, v4, v4)
        _ = v22
    }
    if true {
        v18 = v3
    }
    if ((v4 < float64(v3)) || (true || false)) {
        v23 := GenS1{F0: -140, F1: 9223372036854775807, F2: 42}
        var v24 int8 = 42
        for fi25 := 0; fi25 < 4; fi25++ {
            v18 -= v18
            if false {
                v3 += ((v3 / float32((1e10 * ((v4 - v4) / 1.0)))) + (v3 / v18))
                for fi26 := 0; fi26 < 15; fi26++ {
                    if len(v6) > 0 {
                        v6[0] = 0
                    }
                    _ = fi26
                }
            }
            if true {
                if len(v6) > 0 {
                    v6[0] = 2147483648
                }
                fmt.Println(v5, v18, v4)
            } else {
                switch "" {
                case "foo bar":
                    if true {
                        v5 -= v5
                        for fi27 := 0; fi27 < 11; fi27++ {
                            if ((^(v2)) == v2) {
                                if true {
                                    fmt.Println(v5, fi25)
                                    var v28 float32 = (v3 * (float32(1e10) - (0.0 * (0.0 / v3))))
                                    _ = v28
                                }
                                if ((v2 > v2) || (!true)) {
                                    if len(v1) > 0 {
                                        v1[0] = "foo bar"
                                    }
                                } else {
                                    switch "hello" {
                                    case "a":
                                        _ = len(v1)
                                    }
                                }
                            }
                            _ = fi27
                        }
                        var v29 float32 = v3
                        _ = v29
                    } else {
                        switch fi25 {
                        case -927:
                            v4 = (((v4 + 0.0) / 0.1) - (v4 / v4))
                        }
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                default:
                    fmt.Println(v18)
                }
                if true {
                    for fi30 := 0; fi30 < 16; fi30++ {
                        if (784 >= 700) {
                            if len(v6) > 0 {
                                v6[0] = 1
                            }
                            var v31 uint = 4294967295
                            _ = v31
                        }
                        _ = fi30
                    }
                }
            }
            _ = fi25
        }
        _ = v23
        _ = v24
    }
    v32 := GenS2{F0: 0}
    if ((("\n\t" + ("" + "")) + "foo bar") <= "") {
        if (1 > 223) {
            fmt.Println(v5)
        } else {
            var v33 int32 = 2147483647
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 8; fi34++ {
        if false {
            if true {
                for fi35 := 0; fi35 < 11; fi35++ {
                    var v36 int16 = 849
                    var v37 uint8 = 2
                    v3 += v3
                    var v38 uint64 = (^(v2))
                    _ = fi35
                    _ = v36
                    _ = v37
                    _ = v38
                }
            }
            switch "" {
            case "a":
                var v39 float32 = ((float32(v4) + 0.0) + float32(v4))
                _ = v39
            default:
                var v40 bool = (!true)
                _ = v40
            }
            var v41 uint8 = 255
            _ = v41
        } else {
            v18 = (-2.25 * (v3 / (v18 + (-2.25 / 3.14159))))
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            fmt.Println(v4, v4, v5)
        }
        fmt.Println(fi34)
        _ = fi34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
