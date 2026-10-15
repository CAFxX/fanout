package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 string = "foo bar"
    var v3 uint16 = 0
    v4 := []int{-988, 2147483647, 32767, -998}
    v5 := GenS1{F0: 493, F1: 42}
    var v6 int32 = 42
    _ = len(v1)
    if true {
        v7 := GenS2{F0: (^(v3)), F1: (v6 >> ((133) & 63))}
        _ = v7
    } else {
        v2 = "a"
    }
    fmt.Println(v3)
    if false {
        var v8 uint64 = 1
        v3 += v3
        var v9 int = 128
        _ = v8
        _ = v9
    } else {
        v4 = append(v4, 32768)
    }
    if (-52 == 100) {
        v6 += (v6 | (v6 / (((-(v6))) | 1)))
        for fi10 := 0; fi10 < 6; fi10++ {
            fmt.Println(fi10, v3, v6)
            if (true && ((!false) && (true || false))) {
                v11 := GenS3{F0: ((true && (2 < 246)) && true), F1: 2, F2: (v3 >> ((960) & 63))}
                _ = v11
            }
            _ = fi10
        }
    } else {
        v1[798] = 128
        var v12 uint = 255
        var v13 uint32 = 65535
        _ = v12
        _ = v13
    }
    switch "\n\t" {
    case "hello":
        if (42 >= 915) {
            for fi14 := 0; fi14 < 15; fi14++ {
                for fi15 := 0; fi15 < 8; fi15++ {
                    var v16 string = v2
                    _ = fi15
                    _ = v16
                }
                _ = fi14
            }
            var v17 uint8 = 0
            _ = v17
        }
    case "":
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 uint = 407
            _ = fi18
            _ = v19
        }
    default:
        v20 := GenS1{F0: 42, F1: -32768}
        _ = v20
    }
    switch 32768 {
    case -693:
        if (true && (!false)) {
            switch 0 {
            case 128:
                var v21 int32 = (v6 / ((1) | 1))
                _ = v21
            }
        } else {
            v2 = v2
            fmt.Println(v6, v2, v2)
            fmt.Println(v3, v2, v3)
        }
    case 531:
        var v22 int = 32767
        _ = v22
    }
    var v23 int32 = (v6 + v6)
    if true {
        v2 = v2
        v23 = (v23 * v23)
        var v24 float64 = 3.14159
        _ = v24
    } else {
        switch 32768 {
        case 701:
            var v25 int64 = -9223372036854775808
            _ = v25
        case 128:
            var v26 float32 = -2.25
            _ = v26
        case -1:
            v27 := GenS1{F0: 42, F1: -1}
            _ = v27
        }
    }
    if len(v4) > 0 {
        v4[0] = -356
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
