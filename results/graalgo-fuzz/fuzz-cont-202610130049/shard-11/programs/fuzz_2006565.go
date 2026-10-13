package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint64 }
func genFn0() uint32 {
    switch "hello" {
    case "foo bar":
        var v5 uint64 = 42
        _ = v5
    case "\n\t":
        var v6 string = "a"
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        switch "hello" {
        case "\n\t":
            var v8 int8 = 0
            _ = v8
        default:
            for fi9 := 0; fi9 < 18; fi9++ {
                v10 := GenS1{F0: 129, F1: -2147483648}
                if true {
                    var v11 string = ("\n\t" + ("" + ("a" + "hello")))
                    _ = v11
                }
                _ = fi9
                _ = v10
            }
        }
        if false {
            v12 := GenS1{F0: 193, F1: 42}
            _ = v12
        } else {
            var v13 uint32 = 65535
            _ = v13
        }
        v14 := GenS1{F0: 127, F1: -314}
        _ = fi7
        _ = v14
    }
    var v15 bool = (0 <= 780)
    _ = v15
    return 42
}
func genFn1(p0 uint32, p1 uint64) int32 {
    for fi16 := 0; fi16 < 14; fi16++ {
        if (fi16 == (fi16 ^ int(p0))) {
            v17 := GenS2{F0: (p1 << ((0) & 63)), F1: 42}
            var v18 float32 = -2.25
            _ = v17
            _ = v18
        }
        var v19 float64 = ((0.0 / 1e10) + float64(-2.25))
        _ = fi16
        _ = v19
    }
    v20 := []int{1, -135, 1, 127}
    _ = p0
    _ = p1
    _ = v20
    return -1
}

func main() {
    var v1 int16 = 0
    v2 := make(map[int]uint32)
    v3 := []string{("foo bar" + "hello"), (("" + "") + (("" + ("" + "\n\t")) + ("hello" + ("" + "hello")))), "foo bar", ("" + "hello")}
    var v4 float32 = (3.14159 * 3.14159)
    for fi21 := 0; fi21 < 16; fi21++ {
        v2[(fi21 + fi21)] = 42
        _ = fi21
    }
    var v22 int64 = 42
    _ = len(v3)
    fmt.Println(v22, v22, v1)
    if true {
        if (-2147483648 != int32(v1)) {
            v22 = v22
            v22 -= v22
            if len(v3) > 0 {
                v3[0] = ("" + "hello")
            }
        }
    } else {
        v23 := GenS1{F0: 237, F1: -88}
        if (0 > (v1 & (v1 << ((42) & 63)))) {
            var v24 int32 = -2147483648
            v25 := GenS1{F0: 240, F1: genFn1(42, 1)}
            _ = v24
            _ = v25
        }
        _ = v23
    }
    v3 = append(v3, "foo bar")
    v26 := GenS3{F0: 2}
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
}
