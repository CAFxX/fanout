package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
func genFn0(p0 int32, p1 int) uint32 {
    p1 = 2147483647
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int16 = 32767
        var v8 string = "\n\t"
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 string = "\n\t"
    var v10 int = p1
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 2
}
func genFn1(p0 int8, p1 int16) int64 {
    var v11 float64 = float64(float32(0.0))
    v12 := []int{127}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 9223372036854775807
}

func main() {
    var v1 int64 = -730
    var v2 bool = (!(7 > 343))
    var v3 int16 = 32767
    v4 := make(map[uint32]uint32)
    v5 := []int{715, -761, -818}
    switch -695 {
    case -202:
        v4[462] = 481
    case 9223372036854775807:
        var v13 string = (("foo bar" + "") + "")
        _ = v13
    case 4294967295:
        var v14 bool = (!v2)
        _ = v14
    default:
        for fi15 := 0; fi15 < 10; fi15++ {
            fmt.Println(v2, fi15)
            for fi16 := 0; fi16 < 17; fi16++ {
                v3 += -751
                v17 := GenS1{F0: (("hello" + "") + "a"), F1: fi16, F2: (("\n\t" + ("a" + "hello")) == (("foo bar" + "\n\t") + ("hello" + ("\n\t" + ""))))}
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    }
    v2 = (!(!v2))
    var v18 float32 = 1e10
    v19 := GenS2{F0: 560}
    v1 += v1
    v2 = (42 < genFn0(42, 985))
    fmt.Println(v3, v1, v3)
    var v20 float64 = 0.0
    v1 = v1
    var v21 uint8 = 128
    v22 := GenS3{F0: "hello"}
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 int = (fi23 - fi23)
        if len(v5) > 0 {
            v5[0] = 32767
        }
        _ = fi23
        _ = v24
    }
    _ = len(v5)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
