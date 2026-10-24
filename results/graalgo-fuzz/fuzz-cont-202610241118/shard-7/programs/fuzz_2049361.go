package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint16; F1 bool; F2 int64 }
type GenS3 struct { F0 int }
func genFn0(p0 uint8, p1 uint32) int16 {
    v4 := []int{9223372036854775807, 127, -223}
    v4 = append(v4, 0)
    var v5 int8 = 42
    p1 += 472
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1000
}

func main() {
    var v1 int = -569
    var v2 bool = (!(!false))
    v3 := make(map[uint32]uint32)
    if true {
        switch v1 {
        case -9223372036854775808:
            v6 := []int{(v1 % ((v1) | 1)), (v1 << ((594) & 63)), 818}
            _ = v6
        }
        v7 := []int{(v1 ^ v1), v1, v1, 9223372036854775807}
        if v2 {
            var v8 int64 = -1
            fmt.Println(v2, v8, v1)
            for fi9 := 0; fi9 < 6; fi9++ {
                for fi10 := 0; fi10 < 7; fi10++ {
                    var v11 uint64 = 0
                    v2 = v2
                    var v12 uint16 = 32767
                    v13 := GenS1{F0: fi9, F1: 1000}
                    v11 += (^(v11))
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                v1 = fi9
                _ = fi9
            }
            _ = v8
        } else {
            if false {
                var v14 float32 = (-2.25 - 3.14159)
                var v15 string = "a"
                fmt.Println(v15)
                _ = v14
                _ = v15
            }
            if ((369 < -1) || v2) {
                v16 := v3[2147483648]
                v1 += (v1 / ((v1) | 1))
                _ = v16
            } else {
                if false {
                    fmt.Println(v1, v2)
                    if len(v7) > 0 {
                        v7[0] = (v1 ^ v1)
                    }
                } else {
                    var v17 uint8 = 128
                    _ = v17
                }
                v2 = (!true)
                switch (v1 % ((32768) | 1)) {
                case 498:
                    v2 = (!(32767 <= 509))
                default:
                    v18 := GenS1{F0: (v1 | v1), F1: genFn0(161, 2)}
                    _ = v18
                }
            }
        }
        _ = v7
    }
    v1 = (v1 >> ((2) & 63))
    if true {
        for fi19 := 0; fi19 < 12; fi19++ {
            switch "hello" {
            case "":
                var v20 int = (v1 * (fi19 | (fi19 / (((-(fi19))) | 1))))
                _ = v20
            case "\n\t":
                v21 := []int{v1}
                _ = v21
            default:
                var v22 int16 = 42
                _ = v22
            }
            v23 := GenS1{F0: (v1 | (fi19 + -682)), F1: 867}
            fmt.Println(v2, v2, fi19)
            _ = fi19
            _ = v23
        }
    }
    fmt.Println(v2)
    var v24 uint = 0
    var v25 int8 = 0
    if v2 {
        for fi26 := 0; fi26 < 3; fi26++ {
            if (true && (128 < 2)) {
                v27 := []int{-9223372036854775808, (fi26 << (((^(v24))) & 63)), -1}
                _ = v27
            } else {
                for fi28 := 0; fi28 < 9; fi28++ {
                    v29 := GenS2{F0: 2, F1: (0.1 == 3.14159), F2: 42}
                    v2 = true
                    fmt.Println(v1)
                    _ = fi28
                    _ = v29
                }
                var v30 float64 = (1.5 - float64(3.14159))
                var v31 float64 = float64((-2.25 / 1.0))
                _ = v30
                _ = v31
            }
            for fi32 := 0; fi32 < 18; fi32++ {
                fmt.Println(fi26, v24)
                v33 := GenS2{F0: 980, F1: v2, F2: 124}
                v34 := GenS3{F0: fi26}
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = fi26
        }
    }
    fmt.Println(v24, v24)
    v24 += 978
    for fi35 := 0; fi35 < 1; fi35++ {
        fmt.Println(v2, v1, v25)
        _ = fi35
    }
    if (0 <= (v24 % ((0) | 1))) {
        v36 := []int{(v1 - v1), v1, 4294967295}
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
