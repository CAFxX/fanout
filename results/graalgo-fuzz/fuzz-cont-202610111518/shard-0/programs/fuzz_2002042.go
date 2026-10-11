package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
func genFn0(p0 string, p1 uint32) uint32 {
    p1 = (p1 / ((p1) | 1))
    v7 := GenS1{F0: 4294967295, F1: 1}
    _ = p0
    _ = p1
    _ = v7
    return (p1 & (p1 + (p1 | 4294967295)))
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    var v3 uint8 = 183
    v4 := make(map[int]bool)
    v5 := make(map[string]string)
    var v6 uint8 = v3
    switch 128 {
    case -414:
        if false {
            if false {
                var v8 string = ((("\n\t" + "") + ("foo bar" + ("" + "hello"))) + "a")
                v3 = v6
                v9 := GenS1{F0: 4294967295, F1: v6}
                _ = v8
                _ = v9
            } else {
                v10 := GenS2{F0: 2}
                var v11 uint32 = 2147483647
                for fi12 := 0; fi12 < 6; fi12++ {
                    v13 := GenS2{F0: 828}
                    var v14 uint16 = 791
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = v10
                _ = v11
            }
            for fi15 := 0; fi15 < 14; fi15++ {
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 float32 = ((1.5 * 1e10) - -2.25)
                    _ = fi16
                    _ = v17
                }
                _ = fi15
            }
        }
    default:
        fmt.Println(v6, v3)
    }
    v6 += 42
    switch -2147483648 {
    case 127:
        fmt.Println(v6, v6)
    }
    v6 -= (v3 * 41)
    for fi18 := 0; fi18 < 3; fi18++ {
        _ = len(v1)
        v19 := GenS3{F0: (-(fi18))}
        _ = fi18
        _ = v19
    }
    v1["hello"] = genFn0("a", genFn0("foo bar", 1))
    v3 = (v6 % (((v3 - v6)) | 1))
    fmt.Println(v6, v6)
    fmt.Println(v6)
    v20 := GenS1{F0: 333, F1: (v6 ^ 2)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
