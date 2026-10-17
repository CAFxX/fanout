package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint32 }
func genFn0(p0 int64) int16 {
    fmt.Println(p0)
    v6 := []int{32767, 4294967295, 1}
    if (4294967295 >= 31) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        p0 = p0
        for fi7 := 0; fi7 < 7; fi7++ {
            switch fi7 {
            case -666:
                if (!(!(!true))) {
                    if (-2.25 < 0.0) {
                        fmt.Println(fi7)
                        fmt.Println(fi7, p0)
                        for fi8 := 0; fi8 < 2; fi8++ {
                            fmt.Println(p0, p0, p0)
                            _ = fi8
                        }
                    } else {
                        v9 := GenS1{F0: 0, F1: "\n\t"}
                        _ = v9
                    }
                }
            case 42:
                var v10 uint8 = 42
                _ = v10
            case -1:
                v11 := GenS1{F0: 0, F1: "\n\t"}
                _ = v11
            default:
                var v12 int8 = 101
                _ = v12
            }
            _ = fi7
        }
    } else {
        for fi13 := 0; fi13 < 14; fi13++ {
            for fi14 := 0; fi14 < 2; fi14++ {
                if (18446744073709551615 > 207) {
                    v15 := GenS1{F0: 0, F1: ("\n\t" + "\n\t")}
                    fmt.Println(fi13, fi14, p0)
                    _ = v15
                } else {
                    p0 = p0
                    v16 := GenS1{F0: 336, F1: ("foo bar" + (("foo bar" + "foo bar") + ""))}
                    _ = v16
                }
                fmt.Println(p0)
                var v17 float32 = 0.1
                _ = fi14
                _ = v17
            }
            _ = fi13
        }
        fmt.Println(p0)
        var v18 int8 = 100
        _ = v18
    }
    _ = p0
    _ = v6
    return -32768
}
func genFn1(p0 int, p1 uint16, p2 uint8) int16 {
    switch "foo bar" {
    case "a":
        if (!false) {
            v19 := GenS2{F0: ("hello" + "")}
            switch (p0 ^ (p0 * (p0 * p0))) {
            case -73:
                v20 := GenS3{F0: 42, F1: -2147483648}
                _ = v20
            default:
                v21 := []int{}
                _ = v21
            }
            _ = v19
        }
    default:
        var v22 int = (p0 - (-(p0)))
        _ = v22
    }
    var v23 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return genFn0(1)
}

func main() {
    var v1 int8 = 1
    var v2 uint = 668
    var v3 uint64 = 987
    var v4 int = -2147483648
    var v5 uint = v2
    if true {
        v3 += 4294967295
        for fi24 := 0; fi24 < 2; fi24++ {
            var v25 float64 = float64(0.0)
            v25 += v25
            _ = fi24
            _ = v25
        }
        v4 -= (^(v4))
    }
    if true {
        v5 += 6
        if false {
            switch (v4 / ((1) | 1)) {
            case 372:
                switch v4 {
                case 9223372036854775807:
                    v3 -= v3
                default:
                    var v26 float64 = 0.0
                    _ = v26
                }
            case 679:
                switch "foo bar" {
                case "":
                    v2 += v2
                default:
                    fmt.Println(v3, v3)
                }
            case 32767:
                for fi27 := 0; fi27 < 15; fi27++ {
                    for fi28 := 0; fi28 < 12; fi28++ {
                        fmt.Println(fi28, v2, v4)
                        _ = fi28
                    }
                    _ = fi27
                }
            }
            v4 = (v4 % (((v4 ^ v4)) | 1))
        } else {
            v29 := []int{v4}
            _ = v29
        }
    } else {
        var v30 float64 = ((0.1 - 0.1) * float64(((1e10 + 0.1) + 1.5)))
        _ = v30
    }
    switch -847 {
    case -9223372036854775808:
        v31 := []int{(v4 + 42)}
        _ = v31
    default:
        var v32 uint32 = 1
        _ = v32
    }
    v1 += (^(v1))
    v33 := GenS4{F0: 2, F1: 42, F2: 42}
    v5 = 532
    var v34 int = (v4 - v4)
    switch "hello" {
    case "\n\t":
        v2 -= v2
    case "hello":
        for fi35 := 0; fi35 < 9; fi35++ {
            if true {
                fmt.Println(v2)
                for fi36 := 0; fi36 < 13; fi36++ {
                    v4 -= (-(v34))
                    if true {
                        fmt.Println(fi35)
                        v37 := []int{(v4 - -2147483648)}
                        v34 -= 0
                        v37 = append(v37, (v4 | 128))
                        _ = v37
                    }
                    v3 += (v3 + v3)
                    _ = fi36
                }
            }
            for fi38 := 0; fi38 < 10; fi38++ {
                if (v3 == 18446744073709551615) {
                    fmt.Println(fi35)
                }
                _ = fi38
            }
            if true {
                fmt.Println(v4)
                fmt.Println(v5)
            }
            _ = fi35
        }
    }
    switch (v34 | (v34 - (v4 ^ v4))) {
    case 933:
        fmt.Println(v34, v34)
    case 262:
        v39 := GenS3{F0: 0, F1: 1}
        _ = v39
    case 0:
        v3 = 0
    }
    var v40 float32 = (1e10 + (0.1 + (1.5 / 1.0)))
    v41 := []int{(v34 | 42), (-(v34))}
    fmt.Println(v3)
    for fi42 := 0; fi42 < 20; fi42++ {
        var v43 int8 = (v1 + int8(v3))
        for fi44 := 0; fi44 < 10; fi44++ {
            if (!true) {
                v1 = 127
                v5 = (v2 << (((^(v2))) & 63))
                fmt.Println(v1)
            } else {
                var v45 bool = true
                _ = v45
            }
            fmt.Println(v5, v34, v2)
            _ = fi44
        }
        _ = fi42
        _ = v43
    }
    v40 = float32((float64(v40) - 1.5))
    for fi46 := 0; fi46 < 2; fi46++ {
        for fi47 := 0; fi47 < 4; fi47++ {
            var v48 int32 = -430
            v2 = 274
            if true {
                fmt.Println(v48, v1, fi46)
            }
            _ = fi47
            _ = v48
        }
        _ = fi46
    }
    if (!true) {
        for fi49 := 0; fi49 < 8; fi49++ {
            v40 = 3.14159
            _ = fi49
        }
        for fi50 := 0; fi50 < 4; fi50++ {
            v3 -= 18446744073709551615
            _ = fi50
        }
    }
    if len(v41) >= 2 {
        v41 = v41[1:len(v41)]
    }
    v40 = v40
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
