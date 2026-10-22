package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{9223372036854775807, 2147483647, 1}
    var v3 string = ("foo bar" + "foo bar")
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 9; fi5++ {
        v2 = append(v2, (^(fi5)))
        _ = fi5
    }
    switch ((v3 + v3) + v3) {
    case "foo bar":
        for fi6 := 0; fi6 < 4; fi6++ {
            if len(v2) > 0 {
                v2[0] = -9223372036854775808
            }
            var v7 int = fi6
            _ = fi6
            _ = v7
        }
    case "":
        v4[937] = (!(!(!false)))
    case "\n\t":
        v3 = v3
    default:
        var v8 bool = true
        _ = v8
    }
    delete(v4, 0)
    var v9 uint32 = 0
    fmt.Println(v3, v3)
    v10 := GenS1{F0: ((!false) && (!false)), F1: 4294967295, F2: 0}
    var v11 float64 = ((((1.5 + (((((0.1 + ((((1.5 / 1.0) - -2.25) + 0.0) * 1e10)) / (0.1 * -2.25)) - 0.1) * float64(1e10)) - (0.0 - (-2.25 + (((1e10 + ((0.0 * 1.5) - ((float64(((1.5 * (3.14159 + -2.25)) * (1.5 - 1e10))) - 3.14159) + 1e10))) + float64(0.1)) / ((float64(((3.14159 * 0.0) / -2.25)) * 0.0) - (3.14159 / 1.0))))))) / -2.25) - 1.5) + 0.1)
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 uint64 = 1
        _ = fi12
        _ = v13
    }
    fmt.Println(v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v14 int32 = 42
    for fi15 := 0; fi15 < 5; fi15++ {
        fmt.Println(v14, fi15, fi15)
        _ = fi15
    }
    delete(v4, v9)
    var v16 int = -80
    fmt.Println(v11, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
