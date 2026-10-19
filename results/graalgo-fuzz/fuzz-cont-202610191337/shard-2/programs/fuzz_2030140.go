package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int }
func genFn0(p0 bool, p1 uint32) bool {
    for fi6 := 0; fi6 < 14; fi6++ {
        p0 = (152 > 2)
        var v7 int32 = 1
        v8 := GenS1{F0: 100, F1: 42}
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(p0, p0)
    var v9 uint64 = 42
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint32 = (p1 - (p1 >> ((4294967295) & 63)))
        fmt.Println(v11)
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := []int{}
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    var v14 uint = 42
    _ = p0
    _ = p1
    _ = v14
    _ = v9
    return (9223372036854775807 > 9223372036854775807)
}

func main() {
    v1 := []uint32{}
    var v2 uint16 = 0
    var v3 int32 = 674
    var v4 float32 = 0.0
    var v5 float64 = float64(v4)
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS1{F0: -32, F1: 95}
        var v17 uint16 = (v2 ^ v2)
        _ = fi15
        _ = v16
        _ = v17
    }
    switch 957 {
    case 32768:
        var v18 float32 = float32(v5)
        _ = v18
    }
    v1 = append(v1, 65535)
    switch ("foo bar" + ("foo bar" + ("a" + (("\n\t" + ("hello" + "a")) + "hello")))) {
    case "\n\t":
        v2 += 365
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        if false {
            var v20 uint32 = 726
            _ = v20
        } else {
            v3 = v3
            switch "\n\t" {
            case "":
                v5 = v5
            default:
                v2 += 1
            }
        }
        v2 -= v2
        if len(v1) > 0 {
            v1[0] = 712
        }
        _ = fi19
    }
    if false {
        var v21 int32 = (v3 ^ (v3 >> ((4294967295) & 63)))
        _ = v21
    } else {
        for fi22 := 0; fi22 < 4; fi22++ {
            switch fi22 {
            case 32768:
                v23 := GenS2{F0: fi22}
                _ = v23
            }
            _ = fi22
        }
    }
    v5 -= v5
    v24 := GenS1{F0: 42, F1: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
