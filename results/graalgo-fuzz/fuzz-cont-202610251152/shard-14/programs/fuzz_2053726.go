package main

import (
    "fmt"
)


func main() {
    var v1 string = "a"
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    v4 := make(map[int]bool)
    var v5 float64 = ((-2.25 + 1.5) + 1.5)
    v6 := []int{32768}
    switch ("hello" + v1) {
    case "\n\t":
        v1 = "\n\t"
    case "a":
        var v7 uint16 = 2
        _ = v7
    case "foo bar":
        if len(v6) > 0 {
            v6[0] = -80
        }
    }
    _ = len(v3)
    v8 := v2[2147483648]
    var v9 int16 = 0
    var v10 string = ((v1 + v1) + ("hello" + v1))
    for fi11 := 0; fi11 < 10; fi11++ {
        v5 += v5
        var v12 int64 = -9223372036854775808
        _ = fi11
        _ = v12
    }
    var v13 bool = (((127 == -128) || true) || true)
    v13 = true
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
