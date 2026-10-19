package main

import (
    "fmt"
)

func genFn0(p0 uint, p1 uint8, p2 int64) uint32 {
    v7 := []int{-31, 128, 127, -824}
    fmt.Println(p0)
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int8 = 42
        _ = fi8
        _ = v9
    }
    fmt.Println(p1, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}

func main() {
    v1 := make(map[int]int)
    var v2 string = ((("a" + ("" + ("hello" + (("\n\t" + (("\n\t" + "\n\t") + "foo bar")) + "hello")))) + ("hello" + ("foo bar" + "\n\t"))) + "")
    var v3 float64 = 0.0
    v4 := []uint32{1}
    var v5 uint8 = 20
    var v6 uint16 = 705
    v5 -= 127
    var v10 uint16 = 0
    v3 -= v3
    v10 = (v6 * 924)
    v5 = (^(v5))
    var v11 uint16 = 255
    v6 = 839
    fmt.Println(v6, v2, v5)
    v11 -= 941
    switch v2 {
    case "a":
        if len(v4) > 0 {
            v4[0] = 65535
        }
    case "":
        var v12 int16 = 0
        _ = v12
    case "\n\t":
        for fi13 := 0; fi13 < 11; fi13++ {
            v10 = 32767
            _ = fi13
        }
    }
    v3 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
