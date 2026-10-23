package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int32 }
type GenS2 struct { F0 int64; F1 bool; F2 string }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 string }
func genFn0(p0 int16) int {
    if (!false) {
        fmt.Println(p0)
        v5 := []int{42, -122, -672, 4294967295}
        _ = v5
    }
    p0 = -32768
    v6 := []int{4294967295, -701}
    var v7 uint64 = 825
    _ = p0
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 int32, p1 int64) int8 {
    if (700 < 0) {
        v8 := []int{1, 127, 32768, 1}
        _ = v8
    } else {
        var v9 uint = 42
        var v10 float64 = 1e10
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: p1, F1: -385, F2: p0}
    var v12 float64 = (1e10 + 0.1)
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 70
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    v3 := []string{("foo bar" + "foo bar"), "\n\t", "foo bar"}
    var v4 string = "a"
    switch v4 {
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "\n\t":
        for fi13 := 0; fi13 < 7; fi13++ {
            fmt.Println(v4)
            _ = fi13
        }
    case "a":
        v4 = ((((v4 + (("foo bar" + ((v4 + v4) + ("a" + "\n\t"))) + v4)) + (((v4 + v4) + "hello") + "hello")) + v4) + (v4 + ("\n\t" + "hello")))
    default:
        switch 32767 {
        case -2147483648:
            v14 := GenS2{F0: 42, F1: true, F2: ""}
            _ = v14
        case 4294967295:
            v15 := GenS2{F0: -9223372036854775808, F1: ((("a" + (v4 + v4)) + v4) != ((v4 + (v4 + (v4 + (v4 + v4)))) + v4)), F2: (("" + "\n\t") + "foo bar")}
            _ = v15
        }
    }
    var v16 int = genFn0(1000)
    if false {
        v4 = ((v4 + (v4 + "a")) + "hello")
        for fi17 := 0; fi17 < 7; fi17++ {
            v16 = v16
            for fi18 := 0; fi18 < 10; fi18++ {
                fmt.Println(fi18)
                switch 210 {
                case 32768:
                    v4 = v4
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                var v19 int = v16
                _ = fi18
                _ = v19
            }
            v20 := GenS3{F0: "hello", F1: fi17}
            _ = fi17
            _ = v20
        }
        v16 += v16
    } else {
        for fi21 := 0; fi21 < 11; fi21++ {
            if (1e10 != (1.5 * 1e10)) {
                var v22 uint32 = 2147483648
                fmt.Println(v16, v22)
                _ = v22
            } else {
                var v23 int32 = 42
                var v24 uint = 2
                _ = v23
                _ = v24
            }
            v16 += (fi21 / (((fi21 - genFn0(0))) | 1))
            v4 = "hello"
            _ = fi21
        }
        v4 = v4
        var v25 int = (^(v16))
        _ = v25
    }
    if false {
        for fi26 := 0; fi26 < 9; fi26++ {
            v16 = (fi26 ^ (v16 ^ (v16 / (((fi26 >> ((18446744073709551615) & 63))) | 1))))
            fmt.Println(fi26)
            var v27 float64 = -2.25
            _ = fi26
            _ = v27
        }
        fmt.Println(v4, v16)
        if (1 == 65535) {
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 int = v16
                fmt.Println(v4, v16, v4)
                _ = fi28
                _ = v29
            }
            v30 := GenS2{F0: 1, F1: (v16 > v16), F2: v4}
            _ = v30
        } else {
            fmt.Println(v4, v16)
            if (true || false) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
    } else {
        v4 = ""
        var v31 uint64 = 0
        _ = v31
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        var v33 int64 = -9223372036854775808
        _ = fi32
        _ = v33
    }
    var v34 int64 = -9223372036854775808
    switch (-(v16)) {
    case 9223372036854775807:
        switch "hello" {
        case "foo bar":
            for fi35 := 0; fi35 < 5; fi35++ {
                for fi36 := 0; fi36 < 20; fi36++ {
                    var v37 uint16 = 836
                    fmt.Println(v16)
                    _ = fi36
                    _ = v37
                }
                v34 += v34
                var v38 bool = (!false)
                _ = fi35
                _ = v38
            }
        }
    default:
        switch (v16 | (v16 % (((-(v16))) | 1))) {
        case 1:
            var v39 uint16 = 65535
            _ = v39
        default:
            v16 = v16
        }
    }
    switch v16 {
    case 939:
        for fi40 := 0; fi40 < 17; fi40++ {
            var v41 int = v16
            _ = fi40
            _ = v41
        }
    default:
        fmt.Println(v4, v4)
    }
    v42 := GenS4{F0: v4}
    v34 = (v34 / ((v34) | 1))
    var v43 bool = ((true && false) || (true && true))
    v43 = v43
    for fi44 := 0; fi44 < 20; fi44++ {
        for fi45 := 0; fi45 < 8; fi45++ {
            switch v16 {
            case 4294967295:
                fmt.Println(fi45, v16, fi45)
            case -9223372036854775808:
                var v46 uint = 255
                _ = v46
            case 9223372036854775807:
                v47 := GenS2{F0: v34, F1: v43, F2: ""}
                _ = v47
            }
            _ = fi45
        }
        _ = fi44
    }
    v4 = ((("\n\t" + (v4 + "a")) + (v4 + "hello")) + v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v43
}
