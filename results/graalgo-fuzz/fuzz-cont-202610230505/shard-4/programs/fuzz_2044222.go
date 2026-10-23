package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 string }

func main() {
    var v1 uint16 = 813
    var v2 float32 = 1.5
    v3 := make(map[int]uint32)
    var v4 int64 = 1
    var v5 uint16 = (v1 | v1)
    var v6 int8 = 127
    var v7 uint16 = v5
    v8 := GenS1{F0: 1, F1: (true || (2 < 2)), F2: "a"}
    var v9 int = 2147483647
    v10 := GenS1{F0: 0, F1: (!((255 == 42) || (!false))), F2: ("foo bar" + "a")}
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 int64 = v4
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 uint32 = 2147483648
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := GenS1{F0: 1, F1: (!true), F2: ("foo bar" + ("\n\t" + ("a" + "foo bar")))}
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    var v17 int16 = 1
    v18 := GenS1{F0: 0, F1: true, F2: "foo bar"}
    v5 -= (v7 ^ v7)
    v17 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
