package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint32) int {
    v5 := []int{0, 2147483647, 4294967295}
    if len(v5) > 0 {
        v5[0] = 42
    }
    p0 = (p0 ^ p0)
    _ = p0
    _ = v5
    return -827
}

func main() {
    var v1 uint8 = 1
    v2 := []int{}
    var v3 uint8 = (v1 >> ((1) & 63))
    var v4 uint32 = 598
    v3 = 128
    var v6 int16 = 0
    v7 := GenS1{F0: 4294967295, F1: (v3 >> ((679) & 63)), F2: "a"}
    _ = len(v2)
    var v8 uint32 = v4
    v3 = (v1 / (((v3 / (((v1 | (v3 + (v3 * v1)))) | 1))) | 1))
    var v9 int = -2147483648
    fmt.Println(v8, v6, v9)
    v3 = v3
    fmt.Println(v1)
    for fi10 := 0; fi10 < 5; fi10++ {
        switch "\n\t" {
        case "a":
            var v11 float32 = 1e10
            _ = v11
        }
        v1 -= 146
        _ = fi10
    }
    var v12 int8 = 1
    v12 -= int8(v6)
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 uint64 = 18446744073709551615
        v8 = v8
        v6 = -230
        _ = fi13
        _ = v14
    }
    v8 += 686
    if false {
        v15 := GenS1{F0: 42, F1: (v3 - v1), F2: ""}
        v6 -= (v6 - v6)
        _ = v15
    }
    if false {
        var v16 float64 = 1e10
        for fi17 := 0; fi17 < 6; fi17++ {
            v18 := GenS2{F0: v6}
            for fi19 := 0; fi19 < 20; fi19++ {
                if (!true) {
                    v12 -= (^(v12))
                    if true {
                        if (!true) {
                            v3 = v3
                            var v20 bool = true
                            _ = v20
                        }
                        v21 := GenS3{F0: 2, F1: 32767}
                        _ = v21
                    }
                    v12 += (v12 >> ((65535) & 63))
                }
                var v22 float64 = (((v16 - 1e10) + v16) / 1.5)
                v16 = v22
                _ = fi19
                _ = v22
            }
            _ = fi17
            _ = v18
        }
        v3 = (v1 | (^(v1)))
        _ = v16
    }
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
