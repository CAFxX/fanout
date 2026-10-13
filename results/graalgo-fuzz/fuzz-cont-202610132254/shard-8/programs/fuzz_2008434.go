package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint; F2 int64 }
func genFn0(p0 uint64, p1 uint, p2 int8) string {
    p0 = p0
    fmt.Println(p1, p1, p1)
    var v7 uint16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "\n\t"
}
func genFn1(p0 uint, p1 int16) int32 {
    v8 := GenS1{F0: 1, F1: uint8(p0), F2: uint8(p0)}
    if true {
        v9 := []int{}
        v9 = append(v9, -334)
        _ = v9
    } else {
        p0 += 4294967295
        v10 := []int{42}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v8
    return 322
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (float64(1e10) * -2.25)
    var v3 uint16 = 1
    var v4 float64 = 3.14159
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]bool)
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 uint64 = 4294967295
        v2 = 3.14159
        if (v3 <= 255) {
            fmt.Println(v4, v2, v2)
            if (32767 != 196) {
                var v13 float64 = (v2 - v4)
                fmt.Println(v3)
                v14 := GenS1{F0: 42, F1: 0, F2: 42}
                _ = v13
                _ = v14
            } else {
                fmt.Println(v4, fi11)
                var v15 float32 = (0.1 - 0.1)
                fmt.Println(v4)
                _ = v15
            }
        }
        _ = fi11
        _ = v12
    }
    var v16 float32 = 0.0
    var v17 uint = 524
    for fi18 := 0; fi18 < 18; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            v2 = 0.1
            _ = fi19
        }
        _ = fi18
    }
    var v20 uint16 = (v3 | v3)
    if false {
        if true {
            fmt.Println(v3, v4, v20)
        } else {
            v21 := []int{-61}
            _ = v21
        }
        v22 := []int{-988, -1, -289}
        _ = v22
    }
    if (57 >= 1) {
        fmt.Println(v4, v20)
    } else {
        v3 = (v20 / ((65535) | 1))
        var v23 int = 32768
        if true {
            v24 := GenS2{F0: (147 > 255)}
            _ = v24
        }
        _ = v23
    }
    v3 += 181
    switch 0 {
    case -980:
        v25 := v1["a"]
        _ = v25
    case 32767:
        var v26 uint = v17
        _ = v26
    case 469:
        v27 := GenS1{F0: 1, F1: 42, F2: 42}
        _ = v27
    default:
        for fi28 := 0; fi28 < 20; fi28++ {
            for fi29 := 0; fi29 < 8; fi29++ {
                v5[2147483648] = -822
                var v30 string = "hello"
                _ = fi29
                _ = v30
            }
            _ = fi28
        }
    }
    switch "hello" {
    case "foo bar":
        var v31 int64 = -9223372036854775808
        _ = v31
    case "a":
        for fi32 := 0; fi32 < 8; fi32++ {
            for fi33 := 0; fi33 < 6; fi33++ {
                v34 := GenS3{F0: (v20 >> ((1) & 63)), F1: (v17 | (v17 << ((v17) & 63))), F2: 42}
                if true {
                    v2 += ((0.0 + v2) + v4)
                }
                _ = fi33
                _ = v34
            }
            v6[2147483647] = true
            v35 := GenS2{F0: false}
            _ = fi32
            _ = v35
        }
    }
    if (-128 > -128) {
        for fi36 := 0; fi36 < 1; fi36++ {
            fmt.Println(v16)
            for fi37 := 0; fi37 < 12; fi37++ {
                var v38 float32 = v16
                var v39 bool = false
                v4 = 0.1
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = fi36
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
