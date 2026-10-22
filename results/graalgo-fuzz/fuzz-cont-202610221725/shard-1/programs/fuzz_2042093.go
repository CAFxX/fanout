package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16 }
func genFn0(p0 string, p1 uint) uint32 {
    var v5 float32 = 1e10
    v6 := GenS1{F0: uint16(p1), F1: -1, F2: 255}
    var v7 uint32 = 65535
    v8 := GenS1{F0: 42, F1: 606, F2: 42}
    switch p0 {
    case "\n\t":
        v7 += (v7 ^ 2147483647)
    case "foo bar":
        if (!true) {
            v9 := []int{}
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (^(v7))
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = false
    var v3 int32 = 818
    var v4 uint64 = 4294967295
    switch -927 {
    case 293:
        v3 = v3
    default:
        if v2 {
            v10 := GenS1{F0: 492, F1: -32768, F2: 2}
            var v11 float64 = 0.0
            _ = v10
            _ = v11
        } else {
            v12 := GenS2{F0: v2}
            fmt.Println(v4)
            _ = v12
        }
    }
    fmt.Println(v3, v2)
    var v13 uint16 = 42
    var v14 uint64 = (v4 >> ((4294967295) & 63))
    switch -1 {
    case -115:
        switch "\n\t" {
        case "foo bar":
            for fi15 := 0; fi15 < 16; fi15++ {
                v4 = (v4 / (((v4 * v14)) | 1))
                _ = fi15
            }
        }
    }
    if true {
        switch 781 {
        case 42:
            v16 := GenS3{F0: ((v2 || v2) && v2)}
            _ = v16
        default:
            v17 := GenS1{F0: (v13 * (v13 | v13)), F1: 1000, F2: v13}
            _ = v17
        }
        var v18 uint16 = (v13 | v13)
        _ = v18
    } else {
        v4 = v4
        v19 := v1[2147483647]
        v20 := v1[(v19 + (v19 & (^(v19))))]
        _ = v19
        _ = v20
    }
    switch 127 {
    case 2147483647:
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 float64 = ((-2.25 * -2.25) / 1e10)
            v23 := GenS4{F0: 467}
            _ = fi21
            _ = v22
            _ = v23
        }
    case 1:
        v24 := GenS1{F0: 609, F1: 847, F2: v13}
        _ = v24
    case 32768:
        if true {
            var v25 int64 = 42
            if (false || v2) {
                var v26 uint8 = 226
                v26 -= 0
                _ = v26
            }
            v4 = v4
            _ = v25
        }
    default:
        fmt.Println(v13, v2, v3)
    }
    v27 := []int{32768, 4294967295}
    var v28 float32 = 0.0
    fmt.Println(v13, v4)
    v29 := GenS3{F0: (v2 && false)}
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    v1[uint32(v13)] = 2
    v4 += (v4 % (((v14 >> ((18446744073709551615) & 63))) | 1))
    switch 127 {
    case -2147483648:
        fmt.Println(v13, v3)
    default:
        for fi30 := 0; fi30 < 1; fi30++ {
            v31 := GenS3{F0: true}
            for fi32 := 0; fi32 < 20; fi32++ {
                fmt.Println(v28, v13)
                _ = fi32
            }
            var v33 int = fi30
            _ = fi30
            _ = v31
            _ = v33
        }
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
