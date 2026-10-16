package main

import (
    "fmt"
)


func main() {
    var v1 int32 = 741
    var v2 int8 = 1
    var v3 int64 = 42
    var v4 float64 = 1e10
    var v5 uint8 = 128
    v6 := make(map[string]bool)
    switch ("foo bar" + ("\n\t" + "hello")) {
    case "\n\t":
        fmt.Println(v3)
    case "foo bar":
        if ((v1 | (-(v1))) == (v1 + -2147483648)) {
            var v7 uint16 = 12
            _ = v7
        } else {
            var v8 uint = 0
            _ = v8
        }
    default:
        var v9 int32 = v1
        _ = v9
    }
    delete(v6, "")
    var v10 uint32 = 2
    v10 -= (v10 ^ 2147483647)
    for fi11 := 0; fi11 < 15; fi11++ {
        v4 += v4
        var v12 string = "\n\t"
        v13 := []int{fi11}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int16 = 1
    v6["a"] = false
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
