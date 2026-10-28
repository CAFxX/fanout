package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int32 }
type GenS2 struct { F0 bool; F1 bool; F2 bool }

func main() {
    v1 := []int{4294967295, 9223372036854775807, -9223372036854775808, -929}
    v2 := []string{"hello", "\n\t"}
    var v3 uint16 = 1
    v4 := make(map[string]bool)
    var v5 uint8 = 2
    var v6 uint = 65535
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: true, F1: "hello", F2: -1}
        fmt.Println(v5, v3, v6)
        v3 = v3
        _ = fi7
        _ = v8
    }
    fmt.Println(v5)
    fmt.Println(v3, v6, v6)
    if ("hello" > (("foo bar" + "a") + (((("a" + "hello") + "\n\t") + "\n\t") + (("a" + "") + "\n\t")))) {
        v9 := GenS1{F0: (811 != -9223372036854775808), F1: "foo bar", F2: -1}
        var v10 uint16 = v3
        _ = v10
        _ = v9
    } else {
        if len(v2) > 0 {
            v2[0] = "a"
        }
        var v11 int = 127
        var v12 uint32 = 4294967295
        _ = v11
        _ = v12
    }
    var v13 uint64 = 0
    v6 = (v6 | (v6 | (v6 - v6)))
    v14 := v4[""]
    fmt.Println(v3)
    v6 += 1
    if (!v14) {
        v4[(("hello" + ((("\n\t" + "\n\t") + "foo bar") + ("a" + "a"))) + "a")] = (!v14)
        fmt.Println(v3, v13, v13)
        if v14 {
            fmt.Println(v6, v5, v3)
        }
    } else {
        fmt.Println(v3)
        v13 += v13
        var v15 uint8 = v5
        _ = v15
    }
    v1 = append(v1, -508)
    v16 := GenS2{F0: (-32768 != 0), F1: false, F2: true}
    _ = len(v1)
    fmt.Println(v6, v5, v13)
    v13 -= (v13 % ((v13) | 1))
    var v17 int64 = -884
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
