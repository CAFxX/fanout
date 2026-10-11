package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int8; F1 int; F2 uint16 }
func genFn0(p0 int8) uint {
    v6 := []int{4294967295}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v7 := GenS1{F0: 65535, F1: 32767}
    var v8 uint = 115
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return (v8 % ((v8) | 1))
}

func main() {
    var v1 uint8 = 2
    var v2 uint = 311
    v3 := make(map[int]int)
    var v4 int32 = 0
    v5 := []uint32{42, 2147483647}
    if false {
        if true {
            v2 -= v2
            var v9 int8 = 100
            _ = v9
        }
        var v10 int = 957
        var v11 int8 = 0
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 32767, F1: -245}
    if true {
        if (!true) {
            switch -9223372036854775808 {
            case 466:
                for fi13 := 0; fi13 < 1; fi13++ {
                    v5 = append(v5, 2)
                    _ = fi13
                }
            }
            v14 := GenS1{F0: uint16(v1), F1: 1000}
            _ = v14
        }
        v15 := v3[9223372036854775807]
        v16 := GenS1{F0: uint16(v1), F1: -32768}
        _ = v15
        _ = v16
    }
    if len(v5) > 0 {
        v5[0] = 492
    }
    v2 = v2
    if true {
        v1 = uint8(v2)
        if (4294967295 <= -207) {
            v4 = 1
            if ((171 == 9223372036854775807) || false) {
                v4 += v4
            }
            v4 = -2147483648
        } else {
            v4 += -977
        }
        for fi17 := 0; fi17 < 6; fi17++ {
            v18 := v3[(fi17 + fi17)]
            _ = len(v5)
            _ = fi17
            _ = v18
        }
    } else {
        fmt.Println(v4, v2, v2)
        v19 := v3[2147483647]
        _ = v19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        fmt.Println(fi20)
        v4 = v4
        _ = fi20
    }
    v2 = 74
    var v21 float32 = (0.0 - (float32(((0.0 / 1e10) - 1.5)) + 0.1))
    v22 := GenS1{F0: 1, F1: -688}
    if ("" < ((("foo bar" + "foo bar") + "foo bar") + "\n\t")) {
        switch 128 {
        case 2147483647:
            switch ("a" + "\n\t") {
            case "foo bar":
                v23 := GenS1{F0: 2, F1: 837}
                _ = v23
            }
        case 42:
            fmt.Println(v21)
        case 1:
            var v24 uint64 = 18446744073709551615
            _ = v24
        }
    }
    if true {
        var v25 uint64 = 4294967295
        var v26 float64 = 3.14159
        _ = v25
        _ = v26
    }
    v27 := GenS2{F0: -30, F1: -1, F2: 65535}
    v4 -= v4
    if true {
        v4 -= v4
        var v28 uint32 = 698
        var v29 bool = ("a" > "foo bar")
        _ = v28
        _ = v29
    }
    var v30 uint = v2
    var v31 uint64 = 945
    v32 := GenS2{F0: 35, F1: 741, F2: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
