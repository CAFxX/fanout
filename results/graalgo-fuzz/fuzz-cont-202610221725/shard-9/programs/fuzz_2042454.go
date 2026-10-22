package main

import (
    "fmt"
)

func genFn0(p0 int64) int64 {
    p0 += (-(p0))
    fmt.Println(p0)
    v7 := []int{}
    _ = len(v7)
    var v8 string = "hello"
    _ = p0
    _ = v7
    _ = v8
    return p0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 0
    v3 := []int{-977, 9223372036854775807, 9223372036854775807, 840}
    var v4 int = 0
    var v5 int8 = -105
    var v6 int32 = -1
    var v9 bool = (("a" + "foo bar") == ("foo bar" + ("a" + "a")))
    var v10 int = (v4 & (v4 | (v4 * 32767)))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v11 int64 = 1
    v9 = ((3.14159 / 3.14159) <= -2.25)
    var v12 uint8 = 1
    v4 += (v10 & (v10 * (v10 - (v10 - (v10 ^ int(v6))))))
    v6 = (-(v6))
    fmt.Println(v12, v11, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
    _ = v9
}
