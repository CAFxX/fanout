package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 float64 = (0.0 - (3.14159 / 1e10))
    var v2 string = "hello"
    var v3 string = "foo bar"
    v4 := []int{32767, -9223372036854775808}
    var v5 float64 = ((v1 / (v1 + v1)) * (v1 - v1))
    v6 := GenS1{F0: 170}
    v3 = v2
    var v7 int16 = 781
    var v8 string = v3
    var v9 float32 = 3.14159
    fmt.Println(v5, v2)
    var v10 float64 = (((v5 + v5) * ((((v5 - 1.5) * 1.5) - v1) + (v1 - (v1 - v5)))) + (((v5 + 0.1) * ((0.1 + v5) - 1.5)) - 0.1))
    fmt.Println(v3, v2, v8)
    v7 = v7
    var v11 float64 = v10
    switch 9223372036854775807 {
    case -620:
        fmt.Println(v5, v1, v11)
    default:
        for fi12 := 0; fi12 < 11; fi12++ {
            fmt.Println(v3)
            var v13 string = ("\n\t" + "foo bar")
            _ = fi12
            _ = v13
        }
    }
    var v14 uint64 = 4294967295
    v11 -= (v1 + v10)
    for fi15 := 0; fi15 < 5; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            v9 = (v9 / v9)
            fmt.Println(v2)
            v9 += (((1e10 * (v9 + v9)) * ((v9 * (1.5 + (-2.25 * (v9 - v9)))) + 0.1)) - ((v9 - (float32((v10 * (float64(-2.25) - 1.5))) + v9)) * -2.25))
            v8 = v8
            var v17 float32 = float32(0.0)
            var v18 uint64 = (v14 & 2)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
