package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint8, p2 uint32) uint64 {
    v5 := []int{int(p1), 413, -797, 2147483647}
    var v6 int32 = -260
    var v7 uint32 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return p0
}

func main() {
    var v1 uint8 = 255
    var v2 string = "hello"
    var v3 float64 = (1e10 + 0.1)
    var v4 int16 = 1
    switch ((v2 + "\n\t") + v2) {
    case "\n\t":
        var v8 float64 = v3
        _ = v8
    case "hello":
        v9 := []int{-288, 9223372036854775807, 0, 0}
        _ = v9
    case "a":
        if false {
            v10 := []int{-614}
            if (("\n\t" == (v2 + "")) || false) {
                var v11 int16 = (v4 % ((0) | 1))
                v11 += (v4 / ((v4) | 1))
                _ = v11
            } else {
                var v12 int64 = -9223372036854775808
                fmt.Println(v3, v2, v3)
                _ = v12
            }
            if true {
                for fi13 := 0; fi13 < 10; fi13++ {
                    switch v2 {
                    case "hello":
                        if ((fi13 / ((fi13) | 1)) > (fi13 ^ fi13)) {
                            v1 += v1
                            if false {
                                if false {
                                    var v14 float64 = 0.1
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    var v15 uint8 = (v1 % ((v1) | 1))
                                    _ = v14
                                    _ = v15
                                } else {
                                    if false {
                                        var v16 uint = 1
                                        v16 = 65535
                                        v2 = v2
                                        _ = v16
                                    } else {
                                        v4 = v4
                                    }
                                    v2 = v2
                                    v2 = v2
                                }
                            } else {
                                if false {
                                    for fi17 := 0; fi17 < 6; fi17++ {
                                        v4 -= v4
                                        v3 += float64(3.14159)
                                        _ = fi17
                                    }
                                    v10 = append(v10, fi13)
                                } else {
                                    var v18 uint8 = (v1 / ((v1) | 1))
                                    var v19 uint16 = 0
                                    _ = v18
                                    _ = v19
                                }
                                v2 = (v2 + "")
                                for fi20 := 0; fi20 < 4; fi20++ {
                                    if true {
                                        if len(v10) > 0 {
                                            v10[0] = (-(fi20))
                                        }
                                        v10 = append(v10, (fi20 % ((42) | 1)))
                                        if (v1 != v1) {
                                            var v21 int32 = -866
                                            _ = v21
                                        } else {
                                            if true {
                                                var v22 int32 = 0
                                                if (-1 > int32(v4)) {
                                                    fmt.Println(v3, fi20, v4)
                                                    var v23 uint16 = 773
                                                    if len(v10) >= 2 {
                                                        v10 = v10[1:len(v10)]
                                                    }
                                                    _ = v23
                                                }
                                                _ = v22
                                            } else {
                                                v1 -= (v1 * 175)
                                            }
                                            fmt.Println(v3, v1, fi20)
                                        }
                                    }
                                    v1 = (v1 ^ 255)
                                    v24 := GenS1{F0: (fi20 / ((fi20) | 1))}
                                    _ = fi20
                                    _ = v24
                                }
                            }
                            fmt.Println(v4, v1, v2)
                        }
                    }
                    _ = fi13
                }
            }
            _ = v10
        }
    }
    if false {
        for fi25 := 0; fi25 < 19; fi25++ {
            var v26 uint16 = 1
            _ = fi25
            _ = v26
        }
        var v27 int64 = -9223372036854775808
        _ = v27
    }
    v28 := GenS1{F0: 523}
    var v29 uint32 = 2147483648
    v30 := []int{}
    if true {
        v31 := GenS2{F0: 443, F1: v29, F2: 0}
        v1 = (v1 << ((0) & 63))
        _ = v31
    }
    v32 := GenS3{F0: v1}
    var v33 uint32 = v29
    v34 := GenS4{F0: 2}
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
