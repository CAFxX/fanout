package main

import (
    "fmt"
)

func genFn0(p0 int16) int16 {
    v7 := []int{-770, 127}
    if false {
        _ = len(v7)
        p0 += 848
    }
    p0 += 32767
    p0 += -648
    _ = len(v7)
    _ = p0
    _ = v7
    return p0
}

func main() {
    var v1 int16 = 1
    var v2 string = "foo bar"
    var v3 string = v2
    v4 := []int{-549, 0, 0, -1}
    v5 := []uint32{458, 65535, 2147483648, 42}
    v6 := make(map[uint32]uint32)
    v2 = "foo bar"
    var v8 int32 = -220
    fmt.Println(v3, v1, v1)
    var v9 int = 32768
    v3 = (("\n\t" + "hello") + ("a" + v3))
    var v10 int32 = v8
    _ = len(v4)
    var v11 uint32 = 42
    fmt.Println(v9)
    v2 = ((("\n\t" + (((("a" + v2) + v3) + v3) + (v3 + (v3 + ("hello" + v3))))) + v2) + "foo bar")
    v9 -= v9
    var v12 uint32 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
