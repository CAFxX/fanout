package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 float32 = -2.25
    var v4 string = "\n\t"
    v5 := make(map[string]bool)
    var v6 uint32 = 2147483648
    v7 := v2[v4]
    fmt.Println(v4, v4)
    delete(v2, v4)
    switch ((("hello" + "foo bar") + (v4 + (v4 + "\n\t"))) + ("foo bar" + "")) {
    case "foo bar":
        if (!(42 != 42)) {
            v8 := GenS1{F0: 42, F1: 0, F2: 18446744073709551615}
            if true {
                v7 = 32768
                v9 := v1[(v7 << ((666) & 63))]
                _ = v9
            }
            v10 := v2[v4]
            _ = v10
            _ = v8
        } else {
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 uint16 = 909
                _ = fi11
                _ = v12
            }
        }
    case "hello":
        v13 := []int{(-(v7)), v7, v7}
        _ = v13
    case "\n\t":
        for fi14 := 0; fi14 < 3; fi14++ {
            v7 = (-(fi14))
            _ = fi14
        }
    default:
        v4 = "\n\t"
    }
    v6 = (^(v6))
    if (false && (!true)) {
        var v15 bool = true
        v4 = v4
        _ = v15
    } else {
        v2["a"] = (v7 ^ (v7 & (v7 * v7)))
        var v16 float64 = (0.0 - -2.25)
        v6 = (v6 * (v6 & (^(v6))))
        _ = v16
    }
    var v17 float32 = v3
    v4 = ("foo bar" + (v4 + "\n\t"))
    fmt.Println(v3)
    var v18 int = int(v6)
    var v19 float32 = v3
    v20 := []int{v18}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
