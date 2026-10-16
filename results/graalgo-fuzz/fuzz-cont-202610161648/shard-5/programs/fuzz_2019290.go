package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int8 = 1
    var v2 int = 9223372036854775807
    v3 := make(map[uint32]string)
    var v4 float64 = -2.25
    v5 := []int{(^(v2)), (v2 & v2)}
    var v6 bool = true
    v7 := GenS1{F0: 42}
    switch ("\n\t" + "a") {
    case "hello":
        var v8 string = ("foo bar" + "hello")
        _ = v8
    case "foo bar":
        if v6 {
            var v9 bool = (v4 == (((v4 - 1.5) / -2.25) - 1.5))
            v4 += v4
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 4; fi10++ {
            v1 = (v1 & v1)
            _ = fi10
        }
    }
    v6 = v6
    v1 += v1
    v2 += -814
    v1 += (^(v1))
    _ = len(v3)
    var v11 float32 = 0.1
    v2 = v2
    v12 := v3[546]
    v6 = v6
    fmt.Println(v2, v6)
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
