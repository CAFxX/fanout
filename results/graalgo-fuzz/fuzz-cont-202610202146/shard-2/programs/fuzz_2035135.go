package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int32 }
func genFn0(p0 int16, p1 uint16) int64 {
    for fi5 := 0; fi5 < 9; fi5++ {
        switch "foo bar" {
        case "\n\t":
            if (3.14159 >= -2.25) {
                if ((1 != 18446744073709551615) && true) {
                    var v6 int64 = 342
                    _ = v6
                } else {
                    var v7 uint32 = 65535
                    if false {
                        switch (^(fi5)) {
                        case 127:
                            var v8 uint16 = 42
                            _ = v8
                        case -9223372036854775808:
                            var v9 int16 = (p0 - (p0 & 1))
                            _ = v9
                        default:
                            fmt.Println(p0)
                        }
                        p1 = p1
                    } else {
                        if (true || (0.0 < (1e10 + 0.0))) {
                            v10 := []int{fi5}
                            _ = v10
                        }
                        var v11 uint16 = (p1 ^ (p1 + p1))
                        _ = v11
                    }
                    _ = v7
                }
                var v12 uint = 14
                _ = v12
            }
        case "foo bar":
            p0 += (p0 / ((662) | 1))
        case "":
            v13 := []int{fi5, (fi5 | (fi5 & 42)), 101}
            _ = v13
        default:
            v14 := []int{0}
            _ = v14
        }
        _ = fi5
    }
    p1 = p1
    var v15 float32 = (-2.25 + (float32(1e10) * 1.5))
    if false {
        if true {
            if false {
                p0 += (p0 | -785)
            }
            if true {
                var v16 uint16 = (p1 + p1)
                fmt.Println(p1, p1)
                _ = v16
            }
            var v17 int16 = 0
            _ = v17
        }
        var v18 uint64 = 42
        var v19 int16 = p0
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: "", F1: 2, F2: -2147483648}
    _ = p0
    _ = p1
    _ = v15
    _ = v20
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 182
    var v3 int16 = -1
    var v4 uint = 4294967295
    fmt.Println(v4, v2, v2)
    if false {
        v1[128] = ("a" + "foo bar")
        v4 = 852
    }
    v21 := []int{699, 127, -9223372036854775808, -177}
    fmt.Println(v2)
    switch "foo bar" {
    case "foo bar":
        v4 = (v4 + (^(v4)))
    case "\n\t":
        if true {
            var v22 string = ("" + "hello")
            for fi23 := 0; fi23 < 5; fi23++ {
                var v24 int32 = -2147483648
                v4 -= 255
                v2 = 569
                _ = fi23
                _ = v24
            }
            _ = v22
        } else {
            var v25 uint = v4
            _ = v25
        }
    default:
        var v26 int32 = -222
        _ = v26
    }
    var v27 int64 = -9223372036854775808
    v3 += v3
    switch "" {
    case "\n\t":
        for fi28 := 0; fi28 < 17; fi28++ {
            if false {
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
                fmt.Println(v3, v4, v27)
                for fi29 := 0; fi29 < 5; fi29++ {
                    delete(v1, 128)
                    var v30 uint32 = 1
                    _ = fi29
                    _ = v30
                }
            }
            v27 += v27
            _ = fi28
        }
    default:
        v4 += 364
    }
    fmt.Println(v27)
    for fi31 := 0; fi31 < 9; fi31++ {
        for fi32 := 0; fi32 < 14; fi32++ {
            fmt.Println(fi32)
            v33 := GenS1{F0: "foo bar", F1: 42, F2: -599}
            v4 = (v4 >> (((^(v4))) & 63))
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    fmt.Println(v3, v3, v4)
    var v34 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v34
    _ = v4
}
