package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int16 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0(p0 uint64) uint8 {
    if (!true) {
        for fi6 := 0; fi6 < 8; fi6++ {
            if (104 != 32767) {
                v7 := []int{(fi6 * (fi6 >> ((65535) & 63))), fi6, (fi6 | fi6)}
                _ = v7
            }
            p0 = p0
            v8 := GenS1{F0: -128, F1: 501, F2: 127}
            _ = fi6
            _ = v8
        }
        switch 0 {
        case 4294967295:
            fmt.Println(p0)
        }
        if true {
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 float32 = (-2.25 / 0.1)
                var v11 int32 = -703
                v10 = (0.1 - (1e10 - v10))
                _ = fi9
                _ = v10
                _ = v11
            }
            fmt.Println(p0)
            p0 += (p0 - (p0 << ((1) & 63)))
        } else {
            for fi12 := 0; fi12 < 2; fi12++ {
                fmt.Println(fi12, fi12)
                v13 := GenS2{F0: 42, F1: (!(!((^(fi12)) < (fi12 | (fi12 | fi12)))))}
                _ = fi12
                _ = v13
            }
        }
    } else {
        p0 = 2
        var v14 int16 = 1
        _ = v14
    }
    v15 := []int{9223372036854775807, 1, -2147483648, 32767}
    var v16 int64 = 0
    var v17 int8 = 100
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    return 1
}

func main() {
    var v1 uint32 = 830
    var v2 uint = 65535
    v3 := []string{"foo bar", ("" + "\n\t")}
    var v4 int8 = 103
    var v5 string = ""
    fmt.Println(v1)
    v4 -= 100
    v18 := GenS1{F0: 38, F1: (v2 ^ 1), F2: 42}
    var v19 float32 = 1.5
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 int16 = 1
        v3 = append(v3, v5)
        var v22 uint64 = 0
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        v4 += v4
        _ = fi23
    }
    fmt.Println(v5)
    var v24 uint32 = (v1 + (^(v1)))
    if false {
        var v25 uint32 = 361
        _ = v25
    }
    v26 := GenS2{F0: 127, F1: false}
    v27 := GenS3{F0: 42, F1: (v4 - (v4 + (-(v4))))}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
