package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint; F2 bool }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    if (!false) {
        for fi4 := 0; fi4 < 13; fi4++ {
            for fi5 := 0; fi5 < 1; fi5++ {
                v6 := GenS1{F0: 9223372036854775807, F1: -1, F2: ("" + "")}
                v7 := GenS2{F0: 115}
                v8 := GenS2{F0: 128}
                _ = fi5
                _ = v6
                _ = v7
                _ = v8
            }
            _ = fi4
        }
        if (18446744073709551615 <= 91) {
            var v9 float32 = (v1 / v1)
            v10 := []int{-784, 0, -21}
            var v11 int16 = 32767
            _ = v10
            _ = v11
            _ = v9
        } else {
            v12 := GenS2{F0: 42}
            _ = len(v3)
            _ = v12
        }
        if true {
            var v13 uint16 = 65535
            _ = v13
        } else {
            var v14 string = "foo bar"
            _ = v14
        }
    } else {
        v15 := GenS3{F0: 32768, F1: 200, F2: false}
        v16 := GenS2{F0: 0}
        _ = v15
        _ = v16
    }
    fmt.Println(v1)
    v17 := []int{324, 855, 127, 1}
    if false {
        fmt.Println(v1)
        var v18 string = "hello"
        _ = v18
    } else {
        v19 := v2[1]
        switch -2147483648 {
        case 127:
            if len(v17) > 0 {
                v17[0] = v19
            }
        default:
            v1 += (float32(((float64(v1) * (0.1 + (0.1 - -2.25))) - 3.14159)) * v1)
        }
        _ = v19
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 int = fi20
        var v22 int = (v21 + fi20)
        if true {
            if true {
                v23 := GenS3{F0: (-(fi20)), F1: 65535, F2: (128 == 1)}
                _ = v23
            } else {
                switch (^(v21)) {
                case 2147483647:
                    var v24 uint8 = 1
                    _ = v24
                default:
                    _ = len(v2)
                }
                if len(v17) > 0 {
                    v17[0] = 0
                }
            }
            var v25 int32 = 1
            _ = v25
        } else {
            var v26 float64 = 0.1
            var v27 bool = false
            var v28 uint = 4294967295
            _ = v26
            _ = v27
            _ = v28
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println(v1)
    var v29 uint64 = 676
    v30 := v2[uint32(v29)]
    var v31 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
