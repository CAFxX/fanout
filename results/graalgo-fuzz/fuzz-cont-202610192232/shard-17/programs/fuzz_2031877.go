package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint16, p1 uint16, p2 int8) bool {
    if (1 == 1) {
        var v6 bool = ("foo bar" > "foo bar")
        v6 = (65535 < 2)
        _ = v6
    }
    if (2147483648 == 2147483647) {
        p2 += (p2 * (p2 ^ p2))
        v7 := GenS1{F0: 42, F1: 0, F2: 128}
        if (4294967295 != 0) {
            var v8 int64 = 1
            _ = v8
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return true
}
func genFn1() uint64 {
    var v9 int64 = -276
    v9 = (v9 | (v9 | 0))
    _ = v9
    return 438
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 32768
    v3 := []string{"a", ("\n\t" + ""), "foo bar", "foo bar"}
    var v4 int32 = 992
    v5 := make(map[uint32]int)
    var v10 uint64 = 280
    v5[21] = v2
    var v11 uint32 = 1
    var v12 string = "\n\t"
    switch int(v4) {
    case 32767:
        fmt.Println(v11)
    }
    fmt.Println(v12)
    fmt.Println(v4, v12)
    for fi13 := 0; fi13 < 12; fi13++ {
        v1[127] = 547
        if false {
            var v14 int64 = 628
            _ = v14
        } else {
            v12 = "\n\t"
            var v15 uint64 = v10
            _ = v15
        }
        v5[uint32(v10)] = (fi13 * (v2 * -2147483648))
        _ = fi13
    }
    switch (v12 + v12) {
    case "\n\t":
        var v16 float32 = 1.5
        _ = v16
    case "foo bar":
        fmt.Println(v12, v11)
    case "hello":
        v12 = v12
    default:
        if true {
            switch 32767 {
            case -75:
                v17 := v5[(v11 - (v11 >> ((259) & 63)))]
                _ = v17
            case -98:
                var v18 uint8 = 127
                _ = v18
            }
            v19 := GenS1{F0: 255, F1: 0, F2: uint8(v11)}
            _ = v19
        } else {
            if true {
                v20 := GenS2{F0: v11}
                _ = v20
            }
            v21 := GenS1{F0: 42, F1: v10, F2: 127}
            fmt.Println(v10, v2)
            _ = v21
        }
    }
    v22 := GenS2{F0: (v11 << ((42) & 63))}
    fmt.Println(v12)
    if true {
        for fi23 := 0; fi23 < 14; fi23++ {
            for fi24 := 0; fi24 < 14; fi24++ {
                var v25 int16 = 1
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
        v26 := GenS1{F0: 1, F1: (v10 * v10), F2: 128}
        _ = v26
    } else {
        var v27 string = v12
        _ = v27
    }
    var v28 int32 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
