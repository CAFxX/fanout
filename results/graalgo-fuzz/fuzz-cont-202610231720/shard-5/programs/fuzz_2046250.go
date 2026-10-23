package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS5 struct { F0 string; F1 uint8; F2 string }
func genFn0(p0 int) uint32 {
    var v7 int32 = 409
    p0 += (^(p0))
    _ = p0
    _ = v7
    return 1
}
func genFn1() int {
    var v8 int16 = 1
    fmt.Println(v8)
    v9 := GenS1{F0: 739}
    v10 := GenS2{F0: 0, F1: (v8 * v8)}
    _ = v10
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int8 = -63
    var v2 string = "\n\t"
    v3 := make(map[int]string)
    v4 := []string{(((v2 + "") + "hello") + v2)}
    var v5 float64 = 1e10
    var v6 uint64 = 2
    var v11 uint64 = 905
    for fi12 := 0; fi12 < 19; fi12++ {
        v5 = v5
        v13 := GenS3{F0: 255}
        var v14 uint16 = 255
        _ = fi12
        _ = v13
        _ = v14
    }
    switch -92 {
    case -2147483648:
        for fi15 := 0; fi15 < 13; fi15++ {
            switch (fi15 + genFn1()) {
            case 127:
                if true {
                    var v16 float32 = 3.14159
                    fmt.Println(v16, v6, v11)
                    _ = v16
                }
            case -661:
                if len(v4) > 0 {
                    v4[0] = v2
                }
            case -9223372036854775808:
                var v17 float64 = ((v5 / v5) + ((1.5 - ((float64(-2.25) - 3.14159) + -2.25)) - (v5 * 1e10)))
                _ = v17
            }
            var v18 bool = true
            _ = fi15
            _ = v18
        }
    case -176:
        if (!(true && (int8(v6) < 75))) {
            v19 := GenS4{F0: 4294967295, F1: true, F2: 1}
            v11 = v11
            v5 = ((v5 / float64(float32(v5))) * (((3.14159 * v5) + (v5 + v5)) + v5))
            _ = v19
        }
    case -1:
        v1 += (v1 / ((v1) | 1))
    }
    v20 := GenS2{F0: (^(v6)), F1: 505}
    fmt.Println(v6, v2)
    var v21 uint = 65535
    v11 = (v6 * v6)
    fmt.Println(v1, v21, v1)
    v22 := GenS5{F0: (v2 + ((((v2 + v2) + v2) + (("hello" + "") + v2)) + "hello")), F1: 42, F2: (v2 + "a")}
    var v23 uint8 = 1
    var v24 int16 = -113
    var v25 int32 = -520
    var v26 uint8 = (v23 ^ v23)
    var v27 float64 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
