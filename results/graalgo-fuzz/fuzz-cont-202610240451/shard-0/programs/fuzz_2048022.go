package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 bool }
type GenS2 struct { F0 bool; F1 uint32; F2 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint32 }

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]uint32)
    v3 := []int{246, 9223372036854775807, -422, 272}
    if true {
        var v4 float32 = float32(1.5)
        v5 := GenS1{F0: -1, F1: 262, F2: true}
        var v6 string = ("foo bar" + ((("foo bar" + ("\n\t" + "hello")) + "hello") + "foo bar"))
        _ = v4
        _ = v5
        _ = v6
    }
    if (2 >= 205) {
        v7 := GenS2{F0: false, F1: 252, F2: 667}
        _ = v7
    } else {
        v1 -= (v1 << ((932) & 63))
        if false {
            v8 := GenS2{F0: ((42 < -128) || true), F1: 2147483647, F2: -561}
            _ = v8
        }
        if len(v3) > 0 {
            v3[0] = -950
        }
    }
    var v9 float32 = -2.25
    var v10 int8 = -128
    var v11 uint64 = 0
    var v12 uint8 = (v1 & (v1 << ((42) & 63)))
    fmt.Println(v1, v11, v9)
    if false {
        v1 += (v1 & 56)
    }
    var v13 int64 = 45
    var v14 float64 = 0.0
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 string = "foo bar"
            v18 := GenS3{F0: v10}
            _ = fi16
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 1; fi19++ {
            v2["\n\t"] = uint32(v12)
            _ = fi19
        }
        if len(v3) > 0 {
            v3[0] = (fi15 * int(v1))
        }
        _ = fi15
    }
    var v20 uint8 = (^(v1))
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 uint8 = v1
        _ = fi21
        _ = v22
    }
    if false {
        fmt.Println(v14)
        fmt.Println(v11, v9, v13)
    } else {
        var v23 float64 = (float64(v9) - v14)
        var v24 int8 = (v10 * (-(v10)))
        _ = v23
        _ = v24
    }
    if (true && true) {
        fmt.Println(v13, v14, v10)
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v25 uint16 = 32767
        var v26 bool = false
        _ = v25
        _ = v26
    }
    var v27 bool = ((510 <= 2147483647) && ((1 < -161) || (!false)))
    if ((4294967295 != 32767) && v27) {
        var v28 int16 = 929
        v12 += (^(v1))
        fmt.Println(v10, v1)
        _ = v28
    } else {
        v29 := GenS4{F0: v27, F1: 2147483647}
        v1 = (^(v20))
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 bool = v27
            _ = fi30
            _ = v31
        }
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v9
}
