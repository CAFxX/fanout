package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 87
    v3 := make(map[string]uint32)
    var v4 int16 = 32767
    v5 := GenS1{F0: "\n\t"}
    v6 := GenS1{F0: "a"}
    var v7 int32 = -441
    v2 -= (v2 * 161)
    v8 := GenS1{F0: "a"}
    if true {
        var v9 bool = (10 < 4294967295)
        v2 = 612
        _ = v9
    }
    v10 := GenS1{F0: (("foo bar" + "") + "a")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
