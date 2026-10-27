package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
func genFn0() uint16 {
    v6 := GenS1{F0: ("\n\t" + ("hello" + "hello")), F1: -128}
    var v7 int8 = 42
    v8 := []int{128, 0}
    _ = v6
    _ = v7
    _ = v8
    return 913
}

func main() {
    var v1 uint = 1
    var v2 float32 = (1e10 + (1e10 / 3.14159))
    v3 := make(map[uint32]bool)
    v4 := []int{9223372036854775807, 807, -9223372036854775808, 128}
    var v5 uint64 = uint64(v1)
    v5 = v5
    fmt.Println(v2)
    v2 = float32(-2.25)
    v5 += (v5 ^ v5)
    if false {
        _ = len(v4)
        v9 := GenS1{F0: "hello", F1: int8(v5)}
        _ = v9
    } else {
        v3[2] = false
        v2 -= v2
        for fi10 := 0; fi10 < 2; fi10++ {
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := GenS1{F0: ("hello" + ("\n\t" + "foo bar")), F1: -128}
                v2 -= 0.0
                _ = fi11
                _ = v12
            }
            var v13 int16 = -388
            var v14 string = ("hello" + "\n\t")
            _ = fi10
            _ = v13
            _ = v14
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v5 -= (v5 - (v5 << (((^(v1))) & 63)))
    v2 -= v2
    var v15 int = 127
    var v16 uint64 = 2
    for fi17 := 0; fi17 < 6; fi17++ {
        v5 = 74
        v5 = (v16 | (v5 % ((v5) | 1)))
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
