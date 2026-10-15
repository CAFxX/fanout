package main

import (
    "fmt"
)

func genFn0(p0 int32, p1 int32, p2 uint64) string {
    var v7 int32 = (p0 << ((1) & 63))
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "hello"
}

func main() {
    v1 := []string{"a", ("\n\t" + "foo bar")}
    var v2 uint64 = 63
    var v3 int16 = 1000
    var v4 uint = 42
    v5 := make(map[string]bool)
    v6 := make(map[uint32]bool)
    var v8 bool = true
    if false {
        var v9 int8 = -128
        _ = v9
    }
    v3 = int16(v4)
    v5["\n\t"] = (37 != 127)
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int = (fi10 % (((-(fi10))) | 1))
        _ = fi10
        _ = v11
    }
    switch "hello" {
    case "hello":
        v5[(("a" + "foo bar") + "hello")] = true
    case "foo bar":
        fmt.Println(v8, v4)
    case "":
        var v12 uint16 = 2
        _ = v12
    }
    v2 += v2
    v3 = (v3 & 0)
    fmt.Println(v3, v2)
    if v8 {
        if v8 {
            var v13 int32 = -701
            v4 -= (v4 ^ (v4 | v4))
            for fi14 := 0; fi14 < 4; fi14++ {
                switch (fi14 - (fi14 << (((v4 - 65535)) & 63))) {
                case -790:
                    if (false || (65535 != 2)) {
                        v2 += (v2 % ((v2) | 1))
                        fmt.Println(v8, v2, v4)
                    } else {
                        fmt.Println(fi14)
                    }
                default:
                    fmt.Println(v2)
                }
                var v15 float32 = (float32(((0.1 / 0.1) - 0.1)) / 1e10)
                _ = fi14
                _ = v15
            }
            _ = v13
        } else {
            var v16 string = ""
            _ = v16
        }
        var v17 int64 = 9223372036854775807
        _ = v17
    }
    fmt.Println(v3, v4)
    var v18 uint32 = 965
    var v19 uint16 = 42
    v18 = v18
    var v20 uint = uint(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
