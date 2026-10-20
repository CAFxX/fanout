package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int }
func genFn0(p0 uint32) uint8 {
    v5 := []int{}
    v5 = append(v5, -2147483648)
    _ = p0
    _ = v5
    return 75
}
func genFn1(p0 bool, p1 uint) uint16 {
    var v6 uint64 = 4294967295
    v7 := GenS1{F0: 150}
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(v6)
        var v9 string = "\n\t"
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 434
}

func main() {
    v1 := make(map[int]bool)
    var v2 bool = true
    v3 := make(map[string]uint32)
    v4 := []int{9223372036854775807, 9223372036854775807, 32768, -726}
    switch 32768 {
    case 624:
        v2 = (!v2)
    case -2147483648:
        if len(v4) > 0 {
            v4[0] = -662
        }
    }
    switch 1 {
    case -445:
        v4 = append(v4, 42)
    }
    var v10 int16 = 988
    switch -353 {
    case 32768:
        v11 := GenS1{F0: 2147483647}
        _ = v11
    case 4294967295:
        v12 := GenS1{F0: -9223372036854775808}
        _ = v12
    case 0:
        var v13 float64 = (-2.25 + -2.25)
        _ = v13
    }
    var v14 string = ("hello" + ("a" + "a"))
    v15 := GenS2{F0: -1, F1: 4294967295}
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 int = fi16
        _ = fi16
        _ = v17
    }
    var v18 uint16 = 281
    fmt.Println(v10)
    v2 = v2
    fmt.Println(v2, v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
