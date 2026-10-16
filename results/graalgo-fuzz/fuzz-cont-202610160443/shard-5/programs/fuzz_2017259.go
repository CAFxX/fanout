package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 string }
func genFn0(p0 int64, p1 string, p2 uint8) uint8 {
    v7 := GenS1{F0: 23, F1: p0, F2: ("hello" + (p1 + p1))}
    p2 = p2
    v8 := []int{730}
    var v9 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (p2 >> ((65535) & 63))
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]bool)
    var v3 int16 = 1
    var v4 uint8 = 170
    var v5 int32 = 393
    v6 := []string{"a", "\n\t", ("\n\t" + ("\n\t" + "hello"))}
    v4 += (v4 & (v4 & (v4 & genFn0(-1, ("hello" + "hello"), genFn0(0, "\n\t", v4)))))
    switch "\n\t" {
    case "hello":
        v4 -= 127
    case "a":
        var v10 bool = (1e10 > -2.25)
        _ = v10
    case "foo bar":
        var v11 int16 = v3
        _ = v11
    }
    v4 -= v4
    fmt.Println(v4, v4, v4)
    for fi12 := 0; fi12 < 17; fi12++ {
        v6 = append(v6, "\n\t")
        _ = fi12
    }
    var v13 int64 = 309
    var v14 int64 = v13
    for fi15 := 0; fi15 < 19; fi15++ {
        v6 = append(v6, "a")
        var v16 int = (fi15 * 9223372036854775807)
        _ = fi15
        _ = v16
    }
    v4 = (v4 - v4)
    if len(v6) > 0 {
        v6[0] = "\n\t"
    }
    var v17 int64 = (v13 >> ((2) & 63))
    v13 = v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
