package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 53
    var v2 float64 = 0.1
    var v3 bool = true
    var v4 bool = ("\n\t" >= "hello")
    if (18446744073709551615 > 0) {
        var v5 uint8 = 2
        _ = v5
    }
    var v6 int64 = 0
    var v7 uint8 = 42
    v4 = (270 < -1)
    v7 = v7
    var v8 uint = 255
    for fi9 := 0; fi9 < 17; fi9++ {
        v2 += 3.14159
        v6 += (-(v6))
        _ = fi9
    }
    v1 -= (v1 * (-(v1)))
    var v10 int32 = -2147483648
    var v11 uint64 = 2
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
