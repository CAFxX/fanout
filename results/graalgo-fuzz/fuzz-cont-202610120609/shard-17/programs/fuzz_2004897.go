package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint32 = 881
    var v2 float32 = float32(3.14159)
    var v3 uint64 = 2
    var v4 string = ("" + (("" + "foo bar") + ("\n\t" + "")))
    v5 := []string{}
    if (v4 != (("\n\t" + "foo bar") + "\n\t")) {
        var v6 int8 = 74
        _ = v6
    } else {
        if len(v5) > 0 {
            v5[0] = v4
        }
    }
    v7 := GenS1{F0: 107}
    var v8 uint64 = (v3 + v3)
    var v9 uint8 = 128
    var v10 int64 = 42
    _ = len(v5)
    v1 += (v1 + (v1 + v1))
    for fi11 := 0; fi11 < 4; fi11++ {
        switch v4 {
        case "hello":
            var v12 bool = (-9223372036854775808 != 9223372036854775807)
            _ = v12
        default:
            v4 = ((("hello" + "hello") + ("" + "a")) + v4)
        }
        _ = fi11
    }
    if ((-9223372036854775808 <= v10) || (4294967295 < 42)) {
        switch v4 {
        case "foo bar":
            if (false && true) {
                v1 = (^(v1))
                fmt.Println(v4)
            }
        default:
            var v13 int16 = -32768
            _ = v13
        }
        v4 = "a"
    }
    v3 -= v3
    v8 = (^(v8))
    var v14 float64 = (0.0 - 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
