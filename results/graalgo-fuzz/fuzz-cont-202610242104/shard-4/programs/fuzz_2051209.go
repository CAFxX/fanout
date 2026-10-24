package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int64; F1 uint8 }

func main() {
    var v1 string = (("a" + "") + "\n\t")
    var v2 uint16 = 1
    var v3 int = 2147483647
    v4 := make(map[uint32]uint32)
    for fi5 := 0; fi5 < 16; fi5++ {
        v1 = v1
        v6 := v4[900]
        if (false || false) {
            if true {
                for fi7 := 0; fi7 < 5; fi7++ {
                    var v8 uint32 = uint32(v2)
                    var v9 uint8 = 255
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                v10 := []int{fi5, (^(v3))}
                v11 := GenS1{F0: true, F1: v1}
                _ = v10
                _ = v11
            }
            for fi12 := 0; fi12 < 13; fi12++ {
                fmt.Println(v1, v2)
                _ = fi12
            }
        }
        _ = fi5
        _ = v6
    }
    if (85 > 42) {
        v2 += 32767
        var v13 float32 = -2.25
        v14 := GenS1{F0: false, F1: ""}
        _ = v13
        _ = v14
    } else {
        v15 := GenS1{F0: (127 != 53), F1: v1}
        var v16 uint32 = 2147483647
        var v17 float32 = ((float32(1.5) / (float32(-2.25) * float32(-2.25))) / (-2.25 + 0.0))
        _ = v15
        _ = v16
        _ = v17
    }
    var v18 bool = false
    v19 := v4[2147483647]
    v20 := GenS2{F0: -230, F1: 2}
    v21 := []int{}
    v22 := GenS1{F0: (v19 != v19), F1: ""}
    var v23 uint32 = (v19 ^ 4294967295)
    fmt.Println(v19, v19, v1)
    _ = len(v21)
    v1 = v1
    v19 -= 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
