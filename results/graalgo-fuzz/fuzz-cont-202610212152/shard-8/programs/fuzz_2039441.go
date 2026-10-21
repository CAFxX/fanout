package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (2 > 42)
    var v3 int64 = 9223372036854775807
    var v4 int8 = 0
    v5 := GenS1{F0: 1, F1: 126, F2: 1}
    fmt.Println(v2, v3)
    v6 := []int{int(v4)}
    v6 = append(v6, 128)
    var v7 int32 = 1
    if (42 >= 127) {
        var v8 float64 = 3.14159
        _ = len(v1)
        var v9 uint = 587
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := GenS1{F0: (-(v4)), F1: (-(v4)), F2: 51}
        v2 = v2
        _ = fi10
        _ = v11
    }
    fmt.Println(v3)
    v12 := GenS2{F0: 255}
    switch ("a" + "\n\t") {
    case "\n\t":
        for fi13 := 0; fi13 < 15; fi13++ {
            v1[""] = 0
            if ((fi13 % (((fi13 % ((fi13) | 1))) | 1)) >= 0) {
                for fi14 := 0; fi14 < 11; fi14++ {
                    if (!v2) {
                        v1["hello"] = 2147483648
                    }
                    var v15 uint64 = 42
                    v2 = (1 < 255)
                    _ = fi14
                    _ = v15
                }
                var v16 uint32 = 2147483648
                var v17 float64 = 3.14159
                _ = v16
                _ = v17
            }
            v18 := GenS2{F0: 255}
            _ = fi13
            _ = v18
        }
    }
    fmt.Println(v2, v2)
    for fi19 := 0; fi19 < 15; fi19++ {
        v6 = append(v6, (fi19 % (((^(fi19))) | 1)))
        _ = fi19
    }
    fmt.Println(v4, v3)
    var v20 float64 = -2.25
    if (v2 && ((v4 / ((v4) | 1)) <= (v4 / ((v4) | 1)))) {
        if ((!false) || ((v3 & (v3 % ((-1) | 1))) <= (v3 / ((9223372036854775807) | 1)))) {
            v21 := GenS1{F0: (v4 * (v4 + -128)), F1: 100, F2: 1}
            var v22 uint16 = 1
            _ = v21
            _ = v22
        }
        v20 -= (float64((3.14159 / ((((0.1 / -2.25) - -2.25) / 1.5) + 0.1))) + (v20 * v20))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
