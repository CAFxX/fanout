package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int8; F1 int8; F2 int16 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 int; F1 uint }
func genFn0() bool {
    var v7 int16 = -1
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint8 = 2
        v10 := GenS1{F0: 0, F1: (v9 * (v9 * (v9 & 127)))}
        _ = fi8
        _ = v10
        _ = v9
    }
    v7 = (v7 % (((v7 * v7)) | 1))
    var v11 uint = 4294967295
    switch "foo bar" {
    case "":
        if ((false || false) && false) {
            v11 = 350
            switch 127 {
            case 42:
                v12 := GenS2{F0: 0, F1: 13, F2: (v7 & (v7 ^ v7))}
                _ = v12
            case 781:
                for fi13 := 0; fi13 < 8; fi13++ {
                    v7 += 32767
                    v14 := GenS3{F0: fi13, F1: -1}
                    _ = fi13
                    _ = v14
                }
            }
            var v15 uint8 = 139
            _ = v15
        } else {
            var v16 int64 = 271
            _ = v16
        }
    default:
        if false {
            v17 := []int{2147483647, 643, -9223372036854775808, 679}
            _ = v17
        } else {
            v7 -= v7
        }
    }
    _ = v11
    _ = v7
    return (152 >= 0)
}

func main() {
    var v1 int32 = -722
    v2 := []int{127, -9223372036854775808}
    v3 := make(map[string]uint32)
    var v4 int64 = -74
    var v5 int64 = 42
    var v6 int64 = (v4 + 1)
    v5 -= (-(v5))
    for fi18 := 0; fi18 < 9; fi18++ {
        switch fi18 {
        case 716:
            fmt.Println(v4, v1)
        }
        if true {
            var v19 int16 = 1
            v1 = (v1 * (-(v1)))
            _ = v19
        } else {
            v20 := GenS1{F0: v4, F1: 3}
            v21 := GenS3{F0: 128, F1: fi18}
            _ = v20
            _ = v21
        }
        switch (((("\n\t" + "a") + ("" + ("\n\t" + ("hello" + ("hello" + (("\n\t" + "hello") + (("\n\t" + ("foo bar" + "foo bar")) + ""))))))) + "\n\t") + "hello") {
        case "foo bar":
            var v22 int = fi18
            _ = v22
        case "":
            v23 := GenS4{F0: (v4 / (((-(v4))) | 1)), F1: -2147483648}
            _ = v23
        }
        _ = fi18
    }
    v24 := v3["hello"]
    var v25 uint8 = 128
    v26 := GenS3{F0: 952, F1: 0}
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 float64 = float64(1e10)
        fmt.Println(v28, v5)
        _ = fi27
        _ = v28
    }
    if false {
        fmt.Println(v6, v1, v25)
        v29 := GenS3{F0: 459, F1: 977}
        _ = v29
    }
    switch ("a" + "hello") {
    case "\n\t":
        _ = len(v2)
    case "foo bar":
        v2 = append(v2, 42)
    case "a":
        v30 := GenS4{F0: (v5 & (v4 ^ v6)), F1: v1}
        _ = v30
    }
    fmt.Println(v25, v5, v6)
    var v31 int32 = (v1 * 416)
    var v32 string = "hello"
    v33 := GenS5{F0: -1, F1: 1}
    v32 = v32
    var v34 uint = 229
    for fi35 := 0; fi35 < 9; fi35++ {
        for fi36 := 0; fi36 < 19; fi36++ {
            switch (-(fi35)) {
            case -9:
                v37 := v3[v32]
                _ = v37
            default:
                if len(v2) > 0 {
                    v2[0] = (fi35 / (((^(fi35))) | 1))
                }
            }
            _ = len(v3)
            v31 = (^(v1))
            var v38 int32 = -1
            _ = fi36
            _ = v38
        }
        _ = len(v2)
        var v39 uint16 = 337
        _ = fi35
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
