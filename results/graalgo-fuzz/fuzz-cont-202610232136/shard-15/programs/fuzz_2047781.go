package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint16; F1 int; F2 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 uint; F2 uint }

func main() {
    v1 := []int{}
    var v2 bool = (false || false)
    v3 := []uint32{2, 4294967295}
    v4 := []int{9223372036854775807, 243}
    v5 := []int{127, -1, 990, -577}
    v2 = (!("" != ""))
    var v6 float64 = 0.1
    switch ("hello" + "foo bar") {
    case "foo bar":
        v7 := GenS1{F0: 832, F1: (!true), F2: -491}
        _ = v7
    case "a":
        var v8 uint8 = 210
        _ = v8
    }
    v6 -= 0.1
    var v9 int64 = 42
    v2 = v2
    v4 = append(v4, 1)
    var v10 int16 = 854
    var v11 bool = (432 < 2147483648)
    v12 := GenS1{F0: 966, F1: false, F2: 0}
    for fi13 := 0; fi13 < 15; fi13++ {
        if v2 {
            v14 := GenS2{F0: 42, F1: 65535, F2: 2147483648}
            if v2 {
                var v15 float64 = 1.5
                v16 := GenS3{F0: (("hello" + "") + "foo bar"), F1: 128}
                if (-1 <= int16(fi13)) {
                    if len(v1) > 0 {
                        v1[0] = (fi13 - (fi13 / (((fi13 | (-(fi13)))) | 1)))
                    }
                }
                _ = v15
                _ = v16
            } else {
                for fi17 := 0; fi17 < 20; fi17++ {
                    v18 := GenS4{F0: 255, F1: -1, F2: 1}
                    v11 = (v2 || (127 == 55))
                    switch "\n\t" {
                    case "":
                        switch "hello" {
                        case "foo bar":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case "a":
                            v19 := GenS2{F0: 187, F1: 1, F2: 289}
                            _ = v19
                        case "\n\t":
                            v20 := GenS4{F0: 2, F1: fi17, F2: 358}
                            _ = v20
                        }
                    default:
                        v2 = (v10 >= 42)
                    }
                    v2 = ((0.0 == ((3.14159 / 1.5) * 0.1)) && ((42 > -1) && true))
                    _ = fi17
                    _ = v18
                }
                v9 = -1
            }
            var v21 uint64 = 42
            _ = v14
            _ = v21
        }
        _ = fi13
    }
    if v2 {
        var v22 int64 = v9
        _ = v22
    }
    v4 = append(v4, -154)
    v23 := GenS5{F0: (v9 | v9)}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if ("a" > "hello") {
        fmt.Println(v11, v11, v9)
        var v24 uint32 = 65535
        v25 := GenS2{F0: 1, F1: 873, F2: (v24 * (v24 + 2))}
        _ = v24
        _ = v25
    }
    v26 := GenS6{F0: 200, F1: 18446744073709551615, F2: 732}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
