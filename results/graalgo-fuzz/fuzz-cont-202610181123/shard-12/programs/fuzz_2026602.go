package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
func genFn0() uint {
    var v5 uint64 = 22
    var v6 bool = true
    switch "\n\t" {
    case "hello":
        var v7 int8 = 65
        _ = v7
    default:
        if true {
            if v6 {
                v5 -= (v5 | v5)
                var v8 int32 = 577
                _ = v8
            }
        } else {
            fmt.Println(v6, v5)
            v5 = (^(v5))
        }
    }
    switch -613 {
    case -2147483648:
        var v9 uint16 = 32767
        _ = v9
    case 2147483647:
        v10 := []int{128, -1, -69, 32768}
        _ = v10
    case -773:
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 uint = 16
            fmt.Println(fi11)
            var v13 uint8 = 118
            _ = fi11
            _ = v12
            _ = v13
        }
    default:
        fmt.Println(v5, v6)
    }
    _ = v5
    _ = v6
    return 65535
}

func main() {
    var v1 bool = true
    v2 := make(map[string]bool)
    var v3 int32 = -339
    var v4 float32 = 0.1
    var v14 int32 = v3
    var v15 int = 128
    v16 := GenS1{F0: 42, F1: 181}
    if ((v1 && (863 != 4294967295)) && (v15 > -9223372036854775808)) {
        if (v1 && v1) {
            var v17 uint32 = 2147483648
            _ = v17
        }
        if (388 == 18446744073709551615) {
            fmt.Println(v1)
        }
    } else {
        for fi18 := 0; fi18 < 18; fi18++ {
            v4 = 0.0
            _ = fi18
        }
        for fi19 := 0; fi19 < 15; fi19++ {
            v14 = v14
            _ = fi19
        }
    }
    v15 += v15
    _ = len(v2)
    v20 := []int{v15}
    if v1 {
        if len(v20) > 0 {
            v20[0] = v15
        }
        if ((636 <= 589) || (!(42 < 1))) {
            v14 -= v3
            v4 -= (((3.14159 + ((3.14159 - 3.14159) + v4)) / 1.0) + v4)
            if v1 {
                v3 = v14
                fmt.Println(v4, v1)
                fmt.Println(v15, v3, v1)
            }
        } else {
            fmt.Println(v1)
        }
        if (!((9223372036854775807 >= 9223372036854775807) && (!((1 == 100) && v1)))) {
            v15 += v15
            var v21 int32 = v14
            _ = len(v20)
            _ = v21
        }
    }
    var v22 int8 = 42
    if true {
        v4 = (v4 * 1.5)
    } else {
        for fi23 := 0; fi23 < 6; fi23++ {
            v24 := GenS1{F0: 18446744073709551615, F1: -1}
            if (275 <= 2) {
                var v25 uint8 = 100
                v4 -= (float32(0.0) * v4)
                if ((false && v1) || false) {
                    var v26 uint = 2
                    _ = len(v2)
                    var v27 int8 = v22
                    _ = v26
                    _ = v27
                } else {
                    v20 = append(v20, v15)
                }
                _ = v25
            }
            _ = fi23
            _ = v24
        }
        if (65535 <= 32767) {
            if v1 {
                if ("a" >= (("\n\t" + "\n\t") + ("a" + "a"))) {
                    fmt.Println(v4, v4)
                }
                v28 := GenS1{F0: 42, F1: 0}
                _ = v28
            } else {
                var v29 int32 = v14
                if true {
                    delete(v2, "a")
                }
                switch (v15 << ((65535) & 63)) {
                case -1:
                    v22 = (v22 / (((v22 | v22)) | 1))
                case 409:
                    _ = len(v2)
                }
                _ = v29
            }
        } else {
            v14 -= (v3 + v14)
        }
        var v30 int16 = 33
        _ = v30
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        v15 += (v15 - (^(v15)))
        _ = fi31
    }
    _ = len(v2)
    var v32 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v32
    _ = v4
}
