package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint32 }
func genFn0(p0 int8, p1 bool, p2 int) int8 {
    p0 -= p0
    if false {
        p0 = (p0 | (p0 | (p0 + -93)))
        p1 = (2147483647 < -2147483648)
    }
    if p1 {
        p1 = (138 <= 486)
        fmt.Println(p2)
    }
    _ = p0
    _ = p1
    _ = p2
    return (-(p0))
}

func main() {
    var v1 int32 = 814
    v2 := []string{""}
    v3 := make(map[string]int)
    for fi4 := 0; fi4 < 7; fi4++ {
        var v5 int = fi4
        fmt.Println(v1)
        fmt.Println(v5, fi4)
        _ = fi4
        _ = v5
    }
    v1 = v1
    v6 := GenS1{F0: (("foo bar" + "a") + ""), F1: 1000, F2: 65535}
    v1 = (v1 & -1)
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint64 = 18446744073709551615
        v8 = 303
        _ = fi7
        _ = v8
    }
    fmt.Println(v1)
    switch 32768 {
    case 4294967295:
        v2 = append(v2, "foo bar")
    case 1:
        for fi9 := 0; fi9 < 9; fi9++ {
            for fi10 := 0; fi10 < 2; fi10++ {
                v1 -= 773
                var v11 uint64 = 2
                _ = fi10
                _ = v11
            }
            v12 := GenS1{F0: "", F1: 1, F2: 42}
            _ = fi9
            _ = v12
        }
    case 574:
        v1 = -2147483648
    default:
        fmt.Println(v1)
    }
    var v13 string = ("\n\t" + "\n\t")
    v14 := v3[("hello" + v13)]
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v6
}
