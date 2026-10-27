package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int8; F2 bool }
func genFn0() bool {
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: 791}
        var v8 int32 = 553
        if (true || (42 >= 0)) {
            if false {
                var v9 bool = (fi6 > 313)
                v8 += v8
                _ = v9
            }
        } else {
            v10 := GenS1{F0: 255}
            for fi11 := 0; fi11 < 5; fi11++ {
                if true {
                    if (!((("a" + "") + ("a" + "a")) > "")) {
                        v12 := GenS1{F0: 255}
                        _ = v12
                    } else {
                        var v13 uint8 = 128
                        _ = v13
                    }
                    fmt.Println(v8)
                    v14 := GenS1{F0: 0}
                    _ = v14
                }
                _ = fi11
            }
            _ = v10
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v15 float64 = 1.5
    v15 = 1.5
    v16 := []int{832, 732, 0, 874}
    switch 951 {
    case 128:
        _ = len(v16)
    case -1:
        v17 := GenS1{F0: 320}
        _ = v17
    default:
        v15 = (((v15 + v15) + 1.5) + ((1e10 - -2.25) / v15))
    }
    _ = v15
    _ = v16
    return ((0.1 >= 0.0) || true)
}
func genFn1() int {
    v18 := GenS2{F0: 1}
    var v19 int32 = 0
    v19 = v19
    _ = v18
    _ = v19
    return 2147483647
}

func main() {
    var v1 bool = (true && ((0 >= 42) || (("\n\t" + (((("" + ((("\n\t" + "foo bar") + ("" + ("hello" + "foo bar"))) + "foo bar")) + ("" + ("\n\t" + "foo bar"))) + "a") + "\n\t")) != (("hello" + "\n\t") + ""))))
    var v2 int8 = 82
    v3 := []uint32{0}
    var v4 int8 = (^(v2))
    var v5 bool = (v1 || v1)
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 uint64 = 297
        var v22 int16 = 32767
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        for fi24 := 0; fi24 < 13; fi24++ {
            var v25 float32 = 3.14159
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    if (v5 && ((42 != 1000) || true)) {
        v1 = (9223372036854775807 < -1)
        v4 = v2
    } else {
        for fi26 := 0; fi26 < 9; fi26++ {
            var v27 uint16 = 0
            fmt.Println(v1, fi26, v5)
            _ = fi26
            _ = v27
        }
    }
    var v28 int64 = -72
    var v29 uint = 18446744073709551615
    fmt.Println(v4, v4, v29)
    for fi30 := 0; fi30 < 6; fi30++ {
        var v31 uint8 = 184
        _ = fi30
        _ = v31
    }
    if false {
        var v32 float32 = (float32(1e10) / (-2.25 + float32(0.1)))
        var v33 uint16 = 678
        _ = v32
        _ = v33
    }
    v34 := GenS3{F0: 693, F1: v4, F2: v5}
    v28 = v28
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
