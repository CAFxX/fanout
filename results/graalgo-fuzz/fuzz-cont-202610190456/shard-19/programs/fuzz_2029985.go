package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint16; F1 uint64 }
func genFn0(p0 bool) int32 {
    p0 = false
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint = 0
        _ = fi7
        _ = v8
    }
    var v9 int64 = 623
    p0 = (0 == 0)
    _ = p0
    _ = v9
    return 1
}
func genFn1() int16 {
    var v10 float64 = 0.1
    v10 += (float64((1e10 * ((-2.25 - (-2.25 / (float32(v10) - (1.5 * (float32((-2.25 + 0.0)) * 3.14159))))) + 1.5))) + v10)
    _ = v10
    return 597
}

func main() {
    var v1 uint8 = 128
    var v2 uint32 = 812
    v3 := make(map[string]uint32)
    v4 := make(map[string]string)
    var v5 uint = uint(v1)
    var v6 int = 162
    if (!(!true)) {
        switch int(v1) {
        case 127:
            v11 := []int{(v6 / ((-559) | 1)), v6}
            _ = v11
        default:
            var v12 uint = (v5 << ((v5) & 63))
            _ = v12
        }
        v6 = v6
        v6 = v6
    } else {
        if false {
            v13 := GenS1{F0: ("\n\t" + "\n\t"), F1: v1}
            _ = v13
        }
        if (-1 <= -9223372036854775808) {
            fmt.Println(v2, v2)
            v14 := GenS2{F0: ((0 <= -1) || false), F1: 0}
            for fi15 := 0; fi15 < 14; fi15++ {
                v16 := []int{}
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 int32 = -2147483648
                    _ = fi17
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
            _ = v14
        } else {
            switch (v6 + (v6 % ((32767) | 1))) {
            case 651:
                var v19 float64 = 1e10
                _ = v19
            case -603:
                var v20 float64 = 0.1
                _ = v20
            case 42:
                v21 := GenS1{F0: "a", F1: v1}
                _ = v21
            }
            v22 := []int{v6, v6, (v6 | (^(v6))), v6}
            _ = v22
        }
        v2 -= v2
    }
    v23 := v4[""]
    fmt.Println(v1, v2, v6)
    v1 = 220
    for fi24 := 0; fi24 < 4; fi24++ {
        var v25 uint64 = 0
        _ = fi24
        _ = v25
    }
    switch v6 {
    case 647:
        fmt.Println(v5)
    case 9223372036854775807:
        var v26 int64 = -67
        _ = v26
    case 128:
        v2 = (v2 >> ((v5) & 63))
    default:
        if false {
            v27 := v3[(("" + v23) + "a")]
            v28 := GenS1{F0: ("hello" + (("hello" + ("" + v23)) + "a")), F1: (^(v1))}
            fmt.Println(v5, v5, v6)
            _ = v27
            _ = v28
        } else {
            v1 -= v1
            v5 = v5
        }
    }
    if true {
        v29 := GenS3{F0: 392, F1: 42}
        v1 -= (v1 ^ v1)
        _ = v29
    } else {
        v6 -= (v6 ^ (v6 >> (((v5 / (((v5 - v5)) | 1))) & 63)))
        for fi30 := 0; fi30 < 16; fi30++ {
            v31 := GenS4{F0: 1, F1: 4294967295}
            v32 := GenS1{F0: "a", F1: (^(v1))}
            var v33 uint64 = 940
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
        var v34 int16 = -1
        _ = v34
    }
    for fi35 := 0; fi35 < 17; fi35++ {
        _ = len(v3)
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
