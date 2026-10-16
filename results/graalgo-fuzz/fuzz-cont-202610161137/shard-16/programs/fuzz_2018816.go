package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 uint16; F2 int64 }
func genFn0(p0 int, p1 int, p2 int) int {
    fmt.Println(p1)
    var v5 uint32 = 540
    if true {
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := GenS1{F0: 42, F1: 221}
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 10; fi8++ {
            p0 -= -639
            var v9 string = "a"
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: -9223372036854775808, F1: 0}
        _ = v10
    }
    if (977 == 746) {
        v11 := GenS2{F0: 361, F1: -503}
        _ = v11
    }
    var v12 uint32 = 748
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v5
    return p1
}

func main() {
    v1 := []string{""}
    v2 := []string{("a" + "foo bar")}
    var v3 int32 = 0
    var v4 int = 0
    var v13 uint32 = 42
    v14 := GenS1{F0: -1, F1: 42}
    v15 := GenS3{F0: genFn0((v4 + (v4 % (((v4 + genFn0(128, v4, 2147483647))) | 1))), (v4 << ((4294967295) & 63)), v4)}
    var v16 uint = 2
    v16 -= (^(v16))
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            _ = fi18
        }
        if (557 <= 218) {
            v3 += v3
            v19 := GenS4{F0: v13, F1: 65535, F2: 0}
            _ = v19
        }
        _ = fi17
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 int8 = 0
        var v22 uint16 = 500
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 int32 = (v3 << (((v16 ^ 18446744073709551615)) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
