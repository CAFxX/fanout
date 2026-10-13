package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS2 struct { F0 string; F1 uint }

func main() {
    var v1 int64 = 36
    var v2 uint16 = 881
    var v3 float32 = 1e10
    v4 := make(map[string]bool)
    for fi5 := 0; fi5 < 12; fi5++ {
        if true {
            v6 := GenS1{F0: 0, F1: 137, F2: v2}
            _ = v6
        } else {
            v7 := []int{0}
            switch ("a" + "a") {
            case "hello":
                delete(v4, (((("a" + "foo bar") + "hello") + "") + "a"))
            default:
                var v8 string = "foo bar"
                _ = v8
            }
            _ = v7
        }
        _ = fi5
    }
    var v9 int32 = 42
    v2 += v2
    v10 := GenS2{F0: "foo bar", F1: 2}
    v2 += (v2 & (v2 / ((v2) | 1)))
    var v11 bool = false
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 uint16 = (v2 | v2)
        v13 = v13
        v14 := v4["a"]
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int = 2147483647
    if (false && v11) {
        for fi16 := 0; fi16 < 2; fi16++ {
            fmt.Println(fi16, v11)
            v17 := v4["a"]
            for fi18 := 0; fi18 < 1; fi18++ {
                v3 += -2.25
                v15 = fi18
                _ = fi18
            }
            _ = fi16
            _ = v17
        }
    }
    v19 := []int{(v15 * (^(v15))), 9223372036854775807}
    v20 := GenS1{F0: v1, F1: 33, F2: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
