package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int64 = -928
    var v2 float32 = (0.0 - float32(((0.1 + 1.5) / 0.1)))
    var v3 int = 54
    var v4 string = "a"
    v5 := []int{4294967295, v3}
    var v6 uint16 = 32767
    v4 = v4
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (0 != 4294967295) {
        for fi7 := 0; fi7 < 9; fi7++ {
            v2 = v2
            v8 := GenS1{F0: (-1 < 852)}
            _ = fi7
            _ = v8
        }
    }
    var v9 int8 = -79
    for fi10 := 0; fi10 < 13; fi10++ {
        if false {
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 uint8 = 127
                _ = fi11
                _ = v12
            }
        }
        v9 += 57
        v1 = 1
        _ = fi10
    }
    var v13 float64 = float64(0.0)
    fmt.Println(v1, v2)
    var v14 int64 = v1
    v13 = (0.0 * float64((float32(v13) * float32(0.1))))
    for fi15 := 0; fi15 < 3; fi15++ {
        v13 = ((1.5 + (1.5 - (v13 - (float64(0.0) / 1e10)))) / (float64(v2) + ((v13 - ((0.0 * ((((v13 - v13) / 3.14159) * v13) / v13)) / (float64(((v2 - 1.5) - (v2 / (((v2 * ((1e10 / 0.1) / v2)) * v2) / v2)))) - v13))) * ((v13 * v13) * float64(v2)))))
        var v16 float64 = (1.5 / 1e10)
        v5 = append(v5, (fi15 % ((32768) | 1)))
        _ = fi15
        _ = v16
    }
    var v17 float32 = (v2 + v2)
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
