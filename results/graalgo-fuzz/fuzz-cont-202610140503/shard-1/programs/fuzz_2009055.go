package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int16 }
type GenS2 struct { F0 uint8; F1 int16; F2 int32 }
type GenS3 struct { F0 uint; F1 uint16 }

func main() {
    var v1 int8 = -128
    var v2 int64 = 0
    var v3 uint32 = 2147483647
    v4 := []int{-471, 32768, 1}
    var v5 bool = false
    v3 += (^(v3))
    if v5 {
        if v5 {
            v2 = (v2 & v2)
            if v5 {
                if (!(!true)) {
                    fmt.Println(v3)
                    var v6 uint = 391
                    _ = v6
                }
            }
        } else {
            v3 = v3
            v1 = v1
        }
        v7 := GenS1{F0: 128, F1: 0, F2: -700}
        v5 = v5
        _ = v7
    }
    var v8 string = "hello"
    if (1 != 100) {
        v5 = (310 != 1)
        for fi9 := 0; fi9 < 11; fi9++ {
            for fi10 := 0; fi10 < 13; fi10++ {
                v4 = append(v4, fi10)
                if ("\n\t" <= (((v8 + ((((v8 + v8) + (("" + v8) + v8)) + ((v8 + (v8 + "a")) + (v8 + v8))) + v8)) + (((v8 + "foo bar") + v8) + ((v8 + ("\n\t" + v8)) + (("" + v8) + v8)))) + v8)) {
                    if (!((fi10 + (fi9 / ((-741) | 1))) <= (^(fi9)))) {
                        v11 := GenS1{F0: 168, F1: 0, F2: -32768}
                        v1 += v1
                        _ = v11
                    }
                    var v12 uint8 = 255
                    v2 = (v2 >> ((963) & 63))
                    _ = len(v4)
                    var v13 int16 = 0
                    fmt.Println(v2, v3, v8)
                    _ = v12
                    _ = v13
                } else {
                    if false {
                        v14 := GenS2{F0: 4, F1: 1000, F2: -1}
                        var v15 uint = 255
                        var v16 int8 = v1
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    v1 = 120
                    fmt.Println(v5, v5)
                }
                _ = fi10
            }
            _ = fi9
        }
    }
    v1 -= v1
    var v17 uint32 = (v3 ^ v3)
    for fi18 := 0; fi18 < 12; fi18++ {
        v8 = "foo bar"
        fmt.Println(fi18, fi18, v8)
        switch (("foo bar" + v8) + v8) {
        case "\n\t":
            var v19 bool = (!v5)
            _ = v19
        case "foo bar":
            var v20 uint = 1
            _ = v20
        }
        _ = fi18
    }
    v21 := GenS3{F0: 0, F1: 0}
    switch 42 {
    case -1:
        var v22 uint16 = 32767
        _ = v22
    case 0:
        v23 := GenS1{F0: 127, F1: 131, F2: 1}
        _ = v23
    default:
        fmt.Println(v17, v17, v3)
    }
    var v24 uint64 = 18446744073709551615
    if len(v4) > 0 {
        v4[0] = -382
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
