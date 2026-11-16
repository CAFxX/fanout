package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 float64 = 1e10
    var v2 uint32 = 2147483647
    var v3 float64 = 1.5
    var v4 int32 = -290
    var v5 uint32 = v2
    if true {
        v3 -= 0.1
        v6 := GenS1{F0: v2, F1: v2}
        v7 := []int{42, 46}
        _ = v6
        _ = v7
    }
    if (true && ("\n\t" == "foo bar")) {
        v8 := GenS1{F0: v2, F1: v5}
        var v9 int16 = 42
        v2 = (v2 | (^(v5)))
        _ = v8
        _ = v9
    }
    if (!(!(-861 >= 42))) {
        v5 = 2147483647
        if (((v3 - float64((((0.1 - (((float32(v3) + 1.5) - 0.1) * 3.14159)) / float32(1e10)) - (-2.25 / -2.25)))) / v1) < float64(1.5)) {
            v3 -= 0.1
            for fi10 := 0; fi10 < 18; fi10++ {
                if (!true) {
                    if true {
                        fmt.Println(v3, v4)
                        v11 := GenS2{F0: -38, F1: 0}
                        var v12 int32 = (v4 % (((^(v4))) | 1))
                        _ = v11
                        _ = v12
                    }
                    var v13 int8 = 30
                    v14 := []int{(fi10 + (fi10 % ((fi10) | 1))), -9223372036854775808, (fi10 - fi10)}
                    _ = v13
                    _ = v14
                }
                _ = fi10
            }
        }
    }
    fmt.Println(v4, v4)
    v15 := GenS2{F0: 42, F1: -732}
    var v16 uint = 0
    v17 := GenS1{F0: v5, F1: v5}
    if false {
        v16 -= v16
        var v18 uint16 = 65535
        var v19 float64 = v3
        _ = v18
        _ = v19
    } else {
        var v20 string = ""
        switch ((v20 + "\n\t") + v20) {
        case "a":
            if false {
                v21 := []int{782, 4294967295, -9223372036854775808}
                _ = v21
            }
        }
        v22 := []int{319, 4294967295}
        _ = v20
        _ = v22
    }
    var v23 float32 = float32(-2.25)
    if true {
        v24 := GenS2{F0: 0, F1: int64(v2)}
        var v25 uint = v16
        if (128 < 32768) {
            v26 := []int{-9223372036854775808, -94, 9223372036854775807, 32768}
            v5 -= (v5 % (((v5 ^ (v5 / ((v2) | 1)))) | 1))
            fmt.Println(v2, v4, v25)
            _ = v26
        } else {
            v16 -= (^(v25))
            v27 := GenS1{F0: v2, F1: uint32(v16)}
            v16 -= v16
            _ = v27
        }
        _ = v24
        _ = v25
    }
    v28 := []int{}
    v16 = (v16 | (^(v16)))
    for fi29 := 0; fi29 < 17; fi29++ {
        if len(v28) >= 2 {
            v28 = v28[1:len(v28)]
        }
        _ = fi29
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        fmt.Println(v3, v16, v1)
        _ = fi30
    }
    if (!false) {
        v31 := GenS3{F0: 255, F1: 921}
        _ = v31
    } else {
        v32 := GenS4{F0: v4}
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
