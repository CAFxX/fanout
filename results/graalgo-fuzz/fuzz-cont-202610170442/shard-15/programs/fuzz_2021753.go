package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []string{"hello", "\n\t", ("" + "foo bar"), "a"}
    var v3 uint32 = 65535
    var v4 string = "\n\t"
    v3 += (v3 & (v3 + 42))
    var v5 uint64 = 42
    _ = len(v2)
    fmt.Println(v3, v3, v5)
    fmt.Println(v1)
    var v6 int64 = (v1 << ((uint(v3)) & 63))
    v3 = v3
    v3 = (^(v3))
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint = 20
        v5 -= 658
        v9 := GenS1{F0: (v3 * v3)}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int = -52
    v11 := GenS1{F0: (^(v3))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
