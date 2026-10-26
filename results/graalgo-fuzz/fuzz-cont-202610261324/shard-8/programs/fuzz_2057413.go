package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int8; F1 int8 }

func main() {
    v1 := []uint32{112, 4294967295, 42, 947}
    var v2 uint = 0
    v3 := []int{9223372036854775807, 0}
    if true {
        var v4 string = "foo bar"
        v5 := GenS1{F0: 42}
        _ = v4
        _ = v5
    } else {
        for fi6 := 0; fi6 < 13; fi6++ {
            v2 += (v2 << ((v2) & 63))
            _ = fi6
        }
        for fi7 := 0; fi7 < 4; fi7++ {
            v1 = append(v1, 1)
            v3 = append(v3, (fi7 + (fi7 << ((v2) & 63))))
            _ = fi7
        }
    }
    var v8 uint8 = 0
    var v9 int8 = 0
    v10 := GenS1{F0: 1}
    fmt.Println(v8, v9)
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 int64 = 321
            v8 += (v8 / (((v8 % ((v8) | 1))) | 1))
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 8; fi14++ {
            fmt.Println(v8)
            switch "a" {
            case "a":
                v9 = v9
            default:
                var v15 int = (-(fi14))
                _ = v15
            }
            _ = fi14
        }
        if (558 <= 976) {
            switch ("" + ("foo bar" + "")) {
            case "foo bar":
                v9 -= 100
            case "a":
                var v16 uint8 = (^(v8))
                _ = v16
            case "":
                var v17 string = "hello"
                _ = v17
            default:
                v2 = v2
            }
        }
        _ = fi11
    }
    v2 -= 1
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 uint8 = 22
        var v20 string = ""
        _ = fi18
        _ = v19
        _ = v20
    }
    switch 127 {
    case 604:
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 int32 = 2147483647
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi21
            _ = v22
        }
    case 4294967295:
        v8 = (v8 * v8)
    case -740:
        var v23 int = -2147483648
        _ = v23
    }
    if len(v3) > 0 {
        v3[0] = -133
    }
    v24 := GenS2{F0: -27, F1: (v9 + v9)}
    var v25 float32 = 1.5
    for fi26 := 0; fi26 < 9; fi26++ {
        v2 += 65535
        if (v2 < (^(v2))) {
            var v27 float64 = -2.25
            if (false || false) {
                v28 := GenS3{F0: (v2 | uint(v8)), F1: uint32(v2)}
                _ = v28
            } else {
                for fi29 := 0; fi29 < 11; fi29++ {
                    v30 := GenS2{F0: (-(v9)), F1: 0}
                    v2 = (^(v2))
                    v9 = v9
                    _ = fi29
                    _ = v30
                }
                var v31 int8 = 0
                var v32 uint = (^(v2))
                _ = v31
                _ = v32
            }
            fmt.Println(v25, v2)
            _ = v27
        } else {
            v33 := GenS4{F0: (v9 ^ (v9 - v9)), F1: v9}
            fmt.Println(v8, v2, v25)
            var v34 uint16 = 1
            _ = v33
            _ = v34
        }
        var v35 int16 = 72
        _ = fi26
        _ = v35
    }
    v25 = v25
    v9 = 100
    var v36 int32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v8
    _ = v9
}
