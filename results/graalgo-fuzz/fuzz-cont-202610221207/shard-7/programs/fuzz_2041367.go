package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int }
type GenS2 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int32; F1 uint16; F2 string }
func genFn0() int64 {
    v4 := []int{9223372036854775807, -64}
    var v5 uint = 0
    _ = v4
    _ = v5
    return int64(v5)
}

func main() {
    var v1 uint16 = 2
    var v2 float64 = 0.1
    var v3 uint32 = 172
    var v6 float32 = (1.5 * 3.14159)
    fmt.Println(v2, v2, v3)
    switch 1 {
    case 9223372036854775807:
        v2 = v2
    default:
        v1 = v1
    }
    v7 := GenS1{F0: 9223372036854775807, F1: -313, F2: 127}
    var v8 int8 = -128
    var v9 float32 = (v6 + (float32(((v2 - -2.25) * 3.14159)) - v6))
    if (true && false) {
        if false {
            if false {
                if (1 != 42) {
                    var v10 string = "\n\t"
                    v11 := GenS1{F0: 42, F1: -9223372036854775808, F2: -877}
                    var v12 bool = true
                    _ = v10
                    _ = v11
                    _ = v12
                } else {
                    v13 := GenS2{F0: 65535, F1: 128, F2: v8}
                    var v14 uint64 = 0
                    var v15 int = 2147483647
                    _ = v13
                    _ = v14
                    _ = v15
                }
                fmt.Println(v1, v1, v3)
            } else {
                var v16 string = "\n\t"
                _ = v16
            }
            if true {
                v17 := GenS1{F0: genFn0(), F1: -9223372036854775808, F2: 32768}
                _ = v17
            } else {
                v18 := []int{2147483647}
                if true {
                    v19 := GenS1{F0: -9223372036854775808, F1: 32767, F2: -231}
                    for fi20 := 0; fi20 < 12; fi20++ {
                        switch (("" + "a") + "a") {
                        case "\n\t":
                            var v21 int32 = -1
                            _ = v21
                        case "foo bar":
                            if true {
                                if ((false || (!true)) || ((false || false) || (((v8 >> ((uint(v1)) & 63)) > v8) || true))) {
                                    v18 = append(v18, -916)
                                    v6 += v9
                                } else {
                                    v3 = v3
                                    v2 -= v2
                                    if false {
                                        v22 := GenS3{F0: 1, F1: 1000}
                                        var v23 string = "a"
                                        v24 := GenS4{F0: -832, F1: 0, F2: ("hello" + v23)}
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                    }
                                }
                                v2 -= 0.1
                                v18 = append(v18, -2147483648)
                            }
                        }
                        _ = fi20
                    }
                    _ = v19
                }
                _ = v18
            }
        }
    }
    v3 = v3
    v25 := GenS3{F0: 18446744073709551615, F1: -32768}
    v2 -= 0.0
    var v26 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
