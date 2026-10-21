package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int32 }
type GenS2 struct { F0 uint8; F1 int64; F2 int16 }
func genFn0(p0 uint16, p1 int, p2 int64) int64 {
    p2 -= 0
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}

func main() {
    var v1 string = ""
    var v2 int16 = -960
    var v3 bool = (float64(float32(float64(((((-2.25 / 1e10) + (3.14159 - -2.25)) / float32(3.14159)) * -2.25)))) >= (1e10 + 1e10))
    var v4 uint = 42
    v5 := make(map[int]uint32)
    var v6 bool = (-618 >= -269)
    for fi7 := 0; fi7 < 9; fi7++ {
        v5[fi7] = 0
        v8 := []int{(fi7 * (^(fi7)))}
        v9 := GenS1{F0: 2147483648, F1: fi7, F2: -492}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v6 = v3
        v11 := v5[32767]
        if (32767 <= 764) {
            var v12 uint16 = 42
            switch v1 {
            case "\n\t":
                var v13 uint32 = 2
                _ = v13
            }
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    if v3 {
        fmt.Println(v2, v1, v3)
    } else {
        for fi14 := 0; fi14 < 15; fi14++ {
            if ("foo bar" != "foo bar") {
                for fi15 := 0; fi15 < 17; fi15++ {
                    v2 = 1000
                    v16 := GenS1{F0: 1, F1: fi15, F2: -566}
                    var v17 string = "hello"
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            _ = fi14
        }
    }
    v18 := []int{32768, 0}
    for fi19 := 0; fi19 < 12; fi19++ {
        v6 = v6
        _ = fi19
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 int8 = 42
        if ((4294967295 > 0) && false) {
            var v22 int8 = v21
            var v23 float32 = 0.0
            v24 := GenS1{F0: 398, F1: -273, F2: -2147483648}
            _ = v22
            _ = v23
            _ = v24
        }
        v21 = (v21 | (v21 << (((v4 >> (((^(v4))) & 63))) & 63)))
        _ = fi20
        _ = v21
    }
    v25 := GenS1{F0: 1, F1: 2147483647, F2: 2147483647}
    if ((32767 >= -517) || (229 != 837)) {
        var v26 int16 = v2
        var v27 bool = (!true)
        var v28 uint16 = 1
        _ = v26
        _ = v27
        _ = v28
    } else {
        v1 = (v1 + (((v1 + (v1 + ((v1 + (("\n\t" + v1) + v1)) + "foo bar"))) + (v1 + v1)) + (v1 + v1)))
        v29 := GenS1{F0: 0, F1: -531, F2: 2147483647}
        _ = v29
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        for fi31 := 0; fi31 < 13; fi31++ {
            var v32 int8 = -1
            _ = fi31
            _ = v32
        }
        for fi33 := 0; fi33 < 6; fi33++ {
            _ = len(v5)
            v2 = (^(v2))
            var v34 uint64 = 42
            _ = fi33
            _ = v34
        }
        _ = fi30
    }
    switch 1 {
    case 32767:
        v35 := GenS2{F0: 0, F1: -9223372036854775808, F2: 0}
        _ = v35
    case -9223372036854775808:
        var v36 uint8 = 255
        _ = v36
    case 0:
        var v37 uint32 = 1
        _ = v37
    default:
        fmt.Println(v1, v6, v1)
    }
    v6 = (genFn0(255, 128, -447) == -9223372036854775808)
    fmt.Println(v3, v6, v1)
    var v38 int16 = v2
    v39 := GenS2{F0: 1, F1: 42, F2: (-(v2))}
    var v40 float32 = float32(0.0)
    switch ("hello" + v1) {
    case "a":
        for fi41 := 0; fi41 < 9; fi41++ {
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            v40 -= v40
            var v42 int32 = 781
            _ = fi41
            _ = v42
        }
    case "foo bar":
        v40 = v40
    default:
        fmt.Println(v2, v40, v40)
    }
    var v43 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v5
    _ = v6
}
