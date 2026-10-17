package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 bool; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8; F1 uint16 }

func main() {
    var v1 uint32 = 0
    var v2 float64 = 0.0
    var v3 int64 = 674
    v4 := make(map[int]uint32)
    v5 := make(map[int]string)
    v6 := GenS1{F0: 1}
    fmt.Println(v3, v3)
    for fi7 := 0; fi7 < 20; fi7++ {
        v3 = (v3 & 9223372036854775807)
        _ = fi7
    }
    v2 = 1e10
    v2 -= float64(0.0)
    switch -572 {
    case -785:
        var v8 uint8 = 15
        _ = v8
    case 42:
        fmt.Println(v1, v3)
    case -2147483648:
        if true {
            for fi9 := 0; fi9 < 5; fi9++ {
                v10 := GenS2{F0: uint64(v1), F1: 42}
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: 0}
            var v12 uint32 = (v1 << ((255) & 63))
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 uint = 65535
                var v15 bool = (-60 != 237)
                _ = fi13
                _ = v14
                _ = v15
            }
            switch 1 {
            case 9223372036854775807:
                var v16 string = "hello"
                _ = v16
            case 200:
                v17 := GenS1{F0: -2147483648}
                _ = v17
            }
            v18 := GenS2{F0: 779, F1: 1}
            _ = v18
        }
    default:
        for fi19 := 0; fi19 < 17; fi19++ {
            if false {
                fmt.Println(fi19, fi19, v3)
                var v20 bool = (0 != 0)
                _ = v20
            }
            _ = fi19
        }
    }
    if ((v3 + (-(v3))) >= -1) {
        var v21 bool = (1000 >= -1)
        _ = v21
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        v23 := GenS3{F0: true, F1: 42, F2: (v1 % (((v1 ^ (^(v1)))) | 1))}
        for fi24 := 0; fi24 < 1; fi24++ {
            if false {
                var v25 int64 = 0
                if true {
                    v26 := GenS3{F0: (float64(0.1) > v2), F1: 0, F2: 1}
                    v25 = (v3 % ((-15) | 1))
                    v2 -= (((1e10 - 0.0) + v2) - v2)
                    _ = v26
                }
                if true {
                    v27 := GenS4{F0: 0, F1: 1, F2: 392}
                    v25 = v25
                    _ = v27
                } else {
                    if false {
                        v28 := []int{}
                        var v29 uint64 = 2
                        v30 := GenS2{F0: v29, F1: 32767}
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    var v31 uint8 = 121
                    v32 := GenS1{F0: 2147483647}
                    _ = v31
                    _ = v32
                }
                var v33 string = "foo bar"
                _ = v25
                _ = v33
            } else {
                if (!true) {
                    switch 32767 {
                    case 0:
                        var v34 int64 = v3
                        _ = v34
                    case 1:
                        v2 += v2
                    default:
                        if (65535 <= 732) {
                            delete(v4, (fi22 ^ (^(fi24))))
                        } else {
                            var v35 int8 = 21
                            switch "hello" {
                            case "hello":
                                v36 := []int{fi24, 127, 127, fi24}
                                _ = v36
                            case "a":
                                if ((v35 | 42) <= v35) {
                                    v1 = v1
                                }
                            case "":
                                v35 = 42
                            default:
                                v35 += v35
                            }
                            var v37 float32 = ((3.14159 + 1e10) + ((float32(float64((1.5 / 1e10))) + 1.5) / (-2.25 * (1e10 + (0.0 * (-2.25 * (0.1 * 0.1)))))))
                            _ = v35
                            _ = v37
                        }
                    }
                    if (fi22 >= (fi24 ^ 32767)) {
                        v38 := GenS3{F0: (1e10 <= (float32((1e10 * v2)) / 0.1)), F1: 42, F2: v1}
                        _ = v38
                    } else {
                        var v39 uint64 = 18446744073709551615
                        _ = v39
                    }
                } else {
                    v40 := GenS5{F0: 724}
                    v2 -= float64(((3.14159 * 3.14159) - (0.0 / 1.5)))
                    v2 -= ((v2 / v2) + 1e10)
                    v41 := GenS6{F0: 1, F1: 255}
                    _ = v40
                    _ = v41
                }
            }
            _ = fi24
        }
        _ = fi22
        _ = v23
    }
    var v42 uint8 = 42
    fmt.Println(v42)
    v43 := GenS5{F0: 506}
    v3 = v3
    v42 += (v42 | (v42 ^ v42))
    v2 = v2
    switch "\n\t" {
    case "\n\t":
        v44 := []int{0, 32767, 32767, 0}
        _ = v44
    case "hello":
        var v45 int64 = v3
        _ = v45
    default:
        v42 -= v42
    }
    var v46 int8 = 0
    v5[32768] = "hello"
    var v47 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
