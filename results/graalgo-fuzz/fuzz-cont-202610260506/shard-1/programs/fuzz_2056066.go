package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint8; F1 int16; F2 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64; F1 uint16; F2 int8 }
func genFn0(p0 uint, p1 uint32, p2 uint64) uint {
    var v5 float64 = (3.14159 / 1e10)
    switch 0 {
    case -711:
        switch 271 {
        case 32768:
            if (!(!true)) {
                v6 := GenS1{F0: (p1 / (((^(p1))) | 1))}
                if ((-2.25 / v5) < 0.1) {
                    fmt.Println(p2, p0)
                } else {
                    switch -734 {
                    case 127:
                        v7 := GenS2{F0: 42}
                        _ = v7
                    case -2147483648:
                        var v8 bool = ((float64(1.5) * (v5 * 3.14159)) == -2.25)
                        _ = v8
                    }
                    p2 = 0
                    var v9 uint64 = p2
                    _ = v9
                }
                fmt.Println(p0, p0, v5)
                _ = v6
            } else {
                var v10 int8 = 0
                switch ("foo bar" + "\n\t") {
                case "foo bar":
                    switch "\n\t" {
                    case "hello":
                        var v11 uint32 = p1
                        _ = v11
                    case "a":
                        if (0 > 0) {
                            switch ((("a" + "foo bar") + ("a" + "foo bar")) + ("" + "foo bar")) {
                            case "hello":
                                v5 += (-2.25 - (v5 * 3.14159))
                            case "":
                                for fi12 := 0; fi12 < 17; fi12++ {
                                    v13 := []int{(-(fi12))}
                                    _ = fi12
                                    _ = v13
                                }
                            case "\n\t":
                                v14 := []int{-2147483648, -1, 32767, 9223372036854775807}
                                _ = v14
                            }
                        }
                    }
                case "hello":
                    var v15 int32 = 0
                    _ = v15
                default:
                    var v16 int16 = 32767
                    _ = v16
                }
                _ = v10
            }
        case 658:
            p0 = 255
        }
    case 1:
        switch 42 {
        case 127:
            var v17 uint64 = p2
            _ = v17
        default:
            v18 := GenS3{F0: (false || (32767 != 537)), F1: (p2 / (((p2 << (((p0 << (((p0 - 255)) & 63))) & 63))) | 1))}
            _ = v18
        }
    case 9223372036854775807:
        for fi19 := 0; fi19 < 4; fi19++ {
            fmt.Println(p0)
            p1 += p1
            _ = fi19
        }
    default:
        p2 = (p2 - (p2 >> ((2) & 63)))
    }
    var v20 uint32 = p1
    v21 := []int{42}
    if true {
        if false {
            var v22 int32 = 455
            if false {
                fmt.Println(v22)
                p1 += (p1 << ((p0) & 63))
                var v23 uint = (p0 & p0)
                _ = v23
            }
            _ = v22
        } else {
            var v24 float32 = 0.1
            fmt.Println(v24, v20, v24)
            _ = v24
        }
        fmt.Println(v5)
    } else {
        if (-218 >= -512) {
            v5 = v5
            if (false && (!false)) {
                var v25 float32 = 3.14159
                v26 := GenS4{F0: 182, F1: 117, F2: 18446744073709551615}
                p2 = (p2 << ((4294967295) & 63))
                _ = v25
                _ = v26
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    _ = v5
    return (p0 & 65535)
}
func genFn1() bool {
    var v27 int64 = -425
    v27 += v27
    if true {
        v27 = v27
        var v28 uint8 = 255
        _ = v28
    }
    v29 := []int{42}
    _ = v27
    _ = v29
    return ((2147483647 > 104) || (42 > 127))
}

func main() {
    var v1 int8 = 37
    var v2 int8 = v1
    var v3 string = "\n\t"
    var v4 int = 32767
    var v30 uint64 = 4294967295
    var v31 uint64 = v30
    for fi32 := 0; fi32 < 17; fi32++ {
        v1 = (-(v2))
        _ = fi32
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        var v34 uint32 = 1
        _ = fi33
        _ = v34
    }
    var v35 int32 = 2147483647
    for fi36 := 0; fi36 < 19; fi36++ {
        v3 = "foo bar"
        for fi37 := 0; fi37 < 17; fi37++ {
            v38 := GenS5{F0: 561}
            v2 = v1
            _ = fi37
            _ = v38
        }
        _ = fi36
    }
    for fi39 := 0; fi39 < 14; fi39++ {
        if true {
            v4 = fi39
            if (!false) {
                v40 := GenS6{F0: (v35 ^ (-(v35)))}
                v41 := GenS7{F0: -9223372036854775808, F1: 2, F2: -67}
                _ = v40
                _ = v41
            }
        }
        _ = fi39
    }
    if true {
        v1 -= v2
        var v42 uint = 282
        v43 := []int{(v4 + v4), (v4 ^ -438)}
        _ = v42
        _ = v43
    }
    v35 += v35
    if ((float64((((1.5 + (3.14159 + (1.5 + (3.14159 * 3.14159)))) - float32(-2.25)) / 1.0)) + 1.5) != ((-2.25 + float64(-2.25)) - (3.14159 / -2.25))) {
        if false {
            v4 -= v4
        }
    } else {
        var v44 float64 = 3.14159
        v45 := GenS4{F0: 128, F1: -32768, F2: genFn0(616, 0, v31)}
        _ = v44
        _ = v45
    }
    fmt.Println(v1, v1)
    fmt.Println(v2, v3)
    fmt.Println(v35)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
}
