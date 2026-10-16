package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint32 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int64; F1 string; F2 int32 }
func genFn0(p0 uint32, p1 int32, p2 int64) int16 {
    for fi4 := 0; fi4 < 17; fi4++ {
        v5 := GenS1{F0: 4294967295, F1: (^(fi4)), F2: 2147483647}
        _ = fi4
        _ = v5
    }
    p2 += 9223372036854775807
    p1 = p1
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: 255, F1: (fi6 ^ fi6), F2: p0}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 float32 = 0.0
    var v2 uint8 = 42
    v3 := make(map[string]uint32)
    var v8 int16 = -448
    v9 := GenS2{F0: 849, F1: 116}
    switch "" {
    case "\n\t":
        v2 += (v2 << ((65535) & 63))
    case "foo bar":
        v10 := GenS2{F0: -2147483648, F1: 32768}
        _ = v10
    default:
        switch 0 {
        case -9223372036854775808:
            switch "" {
            case "hello":
                if false {
                    if (!(!false)) {
                        if ("" != "") {
                            v11 := GenS3{F0: 0, F1: "", F2: 1}
                            for fi12 := 0; fi12 < 15; fi12++ {
                                _ = len(v3)
                                _ = fi12
                            }
                            _ = v11
                        } else {
                            var v13 float32 = v1
                            _ = v13
                        }
                        var v14 float32 = 1.5
                        v15 := []int{-9223372036854775808, 128}
                        _ = v14
                        _ = v15
                    }
                    v16 := GenS1{F0: 0, F1: -761, F2: 1}
                    fmt.Println(v1, v8)
                    _ = v16
                }
            }
        }
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS2{F0: 666, F1: fi18}
            switch fi18 {
            case -678:
                var v20 string = ("a" + "")
                _ = v20
            case 979:
                fmt.Println(v2, v1, v1)
            }
            _ = fi18
            _ = v19
        }
        v1 += (0.1 * float32(0.0))
        _ = fi17
    }
    v21 := GenS1{F0: 65535, F1: -693, F2: 65535}
    var v22 int64 = 1
    fmt.Println(v22, v22, v22)
    switch -2147483648 {
    case 354:
        v2 = v2
    default:
        var v23 int8 = -1
        _ = v23
    }
    v1 += v1
    v1 = float32(float64(v1))
    var v24 int64 = v22
    var v25 uint = 2
    v22 -= -704
    v1 -= (((v1 + ((v1 / (v1 / 0.1)) / v1)) - float32(((float64(v1) - 1.5) - 3.14159))) / -2.25)
    var v26 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v8
    _ = v9
}
