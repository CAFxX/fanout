package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint8 }

func main() {
    var v1 string = ("foo bar" + "")
    var v2 uint16 = 42
    var v3 float64 = 0.0
    v4 := make(map[string]string)
    v5 := GenS1{F0: 4294967295, F1: 64, F2: 0}
    var v6 float32 = 1.5
    if false {
        v7 := []int{843, 270, 525}
        v8 := GenS1{F0: 18446744073709551615, F1: 127, F2: 144}
        _ = v7
        _ = v8
    }
    v9 := v4[v1]
    var v10 int8 = -1
    var v11 int16 = 42
    delete(v4, (v9 + v1))
    var v12 int32 = 98
    switch 129 {
    case 1:
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 int16 = (^(v11))
            var v15 uint64 = 709
            for fi16 := 0; fi16 < 10; fi16++ {
                v17 := []int{fi16, (fi16 ^ (fi13 + 1))}
                v3 -= 3.14159
                v11 = v14
                _ = fi16
                _ = v17
            }
            _ = fi13
            _ = v14
            _ = v15
        }
    case -899:
        v18 := GenS1{F0: 18446744073709551615, F1: (v10 >> ((18446744073709551615) & 63)), F2: 220}
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
