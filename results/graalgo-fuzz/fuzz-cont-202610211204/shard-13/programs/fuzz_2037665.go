package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 bool; F1 uint; F2 uint }
type GenS5 struct { F0 uint32; F1 int32; F2 int }
func genFn0() uint16 {
    var v5 float64 = (1e10 + (0.0 - 0.0))
    v5 += ((v5 - 1e10) - (3.14159 - -2.25))
    v6 := []int{}
    if true {
        for fi7 := 0; fi7 < 16; fi7++ {
            if false {
                var v8 uint32 = 635
                v5 -= (float64(0.1) * v5)
                if true {
                    var v9 int = -1
                    if len(v6) > 0 {
                        v6[0] = 135
                    }
                    _ = v9
                } else {
                    var v10 int32 = 58
                    for fi11 := 0; fi11 < 16; fi11++ {
                        v10 -= (v10 & v10)
                        switch fi7 {
                        case -2147483648:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        case -9223372036854775808:
                            v12 := GenS1{F0: 0, F1: ""}
                            _ = v12
                        case 32768:
                            v5 += float64(0.0)
                        }
                        fmt.Println(v10, v8, v8)
                        _ = fi11
                    }
                    var v13 float64 = (float64(-2.25) / v5)
                    _ = v10
                    _ = v13
                }
                _ = v8
            }
            switch -380 {
            case 128:
                v14 := GenS1{F0: 1, F1: ((("hello" + "") + ("a" + ("foo bar" + "a"))) + "a")}
                _ = v14
            case 32768:
                v15 := GenS2{F0: (fi7 & fi7)}
                _ = v15
            default:
                var v16 int32 = -2147483648
                _ = v16
            }
            _ = fi7
        }
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        switch (fi17 >> ((18446744073709551615) & 63)) {
        case -55:
            v5 += v5
        case 4294967295:
            for fi18 := 0; fi18 < 17; fi18++ {
                _ = len(v6)
                var v19 string = "hello"
                _ = fi18
                _ = v19
            }
        case -2147483648:
            var v20 int64 = 9223372036854775807
            _ = v20
        default:
            v5 = 0.1
        }
        v21 := GenS3{F0: 42, F1: 210}
        fmt.Println(v5, fi17)
        _ = fi17
        _ = v21
    }
    _ = v5
    _ = v6
    return 32767
}

func main() {
    var v1 uint = 255
    var v2 int16 = 32767
    v3 := make(map[uint32]int)
    var v4 uint64 = 1
    delete(v3, 2)
    var v22 int16 = v2
    for fi23 := 0; fi23 < 11; fi23++ {
        v4 = (v4 + 281)
        _ = fi23
    }
    var v24 int16 = (-(v2))
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 float64 = -2.25
        v27 := GenS4{F0: false, F1: (v1 % (((v1 >> ((v1) & 63))) | 1)), F2: (^(v1))}
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint64 = 42
    v24 -= v2
    if false {
        v29 := []int{805, -711, -474}
        var v30 uint64 = v28
        v28 -= v28
        _ = v29
        _ = v30
    }
    v2 -= v22
    v31 := GenS5{F0: 0, F1: -1, F2: -723}
    for fi32 := 0; fi32 < 12; fi32++ {
        for fi33 := 0; fi33 < 7; fi33++ {
            var v34 float32 = (0.0 * float32(float64((0.1 + 0.1))))
            v34 += 1e10
            _ = fi33
            _ = v34
        }
        v28 -= 4294967295
        delete(v3, 2147483647)
        _ = fi32
    }
    v35 := GenS5{F0: 1, F1: int32(v24), F2: 119}
    v28 = v28
    v28 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v28
    _ = v3
    _ = v31
    _ = v35
    _ = v4
}
