package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint }
type GenS2 struct { F0 int8; F1 uint32 }
func genFn0(p0 int8, p1 int64, p2 uint) int32 {
    for fi7 := 0; fi7 < 14; fi7++ {
        p0 -= -40
        _ = fi7
    }
    fmt.Println(p2, p0, p1)
    p1 -= -1
    if true {
        fmt.Println(p1, p0, p0)
        p0 = 42
    } else {
        var v8 int32 = -2147483648
        var v9 bool = false
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 uint64 = 42
            _ = fi10
            _ = v11
        }
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 420
}

func main() {
    v1 := []string{"foo bar", "a"}
    v2 := make(map[uint32]bool)
    var v3 float64 = 0.1
    var v4 string = ((("hello" + "hello") + ("foo bar" + "hello")) + "")
    v5 := make(map[uint32]bool)
    var v6 int32 = 42
    v4 = v4
    for fi12 := 0; fi12 < 14; fi12++ {
        fmt.Println(v4, v3)
        for fi13 := 0; fi13 < 2; fi13++ {
            v14 := GenS1{F0: (v6 << ((255) & 63)), F1: 0, F2: 18446744073709551615}
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    fmt.Println(v6, v6, v6)
    var v15 string = (((v4 + v4) + v4) + "hello")
    v16 := GenS2{F0: 0, F1: 520}
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 int8 = 100
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 uint32 = 1
            if false {
                var v21 uint8 = 7
                v20 = (v20 & v20)
                _ = v21
            } else {
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                v6 = (v6 / ((v6) | 1))
            }
            _ = fi19
            _ = v20
        }
        _ = len(v5)
        _ = fi17
        _ = v18
    }
    var v22 uint64 = 407
    v23 := GenS2{F0: -22, F1: 546}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
