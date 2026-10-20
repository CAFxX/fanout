package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
func genFn0() uint8 {
    var v4 uint8 = 42
    v5 := GenS1{F0: 2147483648, F1: (65535 == 0), F2: 0}
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int = (fi6 | (-(fi6)))
        v4 = 0
        _ = fi6
        _ = v7
    }
    fmt.Println(v4)
    var v8 string = (("\n\t" + "hello") + "a")
    _ = v4
    _ = v5
    _ = v8
    return (v4 * (^(v4)))
}
func genFn1(p0 int32) bool {
    for fi9 := 0; fi9 < 13; fi9++ {
        fmt.Println(fi9, p0)
        if (108 >= 0) {
            v10 := GenS2{F0: 65535}
            var v11 int16 = -295
            v12 := []int{(^(fi9)), 32767, (fi9 % (((^(fi9))) | 1))}
            _ = v10
            _ = v11
            _ = v12
        } else {
            var v13 int = 604
            if true {
                fmt.Println(v13, v13, fi9)
                v14 := GenS1{F0: 65535, F1: (15 != 32767), F2: 65535}
                v13 += fi9
                _ = v14
            }
            var v15 uint = 2
            _ = v13
            _ = v15
        }
        fmt.Println(p0, p0)
        _ = fi9
    }
    p0 += 247
    fmt.Println(p0)
    _ = p0
    return false
}

func main() {
    var v1 uint16 = 1
    v2 := []int{128, 32768}
    v3 := make(map[uint32]bool)
    var v16 uint32 = 2147483647
    v1 -= v1
    if true {
        switch 32767 {
        case 4294967295:
            for fi17 := 0; fi17 < 3; fi17++ {
                if len(v2) > 0 {
                    v2[0] = fi17
                }
                _ = fi17
            }
        }
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 float64 = (1e10 * 1e10)
            v19 = v19
            _ = fi18
            _ = v19
        }
        _ = len(v3)
    }
    var v20 int64 = 623
    fmt.Println(v16)
    if (!(-2.25 == 0.0)) {
        switch 0 {
        case 438:
            for fi21 := 0; fi21 < 17; fi21++ {
                var v22 uint = 0
                v20 += -1
                _ = fi21
                _ = v22
            }
        }
        v23 := GenS1{F0: 539, F1: true, F2: v1}
        _ = v23
    } else {
        if true {
            var v24 uint64 = 18446744073709551615
            v16 -= 0
            _ = v24
        } else {
            v16 -= v16
            v25 := GenS2{F0: 0}
            fmt.Println(v20)
            _ = v25
        }
    }
    if ((70 < 127) || true) {
        v26 := GenS3{F0: 0}
        _ = v26
    }
    var v27 uint64 = 4294967295
    fmt.Println(v1, v20, v27)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v27
    _ = v3
}
