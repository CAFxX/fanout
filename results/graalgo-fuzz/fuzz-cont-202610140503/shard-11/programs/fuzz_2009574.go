package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 string }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint64 }
func genFn0(p0 string) uint64 {
    p0 = "foo bar"
    var v4 uint8 = 127
    var v5 string = p0
    v6 := GenS1{F0: 983, F1: 0, F2: ((("foo bar" + "foo bar") + v5) + (p0 + v5))}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := []uint32{0, 4294967295, 124}
    var v2 uint16 = 65535
    v3 := []string{"", "foo bar"}
    var v7 bool = (980 == 1)
    if len(v3) > 0 {
        v3[0] = ""
    }
    v8 := GenS1{F0: 0, F1: (v2 / (((^(v2))) | 1)), F2: "\n\t"}
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 int64 = 9223372036854775807
        if (0.1 > -2.25) {
            var v11 uint16 = (v2 / ((v2) | 1))
            v12 := GenS2{F0: -981, F1: 240}
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 int32 = 42
                _ = len(v1)
                _ = fi13
                _ = v14
            }
            _ = v11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v15 int8 = 42
    var v16 float32 = 3.14159
    v1 = append(v1, 0)
    var v17 uint = 4294967295
    v18 := GenS1{F0: (v17 % ((v17) | 1)), F1: v2, F2: "hello"}
    v17 = v17
    fmt.Println(v7, v2)
    var v19 int64 = -1
    fmt.Println(v15, v17, v19)
    v2 -= v2
    v20 := GenS3{F0: 817, F1: 1}
    for fi21 := 0; fi21 < 15; fi21++ {
        for fi22 := 0; fi22 < 4; fi22++ {
            v15 = v15
            _ = fi22
        }
        _ = fi21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v7
    _ = v8
}
