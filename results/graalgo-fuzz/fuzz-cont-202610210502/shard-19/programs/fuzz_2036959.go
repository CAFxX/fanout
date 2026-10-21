package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int64 }
func genFn0() uint8 {
    if (("foo bar" + (("" + (("hello" + "foo bar") + "foo bar")) + "foo bar")) >= "") {
        v7 := []int{9223372036854775807, 9223372036854775807, 9223372036854775807}
        _ = v7
    }
    var v8 string = "foo bar"
    _ = v8
    return 1
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]int)
    var v5 string = (("" + ("foo bar" + ((("\n\t" + "hello") + "hello") + (("a" + (("foo bar" + ((("hello" + ("a" + "foo bar")) + "\n\t") + "a")) + "")) + "a")))) + ("foo bar" + "hello"))
    var v6 uint8 = 1
    for fi9 := 0; fi9 < 19; fi9++ {
        v5 = v5
        v5 = "hello"
        _ = fi9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS1{F0: true}
        var v12 uint64 = 1
        v13 := GenS2{F0: (-(fi10)), F1: 9223372036854775807}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 uint16 = 795
    v5 = v5
    if false {
        var v15 uint8 = (v6 + v6)
        var v16 uint64 = 42
        var v17 uint32 = 2147483647
        _ = v15
        _ = v16
        _ = v17
    }
    v5 = v5
    v6 = v6
    var v18 uint32 = 658
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
