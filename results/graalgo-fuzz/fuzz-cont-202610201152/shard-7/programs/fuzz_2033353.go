package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 0
    v3 := make(map[string]int)
    if true {
        v2 = (v2 << ((42) & 63))
        var v4 int64 = 0
        if (0 > -117) {
            v5 := []int{482}
            _ = v5
        } else {
            var v6 uint32 = 2147483648
            v4 -= (^(v4))
            if (!true) {
                v2 += (v2 + v2)
                switch 168 {
                case 2147483647:
                    switch 1 {
                    case 32768:
                        var v7 int16 = 32767
                        _ = v7
                    case 315:
                        switch "" {
                        case "hello":
                            var v8 int64 = (v4 & v4)
                            _ = v8
                        case "":
                            var v9 uint8 = 47
                            _ = v9
                        case "\n\t":
                            var v10 float32 = (1.5 - float32(1.5))
                            _ = v10
                        default:
                            var v11 float64 = 0.0
                            _ = v11
                        }
                    default:
                        v2 = v2
                    }
                case 157:
                    for fi12 := 0; fi12 < 5; fi12++ {
                        v4 = v4
                        v2 = v2
                        fmt.Println(v4)
                        _ = fi12
                    }
                }
            }
            _ = v6
        }
        _ = v4
    } else {
        v2 = (v2 % (((v2 & (v2 << ((0) & 63)))) | 1))
        v13 := []int{42, 239}
        _ = v13
    }
    if (-2147483648 > -172) {
        if false {
            v14 := []int{0, -1}
            fmt.Println(v2)
            _ = v14
        } else {
            var v15 float32 = 0.1
            if true {
                v15 = 3.14159
            } else {
                v16 := []int{128, -1, 127}
                v17 := GenS1{F0: 152, F1: 279}
                var v18 uint8 = 168
                _ = v16
                _ = v17
                _ = v18
            }
            var v19 float32 = (v15 * 0.0)
            _ = v15
            _ = v19
        }
    } else {
        v2 = 1
        var v20 uint32 = 505
        v20 += (v20 | 760)
        _ = v20
    }
    v21 := []int{231, -638, 101, -838}
    var v22 float32 = (0.0 + float32(float64(1e10)))
    for fi23 := 0; fi23 < 15; fi23++ {
        fmt.Println(v2)
        _ = len(v21)
        _ = fi23
    }
    switch "hello" {
    case "":
        switch 282 {
        case 0:
            delete(v1, 1)
        }
    default:
        v22 = v22
    }
    if (!(!(false || (true && (65535 >= 4294967295))))) {
        if false {
            fmt.Println(v2, v22)
            fmt.Println(v22)
            var v24 float32 = v22
            _ = v24
        }
        if (!("hello" < "\n\t")) {
            if false {
                v25 := GenS1{F0: 131, F1: 2}
                v26 := v1[9223372036854775807]
                var v27 float32 = float32(0.1)
                _ = v25
                _ = v26
                _ = v27
            } else {
                v28 := GenS2{F0: 42}
                if true {
                    if (-1 >= 2147483647) {
                        var v29 int8 = -80
                        fmt.Println(v2, v2, v22)
                        _ = v29
                    }
                    var v30 uint32 = 4294967295
                    _ = v30
                }
                _ = v28
            }
            v31 := GenS2{F0: 100}
            _ = v31
        }
    } else {
        v2 += v2
    }
    v32 := GenS3{F0: 0}
    v33 := GenS3{F0: 998}
    if (0 <= 254) {
        fmt.Println(v2)
        var v34 bool = false
        delete(v3, "")
        _ = v34
    }
    var v35 uint32 = 212
    fmt.Println(v22)
    v22 -= (0.1 / v22)
    for fi36 := 0; fi36 < 11; fi36++ {
        _ = len(v1)
        v2 = (v2 >> ((532) & 63))
        var v37 bool = ((-65 <= 0) || false)
        _ = fi36
        _ = v37
    }
    v38 := GenS3{F0: 995}
    _ = len(v21)
    v22 -= (3.14159 + 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v38
}
