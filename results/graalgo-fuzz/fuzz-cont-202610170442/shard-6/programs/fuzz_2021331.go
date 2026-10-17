package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 uint; F1 string; F2 string }
type GenS8 struct { F0 uint8 }
func genFn0(p0 int8, p1 int64, p2 int64) int64 {
    var v4 int8 = p0
    switch "foo bar" {
    case "a":
        var v5 int8 = (v4 - -1)
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (-(p2))
}
func genFn1(p0 bool) uint8 {
    v6 := GenS1{F0: -2147483648}
    v7 := []int{-9223372036854775808, 486, 815}
    var v8 int8 = 0
    _ = len(v7)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []int{4294967295, 269, 42}
    var v2 bool = true
    var v3 string = "a"
    var v9 int32 = 42
    var v10 int64 = 0
    v1 = append(v1, -9223372036854775808)
    v11 := GenS1{F0: v9}
    v1 = append(v1, 42)
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := GenS2{F0: 42}
            fmt.Println(v10, v3, v3)
            if (v2 || true) {
                fmt.Println(fi13, fi12)
                var v15 uint = 0
                v9 -= -2147483648
                _ = v15
            } else {
                switch (v3 + v3) {
                case "":
                    switch "hello" {
                    case "\n\t":
                        if len(v1) > 0 {
                            v1[0] = -1
                        }
                    case "hello":
                        v10 += v10
                    default:
                        fmt.Println(v9, v2, v3)
                    }
                case "\n\t":
                    v10 = (v10 << ((0) & 63))
                case "hello":
                    if false {
                        var v16 uint64 = 563
                        _ = v16
                    }
                }
                var v17 string = v3
                _ = v17
            }
            _ = fi13
            _ = v14
        }
        if (100 < -128) {
            v2 = ((0.0 + 0.1) != (0.1 + float64(((0.0 / 1.0) * -2.25))))
        } else {
            var v18 int32 = (-(v9))
            var v19 int = 128
            v20 := GenS3{F0: true}
            _ = v18
            _ = v19
            _ = v20
        }
        _ = fi12
    }
    v2 = (float32(0.1) < -2.25)
    if v2 {
        if ("" > (v3 + "a")) {
            if (2 <= 203) {
                var v21 int16 = 1000
                if (0 >= 846) {
                    v22 := GenS2{F0: 65535}
                    _ = v22
                }
                for fi23 := 0; fi23 < 15; fi23++ {
                    fmt.Println(v2, v9)
                    _ = fi23
                }
                _ = v21
            } else {
                if (487 != 0) {
                    if v2 {
                        fmt.Println(v9, v3)
                        v24 := GenS4{F0: 851, F1: v2}
                        v25 := GenS5{F0: v10}
                        _ = v24
                        _ = v25
                    }
                } else {
                    if (-7 < 0) {
                        var v26 int64 = v10
                        v1 = append(v1, int(v26))
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 6; fi27++ {
                            for fi28 := 0; fi28 < 11; fi28++ {
                                v3 = "hello"
                                _ = fi28
                            }
                            if true {
                                _ = len(v1)
                                v3 = ((v3 + "a") + "\n\t")
                                var v29 uint = 126
                                _ = v29
                            } else {
                                v2 = true
                                for fi30 := 0; fi30 < 20; fi30++ {
                                    var v31 float32 = 3.14159
                                    v32 := GenS3{F0: ((v9 << ((65535) & 63)) >= (v9 + v9))}
                                    if v2 {
                                        fmt.Println(fi30)
                                        v33 := GenS6{F0: fi30, F1: -1}
                                        if len(v1) > 0 {
                                            v1[0] = (^(fi27))
                                        }
                                        var v34 bool = (((genFn0(-62, (-(v10)), v10) <= 1) || false) || v2)
                                        v2 = (!((v10 & v10) < v10))
                                        _ = v33
                                        _ = v34
                                    }
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                }
                                fmt.Println(fi27, v2, v10)
                            }
                            _ = fi27
                        }
                        if v2 {
                            if (v2 && true) {
                                v10 += v10
                                fmt.Println(v9, v2, v10)
                                var v35 int32 = (v9 << ((18446744073709551615) & 63))
                                _ = v35
                            } else {
                                var v36 uint16 = 255
                                v9 = v9
                                _ = v36
                            }
                            switch (v3 + "") {
                            case "foo bar":
                                v10 += v10
                            case "":
                                v2 = v2
                            default:
                                var v37 int32 = (v9 % ((49) | 1))
                                _ = v37
                            }
                            fmt.Println(v2, v3)
                        } else {
                            v38 := GenS7{F0: 1, F1: v3, F2: (v3 + v3)}
                            v9 -= (v9 & v9)
                            var v39 int64 = v10
                            _ = v38
                            _ = v39
                        }
                        v10 -= 808
                    }
                    v40 := GenS2{F0: 0}
                    for fi41 := 0; fi41 < 3; fi41++ {
                        v42 := GenS6{F0: (fi41 ^ fi41), F1: 42}
                        _ = fi41
                        _ = v42
                    }
                    _ = v40
                }
                v43 := GenS8{F0: 56}
                if len(v1) > 0 {
                    v1[0] = 32767
                }
                _ = v43
            }
            switch -9223372036854775808 {
            case 127:
                for fi44 := 0; fi44 < 13; fi44++ {
                    v2 = v2
                    _ = fi44
                }
            default:
                var v45 uint = 855
                _ = v45
            }
        } else {
            fmt.Println(v3, v10, v10)
            var v46 float32 = float32((-2.25 * ((-2.25 * (float64(0.0) / 1.0)) * 0.1)))
            _ = v46
        }
        for fi47 := 0; fi47 < 14; fi47++ {
            var v48 uint8 = genFn1((!false))
            _ = len(v1)
            _ = fi47
            _ = v48
        }
    } else {
        if ((100 != 100) && (!v2)) {
            v2 = (!true)
        }
    }
    v1 = append(v1, -2147483648)
    if (2147483647 == 32767) {
        v49 := GenS6{F0: 128, F1: 100}
        _ = v49
    }
    var v50 int32 = v9
    var v51 uint16 = 42
    fmt.Println(v50, v9)
    v52 := GenS1{F0: (v9 * 1)}
    v2 = (!true)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v50
    _ = v51
    _ = v52
    _ = v9
}
