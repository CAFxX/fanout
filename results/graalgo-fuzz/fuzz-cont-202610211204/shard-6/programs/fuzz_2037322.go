package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 bool; F2 int32 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint; F1 int32; F2 uint }
func genFn0(p0 uint32) int8 {
    v4 := GenS1{F0: 65535}
    v5 := GenS1{F0: 2}
    switch "" {
    case "a":
        if false {
            p0 += p0
            var v6 int16 = -32768
            for fi7 := 0; fi7 < 12; fi7++ {
                v8 := []int{}
                var v9 int32 = 0
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = v6
        } else {
            var v10 bool = true
            _ = v10
        }
    case "foo bar":
        for fi11 := 0; fi11 < 20; fi11++ {
            p0 = p0
            _ = fi11
        }
    }
    _ = p0
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 uint32 = 192
    var v2 uint16 = 2
    var v3 uint16 = 202
    if false {
        v12 := GenS1{F0: 2}
        _ = v12
    } else {
        fmt.Println(v3, v3)
        if true {
            fmt.Println(v3, v2)
        }
    }
    fmt.Println(v3, v3)
    v3 -= (v2 | (v2 << ((497) & 63)))
    v13 := GenS2{F0: v3, F1: true, F2: 562}
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 uint64 = 0
        _ = fi14
        _ = v15
    }
    v3 += 65535
    if (uint8(v3) != 0) {
        v3 += v2
    }
    if false {
        var v16 uint16 = v3
        if true {
            if (true || (!(42 < 2))) {
                v17 := []int{-569}
                _ = v17
            }
            switch 9223372036854775807 {
            case 0:
                if true {
                    v18 := []int{}
                    _ = v18
                } else {
                    var v19 float64 = (float64(3.14159) + ((-2.25 + 0.1) + 1.5))
                    if (!true) {
                        if true {
                            switch 2147483647 {
                            case -172:
                                for fi20 := 0; fi20 < 12; fi20++ {
                                    var v21 uint = 65535
                                    v22 := GenS2{F0: (^(v16)), F1: true, F2: -1}
                                    var v23 bool = false
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                            default:
                                fmt.Println(v3)
                            }
                        }
                        var v24 uint32 = 0
                        _ = v24
                    } else {
                        v25 := GenS3{F0: 313, F1: (v16 - (v2 ^ v16))}
                        _ = v25
                    }
                    _ = v19
                }
            }
        } else {
            v26 := []int{127, 0, -917}
            _ = v26
        }
        v27 := GenS4{F0: 33, F1: 0, F2: 255}
        _ = v16
        _ = v27
    }
    var v28 int = 42
    v28 -= v28
    var v29 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
}
