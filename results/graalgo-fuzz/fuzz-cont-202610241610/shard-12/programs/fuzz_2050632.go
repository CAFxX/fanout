package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int64 }
type GenS2 struct { F0 int8; F1 int8; F2 string }
type GenS3 struct { F0 string }
func genFn0(p0 string) bool {
    var v5 int8 = -128
    var v6 int8 = v5
    _ = p0
    _ = v5
    _ = v6
    return true
}
func genFn1(p0 uint32, p1 int8, p2 int16) bool {
    var v7 uint32 = p0
    fmt.Println(v7, v7)
    fmt.Println(p1)
    v8 := []int{128, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return true
}

func main() {
    var v1 uint = 1
    var v2 int8 = 100
    v3 := make(map[int]int)
    var v4 int64 = 9223372036854775807
    v1 += v1
    if (!(v1 < (v1 + (^(v1))))) {
        if true {
            fmt.Println(v4, v4)
            var v9 int = -2147483648
            var v10 float32 = float32(float64((3.14159 - 0.1)))
            _ = v10
            _ = v9
        } else {
            v1 = v1
        }
        switch "hello" {
        case "foo bar":
            if false {
                v11 := GenS1{F0: 2147483647, F1: 0, F2: v4}
                if false {
                    v12 := []int{}
                    fmt.Println(v2, v4, v4)
                    _ = v12
                } else {
                    var v13 uint64 = 18446744073709551615
                    switch 0 {
                    case -2147483648:
                        fmt.Println(v4, v1)
                    case 0:
                        for fi14 := 0; fi14 < 3; fi14++ {
                            if ((924 <= int32(v2)) || true) {
                                v4 += 0
                                var v15 int8 = v2
                                var v16 uint64 = uint64(v1)
                                _ = v15
                                _ = v16
                            }
                            _ = fi14
                        }
                    default:
                        var v17 int = 32767
                        _ = v17
                    }
                    if false {
                        v18 := GenS1{F0: 42, F1: 127, F2: v4}
                        var v19 uint64 = v13
                        var v20 string = "\n\t"
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        var v21 uint8 = 0
                        _ = len(v3)
                        v4 = 9223372036854775807
                        _ = v21
                    }
                    _ = v13
                }
                if (-2.25 > 0.1) {
                    v1 -= (v1 << ((v1) & 63))
                    fmt.Println(v2, v1, v4)
                    fmt.Println(v1, v4)
                } else {
                    var v22 float64 = 1e10
                    _ = v22
                }
                _ = v11
            }
        case "hello":
            v1 = 65535
        case "":
            v23 := GenS2{F0: (v2 << ((255) & 63)), F1: v2, F2: "foo bar"}
            _ = v23
        default:
            v4 = -136
        }
    } else {
        v1 -= 402
        v24 := []int{128, 9223372036854775807, 4294967295}
        _ = v24
    }
    fmt.Println(v4, v2)
    if true {
        v25 := []int{4294967295, 32768, 407, -1}
        switch "hello" {
        case "foo bar":
            var v26 int = -2147483648
            _ = v26
        default:
            if true {
                if (1 > -1) {
                    v2 += (v2 & v2)
                } else {
                    v27 := GenS3{F0: ""}
                    v28 := GenS1{F0: -2147483648, F1: 164, F2: -413}
                    _ = v27
                    _ = v28
                }
            }
        }
        var v29 float32 = 0.1
        _ = v25
        _ = v29
    } else {
        v1 += v1
        v1 = v1
        v30 := GenS3{F0: "hello"}
        _ = v30
    }
    v31 := []int{-185, -2147483648}
    v1 = v1
    fmt.Println(v4, v4)
    if ((v2 < -128) && (!false)) {
        var v32 uint8 = 0
        for fi33 := 0; fi33 < 12; fi33++ {
            fmt.Println(v1)
            v1 = v1
            _ = fi33
        }
        v2 += (^(v2))
        _ = v32
    }
    if len(v31) > 0 {
        v31[0] = int(v4)
    }
    v34 := GenS1{F0: 0, F1: 0, F2: (v4 << ((255) & 63))}
    if false {
        fmt.Println(v1)
    }
    var v35 int8 = (v2 ^ v2)
    var v36 uint = v1
    if (!(true || false)) {
        if true {
            if len(v31) >= 2 {
                v31 = v31[1:len(v31)]
            }
        }
        if false {
            v35 = (v35 | v35)
            switch 4294967295 {
            case 9223372036854775807:
                switch "hello" {
                case "":
                    if (2147483647 >= -956) {
                        if len(v31) > 0 {
                            v31[0] = 38
                        }
                    }
                default:
                    switch "" {
                    case "\n\t":
                        var v37 uint = (^(v1))
                        _ = v37
                    }
                }
            case 128:
                v35 = v2
            case 32768:
                fmt.Println(v4, v4, v35)
            default:
                var v38 float32 = (0.0 / (0.0 - 3.14159))
                _ = v38
            }
        }
    } else {
        var v39 bool = (!(42 == -1))
        switch 706 {
        case 128:
            v36 = 65535
        }
        _ = v39
    }
    _ = len(v31)
    var v40 uint8 = 42
    var v41 int16 = 60
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v41
}
