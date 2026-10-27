package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int16 {
    var v5 string = "hello"
    fmt.Println(v5)
    v6 := []int{32767, 1, 363, 1}
    _ = v5
    _ = v6
    return -694
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]bool)
    var v4 int = 32768
    v4 -= v4
    v4 -= v4
    if (-128 <= 1) {
        var v7 bool = ((v2 << ((582) & 63)) != v2)
        _ = v7
    }
    v4 -= 1
    if false {
        if true {
            fmt.Println(v4)
            for fi8 := 0; fi8 < 16; fi8++ {
                v4 -= 42
                _ = fi8
            }
            var v9 int32 = 544
            _ = v9
        }
    }
    var v10 int = (v4 >> ((1) & 63))
    switch "\n\t" {
    case "":
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 uint64 = 18446744073709551615
            _ = fi11
            _ = v12
        }
    case "\n\t":
        v13 := GenS1{F0: 0}
        _ = v13
    }
    v14 := []int{v10, v10, (v4 << ((492) & 63)), (v4 >> ((18446744073709551615) & 63))}
    for fi15 := 0; fi15 < 12; fi15++ {
        switch fi15 {
        case -783:
            v10 = v10
        case -72:
            v2 = 311
        case 32767:
            v16 := v3[9223372036854775807]
            _ = v16
        default:
            switch -704 {
            case -1:
                fmt.Println(v2)
            }
        }
        v2 -= (v2 | 9223372036854775807)
        v2 -= v2
        _ = fi15
    }
    var v17 uint16 = 42
    fmt.Println(v10, v4)
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 int = v10
        _ = fi18
        _ = v19
    }
    var v20 float32 = 3.14159
    v21 := GenS1{F0: 687}
    fmt.Println(v20, v4)
    fmt.Println(v2)
    v4 = 1
    v4 -= v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
