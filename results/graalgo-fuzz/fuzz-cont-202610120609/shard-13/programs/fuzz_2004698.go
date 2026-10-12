package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int32 }
func genFn0(p0 int32) int64 {
    var v5 int64 = -43
    v5 = 0
    v5 = v5
    switch 2147483647 {
    case 32767:
        p0 = (p0 - -1)
    case 24:
        var v6 uint32 = 2147483647
        _ = v6
    }
    v7 := GenS1{F0: (v5 - 42)}
    _ = p0
    _ = v5
    _ = v7
    return -9223372036854775808
}
func genFn1() int {
    var v8 string = "foo bar"
    v9 := []int{}
    if true {
        fmt.Println(v8)
    }
    v8 = "hello"
    _ = v8
    _ = v9
    return 128
}

func main() {
    v1 := []string{"foo bar", ("foo bar" + ("\n\t" + "foo bar"))}
    v2 := make(map[uint32]string)
    var v3 string = (("a" + "a") + "\n\t")
    var v4 int16 = -733
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = v3
    var v10 uint = 291
    fmt.Println(v4, v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v10, v10, v4)
    var v11 uint32 = 0
    fmt.Println(v11)
    var v12 int = 963
    for fi13 := 0; fi13 < 7; fi13++ {
        _ = len(v2)
        _ = fi13
    }
    v14 := GenS2{F0: 0, F1: 542, F2: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
