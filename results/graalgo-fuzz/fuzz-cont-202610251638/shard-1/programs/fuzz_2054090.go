package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint8; F2 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 int }

func main() {
    var v1 int32 = 2147483647
    var v2 float64 = -2.25
    v3 := make(map[int]uint32)
    var v4 uint8 = 2
    v5 := make(map[string]bool)
    v4 -= 42
    v5[("" + "\n\t")] = (263 == 2147483647)
    _ = len(v3)
    v6 := GenS1{F0: false}
    v7 := GenS1{F0: ((1.5 + (3.14159 / 0.1)) > 3.14159)}
    v2 = (v2 / 0.1)
    var v8 bool = false
    v9 := GenS2{F0: -1, F1: 26, F2: 32768}
    fmt.Println(v2)
    var v10 uint = 2
    if true {
        v11 := GenS3{F0: 42}
        v12 := GenS1{F0: (!(42 > 1))}
        var v13 uint32 = 4294967295
        _ = v11
        _ = v12
        _ = v13
    }
    v14 := GenS4{F0: 1}
    _ = len(v3)
    switch "a" {
    case "a":
        if (v2 < v2) {
            if v8 {
                var v15 uint64 = 42
                v16 := GenS5{F0: -206, F1: -54}
                for fi17 := 0; fi17 < 8; fi17++ {
                    var v18 int64 = 561
                    var v19 uint = uint(v15)
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = v15
                _ = v16
            } else {
                fmt.Println(v2, v8, v8)
                switch ("" + ("" + (("" + ("hello" + "")) + "hello"))) {
                case "hello":
                    _ = len(v5)
                default:
                    v4 = (v4 % ((v4) | 1))
                }
                v10 -= 879
            }
        } else {
            var v20 int = 9223372036854775807
            _ = v20
        }
    case "\n\t":
        v21 := []int{-910, -9223372036854775808, 127}
        _ = v21
    default:
        if false {
            var v22 uint64 = 18446744073709551615
            if false {
                if (-2147483648 <= -75) {
                    v23 := GenS3{F0: 755}
                    var v24 int8 = -128
                    fmt.Println(v10, v10, v24)
                    _ = v23
                    _ = v24
                } else {
                    var v25 int64 = 440
                    _ = v25
                }
                var v26 int32 = v1
                switch 42 {
                case 4294967295:
                    var v27 string = "hello"
                    _ = v27
                case -1:
                    fmt.Println(v8, v26, v4)
                case -722:
                    v28 := []int{}
                    _ = v28
                default:
                    if (-9223372036854775808 < 127) {
                        for fi29 := 0; fi29 < 12; fi29++ {
                            if (65535 >= 65535) {
                                v30 := []int{(fi29 ^ (^(fi29))), fi29, (fi29 ^ int(v26)), fi29}
                                _ = v30
                            }
                            for fi31 := 0; fi31 < 14; fi31++ {
                                var v32 int32 = (-(v26))
                                v22 = v22
                                v33 := []int{}
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                            _ = fi29
                        }
                    } else {
                        switch 1 {
                        case 9223372036854775807:
                            v8 = v8
                        case 42:
                            v34 := []int{-1, 671, 32767, 127}
                            _ = v34
                        }
                        v8 = ((v8 || v8) && (32767 >= 32767))
                    }
                }
                _ = v26
            }
            v35 := []int{1, 2147483647}
            _ = v22
            _ = v35
        } else {
            v36 := []int{-2147483648, 2147483647}
            v8 = false
            switch -10 {
            case 74:
                _ = len(v36)
            }
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
