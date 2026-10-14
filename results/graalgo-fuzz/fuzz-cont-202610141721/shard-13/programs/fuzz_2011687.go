package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 bool }
func genFn0(p0 int, p1 int16) uint16 {
    fmt.Println(p0)
    p0 += p0
    v4 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v4
    return 32767
}

func main() {
    v1 := []int{-9223372036854775808}
    var v2 int = 4294967295
    v3 := []int{(v2 >> ((65535) & 63)), 128, v2}
    var v5 int = v2
    for fi6 := 0; fi6 < 4; fi6++ {
        v2 = -142
        switch "\n\t" {
        case "foo bar":
            v7 := GenS1{F0: 2}
            _ = v7
        default:
            v8 := GenS2{F0: 255, F1: false}
            _ = v8
        }
        v5 = -525
        _ = fi6
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        if true {
            v5 = (v5 << ((255) & 63))
            v10 := GenS1{F0: 65535}
            for fi11 := 0; fi11 < 6; fi11++ {
                v5 -= (fi11 ^ -22)
                var v12 uint64 = 319
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 float64 = -2.25
            fmt.Println(v5, fi13)
            v1 = append(v1, 42)
            _ = fi13
            _ = v14
        }
        _ = fi9
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        if (2147483647 <= -652) {
            fmt.Println(v2, fi15)
        } else {
            var v16 string = ""
            var v17 uint32 = 2147483648
            _ = v16
            _ = v17
        }
        var v18 uint = 18446744073709551615
        v3 = append(v3, (^(fi15)))
        _ = fi15
        _ = v18
    }
    var v19 uint64 = 4294967295
    v2 -= v5
    fmt.Println(v5, v2, v19)
    switch "foo bar" {
    case "a":
        switch v2 {
        case 9223372036854775807:
            if len(v1) > 0 {
                v1[0] = 127
            }
        case -2147483648:
            v5 += v2
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v5
}
