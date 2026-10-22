package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int8; F1 int64; F2 uint }
type GenS3 struct { F0 uint32; F1 bool; F2 int8 }
type GenS4 struct { F0 int64; F1 string; F2 uint8 }
type GenS5 struct { F0 uint }

func main() {
    var v1 float64 = ((3.14159 * 1.5) + ((0.1 - 0.1) + 3.14159))
    v2 := make(map[uint32]uint32)
    var v3 int16 = 42
    if true {
        var v4 float32 = 1.5
        _ = v4
    }
    var v5 float32 = 1.5
    if (!false) {
        var v6 uint8 = 232
        if false {
            v6 = v6
        }
        if false {
            v7 := GenS1{F0: v3, F1: -9223372036854775808}
            _ = v7
        }
        _ = v6
    } else {
        if true {
            v8 := GenS2{F0: -128, F1: -511, F2: 0}
            _ = v8
        } else {
            _ = len(v2)
            if (0.1 > (v5 - v5)) {
                v9 := GenS2{F0: 127, F1: int64(v3), F2: 0}
                v10 := []int{}
                for fi11 := 0; fi11 < 15; fi11++ {
                    _ = len(v10)
                    v3 = (v3 * (v3 / (((v3 ^ -32768)) | 1)))
                    v12 := GenS2{F0: 127, F1: 1, F2: 996}
                    _ = fi11
                    _ = v12
                }
                _ = v10
                _ = v9
            }
            var v13 int = -251
            _ = v13
        }
    }
    v14 := []int{-344, -630}
    switch ("hello" + ("hello" + "foo bar")) {
    case "foo bar":
        fmt.Println(v1)
    }
    var v15 uint16 = 65535
    var v16 int16 = v3
    v2[2147483648] = 464
    if false {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        v17 := GenS3{F0: 42, F1: (1 < (v15 + 42)), F2: 100}
        var v18 float64 = v1
        _ = v17
        _ = v18
    } else {
        switch "foo bar" {
        case "\n\t":
            v19 := GenS4{F0: 245, F1: ("foo bar" + ("a" + "hello")), F2: 127}
            _ = v19
        }
        if false {
            v20 := GenS5{F0: 42}
            _ = v20
        }
    }
    v3 -= (v16 >> ((4294967295) & 63))
    v21 := GenS1{F0: 0, F1: 0}
    var v22 float32 = v5
    v16 -= v3
    if (-41 > -1) {
        var v23 uint8 = 128
        if true {
            var v24 uint = 65535
            v25 := v2[153]
            _ = v24
            _ = v25
        }
        v14 = append(v14, -9223372036854775808)
        _ = v23
    }
    v26 := v2[1]
    fmt.Println(v15, v1, v22)
    v1 = 1.5
    var v27 float32 = (-2.25 + 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v5
}
