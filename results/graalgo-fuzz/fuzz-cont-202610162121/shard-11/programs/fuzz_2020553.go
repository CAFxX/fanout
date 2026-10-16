package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint }
func genFn0() uint16 {
    var v4 int = -508
    v4 -= (v4 + (v4 << ((42) & 63)))
    _ = v4
    return 32767
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 42
    var v3 uint16 = 2
    if (4294967295 != 2147483647) {
        v1[-690] = 1
    }
    v5 := GenS1{F0: 4294967295, F1: (3.14159 < 0.0)}
    var v6 uint16 = v3
    if (-9223372036854775808 == -9223372036854775808) {
        var v7 uint16 = (v6 * (v3 << ((643) & 63)))
        _ = v7
    } else {
        var v8 bool = false
        delete(v1, 32768)
        _ = v8
    }
    if true {
        switch (("" + "a") + "") {
        case "a":
            v3 += (v6 << ((0) & 63))
        case "foo bar":
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 uint32 = 2
                _ = fi9
                _ = v10
            }
        case "\n\t":
            var v11 int8 = 0
            _ = v11
        default:
            switch "\n\t" {
            case "":
                if ((32767 > 127) && true) {
                    switch "a" {
                    case "":
                        fmt.Println(v6, v2, v2)
                    case "foo bar":
                        v3 -= 65535
                    case "hello":
                        var v12 float64 = 0.1
                        _ = v12
                    default:
                        delete(v1, 487)
                    }
                    v3 = v3
                }
            }
        }
    } else {
        if ((v2 + (^(v2))) <= 1) {
            switch 666 {
            case 0:
                v6 = 2
            default:
                v3 -= (v6 - v6)
            }
        } else {
            var v13 float64 = float64(3.14159)
            var v14 float32 = (((1e10 / -2.25) * (-2.25 + 0.1)) - (0.1 / 3.14159))
            _ = v13
            _ = v14
        }
        v3 = (v3 & v3)
        var v15 float32 = 1e10
        _ = v15
    }
    var v16 int8 = -128
    v17 := GenS2{F0: 99}
    if (!true) {
        var v18 string = "foo bar"
        _ = v18
    } else {
        var v19 int32 = 42
        _ = v19
    }
    if ((v6 << ((255) & 63)) != (v6 >> ((2) & 63))) {
        if true {
            switch ("" + "") {
            case "a":
                if true {
                    for fi20 := 0; fi20 < 20; fi20++ {
                        v16 += 42
                        _ = fi20
                    }
                    var v21 int64 = 1
                    var v22 uint = 126
                    _ = v21
                    _ = v22
                }
            }
        } else {
            v1[4294967295] = 1000
            v23 := GenS3{F0: 839}
            _ = v23
        }
        v24 := GenS4{F0: 255}
        _ = v24
    }
    fmt.Println(v3, v3)
    v25 := GenS2{F0: 1}
    v1[4294967295] = 778
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v5
    _ = v6
}
