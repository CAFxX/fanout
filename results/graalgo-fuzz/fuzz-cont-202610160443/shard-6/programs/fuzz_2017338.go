package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int16 }
type GenS2 struct { F0 int }

func main() {
    v1 := []string{""}
    var v2 uint8 = 41
    v3 := make(map[string]int)
    v4 := GenS1{F0: uint16(v2), F1: 42, F2: -32768}
    v2 += (v2 & 128)
    var v5 uint16 = 2
    var v6 float32 = -2.25
    if (!(18446744073709551615 < 304)) {
        switch 312 {
        case 9223372036854775807:
            v7 := GenS1{F0: 32767, F1: 1, F2: 649}
            _ = v7
        }
    } else {
        if false {
            switch 127 {
            case -1:
                var v8 float64 = (0.0 / 0.1)
                _ = v8
            case 1:
                if (-2.25 != float32(1.5)) {
                    v2 -= v2
                    if (false && (0.1 >= (1.5 / -2.25))) {
                        v6 = 0.1
                        fmt.Println(v5)
                    } else {
                        fmt.Println(v6)
                    }
                } else {
                    v2 -= (v2 << ((2) & 63))
                }
            }
        }
        var v9 int64 = 9223372036854775807
        v6 = v6
        _ = v9
    }
    if false {
        v1 = append(v1, ("hello" + "hello"))
        v10 := GenS1{F0: (v5 % (((v5 ^ v5)) | 1)), F1: -1, F2: 513}
        fmt.Println(v5, v6, v2)
        _ = v10
    }
    fmt.Println(v6, v6)
    if (false && (!(1 >= v2))) {
        v11 := GenS2{F0: -1}
        _ = v11
    }
    if false {
        var v12 int = 267
        switch 2147483647 {
        case 128:
            v1 = append(v1, ("foo bar" + ""))
        case -919:
            if false {
                if len(v1) > 0 {
                    v1[0] = ("a" + ("" + ""))
                }
                var v13 uint16 = (^(v5))
                fmt.Println(v6, v12, v13)
                _ = v13
            } else {
                var v14 bool = false
                _ = v14
            }
        default:
            var v15 string = ("\n\t" + ("\n\t" + (("hello" + "hello") + "hello")))
            _ = v15
        }
        _ = v12
    } else {
        v6 = 0.0
    }
    var v16 uint = 42
    var v17 string = "foo bar"
    v18 := v3[(("" + "hello") + ((v17 + (("\n\t" + "foo bar") + v17)) + ((v17 + v17) + v17)))]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
