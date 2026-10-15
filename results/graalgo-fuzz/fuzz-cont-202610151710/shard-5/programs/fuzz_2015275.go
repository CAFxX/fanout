package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 float32 = -2.25
    var v2 string = "\n\t"
    v3 := make(map[uint32]bool)
    var v4 uint8 = 85
    v2 = v2
    var v5 uint64 = 18446744073709551615
    switch (("a" + ((v2 + v2) + (v2 + v2))) + "foo bar") {
    case "\n\t":
        var v6 uint32 = 1
        _ = v6
    }
    v2 = "a"
    var v7 uint16 = 32767
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint = 623
        v9 = (v9 * v9)
        _ = fi8
        _ = v9
    }
    delete(v3, 1)
    v7 += v7
    fmt.Println(v4, v4, v2)
    if (!true) {
        v4 += v4
        v10 := GenS1{F0: 2}
        fmt.Println(v5, v2, v5)
        _ = v10
    }
    v2 = "foo bar"
    v11 := GenS1{F0: 4294967295}
    var v12 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
