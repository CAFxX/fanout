package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int32; F1 int32 }

func main() {
    var v1 int16 = 1
    var v2 uint32 = 42
    v3 := []uint32{}
    v4 := []string{"hello", "", (("\n\t" + "a") + ("a" + ("a" + "\n\t")))}
    v5 := make(map[string]uint32)
    var v6 int16 = 42
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        v6 += (v6 - (v6 << ((461) & 63)))
    } else {
        v7 := GenS1{F0: true, F1: 127}
        v6 = (v1 * v1)
        v8 := GenS1{F0: (!(493 >= 800)), F1: -128}
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        switch "hello" {
        case "\n\t":
            v1 += (v1 & v6)
        case "a":
            fmt.Println(v1, v6)
        case "":
            v10 := GenS2{F0: 2147483647, F1: 0}
            _ = v10
        }
        _ = fi9
    }
    if true {
        v5["foo bar"] = v2
        if true {
            for fi11 := 0; fi11 < 7; fi11++ {
                v2 = 1
                _ = fi11
            }
            fmt.Println(v1)
            _ = len(v5)
        }
        v1 = (-(v1))
    } else {
        v2 += 2147483648
        v1 = (v1 & v6)
    }
    v6 -= v6
    for fi12 := 0; fi12 < 7; fi12++ {
        if len(v3) > 0 {
            v3[0] = v2
        }
        _ = fi12
    }
    v2 -= 446
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
