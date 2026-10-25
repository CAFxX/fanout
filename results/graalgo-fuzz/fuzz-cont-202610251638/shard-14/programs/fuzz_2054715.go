package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int16; F1 int32 }

func main() {
    var v1 int = -708
    var v2 int = -9223372036854775808
    v3 := []uint32{1, 0, 42}
    v4 := []uint32{1, 65535, 311, 2147483647}
    var v5 int8 = 127
    v6 := []int{(v2 >> ((198) & 63)), (v1 % ((-2147483648) | 1))}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        if true {
            if len(v4) > 0 {
                v4[0] = 65535
            }
        } else {
            if (true || (!(!((32767 >= 32767) && false)))) {
                v8 := GenS1{F0: 1, F1: -333}
                v9 := GenS1{F0: 42, F1: -1}
                _ = v8
                _ = v9
            } else {
                v10 := GenS1{F0: 255, F1: 0}
                fmt.Println(v1, v5)
                _ = v10
            }
            v1 += fi7
        }
        var v11 uint64 = 898
        var v12 int32 = 42
        _ = fi7
        _ = v11
        _ = v12
    }
    if false {
        v1 = v1
        var v13 uint = 4294967295
        if (1 != 2) {
            var v14 float64 = 1e10
            v2 = (v1 * (v2 % ((v2) | 1)))
            var v15 bool = false
            _ = v14
            _ = v15
        }
        _ = v13
    }
    v2 -= v2
    v5 = v5
    v16 := GenS2{F0: "\n\t", F1: v2}
    v1 -= v2
    for fi17 := 0; fi17 < 2; fi17++ {
        v1 += 32767
        _ = fi17
    }
    var v18 uint32 = 1
    fmt.Println(v1, v18, v1)
    switch (v2 + (v2 + (-(v1)))) {
    case 9223372036854775807:
        var v19 string = "hello"
        _ = v19
    case 128:
        v2 -= (v2 & 127)
    }
    v20 := GenS2{F0: "\n\t", F1: v1}
    v21 := GenS3{F0: -1, F1: -1}
    var v22 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
