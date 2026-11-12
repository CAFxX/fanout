package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint }

func main() {
    var v1 int = 9223372036854775807
    var v2 uint64 = 742
    var v3 uint16 = 255
    if (0.1 <= float64(1.5)) {
        v1 += (^(v1))
        fmt.Println(v3)
        v4 := []int{v1, v1}
        _ = v4
    } else {
        for fi5 := 0; fi5 < 13; fi5++ {
            var v6 string = "\n\t"
            _ = fi5
            _ = v6
        }
        v1 = (v1 * (v1 & (v1 - v1)))
        for fi7 := 0; fi7 < 15; fi7++ {
            fmt.Println(v2)
            var v8 string = (("hello" + (("foo bar" + "") + (("" + "\n\t") + "hello"))) + "")
            if (!(42 != 18446744073709551615)) {
                v9 := []int{(fi7 / (((fi7 % ((v1) | 1))) | 1)), 0}
                var v10 uint16 = (v3 | v3)
                for fi11 := 0; fi11 < 3; fi11++ {
                    switch 42 {
                    case 0:
                        if false {
                            switch (fi7 % ((-602) | 1)) {
                            case 4294967295:
                                var v12 float64 = 1.5
                                _ = v12
                            }
                            var v13 float32 = ((1e10 * 0.0) - 3.14159)
                            _ = v13
                        }
                    case 722:
                        switch ("a" + ((("foo bar" + (v8 + "hello")) + v8) + v8)) {
                        case "":
                            if false {
                                var v14 int = (^(fi11))
                                v14 = (fi7 >> ((12) & 63))
                                _ = v14
                            }
                        case "hello":
                            var v15 uint32 = 209
                            _ = v15
                        }
                    case 42:
                        if (true && false) {
                            fmt.Println(v10)
                        }
                    }
                    if ((1.5 / -2.25) == 0.1) {
                        var v16 int64 = 575
                        v17 := GenS1{F0: 1, F1: 962}
                        if (false || true) {
                            v10 += v3
                        } else {
                            if false {
                                v10 -= (v3 ^ 255)
                                v16 = -960
                                v1 += (^(v1))
                                if len(v9) > 0 {
                                    v9[0] = fi7
                                }
                                _ = len(v9)
                                fmt.Println(v3, v8)
                                var v18 uint32 = 65535
                                fmt.Println(fi7, v2)
                                _ = v18
                            } else {
                                v2 += 2
                                if (v2 >= 42) {
                                    var v19 uint64 = (^(v2))
                                    _ = len(v9)
                                    _ = v19
                                }
                                var v20 int32 = int32(v1)
                                _ = v20
                            }
                        }
                        v21 := GenS1{F0: -123, F1: 32767}
                        _ = v16
                        _ = v17
                        _ = v21
                    }
                    _ = fi11
                }
                _ = v10
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    }
    var v22 int32 = 0
    v23 := []int{(v1 + v1), v1}
    var v24 int32 = 0
    fmt.Println(v1, v1)
    var v25 float64 = 3.14159
    v24 = -1
    v22 += (v22 * v22)
    var v26 uint64 = (^(v2))
    v27 := GenS2{F0: 959, F1: 42}
    v24 = v24
    fmt.Println(v24, v1, v26)
    var v28 int32 = (-(v24))
    v1 = v1
    v1 -= -96
    var v29 uint8 = uint8(v2)
    var v30 uint32 = 816
    v2 = (v26 * 511)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
