package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int8; F2 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 int; F2 string }
type GenS6 struct { F0 uint; F1 int64; F2 uint }
type GenS7 struct { F0 int64; F1 uint }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint64; F1 uint; F2 uint }

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    if true {
        v5 := []int{2147483647, 32767, -1}
        if (!false) {
            for fi6 := 0; fi6 < 11; fi6++ {
                if (1 < 185) {
                    switch "foo bar" {
                    case "foo bar":
                        v1 = (v1 - v1)
                    default:
                        var v7 uint64 = 4294967295
                        _ = v7
                    }
                    var v8 float64 = 3.14159
                    v8 = v8
                    _ = v8
                }
                if false {
                    if ((380 != (fi6 | (-(fi6)))) || true) {
                        v1 -= (v1 + 2147483647)
                        v9 := GenS1{F0: (v1 >> ((579) & 63))}
                        _ = v9
                    }
                    if (255 >= 2) {
                        v1 -= v1
                        v10 := GenS1{F0: v1}
                        var v11 uint = 65535
                        _ = v10
                        _ = v11
                    } else {
                        var v12 string = "a"
                        if (1 > -1) {
                            if len(v5) > 0 {
                                v5[0] = 32768
                            }
                        }
                        v12 = v12
                        _ = v12
                    }
                } else {
                    fmt.Println(fi6, v1)
                    v1 = (^(v1))
                }
                _ = fi6
            }
            _ = len(v5)
        } else {
            for fi13 := 0; fi13 < 20; fi13++ {
                if false {
                    if (705 < -247) {
                        if true {
                            v1 -= v1
                        }
                        v14 := GenS2{F0: 0}
                        _ = v14
                    } else {
                        v1 += 364
                    }
                    if true {
                        if false {
                            v1 -= (v1 + v1)
                        } else {
                            if false {
                                for fi15 := 0; fi15 < 17; fi15++ {
                                    v1 += 2147483647
                                    fmt.Println(fi13)
                                    v16 := GenS3{F0: v1, F1: 1, F2: 0}
                                    v1 = v1
                                    _ = fi15
                                    _ = v16
                                }
                                v1 = v1
                                v17 := GenS1{F0: 2147483648}
                                _ = v17
                            }
                            if len(v5) > 0 {
                                v5[0] = 42
                            }
                            if false {
                                if len(v5) > 0 {
                                    v5[0] = fi13
                                }
                            }
                        }
                        var v18 int64 = 0
                        if false {
                            v19 := GenS4{F0: (v18 | v18)}
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v19
                        }
                        _ = v18
                    }
                    switch "foo bar" {
                    case "":
                        if true {
                            switch fi13 {
                            case 2147483647:
                                var v20 int16 = -32768
                                _ = v20
                            case 77:
                                for fi21 := 0; fi21 < 19; fi21++ {
                                    switch fi13 {
                                    case 2147483647:
                                        v1 = v1
                                    case 42:
                                        v1 -= 2
                                    default:
                                        v22 := GenS5{F0: 18446744073709551615, F1: (fi21 | fi21), F2: "hello"}
                                        _ = v22
                                    }
                                    _ = fi21
                                }
                            case 4294967295:
                                var v23 int8 = -94
                                _ = v23
                            default:
                                _ = len(v2)
                            }
                            for fi24 := 0; fi24 < 19; fi24++ {
                                v1 -= 2147483648
                                switch "a" {
                                case "hello":
                                    if true {
                                        switch fi13 {
                                        case 2147483647:
                                            var v25 uint = 147
                                            fmt.Println(v25, fi24, v1)
                                            _ = v25
                                        }
                                        var v26 int = (-(fi24))
                                        if false {
                                            v26 += (v26 % ((v26) | 1))
                                            v26 = (fi13 << ((65535) & 63))
                                        } else {
                                            v5 = append(v5, 2147483647)
                                            fmt.Println(v26, v26)
                                        }
                                        _ = v26
                                    }
                                case "foo bar":
                                    if (110 == 0) {
                                        var v27 int32 = -39
                                        v28 := GenS6{F0: 484, F1: 9223372036854775807, F2: 1}
                                        _ = v27
                                        _ = v28
                                    }
                                }
                                _ = fi24
                            }
                        }
                    default:
                        var v29 uint16 = 255
                        _ = v29
                    }
                } else {
                    v1 = (v1 >> ((1000) & 63))
                }
                fmt.Println(v1)
                for fi30 := 0; fi30 < 15; fi30++ {
                    v31 := GenS7{F0: -1, F1: 2}
                    delete(v4, "hello")
                    _ = fi30
                    _ = v31
                }
                _ = fi13
            }
            v1 = 595
        }
        v1 = 42
        _ = v5
    } else {
        fmt.Println(v1)
    }
    if (!(0 <= 75)) {
        v1 = (v1 ^ v1)
        v32 := []int{-64, 544, 128, 0}
        _ = v32
    } else {
        v33 := GenS8{F0: 100}
        fmt.Println(v1)
        v1 = v1
        _ = v33
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        _ = len(v3)
        _ = fi34
    }
    v1 = (v1 % ((42) | 1))
    if ((-2147483648 <= 1) || true) {
        v35 := GenS9{F0: -9223372036854775808}
        _ = v35
    } else {
        switch 1 {
        case -754:
            var v36 int16 = 42
            _ = v36
        }
        if (-1 >= 1000) {
            v37 := GenS10{F0: true}
            _ = v37
        }
    }
    v1 += v1
    fmt.Println(v1)
    v38 := GenS11{F0: uint64(v1), F1: 4294967295, F2: 18446744073709551615}
    var v39 int64 = -914
    v40 := GenS7{F0: (v39 | -961), F1: 65535}
    var v41 int = -82
    var v42 uint64 = 18446744073709551615
    var v43 int16 = 655
    delete(v2, 42)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
