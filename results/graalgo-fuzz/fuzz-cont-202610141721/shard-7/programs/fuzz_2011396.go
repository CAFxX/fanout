package main

import (
    "fmt"
)


func main() {
    var v1 bool = true
    var v2 int64 = 42
    var v3 string = ("hello" + ("" + "hello"))
    v4 := make(map[string]bool)
    var v5 uint32 = 993
    v6 := []int{42, -318}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        v6 = append(v6, (fi7 * fi7))
        _ = fi7
    }
    v1 = v1
    var v8 uint64 = 4294967295
    for fi9 := 0; fi9 < 1; fi9++ {
        v2 = v2
        _ = fi9
    }
    v2 += v2
    for fi10 := 0; fi10 < 5; fi10++ {
        v6 = append(v6, fi10)
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 float64 = -2.25
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
