package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() int8 {
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 uint8 = 132
            var v8 int32 = -1
            v8 = v8
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 uint64 = 595
        _ = fi5
        _ = v9
    }
    if false {
        var v10 uint32 = 65535
        _ = v10
    } else {
        var v11 int32 = 1
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 uint8 = 127
            v14 := []int{(^(fi12))}
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v11
    }
    if (!true) {
        var v15 uint64 = 0
        v15 = (v15 & (v15 - (v15 / ((818) | 1))))
        fmt.Println(v15)
        _ = v15
    }
    var v16 int32 = -1
    _ = v16
    return 100
}

func main() {
    var v1 float64 = 1e10
    var v2 uint = 224
    var v3 uint8 = 96
    var v4 int64 = -1
    var v17 uint16 = 292
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 uint16 = 694
        v3 += v3
        for fi20 := 0; fi20 < 13; fi20++ {
            fmt.Println(v1, v1)
            var v21 int64 = (v4 - (v4 ^ int64(v3)))
            _ = fi20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v3, v1, v1)
    v22 := GenS1{F0: int8(v2)}
    switch -1 {
    case 128:
        v2 += (v2 / (((^(v2))) | 1))
    case 32768:
        switch -757 {
        case 1:
            var v23 int16 = 241
            _ = v23
        case 42:
            v1 = 0.0
        case 4294967295:
            if (903 <= uint(v3)) {
                var v24 float32 = float32((v1 / v1))
                _ = v24
            }
        default:
            if true {
                var v25 uint = (v2 / (((v2 - v2)) | 1))
                _ = v25
            }
        }
    case 992:
        if false {
            v26 := []int{int(v2), 128}
            if true {
                fmt.Println(v17)
            } else {
                var v27 int32 = -2147483648
                if len(v26) > 0 {
                    v26[0] = 32768
                }
                _ = v27
            }
            var v28 uint = v2
            _ = v26
            _ = v28
        }
    }
    v1 -= v1
    fmt.Println(v1, v1)
    v29 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v4
}
