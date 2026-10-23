package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint8 = 1
    var v2 uint8 = (^(v1))
    var v3 string = "foo bar"
    var v4 uint = 121
    v5 := []int{2147483647}
    var v6 int32 = 0
    var v7 float64 = (1e10 + 0.1)
    v8 := GenS1{F0: v6}
    var v9 string = v3
    v5 = append(v5, 4294967295)
    var v10 float32 = ((0.1 * ((1e10 / float32(v7)) - (3.14159 / ((-2.25 - (1.5 / float32(v7))) * (0.0 - (0.0 * (-2.25 * 0.1))))))) + (1.5 + 3.14159))
    for fi11 := 0; fi11 < 9; fi11++ {
        v2 -= (v1 * v2)
        switch ((v9 + "\n\t") + ("hello" + v3)) {
        case "\n\t":
            if (!false) {
                v6 = 42
                var v12 int = fi11
                v13 := GenS1{F0: 42}
                _ = v12
                _ = v13
            } else {
                v6 -= v6
                v6 = v6
                var v14 int32 = (v6 % ((619) | 1))
                _ = v14
            }
        case "":
            switch 127 {
            case 4294967295:
                if len(v5) > 0 {
                    v5[0] = fi11
                }
            case 595:
                var v15 string = "\n\t"
                _ = v15
            }
        case "a":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = fi11
    }
    var v16 int32 = v6
    var v17 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
