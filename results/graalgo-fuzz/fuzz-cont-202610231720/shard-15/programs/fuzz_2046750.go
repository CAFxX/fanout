package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint32 }
type GenS2 struct { F0 bool; F1 int32; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 uint8; F1 int64 }

func main() {
    v1 := []int{42, -309, -902}
    v2 := make(map[uint32]string)
    var v3 uint32 = 42
    var v4 uint16 = 65535
    var v5 uint8 = 127
    v6 := make(map[string]uint32)
    fmt.Println(v3, v3, v3)
    v7 := GenS1{F0: 42, F1: 106, F2: v3}
    var v8 uint = 0
    _ = len(v6)
    if false {
        var v9 int64 = -675
        _ = v9
    } else {
        if false {
            if false {
                v10 := GenS1{F0: (v8 | (v8 | v8)), F1: 100, F2: v3}
                v3 = (v3 ^ 0)
                v11 := GenS2{F0: true, F1: 666, F2: ""}
                _ = v10
                _ = v11
            } else {
                for fi12 := 0; fi12 < 18; fi12++ {
                    v2[uint32(v4)] = "\n\t"
                    v8 = (v8 << ((238) & 63))
                    if (!true) {
                        v8 -= 0
                        v1 = append(v1, (fi12 | fi12))
                        var v13 uint16 = (v4 / ((v4) | 1))
                        _ = v13
                    }
                    _ = fi12
                }
            }
            if false {
                v8 = v8
                v14 := GenS3{F0: 42}
                _ = v14
            }
        } else {
            if true {
                switch "foo bar" {
                case "a":
                    if (368 == -174) {
                        v4 = 651
                        var v15 uint64 = 42
                        v5 = uint8(v8)
                        _ = v15
                    } else {
                        v16 := GenS2{F0: true, F1: -473, F2: "hello"}
                        v4 -= v4
                        var v17 int64 = -9223372036854775808
                        _ = v16
                        _ = v17
                    }
                default:
                    switch "" {
                    case "":
                        switch "" {
                        case "\n\t":
                            fmt.Println(v5, v8, v3)
                        case "a":
                            var v18 uint64 = 260
                            _ = v18
                        default:
                            var v19 uint = 202
                            _ = v19
                        }
                    case "\n\t":
                        v20 := GenS4{F0: -9223372036854775808}
                        _ = v20
                    }
                }
                v21 := GenS1{F0: v8, F1: 42, F2: v3}
                _ = v21
            } else {
                var v22 int64 = -1
                _ = v22
            }
            v1 = append(v1, 152)
            var v23 bool = (1 > v4)
            _ = v23
        }
        v8 = 2
    }
    v6["a"] = v3
    var v24 int = 32768
    var v25 string = "a"
    if false {
        switch v24 {
        case 9223372036854775807:
            for fi26 := 0; fi26 < 6; fi26++ {
                if false {
                    fmt.Println(fi26)
                    if false {
                        var v27 bool = (-2147483648 <= -2147483648)
                        v5 += (v5 * 145)
                        if ("a" >= "a") {
                            fmt.Println(v25, v5, v8)
                            v27 = false
                        }
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 10; fi28++ {
                            if (false && true) {
                                v8 += (v8 + v8)
                                var v29 int16 = 42
                                fmt.Println(fi28, v5)
                                _ = v29
                            }
                            v30 := GenS3{F0: 405}
                            v1 = append(v1, (fi26 % ((v24) | 1)))
                            _ = fi28
                            _ = v30
                        }
                        fmt.Println(v25, v5, v8)
                        var v31 int = v24
                        _ = v31
                    }
                }
                v32 := GenS1{F0: (v8 & (v8 * (v8 << ((v8) & 63)))), F1: -1, F2: v3}
                for fi33 := 0; fi33 < 7; fi33++ {
                    v34 := v2[65535]
                    _ = fi33
                    _ = v34
                }
                _ = fi26
                _ = v32
            }
        case -9223372036854775808:
            fmt.Println(v5)
        }
        v35 := GenS5{F0: false, F1: -9223372036854775808}
        _ = v35
    } else {
        v36 := GenS6{F0: v5, F1: -9223372036854775808}
        _ = v36
    }
    v4 -= v4
    v3 += 0
    v24 -= (-(v24))
    v37 := GenS6{F0: 128, F1: 217}
    fmt.Println(v4, v5, v3)
    v38 := GenS4{F0: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
