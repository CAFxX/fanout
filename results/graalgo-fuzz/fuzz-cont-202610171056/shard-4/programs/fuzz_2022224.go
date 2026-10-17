package main

import (
    "fmt"
)


func main() {
    var v1 bool = (65535 < 973)
    var v2 string = "\n\t"
    var v3 bool = (!(2147483648 >= 4294967295))
    var v4 uint8 = 212
    v5 := []string{}
    var v6 float64 = float64(((1e10 * 3.14159) - 0.0))
    v4 += (^(v4))
    fmt.Println(v2, v3)
    var v7 int16 = 1000
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        v4 = v4
        _ = fi8
    }
    fmt.Println(v2, v7, v3)
    var v9 string = v2
    var v10 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
