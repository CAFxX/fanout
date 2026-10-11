package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint }

func main() {
    var v1 int = 957
    var v2 uint8 = 1
    v3 := make(map[string]uint32)
    v4 := []string{}
    v5 := []uint32{749, 376, 0, 828}
    v6 := make(map[int]uint32)
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(v1, v2, fi7)
        v8 := GenS1{F0: ("\n\t" + "a"), F1: v1, F2: 18446744073709551615}
        _ = fi7
        _ = v8
    }
    var v9 int = v1
    v2 += 2
    var v10 bool = (!(int16(v2) > 1))
    fmt.Println(v2)
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS1{F0: "", F1: fi11, F2: 982}
            var v14 string = ("hello" + "hello")
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    v15 := GenS1{F0: ("foo bar" + "hello"), F1: (v9 + (v1 << ((2) & 63))), F2: 1}
    var v16 bool = (0.1 == 0.0)
    if len(v4) > 0 {
        v4[0] = "a"
    }
    fmt.Println(v10, v16, v16)
    var v17 string = ("" + "\n\t")
    var v18 int = v1
    if (1.5 > 3.14159) {
        var v19 string = (("foo bar" + (v17 + "a")) + v17)
        _ = v19
    } else {
        for fi20 := 0; fi20 < 16; fi20++ {
            v16 = v16
            _ = fi20
        }
        v16 = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
