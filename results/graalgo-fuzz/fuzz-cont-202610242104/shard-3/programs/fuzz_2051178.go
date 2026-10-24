package main

import (
    "fmt"
)

func genFn0() uint {
    var v6 int = 0
    var v7 uint16 = 255
    v7 += (v7 - (v7 & (v7 ^ v7)))
    _ = v6
    _ = v7
    return 593
}
func genFn1(p0 int32, p1 uint32, p2 uint8) uint8 {
    var v8 bool = (!true)
    p1 = p1
    p2 += 227
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2
}

func main() {
    var v1 int = 0
    v2 := []int{}
    v3 := make(map[string]int)
    v4 := []string{((("\n\t" + "\n\t") + "\n\t") + "hello"), (("\n\t" + "foo bar") + "hello"), "a", ("hello" + "")}
    var v5 bool = false
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 int8 = 77
        v10 = (v10 << ((genFn0()) & 63))
        _ = fi9
        _ = v10
    }
    var v11 uint32 = 1
    v5 = true
    v12 := v3["\n\t"]
    var v13 uint64 = 497
    fmt.Println(v1, v11)
    v12 += (^(v1))
    var v14 uint16 = 255
    v12 = 9223372036854775807
    v15 := v3[(("" + "\n\t") + "a")]
    v1 += 4294967295
    v1 = 4294967295
    var v16 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
