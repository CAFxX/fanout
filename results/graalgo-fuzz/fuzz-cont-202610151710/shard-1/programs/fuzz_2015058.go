package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 864
    var v2 string = ""
    var v3 int8 = 111
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 uint16 = (^(v1))
        v3 = -128
        _ = fi4
        _ = v5
    }
    if true {
        v1 = 318
    }
    fmt.Println(v2, v1, v3)
    v6 := []int{109, 9223372036854775807, 713, 9223372036854775807}
    var v7 uint8 = 255
    if len(v6) > 0 {
        v6[0] = 127
    }
    fmt.Println(v2, v3)
    if true {
        v7 -= v7
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v1, v1)
    var v8 uint8 = 2
    var v9 int64 = 9223372036854775807
    v1 -= (v1 | 0)
    if (false || false) {
        fmt.Println(v7, v3)
        var v10 float64 = (0.1 * -2.25)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v10
    }
    v8 += (v8 - (v8 % ((v7) | 1)))
    if len(v6) > 0 {
        v6[0] = -676
    }
    v6 = append(v6, -1)
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 int64 = (v9 + v9)
        v2 = (v2 + (v2 + "hello"))
        if false {
            var v13 float64 = (1e10 / 1.0)
            for fi14 := 0; fi14 < 9; fi14++ {
                v12 += 42
                v12 += (v9 / ((v12) | 1))
                var v15 float32 = 1.5
                _ = fi14
                _ = v15
            }
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
