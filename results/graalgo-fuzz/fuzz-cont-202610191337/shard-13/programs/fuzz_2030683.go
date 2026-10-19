package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int8; F1 uint; F2 bool }
type GenS4 struct { F0 int32; F1 bool; F2 int64 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint16) int16 {
    v6 := GenS1{F0: 2147483647, F1: p0, F2: 4294967295}
    var v7 uint16 = (p0 | p0)
    _ = p0
    _ = v6
    _ = v7
    return -32768
}

func main() {
    var v1 uint64 = 1
    var v2 int64 = 9223372036854775807
    var v3 int32 = -2147483648
    var v4 uint8 = 28
    var v5 float64 = ((0.0 - -2.25) + 1.5)
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := []int{fi8, 568, fi8}
        fmt.Println(v4)
        if len(v9) > 0 {
            v9[0] = fi8
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(v1, v4, v3)
    switch -284 {
    case 213:
        fmt.Println(v1, v3)
    case 32768:
        v10 := []int{-1, 510, -1, 0}
        _ = v10
    }
    v11 := GenS2{F0: 0, F1: 405}
    v2 -= v2
    v3 += (v3 + v3)
    for fi12 := 0; fi12 < 12; fi12++ {
        switch fi12 {
        case 1:
            if (!(!true)) {
                var v13 int32 = (^(v3))
                if true {
                    fmt.Println(v1, v4, fi12)
                } else {
                    v14 := GenS2{F0: 2, F1: 42}
                    _ = v14
                }
                if (!(65535 > 32767)) {
                    for fi15 := 0; fi15 < 12; fi15++ {
                        fmt.Println(v4)
                        v16 := GenS3{F0: 0, F1: 103, F2: (1 >= 42)}
                        fmt.Println(v13, fi12)
                        _ = fi15
                        _ = v16
                    }
                    if (true || false) {
                        for fi17 := 0; fi17 < 19; fi17++ {
                            v18 := GenS4{F0: (v13 ^ -2147483648), F1: true, F2: (v2 / ((-701) | 1))}
                            var v19 string = "foo bar"
                            switch fi12 {
                            case 4294967295:
                                var v20 uint = 195
                                _ = v20
                            case 32767:
                                var v21 uint = 24
                                _ = v21
                            }
                            var v22 int16 = 0
                            _ = fi17
                            _ = v18
                            _ = v19
                            _ = v22
                        }
                    } else {
                        var v23 uint8 = 127
                        var v24 float64 = v5
                        _ = v23
                        _ = v24
                    }
                }
                _ = v13
            } else {
                fmt.Println(v5, v1)
                for fi25 := 0; fi25 < 4; fi25++ {
                    v26 := GenS4{F0: v3, F1: true, F2: 742}
                    var v27 string = ""
                    var v28 uint64 = 1
                    if true {
                        var v29 uint = 5
                        _ = v29
                    } else {
                        v27 = ""
                        v30 := GenS5{F0: -61}
                        _ = v30
                    }
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                }
            }
        }
        v1 += (^(v1))
        var v31 float32 = (-2.25 - 1e10)
        _ = fi12
        _ = v31
    }
    v3 = v3
    v32 := []int{9223372036854775807, 32767, 2147483647, 42}
    v33 := GenS4{F0: 0, F1: true, F2: 0}
    var v34 int16 = 42
    v5 += v5
    var v35 int64 = v2
    fmt.Println(v4, v5)
    if (!(!(137 == int(v3)))) {
        v32 = append(v32, 505)
    } else {
        v3 = (v3 << ((uint(v4)) & 63))
    }
    v1 -= (^(v1))
    v3 -= (-(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
