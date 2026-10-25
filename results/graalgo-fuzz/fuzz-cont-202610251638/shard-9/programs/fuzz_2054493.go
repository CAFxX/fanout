package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint16 }

func main() {
    v1 := []int{-834}
    v2 := make(map[uint32]bool)
    var v3 int8 = 127
    v4 := make(map[uint32]int)
    v5 := []uint32{2}
    var v6 uint64 = 1
    switch ("a" + "") {
    case "":
        switch 127 {
        case -842:
            for fi7 := 0; fi7 < 17; fi7++ {
                if ((1 <= 32767) && false) {
                    var v8 int32 = 2147483647
                    _ = v8
                } else {
                    for fi9 := 0; fi9 < 12; fi9++ {
                        if false {
                            var v10 string = ("" + "")
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v11 uint = 977
                            _ = v10
                            _ = v11
                        }
                        v3 = v3
                        _ = fi9
                    }
                }
                _ = fi7
            }
        case 1:
            v6 = v6
        case 0:
            v3 -= (v3 & v3)
        }
    default:
        fmt.Println(v3)
    }
    var v12 int8 = v3
    v13 := GenS1{F0: 0, F1: 128, F2: 163}
    _ = len(v5)
    v2[284] = false
    var v14 float32 = (1.5 + (((float32(-2.25) * 3.14159) * (0.0 * 1e10)) - 0.1))
    v15 := v2[0]
    if (((v15 && (v14 <= ((v14 + -2.25) * 0.0))) || false) && (42 == 505)) {
        if v15 {
            v15 = (v15 && v15)
        }
        fmt.Println(v6, v6, v6)
    }
    v16 := GenS2{F0: "a"}
    var v17 int = -9223372036854775808
    v12 -= 31
    if true {
        v15 = (v14 > -2.25)
        for fi18 := 0; fi18 < 20; fi18++ {
            switch "\n\t" {
            case "\n\t":
                v19 := GenS3{F0: 255, F1: 65535}
                _ = v19
            case "a":
                switch "" {
                case "\n\t":
                    v12 -= v12
                case "a":
                    v20 := v2[539]
                    _ = v20
                case "hello":
                    if true {
                        var v21 uint16 = 161
                        v22 := v4[134]
                        _ = v21
                        _ = v22
                    }
                default:
                    v3 -= v3
                }
            }
            _ = fi18
        }
        v23 := GenS3{F0: 4294967295, F1: 1}
        _ = v23
    }
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
