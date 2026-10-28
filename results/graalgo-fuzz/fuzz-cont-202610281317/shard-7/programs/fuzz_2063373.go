package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int16 }
func genFn0(p0 int8, p1 int8, p2 int8) int32 {
    v4 := []int{}
    p1 += 88
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}
func genFn1(p0 uint16, p1 bool, p2 uint) uint16 {
    p2 += 0
    p1 = ((((((1.5 / 1e10) <= -2.25) || p1) && p1) || (42 < -1)) || (-823 <= 0))
    p1 = (255 > 128)
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    v1 := []uint32{0}
    v2 := make(map[string]int)
    v3 := []int{9223372036854775807}
    var v5 uint16 = 1
    if (false && true) {
        for fi6 := 0; fi6 < 4; fi6++ {
            for fi7 := 0; fi7 < 4; fi7++ {
                fmt.Println(v5, v5, fi7)
                _ = fi7
            }
            fmt.Println(fi6, v5)
            v5 += v5
            _ = fi6
        }
    } else {
        var v8 uint64 = 42
        switch "foo bar" {
        case "":
            var v9 uint = 304
            _ = v9
        }
        if true {
            var v10 uint = 0
            var v11 int32 = -532
            _ = v10
            _ = v11
        }
        _ = v8
    }
    var v12 int = 32767
    for fi13 := 0; fi13 < 14; fi13++ {
        v12 -= 32768
        var v14 float64 = (0.1 / 1e10)
        var v15 uint64 = 4294967295
        _ = fi13
        _ = v14
        _ = v15
    }
    switch (v12 * 9223372036854775807) {
    case 344:
        var v16 float32 = (3.14159 / 1.5)
        _ = v16
    case 2147483647:
        var v17 int8 = -1
        _ = v17
    }
    switch "a" {
    case "foo bar":
        switch ("hello" + ("\n\t" + "hello")) {
        case "":
            if false {
                v18 := v2[(((("\n\t" + (("" + "a") + ("" + "hello"))) + "a") + ("foo bar" + "\n\t")) + "")]
                v19 := GenS1{F0: -1, F1: 42, F2: -32768}
                if (false || (0 != 934)) {
                    v20 := GenS1{F0: -22, F1: 32767, F2: 1}
                    _ = v20
                }
                _ = v18
                _ = v19
            }
        default:
            for fi21 := 0; fi21 < 18; fi21++ {
                fmt.Println(fi21, v12)
                _ = fi21
            }
        }
    case "\n\t":
        v22 := v2["a"]
        _ = v22
    default:
        for fi23 := 0; fi23 < 19; fi23++ {
            v12 = v12
            v5 -= v5
            _ = fi23
        }
    }
    if false {
        if false {
            v2["\n\t"] = (v12 - v12)
            var v24 int64 = 75
            _ = v24
        }
        if len(v3) > 0 {
            v3[0] = -657
        }
    }
    fmt.Println(v5, v5)
    switch v12 {
    case 9223372036854775807:
        var v25 uint32 = 65535
        _ = v25
    case 0:
        v3 = append(v3, 4294967295)
    case -726:
        v5 = v5
    default:
        if (0.1 < -2.25) {
            var v26 int16 = -70
            v27 := GenS1{F0: 1, F1: (v26 * (v26 ^ 32767)), F2: (v26 >> ((42) & 63))}
            _ = v26
            _ = v27
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v5
}
