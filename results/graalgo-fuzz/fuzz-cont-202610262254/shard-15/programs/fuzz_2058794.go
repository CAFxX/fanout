package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 bool = false
    v2 := []string{"hello"}
    v3 := []int{}
    v4 := []int{-2147483648}
    v1 = (1 > 32767)
    var v5 uint64 = 18446744073709551615
    if (153 == 128) {
        v4 = append(v4, 127)
        if (192 <= 42) {
            v1 = v1
            v5 = (^(v5))
            v6 := GenS1{F0: 1, F1: 210, F2: 42}
            _ = v6
        }
    }
    _ = len(v2)
    fmt.Println(v1, v5)
    switch ("" + "foo bar") {
    case "a":
        if v1 {
            if (1.5 >= (0.0 * 0.1)) {
                if (v5 == 42) {
                    fmt.Println(v1, v5)
                    var v7 int8 = -128
                    _ = v7
                }
                switch 9223372036854775807 {
                case 4294967295:
                    var v8 bool = v1
                    _ = v8
                case 127:
                    var v9 bool = false
                    _ = v9
                case 9223372036854775807:
                    if v1 {
                        _ = len(v2)
                    }
                }
            }
        } else {
            v5 += (v5 + (v5 | (v5 | 18446744073709551615)))
        }
    case "hello":
        switch 128 {
        case 1:
            switch 1 {
            case 32767:
                fmt.Println(v1, v1)
            case 0:
                var v10 int16 = 1000
                _ = v10
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
    }
    switch "a" {
    case "hello":
        if false {
            v11 := GenS1{F0: 32767, F1: 0, F2: 1}
            _ = v11
        }
    case "":
        fmt.Println(v5)
    case "foo bar":
        switch -463 {
        case 128:
            _ = len(v2)
        case 9223372036854775807:
            _ = len(v4)
        default:
            v1 = v1
        }
    default:
        if v1 {
            v5 = (v5 >> ((65535) & 63))
            v12 := GenS1{F0: 32767, F1: 20, F2: 916}
            v13 := GenS1{F0: 65535, F1: 1, F2: 0}
            _ = v12
            _ = v13
        } else {
            v14 := GenS2{F0: 277}
            var v15 uint8 = 20
            _ = v14
            _ = v15
        }
    }
    switch "foo bar" {
    case "a":
        v4 = append(v4, -9223372036854775808)
    }
    var v16 float32 = 0.0
    for fi17 := 0; fi17 < 1; fi17++ {
        for fi18 := 0; fi18 < 2; fi18++ {
            v19 := GenS2{F0: 797}
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 13; fi20++ {
            if len(v4) > 0 {
                v4[0] = fi20
            }
            var v21 bool = false
            if (!(0.0 != -2.25)) {
                var v22 string = ("\n\t" + "\n\t")
                switch 127 {
                case -460:
                    _ = len(v3)
                case -1:
                    var v23 float64 = 1.5
                    _ = v23
                }
                _ = v22
            } else {
                fmt.Println(v21, fi17)
            }
            _ = fi20
            _ = v21
        }
        if len(v3) > 0 {
            v3[0] = -81
        }
        _ = fi17
    }
    fmt.Println(v5)
    v5 += (v5 >> ((0) & 63))
    v5 = v5
    v5 += v5
    v2 = append(v2, "")
    v24 := GenS1{F0: 42, F1: 2, F2: 333}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
