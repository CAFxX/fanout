package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 bool; F2 string }
func genFn0() int16 {
    var v7 float64 = -2.25
    fmt.Println(v7)
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint16 = 32767
        _ = fi8
        _ = v9
    }
    var v10 int32 = 0
    _ = v10
    _ = v7
    return -32768
}

func main() {
    v1 := []int{}
    var v2 uint32 = 51
    v3 := make(map[uint32]uint32)
    var v4 int8 = -1
    var v5 uint16 = 1
    v6 := make(map[int]string)
    switch -9223372036854775808 {
    case 1:
        if len(v1) > 0 {
            v1[0] = -566
        }
    default:
        switch 9223372036854775807 {
        case -2147483648:
            var v11 uint32 = (^(v2))
            _ = v11
        }
    }
    delete(v6, -376)
    var v12 float64 = (((float64(0.1) / -2.25) * 1.5) / ((float64(3.14159) * float64(-2.25)) - float64(0.0)))
    if true {
        var v13 uint = 502
        var v14 uint16 = (v5 >> ((v13) & 63))
        _ = v13
        _ = v14
    } else {
        var v15 float32 = 0.1
        _ = v15
    }
    v16 := GenS1{F0: v2}
    var v17 uint64 = 176
    var v18 uint = 18446744073709551615
    for fi19 := 0; fi19 < 7; fi19++ {
        for fi20 := 0; fi20 < 5; fi20++ {
            if false {
                v5 = (v5 << ((v18) & 63))
                v4 -= v4
                if false {
                    var v21 float64 = 1.5
                    _ = v21
                } else {
                    v22 := GenS1{F0: v2}
                    switch (^(fi20)) {
                    case -137:
                        var v23 uint = (^(v18))
                        _ = v23
                    case 9223372036854775807:
                        v4 -= (v4 << ((v18) & 63))
                    default:
                        var v24 uint32 = v2
                        if (v12 == ((float64(float32(v12)) * (((v12 - 0.0) - (1.5 / v12)) + (float64((0.0 - 0.0)) * 0.1))) * float64((1.5 / 1.5)))) {
                            fmt.Println(v18)
                            if len(v1) > 0 {
                                v1[0] = fi20
                            }
                        }
                        _ = v24
                    }
                    if false {
                        v25 := v3[(v2 ^ (^(v2)))]
                        switch (^(fi19)) {
                        case 127:
                            switch fi19 {
                            case 1:
                                v26 := GenS1{F0: v25}
                                _ = v26
                            case 32768:
                                var v27 uint8 = 255
                                _ = v27
                            }
                        }
                        _ = v25
                    }
                    var v28 int = 1
                    fmt.Println(v18, v12)
                    fmt.Println(v5, v28, v4)
                    _ = v22
                    _ = v28
                }
            } else {
                var v29 int16 = 32767
                _ = v29
            }
            _ = fi20
        }
        delete(v3, (v2 >> ((519) & 63)))
        _ = fi19
    }
    if true {
        v30 := GenS2{F0: 42, F1: (65535 == (v2 * (^(v2)))), F2: "foo bar"}
        fmt.Println(v17, v12, v2)
        var v31 string = "\n\t"
        _ = v30
        _ = v31
    }
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
