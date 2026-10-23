package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint16 }

func main() {
    var v1 uint32 = 65535
    var v2 int32 = 222
    var v3 int8 = 107
    var v4 uint32 = (^(v1))
    v5 := make(map[string]string)
    v6 := make(map[uint32]int)
    v1 += v1
    var v7 uint8 = 0
    v8 := GenS1{F0: v2, F1: 765, F2: 0}
    v2 -= v2
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := GenS1{F0: (v2 & v2), F1: -32768, F2: 65535}
        fmt.Println(v1, fi9, v2)
        _ = fi9
        _ = v10
    }
    v7 = (v7 / (((^(v7))) | 1))
    fmt.Println(v7, v3)
    if ("hello" != "") {
        var v11 string = ("foo bar" + "foo bar")
        v3 += v3
        v7 += (^(v7))
        _ = v11
    }
    var v12 int = 1
    switch (v12 ^ (v12 >> ((0) & 63))) {
    case -246:
        v13 := GenS1{F0: (v2 % ((-276) | 1)), F1: 830, F2: 700}
        _ = v13
    case 128:
        v14 := []int{}
        _ = v14
    case 0:
        delete(v5, (("" + (("hello" + (("a" + ("foo bar" + "foo bar")) + (("a" + "") + "a"))) + "")) + ("" + ("hello" + "\n\t"))))
    }
    v7 = v7
    v15 := GenS1{F0: (v2 & v2), F1: 42, F2: 2}
    v16 := []int{(v12 << ((634) & 63)), (^(v12)), (v12 % (((^(v12))) | 1))}
    var v17 float32 = (0.0 / 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
