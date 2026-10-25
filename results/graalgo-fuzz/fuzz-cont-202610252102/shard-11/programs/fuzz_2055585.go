package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() string {
    v5 := []int{275, -855}
    if len(v5) > 0 {
        v5[0] = -644
    }
    _ = v5
    return "a"
}

func main() {
    v1 := []string{(("foo bar" + "foo bar") + "a"), "\n\t", "hello", "a"}
    var v2 int = -2147483648
    v3 := []string{}
    v4 := []string{"hello", ""}
    if (42 <= -1) {
        v2 = (-(v2))
    } else {
        v2 += v2
    }
    var v6 int64 = 1
    v2 -= (^(v2))
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint16 = 354
        v2 = (fi7 << ((400) & 63))
        var v9 int64 = v6
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint64 = 248
    v6 -= v6
    var v11 int32 = -167
    if true {
        var v12 uint16 = 32767
        var v13 int8 = -98
        v14 := GenS1{F0: (v12 >> ((42) & 63))}
        _ = v12
        _ = v13
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
