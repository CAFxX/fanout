package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 uint32 }
func genFn0() uint16 {
    switch -144 {
    case 32767:
        var v6 bool = false
        _ = v6
    case 0:
        switch 32767 {
        case 4294967295:
            for fi7 := 0; fi7 < 8; fi7++ {
                if (-366 <= -616) {
                    var v8 uint8 = 107
                    for fi9 := 0; fi9 < 20; fi9++ {
                        v10 := []int{4294967295, (fi9 * (fi9 - fi9)), -36, (fi9 & (fi9 / (((fi9 << ((65535) & 63))) | 1)))}
                        if (!false) {
                            fmt.Println(fi9, fi9, fi7)
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                        }
                        _ = fi9
                        _ = v10
                    }
                    _ = v8
                }
                _ = fi7
            }
        case -2147483648:
            v11 := GenS1{F0: 2147483647}
            _ = v11
        case 42:
            for fi12 := 0; fi12 < 9; fi12++ {
                var v13 int16 = -270
                v13 -= v13
                _ = fi12
                _ = v13
            }
        }
    case -86:
        v14 := []int{951, 4294967295, -962}
        _ = v14
    }
    var v15 uint = 18446744073709551615
    var v16 int = -1
    _ = v15
    _ = v16
    return 2
}
func genFn1(p0 int32, p1 uint64) int32 {
    p1 = (p1 % ((p1) | 1))
    switch 32768 {
    case 46:
        var v17 uint8 = 127
        _ = v17
    }
    v18 := []int{32768}
    for fi19 := 0; fi19 < 19; fi19++ {
        v20 := GenS1{F0: 939}
        for fi21 := 0; fi21 < 5; fi21++ {
            fmt.Println(p0, fi19)
            var v22 uint8 = 187
            v23 := GenS2{F0: (-(p0))}
            _ = fi21
            _ = v22
            _ = v23
        }
        var v24 uint = 65535
        _ = fi19
        _ = v20
        _ = v24
    }
    var v25 float32 = float32(0.1)
    _ = p0
    _ = p1
    _ = v18
    _ = v25
    return (^(p0))
}

func main() {
    var v1 int32 = -447
    var v2 int32 = (-(v1))
    var v3 int64 = 9223372036854775807
    v4 := []int{int(v3), 42}
    var v5 int32 = v2
    _ = len(v4)
    v26 := GenS1{F0: (-(v1))}
    switch 32768 {
    case 128:
        switch "\n\t" {
        case "foo bar":
            _ = len(v4)
        case "a":
            if len(v4) > 0 {
                v4[0] = 127
            }
        case "\n\t":
            fmt.Println(v2, v5)
        }
    case -2147483648:
        _ = len(v4)
    default:
        v27 := GenS3{F0: 0}
        _ = v27
    }
    if false {
        for fi28 := 0; fi28 < 1; fi28++ {
            v1 += (v1 ^ v2)
            switch (fi28 - fi28) {
            case 1:
                var v29 uint8 = 1
                _ = v29
            case 128:
                v2 -= (v2 & (-(v1)))
            case 42:
                var v30 uint = 42
                _ = v30
            default:
                v31 := GenS1{F0: 42}
                _ = v31
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi28
        }
    }
    if false {
        v32 := GenS2{F0: 0}
        _ = v32
    } else {
        v1 = (v5 >> ((255) & 63))
        if (65535 != 2147483647) {
            var v33 bool = true
            _ = v33
        } else {
            v34 := GenS3{F0: 794}
            _ = v34
        }
        for fi35 := 0; fi35 < 5; fi35++ {
            if true {
                for fi36 := 0; fi36 < 14; fi36++ {
                    v37 := GenS4{F0: v5, F1: 906}
                    var v38 float32 = 3.14159
                    _ = fi36
                    _ = v37
                    _ = v38
                }
            }
            if (("foo bar" + "\n\t") >= "foo bar") {
                v3 -= 42
                for fi39 := 0; fi39 < 3; fi39++ {
                    var v40 int = 32768
                    _ = fi39
                    _ = v40
                }
                var v41 float64 = float64((float32((0.0 * (-2.25 / -2.25))) + 3.14159))
                _ = v41
            } else {
                fmt.Println(v3, v1)
            }
            _ = fi35
        }
    }
    for fi42 := 0; fi42 < 20; fi42++ {
        v43 := GenS3{F0: 4294967295}
        _ = fi42
        _ = v43
    }
    v44 := GenS5{F0: 533}
    v4 = append(v4, -914)
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
