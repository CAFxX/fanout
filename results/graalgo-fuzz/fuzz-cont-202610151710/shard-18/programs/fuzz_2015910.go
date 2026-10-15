package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 bool }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int64 = -1
    var v2 float64 = 0.1
    var v3 int16 = 1
    var v4 uint8 = 255
    v1 -= v1
    v1 = (v1 << ((506) & 63))
    var v5 float64 = v2
    if (false && true) {
        v6 := GenS1{F0: (v3 & v3), F1: 0, F2: true}
        if true {
            v7 := GenS2{F0: 15}
            _ = v7
        }
        _ = v6
    } else {
        if false {
            if true {
                if (0 <= 917) {
                    v8 := GenS1{F0: v3, F1: 0, F2: (2 >= 32767)}
                    v4 += (v4 ^ v4)
                    v9 := GenS1{F0: v3, F1: 0, F2: (!(0 <= 0))}
                    _ = v8
                    _ = v9
                }
                var v10 int8 = 37
                _ = v10
            } else {
                v11 := GenS1{F0: (v3 % ((42) | 1)), F1: 166, F2: false}
                fmt.Println(v3)
                v2 += 1e10
                _ = v11
            }
        }
    }
    v2 += (v2 + (v5 - (v2 - float64((3.14159 - 0.1)))))
    fmt.Println(v5)
    if true {
        v12 := []int{9223372036854775807, 148, -2147483648, -1}
        fmt.Println(v5, v4, v1)
        _ = v12
    } else {
        fmt.Println(v2, v3, v1)
        v3 = -32768
        if false {
            fmt.Println(v3, v4, v5)
            v4 -= v4
        } else {
            v13 := []int{-9223372036854775808, -9223372036854775808, -9223372036854775808}
            v5 = (3.14159 / 1e10)
            _ = v13
        }
    }
    var v14 uint64 = 18446744073709551615
    v15 := []int{42}
    var v16 int32 = 0
    if (false && ((((((("foo bar" + ("a" + (("a" + ("a" + ("foo bar" + "a"))) + "hello"))) + ("hello" + (("" + "foo bar") + "hello"))) + "foo bar") + "") + "foo bar") + "a") > "a")) {
        if len(v15) > 0 {
            v15[0] = -9223372036854775808
        }
        for fi17 := 0; fi17 < 15; fi17++ {
            if false {
                switch int(v1) {
                case 2147483647:
                    v3 = v3
                case 61:
                    fmt.Println(v5, fi17, v4)
                default:
                    var v18 uint32 = 0
                    _ = v18
                }
                if (int16(v1) < 1) {
                    if ((255 >= 2) || false) {
                        v19 := GenS1{F0: (v3 | (v3 + v3)), F1: (fi17 / (((fi17 & (^(fi17)))) | 1)), F2: ((v16 - v16) != 2147483647)}
                        v3 += v3
                        _ = v19
                    }
                }
                v3 = (v3 & v3)
            } else {
                v14 = (v14 & v14)
                if (!false) {
                    v20 := GenS1{F0: (v3 & 1), F1: -9223372036854775808, F2: (!false)}
                    _ = v20
                }
                var v21 uint8 = 89
                _ = v21
            }
            var v22 int8 = 9
            for fi23 := 0; fi23 < 20; fi23++ {
                v2 -= v2
                _ = fi23
            }
            _ = fi17
            _ = v22
        }
    }
    v5 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
