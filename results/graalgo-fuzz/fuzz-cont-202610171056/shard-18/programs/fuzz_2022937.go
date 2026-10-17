package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int64 }
func genFn0(p0 uint32) string {
    v7 := []int{128}
    var v8 bool = true
    if v8 {
        switch "" {
        case "a":
            if ((140 < -9223372036854775808) || (-973 >= 1)) {
                fmt.Println(p0, p0)
            }
        case "\n\t":
            for fi9 := 0; fi9 < 9; fi9++ {
                v8 = ((v8 || false) || v8)
                _ = fi9
            }
        default:
            var v10 float64 = float64(0.1)
            _ = v10
        }
        v8 = v8
    } else {
        if false {
            if len(v7) > 0 {
                v7[0] = 128
            }
        } else {
            v8 = true
            var v11 uint64 = 4294967295
            var v12 int8 = 127
            _ = v11
            _ = v12
        }
        var v13 int = -1
        _ = v13
    }
    _ = p0
    _ = v7
    _ = v8
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 255
    var v3 uint16 = 42
    v4 := make(map[uint32]int)
    v5 := make(map[int]string)
    v6 := make(map[string]int)
    v2 = uint8(v3)
    v14 := []int{42, 1, 42, 2147483647}
    var v15 float32 = 3.14159
    v3 = (v3 + v3)
    v16 := GenS1{F0: 65535, F1: -597, F2: 1}
    var v17 float64 = 1e10
    fmt.Println(v3)
    switch 2147483647 {
    case 0:
        switch 171 {
        case -632:
            if false {
                var v18 int16 = 1000
                fmt.Println(v18)
                var v19 uint8 = 128
                _ = v18
                _ = v19
            } else {
                v3 = 215
                var v20 int64 = -694
                _ = v20
            }
        }
    case -107:
        v17 = v17
    default:
        var v21 float64 = v17
        _ = v21
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 int64 = 0
        v17 += (v17 - (v17 - (float64((1.5 + v15)) / 0.1)))
        var v24 int16 = 1
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := v5[630]
    if true {
        var v26 int32 = -1
        _ = v26
    } else {
        var v27 int16 = -750
        delete(v4, 42)
        if false {
            fmt.Println(v2)
            if (v3 > v3) {
                var v28 float64 = v17
                var v29 float64 = 0.0
                _ = v28
                _ = v29
            } else {
                switch 1 {
                case 0:
                    for fi30 := 0; fi30 < 12; fi30++ {
                        if ((-(v27)) <= (v27 | (v27 | v27))) {
                            fmt.Println(v2, v27)
                            switch v25 {
                            case "hello":
                                var v31 uint32 = 0
                                _ = v31
                            case "\n\t":
                                var v32 uint8 = v2
                                _ = v32
                            }
                            var v33 uint = 42
                            _ = v33
                        } else {
                            fmt.Println(v27, v25, v25)
                            fmt.Println(v3, v2, v25)
                            v6[((v25 + ("" + ("hello" + v25))) + v25)] = (fi30 << ((1) & 63))
                        }
                        _ = fi30
                    }
                case 691:
                    v17 -= (v17 + (0.0 - ((((0.1 + v17) - v17) / float64(((-2.25 - (((1.5 + v15) / ((v15 / (v15 - ((v15 - v15) - (v15 / (float32(v17) / ((v15 * 0.1) - float32(3.14159))))))) * v15)) * (((((float32(v17) * 3.14159) + (((v15 / -2.25) - v15) - (1.5 - 3.14159))) * 3.14159) - v15) + float32(((v17 + 3.14159) * v17))))) + 1e10))) * -2.25)))
                }
                if false {
                    var v34 uint64 = 989
                    _ = v34
                } else {
                    fmt.Println(v25)
                    v15 -= v15
                    if len(v14) > 0 {
                        v14[0] = -207
                    }
                }
            }
        }
        _ = v27
    }
    _ = len(v14)
    delete(v4, 52)
    var v35 int8 = -128
    delete(v1, 42)
    for fi36 := 0; fi36 < 6; fi36++ {
        v37 := GenS1{F0: 65535, F1: 127, F2: int64(v2)}
        _ = fi36
        _ = v37
    }
    var v38 uint64 = 0
    switch (v25 + (v25 + ((((("hello" + v25) + v25) + v25) + v25) + v25))) {
    case "":
        v2 -= v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
