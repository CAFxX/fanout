package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int32; F1 int8; F2 int64 }
func genFn0(p0 uint64, p1 bool, p2 int) int16 {
    v6 := []int{p2, -802}
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint32 = 42
        var v9 float64 = ((0.0 * (-2.25 - 0.0)) / (3.14159 * -2.25))
        if p1 {
            var v10 int32 = 963
            for fi11 := 0; fi11 < 6; fi11++ {
                if p1 {
                    v6 = append(v6, -1)
                    var v12 int64 = 9223372036854775807
                    var v13 int32 = v10
                    _ = v12
                    _ = v13
                }
                _ = len(v6)
                p1 = (65535 > 42)
                var v14 float64 = -2.25
                v15 := GenS1{F0: "\n\t", F1: 105}
                _ = fi11
                _ = v14
                _ = v15
            }
            p0 = (p0 & (p0 & p0))
            _ = v10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v16 := GenS1{F0: "foo bar", F1: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v6
    return 668
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 128
    var v3 uint16 = 465
    var v4 bool = ("a" >= ("" + (("a" + "a") + ("foo bar" + "a"))))
    v5 := []uint32{40, 2}
    v17 := GenS1{F0: "a", F1: v2}
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 string = ""
        var v20 uint8 = 226
        v2 = 1
        _ = fi18
        _ = v19
        _ = v20
    }
    v3 = v3
    switch 2147483647 {
    case -253:
        if true {
            if v4 {
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 uint32 = 65535
                    v22 += v22
                    _ = fi21
                    _ = v22
                }
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v23 := GenS1{F0: "foo bar", F1: (^(v2))}
            _ = v23
        }
    case -9223372036854775808:
        v2 += (v2 + 42)
    case 32768:
        for fi24 := 0; fi24 < 6; fi24++ {
            fmt.Println(v2)
            _ = fi24
        }
    default:
        var v25 int32 = 2147483647
        _ = v25
    }
    if ((-9223372036854775808 >= -9223372036854775808) || false) {
        for fi26 := 0; fi26 < 6; fi26++ {
            v3 -= (v3 ^ v3)
            _ = fi26
        }
        delete(v1, (("hello" + "a") + ""))
        for fi27 := 0; fi27 < 14; fi27++ {
            if v4 {
                v28 := GenS1{F0: "a", F1: 0}
                if (-915 >= -2147483648) {
                    var v29 float32 = ((0.1 / 1.0) + (1e10 / float32((float64((0.1 + 1e10)) * ((0.1 - 3.14159) - (((1.5 * (((-2.25 / 1.0) - 1e10) / (0.1 - 0.0))) - 1.5) / (1.5 - float64(0.0))))))))
                    switch "foo bar" {
                    case "\n\t":
                        v30 := GenS2{F0: uint32(v2)}
                        _ = v30
                    default:
                        var v31 uint16 = 0
                        _ = v31
                    }
                    var v32 uint32 = 66
                    _ = v29
                    _ = v32
                }
                if v4 {
                    fmt.Println(fi27)
                    var v33 int64 = -121
                    _ = v33
                } else {
                    v2 = v2
                    for fi34 := 0; fi34 < 17; fi34++ {
                        v2 = (v2 ^ (v2 >> ((332) & 63)))
                        _ = fi34
                    }
                }
                _ = v28
            } else {
                v4 = true
                for fi35 := 0; fi35 < 12; fi35++ {
                    v36 := GenS1{F0: ("a" + ("foo bar" + ("\n\t" + ""))), F1: (^(v2))}
                    _ = fi35
                    _ = v36
                }
            }
            _ = fi27
        }
    }
    switch 9223372036854775807 {
    case 32767:
        v37 := GenS3{F0: (v3 << ((18446744073709551615) & 63)), F1: 4294967295}
        _ = v37
    case 2147483647:
        var v38 uint16 = v3
        _ = v38
    }
    v4 = (!v4)
    v3 -= v3
    for fi39 := 0; fi39 < 10; fi39++ {
        v40 := GenS4{F0: 893, F1: 0, F2: 1}
        _ = fi39
        _ = v40
    }
    if v4 {
        v3 -= 255
        switch ("" + "") {
        case "\n\t":
            var v41 bool = v4
            _ = v41
        default:
            var v42 uint64 = 395
            _ = v42
        }
        for fi43 := 0; fi43 < 14; fi43++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v44 float64 = 0.1
            _ = fi43
            _ = v44
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
