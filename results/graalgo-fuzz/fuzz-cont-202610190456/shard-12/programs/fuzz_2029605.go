package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int64 = 42
    v2 := []string{"", "hello", ("\n\t" + "\n\t")}
    var v3 int32 = 73
    var v4 int = 2147483647
    var v5 uint = 962
    fmt.Println(v3, v1)
    v2 = append(v2, "\n\t")
    v1 = v1
    var v6 float64 = float64(float32(1e10))
    var v7 int32 = -1
    if ((((v1 > v1) && false) && (int16(v4) > -1)) || ((0 >= 44) && false)) {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint8 = 117
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 float32 = -2.25
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    } else {
        var v12 uint64 = 615
        _ = v12
    }
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    if ((v3 & -2147483648) >= v3) {
        v13 := GenS1{F0: 1}
        _ = v13
    }
    _ = len(v2)
    v7 += v3
    var v14 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
