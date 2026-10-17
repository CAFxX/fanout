package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint8; F1 bool; F2 int }
type GenS3 struct { F0 bool; F1 int64; F2 uint8 }
func genFn0(p0 int16, p1 int64, p2 uint32) int32 {
    var v5 int32 = -2147483648
    for fi6 := 0; fi6 < 16; fi6++ {
        v5 = 0
        if (!true) {
            v7 := GenS1{F0: v5, F1: 42, F2: v5}
            fmt.Println(v5, p1)
            _ = v7
        } else {
            switch fi6 {
            case 631:
                v8 := []int{}
                _ = v8
            case -9223372036854775808:
                if false {
                    v5 += (^(v5))
                    v9 := GenS2{F0: 26, F1: (fi6 >= -442), F2: fi6}
                    _ = v9
                }
            default:
                var v10 int32 = (v5 & v5)
                _ = v10
            }
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -2147483648
}

func main() {
    v1 := []uint32{4294967295, 2147483648, 291}
    v2 := make(map[string]bool)
    v3 := make(map[string]uint32)
    var v4 uint8 = 188
    var v11 int8 = 73
    v11 = v11
    if len(v1) > 0 {
        v1[0] = 2
    }
    var v12 string = ("hello" + "a")
    if true {
        var v13 uint16 = 984
        switch -9223372036854775808 {
        case 4294967295:
            if (((v12 + v12) + ((v12 + v12) + v12)) == v12) {
                v14 := GenS2{F0: (^(v4)), F1: true, F2: 127}
                switch 922 {
                case -569:
                    switch "" {
                    case "":
                        v12 = v12
                    case "\n\t":
                        switch "a" {
                        case "foo bar":
                            var v15 int8 = v11
                            _ = v15
                        default:
                            v16 := GenS1{F0: -944, F1: (v4 % (((v4 % ((0) | 1))) | 1)), F2: -114}
                            _ = v16
                        }
                    default:
                        var v17 int8 = -1
                        _ = v17
                    }
                case -9223372036854775808:
                    var v18 string = "\n\t"
                    _ = v18
                case 42:
                    v11 += v11
                }
                if (true && false) {
                    _ = len(v3)
                    delete(v2, v12)
                } else {
                    fmt.Println(v12)
                }
                _ = v14
            } else {
                var v19 int8 = (v11 + v11)
                if false {
                    v12 = v12
                    v4 = 2
                } else {
                    var v20 bool = (true || (v13 == 255))
                    v21 := GenS3{F0: v20, F1: -1, F2: 1}
                    for fi22 := 0; fi22 < 4; fi22++ {
                        v11 = (v11 - (v19 & v19))
                        v2[v12] = (0.0 < -2.25)
                        fmt.Println(v4)
                        _ = fi22
                    }
                    _ = v20
                    _ = v21
                }
                _ = v19
            }
        case 42:
            var v23 int32 = 95
            _ = v23
        case -1:
            v11 += v11
        }
        _ = v13
    } else {
        var v24 uint = 255
        _ = v24
    }
    var v25 bool = true
    if (0 < -450) {
        var v26 uint16 = 375
        _ = v26
    }
    v27 := GenS3{F0: (142 == 42), F1: -1, F2: v4}
    var v28 int64 = 1
    var v29 uint8 = (^(v4))
    if ((2 <= 4294967295) || (v29 < v4)) {
        for fi30 := 0; fi30 < 4; fi30++ {
            if (genFn0(0, v28, 731) < -181) {
                var v31 int16 = 32767
                _ = v31
            } else {
                var v32 int8 = 127
                var v33 float64 = -2.25
                _ = v32
                _ = v33
            }
            _ = fi30
        }
    } else {
        var v34 int64 = (v28 >> ((255) & 63))
        var v35 int = 9223372036854775807
        if false {
            var v36 int = (v35 + 32767)
            _ = v36
        } else {
            var v37 float32 = ((1e10 * 0.0) + 0.1)
            var v38 int32 = 1
            for fi39 := 0; fi39 < 19; fi39++ {
                fmt.Println(v28, v4)
                var v40 uint = 255
                v41 := GenS3{F0: (false && v25), F1: (v34 - -783), F2: v29}
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = v37
            _ = v38
        }
        _ = v34
        _ = v35
    }
    v42 := GenS1{F0: -2147483648, F1: (v29 * 255), F2: 2147483647}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v42
}
