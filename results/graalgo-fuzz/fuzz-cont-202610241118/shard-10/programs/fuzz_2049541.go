package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int32; F2 int }
type GenS4 struct { F0 int16 }
func genFn0(p0 int8) uint64 {
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: 1000, F1: 1}
        _ = fi5
        _ = v6
    }
    if true {
        if true {
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := []int{fi7}
                p0 = p0
                _ = fi7
                _ = v8
            }
            v9 := GenS2{F0: -295}
            _ = v9
        }
    } else {
        if true {
            if (65535 > 490) {
                v10 := GenS1{F0: -715, F1: 2147483647}
                fmt.Println(p0)
                _ = v10
            }
            fmt.Println(p0)
        } else {
            for fi11 := 0; fi11 < 2; fi11++ {
                p0 += (p0 * p0)
                var v12 float32 = 3.14159
                _ = fi11
                _ = v12
            }
            var v13 int64 = -1
            _ = v13
        }
    }
    p0 = p0
    fmt.Println(p0)
    var v14 int64 = -1
    _ = p0
    _ = v14
    return 2
}

func main() {
    var v1 int8 = 42
    var v2 uint64 = 1
    var v3 int64 = 42
    var v4 int32 = 2147483647
    v15 := []int{9223372036854775807, 997}
    v16 := GenS1{F0: 1000, F1: (v4 * (v4 >> ((65535) & 63)))}
    v3 += -9223372036854775808
    var v17 float32 = 1e10
    v15 = append(v15, -637)
    var v18 float32 = (v17 + (v17 + (0.0 / ((float32((((1.5 * float64(v17)) / 3.14159) * (1.5 + -2.25))) * v17) * v17))))
    if (0 != v1) {
        _ = len(v15)
    }
    fmt.Println(v2, v1, v3)
    v2 = v2
    if false {
        v15 = append(v15, 1)
    } else {
        for fi19 := 0; fi19 < 13; fi19++ {
            var v20 uint = 42
            fmt.Println(v20, v4, v17)
            _ = len(v15)
            _ = fi19
            _ = v20
        }
        var v21 float64 = ((((0.1 + 0.1) + (0.1 + 3.14159)) * 0.1) - 1e10)
        switch 2147483647 {
        case 0:
            for fi22 := 0; fi22 < 13; fi22++ {
                v23 := GenS3{F0: v1, F1: v4, F2: 127}
                if len(v15) > 0 {
                    v15[0] = 32767
                }
                _ = fi22
                _ = v23
            }
        case 127:
            v17 += v18
        default:
            v1 += (v1 ^ v1)
        }
        _ = v21
    }
    v24 := GenS4{F0: 1000}
    if true {
        for fi25 := 0; fi25 < 8; fi25++ {
            v15 = append(v15, 397)
            if (!true) {
                switch (fi25 ^ (fi25 | (^(fi25)))) {
                case 4294967295:
                    var v26 uint8 = 128
                    _ = v26
                case 57:
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                case 127:
                    v1 += 42
                default:
                    switch fi25 {
                    case 127:
                        var v27 uint8 = 255
                        _ = v27
                    case -9223372036854775808:
                        fmt.Println(v3, v3)
                    case 2147483647:
                        var v28 uint32 = 4294967295
                        _ = v28
                    default:
                        var v29 bool = (255 == 1)
                        _ = v29
                    }
                }
                for fi30 := 0; fi30 < 9; fi30++ {
                    var v31 int16 = 1
                    var v32 uint = 1
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                switch (fi25 << ((65535) & 63)) {
                case 1:
                    v4 = -1
                }
            } else {
                switch "a" {
                case "":
                    var v33 uint = 0
                    _ = v33
                case "foo bar":
                    var v34 uint8 = 41
                    _ = v34
                case "a":
                    switch "foo bar" {
                    case "":
                        v18 = v17
                    }
                default:
                    fmt.Println(fi25, v1, v1)
                }
            }
            _ = fi25
        }
    } else {
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        v4 = 69
        v1 -= v1
    }
    v2 = genFn0(1)
    switch 32768 {
    case -9223372036854775808:
        switch 2147483647 {
        case 32768:
            var v35 int16 = 1000
            _ = v35
        case -2147483648:
            v17 += float32(0.1)
        }
    case 42:
        v15 = append(v15, -833)
    }
    fmt.Println(v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
