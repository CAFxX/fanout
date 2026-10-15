package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 bool; F1 int16; F2 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 uint16; F2 int64 }

func main() {
    var v1 uint = 0
    var v2 int8 = 0
    var v3 uint32 = 65535
    v4 := []int{128}
    v2 = (v2 << (((^(v1))) & 63))
    v3 = (v3 / ((65535) | 1))
    var v5 float32 = -2.25
    if false {
        var v6 uint32 = 1
        _ = v6
    } else {
        if len(v4) > 0 {
            v4[0] = 127
        }
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := GenS1{F0: (^(v1)), F1: 1}
            for fi9 := 0; fi9 < 20; fi9++ {
                v10 := GenS2{F0: ((v1 >> ((v1) & 63)) == (v1 % ((144) | 1))), F1: -32768, F2: (^(v1))}
                var v11 int8 = (v2 % (((v2 >> ((133) & 63))) | 1))
                var v12 int8 = 127
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            v13 := GenS3{F0: v3}
            _ = fi7
            _ = v13
            _ = v8
        }
    }
    var v14 uint = v1
    v4 = append(v4, -743)
    fmt.Println(v14, v5, v2)
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := GenS2{F0: true, F1: 32767, F2: v14}
            v18 := GenS4{F0: "\n\t"}
            _ = fi16
            _ = v17
            _ = v18
        }
        var v19 int = (^(fi15))
        _ = fi15
        _ = v19
    }
    v1 -= (v1 + 216)
    _ = len(v4)
    if len(v4) > 0 {
        v4[0] = 32768
    }
    fmt.Println(v3)
    switch ("\n\t" + ("hello" + "a")) {
    case "\n\t":
        v20 := GenS1{F0: (v1 % ((1) | 1)), F1: -2147483648}
        _ = v20
    case "":
        for fi21 := 0; fi21 < 8; fi21++ {
            v22 := GenS5{F0: 1, F1: 32767, F2: 42}
            for fi23 := 0; fi23 < 20; fi23++ {
                fmt.Println(fi21, v3, v14)
                if false {
                    var v24 float64 = 0.0
                    _ = v24
                }
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
    }
    v14 -= 733
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
