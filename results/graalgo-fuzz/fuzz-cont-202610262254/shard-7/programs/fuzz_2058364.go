package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int; F1 uint8; F2 uint32 }

func main() {
    var v1 uint8 = 42
    var v2 int16 = 1000
    v3 := make(map[uint32]string)
    v4 := []int{-210, 703}
    if (!false) {
        var v5 float32 = 0.0
        for fi6 := 0; fi6 < 20; fi6++ {
            for fi7 := 0; fi7 < 3; fi7++ {
                v5 += 1.5
                v8 := GenS1{F0: (v2 >> ((18446744073709551615) & 63))}
                _ = fi7
                _ = v8
            }
            var v9 int = fi6
            _ = fi6
            _ = v9
        }
        _ = v5
    } else {
        v4 = append(v4, 127)
        if true {
            var v10 bool = true
            fmt.Println(v1)
            _ = v10
        } else {
            v2 += (v2 & v2)
            if ((-(v2)) <= (^(v2))) {
                var v11 bool = (false && false)
                v11 = (1.5 == 0.1)
                fmt.Println(v11, v1, v2)
                _ = v11
            }
            v12 := GenS2{F0: "\n\t"}
            _ = v12
        }
        v2 = v2
    }
    v3[563] = ("foo bar" + "foo bar")
    switch "" {
    case "foo bar":
        v13 := v3[2]
        _ = v13
    case "hello":
        v14 := v3[4294967295]
        _ = v14
    case "":
        v1 = v1
    default:
        v4 = append(v4, 32767)
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        v1 = 127
        fmt.Println(v2, v2)
        if false {
            v16 := GenS3{F0: "", F1: 255}
            v2 += 42
            v17 := GenS4{F0: (fi15 * (fi15 >> ((256) & 63))), F1: (v1 * (v1 << ((18446744073709551615) & 63))), F2: uint32(v1)}
            _ = v16
            _ = v17
        } else {
            var v18 uint16 = 440
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v18
        }
        _ = fi15
    }
    v1 -= 109
    var v19 uint8 = (v1 ^ 2)
    for fi20 := 0; fi20 < 6; fi20++ {
        if len(v4) > 0 {
            v4[0] = fi20
        }
        fmt.Println(v2, v1, v19)
        _ = fi20
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v22 float64 = 0.0
        for fi23 := 0; fi23 < 15; fi23++ {
            _ = len(v3)
            var v24 int64 = 1
            switch fi23 {
            case 127:
                v2 = v2
            case 0:
                fmt.Println(v1, v1)
            }
            var v25 float64 = v22
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi21
        _ = v22
    }
    v2 = (v2 + v2)
    if (!(-128 > 43)) {
        v19 = 0
        v19 += (v1 + v19)
        var v26 uint = 1
        _ = v26
    } else {
        v27 := GenS2{F0: ("a" + "a")}
        _ = v27
    }
    var v28 bool = true
    switch "" {
    case "":
        var v29 int8 = 35
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
