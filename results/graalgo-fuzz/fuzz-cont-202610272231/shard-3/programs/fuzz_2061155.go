package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]int)
    var v2 string = "hello"
    v3 := make(map[string]int)
    v4 := []uint32{4294967295, 957}
    v5 := []string{((v2 + v2) + v2), (v2 + (v2 + v2)), ((("hello" + ("foo bar" + v2)) + "hello") + (v2 + ((v2 + ("foo bar" + v2)) + v2))), v2}
    fmt.Println(v2)
    fmt.Println(v2)
    var v6 bool = true
    var v7 float64 = (3.14159 / 1.0)
    v4 = append(v4, 441)
    v6 = ((3.14159 - (((3.14159 * 1e10) + 1.5) / 1e10)) != 0.0)
    if (v6 || v6) {
        var v8 int32 = 1
        fmt.Println(v7, v2)
        v7 = 1e10
        _ = v8
    } else {
        for fi9 := 0; fi9 < 17; fi9++ {
            for fi10 := 0; fi10 < 6; fi10++ {
                if true {
                    var v11 uint64 = 418
                    _ = len(v3)
                    v11 = v11
                    _ = v11
                }
                _ = fi10
            }
            v7 -= (-2.25 + v7)
            _ = fi9
        }
        var v12 bool = v6
        if len(v5) > 0 {
            v5[0] = v2
        }
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        v6 = (v6 || true)
        _ = fi13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
