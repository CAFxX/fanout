package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint64 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    var v3 uint64 = 2
    var v4 float32 = 0.1
    v5 := []int{}
    fmt.Println(v4, v3)
    var v6 uint16 = 291
    fmt.Println(v4, v6)
    var v7 uint8 = 0
    if false {
        for fi8 := 0; fi8 < 10; fi8++ {
            v3 += 42
            fmt.Println(v3, v6, fi8)
            _ = fi8
        }
        fmt.Println(v4, v4, v3)
    } else {
        v7 -= 127
        if (float32(0.1) > (-2.25 / (-2.25 - (v4 + (v4 + (v4 + (v4 + float32(1e10)))))))) {
            if false {
                for fi9 := 0; fi9 < 2; fi9++ {
                    v4 = (v4 + v4)
                    _ = fi9
                }
                if false {
                    v2[2147483647] = 513
                    v6 += v6
                    switch ("hello" + "foo bar") {
                    case "foo bar":
                        v10 := GenS1{F0: 368}
                        _ = v10
                    default:
                        v11 := GenS2{F0: 156, F1: 2}
                        _ = v11
                    }
                }
                v7 = v7
            } else {
                if (!false) {
                    if (254 != v7) {
                        var v12 uint64 = 42
                        _ = v12
                    } else {
                        if (2147483647 != -2147483648) {
                            if len(v5) > 0 {
                                v5[0] = -9223372036854775808
                            }
                            _ = len(v5)
                            for fi13 := 0; fi13 < 4; fi13++ {
                                if len(v5) > 0 {
                                    v5[0] = fi13
                                }
                                fmt.Println(v4)
                                v14 := GenS1{F0: 4294967295}
                                _ = fi13
                                _ = v14
                            }
                        } else {
                            delete(v1, 812)
                            fmt.Println(v7, v6)
                        }
                        v6 -= v6
                    }
                    v6 += 1
                } else {
                    fmt.Println(v7)
                    var v15 uint64 = v3
                    v16 := GenS1{F0: 65535}
                    _ = v15
                    _ = v16
                }
            }
            var v17 bool = true
            var v18 float64 = float64((float32(3.14159) - 1.5))
            _ = v17
            _ = v18
        }
    }
    fmt.Println(v6, v4)
    v1[-81] = 128
    v7 = 128
    v2[-2147483648] = 178
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 float32 = v4
        v20 = float32(3.14159)
        _ = fi19
        _ = v20
    }
    var v21 string = "\n\t"
    var v22 int = 32767
    v2[v22] = 4294967295
    v23 := v2[986]
    v7 = 201
    if len(v5) > 0 {
        v5[0] = (-(v22))
    }
    v23 = v23
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
