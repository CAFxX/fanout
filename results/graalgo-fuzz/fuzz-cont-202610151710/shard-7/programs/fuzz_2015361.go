package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 255
    var v2 uint8 = v1
    var v3 uint16 = 908
    var v4 uint = 809
    var v5 uint = 4294967295
    var v6 float64 = (-2.25 - 1.5)
    var v7 int = -188
    for fi8 := 0; fi8 < 13; fi8++ {
        v3 = 65535
        var v9 int16 = 42
        _ = fi8
        _ = v9
    }
    var v10 uint64 = 457
    for fi11 := 0; fi11 < 17; fi11++ {
        if false {
            var v12 int16 = 42
            _ = v12
        } else {
            if false {
                switch "" {
                case "\n\t":
                    var v13 uint16 = v3
                    _ = v13
                }
                v10 = (^(v10))
            }
        }
        _ = fi11
    }
    var v14 float32 = ((0.0 * (0.0 / (3.14159 + (((0.0 - 3.14159) + -2.25) / (1e10 * -2.25))))) * (float32(0.0) + (float32(3.14159) + 0.0)))
    fmt.Println(v7)
    v4 += (v4 | v5)
    v5 = (v4 % (((v5 | v5)) | 1))
    var v15 uint32 = 0
    fmt.Println(v4)
    var v16 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
