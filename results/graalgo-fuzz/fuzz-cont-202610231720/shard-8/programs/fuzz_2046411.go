package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
func genFn0(p0 uint32) bool {
    v6 := []int{-461, 9223372036854775807, 32767, 1}
    var v7 int64 = 42
    _ = p0
    _ = v6
    _ = v7
    return (true || false)
}

func main() {
    v1 := []uint32{1}
    var v2 uint = 65535
    var v3 int16 = 1
    var v4 uint = v2
    v5 := []string{}
    var v8 float64 = 1e10
    fmt.Println(v8, v8, v4)
    v4 = (v4 << ((4294967295) & 63))
    v3 = (v3 >> (((^(v4))) & 63))
    v4 += 838
    fmt.Println(v2)
    v4 = (v2 % ((604) | 1))
    if (0 < -1) {
        var v9 uint64 = 42
        for fi10 := 0; fi10 < 4; fi10++ {
            v1 = append(v1, 65535)
            switch (("" + "a") + ((((("" + "a") + "a") + (("\n\t" + "\n\t") + ("a" + ("foo bar" + ("foo bar" + "\n\t"))))) + "a") + "")) {
            case "a":
                var v11 uint32 = 367
                _ = v11
            default:
                var v12 float32 = 0.1
                _ = v12
            }
            _ = fi10
        }
        _ = v9
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v4 -= (v4 % ((255) | 1))
        _ = fi13
    }
    if (!true) {
        _ = len(v5)
    } else {
        v2 += (v4 / ((v4) | 1))
        v5 = append(v5, (("hello" + "foo bar") + (("hello" + ("a" + "a")) + "\n\t")))
        v14 := GenS1{F0: (!false), F1: 85}
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
