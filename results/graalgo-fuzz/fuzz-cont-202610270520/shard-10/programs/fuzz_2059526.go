package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 string }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{0, 42, 500, 0}
    v3 := []int{-2147483648}
    var v4 bool = (false || (676 >= 908))
    v5 := make(map[string]int)
    v2 = append(v2, 0)
    var v6 string = "\n\t"
    v6 = "foo bar"
    v6 = v6
    if v4 {
        v6 = ((v6 + (v6 + v6)) + (((("" + v6) + v6) + v6) + v6))
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        v6 = (((v6 + (((v6 + (v6 + "a")) + v6) + ("a" + ("foo bar" + (v6 + (((v6 + "foo bar") + (v6 + v6)) + "hello")))))) + (("foo bar" + "\n\t") + "a")) + "foo bar")
        var v8 int16 = -32768
        _ = fi7
        _ = v8
    }
    if (true && v4) {
        v4 = v4
        var v9 uint8 = 0
        _ = v9
    } else {
        fmt.Println(v4, v6)
    }
    var v10 uint32 = 1
    var v11 float32 = (3.14159 + 1.5)
    fmt.Println(v4)
    var v12 int16 = -378
    v2 = append(v2, 65535)
    var v13 int = 706
    var v14 int16 = (v12 - (v12 - -138))
    if (!(943 == 253)) {
        v11 = float32(((0.1 * 1.5) - ((0.1 - 0.1) / 3.14159)))
    } else {
        var v15 uint16 = 1
        var v16 bool = (127 == -128)
        _ = v15
        _ = v16
    }
    fmt.Println(v4)
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 uint64 = 18446744073709551615
        _ = len(v5)
        v19 := GenS1{F0: 42, F1: v10, F2: "a"}
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 int16 = -708
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
