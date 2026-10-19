package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint = 0
    var v2 uint32 = 2
    var v3 int64 = 9223372036854775807
    v4 := GenS1{F0: 0, F1: 1}
    if false {
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            if ((v1 != (^(v1))) || (986 == 42)) {
                switch 32767 {
                case -2147483648:
                    var v5 uint8 = 42
                    _ = v5
                }
                v6 := GenS2{F0: (^(v1))}
                var v7 uint = 212
                _ = v6
                _ = v7
            } else {
                v3 += v3
                var v8 bool = false
                _ = v8
            }
        case "":
            v9 := []int{-605, 206, -493}
            _ = v9
        case "hello":
            v10 := []int{-227, 127, 660}
            _ = v10
        default:
            v1 -= v1
        }
        v2 = 1
        for fi11 := 0; fi11 < 3; fi11++ {
            for fi12 := 0; fi12 < 18; fi12++ {
                v13 := []int{fi11}
                var v14 uint8 = 94
                switch "" {
                case "":
                    var v15 bool = ((true && (0 > 42)) || true)
                    _ = v15
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi11
        }
    }
    v16 := []int{}
    if (!true) {
        v3 = (-(v3))
        v1 = v1
        if (1000 <= 1) {
            v1 += v1
        } else {
            v2 = (v2 | 2147483647)
        }
    }
    v2 -= v2
    v17 := GenS1{F0: 42, F1: -95}
    v3 = v3
    var v18 float64 = 1e10
    var v19 uint16 = 1
    fmt.Println(v1, v18, v19)
    var v20 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
