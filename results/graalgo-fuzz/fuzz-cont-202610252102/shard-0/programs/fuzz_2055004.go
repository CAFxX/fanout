package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint }
type GenS3 struct { F0 int32; F1 int64; F2 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8; F1 int8 }
func genFn0() int64 {
    var v7 int = 2147483647
    switch (v7 >> ((1) & 63)) {
    case 32768:
        for fi8 := 0; fi8 < 2; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                v7 += (fi8 + 32768)
                switch fi9 {
                case -1:
                    fmt.Println(v7, fi9, fi8)
                case 656:
                    fmt.Println(fi9, v7)
                case 2147483647:
                    v7 -= (fi9 - v7)
                default:
                    var v10 bool = (42 != 4294967295)
                    _ = v10
                }
                fmt.Println(fi9, fi8, fi9)
                _ = fi9
            }
            v11 := []int{(fi8 | v7)}
            _ = fi8
            _ = v11
        }
    case 32767:
        if (2 <= 867) {
            fmt.Println(v7)
            for fi12 := 0; fi12 < 4; fi12++ {
                if true {
                    for fi13 := 0; fi13 < 8; fi13++ {
                        var v14 uint64 = 18446744073709551615
                        v7 = (fi13 | v7)
                        _ = fi13
                        _ = v14
                    }
                    v15 := GenS1{F0: 255, F1: 181}
                    _ = v15
                } else {
                    v7 -= 4294967295
                    v16 := []int{fi12, (fi12 ^ (fi12 << ((418) & 63))), 0}
                    _ = v16
                }
                _ = fi12
            }
            switch (-(v7)) {
            case -1:
                v17 := []int{(v7 | v7), (v7 ^ 1), 0}
                _ = v17
            case -2147483648:
                for fi18 := 0; fi18 < 5; fi18++ {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        var v20 uint16 = 2
                        if true {
                            var v21 bool = (0 > 42)
                            v22 := GenS1{F0: 1, F1: 109}
                            var v23 uint64 = 1
                            _ = v21
                            _ = v22
                            _ = v23
                        } else {
                            v24 := []int{}
                            _ = v24
                        }
                        _ = fi19
                        _ = v20
                    }
                    for fi25 := 0; fi25 < 20; fi25++ {
                        var v26 int16 = -1
                        fmt.Println(fi25, fi25, v7)
                        _ = fi25
                        _ = v26
                    }
                    _ = fi18
                }
            default:
                v27 := []int{0}
                _ = v27
            }
        }
    default:
        var v28 uint8 = 42
        _ = v28
    }
    v7 = v7
    v29 := GenS2{F0: 0, F1: 2, F2: 65535}
    v7 += v7
    _ = v29
    _ = v7
    return 42
}
func genFn1(p0 uint16, p1 int8) uint32 {
    v30 := GenS2{F0: -9223372036854775808, F1: p0, F2: 15}
    v31 := GenS3{F0: -591, F1: -9223372036854775808, F2: 127}
    _ = p0
    _ = p1
    _ = v30
    _ = v31
    return 2147483648
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 1
    var v3 bool = false
    v4 := make(map[uint32]bool)
    v5 := make(map[int]int)
    var v6 int64 = -1
    var v32 float64 = 0.0
    v33 := []int{204, 4294967295, -881}
    if len(v33) > 0 {
        v33[0] = -1
    }
    if v3 {
        for fi34 := 0; fi34 < 12; fi34++ {
            v32 += ((float64((float32(float64(0.0)) / 1.5)) * float64(3.14159)) - float64(0.1))
            _ = fi34
        }
    } else {
        v35 := GenS3{F0: v2, F1: v6, F2: 32767}
        _ = v35
    }
    v36 := GenS4{F0: 0}
    v37 := GenS5{F0: 0, F1: 2, F2: 1}
    var v38 float64 = (v32 / v32)
    var v39 float64 = ((v38 / (-2.25 * v32)) * v32)
    v40 := GenS6{F0: 644}
    v41 := GenS7{F0: -128, F1: 127}
    v3 = (914 >= 4294967295)
    var v42 int = -440
    v3 = v3
    var v43 string = ("a" + "foo bar")
    switch 122 {
    case 32767:
        var v44 int32 = (v2 | v2)
        _ = v44
    case 9223372036854775807:
        var v45 int64 = genFn0()
        _ = v45
    default:
        var v46 uint8 = 128
        _ = v46
    }
    v6 = 9223372036854775807
    var v47 bool = (v3 && v3)
    var v48 int64 = (v6 + -9223372036854775808)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
