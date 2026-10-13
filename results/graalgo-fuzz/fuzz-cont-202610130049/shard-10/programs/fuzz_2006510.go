package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 bool; F1 bool }

func main() {
    var v1 int64 = -525
    var v2 int8 = -1
    v3 := make(map[uint32]int)
    v4 := []string{"", ("a" + ("hello" + ("" + "a"))), "a"}
    var v5 uint8 = 151
    v6 := GenS1{F0: 2147483647, F1: "a"}
    if false {
        var v7 uint = 4294967295
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 uint8 = (v5 | v5)
            fmt.Println(v9, fi8, v5)
            v7 -= v7
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        if true {
            _ = len(v4)
            v2 += (v2 ^ (v2 >> ((573) & 63)))
            for fi10 := 0; fi10 < 13; fi10++ {
                v11 := GenS2{F0: (!((1.5 + (-2.25 + 1e10)) >= float64(1e10))), F1: false}
                if (1e10 != float32(float64(float32(float64(0.1))))) {
                    switch ("hello" + (("" + ("\n\t" + "\n\t")) + ("" + ("\n\t" + "hello")))) {
                    case "hello":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case "\n\t":
                        v2 = v2
                    }
                    _ = len(v3)
                } else {
                    for fi12 := 0; fi12 < 8; fi12++ {
                        var v13 int = (fi12 - fi10)
                        _ = fi12
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 20; fi14++ {
                        v5 -= v5
                        _ = fi14
                    }
                    fmt.Println(v2)
                }
                _ = fi10
                _ = v11
            }
        }
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        v2 -= v2
        v16 := GenS2{F0: false, F1: false}
        v1 = (v1 - v1)
        _ = fi15
        _ = v16
    }
    if true {
        switch "a" {
        case "foo bar":
            var v17 int64 = -9223372036854775808
            _ = v17
        default:
            switch 1 {
            case -425:
                var v18 bool = (((v2 ^ v2) == -128) || true)
                _ = v18
            case 9223372036854775807:
                v19 := v3[65535]
                _ = v19
            default:
                v4 = append(v4, "hello")
            }
        }
        v3[1] = 7
    }
    v3[2] = 128
    delete(v3, 42)
    if (9223372036854775807 <= 9223372036854775807) {
        for fi20 := 0; fi20 < 11; fi20++ {
            v5 = 42
            fmt.Println(v2, fi20)
            _ = fi20
        }
        var v21 int8 = v2
        _ = v21
    } else {
        v22 := GenS2{F0: false, F1: ((^(v5)) <= v5)}
        _ = v22
    }
    v2 = (v2 >> ((631) & 63))
    v1 = v1
    var v23 uint16 = 977
    switch "" {
    case "a":
        switch "a" {
        case "hello":
            switch -967 {
            case 960:
                var v24 uint64 = 0
                _ = v24
            case 314:
                var v25 float64 = (0.0 / 1.0)
                _ = v25
            default:
                v5 -= 128
            }
        case "":
            for fi26 := 0; fi26 < 9; fi26++ {
                if (-2.25 >= 3.14159) {
                    v23 -= v23
                    fmt.Println(v1, v1, v2)
                    if (412 < 2147483648) {
                        fmt.Println(v5, fi26)
                        delete(v3, 0)
                        var v27 int8 = v2
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 5; fi28++ {
                            var v29 float32 = (float32((-2.25 - 1.5)) - 0.0)
                            delete(v3, 4294967295)
                            v29 -= (v29 + (v29 + v29))
                            var v30 int32 = 0
                            var v31 uint16 = v23
                            _ = fi28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                    }
                }
                _ = fi26
            }
        default:
            v32 := v3[117]
            _ = v32
        }
    case "\n\t":
        var v33 int32 = 2147483647
        _ = v33
    default:
        var v34 uint = 65535
        _ = v34
    }
    v23 -= (v23 - (v23 >> ((850) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
