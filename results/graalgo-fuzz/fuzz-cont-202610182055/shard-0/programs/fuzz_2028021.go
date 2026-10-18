package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    var v3 int = 9223372036854775807
    var v4 int = v3
    v5 := make(map[string]uint32)
    var v6 uint8 = 0
    switch 9223372036854775807 {
    case 0:
        v7 := GenS1{F0: v6, F1: 65535, F2: 65535}
        _ = v7
    default:
        var v8 int64 = 42
        _ = v8
    }
    v6 = v6
    switch (v4 + 86) {
    case 1:
        if (-83 < -773) {
            if (!(4294967295 >= 1)) {
                v3 += v3
                if (18446744073709551615 >= 338) {
                    for fi9 := 0; fi9 < 15; fi9++ {
                        v4 -= v4
                        for fi10 := 0; fi10 < 19; fi10++ {
                            fmt.Println(v6, v6, v3)
                            _ = fi10
                        }
                        _ = fi9
                    }
                }
                if (!(!true)) {
                    if true {
                        v11 := []int{(v4 / ((42) | 1)), 974, (v4 * v4), -970}
                        fmt.Println(v6, v3)
                        _ = v11
                    }
                } else {
                    v12 := GenS1{F0: (v6 >> ((721) & 63)), F1: 18446744073709551615, F2: 1}
                    var v13 float32 = 0.0
                    _ = v12
                    _ = v13
                }
            }
            var v14 string = "hello"
            _ = v14
        } else {
            v4 = (v3 / (((v3 << ((65535) & 63))) | 1))
            v4 += (v4 & v3)
            v3 += v3
        }
    default:
        for fi15 := 0; fi15 < 11; fi15++ {
            if (((fi15 | v3) > 973) && false) {
                v16 := GenS1{F0: 255, F1: 42, F2: 65535}
                _ = v16
            } else {
                if false {
                    var v17 int32 = -2147483648
                    _ = v17
                }
            }
            _ = fi15
        }
    }
    v3 = v3
    v6 = v6
    v18 := []int{v4, (v4 + 138), (v4 | (v4 - -2147483648)), 2147483647}
    fmt.Println(v6, v3)
    v19 := GenS1{F0: 143, F1: 2, F2: 65535}
    var v20 uint = 549
    v3 += v4
    fmt.Println(v20, v4, v3)
    switch ("hello" + "a") {
    case "a":
        fmt.Println(v3, v6, v6)
    default:
        var v21 uint32 = 645
        _ = v21
    }
    var v22 bool = (4294967295 >= 4294967295)
    if (880 == 42) {
        if (v22 || (-9223372036854775808 != 1)) {
            var v23 string = (("hello" + "hello") + ("a" + ""))
            v3 = -890
            _ = v23
        }
    }
    var v24 int16 = -32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
