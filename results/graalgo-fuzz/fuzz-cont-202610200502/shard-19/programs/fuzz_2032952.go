package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint16; F1 int; F2 bool }
type GenS3 struct { F0 uint8; F1 int8; F2 bool }
func genFn0(p0 string, p1 uint64, p2 uint16) string {
    fmt.Println(p1, p2, p1)
    var v6 float64 = (1e10 * 0.1)
    p1 = 131
    var v7 uint64 = p1
    var v8 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return "hello"
}
func genFn1() int16 {
    var v9 int8 = 1
    fmt.Println(v9)
    _ = v9
    return -668
}

func main() {
    var v1 int16 = -1
    v2 := make(map[int]uint32)
    var v3 int8 = -1
    v4 := []string{(("\n\t" + "foo bar") + ""), ("\n\t" + ("foo bar" + "\n\t")), "\n\t", (("hello" + ("a" + "a")) + (("foo bar" + "hello") + ("\n\t" + "foo bar")))}
    var v5 uint8 = 1
    var v10 uint = 42
    v3 += v3
    v11 := GenS1{F0: 87, F1: ("hello" + "\n\t")}
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := GenS2{F0: 1, F1: fi12, F2: ((true || false) && false)}
        var v14 int16 = int16(v5)
        var v15 uint16 = 0
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    v2[102] = 65535
    v16 := GenS3{F0: (v5 * 115), F1: -128, F2: (-2.25 != 3.14159)}
    fmt.Println(v1, v1)
    fmt.Println(v1, v10)
    v5 = 225
    v3 += 127
    var v17 uint = 911
    if (!false) {
        v10 = (v10 / (((^(v10))) | 1))
        if (genFn1() > 270) {
            v5 += v5
            delete(v2, 9223372036854775807)
            fmt.Println(v17, v3)
        }
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 uint8 = (v5 << (((v17 - 255)) & 63))
        if true {
            v5 -= 127
            v20 := v2[fi18]
            delete(v2, int(v3))
            _ = v20
        }
        _ = fi18
        _ = v19
    }
    v21 := GenS1{F0: v5, F1: "foo bar"}
    v5 = (v5 ^ v5)
    var v22 uint32 = 307
    v23 := GenS1{F0: v5, F1: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
