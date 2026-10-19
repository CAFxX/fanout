package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 933
    v2 := []string{(("foo bar" + "\n\t") + "foo bar"), (("\n\t" + "a") + ("hello" + ("\n\t" + (("\n\t" + "") + "hello"))))}
    v3 := []uint32{653}
    v4 := make(map[int]int)
    v1 = v1
    v1 = (v1 - v1)
    var v5 string = ("\n\t" + ("\n\t" + ""))
    v5 = v5
    var v6 int8 = 100
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 float32 = 1e10
        switch fi7 {
        case -1:
            v6 += v6
        case -9223372036854775808:
            v5 = v5
        }
        _ = fi7
        _ = v8
    }
    v5 = (v5 + v5)
    var v9 int = 127
    var v10 float32 = (0.1 / (1e10 * ((1.5 + (1e10 / 0.1)) - 1.5)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
