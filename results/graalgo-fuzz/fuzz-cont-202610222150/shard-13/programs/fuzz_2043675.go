package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 1
    var v3 uint32 = 0
    var v4 int64 = 42
    var v5 uint16 = 241
    var v6 uint8 = 61
    v5 = v5
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint32 = (v2 | 575)
        var v9 uint16 = v5
        var v10 float64 = ((1e10 * 1.5) + 1e10)
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 int8 = -36
    v12 := v1[v2]
    fmt.Println(v12, v6, v3)
    if (1e10 != 0.1) {
        fmt.Println(v2, v6)
        if ((^(v5)) > 2) {
            v13 := []int{1, 0, 128, 143}
            var v14 uint8 = (^(v6))
            v2 = (v3 >> ((355) & 63))
            _ = v13
            _ = v14
        } else {
            var v15 string = (("a" + ("" + (("foo bar" + "") + ("" + "\n\t")))) + "a")
            if v12 {
                var v16 float64 = (0.1 - 0.1)
                v16 += v16
                switch "\n\t" {
                case "foo bar":
                    v16 = 1.5
                default:
                    v17 := []int{32767, -2147483648, 4294967295, 717}
                    _ = v17
                }
                _ = v16
            }
            _ = v15
        }
        fmt.Println(v5, v11)
    }
    var v18 uint32 = (v2 + (v3 << ((2) & 63)))
    for fi19 := 0; fi19 < 13; fi19++ {
        switch "\n\t" {
        case "hello":
            var v20 uint8 = v6
            _ = v20
        case "foo bar":
            for fi21 := 0; fi21 < 2; fi21++ {
                v22 := GenS1{F0: 415}
                fmt.Println(v2, fi19, v11)
                fmt.Println(v18, v11, fi19)
                v23 := GenS2{F0: -1}
                _ = fi21
                _ = v22
                _ = v23
            }
        default:
            v24 := GenS1{F0: 0}
            _ = v24
        }
        var v25 float32 = (0.1 - (1.5 * ((1.5 + -2.25) * (0.1 + (3.14159 + 3.14159)))))
        _ = fi19
        _ = v25
    }
    var v26 uint16 = (v5 | 739)
    v12 = (v12 && v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
