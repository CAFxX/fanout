package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    var v3 int16 = 582
    var v4 float64 = 1.5
    var v5 float64 = v4
    v6 := make(map[int]uint32)
    v4 = (0.0 - (float64(((-2.25 - 0.0) * 0.1)) + v5))
    if true {
        switch 2147483647 {
        case 1:
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := v2[155]
                for fi9 := 0; fi9 < 18; fi9++ {
                    var v10 int16 = (v3 ^ v3)
                    var v11 uint8 = 0
                    fmt.Println(v11, fi9)
                    var v12 uint32 = 818
                    switch ((("foo bar" + "\n\t") + "") + "foo bar") {
                    case "hello":
                        v13 := v2[42]
                        _ = v13
                    default:
                        v6[(fi7 - 0)] = v12
                    }
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
                v4 += (v4 + v4)
                _ = fi7
                _ = v8
            }
        }
        v14 := []int{32767, 2147483647, -229, 0}
        _ = v14
    }
    var v15 int = -9223372036854775808
    v16 := []int{1, v15, v15, v15}
    fmt.Println(v3, v15, v15)
    v17 := GenS1{F0: 1, F1: "a"}
    var v18 bool = false
    var v19 uint32 = 757
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
