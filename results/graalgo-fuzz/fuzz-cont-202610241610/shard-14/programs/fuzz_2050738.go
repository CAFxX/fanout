package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 30
    var v3 int8 = -128
    var v4 float64 = 1e10
    var v5 uint16 = 65535
    var v6 uint32 = 2
    for fi7 := 0; fi7 < 5; fi7++ {
        switch "\n\t" {
        case "hello":
            v3 -= v3
        case "foo bar":
            for fi8 := 0; fi8 < 14; fi8++ {
                delete(v1, fi8)
                _ = fi8
            }
        case "":
            var v9 int32 = 424
            _ = v9
        }
        v3 = 89
        v2 -= (v2 % ((v2) | 1))
        _ = fi7
    }
    v10 := GenS1{F0: (v5 + 65535), F1: v5}
    var v11 uint = 65535
    var v12 int64 = 42
    v13 := GenS1{F0: (v5 * v5), F1: v5}
    v14 := []int{}
    v4 += (v4 + v4)
    if (9223372036854775807 >= int(v2)) {
        delete(v1, 1)
        delete(v1, 584)
    }
    if false {
        var v15 uint64 = 4294967295
        v15 -= uint64(v2)
        _ = v15
    } else {
        if len(v14) > 0 {
            v14[0] = int(v2)
        }
        v12 = (v12 & (-(v12)))
        switch -150 {
        case -2147483648:
            v16 := GenS2{F0: v2}
            _ = v16
        case 42:
            fmt.Println(v5)
        default:
            var v17 uint8 = (v2 * (^(v2)))
            _ = v17
        }
    }
    v1[42] = -1
    v5 -= 255
    v6 = 0
    for fi18 := 0; fi18 < 20; fi18++ {
        _ = len(v1)
        if false {
            if len(v14) > 0 {
                v14[0] = (fi18 % ((fi18) | 1))
            }
        }
        _ = fi18
    }
    switch 42 {
    case 9223372036854775807:
        v14 = append(v14, 42)
    case -194:
        v2 = 113
    default:
        for fi19 := 0; fi19 < 1; fi19++ {
            v14 = append(v14, (fi19 >> ((275) & 63)))
            _ = fi19
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
