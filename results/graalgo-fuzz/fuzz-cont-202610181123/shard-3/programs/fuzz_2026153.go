package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 string }
type GenS2 struct { F0 int16; F1 uint }

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 4294967295
    var v3 float32 = 0.0
    v4 := GenS1{F0: 32767, F1: 4294967295, F2: ("a" + (("" + "\n\t") + "foo bar"))}
    v5 := GenS1{F0: -967, F1: 626, F2: "hello"}
    v3 += float32(-2.25)
    var v6 uint16 = 65535
    fmt.Println(v3, v2, v2)
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(fi7)
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 int32 = -1
            _ = fi8
            _ = v9
        }
        v2 = (v2 | v2)
        _ = fi7
    }
    var v10 uint32 = 425
    var v11 uint16 = (v6 & 175)
    var v12 int8 = -1
    v6 = v6
    if false {
        for fi13 := 0; fi13 < 13; fi13++ {
            v10 = (^(v10))
            v14 := []int{}
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 float64 = 0.0
            var v17 bool = true
            if (v17 || v17) {
                v18 := v1[fi15]
                _ = v18
            } else {
                for fi19 := 0; fi19 < 15; fi19++ {
                    v20 := GenS2{F0: -560, F1: 4294967295}
                    _ = fi19
                    _ = v20
                }
            }
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
