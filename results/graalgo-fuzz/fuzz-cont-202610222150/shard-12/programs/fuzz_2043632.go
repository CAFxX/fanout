package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int16 }
func genFn0() int16 {
    var v5 uint8 = 255
    fmt.Println(v5)
    var v6 uint = 42
    _ = v5
    _ = v6
    return -32768
}
func genFn1(p0 int64, p1 int, p2 uint) bool {
    v7 := GenS1{F0: p1, F1: 970}
    var v8 string = "hello"
    switch ((v8 + ("\n\t" + v8)) + v8) {
    case "":
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (!(-80 > 0))
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 669
    v3 := make(map[string]string)
    var v4 int32 = 1
    if ((!((!(-91 > 127)) && (-891 >= 629))) && ((!(4294967295 == 167)) || (!(501 > 4294967295)))) {
        if (-9223372036854775808 > 0) {
            v4 = (v4 % ((v4) | 1))
            if (2147483648 >= 65535) {
                switch "" {
                case "a":
                    v9 := GenS2{F0: (^(v2)), F1: -461}
                    _ = v9
                default:
                    v2 = (v2 << ((4294967295) & 63))
                }
                v2 += v2
                if true {
                    v2 -= 1
                    v10 := GenS1{F0: 0, F1: 42}
                    _ = v10
                }
            } else {
                v4 = v4
            }
        }
    }
    v11 := GenS1{F0: -2147483648, F1: 4294967295}
    fmt.Println(v2, v2)
    v12 := GenS1{F0: 9223372036854775807, F1: 2}
    if (199 == 4294967295) {
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 bool = (!true)
            for fi15 := 0; fi15 < 3; fi15++ {
                v16 := []int{42, fi13, fi15, fi15}
                var v17 uint64 = 365
                v16 = append(v16, (fi15 | (fi13 >> ((990) & 63))))
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi13
            _ = v14
        }
        var v18 uint = 42
        fmt.Println(v4, v2)
        _ = v18
    } else {
        fmt.Println(v2, v2)
    }
    v2 = 1
    v2 = v2
    v2 += (^(v2))
    switch "\n\t" {
    case "":
        if true {
            switch -9223372036854775808 {
            case 32768:
                v2 = (^(v2))
            default:
                var v19 float64 = ((((-2.25 / (1e10 + 0.1)) * -2.25) / 1.0) / float64(3.14159))
                _ = v19
            }
            var v20 float64 = 3.14159
            _ = v20
        }
    case "foo bar":
        v4 -= -1
    case "\n\t":
        v4 += (v4 * (v4 & v4))
    default:
        v21 := []int{-9223372036854775808, 73}
        _ = v21
    }
    switch 1 {
    case 42:
        for fi22 := 0; fi22 < 20; fi22++ {
            v23 := GenS2{F0: v2, F1: 743}
            _ = fi22
            _ = v23
        }
    case -732:
        fmt.Println(v4, v2)
    case 614:
        v1[42] = "foo bar"
    default:
        v4 -= (v4 << ((255) & 63))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
