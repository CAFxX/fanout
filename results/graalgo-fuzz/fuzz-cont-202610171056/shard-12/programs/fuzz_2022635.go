package main

import (
    "fmt"
)


func main() {
    v1 := []string{}
    v2 := make(map[string]uint32)
    var v3 uint8 = 0
    var v4 uint = 255
    var v5 int16 = 194
    var v6 uint32 = 717
    v3 = v3
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            v6 = (v6 * (v6 % ((v6) | 1)))
            v2["foo bar"] = (v6 | (v6 * (v6 / (((v6 % ((v6) | 1))) | 1))))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi8
        }
        var v9 float32 = -2.25
        _ = fi7
        _ = v9
    }
    switch 32768 {
    case 42:
        fmt.Println(v4, v6, v5)
    case 957:
        switch 4294967295 {
        case -792:
            v3 -= 128
        case 34:
            var v10 string = ""
            _ = v10
        case 0:
            fmt.Println(v3)
        }
    case -740:
        if true {
            v4 -= v4
        } else {
            v1 = append(v1, "\n\t")
        }
    }
    if false {
        fmt.Println(v3, v4)
        fmt.Println(v5, v4)
        v1 = append(v1, "hello")
    }
    var v11 uint = (v4 - (v4 * 451))
    var v12 int32 = 239
    switch 4294967295 {
    case 2147483647:
        var v13 float32 = 3.14159
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        if (447 < 4294967295) {
            var v15 bool = (false && false)
            _ = v15
        } else {
            v5 = (v5 % ((v5) | 1))
            var v16 float64 = ((float64(0.1) * (((1.5 / 1.0) * (((3.14159 * (1e10 + 0.0)) * 0.1) - 3.14159)) + (-2.25 * 3.14159))) + 0.0)
            _ = v16
        }
        v12 = v12
        _ = fi14
    }
    fmt.Println(v12, v4)
    v12 += -1
    if false {
        var v17 int = -9223372036854775808
        _ = v17
    } else {
        var v18 uint64 = 4294967295
        _ = v18
    }
    if false {
        var v19 float32 = ((1e10 * (1.5 / 3.14159)) + -2.25)
        _ = v19
    }
    v4 -= (^(v4))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
