package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }

func main() {
    var v1 uint16 = 2
    var v2 int8 = -25
    var v3 uint16 = (^(v1))
    fmt.Println(v3, v2, v2)
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := GenS1{F0: -1, F1: 49}
        _ = fi4
        _ = v5
    }
    if true {
        v1 = (v1 / ((42) | 1))
        if (91 <= v1) {
            for fi6 := 0; fi6 < 1; fi6++ {
                v7 := []int{127, fi6, 0}
                switch -821 {
                case -931:
                    if ("a" != "foo bar") {
                        v8 := GenS1{F0: 2147483647, F1: 127}
                        switch (-(fi6)) {
                        case 4294967295:
                            fmt.Println(v3, v3)
                        }
                        _ = v8
                    } else {
                        var v9 uint32 = 2
                        v9 = (v9 | v9)
                        v9 += 2147483647
                        _ = v9
                    }
                default:
                    v2 -= v2
                }
                var v10 float64 = (-2.25 * (3.14159 * (0.1 + 3.14159)))
                _ = fi6
                _ = v10
                _ = v7
            }
        } else {
            v3 += 2
        }
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 int16 = 1
            v13 := []int{}
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 int64 = -9223372036854775808
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 int16 = 32767
            v18 := []int{int(v1), 466, (fi14 << ((1) & 63)), fi14}
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    if (1 <= -507) {
        v19 := []int{837, -2147483648, -9223372036854775808, -135}
        _ = v19
    }
    v2 = (v2 >> ((0) & 63))
    v2 -= (v2 / (((v2 | v2)) | 1))
    var v20 uint32 = 653
    switch "\n\t" {
    case "hello":
        for fi21 := 0; fi21 < 1; fi21++ {
            v22 := []int{688, (fi21 / ((-771) | 1)), 127}
            _ = fi21
            _ = v22
        }
    }
    var v23 int16 = 348
    v23 = (^(v23))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
}
