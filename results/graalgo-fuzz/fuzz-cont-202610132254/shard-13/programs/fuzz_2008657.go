package main

import (
    "fmt"
)


func main() {
    v1 := []int{}
    v2 := []string{"\n\t", "hello"}
    v3 := make(map[string]string)
    var v4 int8 = -105
    v4 += -1
    v4 = (v4 ^ v4)
    v5 := v3[("\n\t" + "")]
    v4 += (v4 << ((0) & 63))
    v5 = (("" + "hello") + v5)
    fmt.Println(v5, v4)
    v5 = "foo bar"
    var v6 uint32 = 2
    v6 = (v6 | 4294967295)
    _ = len(v2)
    v6 += (^(v6))
    fmt.Println("slicelen v1:", len(v1), v1)
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
