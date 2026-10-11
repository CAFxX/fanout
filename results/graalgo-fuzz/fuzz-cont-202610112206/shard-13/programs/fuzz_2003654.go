package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int64, p1 string, p2 uint) int8 {
    v7 := GenS1{F0: p2}
    if (((1 == 2147483647) || (42 <= (^(p2)))) && (!false)) {
        var v8 uint16 = uint16(p2)
        _ = v8
    } else {
        p0 += -979
        p2 += (p2 >> (((p2 * p2)) & 63))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -114
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 518
    var v3 bool = (667 == -316)
    v4 := make(map[string]string)
    v5 := []string{"foo bar", "foo bar", "foo bar", "a"}
    var v6 uint64 = (v2 ^ v2)
    switch "foo bar" {
    case "hello":
        v3 = v3
    }
    v2 += (v2 | (v2 ^ v6))
    v6 = v6
    fmt.Println(v3)
    v2 = v2
    fmt.Println(v2, v6)
    var v9 string = "foo bar"
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    delete(v1, 2147483647)
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
