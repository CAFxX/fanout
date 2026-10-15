package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint32 }
func genFn0() uint16 {
    var v7 uint32 = 1
    v8 := GenS1{F0: (!(1 > 42)), F1: 546}
    _ = v7
    _ = v8
    return 150
}
func genFn1(p0 bool, p1 uint16) uint64 {
    var v9 uint32 = 42
    v10 := GenS1{F0: ((32767 == 0) || p0), F1: -319}
    if p0 {
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 uint = 1
            v9 = (v9 - uint32(p1))
            _ = fi11
            _ = v12
        }
        fmt.Println(p0, p1, v9)
        var v13 int32 = 400
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 245
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = (2 <= 1)
    var v3 int32 = 0
    v4 := []uint32{}
    var v5 float32 = (0.1 + 3.14159)
    v6 := make(map[uint32]bool)
    if len(v4) > 0 {
        v4[0] = 1
    }
    var v14 uint16 = 65535
    v14 -= (v14 / ((2) | 1))
    v6[65535] = (!(134 > 2))
    v15 := GenS2{F0: 2}
    v2 = false
    var v16 int64 = 1
    var v17 string = ((("foo bar" + ((((((("hello" + ("" + ("a" + ("a" + (("" + "") + (("a" + "") + "a")))))) + ("a" + ("\n\t" + "foo bar"))) + "foo bar") + ("\n\t" + "foo bar")) + ("\n\t" + ("foo bar" + "\n\t"))) + "foo bar") + (("\n\t" + "a") + ("a" + "hello")))) + "hello") + "a")
    var v18 int32 = int32(v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
