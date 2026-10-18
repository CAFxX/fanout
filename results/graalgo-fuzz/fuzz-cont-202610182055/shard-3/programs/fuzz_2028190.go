package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint; F1 uint16; F2 int }
type GenS5 struct { F0 int }
func genFn0(p0 uint16, p1 uint32) uint32 {
    p1 = p1
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{329, fi7, fi7, int(p1)}
        if (true || false) {
            var v9 int16 = 492
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 int16 = (v9 >> ((18446744073709551615) & 63))
                p1 = p1
                var v12 uint8 = 127
                _ = fi10
                _ = v11
                _ = v12
            }
            if true {
                v9 = v9
                if len(v8) > 0 {
                    v8[0] = fi7
                }
            } else {
                v9 = (v9 / ((v9) | 1))
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 int16 = 703
                    _ = fi13
                    _ = v14
                }
                p1 = uint32(p0)
            }
            _ = v9
        }
        p0 = (p0 - p0)
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    if ("a" <= ((("hello" + "hello") + "") + "hello")) {
        v15 := GenS1{F0: -2147483648, F1: p0}
        var v16 int64 = 0
        v17 := GenS1{F0: -2147483648, F1: 42}
        _ = v15
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 14; fi18++ {
            p0 += p0
            _ = fi18
        }
        var v19 int64 = -9223372036854775808
        _ = v19
    }
    v20 := []int{}
    _ = p0
    _ = p1
    _ = v20
    return p1
}
func genFn1(p0 int64, p1 uint8, p2 string) string {
    fmt.Println(p1, p0)
    p1 -= (p1 | (p1 / ((p1) | 1)))
    p2 = (p2 + p2)
    if true {
        if (true && false) {
            fmt.Println(p1, p0, p1)
        } else {
            for fi21 := 0; fi21 < 10; fi21++ {
                v22 := GenS1{F0: 829, F1: 42}
                switch "" {
                case "a":
                    switch (fi21 ^ 846) {
                    case -616:
                        v23 := GenS1{F0: 1, F1: 255}
                        _ = v23
                    case 32768:
                        v24 := GenS1{F0: 0, F1: 15}
                        _ = v24
                    }
                }
                _ = fi21
                _ = v22
            }
            var v25 uint64 = 42
            fmt.Println(p1, p2)
            _ = v25
        }
        v26 := []int{}
        _ = v26
    }
    if false {
        v27 := GenS2{F0: (p1 << ((4294967295) & 63)), F1: false}
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    return ((p2 + (p2 + p2)) + "")
}

func main() {
    var v1 int8 = 42
    var v2 float32 = 0.0
    v3 := make(map[string]string)
    var v4 int32 = -382
    v5 := make(map[int]int)
    var v6 float32 = ((((((v2 / 1.0) - 3.14159) / 0.1) / v2) - 3.14159) + (1.5 / v2))
    for fi28 := 0; fi28 < 4; fi28++ {
        v29 := GenS3{F0: fi28, F1: v4}
        _ = fi28
        _ = v29
    }
    fmt.Println(v6, v2, v1)
    v30 := GenS3{F0: 9223372036854775807, F1: 1}
    var v31 uint32 = 65535
    var v32 int64 = int64(v4)
    var v33 string = "foo bar"
    v34 := GenS4{F0: 278, F1: 274, F2: -9223372036854775808}
    v35 := []int{2147483647, 161, 127, 0}
    fmt.Println(v6, v1, v32)
    v36 := GenS5{F0: 2147483647}
    v2 += ((v6 * (v2 * 0.1)) + 0.1)
    if true {
        var v37 float64 = 3.14159
        _ = v37
    } else {
        delete(v5, 9223372036854775807)
        v32 -= int64(v4)
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
