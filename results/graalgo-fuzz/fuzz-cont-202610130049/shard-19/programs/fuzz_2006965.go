package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint16 }
func genFn0() int {
    v4 := GenS1{F0: ((0.0 - 1.5) >= (0.0 + 0.0)), F1: 255, F2: -1}
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 float32 = 1.5
            fmt.Println(v7, fi5, fi6)
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    v8 := []int{4294967295, 728, 42}
    v9 := GenS1{F0: false, F1: 128, F2: 1}
    _ = v4
    _ = v8
    _ = v9
    return -9223372036854775808
}
func genFn1(p0 int32, p1 uint64, p2 uint64) int32 {
    switch genFn0() {
    case -647:
        v10 := []int{}
        _ = v10
    case -2147483648:
        v11 := []int{-1, -9223372036854775808}
        _ = v11
    case -500:
        v12 := GenS1{F0: ((!true) && (!false)), F1: 127, F2: 42}
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        p1 += (p1 / ((0) | 1))
        _ = fi13
    }
    var v14 uint = 973
    v14 = (v14 << ((uint(p1)) & 63))
    v15 := GenS1{F0: false, F1: 127, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 57
}

func main() {
    var v1 uint32 = 0
    var v2 float64 = (0.1 * ((-2.25 - 0.1) / -2.25))
    var v3 int16 = -632
    fmt.Println(v2, v2)
    v3 += 356
    var v16 float64 = v2
    if false {
        var v17 uint32 = (^(v1))
        _ = v17
    } else {
        v18 := GenS2{F0: ((!(-49 < -1)) || false)}
        if false {
            for fi19 := 0; fi19 < 19; fi19++ {
                v16 -= 1.5
                if (2 <= 196) {
                    var v20 int32 = 42
                    if false {
                        fmt.Println(v2, v16)
                        v1 += (^(v1))
                    }
                    _ = v20
                }
                v21 := []int{(fi19 / (((fi19 + genFn0())) | 1)), (fi19 >> ((18446744073709551615) & 63)), (fi19 | 32768)}
                _ = fi19
                _ = v21
            }
            for fi22 := 0; fi22 < 12; fi22++ {
                v2 -= v2
                _ = fi22
            }
            v23 := GenS1{F0: true, F1: 74, F2: 42}
            _ = v23
        } else {
            switch genFn0() {
            case 127:
                v24 := GenS2{F0: (!(!(!(65535 > 4294967295))))}
                _ = v24
            case -438:
                v3 = (v3 + v3)
            default:
                fmt.Println(v3, v3)
            }
            v16 -= 3.14159
        }
        var v25 uint64 = 0
        _ = v18
        _ = v25
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        v27 := GenS2{F0: (!((255 < 0) && true))}
        for fi28 := 0; fi28 < 13; fi28++ {
            if true {
                v29 := GenS1{F0: (146 == 32767), F1: 7, F2: 9223372036854775807}
                var v30 int32 = 2147483647
                var v31 int16 = 1
                _ = v29
                _ = v30
                _ = v31
            }
            v32 := GenS2{F0: (4294967295 > 2)}
            _ = fi28
            _ = v32
        }
        v33 := []int{genFn0(), 823, fi26}
        _ = fi26
        _ = v27
        _ = v33
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        if true {
            var v35 uint16 = 776
            v36 := GenS2{F0: (811 >= -1)}
            fmt.Println(v3)
            _ = v35
            _ = v36
        }
        var v37 uint8 = 94
        _ = fi34
        _ = v37
    }
    switch -1 {
    case 128:
        v38 := []int{}
        _ = v38
    case 9223372036854775807:
        fmt.Println(v16, v2, v1)
    }
    fmt.Println(v2, v3)
    v39 := []int{941, 32768, 9223372036854775807, 0}
    for fi40 := 0; fi40 < 6; fi40++ {
        v41 := GenS2{F0: true}
        switch ("hello" + ("foo bar" + "foo bar")) {
        case "hello":
            switch "foo bar" {
            case "a":
                var v42 int16 = (v3 ^ 32767)
                _ = v42
            default:
                _ = len(v39)
            }
        default:
            var v43 uint32 = v1
            _ = v43
        }
        _ = fi40
        _ = v41
    }
    var v44 uint32 = 42
    v45 := GenS2{F0: true}
    v16 += v16
    v46 := GenS3{F0: -2147483648, F1: 2}
    var v47 uint = 255
    v1 = v1
    switch 919 {
    case 1:
        v39 = append(v39, 943)
    case 380:
        fmt.Println(v44, v47)
    default:
        for fi48 := 0; fi48 < 3; fi48++ {
            v3 -= -221
            var v49 int64 = 0
            _ = fi48
            _ = v49
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v39
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
