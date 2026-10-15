package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 bool }

func main() {
    v1 := []string{(("a" + (("\n\t" + "a") + "foo bar")) + "hello"), "", "foo bar", "hello"}
    v2 := make(map[uint32]int)
    var v3 bool = (((42 == -816) || false) && (0 <= -560))
    v4 := make(map[uint32]int)
    if (81 >= 0) {
        var v5 uint = 65535
        _ = v5
    }
    v6 := GenS1{F0: 9223372036854775807, F1: 128, F2: v3}
    _ = len(v1)
    if (-1 >= -1) {
        fmt.Println(v3)
        v3 = (!v3)
        for fi7 := 0; fi7 < 15; fi7++ {
            v8 := GenS1{F0: (fi7 * -2147483648), F1: 2, F2: (v3 && false)}
            _ = fi7
            _ = v8
        }
    }
    switch "a" {
    case "":
        if (((2147483647 != -662) && true) || ((("" + "hello") + "a") != "a")) {
            var v9 int = -1
            v9 -= (v9 * (^(v9)))
            _ = v9
        }
    }
    var v10 int32 = -1
    var v11 bool = (0 != -1)
    var v12 int8 = -72
    v10 -= (v10 << ((406) & 63))
    var v13 uint16 = 42
    v1 = append(v1, "")
    v12 -= (v12 ^ (v12 << ((363) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
