package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 bool }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 1
    v3 := []uint32{65535, 2147483647, 417}
    switch "foo bar" {
    case "foo bar":
        switch 23 {
        case 4294967295:
            for fi4 := 0; fi4 < 4; fi4++ {
                v2 -= (v2 | (v2 * (v2 - 0)))
                var v5 float32 = 0.1
                _ = fi4
                _ = v5
            }
        }
    case "\n\t":
        v6 := v1[165]
        _ = v6
    case "":
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 int8 = 1
            _ = fi7
            _ = v8
        }
    }
    if ((v2 << ((0) & 63)) >= v2) {
        delete(v1, 42)
    }
    v2 += v2
    if (2 > 1) {
        var v9 uint16 = 769
        _ = v9
    } else {
        var v10 uint8 = 127
        _ = v10
    }
    var v11 int16 = 42
    if false {
        if (255 > 127) {
            v2 -= (v2 & v2)
            v12 := GenS1{F0: 255, F1: 255, F2: (167 == -930)}
            v2 = v2
            _ = v12
        } else {
            v11 = (v11 | -1)
            if true {
                switch 2147483647 {
                case 32767:
                    var v13 float32 = 0.0
                    _ = v13
                default:
                    v2 = 356
                }
            } else {
                var v14 bool = (false || (-127 >= 43))
                var v15 float64 = 0.0
                switch "hello" {
                case "\n\t":
                    v15 = (v15 * (1e10 - ((0.0 + (0.0 * 0.1)) - v15)))
                case "a":
                    var v16 string = "foo bar"
                    _ = v16
                case "hello":
                    var v17 string = "hello"
                    _ = v17
                default:
                    var v18 float64 = 0.1
                    _ = v18
                }
                _ = v14
                _ = v15
            }
        }
    }
    v19 := GenS1{F0: 76, F1: 0, F2: false}
    v20 := GenS1{F0: 2, F1: 255, F2: (18446744073709551615 != 18446744073709551615)}
    var v21 uint16 = 65535
    fmt.Println(v21)
    delete(v1, 1)
    if (!(3.14159 < -2.25)) {
        var v22 int8 = -31
        _ = v22
    } else {
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        var v23 int16 = (v11 >> ((0) & 63))
        _ = v23
    }
    if true {
        if false {
            v21 = 42
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v24 := GenS2{F0: 2147483648}
            _ = v24
        }
    } else {
        switch 4294967295 {
        case 4294967295:
            switch -9223372036854775808 {
            case 1:
                var v25 int = 311
                _ = v25
            case -903:
                v1[42] = ""
            }
        default:
            v26 := GenS1{F0: 2, F1: 42, F2: false}
            _ = v26
        }
        _ = len(v1)
    }
    v27 := GenS1{F0: 1, F1: 44, F2: (!(0.0 != 1e10))}
    switch 9223372036854775807 {
    case 74:
        if false {
            v28 := v1[42]
            _ = v28
        } else {
            var v29 uint32 = 784
            var v30 int = -986
            _ = v29
            _ = v30
        }
    case 0:
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
}
