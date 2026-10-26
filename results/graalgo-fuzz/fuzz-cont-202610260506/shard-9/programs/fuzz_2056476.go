package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[string]string)
    var v2 string = "\n\t"
    var v3 int8 = 71
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    if (2 > 978) {
        v6 := []int{378, 9223372036854775807}
        _ = v6
    }
    fmt.Println(v3, v3)
    v7 := GenS1{F0: (v3 - v3)}
    for fi8 := 0; fi8 < 2; fi8++ {
        v4["foo bar"] = (fi8 >> ((18446744073709551615) & 63))
        _ = fi8
    }
    v2 = (v2 + v2)
    v9 := GenS1{F0: 42}
    for fi10 := 0; fi10 < 14; fi10++ {
        switch v2 {
        case "":
            v11 := []int{(fi10 ^ (fi10 % ((fi10) | 1))), (-(fi10)), fi10, fi10}
            _ = v11
        }
        _ = fi10
    }
    v12 := v5[9223372036854775807]
    switch 2147483647 {
    case 1:
        v3 += v3
    case 128:
        var v13 bool = false
        _ = v13
    default:
        v2 = "foo bar"
    }
    v14 := GenS1{F0: v3}
    v3 = (-(v3))
    var v15 uint32 = 315
    delete(v4, "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
