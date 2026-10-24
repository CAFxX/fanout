package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint64; F1 string; F2 int32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"hello", "a"}
    v3 := make(map[string]uint32)
    var v4 uint8 = 42
    var v5 uint = 653
    var v6 int32 = 430
    v7 := GenS1{F0: "foo bar", F1: 65, F2: (v4 + v4)}
    for fi8 := 0; fi8 < 7; fi8++ {
        _ = len(v3)
        _ = fi8
    }
    fmt.Println(v5, v4, v6)
    if (false || true) {
        v9 := GenS1{F0: ("\n\t" + ""), F1: 97, F2: (v4 ^ (v4 | 255))}
        delete(v3, (("foo bar" + "hello") + ""))
        fmt.Println(v6)
        _ = v9
    } else {
        if (!true) {
            var v10 int32 = v6
            v4 += (v4 + (v4 & v4))
            _ = v10
        } else {
            var v11 bool = true
            if v11 {
                v12 := GenS1{F0: "", F1: -1, F2: v4}
                _ = v12
            } else {
                var v13 int64 = 9223372036854775807
                var v14 uint8 = 207
                _ = v13
                _ = v14
            }
            _ = v11
        }
        v5 -= 2
        v6 = v6
    }
    var v15 uint8 = v4
    var v16 uint64 = uint64(v15)
    if (2147483648 <= 0) {
        var v17 int64 = 9223372036854775807
        v18 := GenS2{F0: v16, F1: ("hello" + ("hello" + "")), F2: 176}
        var v19 uint8 = v15
        _ = v17
        _ = v18
        _ = v19
    } else {
        v5 = 42
        v15 -= (v4 | v15)
        fmt.Println(v15, v4)
    }
    v6 -= v6
    v3[""] = 0
    v6 = (v6 - (v6 % ((v6) | 1)))
    v4 = v4
    var v20 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
