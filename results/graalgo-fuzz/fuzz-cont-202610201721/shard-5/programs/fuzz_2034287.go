package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int16; F1 uint32 }
func genFn0(p0 int32) int32 {
    var v7 int = -494
    var v8 bool = true
    v9 := GenS1{F0: 734, F1: v8}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]string)
    v3 := []uint32{2147483648, 883}
    var v4 int32 = 0
    v5 := []int{42, -580, 0}
    v6 := []int{2147483647, 2147483647, 392, 32767}
    v4 += (-(v4))
    v2[2147483647] = "foo bar"
    for fi10 := 0; fi10 < 19; fi10++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (-193 < -456) {
            if (100 >= -128) {
                v4 += (^(v4))
                v1 = (v1 & 1)
            }
        }
        _ = fi10
    }
    v1 = (v1 ^ (v1 % (((^(v1))) | 1)))
    var v11 bool = true
    v12 := GenS2{F0: -32768, F1: 2147483647}
    v1 = v1
    delete(v2, 1)
    for fi13 := 0; fi13 < 12; fi13++ {
        v11 = true
        switch fi13 {
        case 4294967295:
            delete(v2, 65535)
        case 32768:
            v11 = false
        default:
            if len(v3) > 0 {
                v3[0] = 1
            }
        }
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
