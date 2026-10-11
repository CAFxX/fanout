package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint; F2 bool }
func genFn0(p0 uint8, p1 int8, p2 int64) string {
    v6 := GenS1{F0: (^(p0))}
    if true {
        var v7 int16 = 18
        p2 += (p2 - (p2 + p2))
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "hello"
}

func main() {
    var v1 string = "foo bar"
    var v2 uint = 1
    var v3 uint32 = 2147483648
    var v4 uint16 = 32767
    var v5 int64 = 0
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 string = "a"
        if (0 <= 1) {
            v5 = (v5 >> ((365) & 63))
            fmt.Println(fi8, v4, v5)
            var v10 string = (v1 + "")
            _ = v10
        }
        v4 -= (v4 / ((v4) | 1))
        _ = fi8
        _ = v9
    }
    v1 = ""
    v11 := GenS2{F0: 42, F1: 504, F2: (false || false)}
    v2 = v2
    fmt.Println(v4, v3, v1)
    var v12 uint32 = v3
    var v13 string = (v1 + "hello")
    for fi14 := 0; fi14 < 9; fi14++ {
        if (353 >= -32768) {
            v15 := GenS2{F0: 1, F1: 97, F2: ("foo bar" >= v13)}
            _ = v15
        }
        _ = fi14
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            switch v13 {
            case "\n\t":
                switch v1 {
                case "foo bar":
                    v2 += (v2 / ((v2) | 1))
                case "":
                    var v18 uint16 = 65535
                    _ = v18
                }
            case "foo bar":
                v12 = v3
            case "":
                var v19 uint8 = 42
                _ = v19
            default:
                v20 := GenS2{F0: 2147483647, F1: 2, F2: (float64((((1.5 + 3.14159) * (3.14159 / (1.5 / ((float32(0.0) * 3.14159) - 3.14159)))) + 1e10)) >= (((0.0 * -2.25) - 0.0) + -2.25))}
                _ = v20
            }
            var v21 uint8 = 160
            if true {
                v21 -= v21
            }
            _ = fi17
            _ = v21
        }
        _ = fi16
    }
    var v22 string = ""
    var v23 int64 = (v5 * (-(v5)))
    var v24 int32 = 42
    v3 -= (v3 << ((v2) & 63))
    var v25 uint64 = uint64(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
