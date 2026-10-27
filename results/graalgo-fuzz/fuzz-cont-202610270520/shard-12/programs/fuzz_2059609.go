package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint16; F1 int; F2 int16 }
type GenS6 struct { F0 int8; F1 int }
func genFn0() int64 {
    if true {
        var v4 int32 = 60
        v5 := GenS1{F0: -32768}
        var v6 int16 = -217
        _ = v4
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 14; fi7++ {
            fmt.Println(fi7)
            _ = fi7
        }
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int64 = -863
            v10 := []int{-1, (-(fi8))}
            v11 := GenS1{F0: 1}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    }
    if true {
        switch "hello" {
        case "":
            var v12 int = 128
            _ = v12
        case "\n\t":
            v13 := GenS1{F0: 42}
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 uint = 4294967295
            _ = fi14
            _ = v15
        }
        var v16 int = -368
        switch (^(v16)) {
        case 2147483647:
            v17 := []int{2147483647, v16, -886, (v16 % (((v16 / (((v16 ^ v16)) | 1))) | 1))}
            _ = v17
        }
        _ = v16
    }
    var v18 float32 = 0.0
    _ = v18
    return -9223372036854775808
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint16 = 2
    var v3 bool = true
    v19 := []int{127, 42}
    fmt.Println(v2)
    if (!v3) {
        v3 = ((3.14159 + 1e10) != 1.5)
        v1 = 2
    }
    v20 := GenS2{F0: 9223372036854775807, F1: 33, F2: 1}
    for fi21 := 0; fi21 < 2; fi21++ {
        v22 := GenS1{F0: -1}
        fmt.Println(fi21)
        _ = fi21
        _ = v22
    }
    v23 := GenS2{F0: 192, F1: 42, F2: 127}
    var v24 float32 = ((0.0 / 3.14159) - -2.25)
    if (0 <= genFn0()) {
        v1 += (v1 / (((v1 * 0)) | 1))
        if false {
            v25 := GenS3{F0: -160, F1: 0}
            var v26 int64 = -9223372036854775808
            v3 = ("a" <= "")
            _ = v25
            _ = v26
        }
    } else {
        v27 := GenS2{F0: genFn0(), F1: 0, F2: 74}
        v3 = (v3 && (554 == 1))
        _ = v27
    }
    fmt.Println(v24)
    for fi28 := 0; fi28 < 19; fi28++ {
        v29 := GenS4{F0: -1, F1: (v2 | v2)}
        _ = fi28
        _ = v29
    }
    v30 := GenS5{F0: v2, F1: -9223372036854775808, F2: 1}
    v24 = (((float32(((0.0 * -2.25) / (((1.5 * 1e10) - 0.1) - 1e10))) - float32((((float64((v24 * 0.0)) - (((((1e10 + 3.14159) / (1e10 - float64(3.14159))) * (1e10 * (-2.25 + float64(v24)))) - 0.1) / float64(3.14159))) * 3.14159) / (float64(v24) / (1e10 - 3.14159))))) - 0.0) + (v24 / float32((((3.14159 / (float64(v24) / 3.14159)) / float64(v24)) * float64(((v24 * (float32(1e10) - ((-2.25 / 1.5) * v24))) * (3.14159 * ((((v24 * (((((v24 * v24) / float32(0.1)) - v24) - v24) - (((v24 - v24) - (v24 - float32((1e10 / 0.1)))) - ((1e10 * ((1e10 + 3.14159) + 1.5)) - 1e10)))) * (v24 - v24)) + v24) - v24))))))))
    v31 := GenS6{F0: -102, F1: 527}
    fmt.Println(v3, v2)
    for fi32 := 0; fi32 < 11; fi32++ {
        switch "" {
        case "hello":
            v1 -= 2
        }
        _ = fi32
    }
    fmt.Println(v24, v1, v2)
    for fi33 := 0; fi33 < 9; fi33++ {
        switch fi33 {
        case 1:
            v1 = v1
        case 42:
            v34 := GenS3{F0: -2147483648, F1: -2147483648}
            _ = v34
        }
        v35 := GenS2{F0: -9223372036854775808, F1: uint8(v2), F2: 2}
        v3 = v3
        _ = fi33
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v31
}
