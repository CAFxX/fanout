package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]int)
    v3 := []string{"foo bar", "", "hello", "a"}
    var v4 bool = ("foo bar" < "foo bar")
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]int)
    v4 = (1 <= 590)
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v7 := GenS1{F0: 18446744073709551615}
    if (100 != 45) {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        v4 = v4
    }
    v1 -= (v1 - (v1 + v1))
    v4 = (0 != 32767)
    fmt.Println(v4, v1)
    if (v1 >= (^(v1))) {
        v8 := GenS1{F0: 4294967295}
        switch "hello" {
        case "a":
            v9 := GenS2{F0: -32}
            _ = v9
        case "":
            var v10 int64 = 1
            _ = v10
        }
        _ = v8
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        fmt.Println(fi11, v1, v1)
        _ = fi11
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
