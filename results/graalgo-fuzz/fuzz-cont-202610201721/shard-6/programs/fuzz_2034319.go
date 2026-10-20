package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }

func main() {
    v1 := []int{-111, -30, 2147483647}
    v2 := make(map[uint32]bool)
    var v3 int32 = -1
    var v4 int = -501
    v3 = 42
    fmt.Println(v4, v4)
    var v5 float32 = 1e10
    var v6 float32 = float32((1e10 / (-2.25 * float64((0.0 + ((v5 / float32(3.14159)) + 0.0))))))
    var v7 int8 = -1
    v8 := GenS1{F0: 1}
    _ = len(v1)
    v4 += v4
    _ = len(v1)
    for fi9 := 0; fi9 < 18; fi9++ {
        v5 -= v6
        v10 := GenS1{F0: 0}
        _ = fi9
        _ = v10
    }
    _ = len(v2)
    if true {
        var v11 int8 = (^(v7))
        switch 669 {
        case 42:
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 bool = false
                _ = fi12
                _ = v13
            }
        case 9223372036854775807:
            switch (-(v4)) {
            case 127:
                if (341 == 1) {
                    v4 = (v4 / ((v4) | 1))
                    v14 := GenS1{F0: 4294967295}
                    _ = v14
                }
            }
        default:
            v15 := GenS1{F0: 2}
            _ = v15
        }
        _ = v11
    }
    fmt.Println(v5, v3, v6)
    var v16 float64 = 3.14159
    fmt.Println(v16, v5)
    if (false && false) {
        v6 -= ((float32((-2.25 * 0.1)) - (v6 - v6)) + 0.1)
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 15; fi18++ {
                v6 += (v6 - 0.1)
                if false {
                    v3 -= (v3 - (v3 & (^(v3))))
                    v19 := GenS2{F0: false}
                    _ = v19
                } else {
                    v5 = 3.14159
                    v3 = (v3 & v3)
                }
                _ = fi18
            }
            _ = fi17
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
