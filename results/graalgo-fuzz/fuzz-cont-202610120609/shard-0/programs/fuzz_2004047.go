package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS3 struct { F0 string; F1 int16 }
func genFn0(p0 string) string {
    v7 := GenS1{F0: 255, F1: 348}
    var v8 string = p0
    fmt.Println(p0, v8)
    fmt.Println(p0)
    _ = p0
    _ = v7
    _ = v8
    return "a"
}
func genFn1() uint {
    switch 128 {
    case 200:
        v9 := []int{-2147483648, 2147483647}
        _ = v9
    case 0:
        if (!true) {
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := []int{(-(fi10)), 32768}
                _ = fi10
                _ = v11
            }
            var v12 int = 4294967295
            v12 = 21
            _ = v12
        }
    case -9223372036854775808:
        v13 := GenS2{F0: 255, F1: 238, F2: -1}
        _ = v13
    default:
        v14 := []int{}
        _ = v14
    }
    var v15 int = 32768
    var v16 int = (v15 / ((v15) | 1))
    v17 := GenS2{F0: 0, F1: 2, F2: 944}
    fmt.Println(v15, v15)
    _ = v15
    _ = v16
    _ = v17
    return 18446744073709551615
}

func main() {
    v1 := []uint32{675, 1, 2}
    v2 := make(map[string]string)
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 863
    var v5 int8 = 4
    v6 := []int{}
    var v18 uint = genFn1()
    v19 := GenS3{F0: (("hello" + ("hello" + "")) + "a"), F1: 32767}
    if false {
        if (false || false) {
            if (!false) {
                var v20 int8 = (v5 << ((v18) & 63))
                var v21 int16 = int16(v4)
                v22 := GenS2{F0: 42, F1: 34, F2: 42}
                _ = v20
                _ = v21
                _ = v22
            } else {
                var v23 string = "\n\t"
                fmt.Println(v23, v18, v5)
                v6 = append(v6, 127)
                _ = v23
            }
            var v24 uint16 = 568
            _ = v24
        }
    }
    v18 += (v18 % (((v18 & (v18 & 18446744073709551615))) | 1))
    v5 -= (-(v5))
    if len(v6) > 0 {
        v6[0] = -570
    }
    var v25 int32 = 42
    var v26 int8 = v5
    v26 = v5
    var v27 uint64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
