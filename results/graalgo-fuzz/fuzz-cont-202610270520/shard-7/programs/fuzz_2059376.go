package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 float64 = (1e10 + float64(3.14159))
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    v1 += 0.0
    var v4 uint32 = 117
    if false {
        for fi5 := 0; fi5 < 18; fi5++ {
            _ = len(v3)
            var v6 uint64 = 134
            _ = fi5
            _ = v6
        }
    } else {
        if (4294967295 < 99) {
            switch 127 {
            case 128:
                v2[""] = ((2 < 0) || (((v1 * ((v1 * v1) / 1.0)) - 3.14159) == 0.0))
            default:
                _ = len(v3)
            }
        }
        v4 -= (v4 + (^(v4)))
    }
    if false {
        for fi7 := 0; fi7 < 17; fi7++ {
            v1 = v1
            _ = fi7
        }
    }
    var v8 float64 = (3.14159 * v1)
    v8 = (((((1.5 + (float64(0.0) / (v1 + 0.1))) * (v8 - v8)) / v1) - 0.0) - 1.5)
    v4 = (v4 / (((v4 / (((v4 - (v4 / ((v4) | 1)))) | 1))) | 1))
    var v9 uint32 = 42
    switch "\n\t" {
    case "hello":
        v10 := GenS1{F0: -9223372036854775808}
        _ = v10
    case "foo bar":
        if (((v8 >= (float64(0.1) - ((v8 * (v8 + v1)) * v8))) || false) && (!true)) {
            var v11 uint8 = 128
            _ = v11
        }
    }
    switch 42 {
    case -1:
        var v12 int32 = -47
        _ = v12
    }
    if (1000 == -390) {
        var v13 uint8 = 0
        _ = v13
    } else {
        v14 := []int{-275}
        _ = len(v14)
        _ = len(v14)
        _ = v14
    }
    switch 32767 {
    case 2147483647:
        switch (("hello" + "hello") + "\n\t") {
        case "hello":
            var v15 int16 = 1000
            _ = v15
        case "a":
            v9 = 2147483648
        }
    case -117:
        v4 = (v9 % ((v9) | 1))
    case -9223372036854775808:
        v8 = 0.0
    default:
        v1 -= 0.1
    }
    v4 = v4
    v4 += 108
    for fi16 := 0; fi16 < 7; fi16++ {
        for fi17 := 0; fi17 < 3; fi17++ {
            v18 := GenS1{F0: 0}
            fmt.Println(fi17, fi17)
            v19 := GenS1{F0: 9223372036854775807}
            _ = fi17
            _ = v18
            _ = v19
        }
        v9 -= v4
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
