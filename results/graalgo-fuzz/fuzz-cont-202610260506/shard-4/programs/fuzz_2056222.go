package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 664
    var v2 string = "\n\t"
    var v3 int16 = 1
    v4 := []int{}
    v1 += v1
    fmt.Println(v1, v3)
    var v5 float64 = 3.14159
    var v6 string = v2
    for fi7 := 0; fi7 < 7; fi7++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 float64 = v5
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 int32 = 2147483647
    for fi11 := 0; fi11 < 12; fi11++ {
        if len(v4) > 0 {
            v4[0] = fi11
        }
        _ = fi11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
