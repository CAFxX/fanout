package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
func genFn0() uint16 {
    if (144 < -1) {
        if false {
            var v4 bool = (true || (18446744073709551615 < 18446744073709551615))
            var v5 int16 = 1000
            v6 := []int{42}
            _ = v4
            _ = v5
            _ = v6
        } else {
            var v7 uint8 = 30
            var v8 int = -544
            var v9 int64 = 1
            _ = v7
            _ = v8
            _ = v9
        }
    }
    v10 := GenS1{F0: 541}
    _ = v10
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 0
    v3 := make(map[uint32]int)
    v11 := GenS1{F0: v2}
    v12 := []int{532, 2147483647}
    var v13 int64 = -9223372036854775808
    fmt.Println(v13, v2)
    for fi14 := 0; fi14 < 17; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 float32 = 0.0
            if true {
                fmt.Println(v13, fi15)
                v2 -= (v2 / ((42) | 1))
                fmt.Println(fi15, v2, v13)
                switch "foo bar" {
                case "foo bar":
                    v17 := GenS1{F0: v2}
                    _ = v17
                case "\n\t":
                    fmt.Println(v16)
                case "a":
                    v18 := GenS1{F0: (v2 / ((v2) | 1))}
                    _ = v18
                default:
                    if (319 < 65535) {
                        v19 := GenS1{F0: 1000}
                        v20 := GenS1{F0: 289}
                        fmt.Println(v13, fi14, fi15)
                        _ = v19
                        _ = v20
                    }
                }
            }
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v13 = v13
    if false {
        if (genFn0() <= genFn0()) {
            v12 = append(v12, 1)
            v13 += (^(v13))
        } else {
            v21 := GenS1{F0: v2}
            _ = v21
        }
        switch "a" {
        case "\n\t":
            var v22 string = "hello"
            _ = v22
        case "a":
            v2 = v2
        }
    }
    v23 := GenS1{F0: v2}
    var v24 int32 = 0
    if true {
        v25 := GenS1{F0: (^(v2))}
        var v26 bool = (38 < 1)
        _ = v25
        _ = v26
    }
    switch 32768 {
    case 438:
        for fi27 := 0; fi27 < 12; fi27++ {
            v24 -= (v24 & (v24 % (((v24 * v24)) | 1)))
            _ = fi27
        }
    case 32767:
        v13 += (v13 + -1)
    default:
        var v28 int16 = (v2 - v2)
        _ = v28
    }
    var v29 bool = ((!true) || (false && (int16(v24) != 223)))
    v30 := GenS2{F0: ("foo bar" + ("hello" + "a"))}
    v29 = (-9223372036854775808 > (v13 * (v13 ^ (v13 + 0))))
    v31 := GenS1{F0: 0}
    v32 := GenS2{F0: ""}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
