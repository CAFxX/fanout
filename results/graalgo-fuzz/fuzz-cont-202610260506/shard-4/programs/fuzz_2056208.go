package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 bool; F1 int64 }

func main() {
    var v1 int8 = -18
    var v2 int = 4294967295
    var v3 bool = (-22 != (-(v2)))
    var v4 int16 = 1000
    v5 := []string{}
    if (("a" + (("foo bar" + ("foo bar" + "\n\t")) + ((("hello" + "") + "a") + "\n\t"))) > "hello") {
        v4 = (v4 ^ (v4 ^ v4))
        var v6 bool = v3
        var v7 uint64 = 454
        _ = v6
        _ = v7
    } else {
        fmt.Println(v1)
        if false {
            var v8 int16 = 25
            fmt.Println(v8, v8, v1)
            _ = v8
        }
        v1 = (v1 - v1)
    }
    v9 := GenS1{F0: 42, F1: 448}
    v10 := GenS2{F0: v3, F1: 0}
    v1 = v1
    v3 = ((4294967295 < 608) || (v2 == (v2 ^ -1)))
    if true {
        var v11 uint64 = 2
        if false {
            fmt.Println(v2, v11)
            for fi12 := 0; fi12 < 14; fi12++ {
                v2 = (-(v2))
                v13 := GenS1{F0: 42, F1: 203}
                v2 -= 127
                _ = fi12
                _ = v13
            }
        }
        _ = v11
    } else {
        var v14 uint64 = 426
        v4 += (v4 & -32768)
        _ = v14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        v2 -= 0
        v5 = append(v5, "hello")
        _ = fi15
    }
    if (-417 >= v4) {
        var v16 bool = ((!false) && (((!v3) && v3) || (743 < 2)))
        v5 = append(v5, "")
        fmt.Println(v3, v1)
        _ = v16
    } else {
        v3 = v3
        if len(v5) > 0 {
            v5[0] = (("foo bar" + "foo bar") + "")
        }
    }
    if len(v5) > 0 {
        v5[0] = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
