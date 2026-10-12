package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 string }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 string = "foo bar"
    var v2 uint16 = 32767
    v3 := make(map[int]string)
    var v4 int32 = 863
    v5 := make(map[string]int)
    var v6 uint16 = 425
    for fi7 := 0; fi7 < 10; fi7++ {
        v4 -= (v4 / ((v4) | 1))
        _ = fi7
    }
    v8 := GenS1{F0: -468, F1: 9223372036854775807, F2: "hello"}
    fmt.Println(v2, v4)
    v2 -= v6
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := []int{fi9}
            v1 = "a"
            _ = fi10
            _ = v11
        }
        var v12 float64 = (((float64(float32(1e10)) - (0.0 + 3.14159)) + (float64(3.14159) + 1e10)) / (0.1 - 1.5))
        _ = fi9
        _ = v12
    }
    v13 := GenS2{F0: v2}
    v3[0] = (v1 + "foo bar")
    v1 = v1
    v1 = v1
    switch ("\n\t" + ("" + "hello")) {
    case "\n\t":
        v14 := GenS1{F0: 2147483647, F1: 902, F2: ((((v1 + v1) + (v1 + "")) + v1) + v1)}
        _ = v14
    case "a":
        v4 = -933
    }
    v4 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
