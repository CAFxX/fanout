package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int16 = 381
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    var v4 uint = 343
    delete(v3, "\n\t")
    v3[""] = 4294967295
    v4 += (v4 * v4)
    if ((false && false) && false) {
        v5 := []int{-366, 128}
        for fi6 := 0; fi6 < 6; fi6++ {
            if false {
                delete(v2, 0)
            }
            var v7 uint64 = 0
            switch "hello" {
            case "a":
                v8 := GenS1{F0: v4, F1: ""}
                _ = v8
            }
            _ = fi6
            _ = v7
        }
        _ = v5
    } else {
        var v9 uint32 = 1
        v10 := GenS2{F0: 445}
        var v11 uint32 = v9
        _ = v10
        _ = v11
        _ = v9
    }
    if true {
        if (!(3.14159 > 1.5)) {
            v4 -= 2
        } else {
            switch "" {
            case "":
                v12 := []int{9223372036854775807, 483, -9223372036854775808, 9223372036854775807}
                _ = v12
            case "foo bar":
                v4 += (v4 << ((18446744073709551615) & 63))
            }
            delete(v3, "\n\t")
        }
    }
    fmt.Println(v1, v4)
    if false {
        v13 := []int{128}
        _ = v13
    }
    v14 := []int{127, -9223372036854775808, 0, -1}
    var v15 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
