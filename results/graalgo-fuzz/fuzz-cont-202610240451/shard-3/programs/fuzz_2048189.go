package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0() bool {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int16 = 0
        for fi9 := 0; fi9 < 12; fi9++ {
            fmt.Println(v8, v8, v8)
            v8 = (v8 ^ (v8 + -1))
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 int8 = 1
        var v12 float32 = float32(0.1)
        var v13 int32 = 287
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    switch 627 {
    case -2147483648:
        var v14 int16 = 0
        _ = v14
    case 42:
        v15 := GenS1{F0: 42, F1: 1, F2: -1}
        _ = v15
    case -9223372036854775808:
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := []int{-1, (fi16 | fi16), (^(fi16))}
            v18 := GenS2{F0: 39, F1: 42}
            _ = fi16
            _ = v17
            _ = v18
        }
    default:
        switch "\n\t" {
        case "foo bar":
            if (255 >= 0) {
                var v19 int8 = -123
                v19 = 100
                fmt.Println(v19)
                _ = v19
            }
        default:
            var v20 float32 = 1e10
            _ = v20
        }
    }
    v21 := []int{0, 32767}
    _ = v21
    return false
}

func main() {
    var v1 uint = 4294967295
    var v2 uint8 = 128
    var v3 uint32 = 2147483647
    var v4 float32 = 0.0
    v5 := []int{9223372036854775807, -730}
    v6 := []int{}
    v3 += (^(v3))
    for fi22 := 0; fi22 < 4; fi22++ {
        if ((856 >= 266) && (1 < -128)) {
            v2 += 255
            if false {
                v23 := GenS2{F0: 65535, F1: (v3 << ((4294967295) & 63))}
                var v24 int8 = -128
                _ = v23
                _ = v24
            } else {
                var v25 uint32 = v3
                var v26 float32 = ((((v4 / v4) + v4) - 1e10) / v4)
                _ = v25
                _ = v26
            }
            fmt.Println(v4, v2)
        }
        var v27 int32 = 0
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi22
        _ = v27
    }
    if false {
        if ((v1 / (((v1 >> (((v1 + (^(v1)))) & 63))) | 1)) >= v1) {
            var v28 float32 = v4
            var v29 int32 = 0
            var v30 uint8 = (v2 >> ((v1) & 63))
            _ = v28
            _ = v29
            _ = v30
        } else {
            for fi31 := 0; fi31 < 19; fi31++ {
                var v32 uint = v1
                _ = fi31
                _ = v32
            }
            if len(v6) > 0 {
                v6[0] = 1
            }
            var v33 int16 = 42
            _ = v33
        }
        v34 := GenS1{F0: -2147483648, F1: v3, F2: -16}
        v2 += uint8(v3)
        _ = v34
    }
    for fi35 := 0; fi35 < 17; fi35++ {
        v3 = 1
        var v36 bool = true
        _ = fi35
        _ = v36
    }
    var v37 uint = 2
    v38 := GenS2{F0: (v3 & (v3 & (v3 % (((v3 >> ((v1) & 63))) | 1)))), F1: (^(v3))}
    var v39 string = ("foo bar" + ((("a" + ("foo bar" + "hello")) + "foo bar") + ""))
    v40 := GenS2{F0: v3, F1: (^(v3))}
    v41 := GenS1{F0: 2147483647, F1: (v3 / (((v3 + v3)) | 1)), F2: 100}
    fmt.Println(v3, v1)
    v42 := GenS1{F0: 0, F1: (v3 | 2147483648), F2: 110}
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    v3 = v3
    for fi43 := 0; fi43 < 1; fi43++ {
        v3 += v3
        _ = fi43
    }
    fmt.Println(v37, v3, v4)
    var v44 uint32 = v3
    var v45 float32 = 3.14159
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
