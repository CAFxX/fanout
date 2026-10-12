package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint8 = 42
    var v2 uint16 = 42
    var v3 bool = (!(-2147483648 <= 2147483647))
    var v4 int8 = -128
    v5 := []string{(("foo bar" + "") + "hello"), "\n\t", "foo bar"}
    var v6 int32 = -866
    v1 -= v1
    v4 += (v4 << ((820) & 63))
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint = 4294967295
        v3 = (!false)
        _ = fi7
        _ = v8
    }
    if ((-1 < 131) || true) {
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := GenS1{F0: (v6 * v6)}
            _ = fi9
            _ = v10
        }
    }
    v5 = append(v5, "\n\t")
    if len(v5) > 0 {
        v5[0] = ""
    }
    fmt.Println(v6, v4, v6)
    v11 := GenS1{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
