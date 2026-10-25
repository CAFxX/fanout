package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int8; F2 int16 }
func genFn0(p0 int32) uint16 {
    p0 += -485
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: 255}
        v9 := GenS1{F0: 255}
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(p0)
    var v10 uint64 = 735
    var v11 uint16 = 0
    _ = p0
    _ = v10
    _ = v11
    return 112
}

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[uint32]uint32)
    var v3 int16 = 0
    var v4 float32 = 3.14159
    var v5 float32 = (v4 - v4)
    v6 := make(map[int]uint32)
    var v12 int16 = v3
    var v13 int64 = 306
    v14 := v2[352]
    switch "foo bar" {
    case "":
        v15 := GenS2{F0: 649, F1: -62, F2: v12}
        _ = v15
    case "foo bar":
        v5 -= v4
    }
    var v16 int8 = 0
    if false {
        if true {
            fmt.Println(v3, v16, v13)
            for fi17 := 0; fi17 < 13; fi17++ {
                delete(v6, 32767)
                var v18 string = ("\n\t" + "hello")
                _ = fi17
                _ = v18
            }
            fmt.Println(v5, v3, v14)
        }
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        v3 = v3
        if (688 >= 65535) {
            var v20 int32 = 776
            _ = v20
        }
        _ = fi19
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        _ = len(v1)
        _ = fi21
    }
    v22 := GenS2{F0: 48, F1: (v16 * (v16 << ((255) & 63))), F2: (v12 ^ v3)}
    var v23 int64 = (v13 * (v13 >> ((uint(v14)) & 63)))
    v3 = -292
    for fi24 := 0; fi24 < 14; fi24++ {
        v5 += 1.5
        _ = fi24
    }
    v13 += v13
    v5 = 0.0
    var v25 float64 = 0.0
    var v26 int16 = 1000
    v23 = -879
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
