package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int8; F2 uint }
type GenS4 struct { F0 uint16; F1 int16 }
func genFn0(p0 int64, p1 int, p2 uint32) string {
    v4 := []int{}
    v5 := GenS1{F0: p0, F1: 255}
    p2 = (p2 * p2)
    for fi6 := 0; fi6 < 12; fi6++ {
        fmt.Println(p1, p0, p1)
        switch 2147483647 {
        case 4294967295:
            p1 = (p1 << ((477) & 63))
        }
        var v7 uint8 = 127
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return "foo bar"
}
func genFn1(p0 uint16, p1 int16) uint {
    fmt.Println(p0, p0)
    if (!(("hello" + "hello") != "a")) {
        v8 := []int{2147483647, 128, -893}
        var v9 int16 = -1
        switch 42 {
        case -9223372036854775808:
            if false {
                switch 0 {
                case 0:
                    var v10 uint32 = 407
                    _ = v10
                case -9223372036854775808:
                    if true {
                        _ = len(v8)
                        var v11 int32 = -2147483648
                        _ = v11
                    } else {
                        switch "" {
                        case "a":
                            v12 := GenS2{F0: 42}
                            _ = v12
                        case "foo bar":
                            var v13 int64 = 0
                            _ = v13
                        }
                        var v14 uint8 = uint8(p0)
                        _ = v14
                    }
                }
                var v15 uint32 = 902
                _ = v15
            }
        default:
            for fi16 := 0; fi16 < 20; fi16++ {
                if len(v8) > 0 {
                    v8[0] = (fi16 * fi16)
                }
                _ = fi16
            }
        }
        _ = v8
        _ = v9
    }
    v17 := []int{-418, 714, 32768}
    _ = p0
    _ = p1
    _ = v17
    return 0
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]string)
    var v3 int32 = -2147483648
    for fi18 := 0; fi18 < 10; fi18++ {
        if false {
            v1 += 0
        }
        if true {
            v19 := []int{}
            v1 = v1
            _ = v19
        }
        var v20 int16 = -951
        _ = fi18
        _ = v20
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v2[-1] = (("\n\t" + "foo bar") + "\n\t")
        v22 := []int{(fi21 ^ 0)}
        v1 -= v1
        _ = fi21
        _ = v22
    }
    v1 = (v1 | (^(v1)))
    if true {
        v1 = v1
        v23 := GenS3{F0: (!((881 >= 0) && false)), F1: -24, F2: 42}
        _ = v23
    } else {
        v1 -= (v1 ^ 1)
        switch "a" {
        case "foo bar":
            if true {
                var v24 uint64 = 18446744073709551615
                v3 -= 42
                _ = v24
            } else {
                for fi25 := 0; fi25 < 14; fi25++ {
                    if (42 <= 714) {
                        for fi26 := 0; fi26 < 2; fi26++ {
                            v1 -= 0
                            _ = fi26
                        }
                        v27 := GenS1{F0: -1, F1: 87}
                        _ = v27
                    }
                    _ = fi25
                }
                fmt.Println(v3)
                v28 := GenS4{F0: 1, F1: 0}
                _ = v28
            }
        case "":
            v29 := []int{}
            _ = v29
        }
        v30 := GenS4{F0: 42, F1: 42}
        _ = v30
    }
    fmt.Println(v3, v1)
    if true {
        var v31 uint8 = (^(v1))
        for fi32 := 0; fi32 < 20; fi32++ {
            var v33 float32 = 1e10
            v34 := []int{}
            _ = fi32
            _ = v33
            _ = v34
        }
        v1 = 175
        _ = v31
    }
    var v35 bool = false
    v36 := []int{318, 411, 620, 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
}
