package main

import (
    "fmt"
)

func genFn0(p0 bool, p1 int16, p2 int8) bool {
    v6 := []int{1, -9223372036854775808}
    var v7 uint32 = 2147483648
    var v8 int = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return p0
}

func main() {
    v1 := []int{}
    var v2 string = (((("a" + "") + ("\n\t" + "hello")) + "foo bar") + "hello")
    var v3 int16 = -1
    v4 := make(map[int]int)
    var v5 uint = 18446744073709551615
    var v9 int8 = 1
    var v10 float32 = 1e10
    v9 = v9
    v2 = v2
    v9 = v9
    for fi11 := 0; fi11 < 1; fi11++ {
        v1 = append(v1, (fi11 >> (((v5 - (^(v5)))) & 63)))
        _ = fi11
    }
    fmt.Println(v10, v5)
    fmt.Println(v3, v5)
    var v12 uint64 = 232
    var v13 bool = false
    var v14 uint8 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
