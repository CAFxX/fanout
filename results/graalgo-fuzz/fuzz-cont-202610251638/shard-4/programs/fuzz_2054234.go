package main

import (
    "fmt"
)

func genFn0(p0 int32, p1 int32, p2 uint16) bool {
    p0 = p1
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 int = fi5
        for fi7 := 0; fi7 < 18; fi7++ {
            if false {
                p2 -= p2
            } else {
                p2 -= 1
                if false {
                    fmt.Println(fi5)
                    v6 = fi7
                    v6 += -1
                }
            }
            v6 = 127
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    v8 := []int{-9223372036854775808, 2147483647}
    for fi9 := 0; fi9 < 9; fi9++ {
        switch ("hello" + "a") {
        case "hello":
            _ = len(v8)
        case "foo bar":
            p2 -= p2
        }
        fmt.Println(p2)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return true
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    v4 := make(map[int]uint32)
    delete(v2, -915)
    v1 -= 2147483648
    var v10 bool = (v1 != (v1 % ((120) | 1)))
    v11 := []int{32768, -476}
    v12 := v2[-2147483648]
    var v13 uint = 4294967295
    fmt.Println(v13, v12, v13)
    var v14 int = -613
    var v15 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
