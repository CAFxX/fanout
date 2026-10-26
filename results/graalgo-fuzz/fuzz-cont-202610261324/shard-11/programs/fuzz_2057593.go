package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int8, p1 bool, p2 int32) bool {
    p1 = p1
    v7 := GenS1{F0: 593}
    p1 = ((-(p2)) < 1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (0 > 0)
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{}
    var v3 uint32 = 2147483647
    var v4 int16 = 32767
    v5 := make(map[int]string)
    var v6 float32 = 3.14159
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS2{F0: (fi8 ^ (fi8 - (-(fi8)))), F1: "a"}
        _ = fi8
        _ = v9
    }
    v4 = -32768
    switch -552 {
    case 42:
        var v10 int32 = 245
        _ = v10
    case 4294967295:
        v1["\n\t"] = (!true)
    case 859:
        var v11 float32 = v6
        _ = v11
    default:
        v3 = v3
    }
    var v12 int = 32767
    v13 := GenS1{F0: v4}
    if (65535 < v3) {
        for fi14 := 0; fi14 < 9; fi14++ {
            delete(v5, -9223372036854775808)
            v15 := GenS3{F0: 222}
            fmt.Println(v6)
            _ = fi14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        v17 := GenS4{F0: 1}
        if (3.14159 >= 1.5) {
            var v18 uint32 = v3
            v5[127] = "hello"
            v18 += 605
            _ = v18
        }
        _ = fi16
        _ = v17
    }
    v19 := GenS3{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
