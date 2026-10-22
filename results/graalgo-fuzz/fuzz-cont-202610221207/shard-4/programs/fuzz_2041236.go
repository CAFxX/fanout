package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 42
    var v2 string = "hello"
    var v3 uint = 18446744073709551615
    var v4 float64 = 3.14159
    v1 -= (v1 | (-(v1)))
    fmt.Println(v4, v4)
    var v5 string = "a"
    v4 = 0.1
    v6 := []int{32768, 4294967295, 4294967295}
    v4 = (((v4 + v4) + 1e10) - (float64((1e10 + 1e10)) + 1.5))
    if true {
        if true {
            v3 += v3
        }
    } else {
        for fi7 := 0; fi7 < 10; fi7++ {
            v6 = append(v6, (fi7 >> (((^(v3))) & 63)))
            var v8 string = v2
            var v9 uint32 = 2147483648
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v3 = v3
    for fi10 := 0; fi10 < 9; fi10++ {
        if true {
            var v11 int16 = v1
            _ = v11
        }
        var v12 int64 = 42
        _ = fi10
        _ = v12
    }
    v6 = append(v6, 4294967295)
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
