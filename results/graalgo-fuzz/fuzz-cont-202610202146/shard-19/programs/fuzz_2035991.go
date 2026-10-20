package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 bool }
func genFn0(p0 string, p1 uint16) uint8 {
    v7 := []int{}
    v8 := GenS1{F0: 348}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 128
}

func main() {
    v1 := []uint32{2147483647, 2147483647, 2147483648, 2}
    var v2 string = (("hello" + "\n\t") + "\n\t")
    v3 := make(map[uint32]string)
    var v4 int16 = -1
    v5 := []uint32{4294967295, 2, 2147483648}
    var v6 float64 = (((float64((1.5 - 1.5)) * -2.25) - 0.1) - (0.0 - -2.25))
    v2 = v2
    switch ("a" + v2) {
    case "hello":
        var v9 float32 = 0.0
        _ = v9
    case "a":
        delete(v3, 795)
    case "foo bar":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        v4 += v4
    }
    var v10 float64 = (3.14159 * 1e10)
    v11 := GenS2{F0: 1}
    var v12 int = 0
    if len(v5) > 0 {
        v5[0] = 887
    }
    switch (v12 | (v12 * -1)) {
    case -1:
        switch "foo bar" {
        case "a":
            var v13 float64 = -2.25
            _ = v13
        case "hello":
            for fi14 := 0; fi14 < 10; fi14++ {
                for fi15 := 0; fi15 < 19; fi15++ {
                    var v16 uint32 = 4294967295
                    v6 += (v10 + v10)
                    var v17 int16 = (v4 << ((18446744073709551615) & 63))
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = fi14
            }
        default:
            v18 := GenS3{F0: 893, F1: v4}
            _ = v18
        }
    case 42:
        if ((!(-365 < int64(v4))) || (-128 == 17)) {
            var v19 uint16 = 9
            _ = v19
        } else {
            for fi20 := 0; fi20 < 12; fi20++ {
                _ = len(v3)
                v4 += (v4 % ((v4) | 1))
                for fi21 := 0; fi21 < 3; fi21++ {
                    switch fi20 {
                    case 856:
                        if (1 > 799) {
                            v3[757] = ("a" + "foo bar")
                        }
                    case 42:
                        v5 = append(v5, 158)
                    default:
                        v6 -= (v10 - (0.0 - (1.5 / (float64(-2.25) / 1e10))))
                    }
                    _ = fi21
                }
                _ = fi20
            }
            v22 := v3[4294967295]
            _ = v22
        }
    default:
        v23 := GenS1{F0: 0}
        _ = v23
    }
    var v24 uint8 = 69
    fmt.Println(v4, v24, v24)
    v12 += 42
    if false {
        var v25 int64 = -739
        var v26 int8 = 127
        _ = v25
        _ = v26
    } else {
        for fi27 := 0; fi27 < 18; fi27++ {
            switch (fi27 ^ (^(fi27))) {
            case 127:
                var v28 uint64 = 683
                _ = v28
            case 42:
                for fi29 := 0; fi29 < 3; fi29++ {
                    var v30 uint = 0
                    if (478 < 871) {
                        switch v2 {
                        case "hello":
                            fmt.Println(v4)
                        default:
                            var v31 int64 = -1
                            _ = v31
                        }
                        var v32 uint32 = 884
                        if (1 > int64(v32)) {
                            var v33 uint16 = uint16(v32)
                            v34 := GenS4{F0: 2, F1: true}
                            _ = v33
                            _ = v34
                        }
                        _ = v32
                    } else {
                        v35 := GenS5{F0: true}
                        fmt.Println(v10, v12, v12)
                        _ = len(v1)
                        _ = v35
                    }
                    _ = fi29
                    _ = v30
                }
            case -1:
                _ = len(v3)
            }
            var v36 uint = 42
            fmt.Println(v24, fi27)
            _ = fi27
            _ = v36
        }
    }
    _ = len(v5)
    if len(v5) > 0 {
        v5[0] = 51
    }
    v12 = (^(v12))
    fmt.Println(v12, v12, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
