package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint8) uint32 {
    fmt.Println(p0)
    var v6 int64 = 9223372036854775807
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := []int{fi7}
        for fi9 := 0; fi9 < 14; fi9++ {
            v6 = v6
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            if false {
                fmt.Println(fi9, fi7)
                p0 -= (^(p0))
            }
            _ = fi9
        }
        v10 := GenS1{F0: 18446744073709551615, F1: -949}
        _ = fi7
        _ = v10
        _ = v8
    }
    fmt.Println(p0, p0)
    v6 = (v6 << ((723) & 63))
    _ = p0
    _ = v6
    return 700
}
func genFn1(p0 int) uint32 {
    v11 := []int{-717, 2147483647}
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 int64 = 333
        var v14 uint = 775
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        if (!(65535 != 4294967295)) {
            p0 += (fi15 | 32767)
            if true {
                p0 += -801
            } else {
                p0 = (p0 << ((788) & 63))
            }
        } else {
            _ = len(v11)
        }
        fmt.Println(fi15, p0)
        for fi16 := 0; fi16 < 17; fi16++ {
            if len(v11) > 0 {
                v11[0] = fi15
            }
            if len(v11) > 0 {
                v11[0] = -2147483648
            }
            if true {
                var v17 int16 = 42
                fmt.Println(fi15)
                _ = v17
            }
            _ = fi16
        }
        _ = fi15
    }
    v18 := GenS1{F0: 65535, F1: 242}
    if ("foo bar" != "") {
        v19 := GenS2{F0: 42}
        _ = len(v11)
        _ = v19
    }
    _ = p0
    _ = v11
    _ = v18
    return 2147483648
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[int]bool)
    var v3 uint32 = 4294967295
    v4 := []int{}
    var v5 float32 = (float32((0.0 + ((0.1 / (float64(0.1) / -2.25)) * float64((float32((1e10 * (1.5 * (-2.25 + 1.5)))) + ((float32((-2.25 * (3.14159 * 1.5))) / (v1 + 1e10)) / (v1 - 0.1))))))) * 0.0)
    var v20 int32 = 1
    if (938 < 18446744073709551615) {
        switch 127 {
        case 680:
            var v21 uint16 = 65535
            _ = v21
        case 32768:
            v22 := v2[-436]
            _ = v22
        }
    } else {
        v23 := GenS2{F0: 4294967295}
        _ = v23
    }
    _ = len(v2)
    var v24 string = "hello"
    for fi25 := 0; fi25 < 11; fi25++ {
        v26 := GenS1{F0: 65535, F1: 32767}
        _ = fi25
        _ = v26
    }
    fmt.Println(v3, v1, v24)
    v27 := GenS1{F0: 933, F1: 32767}
    var v28 int8 = int8(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
