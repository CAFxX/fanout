package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint8 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []int{128, -2147483648}
    var v2 uint16 = 1
    var v3 float64 = float64((((-2.25 / 1.0) - 0.1) * 3.14159))
    var v4 int = 288
    var v5 uint64 = 2
    switch "" {
    case "a":
        switch (v4 ^ v4) {
        case -1:
            fmt.Println(v2, v4)
        default:
            v4 = (-(v4))
        }
    case "foo bar":
        v4 = (v4 ^ (v4 + v4))
    default:
        v6 := GenS1{F0: ("a" + "a"), F1: false, F2: 255}
        _ = v6
    }
    v2 = (v2 ^ (v2 << ((65535) & 63)))
    var v7 bool = false
    var v8 uint32 = 992
    v5 -= v5
    var v9 float64 = v3
    if false {
        if true {
            v2 -= 2
        }
    }
    if v7 {
        var v10 int64 = 986
        for fi11 := 0; fi11 < 8; fi11++ {
            v8 -= v8
            _ = fi11
        }
        _ = v10
    }
    v9 = v9
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v12 float32 = 0.1
    switch (v4 * (-(v4))) {
    case -2147483648:
        v2 = (v2 | 0)
    case 253:
        v13 := GenS2{F0: 390}
        _ = v13
    case 32767:
        if v7 {
            var v14 uint16 = 594
            _ = v14
        } else {
            switch ("a" + "\n\t") {
            case "hello":
                var v15 float64 = 0.0
                _ = v15
            case "":
                if ((!(((v3 / v9) != v3) || (v7 || false))) && v7) {
                    for fi16 := 0; fi16 < 10; fi16++ {
                        var v17 bool = (!v7)
                        v12 = (float32(((v3 - 1e10) - v3)) - 0.1)
                        _ = fi16
                        _ = v17
                    }
                    if v7 {
                        for fi18 := 0; fi18 < 20; fi18++ {
                            switch "a" {
                            case "a":
                                var v19 float32 = (0.1 / v12)
                                _ = v19
                            case "":
                                v12 = v12
                            }
                            var v20 int32 = 42
                            v9 -= (((v9 / 1.0) + v3) - (v3 * v3))
                            _ = fi18
                            _ = v20
                        }
                    } else {
                        var v21 string = "a"
                        _ = v21
                    }
                    v1 = append(v1, -693)
                }
            }
            v5 -= 42
            switch v4 {
            case 4294967295:
                var v22 int32 = 2147483647
                _ = v22
            case -535:
                fmt.Println(v9)
            default:
                fmt.Println(v7, v7)
            }
        }
    default:
        var v23 int64 = 42
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
