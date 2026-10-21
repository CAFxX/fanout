package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int32) int8 {
    var v7 int8 = -101
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int64 = -1
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v7
    return -1
}

func main() {
    var v1 uint = 744
    v2 := []string{"\n\t", ("foo bar" + "hello")}
    v3 := make(map[uint32]string)
    var v4 int64 = 1
    var v5 uint32 = 4294967295
    v6 := make(map[uint32]string)
    _ = len(v3)
    fmt.Println(v1)
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 int32 = 0
        for fi12 := 0; fi12 < 14; fi12++ {
            _ = len(v2)
            v11 -= (^(v11))
            _ = fi12
        }
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 uint32 = uint32(v1)
            _ = fi13
            _ = v14
        }
        _ = fi10
        _ = v11
    }
    v15 := GenS1{F0: 42, F1: ("a" < ""), F2: (^(v4))}
    fmt.Println(v4, v1, v4)
    if len(v2) > 0 {
        v2[0] = (("hello" + "\n\t") + "a")
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 float32 = float32(float64((0.1 + ((0.0 + 3.14159) * 1e10))))
        if ((2147483647 >= -1) || (!true)) {
            v18 := GenS2{F0: 0, F1: uint8(v5)}
            v4 += (^(v4))
            _ = v18
        }
        var v19 string = ("hello" + ("foo bar" + ("a" + (((("" + "") + "a") + ("hello" + "\n\t")) + ("" + ("" + ("\n\t" + "hello")))))))
        _ = fi16
        _ = v17
        _ = v19
    }
    v4 = (v4 ^ (v4 >> ((952) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
