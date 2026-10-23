package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint8; F2 int32 }
type GenS3 struct { F0 int; F1 uint16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    var v4 bool = true
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]string)
    v7 := GenS1{F0: 12, F1: 128}
    fmt.Println(v4)
    switch 42 {
    case 692:
        v8 := GenS1{F0: -128, F1: 250}
        _ = v8
    }
    fmt.Println(v4)
    fmt.Println(v4)
    switch "hello" {
    case "":
        v4 = (!((!(v4 || v4)) || (583 > 18446744073709551615)))
    case "hello":
        if (!(v4 || (((1 == 0) && v4) || (255 > 255)))) {
            v9 := v2[4294967295]
            var v10 float64 = (1e10 * 0.0)
            switch 1 {
            case -828:
                fmt.Println(v9)
            case 42:
                v9 = true
            case 128:
                v1[-465] = 65535
            }
            _ = v10
            _ = v9
        } else {
            fmt.Println(v4)
            var v11 uint64 = 1
            switch -9223372036854775808 {
            case 476:
                switch -1 {
                case 4294967295:
                    if ((255 == 0) && v4) {
                        var v12 int8 = -56
                        for fi13 := 0; fi13 < 13; fi13++ {
                            v11 += v11
                            v14 := GenS2{F0: (v12 * v12), F1: 255, F2: 1}
                            v4 = v4
                            _ = fi13
                            _ = v14
                        }
                        if ((v4 && false) || (false || (false && false))) {
                            var v15 int32 = 743
                            v15 -= 42
                            _ = v15
                        } else {
                            var v16 uint = 167
                            for fi17 := 0; fi17 < 16; fi17++ {
                                v11 = (v11 ^ v11)
                                _ = fi17
                            }
                            v18 := []int{0, 174}
                            _ = v16
                            _ = v18
                        }
                        _ = v12
                    }
                case 32767:
                    v19 := []int{989, 559, 32768, -1}
                    _ = v19
                }
            case 32767:
                if true {
                    var v20 int = 32767
                    fmt.Println(v4, v20, v4)
                    _ = v20
                } else {
                    var v21 uint32 = 2
                    v4 = (!false)
                    _ = v21
                }
            case 128:
                v22 := v5[2147483647]
                _ = v22
            default:
                var v23 int8 = -1
                _ = v23
            }
            _ = v11
        }
    case "\n\t":
        var v24 int = 42
        _ = v24
    }
    var v25 float32 = -2.25
    if v4 {
        v25 = -2.25
        v25 -= v25
        v4 = (0.1 < (float64((float32((1e10 - ((0.0 - 1.5) - 3.14159))) - 0.1)) + 1.5))
    } else {
        if v4 {
            var v26 float64 = (float64(float32(1.5)) / 1e10)
            v4 = ("" >= (("hello" + ("hello" + "a")) + ("a" + ("hello" + "foo bar"))))
            v25 = 3.14159
            _ = v26
        } else {
            for fi27 := 0; fi27 < 10; fi27++ {
                for fi28 := 0; fi28 < 1; fi28++ {
                    var v29 int32 = 1
                    v30 := GenS3{F0: -2147483648, F1: 1}
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = fi27
            }
            var v31 int8 = -128
            _ = v31
        }
    }
    var v32 float32 = (v25 + v25)
    var v33 uint = 42
    var v34 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
