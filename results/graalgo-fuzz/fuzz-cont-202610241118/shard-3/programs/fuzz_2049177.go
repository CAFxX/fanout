package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int8 }
type GenS2 struct { F0 bool; F1 uint16 }
func genFn0(p0 bool, p1 int) string {
    fmt.Println(p0, p1)
    p0 = p0
    var v6 uint32 = 43
    _ = p0
    _ = p1
    _ = v6
    return "a"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 42
    v3 := []uint32{4294967295, 42, 4294967295}
    var v4 float32 = 0.0
    var v5 uint8 = 255
    var v7 int32 = 2147483647
    var v8 int = 32768
    for fi9 := 0; fi9 < 15; fi9++ {
        v5 = v5
        var v10 int = v8
        v1["hello"] = 2
        _ = fi9
        _ = v10
    }
    _ = len(v3)
    for fi11 := 0; fi11 < 2; fi11++ {
        if (681 < 314) {
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 uint16 = 251
                _ = len(v3)
                var v14 uint32 = 42
                var v15 bool = false
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v16 int = (v8 * (v8 * (fi12 + (v8 | 1))))
                var v17 float32 = float32(1.5)
                v5 = 1
                v14 += v14
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
                _ = v16
                _ = v17
            }
            var v18 int64 = -184
            _ = v18
        } else {
            var v19 uint32 = 2147483648
            var v20 int16 = 42
            _ = v19
            _ = v20
        }
        _ = fi11
    }
    switch (v8 & v8) {
    case 127:
        v2 = (v2 | v2)
    }
    v4 += (v4 - float32(float64(0.0)))
    var v21 int16 = v2
    v8 += v8
    v2 = (v21 + (v2 << ((261) & 63)))
    for fi22 := 0; fi22 < 7; fi22++ {
        v2 = 1
        var v23 string = "foo bar"
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: (v7 % ((-1) | 1)), F1: 4294967295, F2: 127}
    if (1 <= (v21 & (v2 & (v21 >> ((1) & 63))))) {
        fmt.Println(v8, v5)
        if true {
            var v25 float64 = (-2.25 - 1.5)
            _ = v25
        }
        v26 := GenS1{F0: 2147483647, F1: 1, F2: -128}
        _ = v26
    } else {
        v27 := GenS2{F0: false, F1: 808}
        var v28 int32 = v7
        for fi29 := 0; fi29 < 8; fi29++ {
            v21 = (v21 >> ((18446744073709551615) & 63))
            _ = fi29
        }
        _ = v27
        _ = v28
    }
    var v30 uint32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
