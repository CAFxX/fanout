package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int16 }
type GenS2 struct { F0 int; F1 uint16; F2 uint }

func main() {
    var v1 int = 0
    var v2 uint32 = 65535
    v3 := []int{v1}
    var v4 float32 = (((1.5 / 1e10) - ((0.1 + float32(3.14159)) / (((-2.25 * 1.5) + 3.14159) / 0.1))) + (1.5 / -2.25))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi5 := 0; fi5 < 17; fi5++ {
        v4 = 1.5
        v6 := GenS1{F0: 42, F1: (("\n\t" + "\n\t") + ""), F2: 32767}
        _ = fi5
        _ = v6
    }
    v3 = append(v3, v1)
    v7 := GenS1{F0: int64(v1), F1: ("" + ""), F2: 1000}
    var v8 bool = false
    switch 32767 {
    case -935:
        v1 -= v1
    case 1:
        fmt.Println(v4)
    default:
        _ = len(v3)
    }
    switch "" {
    case "hello":
        v8 = v8
    default:
        var v9 uint16 = 902
        _ = v9
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS1{F0: 415, F1: "a", F2: -1}
        _ = fi10
        _ = v11
    }
    if len(v3) > 0 {
        v3[0] = (v1 / ((v1) | 1))
    }
    fmt.Println(v8, v4, v2)
    v8 = v8
    for fi12 := 0; fi12 < 1; fi12++ {
        if v8 {
            var v13 bool = (36 <= 1)
            v14 := GenS2{F0: 127, F1: 826, F2: 213}
            v2 = 2147483647
            _ = v13
            _ = v14
        } else {
            for fi15 := 0; fi15 < 15; fi15++ {
                v2 = (^(v2))
                _ = fi15
            }
            for fi16 := 0; fi16 < 10; fi16++ {
                _ = len(v3)
                _ = len(v3)
                if v8 {
                    if len(v3) > 0 {
                        v3[0] = (fi16 % (((fi12 + (fi16 + (v1 - fi12)))) | 1))
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi16
            }
            fmt.Println(v8)
        }
        var v17 bool = false
        var v18 int16 = 1
        _ = fi12
        _ = v17
        _ = v18
    }
    if v8 {
        switch "" {
        case "hello":
            v19 := GenS1{F0: 42, F1: "", F2: -1}
            _ = v19
        case "":
            var v20 uint8 = 0
            _ = v20
        case "\n\t":
            for fi21 := 0; fi21 < 3; fi21++ {
                if len(v3) > 0 {
                    v3[0] = 762
                }
                switch "hello" {
                case "\n\t":
                    v22 := GenS1{F0: -428, F1: "a", F2: int16(v2)}
                    _ = v22
                }
                var v23 uint64 = 0
                _ = fi21
                _ = v23
            }
        default:
            var v24 uint32 = v2
            _ = v24
        }
        for fi25 := 0; fi25 < 13; fi25++ {
            if (1e10 < -2.25) {
                v26 := GenS2{F0: v1, F1: 678, F2: 18446744073709551615}
                for fi27 := 0; fi27 < 16; fi27++ {
                    switch (fi27 >> ((527) & 63)) {
                    case 412:
                        v1 -= fi25
                    default:
                        v1 += (v1 & (-(v1)))
                        if (!(v8 || (!v8))) {
                            var v28 int32 = 2147483647
                            fmt.Println(v4, v8, fi27)
                            var v29 int8 = 0
                            var v30 float64 = 1.5
                            switch (fi25 % ((894) | 1)) {
                            case 1:
                                v8 = v8
                            case -2147483648:
                                fmt.Println(v2)
                            case 42:
                                v4 += v4
                            }
                            v4 -= (v4 - ((1.5 + v4) * (v4 / (v4 * float32(v30)))))
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                        var v31 float32 = ((1e10 * (v4 / (-2.25 + (float32(-2.25) + (v4 - v4))))) + (((v4 + v4) / 1e10) + float32(float64(float32(1.5)))))
                        var v32 int32 = 0
                        v32 -= (v32 / (((^(v32))) | 1))
                        _ = v31
                        _ = v32
                    }
                    v33 := GenS1{F0: 0, F1: "foo bar", F2: -926}
                    _ = len(v3)
                    _ = fi27
                    _ = v33
                }
                _ = v26
            }
            _ = len(v3)
            _ = fi25
        }
    } else {
        var v34 int16 = -503
        var v35 float64 = 1.5
        _ = v34
        _ = v35
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        v1 += (^(v1))
        fmt.Println(v1)
    }
    var v36 uint64 = 235
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v7
    _ = v8
}
