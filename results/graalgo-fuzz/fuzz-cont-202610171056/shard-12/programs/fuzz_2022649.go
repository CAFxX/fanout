package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint32 }
type GenS2 struct { F0 bool; F1 int8; F2 int32 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 string; F1 int32; F2 int16 }
type GenS6 struct { F0 uint8; F1 int32 }
func genFn0(p0 int64, p1 int16) bool {
    switch "" {
    case "\n\t":
        v6 := GenS1{F0: (("\n\t" + ("foo bar" + ("a" + ""))) + "a"), F1: p0, F2: 65535}
        _ = v6
    case "foo bar":
        var v7 uint = 549
        _ = v7
    case "hello":
        if true {
            v8 := []int{}
            p0 = p0
            _ = v8
        }
    default:
        v9 := GenS1{F0: ("hello" + "\n\t"), F1: p0, F2: 291}
        _ = v9
    }
    v10 := GenS2{F0: false, F1: 1, F2: 1}
    v11 := GenS1{F0: "\n\t", F1: 507, F2: 42}
    if false {
        p1 = p1
    } else {
        var v12 uint16 = 42
        for fi13 := 0; fi13 < 20; fi13++ {
            v12 = v12
            _ = fi13
        }
        var v14 int64 = -1
        _ = v12
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return (0 < 209)
}
func genFn1(p0 uint64) int16 {
    p0 = p0
    v15 := []int{243, -794}
    _ = p0
    _ = v15
    return 1000
}

func main() {
    var v1 int32 = 0
    var v2 bool = false
    var v3 string = "a"
    v4 := make(map[uint32]uint32)
    var v5 int = 128
    if false {
        if ((!true) || v2) {
            switch v3 {
            case "a":
                v16 := GenS1{F0: "", F1: 42, F2: 181}
                _ = v16
            case "":
                v5 += (v5 + (v5 * 32768))
            case "foo bar":
                var v17 uint64 = 4294967295
                _ = v17
            default:
                if (!(1e10 == (1.5 / -2.25))) {
                    v3 = (v3 + ((v3 + "\n\t") + "a"))
                    delete(v4, 871)
                }
            }
            for fi18 := 0; fi18 < 17; fi18++ {
                var v19 int = (v5 & (fi18 >> ((2) & 63)))
                if true {
                    for fi20 := 0; fi20 < 1; fi20++ {
                        var v21 int = (fi20 * (v5 << ((2) & 63)))
                        _ = fi20
                        _ = v21
                    }
                }
                _ = fi18
                _ = v19
            }
            fmt.Println(v3)
        } else {
            v22 := GenS3{F0: 0, F1: ((v3 + (v3 + "")) + v3)}
            v23 := GenS4{F0: 18446744073709551615, F1: 42}
            var v24 float64 = -2.25
            _ = v22
            _ = v23
            _ = v24
        }
        if v2 {
            var v25 float64 = 0.1
            _ = v25
        }
    }
    v2 = false
    v5 += v5
    if (((v3 != "a") && v2) || (!(42 <= 0))) {
        for fi26 := 0; fi26 < 16; fi26++ {
            v4[614] = 65535
            _ = fi26
        }
        switch (v5 - 2147483647) {
        case 2147483647:
            v1 = v1
        default:
            for fi27 := 0; fi27 < 2; fi27++ {
                for fi28 := 0; fi28 < 11; fi28++ {
                    v29 := GenS5{F0: "hello", F1: (v1 % ((v1) | 1)), F2: 42}
                    fmt.Println(v5, v3, v1)
                    _ = fi28
                    _ = v29
                }
                v3 = "foo bar"
                _ = fi27
            }
        }
    }
    if ((-2.25 * 1e10) != 0.0) {
        var v30 int = (^(v5))
        _ = v30
    }
    var v31 uint = 1
    if v2 {
        if (2 >= 2) {
            v32 := []int{v5, v5, v5, (v5 * 32768)}
            if v2 {
                fmt.Println(v3, v2)
                v2 = true
            }
            for fi33 := 0; fi33 < 15; fi33++ {
                var v34 uint = (v31 | (v31 ^ 128))
                var v35 int16 = 42
                delete(v4, 1)
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = v32
        } else {
            fmt.Println(v2, v2, v31)
            fmt.Println(v3, v31)
            var v36 uint32 = 2147483647
            _ = v36
        }
        v37 := GenS2{F0: v2, F1: -51, F2: v1}
        var v38 int = v5
        _ = v37
        _ = v38
    }
    v2 = (493 != 0)
    if (v2 || v2) {
        if ((!true) || true) {
            var v39 int8 = 1
            var v40 bool = false
            _ = v39
            _ = v40
        }
        for fi41 := 0; fi41 < 12; fi41++ {
            fmt.Println(fi41, v1, v31)
            _ = fi41
        }
    }
    v5 += v5
    _ = len(v4)
    for fi42 := 0; fi42 < 14; fi42++ {
        v4[606] = 65535
        if v2 {
            var v43 int = fi42
            for fi44 := 0; fi44 < 19; fi44++ {
                fmt.Println(v1, v3)
                _ = fi44
            }
            v45 := GenS3{F0: 108, F1: v3}
            _ = v43
            _ = v45
        }
        _ = fi42
    }
    v46 := GenS6{F0: 127, F1: (v1 / ((v1) | 1))}
    var v47 int64 = -1
    if v2 {
        switch v5 {
        case -2147483648:
            var v48 uint16 = 32767
            _ = v48
        case 557:
            v47 = (v47 * v47)
        case 243:
            v49 := []int{-9223372036854775808, (v5 % ((v5) | 1)), v5, int(v31)}
            _ = v49
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
