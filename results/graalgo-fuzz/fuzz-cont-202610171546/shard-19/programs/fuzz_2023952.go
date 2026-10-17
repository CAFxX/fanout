package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint32 }
type GenS2 struct { F0 string; F1 uint32 }
func genFn0(p0 int32) uint32 {
    var v7 string = "hello"
    fmt.Println(p0, v7)
    _ = p0
    _ = v7
    return 480
}
func genFn1(p0 int32) int64 {
    p0 = -1
    if ((1.5 <= 1.5) && (42 == 4294967295)) {
        switch 1 {
        case -1:
            var v8 uint16 = 255
            _ = v8
        default:
            v9 := GenS1{F0: 2147483647, F1: "\n\t", F2: 2147483647}
            _ = v9
        }
        p0 -= p0
    }
    if (42 >= 255) {
        fmt.Println(p0)
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int8 = -49
            v11 -= -128
            if ((3.14159 - -2.25) > 0.1) {
                v11 = v11
                v12 := GenS1{F0: 0, F1: "", F2: 338}
                _ = v12
            } else {
                v13 := GenS1{F0: 2147483647, F1: (((("\n\t" + "\n\t") + "") + "a") + "hello"), F2: 1}
                var v14 uint = 411
                if (true || (true || (fi10 != fi10))) {
                    var v15 int = 4294967295
                    _ = v15
                } else {
                    var v16 uint64 = 42
                    v17 := GenS1{F0: (p0 & (p0 / ((935) | 1))), F1: "a", F2: 65535}
                    _ = v16
                    _ = v17
                }
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 1
    v2 := []uint32{}
    var v3 int = 32767
    var v4 int16 = 1000
    var v5 uint16 = 2
    v6 := make(map[int]uint32)
    var v18 int32 = 1
    if true {
        if (2 >= uint(v5)) {
            var v19 int64 = 271
            if (false && false) {
                v1 = (v1 + v1)
                v20 := GenS2{F0: ("a" + "foo bar"), F1: 0}
                _ = v20
            }
            _ = v19
        }
    } else {
        switch v3 {
        case -9223372036854775808:
            for fi21 := 0; fi21 < 11; fi21++ {
                v6[(v3 / ((fi21) | 1))] = genFn0((v18 << ((76) & 63)))
                fmt.Println(v4)
                _ = fi21
            }
        case 127:
            v1 = v1
        default:
            if (-2.25 > float32(1.5)) {
                v3 -= (-(v3))
            } else {
                var v22 int16 = -32768
                v3 += v3
                for fi23 := 0; fi23 < 13; fi23++ {
                    v5 += (v5 % ((v5) | 1))
                    v24 := GenS1{F0: v18, F1: ("\n\t" + "a"), F2: 26}
                    _ = fi23
                    _ = v24
                }
                _ = v22
            }
        }
        switch "foo bar" {
        case "foo bar":
            var v25 int8 = 127
            _ = v25
        }
        if (!(!(v18 == (v18 & (v18 % (((v18 % (((v18 / (((-(v18))) | 1))) | 1))) | 1)))))) {
            var v26 string = "a"
            var v27 uint64 = 42
            var v28 int8 = 42
            _ = v26
            _ = v27
            _ = v28
        }
    }
    fmt.Println(v4, v3, v4)
    var v29 string = "a"
    v1 -= v1
    _ = len(v2)
    v29 = (("hello" + (v29 + (v29 + v29))) + (v29 + "hello"))
    v5 = (v5 << ((0) & 63))
    _ = len(v2)
    if len(v2) > 0 {
        v2[0] = 2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
