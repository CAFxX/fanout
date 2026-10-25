package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint, p1 int8) string {
    v6 := GenS1{F0: (p0 * (p0 - 485))}
    p0 += (p0 - p0)
    _ = p0
    _ = p1
    _ = v6
    return "\n\t"
}
func genFn1() int16 {
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: 677}
        v9 := GenS1{F0: 863}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 bool = false
    var v11 int32 = 780
    _ = v10
    _ = v11
    return -264
}

func main() {
    var v1 uint64 = 4294967295
    var v2 string = "\n\t"
    var v3 uint16 = 2
    var v4 int = -986
    v5 := []string{}
    v5 = append(v5, (v2 + "hello"))
    v12 := GenS2{F0: 65535}
    v13 := GenS2{F0: uint32(v3)}
    var v14 float64 = (1.5 / 1e10)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v1)
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS1{F0: 763}
        v3 = 678
        if false {
            fmt.Println(v2, v14)
        }
        _ = fi15
        _ = v16
    }
    switch v2 {
    case "hello":
        switch v2 {
        case "foo bar":
            fmt.Println(v3, v3)
        default:
            fmt.Println(v2, v4, v14)
        }
    case "foo bar":
        v2 = "a"
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
