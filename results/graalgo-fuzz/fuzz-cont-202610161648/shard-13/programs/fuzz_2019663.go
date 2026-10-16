package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 string }
func genFn0(p0 uint32, p1 int16) uint32 {
    var v5 float32 = 1e10
    switch ((("" + ("foo bar" + "\n\t")) + "") + "") {
    case "a":
        var v6 uint8 = 127
        _ = v6
    case "":
        var v7 bool = (false && (false && (((1 >= -1) || ("foo bar" != "a")) && true)))
        _ = v7
    default:
        if (2 > 728) {
            v8 := []int{32768, -780}
            switch -649 {
            case 2147483647:
                var v9 uint64 = 42
                _ = v9
            case 9223372036854775807:
                if len(v8) > 0 {
                    v8[0] = 74
                }
            case -1:
                v8 = append(v8, -366)
            }
            _ = v8
        }
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := GenS1{F0: p0, F1: fi10, F2: "a"}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v5
    return (^(p0))
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    var v12 float64 = -2.25
    v13 := GenS1{F0: 292, F1: 0, F2: "foo bar"}
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 uint = 10
        fmt.Println(fi14)
        _ = fi14
        _ = v15
    }
    var v16 int16 = -32768
    v12 += 0.0
    for fi17 := 0; fi17 < 16; fi17++ {
        v18 := GenS1{F0: 415, F1: fi17, F2: "foo bar"}
        fmt.Println(fi17, v16, v16)
        _ = fi17
        _ = v18
    }
    fmt.Println(v12, v16)
    v12 = v12
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
