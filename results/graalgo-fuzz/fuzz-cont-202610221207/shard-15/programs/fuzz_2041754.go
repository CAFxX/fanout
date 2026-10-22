package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int32 }
type GenS2 struct { F0 bool; F1 int; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 uint16; F2 int64 }

func main() {
    var v1 int32 = 42
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 255
    v5 := GenS1{F0: (v1 << ((327) & 63)), F1: false, F2: -1}
    _ = len(v3)
    v1 = 2147483647
    v3[2147483647] = 42
    if false {
        v6 := []int{}
        v4 = 253
        v4 -= v4
        _ = v6
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 float64 = float64(-2.25)
        _ = fi7
        _ = v8
    }
    fmt.Println(v1, v1)
    v9 := []int{-9223372036854775808}
    v4 = v4
    fmt.Println(v1, v1)
    v10 := v2[-1]
    switch (v10 + 42) {
    case -466:
        switch ("\n\t" + "hello") {
        case "":
            var v11 string = "a"
            _ = v11
        case "\n\t":
            if ((42 > 138) && true) {
                v9 = append(v9, (v10 / (((v10 % ((v10) | 1))) | 1)))
                v12 := v3[2147483647]
                _ = v12
            } else {
                v1 -= (v1 + int32(v4))
                v13 := v3[42]
                _ = v13
            }
        default:
            v14 := GenS2{F0: true, F1: (v10 * -73), F2: (v10 + 128)}
            _ = v14
        }
    case -129:
        v15 := GenS3{F0: 255}
        _ = v15
    default:
        var v16 bool = ((v4 * (v4 * v4)) == 127)
        _ = v16
    }
    v1 = (^(v1))
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := GenS4{F0: 1, F1: 42, F2: 42}
        var v19 int8 = 115
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
