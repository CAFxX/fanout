package main

import (
    "fmt"
)

func genFn0() uint8 {
    var v6 bool = (18446744073709551615 != 18446744073709551615)
    v6 = (0.0 > -2.25)
    var v7 int = 453
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    var v3 int64 = 1
    v4 := make(map[string]bool)
    var v5 uint = 278
    var v8 int32 = 42
    for fi9 := 0; fi9 < 9; fi9++ {
        v3 = (-(v3))
        _ = fi9
    }
    switch "\n\t" {
    case "foo bar":
        if false {
            v10 := []int{4294967295, int(v3), 730}
            _ = v10
        } else {
            var v11 uint32 = 65535
            _ = v11
        }
    }
    fmt.Println(v3)
    delete(v1, ("foo bar" + "a"))
    v3 = (v3 >> ((v5) & 63))
    switch "" {
    case "a":
        var v12 uint16 = 0
        _ = v12
    case "":
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 uint8 = genFn0()
            _ = fi13
            _ = v14
        }
    default:
        v5 = v5
    }
    v3 += 32
    var v15 uint8 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
