package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int32; F2 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS5 struct { F0 int8; F1 int8; F2 int }
func genFn0(p0 int8) uint8 {
    var v5 int16 = 1000
    fmt.Println(v5, v5)
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 int64 = -1
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: -290}
    _ = p0
    _ = v5
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 0
    v3 := make(map[string]int)
    var v4 uint16 = 65535
    if (v2 == (v2 % ((v2) | 1))) {
        v9 := GenS2{F0: 128, F1: -2147483648, F2: 560}
        v4 += v4
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := GenS2{F0: 243, F1: int32(fi10), F2: 2147483647}
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        var v12 uint32 = 226
        v4 = (^(v4))
        _ = v12
    }
    v2 = (v2 + v2)
    v13 := GenS3{F0: -207}
    var v14 float32 = 1e10
    v2 += (-(v2))
    fmt.Println(v4)
    if (2147483648 != 0) {
        v15 := []int{}
        _ = v15
    } else {
        fmt.Println(v2, v14, v4)
    }
    var v16 int64 = 42
    v17 := []int{4294967295, 32768, 9223372036854775807, 42}
    var v18 uint8 = 11
    var v19 uint64 = 18446744073709551615
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 bool = true
        var v22 bool = (!(-1 > 2147483647))
        _ = len(v17)
        _ = fi20
        _ = v21
        _ = v22
    }
    if true {
        v17 = append(v17, 128)
        fmt.Println(v14, v19)
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
    } else {
        var v23 uint64 = (v19 * 0)
        var v24 uint = 255
        switch ("\n\t" + "a") {
        case "foo bar":
            for fi25 := 0; fi25 < 17; fi25++ {
                var v26 uint = uint(v19)
                if true {
                    v19 = uint64(v4)
                    for fi27 := 0; fi27 < 1; fi27++ {
                        v26 -= 255
                        _ = fi27
                    }
                }
                var v28 uint = 190
                _ = fi25
                _ = v26
                _ = v28
            }
        case "a":
            switch 0 {
            case -548:
                v18 += (v18 | v18)
            case -2147483648:
                fmt.Println(v16, v4, v24)
            case 127:
                v18 -= (v18 / (((^(v18))) | 1))
            }
        case "":
            for fi29 := 0; fi29 < 7; fi29++ {
                if false {
                    if (v2 > int8(v4)) {
                        if (v4 >= (v4 / ((v4) | 1))) {
                            v30 := GenS4{F0: v23, F1: (v18 ^ genFn0(v2)), F2: 100}
                            v17 = append(v17, fi29)
                            v3[""] = (^(fi29))
                            _ = v30
                        }
                        v31 := GenS1{F0: (v16 + (v16 << ((uint(v18)) & 63)))}
                        var v32 uint64 = (v23 % ((0) | 1))
                        _ = v31
                        _ = v32
                    }
                } else {
                    v24 = (v24 / ((v24) | 1))
                    switch (fi29 - -485) {
                    case 1:
                        v17 = append(v17, (fi29 - (fi29 / ((0) | 1))))
                    case -233:
                        v1[(fi29 + fi29)] = 962
                    }
                }
                _ = fi29
            }
        default:
            v33 := GenS5{F0: v2, F1: v2, F2: 127}
            _ = v33
        }
        _ = v23
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
