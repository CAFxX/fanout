package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 string }
type GenS2 struct { F0 int8; F1 uint16; F2 uint16 }

func main() {
    var v1 int = 127
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    var v4 float32 = 1e10
    v5 := []uint32{}
    v6 := v2[""]
    v7 := GenS1{F0: (("\n\t" + "foo bar") + ""), F1: (("" + ("a" + "")) + ("a" + "foo bar")), F2: "a"}
    v2[("a" + "\n\t")] = v6
    v8 := GenS2{F0: 26, F1: uint16(v6), F2: 0}
    for fi9 := 0; fi9 < 10; fi9++ {
        v6 += (v6 ^ 875)
        v10 := GenS2{F0: 127, F1: 0, F2: 256}
        v11 := GenS1{F0: ("" + ""), F1: "foo bar", F2: "foo bar"}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int16 = -32768
    var v13 uint = 426
    switch (v1 * v1) {
    case -595:
        v3[v6] = v6
    case -577:
        v13 += v13
    default:
        v14 := GenS1{F0: "a", F1: "hello", F2: ("foo bar" + "")}
        _ = v14
    }
    var v15 float64 = (1e10 / 0.1)
    v6 = 65535
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v13 = v13
        if ((v15 * v15) > v15) {
            v15 += 1.5
        } else {
            for fi17 := 0; fi17 < 12; fi17++ {
                v1 -= int(v12)
                if ((v12 ^ (v12 ^ (v12 % ((18) | 1)))) < v12) {
                    var v18 uint64 = 2
                    _ = v18
                }
                _ = fi17
            }
            if len(v5) > 0 {
                v5[0] = 42
            }
        }
        for fi19 := 0; fi19 < 15; fi19++ {
            v20 := GenS2{F0: -128, F1: 0, F2: 255}
            _ = fi19
            _ = v20
        }
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
