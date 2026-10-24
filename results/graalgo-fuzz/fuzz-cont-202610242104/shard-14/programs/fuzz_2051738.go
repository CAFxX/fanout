package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 int32 }

func main() {
    v1 := []string{"\n\t", ("foo bar" + "")}
    var v2 float64 = 1e10
    var v3 int8 = 42
    var v4 int32 = 1
    var v5 float32 = 0.0
    var v6 string = "a"
    var v7 uint64 = 4294967295
    v5 = -2.25
    var v8 uint32 = 751
    if false {
        v9 := GenS1{F0: -128, F1: (v3 - v3), F2: 127}
        v7 = (v7 ^ (v7 % ((v7) | 1)))
        _ = v9
    } else {
        v3 = (v3 - (v3 % (((v3 | (v3 * v3))) | 1)))
        if true {
            v10 := GenS2{F0: (v7 | (^(v7)))}
            if ((71 == (v7 % ((v7) | 1))) && ((0 > 65535) || true)) {
                for fi11 := 0; fi11 < 18; fi11++ {
                    v6 = "\n\t"
                    for fi12 := 0; fi12 < 6; fi12++ {
                        switch fi11 {
                        case -1:
                            var v13 float64 = v2
                            _ = v13
                        default:
                            v14 := GenS3{F0: 255, F1: 0}
                            _ = v14
                        }
                        _ = fi12
                    }
                    _ = fi11
                }
                v3 = 100
            }
            v6 = v6
            _ = v10
        }
        if true {
            for fi15 := 0; fi15 < 19; fi15++ {
                v3 -= -29
                v5 += (v5 - float32(v2))
                _ = fi15
            }
        }
    }
    var v16 int8 = v3
    var v17 uint8 = 127
    var v18 string = v6
    fmt.Println(v18, v8)
    var v19 uint8 = (v17 + (^(v17)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
