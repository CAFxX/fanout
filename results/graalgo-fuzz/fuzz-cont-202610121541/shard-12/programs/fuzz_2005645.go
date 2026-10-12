package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 string; F2 string }

func main() {
    var v1 string = ("hello" + "a")
    var v2 int16 = 42
    var v3 int8 = 1
    var v4 bool = false
    v5 := make(map[uint32]int)
    var v6 uint64 = 4294967295
    switch (v1 + v1) {
    case "hello":
        for fi7 := 0; fi7 < 11; fi7++ {
            v3 = v3
            _ = len(v5)
            _ = fi7
        }
    case "foo bar":
        var v8 int8 = v3
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: 2, F2: 9223372036854775807}
    v10 := GenS1{F0: 91, F1: 648, F2: -9223372036854775808}
    switch 297 {
    case 42:
        var v11 float64 = -2.25
        _ = v11
    case 127:
        var v12 int8 = v3
        _ = v12
    case -1:
        v4 = ((v2 != 1) && (!true))
    }
    var v13 bool = false
    switch "a" {
    case "hello":
        var v14 string = v1
        _ = v14
    default:
        v15 := GenS1{F0: 2, F1: 255, F2: 9223372036854775807}
        _ = v15
    }
    if true {
        if false {
            for fi16 := 0; fi16 < 13; fi16++ {
                v17 := GenS2{F0: 65535, F1: 2, F2: 255}
                _ = fi16
                _ = v17
            }
            delete(v5, 628)
            var v18 int = -2147483648
            _ = v18
        }
    } else {
        fmt.Println(v2, v1)
        if ((((((1.5 / 1e10) / 3.14159) * -2.25) / ((((0.1 + 1.5) + 0.1) / 3.14159) / 1.0)) + 0.0) != 1e10) {
            var v19 string = v1
            fmt.Println(v2)
            _ = v19
        } else {
            var v20 uint8 = 2
            if ((v2 % ((v2) | 1)) == v2) {
                _ = len(v5)
                switch v1 {
                case "\n\t":
                    switch 2147483647 {
                    case 9223372036854775807:
                        var v21 int8 = -128
                        _ = v21
                    case 803:
                        var v22 int32 = 1
                        _ = v22
                    case -404:
                        fmt.Println(v13, v20)
                    default:
                        for fi23 := 0; fi23 < 6; fi23++ {
                            switch fi23 {
                            case 1:
                                v24 := GenS3{F0: v1, F1: 0}
                                _ = v24
                            case 32768:
                                v3 -= v3
                            }
                            switch "a" {
                            case "a":
                                var v25 float64 = (-2.25 + (((-2.25 * 0.0) / 1e10) * 3.14159))
                                _ = v25
                            }
                            for fi26 := 0; fi26 < 15; fi26++ {
                                v20 -= (v20 >> ((829) & 63))
                                var v27 int = (fi23 - (fi23 >> ((140) & 63)))
                                switch (v1 + "a") {
                                case "foo bar":
                                    var v28 float32 = 0.0
                                    _ = v28
                                default:
                                    delete(v5, 0)
                                }
                                _ = fi26
                                _ = v27
                            }
                            _ = fi23
                        }
                    }
                case "a":
                    v29 := GenS2{F0: 42, F1: (v20 - v20), F2: 255}
                    _ = v29
                }
                v30 := GenS4{F0: v1}
                _ = v30
            }
            var v31 float64 = float64(1e10)
            _ = v20
            _ = v31
        }
        var v32 uint32 = 1
        _ = v32
    }
    v33 := GenS5{F0: 1, F1: "", F2: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
