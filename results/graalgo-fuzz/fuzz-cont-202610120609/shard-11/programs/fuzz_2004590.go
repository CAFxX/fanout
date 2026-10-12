package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint64 }

func main() {
    var v1 float32 = float32(float64(1.5))
    var v2 uint16 = 32767
    v3 := make(map[string]uint32)
    var v4 uint16 = v2
    var v5 uint16 = v2
    v2 -= (v4 / ((42) | 1))
    v6 := v3["\n\t"]
    v7 := GenS1{F0: (("" + (("a" + ("a" + "foo bar")) + "")) + ""), F1: 183, F2: 18446744073709551615}
    v8 := []int{}
    if (int8(v6) > -34) {
        if true {
            v4 = 32767
            v4 = 255
        } else {
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 float32 = float32((1e10 - float64(0.1)))
                _ = fi9
                _ = v10
            }
        }
        v1 += -2.25
    } else {
        v5 = v4
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v2 = (v2 >> ((479) & 63))
        _ = fi11
    }
    if len(v8) > 0 {
        v8[0] = 839
    }
    if (!false) {
        for fi12 := 0; fi12 < 1; fi12++ {
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 int = fi12
                v15 := GenS1{F0: "", F1: 251, F2: 565}
                _ = fi13
                _ = v14
                _ = v15
            }
            v1 = 3.14159
            _ = fi12
        }
    } else {
        if ((!false) || false) {
            if false {
                v6 += (v6 >> ((255) & 63))
                var v16 int64 = 744
                for fi17 := 0; fi17 < 10; fi17++ {
                    _ = len(v8)
                    fmt.Println(v5)
                    var v18 string = "a"
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
            v4 += v4
        }
        v5 = v2
    }
    v5 = 255
    v5 = (v4 + v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
