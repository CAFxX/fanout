package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int; F1 uint8; F2 int16 }
type GenS3 struct { F0 int32; F1 uint64; F2 bool }
type GenS4 struct { F0 int64; F1 int8; F2 int16 }
func genFn0(p0 int8, p1 int32) int64 {
    v7 := []int{32767, 237, -702, 4294967295}
    p1 -= 2147483647
    v8 := GenS1{F0: (p1 / ((p1) | 1)), F1: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 591
}

func main() {
    var v1 uint = 4294967295
    var v2 float64 = (0.1 + ((-2.25 * -2.25) - 1e10))
    v3 := make(map[uint32]uint32)
    var v4 string = "foo bar"
    var v5 uint8 = 10
    var v6 int16 = -32768
    for fi9 := 0; fi9 < 13; fi9++ {
        v4 = ("hello" + v4)
        if ((0.1 + (((1e10 / v2) * 0.0) + 1e10)) != float64(3.14159)) {
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 int16 = (v6 >> ((v1) & 63))
                _ = fi10
                _ = v11
            }
        } else {
            var v12 int32 = 1
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 int8 = -128
                v15 := GenS2{F0: int(v14), F1: v5, F2: v6}
                _ = fi13
                _ = v14
                _ = v15
            }
            v1 -= 42
            _ = v12
        }
        var v16 bool = ("" < v4)
        _ = fi9
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 int = (^(fi17))
            v4 = "a"
            if true {
                v20 := []int{v19, -9223372036854775808, fi17}
                delete(v3, 65535)
                _ = v20
            } else {
                v19 -= (-(fi17))
                var v21 float32 = -2.25
                fmt.Println(v5, v2)
                _ = v21
            }
            v2 = (v2 / v2)
            var v22 int64 = 868
            _ = fi18
            _ = v19
            _ = v22
        }
        v1 = 427
        v23 := v3[693]
        _ = fi17
        _ = v23
    }
    var v24 int = -526
    var v25 uint32 = 2147483647
    switch (v24 + v24) {
    case -813:
        var v26 float64 = ((float64((1.5 + 0.0)) / (float64(0.0) + (((1e10 + (((1.5 / v2) * v2) + float64(0.1))) / ((v2 + 0.0) / (v2 * 1e10))) / (-2.25 * -2.25)))) / (((v2 - v2) + (v2 - 0.1)) * float64(3.14159)))
        _ = v26
    case 343:
        v6 = v6
    default:
        if (0.1 != (1e10 / ((3.14159 + v2) - v2))) {
            v27 := []int{(v24 & (-(v24))), (v24 | -265), (v24 << (((^(v1))) & 63)), v24}
            v24 = v24
            _ = v27
        } else {
            v5 = v5
        }
    }
    v25 += 2147483647
    var v28 int32 = 0
    v29 := GenS3{F0: v28, F1: 0, F2: false}
    v1 += v1
    switch (((("" + v4) + "foo bar") + v4) + v4) {
    case "a":
        var v30 int = v24
        _ = v30
    case "":
        switch (v4 + "") {
        case "foo bar":
            v25 += (v25 << ((398) & 63))
        case "hello":
            fmt.Println(v4, v2, v24)
        case "a":
            v5 = v5
        default:
            v28 = (-(v28))
        }
    case "foo bar":
        v4 = ("\n\t" + v4)
    }
    var v31 uint64 = 2
    var v32 float64 = 0.0
    v33 := GenS2{F0: (^(v24)), F1: (v5 << ((v1) & 63)), F2: (v6 | 0)}
    v34 := GenS1{F0: int32(v1), F1: 32767}
    fmt.Println(v4, v31, v1)
    switch 32768 {
    case -2147483648:
        for fi35 := 0; fi35 < 14; fi35++ {
            v36 := GenS4{F0: genFn0(1, (^(v28))), F1: 74, F2: (v6 / (((v6 ^ 22)) | 1))}
            fmt.Println(v32, fi35)
            _ = fi35
            _ = v36
        }
    }
    var v37 uint32 = v25
    v38 := []int{127, (v24 + v24), v24, 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
