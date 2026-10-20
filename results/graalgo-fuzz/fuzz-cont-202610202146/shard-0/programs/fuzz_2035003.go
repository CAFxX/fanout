package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint16, p1 bool) uint32 {
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            p1 = (true || p1)
            var v8 uint8 = 0
            if (!(2 >= 69)) {
                p0 = (p0 - p0)
                v9 := GenS1{F0: -128, F1: -308, F2: 42}
                _ = v9
            } else {
                v10 := []int{fi6, (fi6 << ((65535) & 63)), int(p0)}
                var v11 int8 = 127
                v12 := GenS2{F0: 582, F1: 0}
                _ = v10
                _ = v11
                _ = v12
            }
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    if (4294967295 == 220) {
        fmt.Println(p1, p0)
    }
    if (127 > 114) {
        var v13 int8 = 92
        _ = v13
    } else {
        var v14 int8 = -1
        _ = v14
    }
    var v15 uint16 = p0
    v16 := GenS2{F0: (v15 - v15), F1: 214}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 250
}
func genFn1(p0 int32) uint16 {
    for fi17 := 0; fi17 < 15; fi17++ {
        p0 = (p0 & -1)
        if false {
            var v18 uint32 = 42
            _ = v18
        } else {
            var v19 string = ("foo bar" + ("" + ("hello" + "")))
            fmt.Println(p0, p0)
            _ = v19
        }
        _ = fi17
    }
    p0 = p0
    _ = p0
    return 875
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    v3 := []string{}
    v4 := []uint32{971, 2147483648, 1, 1}
    v5 := []uint32{0, 151, 2, 0}
    for fi20 := 0; fi20 < 16; fi20++ {
        if true {
            v3 = append(v3, ("a" + (("a" + "") + ("hello" + ("a" + ("foo bar" + "hello"))))))
        } else {
            for fi21 := 0; fi21 < 13; fi21++ {
                _ = len(v4)
                v5 = append(v5, 4294967295)
                _ = fi21
            }
            var v22 uint8 = 128
            v23 := GenS1{F0: 0, F1: 1, F2: 0}
            _ = v22
            _ = v23
        }
        var v24 int32 = 42
        fmt.Println(fi20, v24)
        _ = fi20
        _ = v24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        if false {
            var v26 int32 = -2147483648
            v27 := GenS1{F0: 120, F1: v26, F2: 42}
            _ = v26
            _ = v27
        }
        _ = fi25
    }
    var v28 float32 = ((0.1 + 0.0) - (((-2.25 + 0.1) - (float32(-2.25) - 1e10)) / ((float32(3.14159) + ((float32(0.1) - ((((-2.25 * 1e10) - 3.14159) - 0.0) * 0.1)) - float32(((3.14159 / 3.14159) + -2.25)))) + 1.5)))
    v28 = ((1e10 / (1e10 / v28)) * v28)
    switch "\n\t" {
    case "hello":
        fmt.Println(v28)
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        var v30 int16 = 32767
        var v31 float64 = (-2.25 - 1.5)
        _ = fi29
        _ = v30
        _ = v31
    }
    if false {
        switch ("\n\t" + "a") {
        case "":
            var v32 int64 = 9223372036854775807
            _ = v32
        case "foo bar":
            v28 = v28
        case "hello":
            var v33 bool = (false && true)
            _ = v33
        default:
            v2["a"] = 938
        }
        var v34 int = 128
        _ = v34
    }
    v35 := GenS2{F0: 876, F1: 633}
    v36 := GenS1{F0: 114, F1: 2147483647, F2: 886}
    v28 = v28
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
