package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int8; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS6 struct { F0 int }
func genFn0(p0 uint32, p1 int64) uint16 {
    p1 += (p1 << ((0) & 63))
    var v5 float32 = 1e10
    _ = p0
    _ = p1
    _ = v5
    return 1
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]bool)
    var v3 int64 = -1
    v4 := make(map[int]uint32)
    var v6 uint = 0
    switch ((((("hello" + ("hello" + "\n\t")) + "hello") + ("\n\t" + ("" + ((("foo bar" + ("" + "hello")) + "foo bar") + "\n\t")))) + ("" + (("\n\t" + (("a" + "\n\t") + (("a" + (("\n\t" + "hello") + ("foo bar" + "foo bar"))) + "hello"))) + "foo bar"))) + ("" + "\n\t")) {
    case "foo bar":
        v7 := GenS1{F0: 1}
        _ = v7
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        var v8 int16 = 545
        v3 += (v3 >> (((^(v6))) & 63))
        v1 = append(v1, "foo bar")
        _ = v8
    } else {
        delete(v2, 435)
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS2{F0: 0, F1: (v6 / ((4294967295) | 1))}
        _ = fi9
        _ = v10
    }
    v1 = append(v1, "hello")
    v11 := GenS3{F0: 0, F1: 42}
    var v12 int64 = v3
    v13 := GenS4{F0: -1, F1: v12, F2: genFn0(0, v12)}
    if true {
        if (1 >= 2) {
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := GenS5{F0: 4294967295, F1: 4294967295, F2: 42}
                v2[2147483648] = (1000 < -32768)
                var v16 uint = v6
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        var v17 int = -645
        v18 := v2[2]
        _ = v17
        _ = v18
    }
    v4[32768] = 383
    v19 := GenS6{F0: -9223372036854775808}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
