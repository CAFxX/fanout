package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint64, p1 int8) uint16 {
    v6 := GenS1{F0: p0, F1: 436}
    fmt.Println(p0, p1)
    if false {
        fmt.Println(p1)
        p1 = 100
    } else {
        if ((4294967295 <= 561) || true) {
            fmt.Println(p0, p1)
            if (1 != 0) {
                p1 = 1
                if (4294967295 >= -9223372036854775808) {
                    v7 := GenS1{F0: 74, F1: p0}
                    v8 := []int{127}
                    p0 = p0
                    _ = v7
                    _ = v8
                } else {
                    for fi9 := 0; fi9 < 16; fi9++ {
                        var v10 uint = 42
                        _ = fi9
                        _ = v10
                    }
                    p0 -= (p0 >> ((65535) & 63))
                }
            }
            v11 := []int{128, -436}
            _ = v11
        }
        if ("hello" == "foo bar") {
            p1 -= (^(p1))
        }
        fmt.Println(p0, p1)
    }
    switch -728 {
    case 127:
        if false {
            fmt.Println(p0)
            var v12 float32 = 0.0
            for fi13 := 0; fi13 < 11; fi13++ {
                if false {
                    fmt.Println(fi13)
                    v12 = v12
                    var v14 float32 = ((v12 * 0.0) + v12)
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 4; fi15++ {
                        var v16 uint = 2
                        var v17 uint16 = 65535
                        var v18 uint16 = 691
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    v12 = -2.25
                }
                p0 -= p0
                switch ("\n\t" + ("" + "")) {
                case "hello":
                    var v19 int32 = 699
                    _ = v19
                case "a":
                    p0 = (p0 | p0)
                default:
                    for fi20 := 0; fi20 < 8; fi20++ {
                        v21 := GenS1{F0: (p0 >> ((239) & 63)), F1: (p0 * p0)}
                        _ = fi20
                        _ = v21
                    }
                }
                _ = fi13
            }
            _ = v12
        }
    }
    v22 := []int{}
    _ = p0
    _ = p1
    _ = v22
    _ = v6
    return 1
}

func main() {
    var v1 int8 = 100
    var v2 uint8 = 142
    var v3 uint64 = 609
    v4 := make(map[int]uint32)
    var v5 int = 32767
    v1 += -1
    v23 := []int{128, v5, v5, v5}
    for fi24 := 0; fi24 < 19; fi24++ {
        fmt.Println(v3)
        switch "\n\t" {
        case "a":
            v3 += v3
        case "hello":
            fmt.Println(fi24, v5, v2)
        case "foo bar":
            var v25 uint32 = 0
            _ = v25
        default:
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
        }
        _ = fi24
    }
    v5 += (v5 << ((uint(v2)) & 63))
    delete(v4, v5)
    v26 := GenS1{F0: (v3 + (v3 ^ v3)), F1: 984}
    var v27 float32 = ((1.5 / 0.1) / 1.0)
    v28 := GenS1{F0: uint64(v2), F1: v3}
    var v29 uint8 = 1
    if true {
        var v30 uint8 = v2
        if false {
            if (42 <= v1) {
                var v31 float64 = (1e10 - (float64(v27) - 0.1))
                switch (v5 % ((v5) | 1)) {
                case -365:
                    v2 += v2
                }
                _ = v31
            } else {
                v27 -= v27
                for fi32 := 0; fi32 < 10; fi32++ {
                    for fi33 := 0; fi33 < 11; fi33++ {
                        fmt.Println(v1, v27, v2)
                        v30 = v30
                        if (-2.25 <= (0.1 / 1e10)) {
                            var v34 int8 = -128
                            _ = v34
                        }
                        _ = fi33
                    }
                    var v35 int32 = 1
                    fmt.Println(v29)
                    _ = fi32
                    _ = v35
                }
                _ = len(v4)
            }
            var v36 uint8 = (v2 << ((978) & 63))
            _ = v36
        }
        _ = v30
    }
    var v37 uint8 = v2
    var v38 uint8 = v37
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
