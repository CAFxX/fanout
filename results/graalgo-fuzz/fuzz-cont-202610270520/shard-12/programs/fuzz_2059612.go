package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint32 }
func genFn0(p0 uint16, p1 int64, p2 int32) int32 {
    var v4 float32 = ((1e10 + (float32(0.1) - 3.14159)) / (1e10 / 3.14159))
    if (402 >= 0) {
        p2 = p2
    } else {
        v5 := GenS1{F0: 789}
        _ = v5
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p2
}

func main() {
    var v1 int32 = -2147483648
    var v2 bool = true
    var v3 uint = 2
    for fi6 := 0; fi6 < 4; fi6++ {
        if (!v2) {
            v7 := GenS2{F0: ("\n\t" + (("hello" + "hello") + "a")), F1: 534}
            _ = v7
        }
        fmt.Println(v3, fi6, v2)
        _ = fi6
    }
    switch -734 {
    case 667:
        v2 = v2
    case -9223372036854775808:
        var v8 uint64 = 18446744073709551615
        _ = v8
    case 2147483647:
        var v9 float32 = 0.1
        _ = v9
    default:
        switch "hello" {
        case "hello":
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := GenS1{F0: 9223372036854775807}
                v3 = 913
                if v2 {
                    var v12 bool = v2
                    v13 := []int{fi10, 882, -9223372036854775808, fi10}
                    _ = v12
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
        case "foo bar":
            var v14 int32 = (v1 | (-(v1)))
            _ = v14
        case "\n\t":
            v2 = false
        default:
            var v15 string = ""
            _ = v15
        }
    }
    v1 += genFn0(255, 42, 1)
    v16 := []int{-19, -291, -1}
    v3 += 18446744073709551615
    if v2 {
        v17 := GenS2{F0: "foo bar", F1: 42}
        _ = v17
    }
    if len(v16) > 0 {
        v16[0] = 42
    }
    v18 := GenS2{F0: "", F1: 855}
    for fi19 := 0; fi19 < 2; fi19++ {
        v1 -= (v1 * (-(v1)))
        _ = fi19
    }
    fmt.Println(v1)
    switch "foo bar" {
    case "hello":
        var v20 int64 = -258
        _ = v20
    case "foo bar":
        v16 = append(v16, -1)
    }
    fmt.Println(v1, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
}
