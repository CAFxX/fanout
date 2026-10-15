package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int }
func genFn0(p0 uint, p1 uint64, p2 uint32) int32 {
    p2 -= p2
    v4 := GenS1{F0: 2147483648, F1: 42}
    if true {
        p1 = p1
        p2 += p2
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -2147483648
}

func main() {
    var v1 int32 = 655
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v5 string = ("hello" + "a")
    var v6 string = "hello"
    for fi7 := 0; fi7 < 7; fi7++ {
        v5 = "hello"
        _ = fi7
    }
    v8 := GenS2{F0: 109}
    if ((v5 + v6) <= "hello") {
        var v9 uint32 = 4294967295
        v10 := GenS1{F0: v9, F1: -880}
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 8; fi11++ {
            fmt.Println(v5, v5)
            _ = fi11
        }
        _ = len(v3)
        v6 = v5
    }
    if true {
        v6 = ("hello" + "a")
    }
    switch ((v6 + "") + ("foo bar" + (v6 + v6))) {
    case "foo bar":
        if false {
            fmt.Println(v1, v5, v1)
            fmt.Println(v1, v6, v1)
        }
    default:
        var v12 uint16 = 32767
        _ = v12
    }
    v1 = (v1 * (v1 & v1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v8
}
