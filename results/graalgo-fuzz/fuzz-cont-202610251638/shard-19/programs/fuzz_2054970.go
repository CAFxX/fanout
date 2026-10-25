package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint }
type GenS2 struct { F0 uint8; F1 string; F2 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 0
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    var v5 uint16 = 32767
    v6 := []int{67, -2147483648, 866, -1}
    v7 := GenS1{F0: v2, F1: 127, F2: 42}
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            if (false && true) {
                v2 = v2
                var v10 int16 = -1
                switch "" {
                case "a":
                    var v11 float64 = (float64(float32(1e10)) - 0.0)
                    _ = v11
                case "\n\t":
                    var v12 float64 = 1.5
                    _ = v12
                case "foo bar":
                    fmt.Println(v5, fi9)
                default:
                    v5 = v5
                }
                _ = v10
            }
            if true {
                var v13 uint16 = 2
                _ = v13
            }
            _ = len(v4)
            _ = fi9
        }
        var v14 string = "\n\t"
        if len(v6) > 0 {
            v6[0] = fi8
        }
        _ = fi8
        _ = v14
    }
    var v15 int32 = -809
    switch "" {
    case "foo bar":
        v5 = (^(v5))
    case "":
        if (false || false) {
            for fi16 := 0; fi16 < 5; fi16++ {
                v17 := GenS2{F0: 0, F1: ("hello" + ("\n\t" + "\n\t")), F2: v15}
                if len(v6) > 0 {
                    v6[0] = (fi16 >> ((1) & 63))
                }
                v2 -= v2
                _ = fi16
                _ = v17
            }
        } else {
            if true {
                _ = len(v6)
            }
            v5 = (v5 * v5)
        }
    case "a":
        delete(v1, 1)
    }
    v18 := v1[uint32(v5)]
    v19 := GenS1{F0: v2, F1: 2, F2: 1}
    if false {
        var v20 uint64 = 922
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
