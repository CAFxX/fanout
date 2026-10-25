package main

import (
    "fmt"
)


func main() {
    var v1 uint64 = 749
    var v2 int16 = 1000
    var v3 int = 223
    v4 := make(map[int]string)
    var v5 bool = (((281 < 0) || (float32(1e10) <= 0.0)) && (!true))
    _ = len(v4)
    var v6 int16 = (v2 + v2)
    var v7 int16 = (v2 >> ((966) & 63))
    fmt.Println(v2, v2)
    var v8 uint32 = 2147483648
    fmt.Println(v2, v2, v8)
    if ("\n\t" < "foo bar") {
        for fi9 := 0; fi9 < 14; fi9++ {
            v5 = v5
            var v10 uint16 = 105
            _ = fi9
            _ = v10
        }
    }
    v3 = -9223372036854775808
    fmt.Println(v2, v1)
    for fi11 := 0; fi11 < 14; fi11++ {
        v2 += v6
        switch "\n\t" {
        case "hello":
            if ((-2.25 / 1e10) >= (1e10 * 1.5)) {
                v12 := []int{fi11}
                _ = v12
            }
        case "":
            fmt.Println(v2, v8)
        case "\n\t":
            fmt.Println(fi11, v7)
        }
        _ = fi11
    }
    v4[(v3 / ((v3) | 1))] = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
