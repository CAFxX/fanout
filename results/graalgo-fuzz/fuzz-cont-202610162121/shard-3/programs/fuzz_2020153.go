package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }

func main() {
    v1 := []int{42, 160, 9223372036854775807}
    v2 := make(map[uint32]string)
    var v3 int64 = 425
    var v4 float32 = (3.14159 * -2.25)
    v4 += ((((v4 + v4) / (0.0 + v4)) - v4) / -2.25)
    if false {
        for fi5 := 0; fi5 < 9; fi5++ {
            if false {
                v3 -= v3
            }
            v3 = v3
            for fi6 := 0; fi6 < 8; fi6++ {
                fmt.Println(v4, v4)
                _ = fi6
            }
            _ = fi5
        }
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        v3 -= 42
        _ = fi7
    }
    if (3.14159 == 3.14159) {
        var v8 int16 = -1
        fmt.Println(v3, v4)
        var v9 float64 = -2.25
        _ = v8
        _ = v9
    } else {
        if ("" != "foo bar") {
            v10 := v2[65535]
            for fi11 := 0; fi11 < 19; fi11++ {
                for fi12 := 0; fi12 < 20; fi12++ {
                    switch (fi12 + fi12) {
                    case -714:
                        var v13 int = -2147483648
                        v3 -= 1
                        _ = v13
                    case -1:
                        v4 = 1.5
                    case 2147483647:
                        v4 = (v4 * -2.25)
                    }
                    v4 -= v4
                    v14 := GenS1{F0: 71, F1: 1}
                    fmt.Println(v10)
                    _ = fi12
                    _ = v14
                }
                var v15 int8 = -123
                fmt.Println(v4)
                _ = fi11
                _ = v15
            }
            if true {
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 int32 = 2147483647
                    v18 := GenS1{F0: 127, F1: 941}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                switch (v10 + ("foo bar" + v10)) {
                case "\n\t":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    v19 := GenS1{F0: 127, F1: 0}
                    _ = v19
                }
            }
            _ = v10
        }
    }
    var v20 float64 = 1.5
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    v4 += 0.0
    var v21 int = 34
    v22 := GenS1{F0: -128, F1: 2}
    v23 := v2[208]
    var v24 float64 = v20
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
