package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 int8; F1 uint64 }
func genFn0() bool {
    v6 := []int{-1, 841, -2147483648, 32767}
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int64 = 1
        v9 := GenS1{F0: 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = len(v6)
    var v10 bool = true
    _ = v10
    _ = v6
    return v10
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "a"
    var v3 int16 = 15
    v4 := []uint32{0}
    v5 := make(map[uint32]string)
    var v11 uint8 = 42
    for fi12 := 0; fi12 < 1; fi12++ {
        if (1 != 537) {
            var v13 int16 = v3
            if len(v4) > 0 {
                v4[0] = 884
            }
            v14 := GenS1{F0: 1}
            _ = v13
            _ = v14
        }
        v15 := GenS2{F0: fi12}
        _ = fi12
        _ = v15
    }
    var v16 int16 = v3
    v17 := GenS2{F0: 32768}
    v3 -= (v3 * 1000)
    if false {
        if ((40 >= 255) && ((42 <= 0) || false)) {
            v18 := GenS2{F0: -2147483648}
            fmt.Println(v11, v3, v2)
            _ = v18
        }
        v19 := GenS1{F0: 0}
        v16 -= (v3 / ((v16) | 1))
        _ = v19
    }
    v16 = v3
    fmt.Println(v2, v3)
    v20 := GenS2{F0: 951}
    _ = len(v1)
    fmt.Println(v16)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v11)
    if (1 < 65535) {
        delete(v5, 2147483648)
    } else {
        fmt.Println(v16)
        switch int(v11) {
        case 32768:
            if true {
                if true {
                    var v21 float64 = ((1.5 - 1e10) * float64((0.1 / (1.5 * float32(1.5)))))
                    _ = v21
                }
                v22 := GenS3{F0: 42}
                _ = v22
            } else {
                if true {
                    v23 := GenS3{F0: 626}
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        _ = len(v5)
                        switch v2 {
                        case "foo bar":
                            var v25 int64 = 0
                            _ = v25
                        default:
                            v26 := GenS4{F0: v2, F1: -537}
                            _ = v26
                        }
                        _ = fi24
                    }
                }
                if false {
                    var v27 uint8 = v11
                    for fi28 := 0; fi28 < 12; fi28++ {
                        for fi29 := 0; fi29 < 17; fi29++ {
                            var v30 uint = 65535
                            _ = fi29
                            _ = v30
                        }
                        if (!(!(false || ((-456 > -2147483648) && ((2147483647 <= 0) || (889 == 0)))))) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v2 = v2
                            var v31 bool = false
                            _ = v31
                        }
                        v11 = (v27 & (v27 & (v27 | 127)))
                        _ = fi28
                    }
                    _ = v27
                } else {
                    v32 := GenS5{F0: 42, F1: 4294967295}
                    if (0.1 <= 0.1) {
                        for fi33 := 0; fi33 < 14; fi33++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi33
                        }
                        if false {
                            v11 += (^(v11))
                        } else {
                            fmt.Println(v11, v16)
                        }
                    }
                    _ = v32
                }
            }
        }
    }
    v11 = (^(v11))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
