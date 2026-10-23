package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 uint32 }
func genFn0() uint32 {
    var v4 bool = (42 > 1)
    for fi5 := 0; fi5 < 14; fi5++ {
        if false {
            for fi6 := 0; fi6 < 9; fi6++ {
                var v7 int8 = -128
                _ = fi6
                _ = v7
            }
            v4 = true
            v8 := []int{}
            _ = v8
        }
        _ = fi5
    }
    _ = v4
    return 65535
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    var v3 uint = 319
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            v3 = v3
            _ = fi10
        }
        v11 := []int{int(v3)}
        _ = fi9
        _ = v11
    }
    var v12 uint32 = 2147483647
    var v13 uint64 = 795
    v12 += (v12 << ((300) & 63))
    var v14 float32 = -2.25
    for fi15 := 0; fi15 < 8; fi15++ {
        if (!(v14 < 0.1)) {
            var v16 bool = (127 <= -109)
            _ = v16
        } else {
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 uint = 255
                v19 := GenS1{F0: 389}
                _ = fi17
                _ = v18
                _ = v19
            }
            for fi20 := 0; fi20 < 1; fi20++ {
                v21 := GenS2{F0: 1, F1: -128}
                v22 := GenS3{F0: ("hello" + "hello")}
                v23 := GenS2{F0: -9223372036854775808, F1: 118}
                v14 -= (v14 / v14)
                var v24 int = (fi20 * -1)
                v3 = (^(v3))
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
                _ = v24
            }
            fmt.Println(v14, v12)
        }
        _ = len(v1)
        _ = fi15
    }
    var v25 float32 = ((1.5 / (v14 - v14)) / v14)
    var v26 uint32 = (v12 - v12)
    switch 2147483647 {
    case -1:
        if true {
            v27 := []int{-358}
            _ = v27
        } else {
            for fi28 := 0; fi28 < 10; fi28++ {
                v29 := GenS1{F0: 1}
                switch "a" {
                case "a":
                    var v30 int16 = 193
                    _ = v30
                case "hello":
                    v14 += (v25 + v14)
                case "\n\t":
                    v31 := GenS4{F0: 18446744073709551615}
                    _ = v31
                }
                _ = fi28
                _ = v29
            }
            var v32 uint64 = 42
            for fi33 := 0; fi33 < 15; fi33++ {
                var v34 uint32 = (v12 / (((v26 >> ((65) & 63))) | 1))
                v35 := GenS5{F0: false, F1: (v12 | (v26 ^ v26))}
                var v36 uint8 = 255
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
            _ = v32
        }
    case 1:
        fmt.Println(v12, v14)
    case 9223372036854775807:
        v37 := v1[(v26 >> ((v3) & 63))]
        _ = v37
    default:
        for fi38 := 0; fi38 < 9; fi38++ {
            var v39 uint8 = 0
            _ = fi38
            _ = v39
        }
    }
    switch -1 {
    case 2147483647:
        var v40 string = "\n\t"
        _ = v40
    default:
        var v41 int = 0
        _ = v41
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        v13 += v13
        _ = fi42
    }
    var v43 uint64 = (^(v13))
    if (v13 > v43) {
        v44 := GenS4{F0: (v13 + (v13 * v43))}
        _ = v44
    } else {
        var v45 uint64 = (v13 >> ((18446744073709551615) & 63))
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v43
}
