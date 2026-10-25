package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint; F2 int16 }
func genFn0() uint64 {
    var v5 int32 = -2147483648
    switch "a" {
    case "hello":
        v6 := GenS1{F0: "\n\t", F1: 1}
        _ = v6
    case "\n\t":
        v7 := []int{-520, 9223372036854775807, 128, -2147483648}
        _ = v7
    }
    _ = v5
    return 4294967295
}

func main() {
    var v1 int16 = 42
    v2 := []uint32{4294967295, 2147483647, 2147483647}
    var v3 uint = 1
    v4 := make(map[int]int)
    for fi8 := 0; fi8 < 18; fi8++ {
        if true {
            if true {
                switch 4294967295 {
                case 9223372036854775807:
                    switch "foo bar" {
                    case "":
                        for fi9 := 0; fi9 < 10; fi9++ {
                            if len(v2) > 0 {
                                v2[0] = 2147483648
                            }
                            var v10 int32 = -1
                            _ = fi9
                            _ = v10
                        }
                    }
                case -441:
                    v2 = append(v2, 2147483647)
                case 32768:
                    for fi11 := 0; fi11 < 12; fi11++ {
                        v12 := GenS2{F0: 1, F1: 0}
                        var v13 int8 = 0
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
                v3 += 231
                var v14 uint32 = 939
                _ = v14
            } else {
                var v15 uint64 = 0
                v16 := GenS2{F0: 42, F1: 65535}
                for fi17 := 0; fi17 < 13; fi17++ {
                    var v18 bool = true
                    fmt.Println(v15)
                    _ = fi17
                    _ = v18
                }
                _ = v15
                _ = v16
            }
        }
        _ = fi8
    }
    if (!false) {
        for fi19 := 0; fi19 < 20; fi19++ {
            v1 = v1
            _ = fi19
        }
    }
    if false {
        v20 := GenS2{F0: 0, F1: 1}
        _ = v20
    }
    if true {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        v3 -= (^(v3))
    } else {
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 string = "a"
            v22 = v22
            _ = fi21
            _ = v22
        }
        v23 := GenS1{F0: ("hello" + "hello"), F1: 0}
        _ = v23
    }
    switch 32767 {
    case -746:
        var v24 float32 = 1e10
        _ = v24
    case 9223372036854775807:
        fmt.Println(v1, v3)
    default:
        v3 -= 0
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        var v26 string = "a"
        v27 := GenS3{F0: true}
        fmt.Println(fi25, v3, v26)
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint = v3
    for fi29 := 0; fi29 < 6; fi29++ {
        for fi30 := 0; fi30 < 13; fi30++ {
            fmt.Println(fi30)
            _ = fi30
        }
        _ = fi29
    }
    if (((((0.0 + float64((0.1 - 3.14159))) / 0.1) / (3.14159 + (3.14159 + (-2.25 / 0.1)))) * float64(1e10)) == float64(-2.25)) {
        v31 := GenS4{F0: 4294967295, F1: (v28 % ((v28) | 1)), F2: 1000}
        v32 := GenS3{F0: (false || false)}
        _ = v31
        _ = v32
    } else {
        v33 := GenS2{F0: -13, F1: 32767}
        fmt.Println(v28, v3, v3)
        for fi34 := 0; fi34 < 3; fi34++ {
            if false {
                v2 = append(v2, 1)
            } else {
                v1 += (v1 >> ((v28) & 63))
                v1 -= v1
                v3 -= v3
            }
            _ = fi34
        }
        _ = v33
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v28 -= v28
    v1 = (^(v1))
    var v35 uint64 = 977
    if (127 > 127) {
        v35 = (v35 | genFn0())
        var v36 int64 = 1
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v35
    _ = v4
}
