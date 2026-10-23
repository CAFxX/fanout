package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int8 = -128
    v2 := []uint32{2147483647}
    var v3 float32 = 0.0
    v4 := make(map[int]uint32)
    var v5 int32 = -1
    var v6 uint = 2
    fmt.Println(v6, v5, v1)
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: ("a" + "hello")}
        _ = fi7
        _ = v8
    }
    fmt.Println(v6, v3, v3)
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 float32 = float32(1e10)
        _ = fi9
        _ = v10
    }
    fmt.Println(v6, v5, v5)
    if (uint16(v6) == 1) {
        for fi11 := 0; fi11 < 4; fi11++ {
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 int = fi11
                _ = fi12
                _ = v13
            }
            var v14 uint8 = 187
            _ = fi11
            _ = v14
        }
        v2 = append(v2, 58)
        var v15 float32 = v3
        _ = v15
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    var v16 int8 = (-(v1))
    v17 := GenS1{F0: ("foo bar" + "hello")}
    if ((0.1 * 1e10) >= 3.14159) {
        v4[127] = 321
        var v18 int8 = v16
        _ = v18
    } else {
        delete(v4, 4294967295)
        _ = len(v2)
        v19 := GenS1{F0: "hello"}
        _ = v19
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    _ = len(v2)
    v3 = v3
    var v20 int16 = 0
    v6 = (v6 + 1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
