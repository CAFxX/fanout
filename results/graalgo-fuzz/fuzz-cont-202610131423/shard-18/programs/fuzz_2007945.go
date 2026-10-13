package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS3 struct { F0 int32; F1 int16; F2 int }
type GenS4 struct { F0 int16; F1 int8 }
func genFn0(p0 int32, p1 bool) bool {
    v6 := GenS1{F0: 0}
    p1 = p1
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v6
    return p1
}
func genFn1() int32 {
    switch "" {
    case "a":
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 float32 = (0.0 / 1.0)
            v8 += float32((1e10 + -2.25))
            _ = fi7
            _ = v8
        }
    }
    var v9 int8 = -128
    v10 := []int{}
    _ = v10
    _ = v9
    return -2147483648
}

func main() {
    v1 := []uint32{2, 271, 2147483648, 522}
    var v2 uint8 = 255
    v3 := make(map[int]int)
    var v4 int = 500
    var v5 uint8 = 127
    v1 = append(v1, 387)
    if true {
        v11 := GenS1{F0: 2}
        v3[v4] = v4
        _ = v11
    } else {
        v12 := GenS2{F0: 893, F1: 0, F2: 2147483647}
        _ = v12
    }
    if (false || (65535 <= 0)) {
        var v13 float32 = 0.0
        _ = v13
    }
    var v14 int64 = -269
    v2 = (^(v2))
    fmt.Println(v14, v4)
    for fi15 := 0; fi15 < 14; fi15++ {
        v3[v4] = v4
        v14 += -9223372036854775808
        for fi16 := 0; fi16 < 15; fi16++ {
            fmt.Println(fi16, fi16, v5)
            v17 := GenS1{F0: 4294967295}
            v14 -= v14
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v18 := GenS1{F0: 467}
    for fi19 := 0; fi19 < 7; fi19++ {
        if true {
            var v20 bool = true
            _ = v20
        }
        fmt.Println(fi19, fi19, fi19)
        if (false && false) {
            for fi21 := 0; fi21 < 6; fi21++ {
                fmt.Println(v14, v5, v5)
                fmt.Println(v4, v5)
                var v22 int = fi21
                switch (^(fi21)) {
                case 2147483647:
                    v23 := GenS3{F0: -2147483648, F1: 32767, F2: (-(v4))}
                    _ = v23
                }
                _ = fi21
                _ = v22
            }
            v24 := GenS2{F0: 2, F1: genFn1(), F2: 386}
            if (false && false) {
                _ = len(v3)
            } else {
                v2 = v5
                v5 += (^(v2))
            }
            _ = v24
        } else {
            if ("hello" > "foo bar") {
                v14 = v14
            } else {
                v25 := GenS4{F0: -1, F1: -95}
                v14 += v14
                _ = v25
            }
            for fi26 := 0; fi26 < 7; fi26++ {
                v4 = 133
                v4 = fi19
                v2 += v2
                _ = fi26
            }
        }
        _ = fi19
    }
    v27 := v3[v4]
    v4 -= (-(v27))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
