package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int16; F1 int; F2 int }

func main() {
    var v1 int = 42
    var v2 int16 = -1
    var v3 string = ("" + ("\n\t" + "a"))
    fmt.Println(v2)
    v4 := []int{(v1 >> ((4294967295) & 63)), v1, (v1 << ((18446744073709551615) & 63))}
    switch ("foo bar" + v3) {
    case "hello":
        v5 := GenS1{F0: v3, F1: 406}
        _ = v5
    default:
        var v6 int = v1
        _ = v6
    }
    if ((v1 << ((506) & 63)) < v1) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        if false {
            v7 := GenS1{F0: v3, F1: 1}
            v8 := GenS1{F0: v3, F1: 581}
            v1 = v1
            _ = v7
            _ = v8
        } else {
            fmt.Println(v1, v2, v3)
        }
        v9 := GenS2{F0: (v1 | -276), F1: v3}
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v3 = v3
        var v11 uint16 = 1
        _ = fi10
        _ = v11
    }
    switch v3 {
    case "hello":
        v4 = append(v4, (v1 & (v1 * v1)))
    case "":
        for fi12 := 0; fi12 < 9; fi12++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if ((-2147483648 == 888) || true) {
                fmt.Println(v1, v2, v1)
                _ = len(v4)
                fmt.Println(v3)
            }
            v3 = (v3 + v3)
            _ = fi12
        }
    default:
        if len(v4) > 0 {
            v4[0] = v1
        }
    }
    v2 -= (v2 | 32767)
    if false {
        var v13 string = "foo bar"
        for fi14 := 0; fi14 < 19; fi14++ {
            if false {
                switch v1 {
                case -9223372036854775808:
                    v4 = append(v4, v1)
                case 153:
                    v1 = v1
                case 560:
                    var v15 uint16 = 2
                    _ = v15
                default:
                    var v16 int32 = 914
                    _ = v16
                }
            }
            var v17 uint8 = 255
            _ = fi14
            _ = v17
        }
        _ = v13
    }
    if true {
        if false {
            for fi18 := 0; fi18 < 8; fi18++ {
                v19 := GenS3{F0: v2, F1: 1, F2: 128}
                if (false && false) {
                    var v20 int = (^(fi18))
                    _ = v20
                }
                _ = fi18
                _ = v19
            }
            var v21 int16 = (v2 + 1000)
            _ = v21
        } else {
            fmt.Println(v1)
        }
        v1 = v1
    }
    v4 = append(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
