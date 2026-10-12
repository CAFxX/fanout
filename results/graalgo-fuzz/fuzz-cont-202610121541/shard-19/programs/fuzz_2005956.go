package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int8 }
func genFn0(p0 uint, p1 string, p2 uint) uint64 {
    fmt.Println(p2, p0)
    var v6 int32 = 0
    fmt.Println(v6, v6, v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{(("foo bar" + (((("" + "hello") + "") + ((("\n\t" + ("\n\t" + "foo bar")) + (("hello" + (("\n\t" + "") + "")) + "")) + ("hello" + "a"))) + "")) + ("hello" + ("hello" + "\n\t"))), ("" + ""), ("foo bar" + (("" + ("" + "")) + "\n\t")), ""}
    var v3 int16 = 866
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]bool)
    fmt.Println(v3)
    var v7 float64 = 0.1
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 uint32 = 876
        v10 := GenS1{F0: 4294967295, F1: 18446744073709551615, F2: 127}
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(v3)
    var v11 uint8 = 127
    var v12 uint16 = 255
    var v13 float64 = v7
    v3 = 1000
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 int8 = 127
        var v16 int = fi14
        v12 -= (v12 % ((v12) | 1))
        _ = fi14
        _ = v15
        _ = v16
    }
    switch "a" {
    case "\n\t":
        var v17 float32 = (3.14159 + (((0.1 / 1.5) - -2.25) + 0.0))
        _ = v17
    case "foo bar":
        delete(v1, uint32(v11))
    }
    v18 := GenS1{F0: 18446744073709551615, F1: 0, F2: 42}
    fmt.Println(v13, v11, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
