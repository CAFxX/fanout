package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
func genFn0(p0 int, p1 uint64, p2 int32) uint64 {
    v7 := []int{p0, (p0 & (p0 * p0)), p0}
    p2 -= (p2 + (p2 / ((p2) | 1)))
    v8 := GenS1{F0: 0, F1: "\n\t"}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p1
}

func main() {
    var v1 int = 32767
    var v2 int64 = 9223372036854775807
    var v3 int = -2147483648
    var v4 uint = 255
    v5 := []int{(^(v1)), -9223372036854775808}
    var v6 int = v1
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := GenS1{F0: 128, F1: ("foo bar" + ("\n\t" + ""))}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v2 += (v2 + (v2 / ((int64(v4)) | 1)))
    var v12 int32 = -707
    v6 = v3
    v13 := GenS1{F0: 1, F1: ("" + "a")}
    var v14 uint32 = 65535
    var v15 uint8 = 0
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 int16 = 32767
        v2 -= v2
        var v18 float32 = (((((1.5 - (float32(1.5) + 1e10)) - -2.25) - 0.0) - 1e10) / -2.25)
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
