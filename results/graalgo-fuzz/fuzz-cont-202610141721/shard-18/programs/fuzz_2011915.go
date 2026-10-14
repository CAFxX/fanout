package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }

func main() {
    v1 := []int{9223372036854775807, 9223372036854775807, -980}
    var v2 string = "foo bar"
    var v3 float64 = ((float64(-2.25) / 1.5) + (((3.14159 / 0.1) + 1.5) + (((3.14159 - (3.14159 * (3.14159 - -2.25))) / (1e10 + 3.14159)) / (float64((-2.25 - 1e10)) * ((1.5 - 0.1) + 0.1)))))
    var v4 float32 = 1e10
    v5 := make(map[int]string)
    v2 = (("a" + "") + ((v2 + "") + "\n\t"))
    fmt.Println(v2, v3)
    fmt.Println(v3, v2, v3)
    var v6 int16 = 1000
    v6 += -971
    switch 32767 {
    case 731:
        v7 := GenS1{F0: (v6 & (^(v6))), F1: int32(v6)}
        _ = v7
    case -433:
        var v8 int16 = (v6 + v6)
        _ = v8
    case 128:
        v5[0] = "\n\t"
    }
    fmt.Println(v2, v2)
    var v9 bool = false
    var v10 int8 = 127
    v2 = v2
    v6 = v6
    if v9 {
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 uint64 = 293
            _ = fi11
            _ = v12
        }
    } else {
        var v13 uint64 = 323
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
