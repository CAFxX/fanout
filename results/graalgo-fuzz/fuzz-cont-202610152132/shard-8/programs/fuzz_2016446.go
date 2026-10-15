package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 string; F1 int; F2 bool }

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]string)
    var v3 int = 136
    var v4 float64 = 1.5
    v5 := make(map[uint32]int)
    var v6 int64 = 1
    switch v3 {
    case -57:
        if (!(true || (true || (v1 <= 2)))) {
            if (v1 != 2) {
                v7 := v5[900]
                _ = v7
            }
            fmt.Println(v1, v6, v3)
            var v8 float64 = v4
            _ = v8
        }
    case 753:
        v1 += (v1 + v1)
    case 1:
        if true {
            v9 := []int{(v3 >> ((1) & 63)), v3, 446, 518}
            if true {
                var v10 int = v3
                fmt.Println(v4)
                _ = v10
            }
            v11 := GenS1{F0: 65535, F1: (false && ((v4 * -2.25) <= v4))}
            _ = v11
            _ = v9
        }
    default:
        v12 := GenS1{F0: 2, F1: false}
        _ = v12
    }
    v3 = v3
    for fi13 := 0; fi13 < 1; fi13++ {
        v1 += v1
        switch v3 {
        case -439:
            v14 := GenS1{F0: 42, F1: false}
            _ = v14
        case 173:
            v6 = (v6 / ((v6) | 1))
        case -9223372036854775808:
            v15 := GenS2{F0: "foo bar", F1: (v3 & v3), F2: (-1 == -32768)}
            _ = v15
        }
        var v16 uint16 = 42
        _ = fi13
        _ = v16
    }
    if true {
        var v17 uint = 87
        var v18 uint64 = uint64(v17)
        _ = v17
        _ = v18
    }
    fmt.Println(v3, v4)
    var v19 string = (("foo bar" + ("a" + ("" + ("foo bar" + "hello")))) + "\n\t")
    var v20 uint16 = 436
    fmt.Println(v6, v4, v6)
    var v21 int16 = 1
    for fi22 := 0; fi22 < 18; fi22++ {
        v19 = v19
        _ = fi22
    }
    v6 = (v6 | v6)
    v3 = (v3 / ((v3) | 1))
    var v23 uint64 = 475
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
