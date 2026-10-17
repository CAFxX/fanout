package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 1000
    var v2 uint = 18446744073709551615
    var v3 uint16 = 65535
    var v4 uint = v2
    var v5 int32 = -656
    var v6 uint64 = 18446744073709551615
    v6 = v6
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 uint8 = 255
            _ = fi8
            _ = v9
        }
        var v10 float32 = (-2.25 / -2.25)
        _ = fi7
        _ = v10
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := []int{fi11, fi11}
        v2 -= v2
        _ = fi11
        _ = v12
    }
    fmt.Println(v2)
    var v13 float32 = 1e10
    var v14 int64 = 1
    var v15 int64 = 9223372036854775807
    v2 = (^(v4))
    fmt.Println(v14, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
