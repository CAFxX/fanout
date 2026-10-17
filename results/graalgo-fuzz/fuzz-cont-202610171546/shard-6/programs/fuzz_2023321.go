package main

import (
    "fmt"
)


func main() {
    var v1 int = -620
    v2 := make(map[int]bool)
    v3 := []string{"", "\n\t"}
    if true {
        v1 += v1
        v3 = append(v3, "foo bar")
        v1 += 127
    } else {
        if len(v3) > 0 {
            v3[0] = "a"
        }
        var v4 int8 = 0
        _ = v4
    }
    var v5 int8 = -85
    var v6 uint8 = 2
    var v7 uint32 = 228
    switch (v1 * 128) {
    case 32767:
        v6 += 2
    }
    v6 -= 128
    fmt.Println(v1)
    delete(v2, (-(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
