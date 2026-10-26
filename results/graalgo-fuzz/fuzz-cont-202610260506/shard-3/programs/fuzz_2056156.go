package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []uint32{2147483648, 4294967295, 741, 2147483648}
    var v2 string = "\n\t"
    var v3 int = -2147483648
    var v4 uint32 = 4294967295
    v4 -= 2
    v3 += v3
    var v5 int32 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 127 {
    case 127:
        v6 := GenS1{F0: 1}
        _ = v6
    default:
        fmt.Println(v5)
    }
    v2 = (("foo bar" + v2) + ("hello" + v2))
    v2 = (v2 + v2)
    v2 = "\n\t"
    for fi7 := 0; fi7 < 6; fi7++ {
        if len(v1) > 0 {
            v1[0] = v4
        }
        _ = fi7
    }
    fmt.Println(v2, v2)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
