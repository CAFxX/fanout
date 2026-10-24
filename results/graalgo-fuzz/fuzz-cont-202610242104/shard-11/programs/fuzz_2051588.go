package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint32 }
func genFn0() int16 {
    var v6 bool = ((32767 > 441) && (!(!(!false))))
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(fi7, fi7)
        _ = fi7
    }
    var v8 float64 = 3.14159
    _ = v6
    _ = v8
    return -743
}

func main() {
    v1 := []string{"foo bar", ((("a" + "a") + "a") + "")}
    v2 := make(map[int]string)
    var v3 uint64 = 6
    v4 := make(map[string]uint32)
    var v5 int32 = 0
    _ = len(v1)
    v1 = append(v1, ("" + ("foo bar" + "\n\t")))
    switch "foo bar" {
    case "":
        for fi9 := 0; fi9 < 9; fi9++ {
            v3 = 114
            for fi10 := 0; fi10 < 18; fi10++ {
                v1 = append(v1, "foo bar")
                v3 = 2
                v5 += 2147483647
                _ = fi10
            }
            _ = fi9
        }
    case "a":
        var v11 float32 = ((0.0 - float32(1e10)) * 0.1)
        _ = v11
    case "foo bar":
        v12 := GenS1{F0: 255, F1: 0, F2: 2147483647}
        _ = v12
    }
    v5 = -2147483648
    v3 = (v3 % ((622) | 1))
    v5 = v5
    fmt.Println(v5)
    var v13 bool = (-2147483648 <= v5)
    v14 := v4[(("foo bar" + ("" + ("hello" + "a"))) + "")]
    v15 := GenS1{F0: 65535, F1: 255, F2: 4294967295}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
