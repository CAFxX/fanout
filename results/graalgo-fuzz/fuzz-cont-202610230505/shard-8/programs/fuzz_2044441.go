package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 bool }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 1e10
    v3 := make(map[uint32]uint32)
    var v4 string = ""
    var v5 uint16 = 32767
    fmt.Println(v4)
    _ = len(v3)
    var v6 uint32 = 2
    var v7 float64 = 0.0
    fmt.Println(v6, v7, v6)
    v5 += v5
    var v8 int8 = 0
    if true {
        var v9 string = "foo bar"
        fmt.Println(v7, v4, v5)
        _ = v9
    } else {
        if ((v5 / (((v5 / ((65535) | 1))) | 1)) < v5) {
            if true {
                fmt.Println(v8, v4)
            } else {
                for fi10 := 0; fi10 < 8; fi10++ {
                    var v11 uint32 = (v6 >> ((2) & 63))
                    v3[2147483647] = v11
                    _ = fi10
                    _ = v11
                }
            }
            switch 649 {
            case -2147483648:
                var v12 int16 = -430
                _ = v12
            case 32768:
                v4 = ""
            case 127:
                if (false && (!false)) {
                    var v13 uint64 = 834
                    _ = v13
                }
            }
        }
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        if false {
            v2 -= 3.14159
            v15 := GenS1{F0: 363, F1: -1, F2: false}
            if false {
                v1["\n\t"] = (v4 + (v4 + v4))
            }
            _ = v15
        }
        if true {
            for fi16 := 0; fi16 < 9; fi16++ {
                v17 := []int{fi14, fi16, fi16}
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                v4 = ("foo bar" + "")
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 13; fi18++ {
                v2 += v2
                v4 = (v4 + v4)
                var v19 uint = 1
                v4 = v4
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 uint8 = 127
                if false {
                    v22 := v3[v6]
                    _ = v22
                }
                var v23 uint32 = v6
                _ = fi20
                _ = v21
                _ = v23
            }
        } else {
            fmt.Println(v7, fi14)
        }
        var v24 string = "hello"
        _ = fi14
        _ = v24
    }
    switch 120 {
    case 2147483647:
        var v25 int16 = 42
        _ = v25
    case 431:
        v3[(v6 + v6)] = v6
    case 42:
        v5 = (v5 | 255)
    default:
        var v26 string = "hello"
        _ = v26
    }
    var v27 float32 = float32(float64((v2 + 1e10)))
    switch v4 {
    case "foo bar":
        var v28 uint8 = 194
        _ = v28
    case "a":
        for fi29 := 0; fi29 < 16; fi29++ {
            var v30 uint16 = v5
            fmt.Println(v6, v6)
            var v31 float32 = v2
            _ = fi29
            _ = v30
            _ = v31
        }
    default:
        var v32 uint64 = 42
        _ = v32
    }
    var v33 uint8 = 42
    v34 := []int{128}
    if (v7 == v7) {
        var v35 int16 = 245
        var v36 int64 = 9223372036854775807
        switch ("a" + (v4 + v4)) {
        case "a":
            v37 := GenS2{F0: 894}
            _ = v37
        }
        _ = v35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
