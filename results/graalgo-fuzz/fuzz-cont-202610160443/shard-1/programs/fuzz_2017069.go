package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int) bool {
    var v4 int8 = 42
    p0 = p0
    v4 = (v4 << ((4294967295) & 63))
    _ = p0
    _ = v4
    return false
}
func genFn1(p0 uint8, p1 uint64) uint16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 uint = 4294967295
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int16 = -1
        v9 := []int{(-(fi7))}
        for fi10 := 0; fi10 < 16; fi10++ {
            p1 = p1
            _ = fi10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    if (!(true && true)) {
        v11 := GenS1{F0: -9223372036854775808, F1: 32767}
        fmt.Println(p0, p1)
        _ = v11
    } else {
        v12 := GenS1{F0: -1, F1: 532}
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := GenS2{F0: 0, F1: 0}
            var v15 uint8 = p0
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 uint = 2
        _ = v12
        _ = v16
    }
    p1 -= (p1 * p1)
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 32767
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]bool)
    v3 := []string{}
    v1 -= (v1 * (^(v1)))
    for fi17 := 0; fi17 < 9; fi17++ {
        if (900 <= -32768) {
            var v18 int = fi17
            _ = v18
        }
        _ = fi17
    }
    v1 -= v1
    fmt.Println(v1)
    var v19 int8 = -78
    v19 += (v19 & v19)
    fmt.Println(v1, v1)
    fmt.Println(v19, v19)
    var v20 string = (("\n\t" + "hello") + "hello")
    if len(v3) > 0 {
        v3[0] = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
