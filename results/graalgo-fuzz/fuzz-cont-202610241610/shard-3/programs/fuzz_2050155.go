package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int16 }
type GenS2 struct { F0 string; F1 uint16; F2 bool }
type GenS3 struct { F0 uint; F1 string; F2 uint64 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int16; F1 int; F2 uint16 }
type GenS6 struct { F0 uint32; F1 bool }
func genFn0(p0 int16, p1 int16, p2 int64) int64 {
    var v7 int = 9223372036854775807
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint8 = 104
        _ = fi8
        _ = v9
    }
    switch "" {
    case "hello":
        v10 := GenS1{F0: -1, F1: (p2 | p2), F2: -32768}
        _ = v10
    case "\n\t":
        if false {
            var v11 int16 = int16(v7)
            v12 := GenS1{F0: 100, F1: p2, F2: int16(p2)}
            _ = v11
            _ = v12
        }
    case "foo bar":
        for fi13 := 0; fi13 < 16; fi13++ {
            switch ((("" + ("\n\t" + ("" + "a"))) + "foo bar") + ("foo bar" + "a")) {
            case "\n\t":
                p1 += 0
            case "":
                var v14 int32 = 1
                _ = v14
            }
            v15 := GenS1{F0: 24, F1: p2, F2: p0}
            for fi16 := 0; fi16 < 12; fi16++ {
                var v17 uint64 = 472
                p2 = p2
                _ = fi16
                _ = v17
            }
            _ = fi13
            _ = v15
        }
    default:
        fmt.Println(p2, p1)
    }
    v18 := GenS2{F0: "foo bar", F1: 65535, F2: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v7
    return (^(p2))
}
func genFn1(p0 int16) bool {
    v19 := GenS3{F0: 42, F1: "hello", F2: 18446744073709551615}
    if (819 != 602) {
        var v20 int32 = 1
        _ = v20
    } else {
        if false {
            var v21 float64 = -2.25
            v21 = (1e10 + (0.0 / 3.14159))
            if (4294967295 >= 4294967295) {
                var v22 uint16 = 2
                v21 -= 3.14159
                fmt.Println(v22)
                _ = v22
            } else {
                fmt.Println(p0)
            }
            _ = v21
        }
        v23 := GenS1{F0: -57, F1: -625, F2: (p0 % ((61) | 1))}
        _ = v23
    }
    var v24 string = "hello"
    if (true && false) {
        v25 := GenS4{F0: 2147483647, F1: 65535}
        if (!(false || false)) {
            var v26 int8 = 0
            p0 -= 42
            v26 = (v26 ^ 0)
            _ = v26
        }
        var v27 uint64 = 921
        _ = v25
        _ = v27
    } else {
        v24 = (v24 + ((v24 + (("a" + "\n\t") + ("hello" + ("a" + v24)))) + ("a" + v24)))
    }
    var v28 uint32 = 2147483647
    _ = p0
    _ = v19
    _ = v24
    _ = v28
    return true
}

func main() {
    v1 := []string{"a", "\n\t", "\n\t"}
    var v2 int32 = -2147483648
    v3 := make(map[string]uint32)
    var v4 int32 = v2
    v5 := []string{}
    var v6 string = "\n\t"
    if false {
        for fi29 := 0; fi29 < 5; fi29++ {
            switch -892 {
            case 32767:
                for fi30 := 0; fi30 < 5; fi30++ {
                    fmt.Println(fi30)
                    _ = fi30
                }
            case -1:
                var v31 int = (fi29 / ((1) | 1))
                _ = v31
            case -9223372036854775808:
                v4 -= v2
            default:
                for fi32 := 0; fi32 < 6; fi32++ {
                    v33 := GenS5{F0: 946, F1: fi29, F2: 0}
                    _ = fi32
                    _ = v33
                }
            }
            _ = fi29
        }
        v3[("a" + v6)] = 4294967295
        var v34 bool = false
        _ = v34
    }
    fmt.Println(v4)
    _ = len(v3)
    if ((1.5 + 0.0) < 0.1) {
        var v35 uint64 = 101
        if (!false) {
            if true {
                for fi36 := 0; fi36 < 12; fi36++ {
                    v6 = (v6 + "\n\t")
                    fmt.Println(v2)
                    _ = fi36
                }
                if true {
                    fmt.Println(v2)
                    v3[v6] = 2
                    v35 = (^(v35))
                }
            }
        } else {
            var v37 int32 = v2
            v38 := GenS5{F0: -1, F1: -405, F2: 786}
            switch ("foo bar" + v6) {
            case "foo bar":
                switch -304 {
                case 128:
                    var v39 uint = 955
                    _ = v39
                case 42:
                    var v40 uint64 = (v35 ^ v35)
                    _ = v40
                }
            case "":
                v37 -= v4
            case "a":
                v6 = v6
            default:
                v4 -= v2
            }
            _ = v37
            _ = v38
        }
        var v41 uint32 = 2147483648
        _ = v35
        _ = v41
    }
    for fi42 := 0; fi42 < 2; fi42++ {
        if (!(!(!(genFn0(4, 42, -1) < -1)))) {
            for fi43 := 0; fi43 < 8; fi43++ {
                _ = len(v1)
                _ = fi43
            }
            delete(v3, (v6 + ("foo bar" + v6)))
        }
        _ = fi42
    }
    v2 = (-(v4))
    var v44 int16 = -631
    if false {
        v45 := GenS6{F0: 0, F1: (-496 < -460)}
        v2 -= v2
        _ = v45
    }
    fmt.Println(v2, v44, v6)
    var v46 int8 = 42
    if true {
        v44 += (v44 % ((-1) | 1))
        var v47 int64 = -9223372036854775808
        for fi48 := 0; fi48 < 10; fi48++ {
            var v49 string = ("foo bar" + v6)
            _ = len(v3)
            v50 := GenS4{F0: 2147483647, F1: 0}
            _ = fi48
            _ = v49
            _ = v50
        }
        _ = v47
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v46
    _ = v5
    _ = v6
}
