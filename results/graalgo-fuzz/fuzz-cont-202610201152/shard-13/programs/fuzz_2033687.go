package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint16 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]string)
    v3 := []string{"a"}
    var v4 int16 = 42
    v5 := GenS1{F0: 468, F1: ("\n\t" + "foo bar"), F2: 2147483647}
    fmt.Println(v4)
    var v6 int32 = (^(v1))
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int32 = v1
            _ = fi8
            _ = v9
        }
        v6 = -2147483648
        _ = fi7
    }
    v1 = (-(v1))
    v10 := GenS1{F0: 807, F1: ("a" + "a"), F2: 42}
    var v11 uint16 = 1
    for fi12 := 0; fi12 < 14; fi12++ {
        fmt.Println(fi12, fi12)
        _ = fi12
    }
    switch 42 {
    case 146:
        var v13 uint = 42
        _ = v13
    }
    v2[("foo bar" + ("hello" + "foo bar"))] = "a"
    var v14 int = -633
    v15 := GenS2{F0: false, F1: v11}
    v16 := GenS1{F0: 782, F1: "foo bar", F2: 2147483648}
    v1 = (^(v6))
    if len(v3) > 0 {
        v3[0] = ""
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
