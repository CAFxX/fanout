package main

import (
    "fmt"
)


func main() {
    v1 := []int{-805}
    var v2 uint = 675
    v3 := []string{("hello" + "")}
    var v4 int64 = -9223372036854775808
    var v5 uint8 = 0
    v2 = v2
    var v6 uint = v2
    v5 = v5
    fmt.Println(v6, v2)
    for fi7 := 0; fi7 < 1; fi7++ {
        if false {
            var v8 uint32 = 623
            fmt.Println(fi7, v2)
            _ = v8
        }
        _ = fi7
    }
    var v9 bool = (float32((1.5 * 0.1)) == 1e10)
    v2 = v2
    if false {
        var v10 string = ("foo bar" + (((("\n\t" + "a") + "foo bar") + ("hello" + "hello")) + ("\n\t" + "\n\t")))
        for fi11 := 0; fi11 < 9; fi11++ {
            fmt.Println(v6, v9)
            var v12 uint64 = 638
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
