package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    var v4 uint32 = 2147483648
    v1 -= (v1 ^ 28)
    if (42 >= 4294967295) {
        if (960 >= 18446744073709551615) {
            switch "foo bar" {
            case "foo bar":
                delete(v2, ("" + "\n\t"))
            default:
                v5 := v2[("a" + "foo bar")]
                _ = v5
            }
        }
    }
    fmt.Println(v4, v4)
    var v6 uint16 = 0
    v1 -= 1
    var v7 int8 = -31
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint16 = 342
            var v11 int64 = 125
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    var v12 int = 32767
    var v13 float32 = 0.1
    v14 := GenS1{F0: 9223372036854775807}
    v15 := []int{320, v12, v12}
    v7 = (v7 - (v7 % ((v7) | 1)))
    v13 += (1.5 - ((0.1 + v13) / (v13 * (v13 - (-2.25 - v13)))))
    var v16 float64 = 0.0
    v13 = v13
    if true {
        var v17 int16 = -32768
        switch ("foo bar" + "\n\t") {
        case "hello":
            if (!false) {
                delete(v2, ("foo bar" + "hello"))
            }
        }
        _ = v17
    } else {
        if false {
            var v18 string = "hello"
            _ = v18
        } else {
            fmt.Println(v1, v7, v12)
        }
        var v19 uint = 0
        var v20 float32 = ((1.5 + (((v13 * v13) / 1.5) - ((3.14159 + ((v13 - v13) / v13)) / -2.25))) - 1e10)
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
