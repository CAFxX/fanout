package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 int64; F1 int; F2 uint16 }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 int32; F1 int64; F2 uint32 }
type GenS8 struct { F0 bool }
func genFn0(p0 uint16, p1 uint64, p2 uint32) int16 {
    if true {
        var v4 uint = 4294967295
        switch 9223372036854775807 {
        case 2147483647:
            if false {
                v4 -= (v4 - uint(p1))
                v5 := GenS1{F0: 897, F1: -52}
                _ = v5
            } else {
                v6 := GenS2{F0: (v4 - 0), F1: 42}
                var v7 int64 = -1
                _ = v6
                _ = v7
            }
        default:
            v8 := GenS2{F0: (v4 ^ (^(v4))), F1: int(p1)}
            _ = v8
        }
        _ = v4
    } else {
        v9 := []int{42, -636, 32767}
        _ = v9
    }
    var v10 int64 = -9223372036854775808
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return -1
}
func genFn1(p0 int, p1 int16, p2 bool) uint32 {
    v12 := GenS3{F0: -1, F1: p0}
    fmt.Println(p1, p1)
    var v13 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 2
}

func main() {
    v1 := []uint32{2, 178, 1, 1000}
    var v2 int16 = -374
    var v3 uint32 = 880
    switch "hello" {
    case "\n\t":
        var v14 int16 = v2
        _ = v14
    case "hello":
        if len(v1) > 0 {
            v1[0] = genFn1(32768, v2, true)
        }
    case "a":
        v3 += 4294967295
    }
    var v15 float32 = (((0.0 * -2.25) * -2.25) * ((3.14159 - ((0.1 * (0.0 / 1.0)) / 0.1)) - 1.5))
    v15 = v15
    v16 := GenS4{F0: 734, F1: true}
    v1 = append(v1, (v3 / ((2147483648) | 1)))
    var v17 uint = 0
    fmt.Println(v3, v15)
    v17 = (v17 ^ (v17 % (((v17 & (v17 / ((v17) | 1)))) | 1)))
    v15 -= 0.0
    for fi18 := 0; fi18 < 15; fi18++ {
        for fi19 := 0; fi19 < 16; fi19++ {
            fmt.Println(v2)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v20 := GenS1{F0: (^(v17)), F1: 100}
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 1; fi21++ {
            if (uint8(v3) <= 2) {
                if (!(0 >= 1)) {
                    v22 := GenS3{F0: 444, F1: (fi18 - fi18)}
                    if (!(false || false)) {
                        var v23 float32 = float32(1e10)
                        v1 = append(v1, (v3 % ((v3) | 1)))
                        if true {
                            v24 := GenS5{F0: 0, F1: fi18, F2: 42}
                            _ = len(v1)
                            v23 = (0.1 * v15)
                            _ = v24
                        }
                        _ = v23
                    }
                    v17 += v17
                    _ = v22
                }
                v25 := GenS6{F0: -593, F1: 32767}
                _ = v25
            } else {
                switch "hello" {
                case "foo bar":
                    if false {
                        if false {
                            if len(v1) > 0 {
                                v1[0] = v3
                            }
                        }
                        var v26 uint = (v17 ^ v17)
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            fmt.Println(fi18, v3, fi18)
                        }
                        _ = v26
                    } else {
                        var v27 float32 = v15
                        _ = v27
                    }
                case "hello":
                    switch (-(fi21)) {
                    case -2147483648:
                        v28 := GenS3{F0: 466, F1: fi18}
                        _ = v28
                    case 2147483647:
                        v17 += (v17 << ((0) & 63))
                    }
                }
                var v29 uint8 = 42
                var v30 int64 = -9223372036854775808
                _ = v29
                _ = v30
            }
            v15 += float32((1.5 * 0.1))
            switch "a" {
            case "hello":
                v31 := GenS7{F0: int32(v2), F1: 9223372036854775807, F2: v3}
                _ = v31
            case "\n\t":
                v15 = (v15 - 0.1)
            case "a":
                v3 += (v3 * v3)
            default:
                v17 = v17
            }
            _ = fi21
        }
        _ = fi18
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v32 := GenS8{F0: (((100 != 0) || true) && (127 <= 0))}
    var v33 int64 = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
}
