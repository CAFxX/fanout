package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 int32 }
func genFn0(p0 uint) uint8 {
    for fi6 := 0; fi6 < 5; fi6++ {
        if ((0.1 >= (3.14159 + 0.1)) && false) {
            var v7 float64 = (1.5 + ((3.14159 / 1.0) - float64(0.0)))
            _ = v7
        }
        fmt.Println(fi6, p0)
        p0 = 18446744073709551615
        _ = fi6
    }
    if (-785 < -455) {
        if (!false) {
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := GenS1{F0: -32768, F1: 117, F2: 2147483647}
                fmt.Println(p0)
                _ = fi8
                _ = v9
            }
            p0 -= p0
        }
    } else {
        v10 := GenS1{F0: 1000, F1: 1, F2: 1}
        _ = v10
    }
    _ = p0
    return 180
}
func genFn1() int64 {
    var v11 string = ""
    if false {
        for fi12 := 0; fi12 < 6; fi12++ {
            v11 = v11
            _ = fi12
        }
        v13 := GenS1{F0: -1, F1: 1, F2: 2147483648}
        _ = v13
    } else {
        if (353 == 255) {
            for fi14 := 0; fi14 < 13; fi14++ {
                for fi15 := 0; fi15 < 18; fi15++ {
                    var v16 int16 = 1000
                    switch "\n\t" {
                    case "foo bar":
                        v17 := []int{}
                        _ = v17
                    case "\n\t":
                        fmt.Println(fi15, fi14, v11)
                    case "hello":
                        v18 := GenS2{F0: (-(fi15))}
                        v19 := GenS2{F0: (fi15 % ((fi14) | 1))}
                        var v20 uint = 65535
                        v11 = v11
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    v21 := GenS2{F0: (^(fi15))}
                    _ = fi15
                    _ = v16
                    _ = v21
                }
                _ = fi14
            }
            var v22 int16 = 667
            var v23 uint32 = 380
            _ = v22
            _ = v23
        } else {
            fmt.Println(v11)
        }
    }
    v24 := GenS1{F0: 32767, F1: -9223372036854775808, F2: 2}
    var v25 float32 = 1e10
    _ = v11
    _ = v24
    _ = v25
    return -9223372036854775808
}

func main() {
    v1 := []string{((("hello" + "\n\t") + "") + ("" + ("" + "\n\t"))), "\n\t", "", "hello"}
    var v2 float32 = 0.0
    var v3 int = -2147483648
    var v4 bool = true
    var v5 int = (v3 - 1)
    v26 := GenS2{F0: (-(v3))}
    var v27 int8 = -1
    v28 := GenS3{F0: 778}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v29 := GenS4{F0: 2, F1: 0}
    fmt.Println(v5)
    v30 := GenS5{F0: 4294967295, F1: 806}
    if (!(91 <= 1)) {
        _ = len(v1)
        v27 -= -29
    }
    fmt.Println(v5)
    var v31 uint64 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v32 uint8 = 42
    for fi33 := 0; fi33 < 9; fi33++ {
        if v4 {
            v34 := GenS6{F0: 2147483647}
            _ = v34
        }
        for fi35 := 0; fi35 < 14; fi35++ {
            var v36 int64 = -1
            _ = fi35
            _ = v36
        }
        _ = fi33
    }
    v4 = ((!(v27 > v27)) || v4)
    for fi37 := 0; fi37 < 3; fi37++ {
        fmt.Println(v5)
        v2 = v2
        var v38 uint16 = 2
        _ = fi37
        _ = v38
    }
    v27 += (v27 << ((65535) & 63))
    switch (v5 << ((2) & 63)) {
    case 127:
        var v39 uint8 = v32
        _ = v39
    case 0:
        switch v5 {
        case 0:
            var v40 int64 = 0
            _ = v40
        default:
            if len(v1) > 0 {
                v1[0] = "a"
            }
        }
    case 42:
        var v41 uint32 = 4294967295
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
