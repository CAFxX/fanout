package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int32 }
func genFn0() uint8 {
    if (!false) {
        if false {
            v7 := []int{-266}
            _ = v7
        }
    } else {
        var v8 uint32 = 0
        if (false || (!true)) {
            v8 -= (^(v8))
            for fi9 := 0; fi9 < 4; fi9++ {
                for fi10 := 0; fi10 < 6; fi10++ {
                    var v11 int = -9223372036854775808
                    var v12 int = v11
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = fi9
            }
        }
        v8 += 4294967295
        _ = v8
    }
    var v13 float32 = (0.1 - -2.25)
    v14 := GenS1{F0: ("\n\t" + ""), F1: (0 < 0), F2: -2147483648}
    var v15 uint64 = 0
    _ = v13
    _ = v14
    _ = v15
    return 0
}

func main() {
    v1 := []uint32{562, 0, 42, 42}
    v2 := []uint32{42}
    var v3 uint64 = 18446744073709551615
    var v4 int64 = 390
    var v5 uint = 18446744073709551615
    var v6 uint16 = 644
    var v16 float32 = -2.25
    var v17 uint64 = v3
    fmt.Println(v5, v5)
    v17 += 42
    v18 := GenS1{F0: "foo bar", F1: (0 == 1), F2: 2147483647}
    if true {
        if false {
            var v19 uint16 = v6
            _ = v19
        }
        var v20 uint64 = (v3 ^ 501)
        _ = v20
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 -= v17
    fmt.Println(v17, v4, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
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
}
