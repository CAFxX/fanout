package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint8; F1 bool; F2 int32 }

func main() {
    v1 := []uint32{36}
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v4 int8 = 100
    v5 := make(map[int]string)
    v6 := make(map[int]bool)
    v7 := GenS1{F0: int16(v4), F1: "\n\t"}
    v4 += v4
    v4 = (v4 / ((v4) | 1))
    v4 = v4
    var v8 uint64 = 42
    switch 2147483647 {
    case 667:
        var v9 string = "hello"
        _ = v9
    case -477:
        switch -9223372036854775808 {
        case 32767:
            var v10 uint64 = v8
            _ = v10
        default:
            v4 = v4
        }
    case -739:
        if (3.14159 < -2.25) {
            for fi11 := 0; fi11 < 11; fi11++ {
                v4 = v4
                for fi12 := 0; fi12 < 10; fi12++ {
                    v4 = (-(v4))
                    v4 = (v4 * (^(v4)))
                    if len(v1) > 0 {
                        v1[0] = 660
                    }
                    _ = fi12
                }
                _ = fi11
            }
            v13 := GenS1{F0: 0, F1: ("a" + "\n\t")}
            _ = v13
        }
    }
    switch -2147483648 {
    case -867:
        v1 = append(v1, 1)
    case -885:
        switch 32767 {
        case -872:
            switch "a" {
            case "hello":
                if (false || true) {
                    var v14 bool = false
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 13; fi15++ {
                        v16 := GenS1{F0: 0, F1: "foo bar"}
                        _ = fi15
                        _ = v16
                    }
                    if false {
                        var v17 float32 = 0.0
                        v8 = v8
                        v18 := v6[0]
                        _ = v17
                        _ = v18
                    } else {
                        v8 += v8
                    }
                    var v19 uint8 = 100
                    _ = v19
                }
            case "foo bar":
                switch 128 {
                case -625:
                    v4 -= -30
                case 361:
                    for fi20 := 0; fi20 < 17; fi20++ {
                        switch "" {
                        case "foo bar":
                            if (-1 != 0) {
                                var v21 int16 = 1000
                                _ = v21
                            }
                        case "hello":
                            if true {
                                fmt.Println(v4)
                            } else {
                                fmt.Println(fi20, v4, v4)
                                v22 := GenS1{F0: -138, F1: (("\n\t" + "hello") + ("\n\t" + "hello"))}
                                switch (fi20 ^ (fi20 * -524)) {
                                case 2147483647:
                                    v8 = v8
                                case 127:
                                    v4 -= (v4 & v4)
                                case 390:
                                    if false {
                                        var v23 uint32 = 2147483648
                                        var v24 uint8 = 93
                                        _ = v23
                                        _ = v24
                                    }
                                }
                                _ = v22
                            }
                        }
                        _ = fi20
                    }
                case 32768:
                    var v25 float32 = -2.25
                    _ = v25
                }
            }
        default:
            _ = len(v6)
        }
    default:
        var v26 uint = 279
        _ = v26
    }
    if false {
        if (!false) {
            var v27 int8 = v4
            _ = v27
        } else {
            for fi28 := 0; fi28 < 9; fi28++ {
                _ = len(v1)
                var v29 uint = 103
                _ = fi28
                _ = v29
            }
            for fi30 := 0; fi30 < 14; fi30++ {
                fmt.Println(v8, v8, v4)
                _ = fi30
            }
        }
    } else {
        var v31 uint16 = 42
        v32 := GenS2{F0: 42, F1: (3.14159 != 0.0), F2: -2147483648}
        fmt.Println(v4)
        _ = v31
        _ = v32
    }
    delete(v5, 746)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
