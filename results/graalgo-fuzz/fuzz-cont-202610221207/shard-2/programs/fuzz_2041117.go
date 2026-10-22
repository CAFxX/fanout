package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int; F1 uint16; F2 bool }
type GenS3 struct { F0 string; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int64; F1 int; F2 string }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 2
    var v3 int = 128
    var v4 uint32 = 42
    v5 := make(map[int]bool)
    var v6 int16 = 794
    v7 := GenS1{F0: true, F1: (1e10 < -2.25)}
    v8 := GenS2{F0: v3, F1: 893, F2: (!true)}
    v9 := GenS3{F0: "hello", F1: (v4 + v4), F2: v6}
    for fi10 := 0; fi10 < 11; fi10++ {
        if true {
            switch "hello" {
            case "\n\t":
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v3)
                case "hello":
                    v11 := []int{}
                    _ = v11
                case "a":
                    var v12 bool = ((123 > 255) && false)
                    _ = v12
                default:
                    var v13 int16 = 42
                    _ = v13
                }
            default:
                fmt.Println(v6)
            }
            for fi14 := 0; fi14 < 13; fi14++ {
                v15 := GenS1{F0: ((42 == (v6 ^ 32767)) || (40 <= -176)), F1: (127 == -128)}
                v16 := GenS4{F0: 2, F1: 128}
                switch "\n\t" {
                case "hello":
                    if false {
                        v5[fi14] = false
                        var v17 int = (v3 - 2147483647)
                        fmt.Println(v6, v2, fi14)
                        _ = v17
                    }
                }
                var v18 uint64 = (v2 % (((v2 ^ (^(v2)))) | 1))
                v19 := GenS2{F0: 1, F1: 1, F2: (uint16(v2) <= 255)}
                _ = fi14
                _ = v15
                _ = v16
                _ = v18
                _ = v19
            }
            fmt.Println(fi10, fi10, v2)
        }
        if ((-118 > 42) || (730 <= 178)) {
            var v20 int = v3
            var v21 uint = 65535
            var v22 int = 849
            _ = v20
            _ = v21
            _ = v22
        } else {
            if false {
                v23 := GenS5{F0: 840, F1: (v3 >> ((1) & 63)), F2: "\n\t"}
                var v24 float32 = 0.0
                fmt.Println(v24, fi10, v2)
                _ = v23
                _ = v24
            } else {
                if (!((-2147483648 <= 369) || (((v3 | fi10) <= v3) && (false || false)))) {
                    v6 -= -674
                    fmt.Println(v6)
                    var v25 uint64 = 1
                    _ = v25
                }
                for fi26 := 0; fi26 < 20; fi26++ {
                    var v27 string = ("a" + "\n\t")
                    v3 -= (-(v3))
                    v28 := []int{32767, (fi10 / ((2147483647) | 1))}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                v4 = 42
            }
            var v29 int64 = 42
            _ = v29
        }
        delete(v1, "a")
        _ = fi10
    }
    if (-1 <= -76) {
        v6 = v6
        if (true && false) {
            if true {
                v30 := GenS1{F0: true, F1: true}
                _ = v30
            } else {
                var v31 uint8 = 1
                v3 -= v3
                _ = v31
            }
        }
        var v32 uint64 = (v2 ^ (v2 & (v2 ^ (v2 / (((v2 + 567)) | 1)))))
        _ = v32
    }
    var v33 int = 128
    for fi34 := 0; fi34 < 2; fi34++ {
        var v35 int32 = 2147483647
        _ = fi34
        _ = v35
    }
    if true {
        fmt.Println(v2)
        for fi36 := 0; fi36 < 2; fi36++ {
            v37 := []int{4294967295, v33, (v3 - v33), v3}
            if (!("hello" != "")) {
                v38 := GenS5{F0: -9223372036854775808, F1: (^(v3)), F2: ("foo bar" + "hello")}
                fmt.Println(v4, v6)
                var v39 uint64 = (^(v2))
                _ = v38
                _ = v39
            } else {
                fmt.Println(v6)
                delete(v5, 9223372036854775807)
            }
            var v40 uint32 = 65535
            _ = fi36
            _ = v37
            _ = v40
        }
        var v41 uint64 = 700
        _ = v41
    } else {
        fmt.Println(v3, v33, v2)
        delete(v1, "")
    }
    if (255 < 44) {
        delete(v1, "\n\t")
        var v42 uint = 444
        v2 -= v2
        _ = v42
    } else {
        v43 := GenS5{F0: 9223372036854775807, F1: (v33 << ((42) & 63)), F2: ("a" + ("foo bar" + "foo bar"))}
        var v44 uint64 = v2
        _ = v43
        _ = v44
    }
    v6 -= (v6 * 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
