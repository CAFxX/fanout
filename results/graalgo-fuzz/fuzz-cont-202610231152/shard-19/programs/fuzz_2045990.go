package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint8 = 42
    var v2 int32 = 2147483647
    v3 := make(map[string]uint32)
    v4 := []uint32{808}
    v2 += -1
    if (v2 != (v2 << ((0) & 63))) {
        fmt.Println(v2, v1)
        if ((v1 > 1) && true) {
            v1 = v1
            var v5 int32 = 409
            _ = v5
        }
        for fi6 := 0; fi6 < 7; fi6++ {
            switch fi6 {
            case -824:
                fmt.Println(v2, fi6)
            default:
                switch "" {
                case "foo bar":
                    delete(v3, "a")
                case "a":
                    var v7 float64 = (1e10 / (0.0 + -2.25))
                    _ = v7
                case "":
                    fmt.Println(v1)
                default:
                    v1 -= 231
                }
            }
            var v8 uint64 = 42
            _ = fi6
            _ = v8
        }
    } else {
        v9 := GenS1{F0: (821 <= 128)}
        var v10 int16 = -32768
        switch 32767 {
        case -1:
            v1 = v1
        default:
            v11 := v3[("foo bar" + "a")]
            _ = v11
        }
        _ = v10
        _ = v9
    }
    switch "" {
    case "":
        switch 32767 {
        case 0:
            var v12 uint64 = 126
            _ = v12
        case 4294967295:
            v2 -= (v2 % ((-41) | 1))
        case 2147483647:
            _ = len(v4)
        default:
            fmt.Println(v1, v2)
        }
    case "a":
        switch 2147483647 {
        case 4294967295:
            v13 := GenS1{F0: (2 > 0)}
            _ = v13
        case 75:
            var v14 int32 = (-(v2))
            _ = v14
        default:
            v1 = v1
        }
    }
    v3[(("hello" + "\n\t") + ((("a" + "\n\t") + "foo bar") + ("hello" + ("foo bar" + "foo bar"))))] = 4294967295
    v2 -= v2
    v1 = (v1 % ((v1) | 1))
    var v15 string = "\n\t"
    var v16 uint = 109
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
