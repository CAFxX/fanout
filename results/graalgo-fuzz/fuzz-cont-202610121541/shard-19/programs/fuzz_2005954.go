package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int8 }

func main() {
    v1 := make(map[string]bool)
    var v2 int = 1
    var v3 uint8 = 2
    var v4 float32 = 3.14159
    _ = len(v1)
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            v7 := GenS1{F0: -2147483648}
            switch (fi6 << ((849) & 63)) {
            case 159:
                var v8 uint32 = 42
                _ = v8
            case -1:
                v9 := GenS1{F0: 1}
                _ = v9
            default:
                v2 -= v2
            }
            v3 = (^(v3))
            _ = fi6
            _ = v7
        }
        var v10 int64 = 9223372036854775807
        _ = fi5
        _ = v10
    }
    fmt.Println(v2, v4, v2)
    v3 += 127
    var v11 float64 = 0.0
    if false {
        v12 := GenS1{F0: 107}
        _ = len(v1)
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        switch (fi13 * (fi13 + (fi13 / ((fi13) | 1)))) {
        case 128:
            if false {
                switch v2 {
                case 32767:
                    for fi14 := 0; fi14 < 16; fi14++ {
                        v2 -= (^(fi13))
                        _ = fi14
                    }
                case -2147483648:
                    v1[((("hello" + ("a" + "a")) + ("\n\t" + ("hello" + ("foo bar" + ("a" + "hello"))))) + ("\n\t" + "hello"))] = (!true)
                default:
                    v15 := GenS1{F0: -758}
                    _ = v15
                }
            } else {
                for fi16 := 0; fi16 < 15; fi16++ {
                    var v17 int = (fi13 - (^(fi16)))
                    _ = fi16
                    _ = v17
                }
            }
        case 42:
            if (!true) {
                var v18 uint64 = 2
                _ = v18
            }
        }
        if (311 < -1) {
            fmt.Println(v11)
            var v19 int16 = -32768
            if (3.14159 <= v4) {
                for fi20 := 0; fi20 < 19; fi20++ {
                    v4 = v4
                    var v21 int = (fi20 / (((v2 & 32767)) | 1))
                    fmt.Println(fi13)
                    _ = fi20
                    _ = v21
                }
                fmt.Println(v4)
            }
            _ = v19
        } else {
            for fi22 := 0; fi22 < 15; fi22++ {
                switch fi22 {
                case 1:
                    switch ("\n\t" + ("foo bar" + "\n\t")) {
                    case "hello":
                        var v23 int32 = 588
                        _ = v23
                    default:
                        fmt.Println(v11, v3)
                    }
                    v24 := []int{}
                    if ((!true) || false) {
                        var v25 int64 = 42
                        _ = len(v24)
                        _ = v25
                    }
                    _ = v24
                case -601:
                    var v26 uint8 = (v3 - v3)
                    _ = v26
                default:
                    var v27 uint = 65535
                    _ = v27
                }
                _ = fi22
            }
        }
        _ = fi13
    }
    switch "\n\t" {
    case "a":
        v2 = (v2 * (v2 | v2))
    case "hello":
        v11 = (v11 - 1.5)
    default:
        v2 = (^(v2))
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 float32 = v4
        var v30 int = (fi28 * 2147483647)
        _ = fi28
        _ = v29
        _ = v30
    }
    v31 := GenS2{F0: 65535, F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
