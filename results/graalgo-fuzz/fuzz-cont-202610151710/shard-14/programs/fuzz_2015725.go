package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint8 = 255
    var v2 uint64 = 0
    v3 := make(map[string]string)
    var v4 uint8 = (v1 >> ((1) & 63))
    v1 -= (^(v1))
    v5 := v3["foo bar"]
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 int16 = 201
        for fi8 := 0; fi8 < 1; fi8++ {
            if false {
                if false {
                    v3["a"] = v5
                    v3["\n\t"] = v5
                    var v9 float32 = 0.1
                    _ = v9
                }
            } else {
                var v10 float32 = ((0.1 * (0.0 * ((3.14159 - ((-2.25 + (3.14159 - float32((((0.1 / float64((3.14159 - 0.0))) / (3.14159 - float64(1e10))) * 1e10)))) + (((0.0 * 1e10) * 1.5) / -2.25))) / (float32(3.14159) - -2.25)))) / 1.5)
                _ = v10
            }
            v11 := GenS1{F0: 24}
            _ = fi8
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    v4 = (v1 / (((v4 | (v4 + 127))) | 1))
    v12 := []int{}
    fmt.Println(v5, v2)
    if false {
        switch "" {
        case "\n\t":
            var v13 uint = 255
            _ = v13
        }
        v4 -= v1
    }
    var v14 float64 = float64(0.1)
    v15 := GenS1{F0: 100}
    v5 = (v5 + ("hello" + v5))
    var v16 bool = (((0.0 * 1e10) * ((((((-2.25 * (1e10 * 3.14159)) * -2.25) + (3.14159 * 0.1)) - -2.25) + 1e10) / (((1.5 * -2.25) * 3.14159) * 1.5))) == ((1e10 + (1.5 - 1e10)) / (((-2.25 - float32(0.0)) / 1.5) / 1.0)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
