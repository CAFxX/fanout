package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint; F1 uint8; F2 int }
type GenS5 struct { F0 string }
func genFn0() bool {
    var v5 uint64 = 0
    v6 := GenS1{F0: 127, F1: 0, F2: (673 <= 18446744073709551615)}
    _ = v5
    _ = v6
    return (1 < 9223372036854775807)
}
func genFn1(p0 int, p1 uint16, p2 uint8) int16 {
    v7 := GenS2{F0: 1}
    if (0 == 1) {
        var v8 int32 = -353
        if false {
            var v9 float64 = 0.1
            fmt.Println(p2, p0)
            if true {
                v10 := GenS1{F0: 1, F1: 1, F2: (1 > 42)}
                _ = v10
            } else {
                var v11 float64 = (v9 + v9)
                var v12 bool = false
                v13 := []int{(^(p0))}
                _ = v11
                _ = v12
                _ = v13
            }
            _ = v9
        }
        _ = v8
    } else {
        v14 := GenS3{F0: p0, F1: 1000}
        var v15 uint8 = p2
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return int16(p2)
}

func main() {
    var v1 int64 = 47
    v2 := make(map[uint32]int)
    v3 := []uint32{315, 129, 839}
    var v4 int64 = (v1 % (((v1 / ((42) | 1))) | 1))
    var v16 bool = (true && true)
    v4 -= (v4 >> ((772) & 63))
    fmt.Println(v1, v4)
    _ = len(v3)
    for fi17 := 0; fi17 < 14; fi17++ {
        if v16 {
            v18 := GenS4{F0: 65535, F1: 116, F2: -2147483648}
            _ = v18
        }
        v4 = int64(fi17)
        _ = fi17
    }
    if false {
        switch 42 {
        case 4294967295:
            v19 := GenS5{F0: "\n\t"}
            _ = v19
        case -1:
            var v20 uint16 = 979
            _ = v20
        default:
            fmt.Println(v16, v16, v16)
        }
        var v21 uint32 = 0
        _ = v21
    }
    v1 -= (v1 << ((0) & 63))
    fmt.Println(v4, v4)
    for fi22 := 0; fi22 < 5; fi22++ {
        v4 += (-(v1))
        for fi23 := 0; fi23 < 4; fi23++ {
            fmt.Println(v1, v4, fi22)
            var v24 int32 = 42
            _ = fi23
            _ = v24
        }
        v16 = (42 != 1000)
        _ = fi22
    }
    switch "foo bar" {
    case "foo bar":
        var v25 uint8 = 201
        _ = v25
    case "\n\t":
        if (v16 || (!v16)) {
            for fi26 := 0; fi26 < 12; fi26++ {
                _ = len(v3)
                _ = fi26
            }
            switch 872 {
            case 128:
                var v27 float32 = float32((-2.25 + 0.1))
                _ = v27
            case -195:
                if false {
                    var v28 int32 = 42
                    if v16 {
                        v29 := v2[2147483648]
                        fmt.Println(v29)
                        var v30 int = v29
                        _ = v29
                        _ = v30
                    }
                    var v31 int32 = (-(v28))
                    _ = v28
                    _ = v31
                }
            default:
                v4 -= -9223372036854775808
            }
            if (4294967295 <= 2147483647) {
                if false {
                    if true {
                        for fi32 := 0; fi32 < 16; fi32++ {
                            v33 := GenS3{F0: fi32, F1: -32768}
                            delete(v2, 65535)
                            _ = fi32
                            _ = v33
                        }
                        v2[2147483648] = -1
                    } else {
                        var v34 uint8 = 128
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        fmt.Println(v4, v16)
                        _ = v34
                    }
                    _ = len(v3)
                }
                var v35 int = 127
                v1 += v1
                _ = v35
            }
        }
    }
    var v36 float64 = float64(3.14159)
    fmt.Println(v1, v1, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
