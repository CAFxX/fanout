package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int16, p1 string, p2 uint64) bool {
    v5 := []int{-799}
    p1 = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return true
}

func main() {
    var v1 int = 4294967295
    var v2 uint16 = 945
    var v3 uint = 42
    v4 := make(map[string]uint32)
    var v6 uint64 = 18446744073709551615
    switch "\n\t" {
    case "a":
        v2 += (v2 << ((436) & 63))
    case "foo bar":
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := GenS1{F0: -9223372036854775808}
            _ = fi7
            _ = v8
        }
    case "":
        v9 := []int{(v1 >> ((65535) & 63)), 1, v1, v1}
        _ = v9
    }
    v10 := GenS1{F0: -9223372036854775808}
    fmt.Println(v3, v3)
    v11 := []int{}
    v4["hello"] = 2147483648
    var v12 uint8 = 2
    _ = len(v11)
    var v13 uint32 = 449
    if len(v11) > 0 {
        v11[0] = (v1 % (((v1 >> ((255) & 63))) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
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
