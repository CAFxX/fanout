package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 bool; F1 bool }
func genFn0(p0 int8, p1 int16) bool {
    p1 += (p1 << ((18446744073709551615) & 63))
    fmt.Println(p1)
    if (true && (0.1 <= (1e10 / (1e10 + 0.1)))) {
        p0 = p0
        var v4 int32 = -2147483648
        for fi5 := 0; fi5 < 19; fi5++ {
            if true {
                for fi6 := 0; fi6 < 11; fi6++ {
                    if true {
                        var v7 int64 = -765
                        v4 += (v4 ^ v4)
                        _ = v7
                    }
                    fmt.Println(p0, v4, p1)
                    v4 = v4
                    _ = fi6
                }
            }
            var v8 int8 = int8(v4)
            _ = fi5
            _ = v8
        }
        _ = v4
    }
    _ = p0
    _ = p1
    return (!true)
}
func genFn1(p0 int16, p1 uint16) int32 {
    var v9 string = (((("\n\t" + ("a" + ("" + ("\n\t" + "hello")))) + "\n\t") + "foo bar") + "")
    p0 -= 470
    v10 := []int{-2147483648, -9223372036854775808, 127, -178}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return -431
}

func main() {
    v1 := []uint32{65535, 2147483647}
    var v2 uint = 65535
    v3 := []int{169, 2147483647, 32768}
    v11 := GenS1{F0: -32768, F1: 71, F2: 17}
    v2 = 65535
    var v12 bool = ((!false) || (3.14159 > float64(0.0)))
    fmt.Println(v2)
    v13 := GenS2{F0: (^(v2)), F1: 42}
    fmt.Println(v2)
    v2 -= v2
    var v14 uint32 = 4294967295
    var v15 float32 = (((1.5 * 1e10) / (0.1 + (1e10 + 1e10))) - 1.5)
    v16 := GenS3{F0: true, F1: true}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
