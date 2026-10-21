package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint16 }
func genFn0(p0 int16, p1 int, p2 uint) int {
    for fi5 := 0; fi5 < 15; fi5++ {
        p2 = 0
        var v6 float32 = 1.5
        switch p1 {
        case 127:
            p0 += (p0 + 0)
        case 4294967295:
            v7 := []int{int(p0), (fi5 % (((fi5 | p1)) | 1))}
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    v8 := GenS1{F0: 251, F1: "foo bar", F2: 255}
    v9 := []int{(p1 ^ (p1 * p1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return p1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 9223372036854775807
    var v3 int16 = 32767
    var v4 int = v2
    v3 += 0
    v10 := []int{(v4 / ((genFn0((^(v3)), 42, 1)) | 1)), 9223372036854775807, v2}
    switch ("foo bar" + (("" + "\n\t") + (("a" + "foo bar") + "hello"))) {
    case "":
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 uint64 = 1
            var v13 uint = 255
            _ = fi11
            _ = v12
            _ = v13
        }
    default:
        if (!(((498 != 18446744073709551615) && (!true)) && false)) {
            v4 = (v2 * 173)
            v14 := GenS1{F0: 78, F1: "a", F2: 1}
            _ = v14
        }
    }
    var v15 int32 = 2147483647
    v3 = (v3 * (v3 | (^(v3))))
    fmt.Println(v2)
    v15 -= v15
    var v16 float32 = -2.25
    v4 = -2147483648
    var v17 int32 = v15
    var v18 float32 = (1e10 * 1.5)
    if (!false) {
        if (4 >= 2) {
            fmt.Println(v16)
            if (false || true) {
                for fi19 := 0; fi19 < 17; fi19++ {
                    for fi20 := 0; fi20 < 5; fi20++ {
                        fmt.Println(v17, v15, v18)
                        v21 := GenS1{F0: 42, F1: "\n\t", F2: 32767}
                        _ = fi20
                        _ = v21
                    }
                    var v22 int64 = 0
                    _ = fi19
                    _ = v22
                }
                v18 = (-2.25 / (0.0 - (v18 - v18)))
                if false {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    v23 := GenS1{F0: 128, F1: "hello", F2: 0}
                    _ = v23
                } else {
                    fmt.Println(v17)
                }
            }
            var v24 uint8 = 85
            _ = v24
        }
        var v25 int32 = v17
        _ = v25
    }
    var v26 int8 = 5
    switch (v4 | (v4 * 127)) {
    case -997:
        fmt.Println(v2, v3)
    default:
        var v27 int32 = (v15 % (((v15 * (v15 ^ (^(v17))))) | 1))
        _ = v27
    }
    if false {
        v28 := v1[170]
        _ = v28
    }
    var v29 float32 = v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
}
