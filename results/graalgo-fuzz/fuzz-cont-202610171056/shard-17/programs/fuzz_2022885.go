package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint }

func main() {
    var v1 float64 = (3.14159 * 0.1)
    var v2 bool = (4294967295 == 1)
    v3 := []int{495, -330, 0}
    var v4 float64 = (v1 * v1)
    v5 := GenS1{F0: -1, F1: 0, F2: 16}
    switch (((("" + "hello") + (("foo bar" + ("hello" + "a")) + ("foo bar" + "\n\t"))) + "a") + "a") {
    case "hello":
        var v6 uint8 = 116
        _ = v6
    default:
        if v2 {
            if len(v3) > 0 {
                v3[0] = 899
            }
            if (-29 >= -128) {
                v1 = v4
            }
        }
    }
    v2 = true
    var v7 int = 4294967295
    for fi8 := 0; fi8 < 4; fi8++ {
        v1 += (float64(-2.25) * (v4 / (v4 * float64(0.1))))
        var v9 uint32 = 42
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: -14, F1: 9223372036854775807, F2: 65535}
        _ = fi10
        _ = v11
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v2)
    case "":
        var v12 int8 = 0
        _ = v12
    }
    if (127 != 127) {
        if (v4 != 1e10) {
            v13 := GenS1{F0: -2147483648, F1: 9223372036854775807, F2: 2}
            _ = v13
        } else {
            v3 = append(v3, 42)
            var v14 int32 = -2147483648
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 4; fi15++ {
            for fi16 := 0; fi16 < 15; fi16++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi16
            }
            var v17 bool = ((((0.0 * -2.25) / (((0.1 / 0.1) + 3.14159) + 1e10)) == -2.25) && (v7 >= fi15))
            _ = fi15
            _ = v17
        }
    }
    v2 = (!(-1 <= 42))
    for fi18 := 0; fi18 < 4; fi18++ {
        v1 = float64(1.5)
        v1 = -2.25
        _ = fi18
    }
    v2 = v2
    _ = len(v3)
    var v19 string = ""
    var v20 bool = (((((0.1 + 1e10) - 1e10) - 1e10) - 1e10) != 3.14159)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
