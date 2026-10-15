package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }

func main() {
    var v1 uint64 = 369
    var v2 uint64 = 42
    v3 := []uint32{}
    v4 := make(map[string]bool)
    v5 := make(map[uint32]bool)
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 int8 = -128
            var v9 uint8 = 0
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 string = "hello"
        _ = fi6
        _ = v10
    }
    v11 := GenS1{F0: 812, F1: "foo bar"}
    for fi12 := 0; fi12 < 17; fi12++ {
        if (525 >= 0) {
            var v13 int32 = 485
            for fi14 := 0; fi14 < 16; fi14++ {
                switch "" {
                case "":
                    if len(v3) > 0 {
                        v3[0] = 893
                    }
                case "\n\t":
                    var v15 int64 = 1
                    _ = v15
                }
                _ = fi14
            }
            delete(v5, 4294967295)
            _ = v13
        }
        v4[""] = false
        var v16 uint64 = (^(v1))
        _ = fi12
        _ = v16
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 string = "hello"
        var v19 int16 = 0
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(v1, v2)
    switch -2147483648 {
    case 647:
        for fi20 := 0; fi20 < 17; fi20++ {
            var v21 float64 = 3.14159
            for fi22 := 0; fi22 < 14; fi22++ {
                if true {
                    var v23 string = "a"
                    _ = v23
                } else {
                    switch fi22 {
                    case -9223372036854775808:
                        var v24 int64 = 234
                        _ = v24
                    }
                }
                _ = fi22
            }
            _ = fi20
            _ = v21
        }
    }
    if (3.14159 != (3.14159 / ((1e10 / 1.0) / 1.0))) {
        for fi25 := 0; fi25 < 6; fi25++ {
            v1 += v2
            var v26 uint8 = 92
            _ = fi25
            _ = v26
        }
    } else {
        v1 += v1
        switch (("a" + "hello") + ("" + "hello")) {
        case "a":
            fmt.Println(v2)
        case "\n\t":
            var v27 int16 = 350
            _ = v27
        case "":
            var v28 string = ("hello" + ("" + "\n\t"))
            _ = v28
        default:
            if (!false) {
                v2 = v2
                switch "" {
                case "\n\t":
                    v1 -= v1
                case "a":
                    v3 = append(v3, 487)
                case "hello":
                    if len(v3) > 0 {
                        v3[0] = 4294967295
                    }
                default:
                    switch 1 {
                    case 127:
                        switch 32768 {
                        case 1:
                            v2 = 609
                        case -34:
                            var v29 string = "\n\t"
                            _ = v29
                        case -1:
                            for fi30 := 0; fi30 < 4; fi30++ {
                                var v31 int8 = 0
                                var v32 uint64 = v1
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        }
                    default:
                        var v33 bool = true
                        _ = v33
                    }
                }
            }
        }
        var v34 int = 42
        _ = v34
    }
    var v35 uint64 = 2
    v35 += (v2 << ((18446744073709551615) & 63))
    fmt.Println(v1)
    v35 += v35
    fmt.Println(v35, v2)
    v35 += (^(v1))
    fmt.Println(v2, v2, v35)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
