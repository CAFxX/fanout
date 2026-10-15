package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint, p1 uint32) int16 {
    p0 = 12
    v7 := []int{32767, 128, -2147483648}
    _ = p0
    _ = p1
    _ = v7
    return 285
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "a"
    var v3 int32 = -2147483648
    v4 := []uint32{}
    v5 := []uint32{2147483647, 2147483647, 2147483647, 4294967295}
    var v6 int32 = (v3 - v3)
    var v8 int32 = 0
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: true}
        if (((-2147483648 == (v6 | v3)) || true) || true) {
            var v11 float64 = 3.14159
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    v12 := GenS1{F0: true}
    v13 := v1[316]
    fmt.Println(v6, v13, v13)
    for fi14 := 0; fi14 < 16; fi14++ {
        switch v13 {
        case "hello":
            var v15 uint32 = 925
            _ = v15
        case "a":
            for fi16 := 0; fi16 < 1; fi16++ {
                _ = len(v1)
                _ = fi16
            }
        case "\n\t":
            for fi17 := 0; fi17 < 7; fi17++ {
                switch -2147483648 {
                case 2147483647:
                    var v18 string = v2
                    _ = v18
                default:
                    var v19 bool = (int64(fi14) < 9223372036854775807)
                    v4 = append(v4, 493)
                    _ = v19
                }
                if (v2 < "a") {
                    v6 = (v6 / ((v6) | 1))
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v4 = append(v4, 686)
                }
                _ = fi17
            }
        }
        v20 := GenS1{F0: false}
        _ = fi14
        _ = v20
    }
    v8 = (v3 % ((v8) | 1))
    if false {
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 float32 = ((0.1 - 0.1) * 1.5)
            switch fi21 {
            case 9223372036854775807:
                var v23 string = ("" + ((v2 + (v2 + v2)) + v2))
                _ = v23
            }
            var v24 string = (((v2 + v2) + v13) + "")
            _ = fi21
            _ = v22
            _ = v24
        }
        v25 := GenS2{F0: 4294967295}
        _ = v25
    }
    if true {
        v26 := GenS3{F0: 255, F1: 0}
        var v27 int8 = 0
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        switch 32767 {
        case 0:
            if ((^(v6)) < int32(fi28)) {
                for fi29 := 0; fi29 < 15; fi29++ {
                    var v30 int = fi28
                    _ = fi29
                    _ = v30
                }
                v13 = (("hello" + v13) + v13)
                v3 = 2147483647
            } else {
                var v31 bool = true
                _ = v31
            }
        case 127:
            fmt.Println(v13, v13)
        default:
            var v32 uint8 = 50
            _ = v32
        }
        _ = len(v5)
        for fi33 := 0; fi33 < 15; fi33++ {
            var v34 int64 = -150
            _ = fi33
            _ = v34
        }
        _ = fi28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
