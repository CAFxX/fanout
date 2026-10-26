package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int32 = 0
    var v2 int64 = -9223372036854775808
    var v3 uint64 = 42
    v4 := make(map[uint32]bool)
    if (-1 >= 0) {
        v3 -= v3
        v1 += v1
        if false {
            for fi5 := 0; fi5 < 6; fi5++ {
                fmt.Println(fi5)
                _ = fi5
            }
            v4[2147483648] = false
            if false {
                v1 = v1
                var v6 int16 = -1
                _ = v6
            } else {
                if false {
                    v3 = (v3 | v3)
                    v3 = v3
                    if (true && true) {
                        for fi7 := 0; fi7 < 12; fi7++ {
                            if (1.5 == float32(float64(0.1))) {
                                if ((-(v2)) <= 0) {
                                    v3 -= v3
                                } else {
                                    v8 := []int{}
                                    _ = v8
                                }
                                v3 = (v3 + (v3 + v3))
                            } else {
                                v9 := GenS1{F0: (v1 | v1), F1: v1}
                                v10 := GenS1{F0: 1, F1: (v1 % ((v1) | 1))}
                                _ = v10
                                _ = v9
                            }
                            _ = fi7
                        }
                    } else {
                        v11 := []int{32767, -812}
                        _ = v11
                    }
                } else {
                    if false {
                        if false {
                            fmt.Println(v2)
                        } else {
                            v1 = (v1 & -585)
                            var v12 uint64 = v3
                            _ = v12
                        }
                        if ((v3 / ((v3) | 1)) < v3) {
                            for fi13 := 0; fi13 < 14; fi13++ {
                                v1 = (v1 - (v1 - v1))
                                _ = fi13
                            }
                            v14 := []int{0, 32767}
                            _ = v14
                        }
                    }
                    v15 := []int{-2147483648, 127}
                    _ = v15
                }
                switch 195 {
                case -9223372036854775808:
                    fmt.Println(v1, v1)
                }
            }
        }
    }
    fmt.Println(v2, v2)
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := []int{int(v2)}
        _ = fi16
        _ = v17
    }
    var v18 int64 = (v2 >> ((0) & 63))
    fmt.Println(v2, v2)
    if false {
        switch int(v1) {
        case 4294967295:
            v18 += v18
        }
    } else {
        v19 := []int{1, 4294967295, -2147483648}
        v1 = (v1 & (v1 >> ((65535) & 63)))
        _ = v19
    }
    v1 -= (v1 >> ((380) & 63))
    var v20 uint = 276
    v21 := GenS1{F0: v1, F1: 1}
    v2 -= -483
    var v22 int8 = 100
    fmt.Println(v18, v2, v18)
    switch "a" {
    case "":
        var v23 bool = (1 <= 0)
        _ = v23
    }
    v20 = (v20 / ((18446744073709551615) | 1))
    if false {
        var v24 uint = (v20 & 163)
        _ = v24
    } else {
        v2 += v2
        var v25 int64 = (v18 >> ((v20) & 63))
        for fi26 := 0; fi26 < 13; fi26++ {
            if true {
                v27 := GenS1{F0: v1, F1: v1}
                v28 := GenS1{F0: v1, F1: v1}
                if (false && true) {
                    switch (fi26 | fi26) {
                    case 127:
                        for fi29 := 0; fi29 < 15; fi29++ {
                            v3 -= (v3 / ((42) | 1))
                            if true {
                                var v30 uint16 = 545
                                v1 = int32(v30)
                                _ = v30
                            } else {
                                var v31 uint16 = 42
                                v25 = v18
                                _ = v31
                            }
                            _ = fi29
                        }
                    case 128:
                        if ((!(!true)) || false) {
                            var v32 int = int(v3)
                            _ = v32
                        }
                    }
                    v20 = uint(v3)
                    v1 += int32(v20)
                } else {
                    fmt.Println(v2, v25, v1)
                    var v33 uint16 = 255
                    _ = v33
                }
                _ = v27
                _ = v28
            } else {
                var v34 float64 = (-2.25 * 0.0)
                v35 := GenS2{F0: 0}
                delete(v4, 4294967295)
                _ = v34
                _ = v35
            }
            if false {
                fmt.Println(v18)
                var v36 float64 = (1.5 - 1.5)
                var v37 int = fi26
                _ = v36
                _ = v37
            } else {
                v38 := GenS2{F0: 1}
                var v39 uint32 = 2
                var v40 float64 = 3.14159
                _ = v38
                _ = v39
                _ = v40
            }
            v22 = 0
            _ = fi26
        }
        _ = v25
    }
    if (false && false) {
        var v41 string = "\n\t"
        _ = v41
    } else {
        v1 = (^(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
