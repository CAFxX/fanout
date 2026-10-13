package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int32) int16 {
    p0 = p0
    p0 += -58
    p0 += p0
    p0 -= 0
    var v7 bool = (42 < -1)
    _ = p0
    _ = v7
    return 1000
}

func main() {
    var v1 float64 = (float64(0.1) * 1.5)
    var v2 uint64 = 2
    v3 := []string{}
    v4 := make(map[string]uint32)
    var v5 float64 = (v1 * (v1 * (3.14159 * (v1 * v1))))
    var v6 uint = 65535
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            if false {
                var v10 int32 = 0
                v11 := GenS1{F0: v2}
                var v12 int = fi8
                _ = v10
                _ = v11
                _ = v12
            } else {
                var v13 int8 = 127
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v13
            }
            _ = fi9
        }
        for fi14 := 0; fi14 < 4; fi14++ {
            _ = len(v3)
            if false {
                v15 := GenS1{F0: (^(v2))}
                v2 = (v2 % ((v2) | 1))
                fmt.Println(v2)
                _ = v15
            } else {
                if false {
                    switch (fi14 & fi14) {
                    case 127:
                        _ = len(v4)
                    default:
                        var v16 int32 = 42
                        _ = v16
                    }
                }
                v17 := GenS1{F0: (^(v2))}
                _ = v17
            }
            _ = fi14
        }
        _ = fi8
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        if (42 < v6) {
            switch (("hello" + "hello") + (("" + "foo bar") + (("hello" + "a") + ""))) {
            case "":
                fmt.Println(v5, v2)
            case "foo bar":
                v1 += v1
            default:
                v2 = (^(v2))
            }
        } else {
            if true {
                var v19 int8 = 127
                _ = v19
            } else {
                v5 += v1
                var v20 float32 = (1e10 * 1e10)
                _ = v20
            }
        }
        v21 := GenS1{F0: v2}
        fmt.Println(v1)
        _ = fi18
        _ = v21
    }
    fmt.Println(v1, v5, v6)
    switch 0 {
    case 115:
        v22 := GenS1{F0: 1}
        _ = v22
    case 733:
        v23 := GenS1{F0: 1}
        _ = v23
    }
    v5 += ((float64((float32(1.5) - 0.1)) / v5) + (v1 * 1.5))
    v6 = (^(v6))
    var v24 int64 = -1
    var v25 int = -9223372036854775808
    fmt.Println(v25)
    v6 -= v6
    v2 = v2
    if (!true) {
        for fi26 := 0; fi26 < 10; fi26++ {
            v4["hello"] = 2147483648
            v27 := GenS1{F0: (v2 * (v2 + (v2 & (^(v2)))))}
            _ = fi26
            _ = v27
        }
        v6 -= (v6 % ((567) | 1))
    } else {
        var v28 int = v25
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
