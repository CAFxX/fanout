package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
func genFn0(p0 int16, p1 int32) bool {
    var v6 uint64 = 985
    switch "a" {
    case "\n\t":
        for fi7 := 0; fi7 < 15; fi7++ {
            v6 = (v6 - 2)
            v8 := GenS1{F0: 2, F1: 335}
            v9 := GenS1{F0: 4294967295, F1: 255}
            _ = fi7
            _ = v8
            _ = v9
        }
    case "hello":
        p0 = p0
    case "foo bar":
        var v10 float32 = 0.0
        _ = v10
    }
    var v11 uint16 = 42
    v6 += 68
    fmt.Println(v11, v6, v6)
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return (1.5 != 1e10)
}

func main() {
    var v1 int64 = 840
    v2 := make(map[uint32]uint32)
    var v3 int = 128
    var v4 int64 = (v1 % (((v1 + v1)) | 1))
    var v5 uint = 1
    v5 += (v5 * (v5 | 889))
    var v12 uint64 = 525
    if false {
        v1 = (-(v4))
    } else {
        var v13 float32 = -2.25
        var v14 int = 9223372036854775807
        fmt.Println(v5)
        _ = v13
        _ = v14
    }
    v2[65535] = 518
    var v15 uint8 = 255
    v16 := v2[4294967295]
    v17 := []int{(^(v3))}
    v18 := GenS1{F0: v16, F1: 1}
    for fi19 := 0; fi19 < 5; fi19++ {
        v1 -= (v1 >> ((v5) & 63))
        if ((v4 * v4) > (^(v1))) {
            var v20 int64 = v4
            _ = v20
        }
        v12 = 614
        _ = fi19
    }
    v21 := GenS1{F0: (v16 / (((v16 >> ((1) & 63))) | 1)), F1: 255}
    fmt.Println(v5)
    if false {
        var v22 int = (v3 >> ((186) & 63))
        for fi23 := 0; fi23 < 3; fi23++ {
            v16 = 145
            for fi24 := 0; fi24 < 19; fi24++ {
                var v25 int = 151
                v26 := v2[v16]
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = fi23
        }
        _ = v22
    } else {
        v27 := GenS1{F0: v16, F1: v5}
        if (!false) {
            v4 = v1
            _ = len(v17)
        } else {
            var v28 int64 = (v4 * (v4 >> ((350) & 63)))
            if (("" + "a") <= "") {
                v4 = v1
                var v29 string = ""
                var v30 bool = (-748 != v3)
                _ = v29
                _ = v30
            }
            v31 := GenS1{F0: (^(v16)), F1: v5}
            _ = v28
            _ = v31
        }
        _ = v27
    }
    v32 := v2[(^(v16))]
    v33 := v2[v16]
    for fi34 := 0; fi34 < 20; fi34++ {
        v16 = v33
        var v35 float64 = ((0.1 / (0.1 * 1.5)) / -2.25)
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
