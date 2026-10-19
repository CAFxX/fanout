package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 bool; F2 int8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint32) int {
    v4 := []int{int(p0), 32768}
    if false {
        var v5 uint16 = 1
        for fi6 := 0; fi6 < 16; fi6++ {
            fmt.Println(p0)
            var v7 uint = 255
            var v8 uint32 = p0
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = v5
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int64 = 1
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 5; fi11++ {
            fmt.Println(fi11)
            var v12 float64 = (0.1 + 1e10)
            _ = fi11
            _ = v12
        }
        var v13 uint = 1
        _ = v13
    }
    switch 56 {
    case 128:
        for fi14 := 0; fi14 < 13; fi14++ {
            fmt.Println(p0)
            p0 += p0
            if (!false) {
                p0 = p0
            } else {
                p0 -= (p0 * p0)
                for fi15 := 0; fi15 < 3; fi15++ {
                    switch fi14 {
                    case 127:
                        var v16 float64 = ((1.5 / (1e10 + float64((((1e10 - (((float32(3.14159) - 0.0) + 1e10) * 0.0)) - ((0.1 / -2.25) / 0.1)) - (float32(1.5) - (1.5 / 3.14159)))))) - (-2.25 - 0.0))
                        _ = v16
                    }
                    v4 = append(v4, (fi15 ^ (fi15 ^ (fi15 & (fi14 & 128)))))
                    _ = fi15
                }
            }
            _ = fi14
        }
    case -732:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 32768:
        _ = len(v4)
    }
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    _ = p0
    _ = v4
    return 9223372036854775807
}

func main() {
    var v1 float64 = -2.25
    var v2 uint16 = 32767
    var v3 int8 = 127
    fmt.Println(v3)
    var v17 uint32 = 0
    for fi18 := 0; fi18 < 20; fi18++ {
        v19 := []int{(fi18 >> ((872) & 63)), int(v2), genFn0(42), (fi18 ^ (fi18 << ((1) & 63)))}
        v19 = append(v19, fi18)
        v20 := GenS1{F0: 128}
        _ = fi18
        _ = v19
        _ = v20
    }
    v2 = (v2 & v2)
    var v21 uint8 = 42
    v21 += v21
    v22 := GenS2{F0: 255, F1: (-394 < -387), F2: v3}
    v2 += 865
    for fi23 := 0; fi23 < 14; fi23++ {
        v3 += v3
        v24 := GenS3{F0: 471}
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        fmt.Println(v1, v1, v17)
        switch (fi25 << ((0) & 63)) {
        case 502:
            v21 = 128
        case 128:
            v21 = v21
        default:
            v21 -= v21
        }
        for fi26 := 0; fi26 < 4; fi26++ {
            v27 := []int{(^(fi26)), fi25, fi26, (fi26 - 2147483647)}
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    var v28 string = "hello"
    if true {
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 uint8 = 128
            fmt.Println(v21, fi29)
            v30 = (v30 >> ((255) & 63))
            _ = fi29
            _ = v30
        }
    }
    v17 = v17
    v28 = ((((v28 + "a") + (v28 + "\n\t")) + "hello") + (((v28 + v28) + "\n\t") + (("hello" + v28) + "\n\t")))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v3
}
