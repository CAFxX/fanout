package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint16, p1 uint) int16 {
    p1 -= (p1 >> (((p1 % (((p1 & 1)) | 1))) & 63))
    var v4 string = "a"
    _ = p0
    _ = p1
    _ = v4
    return 615
}

func main() {
    var v1 uint16 = 632
    v2 := make(map[string]int)
    v3 := make(map[int]string)
    if (!false) {
        if true {
            for fi5 := 0; fi5 < 6; fi5++ {
                if false {
                    var v6 uint8 = 42
                    v7 := GenS1{F0: (^(v1)), F1: -377, F2: 0}
                    _ = v6
                    _ = v7
                }
                var v8 int64 = -964
                v9 := GenS2{F0: false}
                _ = fi5
                _ = v8
                _ = v9
            }
            fmt.Println(v1)
            var v10 int = -9223372036854775808
            _ = v10
        }
        if (!false) {
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := GenS3{F0: "", F1: 1000}
                _ = fi11
                _ = v12
            }
            v1 += (v1 / ((v1) | 1))
        } else {
            var v13 int16 = -245
            for fi14 := 0; fi14 < 12; fi14++ {
                switch (fi14 | 494) {
                case -1:
                    v1 = (v1 | 274)
                case 1:
                    v1 += 113
                case 127:
                    _ = len(v3)
                default:
                    v1 -= v1
                }
                _ = fi14
            }
            _ = v13
        }
        fmt.Println(v1)
    }
    var v15 uint16 = (v1 * (^(v1)))
    var v16 uint = 18446744073709551615
    var v17 int32 = -1
    if (42 != 114) {
        v16 = v16
    } else {
        if ((v16 < 823) && false) {
            var v18 bool = (-2.25 > -2.25)
            v19 := []int{2147483647, 32767}
            fmt.Println(v16, v15, v18)
            _ = v18
            _ = v19
        }
        var v20 int32 = v17
        v20 += v17
        _ = v20
    }
    v15 += v1
    var v21 int16 = -481
    if ((false && (!false)) || (-1 < -1)) {
        var v22 int32 = 896
        var v23 uint8 = 1
        _ = v22
        _ = v23
    } else {
        v24 := GenS4{F0: "", F1: 186}
        switch 0 {
        case 548:
            var v25 uint = 18446744073709551615
            _ = v25
        case 9223372036854775807:
            switch "hello" {
            case "a":
                v16 = 255
            case "":
                switch "hello" {
                case "hello":
                    v26 := GenS5{F0: 255, F1: (v21 | v21)}
                    _ = v26
                }
            default:
                if (genFn0((v1 & (v15 / ((v1) | 1))), (^(v16))) != 1000) {
                    _ = len(v2)
                    var v27 uint16 = 174
                    _ = v27
                }
            }
        case -424:
            v1 = v15
        default:
            v28 := []int{}
            _ = v28
        }
        if false {
            if false {
                fmt.Println(v1, v17)
                var v29 int16 = (-(v21))
                _ = v29
            }
        } else {
            var v30 int8 = -1
            v31 := GenS1{F0: 255, F1: (v17 / (((v17 >> ((v16) & 63))) | 1)), F2: v16}
            for fi32 := 0; fi32 < 8; fi32++ {
                v30 += v30
                if false {
                    var v33 int16 = genFn0(32767, v16)
                    fmt.Println(fi32, fi32)
                    _ = v33
                }
                _ = fi32
            }
            _ = v30
            _ = v31
        }
        _ = v24
    }
    v15 = (v1 >> (((v16 >> ((v16) & 63))) & 63))
    fmt.Println(v15)
    switch -179 {
    case 1:
        v21 += genFn0(42, (v16 >> ((uint(v1)) & 63)))
    }
    var v34 uint = v16
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v34
}
