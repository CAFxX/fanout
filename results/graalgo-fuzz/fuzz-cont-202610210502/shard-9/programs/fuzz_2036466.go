package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int16; F2 uint64 }
type GenS3 struct { F0 int; F1 uint }
func genFn0(p0 uint8, p1 uint32) int {
    if true {
        var v7 float64 = (0.1 - 0.0)
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 bool = true
            v10 := []int{(fi8 * (^(fi8))), -749, 4294967295}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    }
    var v11 int = -830
    p1 = (p1 + p1)
    _ = p0
    _ = p1
    _ = v11
    return v11
}

func main() {
    var v1 uint16 = 42
    var v2 int16 = -32768
    var v3 int16 = -399
    var v4 uint8 = 255
    v5 := []int{565}
    v6 := []uint32{1, 2}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v1 -= (v1 << ((411) & 63))
    v4 += (v4 / ((v4) | 1))
    v12 := GenS1{F0: 70}
    if false {
        var v13 int = 0
        v14 := GenS2{F0: 18446744073709551615, F1: v3, F2: 42}
        _ = v13
        _ = v14
    } else {
        var v15 int8 = 54
        for fi16 := 0; fi16 < 11; fi16++ {
            v3 += (v3 << ((42) & 63))
            _ = fi16
        }
        _ = v15
    }
    switch "" {
    case "hello":
        if (888 > 2147483647) {
            if (!(65535 < 589)) {
                fmt.Println(v4, v4, v3)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v1 += uint16(v4)
            }
            var v17 int64 = -9223372036854775808
            _ = v17
        }
    case "":
        var v18 uint64 = 4294967295
        _ = v18
    default:
        if len(v5) > 0 {
            v5[0] = -9223372036854775808
        }
    }
    if (float64(-2.25) != 3.14159) {
        v5 = append(v5, 0)
    }
    _ = len(v6)
    switch -2147483648 {
    case 300:
        v19 := GenS3{F0: genFn0(v4, 42), F1: 18446744073709551615}
        _ = v19
    }
    if (42 >= v3) {
        v20 := GenS1{F0: 2}
        _ = v20
    }
    var v21 int16 = 394
    if len(v6) > 0 {
        v6[0] = uint32(v1)
    }
    var v22 int32 = -2147483648
    if true {
        var v23 bool = true
        if ((!(v2 < (v2 >> ((1) & 63)))) || (!v23)) {
            var v24 int8 = -125
            _ = v24
        } else {
            v23 = v23
        }
        _ = v23
    }
    if (true || true) {
        for fi25 := 0; fi25 < 16; fi25++ {
            if (!((1e10 - float32((0.0 - (0.1 / 1.0)))) <= float32(float64(1.5)))) {
                var v26 int16 = 32767
                _ = v26
            }
            v1 = (v1 ^ 138)
            var v27 uint = 321
            _ = fi25
            _ = v27
        }
    }
    if true {
        var v28 int32 = (^(v22))
        var v29 int8 = 1
        _ = v28
        _ = v29
    } else {
        var v30 string = "a"
        if false {
            v22 -= v22
            for fi31 := 0; fi31 < 2; fi31++ {
                v32 := GenS3{F0: 1, F1: 4294967295}
                v30 = v30
                fmt.Println(v3, v22, fi31)
                _ = fi31
                _ = v32
            }
        }
        _ = v30
    }
    var v33 float64 = 1e10
    var v34 float32 = (((float32(0.0) - (((0.1 + 0.1) - 0.1) / 1e10)) - float32(v33)) / -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
