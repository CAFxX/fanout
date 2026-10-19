package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int64; F2 int32 }
func genFn0(p0 uint32) uint8 {
    var v7 uint = 3
    v8 := []int{-146, 127}
    _ = p0
    _ = v7
    _ = v8
    return 128
}

func main() {
    v1 := []uint32{0, 65535, 1}
    v2 := make(map[string]uint32)
    var v3 uint16 = 162
    var v4 int64 = 0
    var v5 int = 42
    var v6 string = "hello"
    if false {
        v6 = "foo bar"
        v6 = v6
        v5 = (v5 / ((v5) | 1))
    }
    v9 := GenS1{F0: false}
    if true {
        for fi10 := 0; fi10 < 10; fi10++ {
            v4 += 938
            for fi11 := 0; fi11 < 10; fi11++ {
                v12 := GenS1{F0: false}
                if true {
                    switch (fi10 << ((255) & 63)) {
                    case -1:
                        v13 := GenS1{F0: false}
                        v6 = (("\n\t" + v6) + (v6 + ((v6 + (v6 + v6)) + v6)))
                        _ = v13
                    default:
                        if (true && (-704 > 1)) {
                            v3 = v3
                            var v14 int16 = 42
                            _ = v14
                        }
                    }
                }
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
        v3 += 42
        v15 := GenS1{F0: false}
        _ = v15
    }
    v5 += v5
    v16 := GenS2{F0: true, F1: -200, F2: 179}
    v17 := v2[v6]
    var v18 uint16 = v3
    v3 += 488
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
