package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int16 }
func genFn0() int64 {
    v6 := []int{9223372036854775807}
    v7 := GenS1{F0: (!(((-2.25 / 1.0) - ((0.0 + ((0.0 / (float64(((3.14159 / (float32((1e10 + ((1e10 / 1e10) - 1.5))) + (3.14159 * 1e10))) - (float32(float64(float32(float64((1.5 + 1.5))))) - 0.0))) + (1.5 - float64((3.14159 + float32(3.14159)))))) - (float64((0.0 * (float32(float64((0.0 * float32(1.5)))) - (-2.25 - 1.5)))) + float64(0.1)))) / 0.1)) != (3.14159 + 1e10)))}
    var v8 bool = true
    _ = v6
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    v1 := []uint32{2147483647, 4294967295}
    v2 := make(map[int]uint32)
    var v3 int8 = 42
    v4 := []string{("\n\t" + (("hello" + "") + ((("" + "a") + "hello") + ("" + ("hello" + "a"))))), "foo bar", ("hello" + "hello")}
    var v5 uint16 = 32767
    for fi9 := 0; fi9 < 9; fi9++ {
        switch (fi9 << ((42) & 63)) {
        case 870:
            fmt.Println(v3, fi9, v5)
        case 1:
            if (!false) {
                v3 = (v3 / (((v3 ^ (-(v3)))) | 1))
            } else {
                var v10 uint16 = 42
                fmt.Println(v5)
                v10 += (v10 / ((633) | 1))
                _ = v10
            }
        }
        v3 += (v3 % ((-1) | 1))
        _ = fi9
    }
    v5 = (v5 / ((v5) | 1))
    v5 = v5
    var v11 uint64 = 2
    _ = len(v1)
    var v12 int64 = 463
    v13 := GenS1{F0: false}
    v14 := GenS2{F0: 243, F1: 1}
    v15 := GenS1{F0: ((-(v3)) != (v3 / (((v3 >> ((478) & 63))) | 1)))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
