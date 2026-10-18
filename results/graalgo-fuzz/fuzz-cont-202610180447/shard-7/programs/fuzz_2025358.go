package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]int)
    var v3 int16 = 42
    var v4 float64 = (3.14159 * 1.5)
    var v5 int16 = v3
    v5 += v3
    switch "\n\t" {
    case "hello":
        switch 2147483647 {
        case 2147483647:
            for fi6 := 0; fi6 < 5; fi6++ {
                var v7 uint = 28
                _ = fi6
                _ = v7
            }
        }
    case "\n\t":
        v8 := GenS1{F0: 1}
        _ = v8
    case "foo bar":
        v5 += v3
    }
    v4 += -2.25
    var v9 uint64 = 443
    v10 := GenS1{F0: 191}
    var v11 bool = false
    var v12 bool = true
    var v13 bool = (1 <= 582)
    v14 := []int{}
    fmt.Println(v12, v12)
    var v15 bool = (((((-2.25 + (((1e10 + (1.5 * 0.1)) + 1e10) / (0.1 * 1e10))) - 1.5) + 1.5) - (((1.5 * 0.1) * (0.0 * 1.5)) / 1e10)) <= ((-2.25 * -2.25) + 3.14159))
    for fi16 := 0; fi16 < 13; fi16++ {
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := GenS1{F0: 1}
            _ = fi17
            _ = v18
        }
        v4 = v4
        v13 = v13
        _ = fi16
    }
    var v19 uint = 2
    v20 := GenS1{F0: v19}
    var v21 int32 = -2147483648
    fmt.Println(v9, v9)
    if v12 {
        switch "a" {
        case "\n\t":
            fmt.Println(v5, v15, v4)
        case "":
            var v22 int = -2147483648
            _ = v22
        case "hello":
            if (65535 == 255) {
                v23 := v2[""]
                var v24 int32 = (v21 + (v21 + v21))
                _ = v23
                _ = v24
            }
        default:
            var v25 bool = ((100 != 68) || (9223372036854775807 > 633))
            _ = v25
        }
        switch "" {
        case "":
            switch 4294967295 {
            case 128:
                var v26 uint64 = (v9 * v9)
                _ = v26
            }
        }
    } else {
        var v27 uint64 = 687
        var v28 string = ("foo bar" + "hello")
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
