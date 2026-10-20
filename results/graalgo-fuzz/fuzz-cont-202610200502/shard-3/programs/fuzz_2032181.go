package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint32; F1 int16; F2 uint }
type GenS3 struct { F0 uint16; F1 bool; F2 int64 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 uint32, p1 uint, p2 uint32) uint64 {
    fmt.Println(p2, p2, p0)
    switch "foo bar" {
    case "hello":
        var v6 float64 = ((((float64((1.5 * (-2.25 / 3.14159))) / 1.5) / float64(1.5)) * (1e10 - 1e10)) * (-2.25 - (1.5 * ((0.1 * ((float64(((3.14159 * 3.14159) * float32(1e10))) + (1.5 * 1e10)) / -2.25)) * float64((float32((1e10 + (0.0 * (3.14159 * (float64(-2.25) / 1.0))))) + (3.14159 - (1e10 / 1e10))))))))
        _ = v6
    case "\n\t":
        var v7 bool = true
        _ = v7
    case "a":
        if (!(false || (42 < -87))) {
            if true {
                v8 := []int{534}
                p2 = p2
                p2 = p0
                _ = v8
            }
            if (false && true) {
                p1 += 65535
                for fi9 := 0; fi9 < 11; fi9++ {
                    v10 := GenS1{F0: p1, F1: true}
                    _ = fi9
                    _ = v10
                }
                if false {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        switch "" {
                        case "":
                            fmt.Println(p0)
                        default:
                            switch 4294967295 {
                            case 9223372036854775807:
                                p0 = p0
                            case -9223372036854775808:
                                p0 = (^(p0))
                            case 128:
                                p1 = (p1 * (p1 & (p1 << ((p1) & 63))))
                            }
                        }
                        switch (fi11 / ((32768) | 1)) {
                        case 9223372036854775807:
                            v12 := GenS2{F0: 42, F1: 0, F2: (p1 % (((p1 << ((p1) & 63))) | 1))}
                            _ = v12
                        case 32768:
                            switch "hello" {
                            case "":
                                var v13 int64 = -612
                                _ = v13
                            case "\n\t":
                                v14 := []int{(fi11 / ((fi11) | 1)), (^(fi11))}
                                _ = v14
                            default:
                                var v15 float64 = float64(((1e10 / 1.0) * (((0.0 + 0.0) + -2.25) / 1.5)))
                                _ = v15
                            }
                        }
                        _ = fi11
                    }
                    p0 += p0
                    if true {
                        var v16 int = -369
                        _ = v16
                    }
                } else {
                    p0 += 2
                    switch 32767 {
                    case 1:
                        v17 := GenS2{F0: 1, F1: 0, F2: (p1 + 478)}
                        _ = v17
                    case -65:
                        switch (((("\n\t" + "") + "hello") + ("a" + "foo bar")) + "hello") {
                        case "":
                            v18 := GenS3{F0: 0, F1: (!false), F2: -771}
                            _ = v18
                        case "hello":
                            var v19 float64 = 3.14159
                            _ = v19
                        }
                    }
                }
            } else {
                var v20 uint32 = 536
                p2 = p2
                v21 := GenS3{F0: 42, F1: ((^(p1)) > 0), F2: 9223372036854775807}
                _ = v20
                _ = v21
            }
            p0 += (p0 << (((p1 & p1)) & 63))
        }
    default:
        v22 := GenS1{F0: 42, F1: false}
        _ = v22
    }
    v23 := GenS4{F0: true, F1: -77}
    var v24 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v24
    return 2
}

func main() {
    var v1 int8 = 1
    var v2 int16 = 1
    v3 := []uint32{4294967295}
    var v4 int64 = -572
    var v5 uint32 = 2147483647
    v1 += -128
    if true {
        v25 := GenS5{F0: 0, F1: 0}
        _ = v25
    }
    v2 = 0
    var v26 int = 2147483647
    v3 = append(v3, v5)
    v27 := GenS1{F0: 753, F1: (42 >= 221)}
    switch v26 {
    case 4294967295:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        var v28 bool = (!false)
        _ = v28
    }
    v29 := GenS6{F0: 881}
    v5 += v5
    var v30 float32 = -2.25
    if (true && ((v26 >= 240) || (0.0 > 0.1))) {
        var v31 int16 = v2
        _ = v31
    } else {
        if (255 == 32767) {
            v3 = append(v3, 699)
            var v32 float64 = ((1.5 - 1.5) - (((float64(1e10) / (1.5 * (((float64(v30) / 1.0) / (0.0 + 3.14159)) - (-2.25 * 1.5)))) * 3.14159) + ((3.14159 + (0.0 - (1.5 / 1.0))) + float64(v30))))
            v32 = 3.14159
            _ = v32
        }
        for fi33 := 0; fi33 < 16; fi33++ {
            var v34 bool = true
            if v34 {
                fmt.Println(fi33, v1)
                var v35 uint16 = 195
                _ = v35
            }
            v36 := GenS7{F0: 757}
            _ = fi33
            _ = v34
            _ = v36
        }
    }
    fmt.Println(v5, v4)
    _ = len(v3)
    v37 := GenS1{F0: 397, F1: (!(v2 >= v2))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v4
    _ = v5
}
