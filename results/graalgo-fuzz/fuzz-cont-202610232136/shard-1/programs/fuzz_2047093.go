package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS2 struct { F0 string; F1 uint64; F2 int32 }
func genFn0(p0 uint32) int8 {
    var v5 uint32 = (p0 ^ (p0 | 786))
    v6 := []int{-573}
    _ = p0
    _ = v5
    _ = v6
    return -128
}
func genFn1(p0 uint64) uint16 {
    switch "a" {
    case "":
        v7 := GenS1{F0: 4294967295, F1: -1, F2: (p0 - (p0 / ((178) | 1)))}
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 5; fi8++ {
            if (4294967295 != 65535) {
                for fi9 := 0; fi9 < 17; fi9++ {
                    var v10 int64 = -1
                    v10 -= (v10 * (^(v10)))
                    _ = fi9
                    _ = v10
                }
                for fi11 := 0; fi11 < 4; fi11++ {
                    var v12 bool = true
                    _ = fi11
                    _ = v12
                }
            }
            _ = fi8
        }
    }
    p0 = (p0 >> ((65535) & 63))
    p0 += p0
    v13 := GenS2{F0: ("foo bar" + ("a" + ((("\n\t" + "a") + "") + ("" + ("a" + "a"))))), F1: 1, F2: -2147483648}
    _ = p0
    _ = v13
    return 1
}

func main() {
    var v1 int64 = 0
    var v2 float32 = -2.25
    var v3 uint16 = 42
    var v4 int = 4294967295
    v4 -= v4
    var v14 uint = 0
    v1 = v1
    var v15 int16 = 32767
    v1 += v1
    v2 -= 0.1
    fmt.Println(v3, v14, v15)
    v16 := []int{(v4 & (v4 % (((v4 - v4)) | 1))), v4}
    v3 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
