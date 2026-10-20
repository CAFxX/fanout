package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint64 }
func genFn0(p0 int16, p1 int16, p2 uint16) uint32 {
    if true {
        fmt.Println(p2)
        p0 -= (p0 << ((18446744073709551615) & 63))
        p1 -= (p1 + -1)
    }
    if (!(float32(0.0) > 1.5)) {
        fmt.Println(p1, p0, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    var v1 float32 = float32((0.0 / 3.14159))
    v2 := []int{2147483647, 463, 9223372036854775807}
    var v3 float64 = 1.5
    v4 := []string{"foo bar", "foo bar", "a", ("hello" + "a")}
    var v5 int32 = 1
    var v6 uint32 = 65535
    var v7 uint32 = 2147483647
    v8 := GenS1{F0: 412, F1: 42, F2: 1}
    switch -768 {
    case 308:
        var v9 int = 478
        _ = v9
    }
    var v10 int8 = -73
    v7 = genFn0(1000, -32768, 799)
    if (!false) {
        if (false || true) {
            if false {
                v3 += (0.0 - v3)
                fmt.Println(v10, v3, v1)
            } else {
                for fi11 := 0; fi11 < 19; fi11++ {
                    if ((false && (false || false)) && true) {
                        v6 = v6
                    } else {
                        v12 := GenS2{F0: 0, F1: -32768, F2: 42}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if (762 != 42) {
                            v13 := GenS2{F0: 115, F1: -50, F2: 0}
                            _ = v13
                        } else {
                            v3 = v3
                        }
                        _ = v12
                    }
                    _ = fi11
                }
                fmt.Println(v6, v10, v7)
            }
        }
        fmt.Println(v1, v7)
        for fi14 := 0; fi14 < 4; fi14++ {
            if (4294967295 == 0) {
                if len(v2) > 0 {
                    v2[0] = fi14
                }
                var v15 float64 = v3
                _ = v15
            }
            if true {
                var v16 uint64 = 4294967295
                var v17 string = "a"
                v10 = (v10 | v10)
                _ = v16
                _ = v17
            }
            _ = fi14
        }
    } else {
        v2 = append(v2, 42)
        if (-929 > 1) {
            fmt.Println(v6, v3, v6)
        }
    }
    switch -1 {
    case -313:
        v2 = append(v2, 9223372036854775807)
    case -282:
        var v18 float32 = (1e10 - float32(1e10))
        _ = v18
    default:
        _ = len(v4)
    }
    v19 := GenS1{F0: v5, F1: -1, F2: 0}
    switch "foo bar" {
    case "":
        var v20 string = ""
        _ = v20
    case "\n\t":
        v2 = append(v2, 128)
    }
    v21 := GenS1{F0: -2147483648, F1: (v10 & v10), F2: 233}
    var v22 int = 939
    if ((!(!(((1 >= 1) && (false && (!(v22 == v22)))) && true))) && false) {
        for fi23 := 0; fi23 < 6; fi23++ {
            v1 -= 1e10
            var v24 int16 = -32768
            _ = fi23
            _ = v24
        }
        v10 = 127
    } else {
        fmt.Println(v7)
    }
    v1 = ((0.1 / -2.25) / 1e10)
    if (42 == uint16(v7)) {
        var v25 uint16 = 511
        if true {
            var v26 string = ("a" + "")
            _ = v26
        } else {
            if len(v2) > 0 {
                v2[0] = v22
            }
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                fmt.Println(v10, v5, v7)
            }
            var v27 float64 = 0.0
            _ = v27
        }
        _ = v25
    } else {
        var v28 uint64 = 1
        v4 = append(v4, ("" + "\n\t"))
        switch (-(v22)) {
        case 42:
            v28 += v28
        default:
            fmt.Println(v10)
        }
        _ = v28
    }
    v2 = append(v2, 32768)
    v5 += v5
    var v29 float32 = 3.14159
    fmt.Println(v1, v22, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
