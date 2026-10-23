package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 string }

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]int)
    var v3 uint = 859
    v4 := make(map[int]string)
    var v5 int8 = 127
    v6 := []uint32{12, 42}
    v5 = v5
    v3 = 512
    if true {
        fmt.Println(v1, v3)
        v3 = v3
    }
    if true {
        for fi7 := 0; fi7 < 18; fi7++ {
            fmt.Println(fi7, v1, fi7)
            _ = len(v6)
            v8 := GenS1{F0: -9223372036854775808, F1: 926}
            _ = fi7
            _ = v8
        }
    } else {
        v6 = append(v6, 2)
        v9 := GenS2{F0: ("\n\t" + (("\n\t" + "\n\t") + "foo bar"))}
        _ = len(v6)
        _ = v9
    }
    v1 = (v1 * 1)
    v6 = append(v6, 42)
    v10 := GenS1{F0: 42, F1: uint32(v3)}
    v4[128] = ((("" + "") + (("foo bar" + (("hello" + "foo bar") + ("" + (("\n\t" + "") + "hello")))) + "hello")) + "")
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v11 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
