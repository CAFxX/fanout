package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int16; F2 uint16 }
type GenS3 struct { F0 int; F1 string }
func genFn0() int64 {
    var v6 int16 = 916
    if true {
        if (!(42 > 1)) {
            v6 = v6
        }
    } else {
        v6 = v6
        fmt.Println(v6)
    }
    var v7 uint = 18446744073709551615
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint16 = 344
    var v2 uint32 = 0
    var v3 int64 = 0
    v4 := []int{-1, 128, 88}
    var v5 string = "hello"
    if ((v3 & genFn0()) < -1) {
        v3 -= (v3 ^ (v3 - genFn0()))
    } else {
        v1 += (v1 - (v1 & v1))
    }
    v8 := GenS1{F0: 100}
    if false {
        if true {
            v9 := GenS2{F0: (v3 * (-(v3))), F1: -1, F2: 913}
            v5 = v5
            _ = v9
        } else {
            for fi10 := 0; fi10 < 14; fi10++ {
                var v11 int16 = -556
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi10
                _ = v11
            }
            if (42 <= -128) {
                var v12 int = 4294967295
                v2 = (v2 << ((0) & 63))
                _ = v12
            } else {
                v13 := GenS2{F0: 9223372036854775807, F1: 0, F2: uint16(v2)}
                v1 += v1
                _ = v13
            }
        }
        var v14 uint64 = 18446744073709551615
        var v15 int64 = (v3 % (((v3 + genFn0())) | 1))
        _ = v14
        _ = v15
    } else {
        var v16 int = 32767
        for fi17 := 0; fi17 < 10; fi17++ {
            _ = len(v4)
            v5 = v5
            v2 = 2147483647
            _ = fi17
        }
        _ = v16
    }
    fmt.Println(v1, v3)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    if true {
        if (v5 == v5) {
            var v18 uint16 = 0
            fmt.Println(v3)
            v19 := GenS1{F0: 127}
            _ = v18
            _ = v19
        }
        if true {
            v3 = (v3 ^ -9223372036854775808)
            fmt.Println(v2)
            var v20 int = 0
            _ = v20
        }
        v21 := GenS3{F0: 907, F1: "a"}
        _ = v21
    }
    var v22 float64 = (3.14159 / 0.1)
    if len(v4) > 0 {
        v4[0] = 128
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
