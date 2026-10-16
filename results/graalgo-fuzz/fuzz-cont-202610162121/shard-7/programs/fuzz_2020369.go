package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int8 }
func genFn0() int8 {
    v4 := []int{-67, 74, 513}
    v5 := GenS1{F0: 9223372036854775807}
    _ = v4
    _ = v5
    return 100
}

func main() {
    v1 := []int{0, 42, -460, 32767}
    v2 := make(map[string]bool)
    var v3 int = 725
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int = (v3 << ((4294967295) & 63))
        if len(v1) > 0 {
            v1[0] = 0
        }
        v8 := GenS1{F0: 9223372036854775807}
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(v3)
    var v9 uint32 = 1
    if len(v1) > 0 {
        v1[0] = (v3 | v3)
    }
    var v10 float64 = 1e10
    for fi11 := 0; fi11 < 1; fi11++ {
        v10 += 0.1
        var v12 int = v3
        _ = fi11
        _ = v12
    }
    var v13 string = "\n\t"
    v9 -= (v9 << ((1) & 63))
    fmt.Println(v13)
    v14 := GenS1{F0: -9223372036854775808}
    switch "a" {
    case "":
        fmt.Println(v10)
    case "a":
        v15 := v2[("foo bar" + "hello")]
        _ = v15
    case "foo bar":
        v10 = v10
    default:
        v10 -= 0.1
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        for fi17 := 0; fi17 < 18; fi17++ {
            v13 = ""
            v18 := GenS1{F0: 42}
            v19 := GenS2{F0: (((-(fi17)) >= (fi16 ^ fi17)) && (465 < 2147483647)), F1: 1}
            _ = fi17
            _ = v18
            _ = v19
        }
        v1 = append(v1, 127)
        _ = fi16
    }
    v20 := GenS2{F0: false, F1: genFn0()}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v9
}
