package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }

func main() {
    var v1 uint32 = 0
    var v2 int16 = 0
    var v3 float64 = (3.14159 + 1e10)
    var v4 int32 = -768
    fmt.Println(v3, v4)
    if ((v1 - (v1 & v1)) <= v1) {
        switch "foo bar" {
        case "a":
            var v5 int = 4294967295
            _ = v5
        }
        var v6 int16 = 611
        v7 := GenS1{F0: 278, F1: -702}
        _ = v6
        _ = v7
    }
    if (-9223372036854775808 <= -1) {
        switch -2147483648 {
        case -2147483648:
            var v8 uint = 402
            _ = v8
        case 9223372036854775807:
            if (v4 >= v4) {
                var v9 uint = 1
                _ = v9
            }
        case 127:
            fmt.Println(v2, v2)
        default:
            v10 := GenS1{F0: 18446744073709551615, F1: 340}
            _ = v10
        }
        var v11 uint8 = 0
        for fi12 := 0; fi12 < 13; fi12++ {
            fmt.Println(v3, v3, v2)
            v2 = -971
            _ = fi12
        }
        _ = v11
    } else {
        var v13 uint16 = 42
        v14 := GenS1{F0: 449, F1: v4}
        _ = v13
        _ = v14
    }
    v3 = 1e10
    fmt.Println(v1, v1)
    var v15 bool = true
    var v16 uint8 = 42
    for fi17 := 0; fi17 < 14; fi17++ {
        if (255 == 312) {
            if ((2 < 1) && (false && (false || false))) {
                v3 = v3
            }
        } else {
            var v18 int = fi17
            var v19 uint8 = v16
            switch v18 {
            case -22:
                for fi20 := 0; fi20 < 13; fi20++ {
                    var v21 int64 = 42
                    _ = fi20
                    _ = v21
                }
            case 32767:
                fmt.Println(v1, v18)
            default:
                v16 = 1
            }
            _ = v18
            _ = v19
        }
        v22 := []int{597}
        for fi23 := 0; fi23 < 2; fi23++ {
            switch (fi17 + fi17) {
            case 42:
                var v24 int = 832
                _ = v24
            }
            if len(v22) > 0 {
                v22[0] = fi17
            }
            if len(v22) > 0 {
                v22[0] = fi23
            }
            _ = fi23
        }
        _ = fi17
        _ = v22
    }
    var v25 uint = 2
    v3 = v3
    var v26 int = 0
    switch v26 {
    case 0:
        switch "hello" {
        case "foo bar":
            fmt.Println(v3, v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
