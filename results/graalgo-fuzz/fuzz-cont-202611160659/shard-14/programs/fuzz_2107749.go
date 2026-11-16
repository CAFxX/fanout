package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int16) int {
    if true {
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := []int{(fi6 - -452), (^(fi6)), fi6, (fi6 ^ 32768)}
            fmt.Println(p0, fi6)
            for fi8 := 0; fi8 < 9; fi8++ {
                if (!((8 == 284) && true)) {
                    p0 = (-(p0))
                    if len(v7) > 0 {
                        v7[0] = (^(fi6))
                    }
                } else {
                    v9 := GenS1{F0: 1, F1: 42}
                    p0 = p0
                    _ = v9
                }
                fmt.Println(p0, p0, fi8)
                _ = fi8
            }
            _ = fi6
            _ = v7
        }
        p0 -= (p0 / ((42) | 1))
    }
    switch "\n\t" {
    case "hello":
        for fi10 := 0; fi10 < 4; fi10++ {
            for fi11 := 0; fi11 < 12; fi11++ {
                switch "\n\t" {
                case "hello":
                    v12 := []int{(fi10 & 2147483647), fi11}
                    _ = v12
                case "foo bar":
                    var v13 uint32 = 984
                    _ = v13
                case "a":
                    var v14 uint = 0
                    _ = v14
                default:
                    fmt.Println(fi10, fi10, p0)
                }
                v15 := []int{fi10, 827, -2147483648}
                _ = fi11
                _ = v15
            }
            _ = fi10
        }
    case "foo bar":
        switch "foo bar" {
        case "hello":
            v16 := []int{127, 381, 1, 32768}
            _ = v16
        default:
            v17 := []int{-294, 2147483647, 32768}
            _ = v17
        }
    case "":
        v18 := GenS2{F0: 1}
        _ = v18
    default:
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 uint64 = 1
            if (-9223372036854775808 >= -1) {
                for fi21 := 0; fi21 < 18; fi21++ {
                    v20 += v20
                    fmt.Println(fi21)
                    v22 := GenS2{F0: 0}
                    var v23 int = 128
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
            _ = fi19
            _ = v20
        }
    }
    var v24 uint32 = 2147483648
    var v25 uint64 = 18446744073709551615
    var v26 uint = 65535
    _ = p0
    _ = v24
    _ = v25
    _ = v26
    return 127
}

func main() {
    var v1 int64 = 42
    v2 := make(map[uint32]int)
    v3 := []int{0, 1, 0}
    var v4 uint16 = 32767
    var v5 uint8 = 42
    fmt.Println(v1)
    v4 += v4
    for fi27 := 0; fi27 < 8; fi27++ {
        v4 += v4
        var v28 uint = 4294967295
        var v29 uint8 = (v5 >> ((18446744073709551615) & 63))
        _ = fi27
        _ = v28
        _ = v29
    }
    v1 -= (v1 / ((v1) | 1))
    if ((false && false) && (true || true)) {
        if (42 >= 2147483648) {
            for fi30 := 0; fi30 < 4; fi30++ {
                v1 = 0
                var v31 int8 = 100
                _ = fi30
                _ = v31
            }
        }
        var v32 uint8 = (v5 % ((255) | 1))
        for fi33 := 0; fi33 < 2; fi33++ {
            v34 := GenS2{F0: 4294967295}
            var v35 int8 = 1
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = v32
    } else {
        var v36 float64 = 1.5
        _ = v36
    }
    v37 := v2[42]
    v4 = v4
    if ((0 >= 42) && true) {
        fmt.Println(v1)
    } else {
        v4 = (v4 >> ((255) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
