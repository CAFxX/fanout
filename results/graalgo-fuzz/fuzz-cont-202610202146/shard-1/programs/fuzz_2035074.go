package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int8; F1 uint64 }
func genFn0() uint32 {
    var v4 int64 = 9223372036854775807
    v5 := GenS1{F0: 18446744073709551615}
    v4 += (v4 + 0)
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            switch 1 {
            case 32767:
                var v8 bool = ((1.5 == 1e10) && (!(!true)))
                v9 := GenS2{F0: (18446744073709551615 >= 1), F1: 127}
                v10 := GenS1{F0: 1}
                _ = v10
                _ = v8
                _ = v9
            case 4294967295:
                fmt.Println(fi6, fi6)
            }
            _ = fi7
        }
        v4 = v4
        v4 = v4
        _ = fi6
    }
    switch 42 {
    case 32768:
        v11 := GenS3{F0: -60, F1: 2}
        _ = v11
    default:
        v4 = (v4 & -970)
    }
    _ = v4
    _ = v5
    return 742
}
func genFn1(p0 uint8) uint32 {
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 int16 = 0
        p0 = (p0 / ((p0) | 1))
        v14 := []int{(fi12 & (fi12 >> ((656) & 63))), (fi12 | (fi12 / ((fi12) | 1))), fi12, fi12}
        _ = fi12
        _ = v13
        _ = v14
    }
    p0 = 42
    v15 := []int{}
    var v16 float32 = 0.0
    _ = p0
    _ = v15
    _ = v16
    return 4294967295
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 1
    v3 := []string{"", "foo bar", "hello", "\n\t"}
    for fi17 := 0; fi17 < 6; fi17++ {
        if false {
            if (!(2 < 665)) {
                if (!(!(0.0 <= 0.0))) {
                    v18 := v1["foo bar"]
                    _ = v18
                } else {
                    var v19 uint16 = 32767
                    fmt.Println(v2, v2, v19)
                    v2 += v2
                    _ = v19
                }
                v2 -= v2
            }
        }
        v20 := GenS3{F0: 74, F1: 313}
        if true {
            var v21 uint = 65535
            _ = v21
        }
        _ = fi17
        _ = v20
    }
    v2 = (^(v2))
    var v22 string = "hello"
    var v23 uint32 = 4294967295
    var v24 string = v22
    var v25 bool = (true || true)
    var v26 uint32 = genFn1(42)
    var v27 uint8 = 179
    var v28 int32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
