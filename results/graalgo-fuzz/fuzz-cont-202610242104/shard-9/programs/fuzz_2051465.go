package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int }
type GenS2 struct { F0 uint8; F1 int64 }
func genFn0(p0 int) int64 {
    v6 := GenS1{F0: 42, F1: "a", F2: p0}
    if false {
        v7 := GenS1{F0: 0, F1: "hello", F2: (p0 + p0)}
        var v8 uint64 = 512
        _ = v7
        _ = v8
    }
    if ((-1 <= 1) || (388 != -31)) {
        for fi9 := 0; fi9 < 12; fi9++ {
            p0 += 0
            var v10 float64 = 3.14159
            v11 := GenS1{F0: -73, F1: "\n\t", F2: (fi9 & (fi9 - (^(p0))))}
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    p0 += (p0 / (((-(p0))) | 1))
    var v12 int64 = 9223372036854775807
    _ = p0
    _ = v12
    _ = v6
    return (^(v12))
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]string)
    var v3 uint8 = 42
    var v4 int64 = -1
    v5 := []uint32{}
    v1 = v1
    v5 = append(v5, 791)
    fmt.Println(v1, v4)
    v2["hello"] = ((v1 + "hello") + v1)
    var v13 int32 = 42
    var v14 int16 = -32768
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 int64 = 9223372036854775807
        _ = fi15
        _ = v16
    }
    fmt.Println(v13, v14)
    if true {
        for fi17 := 0; fi17 < 18; fi17++ {
            fmt.Println(v14, v13, v1)
            v18 := GenS1{F0: 0, F1: (("" + (v1 + ((("\n\t" + "hello") + v1) + (((v1 + (v1 + v1)) + v1) + v1)))) + ((v1 + "\n\t") + v1)), F2: (-(fi17))}
            v13 -= 1
            _ = fi17
            _ = v18
        }
        v19 := GenS2{F0: (v3 % (((v3 - (v3 % (((v3 | 0)) | 1)))) | 1)), F1: (v4 / ((int64(v3)) | 1))}
        _ = v19
    } else {
        var v20 uint64 = 0
        _ = v20
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        v1 = v1
        var v22 float64 = ((1.5 / (1e10 + -2.25)) / ((3.14159 * (1e10 - ((0.0 / -2.25) / (3.14159 / -2.25)))) / 1.0))
        _ = fi21
        _ = v22
    }
    if (1e10 != float64((((1.5 * 0.1) - -2.25) * -2.25))) {
        for fi23 := 0; fi23 < 2; fi23++ {
            var v24 float64 = float64(((3.14159 + 0.0) + 3.14159))
            for fi25 := 0; fi25 < 16; fi25++ {
                v24 -= -2.25
                var v26 int = int(v13)
                _ = fi25
                _ = v26
            }
            _ = fi23
            _ = v24
        }
    } else {
        var v27 int16 = (v14 * (v14 & 42))
        for fi28 := 0; fi28 < 20; fi28++ {
            var v29 float32 = float32((1.5 - 3.14159))
            v5 = append(v5, 4294967295)
            _ = fi28
            _ = v29
        }
        _ = v27
    }
    var v30 bool = (1 == v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
