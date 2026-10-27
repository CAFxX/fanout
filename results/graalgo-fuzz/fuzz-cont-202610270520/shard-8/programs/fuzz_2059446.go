package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 3.14159
    v3 := make(map[string]int)
    var v4 int16 = 42
    var v5 int = 127
    v6 := make(map[string]uint32)
    v2 = v2
    var v7 int = (v5 % ((v5) | 1))
    v6["\n\t"] = 2
    _ = len(v6)
    if (false && ((1 < 165) && false)) {
        var v8 bool = false
        _ = v8
    } else {
        v1[136] = (v7 >> ((255) & 63))
        switch "\n\t" {
        case "\n\t":
            v2 += v2
        default:
            var v9 float64 = 0.0
            _ = v9
        }
    }
    var v10 float32 = 3.14159
    v11 := GenS1{F0: (1 >= 255)}
    if true {
        v4 = (^(v4))
        v12 := GenS1{F0: true}
        if (v2 < float32(((((((0.0 / 1.0) / 0.1) - ((0.1 * 3.14159) + (3.14159 + (1.5 + 0.1)))) * (3.14159 * 0.1)) / (((0.1 / 0.1) + (((((0.0 * 1.5) / -2.25) / ((((1.5 - -2.25) / ((float64(v10) * (0.0 + -2.25)) * 0.0)) - 3.14159) + ((3.14159 * 3.14159) + (((-2.25 + 3.14159) * 1.5) * ((3.14159 + 0.0) * 3.14159))))) / 1.0) * (3.14159 - ((0.1 / 1.0) / 0.1)))) * -2.25)) - 3.14159))) {
            v10 = 1.5
            v13 := []int{(-(v5)), 0, 32768}
            _ = v13
        }
        _ = v12
    } else {
        if (false || false) {
            switch v5 {
            case 127:
                if (84 < 2) {
                    var v14 string = ("a" + "a")
                    if (v5 != v5) {
                        var v15 int8 = 100
                        v16 := GenS1{F0: ((v14 > (("foo bar" + (v14 + v14)) + "a")) && false)}
                        _ = v15
                        _ = v16
                    }
                    _ = v14
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
