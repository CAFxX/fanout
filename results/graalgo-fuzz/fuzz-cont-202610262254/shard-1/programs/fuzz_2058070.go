package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    v3 := make(map[int]string)
    var v4 string = "foo bar"
    v5 := make(map[int]bool)
    v6 := make(map[string]uint32)
    v4 = ("foo bar" + v4)
    var v7 uint16 = 65535
    switch -984 {
    case -176:
        v4 = ((("\n\t" + "foo bar") + "") + "foo bar")
    }
    v8 := []int{32767, -9223372036854775808}
    var v9 int16 = 42
    v9 = (v9 << ((263) & 63))
    var v10 int = 4294967295
    var v11 int32 = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
