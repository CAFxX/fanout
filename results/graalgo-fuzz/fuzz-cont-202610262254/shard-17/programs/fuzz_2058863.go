package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 755
    v3 := make(map[int]string)
    v4 := []string{"", "a", "foo bar", ("a" + "hello")}
    var v5 float64 = 3.14159
    v6 := make(map[string]uint32)
    for fi7 := 0; fi7 < 16; fi7++ {
        v2 = (v2 & v2)
        _ = fi7
    }
    var v8 string = "hello"
    v5 -= -2.25
    v8 = v8
    v2 += 42
    switch -9223372036854775808 {
    case -2147483648:
        v5 += (v5 - (0.0 * v5))
    case 42:
        if true {
            var v9 bool = false
            fmt.Println(v2)
            _ = v9
        }
    }
    v10 := v1[32768]
    var v11 int8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
