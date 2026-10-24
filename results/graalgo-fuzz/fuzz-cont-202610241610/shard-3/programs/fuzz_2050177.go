package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint32 = 275
    var v2 uint32 = v1
    var v3 int16 = 1000
    var v4 int64 = 982
    var v5 uint8 = 128
    var v6 int = 127
    switch "hello" {
    case "":
        switch (^(v6)) {
        case -9223372036854775808:
            v7 := []int{}
            _ = v7
        case 42:
            v8 := GenS1{F0: 38, F1: 100}
            _ = v8
        case 32768:
            v2 = v1
        default:
            v1 = uint32(v5)
        }
    case "a":
        var v9 uint = 670
        _ = v9
    case "foo bar":
        v10 := []int{v6, (v6 / ((-2147483648) | 1)), 2147483647}
        _ = v10
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := GenS1{F0: -57, F1: -128}
        _ = fi11
        _ = v12
    }
    v13 := []int{4294967295, (v6 & v6), v6}
    v4 += v4
    if ((0.1 + (0.0 - (3.14159 - -2.25))) > 1e10) {
        v13 = append(v13, (v6 >> ((405) & 63)))
    } else {
        var v14 uint64 = 497
        for fi15 := 0; fi15 < 7; fi15++ {
            if (255 > 42) {
                var v16 float32 = (0.0 + ((((((0.1 - 1e10) * float32(float64(1.5))) * 3.14159) / 0.1) / (-2.25 + 0.0)) - 3.14159))
                if false {
                    var v17 string = ("hello" + "a")
                    _ = v17
                }
                _ = v16
            }
            _ = fi15
        }
        v5 += v5
        _ = v14
    }
    v2 -= v1
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 string = "foo bar"
        fmt.Println(v2, v2)
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: 194}
    switch "foo bar" {
    case "a":
        var v21 uint32 = 195
        _ = v21
    case "hello":
        v4 = (v4 >> ((2) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
