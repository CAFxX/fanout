package main

import (
    "fmt"
)

func genFn0() bool {
    for fi4 := 0; fi4 < 19; fi4++ {
        fmt.Println(fi4)
        var v5 uint64 = 682
        _ = fi4
        _ = v5
    }
    var v6 int32 = -1
    _ = v6
    return (-1 < -1)
}

func main() {
    var v1 uint = 456
    v2 := make(map[string]string)
    v3 := []string{}
    v1 -= 407
    v1 = (v1 + v1)
    v7 := v2["a"]
    v1 += v1
    v1 = 18446744073709551615
    var v8 uint16 = 65535
    if (255 >= 32767) {
        v8 += 65535
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
