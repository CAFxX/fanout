package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int; F1 int }
func genFn0(p0 int32, p1 int64) uint16 {
    v6 := []int{128}
    var v7 int64 = p1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 618
}
func genFn1() uint8 {
    var v8 uint16 = 42
    v9 := GenS1{F0: 2147483648}
    v8 = genFn0(2147483647, -9223372036854775808)
    var v10 int64 = 9223372036854775807
    _ = v10
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 string = "a"
    var v2 int = -685
    v3 := make(map[int]int)
    v4 := make(map[string]int)
    v5 := make(map[uint32]bool)
    v11 := []int{}
    if (v1 != v1) {
        v12 := GenS2{F0: 0}
        _ = v12
    }
    if false {
        v1 = v1
        switch ("\n\t" + "") {
        case "hello":
            fmt.Println(v2, v2)
        case "foo bar":
            v2 = 42
        default:
            if (42 < 42) {
                for fi13 := 0; fi13 < 12; fi13++ {
                    v2 = 127
                    var v14 uint64 = 839
                    if false {
                        if (!(!false)) {
                            if (127 <= genFn1()) {
                                v2 = fi13
                                fmt.Println(v1, v14)
                                for fi15 := 0; fi15 < 12; fi15++ {
                                    switch v1 {
                                    case "a":
                                        var v16 int64 = int64(fi13)
                                        _ = v16
                                    }
                                    v11 = append(v11, (v2 % (((^(fi13))) | 1)))
                                    _ = fi15
                                }
                            }
                            for fi17 := 0; fi17 < 9; fi17++ {
                                var v18 int32 = 1
                                v19 := GenS3{F0: 1, F1: 127}
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                            v14 = (^(v14))
                        }
                        v11 = append(v11, fi13)
                    } else {
                        var v20 int8 = 100
                        v20 -= (v20 - (v20 % ((v20) | 1)))
                        fmt.Println(v1, v2)
                        _ = v20
                    }
                    _ = fi13
                    _ = v14
                }
                v2 += (v2 - v2)
                v21 := GenS2{F0: 360}
                _ = v21
            }
        }
    }
    v2 += 0
    v22 := GenS4{F0: (v2 + (v2 * (^(v2)))), F1: v2}
    _ = len(v3)
    if ((255 <= 396) || (true && false)) {
        var v23 float64 = (0.0 - (0.1 * 3.14159))
        _ = v23
    } else {
        fmt.Println(v2, v2)
    }
    var v24 int = v2
    fmt.Println(v2)
    for fi25 := 0; fi25 < 8; fi25++ {
        v24 += -9223372036854775808
        _ = fi25
    }
    v26 := GenS1{F0: 65535}
    v11 = append(v11, v24)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
