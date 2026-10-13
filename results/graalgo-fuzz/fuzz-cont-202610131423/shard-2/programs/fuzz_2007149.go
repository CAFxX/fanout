package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int }

func main() {
    var v1 int = -1
    v2 := []string{"hello", "", ("hello" + "foo bar")}
    v3 := make(map[int]uint32)
    v4 := []string{"a", (("foo bar" + "\n\t") + "a"), "foo bar", "foo bar"}
    var v5 bool = (true && ((-489 == -449) || ((true && (true && false)) || (!(-1 <= 0)))))
    fmt.Println(v5, v1)
    if false {
        var v6 uint64 = 42
        _ = v6
    }
    v5 = false
    v5 = v5
    var v7 string = "a"
    v8 := GenS1{F0: -9223372036854775808, F1: -79}
    v9 := GenS2{F0: v1}
    v7 = ""
    if (2 >= 206) {
        v5 = (-1 >= 9223372036854775807)
        var v10 int32 = 347
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint16 = 1
            var v13 bool = (("foo bar" >= "") && (v5 && v5))
            delete(v3, (^(fi11)))
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
    }
    fmt.Println(v5)
    v7 = v7
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
