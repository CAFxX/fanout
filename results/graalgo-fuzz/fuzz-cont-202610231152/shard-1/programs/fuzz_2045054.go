package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 732
    v3 := []string{"\n\t", "foo bar"}
    _ = len(v3)
    if (42 <= 42) {
        fmt.Println(v2)
        var v4 int32 = 85
        v2 = (v2 & 65535)
        _ = v4
    }
    v2 = v2
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: 65535}
        v7 := GenS1{F0: (v2 & v2)}
        v2 += v2
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(v2)
    switch ((("hello" + "\n\t") + "\n\t") + "hello") {
    case "a":
        v2 = 18446744073709551615
    default:
        var v8 uint32 = 0
        _ = v8
    }
    var v9 uint = 255
    var v10 uint32 = 91
    fmt.Println(v10)
    v11 := GenS1{F0: 2}
    for fi12 := 0; fi12 < 12; fi12++ {
        v10 += v10
        v10 -= v10
        _ = fi12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
