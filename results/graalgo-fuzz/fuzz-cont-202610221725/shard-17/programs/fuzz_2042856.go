package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint64; F1 string }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    var v4 int8 = 100
    var v5 uint8 = 127
    v6 := GenS1{F0: 393, F1: 186}
    v7 := GenS1{F0: 498, F1: 42}
    switch 32767 {
    case 4294967295:
        var v8 int64 = -1
        _ = v8
    case 9223372036854775807:
        fmt.Println(v5, v5)
    case 227:
        for fi9 := 0; fi9 < 10; fi9++ {
            delete(v1, "foo bar")
            _ = fi9
        }
    default:
        for fi10 := 0; fi10 < 15; fi10++ {
            switch ("a" + (("foo bar" + "a") + "foo bar")) {
            case "hello":
                for fi11 := 0; fi11 < 8; fi11++ {
                    var v12 uint8 = v5
                    _ = fi11
                    _ = v12
                }
            case "foo bar":
                switch "" {
                case "a":
                    var v13 int64 = -1
                    _ = v13
                case "\n\t":
                    if true {
                        v14 := []int{1, 4294967295, fi10, fi10}
                        for fi15 := 0; fi15 < 7; fi15++ {
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            _ = fi15
                        }
                        for fi16 := 0; fi16 < 17; fi16++ {
                            v17 := GenS1{F0: 2147483648, F1: 65535}
                            var v18 uint32 = 42
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        _ = v14
                    }
                }
            default:
                v19 := []int{(fi10 ^ (fi10 << ((289) & 63))), (fi10 & (-(fi10))), (fi10 & 0)}
                _ = v19
            }
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 float32 = float32(1.5)
                if (-2.25 <= ((1e10 - 0.1) / (1e10 / 0.1))) {
                    fmt.Println(fi20, v5)
                    var v22 bool = false
                    _ = v22
                }
                v5 -= 42
                _ = fi20
                _ = v21
            }
            _ = fi10
        }
    }
    if ((!(32767 >= 32767)) && (!(uint64(v5) <= 4294967295))) {
        var v23 int = 1
        _ = v23
    }
    var v24 string = (("hello" + "\n\t") + "")
    v25 := []int{-331, -758}
    for fi26 := 0; fi26 < 17; fi26++ {
        if ((0.0 / 1.5) >= 0.1) {
            fmt.Println(v4, v5, v5)
            if false {
                v4 -= -128
            }
            v5 += v5
        } else {
            v27 := GenS1{F0: 2, F1: 65535}
            if false {
                var v28 uint8 = v5
                _ = v28
            } else {
                fmt.Println(v5, v5)
                var v29 uint32 = 748
                _ = v29
            }
            v30 := GenS1{F0: 228, F1: 546}
            _ = v27
            _ = v30
        }
        for fi31 := 0; fi31 < 12; fi31++ {
            v5 -= v5
            _ = fi31
        }
        v4 = (v4 ^ v4)
        _ = fi26
    }
    switch "hello" {
    case "\n\t":
        v32 := GenS2{F0: 18446744073709551615, F1: v24}
        _ = v32
    default:
        v5 = (^(v5))
    }
    v33 := GenS1{F0: 669, F1: 42}
    v34 := GenS1{F0: 2147483647, F1: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
