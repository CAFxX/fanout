package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
func genFn0(p0 string, p1 int32, p2 uint) int64 {
    fmt.Println(p0, p1, p0)
    var v4 float64 = float64((1e10 * 1.5))
    var v5 int32 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 int32 = 1
    var v2 string = "a"
    v3 := make(map[uint32]bool)
    if (-2.25 >= 0.1) {
        fmt.Println(v2, v1)
        if true {
            v6 := GenS1{F0: (v1 * (v1 / ((v1) | 1))), F1: 446, F2: 42}
            if true {
                if (-73 != 32767) {
                    v1 = v1
                }
                switch "a" {
                case "":
                    var v7 float32 = 1.5
                    _ = v7
                case "hello":
                    for fi8 := 0; fi8 < 7; fi8++ {
                        v1 += (v1 & v1)
                        switch v2 {
                        case "hello":
                            var v9 uint = 0
                            _ = v9
                        }
                        if (1.5 > float64((0.1 / (float32(1.5) - 0.1)))) {
                            var v10 int64 = 0
                            v11 := []int{fi8, fi8}
                            if len(v11) > 0 {
                                v11[0] = (fi8 - fi8)
                            }
                            _ = v10
                            _ = v11
                        } else {
                            switch -9223372036854775808 {
                            case -92:
                                switch fi8 {
                                case 9223372036854775807:
                                    fmt.Println(fi8, fi8, v1)
                                case 0:
                                    var v12 uint = 255
                                    _ = v12
                                case 32768:
                                    v2 = v2
                                }
                            case 4294967295:
                                var v13 int64 = -9223372036854775808
                                _ = v13
                            default:
                                switch fi8 {
                                case 852:
                                    v1 -= 635
                                case 32767:
                                    var v14 bool = (true || true)
                                    _ = v14
                                case -75:
                                    v15 := GenS1{F0: v1, F1: fi8, F2: 2147483648}
                                    _ = v15
                                default:
                                    var v16 uint16 = 42
                                    _ = v16
                                }
                            }
                        }
                        _ = fi8
                    }
                default:
                    var v17 int16 = 42
                    _ = v17
                }
            }
            _ = v6
        }
    } else {
        v18 := GenS1{F0: (v1 * v1), F1: 762, F2: 65535}
        _ = v18
    }
    var v19 uint8 = 127
    var v20 int16 = -42
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 string = (v2 + v2)
        v22 = v2
        v1 += (v1 >> ((172) & 63))
        _ = fi21
        _ = v22
    }
    switch -9223372036854775808 {
    case 128:
        v23 := GenS1{F0: v1, F1: 0, F2: 986}
        _ = v23
    default:
        v24 := GenS2{F0: 2, F1: 65535, F2: false}
        _ = v24
    }
    v20 = v20
    v25 := GenS3{F0: 4294967295}
    fmt.Println(v1, v2)
    var v26 uint = 65535
    if (-2.25 >= float64(0.1)) {
        switch 127 {
        case -566:
            v27 := GenS3{F0: 2}
            _ = v27
        case 9223372036854775807:
            if true {
                if false {
                    v2 = v2
                    v20 = (v20 ^ v20)
                    v2 = ("foo bar" + ("hello" + v2))
                } else {
                    for fi28 := 0; fi28 < 3; fi28++ {
                        var v29 int = int(v1)
                        _ = fi28
                        _ = v29
                    }
                    v30 := GenS4{F0: v1}
                    var v31 int16 = (-(v20))
                    _ = v30
                    _ = v31
                }
            } else {
                fmt.Println(v1, v19, v20)
            }
        case 924:
            for fi32 := 0; fi32 < 18; fi32++ {
                v33 := GenS3{F0: 42}
                v26 += 830
                var v34 uint64 = 369
                _ = fi32
                _ = v33
                _ = v34
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
}
