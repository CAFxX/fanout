package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 0
    v2 := []int{}
    v3 := make(map[uint32]bool)
    v4 := v3[2147483648]
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        v1 += -128
    }
    var v5 int8 = (v1 - 1)
    if len(v2) > 0 {
        v2[0] = 596
    }
    var v6 uint8 = 127
    v4 = (255 > 1)
    v1 += (-(v1))
    delete(v3, 149)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
