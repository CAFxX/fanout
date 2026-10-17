package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint16 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int16; F1 uint32 }
func genFn0(p0 int64) uint64 {
    if false {
        v5 := GenS1{F0: -1, F1: (!true), F2: 0}
        p0 += -1
        fmt.Println(p0)
        _ = v5
    }
    switch "a" {
    case "a":
        switch (("a" + "hello") + "") {
        case "foo bar":
            p0 = p0
        case "a":
            switch -341 {
            case -2147483648:
                var v6 string = ("a" + "hello")
                _ = v6
            default:
                fmt.Println(p0)
            }
        case "hello":
            if false {
                v7 := []int{2147483647, 1}
                for fi8 := 0; fi8 < 10; fi8++ {
                    if false {
                        p0 = p0
                    }
                    var v9 string = (("foo bar" + "hello") + "")
                    for fi10 := 0; fi10 < 9; fi10++ {
                        var v11 uint64 = 0
                        if (-1 > -277) {
                            var v12 int = (fi10 & (-(fi10)))
                            _ = v12
                        }
                        _ = fi10
                        _ = v11
                    }
                    _ = fi8
                    _ = v9
                }
                var v13 uint32 = 589
                _ = v13
                _ = v7
            }
        default:
            var v14 int8 = -16
            _ = v14
        }
    case "foo bar":
        var v15 bool = (1 >= -614)
        _ = v15
    case "hello":
        p0 -= (^(p0))
    default:
        var v16 uint32 = 2147483647
        _ = v16
    }
    var v17 uint32 = 0
    if true {
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := []int{}
            _ = fi18
            _ = v19
        }
        var v20 uint16 = 65535
        for fi21 := 0; fi21 < 4; fi21++ {
            if false {
                var v22 uint16 = v20
                _ = v22
            }
            _ = fi21
        }
        _ = v20
    } else {
        if false {
            v23 := GenS2{F0: 0, F1: (!true)}
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v25 := GenS1{F0: -2147483648, F1: (!false), F2: 65535}
        p0 = (p0 >> ((255) & 63))
        if ((true || ((false || true) || true)) || (803 == -232)) {
            fmt.Println(p0, fi24)
            v26 := GenS2{F0: 1, F1: false}
            if true {
                for fi27 := 0; fi27 < 20; fi27++ {
                    fmt.Println(p0)
                    _ = fi27
                }
                v28 := GenS2{F0: 32767, F1: ((0 < 18446744073709551615) && true)}
                if (42 >= 255) {
                    var v29 uint32 = v17
                    if false {
                        for fi30 := 0; fi30 < 11; fi30++ {
                            var v31 int8 = -1
                            _ = fi30
                            _ = v31
                        }
                        v17 = (v17 / ((v17) | 1))
                    }
                    _ = v29
                }
                _ = v28
            }
            _ = v26
        }
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = v17
    return 2
}
func genFn1(p0 int, p1 int64, p2 int8) int16 {
    for fi32 := 0; fi32 < 3; fi32++ {
        fmt.Println(p2, p1)
        if true {
            p0 += p0
        }
        var v33 bool = false
        _ = fi32
        _ = v33
    }
    var v34 uint16 = 426
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    return -1
}

func main() {
    var v1 float32 = -2.25
    var v2 uint32 = 4294967295
    var v3 float64 = (-2.25 * (3.14159 * 0.0))
    v4 := []string{"\n\t", "foo bar"}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 = v3
    if false {
        fmt.Println(v2, v2)
    }
    v1 = (float32(((v3 - v3) / ((float64(v1) / 3.14159) - v3))) / v1)
    var v35 float64 = float64(0.1)
    for fi36 := 0; fi36 < 12; fi36++ {
        var v37 uint32 = 4294967295
        v1 -= v1
        _ = fi36
        _ = v37
    }
    var v38 uint = 0
    v39 := GenS2{F0: -1, F1: (9223372036854775807 != 32768)}
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (18446744073709551615 != v38) {
            v2 = v2
            switch "foo bar" {
            case "hello":
                if false {
                    v40 := GenS1{F0: -301, F1: true, F2: 32767}
                    for fi41 := 0; fi41 < 14; fi41++ {
                        var v42 uint = (v38 % (((v38 | v38)) | 1))
                        v1 += v1
                        _ = fi41
                        _ = v42
                    }
                    var v43 uint64 = 18446744073709551615
                    _ = v40
                    _ = v43
                }
            case "\n\t":
                for fi44 := 0; fi44 < 6; fi44++ {
                    fmt.Println(v3, v2)
                    _ = fi44
                }
            case "foo bar":
                v38 -= v38
            }
            var v45 uint64 = 4294967295
            _ = v45
        }
    }
    v46 := GenS3{F0: 32767, F1: v2}
    v3 = ((v3 * 1e10) * v3)
    fmt.Println(v2, v3, v2)
    for fi47 := 0; fi47 < 7; fi47++ {
        var v48 string = "foo bar"
        _ = fi47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v46
}
