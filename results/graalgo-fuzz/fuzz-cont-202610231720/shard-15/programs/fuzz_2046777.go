package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
func genFn0() uint {
    var v6 float32 = (3.14159 - ((1e10 + 3.14159) / -2.25))
    v6 -= v6
    var v7 uint32 = 4294967295
    fmt.Println(v7, v6)
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1(p0 int32) uint16 {
    var v8 uint64 = 42
    var v9 uint8 = 140
    _ = p0
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int64 = 702
    v2 := []uint32{4294967295}
    v3 := []uint32{65535, 117, 2147483647}
    v4 := []string{"a", "hello", "\n\t", "\n\t"}
    var v5 int = 1
    var v10 uint16 = genFn1(-673)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v10 = v10
    for fi11 := 0; fi11 < 8; fi11++ {
        if true {
            var v12 int16 = 164
            var v13 int64 = v1
            var v14 uint = 4294967295
            _ = v12
            _ = v13
            _ = v14
        } else {
            v15 := GenS1{F0: 128}
            _ = v15
        }
        _ = fi11
    }
    var v16 uint = 129
    v10 -= (v10 % ((v10) | 1))
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 float32 = float32(3.14159)
        if (!((v10 * (v10 + v10)) > v10)) {
            _ = len(v2)
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 uint16 = (^(v10))
                v18 += ((((((0.0 - (((1e10 - v18) + -2.25) * -2.25)) * 0.0) * v18) / v18) - v18) / float32(float64(1.5)))
                fmt.Println(fi19, v16)
                _ = fi19
                _ = v20
            }
            var v21 uint8 = 1
            _ = v21
        } else {
            v22 := GenS2{F0: "hello"}
            _ = v22
        }
        _ = fi17
        _ = v18
    }
    var v23 uint = (v16 >> ((255) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
