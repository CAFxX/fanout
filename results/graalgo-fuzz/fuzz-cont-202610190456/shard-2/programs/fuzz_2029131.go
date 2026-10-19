package main

import (
    "fmt"
)


func main() {
    var v1 string = "\n\t"
    var v2 float64 = float64((0.0 / 1.0))
    var v3 string = ("\n\t" + ("hello" + "a"))
    v4 := []int{32767}
    var v5 uint8 = 255
    v6 := []string{}
    if false {
        var v7 uint16 = 263
        _ = v7
    } else {
        v1 = (v3 + v1)
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 float32 = 0.0
            v6 = append(v6, v1)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi8
            _ = v9
        }
    }
    var v10 int64 = int64(v5)
    var v11 int64 = v10
    v5 = v5
    var v12 float32 = (-2.25 / 1e10)
    v4 = append(v4, -1)
    var v13 string = v3
    fmt.Println(v2)
    v5 = (v5 / ((196) | 1))
    v10 += (v10 + -9223372036854775808)
    if (100 == -1) {
        fmt.Println(v5)
    } else {
        v6 = append(v6, (v1 + (v3 + v3)))
    }
    v3 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
