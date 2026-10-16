package main

import (
    "fmt"
)


func main() {
    v1 := []int{9223372036854775807}
    var v2 uint = 255
    v3 := []string{}
    var v4 int16 = -32768
    var v5 float64 = 3.14159
    fmt.Println(v5, v5, v2)
    var v6 int32 = -2147483648
    fmt.Println(v6, v2, v5)
    fmt.Println(v2)
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint8 = 92
        var v9 int16 = v4
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            v5 += (v5 + (float64((((0.0 * -2.25) * 0.0) + 1.5)) / v5))
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
            var v12 float32 = 1.5
            _ = fi11
            _ = v12
        }
        if true {
            fmt.Println(v2, v6)
        }
        _ = fi10
    }
    _ = len(v1)
    var v13 bool = false
    if (0 == 4294967295) {
        v3 = append(v3, "a")
    }
    var v14 uint64 = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
