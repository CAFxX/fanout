package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }

func main() {
    var v1 uint16 = 65535
    var v2 uint32 = 65535
    v3 := make(map[string]uint32)
    var v4 uint32 = (^(v2))
    var v5 uint8 = 1
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 uint16 = uint16(v5)
        _ = fi6
        _ = v7
    }
    if (float32((0.0 / -2.25)) == (0.1 - (float32(1.5) - ((1.5 - (((-2.25 - -2.25) + (0.1 * -2.25)) + (3.14159 * (0.0 * -2.25)))) + 1.5)))) {
        if false {
            v1 += v1
            for fi8 := 0; fi8 < 1; fi8++ {
                v9 := GenS1{F0: -1, F1: (v5 ^ (v5 ^ 29))}
                _ = fi8
                _ = v9
            }
        }
    }
    var v10 string = ("\n\t" + ("\n\t" + "\n\t"))
    if false {
        delete(v3, "")
    } else {
        var v11 uint16 = (^(v1))
        v5 -= 0
        var v12 float32 = 0.1
        _ = v11
        _ = v12
    }
    fmt.Println(v5)
    for fi13 := 0; fi13 < 5; fi13++ {
        v2 = 42
        v4 = v4
        _ = fi13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        switch fi14 {
        case 32768:
            var v15 string = v10
            _ = v15
        case -1:
            var v16 uint = uint(v5)
            _ = v16
        case 127:
            v1 = (v1 % ((v1) | 1))
        default:
            var v17 uint8 = 127
            _ = v17
        }
        _ = fi14
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        if false {
            for fi19 := 0; fi19 < 10; fi19++ {
                v5 = 1
                v20 := []int{2147483647, (^(fi18)), (fi18 | 114), (fi18 % ((fi18) | 1))}
                var v21 string = v10
                v5 = (v5 % ((1) | 1))
                var v22 int = (fi19 | fi19)
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
        }
        _ = fi18
    }
    v23 := []int{9223372036854775807, -1}
    if false {
        v24 := GenS1{F0: 9223372036854775807, F1: (v5 * (v5 % ((127) | 1)))}
        switch v10 {
        case "foo bar":
            var v25 int = -9223372036854775808
            _ = v25
        case "hello":
            switch v10 {
            case "\n\t":
                for fi26 := 0; fi26 < 19; fi26++ {
                    v5 += (^(v5))
                    var v27 uint32 = v2
                    _ = fi26
                    _ = v27
                }
            case "foo bar":
                var v28 int32 = -441
                _ = v28
            default:
                switch ("\n\t" + (("" + v10) + v10)) {
                case "a":
                    for fi29 := 0; fi29 < 3; fi29++ {
                        _ = len(v23)
                        _ = fi29
                    }
                }
            }
        default:
            v10 = v10
        }
        v30 := GenS1{F0: 42, F1: (v5 | 94)}
        _ = v24
        _ = v30
    } else {
        var v31 uint16 = v1
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        switch 127 {
        case 9223372036854775807:
            var v32 int8 = -89
            _ = v32
        default:
            fmt.Println(v5, v10)
        }
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
