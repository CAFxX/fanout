package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 string }

func main() {
    var v1 string = "\n\t"
    v2 := []int{670}
    var v3 float64 = float64((0.0 * (((1e10 * 0.0) - 3.14159) + 0.1)))
    var v4 bool = true
    var v5 bool = ((0.0 / 3.14159) != -2.25)
    var v6 float32 = 3.14159
    v7 := GenS1{F0: 69}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!(!(1 != 32767))) {
        v8 := GenS1{F0: 99}
        for fi9 := 0; fi9 < 9; fi9++ {
            v6 = v6
            if (!false) {
                v5 = false
                v5 = ((v1 + ((("a" + "\n\t") + v1) + "\n\t")) != "")
                v1 = "hello"
            }
            for fi10 := 0; fi10 < 15; fi10++ {
                fmt.Println(v6)
                _ = fi10
            }
            _ = fi9
        }
        v3 = ((-2.25 + 1e10) - ((-2.25 + v3) - v3))
        _ = v8
    } else {
        var v11 uint = 377
        _ = v11
    }
    v12 := GenS2{F0: 462, F1: 2, F2: "foo bar"}
    var v13 string = v1
    _ = len(v2)
    var v14 string = v13
    v2 = append(v2, 32767)
    if len(v2) > 0 {
        v2[0] = 307
    }
    fmt.Println(v13, v5, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
