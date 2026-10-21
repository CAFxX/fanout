package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint8 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint8; F1 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 32767
    var v3 uint = 4294967295
    var v4 int64 = 1
    v5 := []int{4294967295, 32768, 128, -366}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (!true) {
        var v6 string = "\n\t"
        if true {
            for fi7 := 0; fi7 < 10; fi7++ {
                v4 -= v4
                _ = fi7
            }
            delete(v1, 2147483648)
        }
        var v8 uint8 = 128
        _ = v6
        _ = v8
    } else {
        var v9 int32 = 1
        if (v4 < v4) {
            v10 := GenS1{F0: 198, F1: "foo bar", F2: 255}
            v5 = append(v5, 128)
            var v11 int32 = (v9 ^ v9)
            _ = v10
            _ = v11
        } else {
            var v12 int = 32768
            v13 := GenS1{F0: 25, F1: ("foo bar" + "a"), F2: 1}
            _ = v12
            _ = v13
        }
        _ = v9
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 uint8 = 127
        _ = fi14
        _ = v15
    }
    fmt.Println(v4, v4)
    v2 -= v2
    var v16 uint32 = 356
    var v17 uint8 = 2
    v2 += (v2 + 0)
    if true {
        v18 := GenS2{F0: -17, F1: 1}
        v19 := GenS3{F0: (v17 << (((v3 >> ((uint(v16)) & 63))) & 63)), F1: (!(false || false))}
        for fi20 := 0; fi20 < 3; fi20++ {
            v3 -= (v3 - (v3 * uint(v16)))
            var v21 uint = (v3 ^ v3)
            _ = fi20
            _ = v21
        }
        _ = v18
        _ = v19
    }
    v22 := v1[(v16 >> (((v3 & (v3 << (((^(v3))) & 63)))) & 63))]
    if false {
        var v23 uint = 0
        var v24 float64 = (float64(-2.25) + float64(-2.25))
        _ = v23
        _ = v24
    } else {
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
