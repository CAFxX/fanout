package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }

func main() {
    v1 := []uint32{1, 913, 4294967295}
    var v2 uint64 = 1
    v3 := []string{"foo bar", "", "", "hello"}
    v4 := []int{32767, 2147483647, 617}
    fmt.Println(v2)
    var v5 float64 = (-2.25 * (-2.25 - -2.25))
    if len(v3) > 0 {
        v3[0] = ("\n\t" + ("a" + ("\n\t" + "a")))
    }
    fmt.Println(v2)
    v6 := GenS1{F0: 78, F1: 1}
    switch (("a" + "hello") + "a") {
    case "":
        for fi7 := 0; fi7 < 15; fi7++ {
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            var v8 uint32 = 860
            v9 := GenS1{F0: 799, F1: 0}
            _ = fi7
            _ = v8
            _ = v9
        }
    default:
        var v10 int8 = -128
        _ = v10
    }
    fmt.Println(v5, v5)
    var v11 float64 = v5
    var v12 uint32 = 910
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
