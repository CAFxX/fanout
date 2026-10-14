package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int = 32768
    var v2 int = (v1 - 2147483647)
    var v3 int64 = 42
    var v4 uint64 = 18446744073709551615
    fmt.Println(v4)
    fmt.Println(v3, v2)
    if (true && true) {
        for fi5 := 0; fi5 < 8; fi5++ {
            switch (v2 & 830) {
            case 1:
                switch "a" {
                case "foo bar":
                    var v6 int8 = -98
                    _ = v6
                case "a":
                    v4 += v4
                case "hello":
                    if false {
                        if ((^(v3)) >= 1) {
                            v1 = fi5
                            for fi7 := 0; fi7 < 19; fi7++ {
                                if (("foo bar" + "foo bar") <= "hello") {
                                    var v8 int16 = 42
                                    _ = v8
                                }
                                _ = fi7
                            }
                        }
                        v9 := []int{}
                        for fi10 := 0; fi10 < 2; fi10++ {
                            v9 = append(v9, 127)
                            var v11 string = "\n\t"
                            _ = fi10
                            _ = v11
                        }
                        _ = v9
                    }
                default:
                    v3 -= 1
                }
            }
            fmt.Println(v2, fi5)
            var v12 int64 = 9223372036854775807
            _ = fi5
            _ = v12
        }
        v13 := []int{(v1 & 2147483647), (v2 / (((v1 >> ((65535) & 63))) | 1)), (v1 & v2)}
        var v14 uint = 18446744073709551615
        _ = v13
        _ = v14
    }
    if ("" >= ("\n\t" + "foo bar")) {
        var v15 float32 = 0.0
        if true {
            if (1.5 < 0.0) {
                v4 -= (v4 ^ (v4 * v4))
                var v16 uint32 = 42
                for fi17 := 0; fi17 < 6; fi17++ {
                    v18 := []int{-2147483648, -9223372036854775808}
                    _ = len(v18)
                    for fi19 := 0; fi19 < 9; fi19++ {
                        var v20 int16 = 1
                        fmt.Println(v15)
                        _ = fi19
                        _ = v20
                    }
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
            var v21 float64 = (float64(v15) - (1e10 - 0.0))
            _ = v21
        } else {
            if false {
                var v22 float32 = 1.5
                _ = v22
            } else {
                if (v2 == 84) {
                    v2 += (v1 ^ (v1 - (v2 >> ((1) & 63))))
                    var v23 int16 = -272
                    _ = v23
                } else {
                    var v24 uint16 = 0
                    var v25 int32 = 1
                    var v26 bool = (("hello" + "") < "a")
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
        }
        var v27 uint8 = 1
        _ = v15
        _ = v27
    }
    switch 32767 {
    case 847:
        var v28 bool = false
        _ = v28
    case -373:
        var v29 uint = 65535
        _ = v29
    }
    v30 := GenS1{F0: v2, F1: 2147483648, F2: v1}
    v3 = (v3 ^ (v3 * (v3 & (v3 * (v3 ^ int64(v4))))))
    v31 := []int{(v2 % ((2147483647) | 1)), v1, v2, (v2 << ((2) & 63))}
    fmt.Println(v2, v3)
    _ = len(v31)
    v3 = v3
    v32 := GenS2{F0: 255}
    var v33 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
