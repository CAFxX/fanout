package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{602, 1, 2147483647}
    v3 := []uint32{633, 0, 1, 0}
    var v4 bool = (!false)
    var v5 uint32 = 1
    var v6 uint = 1
    for fi7 := 0; fi7 < 8; fi7++ {
        if (v4 || (-1 >= 1)) {
            v8 := v1["hello"]
            v4 = false
            _ = v8
        } else {
            v6 = (v6 / (((v6 << ((0) & 63))) | 1))
            var v9 uint8 = 127
            _ = v9
        }
        v6 = v6
        _ = fi7
    }
    v5 = 731
    var v10 string = ""
    v11 := GenS1{F0: v6, F1: 341}
    fmt.Println(v6)
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: v6, F1: -2147483648}
        v5 += (v5 >> (((v6 / ((v6) | 1))) & 63))
        _ = fi12
        _ = v13
    }
    var v14 float32 = (0.0 / (float32((0.0 / 1e10)) - -2.25))
    v15 := v1[(v10 + ("\n\t" + "foo bar"))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
