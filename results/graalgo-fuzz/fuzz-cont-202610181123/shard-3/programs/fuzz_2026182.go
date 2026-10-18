package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint; F1 int8; F2 int32 }

func main() {
    var v1 uint = 2
    var v2 int8 = 100
    var v3 uint16 = 1
    var v4 int32 = -1
    if (42 < 8) {
        switch 1 {
        case 0:
            v4 = (v4 - (v4 & -805))
        case -404:
            for fi5 := 0; fi5 < 12; fi5++ {
                var v6 uint = v1
                switch ("" + "hello") {
                case "foo bar":
                    var v7 string = "hello"
                    _ = v7
                case "":
                    v8 := []int{}
                    _ = v8
                case "\n\t":
                    if ((174 < 2) && ((false || (!true)) || false)) {
                        for fi9 := 0; fi9 < 1; fi9++ {
                            fmt.Println(fi9, v2)
                            var v10 int32 = (v4 - 1)
                            var v11 int64 = 9223372036854775807
                            var v12 uint16 = v3
                            v12 += (v12 % (((v3 - (^(v12)))) | 1))
                            _ = fi9
                            _ = v10
                            _ = v11
                            _ = v12
                        }
                        v6 -= v1
                        var v13 int32 = (v4 << ((v1) & 63))
                        _ = v13
                    }
                }
                _ = fi5
                _ = v6
            }
        }
        fmt.Println(v3, v1, v3)
    } else {
        v14 := GenS1{F0: -128, F1: v1}
        fmt.Println(v1, v1)
        v4 += (-(v4))
        _ = v14
    }
    var v15 bool = true
    var v16 int64 = 42
    v17 := GenS1{F0: v2, F1: v1}
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 int = fi18
        v3 += (v3 | v3)
        _ = fi18
        _ = v19
    }
    switch -273 {
    case -9223372036854775808:
        var v20 bool = true
        _ = v20
    case 2147483647:
        if (v16 != -880) {
            if false {
                v1 -= v1
                var v21 float64 = ((0.0 * 3.14159) + ((1.5 * 3.14159) + (((1e10 / ((-2.25 * (0.1 - (1.5 / (((float64(3.14159) * (-2.25 / 3.14159)) * 1.5) / 0.1)))) / 1e10)) * 0.0) + -2.25)))
                _ = v21
            } else {
                v22 := []int{-1, 670}
                v23 := GenS2{F0: 314, F1: (-(v2)), F2: 464}
                _ = v22
                _ = v23
            }
        } else {
            v15 = (544 <= -1)
        }
    default:
        switch ("\n\t" + ("hello" + "foo bar")) {
        case "\n\t":
            v1 += (^(v1))
        }
    }
    var v24 uint8 = 0
    v25 := []int{152, -165, -153, 787}
    fmt.Println(v15)
    var v26 uint8 = 2
    for fi27 := 0; fi27 < 10; fi27++ {
        v3 = 42
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
