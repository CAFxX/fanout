package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() uint {
    v4 := []int{-264, 42, -1, -2147483648}
    var v5 float64 = 0.1
    fmt.Println(v5)
    var v6 uint64 = 1
    var v7 int = 42
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 int8 = -65
    var v2 uint8 = 57
    v3 := make(map[int]uint32)
    v8 := []int{128, 127, -481, 128}
    _ = len(v3)
    switch 678 {
    case -184:
        switch 2 {
        case -9223372036854775808:
            switch (("hello" + "foo bar") + (("" + "") + "")) {
            case "hello":
                switch "foo bar" {
                case "":
                    fmt.Println(v1, v1)
                default:
                    v2 -= v2
                }
            }
        case -2147483648:
            switch 4294967295 {
            case 128:
                v1 += v1
            case -754:
                v9 := GenS1{F0: (0 < 2)}
                _ = v9
            case 0:
                switch 2147483647 {
                case 128:
                    var v10 int16 = -1
                    _ = v10
                case 525:
                    switch -1 {
                    case 4294967295:
                        v1 = (v1 | (^(v1)))
                    case 32767:
                        v2 += 235
                    }
                case 4294967295:
                    _ = len(v8)
                default:
                    v11 := GenS1{F0: (-1 > -1)}
                    _ = v11
                }
            }
        case -670:
            var v12 int64 = -99
            _ = v12
        }
    default:
        v2 -= v2
    }
    if ((false || (2 == 799)) || ((1.5 * (((((0.0 + -2.25) * 1.5) / 1.5) * 1.5) * -2.25)) == 1.5)) {
        for fi13 := 0; fi13 < 14; fi13++ {
            for fi14 := 0; fi14 < 16; fi14++ {
                fmt.Println(v1, fi14, v1)
                var v15 uint16 = 0
                v1 -= (v1 ^ v1)
                v1 = 127
                _ = fi14
                _ = v15
            }
            fmt.Println(fi13, v1, v1)
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 int = (fi16 % ((fi13) | 1))
                var v18 float64 = (0.1 / 1.5)
                if true {
                    var v19 bool = false
                    _ = v19
                }
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi13
        }
    }
    v20 := GenS1{F0: (!(!true))}
    var v21 int16 = int16(v1)
    fmt.Println(v1, v1)
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 int64 = -262
        fmt.Println(v2)
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        v21 = v21
        if false {
            for fi25 := 0; fi25 < 17; fi25++ {
                var v26 uint32 = 65535
                _ = fi25
                _ = v26
            }
        } else {
            var v27 uint16 = 65535
            _ = v27
        }
        v1 = (v1 ^ (^(v1)))
        _ = fi24
    }
    if len(v8) > 0 {
        v8[0] = 1
    }
    var v28 uint = 4294967295
    var v29 uint8 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v8
}
