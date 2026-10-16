package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint; F2 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 int64 }

func main() {
    var v1 uint64 = 0
    var v2 uint32 = 65535
    var v3 int64 = 42
    if ((v3 - (^(v3))) > (v3 * -673)) {
        v4 := GenS1{F0: v2, F1: 0}
        v5 := []int{128, -535, -2147483648, 127}
        _ = v4
        _ = v5
    }
    v6 := GenS1{F0: 2, F1: -128}
    var v7 uint = 130
    v2 -= v2
    if ((-1 != -1) || false) {
        fmt.Println(v1, v7)
        for fi8 := 0; fi8 < 2; fi8++ {
            v3 -= int64(v7)
            _ = fi8
        }
    } else {
        v7 -= 2
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        v1 = 42
        var v10 uint8 = 42
        v11 := GenS1{F0: (v2 & (^(v2))), F1: 100}
        _ = fi9
        _ = v10
        _ = v11
    }
    v3 = (v3 - -9223372036854775808)
    var v12 uint8 = 42
    if true {
        var v13 uint64 = v1
        var v14 uint64 = v1
        _ = v13
        _ = v14
    } else {
        if false {
            v3 = 42
            v7 = v7
        }
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        if ((!(-979 <= 0)) && (((^(v1)) == (v1 / (((^(v1))) | 1))) || (true || (!(!(255 > 669)))))) {
            var v16 float32 = 3.14159
            var v17 int64 = 0
            var v18 int8 = 0
            _ = v16
            _ = v17
            _ = v18
        } else {
            v19 := GenS1{F0: (v2 >> (((^(v7))) & 63)), F1: -1}
            for fi20 := 0; fi20 < 7; fi20++ {
                if (v2 >= v2) {
                    var v21 uint = (v7 % (((v7 - 948)) | 1))
                    _ = v21
                }
                _ = fi20
            }
            v22 := []int{(fi15 | (fi15 % ((fi15) | 1))), (fi15 >> (((^(v7))) & 63)), fi15}
            _ = v19
            _ = v22
        }
        _ = fi15
    }
    var v23 uint8 = 1
    var v24 int64 = v3
    v25 := GenS2{F0: 42, F1: (v7 << (((^(v7))) & 63)), F2: -9223372036854775808}
    if false {
        var v26 float32 = 1e10
        v27 := GenS3{F0: (v3 - v24)}
        _ = v26
        _ = v27
    } else {
        v3 = (-(v24))
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        v29 := GenS1{F0: 188, F1: -128}
        for fi30 := 0; fi30 < 7; fi30++ {
            if ("" >= ("" + ("hello" + "a"))) {
                v31 := GenS4{F0: v2, F1: v24}
                _ = v31
            }
            _ = fi30
        }
        _ = fi28
        _ = v29
    }
    v23 -= (v23 | 194)
    if false {
        v32 := []int{-551, 259}
        v12 = 1
        var v33 float32 = ((float32(1.5) * -2.25) - (float32(1e10) + (((3.14159 / 1.0) + (3.14159 / 0.1)) * (3.14159 * ((1e10 / 1.5) - 0.0)))))
        _ = v32
        _ = v33
    } else {
        switch ("a" + "a") {
        case "a":
            v34 := []int{4294967295, -617, 128}
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v6
    _ = v7
}
