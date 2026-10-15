package main

import (
    "fmt"
)

func genFn0(p0 uint8) int16 {
    p0 -= p0
    p0 = p0
    _ = p0
    return 240
}

func main() {
    var v1 int8 = -1
    var v2 uint32 = 200
    var v3 int = 128
    v4 := []uint32{v2, 785}
    var v5 uint32 = (v2 * (v2 % ((v2) | 1)))
    v3 -= (v3 ^ (v3 | int(v5)))
    if (true || false) {
        var v6 int64 = 9223372036854775807
        _ = v6
    } else {
        for fi7 := 0; fi7 < 11; fi7++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v8 int8 = v1
            _ = fi7
            _ = v8
        }
        switch "foo bar" {
        case "hello":
            var v9 int64 = 42
            _ = v9
        default:
            var v10 string = "\n\t"
            _ = v10
        }
    }
    var v11 float64 = 1e10
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (!(((false && false) || false) && false)) {
        for fi12 := 0; fi12 < 4; fi12++ {
            fmt.Println(v3, v1, v3)
            var v13 int32 = -208
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v4 = append(v4, (^(v2)))
        _ = fi14
    }
    var v15 uint32 = 705
    fmt.Println(v1)
    if len(v4) > 0 {
        v4[0] = 0
    }
    var v16 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
