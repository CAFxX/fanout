package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint16 }

func main() {
    var v1 int8 = -128
    var v2 bool = (false || (!true))
    v3 := []string{"foo bar", ("a" + "foo bar")}
    var v4 int64 = 9223372036854775807
    var v5 uint64 = 84
    v4 -= (v4 | (v4 << ((255) & 63)))
    var v6 string = ""
    if false {
        v2 = (int32(v1) == int32(v5))
        var v7 uint32 = 65535
        for fi8 := 0; fi8 < 13; fi8++ {
            v4 += (v4 / (((v4 & v4)) | 1))
            for fi9 := 0; fi9 < 20; fi9++ {
                _ = len(v3)
                _ = len(v3)
                fmt.Println(v7)
                _ = fi9
            }
            _ = fi8
        }
        _ = v7
    } else {
        switch ("\n\t" + v6) {
        case "":
            if len(v3) > 0 {
                v3[0] = "hello"
            }
        case "\n\t":
            fmt.Println(v2)
        case "a":
            for fi10 := 0; fi10 < 20; fi10++ {
                if v2 {
                    v4 = 9223372036854775807
                    var v11 int16 = 32767
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 3; fi12++ {
                        var v13 uint8 = 19
                        v1 = 1
                        var v14 int64 = -908
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    var v15 float64 = (0.1 * (float64((3.14159 - 0.1)) / 1.0))
                    v1 += v1
                    _ = v15
                }
                _ = len(v3)
                fmt.Println(v4, v4, v2)
                _ = fi10
            }
        default:
            var v16 uint16 = 65535
            _ = v16
        }
        switch -2147483648 {
        case 2147483647:
            fmt.Println(v2)
        }
    }
    var v17 uint = 1
    var v18 string = v6
    var v19 float32 = (1e10 / (-2.25 + 1e10))
    var v20 int16 = 32767
    if false {
        fmt.Println(v5)
        fmt.Println(v1, v1)
    }
    v21 := GenS1{F0: (v17 * 0), F1: 0, F2: 678}
    fmt.Println(v1, v18)
    v18 = v18
    switch 32768 {
    case -407:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 0:
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    case -2147483648:
        _ = len(v3)
    default:
        v4 = v4
    }
    v2 = false
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
