package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0(p0 uint8) int8 {
    if (!((32767 < 873) && false)) {
        var v5 uint16 = 496
        var v6 uint8 = (p0 << ((1) & 63))
        fmt.Println(v6, p0, p0)
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(p0, p0)
            var v8 int32 = 0
            v9 := []int{}
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := []int{-2147483648, 32768}
        _ = v10
    }
    fmt.Println(p0)
    switch int(p0) {
    case 127:
        v11 := GenS1{F0: 829, F1: 42}
        _ = v11
    case 4294967295:
        p0 = p0
    case 32768:
        for fi12 := 0; fi12 < 11; fi12++ {
            v13 := GenS1{F0: 420, F1: 333}
            if false {
                var v14 uint = 42
                _ = v14
            }
            p0 -= (p0 / ((194) | 1))
            _ = fi12
            _ = v13
        }
    }
    v15 := GenS1{F0: 342, F1: 255}
    _ = p0
    _ = v15
    return 59
}
func genFn1() int8 {
    v16 := GenS1{F0: 2, F1: 860}
    if (false && (0 >= 100)) {
        v17 := GenS1{F0: 42, F1: 747}
        _ = v17
    } else {
        var v18 int32 = -401
        v18 -= (-(v18))
        _ = v18
    }
    v19 := []int{-1, -686}
    var v20 float32 = 3.14159
    v20 -= float32(((0.1 + ((1e10 / (1.5 / (-2.25 - 0.1))) * 1e10)) - (0.1 / 0.1)))
    _ = v16
    _ = v19
    _ = v20
    return 42
}

func main() {
    var v1 string = ""
    v2 := []uint32{4294967295, 2147483648, 0, 4294967295}
    var v3 int8 = 127
    v4 := []int{112}
    if (!(((1.5 * 1.5) - 1.5) != 1e10)) {
        if true {
            if (272 < 1) {
                var v21 uint32 = 331
                _ = v21
            }
            var v22 string = v1
            _ = v22
        } else {
            fmt.Println(v3, v1)
        }
    }
    fmt.Println(v3)
    if (660 != 2) {
        if true {
            var v23 uint16 = 32767
            var v24 int64 = 42
            _ = v23
            _ = v24
        } else {
            _ = len(v4)
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 bool = true
                switch fi25 {
                case -1:
                    var v27 uint64 = 42
                    _ = v27
                case 1:
                    for fi28 := 0; fi28 < 18; fi28++ {
                        var v29 int64 = 9223372036854775807
                        if v26 {
                            v30 := GenS1{F0: 2, F1: 472}
                            if (!v26) {
                                v26 = v26
                                switch v1 {
                                case "foo bar":
                                    fmt.Println(v29)
                                }
                                v31 := GenS2{F0: 0, F1: 2147483648, F2: 18446744073709551615}
                                _ = v31
                            }
                            v26 = true
                            _ = v30
                        }
                        _ = fi28
                        _ = v29
                    }
                }
                _ = fi25
                _ = v26
            }
        }
        v3 = 1
    }
    var v32 float32 = 1e10
    switch 9223372036854775807 {
    case -418:
        fmt.Println(v1, v3)
    default:
        v33 := GenS3{F0: 4294967295, F1: -9223372036854775808}
        _ = v33
    }
    v2 = append(v2, 440)
    var v34 uint64 = 2
    var v35 uint16 = 65535
    var v36 string = v1
    var v37 uint = 4294967295
    var v38 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
