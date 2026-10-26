package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 bool = ((-2147483648 >= 42) || false)
    var v3 float64 = -2.25
    v4 := []uint32{2, 62, 2147483648, 4294967295}
    delete(v1, ("foo bar" + "hello"))
    v5 := GenS1{F0: -1, F1: 1}
    v1["\n\t"] = (("foo bar" + ("a" + "a")) + (("\n\t" + "a") + ("" + "foo bar")))
    v3 += v3
    switch -1 {
    case 32767:
        switch 32768 {
        case 960:
            v2 = (0 <= 906)
        case 32767:
            fmt.Println(v3)
        default:
            var v6 uint8 = 2
            _ = v6
        }
    case -1:
        var v7 string = ""
        _ = v7
    default:
        if (0 >= 1) {
            if v2 {
                v8 := GenS2{F0: -495}
                _ = v8
            } else {
                var v9 float64 = v3
                v3 -= v3
                _ = v9
            }
            v2 = v2
            var v10 int = -495
            _ = v10
        } else {
            fmt.Println(v2, v2)
        }
    }
    var v11 float32 = 3.14159
    v3 -= -2.25
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            v14 := v1["\n\t"]
            v15 := GenS1{F0: 32767, F1: 71}
            fmt.Println(v14)
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 uint32 = 2147483647
        _ = fi12
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
