package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int8 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]int)
    var v3 float32 = ((-2.25 - 1e10) * 1e10)
    var v4 uint64 = 4294967295
    var v5 int32 = 922
    v6 := GenS1{F0: 42, F1: (^(v4)), F2: 42}
    if (!false) {
        v5 -= (-(v5))
    }
    var v7 string = "a"
    var v8 int8 = -128
    v9 := []int{32767, 1}
    if true {
        switch (((v7 + v7) + v7) + "a") {
        case "\n\t":
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
    }
    switch 620 {
    case -765:
        v10 := GenS2{F0: 9, F1: (v4 >> ((18446744073709551615) & 63)), F2: 325}
        _ = v10
    }
    v11 := GenS2{F0: -9223372036854775808, F1: v4, F2: (^(v4))}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    var v12 int = 337
    if false {
        var v13 uint64 = (v4 - (v4 << ((v1) & 63)))
        if false {
            v14 := GenS3{F0: v8}
            v4 += (v4 * (v13 & v4))
            _ = v14
        } else {
            var v15 int32 = (^(v5))
            var v16 int16 = -117
            v8 += (v8 & -28)
            _ = v15
            _ = v16
        }
        _ = v13
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 uint8 = 128
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    v4 = v4
    v1 -= v1
    var v20 uint16 = 32767
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
