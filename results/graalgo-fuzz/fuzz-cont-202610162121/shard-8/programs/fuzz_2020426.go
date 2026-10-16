package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 string; F2 uint32 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int64 }
func genFn0() int8 {
    v4 := []int{4294967295}
    if (255 > 466) {
        _ = len(v4)
        v5 := GenS1{F0: 1, F1: -515}
        _ = v5
    }
    _ = v4
    return 17
}

func main() {
    var v1 int16 = 200
    var v2 int32 = 1
    var v3 float64 = 0.0
    v6 := []int{-9223372036854775808}
    v3 += (v3 / v3)
    if len(v6) > 0 {
        v6[0] = -9223372036854775808
    }
    var v7 int16 = (v1 - v1)
    v1 = -179
    v6 = append(v6, 4294967295)
    for fi8 := 0; fi8 < 6; fi8++ {
        switch "hello" {
        case "\n\t":
            if (127 > 127) {
                for fi9 := 0; fi9 < 1; fi9++ {
                    _ = len(v6)
                    _ = fi9
                }
                fmt.Println(fi8, v3)
                for fi10 := 0; fi10 < 13; fi10++ {
                    var v11 float32 = 0.0
                    v12 := GenS1{F0: 1, F1: (-(fi8))}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
        case "a":
            v6 = append(v6, fi8)
        case "":
            v13 := GenS1{F0: 18446744073709551615, F1: (fi8 | fi8)}
            _ = v13
        default:
            v6 = append(v6, fi8)
        }
        if (!(v2 >= (v2 / (((^(v2))) | 1)))) {
            v14 := GenS1{F0: 18446744073709551615, F1: (fi8 << ((255) & 63))}
            var v15 float64 = -2.25
            var v16 uint8 = 2
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi8
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 bool = false
        fmt.Println(v1, v18, fi17)
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        if true {
            v20 := GenS2{F0: true}
            _ = v20
        } else {
            v7 += v1
            v21 := GenS2{F0: false}
            v1 -= 0
            _ = v21
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if ((!false) && false) {
            if false {
                switch (^(fi19)) {
                case 540:
                    v7 += v7
                default:
                    v22 := GenS3{F0: 523, F1: "", F2: 2}
                    _ = v22
                }
            }
        }
        _ = fi19
    }
    v6 = append(v6, 672)
    v1 -= v1
    switch ("a" + ("hello" + "hello")) {
    case "a":
        v23 := GenS4{F0: 2, F1: 547, F2: -9223372036854775808}
        _ = v23
    }
    var v24 bool = false
    v2 += v2
    var v25 int64 = 745
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v6
    _ = v7
}
