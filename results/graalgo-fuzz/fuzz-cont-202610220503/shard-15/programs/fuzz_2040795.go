package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int32 = -1
    var v2 int16 = 32767
    var v3 uint = 899
    var v4 uint32 = 1
    var v5 int = 127
    var v6 string = ("" + "\n\t")
    v5 = (v5 & -126)
    if (!(((1.5 - 1e10) - 0.1) > 1.5)) {
        v7 := []int{v5, 32768, (-(v5))}
        v3 += v3
        _ = v7
    } else {
        var v8 uint16 = 420
        _ = v8
    }
    v9 := []int{}
    var v10 uint8 = 0
    switch (-(v5)) {
    case 0:
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    v9 = append(v9, (-(v5)))
    switch v6 {
    case "hello":
        for fi11 := 0; fi11 < 8; fi11++ {
            switch -2147483648 {
            case 127:
                v5 += v5
            default:
                var v12 int16 = v2
                _ = v12
            }
            _ = fi11
        }
    case "foo bar":
        v13 := GenS1{F0: v3}
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
