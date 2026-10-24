package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 float64 = 0.1
    var v2 bool = true
    v3 := []string{"hello", ("\n\t" + ("\n\t" + "\n\t")), "a", ("foo bar" + ("hello" + ((("a" + "") + (("a" + "foo bar") + "foo bar")) + ("hello" + "a"))))}
    v4 := make(map[int]int)
    var v5 bool = (4294967295 != 0)
    v5 = v2
    if (1 == 128) {
        var v6 int64 = 9223372036854775807
        _ = v6
    } else {
        if (((2 < 2) || v2) || (127 <= 0)) {
            if false {
                var v7 int = -1
                v3 = append(v3, "a")
                for fi8 := 0; fi8 < 14; fi8++ {
                    var v9 uint8 = 127
                    _ = fi8
                    _ = v9
                }
                _ = v7
            } else {
                v1 += (((float64((3.14159 * 1.5)) / v1) - -2.25) / float64(float32(v1)))
            }
            v10 := v4[297]
            _ = v10
        }
        switch "\n\t" {
        case "":
            fmt.Println(v1, v1)
        default:
            v1 = float64(1.5)
        }
        if (2 >= 4294967295) {
            var v11 int16 = -1
            _ = v11
        }
    }
    switch 127 {
    case 1:
        for fi12 := 0; fi12 < 15; fi12++ {
            v1 -= ((0.1 + float64(float32(v1))) / float64(1e10))
            _ = fi12
        }
    default:
        var v13 float32 = float32(0.0)
        _ = v13
    }
    switch ("a" + "hello") {
    case "":
        var v14 float32 = (0.0 / 1.0)
        _ = v14
    }
    v15 := v4[191]
    for fi16 := 0; fi16 < 12; fi16++ {
        fmt.Println(v15)
        _ = fi16
    }
    var v17 int64 = -706
    fmt.Println(v2, v15, v15)
    fmt.Println(v15, v2)
    var v18 uint8 = 255
    v19 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
