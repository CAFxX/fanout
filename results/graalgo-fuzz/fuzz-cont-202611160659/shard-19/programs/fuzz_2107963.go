package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int64; F1 int16; F2 bool }
type GenS4 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint; F1 int64 }
func genFn0(p0 string, p1 uint32, p2 bool) int {
    var v5 uint16 = 0
    p0 = "a"
    v6 := []int{0, -2147483648, 32767, -970}
    var v7 uint32 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 128
}
func genFn1(p0 string) int8 {
    if false {
        var v8 int32 = 2147483647
        var v9 float32 = 3.14159
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: 1, F1: (-254 < -1)}
        v11 := []int{-643, -517}
        v11 = append(v11, 1)
        _ = v10
        _ = v11
    }
    fmt.Println(p0)
    switch -728 {
    case -9223372036854775808:
        p0 = (p0 + p0)
    case 0:
        switch 101 {
        case 4294967295:
            if true {
                for fi12 := 0; fi12 < 7; fi12++ {
                    fmt.Println(fi12, fi12)
                    switch (p0 + ("a" + (p0 + p0))) {
                    case "":
                        for fi13 := 0; fi13 < 3; fi13++ {
                            p0 = p0
                            var v14 float64 = 0.0
                            var v15 uint32 = 4294967295
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    case "foo bar":
                        if true {
                            v16 := GenS1{F0: 100, F1: true}
                            p0 = p0
                            p0 = p0
                            _ = v16
                        } else {
                            v17 := []int{}
                            p0 = p0
                            for fi18 := 0; fi18 < 15; fi18++ {
                                var v19 uint32 = 0
                                var v20 uint64 = 4294967295
                                if true {
                                    var v21 uint16 = 65535
                                    var v22 float32 = (0.1 / ((1e10 / 1.0) / 0.1))
                                    _ = v21
                                    _ = v22
                                }
                                v23 := GenS2{F0: -372, F1: 42}
                                _ = fi18
                                _ = v19
                                _ = v20
                                _ = v23
                            }
                            _ = v17
                        }
                    case "a":
                        var v24 bool = (!(true && false))
                        _ = v24
                    default:
                        var v25 float32 = -2.25
                        _ = v25
                    }
                    _ = fi12
                }
            } else {
                v26 := GenS2{F0: -1, F1: 475}
                _ = v26
            }
        case 9223372036854775807:
            var v27 bool = true
            _ = v27
        }
    default:
        for fi28 := 0; fi28 < 12; fi28++ {
            if (127 >= 42) {
                var v29 uint32 = 65535
                v29 -= v29
                _ = v29
            } else {
                fmt.Println(fi28, fi28)
                if false {
                    for fi30 := 0; fi30 < 3; fi30++ {
                        if true {
                            if (42 < 552) {
                                if true {
                                    if true {
                                        v31 := []int{fi30, (fi28 % ((fi30) | 1)), (^(fi30))}
                                        fmt.Println(fi30)
                                        if len(v31) >= 2 {
                                            v31 = v31[1:len(v31)]
                                        }
                                        _ = v31
                                    } else {
                                        p0 = (p0 + "hello")
                                        var v32 uint = 1
                                        _ = v32
                                    }
                                    v33 := GenS3{F0: -9223372036854775808, F1: 1, F2: (((-9223372036854775808 > -9223372036854775808) && (1 != 32767)) && false)}
                                    var v34 int64 = -442
                                    _ = v33
                                    _ = v34
                                } else {
                                    v35 := GenS4{F0: true, F1: 853, F2: 4294967295}
                                    _ = v35
                                }
                                if true {
                                    v36 := []int{fi28}
                                    var v37 string = p0
                                    v38 := GenS3{F0: -859, F1: 0, F2: ((100 < 0) && (4294967295 >= 2))}
                                    v39 := GenS3{F0: 1, F1: 376, F2: false}
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                            }
                        } else {
                            fmt.Println(fi28, p0, fi28)
                            v40 := GenS4{F0: false, F1: 1, F2: 18446744073709551615}
                            _ = v40
                        }
                        _ = fi30
                    }
                }
                var v41 float32 = 1.5
                _ = v41
            }
            var v42 int8 = -1
            _ = fi28
            _ = v42
        }
    }
    p0 = p0
    _ = p0
    return -128
}

func main() {
    var v1 uint32 = 2
    var v2 float64 = (3.14159 / ((3.14159 - (float64(((3.14159 + (float32(3.14159) + 1e10)) * 0.1)) / (0.1 - 0.0))) / 1.5))
    v3 := make(map[int]int)
    v4 := []int{int(v1), 32768}
    v1 = v1
    fmt.Println(v1, v1)
    v43 := GenS2{F0: 9223372036854775807, F1: 42}
    if (true || (2147483647 < (v1 + v1))) {
        fmt.Println(v1, v2)
    }
    v1 = (v1 * v1)
    if (-668 >= 1000) {
        if false {
            v2 -= v2
            v44 := v3[-2147483648]
            _ = v44
        }
    }
    v45 := GenS4{F0: false, F1: 32767, F2: 398}
    var v46 int64 = -1
    if false {
        var v47 string = ("" + "a")
        var v48 int = 2147483647
        _ = v47
        _ = v48
    }
    if false {
        v1 = 2147483647
    }
    v49 := GenS2{F0: v46, F1: 18446744073709551615}
    for fi50 := 0; fi50 < 5; fi50++ {
        v46 = v46
        v51 := GenS5{F0: 42, F1: (-(v46))}
        _ = fi50
        _ = v51
    }
    var v52 int8 = int8(v46)
    v52 = v52
    v52 -= (-(v52))
    v52 = v52
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v45
    _ = v46
    _ = v49
    _ = v52
}
