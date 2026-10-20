package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS2 struct { F0 int64; F1 int8; F2 uint }
type GenS3 struct { F0 bool }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []int{-316, -9223372036854775808}
    v3 := make(map[uint32]uint32)
    var v4 float64 = float64((0.1 - float32(0.0)))
    var v5 bool = (((0.0 - -2.25) / 1e10) <= 1.5)
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 string = ("hello" + "\n\t")
        var v8 int8 = 1
        if true {
            v7 = "foo bar"
            v1 = v1
        } else {
            v5 = ((1e10 * ((0.1 + float32(1.5)) + 1e10)) <= 1e10)
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if v5 {
        v1 += (v1 * -176)
        var v9 int64 = v1
        _ = v9
    }
    v10 := GenS1{F0: 127, F1: (-(v1)), F2: 179}
    _ = len(v2)
    var v11 int16 = -32768
    if (1 <= 42) {
        switch (("foo bar" + "foo bar") + ("a" + "")) {
        case "hello":
            fmt.Println(v5, v1, v11)
        default:
            var v12 string = "\n\t"
            _ = v12
        }
        if false {
            v5 = v5
            v5 = ("\n\t" <= "")
        } else {
            var v13 uint32 = 734
            _ = v13
        }
        _ = len(v3)
    } else {
        if v5 {
            v14 := GenS2{F0: v1, F1: 42, F2: 2}
            var v15 int32 = 128
            v1 += 0
            _ = v14
            _ = v15
        } else {
            if true {
                var v16 int64 = v1
                _ = v16
            }
            var v17 float64 = (v4 - float64(float32(0.0)))
            _ = len(v3)
            _ = v17
        }
        switch 127 {
        case 32768:
            var v18 float64 = (1.5 + -2.25)
            _ = v18
        case 128:
            switch -9223372036854775808 {
            case 32768:
                v19 := GenS3{F0: (v5 && (525 > 2))}
                _ = v19
            }
        default:
            var v20 float64 = 0.0
            _ = v20
        }
    }
    if (!false) {
        switch (("hello" + ("foo bar" + "\n\t")) + "hello") {
        case "hello":
            v11 += (v11 % (((v11 ^ -32768)) | 1))
        case "foo bar":
            var v21 float32 = (0.1 / 1.0)
            _ = v21
        case "":
            v4 = (1.5 * ((v4 * 0.0) + 1.5))
        default:
            v5 = (1 == -722)
        }
        v22 := v3[518]
        _ = v22
    }
    var v23 uint = 815
    var v24 float64 = v4
    var v25 bool = v5
    var v26 uint16 = 459
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
