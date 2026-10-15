package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int64; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint32; F1 uint }
func genFn0(p0 int64, p1 string) int16 {
    p1 = p1
    var v7 int8 = 0
    fmt.Println(v7)
    _ = p0
    _ = p1
    _ = v7
    return 32767
}

func main() {
    var v1 bool = false
    var v2 int8 = -128
    v3 := []string{"\n\t", ("a" + ""), "\n\t"}
    var v4 string = "\n\t"
    var v5 uint = 4294967295
    var v6 float32 = ((-2.25 / 1.0) / (((1.5 * -2.25) + ((1e10 / 1e10) + 0.0)) - ((3.14159 - -2.25) / 0.1)))
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: -1, F1: 1}
        _ = fi8
        _ = v9
    }
    v1 = true
    v10 := GenS1{F0: -1, F1: 0}
    v6 -= 1.5
    for fi11 := 0; fi11 < 16; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            v1 = v1
            _ = fi12
        }
        for fi13 := 0; fi13 < 3; fi13++ {
            if (42 != genFn0(42, "hello")) {
                v1 = (23 < 231)
            }
            v14 := GenS2{F0: -9223372036854775808, F1: v5, F2: 1}
            _ = fi13
            _ = v14
        }
        _ = fi11
    }
    v5 = v5
    if len(v3) > 0 {
        v3[0] = v4
    }
    v15 := GenS3{F0: 65535, F1: v5}
    var v16 uint = (v5 << ((v5) & 63))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v17 float64 = 0.0
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
